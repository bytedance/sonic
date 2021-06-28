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
	"testing"

	"github.com/davecgh/go-spew/spew"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _BindingValue TwitterStruct

func init() {
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func decode(s string, v interface{}) (int, error) {
	d := NewDecoder(s)
	err := d.Decode(v)
	if err != nil {
		return 0, err
	}
	return d.i, err
}

func TestDecoder_Basic(t *testing.T) {
	var v int
	pos, err := decode("12345", &v)
	assert.NoError(t, err)
	assert.Equal(t, 5, pos)
	assert.Equal(t, 12345, v)
}

func TestDecoder_Generic(t *testing.T) {
	var v interface{}
	pos, err := decode(TwitterJson, &v)
	assert.NoError(t, err)
	assert.Equal(t, len(TwitterJson), pos)
	spew.Dump(v)
}

func TestDecoder_Binding(t *testing.T) {
	var v TwitterStruct
	pos, err := decode(TwitterJson, &v)
	assert.NoError(t, err)
	assert.Equal(t, len(TwitterJson), pos)
	assert.Equal(t, _BindingValue, v, 0)
	spew.Dump(v)
}

func TestDecoder_MapWithIndirectElement(t *testing.T) {
	var v map[string]struct{ A [129]byte }
	_, err := decode(`{"":{"A":[1,2,3,4,5]}}`, &v)
	if x, ok := err.(SyntaxError); ok {
		println(x.Description())
	}
	require.NoError(t, err)
	assert.Equal(t, [129]byte{1, 2, 3, 4, 5}, v[""].A)
}

func BenchmarkDecoder_Generic_Sonic(b *testing.B) {
	var w interface{}
	_, _ = decode(TwitterJson, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_, _ = decode(TwitterJson, &v)
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

func BenchmarkDecoder_Binding_Sonic(b *testing.B) {
	var w TwitterStruct
	_, _ = decode(TwitterJson, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_, _ = decode(TwitterJson, &v)
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

func BenchmarkDecoder_Parallel_Generic_Sonic(b *testing.B) {
	var w interface{}
	_, _ = decode(TwitterJson, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_, _ = decode(TwitterJson, &v)
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

func BenchmarkDecoder_Parallel_Binding_Sonic(b *testing.B) {
	var w TwitterStruct
	_, _ = decode(TwitterJson, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_, _ = decode(TwitterJson, &v)
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
