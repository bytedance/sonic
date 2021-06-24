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
    `fmt`
    `reflect`
    `strconv`
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/cpu`
    `github.com/bytedance/sonic/internal/jit`
    `github.com/twitchyliquid64/golang-asm/obj`
    `github.com/twitchyliquid64/golang-asm/obj/x86`

    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/rt`
)

/** Register Allocations
 *
 *  State Registers:
 *
 *      %rbx : stack base
 *      %rdi : result pointer
 *      %rsi : result length
 *      %rdx : result capacity
 *      %r12 : sp->p
 *      %r13 : sp->q
 *      %r14 : sp->x
 *      %r15 : sp->f
 *
 *  Error Registers:
 *
 *      %r10 : error type register
 *      %r11 : error pointer register
 */

/** Function Prototype & Stack Map
 *
 *  func (buf *[]byte, p unsafe.Pointer, sb *_Stack) (err error)
 *
 *  buf    :   (FP)
 *  p      :  8(FP)
 *  sb     : 16(FP)
 *  err.vt : 24(FP)
 *  err.vp : 32(FP)
 */

const (
    _S_cond = iota
    _S_init
)

const (
    _FP_args  = 40  // 40 bytes for passing arguments to this function
    _FP_fargs = 64  // 64 bytes for passing arguments to other Go functions
    _FP_saves = 64  // 64 bytes for saving the registers before CALL instructions
)

const (
    _FP_offs = _FP_fargs + _FP_saves
    _FP_size = _FP_offs + 8     // 8 bytes for the parent frame pointer
    _FP_base = _FP_size + 8     // 8 bytes for the return address
)

const (
    _FM_exp32 = 0x7f800000
    _FM_exp64 = 0x7ff0000000000000
)

const (
    _IM_null = 0x6c6c756e           // 'null'
    _IM_true = 0x65757274           // 'true'
    _IM_fals = 0x736c6166           // 'fals' ('false' without the 'e')
    _IM_open = 0x00225c22           // '"\"∅'
    _IM_mulv = -0x5555555555555555
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
)

var (
    _AX = jit.Reg("AX")
    _CX = jit.Reg("CX")
    _DX = jit.Reg("DX")
    _DI = jit.Reg("DI")
    _SI = jit.Reg("SI")
    _BP = jit.Reg("BP")
    _SP = jit.Reg("SP")
    _R8 = jit.Reg("R8")
)

var (
    _X0 = jit.Reg("X0")
    _Y0 = jit.Reg("Y0")
    _Y1 = jit.Reg("Y1")
    _Y2 = jit.Reg("Y2")
    _Y3 = jit.Reg("Y3")
    _Y4 = jit.Reg("Y4")
    _Y5 = jit.Reg("Y5")
)

var (
	_ST = jit.Reg("BX")
	_RP = jit.Reg("DI")
	_RL = jit.Reg("SI")
	_RC = jit.Reg("DX")
)

var (
    _LR = jit.Reg("R9")
	_ET = jit.Reg("R10")
	_EP = jit.Reg("R11")
)

var (
    _SP_p = jit.Reg("R12")
    _SP_q = jit.Reg("R13")
    _SP_x = jit.Reg("R14")
    _SP_f = jit.Reg("R15")
)

var (
    _ARG_rb = jit.Ptr(_SP, _FP_base)
    _ARG_vp = jit.Ptr(_SP, _FP_base + 8)
    _ARG_sb = jit.Ptr(_SP, _FP_base + 16)
)

var (
    _RET_et = jit.Ptr(_SP, _FP_base + 24)
    _RET_ep = jit.Ptr(_SP, _FP_base + 32)
)

var (
    _REG_ffi = []obj.Addr{_RP, _RL, _RC}
    _REG_enc = []obj.Addr{_ST, _SP_x, _SP_f, _SP_p, _SP_q}
    _REG_jsr = []obj.Addr{_ST, _SP_x, _SP_f, _SP_p, _SP_q, _LR}
    _REG_all = []obj.Addr{_ST, _SP_x, _SP_f, _SP_p, _SP_q, _RP, _RL, _RC}
)

type _Assembler struct {
    jit.BaseAssembler
    p *_Program
    x int
}

func newAssembler(p *_Program) *_Assembler {
    return new(_Assembler).Init(p)
}

/** Assembler Interface **/

func (self *_Assembler) Load() _Encoder {
    return ptoenc(self.BaseAssembler.Load("json_encoder", _FP_size, _FP_args))
}

func (self *_Assembler) Init(p *_Program) *_Assembler {
    self.p = p
    self.BaseAssembler.Init(self.compile)
    return self
}

func (self *_Assembler) compile() {
    self.prologue()
    self.instrs()
    self.epilogue()
    self.builtins()
}

/** Assembler Stages **/

var _OpFuncTab = [256]func(*_Assembler, *_Instr) {
    _OP_null           : (*_Assembler)._asm_OP_null,
    _OP_bool           : (*_Assembler)._asm_OP_bool,
    _OP_i8             : (*_Assembler)._asm_OP_i8,
    _OP_i16            : (*_Assembler)._asm_OP_i16,
    _OP_i32            : (*_Assembler)._asm_OP_i32,
    _OP_i64            : (*_Assembler)._asm_OP_i64,
    _OP_u8             : (*_Assembler)._asm_OP_u8,
    _OP_u16            : (*_Assembler)._asm_OP_u16,
    _OP_u32            : (*_Assembler)._asm_OP_u32,
    _OP_u64            : (*_Assembler)._asm_OP_u64,
    _OP_f32            : (*_Assembler)._asm_OP_f32,
    _OP_f64            : (*_Assembler)._asm_OP_f64,
    _OP_str            : (*_Assembler)._asm_OP_str,
    _OP_bin            : (*_Assembler)._asm_OP_bin,
    _OP_quote          : (*_Assembler)._asm_OP_quote,
    _OP_number         : (*_Assembler)._asm_OP_number,
    _OP_eface          : (*_Assembler)._asm_OP_eface,
    _OP_iface          : (*_Assembler)._asm_OP_iface,
    _OP_byte           : (*_Assembler)._asm_OP_byte,
    _OP_text           : (*_Assembler)._asm_OP_text,
    _OP_deref          : (*_Assembler)._asm_OP_deref,
    _OP_index          : (*_Assembler)._asm_OP_index,
    _OP_load           : (*_Assembler)._asm_OP_load,
    _OP_save           : (*_Assembler)._asm_OP_save,
    _OP_drop           : (*_Assembler)._asm_OP_drop,
    _OP_drop_2         : (*_Assembler)._asm_OP_drop_2,
    _OP_recurse        : (*_Assembler)._asm_OP_recurse,
    _OP_is_nil         : (*_Assembler)._asm_OP_is_nil,
    _OP_is_nil_p1      : (*_Assembler)._asm_OP_is_nil_p1,
    _OP_is_zero_1      : (*_Assembler)._asm_OP_is_zero_1,
    _OP_is_zero_2      : (*_Assembler)._asm_OP_is_zero_2,
    _OP_is_zero_4      : (*_Assembler)._asm_OP_is_zero_4,
    _OP_is_zero_8      : (*_Assembler)._asm_OP_is_zero_8,
    _OP_is_zero_map    : (*_Assembler)._asm_OP_is_zero_map,
    _OP_is_zero_mem    : (*_Assembler)._asm_OP_is_zero_mem,
    _OP_is_zero_safe   : (*_Assembler)._asm_OP_is_zero_safe,
    _OP_goto           : (*_Assembler)._asm_OP_goto,
    _OP_map_iter       : (*_Assembler)._asm_OP_map_iter,
    _OP_map_check_key  : (*_Assembler)._asm_OP_map_check_key,
    _OP_map_value_next : (*_Assembler)._asm_OP_map_value_next,
    _OP_slice_len      : (*_Assembler)._asm_OP_slice_len,
    _OP_slice_next     : (*_Assembler)._asm_OP_slice_next,
    _OP_marshal        : (*_Assembler)._asm_OP_marshal,
    _OP_marshal_p      : (*_Assembler)._asm_OP_marshal_p,
    _OP_marshal_text   : (*_Assembler)._asm_OP_marshal_text,
    _OP_marshal_text_p : (*_Assembler)._asm_OP_marshal_text_p,
    _OP_cond_set       : (*_Assembler)._asm_OP_cond_set,
    _OP_cond_testc     : (*_Assembler)._asm_OP_cond_testc,
}

func (self *_Assembler) instr(v *_Instr) {
    if fn := _OpFuncTab[v.op()]; fn != nil {
        fn(self, v)
    } else {
        panic(fmt.Sprintf("invalid opcode: %d", v.op()))
    }
}

func (self *_Assembler) instrs() {
    for i, v := range self.p.ins {
        self.Mark(i)
        self.instr(&v)
    }
}

func (self *_Assembler) builtins() {
    self.more_space()
    self.error_too_deep()
    self.error_invalid_number()
    self.error_nan_or_infinite()
}

func (self *_Assembler) epilogue() {
    self.Mark(len(self.p.ins))
    self.Emit("XORL", _ET, _ET)
    self.Emit("XORL", _EP, _EP)
    self.Link(_LB_error)
    self.Emit("MOVQ", _ARG_rb, _AX)                 // MOVQ rb<>+0(FP), AX
    self.Emit("MOVQ", _RL, jit.Ptr(_AX, 8))         // MOVQ RL, 8(AX)
    self.Emit("MOVQ", _ET, _RET_et)                 // MOVQ ET, et<>+24(FP)
    self.Emit("MOVQ", _EP, _RET_ep)                 // MOVQ EP, ep<>+32(FP)
    self.Emit("MOVQ", jit.Ptr(_SP, _FP_offs), _BP)  // MOVQ _FP_offs(SP), BP
    self.Emit("ADDQ", jit.Imm(_FP_size), _SP)       // ADDQ $_FP_size, SP
    self.Emit("RET")                                // RET
}

func (self *_Assembler) prologue() {
    self.Emit("SUBQ", jit.Imm(_FP_size), _SP)       // SUBQ $_FP_size, SP
    self.Emit("MOVQ", _BP, jit.Ptr(_SP, _FP_offs))  // MOVQ BP, _FP_offs(SP)
    self.Emit("LEAQ", jit.Ptr(_SP, _FP_offs), _BP)  // LEAQ _FP_offs(SP), BP
    self.load_buffer()                              // LOAD {buf}
    self.Emit("MOVQ", _ARG_vp, _SP_p)               // MOVQ vp<>+8(FP), SP.p
    self.Emit("MOVQ", _ARG_sb, _ST)                 // MOVQ sb<>+16(FP), ST
    self.Emit("XORL", _SP_x, _SP_x)                 // XORL SP.x, SP.x
    self.Emit("XORL", _SP_f, _SP_f)                 // XORL SP.f, SP.f
    self.Emit("XORL", _SP_q, _SP_q)                 // XORL SP.q, SP.q
}

/** Assembler Inline Functions **/

func (self *_Assembler) xsave(reg ...obj.Addr) {
    for i, v := range reg {
        if i > _FP_saves / 8 - 1 {
            panic("too many registers to save")
        } else {
            self.Emit("MOVQ", v, jit.Ptr(_SP, _FP_fargs + int64(i) * 8))
        }
    }
}

func (self *_Assembler) xload(reg ...obj.Addr) {
    for i, v := range reg {
        if i > _FP_saves / 8 - 1 {
            panic("too many registers to load")
        } else {
            self.Emit("MOVQ", jit.Ptr(_SP, _FP_fargs + int64(i) * 8), v)
        }
    }
}

func (self *_Assembler) rbuf_di() {
    if _RP.Reg != x86.REG_DI {
        panic("register allocation messed up: RP != DI")
    } else {
        self.Emit("ADDQ", _RL, _RP)
    }
}

func (self *_Assembler) store_int(nd int, fn obj.Addr, ins string) {
    self.check_size(nd)
    self.save_c()                           // SAVE   $C_regs
    self.rbuf_di()                          // MOVQ   RP, DI
    self.Emit(ins, jit.Ptr(_SP_p, 0), _SI)  // $ins   (SP.p), SI
    self.call_c(fn)                         // CALL_C $fn
    self.Emit("ADDQ", _AX, _RL)             // ADDQ   AX, RL
}

func (self *_Assembler) store_str(s string) {
    i := 0
    m := rt.Str2Mem(s)

    /* 8-byte stores */
    for i <= len(m) - 8 {
        self.Emit("MOVQ", jit.Imm(rt.Get64(m[i:])), _AX)        // MOVQ $s[i:], AX
        self.Emit("MOVQ", _AX, jit.Sib(_RP, _RL, 1, int64(i)))  // MOVQ AX, i(RP)(RL)
        i += 8
    }

    /* 4-byte stores */
    if i <= len(m) - 4 {
        self.Emit("MOVL", jit.Imm(int64(rt.Get32(m[i:]))), jit.Sib(_RP, _RL, 1, int64(i)))  // MOVL $s[i:], i(RP)(RL)
        i += 4
    }

    /* 2-byte stores */
    if i <= len(m) - 2 {
        self.Emit("MOVW", jit.Imm(int64(rt.Get16(m[i:]))), jit.Sib(_RP, _RL, 1, int64(i)))  // MOVW $s[i:], i(RP)(RL)
        i += 2
    }

    /* last byte */
    if i < len(m) {
        self.Emit("MOVB", jit.Imm(int64(m[i])), jit.Sib(_RP, _RL, 1, int64(i)))     // MOVB $s[i:], i(RP)(RL)
    }
}

func (self *_Assembler) check_size(n int) {
    self.check_size_rl(jit.Ptr(_RL, int64(n)))
}

func (self *_Assembler) check_size_r(r obj.Addr, d int) {
    self.check_size_rl(jit.Sib(_RL, r, 1, int64(d)))
}

func (self *_Assembler) check_size_rl(v obj.Addr) {
    idx := self.x
    key := _LB_more_space_return + strconv.Itoa(idx)

    /* the following code relies on LR == R9 to work */
    if _LR.Reg != x86.REG_R9 {
        panic("register allocation messed up: LR != R9")
    }

    /* check for buffer capacity */
    self.x++
    self.Emit("LEAQ", v, _AX)           // LEAQ $v, AX
    self.Emit("CMPQ", _AX, _RC)         // CMPQ AX, RC
    self.Sjmp("JBE" , key)              // JBE  _more_space_return_{n}
    self.Byte(0x4c, 0x8d, 0x0d)         // LEAQ ?(PC), R9
    self.Sref(key, 4)                   // .... &key
    self.Sjmp("JMP" , _LB_more_space)   // JMP  _more_space
    self.Link(key)                      // _more_space_return_{n}:
}

/** State Stack Helpers **/

const (
    _StateSize  = int64(unsafe.Sizeof(_State{}))
    _StackLimit = _MaxStack * _StateSize
)

func (self *_Assembler) save_state() {
    self.Emit("MOVQ", jit.Ptr(_ST, 0), _AX)             // MOVQ (ST), AX
    self.Emit("LEAQ", jit.Ptr(_AX, _StateSize), _R8)    // LEAQ _StateSize(AX), R8
    self.Emit("CMPQ", _R8, jit.Imm(_StackLimit))        // CMPQ R8, $_StackLimit
    self.Sjmp("JA"  , _LB_error_too_deep)               // JA   _error_too_deep
    self.Emit("MOVQ", _SP_x, jit.Sib(_ST, _AX, 1, 8))   // MOVQ SP.x, 8(ST)(AX)
    self.Emit("MOVQ", _SP_f, jit.Sib(_ST, _AX, 1, 16))  // MOVQ SP.f, 16(ST)(AX)
    self.Emit("MOVQ", _SP_p, jit.Sib(_ST, _AX, 1, 24))  // MOVQ SP.p, 24(ST)(AX)
    self.Emit("MOVQ", _SP_q, jit.Sib(_ST, _AX, 1, 32))  // MOVQ SP.q, 32(ST)(AX)
    self.Emit("MOVQ", _R8, jit.Ptr(_ST, 0))             // MOVQ R8, (ST)
}

func (self *_Assembler) drop_state(decr int64) {
    self.Emit("MOVQ" , jit.Ptr(_ST, 0), _AX)                // MOVQ  (ST), AX
    self.Emit("SUBQ" , jit.Imm(decr), _AX)                  // SUBQ  $decr, AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_ST, 0))                // MOVQ  AX, (ST)
    self.Emit("MOVQ" , jit.Sib(_ST, _AX, 1, 8), _SP_x)      // MOVQ  8(ST)(AX), SP.x
    self.Emit("MOVQ" , jit.Sib(_ST, _AX, 1, 16), _SP_f)     // MOVQ  16(ST)(AX), SP.f
    self.Emit("MOVQ" , jit.Sib(_ST, _AX, 1, 24), _SP_p)     // MOVQ  24(ST)(AX), SP.p
    self.Emit("MOVQ" , jit.Sib(_ST, _AX, 1, 32), _SP_q)     // MOVQ  32(ST)(AX), SP.q
    self.Emit("PXOR" , _X0, _X0)                            // PXOR  X0, X0
    self.Emit("MOVOU", _X0, jit.Sib(_ST, _AX, 1, 8))        // MOVOU X0, 8(ST)(AX)
    self.Emit("MOVOU", _X0, jit.Sib(_ST, _AX, 1, 24))       // MOVOU X0, 24(ST)(AX)
}

/** Buffer Helpers **/

func (self *_Assembler) add_char(ch byte) {
    self.Emit("MOVB", jit.Imm(int64(ch)), jit.Sib(_RP, _RL, 1, 0))  // MOVB $ch, (RP)(RL)
    self.Emit("ADDQ", jit.Imm(1), _RL)                              // ADDQ $1, RL
}

func (self *_Assembler) add_long(ch uint32, n int64) {
    self.Emit("MOVL", jit.Imm(int64(ch)), jit.Sib(_RP, _RL, 1, 0))  // MOVL $ch, (RP)(RL)
    self.Emit("ADDQ", jit.Imm(n), _RL)                              // ADDQ $n, RL
}

func (self *_Assembler) prep_buffer() {
    self.Emit("MOVQ", _ARG_rb, _AX)             // MOVQ rb<>+0(FP), AX
    self.Emit("MOVQ", _RL, jit.Ptr(_AX, 8))     // MOVQ RL, 8(AX)
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))     // MOVQ AX, (SP)
}

func (self *_Assembler) prep_buffer_c() {
    self.Emit("MOVQ", _ARG_rb, _DI)             // MOVQ rb<>+0(FP), DI
    self.Emit("MOVQ", _RL, jit.Ptr(_DI, 8))     // MOVQ RL, 8(DI)
}

func (self *_Assembler) save_buffer() {
    self.Emit("MOVQ", _ARG_rb, _AX)             // MOVQ rb<>+0(FP), AX
    self.Emit("MOVQ", _RP, jit.Ptr(_AX,  0))    // MOVQ RP, (AX)
    self.Emit("MOVQ", _RL, jit.Ptr(_AX,  8))    // MOVQ RL, 8(AX)
    self.Emit("MOVQ", _RC, jit.Ptr(_AX, 16))    // MOVQ RC, 16(AX)
}

func (self *_Assembler) load_buffer() {
    self.Emit("MOVQ", _ARG_rb, _AX)             // MOVQ rb<>+0(FP), AX
    self.Emit("MOVQ", jit.Ptr(_AX,  0), _RP)    // MOVQ (AX), RP
    self.Emit("MOVQ", jit.Ptr(_AX,  8), _RL)    // MOVQ 8(AX), RL
    self.Emit("MOVQ", jit.Ptr(_AX, 16), _RC)    // MOVQ 16(AX), RC
}

/** Function Interface Helpers **/

var (
    _F_assertI2I = jit.Func(assertI2I)
)

func (self *_Assembler) call(pc obj.Addr) {
    self.Emit("MOVQ", pc, _AX)  // MOVQ $pc, AX
    self.Rjmp("CALL", _AX)      // CALL AX
}

func (self *_Assembler) save_c() {
    self.xsave(_REG_ffi...)     // SAVE $REG_ffi
}

func (self *_Assembler) call_c(pc obj.Addr) {
    self.call(pc)               // CALL $pc
    self.xload(_REG_ffi...)     // LOAD $REG_ffi
}

func (self *_Assembler) call_go(pc obj.Addr) {
    self.xsave(_REG_all...)     // SAVE $REG_all
    self.call(pc)               // CALL $pc
    self.xload(_REG_all...)     // LOAD $REG_all
}

func (self *_Assembler) call_encoder(pc obj.Addr) {
    self.xsave(_REG_enc...)     // SAVE $REG_enc
    self.call(pc)               // CALL $pc
    self.xload(_REG_enc...)     // LOAD $REG_enc
    self.load_buffer()          // LOAD {buf}
}

func (self *_Assembler) call_marshaler(fn obj.Addr, it *rt.GoType, vt reflect.Type) {
    switch vt.Kind() {
        case reflect.Interface : self.call_marshaler_i(fn, it)
        case reflect.Ptr       : self.call_marshaler_v(fn, it, vt, true)
        default                : self.call_marshaler_v(fn, it, vt, false)
    }
}

func (self *_Assembler) call_marshaler_i(fn obj.Addr, it *rt.GoType) {
    self.Emit("MOVQ" , jit.Gtype(it), _AX)                          // MOVQ    $it, AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))                        // MOVQ    AX, (SP)
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _AX)                      // MOVQ    (SP.p), AX
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 8), _CX)                      // MOVQ    8(SP.p), CX
    self.Emit("TESTQ", _AX, _AX)                                    // TESTQ   AX, AX
    self.Sjmp("JZ"   , "_null_{n}")                                 // JZ      _null_{n}
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 8))                        // MOVQ    AX, 8(SP)
    self.Emit("MOVQ" , _CX, jit.Ptr(_SP, 16))                       // MOVQ    CX, 16(SP)
    self.call_go(_F_assertI2I)                                      // CALL_GO assertI2I
    self.prep_buffer()                                              // MOVE    {buf}, (SP)
    self.Emit("MOVOU", jit.Ptr(_SP, 24), _X0)                       // MOVOU   24(SP), X0
    self.Emit("MOVOU", _X0, jit.Ptr(_SP, 8))                        // MOVOU   X0, 8(SP)
    self.call_encoder(fn)                                           // CALL    $fn
    self.Emit("MOVQ" , jit.Ptr(_SP, 24), _ET)                       // MOVQ    24(SP), ET
    self.Emit("MOVQ" , jit.Ptr(_SP, 32), _EP)                       // MOVQ    32(SP), EP
    self.Emit("TESTQ", _ET, _ET)                                    // TESTQ   ET, ET
    self.Sjmp("JNZ"  , _LB_error)                                   // JNZ     _error
    self.Sjmp("JMP"  , "_done_{n}")                                 // JMP     _done_{n}
    self.Link("_null_{n}")                                          // _null_{n}:
    self.check_size(4)                                              // SIZE    $4
    self.Emit("MOVL", jit.Imm(_IM_null), jit.Sib(_RP, _RL, 1, 0))   // MOVL    $'null', (RP)(RL*1)
    self.Emit("ADDQ", jit.Imm(4), _RL)                              // ADDQ    $4, RL
    self.Link("_done_{n}")                                          // _done_{n}:
}

func (self *_Assembler) call_marshaler_v(fn obj.Addr, it *rt.GoType, vt reflect.Type, deref bool) {
    self.prep_buffer()                          // MOVE {buf}, (SP)
    self.Emit("MOVQ", jit.Itab(it, vt), _AX)    // MOVQ $(itab(it, vt)), AX
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 8))     // MOVQ AX, 8(SP)

    /* dereference the pointer if needed */
    if !deref {
        self.Emit("MOVQ", _SP_p, jit.Ptr(_SP, 16))  // MOVQ SP.p, 16(SP)
    } else {
        self.Emit("MOVQ", jit.Ptr(_SP_p, 0), _AX)   // MOVQ (SP.p), AX
        self.Emit("MOVQ", _AX, jit.Ptr(_SP, 16))    // MOVQ AX, 16(SP)
    }

    /* call the encoder, and perform error checks */
    self.call_encoder(fn)                       // CALL  $fn
    self.Emit("MOVQ" , jit.Ptr(_SP, 24), _ET)   // MOVQ  24(SP), ET
    self.Emit("MOVQ" , jit.Ptr(_SP, 32), _EP)   // MOVQ  32(SP), EP
    self.Emit("TESTQ", _ET, _ET)                // TESTQ ET, ET
    self.Sjmp("JNZ"  , _LB_error)               // JNZ   _error
}

/** Builtin: _more_space **/

var (
    _T_byte      = jit.Type(byteType)
    _F_growslice = jit.Func(growslice)
)

func (self *_Assembler) more_space() {
    self.Link(_LB_more_space)
    self.Emit("MOVQ", _T_byte, jit.Ptr(_SP, 0))     // MOVQ $_T_byte, (SP)
    self.Emit("MOVQ", _RP, jit.Ptr(_SP, 8))         // MOVQ RP, 8(SP)
    self.Emit("MOVQ", _RL, jit.Ptr(_SP, 16))        // MOVQ RL, 16(SP)
    self.Emit("MOVQ", _RC, jit.Ptr(_SP, 24))        // MOVQ RC, 24(SP)
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 32))        // MOVQ AX, 32(SP)
    self.xsave(_REG_jsr...)                         // SAVE $REG_jsr
    self.call(_F_growslice)                         // CALL $pc
    self.xload(_REG_jsr...)                         // LOAD $REG_jsr
    self.Emit("MOVQ", jit.Ptr(_SP, 40), _RP)        // MOVQ 40(SP), RP
    self.Emit("MOVQ", jit.Ptr(_SP, 48), _RL)        // MOVQ 48(SP), RL
    self.Emit("MOVQ", jit.Ptr(_SP, 56), _RC)        // MOVQ 56(SP), RC
    self.save_buffer()                              // SAVE {buf}
    self.Rjmp("JMP" , _LR)                          // JMP  LR
}

/** Builtin Errors **/

var (
    _V_ERR_too_deep               = jit.Imm(int64(uintptr(unsafe.Pointer(_ERR_too_deep))))
    _V_ERR_nan_or_infinite        = jit.Imm(int64(uintptr(unsafe.Pointer(_ERR_nan_or_infinite))))
    _I_json_UnsupportedValueError = jit.Itab(rt.UnpackType(errorType), jsonUnsupportedValueType)
)

func (self *_Assembler) error_too_deep() {
    self.Link(_LB_error_too_deep)
    self.Emit("MOVQ", _V_ERR_too_deep, _EP)                 // MOVQ $_V_ERR_too_deep, EP
    self.Emit("MOVQ", _I_json_UnsupportedValueError, _ET)   // MOVQ $_I_json_UnsupportedValuError, ET
    self.Sjmp("JMP" , _LB_error)                            // JMP  _error
}

func (self *_Assembler) error_invalid_number() {
    self.Link(_LB_error_invalid_number)
    self.call_go(_F_error_number)               // CALL_GO error_number
    self.Emit("MOVQ", jit.Ptr(_SP, 16), _ET)    // MOVQ    16(SP), ET
    self.Emit("MOVQ", jit.Ptr(_SP, 24), _EP)    // MOVQ    24(SP), EP
    self.Sjmp("JMP" , _LB_error)                // JMP     _error
}

func (self *_Assembler) error_nan_or_infinite()  {
    self.Link(_LB_error_nan_or_infinite)
    self.Emit("MOVQ", _V_ERR_nan_or_infinite, _EP)          // MOVQ $_V_ERR_nan_or_infinite, EP
    self.Emit("MOVQ", _I_json_UnsupportedValueError, _ET)   // MOVQ $_I_json_UnsupportedValuError, ET
    self.Sjmp("JMP" , _LB_error)                            // JMP  _error
}

/** String Encoding Routine **/

func (self *_Assembler) open_quote(doubleQuote bool) {
    if !doubleQuote {
        self.check_size_r(_AX, 2)   // SIZE $2
        self.add_char('"')          // CHAR $'"'
    } else {
        self.check_size_r(_AX, 6)   // SIZE $6
        self.add_long(_IM_open, 3)  // TEXT $`"\"`
    }
}

func (self *_Assembler) close_quote(doubleQuote bool) {
    if !doubleQuote {
        self.check_size(1)          // SIZE $1
        self.Link("_str_end_{n}")   // _str_end_{n}:
        self.add_char('"')          // CHAR $'"'
    } else {
        self.check_size(3)                  // SIZE $3
        self.Link("_str_end_{n}")           // _str_end_{n}:
        self.store_str(`\""`)               // TEXT $`\""`
        self.Emit("ADDQ", jit.Imm(3), _RL)  // ADDQ $3, RL
    }
}

func (self *_Assembler) encode_string(fn obj.Addr, doubleQuote bool) {
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 8), _AX)          // MOVQ    8(SP.p), AX
    self.open_quote(doubleQuote)                        // QOPEN   $doubleQuote
    self.Emit("CMPQ" , jit.Ptr(_SP_p, 8), jit.Imm(0))   // CMPQ    8(SP.p), $0
    self.Sjmp("JE"   , "_str_end_{n}")                  // JE      _str_end_{n}
    self.save_c()                                       // SAVE    $REG_ffi
    self.Emit("MOVQ" , _SP_p, _DI)                      // MOVQ    SP.p, DI
    self.Emit("XORL" , _SI, _SI)                        // XORL    SI, SI
    self.call_c(_F_lquote)                              // CALL    lquote
    self.Emit("CMPQ" , _AX, jit.Ptr(_SP_p, 8))          // CMPQ    AX, 8(SP.p)
    self.Sjmp("JNE"  , "_str_quote_{n}")                // JNE     _str_quote_{n}
    self.Emit("LEAQ" , jit.Sib(_RP, _RL, 1, 0), _AX)    // LEAQ    (RP)(RL), AX
    self.Emit("ADDQ" , jit.Ptr(_SP_p, 8), _RL)          // ADDQ    8(SP.p), RL
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))            // MOVQ    AX, 0(SP)
    self.Emit("MOVOU", jit.Ptr(_SP_p, 0), _X0)          // MOVOU   (SP.p), X0
    self.Emit("MOVOU", _X0, jit.Ptr(_SP, 8))            // MOVOU   X0, 8(SP)
    self.call_go(_F_memmove)                            // CALL_GO memmove
    self.Sjmp("JMP"  , "_str_end_{n}")                  // JMP     _str_end_{n}
    self.Link("_str_quote_{n}")                         // _str_quote_{n}:
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 8))            // MOVQ    AX, 8(SP)
    self.prep_buffer()                                  // MOVE    {buf}, (SP)
    self.Emit("MOVOU", jit.Ptr(_SP_p, 0), _X0)          // MOVOU   (SP.p), X0
    self.Emit("MOVOU", _X0, jit.Ptr(_SP, 16))           // MOVOU   X0, 16(SP)
    self.call_encoder(fn)                               // CALL    $fn
    self.close_quote(doubleQuote)                       // QCLOSE  $doubleQuote
}

/** Zero Value Check Routine **/

func (self *_Assembler) check_zero(nb int, dest int) {
    i := int64(0)
    e := int64(nb)

    /* special case: zero-sized value, always empty */
    if e == 0 {
        return
    }

    /* default instructions for AVX2 */
    vclear := func(v obj.Addr)       { self.Emit("VPXOR"   , v, v, v) }
    vset1a := func(a, b obj.Addr)    { self.Emit("VPCMPEQB", a, a, b) }
    vandpb := func(b, a, r obj.Addr) { self.Emit("VPAND"   , b, a, r) }
    vcmpeq := func(b, a, r obj.Addr) { self.Emit("VPCMPEQB", b, a, r) }

    /* fall-back instructions for AVX */
    if !cpu.HasAVX2 {
        vclear = func(v obj.Addr)       { self.Emit("VXORPS", v, v, v) }
        vset1a = func(a, b obj.Addr)    { self.Emit("VCMPPS", a, a, b, jit.Imm(0x0f)) }
        vandpb = func(b, a, r obj.Addr) { self.Emit("VANDPS", b, a, r) }
        vcmpeq = func(b, a, r obj.Addr) { self.Emit("VCMPPS", b, a, r, jit.Imm(0x00)) }
    }

    /* if n is less than 32 byte, only scalar code will be used;
     * otherwise AVX is used, so clear Y0, and set Y1 to all 1s */
    if e >= 32 {
        vclear(_Y0)         // CLEAR Y0
        vset1a(_Y0, _Y1)    // SET1A Y0, Y1
    }

    /* 128-byte tests */
    for i <= e - 128 {
        vcmpeq(jit.Ptr(_SP_p, i +  0), _Y0, _Y2)    // CMPEQ  i+0(SP.p), Y0, Y2
        vcmpeq(jit.Ptr(_SP_p, i + 32), _Y0, _Y3)    // CMPEQ  i+32(SP.p), Y0, Y3
        vcmpeq(jit.Ptr(_SP_p, i + 64), _Y0, _Y4)    // CMPEQ  i+64(SP.p), Y0, Y4
        vcmpeq(jit.Ptr(_SP_p, i + 96), _Y0, _Y5)    // CMPEQ  i+96(SP.p), Y0, Y5
        vandpb(_Y3, _Y2, _Y2)                       // ANDPB  Y3, Y2, Y2
        vandpb(_Y5, _Y4, _Y3)                       // ANDPB  Y5, Y4, Y3
        vandpb(_Y2, _Y3, _Y3)                       // ANDPB  Y2, Y3, Y3
        self.Emit("VPTEST", _Y1, _Y3)               // VPTEST Y1, Y3
        self.Sjmp("JNC"   , "_not_zero_z_{n}")      // JNC    _not_zero_z_{n}
        i += 128
    }

    /* 32-byte tests */
    for i <= e - 32 {
        vcmpeq(jit.Ptr(_SP_p, i), _Y0, _Y2)     // CMPEQ  i(SP.p), Y0, Y2
        self.Emit("VPTEST", _Y1, _Y2)           // VPTEST Y1, Y2
        self.Sjmp("JNC"   , "_not_zero_z_{n}")  // JNC    _not_zero_z_{n}
        i += 32
    }

    /* VZEROUPPER to avoid AVX-SSE transition penalty */
    if e >= 32 {
        self.Emit("VZEROUPPER")
    }

    /* 8-byte tests */
    for i <= e - 8 {
        self.Emit("CMPQ", jit.Ptr(_SP_p, i), jit.Imm(0))    // CMPQ i(SP.p), $0
        self.Sjmp("JNE" , "_not_zero_{n}")                  // JNE  _not_zero_{n}
        i += 8
    }

    /* 4 byte test */
    if i <= e - 4 {
        self.Emit("CMPL", jit.Ptr(_SP_p, i), jit.Imm(0))    // CMPL i(SP.p), $0
        self.Sjmp("JNE" , "_not_zero_{n}")                  // JNE  _not_zero_{n}
        i += 4
    }

    /* 2 byte test */
    if i <= e - 2 {
        self.Emit("CMPW", jit.Ptr(_SP_p, i), jit.Imm(0))    // CMPW i(SP.p), $0
        self.Sjmp("JNE" , "_not_zero_{n}")                  // JNE  _not_zero_{n}
        i += 2
    }

    /* the last byte */
    if i < e {
        self.Emit("CMPB", jit.Ptr(_SP_p, i), jit.Imm(0))    // CMPB i(SP.p), $0
        self.Sjmp("JNE" , "_not_zero_{n}")                  // JNE  _not_zero_{n}
    }

    /* value is not zero */
    if e < 32 {
        self.Xjmp("JMP", dest)
        self.Link("_not_zero_{n}")
        return
    }

    /* VZEROUPPER to avoid AVX-SSE transition penalty */
    self.Xjmp("JMP", dest)
    self.Link("_not_zero_z_{n}")
    self.Emit("VZEROUPPER")
    self.Link("_not_zero_{n}")
}

/** OpCode Assembler Functions **/

var (
    _T_map_Iterator           = rt.UnpackType(mapIteratorType)
    _T_map_PIterator          = rt.UnpackType(mapPIteratorType)
    _T_json_Marshaler         = rt.UnpackType(jsonMarshalerType)
    _T_encoding_TextMarshaler = rt.UnpackType(encodingTextMarshalerType)
)

var (
    _P_iteratorPool = new(sync.Pool)
    _N_iteratorPool = jit.Imm(int64(unsafe.Sizeof(rt.GoMapIterator{})))
    _V_iteratorPool = jit.Imm(int64(uintptr(unsafe.Pointer(_P_iteratorPool))))
)

var (
    _F_f64toa    = jit.Imm(int64(native.S_f64toa))
    _F_i64toa    = jit.Imm(int64(native.S_i64toa))
    _F_u64toa    = jit.Imm(int64(native.S_u64toa))
    _F_lquote    = jit.Imm(int64(native.S_lquote))
    _F_b64encode = jit.Imm(int64(_subr__b64encode))
)

var (
    _F_memmove              = jit.Func(memmove)
    _F_newobject            = jit.Func(newobject)
    _F_isZeroTyped          = jit.Func(isZeroTyped)
    _F_mapiternext          = jit.Func(mapiternext)
    _F_mapiterinit          = jit.Func(mapiterinit)
    _F_error_number         = jit.Func(error_number)
    _F_isValidNumber        = jit.Func(isValidNumber)
    _F_memclrNoHeapPointers = jit.Func(memclrNoHeapPointers)
)

var (
    _F_sync_Pool_Get = jit.Func((*sync.Pool).Get)
    _F_sync_Pool_Put = jit.Func((*sync.Pool).Put)
)

var (
    _F_encodeQuote         obj.Addr
    _F_encodeDoubleQuote   obj.Addr
    _F_encodeTypedPointer  obj.Addr
    _F_encodeJsonMarshaler obj.Addr
    _F_encodeTextMarshaler obj.Addr
)

func init() {
    _F_encodeQuote         = jit.Func(encodeQuote)
    _F_encodeDoubleQuote   = jit.Func(encodeDoubleQuote)
    _F_encodeTypedPointer  = jit.Func(encodeTypedPointer)
    _F_encodeJsonMarshaler = jit.Func(encodeJsonMarshaler)
    _F_encodeTextMarshaler = jit.Func(encodeTextMarshaler)
}

func (self *_Assembler) _asm_OP_null(_ *_Instr) {
    self.check_size(4)
    self.Emit("MOVL", jit.Imm(_IM_null), jit.Sib(_RP, _RL, 1, 0))  // MOVL $'null', (RP)(RL*1)
    self.Emit("ADDQ", jit.Imm(4), _RL)                             // ADDQ $4, RL
}

func (self *_Assembler) _asm_OP_bool(_ *_Instr) {
    self.Emit("CMPB", jit.Ptr(_SP_p, 0), jit.Imm(0))                // CMPB (SP.p), $0
    self.Sjmp("JE"  , "_false_{n}")                                 // JE   _false_{n}
    self.check_size(4)                                              // SIZE $4
    self.Emit("MOVL", jit.Imm(_IM_true), jit.Sib(_RP, _RL, 1, 0))   // MOVL $'true', (RP)(RL*1)
    self.Emit("ADDQ", jit.Imm(4), _RL)                              // ADDQ $4, RL
    self.Sjmp("JMP" , "_end_{n}")                                   // JMP  _end_{n}
    self.Link("_false_{n}")                                         // _false_{n}:
    self.check_size(5)                                              // SIZE $5
    self.Emit("MOVL", jit.Imm(_IM_fals), jit.Sib(_RP, _RL, 1, 0))   // MOVL $'fals', (RP)(RL*1)
    self.Emit("MOVB", jit.Imm('e'), jit.Sib(_RP, _RL, 1, 4))        // MOVB $'e', 4(RP)(RL*1)
    self.Emit("ADDQ", jit.Imm(5), _RL)                              // ADDQ $5, RL
    self.Link("_end_{n}")                                           // _end_{n}:
}

func (self *_Assembler) _asm_OP_i8(_ *_Instr) {
    self.store_int(4, _F_i64toa, "MOVBQSX")
}

func (self *_Assembler) _asm_OP_i16(_ *_Instr) {
    self.store_int(6, _F_i64toa, "MOVWQSX")
}

func (self *_Assembler) _asm_OP_i32(_ *_Instr) {
    self.store_int(11, _F_i64toa, "MOVLQSX")
}

func (self *_Assembler) _asm_OP_i64(_ *_Instr) {
    self.store_int(21, _F_i64toa, "MOVQ")
}

func (self *_Assembler) _asm_OP_u8(_ *_Instr) {
    self.store_int(3, _F_u64toa, "MOVBQZX")
}

func (self *_Assembler) _asm_OP_u16(_ *_Instr) {
    self.store_int(5, _F_u64toa, "MOVWQZX")
}

func (self *_Assembler) _asm_OP_u32(_ *_Instr) {
    self.store_int(10, _F_u64toa, "MOVLQZX")
}

func (self *_Assembler) _asm_OP_u64(_ *_Instr) {
    self.store_int(20, _F_u64toa, "MOVQ")
}

func (self *_Assembler) _asm_OP_f32(_ *_Instr) {
    self.check_size(32)
    self.Emit("MOVL"    , jit.Ptr(_SP_p, 0), _AX)       // MOVL     (SP.p), AX
    self.Emit("ANDL"    , jit.Imm(_FM_exp32), _AX)      // ANDL     $_FM_exp32, AX
    self.Emit("XORL"    , jit.Imm(_FM_exp32), _AX)      // XORL     $_FM_exp32, AX
    self.Sjmp("JZ"      , _LB_error_nan_or_infinite)    // JZ       _error_nan_or_infinite
    self.save_c()                                       // SAVE     $C_regs
    self.rbuf_di()                                      // MOVQ     RP, DI
    self.Emit("MOVSS"   , jit.Ptr(_SP_p, 0), _X0)       // MOVSS    (SP.p), X0
    self.Emit("CVTSS2SD", _X0, _X0)                     // CVTSS2SD X0, X0
    self.call_c(_F_f64toa)                              // CALL_C   f64toa
    self.Emit("ADDQ"    , _AX, _RL)                     // ADDQ     AX, RL
}

func (self *_Assembler) _asm_OP_f64(_ *_Instr) {
    self.check_size(32)
    self.Emit("MOVQ"  , jit.Ptr(_SP_p, 0), _AX)     // MOVQ   (SP.p), AX
    self.Emit("MOVQ"  , jit.Imm(_FM_exp64), _CX)    // MOVQ   $_FM_exp64, CX
    self.Emit("ANDQ"  , _CX, _AX)                   // ANDQ   CX, AX
    self.Emit("XORQ"  , _CX, _AX)                   // XORQ   CX, AX
    self.Sjmp("JZ"    , _LB_error_nan_or_infinite)  // JZ     _error_nan_or_infinite
    self.save_c()                                   // SAVE   $C_regs
    self.rbuf_di()                                  // MOVQ   RP, DI
    self.Emit("MOVSD" , jit.Ptr(_SP_p, 0), _X0)     // MOVSD  (SP.p), X0
    self.call_c(_F_f64toa)                          // CALL_C f64toa
    self.Emit("ADDQ"  , _AX, _RL)                   // ADDQ   AX, RL
}

func (self *_Assembler) _asm_OP_str(_ *_Instr) {
    self.encode_string(_F_encodeQuote, false)
}

func (self *_Assembler) _asm_OP_bin(_ *_Instr) {
    self.Emit("MOVQ", jit.Ptr(_SP_p, 8), _AX)           // MOVQ 8(SP.p), AX
    self.Emit("ADDQ", jit.Imm(2), _AX)                  // ADDQ $2, AX
    self.Emit("MOVQ", jit.Imm(_IM_mulv), _CX)           // MOVQ $_MF_mulv, CX
    self.Emit("MOVQ", _DX, _R8)                         // MOVQ DX, R8
    self.From("MULQ", _CX)                              // MULQ CX
    self.Emit("LEAQ", jit.Sib(_DX, _DX, 1, 1), _AX)     // LEAQ 1(DX)(DX), AX
    self.Emit("ORQ" , jit.Imm(2), _AX)                  // ORQ  $2, AX
    self.Emit("MOVQ", _R8, _DX)                         // MOVQ R8, DX
    self.check_size_r(_AX, 0)                           // SIZE AX
    self.add_char('"')                                  // CHAR $'"'
    self.save_c()                                       // SAVE $REG_ffi
    self.prep_buffer_c()                                // MOVE {buf}, DI
    self.Emit("MOVQ", _SP_p, _SI)                       // MOVQ SP.p, SI
    self.Emit("XORL", _DX, _DX)                         // XORL DX, DX
    self.call_c(_F_b64encode)                           // CALL b64encode
    self.load_buffer()                                  // LOAD {buf}
    self.add_char('"')                                  // CHAR $'"'
}

func (self *_Assembler) _asm_OP_quote(_ *_Instr) {
    self.encode_string(_F_encodeDoubleQuote, true)
}

func (self *_Assembler) _asm_OP_number(_ *_Instr) {
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _AX)          // MOVQ    (SP.p), AX
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 8), _CX)          // MOVQ    (SP.p), CX
    self.Emit("TESTQ", _CX, _CX)                        // TESTQ   CX, CX
    self.Sjmp("JZ"   , "_empty_{n}")                    // JZ      _empty_{n}
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))            // MOVQ    AX, (SP)
    self.Emit("MOVQ" , _CX, jit.Ptr(_SP, 8))            // MOVQ    CX, 8(SP)
    self.call_go(_F_isValidNumber)                      // CALL_GO isValidNumber
    self.Emit("CMPB" , jit.Ptr(_SP, 16), jit.Imm(0))    // CMPB    16(SP), $0
    self.Sjmp("JE"   , _LB_error_invalid_number)        // JE      _error_invalid_number
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 8), _AX)          // MOVQ    8(SP.p), AX
    self.check_size_r(_AX, 0)                           // SIZE    AX
    self.Emit("LEAQ" , jit.Sib(_RP, _RL, 1, 0), _AX)    // LEAQ    (RP)(RL), AX
    self.Emit("ADDQ" , jit.Ptr(_SP_p, 8), _RL)          // ADDQ    8(SP.p), RL
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))            // MOVQ    AX, (SP)
    self.Emit("MOVOU", jit.Ptr(_SP_p, 0), _X0)          // MOVOU   (SP.p), X0
    self.Emit("MOVOU", _X0, jit.Ptr(_SP, 8))            // MOVOU   X0, 8(SP)
    self.call_go(_F_memmove)                            // CALL_GO memmove
    self.Sjmp("JMP"  , "_done_{n}")                     // JMP     _done_{n}
    self.Link("_empty_{n}")                             // _empty_{n}:
    self.check_size(1)                                  // SIZE    $1
    self.add_char('0')                                  // CHAR    $'0'
    self.Link("_done_{n}")                              // _done_{n}:
}

func (self *_Assembler) _asm_OP_eface(_ *_Instr) {
    self.prep_buffer()                          // MOVE  {buf}, (SP)s
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _AX)  // MOVQ  (SP.p), AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 8))    // MOVQ  AX, 8(SP)
    self.Emit("LEAQ" , jit.Ptr(_SP_p, 8), _AX)  // LEAQ  8(SP.p), AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 16))   // MOVQ  AX, 16(SP)
    self.Emit("MOVQ" , _ST, jit.Ptr(_SP, 24))   // MOVQ  ST, 24(SP)
    self.call_encoder(_F_encodeTypedPointer)    // CALL  encodeTypedPointer
    self.Emit("MOVQ" , jit.Ptr(_SP, 32), _ET)   // MOVQ  32(SP), ET
    self.Emit("MOVQ" , jit.Ptr(_SP, 40), _EP)   // MOVQ  40(SP), EP
    self.Emit("TESTQ", _ET, _ET)                // TESTQ ET, ET
    self.Sjmp("JNZ"  , _LB_error)               // JNZ   _error
}

func (self *_Assembler) _asm_OP_iface(_ *_Instr) {
    self.prep_buffer()                          // MOVE  {buf}, (SP)
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _AX)  // MOVQ  (SP.p), AX
    self.Emit("MOVQ" , jit.Ptr(_AX, 8), _AX)    // MOVQ  8(AX), AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 8))    // MOVQ  AX, 8(SP)
    self.Emit("LEAQ" , jit.Ptr(_SP_p, 8), _AX)  // LEAQ  8(SP.p), AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 16))   // MOVQ  AX, 16(SP)
    self.Emit("MOVQ" , _ST, jit.Ptr(_SP, 24))   // MOVQ  ST, 24(SP)
    self.call_encoder(_F_encodeTypedPointer)    // CALL  encodeTypedPointer
    self.Emit("MOVQ" , jit.Ptr(_SP, 32), _ET)   // MOVQ  32(SP), ET
    self.Emit("MOVQ" , jit.Ptr(_SP, 40), _EP)   // MOVQ  40(SP), EP
    self.Emit("TESTQ", _ET, _ET)                // TESTQ ET, ET
    self.Sjmp("JNZ"  , _LB_error)               // JNZ   _error
}

func (self *_Assembler) _asm_OP_byte(p *_Instr) {
    self.check_size(1)
    self.Emit("MOVB", jit.Imm(p.i64()), jit.Sib(_RP, _RL, 1, 0))    // MOVL p.vi(), (RP)(RL*1)
    self.Emit("ADDQ", jit.Imm(1), _RL)                              // ADDQ $1, RL
}

func (self *_Assembler) _asm_OP_text(p *_Instr) {
    self.check_size(len(p.vs()))
    self.store_str(p.vs())
    self.Emit("ADDQ", jit.Imm(int64(len(p.vs()))), _RL)     // ADDQ $len(p.vs()), RL
}

func (self *_Assembler) _asm_OP_deref(_ *_Instr) {
    self.Emit("MOVQ", jit.Ptr(_SP_p, 0), _SP_p)     // MOVQ (SP.p), SP.p
}

func (self *_Assembler) _asm_OP_index(p *_Instr) {
    self.Emit("MOVQ", jit.Imm(p.i64()), _AX)    // MOVQ $p.vi(), AX
    self.Emit("ADDQ", _AX, _SP_p)               // ADDQ AX, SP.p
}

func (self *_Assembler) _asm_OP_load(_ *_Instr) {
    self.Emit("MOVQ", jit.Ptr(_ST, 0), _AX)                 // MOVQ (ST), AX
    self.Emit("MOVQ", jit.Sib(_ST, _AX, 1, -24), _SP_x)     // MOVQ -24(ST)(AX), SP.x
    self.Emit("MOVQ", jit.Sib(_ST, _AX, 1, -8), _SP_p)      // MOVQ -8(ST)(AX), SP.p
    self.Emit("MOVQ", jit.Sib(_ST, _AX, 1, 0), _SP_q)       // MOVQ (ST)(AX), SP.q
}

func (self *_Assembler) _asm_OP_save(_ *_Instr) {
    self.save_state()
}

func (self *_Assembler) _asm_OP_drop(_ *_Instr) {
    self.drop_state(_StateSize)
}

func (self *_Assembler) _asm_OP_drop_2(_ *_Instr) {
    self.drop_state(_StateSize * 2)                     // DROP  $(_StateSize * 2)
    self.Emit("MOVOU", _X0, jit.Sib(_ST, _AX, 1, 56))   // MOVOU X0, 56(ST)(AX)
}

func (self *_Assembler) _asm_OP_recurse(p *_Instr) {
    self.prep_buffer()                          // MOVE {buf}, (SP)
    self.Emit("MOVQ", jit.Type(p.vt()), _AX)    // MOVQ $(type(p.vt())), AX
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 8))     // MOVQ AX, 8(SP)

    /* check for indirection */
    if p.vk() == reflect.Ptr {
        self.Emit("MOVQ", _SP_p, _AX)               // MOVQ SP.p, AX
    } else {
        self.Emit("MOVQ", _SP_p, jit.Ptr(_SP, 48))  // MOVQ SP.p, 48(SP)
        self.Emit("LEAQ", jit.Ptr(_SP, 48), _AX)    // LEAQ 48(SP), AX
    }

    /* call the encoder */
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 16))   // MOVQ  AX, 16(SP)
    self.Emit("MOVQ" , _ST, jit.Ptr(_SP, 24))   // MOVQ  ST, 24(SP)
    self.call_encoder(_F_encodeTypedPointer)    // CALL  encodeTypedPointer
    self.Emit("MOVQ" , jit.Ptr(_SP, 32), _ET)   // MOVQ  32(SP), ET
    self.Emit("MOVQ" , jit.Ptr(_SP, 40), _EP)   // MOVQ  40(SP), EP
    self.Emit("TESTQ", _ET, _ET)                // TESTQ ET, ET
    self.Sjmp("JNZ"  , _LB_error)               // JNZ   _error
}

func (self *_Assembler) _asm_OP_is_nil(p *_Instr) {
    self.Emit("CMPQ", jit.Ptr(_SP_p, 0), jit.Imm(0))    // CMPQ (SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_nil_p1(p *_Instr) {
    self.Emit("CMPQ", jit.Ptr(_SP_p, 8), jit.Imm(0))    // CMPQ 8(SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_1(p *_Instr) {
    self.Emit("CMPB", jit.Ptr(_SP_p, 0), jit.Imm(0))    // CMPB (SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_2(p *_Instr) {
    self.Emit("CMPW", jit.Ptr(_SP_p, 0), jit.Imm(0))    // CMPW (SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_4(p *_Instr) {
    self.Emit("CMPL", jit.Ptr(_SP_p, 0), jit.Imm(0))    // CMPL (SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_8(p *_Instr) {
    self.Emit("CMPQ", jit.Ptr(_SP_p, 0), jit.Imm(0))    // CMPQ (SP.p), $0
    self.Xjmp("JE"  , p.vi())                           // JE   p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_map(p *_Instr) {
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _AX)          // MOVQ  (SP.p), AX
    self.Emit("TESTQ", _AX, _AX)                        // TESTQ AX, AX
    self.Xjmp("JZ"   , p.vi())                          // JZ    p.vi()
    self.Emit("CMPQ" , jit.Ptr(_AX, 0), jit.Imm(0))     // CMPQ  (AX), $0
    self.Xjmp("JE"   , p.vi())                          // JE    p.vi()
}

func (self *_Assembler) _asm_OP_is_zero_mem(p *_Instr) {
    self.check_zero(p.vlen(), p.vi())
}

func (self *_Assembler) _asm_OP_is_zero_safe(p *_Instr) {
    self.check_zero(p.vlen(), p.vi())                   // CHECKZ  $p.vlen(), p.vi()
    self.Emit("MOVQ", jit.Type(p.vt()), _AX)            // MOVQ    $p.vt(), AX
    self.Emit("MOVQ", _SP_p, jit.Ptr(_SP, 0))           // MOVQ    SP.p, (SP)
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 8))             // MOVQ    AX, 8(SP)
    self.call_go(_F_isZeroTyped)                        // CALL_GO isZeroTyped
    self.Emit("CMPQ", jit.Ptr(_SP, 16), jit.Imm(0))     // CMPQ    16(SP), $0
    self.Xjmp("JNE" , p.vi())                           // JNE     p.vi()
}

func (self *_Assembler) _asm_OP_goto(p *_Instr) {
    self.Xjmp("JMP", p.vi())
}

func (self *_Assembler) _asm_OP_map_iter(p *_Instr) {
    self.Emit("MOVQ" , _V_iteratorPool, _AX)                // MOVQ    $&iteratorPool, AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))                // MOVQ    AX, (SP)
    self.call_go(_F_sync_Pool_Get)                          // CALL_GO (*sync.Pool).Get
    self.Emit("MOVQ" , jit.Ptr(_SP, 16), _SP_q)             // MOVQ    16(SP), SP.q
    self.Emit("TESTQ", _SP_q, _SP_q)                        // TESTQ   SP.q, SP.q
    self.Sjmp("JZ"   , "_new_iter_{n}")                     // JZ      _new_iter_{n}
    self.Emit("MOVL" , _N_iteratorPool, _AX)                // MOVL    ${size(GoMapIterator)}, AX
    self.Emit("MOVQ" , _SP_q, jit.Ptr(_SP, 0))              // MOVQ    SP.q, (SP)
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 8))                // MOVQ    AX, 8(SP)
    self.call_go(_F_memclrNoHeapPointers)                   // CALL_GO memclrNoHeapPointers
    self.Sjmp("JMP"  , "_init_iter_{n}")                    // JMP     _init_iter_{n}
    self.Link("_new_iter_{n}")                              // _new_iter_{n}:
    self.Emit("MOVQ" , jit.Gtype(_T_map_Iterator), _AX)     // MOVQ    ${type(GoMapIterator)}, AX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))                // MOVQ    AX, (SP)
    self.call_go(_F_newobject)                              // CALL_GO newobject
    self.Emit("MOVQ" , jit.Ptr(_SP, 8), _SP_q)              // MOVQ    8(SP), SP.q
    self.Link("_init_iter_{n}")                             // _init_iter_{n}:
    self.Emit("MOVQ" , jit.Type(p.vt()), _AX)               // MOVQ    $p.vt(), AX
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _CX)              // MOVQ    (SP.p), CX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))                // MOVQ    AX, (SP)
    self.Emit("MOVQ" , _CX, jit.Ptr(_SP, 8))                // MOVQ    CX, 8(SP)
    self.Emit("MOVQ" , _SP_q, jit.Ptr(_SP, 16))             // MOVQ    SP.q, 16(SP)
    self.call_go(_F_mapiterinit)                            // CALL_GO mapiterinit
}

func (self *_Assembler) _asm_OP_map_check_key(p *_Instr) {
    self.Emit("MOVQ" , jit.Ptr(_SP_q, 0), _SP_p)            // MOVQ    (SP.q), SP.p
    self.Emit("TESTQ", _SP_p, _SP_p)                        // TESTQ   SP.p, SP.p
    self.Sjmp("JNZ"  , "_map_next_{n}")                     // JNZ     _map_next_{n}
    self.Emit("MOVQ" , _V_iteratorPool, _AX)                // MOVQ    $&iteratorPool, AX
    self.Emit("MOVQ" , jit.Gtype(_T_map_PIterator), _CX)    // MOVQ    ${type(*GoMapIterator)}, CX
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 0))                // MOVQ    AX, (SP)
    self.Emit("MOVQ" , _CX, jit.Ptr(_SP, 8))                // MOVQ    CX, 8(SP)
    self.Emit("MOVQ" , _SP_q, jit.Ptr(_SP, 16))             // MOVQ    SP.q, 16(SP)
    self.call_go(_F_sync_Pool_Put)                          // CALL_GO (*sync.Pool).Put
    self.Emit("XORL" , _SP_q, _SP_q)                        // XORL    SP.q, SP.q
    self.Xjmp("JMP"  , p.vi())                              // JMP     p.vi()
    self.Link("_map_next_{n}")                              // _map_next_{n}:
}

func (self *_Assembler) _asm_OP_map_value_next(_ *_Instr) {
    self.Emit("MOVQ", jit.Ptr(_SP_q, 8), _SP_p)     // MOVQ    8(SP.q), SP.p
    self.Emit("MOVQ", _SP_q, jit.Ptr(_SP, 0))       // MOVQ    SP.q, (SP)
    self.call_go(_F_mapiternext)                    // CALL_GO mapiternext
}

func (self *_Assembler) _asm_OP_slice_len(_ *_Instr) {
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 8), _SP_x)        // MOVQ  8(SP.p), SP.x
    self.Emit("MOVQ" , jit.Ptr(_SP_p, 0), _SP_p)        // MOVQ  (SP.p), SP.p
    self.Emit("ORQ"  , jit.Imm(1 << _S_init), _SP_f)    // ORQ   $(1<<_S_init), SP.f
}

func (self *_Assembler) _asm_OP_slice_next(p *_Instr) {
    self.Emit("TESTQ"  , _SP_x, _SP_x)                          // TESTQ   SP.x, SP.x
    self.Xjmp("JZ"     , p.vi())                                // JZ      p.vi()
    self.Emit("SUBQ"   , jit.Imm(1), _SP_x)                     // SUBQ    $1, SP.x
    self.Emit("BTRQ"   , jit.Imm(_S_init), _SP_f)               // BTRQ    $_S_init, SP.f
    self.Emit("LEAQ"   , jit.Ptr(_SP_p, int64(p.vlen())), _AX)  // LEAQ    $(p.vlen())(SP.p), AX
    self.Emit("CMOVQCC", _AX, _SP_p)                            // CMOVQNC AX, SP.p
}

func (self *_Assembler) _asm_OP_marshal(p *_Instr) {
    self.call_marshaler(_F_encodeJsonMarshaler, _T_json_Marshaler, p.vt())
}

func (self *_Assembler) _asm_OP_marshal_p(p *_Instr) {
    if p.vk() != reflect.Ptr {
        panic("marshal_p: invalid type")
    } else {
        self.call_marshaler_v(_F_encodeJsonMarshaler, _T_json_Marshaler, p.vt(), false)
    }
}

func (self *_Assembler) _asm_OP_marshal_text(p *_Instr) {
    self.call_marshaler(_F_encodeTextMarshaler, _T_encoding_TextMarshaler, p.vt())
}

func (self *_Assembler) _asm_OP_marshal_text_p(p *_Instr) {
    if p.vk() != reflect.Ptr {
        panic("marshal_text_p: invalid type")
    } else {
        self.call_marshaler_v(_F_encodeTextMarshaler, _T_encoding_TextMarshaler, p.vt(), false)
    }
}

func (self *_Assembler) _asm_OP_cond_set(_ *_Instr) {
    self.Emit("ORQ", jit.Imm(1 << _S_cond), _SP_f)  // ORQ $(1<<_S_cond), SP.f
}

func (self *_Assembler) _asm_OP_cond_testc(p *_Instr) {
    self.Emit("BTRQ", jit.Imm(_S_cond), _SP_f)      // BTRQ $_S_cond, SP.f
    self.Xjmp("JC"  , p.vi())
}
