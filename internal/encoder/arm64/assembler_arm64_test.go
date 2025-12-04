/*
 * Copyright 2025 Huawei Technologies Co., Ltd.
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

package arm64_test

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"

	"fmt"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/internal/encoder"
	"github.com/bytedance/sonic/internal/encoder/arm64"
	"github.com/bytedance/sonic/internal/encoder/ir"
	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	encoder.ForceUseJit()
	m.Run()
}

type testOps struct {
	key string
	ins ir.Program
	exp string
	err error
	val interface{}
}

func testOpCode(t *testing.T, v interface{}, ex string, err error, ins ir.Program) {
	p := ins
	m := []byte(nil)
	s := new(vars.Stack)
	a := arm64.NewAssembler(p)
	f := a.Load()
	e := f(&m, rt.UnpackEface(v).Value, s, 0)
	if err != nil {
		assert.EqualError(t, e, err.Error())
	} else {
		assert.Nil(t, e)
		assert.Equal(t, ex, string(m))
	}
}

type IfaceValue int

func (IfaceValue) Error() string {
	return "not really implemented"
}

func mustCompile(t interface{}) ir.Program {
	p, err := encoder.NewCompiler().Compile(reflect.TypeOf(t), !rt.UnpackEface(t).Type.Indirect())
	if err != nil {
		panic(err)
	}
	return p
}

type JsonMarshalerValue int

func (JsonMarshalerValue) MarshalJSON() ([]byte, error) {
	return []byte("123456789"), nil
}

type RecursiveValue struct {
	A int                       `json:"a"`
	P *RecursiveValue           `json:"p,omitempty"`
	Q []RecursiveValue          `json:"q"`
	R map[string]RecursiveValue `json:"r"`
	Z int                       `json:"z"`
}

func TestAssembler_OpCode(t *testing.T) {
	var iface error = IfaceValue(12345)
	var eface interface{} = 12345
	var jval = new(JsonMarshalerValue)
	var jifv json.Marshaler = JsonMarshalerValue(0)
	var jifp json.Marshaler = jval
	var rec = &RecursiveValue{
		A: 123,
		Z: 456,
		P: &RecursiveValue{
			A: 789,
			Z: 666,
			P: &RecursiveValue{
				A: 777,
				Z: 888,
				Q: []RecursiveValue{{
					A: 999,
					Z: 222,
					R: map[string]RecursiveValue{
						"xxx": {
							A: 333,
						},
					},
				}},
			},
		},
	}
	tests := []testOps{
		{
			key: "_OP_null",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_null)},
			exp: "null",
			val: nil,
		}, {
			key: "_OP_i16",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i16)},
			exp: "-32768",
			val: int16(-32768),
		},
		{
			key: "_OP_i32",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i32)},
			exp: "-2147483648",
			val: int32(-2147483648),
		},
		{
			key: "_OP_i64",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i64)},
			exp: "-9223372036854775808",
			val: int64(math.MinInt64),
		},
		{
			key: "_OP_u8",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u8)},
			exp: "255",
			val: uint8(255),
		},
		{
			key: "_OP_u16",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u16)},
			exp: "65535",
			val: uint16(65535),
		},
		{
			key: "_OP_u32",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u32)},
			exp: "4294967295",
			val: uint32(4294967295),
		},
		{
			key: "_OP_str",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_str)},
			exp: `"Cartoonist, Illustrator, and T-Shirt connoisseur"`,
			val: "Cartoonist, Illustrator, and T-Shirt connoisseur",
		},
		{
			key: "_OP_str/empty",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_str)},
			exp: `""`,
			val: "",
		},
		{
			key: "_OP_quote",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_quote)},
			exp: `"\"test\""`,
			val: "test",
		},
		{
			key: "_OP_quote/escape",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_quote)},
			exp: `"\"hello\\n\\t\\rworld\""`,
			val: "hello\n\t\rworld",
		},
		{
			key: "_OP_number",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_number)},
			exp: "1.2345",
			val: "1.2345",
		}, {
			key: "_OP_number/invalid",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_number)},
			err: vars.Error_number("not a number"),
			val: "not a number",
		}, {
			key: "_OP_f64",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f64)},
			exp: "-2.2250738585072014e-308",
			val: -2.2250738585072014e-308,
		}, {
			key: "_OP_f64/nan",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f64)},
			err: vars.ERR_nan_or_infinite,
			val: math.NaN(),
		}, {
			key: "_OP_f64/+inf",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f64)},
			err: vars.ERR_nan_or_infinite,
			val: math.Inf(1),
		}, {
			key: "_OP_f64/-inf",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f64)},
			err: vars.ERR_nan_or_infinite,
			val: math.Inf(-1),
		},
		{
			key: "_OP_f32",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f32)},
			exp: "-12.5",
			val: float32(-12.5),
		}, {
			key: "_OP_f32/nan",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f32)},
			err: vars.ERR_nan_or_infinite,
			val: float32(math.NaN()),
		}, {
			key: "_OP_f32/+inf",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f32)},
			err: vars.ERR_nan_or_infinite,
			val: float32(math.Inf(1)),
		}, {
			key: "_OP_f32/-inf",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_f32)},
			err: vars.ERR_nan_or_infinite,
			val: float32(math.Inf(-1)),
		},
		{
			key: "_OP_empty_obj",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_empty_obj)},
			exp: "null",
			val: nil,
		},
		{
			key: "_OP_i8",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i8)},
			exp: "-128",
			val: int8(-128),
		},
		{
			key: "_OP_u64",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u64)},
			exp: "18446744073709551615",
			val: uint64(18446744073709551615),
		},
		{
			key: "_OP_str",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_str)},
			exp: "\"hello world\"",
			val: "hello world",
		},
		{
			key: "_OP_byte",
			ins: []ir.Instr{ir.NewInsVi(ir.OP_byte, 'x')},
			exp: "x",
			val: nil,
		},
		{
			key: "_OP_quote",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_quote)},
			exp: "\"\\\"hello world\\\"\"",
			val: "hello world",
		},
		{
			key: "_OP_bool/true",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_bool)},
			exp: "true",
			val: true,
		}, {
			key: "_OP_bool/false",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_bool)},
			exp: "false",
			val: false,
		}, {
			key: "_OP_eface",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_eface)},
			exp: `12345`,
			val: &eface,
		}, {
			key: "_OP_iface",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_iface)},
			exp: `12345`,
			val: &iface,
		}, {
			key: "_OP_map_[iter,next,value]",
			ins: mustCompile(map[string]map[int64]int{}),
			exp: `{"asdf":{"-9223372036854775808":1234}}`,
			val: &map[string]map[int64]int{"asdf": {math.MinInt64: 1234}},
		}, {
			key: "_OP_marshal[_text]",
			ins: []ir.Instr{ir.NewInsVt(ir.OP_marshal, reflect.TypeOf(JsonMarshalerValue(0)))},
			exp: "123456789",
			val: new(JsonMarshalerValue),
		}, {
			key: "_OP_marshal[_text]/ptr",
			ins: []ir.Instr{ir.NewInsVt(ir.OP_marshal, reflect.TypeOf(new(JsonMarshalerValue)))},
			exp: "123456789",
			val: &jval,
		}, {
			key: "_OP_marshal[_text]/iface_v",
			ins: []ir.Instr{ir.NewInsVt(ir.OP_marshal, vars.JsonMarshalerType)},
			exp: "123456789",
			val: &jifv,
		}, {
			key: "_OP_marshal[_text]/iface_p",
			ins: []ir.Instr{ir.NewInsVt(ir.OP_marshal, vars.JsonMarshalerType)},
			exp: "123456789",
			val: &jifp,
		}, {
			key: "_OP_slice_[len,next]",
			ins: mustCompile([][]int{}),
			exp: `[[1,2,3],[4,5,6]]`,
			val: &[][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			key: "_OP_bin",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_bin)},
			exp: `"AQIDBAU="`,
			val: []byte{1, 2, 3, 4, 5},
		}, {
			key: "_OP_bin/empty",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_bin)},
			exp: `""`,
			val: []byte{},
		}, {
			key: "_OP_recurse",
			ins: mustCompile(rec),
			exp: `{"a":123,"p":{"a":789,"p":{"a":777,"q":[{"a":999,"q":null,"r":{"` +
				`xxx":{"a":333,"q":null,"r":null,"z":0}},"z":222}],"r":null,"z":8` +
				`88},"q":null,"r":null,"z":666},"q":null,"r":null,"z":456}`,
			val: &rec,
		},
	}
	for _, tv := range tests {
		t.Run(tv.key, func(t *testing.T) {
			fmt.Printf("Execute Test: \"%s\"\n", tv.key)
			testOpCode(t, tv.val, tv.exp, tv.err, tv.ins)
		})
	}
}

func TestAssembler_unsupported(t *testing.T) {
	v := complex(1.0, 2.0)
	_, err := sonic.Marshal(v)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "unsupported type")
}

type HugeStruct_int8 struct {
	Field0 int8 `json:"filed0,omitempty"`
}

func TestAssembler_i8(t *testing.T) {
	v := HugeStruct_int8{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_int8
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_int16 struct {
	Field0 int16 `json:"filed0,omitempty"`
}

func TestAssembler_i16(t *testing.T) {
	v := HugeStruct_int16{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_int16
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_int32 struct {
	Field0 int32 `json:"filed0,omitempty"`
}

func TestAssembler_i32(t *testing.T) {
	v := HugeStruct_int32{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_int32
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_int64 struct {
	Field0 int64 `json:"filed0,omitempty"`
}

func TestAssembler_i64(t *testing.T) {
	v := HugeStruct_int64{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_int64
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_uint8 struct {
	Field0 uint8 `json:"filed0,omitempty"`
}

func TestAssembler_u8(t *testing.T) {
	v := HugeStruct_uint8{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_uint8
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_uint16 struct {
	Field0 uint16 `json:"filed0,omitempty"`
}

func TestAssembler_u16(t *testing.T) {
	v := HugeStruct_uint16{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_uint16
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_uint32 struct {
	Field0 uint32 `json:"filed0,omitempty"`
}

func TestAssembler_u32(t *testing.T) {
	v := HugeStruct_uint32{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_uint32
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_uint64 struct {
	Field0 uint64 `json:"filed0,omitempty"`
}

func TestAssembler_u64(t *testing.T) {
	v := HugeStruct_uint64{
		Field0: 1,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_uint64
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_float32 struct {
	Field0 float32 `json:"filed0,omitempty"`
}

func TestAssembler_f32(t *testing.T) {
	v := HugeStruct_float32{
		Field0: 1.5,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_float32
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}
func TestAssembler_f32_maximum(t *testing.T) {
	v := HugeStruct_float32{
		Field0: math.MaxFloat32,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_float32
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.NotNil(t, err)
	require.ErrorContains(t, err, "Mismatch type float32 with value number")
}

func TestAssembler_f32_minimum(t *testing.T) {
	v := HugeStruct_float32{
		Field0: -math.MaxFloat32,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_float32
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.NotNil(t, err)
	require.ErrorContains(t, err, "Mismatch type float32 with value number")
}

type HugeStruct_float64 struct {
	Field0 float64 `json:"filed0,omitempty"`
}

func TestAssembler_f64(t *testing.T) {
	v := HugeStruct_float64{
		Field0: 1.5,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_float64
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_string struct {
	Field0 string `json:"filed0,omitempty"`
}

func TestAssembler_str(t *testing.T) {
	v := HugeStruct_string{
		Field0: "test",
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_string
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_quote struct {
	Field0 string `json:"filed0,omitempty,string"`
}

func TestAssembler_quote(t *testing.T) {
	v := HugeStruct_quote{
		Field0: "test",
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_quote
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_byte struct {
	Field0 byte `json:"filed0,omitempty"`
}

func TestAssembler_byte(t *testing.T) {
	v := HugeStruct_byte{
		Field0: 65, // 'A'
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_byte
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_bool struct {
	Field0 bool `json:"filed0,omitempty"`
	Field1 bool `json:"filed1,omitempty"`
}

func TestAssembler_bool(t *testing.T) {
	v := HugeStruct_bool{
		Field0: true,
		Field1: false,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_bool
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct_star struct {
	Field0 *int64 `json:"filed1,omitempty"`
}

func TestAssembler_star(t *testing.T) {
	var num int64 = 10
	v := HugeStruct_star{
		Field0: &num,
	}

	s, err := sonic.Marshal(v)

	var v2 HugeStruct_star
	if err == nil {
		err = sonic.Unmarshal(s, &v2)
	}
	require.Equal(t, v.Field0, v2.Field0)
}

type HugeStruct struct {
	FieldInt8    int8    `json:"field_int8,omitempty"`
	FieldInt16   int16   `json:"field_int16,omitempty"`
	FieldInt32   int32   `json:"field_int32,omitempty"`
	FieldInt64   int64   `json:"field_int64,omitempty"`
	FieldUint8   uint8   `json:"field_uint8,omitempty"`
	FieldUint16  uint16  `json:"field_uint16,omitempty"`
	FieldUint32  uint32  `json:"field_uint32,omitempty"`
	FieldUint64  uint64  `json:"field_uint64,omitempty"`
	FieldFloat64 float64 `json:"field_float64,omitempty"`
	FieldString  string  `json:"field_string,omitempty"`
	FieldQuote   string  `json:"field_quote,omitempty,string"`
	FieldByte    byte    `json:"field_byte,omitempty"`
	FieldBool    bool    `json:"field_bool,omitempty"`
	FieldIntPtr  *int    `json:"field_int_ptr,omitempty"`
}

func TestAssembler_All_EdgeCases(t *testing.T) {
	// 测试边界值和零值
	zero := 0
	negative := -1
	maxInt := 2147483647

	tests := []struct {
		name string
		v    HugeStruct
	}{
		{
			name: "max_values",
			v: HugeStruct{
				FieldInt8:    127,
				FieldInt16:   32767,
				FieldInt32:   2147483647,
				FieldInt64:   9223372036854775807,
				FieldUint8:   255,
				FieldUint16:  65535,
				FieldUint32:  4294967295,
				FieldUint64:  18446744073709551615,
				FieldFloat64: math.MaxFloat64,
				FieldString:  "Maximum",
				FieldQuote:   "Maximum",
				FieldByte:    255,
				FieldBool:    true,
				FieldIntPtr:  &maxInt,
			},
		},
		{
			name: "min_values",
			v: HugeStruct{
				FieldInt8:    -128,
				FieldInt16:   -32768,
				FieldInt32:   -2147483648,
				FieldInt64:   -9223372036854775808,
				FieldUint8:   0,
				FieldUint16:  0,
				FieldUint32:  0,
				FieldUint64:  0,
				FieldFloat64: -math.MaxFloat64,
				FieldString:  "Minimum",
				FieldQuote:   "Minimum",
				FieldByte:    0,
				FieldBool:    false,
				FieldIntPtr:  &negative,
			},
		},
		{
			name: "zero_values",
			v: HugeStruct{
				FieldInt8:    0,
				FieldInt16:   0,
				FieldInt32:   0,
				FieldInt64:   0,
				FieldUint8:   0,
				FieldUint16:  0,
				FieldUint32:  0,
				FieldUint64:  0,
				FieldFloat64: 0,
				FieldString:  "",
				FieldQuote:   "",
				FieldByte:    0,
				FieldBool:    false,
				FieldIntPtr:  &zero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := sonic.Marshal(tt.v)
			require.NoError(t, err)

			var v2 HugeStruct
			err = sonic.Unmarshal(s, &v2)
			require.NoError(t, err)

			require.Equal(t, tt.v.FieldInt8, v2.FieldInt8)
			require.Equal(t, tt.v.FieldInt16, v2.FieldInt16)
			require.Equal(t, tt.v.FieldInt32, v2.FieldInt32)
			require.Equal(t, tt.v.FieldInt64, v2.FieldInt64)
			require.Equal(t, tt.v.FieldUint8, v2.FieldUint8)
			require.Equal(t, tt.v.FieldUint16, v2.FieldUint16)
			require.Equal(t, tt.v.FieldUint32, v2.FieldUint32)
			require.Equal(t, tt.v.FieldUint64, v2.FieldUint64)
			require.Equal(t, tt.v.FieldFloat64, v2.FieldFloat64)
			require.Equal(t, tt.v.FieldString, v2.FieldString)
			require.Equal(t, tt.v.FieldQuote, v2.FieldQuote)
			require.Equal(t, tt.v.FieldByte, v2.FieldByte)
			require.Equal(t, tt.v.FieldBool, v2.FieldBool)
			require.Equal(t, *tt.v.FieldIntPtr, *v2.FieldIntPtr)
		})
	}
}
