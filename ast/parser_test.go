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
    `testing`

    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
    `github.com/tidwall/gjson`
)

func runDecoderTest(t *testing.T, src string, expect interface{}) {
    vv, err := NewParser(src).Parse()
    if err != 0 { panic(err) }
    //fmt.Printf("%s -> %s :: %v\n", src, reflect.TypeOf(vv), vv)
    assert.Equal(t, expect, vv.Interface())
}

func runDecoderTestUseNumber(t *testing.T, src string, expect interface{}) {
    vv, err := NewParser(src).Parse()
    if err != 0 { panic(err) }
    //fmt.Printf("%s -> %s :: %v\n", src, reflect.TypeOf(vv), vv)
    vvv := vv.InterfaceUseNumber()
    switch vvv.(type) {
    case json.Number:
        assert.Equal(t, expect, n2f64(vvv.(json.Number)))
    case []interface{}:
        x := vvv.([]interface{})
        for i, e := range x {
            if ev,ok := e.(json.Number);ok {
                x[i] = n2f64(ev)
            }
        }
        assert.Equal(t, expect, x)
    case map[string]interface{}:
        x := vvv.(map[string]interface{})
        for k,v := range x {
            if ev, ok := v.(json.Number); ok {
                x[k] = n2f64(ev)
            }
        }
        assert.Equal(t, expect, x)
    }
}

func n2f64(i json.Number) float64{ 
    x, err := i.Float64()
    if err != nil {
        panic(err)
    }
    return x
}

func runDecoderTestUseInt64(t *testing.T, src string, expect interface{}) {
    vv, err := NewParser(src).Parse()
    if err != 0 { panic(err) }
    //fmt.Printf("%s -> %s :: %v\n", src, reflect.TypeOf(vv), vv)
    assert.Equal(t, expect, vv.InterfaceUseInt64())
}

func TestParser_Basic(t *testing.T) {
    runDecoderTest(t, `null`, nil)
    runDecoderTest(t, `true`, true)
    runDecoderTest(t, `false`, false)
    runDecoderTest(t, `"hello, world \\ \/ \b \f \n \r \t \u666f 测试中文"`, "hello, world \\ / \b \f \n \r \t \u666f 测试中文")
    runDecoderTest(t, `"\ud83d\ude00"`, "😀")
    runDecoderTest(t, `0`, float64(0))
    runDecoderTest(t, `-0`, float64(0))
    runDecoderTest(t, `123456`, float64(123456))
    runDecoderTest(t, `-12345`, float64(-12345))
    runDecoderTest(t, `0.2`, 0.2)
    runDecoderTest(t, `1.2`, 1.2)
    runDecoderTest(t, `-0.2`, -0.2)
    runDecoderTest(t, `-1.2`, -1.2)
    runDecoderTest(t, `0e12`, 0e12)
    runDecoderTest(t, `0e+12`, 0e+12)
    runDecoderTest(t, `0e-12`, 0e-12)
    runDecoderTest(t, `-0e12`, -0e12)
    runDecoderTest(t, `-0e+12`, -0e+12)
    runDecoderTest(t, `-0e-12`, -0e-12)
    runDecoderTest(t, `2e12`, 2e12)
    runDecoderTest(t, `2E12`, 2e12)
    runDecoderTest(t, `2e+12`, 2e+12)
    runDecoderTest(t, `2e-12`, 2e-12)
    runDecoderTest(t, `-2e12`, -2e12)
    runDecoderTest(t, `-2e+12`, -2e+12)
    runDecoderTest(t, `-2e-12`, -2e-12)
    runDecoderTest(t, `0.2e12`, 0.2e12)
    runDecoderTest(t, `0.2e+12`, 0.2e+12)
    runDecoderTest(t, `0.2e-12`, 0.2e-12)
    runDecoderTest(t, `-0.2e12`, -0.2e12)
    runDecoderTest(t, `-0.2e+12`, -0.2e+12)
    runDecoderTest(t, `-0.2e-12`, -0.2e-12)
    runDecoderTest(t, `1.2e12`, 1.2e12)
    runDecoderTest(t, `1.2e+12`, 1.2e+12)
    runDecoderTest(t, `1.2e-12`, 1.2e-12)
    runDecoderTest(t, `-1.2e12`, -1.2e12)
    runDecoderTest(t, `-1.2e+12`, -1.2e+12)
    runDecoderTest(t, `-1.2e-12`, -1.2e-12)
    runDecoderTest(t, `-1.2E-12`, -1.2e-12)
    runDecoderTest(t, `[]`, []interface{}{})
    runDecoderTest(t, `{}`, map[string]interface{}{})
    runDecoderTest(t, `["asd", "123", true, false, null, 2.4, 1.2e15]`, []interface{}{"asd", "123", true, false, nil, 2.4, 1.2e15})
    runDecoderTest(t, `{"asdf": "qwer", "zxcv": true}`, map[string]interface{}{"asdf": "qwer", "zxcv": true})
    runDecoderTest(t, `{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`, map[string]interface{}{"a":"123", "b":true, "c":false, "d":nil, "e":float64(2.4), "f":float64(1.2e15), "g":float64(1)})

    runDecoderTestUseNumber(t, `null`, nil)
    runDecoderTestUseNumber(t, `true`, true)
    runDecoderTestUseNumber(t, `false`, false)
    runDecoderTestUseNumber(t, `"hello, world \\ \/ \b \f \n \r \t \u666f 测试中文"`, "hello, world \\ / \b \f \n \r \t \u666f 测试中文")
    runDecoderTestUseNumber(t, `"\ud83d\ude00"`, "😀")
    runDecoderTestUseNumber(t, `0`, float64(0))
    runDecoderTestUseNumber(t, `-0`, float64(0))
    runDecoderTestUseNumber(t, `123456`, float64(123456))
    runDecoderTestUseNumber(t, `-12345`, float64(-12345))
    runDecoderTestUseNumber(t, `0.2`, float64(0.2))
    runDecoderTestUseNumber(t, `1.2`, float64(1.2))
    runDecoderTestUseNumber(t, `-0.2`, float64(-0.2))
    runDecoderTestUseNumber(t, `-1.2`, float64(-1.2))
    runDecoderTestUseNumber(t, `0e12`, float64(0e12))
    runDecoderTestUseNumber(t, `0e+12`, float64(0e+12))
    runDecoderTestUseNumber(t, `0e-12`, float64(0e-12))
    runDecoderTestUseNumber(t, `-0e12`, float64(-0e12))
    runDecoderTestUseNumber(t, `-0e+12`, float64(-0e+12))
    runDecoderTestUseNumber(t, `-0e-12`, float64(-0e-12))
    runDecoderTestUseNumber(t, `2e12`, float64(2e12))
    runDecoderTestUseNumber(t, `2E12`, float64(2e12))
    runDecoderTestUseNumber(t, `2e+12`, float64(2e+12))
    runDecoderTestUseNumber(t, `2e-12`, float64(2e-12))
    runDecoderTestUseNumber(t, `-2e12`, float64(-2e12))
    runDecoderTestUseNumber(t, `-2e+12`, float64(-2e+12))
    runDecoderTestUseNumber(t, `-2e-12`, float64(-2e-12))
    runDecoderTestUseNumber(t, `0.2e12`, float64(0.2e12))
    runDecoderTestUseNumber(t, `0.2e+12`, float64(0.2e+12))
    runDecoderTestUseNumber(t, `0.2e-12`, float64(0.2e-12))
    runDecoderTestUseNumber(t, `-0.2e12`, float64(-0.2e12))
    runDecoderTestUseNumber(t, `-0.2e+12`, float64(-0.2e+12))
    runDecoderTestUseNumber(t, `-0.2e-12`, float64(-0.2e-12))
    runDecoderTestUseNumber(t, `1.2e12`, float64(1.2e12))
    runDecoderTestUseNumber(t, `1.2e+12`, float64(1.2e+12))
    runDecoderTestUseNumber(t, `1.2e-12`, float64(1.2e-12))
    runDecoderTestUseNumber(t, `-1.2e12`, float64(-1.2e12))
    runDecoderTestUseNumber(t, `-1.2e+12`, float64(-1.2e+12))
    runDecoderTestUseNumber(t, `-1.2e-12`, float64(-1.2e-12))
    runDecoderTestUseNumber(t, `-1.2E-12`, float64(-1.2e-12))
    runDecoderTestUseNumber(t, `["asd", "123", true, false, null, 2.4, 1.2e15, 1]`, []interface{}{"asd", "123", true, false, nil, float64(2.4), float64(1.2e15), float64(1)})
    runDecoderTestUseNumber(t, `{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`, map[string]interface{}{"a":"123", "b":true, "c":false, "d":nil, "e":float64(2.4), "f":float64(1.2e15), "g":float64(1)})

    runDecoderTestUseInt64(t, `null`, nil)
    runDecoderTestUseInt64(t, `true`, true)
    runDecoderTestUseInt64(t, `false`, false)
    runDecoderTestUseInt64(t, `"hello, world \\ \/ \b \f \n \r \t \u666f 测试中文"`, "hello, world \\ / \b \f \n \r \t \u666f 测试中文")
    runDecoderTestUseInt64(t, `"\ud83d\ude00"`, "😀")
    runDecoderTestUseInt64(t, `0`, int64(0))
    runDecoderTestUseInt64(t, `-0`, int64(0))
    runDecoderTestUseInt64(t, `123456`, int64(123456))
    runDecoderTestUseInt64(t, `-12345`, int64(-12345))
    runDecoderTestUseInt64(t, `0.2`, float64(0.2))
    runDecoderTestUseInt64(t, `1.2`, float64(1.2))
    runDecoderTestUseInt64(t, `-0.2`, float64(-0.2))
    runDecoderTestUseInt64(t, `-1.2`, float64(-1.2))
    runDecoderTestUseInt64(t, `0e12`, float64(0e12))
    runDecoderTestUseInt64(t, `0e+12`, float64(0e+12))
    runDecoderTestUseInt64(t, `0e-12`, float64(0e-12))
    runDecoderTestUseInt64(t, `-0e12`, float64(-0e12))
    runDecoderTestUseInt64(t, `-0e+12`, float64(-0e+12))
    runDecoderTestUseInt64(t, `-0e-12`, float64(-0e-12))
    runDecoderTestUseInt64(t, `2e12`, float64(2e12))
    runDecoderTestUseInt64(t, `2E12`, float64(2e12))
    runDecoderTestUseInt64(t, `2e+12`, float64(2e+12))
    runDecoderTestUseInt64(t, `2e-12`, float64(2e-12))
    runDecoderTestUseInt64(t, `-2e12`, float64(-2e12))
    runDecoderTestUseInt64(t, `-2e+12`, float64(-2e+12))
    runDecoderTestUseInt64(t, `-2e-12`, float64(-2e-12))
    runDecoderTestUseInt64(t, `0.2e12`, float64(0.2e12))
    runDecoderTestUseInt64(t, `0.2e+12`, float64(0.2e+12))
    runDecoderTestUseInt64(t, `0.2e-12`, float64(0.2e-12))
    runDecoderTestUseInt64(t, `-0.2e12`, float64(-0.2e12))
    runDecoderTestUseInt64(t, `-0.2e+12`, float64(-0.2e+12))
    runDecoderTestUseInt64(t, `-0.2e-12`, float64(-0.2e-12))
    runDecoderTestUseInt64(t, `1.2e12`, float64(1.2e12))
    runDecoderTestUseInt64(t, `1.2e+12`, float64(1.2e+12))
    runDecoderTestUseInt64(t, `1.2e-12`, float64(1.2e-12))
    runDecoderTestUseInt64(t, `-1.2e12`, float64(-1.2e12))
    runDecoderTestUseInt64(t, `-1.2e+12`, float64(-1.2e+12))
    runDecoderTestUseInt64(t, `-1.2e-12`, float64(-1.2e-12))
    runDecoderTestUseInt64(t, `-1.2E-12`, float64(-1.2e-12))
    runDecoderTestUseInt64(t, `["asd", "123", true, false, null, 2.4, 1.2e15, 1]`, []interface{}{"asd", "123", true, false, nil, float64(2.4), float64(1.2e15), int64(1)})
    runDecoderTestUseInt64(t, `{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`, map[string]interface{}{"a":"123", "b":true, "c":false, "d":nil, "e":float64(2.4), "f":float64(1.2e15), "g":int64(1)})
}

func BenchmarkParser_StdLib(b *testing.B) {
    var bv = []byte(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var out interface{}
        _ = json.Unmarshal(bv, &out)
    }
}

func BenchmarkParser_JsonIter(b *testing.B) {
    var bv = []byte(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var out interface{}
        _ = jsoniter.Unmarshal(bv, &out)
    }
}

func BenchmarkParser_Sonic(b *testing.B) {
    _, _, _ = Loads(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _, _ = Loads(_TwitterJson)
    }
}

func BenchmarkParser_Parallel_StdLib(b *testing.B) {
    var bv = []byte(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var out interface{}
            _ = json.Unmarshal(bv, &out)
        }
    })
}

func BenchmarkParser_Parallel_JsonIter(b *testing.B) {
    var bv = []byte(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var out interface{}
            _ = jsoniter.Unmarshal(bv, &out)
        }
    })
}

func BenchmarkParser_Parallel_Sonic(b *testing.B) {
    _, _, _ = Loads(_TwitterJson)
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _, _ = Loads(_TwitterJson)
        }
    })
}

func BenchmarkGetOne_Gjson(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ast := gjson.Get(_TwitterJson, "statuses.2.id")
        node := ast.Int()
        if node != 249289491129438208 {
            b.Fatal(node)
        }
    }
}

func BenchmarkGetOne_Jsoniter(b *testing.B) {
    data := []byte(_TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 2, "id")
        node := ast.ToInt()
        if node != 249289491129438208 {
            b.Fail()
        }
    }
}

func BenchmarkGetOne_Sonic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ast, _ := NewParser(_TwitterJson).Parse()
        node := ast.Get("statuses").Index(2).Get("id").Int64()
        if node != 249289491129438208 {
            b.Fail()
        }
    }
}

func BenchmarkGetSeven_Gjson(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ast := gjson.Get(_TwitterJson, "statuses.3.id")
        ast = gjson.Get(_TwitterJson, "statuses.3.user.entities.description")
        ast = gjson.Get(_TwitterJson, "statuses.3.user.entities.url.urls")
        ast = gjson.Get(_TwitterJson, "statuses.3.user.entities.url")
        ast = gjson.Get(_TwitterJson, "statuses.3.user.created_at")
        ast = gjson.Get(_TwitterJson, "statuses.3.user.name")
        ast = gjson.Get(_TwitterJson, "statuses.3.text")
        if ast.Value() == nil {
            b.Fail()
        }
    }
}

func BenchmarkGetSeven_Jsoniter(b *testing.B) {
    data := []byte(_TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 3, "id")
        ast = jsoniter.Get(data, "statuses",  3, "user", "entities","description")
        ast = jsoniter.Get(data, "statuses", 3, "user", "entities","url","urls")
        ast = jsoniter.Get(data, "statuses",  3, "user", "entities","url")
        ast = jsoniter.Get(data, "statuses",  3, "user", "created_at")
        ast = jsoniter.Get(data, "statuses",  3, "user", "name")
        ast = jsoniter.Get(data, "statuses",  3, "text")
        if ast == nil {
            b.Fail()
        }
    }
}

func BenchmarkGetSeven_SonicParser(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ast, _ := NewParser(_TwitterJson).Parse()
        node := ast.GetByPath( "statuses", 3, "id")
        node = ast.GetByPath("statuses",  3, "user", "entities","description")
        node = ast.GetByPath("statuses", 3, "user", "entities","url","urls")
        node = ast.GetByPath("statuses",  3, "user", "entities","url")
        node = ast.GetByPath("statuses",  3, "user", "created_at")
        node = ast.GetByPath("statuses",  3, "user", "name")
        node = ast.GetByPath("statuses",  3, "text")
        if node == nil {
            b.Fail()
        }
    }
}
