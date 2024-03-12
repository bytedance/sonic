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
    `runtime`
    `sync`
    `testing`
    `strings`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

func TestGC_Encode(t *testing.T) {
    if debugSyncGC {
        return
    }
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    root.LoadAll()
    _, err = root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
            defer wg.Done()
            root, err := NewSearcher(_TwitterJson).GetByPath()
            if err != nil {
                t.Error(err)
                return
            }
            root.Load()
            _, err = root.MarshalJSON()
            if err != nil {
                t.Error(err)
                return
            }
            runtime.GC()
        }(wg)
    }
    wg.Wait()
}

func TestEncodeValue(t *testing.T) {
    obj := new(_TwitterStruct)
    if err := json.Unmarshal([]byte(_TwitterJson), obj); err != nil {
        t.Fatal(err)
    }
    // buf, err := encoder.Encode(obj, encoder.EscapeHTML|encoder.SortMapKeys)
    buf, err := json.Marshal(obj)
    if err != nil {
        t.Fatal(err)
    }
    quote, err := json.Marshal(_TwitterJson)
    if err != nil {
        t.Fatal(err)
    }
    type Case struct {
        node Node
        exp string
        err bool
    }
    input := []Case{
        {NewNull(), "null", false},
        {NewBool(true), "true", false},
        {NewBool(false), "false", false},
        {NewNumber("0.0"), "0.0", false},
        {NewString(""), `""`, false},
        {NewString(`\"\"`), `"\\\"\\\""`, false},
        {NewString(_TwitterJson), string(quote), false},
        {NewArray([]Node{}), "[]", false},
        {NewArray([]Node{NewString(""), NewNull()}), `["",null]`, false},
        {NewArray([]Node{NewBool(true), NewString("true"), NewString("\t")}), `[true,"true","\t"]`, false},
        {NewObject([]Pair{Pair{"a", NewNull()}, Pair{"b", NewNumber("0")}}), `{"a":null,"b":0}`, false},
        {NewObject([]Pair{Pair{"\ta", NewString("\t")}, Pair{"\bb", NewString("\b")}, Pair{"\nb", NewString("\n")}, Pair{"\ra", NewString("\r")}}),`{"\ta":"\t","\u0008b":"\u0008","\nb":"\n","\ra":"\r"}`, false},
        {NewObject([]Pair{}), `{}`, false},
        {NewObject([]Pair{Pair{Key: "", Value: NewNull()}}), `{"":null}`, false},
        {NewBytes([]byte("hello, world")), `"aGVsbG8sIHdvcmxk"`, false},
        {NewAny(obj), string(buf), false},
        {NewRaw(`[{ }]`), "[{}]", false},
        {Node{}, "", true},
        {Node{t: types.ValueType(1)}, "", true},
    }
    for i, c := range input {
        t.Log(i)
        buf, err := json.Marshal(&c.node)
        if c.err {
            if err == nil {
                t.Fatal(i)
            }
            continue
        }
        if err != nil {
            t.Fatal(i, err)
        }
        assert.Equal(t, c.exp, string(buf))
    }
}

func TestEncodeNode(t *testing.T) {
    data := `{"a":[{},[],-0.1,true,false,null,""],"b":0,"c":true,"d":false,"e":null,"g":""}`
    root, e := NewSearcher(data).GetByPath()
    if e != nil {
        t.Fatal(root)
    }
    ret, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    if string(ret) != data {
        t.Fatal(string(ret))
    }
    root.skipAllKey()
    ret, err = root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    if string(ret) != data {
        t.Fatal(string(ret))
    }
    root.loadAllKey()
    ret, err = root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    if string(ret) != data {
        t.Fatal(string(ret))
    }
}

type SortableNode struct {
    sorted bool
	*Node
}

func (j *SortableNode) UnmarshalJSON(data []byte) (error) {
    j.Node = new(Node)
	return j.Node.UnmarshalJSON(data)
}

func (j *SortableNode) MarshalJSON() ([]byte, error) {
    if !j.sorted {
        j.Node.SortKeys(true)
        j.sorted = true
    }
	return j.Node.MarshalJSON()
}

func TestMarshalSort(t *testing.T) {
    var data = `{"d":3,"a":{"c":1,"b":2},"e":null}`
    var obj map[string]*SortableNode
    require.NoError(t, json.Unmarshal([]byte(data), &obj))
    out, err := json.Marshal(obj)
    require.NoError(t, err)
    require.Equal(t, `{"a":{"b":2,"c":1},"d":3,"e":null}`, string(out))
    out, err = json.Marshal(obj)
    require.NoError(t, err)
    require.Equal(t, `{"a":{"b":2,"c":1},"d":3,"e":null}`, string(out))
}

func BenchmarkEncodeRaw_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewSearcher(data).GetByPath()
    if e != nil {
        b.Fatal(root)
    }
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _, err := root.MarshalJSON()
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkEncodeSkip_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewParser(data).Parse()
    if e != 0 {
        b.Fatal(root)
    }
    root.skipAllKey()
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _, err := root.MarshalJSON()
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkEncodeLoad_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewParser(data).Parse()
    if e != 0 {
        b.Fatal(root)
    }
    root.loadAllKey()
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _, err := root.MarshalJSON()
        if err != nil {
            b.Fatal(err)
        }
    }
}

func TestEncodeNone(t *testing.T) {
    n := NewObject([]Pair{{Key:"a", Value:Node{}}})
    out, err := n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, "{}", string(out))
    n = NewObject([]Pair{{Key:"a", Value:NewNull()}, {Key:"b", Value:Node{}}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, `{"a":null}`, string(out))

    n = NewArray([]Node{Node{}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, "[]", string(out))
    n = NewArray([]Node{NewNull(), Node{}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, `[null]`, string(out))
}


type Path = []interface{}

type testGetApi struct {
    json      string
    path      Path
}

type checkError func(error) bool

func isSyntaxError(err error) bool {
    if err == nil {
        return false
    }
    return strings.HasPrefix(err.Error(), `"Syntax error at index`)
}

func isEmptySource(err error) bool {
    if err == nil {
        return false
    }
    return strings.Contains(err.Error(), "no sources available")
}

func isErrNotExist(err error) bool {
    return err == ErrNotExist
}

func isErrUnsupportType(err error) bool {
    return err == ErrUnsupportType
}

func testSyntaxJson(t *testing.T, json string, path ...interface{}) {
    search := NewSearcher(json)
    _, err := search.GetByPath(path...)
    assert.True(t, isSyntaxError(err))
}

func TestGetFromEmptyJson(t *testing.T) {
    tests := []testGetApi {
        { "", nil },
        { "", Path{}},
        { "", Path{""}},
        { "", Path{0}},
        { "", Path{"", ""}},
    }
    for _, test := range tests {
        f := func(t *testing.T) {
            search := NewSearcher(test.json)
            _, err := search.GetByPath(test.path...)
            assert.True(t, isEmptySource(err))
        }
        t.Run(test.json, f)
    }
}

func TestGetFromSyntaxError(t *testing.T) {
    tests := []testGetApi {
        { " \r\n\f\t", Path{} },
        { "123.", Path{} },
        { "+124", Path{} },
        { "-", Path{} },
        { "-e123", Path{} },
        { "-1.e123", Path{} },
        { "-12e456.1", Path{} },
        { "-12e.1", Path{} },
        { "[", Path{} },
        { "{", Path{} },
        { "[}", Path{} },
        { "{]", Path{} },
        { "{,}", Path{} },
        { "[,]", Path{} },
        { "tru", Path{} },
        { "fals", Path{} },
        { "nul", Path{} },
        { `{"a":"`, Path{"a"} },
        { `{"`, Path{} },
        { `"`, Path{} },
        { `"\"`, Path{} },
        { `"\\\"`, Path{} },
        { `"hello`, Path{} },
        { `{{}}`, Path{} },
        { `{[]}`, Path{} },
        { `{:,}`, Path{} },
        { `{test:error}`, Path{} },
        { `{":true}`, Path{} },
        { `{"" false}`, Path{} },
        { `{ "" : "false }`, Path{} },
        { `{"":"",}`, Path{} },
        { `{ " test : true}`, Path{} },
        { `{ "test" : tru }`, Path{} },
        { `{ "test" : true , }`, Path{} },
        { `{ {"test" : true , } }`, Path{} },
        { `{"test":1. }`, Path{} },
        { `{"\\\""`, Path{} },
        { `{"\\\"":`, Path{} },
        { `{"\\\":",""}`, Path{} },
        { `[{]`, Path{} },
        { `[tru]`, Path{} },
        { `[-1.]`, Path{} },
        { `[[]`, Path{} },
        { `[[],`, Path{} },
        { `[ true , false , [ ]`, Path{} },
        { `[true, false, [],`, Path{} },
        { `[true, false, [],]`, Path{} },
        { `{"key": [true, false, []], "key2": {{}}`, Path{} },
    }

    for _, test := range tests {
        f := func(t *testing.T) {
            testSyntaxJson(t, test.json, test.path...)
            path := append(Path{"key"}, test.path...)
            testSyntaxJson(t, `{"key":` + test.json, path...)
            path  = append(Path{""}, test.path...)
            testSyntaxJson(t, `{"":` + test.json, path...)
            path  = append(Path{1}, test.path...)
            testSyntaxJson(t, `["",` + test.json, path...)
        }
        t.Run(test.json, f)
    }
}

// NOTE: GetByPath API not validate the undemanded fields for performance.
func TestGetWithInvalidUndemandedField(t *testing.T) {
    type Any = interface{}
    tests := []struct {
        json string
        path Path
        exp  Any
    } {
        { "-0xyz", Path{}, Any(float64(-0))},
        { "-12e4xyz", Path{}, Any(float64(-12e4))},
        { "truex",  Path{}, Any(true)},
        { "false,", Path{}, Any(false)},
        { `{"a":{,xxx},"b":true}`, Path{"b"}, Any(true)},
        { `{"a":[,xxx],"b":true}`, Path{"b"}, Any(true)},
    }

    for _, test := range tests {
        f := func(t *testing.T) {
            search := NewSearcher(test.json)
            node, err := search.GetByPath(test.path...)
            assert.NoError(t, err)
            v, err := node.Interface()
            assert.NoError(t, err)
            assert.Equal(t, v, test.exp)
        }
        t.Run(test.json, f)
    }
}

func TestGet_InvalidPathType(t *testing.T) {
    assert.Panics(t, assert.PanicTestFunc(func() {
        data := `{"a":[{"b":true}]}`
        s := NewSearcher(data)
        s.GetByPath("a", true)

        s = NewSearcher(data)
        s.GetByPath("a", nil)

        s = NewSearcher(data)
        s.GetByPath("a", -1)
    }))
}
