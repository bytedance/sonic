/*
 * Copyright 2025 ByteDance Inc.
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
	"github.com/stretchr/testify/require"
	msgpack "github.com/vmihailenco/msgpack/v5"
)

var (
	twitterJSONBytes       = []byte(TwitterJson)
	twitterStructForBench  TwitterStruct
	twitterGenericForBench interface{}
	twitterMsgpackBinding  []byte
	twitterMsgpackGeneric  []byte
)

func init() {
	if err := json.Unmarshal(twitterJSONBytes, &twitterStructForBench); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(twitterJSONBytes, &twitterGenericForBench); err != nil {
		panic(err)
	}
	var err error
	twitterMsgpackBinding, err = msgpack.Marshal(&twitterStructForBench)
	if err != nil {
		panic(err)
	}
	twitterMsgpackGeneric, err = msgpack.Marshal(twitterGenericForBench)
	if err != nil {
		panic(err)
	}
}

func BenchmarkWithMsgPackDecode_Generic_Sonic_JSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		if err := sonic.Unmarshal(twitterJSONBytes, &v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackDecode_Generic_Msgpack(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		if err := msgpack.Unmarshal(twitterMsgpackGeneric, &v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackDecode_Binding_Sonic_JSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		if err := sonic.Unmarshal(twitterJSONBytes, &v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackDecode_Binding_Msgpack(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		if err := msgpack.Unmarshal(twitterMsgpackBinding, &v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackEncode_Generic_Sonic_JSON(b *testing.B) {
	if _, err := sonic.Marshal(twitterGenericForBench); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sonic.Marshal(twitterGenericForBench); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackEncode_Generic_Msgpack(b *testing.B) {
	if _, err := msgpack.Marshal(twitterGenericForBench); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := msgpack.Marshal(twitterGenericForBench); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackEncode_Binding_Sonic_JSON(b *testing.B) {
	if _, err := sonic.Marshal(&twitterStructForBench); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sonic.Marshal(&twitterStructForBench); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWithMsgPackEncode_Binding_Msgpack(b *testing.B) {
	if _, err := msgpack.Marshal(&twitterStructForBench); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := msgpack.Marshal(&twitterStructForBench); err != nil {
			b.Fatal(err)
		}
	}
}

func TestMsgpackDecodeMatchesSonicGeneric(t *testing.T) {
	var sonicGeneric interface{}
	require.NoError(t, sonic.Unmarshal(twitterJSONBytes, &sonicGeneric))

	var msgpackGeneric interface{}
	require.NoError(t, msgpack.Unmarshal(twitterMsgpackGeneric, &msgpackGeneric))

	require.Equal(t, sonicGeneric, msgpackGeneric)
}

func TestMsgpackDecodeMatchesSonicBinding(t *testing.T) {
	var sonicStruct TwitterStruct
	require.NoError(t, sonic.Unmarshal(twitterJSONBytes, &sonicStruct))

	var msgpackStruct TwitterStruct
	require.NoError(t, msgpack.Unmarshal(twitterMsgpackBinding, &msgpackStruct))

	require.Equal(t, sonicStruct, msgpackStruct)
}

func TestMsgpackEncodeMatchesSonicGeneric(t *testing.T) {
	sonicBytes, err := sonic.Marshal(twitterGenericForBench)
	require.NoError(t, err)

	var sonicDecoded interface{}
	require.NoError(t, json.Unmarshal(sonicBytes, &sonicDecoded))

	msgpackBytes, err := msgpack.Marshal(twitterGenericForBench)
	require.NoError(t, err)

	var msgpackDecoded interface{}
	require.NoError(t, msgpack.Unmarshal(msgpackBytes, &msgpackDecoded))

	require.Equal(t, sonicDecoded, msgpackDecoded)
}

func TestMsgpackEncodeMatchesSonicBinding(t *testing.T) {
	sonicBytes, err := sonic.Marshal(&twitterStructForBench)
	require.NoError(t, err)

	var sonicDecoded TwitterStruct
	require.NoError(t, sonic.Unmarshal(sonicBytes, &sonicDecoded))

	msgpackBytes, err := msgpack.Marshal(&twitterStructForBench)
	require.NoError(t, err)

	var msgpackDecoded TwitterStruct
	require.NoError(t, msgpack.Unmarshal(msgpackBytes, &msgpackDecoded))

	require.Equal(t, sonicDecoded, msgpackDecoded)
}
