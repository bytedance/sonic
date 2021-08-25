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
	`fmt`
	`reflect`
	`strconv`
	`testing`

	`github.com/bytedance/sonic/internal/native/types`
	jsoniter `github.com/json-iterator/go`
	`github.com/stretchr/testify/assert`
)

var parallelism = 4

func TestTypeCast(t *testing.T) {
    type tcase struct {
        method string
        node Node
        exp interface{}
        err error
    }
    lazyArray, _ := NewParser("[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]").Parse()
    lazyObject, _ := NewParser(`{"0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"10":10,"11":11,"12":12,"13":13,"14":14,"15":15,"16":16}`).Parse()
    var cases = []tcase{
        {"Raw", Node{}, "", ErrUnsupportType},
        {"Raw", newRawNode("[ ]", types.V_ARRAY), "[ ]", nil},
        {"Bool", Node{}, false, ErrNotExist},
        {"Bool", newRawNode("true", types.V_TRUE), true, nil},
        {"Bool", newRawNode("false", types.V_FALSE), false, nil},
        {"Int64", Node{}, int64(0), ErrNotExist},
        {"Int64", newRawNode("0", _V_NUMBER), int64(0), nil},
        {"Float64", Node{}, float64(0), ErrNotExist},
        {"Float64", newRawNode("0.0", _V_NUMBER), float64(0.0), nil},
        {"Number", Node{}, json.Number(""), ErrNotExist},
        {"Number", newRawNode("0.0", _V_NUMBER), json.Number("0.0"), nil},
        {"Number", newRawNode("true", types.V_TRUE), json.Number("1"), nil},
        {"Number", newRawNode("false", types.V_FALSE), json.Number("0"), nil},
        {"String", Node{}, "", ErrNotExist},
        {"String", newRawNode(`""`, types.V_STRING), ``, nil},
        {"String", newRawNode(`0.0`, _V_NUMBER), "0.0", nil},
        {"String", newRawNode(`null`, types.V_NULL), "null", nil},
        {"String", newRawNode(`true`, types.V_TRUE), "true", nil},
        {"String", newRawNode(`false`, types.V_FALSE), "false", nil},
        {"Len", NewNull(), 0, ErrUnsupportType},
        {"Len", newRawNode(`"1"`, types.V_STRING), 1, nil},
        {"Len", newRawNode(`[1]`, types.V_ARRAY), 0, nil},
        {"Len", NewArray([]Node{NewNull()}), 1, nil},
        {"Len", lazyArray, 0, nil},
        {"Len", newRawNode(`{"a":1}`, types.V_OBJECT), 0, nil},
        {"Len", lazyObject, 0, nil},
        {"Cap", NewNull(), 0, ErrUnsupportType},
        {"Cap", newRawNode(`[1]`, types.V_ARRAY), _DEFAULT_NODE_CAP, nil},
        {"Cap", NewObject([]Pair{{"",NewNull()}}), 1, nil},
        {"Cap", newRawNode(`{"a":1}`, types.V_OBJECT), _DEFAULT_NODE_CAP, nil},
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
        fmt.Println(c)
        rt := reflect.ValueOf(&c.node)
        m := rt.MethodByName(c.method)
        rets := m.Call([]reflect.Value{})
        if len(rets) != 2 {
            t.Fatal(i, rets)
        }
        if rets[0].Interface() != c.exp {
            t.Fatal(i, rets[0].Interface())
        }
        if rets[1].Interface() != c.err {
            t.Fatal(i, rets[1].Interface())
        }
    }
}

func TestCheckError(t *testing.T) {
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
    str, loop := getTestIteratorSample()

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
    str, loop := getTestIteratorSample()
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
    str, _ := getTestIteratorSample()
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

func BenchmarkNodeRaw(b *testing.B) {
    root, derr := NewSearcher(_TwitterJson).GetByPath("search_metadata")
    if derr != nil {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    b.SetParallelism(parallelism)
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            root.Raw()
        }
    })
}

func BenchmarkNodeGetByPath(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
    b.SetParallelism(parallelism)
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
}

func BenchmarkStructGetByPath_Jsoniter(b *testing.B) {
    var root = _TwitterStruct{}
    err := jsoniter.Unmarshal([]byte(_TwitterJson), &root)
    if err != nil {
        b.Fatalf("unmarshal failed: %v", err)
    }

    b.SetParallelism(parallelism)
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = root.Statuses[3].Entities.Hashtags[0].Text
        }
    })
}

func BenchmarkNodeGet(b *testing.B) {
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
        node.Get("text")
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
        _ = m["text"]
    }
}

func BenchmarkNodeSet(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0)
    b.SetParallelism(parallelism)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node.Set("test1", NewNumber("1"))
    }
}

func BenchmarkMapSet(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    node, _ := root.Get("statuses").Index(3).Get("entities").Get("hashtags").Index(0).Map()
    b.SetParallelism(parallelism)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node["test1"] = map[string]int{"test1": 1}
    }
}

func BenchmarkNodeAdd(b *testing.B) {
    data := `{"statuses":[]}`
    _, derr := NewParser(data).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    b.SetParallelism(parallelism)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        root, _ := NewParser(data).Parse()
        node := root.Get("statuses")
        node.Add(NewObject([]Pair{{"test", NewNumber("1")}}))
    }
}

func BenchmarkSliceAdd(b *testing.B) {
    data := `{"statuses":[]}`
    _, derr := NewParser(data).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    b.SetParallelism(parallelism)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        root, _ := NewParser(data).Parse()
        node, _ := root.Get("statuses").Array()
        node = append(node, map[string]interface{}{"test": 1})
    }
}
