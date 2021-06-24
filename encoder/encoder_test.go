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
    `encoding/json`
    `strconv`
    `testing`

    gojson `github.com/goccy/go-json`
    `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
)

func runEncoderTest(t *testing.T, fn func(string)string, exp string, arg string) {
    assert.Equal(t, exp, fn(arg))
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
    r, e := Encode(v)
    assert.Nil(t, e)
    println(string(r))
}

type MarshalerImpl struct {
    X int
}

func (self *MarshalerImpl) MarshalJSON() ([]byte, error) {
    return []byte(strconv.Itoa(self.X)), nil
}

type MarshalerStruct struct {
    V MarshalerImpl
}

func TestEncoder_Marshaler(t *testing.T) {
    v := MarshalerStruct{V: MarshalerImpl{X: 12345}}
    ret, err := Encode(&v)
    assert.Nil(t, err)
    assert.Equal(t, `{"V":12345}`, string(ret))
    ret, err = Encode(v)
    assert.Nil(t, err)
    assert.Equal(t, `{"V":{"X":12345}}`, string(ret))
}

type RawMessageStruct struct {
    X json.RawMessage
}

func TestEncoder_RawMessage(t *testing.T) {
    rms := RawMessageStruct{
        X: json.RawMessage("123456"),
    }
    ret, err := Encode(&rms)
    assert.Nil(t, err)
    assert.Equal(t, `{"X":123456}`, string(ret))
}

var _GenericValue interface{}
var _BindingValue TwitterStruct

func init() {
    _ = json.Unmarshal([]byte(TwitterJson), &_GenericValue)
    _ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func TestEncoder_Generic(t *testing.T) {
    v, e := Encode(_GenericValue)
    assert.Nil(t, e)
    println(string(v))
}

func TestEncoder_Binding(t *testing.T) {
    v, e := Encode(_BindingValue)
    assert.Nil(t, e)
    println(string(v))
}

func BenchmarkEncoder_Generic_Sonic(b *testing.B) {
    _, _ = Encode(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(_GenericValue)
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
    _, _ = Encode(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Encode(&_BindingValue)
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
    _, _ = Encode(_GenericValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(_GenericValue)
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
    _, _ = Encode(&_BindingValue)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _ = Encode(&_BindingValue)
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
