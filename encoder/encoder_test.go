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

package encoder

import (
    `bytes`
    `encoding/json`
    `runtime`
    `runtime/debug`
    `strconv`
    `sync`
    `testing`
    `time`
    `unicode/utf8`

    gojson `github.com/goccy/go-json`
    `github.com/json-iterator/go`
    `github.com/stretchr/testify/require`
    `github.com/bytedance/sonic/internal/rt`
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
    out, err := Encode(_GenericValue, 0)
    if err != nil {
        t.Fatal(err)
    }
    n := len(out)
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup, size int)  {
            defer wg.Done()
            out, err := Encode(_GenericValue, 0)
            if err != nil {
                t.Fatal(err)
            }
            if len(out) != size {
                t.Fatal(len(out), size)
            }
            runtime.GC()
        }(wg, n)
    }
    wg.Wait()
}

func runEncoderTest(t *testing.T, fn func(string)string, exp string, arg string) {
    require.Equal(t, exp, fn(arg))
}

func TestEncoder_String(t *testing.T) {
    runEncoderTest(t, Quote, `""`                                                 , "")
    runEncoderTest(t, Quote, `"hello, world"`                                     , "hello, world")
    runEncoderTest(t, Quote, `"hello啊啊啊aa"`                                    , "hello啊啊啊aa")
    runEncoderTest(t, Quote, `"hello\\\"world"`                                   , "hello\\\"world")
    runEncoderTest(t, Quote, `"hello\n\tworld"`                                   , "hello\n\tworld")
    runEncoderTest(t, Quote, `"hello\u0000\u0001world"`                           , "hello\x00\x01world")
    runEncoderTest(t, Quote, `"hello\u0000\u0001world"`                           , "hello\x00\x01world")
    runEncoderTest(t, Quote, `"Cartoonist, Illustrator, and T-Shirt connoisseur"` , "Cartoonist, Illustrator, and T-Shirt connoisseur")
}

type StringStruct struct {
    X *int        `json:"x,string,omitempty"`
    Y []int       `json:"y"`
    Z json.Number `json:"z,string"`
    W string      `json:"w,string"`
}

func TestEncoder_FieldStringize(t *testing.T) {
    x := 12345
    v := StringStruct{X: &x, Y: []int{1, 2, 3}, Z: "4567456", W: "asdf"}
    r, e := Encode(v, 0)
    require.NoError(t, e)
    println(string(r))
}

type MarshalerImpl struct {
    X int
}

func (self *MarshalerImpl) MarshalJSON() ([]byte, error) {
    ret := []byte(strconv.Itoa(self.X))
    return append(ret, "    "...), nil
}

type MarshalerStruct struct {
    V MarshalerImpl
}

func TestEncoder_Marshaler(t *testing.T) {
    v := MarshalerStruct{V: MarshalerImpl{X: 12345}}
    ret, err := Encode(&v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":12345    }`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":{"X":12345}}`, string(ret))

    ret2, err2 := Encode(&v, 0)
    require.NoError(t, err2)
    require.Equal(t, `{"V":12345    }`, string(ret2))
    ret3, err3 := Encode(v, CompactMarshaler)
    require.NoError(t, err3)
    require.Equal(t, `{"V":{"X":12345}}`, string(ret3))
}

type MarshalerErrorStruct struct {
    V MarshalerImpl
}

func (self *MarshalerErrorStruct) MarshalJSON() ([]byte, error) {
    return []byte(`[""] {`), nil
}

func TestMarshalerError(t *testing.T) {
    v := MarshalerErrorStruct{}
    ret, err := Encode(&v, 0)
    require.EqualError(t, err, `invalid Marshaler output json syntax at 5: "[\"\"] {"`)
    require.Equal(t, []byte(nil), ret)
}

type RawMessageStruct struct {
    X json.RawMessage
}

func TestEncoder_RawMessage(t *testing.T) {
    rms := RawMessageStruct{
        X: json.RawMessage("123456    "),
    }
    ret, err := Encode(&rms, 0)
    require.NoError(t, err)
    require.Equal(t, `{"X":123456    }`, string(ret))

    ret, err = Encode(&rms, CompactMarshaler)
    require.NoError(t, err)
    require.Equal(t, `{"X":123456}`, string(ret))
}

type TextMarshalerImpl struct {
    X string
}

func (self *TextMarshalerImpl) MarshalText() ([]byte, error) {
    return []byte(self.X), nil
}

type TextMarshalerStruct struct {
    V TextMarshalerImpl
}

func TestEncoder_TextMarshaler(t *testing.T) {
    v := TextMarshalerStruct{V: TextMarshalerImpl{X: (`{"a"}`)}}
    ret, err := Encode(&v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":"{\"a\"}"}`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret))

    ret2, err2 := Encode(&v, NoQuoteTextMarshaler)
    require.NoError(t, err2)
    require.Equal(t, `{"V":{"a"}}`, string(ret2))
    ret3, err3 := Encode(v, NoQuoteTextMarshaler)
    require.NoError(t, err3)
    require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret3))
}

func TestEncoder_Marshal_EscapeHTML(t *testing.T) {
    v := map[string]TextMarshalerImpl{"&&":{"<>"}}
    ret, err := Encode(v, EscapeHTML)
    require.NoError(t, err)
    require.Equal(t, `{"\u0026\u0026":{"X":"\u003c\u003e"}}`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"&&":{"X":"<>"}}`, string(ret))

    // “ is \xe2\x80\x9c, and ” is \xe2\x80\x9d,
    // similar as HTML escaped chars \u2028(\xe2\x80\xa8) and \u2029(\xe2\x80\xa9)
    m := map[string]string{"test": "“123”"}
    ret, err = Encode(m, EscapeHTML)
    require.Equal(t, string(ret), `{"test":"“123”"}`)
    require.NoError(t, err)

    m = map[string]string{"K": "\u2028\u2028\xe2"}
    ret, err = Encode(m, EscapeHTML)
    require.Equal(t, string(ret), "{\"K\":\"\\u2028\\u2028\xe2\"}")
    require.NoError(t, err)
}

func TestEncoder_EscapeHTML(t *testing.T) {
    // test data from libfuzzer
    test := []string{
        "&&&&&&&&&&&&&&&&&&&&&&&\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&",
        "{\"\"\u2028\x94\xe2\x00\x00\x00\x00\x00\x00\x00\x00\u2028\x80\u2028\x80\u2028\xe2\u2028\x8a\u2028⑀\xa8\x8a\xa8\xe2\u2028\xe2\u2028\xe2\u2028\xe2\u2000\x8d\xe2\u2028\xe2\u2028\xe2\xe2\xa8\"}",
    }
    for _, s := range(test) {
        data := []byte(s)
        sdst := HTMLEscape(nil, data)
        var dst bytes.Buffer
        json.HTMLEscape(&dst, data)
        require.Equal(t, string(sdst), dst.String())
    }
}

type validTest struct {
    data string
    pos  int
}

func TestValidJsonOk(t *testing.T) {
    tests := []validTest{
        { `null`, 0}, 
        { ` 123`, 1},
        { "  \"\u2028\\\"\\u2029 你好\"", 2},
        { `   {"key":123}`, 3},
        { `    [true, false, {}, [], [{}]]`, 4},
        {"{\"S\":\"\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\\f\"}", 0},

         /* not valid the escaped surrogate half */
        { "\"\\uDE1D\"", 0},
        { "\"\\uDB00\\uDB00\"", 0},
    }
    for _, in := range(tests) {
        ok, start := Valid([]byte(in.data))
        jok := json.Valid([]byte(in.data)) && utf8.Valid([]byte(in.data))
        require.True(t, ok)
        require.True(t, jok)
        require.Equal(t, start, in.pos)
    }
}

func TestValidJsonFalse(t *testing.T) {
    tests := []validTest {
        { `nul`, 2},
        { `fale`, 3},
        { `-1.23e`, 4},
        {  "\u2028\\\"\\u2029 你好", 0},
        { "\"\xf8\x90\x90\x80\x80\"",  0},
        { "{\"\xff\":123}",  1},
        { "{{}}", 1},
        { `"\"\x\""`, 4},
        { `"\"\uD88X\""`, 8},
    }
    for _, in := range(tests) {
        println(in.data)
        ok, start := Valid([]byte(in.data))
        jok := json.Valid([]byte(in.data)) && utf8.Valid([]byte(in.data))
        require.False(t, jok)
        require.False(t, ok)
        require.Equal(t, start, in.pos)
    }
}

func TestEncoder_Marshal_EscapeHTML_LargeJson(t *testing.T) {
    buf1, err1 := Encode(&_BindingValue, SortMapKeys | EscapeHTML)
    require.NoError(t, err1)
    buf2, err2 :=json.Marshal(&_BindingValue)
    require.NoError(t, err2)
    require.Equal(t, buf1, buf2)
}

var _GenericValue interface{}
var _BindingValue TwitterStruct

func init() {
    _ = json.Unmarshal([]byte(TwitterJson), &_GenericValue)
    _ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func TestEncoder_Generic(t *testing.T) {
    v, e := Encode(_GenericValue, 0)
    require.NoError(t, e)
    println(string(v))
}

func TestEncoder_Binding(t *testing.T) {
    v, e := Encode(_BindingValue, 0)
    require.NoError(t, e)
    println(string(v))
}

func TestEncoder_MapSortKey(t *testing.T) {
    m := map[string]string {
        "C": "third",
        "D": "forth",
        "A": "first",
        "F": "sixth",
        "E": "fifth",
        "B": "second",
    }
    v, e := Encode(m, SortMapKeys)
    require.NoError(t, e)
    require.Equal(t, `{"A":"first","B":"second","C":"third","D":"forth","E":"fifth","F":"sixth"}`, string(v))
}

func BenchmarkEncoder_Generic_Sonic(b *testing.B) {
    _, _ = Encode(_GenericValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(_GenericValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    }
}

func BenchmarkEncoder_Generic_Sonic_Fast(b *testing.B) {
    _, _ = Encode(_GenericValue, 0)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(_GenericValue, 0)
    }
}

func BenchmarkEncoder_Generic_JsonIter(b *testing.B) {
    _, _ = jsoniter.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = jsoniter.Marshal(_GenericValue)
    }
}

func BenchmarkEncoder_Generic_GoJson(b *testing.B) {
    _, _ = gojson.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = gojson.Marshal(_GenericValue)
    }
}

func BenchmarkEncoder_Generic_StdLib(b *testing.B) {
    _, _ = json.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = json.Marshal(_GenericValue)
    }
}

func BenchmarkEncoder_Binding_Sonic(b *testing.B) {
    _, _ = Encode(&_BindingValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(&_BindingValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    }
}

func BenchmarkEncoder_Binding_Sonic_Fast(b *testing.B) {
    _, _ = Encode(&_BindingValue, NoQuoteTextMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(&_BindingValue, NoQuoteTextMarshaler)
    }
}

func BenchmarkEncoder_Binding_JsonIter(b *testing.B) {
    _, _ = jsoniter.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = jsoniter.Marshal(&_BindingValue)
    }
}

func BenchmarkEncoder_Binding_GoJson(b *testing.B) {
    _, _ = gojson.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = gojson.Marshal(&_BindingValue)
    }
}

func BenchmarkEncoder_Binding_StdLib(b *testing.B) {
    _, _ = json.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = json.Marshal(&_BindingValue)
    }
}

func BenchmarkEncoder_Parallel_Generic_Sonic(b *testing.B) {
    _, _ = Encode(_GenericValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(_GenericValue, SortMapKeys | EscapeHTML | CompactMarshaler)
        }
    })
}

func BenchmarkEncoder_Parallel_Generic_Sonic_Fast(b *testing.B) {
    _, _ = Encode(_GenericValue, NoQuoteTextMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(_GenericValue, NoQuoteTextMarshaler)
        }
    })
}

func BenchmarkEncoder_Parallel_Generic_JsonIter(b *testing.B) {
    _, _ = jsoniter.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = jsoniter.Marshal(_GenericValue)
        }
    })
}

func BenchmarkEncoder_Parallel_Generic_GoJson(b *testing.B) {
    _, _ = gojson.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = gojson.Marshal(_GenericValue)
        }
    })
}

func BenchmarkEncoder_Parallel_Generic_StdLib(b *testing.B) {
    _, _ = json.Marshal(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = json.Marshal(_GenericValue)
        }
    })
}

func BenchmarkEncoder_Parallel_Binding_Sonic(b *testing.B) {
    _, _ = Encode(&_BindingValue, SortMapKeys | EscapeHTML | CompactMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(&_BindingValue, SortMapKeys | EscapeHTML | CompactMarshaler)
        }
    })
}

func BenchmarkEncoder_Parallel_Binding_Sonic_Fast(b *testing.B) {
    _, _ = Encode(&_BindingValue, NoQuoteTextMarshaler)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(&_BindingValue, NoQuoteTextMarshaler)
        }
    })
}

func BenchmarkEncoder_Parallel_Binding_JsonIter(b *testing.B) {
    _, _ = jsoniter.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = jsoniter.Marshal(&_BindingValue)
        }
    })
}

func BenchmarkEncoder_Parallel_Binding_GoJson(b *testing.B) {
    _, _ = gojson.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = gojson.Marshal(&_BindingValue)
        }
    })
}

func BenchmarkEncoder_Parallel_Binding_StdLib(b *testing.B) {
    _, _ = json.Marshal(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = json.Marshal(&_BindingValue)
        }
    })
}

func BenchmarkHTMLEscape_Sonic(b *testing.B) {
    jsonByte := []byte(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    var buf []byte
    for i := 0; i < b.N; i++ {
        buf = HTMLEscape(nil, jsonByte)
    }
    _ = buf
}

func BenchmarkHTMLEscape_StdLib(b *testing.B) {
    jsonByte := []byte(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    var buf []byte
    for i := 0; i < b.N; i++ {
        out := bytes.NewBuffer(make([]byte, 0, len(TwitterJson) * 6 / 5))
        json.HTMLEscape(out, jsonByte)
        buf = out.Bytes()
    }
    _ = buf
}


func BenchmarkValidate_Sonic(b *testing.B) {
    var data = rt.Str2Mem(TwitterJson)
    ok, s := Valid(data)
    if !ok {
        b.Fatal(s)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _, _ = Valid(data)
    }
}

func BenchmarkValidate_Std(b *testing.B) {
    var data = rt.Str2Mem(TwitterJson)
    if !json.Valid(data) {
        b.Fatal()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        _ = json.Valid(data)
    }
}

func BenchmarkCompact_Std(b *testing.B) {
    var data = rt.Str2Mem(TwitterJson)
    var dst = bytes.NewBuffer(nil)
    if err := json.Compact(dst, data); err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        dst.Reset()
        _ = json.Compact(dst, data)
    }
}