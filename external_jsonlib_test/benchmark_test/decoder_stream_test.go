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

package benchmark_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/bytedance/sonic/decoder"
	jsoniter "github.com/json-iterator/go"
)

var (
    defaultBufferSize uint = 4096
    _Single_JSON           = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"} { `
    _Double_JSON           = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"}    {"11111":"` + strings.Repeat("2", int(defaultBufferSize)) + `"}`
)

type HaltReader struct {
    halts map[int]bool
    buf   string
    p     int
}

func NewHaltReader(buf string, halts map[int]bool) *HaltReader {
    return &HaltReader{
        halts: halts,
        buf:   buf,
        p:     0,
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

var testHalts = func() map[int]bool {
    return map[int]bool{
        1:  true,
        10: true,
        20: true}
}

func BenchmarkDecodeStream_Jsoniter(b *testing.B) {
    b.Run("single", func(b *testing.B) {
        var str = _Single_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("halt", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}


func BenchmarkDecodeStream_Std(b *testing.B) {
    b.Run("single", func(b *testing.B) {
        var str = _Single_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("halt", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}



func BenchmarkDecodeStream_Sonic(b *testing.B) {
    b.Run("single", func(b *testing.B) {
        var str = _Single_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := decoder.NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := decoder.NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("halt", func(b *testing.B) {
        var str = _Double_JSON
        b.SetBytes(int64(len(str)))
        for i := 0; i < b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := decoder.NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}

