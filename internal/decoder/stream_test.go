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
    `bytes`
    `encoding/json`
    `io`
    `io/ioutil`
    `strings`
    `testing`

    `github.com/bytedance/sonic/option`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

var (
    DefaultBufferSize = option.DefaultDecoderBufferSize
    _Single_JSON = `{`+`"aaaaa":"` + strings.Repeat("b", int(DefaultBufferSize)) + `"}` + strings.Repeat(" ", int(DefaultBufferSize)) + `{`
    _Double_JSON = `{"aaaaa":"` + strings.Repeat("b", int(DefaultBufferSize)) + `"}` + strings.Repeat(" ", int(DefaultBufferSize)) + `{"11111":"` + strings.Repeat("2", int(DefaultBufferSize)) + `"}`     
    _Triple_JSON = `{"aaaaa":"` + strings.Repeat("b", int(DefaultBufferSize)) + `"}{ } {"11111":"` + 
    strings.Repeat("2", int(DefaultBufferSize))+`"} b {}`
    _SMALL_JSON = `{"a":"b"} [1] {"c":"d"} [2]`
)

func TestStreamError(t *testing.T) {
    var qs = []string{
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`""}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":]`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":1x`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":1x}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":1x]`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":t`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":t}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":true]`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":f`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":f}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":false]`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":n`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":n}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":null]`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":"`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":"a`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":"a}`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":"a"`,
        `{`+strings.Repeat(" ", int(DefaultBufferSize))+`"":"a"]`,
    }

    for i, q := range qs {
        var qq = []byte(q[:int(DefaultBufferSize)]+strings.Repeat(" ", i*100)+q[int(DefaultBufferSize):])
        var obj interface{}
        require.NotNil(t, NewStreamDecoder(bytes.NewReader(qq)).Decode(&obj))
    }
}

func TestDecodeEmpty(t *testing.T) {
    var str = ``
    var r1 = strings.NewReader(str)
    var v1 interface{}
    var d1 = json.NewDecoder(r1)
    var r2 = strings.NewReader(str)
    var v2 interface{}
    var d2 = NewStreamDecoder(r2)
    es1 := d1.Decode(&v1)
    ee1 := d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    assert.Equal(t, v1, v2)
}

func TestDecodeRecurse(t *testing.T) {
    var str = _Single_JSON
    var r1 = bytes.NewBufferString(str)
    var v1 interface{}
    var d1 = json.NewDecoder(r1)
    var r2 = bytes.NewBufferString(str)
    var v2 interface{}
    var d2 = NewStreamDecoder(r2)

    require.Equal(t, d1.More(), d2.More())
    es1 := d1.Decode(&v1)
    ee1 := d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    assert.Equal(t, v1, v2)

    require.Equal(t, d1.More(), d2.More())
    r1.WriteString(str[1:])
    r2.WriteString(str[1:])

    require.Equal(t, d1.More(), d2.More())
    es1 = d1.Decode(&v1)
    ee1 = d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    println(es1)
    assert.Equal(t, v1, v2)
    require.Equal(t, d1.More(), d2.More())
}

type HaltReader struct {
    halts map[int]bool
    buf string
    p int
}

func NewHaltReader(buf string, halts map[int]bool) *HaltReader {
    return &HaltReader{
        halts: halts,
        buf: buf,
        p: 0,
    }
}

func (self *HaltReader) Read(p []byte) (int, error) {
    t := 0
    for ; t < len(p); {
        if self.p >= len(self.buf) {
            return t, io.EOF
        }
        if b, ok := self.halts[self.p]; b {
            self.halts[self.p] = false
            return t, nil
        } else if ok {
            delete(self.halts, self.p)
            return 0, nil
        }
        p[t] = self.buf[self.p]
        self.p++
        t++
    }
    return t, nil
}

func (self *HaltReader) Reset(buf string) {
    self.p = 0
    self.buf = buf
}

var testHalts = func () map[int]bool {
    return map[int]bool{
        1: true,
        10:true,
        20: true}
}

func TestBuffered(t *testing.T) {
    var str = _Triple_JSON
    var r1 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var d1 = json.NewDecoder(r1)
  
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)

    require.Equal(t, d1.More(), d2.More())
    require.Nil(t, d1.Decode(&v1))
    require.Nil(t, d2.Decode(&v2))
    left1, err1 := ioutil.ReadAll(d1.Buffered())
    require.Nil(t, err1)
    left2, err2 := ioutil.ReadAll(d2.Buffered())
    require.Nil(t, err2)
    require.Equal(t, d1.InputOffset(), d2.InputOffset())
    min := len(left1)
    if min > len(left2) {
        min = len(left2)
    }
    require.Equal(t, left1[:min], left2[:min])

    require.Equal(t, d1.More(), d2.More())
    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    println(str[d1.InputOffset()-5:d1.InputOffset()+5])
    assert.Equal(t, d1.InputOffset(), d2.InputOffset()-1)

    require.Equal(t, d1.More(), d2.More())
    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    println(str[d1.InputOffset()-5:d1.InputOffset()+5])
    assert.Equal(t, d1.InputOffset(), d2.InputOffset()-1)
}

func BenchmarkDecodeStream_Std(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        var r1 = bytes.NewBufferString(str)
        dc := json.NewDecoder(r1)
        for i:=0; i<b.N; i++ {
            var v1 map[string]interface{}
            e := dc.Decode(&v1)
            if e != nil {
                b.Fatal(e)
            }
            r1.WriteString(str[1:])
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            if dc.More() {
                _ = dc.Decode(&v1)
            }
        }
    })

    b.Run("4x", func (b *testing.B) {
        var str = _Double_JSON + strings.Repeat(" ", int(DefaultBufferSize-10)) + _Double_JSON
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            for dc.More() {
                e := dc.Decode(&v1)
                if e != nil {
                    b.Fatal(e)
                }
            }
        }
    })
    
    b.Run("halt", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("small", func (b *testing.B) {
        var str = _SMALL_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 interface{}
            dc := json.NewDecoder(r1)
            for dc.More() {
                e := dc.Decode(&v1)
                if e != nil {
                    b.Fatal(e)
                }
            }
        }
    })
}

// func BenchmarkDecodeError_Sonic(b *testing.B) {
//     var str = `\b测试1234`
//     for i:=0; i<b.N; i++ {
//         var v1 map[string]interface{}
//         _ = NewDecoder(str).Decode(&v1)
//     }
// }

func BenchmarkDecodeStream_Sonic(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        var r1 = bytes.NewBufferString(str)
        dc := NewStreamDecoder(r1)
        for i:=0; i<b.N; i++ {
            var v1 map[string]interface{}
            e := dc.Decode(&v1)
            if e != nil {
                b.Fatal(e)
            }
            r1.WriteString(str[1:])
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            if dc.More() {
                _ = dc.Decode(&v1)
            }
        }
    })

    b.Run("4x", func (b *testing.B) {
        var str = _Double_JSON + strings.Repeat(" ", int(DefaultBufferSize-10)) + _Double_JSON
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            // println("loop")
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            for dc.More() {
                e := dc.Decode(&v1)
                if e != nil {
                    b.Fatal(e)
                }
            }
        }
    })

    b.Run("halt", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("small", func (b *testing.B) {
        var str = _SMALL_JSON
        for i:=0; i<b.N; i++ {
            // println("one loop")
            var r1 = strings.NewReader(str)
            var v1 interface{}
            dc := NewStreamDecoder(r1)
            for dc.More() {
                e := dc.Decode(&v1)
                if e != nil {
                    b.Fatal(e)
                }
            }
        }
    })
}