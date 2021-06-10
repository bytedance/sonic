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
    `reflect`
    `unsafe`

    `github.com/bytedance/sonic/internal/jit`
    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/twitchyliquid64/golang-asm/obj`
)

const (
    _VD_args   = 56     // 56 bytes for passing arguments to other Go functions
    _VD_saves  = 40     // 40 bytes for saving the registers before CALL instructions
    _VD_locals = 48     // 48 bytes for local variables
)

const (
    _VD_offs = _VD_args + _VD_saves + _VD_locals
    _VD_size = _VD_offs + 8     // 8 bytes for the parent frame pointer
)

const (
    _LB_done         = "_done"
    _LB_esc_error    = "_esc_error"
    _LB_type_error   = "_type_error"
    _LB_value_error  = "_value_error"
    _LB_switch_table = "_switch_table"
)

var (
    _RT = jit.Reg("R8")
    _RV = jit.Reg("R9")
)

var (
    _VAR_ss_Sp = jit.Ptr(_SP, _VD_args + _VD_saves)
    _VAR_ss_Sn = jit.Ptr(_SP, _VD_args + _VD_saves + 8)
)

var (
    _VAR_vv    = _VAR_vv_Vt
    _VAR_vv_Vt = jit.Ptr(_SP, _VD_args + _VD_saves + 24)
    _VAR_vv_Dv = jit.Ptr(_SP, _VD_args + _VD_saves + 32)
    _VAR_vv_Iv = jit.Ptr(_SP, _VD_args + _VD_saves + 40)
    _VAR_vv_Ep = jit.Ptr(_SP, _VD_args + _VD_saves + 48)
)

//go:noescape
//goland:noinspection GoUnusedParameter
func decodeArray() unsafe.Pointer

//go:noescape
//goland:noinspection GoUnusedParameter
func decodeObject() unsafe.Pointer

type _ValueDecoder struct {
    jit.BaseAssembler
}

func (self *_ValueDecoder) load() uintptr {
    self.Init(self.compile)
    return *(*uintptr)(self.Load("decode_value", _VD_size, 0))
}

func (self *_ValueDecoder) compile() {
    self.prologue()
    self.instrs()
    self.epilogue()
    self.errors()
    self.tables()
}

func (self *_ValueDecoder) epilogue() {
    self.Link(_LB_done)                             // _done:
    self.Emit("XORL", _EP, _EP)                     // XORL EP, EP
    self.Link(_LB_error)                            // _error:
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_offs), _BP)  // MOVQ _VD_offs(SP), BP
    self.Emit("ADDQ", jit.Imm(_VD_size), _SP)       // ADDQ $_VD_size, SP
    self.Emit("RET")
}

func (self *_ValueDecoder) prologue() {
    self.Emit("SUBQ", jit.Imm(_VD_size), _SP)       // SUBQ $_VD_size, SP
    self.Emit("MOVQ", _BP, jit.Ptr(_SP, _VD_offs))  // MOVQ BP, _VD_offs(SP)
    self.Emit("LEAQ", jit.Ptr(_SP, _VD_offs), _BP)  // LEAQ _VD_offs(SP), BP
}

/** Decoder Assembler **/

var (
    _Vp_zero  = unsafe.Pointer(&struct{}{})
    _Vp_true  = rt.UnpackEface(true).Value
    _Vp_false = rt.UnpackEface(false).Value
)

var (
    _V_max   = jit.Imm(int64(types.V_MAX))
    _V_eof   = jit.Imm(int64(types.ERR_EOF))
    _F_value = jit.Imm(int64(native.S_value))
)

var (
    _V_zero  = jit.Imm(int64(uintptr(_Vp_zero)))
    _V_true  = jit.Imm(int64(uintptr(_Vp_true)))
    _V_false = jit.Imm(int64(uintptr(_Vp_false)))
)

var (
    _T_bool        = jit.Type(reflect.TypeOf(true))
    _T_string      = jit.Type(reflect.TypeOf(""))
    _T_float64     = jit.Type(reflect.TypeOf(0.0))
    _T_json_Number = jit.Type(jsonNumberType)
)

var (
    _T_iface_sl  = jit.Type(reflect.TypeOf([]interface{}(nil)))
    _T_iface_map = jit.Type(reflect.TypeOf(map[string]interface{}(nil)))
)

var (
    _F_convTslice  = jit.Func(convTslice)
    _F_convTstring = jit.Func(convTstring)
)

var (
    _F_decodeArray        = jit.Func(decodeArray)
    _F_decodeObject       = jit.Func(decodeObject)
    _F_throw_invalid_type = jit.Func(throw_invalid_type)
)

const (
    _SW_case_V_EOF     = _LB_error
    _SW_case_V_NULL    = _LB_done
    _SW_case_V_TRUE    = "_case_V_TRUE"
    _SW_case_V_FALSE   = "_case_V_FALSE"
    _SW_case_V_ARRAY   = "_case_V_ARRAY"
    _SW_case_V_OBJECT  = "_case_V_OBJECT"
    _SW_case_V_STRING  = "_case_V_STRING"
    _SW_case_V_DOUBLE  = "_case_V_DOUBLE"
    _SW_case_V_INTEGER = "_case_V_INTEGER"
)

func (self *_ValueDecoder) call(pc obj.Addr) {
    self.Emit("MOVQ", pc, _AX)  // MOVQ ${pc}, AX
    self.Rjmp("CALL", _AX)      // CALL AX
}

func (self *_ValueDecoder) call_go(pc obj.Addr) {
    self.Emit("MOVQ", _IP, jit.Ptr(_SP, _VD_args))          // MOVQ IP, args+0(SP)
    self.Emit("MOVQ", _IL, jit.Ptr(_SP, _VD_args + 8))      // MOVQ IL, args+8(SP)
    self.Emit("MOVQ", _IC, jit.Ptr(_SP, _VD_args + 16))     // MOVQ IC, args+16(SP)
    self.Emit("MOVQ", _ST, jit.Ptr(_SP, _VD_args + 24))     // MOVQ ST, args+24(SP)
    self.Emit("MOVQ", _VP, jit.Ptr(_SP, _VD_args + 32))     // MOVQ VP, args+24(SP)
    self.call(pc)
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_args), _IP)          // MOVQ args+0(SP), IP
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_args + 8), _IL)      // MOVQ args+8(SP), IL
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_args + 16), _IC)     // MOVQ args+16(SP), IC
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_args + 24), _ST)     // MOVQ args+24(SP), ST
    self.Emit("MOVQ", jit.Ptr(_SP, _VD_args + 32), _VP)     // MOVQ args+32(SP), VP
}

func (self *_ValueDecoder) errors() {
    self.Link(_LB_esc_error)                    // _esc_error:
    self.Emit("MOVQ", _VAR_ss_Sn, _CX)          // MOVQ ss.Sn, CX
    self.Emit("SUBQ", _VAR_vv_Ep, _CX)          // SUBQ vv.Ep, CX
    self.Emit("SUBQ", _CX, _IC)                 // SUBQ CX, IC
    self.Emit("SUBQ", jit.Imm(1), _IC)          // SUBQ $1, IC
    self.Link(_LB_value_error)                  // _value_error:
    self.Emit("NEGQ", _AX)                      // NEGQ AX
    self.Emit("MOVQ", _AX, _EP)                 // MOVQ AX, EP
    self.Sjmp("JMP" , _LB_error)                // JMP  _error
    self.Link(_LB_type_error)                   // _type_error:
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))     // MOVQ AX, (SP)
    self.call(_F_throw_invalid_type)            // CALL throw_invalid_type
    self.Emit("UD2")                            // UD2
}

func (self *_ValueDecoder) tables() {
    self.Link(_LB_switch_table)         // _switch_table:
    self.Sref(_SW_case_V_EOF, 0)        // SREF &_error, $0
    self.Sref(_SW_case_V_NULL, -4)      // SREF &_done, $-4
    self.Sref(_SW_case_V_TRUE, -8)      // SREF &_case_V_TRUE, $-8
    self.Sref(_SW_case_V_FALSE, -12)    // SREF &_case_V_FALSE, $-12
    self.Sref(_SW_case_V_ARRAY, -16)    // SREF &_case_V_ARRAY, $-16
    self.Sref(_SW_case_V_OBJECT, -20)   // SREF &_case_V_OBJECT, $-20
    self.Sref(_SW_case_V_STRING, -24)   // SREF &_case_V_STRING, $-24
    self.Sref(_SW_case_V_DOUBLE, -28)   // SREF &_case_V_DOUBLE, $-28
    self.Sref(_SW_case_V_INTEGER, -32)  // SREF &_case_V_INTEGER, $-32
}

func (self *_ValueDecoder) instrs() {
    self.Emit("MOVQ", _IP, _DI)         // MOVQ IP, DI
    self.Emit("MOVQ", _IL, _SI)         // MOVQ IL, SI
    self.Emit("MOVQ", _IC, _DX)         // MOVQ IC, DX
    self.Emit("LEAQ", _VAR_vv, _CX)     // LEAQ vv, CX
    self.call(_F_value)                 // CALL value
    self.Emit("MOVQ", _AX, _IC)         // MOVQ AX, IC
    self.Emit("MOVQ", _VAR_vv_Vt, _AX)  // MOVQ vv.Vt, AX

    /* check for errors & type range */
    self.Emit("TESTQ", _AX, _AX)            // TESTQ AX, AX
    self.Sjmp("JS"   , _LB_value_error)     // JS    _value_error
    self.Sjmp("JZ"   , _LB_type_error)      // JZ    _type_error
    self.Emit("CMPQ" , _AX, _V_max)         // CMPQ  AX, ${native.V_MAX}
    self.Sjmp("JA"   , _LB_type_error)      // JA    _type_error
    self.Emit("XORL" , _RT, _RT)            // XORL  RT, RT
    self.Emit("XORL" , _RV, _RV)            // XORL  RV, RV
    self.Emit("MOVQ" , _V_eof, _EP)         // MOVQ  ${native.ERR_EOF}, EP

    /* jump table selector */
    self.Byte(0x48, 0x8d, 0x3d)                             // LEAQ    ?(PC), DI
    self.Sref(_LB_switch_table, 4)                          // ....    &_switch_table
    self.Emit("MOVLQSX", jit.Sib(_DI, _AX, 4, -4), _AX)     // MOVLQSX -4(DI)(AX*4), AX
    self.Emit("ADDQ"   , _DI, _AX)                          // ADDQ    DI, AX
    self.Rjmp("JMP"    , _AX)                               // JMP     AX

    /* V_TRUE */
    self.Link(_SW_case_V_TRUE)
    self.Emit("MOVQ", _T_bool, _RT)     // MOVQ ${type(bool)}, RT
    self.Emit("MOVQ", _V_true, _RV)     // MOVQ ${&true}, RV
    self.Sjmp("JMP" , _LB_done)         // JMP  _done

    /* V_FALSE */
    self.Link(_SW_case_V_FALSE)
    self.Emit("MOVQ", _T_bool, _RT)     // MOVQ ${type(bool)}, RT
    self.Emit("MOVQ", _V_false, _RV)    // MOVQ ${&false}, RV
    self.Sjmp("JMP" , _LB_done)         // JMP  _done

    /* V_ARRAY */
    self.Link(_SW_case_V_ARRAY)
    self.call(_F_decodeArray)                       // CALL    decodeArray
    self.Emit("MOVQ"   , _V_zero, _CX)              // MOVQ    ${zero}, CX
    self.Emit("TESTQ"  , _AX, _AX)                  // TESTQ   AX, AX
    self.Sjmp("JS"     , _LB_value_error)           // JS      _value_error
    self.Emit("CMOVQEQ", _CX, _RV)                  // CMOVQEQ CX, RV
    self.Emit("MOVQ"   , _RV, jit.Ptr(_SP, 0))      // MOVQ    RV, (SP)
    self.Emit("MOVQ"   , _AX, jit.Ptr(_SP, 8))      // MOVQ    AX, 8(SP)
    self.Emit("MOVQ"   , _AX, jit.Ptr(_SP, 16))     // MOVQ    AX, 16(SP)
    self.call_go(_F_convTslice)                     // CALL_GO convTslice
    self.Emit("MOVQ"   , _T_iface_sl, _RT)          // MOVQ    ${type([]interface{})}, RT
    self.Emit("MOVQ"   , jit.Ptr(_SP, 24), _RV)     // MOVQ    24(SP), RV
    self.Sjmp("JMP"    , _LB_done)                  // JMP     _done

    /* V_OBJECT */
    self.Link(_SW_case_V_OBJECT)
    self.call(_F_decodeObject)              // CALL  decodeObject
    self.Emit("TESTQ", _AX, _AX)            // TESTQ AX, AX
    self.Sjmp("JNZ"  , _LB_value_error)     // JNZ   _value_error
    self.Emit("MOVQ" , _T_iface_map, _RT)   // MOVQ  ${type(map[string]interface{})}, RT
    self.Sjmp("JMP"  , _LB_done)            // JMP   _done

    /* V_STRING */
    self.Link(_SW_case_V_STRING)
    self.Emit("MOVQ" , _VAR_vv_Iv, _CX)                         // MOVQ    vv.Iv, CX
    self.Emit("MOVQ" , _IP, _DI)                                // MOVQ    IP, DI
    self.Emit("MOVQ" , _IC, _SI)                                // MOVQ    IC, SI
    self.Emit("ADDQ" , _CX, _DI)                                // ADDQ    CX, DI
    self.Emit("SUBQ" , _CX, _SI)                                // SUBQ    CX, SI
    self.Emit("SUBQ" , jit.Imm(1), _SI)                         // SUBQ    $1, SI
    self.Emit("CMPQ" , _VAR_vv_Ep, jit.Imm(-1))                 // CMPQ    vv.Ep, $-1
    self.Sjmp("JE"   , "_noescape")                             // JE      _noescape
    self.Emit("XORL" , _AX, _AX)                                // XORL    AX, AX
    self.Emit("MOVQ" , _T_byte, _CX)                            // MOVQ    ${type(byte)}, CX
    self.Emit("MOVQ" , _DI, _VAR_ss_Sp)                         // MOVQ    DI, ss.Sp
    self.Emit("MOVQ" , _SI, _VAR_ss_Sn)                         // MOVQ    SI, ss.Sn
    self.Emit("MOVQ" , _SI, jit.Ptr(_SP, 0))                    // MOVQ    SI, (SP)
    self.Emit("MOVQ" , _CX, jit.Ptr(_SP, 8))                    // MOVQ    CX, 8(SP)
    self.Emit("MOVQ" , _AX, jit.Ptr(_SP, 16))                   // MOVQ    AX, 16(SP)
    self.call_go(_F_mallocgc)                                   // CALL_GO mallocgc
    self.Emit("MOVQ" , _VAR_ss_Sp, _DI)                         // MOVQ    ss.Sp, DI
    self.Emit("MOVQ" , _VAR_ss_Sn, _SI)                         // MOVQ    ss.Sn, SI
    self.Emit("MOVQ" , jit.Ptr(_SP, 24), _DX)                   // MOVQ    24(SP), DX
    self.Emit("MOVQ" , _DX, _VAR_ss_Sp)                         // MOVQ    DX, ss.Sp
    self.Emit("LEAQ" , _VAR_vv_Ep, _CX)                         // LEAQ    vv.Ep, CX
    self.Emit("XORL" , _R8, _R8)                                // XORL    R8, R8
    self.Emit("BTQ"  , jit.Imm(_F_disable_urc), _VP)            // BTQ     ${_F_disable_urc}, VP
    self.Emit("SETCC", _R8)                                     // SETCC   R8
    self.Emit("SHLQ" , jit.Imm(types.B_UNICODE_REPLACE), _R8)   // SHLQ    ${types.B_UNICODE_REPLACE}, R8
    self.call(_F_unquote)                                       // CALL    unquote
    self.Emit("TESTQ", _AX, _AX)                                // TESTQ   AX, AX
    self.Sjmp("JS"   , _LB_esc_error)                           // JS      _esc_error
    self.Emit("MOVQ" , _AX, _SI)                                // MOVQ    AX, SI
    self.Emit("MOVQ" , _VAR_ss_Sp, _DI)                         // MOVQ    ss.Sp, DI
    self.Link("_noescape")                                      // _noescape:
    self.Emit("MOVQ" , _DI, jit.Ptr(_SP, 0))                    // MOVQ    DI, (SP)
    self.Emit("MOVQ" , _SI, jit.Ptr(_SP, 8))                    // MOVQ    SI, 8(SP)
    self.call_go(_F_convTstring)                                // CALL_GO convTstring
    self.Emit("MOVQ" , _T_string, _RT)                          // MOVQ    ${type(string)}, RT
    self.Emit("MOVQ" , jit.Ptr(_SP, 16), _RV)                   // MOVQ    16(SP), RV
    self.Sjmp("JMP"  , _LB_done)                                // JMP     _done

    /* V_DOUBLE */
    self.Link(_SW_case_V_DOUBLE)
    self.Emit("BTQ"  , jit.Imm(_F_use_number), _VP)     // BTQ     ${_F_use_number}, VP
    self.Sjmp("JC"   , "_use_number")                   // JC      _use_number
    self.Emit("MOVSD", _VAR_vv_Dv, _X0)                 // MOVSD   st.Dv, X0
    self.Emit("MOVSD", _X0, jit.Ptr(_SP, 0))            // MOVSD   X0, (SP)
    self.call_go(_F_convT64)                            // CALL_GO convT64
    self.Emit("MOVQ" , _T_float64, _RT)                 // MOVQ    ${type(float64)}, RT
    self.Emit("MOVQ" , jit.Ptr(_SP, 8), _RV)            // MOVQ    8(SP), RV
    self.Sjmp("JMP"  , _LB_done)                        // JMP     _done

    /* V_INTEGER */
    self.Link(_SW_case_V_INTEGER)
    self.Emit("BTQ" , jit.Imm(_F_use_int64), _VP)   // BTQ     ${_F_use_int64}, VP
    self.Sjmp("JNC" , _SW_case_V_DOUBLE)            // JNC     _case_V_DOUBLE
    self.Emit("MOVQ", _VAR_vv_Iv, _X0)              // MOVQ    st.Iv, AX
    self.Emit("MOVQ", _X0, jit.Ptr(_SP, 0))         // MOVQ    AX, (SP)
    self.call_go(_F_convT64)                        // CALL_GO convT64
    self.Emit("MOVQ", jit.Gtype(_T_int64), _RT)     // MOVQ    ${type(int64)}, RT
    self.Emit("MOVQ", jit.Ptr(_SP, 8), _RV)         // MOVQ    8(SP), RV
    self.Sjmp("JMP" , _LB_done)                     // JMP     _done

    /* case when `UseNumber` is set */
    self.Link("_use_number")                            // _use_number:
    self.Emit("MOVQ", _VAR_vv_Ep, _SI)                  // MOVQ ${p}, SI
    self.Emit("LEAQ", jit.Sib(_IP, _SI, 1, 0), _DI)     // LEAQ (IP)(SI), DI
    self.Emit("NEGQ", _SI)                              // NEGQ SI
    self.Emit("LEAQ", jit.Sib(_IC, _SI, 1, 0), _SI)     // LEAQ (IC)(SI), SI
    self.Emit("MOVQ", _DI, jit.Ptr(_SP, 0))             // MOVQ DI, (SP)
    self.Emit("MOVQ", _SI, jit.Ptr(_SP, 8))             // MOVQ SI, 8(SP)
    self.call_go(_F_convTstring)                        // CALL_GO convTstring
    self.Emit("MOVQ", _T_json_Number, _RT)              // MOVQ ${type(json.Number)}, RT
    self.Emit("MOVQ", jit.Ptr(_SP, 16), _RV)            // MOVQ 16(SP), RV
}

// These are referenced in `generic_amd64.s`
//goland:noinspection GoUnusedGlobalVariable
var (
    _type_byte   = rt.UnpackType(reflect.TypeOf(byte(0)))
    _type_eface  = rt.UnpackType(reflect.TypeOf((*interface{})(nil)).Elem())
    _type_strmap = rt.UnpackType(reflect.TypeOf(map[string]interface{}(nil)))
)

/** Generic Decoder **/

var (
    _subr_decode_value = new(_ValueDecoder).load()
)
