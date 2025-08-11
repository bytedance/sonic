/**
 * Copyright 2025 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package issue_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

type ITM int
type ITM8 int8
type ITM16 int16
type ITM32 int32
type ITM64 int64
type UTM uint
type UTM8 uint8
type UTM16 uint16
type UTM32 uint32
type UTM64 uint64
type FTM32 float32
type FTM64 float64
type StrTM string
type Str2TM string
type CustomTM chan int
type STM struct{ ID int }
type PTM struct{ ID int }

func (i ITM) MarshalText() ([]byte, error) {
	return []byte("u" + strconv.Itoa(int(i))), nil
}

func (i ITM8) MarshalText() ([]byte, error) {
	return []byte("u8" + strconv.FormatInt(int64(i), 10)), nil
}

func (i ITM16) MarshalText() ([]byte, error) {
	return []byte("u16" + strconv.FormatInt(int64(i), 10)), nil
}

func (i ITM32) MarshalText() ([]byte, error) {
	return []byte("u32" + strconv.FormatInt(int64(i), 10)), nil
}

func (i ITM64) MarshalText() ([]byte, error) {
	return []byte("u64" + strconv.FormatInt(int64(i), 10)), nil
}

func (i UTM) MarshalText() ([]byte, error) {
	return []byte("u" + strconv.FormatUint(uint64(i), 10)), nil
}

func (i UTM8) MarshalText() ([]byte, error) {
	return []byte("u8" + strconv.FormatUint(uint64(i), 10)), nil
}

func (i UTM16) MarshalText() ([]byte, error) {
	return []byte("u16" + strconv.FormatUint(uint64(i), 10)), nil
}

func (i UTM32) MarshalText() ([]byte, error) {
	return []byte("u32" + strconv.FormatUint(uint64(i), 10)), nil
}

func (i UTM64) MarshalText() ([]byte, error) {
	return []byte("u64" + strconv.FormatUint(uint64(i), 10)), nil
}

func (f FTM32) MarshalText() ([]byte, error) {
	return []byte("f32" + strconv.FormatFloat(float64(f), 'f', -1, 32)), nil
}

func (f FTM64) MarshalText() ([]byte, error) {
	return []byte("f64" + strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}

func (s STM) MarshalText() ([]byte, error) {
	return []byte("struct" + strconv.Itoa(s.ID)), nil
}

func (p *PTM) MarshalText() ([]byte, error) {
	return []byte("ptr" + strconv.Itoa(p.ID)), nil
}

func (c CustomTM) MarshalText() ([]byte, error) {
	if c == nil {
		return []byte("customnil"), nil
	}
	return []byte("custom"), nil
}

func (s *StrTM) MarshalText() ([]byte, error) {
	return []byte("str" + string(*s)), nil
}

func (s Str2TM) MarshalText() ([]byte, error) {
	return []byte("str" + string(s)), nil
}

func TestIssue827_AllTextMarshalerTypes(t *testing.T) {
	p1 := PTM{ID: 1}
	p2 := PTM{ID: 2}

	s1 := Str2TM("one")
	s2 := Str2TM("two")

	testCases := []struct {
		name string
		data interface{}
	}{
		{"int", map[ITM]int{1: 1, 2: 2}},
		{"int8", map[ITM8]int{1: 1, 2: 2}},
		{"int16", map[ITM16]int{1: 1, 2: 2}},
		{"int32", map[ITM32]int{1: 1, 2: 2}},
		{"int64", map[ITM64]int{1: 1, 2: 2}},
		{"uint", map[UTM]int{1: 1, 2: 2}},
		{"uint8", map[UTM8]int{1: 1, 2: 2}},
		{"uint16", map[UTM16]int{1: 1, 2: 2}},
		{"uint32", map[UTM32]int{1: 1, 2: 2}},
		{"uint64", map[UTM64]int{1: 1, 2: 2}},
		{"float32", map[FTM32]int{1.5: 1, 2.7: 2}},
		{"float64", map[FTM64]int{1.5: 1, 2.7: 2}},
		{"struct", map[STM]int{{ID: 1}: 1, {ID: 2}: 2}},
		{"pointer", map[*PTM]int{&p1: 1, &p2: 2, nil: 3}},
		{"string", map[StrTM]int{"one": 1, "two": 2}},
		{"string2", map[*Str2TM]int{&s1: 1, &s2: 2, nil: 3}},
		{"string3", map[Str2TM]int{s1: 1, s2: 2}},
		{"custom", map[CustomTM]int{nil: 1, make(chan int): 2}},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.data)
			require.NoError(t, err)
			expected := string(jsonData)
			println(tc.name, "expected:", expected)
			
			sonicData, err := sonic.Marshal(tc.data)
			require.NoError(t, err)
			
			sonicStdData, err := sonic.ConfigStd.Marshal(tc.data)
			require.NoError(t, err)
			require.Equal(t, expected, string(sonicStdData))

			// compare the output of sonic.Marshal
			var sonicMap, jsonMap map[string]int
			require.NoError(t, json.Unmarshal(sonicData, &sonicMap))
			require.NoError(t, json.Unmarshal(jsonData, &jsonMap))
			require.Equal(t, jsonMap, sonicMap)
			
		})
	}
}
