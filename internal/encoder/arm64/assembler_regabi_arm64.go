//go:build arm64 && go1.17 && !go1.26
// +build arm64,go1.17,!go1.26

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

package arm64

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/bytedance/sonic/internal/encoder/alg"
	"github.com/bytedance/sonic/internal/encoder/ir"
	"github.com/bytedance/sonic/internal/encoder/prim"
	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/twitchyliquid64/golang-asm/obj"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/rt"
)

/** Register Allocations
 *
 *  State Registers:
 *
 *      %R29 : stack base
 *      %R20 : result pointer
 *      %R21 : result length
 *      %R22 : result capacity
 *      %R11 : sp->p
 *      %R12 : sp->q
 *      %R13 : sp->x
 *      %R14 : sp->f
 *
 *  Error Registers:
 *
 *      %R0 : error type register
 *      %R1 : error pointer register
 */

/** Function Prototype & Stack Map
 *
 *  func (buf *[]byte, p unsafe.Pointer, sb *_Stack, fv uint64) (err error)
 *
 *  buf    :  8(FP)
 *  p      : 16(FP)
 *  sb     : 24(FP)
 *  fv     : 32(FP)
 *  err.vt : 40(FP)
 *  err.vp : 48(FP)
 */

const (
	_S_cond = iota
	_S_init
)

var (
	// true: emit brk in control flow, so gdb automatically stops
	_DBG_BRK_EMIT = false
	// true: enable emit debug insn
	_DBG_BRK_ALL = false
)

const (
	_REG_BYTES = 8
	// 10 reg slots for other callee fns
	_FP_fargs = _REG_BYTES * 10
	_FP_args  = _FP_fargs
	// 10 reg slots for xsave op
	_FP_saves = _REG_BYTES * 10
	// 4 reg slots for other local var save & recover
	_FP_locals = _REG_BYTES * 4
	_locals    = _FP_locals
	// frame size = callee argument slots + save argument slots + local var slots + (previous) frame pointer slot + (current) return address slot
	_FP_size = _FP_fargs + _FP_saves + _FP_locals + _REG_BYTES*2
	// frame pointer offset = (current) stack top - 1 reg slot
	_FP_offs = _REG_BYTES * -1
	// (previous) frame pointer offset = frame size + (current) frame pointer offset
	_FP_base = _FP_size - _REG_BYTES*1
	// (previous) argument slot = frame size + return address slot
	// (previous) argument slot = frame pointer offset + return address slot + frame pointer slot
	_fargs_base = _FP_base + _REG_BYTES*2
	FP_offs = _FP_size - _REG_BYTES*2
)

const (
	_ARG_rb_off = _fargs_base
	_ARG_vp_off = _fargs_base + 8
	_ARG_sb_off = _fargs_base + 16
	_ARG_fv_off = _fargs_base + 24
)

var (
	_SP = jit.Reg("RSP")
)

var (
	// input argument: result buffer (for output)
	_ARG_rb = jit.Ptr(_SP, _ARG_rb_off)
	// input argument: value pointer (for input)
	_ARG_vp = jit.Ptr(_SP, _ARG_vp_off)
	// input argument: virtual machine stack base
	_ARG_sb = jit.Ptr(_SP, _ARG_sb_off)
	// input argument: bit flag
	_ARG_fv = jit.Ptr(_SP, _ARG_fv_off)
)

var (
	// virtual machine stack base/top
	_ST = jit.Reg("R19")
	// result buffer pointer
	_RP = jit.Reg("R20")
	// result buffer length
	_RL = jit.Reg("R21")
	// result buffer capacity
	_RC = jit.Reg("R22")
)

const (
	_VAR_sp_off = _FP_fargs + _FP_saves + _REG_BYTES*1
	_VAR_dn_off = _FP_fargs + _FP_saves + _REG_BYTES*2
	_VAR_vp_off = _FP_fargs + _FP_saves + _REG_BYTES*3
)

var (
	// local var: result buffer length
	_VAR_sp = jit.Ptr(_SP, _VAR_sp_off)
	// local var:
	//  input argument usage: the remain length of result buffer
	// output argument usage: the length of usage, result buffer
	_VAR_dn = jit.Ptr(_SP, _VAR_dn_off)
	// local var: value object pointer
	_VAR_vp = jit.Ptr(_SP, _VAR_vp_off)
)

// reg for input/output arguments, golang internal abi
var (
	_R0  = jit.Reg("R0")
	_R1  = jit.Reg("R1")
	_R2  = jit.Reg("R2")
	_R3  = jit.Reg("R3")
	_R4  = jit.Reg("R4")
	_R5  = jit.Reg("R5")
	_R6  = jit.Reg("R6")
	_R7  = jit.Reg("R7")
	_R8  = jit.Reg("R8")
	_R9  = jit.Reg("R9")
	_R10 = jit.Reg("R10")
	_R11 = jit.Reg("R11")
	_R12 = jit.Reg("R12")
	_R13 = jit.Reg("R13")
	_R14 = jit.Reg("R14")
	_R15 = jit.Reg("R15")
)

var (
	_R25 = jit.Reg("R25")
	// frame pointer reg
	_R29 = jit.Reg("R29")
	// link reg, return address reg
	_R30 = jit.Reg("R30")
)

var (
	// c abi: caller save
	_SP_p = jit.Reg("R11")
	// c abi: caller save
	_SP_q = jit.Reg("R12")
	// c abi: caller save
	_SP_x = jit.Reg("R13")
	// c abi: caller save
	_SP_f = jit.Reg("R14")
)

var (
	// output argument: error type
	_ET = jit.Reg("R0")
	// output argument: error info pointer
	_EP = jit.Reg("R1")
)

var (
	// c abi: caller save
	_T0 = jit.Reg("R23")
	// c abi: caller save
	_T1 = jit.Reg("R24")
	// c abi: caller save
	_T2 = jit.Reg("R25")
	// warning: carefully used
	_T3 = _R15
)

const (
	_FM_exp32 = 0x7f800000
	_FM_exp64 = 0x7ff0000000000000
)

// Strings in hex format
const (
	// null
	_IM_null = 0x6c6c756e
	// true
	_IM_true = 0x65757274
	// false
	_IM_fals = 0x736c6166
	// '"\"∅'
	_IM_open = 0x00225c22
	// []
	_IM_array = 0x5d5b
	// {}
	_IM_object = 0x7d7b
)

const (
	_LB_more_space        = "_more_space"
	_LB_more_space_return = "_more_space_return_"
)

const (
	_LB_error                 = "_error"
	_LB_error_too_deep        = "_error_too_deep"
	_LB_error_invalid_number  = "_error_invalid_number"
	_LB_error_nan_or_infinite = "_error_nan_or_infinite"
	_LB_panic                 = "_panic"
)

var (
	_F0 = jit.Reg("F0") // 浮点数寄存器
	_X0 = jit.Reg("V0") // 128bit -> NEON向量寄存器
	_Y0 = jit.Reg("V0") // 256bit -> SVE向量寄存器但是由于Reg中没有Z所以还是用V
)

var (
	_ZR = jit.Reg("ZR")
	_LR = _R30
)

var (
	_REG_ffi = []obj.Addr{_SP_p, _SP_q, _SP_x, _SP_f, _T2}
	_REG_b64 = []obj.Addr{_SP_p, _SP_q, _SP_x, _SP_f, _ST}
	_REG_all = []obj.Addr{_ST, _SP_x, _SP_f, _SP_p, _SP_q, _RP, _RL, _RC}
	_REG_ms  = []obj.Addr{_ST, _SP_x, _SP_f, _SP_q, _SP_p, _LR}
	_REG_enc = []obj.Addr{_ST, _SP_x, _SP_f, _SP_p, _SP_q, _RP, _RL, _RC}
)

type Assembler struct {
	Name string
	jit.BaseAssembler
	p ir.Program
	x int
}

func NewAssembler(p ir.Program) *Assembler {
	return new(Assembler).Init(p)
}

/** Assembler Interface **/

func (self *Assembler) Load() vars.Encoder {
	return ptoenc(self.BaseAssembler.Load("encode_"+self.Name, _FP_size, _FP_args, vars.ArgPtrs, vars.LocalPtrs))
}

func (self *Assembler) Init(p ir.Program) *Assembler {
	self.p = p
	self.BaseAssembler.Init(self.compile)
	return self
}

func (self *Assembler) compile() {
	self.prologue()
	self.instrs()
	self.epilogue()
	self.builtins()
}

/** Assembler Stages **/

var _OpFuncTab = [256]func(*Assembler, *ir.Instr){
	ir.OP_null:           (*Assembler)._asm_OP_null,
	ir.OP_empty_arr:      (*Assembler)._asm_OP_empty_arr,
	ir.OP_empty_obj:      (*Assembler)._asm_OP_empty_obj,
	ir.OP_bool:           (*Assembler)._asm_OP_bool,
	ir.OP_i8:             (*Assembler)._asm_OP_i8,
	ir.OP_i16:            (*Assembler)._asm_OP_i16,
	ir.OP_i32:            (*Assembler)._asm_OP_i32,
	ir.OP_i64:            (*Assembler)._asm_OP_i64,
	ir.OP_u8:             (*Assembler)._asm_OP_u8,
	ir.OP_u16:            (*Assembler)._asm_OP_u16,
	ir.OP_u32:            (*Assembler)._asm_OP_u32,
	ir.OP_u64:            (*Assembler)._asm_OP_u64,
	ir.OP_f32:            (*Assembler)._asm_OP_f32,
	ir.OP_f64:            (*Assembler)._asm_OP_f64,
	ir.OP_str:            (*Assembler)._asm_OP_str,
	ir.OP_bin:            (*Assembler)._asm_OP_bin,
	ir.OP_quote:          (*Assembler)._asm_OP_quote,
	ir.OP_number:         (*Assembler)._asm_OP_number,
	ir.OP_eface:          (*Assembler)._asm_OP_eface,
	ir.OP_iface:          (*Assembler)._asm_OP_iface,
	ir.OP_byte:           (*Assembler)._asm_OP_byte,
	ir.OP_text:           (*Assembler)._asm_OP_text,
	ir.OP_deref:          (*Assembler)._asm_OP_deref,
	ir.OP_index:          (*Assembler)._asm_OP_index,
	ir.OP_load:           (*Assembler)._asm_OP_load,
	ir.OP_save:           (*Assembler)._asm_OP_save,
	ir.OP_drop:           (*Assembler)._asm_OP_drop,
	ir.OP_drop_2:         (*Assembler)._asm_OP_drop_2,
	ir.OP_recurse:        (*Assembler)._asm_OP_recurse,
	ir.OP_is_nil:         (*Assembler)._asm_OP_is_nil,
	ir.OP_is_nil_p1:      (*Assembler)._asm_OP_is_nil_p1,
	ir.OP_is_zero_1:      (*Assembler)._asm_OP_is_zero_1,
	ir.OP_is_zero_2:      (*Assembler)._asm_OP_is_zero_2,
	ir.OP_is_zero_4:      (*Assembler)._asm_OP_is_zero_4,
	ir.OP_is_zero_8:      (*Assembler)._asm_OP_is_zero_8,
	ir.OP_is_zero_map:    (*Assembler)._asm_OP_is_zero_map,
	ir.OP_goto:           (*Assembler)._asm_OP_goto,
	ir.OP_map_iter:       (*Assembler)._asm_OP_map_iter,
	ir.OP_map_stop:       (*Assembler)._asm_OP_map_stop,
	ir.OP_map_check_key:  (*Assembler)._asm_OP_map_check_key,
	ir.OP_map_write_key:  (*Assembler)._asm_OP_map_write_key,
	ir.OP_map_value_next: (*Assembler)._asm_OP_map_value_next,
	ir.OP_slice_len:      (*Assembler)._asm_OP_slice_len,
	ir.OP_slice_next:     (*Assembler)._asm_OP_slice_next,
	ir.OP_marshal:        (*Assembler)._asm_OP_marshal,
	ir.OP_marshal_p:      (*Assembler)._asm_OP_marshal_p,
	ir.OP_marshal_text:   (*Assembler)._asm_OP_marshal_text,
	ir.OP_marshal_text_p: (*Assembler)._asm_OP_marshal_text_p,
	ir.OP_cond_set:       (*Assembler)._asm_OP_cond_set,
	ir.OP_cond_testc:     (*Assembler)._asm_OP_cond_testc,
	ir.OP_unsupported:    (*Assembler)._asm_OP_unsupported,
	ir.OP_is_zero:        (*Assembler)._asm_OP_is_zero,
}

func (self *Assembler) instr(v *ir.Instr) {
	if fn := _OpFuncTab[v.Op()]; fn != nil {
		fn(self, v)
	} else {
		panic(fmt.Sprintf("invalid opcode: %d", v.Op()))
	}
}

func (self *Assembler) instrs() {
	for i, v := range self.p {
		self.Mark(i)
		self.instr(&v)
	}
}

func (self *Assembler) builtins() {
	self.more_space()
	self.error_too_deep()
	self.error_invalid_number()
	self.error_nan_or_infinite()
	self.go_panic()
}

func (self *Assembler) epilogue() {
	self.Mark(len(self.p))
	self.Emit("MOVD", _ZR, _ET)
	self.Emit("MOVD", _ZR, _EP)
	self.Link(_LB_error)
	self.Emit("MOVD", _ARG_rb, _T0)
	self.Emit("MOVD", _RL, jit.Ptr(_T0, 8))
	self.Emit("MOVD", _ZR, _ARG_rb)
	self.Emit("MOVD", _ZR, _ARG_vp)
	self.Emit("MOVD", _ZR, _ARG_sb)
	self.Emit("MOVD", jit.Ptr(_SP, _FP_offs), _R29)
	self.Emit("MOVD", jit.Ptr(_SP, 0), _LR)
	self.Emit("ADD", jit.Imm(_FP_size), _SP)
	self.Emit("RET", _LR)
}

func (self *Assembler) prologue() {
	self.Emit("NOP")
	self.Emit("SUB", jit.Imm(_FP_size), _SP)
	self.Emit("MOVD", _R29, jit.Ptr(_SP, _FP_offs))
	self.Emit("MOVD", _LR, jit.Ptr(_SP, 0))
	self.Emit("MOVD", jit.Imm(_FP_offs), _T0)
	self.Emit("MOVD", _SP, _T1)
	self.EmitAdd(_R29, _T0, _T1)
	self.Emit("MOVD", _R0, _ARG_rb)
	self.Emit("MOVD", _R1, _ARG_vp)
	self.Emit("MOVD", _R2, _ARG_sb)
	self.Emit("MOVD", _R3, _ARG_fv)
	self.Emit("MOVD", jit.Ptr(_R0, 0), _RP)
	self.Emit("MOVD", jit.Ptr(_R0, 8), _RL)
	self.Emit("MOVD", jit.Ptr(_R0, 16), _RC)
	self.Emit("MOVD", _R1, _SP_p)
	self.Emit("MOVD", _R2, _ST)
	self.Emit("MOVD", _ZR, _SP_x)
	self.Emit("MOVD", _ZR, _SP_f)
	self.Emit("MOVD", _ZR, _SP_q)
}

/** Assembler Inline Functions **/

func (self *Assembler) xsave(reg ...obj.Addr) {
	for i, v := range reg {
		if i > _FP_saves/8-1 {
			panic("too many registers to save")
		} else {
			self.EmitStur(jit.Ptr, v, _SP, _FP_fargs+int64(i)*8+8)
		}
	}
}

func (self *Assembler) xload(reg ...obj.Addr) {
	for i, v := range reg {
		if i > _FP_saves/8-1 {
			panic("too many registers to load")
		} else {
			self.EmitLdur(jit.Ptr, _SP, _FP_fargs+int64(i)*8+8, v)
		}
	}
}

func (self *Assembler) rbuf_di() {
	self.EmitAdd(_R0, _RP, _RL)
}

func (self *Assembler) store_int(nd int, fn obj.Addr, ins string) {
	self.check_size(nd)
	self.save_c()
	self.rbuf_di()
	self.Emit(ins, jit.Ptr(_SP_p, 0), _R1)
	self.call_c(fn)
	self.EmitAdd(_RL, _R0, _RL)
}

func (self *Assembler) write_imm_to_mem(op string, imm obj.Addr, base0 obj.Addr, base1 obj.Addr, base2_imm int64, script_reg_v obj.Addr, script_reg_a obj.Addr) {
	self.Emit("MOVD", imm, script_reg_v)
	self.EmitStrriWithInsn(jit.Ptr, op, script_reg_v, base0, base1, base2_imm, script_reg_a)
}

func (self *Assembler) store_str(s string) {
	i := 0
	m := rt.Str2Mem(s)

	/* 8-byte stores */
	for i <= len(m)-8 {
		self.write_imm_to_mem("MOVD", jit.Imm(rt.Get64(m[i:])), _RP, _RL, int64(i), _T0, _T1)
		i += 8
	}

	/* 4-byte stores */
	if i <= len(m)-4 {
		self.write_imm_to_mem("MOVW", jit.Imm(int64(rt.Get32(m[i:]))), _RP, _RL, int64(i), _T0, _T1)
		i += 4
	}

	/* 2-byte stores */
	if i <= len(m)-2 {
		self.write_imm_to_mem("MOVH", jit.Imm(int64(rt.Get16(m[i:]))), _RP, _RL, int64(i), _T0, _T1)
		i += 2
	}

	/* last byte */
	if i < len(m) {
		self.write_imm_to_mem("MOVB", jit.Imm(int64(m[i])), _RP, _RL, int64(i), _T0, _T1)
	}
}

func (self *Assembler) check_size(n int) {
	self.Emit("MOVD", jit.Imm(int64(n)), _T1)
	self.check_size_r_auto(_T1)
}

func (self *Assembler) check_size_r_auto(r obj.Addr) {
	idx := self.x
	key := _LB_more_space_return + strconv.Itoa(idx)

	/* check for buffer capacity */
	self.x++
	self.EmitAdd(_T0, _RL, r)
	self.EmitCmpq(_RC, _T0) // CMPQ AX, RC
	self.Sjmp("BLE", key)   // BLE  _more_space_return_{n}
	self.slice_grow(key)    // GROW $key
	self.Link(key)          // _more_space_return_{n}:
}

func (self *Assembler) slice_grow(ret string) {
	self.SrefRm(ret, 0, 30)
	self.Sjmp("B", _LB_more_space)
}

/** State Stack Helpers **/

func (self *Assembler) save_state() {
	self.EmitLdur(jit.Ptr, _ST, 0, _T0)
	self.EmitAdd(_T1, _T0, jit.Imm(vars.StateSize))     // ADD X1, X0, #vars.StateSize
	self.EmitCmpqLiteral(jit.Imm(vars.StackLimit), _T1) // CMP #vars.StackLimit, X1
	self.Emit("MOVD", _T1, _T3)
	self.Sjmp("BHS", _LB_error_too_deep) // B.HS _error_too_deep

	// save current states
	self.EmitStrri(jit.Ptr, _SP_x, _ST, _T0, 8, _T1)  // MOVQ SP.x, 8(ST)(CX)
	self.EmitStrri(jit.Ptr, _SP_f, _ST, _T0, 16, _T1) // MOVQ SP.f, 16(ST)(CX)
	self.WritePtr(0, _SP_p, _ST, _T0, 24, _T1)
	self.WritePtr(1, _SP_q, _ST, _T0, 32, _T1)
	self.EmitStur(jit.Ptr, _T3, _ST, 0)
}

func (self *Assembler) drop_state(decr int64) {
	self.EmitLdur(jit.Ptr, _ST, 0, _T2)
	self.Emit("SUB", jit.Imm(decr), _T2)
	self.EmitStur(jit.Ptr, _T2, _ST, 0)
	self.EmitLdrri(jit.Ptr, _ST, _T2, 8, _SP_x, _T0)
	self.EmitLdrri(jit.Ptr, _ST, _T2, 16, _SP_f, _T0)
	self.EmitLdrri(jit.Ptr, _ST, _T2, 24, _SP_p, _T0)
	self.EmitLdrri(jit.Ptr, _ST, _T2, 32, _SP_q, _T0)
	self.EmitStrri(jit.Ptr, _ZR, _ST, _T2, 8, _T0)
	self.EmitStrri(jit.Ptr, _ZR, _ST, _T2, 16, _T0)
	self.EmitStrri(jit.Ptr, _ZR, _ST, _T2, 24, _T0)
	self.EmitStrri(jit.Ptr, _ZR, _ST, _T2, 32, _T0)
}

/** Buffer Helpers **/

func (self *Assembler) add_char(ch byte) {
	self.write_imm_to_mem("MOVB", jit.Imm(int64(ch)), _RP, _RL, 0, _T0, _T1)
	self.Emit("ADD", jit.Imm(1), _RL)
}

func (self *Assembler) add_long(ch uint32, n int64) {
	self.write_imm_to_mem("MOVW", jit.Imm(int64(ch)), _RP, _RL, 0, _T0, _T1)
	self.Emit("ADD", jit.Imm(n), _RL)
}

func (self *Assembler) add_text(ss string) {
	self.store_str(ss)
	self.EmitAdd(_RL, _RL, jit.Imm(int64(len(ss))))
}

// get *buf at AX
func (self *Assembler) prep_buffer() {
	self.Emit("MOVD", _ARG_rb, _R0)         // MOVQ rb<>+0(FP), AX
	self.Emit("MOVD", _RL, jit.Ptr(_R0, 8)) // MOVQ RL, 8(AX)
}

func (self *Assembler) save_buffer() {
	self.Emit("MOVD", _ARG_rb, _T0)          // MOVQ rb<>+0(FP), CX
	self.Emit("MOVD", _RP, jit.Ptr(_T0, 0))  // MOVQ RP, (CX)
	self.Emit("MOVD", _RL, jit.Ptr(_T0, 8))  // MOVQ RL, 8(CX)
	self.Emit("MOVD", _RC, jit.Ptr(_T0, 16)) // MOVQ RC, 16(CX)
}

// get *buf at AX
func (self *Assembler) load_buffer() {
	self.Emit("MOVD", _ARG_rb, _R0)          // MOVQ rb<>+0(FP), AX
	self.Emit("MOVD", jit.Ptr(_R0, 0), _RP)  // MOVQ (AX), RP
	self.Emit("MOVD", jit.Ptr(_R0, 8), _RL)  // MOVQ 8(AX), RL
	self.Emit("MOVD", jit.Ptr(_R0, 16), _RC) // MOVQ 16(AX), RC
}

/** Function Interface Helpers **/

func (self *Assembler) call(pc obj.Addr) {
	self.Emit("MOVD", pc, _LR) // MOVQ $pc, AX
	self.Rjmp("CALL", _LR)     // CALL AX
}

func (self *Assembler) save_c() {
	self.xsave(_REG_ffi...) // SAVE $REG_ffi
}

func (self *Assembler) call_b64(pc obj.Addr) {
	self.xsave(_REG_b64...) // SAVE $REG_all
	self.call(pc)           // CALL $pc
	self.xload(_REG_b64...) // LOAD $REG_ffi
}

func (self *Assembler) call_c(pc obj.Addr) {
	self.call(pc)           // CALL $pc
	self.xload(_REG_ffi...) // LOAD $REG_ffi
}

func (self *Assembler) call_go(pc obj.Addr) {
	self.xsave(_REG_all...) // SAVE $REG_all
	self.call(pc)           // CALL $pc
	self.xload(_REG_all...) // LOAD $REG_all
}

func (self *Assembler) call_more_space(pc obj.Addr) {
	self.xsave(_REG_ms...) // SAVE $REG_all
	self.call(pc)          // CALL $pc
	self.xload(_REG_ms...) // LOAD $REG_all
}

func (self *Assembler) call_encoder(pc obj.Addr) {
	self.xsave(_REG_enc...) // SAVE $REG_all
	self.call(pc)           // CALL $pc
	self.xload(_REG_enc...) // LOAD $REG_all
}

func (self *Assembler) call_marshaler(fn obj.Addr, it *rt.GoType, vt reflect.Type) {
	switch vt.Kind() {
	case reflect.Interface:
		self.call_marshaler_i(fn, it)
	case reflect.Ptr, reflect.Map:
		self.call_marshaler_v(fn, it, vt, true)
	// struct/array of 1 direct iface type can be direct
	default:
		self.call_marshaler_v(fn, it, vt, !rt.UnpackType(vt).Indirect())
	}
}

var (
	_F_assertI2I = jit.Func(rt.AssertI2I)
)

func (self *Assembler) call_marshaler_i(fn obj.Addr, it *rt.GoType) {
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R0) // MOVQ    (SP.p), AX
	self.EmitCbzS("CBZ", _R0, "_null_{n}")
	self.Emit("MOVD", _R0, _R1)               // MOVQ    AX, BX
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R2) // MOVQ    8(SP.p), CX
	self.Emit("MOVD", jit.Gtype(it), _R0)     // MOVQ    $it, AX

	self.call_go(_F_assertI2I) // CALL_GO assertI2I
	self.EmitCbzS("CBZ", _R0, "_null_{n}")
	self.Emit("MOVD", _R1, _R2) // MOVQ   BX, CX
	self.Emit("MOVD", _R0, _R1) // MOVQ   AX, BX

	self.prep_buffer()
	self.Emit("MOVD", _ARG_fv, _R3) // MOVQ   ARG.fv, DI
	self.call_go(fn)                // CALL    $fn
	self.EmitCbzS("CBNZ", _ET, _LB_error)

	self.load_buffer()
	self.Sjmp("B", "_done_{n}")                                             // JMP     _done_{n}
	self.Link("_null_{n}")                                                  // _null_{n}:
	self.check_size(4)                                                      // SIZE    $4
	self.write_imm_to_mem("MOVW", jit.Imm(_IM_null), _RP, _RL, 0, _T0, _T1) // MOVL    $'null', (RP)(RL*1)
	self.Emit("ADD", jit.Imm(4), _RL)                                       // ADDQ    $4, RL
	self.Link("_done_{n}")                                                  // _done_{n}:
}

func (self *Assembler) call_marshaler_v(fn obj.Addr, it *rt.GoType, vt reflect.Type, deref bool) {
	self.prep_buffer()                       // MOVE {buf}, (SP)
	self.Emit("MOVD", jit.Itab(it, vt), _R1) // MOVQ $(itab(it, vt)), BX

	/* dereference the pointer if needed */
	if !deref {
		self.Emit("MOVD", _SP_p, _R2) // MOVQ SP.p, CX
	} else {
		self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R2) // MOVQ 0(SP.p), CX
	}

	/* call the encoder, and perform error checks */
	self.Emit("MOVD", _ARG_fv, _R3) // MOVQ   ARG.fv, DI
	self.call_go(fn)                // CALL  $fn
	self.EmitCbzS("CBNZ", _ET, _LB_error)

	self.load_buffer()
}

/** Builtin: _more_space **/

var (
	_T_byte      = jit.Type(vars.ByteType)
	_F_growslice = jit.Func(rt.GrowSlice)

	_T_json_Marshaler         = rt.UnpackType(vars.JsonMarshalerType)
	_T_encoding_TextMarshaler = rt.UnpackType(vars.EncodingTextMarshalerType)
)

// input: _T0
func (self *Assembler) more_space() {
	self.Link(_LB_more_space)
	self.Emit("MOVD", _T_byte, _R0)
	self.Emit("MOVD", _RP, _R1)
	self.Emit("MOVD", _RL, _R2)
	self.Emit("MOVD", _RC, _R3)
	self.Emit("MOVD", _T0, _R4)
	self.call_more_space(_F_growslice)
	self.Emit("MOVD", _R0, _RP)
	self.Emit("MOVD", _R1, _RL)
	self.Emit("MOVD", _R2, _RC)
	self.save_buffer()
	self.Rjmp("B", _LR)
}

/** Builtin Errors **/

var (
	_V_ERR_too_deep               = jit.Imm(int64(uintptr(unsafe.Pointer(vars.ERR_too_deep))))
	_V_ERR_nan_or_infinite        = jit.Imm(int64(uintptr(unsafe.Pointer(vars.ERR_nan_or_infinite))))
	_I_json_UnsupportedValueError = jit.Itab(rt.UnpackType(vars.ErrorType), vars.JsonUnsupportedValueType)
)

func (self *Assembler) error_too_deep() {
	self.Link(_LB_error_too_deep)
	self.Emit("MOVD", _V_ERR_too_deep, _EP)
	self.Emit("MOVD", _I_json_UnsupportedValueError, _ET)
	self.Sjmp("B", _LB_error)
}

func (self *Assembler) error_invalid_number() {
	self.Link(_LB_error_invalid_number)
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R0)
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R1)
	self.call_go(_F_error_number)
	self.Sjmp("B", _LB_error)
}

func (self *Assembler) error_nan_or_infinite() {
	self.Link(_LB_error_nan_or_infinite)
	self.Emit("MOVD", _V_ERR_nan_or_infinite, _EP)
	self.Emit("MOVD", _I_json_UnsupportedValueError, _ET)
	self.Sjmp("B", _LB_error)
}

/** String Encoding Routine **/

var (
	_F_quote = jit.Imm(int64(native.S_quote))
	_F_panic = jit.Func(vars.GoPanic)
)

func (self *Assembler) go_panic() {
	self.Link(_LB_panic)
	self.Emit("MOVD", _SP_p, _R1)
	self.Emit("MOVD", _RP, _R2)
	self.Emit("MOVD", _RL, _R3)
	self.call_go(_F_panic)
}

func (self *Assembler) encode_string(doubleQuote bool) {
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _T2)
	self.EmitCbzS("CBZ", _T2, "_str_empty_{n}")
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzS("CBNZ", _T0, "_str_next_{n}")
	self.Emit("MOVD", jit.Imm(int64(vars.PanicNilPointerOfNonEmptyString)), _R0)
	self.Sjmp("B", _LB_panic)
	self.Link("_str_next_{n}")

	// opening quote, check for double quote
	if !doubleQuote {
		self.EmitAdd(_T0, _T2, jit.Imm(2))
		self.check_size_r_auto(_T0)
		self.add_char('"')
	} else {
		self.EmitAdd(_T0, _T2, jit.Imm(6))
		self.check_size_r_auto(_T0)
		self.add_long(_IM_open, 3)
	}

	// quoting loop
	self.Emit("MOVD", _ZR, _T2)
	self.Emit("MOVD", _T2, _VAR_sp) // count
	self.Link("_str_loop_{n}")      // _str_loop_{n}:
	self.save_c()                   // SAVE $REG_ffi

	// load the output buffer first, and then input buffer,
	// because the parameter registers collide with RP / RL / RC
	self.Emit("MOVD", _RC, _T0)     // MOVQ RC, CX
	self.Emit("SUB", _RL, _T0)      // SUBQ RL, CX
	self.Emit("MOVD", _T0, _VAR_dn) // MOVQ CX, dn

	self.EmitAdd(_R2, _RP, _RL) // LEAQ (RP)(RL), DX
	self.Emit("MOVD", jit.Imm(_VAR_dn_off), _T0)
	self.Emit("MOVD", _SP, _R3)
	self.EmitAdd(_R3, _T0, _R3)

	self.Emit("MOVD", _VAR_sp, _T2)

	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R0)
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R1)
	self.EmitAdd(_R0, _R0, _T2)
	self.Emit("SUB", _T2, _R1)

	// set the flags based on `doubleQuote` (R4)
	if !doubleQuote {
		// R4 = 0
		self.Emit("MOVD", _ZR, _R4)
	} else {
		// R4 = ${types.F_DOUBLE_UNQUOTE}
		self.Emit("MOVW", jit.Imm(types.F_DOUBLE_UNQUOTE), _R4)
	}

	// call the native quoter
	self.call_c(_F_quote)
	self.Emit("MOVD", _VAR_dn, _T0)
	self.Emit("ADD", _T0, _RL)

	self.Emit("MOVD", _R0, _T2)
	self.EmitTbzS("TBNZ", jit.Imm(int64(63)), _R0, "_str_space_{n}")

	// close the string, check for double quote
	if !doubleQuote {
		// size $1
		// char $'"'
		// B    _str_end_{n}
		self.check_size(1)
		self.add_char('"')
		self.Sjmp("B", "_str_end_{n}")
	} else {
		self.check_size(3)
		self.add_text("\\\"\"")
		self.Sjmp("B", "_str_end_{n}")
	}

	// not enough space to contain the quoted string
	self.Link("_str_space_{n}") // _str_space_{n}:
	self.Emit("MVN", _T2, _T2)
	self.Emit("MOVD", _VAR_sp, _T0)
	self.Emit("ADD", _T2, _T0)
	self.Emit("MOVD", _T0, _VAR_sp)
	self.EmitAdd(_T2, _RC, _RC)
	self.Emit("MOVD", _T2, _T0)
	self.slice_grow("_str_loop_{n}") // GROW _str_loop_{n}

	// empty string, check for double quote
	if !doubleQuote {
		// _str_empty_{n}:
		// SIZE $2
		// TEXT $'""'
		// _str_end_{n}:
		self.Link("_str_empty_{n}")
		self.check_size(2)
		self.add_text("\"\"")
		self.Link("_str_end_{n}")
	} else {
		self.Link("_str_empty_{n}")
		self.check_size(6)
		self.add_text("\"\\\"\\\"\"")
		self.Link("_str_end_{n}")
	}
}

/** OpCode Assembler Functions **/

var (
	_F_f64toa    = jit.Imm(int64(native.S_f64toa))
	_F_f32toa    = jit.Imm(int64(native.S_f32toa))
	_F_i64toa    = jit.Imm(int64(native.S_i64toa))
	_F_u64toa    = jit.Imm(int64(native.S_u64toa))
	_F_b64encode = jit.Imm(int64(rt.SubrB64Encode))
)

var (
	_F_memmove       = jit.Func(rt.Memmove)
	_F_error_number  = jit.Func(vars.Error_number)
	_F_isValidNumber = jit.Func(alg.IsValidNumber)
)

var (
	_F_iteratorStop  = jit.Func(alg.IteratorStop)
	_F_iteratorNext  = jit.Func(alg.IteratorNext)
	_F_iteratorStart = jit.Func(alg.IteratorStart)
)

var (
	_F_encodeTypedPointer  obj.Addr
	_F_encodeJsonMarshaler obj.Addr
	_F_encodeTextMarshaler obj.Addr
)

func init() {
	_F_encodeJsonMarshaler = jit.Func(prim.EncodeJsonMarshaler)
	_F_encodeTextMarshaler = jit.Func(prim.EncodeTextMarshaler)
	_F_encodeTypedPointer = jit.Func(EncodeTypedPointer)
}

func (self *Assembler) _asm_OP_null(_ *ir.Instr) {
	self._debug_tag(0x01) // 0x001
	self.check_size(4)
	self.write_imm_to_mem("MOVW", jit.Imm(_IM_null), _RP, _RL, int64(0), _T0, _T1)
	self.EmitAdd(_RL, _RL, jit.Imm(4))
}

func (self *Assembler) _asm_OP_empty_arr(_ *ir.Instr) {
	self._debug_tag(0x02) // 0x002
	self.EmitTbzSLdriPtr(jit.Ptr, "TBNZ", jit.Imm(int64(alg.BitNoNullSliceOrMap)), _ARG_fv, "_empty_arr_{n}", _T0)
	self._asm_OP_null(nil)
	self.Sjmp("B", "_empty_arr_end_{n}")
	self.Link("_empty_arr_{n}")
	self.check_size(2)
	self.Emit("MOVD", jit.Imm(_IM_array), _T0)
	self.Emit("MOVH", _T0, jit.Sib(_RP, _RL, 1, 0))
	self.Emit("ADD", jit.Imm(2), _RL)
	self.Link("_empty_arr_end_{n}")
}

func (self *Assembler) _asm_OP_empty_obj(_ *ir.Instr) {
	self._debug_tag(0x03) // 0x003
	self.EmitTbzSLdriPtr(jit.Ptr, "TBNZ", jit.Imm(int64(alg.BitNoNullSliceOrMap)), _ARG_fv, "_empty_obj_{n}", _T0)
	self._asm_OP_null(nil)
	self.Sjmp("B", "_empty_obj_end_{n}")
	self.Link("_empty_obj_{n}")
	self.check_size(2)
	self.write_imm_to_mem("MOVH", jit.Imm(_IM_object), _RP, _RL, 0, _T0, _T1)
	self.Emit("ADD", jit.Imm(2), _RL)
	self.Link("_empty_obj_end_{n}")
}

func (self *Assembler) _asm_OP_bool(_ *ir.Instr) {
	self._debug_tag(0x04) // 0x004
	self.Emit("MOVB", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzS("CBZ", _T0, "_false_{n}")
	self.check_size(4) // SIZE $4
	self.Emit("MOVD", jit.Imm(_IM_true), _T0)
	self.Emit("MOVW", _T0, jit.Sib(_RP, _RL, 1, 0)) // MOVL $'true', (RP)(RL*1)
	self.Emit("ADD", jit.Imm(4), _RL)               // ADDQ $4, RL
	self.Sjmp("B", "_end_{n}")                      // JMP  _end_{n}
	self.Link("_false_{n}")                         // _false_{n}:
	self.check_size(5)                              // SIZE $5
	self.Emit("MOVD", jit.Imm(_IM_fals), _T0)
	self.Emit("MOVW", _T0, jit.Sib(_RP, _RL, 1, 0)) // MOVL $'fals', (RP)(RL*1)
	self.EmitAdd(_RL, _RL, jit.Imm(4))
	self.Emit("MOVD", jit.Imm('e'), _T0)            // MOVB $'e', 4(RP)(RL*1)
	self.Emit("MOVB", _T0, jit.Sib(_RP, _RL, 1, 0)) // MOVB $'e', 4(RP)(RL*1)
	self.EmitAdd(_RL, _RL, jit.Imm(1))              // ADDQ $5, RL
	self.Link("_end_{n}")                           // _end_{n}:
}

func (self *Assembler) _asm_OP_i8(_ *ir.Instr) {
	self._debug_tag(0x05) // 0x005
	self.store_int(4, _F_i64toa, "MOVB")
}

func (self *Assembler) _asm_OP_i16(_ *ir.Instr) {
	self._debug_tag(0x06) // 0x006
	self.store_int(6, _F_i64toa, "MOVH")
}

func (self *Assembler) _asm_OP_i32(_ *ir.Instr) {
	self._debug_tag(0x07) // 0x007
	self.store_int(17, _F_i64toa, "MOVW")
}

func (self *Assembler) _asm_OP_i64(_ *ir.Instr) {
	self._debug_tag(0x08) // 0x008
	self.store_int(21, _F_i64toa, "MOVD")
}

func (self *Assembler) _asm_OP_u8(_ *ir.Instr) {
	self._debug_tag(0x09) // 0x009
	self.store_int(3, _F_u64toa, "MOVBU")
}

func (self *Assembler) _asm_OP_u16(_ *ir.Instr) {
	self._debug_tag(0x0a) // 0x00a
	self.store_int(5, _F_u64toa, "MOVHU")
}

func (self *Assembler) _asm_OP_u32(_ *ir.Instr) {
	self._debug_tag(0x0b) // 0x00b
	self.store_int(16, _F_u64toa, "MOVWU")
}

func (self *Assembler) _asm_OP_u64(_ *ir.Instr) {
	self._debug_tag(0x0c) // 0x00c
	self.store_int(20, _F_u64toa, "MOVD")
}

func (self *Assembler) _asm_OP_f32(_ *ir.Instr) {
	self._debug_tag(0x0d) // 0x00d
	self.check_size(32)
	self.Emit("MOVW", jit.Ptr(_SP_p, 0), _T0)                                                                            // MOVL     (SP.p), AX
	self.Emit("AND", jit.Imm(_FM_exp32), _T0)                                                                            // ANDL     $_FM_exp32, AX
	self.Emit("EOR", jit.Imm(_FM_exp32), _T0)                                                                            // XORL     $_FM_exp32, AX
	self.EmitCbzS("CBNZ", _T0, "_encode_normal_f32_{n}")                                                                 // JNZ      _encode_normal_f32_{n};
	self.EmitTbzSLdriPtr(jit.Ptr, "TBZ", jit.Imm(alg.BitEncodeNullForInfOrNan), _ARG_fv, _LB_error_nan_or_infinite, _T1) // BTQ ${BitEncodeNullForInfOrNan}, fv; JNC  _error_nan_or_infinite
	self._asm_OP_null(nil)
	self.Sjmp("JMP", "_encode_f32_end_{n}") // JMP      _encode_f32_end_{n}
	self.Link("_encode_normal_f32_{n}")
	self.save_c()                              // SAVE     $C_regs
	self.rbuf_di()                             // MOVQ     RP, DI
	self.Emit("FMOVS", jit.Ptr(_SP_p, 0), _F0) // MOVSS    (SP.p), X0
	self.call_c(_F_f32toa)                     // CALL_C   f32toa
	self.Emit("ADD", _R0, _RL)                 // ADDQ   AX, RL
	self.Link("_encode_f32_end_{n}")
}

func (self *Assembler) _asm_OP_f64(_ *ir.Instr) {
	self._debug_tag(0x0e) // 0x00e
	self.check_size(32)
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _T0)  // MOVQ   (SP.p), AX
	self.Emit("MOVD", jit.Imm(_FM_exp64), _T1) // MOVQ   $_FM_exp64, CX
	self.Emit("AND", _T1, _T0)                 // ANDQ   CX, AX
	self.Emit("EOR", _T1, _T0)                 // XORQ   CX, AX

	self.EmitCbzS("CBNZ", _T0, "_encode_normal_f64_{n}") // JNZ    _encode_normal_f64_{n}

	self.EmitTbzSLdriPtr(jit.Ptr, "TBZ", jit.Imm(alg.BitEncodeNullForInfOrNan), _ARG_fv, _LB_error_nan_or_infinite, _T2) // BTQ ${BitEncodeNullForInfOrNan}, fv; JNC  _error_nan_or_infinite

	self._asm_OP_null(nil)

	self.Sjmp("JMP", "_encode_f64_end_{n}") // JMP    _encode_f64_end_{n}

	self.Link("_encode_normal_f64_{n}")
	self.save_c()                              // SAVE   $C_regs
	self.rbuf_di()                             // MOVQ   RP, DI
	self.Emit("FMOVD", jit.Ptr(_SP_p, 0), _F0) // MOVSD  (SP.p), X0
	self.call_c(_F_f64toa)                     // CALL_C f64toa
	self.Emit("ADD", _R0, _RL)                 // ADDQ   AX, RL
	self.Link("_encode_f64_end_{n}")
}

func (self *Assembler) _asm_OP_str(_ *ir.Instr) {
	self._debug_tag(0x0f) // 0x00f
	self.encode_string(false)
}

func (self *Assembler) _asm_OP_bin(_ *ir.Instr) {
	self._debug_tag(0x10) // 0x010
	self.Emit("MOVD", _RP, _R0)
	self.Emit("MOVD", _RL, _R1)
	self.Emit("MOVD", _RC, _R2)
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R3)
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R4)
	self.Emit("MOVD", jit.Ptr(_SP_p, 16), _R5)
	self.call_b64(_F_b64encode)
	self.Emit("MOVD", _R0, _RP)
	self.Emit("MOVD", _R1, _RL)
	self.Emit("MOVD", _R2, _RC)
	self.save_buffer()
}

func (self *Assembler) _asm_OP_quote(_ *ir.Instr) {
	self._debug_tag(0x11) // 0x011
	self.encode_string(true)
}

func (self *Assembler) _asm_OP_number(_ *ir.Instr) {
	self._debug_tag(0x12)                     // 0x012
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R1) // MOVQ    (SP.p), BX
	self.EmitCbzS("CBZ", _R1, "_empty_{n}")   // TESTQ   BX, BX; JZ _empty_{n}
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R0) // MOVQ    (SP.p), AX

	self.EmitCbzS("CBNZ", _R0, "_number_next_{n}") // TESTQ   AX, AX; JNZ _number_next_{n}
	self.Emit("MOVD", jit.Imm(int64(vars.PanicNilPointerOfNonEmptyString)), _R0)
	self.Sjmp("B", _LB_panic)
	self.Link("_number_next_{n}")
	self.call_go(_F_isValidNumber) // CALL_GO isValidNumber

	self.Emit("AND", jit.Imm(0xff), _R0)
	self.EmitCbzS("CBZ", _R0, _LB_error_invalid_number)

	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _T0) // MOVQ    (SP.p), BX
	self.check_size_r_auto(_T0)               // SIZE    BX

	self.EmitAdd(_T1, _RP, _RL)
	self.Emit("MOVD", _T1, _R0)

	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _T2) // ADDQ    8(SP.p), RL
	self.EmitAdd(_RL, _T2, _RL)

	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R1) // MOVOU   (SP.p), BX
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _R2) // MOVOU   X0, 8(SP)

	self.call_go(_F_memmove) // CALL_GO memmove

	self.Emit("MOVD", _ARG_rb, _T0) // MOVQ rb<>+0(FP), AX

	self.Emit("MOVD", _RL, jit.Ptr(_T0, 8)) // MOVQ RL, 8(AX)

	self.Sjmp("JMP", "_done_{n}") // JMP     _done_{n}
	self.Link("_empty_{n}")       // _empty_{n}
	self.check_size(1)            // SIZE    $1
	self.add_char('0')            // CHAR    $'0'
	self.Link("_done_{n}")        // _done_{n}:
}

func (self *Assembler) _asm_OP_eface(_ *ir.Instr) {
	self.prep_buffer()                        // MOVE  {buf}, AX
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R1) // MOVQ  (SP.p), BX
	self.EmitAdd(_R2, _SP_p, jit.Imm(8))      // LEAQ  8(SP.p), CX
	self.Emit("MOVD", _ST, _R3)               // MOVQ  ST, DI
	self.Emit("MOVD", _ARG_fv, _R4)           // MOVQ  fv, AX
	self.call_encoder(_F_encodeTypedPointer)  // CALL  encodeTypedPointer
	self.EmitCbzS("CBNZ", _ET, _LB_error)     // TESTQ ET, ET; JNZ   _error
	self.load_buffer()
}

func (self *Assembler) _asm_OP_iface(_ *ir.Instr) {
	self.prep_buffer()                        // MOVE  {buf}, AX
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R2) // MOVQ  (SP.p), CX
	self.Emit("MOVD", jit.Ptr(_R2, 8), _R1)   // MOVQ  8(CX), BX
	self.EmitAdd(_R2, _SP_p, jit.Imm(8))      // LEAQ  8(SP.p), CX
	self.Emit("MOVD", _ST, _R3)               // MOVQ  ST, DI
	self.Emit("MOVD", _ARG_fv, _R4)           // MOVQ  fv, AX
	self.call_encoder(_F_encodeTypedPointer)  // CALL  encodeTypedPointer
	self.EmitCbzS("CBNZ", _ET, _LB_error)     // TESTQ ET, ET; JNZ   _error
	self.load_buffer()
}

func (self *Assembler) _asm_OP_byte(p *ir.Instr) {
	self._debug_tag(0x15) // 0x015
	self.check_size(1)
	self.write_imm_to_mem("MOVB", jit.Imm(p.I64()), _RP, _RL, 0, _T0, _T1)
	self.EmitAdd(_RL, _RL, jit.Imm(1))
	self.save_buffer()
}

func (self *Assembler) _asm_OP_text(p *ir.Instr) {
	self._debug_tag(0x16)        // 0x016
	self.check_size(len(p.Vs())) // SIZE ${len(p.Vs())}
	self.add_text(p.Vs())        // TEXT ${p.Vs()}
}

func (self *Assembler) _asm_OP_deref(_ *ir.Instr) {
	self._debug_tag(0x17)                       // 0x017
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _SP_p) // MOVQ (SP.p), SP.p
}

func (self *Assembler) _asm_OP_index(p *ir.Instr) {
	self._debug_tag(0x18)                    // 0x018
	self.Emit("MOVD", jit.Imm(p.I64()), _T0) // MOVQ $p.Vi(), AX
	self.Emit("ADD", _T0, _SP_p)             // ADDQ AX, SP.p
}

func (self *Assembler) _asm_OP_load(_ *ir.Instr) {
	self._debug_tag(0x19)                   // 0x019
	self.Emit("MOVD", jit.Ptr(_ST, 0), _T0) // MOVQ (ST), AX
	self.EmitLdrri(jit.Ptr, _ST, _T0, -24, _SP_x, _T1)
	self.EmitLdrri(jit.Ptr, _ST, _T0, -8, _SP_p, _T1)
	self.EmitLdrri(jit.Ptr, _ST, _T0, 0, _SP_q, _T1)
}

func (self *Assembler) _asm_OP_save(_ *ir.Instr) {
	self._debug_tag(0x1a) // 0x01a
	self.save_state()
}

func (self *Assembler) _asm_OP_drop(_ *ir.Instr) {
	self._debug_tag(0x1b) // 0x01b
	self.drop_state(vars.StateSize)
}

func (self *Assembler) _asm_OP_drop_2(_ *ir.Instr) {
	self._debug_tag(0x1c) // 0x01c
	self.drop_state(vars.StateSize * 2)
	self.EmitStrriWithInsn(jit.Ptr, "MOVD", _ZR, _ST, _T2, 56, _T0)
	self.EmitStrriWithInsn(jit.Ptr, "MOVD", _ZR, _ST, _T2, 64, _T0)
}

func (self *Assembler) _asm_OP_recurse(p *ir.Instr) {
	self._debug_tag(0x1d) // 0x01d
	self.prep_buffer()    // MOVE {buf}, (SP)
	vt, pv := p.Vp()
	self.Emit("MOVD", jit.Type(vt), _R1) // MOVQ $(type(p.Vt())), BX

	/* check for indirection */
	if !rt.UnpackType(vt).Indirect() {
		self.Emit("MOVD", _SP_p, _R2) // MOVQ SP.p, CX
	} else {
		self.Emit("MOVD", _SP_p, _VAR_vp) // MOVQ SP.p, VAR.vp
		self.Emit("MOVD", _SP, _T0)
		self.EmitAdd(_R2, _T0, jit.Imm(_VAR_vp_off)) // LEAQ VAR.vp, CX
	}

	/* call the encoder */
	self.Emit("MOVD", _ST, _R3)     // MOVQ  ST, DI
	self.Emit("MOVD", _ARG_fv, _R4) // MOVQ  $fv, SI
	if pv {
		mask := int64(-1 << alg.BitPointerValue) // As mask, uint64(1<<63) = int64(-1<<63)
		self.EmitTst(jit.Imm(mask), _R4)
		self.Emit("ORR", jit.Imm(int64(mask)), _R4)
	}

	self.call_encoder(_F_encodeTypedPointer) // CALL  encodeTypedPointer
	self.EmitCbzS("CBNZ", _ET, _LB_error)
	self.load_buffer()
}

func (self *Assembler) _asm_OP_is_nil(p *ir.Instr) {
	self._debug_tag(0x1e)                            // 0x01e
	self.EmitCmpqLdrPtr(jit.Ptr(_SP_p, 0), _ZR, _T0) // CMPQ (SP.p), $0 -> CMP xzr, (SP.p)
	self.Xjmp("BEQ", p.Vi())                         // JE   p.Vi()
}

func (self *Assembler) _asm_OP_is_nil_p1(p *ir.Instr) {
	self._debug_tag(0x1f)                            // 0x01f
	self.EmitCmpqLdrPtr(jit.Ptr(_SP_p, 8), _ZR, _T0) // CMPQ 8(Sp.p), $0 -> CMP xzr, 8(SP.p)
	self.Xjmp("BEQ", p.Vi())                         // JE   p.Vi()
}

func (self *Assembler) _asm_OP_is_zero_1(p *ir.Instr) {
	self._debug_tag(0x20) // 0x020
	self.Emit("MOVB", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzX("CBZ", _T0, p.Vi())

}

func (self *Assembler) _asm_OP_is_zero_2(p *ir.Instr) {
	self._debug_tag(0x21) // 0x021
	self.Emit("MOVH", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzX("CBZ", _T0, p.Vi())
}

func (self *Assembler) _asm_OP_is_zero_4(p *ir.Instr) {
	self._debug_tag(0x22) // 0x022
	self.Emit("MOVW", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzX("CBZ", _T0, p.Vi())
}

func (self *Assembler) _asm_OP_is_zero_8(p *ir.Instr) {
	self._debug_tag(0x23) // 0x023
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _T0)
	self.EmitCbzX("CBZ", _T0, p.Vi())
}

func (self *Assembler) _asm_OP_is_zero_map(p *ir.Instr) {
	self._debug_tag(0x24) // 0x024
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R0)
	self.EmitCbzX("CBZ", _R0, p.Vi())
	self.Emit("MOVD", jit.Ptr(_R0, 0), _T0)
	self.EmitCbzX("CBZ", _T0, p.Vi())
}

var (
	_F_is_zero      = jit.Func(prim.IsZero)
	_T_reflect_Type = rt.UnpackIface(reflect.Type(nil))
)

func (self *Assembler) _asm_OP_is_zero(p *ir.Instr) {
	self._debug_tag(0x25) // 0x025
	fv := p.VField()
	self.Emit("MOVD", _SP_p, _R0)
	self.Emit("MOVD", jit.ImmPtr(unsafe.Pointer(fv)), _R1)
	self.call_go(_F_is_zero)
	self.EmitCbzX("CBNZ", _R0, p.Vi())
}

func (self *Assembler) _asm_OP_cond_set(_ *ir.Instr) {
	self._debug_tag(0x26)                        // 0x026
	self.Emit("ORR", jit.Imm(1<<_S_cond), _SP_f) // EOR $(1<<_S_cond), SP.f
}

func (self *Assembler) _asm_OP_cond_testc(p *ir.Instr) {
	self._debug_tag(0x27) // 0x027
	self.EmitTst(jit.Imm(1<<_S_cond), _SP_f)
	self.Emit("BIC", jit.Imm(1<<_S_cond), _SP_f)
	self.Xjmp("BNE", p.Vi())
}

func (self *Assembler) _asm_OP_goto(p *ir.Instr) {
	self._debug_tag(0x28) // 0x028
	self.Xjmp("JMP", p.Vi())
}

func (self *Assembler) _asm_OP_map_iter(p *ir.Instr) {
	self._debug_tag(0x29)                     // 0x029
	self.Emit("MOVD", jit.Type(p.Vt()), _R0)  // MOVD    $p.Vt(), AX
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _R1) // MOVD    (SP.p), BX
	self.Emit("MOVD", _ARG_fv, _R2)           // MOVD    fv, CX
	self.call_go(_F_iteratorStart)            // CALL_GO iteratorStart
	self.Emit("MOVD", _R0, _SP_q)             // MOVD    AX, SP.q
	self.Emit("MOVD", _R1, _ET)               // MOVD    32(SP), ET
	self.Emit("MOVD", _R2, _EP)               // MOVD    40(SP), EP
	self.EmitCbzS("CBNZ", _ET, _LB_error)
}

func (self *Assembler) _asm_OP_map_stop(_ *ir.Instr) {
	self._debug_tag(0x2a)          // 0x02a
	self.Emit("MOVD", _SP_q, _R0)  // MOVQ    SP.q, AX
	self.call_go(_F_iteratorStop)  // CALL_GO iteratorStop
	self.Emit("EOR", _SP_q, _SP_q) // XORL    SP.q, SP.q
}

func (self *Assembler) _asm_OP_map_check_key(p *ir.Instr) {
	self._debug_tag(0x2b)                       // 0x02b
	self.Emit("MOVD", jit.Ptr(_SP_q, 0), _SP_p) // MOVQ    (SP.q), SP.p
	self.EmitCbzX("CBZ", _SP_p, p.Vi())
}

func (self *Assembler) _asm_OP_map_write_key(p *ir.Instr) {
	self._debug_tag(0x2c) // 0x02c
	self.Emit("MOVD", _ARG_fv, _T0)
	self.EmitTbzS("TBZ", jit.Imm(alg.BitSortMapKeys), _T0, "_unordered_key_{n}")
	self.encode_string(false)       // STR $false
	self.Xjmp("B", p.Vi())          // JMP ${p.Vi()}
	self.Link("_unordered_key_{n}") // _unordered_key_{n}:
}

func (self *Assembler) _asm_OP_map_value_next(_ *ir.Instr) {
	self._debug_tag(0x2d)                       // 0x02d
	self.Emit("MOVD", jit.Ptr(_SP_q, 8), _SP_p) // MOVQ    8(SP.q), SP.p
	self.Emit("MOVD", _SP_q, _R0)               // MOVQ    SP.q, AX
	self.call_go(_F_iteratorNext)               // CALL_GO iteratorNext
}

func (self *Assembler) _asm_OP_slice_len(_ *ir.Instr) {
	self._debug_tag(0x2e)                        // 0x02e
	self.Emit("MOVD", jit.Ptr(_SP_p, 8), _SP_x)  // MOVD 8(SP.p), SP.x
	self.Emit("MOVD", jit.Ptr(_SP_p, 0), _SP_p)  // MOVD (SP.p), SP.p
	self.Emit("ORR", jit.Imm(1<<_S_init), _SP_f) // ORR  $(1<<_S_init), SP.f
}

func (self *Assembler) _asm_OP_slice_next(p *ir.Instr) {
	self._debug_tag(0x2f)                        // 0x02e
	self.EmitCbzX("CBZ", _SP_x, p.Vi())          // CBZ     SP.x, p.Vi()
	self.Emit("SUB", jit.Imm(1), _SP_x)          // SUB
	self.EmitTst(jit.Imm(1<<_S_init), _SP_f)     // TST     $(1<<_S_init), SP.f
	self.Emit("BIC", jit.Imm(1<<_S_init), _SP_f) // BIC     $(1<<_S_init), SP.f
	self.Emit("MOVD", jit.Imm(int64(p.Vlen())), _T0)
	self.EmitAdd(_R0, _T0, _SP_p) // ADD     $(p.vlen()), SP.p
	self.Sjmp("BNE", "_slice_done_{n}")
	self.Emit("MOVD", _R0, _SP_p)
	self.Link("_slice_done_{n}")
}

func (self *Assembler) _asm_OP_marshal(p *ir.Instr) {
	self._debug_tag(0x30) // 0x030
	self.call_marshaler(_F_encodeJsonMarshaler, _T_json_Marshaler, p.Vt())
}

func (self *Assembler) _asm_OP_marshal_p(p *ir.Instr) {
	self._debug_tag(0x31) // 0x031
	if p.Vk() != reflect.Ptr {
		panic("marshal_p: invalid type")
	} else {
		self.call_marshaler_v(_F_encodeJsonMarshaler, _T_json_Marshaler, p.Vt(), false)
	}
}

func (self *Assembler) _asm_OP_marshal_text(p *ir.Instr) {
	self._debug_tag(0x32) // 0x032
	self.call_marshaler(_F_encodeTextMarshaler, _T_encoding_TextMarshaler, p.Vt())
}

func (self *Assembler) _asm_OP_marshal_text_p(p *ir.Instr) {
	self._debug_tag(0x33) // 0x033
	if p.Vk() != reflect.Ptr {
		panic("marshal_text_p: invalid type")
	} else {
		self.call_marshaler_v(_F_encodeTextMarshaler, _T_encoding_TextMarshaler, p.Vt(), false)
	}
}

var _F_error_unsupported = jit.Func(vars.Error_unsuppoted)

func (self *Assembler) _asm_OP_unsupported(i *ir.Instr) {
	self._debug_tag(0x34) // 0x034
	typ := int64(uintptr(unsafe.Pointer(i.GoType())))
	self.Emit("MOVD", jit.Imm(typ), _R0)
	self.call_go(_F_error_unsupported)
	self.Sjmp("JMP", _LB_error)
}

var (
	_debug_label_counter = 0
)

func (self *Assembler) _debug_tag(i int64) {
	if !_DBG_BRK_ALL {
		return
	}
	var label string = "_brk_" + strconv.Itoa(_debug_label_counter)
	if !_DBG_BRK_EMIT {
		self.Sjmp("B", label)
	}
	self.EmitBrk(jit.Imm(i))
	if !_DBG_BRK_EMIT {
		self.Link(label)
	}
	_debug_label_counter += 1
}
