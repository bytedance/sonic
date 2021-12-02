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
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	go func() {
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
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, size int) {
			defer wg.Done()
			out, err := Encode(_GenericValue, 0)
			if err != nil {
				t.Fatal(err)
			}
			if len(out) != size {
				t.Fatal(len(out), size)
			}
			runtime.GC()
			debug.FreeOSMemory()
		}(wg, n)
	}
	wg.Wait()
}

func runEncoderTest(t *testing.T, fn func(string) string, exp string, arg string) {
	require.Equal(t, exp, fn(arg))
}

func TestEncoder_String(t *testing.T) {
	runEncoderTest(t, Quote, `""`, "")
	runEncoderTest(t, Quote, `"hello, world"`, "hello, world")
	runEncoderTest(t, Quote, `"hello啊啊啊aa"`, "hello啊啊啊aa")
	runEncoderTest(t, Quote, `"hello\\\"world"`, "hello\\\"world")
	runEncoderTest(t, Quote, `"hello\n\tworld"`, "hello\n\tworld")
	runEncoderTest(t, Quote, `"hello\u0000\u0001world"`, "hello\x00\x01world")
	runEncoderTest(t, Quote, `"hello\u0000\u0001world"`, "hello\x00\x01world")
	runEncoderTest(t, Quote, `"Cartoonist, Illustrator, and T-Shirt connoisseur"`, "Cartoonist, Illustrator, and T-Shirt connoisseur")
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
	return []byte(strconv.Itoa(self.X)), nil
}

type MarshalerStruct struct {
	V MarshalerImpl
}

func TestEncoder_Marshaler(t *testing.T) {
	v := MarshalerStruct{V: MarshalerImpl{X: 12345}}
	ret, err := Encode(&v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":12345}`, string(ret))
	ret, err = Encode(v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":{"X":12345}}`, string(ret))
}

type RawMessageStruct struct {
	X json.RawMessage
}

func TestEncoder_RawMessage(t *testing.T) {
	rms := RawMessageStruct{
		X: json.RawMessage("123456"),
	}
	ret, err := Encode(&rms, 0)
	require.NoError(t, err)
	require.Equal(t, `{"X":123456}`, string(ret))
}

var (
	_mediumData         []byte
	_largeData          []byte
	_mediumGenericValue interface{}
	_largeGenericValue  interface{}
	_mediumBindingValue TwitterStruct
	_largeBindingValue  TwitterStruct

	_GenericValue interface{}
	_BindingValue TwitterStruct
)

func init() {
	mediumData := []byte(TwitterJson)
	largeData, err := os.ReadFile(filepath.Join("..", "testdata", "twitter.json"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(mediumData, &_GenericValue); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(mediumData, &_BindingValue); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(mediumData, &_mediumGenericValue); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(mediumData, &_mediumBindingValue); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(largeData, &_largeGenericValue); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(largeData, &_largeBindingValue); err != nil {
		panic(err)
	}
	_mediumData = mediumData
	_largeData = largeData
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
	m := map[string]string{
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

func encodeWithSonic(b *testing.B, data []byte, v interface{}) {
	_, _ = Encode(v, 0)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Encode(v, 0)
	}
}

func encodeParallelWithSonic(b *testing.B, data []byte, v interface{}) {
	_, _ = Encode(v, 0)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Encode(v, 0)
		}
	})
}

func encodeWithSonicSortedMap(b *testing.B, data []byte, v interface{}) {
	_, _ = Encode(v, SortMapKeys)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Encode(v, SortMapKeys)
	}
}

func encodeParallelWithSonicSortedMap(b *testing.B, data []byte, v interface{}) {
	_, _ = Encode(v, SortMapKeys)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Encode(v, SortMapKeys)
		}
	})
}

func encodeWithJsonIter(b *testing.B, data []byte, v interface{}) {
	_, _ = jsoniter.Marshal(v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(v)
	}
}

func encodeParallelWithJsonIter(b *testing.B, data []byte, v interface{}) {
	_, _ = jsoniter.Marshal(v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = jsoniter.Marshal(v)
		}
	})
}

func encodeWithGoJson(b *testing.B, data []byte, v interface{}) {
	_, _ = gojson.Marshal(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gojson.MarshalWithOption(v, gojson.UnorderedMap())
	}
}

func encodeParallelWithGoJson(b *testing.B, data []byte, v interface{}) {
	_, _ = gojson.Marshal(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = gojson.MarshalWithOption(v, gojson.UnorderedMap())
		}
	})
}

func encodeWithStdLib(b *testing.B, data []byte, v interface{}) {
	_, _ = json.Marshal(v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(v)
	}
}

func encodeParallelWithStdLib(b *testing.B, data []byte, v interface{}) {
	_, _ = json.Marshal(v)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = json.Marshal(v)
		}
	})
}

func BenchmarkEncoder_Generic_Sonic(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithSonic(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithSonic(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Generic_SonicSorted(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithSonicSortedMap(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithSonicSortedMap(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Generic_JsonIter(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithJsonIter(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithJsonIter(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Generic_GoJson(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithGoJson(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithGoJson(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Generic_StdLib(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithStdLib(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithStdLib(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Binding_Sonic(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithSonic(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithSonic(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Binding_SonicSorted(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithSonicSortedMap(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithSonicSortedMap(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Binding_JsonIter(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithJsonIter(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithJsonIter(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Binding_GoJson(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithGoJson(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithGoJson(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Binding_StdLib(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeWithStdLib(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeWithStdLib(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Parallel_Generic_Sonic(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithSonic(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithSonic(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Parallel_Generic_SonicSorted(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithSonicSortedMap(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithSonicSortedMap(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Parallel_Generic_JsonIter(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithJsonIter(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithJsonIter(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Parallel_Generic_GoJson(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithGoJson(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithGoJson(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Parallel_Generic_StdLib(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithStdLib(b, _mediumData, &_mediumGenericValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithStdLib(b, _largeData, &_largeGenericValue)
	})
}

func BenchmarkEncoder_Parallel_Binding_Sonic(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithSonic(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithSonic(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Parallel_Binding_SonicSorted(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithSonicSortedMap(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithSonicSortedMap(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Parallel_Binding_JsonIter(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithJsonIter(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithJsonIter(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Parallel_Binding_GoJson(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithGoJson(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithGoJson(b, _largeData, &_largeBindingValue)
	})
}

func BenchmarkEncoder_Parallel_Binding_StdLib(b *testing.B) {
	b.Run("Medium", func(b *testing.B) {
		encodeParallelWithStdLib(b, _mediumData, &_mediumBindingValue)
	})
	b.Run("Large", func(b *testing.B) {
		encodeParallelWithStdLib(b, _largeData, &_largeBindingValue)
	})
}
