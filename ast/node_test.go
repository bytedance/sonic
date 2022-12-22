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
   `encoding/json`
   `errors`
   `fmt`
   `reflect`
   `runtime`
   `runtime/debug`
   `strconv`
   `testing`

   `github.com/bytedance/sonic/internal/native/types`
   `github.com/bytedance/sonic/internal/rt`
   `github.com/stretchr/testify/assert`
)


func TestNodeSortKeys(t *testing.T) {
    var src = `{"b":1,"a":2,"c":3}`
    root, err := NewSearcher(src).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    obj, err := root.MapUseNumber()
    if err != nil {
        t.Fatal(err)
    }
    exp, err := json.Marshal(obj)
    if err != nil {
        t.Fatal(err)
    }
    if err := root.SortKeys(true); err != nil {
        t.Fatal(err)
    }
    act, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    assert.Equal(t, len(exp), len(act))
    assert.Equal(t, string(exp), string(act))
}

func BenchmarkNodeSortKeys(b *testing.B) {
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        b.Fatal(err)
    }
    if err := root.LoadAll(); err != nil {
        b.Fatal(err)
    }
    
    b.Run("single", func(b *testing.B) {
        r := root.Get("statuses")
        if r.Check() != nil {
            b.Fatal(r.Error())
        }
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = root.SortKeys(false)
        }
    })
    b.Run("recurse", func(b *testing.B) {
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = root.SortKeys(true)
        }
    })
}

//go:noinline
func stackObj() interface{} {
    var a int = 1
    return rt.UnpackEface(a).Pack()
}

func TestStackAny(t *testing.T) {
    var obj = stackObj()
    any := NewAny(obj)
    fmt.Printf("any: %#v\n", any)
    runtime.GC()
    debug.FreeOSMemory()
    println("finish GC")
    buf, err := any.MarshalJSON()
    println("finish marshal")
    if err != nil {
        t.Fatal(err)
    }
    if string(buf) != `1` {
        t.Fatal(string(buf))
    }
}

func TestLoadAll(t *testing.T) {
    e := Node{}
    err := e.Load()
    if err != nil {
        t.Fatal(err)
    }
    err = e.LoadAll()
    if err != nil {
        t.Fatal(err)
    }

    root, err := NewSearcher(`{"a":{"1":[1],"2":2},"b":[{"1":1},2],"c":[1,2]}`).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    if err = root.Load(); err != nil {
        t.Fatal(err)
    }
    if root.len() != 3 {
        t.Fatal(root.len())
    }

    c := root.Get("c")
    if !c.IsRaw() {
        t.Fatal(err)
    }
    err = c.LoadAll()
    if err != nil {
        t.Fatal(err)
    }
    if c.len() != 2 {
        t.Fatal(c.len())
    }
    c1 := c.nodeAt(0)
    if n, err := c1.Int64(); err != nil || n != 1 {
        t.Fatal(n, err)
    }

    a := root.pairAt(0)
    if a.Key != "a" {
        t.Fatal(a.Key)
    } else if !a.Value.IsRaw() {
        t.Fatal(a.Value.itype())
    } else if n, err := a.Value.Len(); n != 0 || err != nil {
        t.Fatal(n, err)
    }
    if err := a.Value.Load(); err != nil {
        t.Fatal(err)
    }
    if a.Value.len() != 2 {
        t.Fatal(a.Value.len())
    }
    a1 := a.Value.Get("1")
    if !a1.IsRaw() {
        t.Fatal(a1)
    }
    a.Value.LoadAll()
    if a1.t != types.V_ARRAY || a1.len() != 1 {
        t.Fatal(a1.t, a1.len())
    }

    b := root.pairAt(1)
    if b.Key != "b" {
        t.Fatal(b.Key)
    } else if !b.Value.IsRaw() {
        t.Fatal(b.Value.itype())
    } else if n, err := b.Value.Len(); n != 0 || err != nil {
        t.Fatal(n, err)
    }
    if err := b.Value.Load(); err != nil {
        t.Fatal(err)
    }
    if b.Value.len() != 2 {
        t.Fatal(b.Value.len())
    }
    b1 := b.Value.Index(0)
    if !b1.IsRaw() {
        t.Fatal(b1)
    }
    b.Value.LoadAll()
    if b1.t != types.V_OBJECT || b1.len() != 1 {
        t.Fatal(a1.t, a1.len())
    }
}

func TestIndexPair(t *testing.T) {
    root, _ := NewParser(`{"a":1,"b":2}`).Parse()
    a := root.IndexPair(0)
    if a == nil || a.Key != "a" {
        t.Fatal(a)
    }
    b := root.IndexPair(1)
    if b == nil || b.Key != "b" {
        t.Fatal(b)
    }
    c := root.IndexPair(2)
    if c != nil {
        t.Fatal(c)
    }
}

func TestIndexOrGet(t *testing.T) {
    root, _ := NewParser(`{"a":1,"b":2}`).Parse()
    a := root.IndexOrGet(0, "a")
    if v, err := a.Int64(); err != nil || v != int64(1) {
        t.Fatal(a)
    }
    a = root.IndexOrGet(0, "b")
    if v, err := a.Int64(); err != nil || v != int64(2) {
        t.Fatal(a)
    }
    a = root.IndexOrGet(0, "c")
    if a.Valid()  {
        t.Fatal(a)
    }
}

func TestTypeCast(t *testing.T) {
    type tcase struct {
        method string
        node Node
        exp interface{}
        err error
    }
    var nonEmptyErr error = errors.New("")
    a1 := NewAny(1)
    lazyArray, _ := NewParser("[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]").Parse()
    lazyObject, _ := NewParser(`{"0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"10":10,"11":11,"12":12,"13":13,"14":14,"15":15,"16":16}`).Parse()
    var cases = []tcase{
        {"Interface", Node{}, interface{}(nil), ErrUnsupportType},
        {"Interface", NewAny(NewNumber("1")), float64(1), nil},
        {"Interface", NewAny(int64(1)), int64(1), nil},
        {"Interface", NewNumber("1"), float64(1), nil},
        {"InterfaceUseNode", Node{}, Node{}, nil},
        {"InterfaceUseNode", a1, a1, nil},
        {"InterfaceUseNode", NewNumber("1"), NewNumber("1"), nil},
        {"InterfaceUseNumber", Node{}, interface{}(nil), ErrUnsupportType},
        {"InterfaceUseNumber", NewAny(1), 1, nil},
        {"InterfaceUseNumber", NewNumber("1"), json.Number("1"), nil},
        {"Map", Node{}, map[string]interface{}(nil), ErrUnsupportType},
        {"Map", NewAny(map[string]Node{"a":NewNumber("1")}), map[string]interface{}(nil), ErrUnsupportType},
        {"Map", NewAny(map[string]interface{}{"a":1}), map[string]interface{}{"a":1}, nil},
        {"Map", NewObject([]Pair{{"a",NewNumber("1")}}), map[string]interface{}{"a":float64(1.0)}, nil},
        {"MapUseNode", Node{}, map[string]Node(nil), ErrUnsupportType},
        {"MapUseNode", NewAny(map[string]interface{}{"a":1}), map[string]Node(nil), ErrUnsupportType},
        {"MapUseNode", NewAny(map[string]Node{"a":NewNumber("1")}), map[string]Node{"a":NewNumber("1")}, nil},
        {"MapUseNode", NewObject([]Pair{{"a",NewNumber("1")}}), map[string]Node{"a":NewNumber("1")}, nil},
        {"MapUseNumber", Node{}, map[string]interface{}(nil), ErrUnsupportType},
        {"MapUseNumber", NewAny(map[string]interface{}{"a":1}), map[string]interface{}{"a":1}, nil},
        {"MapUseNumber", NewObject([]Pair{{"a",NewNumber("1")}}), map[string]interface{}{"a":json.Number("1")}, nil},
        {"Array", Node{}, []interface{}(nil), ErrUnsupportType},
        {"Array", NewAny([]interface{}{1}), []interface{}{1}, nil},
        {"Array", NewArray([]Node{NewNumber("1")}), []interface{}{float64(1.0)}, nil},
        {"ArrayUseNode", Node{}, []Node(nil), ErrUnsupportType},
        {"ArrayUseNode", NewAny([]interface{}{1}), []Node(nil), ErrUnsupportType},
        {"ArrayUseNode", NewAny([]Node{NewNumber("1")}), []Node{NewNumber("1")}, nil},
        {"ArrayUseNode", NewArray([]Node{NewNumber("1")}), []Node{NewNumber("1")}, nil},
        {"ArrayUseNumber", Node{}, []interface{}(nil), ErrUnsupportType},
        {"ArrayUseNumber", NewAny([]interface{}{1}), []interface{}{1}, nil},
        {"ArrayUseNumber", NewAny([]Node{NewNumber("1")}), []interface{}(nil), ErrUnsupportType},
        {"ArrayUseNumber", NewArray([]Node{NewNumber("1")}), []interface{}{json.Number("1")}, nil},
        {"Raw", Node{}, "", ErrNotExist},
        {"Raw", NewAny(""), `""`, nil},
        {"Raw", NewRaw("[ ]"), "[ ]", nil},
        {"Raw", NewBool(true), "true", nil},
        {"Raw", NewNumber("-0.0"), "-0.0", nil},
        {"Raw", NewString(""), `""`, nil},
        {"Raw", NewBytes([]byte("hello, world")), `"aGVsbG8sIHdvcmxk"`, nil},
        {"Bool", Node{}, false, ErrUnsupportType},
        {"Bool", NewAny(true), true, nil},
        {"Bool", NewAny(false), false, nil},
        {"Bool", NewAny(int(0)), false, nil},
        {"Bool", NewAny(int8(1)), true, nil},
        {"Bool", NewAny(int16(1)), true, nil},
        {"Bool", NewAny(int32(1)), true, nil},
        {"Bool", NewAny(int64(1)), true, nil},
        {"Bool", NewAny(uint(1)), true, nil},
        {"Bool", NewAny(uint16(1)), true, nil},
        {"Bool", NewAny(uint32(1)), true, nil},
        {"Bool", NewAny(uint64(1)), true, nil},
        {"Bool", NewAny(float64(0)), false, nil},
        {"Bool", NewAny(float32(1)), true, nil},
        {"Bool", NewAny(float64(1)), true, nil},
        {"Bool", NewAny(json.Number("0")), false, nil},
        {"Bool", NewAny(json.Number("1")), true, nil},
        {"Bool", NewAny(json.Number("1.1")), true, nil},
        {"Bool", NewAny(json.Number("+x1.1")), false, nonEmptyErr},
        {"Bool", NewAny(string("0")), false, nil},
        {"Bool", NewAny(string("t")), true, nil},
        {"Bool", NewAny([]byte{0}), false, nonEmptyErr},
        {"Bool", NewRaw("true"), true, nil},
        {"Bool", NewRaw("false"), false, nil},
        {"Bool", NewRaw("null"), false, nil},
        {"Bool", NewString(`true`), true, nil},
        {"Bool", NewString(`false`), false, nil},
        {"Bool", NewString(``), false, nonEmptyErr},
        {"Bool", NewNumber("2"), true, nil},
        {"Bool", NewNumber("-2.1"), true, nil},
        {"Bool", NewNumber("-x-2.1"), false, nonEmptyErr},
        {"Int64", NewRaw("true"), int64(1), nil},
        {"Int64", NewRaw("false"), int64(0), nil},
        {"Int64", NewRaw("\"1\""), int64(1), nil},
        {"Int64", NewRaw("\"1.1\""), int64(1), nil},
        {"Int64", NewRaw("\"1.0\""), int64(1), nil},
        {"Int64", NewNumber("+x.0"), int64(0), nonEmptyErr},
        {"Int64", NewAny(false), int64(0), nil},
        {"Int64", NewAny(true), int64(1), nil},
        {"Int64", NewAny(int(1)), int64(1), nil},
        {"Int64", NewAny(int8(1)), int64(1), nil},
        {"Int64", NewAny(int16(1)), int64(1), nil},
        {"Int64", NewAny(int32(1)), int64(1), nil},
        {"Int64", NewAny(int64(1)), int64(1), nil},
        {"Int64", NewAny(uint(1)), int64(1), nil},
        {"Int64", NewAny(uint8(1)), int64(1), nil},
        {"Int64", NewAny(uint32(1)), int64(1), nil},
        {"Int64", NewAny(uint64(1)), int64(1), nil},
        {"Int64", NewAny(float32(1)), int64(1), nil},
        {"Int64", NewAny(float64(1)), int64(1), nil},
        {"Int64", NewAny("1"), int64(1), nil},
        {"Int64", NewAny("1.1"), int64(1), nil},
        {"Int64", NewAny("+1x.1"), int64(0), nonEmptyErr},
        {"Int64", NewAny(json.Number("1")), int64(1), nil},
        {"Int64", NewAny(json.Number("1.1")), int64(1), nil},
        {"Int64", NewAny(json.Number("+1x.1")), int64(0), nonEmptyErr},
        {"Int64", NewAny([]byte{0}), int64(0), ErrUnsupportType},
        {"Int64", Node{}, int64(0), ErrUnsupportType},
        {"Int64", NewRaw("0"), int64(0), nil},
        {"Int64", NewRaw("null"), int64(0), nil},
        {"StrictInt64", NewRaw("true"), int64(0), ErrUnsupportType},
        {"StrictInt64", NewRaw("false"), int64(0), ErrUnsupportType},
        {"StrictInt64", NewAny(int(0)), int64(0), nil},
        {"StrictInt64", NewAny(int8(0)), int64(0), nil},
        {"StrictInt64", NewAny(int16(0)), int64(0), nil},
        {"StrictInt64", NewAny(int32(0)), int64(0), nil},
        {"StrictInt64", NewAny(int64(0)), int64(0), nil},
        {"StrictInt64", NewAny(uint(0)), int64(0), nil},
        {"StrictInt64", NewAny(uint8(0)), int64(0), nil},
        {"StrictInt64", NewAny(uint32(0)), int64(0), nil},
        {"StrictInt64", NewAny(uint64(0)), int64(0), nil},
        {"StrictInt64", Node{}, int64(0), ErrUnsupportType},
        {"StrictInt64", NewRaw("0"), int64(0), nil},
        {"StrictInt64", NewRaw("null"), int64(0), ErrUnsupportType},
        {"Float64", NewRaw("true"), float64(1), nil},
        {"Float64", NewRaw("false"), float64(0), nil},
        {"Float64", NewRaw("\"1.0\""), float64(1.0), nil},
        {"Float64", NewRaw("\"xx\""), float64(0), nonEmptyErr},
        {"Float64", Node{}, float64(0), ErrUnsupportType},
        {"Float64", NewAny(false), float64(0), nil},
        {"Float64", NewAny(true), float64(1), nil},
        {"Float64", NewAny(int(1)), float64(1), nil},
        {"Float64", NewAny(int8(1)), float64(1), nil},
        {"Float64", NewAny(int16(1)), float64(1), nil},
        {"Float64", NewAny(int32(1)), float64(1), nil},
        {"Float64", NewAny(int64(1)), float64(1), nil},
        {"Float64", NewAny(uint(1)), float64(1), nil},
        {"Float64", NewAny(uint8(1)), float64(1), nil},
        {"Float64", NewAny(uint32(1)), float64(1), nil},
        {"Float64", NewAny(uint64(1)), float64(1), nil},
        {"Float64", NewAny(float32(1)), float64(1), nil},
        {"Float64", NewAny(float64(1)), float64(1), nil},
        {"Float64", NewAny("1.1"), float64(1.1), nil},
        {"Float64", NewAny("+1x.1"), float64(0), nonEmptyErr},
        {"Float64", NewAny(json.Number("0")), float64(0), nil},
        {"Float64", NewAny(json.Number("x")), float64(0), nonEmptyErr},
        {"Float64", NewAny([]byte{0}), float64(0), ErrUnsupportType},
        {"Float64", NewRaw("0.0"), float64(0.0), nil},
        {"Float64", NewRaw("1"), float64(1.0), nil},
        {"Float64", NewRaw("null"), float64(0.0), nil},
        {"StrictFloat64", NewRaw("true"), float64(0), ErrUnsupportType},
        {"StrictFloat64", NewRaw("false"), float64(0), ErrUnsupportType},
        {"StrictFloat64", Node{}, float64(0), ErrUnsupportType},
        {"StrictFloat64", NewAny(float32(0)), float64(0), nil},
        {"StrictFloat64", NewAny(float64(0)), float64(0), nil},
        {"StrictFloat64", NewRaw("0.0"), float64(0.0), nil},
        {"StrictFloat64", NewRaw("null"), float64(0.0), ErrUnsupportType},
        {"Number", Node{}, json.Number(""), ErrUnsupportType},
        {"Number", NewAny(false), json.Number("0"), nil},
        {"Number", NewAny(true), json.Number("1"), nil},
        {"Number", NewAny(int(1)), json.Number("1"), nil},
        {"Number", NewAny(int8(1)), json.Number("1"), nil},
        {"Number", NewAny(int16(1)), json.Number("1"), nil},
        {"Number", NewAny(int32(1)), json.Number("1"), nil},
        {"Number", NewAny(int64(1)), json.Number("1"), nil},
        {"Number", NewAny(uint(1)), json.Number("1"), nil},
        {"Number", NewAny(uint8(1)), json.Number("1"), nil},
        {"Number", NewAny(uint32(1)), json.Number("1"), nil},
        {"Number", NewAny(uint64(1)), json.Number("1"), nil},
        {"Number", NewAny(float32(1)), json.Number("1"), nil},
        {"Number", NewAny(float64(1)), json.Number("1"), nil},
        {"Number", NewAny("1.1"), json.Number("1.1"), nil},
        {"Number", NewAny("+1x.1"), json.Number(""), nonEmptyErr},
        {"Number", NewAny(json.Number("0")), json.Number("0"), nil},
        {"Number", NewAny(json.Number("x")), json.Number("x"), nil},
        {"Number", NewAny(json.Number("+1x.1")), json.Number("+1x.1"), nil},
        {"Number", NewAny([]byte{0}), json.Number(""), ErrUnsupportType},
        {"Number", NewRaw("x"), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("0.0"), json.Number("0.0"), nil},
        {"Number", NewRaw("\"1\""), json.Number("1"), nil},
        {"Number", NewRaw("\"1.1\""), json.Number("1.1"), nil},
        {"Number", NewRaw("\"0.x0\""), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("{]"), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("true"), json.Number("1"), nil},
        {"Number", NewRaw("false"), json.Number("0"), nil},
        {"Number", NewRaw("null"), json.Number("0"), nil},
        {"StrictNumber", NewRaw("true"), json.Number(""), ErrUnsupportType},
        {"StrictNumber", NewRaw("false"), json.Number(""), ErrUnsupportType},
        {"StrictNumber", Node{}, json.Number(""), ErrUnsupportType},
        {"StrictNumber", NewAny(json.Number("0")), json.Number("0"), nil},
        {"StrictNumber", NewRaw("0.0"), json.Number("0.0"), nil},
        {"StrictNumber", NewRaw("null"), json.Number(""), ErrUnsupportType},
        {"String", Node{}, "", ErrUnsupportType},
        {"String", NewAny(`\u263a`), `\u263a`, nil},
        {"String", NewRaw(`"\u263a"`), `☺`, nil},
        {"String", NewString(`\u263a`), `\u263a`, nil},
        {"String", NewRaw(`0.0`), "0.0", nil},
        {"String", NewRaw(`true`), "true", nil},
        {"String", NewRaw(`false`), "false", nil},
        {"String", NewRaw(`null`), "", nil},
        {"String", NewAny(false), "false", nil},
        {"String", NewAny(true), "true", nil},
        {"String", NewAny(int(1)), "1", nil},
        {"String", NewAny(int8(1)), "1", nil},
        {"String", NewAny(int16(1)), "1", nil},
        {"String", NewAny(int32(1)), "1", nil},
        {"String", NewAny(int64(1)), "1", nil},
        {"String", NewAny(uint(1)), "1", nil},
        {"String", NewAny(uint8(1)), "1", nil},
        {"String", NewAny(uint32(1)), "1", nil},
        {"String", NewAny(uint64(1)), "1", nil},
        {"String", NewAny(float32(1)), "1", nil},
        {"String", NewAny(float64(1)), "1", nil},
        {"String", NewAny("1.1"), "1.1", nil},
        {"String", NewAny("+1x.1"), "+1x.1", nil},
        {"String", NewAny(json.Number("0")), ("0"), nil},
        {"String", NewAny(json.Number("x")), ("x"), nil},
        {"String", NewAny([]byte{0}), (""), ErrUnsupportType},
        {"StrictString", Node{}, "", ErrUnsupportType},
        {"StrictString", NewAny(`\u263a`), `\u263a`, nil},
        {"StrictString", NewRaw(`"\u263a"`), `☺`, nil},
        {"StrictString", NewString(`\u263a`), `\u263a`, nil},
        {"StrictString", NewRaw(`0.0`), "", ErrUnsupportType},
        {"StrictString", NewRaw(`true`), "", ErrUnsupportType},
        {"StrictString", NewRaw(`false`), "", ErrUnsupportType},
        {"StrictString", NewRaw(`null`), "", ErrUnsupportType},
        {"Len", Node{}, 0, nil},
        {"Len", NewAny(0), 0, ErrUnsupportType},
        {"Len", NewNull(), 0, nil},
        {"Len", NewRaw(`"1"`), 1, nil},
        {"Len", NewRaw(`[1]`), 0, nil},
        {"Len", NewArray([]Node{NewNull()}), 1, nil},
        {"Len", lazyArray, 0, nil},
        {"Len", NewRaw(`{"a":1}`), 0, nil},
        {"Len", lazyObject, 0, nil},
        {"Cap", Node{}, 0, nil},
        {"Cap", NewAny(0), 0, ErrUnsupportType},
        {"Cap", NewNull(), 0, nil},
        {"Cap", NewRaw(`[1]`), _DEFAULT_NODE_CAP, nil},
        {"Cap", NewObject([]Pair{{"",NewNull()}}), 1, nil},
        {"Cap", NewRaw(`{"a":1}`), _DEFAULT_NODE_CAP, nil},
    }
    lazyArray.skipAllIndex()
    lazyObject.skipAllKey()
    cases = append(cases, 
        tcase{"Len", lazyObject, 17, nil},
        tcase{"Len", lazyObject, 17, nil},
        tcase{"Cap", lazyObject, _DEFAULT_NODE_CAP*2, nil},
        tcase{"Cap", lazyObject, _DEFAULT_NODE_CAP*2, nil},
    )

    for i, c := range cases {
        fmt.Println(i, c)
        rt := reflect.ValueOf(&c.node)
        m := rt.MethodByName(c.method)
        rets := m.Call([]reflect.Value{})
        if len(rets) != 2 {
            t.Error(i, rets)
        }
        if !reflect.DeepEqual(rets[0].Interface(), c.exp) {
            t.Error(i, rets[0].Interface(), c.exp)
        }
        v := rets[1].Interface();
        if c.err == nonEmptyErr {
            if reflect.ValueOf(v).IsNil() {
                t.Error(i, v)
            }
        } else if  v != c.err {
            t.Error(i, v)
        }
    }
}

func TestCheckError(t *testing.T) {
    empty := Node{}
    if !empty.Valid() || empty.Check() != nil || empty.Error() != "" {
        t.Fatal()
    }
 
    n := newRawNode("[hello]", types.V_ARRAY)
    n.parseRaw(false)
    if n.Check() != nil {
        t.Fatal(n.Check())
    }
    n = newRawNode("[hello]", types.V_ARRAY)
    n.parseRaw(true)
    p := NewParser("[hello]")
    p.noLazy = true
    p.skipValue = false
    _, x := p.Parse()
    if n.Error() != newSyntaxError(p.syntaxError(x)).Error() {
        t.Fatal(n.Check())
    }


    s, err := NewParser(`{"a":{}, "b":talse, "c":{}}`).Parse()
    if err != 0 {
        t.Fatal(err)
    }

    root := s.GetByPath()
    // fmt.Println(root.Check())
    a := root.Get("a")
    if a.Check() != nil {
        t.Fatal(a.Check())
    }
    c := root.Get("c")
    if c.Check() == nil {
        t.Fatal()
    }
    fmt.Println(c.Check())

    _, e := a.Properties()
    if e != nil {
        t.Fatal(e)
    }
    exist, e := a.Set("d", newRawNode("x", types.V_OBJECT))
    if exist || e != nil {
        t.Fatal(err)
    }
    if a.len() != 1 {
        t.Fail()
    }
    d := a.Get("d").Get("")
    if d.Check() == nil {
        t.Fatal(d) 
    }
    exist, e = a.Set("e", newRawNode("[}", types.V_ARRAY))
    if e != nil {
        t.Fatal(e)
    }
    if a.len() != 2 {
        t.Fail()
    }
    d = a.Index(1).Index(0)
    if d.Check() == nil {
        t.Fatal(d)
    }


    it, e := root.Interface()
    if e == nil {
        t.Fatal(it)
    }
    fmt.Println(e)
}

func TestIndex(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    status := root.GetByPath("statuses", 0)
    x, _ := status.Index(4).String()
    y, _ := status.Get("id_str").String()
    if x != y {
        t.Fail()
    }
}

func TestUnset(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    entities := root.GetByPath("statuses", 0, "entities")
    if !entities.Exists() {
        t.Fatal()
    }
    exist, err := entities.Unset("urls")
    if !exist || err != nil {
        t.Fatal()
    }
    e := entities.Get("urls")
    if e.Exists() {
        t.Fatal()
    }
    if entities.len() != 2 {
        t.Fatal()
    }

    entities.Set("urls", NewString("a"))
    e = entities.Get("urls")
    x, _ := e.String()
    if !e.Exists() || x != "a" {
        t.Fatal()
    }
    exist, err = entities.UnsetByIndex(entities.len()-1)
    if !exist || err != nil {
        t.Fatal()
    }
    e = entities.Get("urls")
    if e.Exists() {
        t.Fatal()
    }

    hashtags := entities.Get("hashtags").Index(0)
    hashtags.Set("text2", newRawNode(`{}`, types.V_OBJECT))
    exist, err = hashtags.Unset("indices")
    if !exist || err != nil || hashtags.len() != 2 {
        t.Fatal()
    }
    y, _ := hashtags.Get("text").String()
    if y != "freebandnames" {
        t.Fatal()
    }
    if hashtags.Get("text2").Type() != V_OBJECT {
        t.Fatal()
    }

    ums := entities.Get("user_mentions")
    ums.Add(NewNull())
    ums.Add(NewBool(true))
    ums.Add(NewBool(false))
    if ums.len() != 3 {
        t.Fatal()
    }
    exist, err = ums.UnsetByIndex(1)
    if !exist || err != nil {
        t.Fatal()
    }
    v1, _ := ums.Index(0).Interface()
    v2, _ := ums.Index(1).Interface()
    if v1 != nil || v2 != false {
        t.Fatal()
    } 

}

func TestUnsafeNode(t *testing.T) {
    str, loop := getTestIteratorSample(_DEFAULT_NODE_CAP)

    root, err := NewSearcher(str).GetByPath("array")
    if err != nil {
        t.Fatal(err)
    }
    a, _ := root.UnsafeArray()
    if len(a) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(a))
    }
    for i := int64(0); i<int64(loop); i++{
        in := a[i]
        x, _ := in.Int64()
        if x != i {
            t.Fatalf("exp:%v, got:%v", i, x)
        }
    }
    
    root, err = NewSearcher(str).GetByPath("object")
    if err != nil {
        t.Fatal(err)
    }
    b, _ := root.UnsafeMap()
    if len(b) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(b))
    }
    for i := int64(0); i<int64(loop); i++ {
        k := `k`+strconv.Itoa(int(i))
        if k != b[i].Key {
            t.Fatalf("unexpected element: %#v", b[i])
        }
        x, _ := b[i].Value.Int64()
        if x != i {
            t.Fatalf("exp:%v, got:%v", i, x)
        }
    }
}

func TestUseNode(t *testing.T) {
    str, loop := getTestIteratorSample(_DEFAULT_NODE_CAP)
    root, e := NewParser(str).Parse()
    if e != 0 {
        t.Fatal(e)
    }
    _, er := root.InterfaceUseNode()
    if er != nil {
        t.Fatal(er)
    }

    root, err := NewSearcher(str).GetByPath("array")
    if err != nil {
        t.Fatal(err)
    }
    a, _ := root.ArrayUseNode()
    if len(a) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(a))
    }
    for i := int64(0); i<int64(loop); i++{
        in := a[i]
        a, _ := in.Int64()
        if a != i {
            t.Fatalf("exp:%v, got:%v", i, a)
        }
    }

    root, err = NewSearcher(str).GetByPath("array")
    if err != nil {
        t.Fatal(err)
    }
    x, _ := root.InterfaceUseNode()
    a = x.([]Node)
    if len(a) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(a))
    }
    for i := int64(0); i<int64(loop); i++{
        in := a[i]
        a, _ := in.Int64()
        if a != i {
            t.Fatalf("exp:%v, got:%v", i, a)
        }
    }
    
    root, err = NewSearcher(str).GetByPath("object")
    if err != nil {
        t.Fatal(err)
    }
    b, _ := root.MapUseNode()
    if len(b) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(b))
    }
    for i := int64(0); i<int64(loop); i++ {
        k := `k`+strconv.Itoa(int(i))
        xn, ok := b[k]
        if !ok {
            t.Fatalf("unexpected element: %#v", xn)
        }
        a, _ := xn.Int64()
        if a != i {
            t.Fatalf("exp:%v, got:%v", i, a)
        }
    }
    
    root, err = NewSearcher(str).GetByPath("object")
    if err != nil {
        t.Fatal(err)
    }
    x, _ = root.InterfaceUseNode()
    b = x.(map[string]Node)
    if len(b) != loop {
        t.Fatalf("exp:%v, got:%v", loop, len(b))
    }
    for i := int64(0); i<int64(loop); i++ {
        k := `k`+strconv.Itoa(int(i))
        xn, ok := b[k]
        if !ok {
            t.Fatalf("unexpected element: %#v", xn)
        }
        a, _ := xn.Int64()
        if a != i {
            t.Fatalf("exp:%v, got:%v", i, a)
        }
    }
}

func TestUseNumber(t *testing.T) {
    str, _ := getTestIteratorSample(_DEFAULT_NODE_CAP)
    root, e := NewParser(str).Parse()
    if e != 0 {
        t.Fatal(e)
    }
    _, er := root.InterfaceUseNumber()
    if er != nil {
        t.Fatal(er)
    }

    node, err := NewParser("1061346755812312312313").Parse()
    if err != 0 {
        t.Fatal(err)
    }
    if node.Type() != V_NUMBER {
        t.Fatalf("wrong type: %v", node.Type())
    }
    x, _ := node.InterfaceUseNumber()
    iv := x.(json.Number)
    if iv.String() != "1061346755812312312313" {
        t.Fatalf("exp:%#v, got:%#v", "1061346755812312312313", iv.String())
    }
    x, _ = node.Interface()
    ix := x.(float64)
    if ix != float64(1061346755812312312313) {
        t.Fatalf("exp:%#v, got:%#v", float64(1061346755812312312313), ix)
    }
    xj, _ := node.Number()
    ij, _ := xj.Int64()
    jj, _ := json.Number("1061346755812312312313").Int64()
    if ij != jj {
        t.Fatalf("exp:%#v, got:%#v", jj, ij)
    }
}

func TestMap(t *testing.T) {
    node, err := NewParser(`{"a":-0, "b":1, "c":-1.2, "d":-1.2e-10}`).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    m, _ := node.Map()
    assert.Equal(t, m, map[string]interface{}{
        "a": float64(0),    
        "b": float64(1),    
        "c": -1.2,
        "d": -1.2e-10,
    })
    m1, _ := node.MapUseNumber()
    assert.Equal(t, m1, map[string]interface{}{
        "a": json.Number("-0"),    
        "b": json.Number("1"),    
        "c": json.Number("-1.2"),    
        "d": json.Number("-1.2e-10"),    
    })
}

func TestArray(t *testing.T) {
    node, err := NewParser(`[-0, 1, -1.2, -1.2e-10]`).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    m, _ := node.Array()
    assert.Equal(t, m, []interface{}{
        float64(0),    
        float64(1),
        -1.2,
        -1.2e-10,
    })
    m1, _ := node.ArrayUseNumber()
    assert.Equal(t, m1, []interface{}{
        json.Number("-0"),    
        json.Number("1"),    
        json.Number("-1.2"),    
        json.Number("-1.2e-10"),    
    })
}

func TestNodeRaw(t *testing.T) {
    root, derr := NewSearcher(_TwitterJson).GetByPath("search_metadata")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ := root.Raw()
    var comp = `{
    "max_id": 250126199840518145,
    "since_id": 24012619984051000,
    "refresh_url": "?since_id=250126199840518145&q=%23freebandnames&result_type=mixed&include_entities=1",
    "next_results": "?max_id=249279667666817023&q=%23freebandnames&count=4&include_entities=1&result_type=mixed",
    "count": 4,
    "completed_in": 0.035,
    "since_id_str": "24012619984051000",
    "query": "%23freebandnames",
    "max_id_str": "250126199840518145"
  }`
    if val != comp {
        t.Fatalf("exp: %+v, got: %+v", comp, val)
    }

    root, derr = NewSearcher(_TwitterJson).GetByPath("statuses", 0, "entities", "hashtags")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ = root.Raw()
    comp = `[
          {
            "text": "freebandnames",
            "indices": [
              20,
              34
            ]
          }
        ]`
    if val != comp {
        t.Fatalf("exp: \n%s\n, got: \n%s\n", comp, val)
    }

    var data = `{"k1" :   { "a" : "bc"} , "k2" : [1,2 ] , "k3":{} , "k4":[]}`
    root, derr = NewSearcher(data).GetByPath("k1")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ = root.Raw()
    comp = `{ "a" : "bc"}`
    if val != comp {
        t.Fatalf("exp: %+v, got: %+v", comp, val)
    }

    root, derr = NewSearcher(data).GetByPath("k2")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ = root.Raw()
    comp = `[1,2 ]`
    if val != comp {
        t.Fatalf("exp: %+v, got: %+v", comp, val)
    }

    root, derr = NewSearcher(data).GetByPath("k3")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ = root.Raw()
    comp = `{}`
    if val != comp {
        t.Fatalf("exp: %+v, got: %+v", comp, val)
    }

    root, derr = NewSearcher(data).GetByPath("k4")
    if derr != nil {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ = root.Raw()
    comp = `[]`
    if val != comp {
        t.Fatalf("exp: %+v, got: %+v", comp, val)
    }
}

func TestNodeGet(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ := root.Get("search_metadata").Get("max_id").Int64()
    if val != int64(250126199840518145) {
        t.Fatalf("exp: %+v, got: %+v", 250126199840518145, val)
    }
}

func TestNodeIndex(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ := root.Get("statuses").Index(3).Get("id_str").String()
    if val != "249279667666817024" {
        t.Fatalf("exp: %+v, got: %+v", "249279667666817024", val)
    }
}

func TestNodeGetByPath(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    val, _ := root.GetByPath("statuses", 3, "id_str").String()
    if val != "249279667666817024" {
        t.Fatalf("exp: %+v, got: %+v", "249279667666817024", val)
    }
}

func TestNodeSet(t *testing.T) {
    empty := Node{}
    err := empty.Add(Node{})
    if err != nil {
        t.Fatal(err)
    }
    empty2 := empty.Index(0)
    if empty2.Check() != nil {
        t.Fatal(err)
    }
    exist, err := empty2.SetByIndex(1, Node{})
    if exist || err == nil {
        t.Fatal(exist, err)
    }
    empty3 := empty.Index(0)
    if empty3.Check() != nil {
        t.Fatal(err)
    }
    exist, err = empty3.Set("a", NewNumber("-1"))
    if exist || err != nil {
        t.Fatal(exist, err)
    }
    if n, e := empty.Index(0).Get("a").Int64(); e != nil || n != -1 {
        t.Fatal(n, e)
    }

    empty = NewNull()
    err = empty.Add(NewNull())
    if err != nil {
        t.Fatal(err)
    }
    empty2 = empty.Index(0)
    if empty2.Check() != nil {
        t.Fatal(err)
    }
    exist, err = empty2.SetByIndex(1, NewNull())
    if exist || err == nil {
        t.Fatal(exist, err)
    }
    empty3 = empty.Index(0)
    if empty3.Check() != nil {
        t.Fatal(err)
    }
    exist, err = empty3.Set("a", NewNumber("-1"))
    if exist || err != nil {
        t.Fatal(exist, err)
    }
    if n, e := empty.Index(0).Get("a").Int64(); e != nil || n != -1 {
        t.Fatal(n, e)
    }
    
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    app,_ := NewParser("111").Parse()
    root.GetByPath("statuses", 3).Set("id_str", app)
    val, _ := root.GetByPath("statuses", 3, "id_str").Int64()
    if val != 111 {
        t.Fatalf("exp: %+v, got: %+v", 111, val)
    }
    for i := root.GetByPath("statuses", 3).cap(); i >= 0; i-- {
        root.GetByPath("statuses", 3).Set("id_str"+strconv.Itoa(i), app)
    }
    val, _ = root.GetByPath("statuses", 3, "id_str0").Int64()
    if val != 111 {
        t.Fatalf("exp: %+v, got: %+v", 111, val)
    }

    nroot, derr := NewParser(`{"a":[0.1,true,0,"name",{"b":"c"}]}`).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    root.GetByPath("statuses", 3).Set("id_str2", nroot)
    val2, _ := root.GetByPath("statuses", 3, "id_str2", "a", 4, "b").String()
    if val2 != "c" {
        t.Fatalf("exp:%+v, got:%+v", "c", val2)
    }
}

func TestNodeAny(t *testing.T) {
    empty := Node{}
    _,err := empty.SetAny("any", map[string]interface{}{"a": []int{0}})
    if err != nil {
        t.Fatal(err)
    }
    if m, err := empty.Get("any").Interface(); err != nil {
        t.Fatal(err)
    } else if v, ok := m.(map[string]interface{}); !ok {
        t.Fatal(v)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]}}` {
        t.Fatal(string(buf))
    }
    if _, err := empty.Set("any2", Node{}); err != nil {
        t.Fatal(err)
    }
    if err := empty.Get("any2").AddAny(nil); err != nil {
        t.Fatal(err)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]},"any2":[null]}` {
        t.Fatal(string(buf))
    }
    if _, err := empty.Get("any2").SetAnyByIndex(0, NewNumber("-0.0")); err != nil {
        t.Fatal(err)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]},"any2":[-0.0]}` {
        t.Fatal(string(buf))
    }
}

func TestNodeSetByIndex(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    app, _ := NewParser("111").Parse()
    st := root.GetByPath("statuses")
    st.SetByIndex(0, app)
    st = root.GetByPath("statuses") 
    val := st.Index(0)
    x, _ := val.Int64()
    if x != 111 {
        t.Fatalf("exp: %+v, got: %+v", 111, val)
    }

    nroot, derr := NewParser(`{"a":[0.1,true,0,"name",{"b":"c"}]}`).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    root.GetByPath("statuses").SetByIndex(0, nroot)
    val2, _ := root.GetByPath("statuses", 0, "a", 4, "b").String()
    if val2 != "c" {
        t.Fatalf("exp:%+v, got:%+v", "c", val2)
    }
}

func TestNodeAdd(t *testing.T) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    app, _ := NewParser("111").Parse()

    for i := root.GetByPath("statuses").cap(); i >= 0; i-- {
        root.GetByPath("statuses").Add(app)
    }
    val, _ := root.GetByPath("statuses", 4).Int64()
    if val != 111 {
        t.Fatalf("exp: %+v, got: %+v", 111, val)
    }
    val, _ = root.GetByPath("statuses", root.GetByPath("statuses").len()-1).Int64()
    if val != 111 {
        t.Fatalf("exp: %+v, got: %+v", 111, val)
    }

    nroot, derr := NewParser(`{"a":[0.1,true,0,"name",{"b":"c"}]}`).Parse()
    if derr != 0 {
        t.Fatalf("decode failed: %v", derr.Error())
    }
    root.GetByPath("statuses").Add(nroot)
    val2, _ := root.GetByPath("statuses", root.GetByPath("statuses").len()-1, "a", 4, "b").String()
    if val2 != "c" {
        t.Fatalf("exp:%+v, got:%+v", "c", val2)
    }
}

func BenchmarkLoadNode(b *testing.B) {
    b.Run("Interface()", func(b *testing.B) {
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0)
                if err != nil {
                    b.Fatal(err)
                }
                _, _ = root.Interface()
            }
        })
    })

    b.Run("LoadAll()", func(b *testing.B) {
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0)
                if err != nil {
                    b.Fatal(err)
                }
                _ = root.LoadAll()
            }
        })
    })

    b.Run("InterfaceUseNode()", func(b *testing.B) {
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0)
                if err != nil {
                    b.Fatal(err)
                }
                _, _ = root.InterfaceUseNode()
            }
        })
    })

    b.Run("Load()", func(b *testing.B) {
        b.SetBytes(int64(len(_TwitterJson)))
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0)
                if err != nil {
                    b.Fatal(err)
                }
                _ = root.Load()
            }
        })
    })
}

func BenchmarkNodeGetByPath(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
}

func BenchmarkStructGetByPath(b *testing.B) {
    var root = _TwitterStruct{}
    err := json.Unmarshal([]byte(_TwitterJson), &root)
    if err != nil {
        b.Fatalf("unmarshal failed: %v", err)
    }

    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = root.Statuses[3].Entities.Hashtags[0].Text
        }
    })
}

func BenchmarkNodeIndex(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.Index(2)
    }
}

func BenchmarkStructIndex(b *testing.B) {
    type T struct {
        A Node
        B Node
        C Node
        D Node
        E Node
    }
    var obj = new(T)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = obj.C
    }
}

func BenchmarkSliceIndex(b *testing.B) {
    var obj = []Node{Node{},Node{},Node{},Node{},Node{}}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = obj[2]
    }
}

func BenchmarkMapIndex(b *testing.B) {
    var obj = map[string]interface{}{"test1":Node{}, "test2":Node{}, "test3":Node{}, "test4":Node{}, "test5":Node{}}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for k := range obj {
            if k == "test3" {
                break 
            }
        }
    }
}

func BenchmarkNodeGet(b *testing.B) {
    var N = 5
    var half = "test" + strconv.Itoa(N/2+1)
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    for i:=0; i<N; i++ {
        node.Set("test"+strconv.Itoa(i), NewNumber(strconv.Itoa(i)))
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = node.Get(half)
    }
}

func BenchmarkSliceGet(b *testing.B) {
    var obj = []string{"test1", "test2", "test3", "test4", "test5"}
    str := "test3"
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, k := range obj {
            if k == str {
                break 
            }
        }
    }
}

func BenchmarkMapGet(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    m, _ := node.Map()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = m["test3"]
    }
}

func BenchmarkNodeSet(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    n := NewNull()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.Set("test3", n)
    }
}

func BenchmarkMapSet(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    m, _ := node.Map()
    n := NewNull()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m["test3"] = n
    }
}

func BenchmarkNodeSetByIndex(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags")
    node.Add(NewNumber("1"))
    node.Add(NewNumber("2"))
    node.Add(NewNumber("3"))
    node.Add(NewNumber("4"))
    node.Add(NewNumber("5"))
    n := NewNull()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.SetByIndex(2, n)
    }
}

func BenchmarkSliceSetByIndex(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags")
    node.Add(NewNumber("1"))
    node.Add(NewNumber("2"))
    node.Add(NewNumber("3"))
    node.Add(NewNumber("4"))
    node.Add(NewNumber("5"))
    m, _ := node.Array()
    n := NewNull()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m[2] = n
    }
}

func BenchmarkStructSetByIndex(b *testing.B) {
    type T struct {
        A Node
        B Node
        C Node
        D Node
        E Node
    }
    var obj = new(T)
    n := NewNull()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        obj.C = n
    }
}

func BenchmarkNodeUnset(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.Unset("test3")
    }
}

func BenchmarkMapUnset(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    node.Set("test1", NewNumber("1"))
    node.Set("test2", NewNumber("2"))
    node.Set("test3", NewNumber("3"))
    node.Set("test4", NewNumber("4"))
    node.Set("test5", NewNumber("5"))
    m, _ := node.Map()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        delete(m, "test3")
    }
}

func BenchmarkNodUnsetByIndex(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags")
    node.Add(NewNumber("1"))
    node.Add(NewNumber("2"))
    node.Add(NewNumber("3"))
    node.Add(NewNumber("4"))
    node.Add(NewNumber("5"))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.UnsetByIndex(2)
    }
}

func BenchmarkSliceUnsetByIndex(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags")
    node.Add(NewNumber("1"))
    node.Add(NewNumber("2"))
    node.Add(NewNumber("3"))
    node.Add(NewNumber("4"))
    node.Add(NewNumber("5"))
    m, _ := node.Array()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for i:=3; i<5; i++ {
            m[i-1] = m[i]
        }
    }
}

func BenchmarkNodeAdd(b *testing.B) {
    n := NewObject([]Pair{{"test", NewNumber("1")}})
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node := NewArray([]Node{})
        node.Add(n)
    }
}

func BenchmarkSliceAdd(b *testing.B) {
    n := NewObject([]Pair{{"test", NewNumber("1")}})
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node := []Node{}
        node = append(node, n)
    }
}

func BenchmarkMapAdd(b *testing.B) {
    n := NewObject([]Pair{{"test", NewNumber("1")}})
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node := map[string]Node{}
        node["test3"] = n
    }
}