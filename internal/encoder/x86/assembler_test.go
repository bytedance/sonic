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

package x86_test

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"reflect"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/encoder"
	"github.com/bytedance/sonic/internal/encoder/ir"
	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/bytedance/sonic/internal/encoder/x86"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	encoder.ForceUseJit()
	m.Run()
}

func TestAssembler_CompileAndLoad(t *testing.T) {
	p, err := encoder.NewCompiler().Compile(reflect.TypeOf((*bool)(nil)), true)
	assert.Nil(t, err)
	a := x86.NewAssembler(p)
	f := a.Load()
	s := vars.NewStack()
	b := []byte(nil)

	/* true */
	v := true
	u := &v
	e := f(&b, unsafe.Pointer(&u), s, 0)
	assert.Nil(t, e)
	println(cap(b))
	println(hex.Dump(b))

	/* false */
	v = false
	u = &v
	b = b[:0]
	e = f(&b, unsafe.Pointer(&u), s, 0)
	assert.Nil(t, e)
	println(cap(b))
	println(hex.Dump(b))

	/* nil */
	u = nil
	b = b[:0]
	e = f(&b, unsafe.Pointer(&u), s, 0)
	assert.Nil(t, e)
	println(cap(b))
	println(hex.Dump(b))
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
	a := x86.NewAssembler(p)
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

func mustCompile(t interface{}) ir.Program {
	p, err := encoder.NewCompiler().Compile(reflect.TypeOf(t), !rt.UnpackEface(t).Type.Indirect())
	if err != nil {
		panic(err)
	}
	return p
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
			key: "_OP_i8",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i8)},
			exp: "-128",
			val: int8(-128),
		}, {
			key: "_OP_i16",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i16)},
			exp: "-32768",
			val: int16(-32768),
		}, {
			key: "_OP_i32",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i32)},
			exp: "-2147483648",
			val: int32(-2147483648),
		}, {
			key: "_OP_i64",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_i64)},
			exp: "-9223372036854775808",
			val: int64(math.MinInt64),
		}, {
			key: "_OP_u8",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u8)},
			exp: "255",
			val: uint8(255),
		}, {
			key: "_OP_u16",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u16)},
			exp: "65535",
			val: uint16(65535),
		}, {
			key: "_OP_u32",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u32)},
			exp: "4294967295",
			val: uint32(4294967295),
		}, {
			key: "_OP_u64",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_u64)},
			exp: "18446744073709551615",
			val: uint64(18446744073709551615),
		}, {
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
		}, {
			key: "_OP_str",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_str)},
			exp: `"Cartoonist, Illustrator, and T-Shirt connoisseur"`,
			val: "Cartoonist, Illustrator, and T-Shirt connoisseur",
		}, {
			key: "_OP_str/empty",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_str)},
			exp: `""`,
			val: "",
		}, {
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
			key: "_OP_quote",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_quote)},
			exp: `"\"test\""`,
			val: "test",
		}, {
			key: "_OP_quote/escape",
			ins: []ir.Instr{ir.NewInsOp(ir.OP_quote)},
			exp: `"\"hello\\n\\t\\rworld\""`,
			val: "hello\n\t\rworld",
		}, {
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
			key: "_OP_byte",
			ins: []ir.Instr{ir.NewInsVi(ir.OP_byte, 'x')},
			exp: "x",
			val: nil,
		}, {
			key: "_OP_text",
			ins: []ir.Instr{ir.NewInsVs(ir.OP_text, "hello, world !!")},
			exp: "hello, world !!",
			val: nil,
		}, {
			key: "_OP_map_[iter,next,value]",
			ins: mustCompile(map[string]map[int64]int{}),
			exp: `{"asdf":{"-9223372036854775808":1234}}`,
			val: &map[string]map[int64]int{"asdf": {math.MinInt64: 1234}},
		}, {
			key: "_OP_slice_[len,next]",
			ins: mustCompile([][]int{}),
			exp: `[[1,2,3],[4,5,6]]`,
			val: &[][]int{{1, 2, 3}, {4, 5, 6}},
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
		},
		{
			key: "_OP_recurse",
			ins: mustCompile(rec),
			exp: `{"a":123,"p":{"a":789,"p":{"a":777,"q":[{"a":999,"q":null,"r":{"` +
				`xxx":{"a":333,"q":null,"r":null,"z":0}},"z":222}],"r":null,"z":8` +
				`88},"q":null,"r":null,"z":666},"q":null,"r":null,"z":456}`,
			val: &rec,
		}}
	for _, tv := range tests {
		t.Run(tv.key, func(t *testing.T) {
			testOpCode(t, tv.val, tv.exp, tv.err, tv.ins)
		})
	}
}

func TestAssembler_StringMoreSpace(t *testing.T) {
	p := ir.Program{ir.NewInsOp(ir.OP_str)}
	m := make([]byte, 0, 8)
	s := new(vars.Stack)
	a := x86.NewAssembler(p)
	f := a.Load()
	v := "\u0001\u0002\u0003\u0004\u0005\u0006\u0007\u0008\u0009\u000a\u000b\u000c\u000d\u000e\u000f\u0010"
	e := f(&m, unsafe.Pointer(&v), s, 0)
	assert.Nil(t, e)
	spew.Dump(m)
}
