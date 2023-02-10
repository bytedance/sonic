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

package decoder

import (
    `encoding/json`
    `runtime`
    `runtime/debug`
    `strings`
    `sync`
    `testing`
    `time`
    `reflect`

    `github.com/bytedance/sonic/internal/rt`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
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

func TestGC(t *testing.T) {
    if debugSyncGC {
        return 
    }
    var w interface{}
    out, err := decode(TwitterJson, &w, true)
    if err != nil {
        t.Fatal(err)
    }
    if out != len(TwitterJson) {
        t.Fatal(out)
    }
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
            defer wg.Done()
            var w interface{}
            out, err := decode(TwitterJson, &w, true)
            if err != nil {
                t.Fatal(err)
            }
            if out != len(TwitterJson) {
                t.Fatal(out)
            }
            runtime.GC()
        }(wg)
    }
    wg.Wait()
}

var _BindingValue TwitterStruct

func init() {
    _ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}


func TestSkipMismatchTypeError(t *testing.T) {
    t.Run("struct", func(t *testing.T) {
        println("TestSkipError")
        type skiptype struct {
            A int `json:"a"`
            B string `json:"b"`

            Pass *int `json:"pass"`

            C struct{

                Pass4 interface{} `json:"pass4"`

                D struct{
                    E float32 `json:"e"`
                } `json:"d"`

                Pass2 int `json:"pass2"`

            } `json:"c"`

            E bool `json:"e"`
            F []int `json:"f"`
            G map[string]int `json:"g"`
            H bool `json:"h,string"`

            Pass3 int `json:"pass2"`

            I json.Number `json:"i"`
        }
        var obj, obj2 = &skiptype{Pass:new(int)}, &skiptype{Pass:new(int)}
        var data = `{"a":"","b":1,"c":{"d":true,"pass2":1,"pass4":true},"e":{},"f":"","g":[],"pass":null,"h":"1.0","i":true,"pass3":1}`
        d := NewDecoder(data)
        err := d.Decode(obj)
        err2 := json.Unmarshal([]byte(data), obj2)
        println(err2.Error())
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        assert.Equal(t, obj2, obj)
        if te, ok := err.(*MismatchTypeError); ok {
            assert.Equal(t, reflect.TypeOf(obj.I), te.Type)
            assert.Equal(t, strings.Index(data, `"i":t`)+4, te.Pos)
            println(err.Error())
        } else {
            t.Fatal("invalid error")
        }
    })
    t.Run("array", func(t *testing.T) {
        var obj, obj2 = &[]int{}, &[]int{}
        var data = `["",1,true]`
        d := NewDecoder(data)
        err := d.Decode(obj)
        err2 := json.Unmarshal([]byte(data), obj2)
        // println(err2.Error())
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        assert.Equal(t, obj2, obj)
    })
    t.Run("map", func(t *testing.T) {
        var obj, obj2 = &map[int]int{}, &map[int]int{}
        var data = `{"true" : { },"1":1,"2" : true,"3":3}`
        d := NewDecoder(data)
        err := d.Decode(obj)
        err2 := json.Unmarshal([]byte(data), obj2)
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        assert.Equal(t, obj2, obj)
    })
    t.Run("map error", func(t *testing.T) {
        var obj, obj2 = &map[int]int{}, &map[int]int{}
        var data = `{"true" : { ],"1":1,"2" : true,"3":3}`
        d := NewDecoder(data)
        err := d.Decode(obj)
        err2 := json.Unmarshal([]byte(data), obj2)
        println(err.Error())
        println(err2.Error())
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        // assert.Equal(t, obj2, obj)
    })
}

func TestDecodeCorrupt(t *testing.T) {
    var ds = []string{
        `{,}`,
        `{,"a"}`,
        `{"a":}`,
        `{"a":1,}`,
        `{"a":1,"b"}`,
        `{"a":1,"b":}`,
        `{,"a":1 "b":2}`,
        `{"a",:1 "b":2}`,
        `{"a":,1 "b":2}`,
        `{"a":1 "b",:2}`,
        `{"a":1 "b":,2}`,
        `{"a":1 "b":2,}`,
        `{"a":1 "b":2}`,
        `[,]`,
        `[,1]`,
        `[1,]`,
        `[,1,2]`,
        `[1,2,]`,
    }
    for _, d := range ds {
        var o interface{}
        _, err := decode(d, &o, false)
        if err == nil {
            t.Fatalf("%#v", d)
        }
        if !strings.Contains(err.Error(), "invalid char"){
            t.Fatal(err.Error())
        }
    }
}

func decode(s string, v interface{}, copy bool) (int, error) {
    d := NewDecoder(s)
    if copy {
        d.CopyString()
    }
    err := d.Decode(v)
    if err != nil {
        return 0, err
    }
    return d.i, err
}

func TestCopyString(t *testing.T) {
    var data []byte
    var dc *Decoder
    var err error
    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    dc.CopyString()
    var obj struct{
        A string
        B string
    }
    err = dc.Decode(&obj)
    if err != nil {
        t.Fatal(err)
    }
    data[6] = '1'
    if obj.A != "0" {
        t.Fatal(obj)
    }
    data[14] = '0'
    if obj.B != "1" {
        t.Fatal(obj)
    }

    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    err = dc.Decode(&obj)
    if err != nil {
        t.Fatal(err)
    }
    data[6] = '1'
    if obj.A != "1" {
        t.Fatal(obj)
    }
    data[14] = '0'
    if obj.B != "0" {
        t.Fatal(obj)
    }

    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    dc.CopyString()
    m := map[string]interface{}{}
    err = dc.Decode(&m)
    if err != nil {
        t.Fatal(err)
    }
    data[2] = 'C'
    data[6] = '1'
    if m["A"] != "0" {
        t.Fatal(m)
    }
    data[10] = 'D'
    data[14] = '0'
    if m["B"] != "1" {
        t.Fatal(m)
    }

    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    m = map[string]interface{}{}
    err = dc.Decode(&m)
    if err != nil {
        t.Fatal(err)
    }
    data[6] = '1'
    if m["A"] != "1" {
        t.Fatal(m)
    }
    data[14] = '0'
    if m["B"] != "0" {
        t.Fatal(m)
    }

    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    dc.CopyString()
    var x interface{}
    err = dc.Decode(&x)
    if err != nil {
        t.Fatal(err)
    }
    data[2] = 'C'
    data[6] = '1'
    m = x.(map[string]interface{})
    if m["A"] != "0" {
        t.Fatal(m)
    }
    data[10] = 'D'
    data[14] = '0'
    if m["B"] != "1" {
        t.Fatal(m)
    }

    data = []byte(`{"A":"0","B":"1"}`)
    dc = NewDecoder(rt.Mem2Str(data))
    dc.UseNumber()
    var y interface{}
    err = dc.Decode(&y)
    if err != nil {
        t.Fatal(err)
    }
    m = y.(map[string]interface{})
    data[6] = '1'
    if m["A"] != "1" {
        t.Fatal(m)
    }
    data[14] = '0'
    if m["B"] != "0" {
        t.Fatal(m)
    }
}

func TestDecoder_Basic(t *testing.T) {
    var v int
    pos, err := decode("12345", &v, false)
    assert.NoError(t, err)
    assert.Equal(t, 5, pos)
    assert.Equal(t, 12345, v)
}

func TestDecoder_Generic(t *testing.T) {
    var v interface{}
    pos, err := decode(TwitterJson, &v, false)
    assert.NoError(t, err)
    assert.Equal(t, len(TwitterJson), pos)
}

func TestDecoder_Binding(t *testing.T) {
    var v TwitterStruct
    pos, err := decode(TwitterJson, &v, false)
    assert.NoError(t, err)
    assert.Equal(t, len(TwitterJson), pos)
    assert.Equal(t, _BindingValue, v, 0)
    spew.Dump(v)
}

func TestDecoder_SetOption(t *testing.T) {
    var v interface{}
    d := NewDecoder("123")
    d.SetOptions(OptionUseInt64)
    err := d.Decode(&v)
    assert.NoError(t, err)
    assert.Equal(t, v, int64(123))
}

func TestDecoder_MapWithIndirectElement(t *testing.T) {
    var v map[string]struct { A [129]byte }
    _, err := decode(`{"":{"A":[1,2,3,4,5]}}`, &v, false)
    if x, ok := err.(SyntaxError); ok {
        println(x.Description())
    }
    require.NoError(t, err)
    assert.Equal(t, [129]byte{1, 2, 3, 4, 5}, v[""].A)
}

func BenchmarkDecoder_Generic_Sonic(b *testing.B) {
    var w interface{}
    _, _ = decode(TwitterJson, &w, true)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _, _ = decode(TwitterJson, &v, true)
    }
}

func BenchmarkDecoder_Generic_Sonic_Fast(b *testing.B) {
    var w interface{}
    _, _ = decode(TwitterJson, &w, false)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _, _ = decode(TwitterJson, &v, false)
    }
}

func BenchmarkDecoder_Generic_StdLib(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = json.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _ = json.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Binding_Sonic(b *testing.B) {
    var w TwitterStruct
    _, _ = decode(TwitterJson, &w, true)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _, _ = decode(TwitterJson, &v, true)
    }
}

func BenchmarkDecoder_Binding_Sonic_Fast(b *testing.B) {
    var w TwitterStruct
    _, _ = decode(TwitterJson, &w, false)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _, _ = decode(TwitterJson, &v, false)
    }
}

func BenchmarkDecoder_Binding_StdLib(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = json.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _ = json.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Parallel_Generic_Sonic(b *testing.B) {
    var w interface{}
    _, _ = decode(TwitterJson, &w, true)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _, _ = decode(TwitterJson, &v, true)
        }
    })
}

func BenchmarkDecoder_Parallel_Generic_Sonic_Fast(b *testing.B) {
    var w interface{}
    _, _ = decode(TwitterJson, &w, false)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _, _ = decode(TwitterJson, &v, false)
        }
    })
}

func BenchmarkDecoder_Parallel_Generic_StdLib(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = json.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _ = json.Unmarshal(m, &v)
        }
    })
}

func BenchmarkDecoder_Parallel_Binding_Sonic(b *testing.B) {
    var w TwitterStruct
    _, _ = decode(TwitterJson, &w, true)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _, _ = decode(TwitterJson, &v, true)
        }
    })
}

func BenchmarkDecoder_Parallel_Binding_Sonic_Fast(b *testing.B) {
    var w TwitterStruct
    _, _ = decode(TwitterJson, &w, false)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _, _ = decode(TwitterJson, &v, false)
        }
    })
}

func BenchmarkDecoder_Parallel_Binding_StdLib(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = json.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _ = json.Unmarshal(m, &v)
        }
    })
}

func BenchmarkSkip_Sonic(b *testing.B) {
    var data = rt.Str2Mem(TwitterJson)
    if ret, _ := Skip(data); ret < 0 {
        b.Fatal()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _, _ = Skip(data)
    }
}