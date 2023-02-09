// +build amd64,go1.15,!go1.21

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

package sonic

import (
    `bytes`
    `encoding/hex`
    `encoding/json`
    `fmt`
    `math/rand`
    `reflect`
    `strings`
    `testing`
    `time`

    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/assert`
    `github.com/bytedance/sonic/ast`
)

func Parse(src string) (*ast.Node, error) {
    node, err := ast.NewParser(src).Parse()
    if err != 0 {
        return &node, fmt.Errorf("parsing error: %v", err)
    }
    return &node, nil
}

func assertCond(cond bool) {
    if !cond {
        panic("assertCond failed")
    }
}

func TestExampleSearch(t *testing.T) {
    data := []byte(` { "xx" : [] ,"yy" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ] } `)

    node, e := Get(data, "test", 0)
    x, _ := node.Bool()
    if e != nil || x != true {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = Get(data, "test", 1)
    a, _ := node.Float64()
    if e != nil || a != 0.1 {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = Get(data, "test", 2)
    b, _ := node.String()
    if e != nil || b != "abc" {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = Get(data, "test", 3)
    arr, _ := node.Array()
    if e != nil || arr[0] != "h" {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = Get(data, "test", 4, "a")
    c, _ := node.String()
    if e != nil || c != "bc" {
        t.Fatalf("node: %v, err: %v", node, e)
    }
}

func TestExampleSearchEscapedKey(t *testing.T) {
    data := []byte(`
{
    "xx" : [] ,
    "yy" :{ }, 
    "test\"" : [
        true ,
        0.1 ,
        "abc",
        [
            "h"
        ], 
        {
            "a\u0008": {},
            "b\\\\": null,
            "\u2028\u2028": "\u2028\u2029",
            "\u0026":2,
            "0":1
        } 
    ],
    ",,,,,,,,,,(,,15": ",",
    ",,,,,,,,,,(,,,16": "a,]}",
    ",,,,,,,,,,(,,,,17": 1,
    ",,,,,,,,,,(,,,,,,,,,,,,,,,,(,,,,34": "c"
} `)

    type getTest struct{
        path []interface{}
        expect interface{}
    }

    tests := []getTest {
        {[]interface{}{"test\"", 0}, true},
        {[]interface{}{"test\"", 1}, 0.1},
        {[]interface{}{"test\"", 2}, "abc"},
        {[]interface{}{"test\"", 3}, []interface{}{"h"}},
        {[]interface{}{"test\"", 4, "a\u0008"}, map[string]interface{}{}},
        {[]interface{}{"test\"", 4, "b\\\\"}, nil},
        {[]interface{}{"test\"", 4, "\u2028\u2028"}, "\u2028\u2029"},
        {[]interface{}{"test\"", 4, "0"}, float64(1)},
        {[]interface{}{",,,,,,,,,,(,,15"}, ","},
        {[]interface{}{",,,,,,,,,,(,,,16"}, "a,]}"},
        {[]interface{}{",,,,,,,,,,(,,,,17"}, float64(1)},
        {[]interface{}{",,,,,,,,,,(,,,,,,,,,,,,,,,,(,,,,34"}, "c"},
    }

    for _, test := range(tests) {
        node, err := Get(data, test.path...)
        assert.NoErrorf(t, err, "get return errors")
        got, err := node.Interface()
        assert.NoErrorf(t, err, "get convert errors")
        assert.Equalf(t, test.expect, got, "get result is wrong from path %#v", test.path)
    }
}

func TestExampleSearchErr(t *testing.T) {
    data := []byte(` { "xx" : [] ,"yy" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ] } `)
    node, e := Get(data, "zz")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "xx", 4)
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "yy", "a")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "test", 4, "x")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)
}

func TestExampleSearchEscapedKeyError(t *testing.T) {
    data := []byte(` { "xx" : [] ,"yy" :{ }, "x\u0008" : [] ,"y\\\"y" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ] } `)
    node, e := Get(data, "zz")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "x\u0008", 4)
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "yy", "a")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)

    node, e = Get(data, "test", 4, "x")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = Get(data, "y\\\"y", 4, "x")
    if e == nil {
        t.Fatalf("node: %v, err: %v", node, e)
    }
    fmt.Println(e)
}

func TestRandomData(t *testing.T) {
    var lstr string
    defer func() {
        if v := recover(); v != nil {
            println("'" + hex.Dump([]byte(lstr)) + "'")
            println(lstr)
            panic(v)
        }
    }()
    data := []byte(`"�-mp�`)
    _, err := ast.NewParser(string(data)).Parse()
    if err != 0 {
        fmt.Println(hex.Dump(data))
        fmt.Println(string(data))
    }
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 200)
    for i := 0; i < 1000000; i++ {
        n, ee := rand.Read(b[:rand.Int()%len(b)])
        if ee != nil {
            t.Fatalf("get random bytes failed: %v,", ee)
            return
        }
        lstr = string(b[:n])
        _, _ = ast.NewParser(lstr).Parse()
    }
}

func TestRandomValidStrings(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 200)
    for i := 0; i < 1000; i++ {
        n, err := rand.Read(b[:rand.Int()%len(b)])
        if err != nil {
            t.Fatal("get random data failed:", err)
        }
        sm, err := json.Marshal(string(b[:n]))
        if err != nil {
            t.Fatal("marshal data failed:",err)
        }
        var su string
        if err := json.Unmarshal(sm, &su); err != nil {
            t.Fatal("unmarshal data failed:",err)
        }
        token, err := GetFromString(`{"str":`+string(sm)+`}`, "str")
        if err != nil {
            spew.Dump(string(sm))
            t.Fatal("search data failed:",err)
        }
        x, _ := token.Interface()
        st, ok := x.(string)
        if !ok {
            t.Fatalf("type mismatch, exp: %v, got: %v", su, x)
        }
        if st != su {
            t.Fatalf("string mismatch, exp: %v, got: %v", su, x)
        }
    }
}


func TestEmoji(t *testing.T) {
    var input = []byte(`{"utf8":"Example emoji, KO: \ud83d\udd13, \ud83c\udfc3 ` +
        `OK: \u2764\ufe0f "}`)
    value, err := Get(input, "utf8")
    if err != nil {
        t.Fatal(err)
    }
    var v map[string]interface{}
    if err := json.Unmarshal(input, &v); err != nil {
        t.Fatal(err)
    }
    s, _ := v["utf8"].(string)
    x, _ := value.String()
    if x != s {
        t.Fatalf("expected '%v', got '%v'", s, x)
    }
}

func testEscapePath(t *testing.T, json, expect string, path ...interface{}) {
    n, e := Get([]byte(json), path...)
    if e != nil {
        t.Fatal(e)
    }
    x, _ := n.String()
    if x != expect {
        x, _ := n.Interface()
        t.Fatalf("expected '%v', got '%v'", expect, x)
    }
}

func TestEscapePath(t *testing.T) {
    data := `{
        "test":{
            "*":"valZ",
            "*v":"val0",
            "keyv*":"val1",
            "key*v":"val2",
            "keyv?":"val3",
            "key?v":"val4",
            "keyv.":"val5",
            "key.v":"val6",
            "keyk*":{"key?":"val7"}
        }
    }`

    testEscapePath(t, data, "valZ", "test", "*")
    testEscapePath(t, data, "val0", "test", "*v")
    testEscapePath(t, data, "val1", "test", "keyv*")
    testEscapePath(t, data, "val2", "test", "key*v")
    testEscapePath(t, data, "val3", "test", "keyv?")
    testEscapePath(t, data, "val4", "test", "key?v")
    testEscapePath(t, data, "val5", "test", "keyv.")
    testEscapePath(t, data, "val6", "test", "key.v")
    testEscapePath(t, data, "val7", "test", "keyk*", "key?")
}

func TestParseAny(t *testing.T) {
    n, e := Parse("100")
    assertCond(e == nil)
    if n == nil {
        panic("n is nil")
    }
    x, _ := n.Float64()
    assertCond(x == 100)
    n, e = Parse("true")
    assertCond(e == nil)
    if n == nil {
        panic("n is nil")
    }

    a, _ := n.Bool()
    assertCond(a)
    n, e = Parse("false")
    assertCond(e == nil)
    if n == nil {
        panic("n is nil")
    }
    b, _ := n.Bool()
    assertCond(b == false)
    n, e = Parse("yikes")
    assertCond(e != nil)
}

func TestTime(t *testing.T) {
    data := []byte(`{
      "code": 0,
      "msg": "",
      "data": {
        "sz002024": {
          "qfqday": [
            [
              "2014-01-02",
              "8.93",
              "9.03",
              "9.17",
              "8.88",
              "621143.00"
            ],
            [
              "2014-01-03",
              "9.03",
              "9.30",
              "9.47",
              "8.98",
              "1624438.00"
            ]
          ]
        }
      }
    }`)

    var num []string
    n, e := Get(data, "data", "sz002024", "qfqday", 0)
    if e != nil {
        t.Fatal(e)
    }

    arr, _ := n.Array()
    for _, v := range arr {
        s := v.(string)
        num = append(num, s)
    }
    if fmt.Sprintf("%v", num) != "[2014-01-02 8.93 9.03 9.17 8.88 621143.00]" {
        t.Fatalf("invalid result")
    }
}

var exampleJSON = `{
    "widget": {
        "debug": "on",
        "window": {
            "title": "Sample Konfabulator Widget",
            "name": "main_window",
            "width": 500,
            "height": 500
        },
        "image": {
            "src": "Images/Sun.png",
            "hOffset": 250,
            "vOffset": 250,
            "alignment": "center"
        },
        "text": {
            "data": "Click Here",
            "size": 36,
            "style": "bold",
            "vOffset": 100,
            "alignment": "center",
            "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
        }
    }
}`

func TestUnmarshalMap(t *testing.T) {
    n, err := Parse(exampleJSON)
    if err != nil || n == nil {
        t.Fatal(err)
    }
    m1, _ := n.Map()
    var m2 map[string]interface{}
    if err := json.Unmarshal([]byte(exampleJSON), &m2); err != nil {
        t.Fatal(err)
    }
    b1, err := json.Marshal(m1)
    if err != nil {
        t.Fatal(err)
    }
    b2, err := json.Marshal(m2)
    if err != nil {
        t.Fatal(err)
    }
    if !bytes.Equal(b1, b2) {
        t.Fatalf("b1 != b2\n b1: %v\nb2: %v", string(b1), string(b2))
    }
}

func GetMany(src2 string, path ...string) (ret []string) {
    src := []byte(src2)
    for _, p := range path {
        pathes := strings.Split(p, ".")
        if len(pathes) == 0 {
            panic(fmt.Sprintf("invalid path: %v", p))
        }
        ps := make([]interface{}, 0, len(pathes))
        for _, p := range pathes {
            ps = append(ps, p)
        }
        n, e := Get(src, ps...)
        if e != nil {
            ret = append(ret, "")
            continue
        }
        x, _ := n.Interface()
        ret = append(ret, fmt.Sprintf("%v", x))
    }
    return
}

func get(str string, path string) *ast.Node {
    src := []byte(str)
    pathes := strings.Split(path, ".")
    if len(pathes) == 0 {
        panic(fmt.Sprintf("invalid path: %v", path))
    }
    ps := make([]interface{}, 0, len(pathes))
    for _, p := range pathes {
        ps = append(ps, p)
    }
    n, e := Get(src, ps...)
    if e != nil {
        return nil
    }
    return &n
}

func TestSingleArrayValue(t *testing.T) {
    var data = []byte(`{"key": "value","key2":[1,2,3,4,"A"]}`)
    array, _ := get(string(data), "key2").Array()
    if len(array) != 5 {
        t.Fatalf("got '%v', expected '%v'", len(array), 5)
    }

    _, e := Get(data, "key3")
    if e == nil {
        t.Fatalf("got '%v', expected '%v'", e, nil)
    }
}

var manyJSON = `  {
    "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
        "a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"hello":"world"
        }}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}},
        "position":{"type":"Point","coordinates":[-115.24,33.09]},
        "loves":["world peace"],
        "name":{"last":"Anderson","first":"Nancy"},
        "age":31,
        "x":{"a":"emptya","b":"emptyb"},
        "name_last":"Yellow",
        "name_first":"Cat",
}`

func TestManyBasic(t *testing.T) {
    testMany := func(shouldFallback bool, expect string, paths ...string) {
        println()
        rs := GetMany(
            manyJSON,
            paths...,
        )
        // if len(rs) != len(paths) {
        //     t.Fatalf("expected %v, got %v", len(paths), len(rs))
        // }
        var results = "[" + strings.Join(rs, " ") + "]"
        if results != expect {
            fmt.Printf("%v\n", paths)
            t.Fatalf("expected %v, got %v", expect, results)
        }
    }
    testMany(false, "[Point]", "position.type")
    testMany(false, `[emptya [world peace] 31]`, "x.a", "loves", "age")
    testMany(false, `[[world peace]]`, "loves")
    testMany(false, `[map[first:Nancy last:Anderson] Nancy]`, "name",
        "name.first")
    testMany(true, `[]`, strings.Repeat("a.", 40)+"hello")
    res := get(manyJSON, strings.Repeat("a.", 48)+"a")
    assertCond(res != nil)
    x, _ := res.Interface()
    testMany(true, "["+fmt.Sprint(x)+"]", strings.Repeat("a.", 48)+"a")
    // these should fallback
    testMany(true, `[Cat Nancy]`, "name_first", "name.first")
    testMany(true, `[world]`, strings.Repeat("a.", 70)+"hello")
}

func testMany(t *testing.T, json string, paths, expected []string) {
    testManyAny(t, json, paths, expected)
    testManyAny(t, json, paths, expected)
}

func testManyAny(t *testing.T, json string, paths, expected []string) {
    var result []string
    for i := 0; i < 2; i++ {
        var which string
        if i == 0 {
            which = "Get"
            result = nil
            for j := 0; j < len(expected); j++ {
                x, _ := get(json, paths[j]).Interface()
                result = append(result, fmt.Sprintf("%v", x))
            }
        } else if i == 1 {
            which = "GetMany"
            result = GetMany(json, paths...)
        }
        if result == nil {
            panic("result is nil")
        }
        for j := 0; j < len(expected); j++ {
            if result[j] != expected[j] {
                t.Fatalf("Using key '%s' for '%s'\nexpected '%v', got '%v'",
                    paths[j], which, expected[j], result[j])
            }
        }
    }
}

func TestNested(t *testing.T) {
    data := `{ "name": "FirstName", "name1": "FirstName1", ` +
        `"address": "address1", "addressDetails": "address2", }`
    paths := []string{"name", "name1", "address", "addressDetails"}
    expected := []string{"FirstName", "FirstName1", "address1", "address2"}
    t.Run("SingleMany", func(t *testing.T) {
        testMany(t, data, paths,
            expected)
    })
}

func TestMultiLevelFields(t *testing.T) {
    data := `{ "Level1Field1":3, 
               "Level1Field4":4, 
               "Level1Field2":{ "Level2Field1":[ "value1", "value2" ], 
               "Level2Field2":{ "Level3Field1":[ { "key1":"value1" } ] } } }`
    paths := []string{"Level1Field1", "Level1Field2.Level2Field1",
        "Level1Field2.Level2Field2.Level3Field1", "Level1Field4"}
    expected := []string{"3", `[value1 value2]`,
        `[map[key1:value1]]`, "4"}
    t.Run("SingleMany", func(t *testing.T) {
        testMany(t, data, paths,
            expected)
    })
}

func TestRandomMany(t *testing.T) {
    var lstr string
    defer func() {
        if v := recover(); v != nil {
            println("'" + hex.EncodeToString([]byte(lstr)) + "'")
            println("'" + lstr + "'")
            panic(v)
        }
    }()
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 512)
    for i := 0; i < 5000; i++ {
        n, err := rand.Read(b[:rand.Int()%len(b)])
        if err != nil {
            t.Fatal(err)
        }
        lstr = string(b[:n])
        paths := make([]string, rand.Int()%64)
        for i := range paths {
            var b []byte
            n := rand.Int() % 5
            for j := 0; j < n; j++ {
                if j > 0 {
                    b = append(b, '.')
                }
                nn := rand.Int() % 10
                for k := 0; k < nn; k++ {
                    b = append(b, 'a'+byte(rand.Int()%26))
                }
            }
            paths[i] = string(b)
        }
        GetMany(lstr, paths...)
    }
}

func TestGetMany(t *testing.T) {
    data := `{"bar": {"id": 99, "mybar": "my mybar" }, "foo": ` +
        `{"myfoo": [605]}}`
    paths := []string{"foo.myfoo", "bar.id", "bar.mybar", "bar.mybarx"}
    expected := []string{"[605]", "99", "my mybar", ""}
    results := GetMany(data, paths...)
    if len(expected) != len(results) {
        t.Fatalf("expected %v, got %v", len(expected), len(results))
    }
    for i, path := range paths {
        if results[i] != expected[i] {
            t.Fatalf("expected '%v', got '%v' for path '%v'", expected[i],
                results[i], path)
        }
    }
}

func TestGetMany2(t *testing.T) {
    data := `{"bar": {"id": 99, "xyz": "my xyz"}, "foo": {"myfoo": [605]}}`
    paths := []string{"foo.myfoo", "bar.id", "bar.xyz", "bar.abc"}
    expected := []string{"[605]", "99", "my xyz", ""}
    results := GetMany(data, paths...)
    if len(expected) != len(results) {
        t.Fatalf("expected %v, got %v", len(expected), len(results))
    }
    for i, path := range paths {
        if results[i] != expected[i] {
            t.Fatalf("expected '%v', got '%v' for path '%v'", expected[i],
                results[i], path)
        }
    }
}

func TestNullArray(t *testing.T) {
    n, _ := get(`{"data":null}`, "data").Interface()
    if n != nil {
        t.Fatalf("expected '%v', got '%v'", nil, n)
    }
    n = get(`{}`, "data")
    if reflect.DeepEqual(n, nil) {
        t.Fatalf("expected '%v', got '%v'", nil, n)
    }
    n = get(`{"data":[]}`, "data")
    if reflect.DeepEqual(n, &ast.Node{}) {
        t.Fatalf("expected '%v', got '%v'", nil, n)
    }
    arr, _ := get(`{"data":[null]}`, "data").Array()
    n = len(arr)
    if n != 1 {
        t.Fatalf("expected '%v', got '%v'", 1, n)
    }
}

func TestGetMany3(t *testing.T) {
    var r []string
    data := `{"MarketName":null,"Nounce":6115}`
    r = GetMany(data, "Nounce", "Buys", "Sells", "Fills")
    if strings.Replace(fmt.Sprintf("%v", r), " ", "", -1) != "[6115]" {
        t.Fatalf("expected '%v', got '%v'", "[6115]",
            strings.Replace(fmt.Sprintf("%v", r), " ", "", -1))
    }
    r = GetMany(data, "Nounce", "Buys", "Sells")
    if strings.Replace(fmt.Sprintf("%v", r), " ", "", -1) != "[6115]" {
        t.Fatalf("expected '%v', got '%v'", "[6115]",
            strings.Replace(fmt.Sprintf("%v", r), " ", "", -1))
    }
    r = GetMany(data, "Nounce")
    if strings.Replace(fmt.Sprintf("%v", r), " ", "", -1) != "[6115]" {
        t.Fatalf("expected '%v', got '%v'", "[6115]",
            strings.Replace(fmt.Sprintf("%v", r), " ", "", -1))
    }
}

func TestGetMany4(t *testing.T) {
    data := `{"one": {"two": 2, "three": 3}, "four": 4, "five": 5}`
    results := GetMany(data, "four", "five", "one.two", "one.six")
    expected := []string{"4", "5", "2", ""}
    for i, r := range results {
        if r != expected[i] {
            t.Fatalf("expected %v, got %v", expected[i], r)
        }
    }
}

func TestGetNotExist(t *testing.T) {
    var dataStr = `{"m1":{"m2":3}}`
    ret, err := GetFromString(dataStr, "not_exist", "m3")
    if err == nil || ret.Exists() {
        t.Fatal("Get exist!")
    }
    if ret.Type() != ast.V_NONE {
        t.Fatal(ret.Type())
    }
    ret, err = GetFromString(dataStr)
    if !ret.IsRaw() || ret.Type() != ast.V_OBJECT {
        t.Fatal(ret.Type())
    }
    v11 := ret.Get("not_exist")
    if v11.Exists() {
        t.Fatal()
    }
    v2 := ret.GetByPath("m1", "m2")
    if !v2.Exists() || !v2.IsRaw() {
        t.Fatal(ret.Type())
    }
    x, _ := v2.Int64()
    if x != 3 {
        t.Fatal(x)
    }
}