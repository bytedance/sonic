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
    `os`
    `runtime`
    `runtime/debug`
    `sync`
    `testing`
    `time`

    `github.com/stretchr/testify/assert`
)

var (
    debugSyncGC  = os.Getenv("SONIC_SYNC_GC")  != ""
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

func TestMain(m *testing.M) {
    go func ()  {
        if !debugAsyncGC {
            return
        }
        println("Begin GC looping...")
        for {
            runtime.GC()
            debug.FreeOSMemory()
        }
        println("stop GC looping!")
    }()
    time.Sleep(time.Millisecond)
    m.Run()
}

func TestGC_Parse(t *testing.T) {
    if debugSyncGC {
        return
    }
    _, _, err := Loads(_TwitterJson)
    if err != nil {
        t.Fatal(err)
    }
    wg := &sync.WaitGroup{}
    N := 1000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
            defer wg.Done()
            _, _, err := Loads(_TwitterJson)
            if err != nil {
                t.Fatal(err)
            }
            runtime.GC()
        }(wg)
    }
    wg.Wait()
}

func runDecoderTest(t *testing.T, src string, expect interface{}) {
    vv, err := NewParser(src).Parse()
    if err != 0 { panic(err) }
    x, _ := vv.Interface()
    assert.Equal(t, expect, x)
}

func runDecoderTestUseNumber(t *testing.T, src string, expect interface{}) {
    vv, err := NewParser(src).Parse()
    if err != 0 { panic(err) }
    vvv, _ := vv.InterfaceUseNumber()
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
    runDecoderTest(t, `{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`, map[string]interface{}{"a":"123", "b":true, "c":false, "d":nil, "e": 2.4, "f": 1.2e15, "g":float64(1)})

    runDecoderTestUseNumber(t, `null`, nil)
    runDecoderTestUseNumber(t, `true`, true)
    runDecoderTestUseNumber(t, `false`, false)
    runDecoderTestUseNumber(t, `"hello, world \\ \/ \b \f \n \r \t \u666f 测试中文"`, "hello, world \\ / \b \f \n \r \t \u666f 测试中文")
    runDecoderTestUseNumber(t, `"\ud83d\ude00"`, "😀")
    runDecoderTestUseNumber(t, `0`, float64(0))
    runDecoderTestUseNumber(t, `-0`, float64(0))
    runDecoderTestUseNumber(t, `123456`, float64(123456))
    runDecoderTestUseNumber(t, `-12345`, float64(-12345))
    runDecoderTestUseNumber(t, `0.2`, 0.2)
    runDecoderTestUseNumber(t, `1.2`, 1.2)
    runDecoderTestUseNumber(t, `-0.2`, -0.2)
    runDecoderTestUseNumber(t, `-1.2`, -1.2)
    runDecoderTestUseNumber(t, `0e12`, 0e12)
    runDecoderTestUseNumber(t, `0e+12`, 0e+12)
    runDecoderTestUseNumber(t, `0e-12`, 0e-12)
    runDecoderTestUseNumber(t, `-0e12`, -0e12)
    runDecoderTestUseNumber(t, `-0e+12`, -0e+12)
    runDecoderTestUseNumber(t, `-0e-12`, -0e-12)
    runDecoderTestUseNumber(t, `2e12`, 2e12)
    runDecoderTestUseNumber(t, `2E12`, 2e12)
    runDecoderTestUseNumber(t, `2e+12`, 2e+12)
    runDecoderTestUseNumber(t, `2e-12`, 2e-12)
    runDecoderTestUseNumber(t, `-2e12`, -2e12)
    runDecoderTestUseNumber(t, `-2e+12`, -2e+12)
    runDecoderTestUseNumber(t, `-2e-12`, -2e-12)
    runDecoderTestUseNumber(t, `0.2e12`, 0.2e12)
    runDecoderTestUseNumber(t, `0.2e+12`, 0.2e+12)
    runDecoderTestUseNumber(t, `0.2e-12`, 0.2e-12)
    runDecoderTestUseNumber(t, `-0.2e12`, -0.2e12)
    runDecoderTestUseNumber(t, `-0.2e+12`, -0.2e+12)
    runDecoderTestUseNumber(t, `-0.2e-12`, -0.2e-12)
    runDecoderTestUseNumber(t, `1.2e12`, 1.2e12)
    runDecoderTestUseNumber(t, `1.2e+12`, 1.2e+12)
    runDecoderTestUseNumber(t, `1.2e-12`, 1.2e-12)
    runDecoderTestUseNumber(t, `-1.2e12`, -1.2e12)
    runDecoderTestUseNumber(t, `-1.2e+12`, -1.2e+12)
    runDecoderTestUseNumber(t, `-1.2e-12`, -1.2e-12)
    runDecoderTestUseNumber(t, `-1.2E-12`, -1.2e-12)
    runDecoderTestUseNumber(t, `["asd", "123", true, false, null, 2.4, 1.2e15, 1]`, []interface{}{"asd", "123", true, false, nil, 2.4, 1.2e15, float64(1)})
    runDecoderTestUseNumber(t, `{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`, map[string]interface{}{"a":"123", "b":true, "c":false, "d":nil, "e": 2.4, "f": 1.2e15, "g":float64(1)})
}

func TestLoads(t *testing.T) {
    _,i,e := Loads(`{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`)
    if e != nil {
        t.Fatal(e)
    }
    assert.Equal(t, map[string]interface{}{"a": "123", "b": true, "c": false, "d": nil, "e": 2.4, "f": 1.2e15, "g": float64(1)}, i)
    _,i,e = LoadsUseNumber(`{"a": "123", "b": true, "c": false, "d": null, "e": 2.4, "f": 1.2e15, "g": 1}`)
    if e != nil {
        t.Fatal(e)
    }
    assert.Equal(t, map[string]interface{}{"a": "123", "b": true, "c": false, "d": nil, "e": json.Number("2.4"), "f": json.Number("1.2e15"), "g": json.Number("1")}, i)
}

func TestParsehNotExist(t *testing.T) {
    s,err := NewParser(` { "xx" : [ 0, "" ] ,"yy" :{ "2": "" } } `).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    node := s.GetByPath("xx", 2)
    if node.Exists() {
        t.Fatalf("node: %v", node)
    }
    node = s.GetByPath("xx", 1)
    if !node.Exists() {
        t.Fatalf("node: %v", nil)
    }
    node = s.GetByPath("yy", "3")
    if node.Exists() {
        t.Fatalf("node: %v", node)
    }
    node = s.GetByPath("yy", "2")
    if !node.Exists() {
        t.Fatalf("node: %v", nil)
    }
}

func BenchmarkParser_Sonic(b *testing.B) {
    r, err := NewParser(_TwitterJson).Parse()
    if err != 0 {
        b.Fatal(err)
    }
    if err := r.LoadAll(); err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r, _ = NewParser(_TwitterJson).Parse()
        _ = r.LoadAll()
    }
}

func BenchmarkParser_Parallel_Sonic(b *testing.B) {
    r, _ := NewParser(_TwitterJson).Parse()
    if err := r.LoadAll(); err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            r, _ := NewParser(_TwitterJson).Parse()
            _ = r.LoadAll()
        }
    })
}

func BenchmarkParseOne_Sonic(b *testing.B) {
    ast, _ := NewParser(_TwitterJson).Parse()
    node, _ := ast.Get("statuses").Index(2).Get("id").Int64()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast, _ := NewParser(_TwitterJson).Parse()
        _, _ = ast.Get("statuses").Index(2).Get("id").Int64()
    }
}

func BenchmarkParseOne_Parallel_Sonic(b *testing.B) {
    ast, _ := NewParser(_TwitterJson).Parse()
    node, _ := ast.Get("statuses").Index(2).Get("id").Int64()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast, _ := NewParser(_TwitterJson).Parse()
            _, _ = ast.Get("statuses").Index(2).Get("id").Int64()
        }
    })
}

func BenchmarkParseSeven_Sonic(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast, _ := NewParser(_TwitterJson).Parse()
        node := ast.GetByPath("statuses", 3, "id")
        node = ast.GetByPath("statuses",  3, "user", "entities","description")
        node = ast.GetByPath("statuses",  3, "user", "entities","url","urls")
        node = ast.GetByPath("statuses",  3, "user", "entities","url")
        node = ast.GetByPath("statuses",  3, "user", "created_at")
        node = ast.GetByPath("statuses",  3, "user", "name")
        node = ast.GetByPath("statuses",  3, "text")
        if node.Check() != nil {
            b.Fail()
        }
    }
}

func BenchmarkParseSeven_Parallel_Sonic(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast, _ := NewParser(_TwitterJson).Parse()
            node := ast.GetByPath("statuses", 3, "id")
            node = ast.GetByPath("statuses",  3, "user", "entities","description")
            node = ast.GetByPath("statuses",  3, "user", "entities","url","urls")
            node = ast.GetByPath("statuses",  3, "user", "entities","url")
            node = ast.GetByPath("statuses",  3, "user", "created_at")
            node = ast.GetByPath("statuses",  3, "user", "name")
            node = ast.GetByPath("statuses",  3, "text")
            if node.Check() != nil {
                b.Fail()
            }
        }
    })
}