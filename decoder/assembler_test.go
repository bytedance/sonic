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
    `encoding/base64`
    `encoding/json`
    `reflect`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/internal/caching`
    `github.com/bytedance/sonic/internal/jit`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

func TestAssembler_PrologueAndEpilogue(t *testing.T) {
    a := newAssembler(nil)
    _, e := a.Load()("", 0, nil, nil, 0, "", nil)
    assert.Nil(t, e)
}

var utextVar []byte
type UtextValue int

func (UtextValue) UnmarshalText(text []byte) error {
    utextVar = text
    return nil
}

var ujsonVar []byte
type UjsonValue int

func (UjsonValue) UnmarshalJSON(json []byte) error {
    ujsonVar = json
    return nil
}

type UtextStruct struct {
    V string
}

func (self *UtextStruct) UnmarshalText(text []byte) error {
    self.V = string(text)
    return nil
}

type UjsonStruct struct {
    V string
}

func (self *UjsonStruct) UnmarshalJSON(v []byte) error {
    self.V = string(v)
    return nil
}

const (
    _OP_dbg_get_sr _Op = 253
    _OP_dbg_set_sr _Op = 254
    _OP_dbg_break  _Op = 255
)

func (self *_Assembler) _asm_OP_dbg_get_sr(_ *_Instr) {
    self.Emit("MOVQ", _VAR_sr, _AX)
    self.Emit("MOVQ", _AX, jit.Ptr(_VP, 0))
}

func (self *_Assembler) _asm_OP_dbg_set_sr(p *_Instr) {
    self.Emit("MOVQ", jit.Imm(p.i64()), _AX)
    self.Emit("MOVQ", _AX, _VAR_sr)
}

func (self *_Assembler) _asm_OP_dbg_break(_ *_Instr) {
    self.Byte(0xcc)
}

func init() {
    _OpNames[_OP_dbg_get_sr] = "dbg_get_sr"
    _OpNames[_OP_dbg_set_sr] = "dbg_set_sr"
    _OpNames[_OP_dbg_break]  = "dbg_break"
    _OpFuncTab[_OP_dbg_get_sr] = (*_Assembler)._asm_OP_dbg_get_sr
    _OpFuncTab[_OP_dbg_set_sr] = (*_Assembler)._asm_OP_dbg_set_sr
    _OpFuncTab[_OP_dbg_break]  = (*_Assembler)._asm_OP_dbg_break
}

type testOps struct {
    key string
    ins _Program
    src string
    pos int
    opt uint64
    vfn func(i int, v interface{})
    exp interface{}
    err error
    val interface{}
}

func testOpCode(t *testing.T, ops *testOps) {
    p := ops.ins
    k := new(_Stack)
    a := newAssembler(p)
    f := a.Load()
    i, e := f(ops.src, ops.pos, rt.UnpackEface(ops.val).Value, k, ops.opt, "", nil)
    if ops.err != nil {
        assert.EqualError(t, e, ops.err.Error())
    } else {
        assert.NoError(t, e)
        if ops.vfn != nil {
            if ops.val == nil {
                ops.vfn(i, nil)
            } else {
                ops.vfn(i, reflect.Indirect(reflect.ValueOf(ops.val)).Interface())
            }
        } else {
            if ops.val == nil {
                assert.Nil(t, ops.exp)
            } else {
                assert.Equal(t, ops.exp, reflect.Indirect(reflect.ValueOf(ops.val)).Interface())
            }
        }
    }
}

func TestAssembler_OpCode(t *testing.T) {
    tests := []testOps{
    {
        key: "_OP_any/stdlib",
        ins: []_Instr{newInsOp(_OP_any)},
        src: `{"a": [1, 2, 3]}`,
        exp: map[string]interface{}{"a": []interface{}{1.0, 2.0, 3.0}},
        val: new(interface{}),
    }, 
    {
        key: "_OP_any/use_int64",
        ins: []_Instr{newInsOp(_OP_any)},
        src: `{"a": [1, 2, 3]}`,
        opt: 1 << _F_use_int64,
        exp: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}},
        val: new(interface{}),
    }, 
    {
        key: "_OP_any/use_number",
        ins: []_Instr{newInsOp(_OP_any)},
        src: `{"a": [1, 2, 3]}`,
        opt: 1 << _F_use_number,
        exp: map[string]interface{}{"a": []interface{}{json.Number("1"), json.Number("2"), json.Number("3")}},
        val: new(interface{}),
    }, 
    {
        key: "_OP_str/plain",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `hello, world"`,
        exp: "hello, world",
        val: new(string),
    }, {
        key: "_OP_str/unquote",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `hello, world \\ \" \/ \b \f \n \r \t \u666f 测试中文 \ud83d\ude00"`,
        exp: "hello, world \\ \" / \b \f \n \r \t 景 测试中文 😀",
        val: new(string),
    }, {
        key: "_OP_str/unquote_unirep",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `hello\ud800world"`,
        exp: "hello\ufffdworld",
        val: new(string),
    }, {
        key: "_OP_str/error_eof",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `12345`,
        err: SyntaxError{Src: `12345`, Pos: 5, Code: types.ERR_EOF},
        val: new(string),
    }, {
        key: "_OP_str/error_invalid_escape",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `12\g345"`,
        err: SyntaxError{Src: `12\g345"`, Pos: 3, Code: types.ERR_INVALID_ESCAPE},
        val: new(string),
    }, {
        key: "_OP_str/error_invalid_unicode",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `hello\ud800world"`,
        opt: 1 << _F_disable_urc,
        err: SyntaxError{Src: `hello\ud800world"`, Pos: 7, Code: types.ERR_INVALID_UNICODE},
        val: new(string),
    }, {
        key: "_OP_str/error_invalid_char",
        ins: []_Instr{newInsOp(_OP_str)},
        src: `12\u1ggg345"`,
        err: SyntaxError{Src: `12\u1ggg345"`, Pos: 5, Code: types.ERR_INVALID_CHAR},
        val: new(string),
    }, {
        key: "_OP_bin",
        ins: []_Instr{newInsOp(_OP_bin)},
        src: `aGVsbG8sIHdvcmxk"`,
        exp: []byte("hello, world"),
        val: new([]byte),
    }, {
        key: "_OP_bin/error_eof",
        ins: []_Instr{newInsOp(_OP_bin)},
        src: `aGVsbG8sIHdvcmxk`,
        err: SyntaxError{Src: `aGVsbG8sIHdvcmxk`, Pos: 16, Code: types.ERR_EOF},
        val: new([]byte),
    }, {
        key: "_OP_bin/error_corrupt_input",
        ins: []_Instr{newInsOp(_OP_bin)},
        src: `aGVsbG8!sIHdvcmxk"`,
        err: base64.CorruptInputError(7),
        val: new([]byte),
    }, {
        key: "_OP_bool/true",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "true",
        exp: true,
        val: new(bool),
    }, 
    {
        key: "_OP_bool/skip",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: `"true"`,
        exp: nil,
        val: new(bool),
        err: &MismatchTypeError{Src: `"true"`, Pos: 0, Type: reflect.TypeOf(true)},
    }, 
    {
        key: "_OP_bool/false",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "false",
        exp: false,
        val: new(bool),
    }, {
        key: "_OP_bool/false_pos",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "false",
        vfn: func(i int, v interface{}) { require.False(t, v.(bool)); assert.Equal(t, 5, i) },
        val: new(bool),
    }, {
        key: "_OP_bool/error_eof_1",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "tru",
        err: SyntaxError{Src: `tru`, Pos: 3, Code: types.ERR_EOF},
        val: new(bool),
    }, {
        key: "_OP_bool/error_eof_2",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "fals",
        err: SyntaxError{Src: `fals`, Pos: 4, Code: types.ERR_EOF},
        val: new(bool),
    }, {
        key: "_OP_bool/error_invalid_char_1",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "falxe",
        err: SyntaxError{Src: `falxe`, Pos: 3, Code: types.ERR_INVALID_CHAR},
        val: new(bool),
    }, {
        key: "_OP_bool/error_invalid_char_2",
        ins: []_Instr{newInsOp(_OP_bool)},
        src: "falsx",
        err: SyntaxError{Src: `falsx`, Pos: 4, Code: types.ERR_INVALID_CHAR},
        val: new(bool),
    },
    {
        key: "_OP_num/positive",
        ins: []_Instr{newInsOp(_OP_num)},
        src: "1.234e5",
        exp: json.Number("1.234e5"),
        val: new(json.Number),
    }, {
        key: "_OP_num/negative",
        ins: []_Instr{newInsOp(_OP_num)},
        src: "-1.234e5",
        exp: json.Number("-1.234e5"),
        val: new(json.Number),
    }, {
        key: "_OP_num/error_eof",
        ins: []_Instr{newInsOp(_OP_num)},
        src: "-",
        err: SyntaxError{Src: `-`, Pos: 1, Code: types.ERR_INVALID_CHAR},
        val: new(json.Number),
    }, {
        key: "_OP_num/error_invalid_char",
        ins: []_Instr{newInsOp(_OP_num)},
        src: "xxx",
        err: SyntaxError{Src: `xxx`, Pos: 1, Code: types.ERR_INVALID_CHAR},
        val: new(json.Number),
    }, {
        key: "_OP_i8",
        ins: []_Instr{newInsOp(_OP_i8)},
        src: "123",
        exp: int8(123),
        val: new(int8),
    }, {
        key: "_OP_i8/error_overflow",
        ins: []_Instr{newInsOp(_OP_i8)},
        src: "1234",
        err: error_value("1234", reflect.TypeOf(int8(0))),
        val: new(int8),
    }, 
    {
        key: "_OP_i8/error_wrong_type",
        ins: []_Instr{newInsOp(_OP_i8)},
        src: "12.34",
        err: &MismatchTypeError{Src: `12.34`, Pos: 0, Type: int8Type},
        val: new(int8),
    }, {
        key: "_OP_u8",
        ins: []_Instr{newInsOp(_OP_u8)},
        src: "234",
        exp: uint8(234),
        val: new(uint8),
    }, {
        key: "_OP_u8/error_overflow",
        ins: []_Instr{newInsOp(_OP_u8)},
        src: "1234",
        err: error_value("1234", reflect.TypeOf(uint8(0))),
        val: new(uint8),
    }, {
        key: "_OP_u8/error_underflow",
        ins: []_Instr{newInsOp(_OP_u8)},
        src: "-123",
        err: &MismatchTypeError{Src: `-123`, Pos: 0, Type: uint8Type},
        val: new(uint8),
    }, {
        key: "_OP_u8/error_wrong_type",
        ins: []_Instr{newInsOp(_OP_u8)},
        src: "12.34",
        err: &MismatchTypeError{Src: `12.34`, Pos: 0, Type: uint8Type},
        val: new(uint8),
    }, {
        key: "_OP_f32",
        ins: []_Instr{newInsOp(_OP_f32)},
        src: "1.25e20",
        exp: float32(1.25e20),
        val: new(float32),
    }, {
        key: "_OP_f32/overflow",
        ins: []_Instr{newInsOp(_OP_f32)},
        src: "1.25e50",
        err: error_value("1.25e50", reflect.TypeOf(float32(0))),
        val: new(float32),
    }, {
        key: "_OP_f32/underflow",
        ins: []_Instr{newInsOp(_OP_f32)},
        src: "-1.25e50",
        err: error_value("-1.25e50", reflect.TypeOf(float32(0))),
        val: new(float32),
    }, {
        key: "_OP_f64",
        ins: []_Instr{newInsOp(_OP_f64)},
        src: "1.25e123",
        exp: 1.25e123,
        val: new(float64),
    }, {
        key: "_OP_unquote/plain",
        ins: []_Instr{newInsOp(_OP_unquote)},
        src: `\"hello, world\""`,
        exp: "hello, world",
        val: new(string),
    }, {
        key: "_OP_unquote/unquote",
        ins: []_Instr{newInsOp(_OP_unquote)},
        src: `\"hello, world \\\\ \\\" \\/ \\b \\f \\n \\r \\t \\u666f 测试中文 \\ud83d\\ude00\""`,
        exp: "hello, world \\ \" / \b \f \n \r \t 景 测试中文 😀",
        val: new(string),
    }, {
        key: "_OP_unquote/error_invalid_end",
        ins: []_Instr{newInsOp(_OP_unquote)},
        src: `\"te\\\"st"`,
        err: SyntaxError{Src: `\"te\\\"st"`, Pos: 8, Code: types.ERR_INVALID_CHAR},
        val: new(string),
    }, {
        key: "_OP_nil_1",
        ins: []_Instr{newInsOp(_OP_nil_1)},
        src: "",
        exp: 0,
        val: (func() *int { v := new(int); *v = 123; return v })(),
    }, {
        key: "_OP_nil_2",
        ins: []_Instr{newInsOp(_OP_nil_2)},
        src: "",
        exp: error(nil),
        val: (func() *error { v := new(error); *v = types.ERR_EOF; return v })(),
    }, {
        key: "_OP_nil_3",
        ins: []_Instr{newInsOp(_OP_nil_3)},
        src: "",
        exp: []byte(nil),
        val: &[]byte{1, 2, 3},
    }, {
        key: "_OP_deref",
        ins: []_Instr{newInsVt(_OP_deref, reflect.TypeOf(0))},
        src: "",
        vfn: func(_ int, v interface{}) { require.NotNil(t, v); assert.NotNil(t, v.(*int)) },
        val: new(*int),
    }, {
        key: "_OP_map_init",
        ins: []_Instr{newInsOp(_OP_map_init)},
        src: "",
        vfn: func(_ int, v interface{}) { require.NotNil(t, v); assert.NotNil(t, v.(map[string]int)) },
        val: new(map[string]int),
    }, {
        key: "_OP_map_key_i8",
        ins: []_Instr{newInsVt(_OP_map_key_i8, reflect.TypeOf(map[int8]int{}))},
        src: `123`,
        exp: map[int8]int{123: 0},
        val: map[int8]int{},
    }, {
        key: "_OP_map_key_i32",
        ins: []_Instr{newInsVt(_OP_map_key_i32, reflect.TypeOf(map[int32]int{}))},
        src: `123456789`,
        exp: map[int32]int{123456789: 0},
        val: map[int32]int{},
    }, {
        key: "_OP_map_key_i64",
        ins: []_Instr{newInsVt(_OP_map_key_i64, reflect.TypeOf(map[int64]int{}))},
        src: `123456789123456789`,
        exp: map[int64]int{123456789123456789: 0},
        val: map[int64]int{},
    }, {
        key: "_OP_map_key_u8",
        ins: []_Instr{newInsVt(_OP_map_key_u8, reflect.TypeOf(map[uint8]int{}))},
        src: `123`,
        exp: map[uint8]int{123: 0},
        val: map[uint8]int{},
    }, {
        key: "_OP_map_key_u32",
        ins: []_Instr{newInsVt(_OP_map_key_u32, reflect.TypeOf(map[uint32]int{}))},
        src: `123456789`,
        exp: map[uint32]int{123456789: 0},
        val: map[uint32]int{},
    }, {
        key: "_OP_map_key_u64",
        ins: []_Instr{newInsVt(_OP_map_key_u64, reflect.TypeOf(map[uint64]int{}))},
        src: `123456789123456789`,
        exp: map[uint64]int{123456789123456789: 0},
        val: map[uint64]int{},
    }, {
        key: "_OP_map_key_f32",
        ins: []_Instr{newInsVt(_OP_map_key_f32, reflect.TypeOf(map[float32]int{}))},
        src: `1.25`,
        exp: map[float32]int{1.25: 0},
        val: map[float32]int{},
    }, {
        key: "_OP_map_key_f64",
        ins: []_Instr{newInsVt(_OP_map_key_f64, reflect.TypeOf(map[float64]int{}))},
        src: `1.25`,
        exp: map[float64]int{1.25: 0},
        val: map[float64]int{},
    }, {
        key: "_OP_map_key_str/plain",
        ins: []_Instr{newInsVt(_OP_map_key_str, reflect.TypeOf(map[string]int{}))},
        src: `foo"`,
        exp: map[string]int{"foo": 0},
        val: map[string]int{},
    }, {
        key: "_OP_map_key_str/unquote",
        ins: []_Instr{newInsVt(_OP_map_key_str, reflect.TypeOf(map[string]int{}))},
        src: `foo\nbar"`,
        exp: map[string]int{"foo\nbar": 0},
        val: map[string]int{},
    }, 
    {
        key: "_OP_map_key_utext/value",
        ins: []_Instr{newInsVt(_OP_map_key_utext, reflect.TypeOf(map[UtextValue]int{}))},
        src: `foo"`,
        vfn: func(_ int, v interface{}) {
            m := v.(map[UtextValue]int)
            assert.Equal(t, 1, len(m))
            for k := range m {
                assert.Equal(t, UtextValue(0), k)
            }
            assert.Equal(t, []byte("foo"), utextVar)
        },
        val: map[UtextValue]int{},
    }, 
    {
        key: "_OP_map_key_utext/pointer",
        ins: []_Instr{newInsVt(_OP_map_key_utext, reflect.TypeOf(map[*UtextStruct]int{}))},
        src: `foo"`,
        vfn: func(_ int, v interface{}) {
            m := v.(map[*UtextStruct]int)
            assert.Equal(t, 1, len(m))
            for k := range m {
                assert.Equal(t, "foo", k.V)
            }
        },
        val: map[*UtextStruct]int{},
    }, 
    {
        key: "_OP_map_key_utext_p",
        ins: []_Instr{newInsVt(_OP_map_key_utext_p, reflect.TypeOf(map[UtextStruct]int{}))},
        src: `foo"`,
        exp: map[UtextStruct]int{UtextStruct{V: "foo"}: 0},
        val: map[UtextStruct]int{},
    }, 
    {
        key: "_OP_array_skip",
        ins: []_Instr{newInsOp(_OP_array_skip)},
        src: `[1,2.0,true,false,null,"asdf",{"qwer":[1,2,3,4]}]`,
        pos: 1,
        vfn: func(i int, _ interface{}) { assert.Equal(t, 49, i) },
        val: nil,
    },{
        key: "_OP_slice_init",
        ins: []_Instr{newInsVt(_OP_slice_init, reflect.TypeOf(0))},
        src: "",
        vfn: func(_ int, v interface{}) {
            require.NotNil(t, v)
            assert.Equal(t, 0, len(v.([]int)))
            assert.Equal(t, _MinSlice, cap(v.([]int)))
        },
        val: new([]int),
    }, {
        key: "_OP_slice_append",
        ins: []_Instr{newInsVt(_OP_slice_append, reflect.TypeOf(0)), newInsOp(_OP_nil_1)},
        src: "",
        exp: []int{123, 0},
        val: &[]int{123},
    }, {
        key: "_OP_object_skip",
        ins: []_Instr{newInsOp(_OP_object_skip)},
        src: `{"zxcv":[1,2.0],"asdf":[true,false,null,"asdf",{"qwer":345}]}`,
        pos: 1,
        vfn: func(i int, _ interface{}) { assert.Equal(t, 61, i) },
        val: nil,
    }, {
        key: "_OP_object_next",
        ins: []_Instr{newInsOp(_OP_object_next)},
        src: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`,
        vfn: func(i int, _ interface{}) { assert.Equal(t, 52, i) },
        val: nil,
    }, {
        key: "_OP_struct_field",
        ins: []_Instr{
            newInsVf(_OP_struct_field, (func() *caching.FieldMap {
                ret := caching.CreateFieldMap(2)
                ret.Set("bab", 1)
                ret.Set("bac", 2)
                ret.Set("bad", 3)
                return ret
            })()),
            newInsOp(_OP_dbg_get_sr),
        },
        src: `bac"`,
        exp: 2,
        val: new(int),
    }, {
        key: "_OP_struct_field/case_insensitive",
        ins: []_Instr{
            newInsVf(_OP_struct_field, (func() *caching.FieldMap {
                ret := caching.CreateFieldMap(2)
                ret.Set("Bac", 2)
                ret.Set("BAC", 1)
                ret.Set("baC", 3)
                return ret
            })()),
            newInsOp(_OP_dbg_get_sr),
        },
        src: `bac"`,
        exp: 1,
        val: new(int),
    }, {
        key: "_OP_struct_field/not_found",
        ins: []_Instr{
            newInsVf(_OP_struct_field, (func() *caching.FieldMap {
                ret := caching.CreateFieldMap(2)
                ret.Set("bab", 1)
                ret.Set("bac", 2)
                ret.Set("bad", 3)
                return ret
            })()),
            newInsOp(_OP_dbg_get_sr),
        },
        src: `bae"`,
        exp: -1,
        val: new(int),
    }, {
        key: "_OP_unmarshal/value",
        ins: []_Instr{newInsVt(_OP_unmarshal, reflect.TypeOf(UjsonValue(0)))},
        src: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`,
        vfn: func(_ int, v interface{}) {
            assert.Equal(t, []byte(`{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`), ujsonVar)
        },
        val: new(UjsonValue),
    }, {
        key: "_OP_unmarshal/pointer",
        ins: []_Instr{newInsVt(_OP_unmarshal, reflect.TypeOf(new(UjsonStruct)))},
        src: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`,
        exp: &UjsonStruct{V: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`},
        val: new(*UjsonStruct),
    }, {
        key: "_OP_unmarshal_p",
        ins: []_Instr{newInsVt(_OP_unmarshal_p, reflect.TypeOf(new(UjsonStruct)))},
        src: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`,
        exp: UjsonStruct{V: `{"asdf":[1,2.0,true,false,null,"asdf",{"qwer":345}]}`},
        val: new(UjsonStruct),
    }, {
        key: "_OP_unmarshal_text/value",
        ins: []_Instr{newInsVt(_OP_unmarshal_text, reflect.TypeOf(UtextValue(0)))},
        src: `hello\n\r\tworld"`,
        vfn: func(_ int, v interface{}) {
            assert.Equal(t, []byte("hello\n\r\tworld"), utextVar)
        },
        val: new(UtextValue),
    }, {
        key: "_OP_unmarshal_text/pointer",
        ins: []_Instr{newInsVt(_OP_unmarshal_text, reflect.TypeOf(new(UtextStruct)))},
        src: `hello\n\r\tworld"`,
        exp: &UtextStruct{V: "hello\n\r\tworld"},
        val: new(*UtextStruct),
    }, {
        key: "_OP_unmarshal_text_p",
        ins: []_Instr{newInsVt(_OP_unmarshal_text_p, reflect.TypeOf(new(UtextStruct)))},
        src: `hello\n\r\tworld"`,
        exp: UtextStruct{V: "hello\n\r\tworld"},
        val: new(UtextStruct),
    }, {
        key: "_OP_lspace",
        ins: []_Instr{newInsOp(_OP_lspace)},
        src: " \t\r\na",
        vfn: func(i int, _ interface{}) { assert.Equal(t, 4, i) },
        val: nil,
    }, {
        key: "_OP_lspace/error",
        ins: []_Instr{newInsOp(_OP_lspace)},
        src: "",
        err: SyntaxError{Src: ``, Pos: 0, Code: types.ERR_EOF},
        val: nil,
    }, {
        key: "_OP_match_char/correct",
        ins: []_Instr{newInsVb(_OP_match_char, 'a')},
        src: "a",
        exp: nil,
        val: nil,
    }, {
        key: "_OP_match_char/error",
        ins: []_Instr{newInsVb(_OP_match_char, 'b')},
        src: "a",
        err: SyntaxError{Src: `a`, Pos: 0, Code: types.ERR_INVALID_CHAR},
        val: nil,
    }, {
        key: "_OP_switch",
        ins: []_Instr{
            newInsVi(_OP_dbg_set_sr, 1),
            newInsVs(_OP_switch, []int{4, 6, 8}),
            newInsOp(_OP_i8),
            newInsVi(_OP_goto, 9),
            newInsOp(_OP_i16),
            newInsVi(_OP_goto, 9),
            newInsOp(_OP_i32),
            newInsVi(_OP_goto, 9),
            newInsOp(_OP_u8),
        },
        src: "-1234567",
        exp: int32(-1234567),
        val: new(int32),
    },
    }
    for _, tv := range tests {
        t.Run(tv.key, func(t *testing.T) {
            testOpCode(t, &tv)
        })
    }
}

type JsonStruct struct {
    A int
    B string
    C map[string]int
    D []int
}

func TestAssembler_DecodeStruct(t *testing.T) {
    var v JsonStruct
    s := `{"A": 123, "B": "asdf", "C": {"qwer": 4567}, "D": [1, 2, 3, 4, 5]}`
    p, err := newCompiler().compile(reflect.TypeOf(v))
    require.NoError(t, err)
    k := new(_Stack)
    a := newAssembler(p)
    f := a.Load()
    pos, err := f(s, 0, unsafe.Pointer(&v), k, 0, "", nil)
    require.NoError(t, err)
    assert.Equal(t, len(s), pos)
    assert.Equal(t, JsonStruct{
        A: 123,
        B: "asdf",
        C: map[string]int{"qwer": 4567},
        D: []int{1, 2, 3, 4, 5},
    }, v)
}

type Tx struct {
    x int
}

func TestAssembler_DecodeStruct_SinglePrivateField(t *testing.T) {
    var v Tx
    s := `{"x": 1}`
    p, err := newCompiler().compile(reflect.TypeOf(v))
    require.NoError(t, err)
    k := new(_Stack)
    a := newAssembler(p)
    f := a.Load()
    pos, err := f(s, 0, unsafe.Pointer(&v), k, 0, "", nil)
    require.NoError(t, err)
    assert.Equal(t, len(s), pos)
    assert.Equal(t, Tx{}, v)
}

func TestAssembler_DecodeByteSlice_Bin(t *testing.T) {
    var v []byte
    s := `"aGVsbG8sIHdvcmxk"`
    p, err := newCompiler().compile(reflect.TypeOf(v))
    require.NoError(t, err)
    k := new(_Stack)
    a := newAssembler(p)
    f := a.Load()
    pos, err := f(s, 0, unsafe.Pointer(&v), k, 0, "", nil)
    require.NoError(t, err)
    assert.Equal(t, len(s), pos)
    assert.Equal(t, []byte("hello, world"), v)
}

func TestAssembler_DecodeByteSlice_List(t *testing.T) {
    var v []byte
    s := `[104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100]`
    p, err := newCompiler().compile(reflect.TypeOf(v))
    require.NoError(t, err)
    k := new(_Stack)
    a := newAssembler(p)
    f := a.Load()
    pos, err := f(s, 0, unsafe.Pointer(&v), k, 0, "", nil)
    require.NoError(t, err)
    assert.Equal(t, len(s), pos)
    assert.Equal(t, []byte("hello, world"), v)
}