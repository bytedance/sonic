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

package unit_test

import (
    `io`
    `strings`
    `testing`

    `github.com/bytedance/sonic/decoder`
    jsoniter "github.com/json-iterator/go"
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

var (
    defaultBufferSize uint = 4096
    _Single_JSON           = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"} { `
    _Triple_JSON           = `{"aaaaa":"` + strings.Repeat("b", int(defaultBufferSize)) + `"}{ } {"11111":"` +
        strings.Repeat("2", int(defaultBufferSize)) + `"} b {}`
)

func TestDecodeSingle(t *testing.T) {
    var str = _Single_JSON

    var r1 = strings.NewReader(str)
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    var r2 = strings.NewReader(str)
    var v2 map[string]interface{}
    var d2 = decoder.NewStreamDecoder(r2)

    require.Equal(t, d1.More(), d2.More())
    es1 := d1.Decode(&v1)
    ee1 := d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())
}

func TestDecodeMulti(t *testing.T) {
    var str = _Triple_JSON

    var r1 = strings.NewReader(str)
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    var r2 = strings.NewReader(str)
    var v2 map[string]interface{}
    var d2 = decoder.NewStreamDecoder(r2)

    require.Equal(t, d1.More(), d2.More())
    es1 := d1.Decode(&v1)
    ee1 := d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())
    // fmt.Printf("v:%#v\n", v1)

    require.Equal(t, d1.More(), d2.More())
    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)

    require.Equal(t, d1.More(), d2.More())
    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
}

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

func TestDecodeHalt(t *testing.T) {
    var str = _Triple_JSON
    var r1 = NewHaltReader(str, testHalts())
    var r2 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var v2 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    var d2 = decoder.NewStreamDecoder(r2)

    require.Equal(t, d1.More(), d2.More())
    err1 := d1.Decode(&v1)
    err2 := d2.Decode(&v2)
    assert.Equal(t, err1, err2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    require.Equal(t, d1.More(), d2.More())
    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)

    require.Equal(t, d1.More(), d2.More())
    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
}

func TestMore(t *testing.T) {
    var str = _Triple_JSON
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = decoder.NewStreamDecoder(r2)
    var r1 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    require.Nil(t, d1.Decode(&v1))
    require.Nil(t, d2.Decode(&v2))
    require.Equal(t, d1.More(), d2.More())

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    require.Equal(t, d1.More(), d2.More())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    require.Equal(t, d1.More(), d2.More())

    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)
    require.Equal(t, d1.More(), d2.More())

    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
    require.Equal(t, d1.More(), d2.More())
}
