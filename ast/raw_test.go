package ast

import (
    `encoding/json`
    `errors`
    `fmt`
    `reflect`
    `strconv`
    `sync`
    `testing`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`
)

var concurrency = 1000

func TestCast(t *testing.T) {
    v, err := strconv.ParseInt("1.1", 0, 0)
    spew.Dump(v, err)
}

func TestSpace(t *testing.T) {
    v, err := NewSearcher(`[ true , 1 , 1.1 , false , null , "a" , [ 2 ] , { "a" : 2 } , 2 ]`).GetValueByPath()
    require.NoError(t, err)

    vv, e := v.Index(0).Interface()
    require.NoError(t, e)
    require.Equal(t, true, vv)

    vv, e = v.Index(1).Int64()
    require.NoError(t, e)
    require.Equal(t, int64(1), vv)

    vv, e = v.Index(2).Interface()
    require.NoError(t, e)
    require.Equal(t, float64(1.1), vv)

    vv, e = v.Index(3).Interface()
    require.NoError(t, e)
    require.Equal(t, false, vv)

    vv, e = v.Index(4).Interface()
    require.NoError(t, e)
    require.Equal(t, nil, vv)

    vv, e = v.Index(5).Interface()
    require.NoError(t, e)
    require.Equal(t, "a", vv)

    vv, e = v.Index(6).Interface()
    require.NoError(t, e)
    require.Equal(t, []interface{}{float64(2)}, vv)

    vv, e = v.Index(7).Interface()
    require.NoError(t, e)
    require.Equal(t, map[string]interface{}{"a":float64(2)}, vv)
}

func TestValueAPI(t *testing.T) {
    var nonEmptyErr error = errors.New("")

    var cases = []struct {
        method string
        js     Node
        exp    interface{}
        err error
    }{
        {"Len", NewAny(""), 0, nil},
        {"Len", NewAny("a"), 1,nil},
        {"Len", NewAny(1), -1, ErrUnsupportType},
        {"Len", NewAny(true), -1,ErrUnsupportType},
        {"Len", NewAny(nil), -1,ErrUnsupportType},
        {"Len", NewAny([]int{}), 0,nil},
        {"Len", NewAny([]int{1}), 1,nil},
        {"Len", NewAny([]int{1,2}), 2,nil},
        {"Len", NewAny(map[string]string{}), 0,nil},
        {"Len", NewAny(map[string]int{"a":1}), 1,nil},
        {"Len", NewAny(map[string]int{"a":1,"b":2}), 2,nil},
        {"Len", NewRaw(`{"a":[1,2]}`), 1, nil},
        {"Len", NewRaw(`[1,{"a":2}]`), 2, nil},
        {"Interface", Node{}, interface{}(nil), nonEmptyErr},
        {"Interface", NewNumber("1"), float64(1), nil},
        {"Interface", NewAny(int64(1)), float64(1), nil},
        {"Interface", NewAny(float64(1)), float64(1), nil},
        {"InterfaceUseNode", NewNumber("1"), float64(1), nil},
        {"InterfaceUseNode", NewAny(int64(1)), float64(1), nil},
        {"InterfaceUseNode", NewAny(float64(1)), float64(1), nil},
        {"InterfaceUseNode", NewRaw(`[1]`), []Value{NewValue(1)}, nil},
        {"InterfaceUseNode", NewRaw(`{"a":1}`), map[string]Value{"a":NewValue(1)}, nil},
        {"InterfaceUseNode", NewRaw(`{"a":1,"b":[2]}`), map[string]Value{"a":NewValue(1), "b":NewValueJSON(`[2]`)}, nil},
        {"Bool", Node{}, false, nonEmptyErr},
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
        {"Int64", NewRaw("\"1.1\""), int64(0), nonEmptyErr},
        {"Int64", NewRaw("\"1.0\""), int64(0), nonEmptyErr},
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
        {"Int64", NewAny("1.1"), int64(0), nonEmptyErr},
        {"Int64", NewAny("+1x.1"), int64(0), nonEmptyErr},
        {"Int64", NewAny(json.Number("1")), int64(1), nil},
        {"Int64", NewAny(json.Number("1.1")), int64(0), nonEmptyErr},
        {"Int64", NewAny(json.Number("+1x.1")), int64(0), nonEmptyErr},
        {"Int64", NewAny([]byte{0}), int64(0), nonEmptyErr},
        {"Int64", Node{}, int64(0), nonEmptyErr},
        {"Int64", NewRaw("0"), int64(0), nil},
        {"Int64", NewRaw("null"), int64(0), nil},
        {"Float64", NewRaw("true"), float64(1), nil},
        {"Float64", NewRaw("false"), float64(0), nil},
        {"Float64", NewRaw("\"1.0\""), float64(1.0), nil},
        {"Float64", NewRaw("\"xx\""), float64(0), nonEmptyErr},
        {"Float64", Node{}, float64(0), nonEmptyErr},
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
        {"Float64", NewAny([]byte{0}), float64(0), nonEmptyErr},
        {"Float64", NewRaw("0.0"), float64(0.0), nil},
        {"Float64", NewRaw("1"), float64(1.0), nil},
        {"Float64", NewRaw("null"), float64(0.0), nil},
        {"Number", Node{}, json.Number(""), nonEmptyErr},
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
        {"Number", NewAny(json.Number("x")), json.Number(""), nonEmptyErr},
        {"Number", NewAny(json.Number("+1x.1")), json.Number(""), nonEmptyErr},
        {"Number", NewAny([]byte{0}), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("x"), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("0.0"), json.Number("0.0"), nil},
        {"Number", NewRaw("\"1\""), json.Number("1"), nil},
        {"Number", NewRaw("\"1.1\""), json.Number("1.1"), nil},
        {"Number", NewRaw("\"0.x0\""), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("{]"), json.Number(""), nonEmptyErr},
        {"Number", NewRaw("true"), json.Number("1"), nil},
        {"Number", NewRaw("false"), json.Number("0"), nil},
        {"Number", NewRaw("null"), json.Number("0"), nil},
        {"String", Node{}, "", nonEmptyErr},
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
        {"String", NewAny(json.Number("x")), (""), nonEmptyErr},
        {"String", NewAny([]byte{0}), "AA==", nil},
    }
    for i, c := range cases {
        fmt.Println(i)
        node := NewValue(&c.js)
        println(node.js, node.t)
        rt := reflect.ValueOf(&node)
        m := rt.MethodByName(c.method)
        rets := m.Call([]reflect.Value{})
        if len(rets) != 2 {
            t.Fatal(i, rets)
        }
        if !reflect.DeepEqual(rets[0].Interface(), c.exp) {
            t.Fatal(i, c.exp, "::", rets[0].Interface())
        }
        v := rets[1].Interface()
        if c.err == nonEmptyErr {
            if rets[1].IsNil() {
                t.Fatal(i, v)
            }
        } else if v != c.err {
            t.Fatal(i, v)
        }
    }
}

// IndexVal
type IndexVal struct {
    Index int
    Val Value
}

type keyVal struct {
    Key string
    Val Value
}

func TestGetMany(t *testing.T) {
    var cases = []struct {
        name string
        js   string
        kvs  interface{} // []KeyVal or []IndexVal
        err  error
    }{
        {"Get fail", `{}`, []keyVal{{"b", Value{}}}, nil},
        {"Get 0", `{"a":1, "b":2, "c":3}`, []keyVal{}, nil},
        {"Get 1", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", value(`2`)}}, nil},
        {"Get 2", `{"a":1, "b":2, "c":3}`, []keyVal{{"a", value(`1`)}, {"c", value(`3`)}}, nil},
        {"Get 2", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", value(`2`)}, {"a", value(`1`)}}, nil},
        {"Get 3", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", value(`2`)}, {"c", value(`3`)}, {"a", value(`1`)}}, nil},
        {"Get 3", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", value(`2`)}, {"c", value(`3`)}, {"d", Value{}}, {"a", value(`1`)}}, nil},
        {"Index fail", `[]`, []IndexVal{{1, Value{}}}, nil},
        {"Index 0", `[1, 2, 3]`, []IndexVal{{1, value(`2`)}}, nil},
        {"Index 1", `[1, 2, 3]`, []IndexVal{{1, value(`2`)}}, nil},
        {"Index 2", `[1, 2, 3]`, []IndexVal{{0, value(`1`)}, {2, value(`3`)}}, nil},
        {"Index 2", `[1, 2, 3]`, []IndexVal{{1, value(`2`)}, {0, value(`1`)}}, nil},
        {"Index 3", `[1, 2, 3]`, []IndexVal{{1, value(`2`)}, {2, value(`3`)}, {0, value(`1`)}}, nil},
        {"Index 3", `[1, 2, 3]`, []IndexVal{{1, value(`2`)}, {2, value(`3`)}, {3, Value{}}, {0, value(`1`)}}, nil},
    }
    for i, c := range cases {
        fmt.Println(i, c)
        node := NewValueJSON(c.js)
        var err error
        if kvs, ok := c.kvs.([]keyVal); ok {
            keys := []string{}
            vals := []Value{}
            for _, kv := range kvs {
                keys = append(keys, kv.Key)
                vals = append(vals, kv.Val)
            }
            act := make([]Value, len(keys))
            err = node.GetMany(keys, act)
            require.Equal(t, vals, act)
        } else  if ids, ok := c.kvs.([]IndexVal); ok {
            keys := []int{}
            vals := []Value{}
            for _, kv := range ids {
                keys = append(keys, kv.Index)
                vals = append(vals, kv.Val)
            }
            act := make([]Value, len(keys))
            err = node.IndexMany(keys, act)
            require.Equal(t, vals, act)
        }
        if err != nil && c.err.Error() != err.Error() {
            t.Fatal(err)
        }
    }
}

func TestSetMany(t *testing.T) {
    var cases = []struct{
        name string
        js string
        kvs interface{}
        exp string
        err string
    }{
        {"replace 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", value(`11`)}}, `{ "a" : 11, "b" : 2}`, ""},
        {"replace 2", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", value(`11`)}, {"b", value(`22`)}}, `{ "a" : 11, "b" : 22}`, ""},
        {"replace repeated", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", value(`11`)}, {"a", value(`22`)}}, `{ "a" : 11, "b" : 2,"a":22}`, ""},
        {"insert empty", `{ } `, []keyVal{{"c", value(`33`)}}, `{ "c":33}`, ""},
        {"insert repeated", `{ } `, []keyVal{{"c", value(`33`)}, {"c", value(`33`)}}, `{ "c":33,"c":33}`, ""},
        {"insert 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"c", value(`33`)}}, `{ "a" : 1, "b" : 2,"c":33}`, ""},
        {"insert 2", `{ "a" : 1, "b" : 2} `, []keyVal{{"c", value(`33`)},{"d", value(`44`)}}, `{ "a" : 1, "b" : 2,"c":33,"d":44}`, ""},
        {"replace 1, insert 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", value(`11`)}, {"c", value(`33`)}}, `{ "a" : 11, "b" : 2,"c":33}`, ""},
        
        {"replace 1", `[ 1, 2] `, []IndexVal{{0, value(`11`)}}, `[ 11, 2]`, ""},
        {"replace 2", `[ 1, 2] `, []IndexVal{{0, value(`11`)}, {0, value(`22`)}}, `[ 11, 2]`, "value not exists"},
        {"replace repeated", `[ 1, 2] `, []IndexVal{{0, value(`11`)}, {1, value(`22`)}}, `[ 11, 22]`, ""},
        {"insert empty", `[ ] `, []IndexVal{{2, value(`33`)}}, `[ ]`, "value not exists"},
        {"insert 1", `[ 1, 2] `, []IndexVal{{2, value(`33`)}}, `[ 1, 2]`, "value not exists"},
        {"insert 2", `[ 1, 2] `, []IndexVal{{2, value(`33`)},{3, value(`44`)}}, `[ 1, 2]`, "value not exists"},
        {"insert repeated", `[ 1, 2] `, []IndexVal{{2, value(`33`)},{2, value(`44`)}}, `[ 1, 2]`, "value not exists"},
        {"replace 1, insert 1", `[ 1, 2] `, []IndexVal{{0, value(`11`)}, {2, value(`33`)}}, `[ 11, 2]`, "value not exists"},
    }

    for i, c := range cases {
        println(i, c.name)
        node := NewValueJSON(c.js)
        var err error
        if kvs, ok := c.kvs.([]keyVal); ok {
            keys := []string{}
            vals := []Value{}
            for _, kv := range kvs {
                keys = append(keys, kv.Key)
                vals = append(vals, kv.Val)
            }
            _, err = node.SetMany(keys, vals)
            require.Equal(t, c.exp, node.raw())
        } else  if ids, ok := c.kvs.([]IndexVal); ok {
            keys := []int{}
            vals := []Value{}
            for _, kv := range ids {
                keys = append(keys, kv.Index)
                vals = append(vals, kv.Val)
            }
            _, err = node.SetManyByIndex(keys, vals)
            require.Equal(t, c.exp, node.raw())
        }
        if err != nil && c.err != err.Error() {
            t.Fatal(err)
        }
    }
}

func TestForEachRaw(t *testing.T) {
    val := _TwitterJson
    node, err := NewSearcher(val).GetValueByPath()
    require.Nil(t, err)
    nodes := []Value{}

    var dfs func(key string, node Value) bool
    var dfs2 func(i int, node Value) bool
    dfs = func(key string, node Value) bool {
        if node.Type() == V_OBJECT  {
            if err := node.ForEachKV(dfs); err != nil {
                panic(err)
            }
        }
        if node.Type() == V_ARRAY  {
            if err := node.ForEachElem(dfs2); err != nil {
                panic(err)
            }
        }
        nodes = append(nodes, node)
        return true
    }
    dfs2 = func(i int, node Value) bool {
        if node.Type() == V_OBJECT  {
            if err := node.ForEachKV(dfs); err != nil {
                panic(err)
            }
        }
        if node.Type() == V_ARRAY  {
            if err := node.ForEachElem(dfs2); err != nil {
                panic(err)
            }
        }
        nodes = append(nodes, node)
        return true
    }
    
    require.NoError(t, node.ForEachKV(dfs))
    require.NotEmpty(t, nodes)
}

func TestRawNode(t *testing.T) {
    _, err := NewSearcher(` { ] `).GetValueByPath()
    require.Error(t, err)
    d1 := ` {"a":1,"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"d":"{\"你好\":\"hello\"}"} `
    root, err := NewSearcher(d1).GetValueByPath()
    require.NoError(t, err)
    v1, err := root.GetByPath("a").Int64()
    require.NoError(t, err)
    require.Equal(t, int64(1), v1)
    v2, err := root.GetByPath("b", 1).Int64()
    require.NoError(t, err)
    require.Equal(t, int64(1), v2)
    v3, err := root.GetByPath("c", "f").Int64()
    require.NoError(t, err)
    require.Equal(t, int64(1), v3)
    v4, err := root.GetByPath("a").Interface()
    require.NoError(t, err)
    require.Equal(t, float64(1), v4)
    v5, err := root.GetByPath("b").Interface()
    require.NoError(t, err)
    require.Equal(t, []interface{}{float64(1), float64(1), float64(1)}, v5)
    v6, err := root.GetByPath("c").Interface()
    require.NoError(t, err)
    require.Equal(t, map[string]interface{}{"d": float64(1), "e": float64(1), "f": float64(1)}, v6)
    v7, err := root.GetByPath("d").String()
    require.NoError(t, err)
    require.Equal(t, `{"你好":"hello"}`, v7)
}

func TestConcurrentGetByPath(t *testing.T) {
    cont, err := NewSearcher(`{"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"a":1}`).GetValueByPath()
    if err != nil {
        t.Fatal(err)
    }
    c := make(chan struct{}, 7)
    wg := sync.WaitGroup{}

    var helper = func(ps ...interface{}){
        wg.Add(1)
        defer wg.Done()
        <-c
        v := cont.GetByPath(ps...)
        require.NoError(t, v.Check())
        vv, _ := v.Int64()
        require.Equal(t, int64(1), vv)
    }
    for i := 0; i < concurrency; i++ {
        go func() {
            helper("b", 1)
        }()
        go func() {
            helper("b", 0)
        }()
        go func() {
            helper("b", 2)
        }()
        go func() {
            helper("c", "d")
        }()
        go func() {
            helper("c", "f")
        }()
        go func() {
            helper("c", "f")
        }()
        go func() {
            helper("a")
        }()
    }

    for i := 0; i < 7*concurrency; i++ {
        c <- struct{}{}
    }
    
    wg.Wait()
}

func TestRawNode_Set(t *testing.T) {
    tests := []struct{
        name string
        js string
        key interface{}
        val Value
        exist bool
        err string
        out string
    }{
        {"exist object",`{"a":1}`,"a",NewValueJSON(`2`),true,"",`{"a":2}`},
        {"not-exist object space",`{"b":1}`,"a",NewValueJSON(`2`),false,"",`{"b":1,"a":2}`},
        {"not-exist object",`{"b":1}`,"a",NewValueJSON(`2`),false,"",`{"b":1,"a":2}`},
        {"empty object",`{}`,"a",NewValueJSON(`2`),false,"",`{"a":2}`},
        {"exist array",`[1]`,0,NewValueJSON(`2`),true,"",`[2]`},
        {"not exist array",`[1]`,1,NewValueJSON(`2`),false,"value not exists",`[1]`},
        {"not exist array over",`[1]`,99,NewValueJSON(`2`),false,"value not exists",`[1]`},
        {"empty array",`[]`,1,NewValueJSON(`2`),false,"value not exists",`[]`},
    }
    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        var exist bool
        var err error
        if key, ok:= c.key.(string); ok{
            exist, err = root.Set(key, c.val)
        } else if id, ok := c.key.(int); ok {
            exist, err = root.SetByIndex(id, c.val)
        }
        if exist != c.exist {
            t.Fatal()
        }
        if err != nil && err.Error() != c.err {
            t.Fatal(err)
        }
        require.Equal(t, c.out, root.raw())
    }
}

func TestRawNode_SetByPath(t *testing.T) {
    tests := []struct{
        name string
        js string
        allowInsert bool
        paths []interface{}
        val Value
        exist bool
        err string
        out string
    }{
        {"exist object",`{"a":1}`,true,[]interface{}{"a"},NewValueJSON(`2`),true,"",`{"a":2}`},
        {"not-exist object",`{"b":1}`,true,[]interface{}{"a"},NewValueJSON(`2`),false,"",`{"b":1,"a":2}`},
        {"empty object",`{}`,true,[]interface{}{"a"},NewValueJSON(`2`),false,"",`{"a":2}`},
        {"empty object 2",`{}`,true,[]interface{}{"a",1},NewValueJSON(`2`),false,"",`{"a":[2]}`},
        {"empty object 3",`{}`,true,[]interface{}{"a",1,"a"},NewValueJSON(`2`),false,"",`{"a":[{"a":2}]}`},
        {"exist array",`[1]`,true,[]interface{}{0},NewValueJSON(`2`),true,"",`[2]`},
        {"not exist array",`[1]`,true,[]interface{}{1},NewValueJSON(`2`),false,"",`[1,2]`},
        {"not exist array",`[1]`,false,[]interface{}{1},NewValueJSON(`2`),false,"value not exists",`[1]`},
        {"empty array",`[]`,true,[]interface{}{1},NewValueJSON(`2`),false,"",`[2]`},
        {"empty array",`[]`,false,[]interface{}{1},NewValueJSON(`2`),false,"value not exists",`[]`},
        {"empty array 2",`[]`,true,[]interface{}{1,1},NewValueJSON(`2`),false,"",`[[2]]`},
        {"empty array 3",`[]`,true,[]interface{}{1,"a",1},NewValueJSON(`2`),false,"",`[{"a":[2]}]`},
        {"empty array 3",`[]`,true,[]interface{}{1,"a","a"},NewValueJSON(`2`),false,"",`[{"a":{"a":2}}]`},
    }
    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        exist, err := root.SetByPath(c.val, c.allowInsert, c.paths...)
        if err != nil && err.Error() != c.err {
            t.Fatal(err)
        }
        require.Equal(t, c.out, root.raw())
        require.Equal(t, c.exist, exist)
    }
}

func TestRawNode_UnsetByPath(t *testing.T) {
    tests := []struct{
        name string
        js string
        paths []interface{}
        exist bool
        err string
        out string
    }{
        {"exist object",`{"a":1}`,[]interface{}{"a"},true,"",`{}`},
        {"exist object",`{"a":1,"b":2,"c":3}`,[]interface{}{"a"},true,"",`{"b":2,"c":3}`},
        {"exist object",`{"a":1,"b":2,"c":3}`,[]interface{}{"b"},true,"",`{"a":1,"c":3}`},
        {"exist object",`{"a":1,"b":2,"c":3}`,[]interface{}{"c"},true,"",`{"a":1,"b":2}`},
        {"not-exist object",`{"b":1}`,[]interface{}{"a"},false,"",`{"b":1}`},
        {"empty object",`{}`,[]interface{}{"a"},false,"",`{}`},
        {"exist object 2",`{"a":{"a":1}}`,[]interface{}{"a","a"},true,"",`{"a":{}}`},
        {"exist object 2",`[{"a":1}]`,[]interface{}{0,"a"},true,"",`[{}]`},
        
        {"not exist array",`[1]`,[]interface{}{1},false,"",`[1]`},
        {"empty array",`[]`,[]interface{}{1},false,"",`[]`},
        {"exist array",`[[1,2,3]]`,[]interface{}{0,0},true,"",`[[2,3]]`},
        {"exist array",`[[1,2,3]]`,[]interface{}{0,1},true,"",`[[1,3]]`},
        {"exist array",`[[1,2,3]]`,[]interface{}{0,2},true,"",`[[1,2]]`},
        {"exist array 2",`[[1]]`,[]interface{}{0,0},true,"",`[[]]`},
        {"exist array 2",`{"a":[1]}`,[]interface{}{"a",0},true,"",`{"a":[]}`},
    }
    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        exist, err := root.UnsetByPath(c.paths...)
        if err != nil && err.Error() != c.err {
            t.Fatal(err)
        }
        if out := root.raw(); err != nil {
            t.Fatal()
        } else {
            require.Equal(t, c.out, out)
        }
        require.Equal(t, c.exist, exist)
    }
}

func TestRawNode_Unset(t *testing.T) {
    tests := []struct{
        name string
        js string
        key interface{}
        exist bool
        err string
        out string
    }{
        {"exist object",`{"a":1}`,"a",true,"",`{}`},
        {"exist object space",`{ "a":1 }`,"a",true, "",`{ }`},
        {"exist object comma",`{ "a":1 , "b":2 }`,"a",true, "",`{  "b":2 }`},
        {"empty object",`{ }`,"a",false, "",`{ }`},
        {"not-exist object",`{"b":1}`,"a",false, "",`{"b":1}`},
        {"exist array",`[1]`,0,true, "",`[]`},
        {"exist array space",`[ 1 ]`,0,true, "",`[ ]`},
        {"exist array comma",`[ 1, 2 ]`,0,true, "",`[  2 ]`},
        {"not exist array",`[1]`,1,false, "",`[1]`},
        {"empty array",`[ ]`,0,false, "",`[ ]`},
    }

    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        var exist bool
        var err error
        if key, ok:= c.key.(string); ok{
            exist, err = root.Unset(key)
        } else if id, ok := c.key.(int); ok {
            exist, err = root.UnsetByIndex(id)
        }

        if err != nil && err.Error() != c.err {
            t.Fatal(err.Error())
        }
        if out := root.raw(); err != nil {
            t.Fatal()
        } else {
            require.Equal(t, c.out, out)
        }
        if exist != c.exist {
            t.Fatal()
        }
    }
}

func TestRawNode_UnsetMany(t *testing.T) {
    tests := []struct{
        name string
        js string
        key interface{}
        exist bool
        err string
        out string
    }{
        {"empty object",`{ }`,[]string{"a","c"},false, "",`{ }`},
        {"1-1 object",`{"a":1}`,[]string{"a"},true, "",`{}`},
        {"1-2 object",`{"a":1}`,[]string{"a","c"},true, "",`{}`},
        {"2-1 object",`{"a":1,"c":3}`,[]string{"a"},true,"",`{"c":3}`},
        {"2-1 object",`{"a":1,"c":3}`,[]string{"c"},true,"",`{"a":1}`},
        {"2-2 object",`{"a":1,"c":3}`,[]string{"a","c"},true,"",`{}`},
        {"3-2 object",`{"a":1,"b":2,"c":3}`,[]string{"a","c"},true,"",`{"b":2}`},
        {"3-2 object",`{"a":1,"b":2,"c":3}`,[]string{"a","b"},true,"",`{"c":3}`},
        {"3-2 object",`{"a":1,"b":2,"c":3}`,[]string{"b","c"},true, "",`{"a":1}`},
        {"3-3 object",`{"a":1,"b":2,"c":3}`,[]string{"a","b","c"},true, "",`{}`},

        {"empty array",`[ ]`,[]int{0, 2},false, "",`[ ]`},
        {"1-1 array",`[1]`,[]int{0},true, "",`[]`},
        {"1-2 array",`[1]`,[]int{0,1},true, "",`[]`},
        {"2-1 array",`[1,2]`,[]int{0},true,"",`[2]`},
        {"2-1 array",`[1,2]`,[]int{1},true,"",`[1]`},
        {"2-2 array",`[1,2]`,[]int{0,1},true,"",`[]`},
        {"3-2 array",`[1,2,3]`,[]int{0,2},true,"",`[2]`},
        {"3-2 array",`[1,2,3]`,[]int{0,1},true,"",`[3]`},
        {"3-2 array",`[1,2,3]`,[]int{1,2},true, "",`[1]`},
        {"3-3 array",`[1,2,3]`,[]int{0,1,2},true, "",`[]`},
    }

    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        var err error
        var exist bool
        if keys, ok := c.key.([]string); ok{
            exist, err = root.UnsetMany(keys)
        } else if ids, ok := c.key.([]int); ok {
            exist, err = root.UnsetManyByIndex(ids)
        }
        require.Equal(t, c.exist, exist)
        if err != nil && err.Error() != c.err {
            t.Fatal(err.Error())
        }
        if out := root.raw(); err != nil {
            t.Fatal()
        } else {
            require.Equal(t, c.out, out)
        }
    }
}

func TestValue_Add(t *testing.T) {
    tests := []struct{
        name string
        js string
        adds []interface{}
        from int
        err string
        out string
    }{
        {"empty +1", `[]`, []interface{}{1}, 0, "", `[1]`},
        {"empty +1", `[]`, []interface{}{1}, 1, "", `[1]`},
        {"empty +1", `[]`, []interface{}{1}, -1, "", `[1]`},
        {"1+1", `[0]`, []interface{}{1}, 0, "", `[1,0]`},
        {"1+1", `[0]`, []interface{}{1}, 1, "", `[0,1]`},
        {"1+1", `[0]`, []interface{}{1}, -1, "", `[0,1]`},
        {"2+1", `[-1,0]`, []interface{}{1}, 0, "", `[1,-1,0]`},
        {"2+1", `[-1,0]`, []interface{}{1}, 1, "", `[-1,1,0]`},
        {"2+1", `[-1,0]`, []interface{}{1}, 2, "", `[-1,0,1]`},
        {"2+1", `[-1,0]`, []interface{}{1}, -1, "", `[-1,0,1]`},
    }
    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        vals := []Value{}
        for _, val := range c.adds {
            vals = append(vals, NewValue(val))
        }
        require.NoError(t, root.AddMany(c.from, vals))
        require.Equal(t, c.out, root.js)
    }
}

func TestValue_Add_Pop(t *testing.T) {
    tests := []struct{
        name string
        js string
        adds []interface{}
        pops int
        err string
        out string
    }{
        {"empty+1-0", `[]`, []interface{}{1}, 0, "", `[1]`},
        {"empty+1-1", `[]`, []interface{}{1}, 1, "", `[]`},
        {"empty+1-2", `[]`, []interface{}{1}, 2, "", `[]`},
        {"empty+2-0", `[]`, []interface{}{1,2}, 0, "", `[1,2]`},
        {"empty+2-1", `[]`, []interface{}{1,2}, 1, "", `[1]`},
        {"1+1-0", `[0]`, []interface{}{1}, 0, "", `[0,1]`},
        {"1+1-1", `[0]`, []interface{}{1}, 1, "", `[0]`},
        {"1+1-2", `[0]`, []interface{}{1}, 2, "", `[]`},
        {"1+1-3", `[0]`, []interface{}{1}, 3, "", `[]`},
        {"1+2-0", `[0]`, []interface{}{1,2}, 0, "", `[0,1,2]`},
        {"1+2-1", `[0]`, []interface{}{1,2}, 1, "", `[0,1]`},
    }
    for _, c := range tests {
        println(c.name)
        root := NewValueJSON(c.js)
        vals := []Value{}
        for i, val := range c.adds {
            vals = append(vals, NewValue(val))
            println("add", i)
            if err := root.AddAny(-1, val); err != nil {
                t.Fatal(err)
            }
            println(root.js)
            if err := root.Pop(); err != nil {
                t.Fatal(err)
            }
            println(root.js)
        }
        require.Equal(t, c.js, root.js)
        require.NoError(t, root.AddMany(-1, vals))
        require.NoError(t, root.PopMany(c.pops))
        require.Equal(t, c.out, root.js)
    }
}

func BenchmarkGetByPath_ReuseNode(b *testing.B) {
    b.Run("Node", func(b *testing.B) {
        root := NewRaw(_TwitterJson)
        _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
    b.Run("Value", func(b *testing.B) {
        cont := NewValueJSON(_TwitterJson)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _, _ = cont.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
}

func BenchmarkNodesGetByPath_NewNode(b *testing.B) {
    b.Run("Node", func(b *testing.B) {
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            root := newRawNode(_TwitterJson, types.V_OBJECT)
            _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
    b.Run("Value", func(b *testing.B) {
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            cont := value(_TwitterJson)
            _, _ = cont.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
        }
    })
}
