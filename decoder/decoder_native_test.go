//go:build (amd64 && go1.17 && !go1.24) || (arm64 && go1.20 && !go1.24)
// +build amd64,go1.17,!go1.24 arm64,go1.20,!go1.24

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
	"encoding/json"
	"fmt"
	_ "reflect"
	"strings"
	_ "strings"
	"testing"
	"time"

	"github.com/bytedance/sonic/internal/rt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func BenchmarkSkipValidate(b *testing.B) {
    type skiptype struct {
        A int `json:"a"` // mismatched
        B string `json:"-"` // ommited
        C [1]int `json:"c"` // fast int
        D struct {} `json:"d"` // empty struct
        E map[string]int `json:"e"` // mismatched elem
        F json.RawMessage `json:"f"` // unmarshaler
        // Unknonwn
    }
    type C struct {
        name string
        json string
        expTime float64
    }
    var sam = map[int]interface{}{}
    for i := 0; i < 1; i++ {
        sam[i] = _BindingValue
    }
    comptd, err := json.Marshal(sam)
    if err != nil {
        b.Fatal("invalid json")
    }
    compt := string(comptd)
    var cases = []C{
        {"mismatched", `{"a":`+compt+`}`, 5},
        {"ommited", `{"b":`+compt+`}`, 5},
        {"number", `{"c":[`+strings.Repeat("-1.23456e-19,", 1000)+`1]}`, 1.5},
        {"unknown", `{"unknown":`+compt+`}`, 5},
        {"empty", `{"d":`+compt+`}`, 5},
        {"mismatched elem", `{"e":`+compt+`}`, 5},
        {"unmarshaler", `{"f":`+compt+`}`, 3},
    }
    _ = NewDecoder(`{}`).Decode(&skiptype{})

    var avg1, avg2 time.Duration
    for _, c := range cases {
        b.Run(c.name, func(b *testing.B) {
            b.Run("validate", func(b *testing.B) {
                b.ResetTimer()
                t1 := time.Now()
                for i := 0; i < b.N; i++ {
                    var obj1 = &skiptype{}
                    // validate skip
                    d := NewDecoder(c.json)
                    _ = d.Decode(obj1)
                }
                d1 := time.Since(t1)
                avg1 = d1/time.Duration(b.N)
            })
            b.Run("fast", func(b *testing.B) {
                b.ResetTimer()
                t2 := time.Now()
                for i := 0; i < b.N; i++ {
                    var obj2 = &skiptype{}
                    // fask skip
                    d := NewDecoder(c.json)
                    d.SetOptions(OptionNoValidateJSON)
                    _ = d.Decode(obj2)
                }
                d2 := time.Since(t2)
                avg2 = d2/time.Duration(b.N)
            })
            // fast skip must be expTime x faster
            require.True(b, float64(avg1)/float64(avg2) >  c.expTime, fmt.Sprintf("%v/%v=%v", avg1, avg2, float64(avg1)/float64(avg2)))
        })
    }
}


func TestSkipMismatchTypeAmd64Error(t *testing.T) {
    // t.Run("struct", func(t *testing.T) {
    //     println("TestSkipError")
    //     type skiptype struct {
    //         A int `json:"a"`
    //         B string `json:"b"`

    //         Pass *int `json:"pass"`

    //         C struct{

    //             Pass4 interface{} `json:"pass4"`

    //             D struct{
    //                 E float32 `json:"e"`
    //             } `json:"d"`

    //             Pass2 int `json:"pass2"`

    //         } `json:"c"`

    //         E bool `json:"e"`
    //         F []int `json:"f"`
    //         G map[string]int `json:"g"`
    //         H bool `json:"h,string"`

    //         Pass3 int `json:"pass2"`

    //         I json.Number `json:"i"`
    //     }
    //     var obj, obj2 = &skiptype{Pass:new(int)}, &skiptype{Pass:new(int)}
    //     var data = `{"a":"","b":1,"c":{"d":true,"pass2":1,"pass4":true},"e":{},"f":"","g":[],"pass":null,"h":"1.0","i":true,"pass3":1}`
    //     d := NewDecoder(data)
    //     err := d.Decode(obj)
    //     err2 := json.Unmarshal([]byte(data), obj2)
    //     println(err2.Error())
    //     assert.Equal(t, err2 == nil, err == nil)
    //     // assert.Equal(t, len(data), d.i)
    //     assert.Equal(t, obj2, obj)
    //     if te, ok := err.(*MismatchTypeError); ok {
    //         assert.Equal(t, reflect.TypeOf(obj.I), te.Type)
    //         assert.Equal(t, strings.Index(data, `"i":t`)+4, te.Pos)
    //         println(err.Error())
    //     } else {
    //         t.Fatal("invalid error")
    //     }
    // })
    t.Run("short array", func(t *testing.T) {
        var obj, obj2 = &[]int{}, &[]int{}
        var data = `[""]`
        d := NewDecoder(data)
        err := d.Decode(obj)
        err2 := json.Unmarshal([]byte(data), obj2)
        // println(err2.Error())
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        assert.Equal(t, obj2, obj)
    })

    t.Run("int ", func(t *testing.T) {
        var obj int = 123
        var obj2 int = 123
        var data = `[""]`
        d := NewDecoder(data)
        err := d.Decode(&obj)
        err2 := json.Unmarshal([]byte(data), &obj2)
        println(err.Error(), obj, obj2)
        assert.Equal(t, err2 == nil, err == nil)
        // assert.Equal(t, len(data), d.i)
        assert.Equal(t, obj2, obj)
    })

    t.Run("array", func(t *testing.T) {
        var obj, obj2 = &[]int{}, &[]int{}
        var data = `["",true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true,true]`
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

func TestDecoder_SetOption(t *testing.T) {
    var v interface{}
    d := NewDecoder("123")
    d.SetOptions(OptionUseInt64)
    err := d.Decode(&v)
    assert.NoError(t, err)
    assert.Equal(t, v, int64(123))
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