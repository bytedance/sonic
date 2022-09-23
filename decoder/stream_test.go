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

    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

var (
    _Single_JSON = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"} { `
    _Double_JSON = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"}    {"11111":"` + strings.Repeat("2", int(defaultBufferSize)) + `"}`     
    _Triple_JSON = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"}{ } {"11111":"` + 
    strings.Repeat("2", int(defaultBufferSize))+`"} b {}`
)

func TestStreamError(t *testing.T) {
    var qs = []string{
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`""}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":]`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":1x`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":1x}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":1x]`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":t`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":t}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":true]`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":f`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":f}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":false]`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":n`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":n}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":null]`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":"`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":"a`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":"a}`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":"a"`,
        `{`+strings.Repeat(" ", int(defaultBufferSize))+`"":"a"]`,
    }

    for i, q := range qs {
        var qq = []byte(q[:int(defaultBufferSize)]+strings.Repeat(" ", i*100)+q[int(defaultBufferSize):])
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
    require.Nil(t, d1.Decode(&v1))
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)
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

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, d1.InputOffset(), d2.InputOffset())
}

func BenchmarkDecodeStream_Std(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
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
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
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
}