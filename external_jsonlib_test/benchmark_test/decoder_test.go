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
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

func init() {
    _ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
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

func BenchmarkDecoder_Generic_JsonIter(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = jsoniter.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _ = jsoniter.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Generic_GoJson(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = gojson.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _ = gojson.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Sonic_GoJson(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = sonic.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v interface{}
        _ = sonic.Unmarshal(m, &v)
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

func BenchmarkDecoder_Binding_JsonIter(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = jsoniter.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _ = jsoniter.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Binding_GoJson(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = gojson.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _ = gojson.Unmarshal(m, &v)
    }
}

func BenchmarkDecoder_Binding_Sonic(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = sonic.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var v TwitterStruct
        _ = sonic.Unmarshal(m, &v)
    }
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

func BenchmarkDecoder_Parallel_Generic_JsonIter(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = jsoniter.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _ = jsoniter.Unmarshal(m, &v)
        }
    })
}

func BenchmarkDecoder_Parallel_Generic_GoJson(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = gojson.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _ = gojson.Unmarshal(m, &v)
        }
    })
}

func BenchmarkDecoder_Parallel_Generic_Sonic(b *testing.B) {
    var w interface{}
    m := []byte(TwitterJson)
    _ = sonic.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v interface{}
            _ = sonic.Unmarshal(m, &v)
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

func BenchmarkDecoder_Parallel_Binding_JsonIter(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = jsoniter.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _ = jsoniter.Unmarshal(m, &v)
        }
    })
}

func BenchmarkDecoder_Parallel_Binding_GoJson(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = gojson.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _ = gojson.Unmarshal(m, &v)
        }
    })
}

func BenchmarkDecoder_Parallel_Binding_Sonic(b *testing.B) {
    var w TwitterStruct
    m := []byte(TwitterJson)
    _ = sonic.Unmarshal(m, &w)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var v TwitterStruct
            _ = sonic.Unmarshal(m, &v)
        }
    })
}
