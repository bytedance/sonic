/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
    `bufio`
    `encoding/json`
    `fmt`
    `io`
    `os`
    `sort`
    `strings`
    `testing`

    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

var visitorTestCases = []struct {
    name string
    jsonStr string
} {
    {"default", _TwitterJson},
    {"issue_case01", "[1193.6419677734375]"},
    {"issue653", `{"v0": 0, "m0": {}, "v1": 1, "a0": [], "v2": 2}`},
}

type visitorNodeDiffTest struct {
    t   *testing.T
    str string

    tracer io.Writer

    cursor Node
    stk    visitorNodeStack
    sp     uint8
}

type visitorNodeStack = [256]struct {
    Node   Node
    Object map[string]Node
    Array  []Node

    ObjectKey string
}

func (self *visitorNodeDiffTest) incrSP() {
    self.t.Helper()
    self.sp++
    require.NotZero(self.t, self.sp, "stack overflow")
}

func (self *visitorNodeDiffTest) debugStack() string {
    var buf strings.Builder
    buf.WriteString("[")
    for i := uint8(0); i < self.sp; i++ {
        if i != 0 {
            buf.WriteString(", ")
        }
        if self.stk[i].Array != nil {
            buf.WriteString("Array")
        } else if self.stk[i].Object != nil {
            buf.WriteString("Object")
        } else {
            fmt.Fprintf(&buf, "Key(%q)", self.stk[i].ObjectKey)
        }
    }
    buf.WriteString("]")
    return buf.String()
}

func (self *visitorNodeDiffTest) requireType(got int) {
    self.t.Helper()
    want := self.cursor.Type()
    require.EqualValues(self.t, want, got)
}

func (self *visitorNodeDiffTest) toArrayIndex(array Node, i int) {
    // set cursor to next Value if existed
    self.t.Helper()
    n, err := array.Len()
    require.NoError(self.t, err)
    if i < n {
        self.cursor = *array.Index(i)
        require.NoError(self.t, self.cursor.Check())
    }
}

func (self *visitorNodeDiffTest) onValueEnd() {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnValueEnd: %s\n", self.debugStack())
    }
    // cursor should point to the Value now
    self.t.Helper()
    if self.sp == 0 {
        if self.tracer != nil {
            fmt.Fprintf(self.tracer, "EOF\n\n")
        }
        return
    }
    // [..., Array, sp]
    if array := self.stk[self.sp-1].Array; array != nil {
        array = append(array, self.cursor)
        self.stk[self.sp-1].Array = array
        self.toArrayIndex(self.stk[self.sp-1].Node, len(array))
        return
    }
    // [..., Object, ObjectKey, sp]
    require.GreaterOrEqual(self.t, self.sp, uint8(2))
    require.NotNil(self.t, self.stk[self.sp-2].Object)
    require.Nil(self.t, self.stk[self.sp-1].Object)
    require.Nil(self.t, self.stk[self.sp-1].Array)
    self.stk[self.sp-2].Object[self.stk[self.sp-1].ObjectKey] = self.cursor
    self.cursor = self.stk[self.sp-2].Node // reset cursor to Object
    self.sp--                              // pop ObjectKey
}

func (self *visitorNodeDiffTest) OnNull() error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnNull\n")
    }
    self.requireType(V_NULL)
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnBool(v bool) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnBool: %t\n", v)
    }
    if v {
        self.requireType(V_TRUE)
    } else {
        self.requireType(V_FALSE)
    }
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnString(v string) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnString: %q\n", v)
    }
    self.requireType(V_STRING)
    want, err := self.cursor.StrictString()
    require.NoError(self.t, err)
    require.EqualValues(self.t, want, v)
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnInt64(v int64, n json.Number) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnInt64: %d (%q)\n", v, n)
    }
    self.requireType(V_NUMBER)
    want, err := self.cursor.StrictInt64()
    require.NoError(self.t, err)
    require.EqualValues(self.t, want, v)
    nv, err := n.Int64()
    require.NoError(self.t, err)
    require.EqualValues(self.t, want, nv)
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnFloat64(v float64, n json.Number) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnFloat64: %f (%q)\n", v, n)
    }
    self.requireType(V_NUMBER)
    want, err := self.cursor.StrictFloat64()
    require.NoError(self.t, err)
    require.EqualValues(self.t, want, v)
    nv, err := n.Float64()
    require.NoError(self.t, err)
    require.EqualValues(self.t, want, nv)
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnObjectBegin(capacity int) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnObjectBegin: %d\n", capacity)
    }
    self.requireType(V_OBJECT)
    self.stk[self.sp].Node = self.cursor
    self.stk[self.sp].Object = make(map[string]Node, capacity)
    self.incrSP()
    return nil
}

func (self *visitorNodeDiffTest) OnObjectKey(key string) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnObjectKey: %q %s\n", key, self.debugStack())
    }
    require.NotNil(self.t, self.stk[self.sp-1].Object)
    node := self.stk[self.sp-1].Node
    self.stk[self.sp].ObjectKey = key
    self.incrSP()
    self.cursor = *node.Get(key) // set cursor to Value
    require.NoError(self.t, self.cursor.Check())
    return nil
}

func (self *visitorNodeDiffTest) OnObjectEnd() error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnObjectEnd\n")
    }
    object := self.stk[self.sp-1].Object
    require.NotNil(self.t, object)

    node := self.stk[self.sp-1].Node
    ps, err := node.unsafeMap()
    var pairs = make([]Pair, ps.Len())
    ps.ToSlice(pairs)
    require.NoError(self.t, err)

    keysGot := make([]string, 0, len(object))
    for key := range object {
        keysGot = append(keysGot, key)
    }
    keysWant := make([]string, 0, len(pairs))
    for _, pair := range pairs {
        keysWant = append(keysWant, pair.Key)
    }
    sort.Strings(keysGot)
    sort.Strings(keysWant)
    require.EqualValues(self.t, keysWant, keysGot)

    for _, pair := range pairs {
        typeGot := object[pair.Key].Type()
        typeWant := pair.Value.Type()
        require.EqualValues(self.t, typeWant, typeGot)
    }

    // pop Object
    self.sp--
    self.stk[self.sp].Node = Node{}
    self.stk[self.sp].Object = nil

    self.cursor = node // set cursor to this Object
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) OnArrayBegin(capacity int) error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnArrayBegin: %d\n", capacity)
    }
    self.requireType(V_ARRAY)
    self.stk[self.sp].Node = self.cursor
    self.stk[self.sp].Array = make([]Node, 0, capacity)
    self.incrSP()
    self.toArrayIndex(self.stk[self.sp-1].Node, 0)
    return nil
}

func (self *visitorNodeDiffTest) OnArrayEnd() error {
    if self.tracer != nil {
        fmt.Fprintf(self.tracer, "OnArrayEnd\n")
    }
    array := self.stk[self.sp-1].Array
    require.NotNil(self.t, array)

    node := self.stk[self.sp-1].Node
    vs, err := node.unsafeArray()
    require.NoError(self.t, err)
    var values = make([]Node, vs.Len())
    vs.ToSlice(values)

    require.EqualValues(self.t, len(values), len(array))

    for i, n := 0, len(values); i < n; i++ {
        typeGot := array[i].Type()
        typeWant := values[i].Type()
        require.EqualValues(self.t, typeWant, typeGot)
    }

    // pop Array
    self.sp--
    self.stk[self.sp].Node = Node{}
    self.stk[self.sp].Array = nil

    self.cursor = node // set cursor to this Array
    self.onValueEnd()
    return nil
}

func (self *visitorNodeDiffTest) Run(t *testing.T, str string,
    tracer io.Writer) {
    self.t = t
    self.str = str
    self.tracer = tracer

    self.t.Helper()

    self.cursor = NewRaw(self.str)
    require.NoError(self.t, self.cursor.LoadAll())

    self.stk = visitorNodeStack{}
    self.sp = 0

    require.NoError(self.t, Preorder(self.str, self, nil))
}

func TestVisitor_NodeDiff(t *testing.T) {
    var suite visitorNodeDiffTest

    newTracer := func(t *testing.T) io.Writer {
        const EnableTracer = false
        if !EnableTracer {
            return nil
        }
        basename := strings.ReplaceAll(t.Name(), "/", "_")
        fp, err := os.Create(fmt.Sprintf("../output/%s.log", basename))
        require.NoError(t, err)
        writer := bufio.NewWriter(fp)
        t.Cleanup(func() {
            _ = writer.Flush()
            _ = fp.Close()
        })
        return writer
    }

    for _, c := range visitorTestCases {
        t.Run(c.name, func(t *testing.T) {
            suite.Run(t, c.jsonStr, newTracer(t))
        })
    }
}

type visitorUserNode interface {
    UserNode()
}

type (
    visitorUserNull    struct{}
    visitorUserBool    struct{ Value bool }
    visitorUserInt64   struct{ Value int64 }
    visitorUserFloat64 struct{ Value float64 }
    visitorUserString  struct{ Value string }
    visitorUserObject  struct{ Value map[string]visitorUserNode }
    visitorUserArray   struct{ Value []visitorUserNode }
)

func (*visitorUserNull) UserNode()    {}
func (*visitorUserBool) UserNode()    {}
func (*visitorUserInt64) UserNode()   {}
func (*visitorUserFloat64) UserNode() {}
func (*visitorUserString) UserNode()  {}
func (*visitorUserObject) UserNode()  {}
func (*visitorUserArray) UserNode()   {}

func compareUserNode(tb testing.TB, lhs, rhs visitorUserNode) bool {
    switch lhs := lhs.(type) {
    case *visitorUserNull:
        _, ok := rhs.(*visitorUserNull)
        return assert.True(tb, ok)
    case *visitorUserBool:
        rhs, ok := rhs.(*visitorUserBool)
        return assert.True(tb, ok) && assert.Equal(tb, lhs.Value, rhs.Value)
    case *visitorUserInt64:
        rhs, ok := rhs.(*visitorUserInt64)
        return assert.True(tb, ok) && assert.Equal(tb, lhs.Value, rhs.Value)
    case *visitorUserFloat64:
        rhs, ok := rhs.(*visitorUserFloat64)
        return assert.True(tb, ok) && assert.Equal(tb, lhs.Value, rhs.Value)
    case *visitorUserString:
        rhs, ok := rhs.(*visitorUserString)
        return assert.True(tb, ok) && assert.Equal(tb, lhs.Value, rhs.Value)
    case *visitorUserObject:
        rhs, ok := rhs.(*visitorUserObject)
        if !(assert.True(tb, ok) && assert.Equal(tb, len(lhs.Value), len(rhs.Value))) {
            return false
        }
        for key, lhs := range lhs.Value {
            rhs, ok := rhs.Value[key]
            if !(assert.True(tb, ok) && assert.True(tb, compareUserNode(tb, lhs, rhs))) {
                return false
            }
        }
        return true
    case *visitorUserArray:
        rhs, ok := rhs.(*visitorUserArray)
        if !(assert.True(tb, ok) && assert.Equal(tb, len(lhs.Value), len(rhs.Value))) {
            return false
        }
        for i, n := 0, len(lhs.Value); i < n; i++ {
            if !assert.True(tb, compareUserNode(tb, lhs.Value[i], rhs.Value[i])) {
                return false
            }
        }
        return true
    default:
        tb.Fatalf("unexpected type of UserNode: %T", lhs)
        return false
    }
}

type visitorUserNodeDecoder interface {
    Reset()
    Decode(str string) (visitorUserNode, error)
}

var _ visitorUserNodeDecoder = (*visitorUserNodeASTDecoder)(nil)

type visitorUserNodeASTDecoder struct{}

func (self *visitorUserNodeASTDecoder) Reset() {}

func (self *visitorUserNodeASTDecoder) Decode(str string) (visitorUserNode, error) {
    root := NewRaw(str)
    if err := root.LoadAll(); err != nil {
        return nil, err
    }
    return self.decodeValue(&root)
}

func (self *visitorUserNodeASTDecoder) decodeValue(root *Node) (visitorUserNode, error) {
    switch typ := root.Type(); typ {
    // embed (*Node).Check
    case V_NONE:
        return nil, ErrNotExist
    case V_ERROR:
        return nil, root

    case V_NULL:
        return &visitorUserNull{}, nil
    case V_TRUE:
        return &visitorUserBool{Value: true}, nil
    case V_FALSE:
        return &visitorUserBool{Value: false}, nil

    case V_STRING:
        value, err := root.StrictString()
        if err != nil {
            return nil, err
        }
        return &visitorUserString{Value: value}, nil

    case V_NUMBER:
        value, err := root.StrictNumber()
        if err != nil {
            return nil, err
        }
        i64, ierr := value.Int64()
        if ierr == nil {
            return &visitorUserInt64{Value: i64}, nil
        }
        f64, ferr := value.Float64()
        if ferr == nil {
            return &visitorUserFloat64{Value: f64}, nil
        }
        return nil, fmt.Errorf("invalid number: %v, ierr: %v, ferr: %v",
            value, ierr, ferr)

    case V_ARRAY:
        nodes, err := root.unsafeArray()
        if err != nil {
            return nil, err
        }
        values := make([]visitorUserNode, nodes.Len())
        for i := 0; i<nodes.Len(); i++ {
            n := nodes.At(i)
            value, err := self.decodeValue(n)
            if err != nil {
                return nil, err
            }
            values[i] = value
        }
        return &visitorUserArray{Value: values}, nil

    case V_OBJECT:
        pairs, err := root.unsafeMap()
        if err != nil {
            return nil, err
        }
        values := make(map[string]visitorUserNode, pairs.Len())
        for i := 0; i < pairs.Len(); i++ {
            value, err := self.decodeValue(&pairs.At(i).Value)
            if err != nil {
                return nil, err
            }
            values[pairs.At(i).Key] = value
        }
        return &visitorUserObject{Value: values}, nil

    case V_ANY:
        fallthrough
    default:
        return nil, fmt.Errorf("unexpected Node type: %v", typ)
    }
}

var _ visitorUserNodeDecoder = (*visitorUserNodeVisitorDecoder)(nil)

type visitorUserNodeVisitorDecoder struct {
    stk visitorUserNodeStack
    sp  uint8
}

type visitorUserNodeStack = [256]struct {
    val visitorUserNode
    obj map[string]visitorUserNode
    arr []visitorUserNode
    key string
}

func (self *visitorUserNodeVisitorDecoder) Reset() {
    self.stk = visitorUserNodeStack{}
    self.sp = 0
}

func (self *visitorUserNodeVisitorDecoder) Decode(str string) (visitorUserNode, error) {
    if err := Preorder(str, self, nil); err != nil {
        return nil, err
    }
    return self.result()
}

func (self *visitorUserNodeVisitorDecoder) result() (visitorUserNode, error) {
    if self.sp != 1 {
        return nil, fmt.Errorf("incorrect sp: %d", self.sp)
    }
    return self.stk[0].val, nil
}

func (self *visitorUserNodeVisitorDecoder) incrSP() error {
    self.sp++
    if self.sp == 0 {
        return fmt.Errorf("reached max depth: %d", len(self.stk))
    }
    return nil
}

func (self *visitorUserNodeVisitorDecoder) OnNull() error {
    self.stk[self.sp].val = &visitorUserNull{}
    if err := self.incrSP(); err != nil {
        return err
    }
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnBool(v bool) error {
    self.stk[self.sp].val = &visitorUserBool{Value: v}
    if err := self.incrSP(); err != nil {
        return err
    }
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnString(v string) error {
    self.stk[self.sp].val = &visitorUserString{Value: v}
    if err := self.incrSP(); err != nil {
        return err
    }
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnInt64(v int64, n json.Number) error {
    self.stk[self.sp].val = &visitorUserInt64{Value: v}
    if err := self.incrSP(); err != nil {
        return err
    }
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnFloat64(v float64, n json.Number) error {
    self.stk[self.sp].val = &visitorUserFloat64{Value: v}
    if err := self.incrSP(); err != nil {
        return err
    }
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnObjectBegin(capacity int) error {
    self.stk[self.sp].obj = make(map[string]visitorUserNode, capacity)
    return self.incrSP()
}

func (self *visitorUserNodeVisitorDecoder) OnObjectKey(key string) error {
    self.stk[self.sp].key = key
    return self.incrSP()
}

func (self *visitorUserNodeVisitorDecoder) OnObjectEnd() error {
    self.stk[self.sp-1].val = &visitorUserObject{Value: self.stk[self.sp-1].obj}
    self.stk[self.sp-1].obj = nil
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) OnArrayBegin(capacity int) error {
    self.stk[self.sp].arr = make([]visitorUserNode, 0, capacity)
    return self.incrSP()
}

func (self *visitorUserNodeVisitorDecoder) OnArrayEnd() error {
    self.stk[self.sp-1].val = &visitorUserArray{Value: self.stk[self.sp-1].arr}
    self.stk[self.sp-1].arr = nil
    return self.onValueEnd()
}

func (self *visitorUserNodeVisitorDecoder) onValueEnd() error {
    if self.sp == 1 {
        return nil
    }
    // [..., Array, Value, sp]
    if self.stk[self.sp-2].arr != nil {
        self.stk[self.sp-2].arr = append(self.stk[self.sp-2].arr, self.stk[self.sp-1].val)
        self.sp--
        return nil
    }
    // [..., Object, ObjectKey, Value, sp]
    self.stk[self.sp-3].obj[self.stk[self.sp-2].key] = self.stk[self.sp-1].val
    self.sp -= 2
    return nil
}

func testUserNodeDiff(t *testing.T, d1, d2 visitorUserNodeDecoder, str string) {
    t.Helper()
    d1.Reset()
    n1, err := d1.Decode(_TwitterJson)
    require.NoError(t, err)

    d2.Reset()
    n2, err := d2.Decode(_TwitterJson)
    require.NoError(t, err)

    require.True(t, compareUserNode(t, n1, n2))
}

func TestVisitor_UserNodeDiff(t *testing.T) {
    var d1 visitorUserNodeASTDecoder
    var d2 visitorUserNodeVisitorDecoder

    for _, c := range visitorTestCases {
        t.Run(c.name, func(t *testing.T) {
            testUserNodeDiff(t, &d1, &d2, c.jsonStr)
        })
    }
}

type skipVisitor struct {
    sp int
    Skip int
    inSkip bool
    CountSkip int
}

func (self *skipVisitor) OnNull() error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnFloat64(v float64, n json.Number) error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnInt64(v int64, n json.Number) error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnBool(v bool) error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnString(v string) error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnObjectBegin(capacity int) error {
    println("self.sp", self.sp)
    if self.sp == self.Skip {
        self.inSkip = true
        self.CountSkip++
        println("op skip")
        self.sp++
        return VisitOPSkip
    }
    self.sp++
    return nil
}

func (self *skipVisitor) OnObjectKey(key string) error {
    if self.sp == self.Skip+1 && self.inSkip  {
        panic("unexpected key")
    }
    return nil
}

func (self *skipVisitor) OnObjectEnd() error {
    if self.sp == self.Skip + 1 {
        if !self.inSkip {
            panic("not in skip")
        }
        self.inSkip = false
        println("finish op skip")
    }
    self.sp--
    return nil
}

func (self *skipVisitor) OnArrayBegin(capacity int) error {
    println("arr self.sp", self.sp)
    if self.sp == self.Skip {
        self.inSkip = true
        self.CountSkip++
        println("arr op skip")
        self.sp++
        return VisitOPSkip
    }
    self.sp++
    return nil
}

func (self *skipVisitor) OnArrayEnd() error {
    println("arr self.sp", self.sp)
    if self.sp == self.Skip + 1 {
        if !self.inSkip {
            panic("arr not in skip")
        }
        self.inSkip = false
        println("arr finish op skip")
    }
    self.sp--
    return nil
}

func TestVisitor_OpSkip(t *testing.T) {
    var suite skipVisitor
    suite.Skip = 1
    Preorder(`{ "a": [ null ] , "b": 1, "c": { "1" : 1 } }`, &suite, nil)
    if suite.CountSkip != 2 {
        t.Fatal(suite.CountSkip)
    }
}

func BenchmarkVisitor_UserNode(b *testing.B) {
    const str = _TwitterJson
    b.Run("AST", func(b *testing.B) {
        var d visitorUserNodeASTDecoder
        b.ResetTimer()
        for k := 0; k < b.N; k++ {
            d.Reset()
            _, err := d.Decode(str)
            require.NoError(b, err)
            b.SetBytes(int64(len(str)))
        }
    })
    b.Run("Visitor", func(b *testing.B) {
        var d visitorUserNodeVisitorDecoder
        b.ResetTimer()
        for k := 0; k < b.N; k++ {
            d.Reset()
            _, err := d.Decode(str)
            require.NoError(b, err)
            b.SetBytes(int64(len(str)))
        }
    })
}
