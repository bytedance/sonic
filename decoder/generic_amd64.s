//
// Copyright 2021 ByteDance Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

// Register Allocations:
//  AX      result index (if any), or negative for errors
//  R9      result value (slice array pointer or map pointer)
//  R12     parameter string buffer
//  R13     parameter string length
//  R14     parameter index
//  R15     decoder flags

#define ST                  BX
#define RT                  R8
#define RV                  R9
#define RE                  R11
#define PS                  R12
#define PN                  R13
#define PI                  R14
#define FL                  R15

#define ERR_EOF             $-1
#define ERR_INVALID_CHAR    $-2

#define lspace(to) \
	MOVQ  PS, DI                                                           \
	MOVQ  PN, SI                                                           \
	MOVQ  PI, DX                                                           \
	MOVQ  github·com∕bytedance∕sonic∕internal∕native·_subr__lspace(SB), AX \
	CALL  AX                                                               \
	MOVQ  AX, PI                                                           \
	TESTQ AX, AX                                                           \
	JNS   to                                                               \
	RET                                                                    \
to:

#define match_eof(to) \
	MOVQ ERR_EOF, AX \
	CMPQ PI, PN      \
	JBE  to          \
	MOVQ PN, PI      \
	RET              \
to:

#define match_empty(to, ch) \
	CMPB (PS)(PI*1), ch \
	JNE  to             \
	XORQ AX, AX         \
	XORQ RV, RV         \
	ADDQ $1, PI         \
	RET                 \
to:

#define match_delim(done, next, ch) \
	MOVBLZX (PS)(PI*1), AX       \
	ADDQ    $1, PI               \
	CMPL    AX, ch               \
	JE      done                 \
	CMPL    AX, $','             \
	JE      next                 \
	MOVQ    ERR_INVALID_CHAR, AX \
	SUBQ    $1, PI               \
	RET                          \
done:

#define check_char(to, ch) \
	MOVBLZX (PS)(PI*1), CX       \
	ADDQ    $1, PI               \
	MOVQ    ERR_INVALID_CHAR, AX \
	CMPL    CX, ch               \
	JE      to                   \
	SUBQ    $1, PI               \
	RET                          \
to:

#define check_empty(to, ch) \
	match_eof(_check_empty) \
	match_empty(to, ch)

#define check_delim(done, next, ch) \
match_eof(_check_delim)     \
match_delim(done, next, ch) \

// Generic Array Decoder

#define ARR_st  fl-80(SP)
#define ARR_fl  fl-72(SP)
#define ARR_ps  ps-64(SP)
#define ARR_pn  ps-56(SP)
#define ARR_pi  pi-48(SP)
#define ARR_sp  sp-40(SP)
#define ARR_sl  sl-32(SP)
#define ARR_sc  sc-24(SP)
#define ARR_rt  sl-16(SP)
#define ARR_rv  sc-8(SP)

TEXT ·decodeArray(SB), NOSPLIT, $144 - 8
	NO_LOCAL_POINTERS
	lspace(_check_array_empty)
	check_empty(_make_slice, $']')

	// set slice initial capacity to 16
	XORQ CX, CX
	MOVL $16, DX
	MOVQ CX, ARR_sl
	MOVQ DX, ARR_sc

	// allocate memory for slice
	MOVQ ST, ARR_st
	MOVQ FL, ARR_fl
	MOVQ PS, ARR_ps
	MOVQ PN, ARR_pn
	MOVQ PI, ARR_pi
	MOVQ ·_type_eface(SB), AX
	MOVQ AX, (SP)
	MOVQ CX, 8(SP)
	MOVQ DX, 16(SP)
	CALL runtime·makeslice(SB)
	MOVQ 24(SP), AX
	MOVQ AX, ARR_sp
	MOVQ AX, ret+0(FP)
	MOVQ ARR_st, ST
	MOVQ ARR_fl, FL
	MOVQ ARR_ps, PS
	MOVQ ARR_pn, PN
	MOVQ ARR_pi, PI
	GO_RESULTS_INITIALIZED

_parse_loop:
	MOVQ  ·_subr_decode_value(SB), AX
	CALL  AX
	TESTQ RE, RE
	JNZ   _parsing_error

	// check for slice space
	MOVQ ARR_sp, DI
	MOVQ ARR_sl, CX
	CMPQ CX, ARR_sc
	JAE  _more_space

_append_slice:
	LEAQ (DI)(CX*8), DI
	LEAQ (DI)(CX*8), DI
	MOVQ RT, (DI)
	MOVQ RV, 8(DI)
	ADDQ $1, CX
	MOVQ CX, ARR_sl

	// check for the delimiter
	lspace(_check_array_delim)
	check_delim(_done, _parse_loop, $']')

	// set the result register and length
	MOVQ ARR_sp, RV
	MOVQ ARR_sl, AX
	RET

_parsing_error:
	MOVQ RE, AX
	NEGQ AX
	RET

_more_space:
	MOVQ ST, ARR_st
	MOVQ FL, ARR_fl
	MOVQ PS, ARR_ps
	MOVQ PN, ARR_pn
	MOVQ PI, ARR_pi
	MOVQ RT, ARR_rt
	MOVQ RV, ARR_rv
	MOVQ ARR_sc, DX
	MOVQ ·_type_eface(SB), AX
	MOVQ AX, (SP)
	MOVQ DI, 8(SP)
	MOVQ CX, 16(SP)
	MOVQ DX, 24(SP)
	SHLQ $1, DX
	MOVQ DX, 32(SP)
	CALL runtime·growslice(SB)
	MOVQ 40(SP), DI
	MOVQ 48(SP), CX
	MOVQ 56(SP), DX
	MOVQ DI, ARR_sp
	MOVQ CX, ARR_sl
	MOVQ DX, ARR_sc
	MOVQ ARR_st, ST
	MOVQ ARR_fl, FL
	MOVQ ARR_ps, PS
	MOVQ ARR_pn, PN
	MOVQ ARR_pi, PI
	MOVQ ARR_rt, RT
	MOVQ ARR_rv, RV
	JMP  _append_slice

// Generic Object Decoder

#define OBJ_ss      OBJ_ss_Sp
#define OBJ_ss_Sp   ss_Sp-72(SP)
#define OBJ_ss_Sn   ss_Sn-64(SP)

#define OBJ_st  ps-56(SP)
#define OBJ_fl  ps-48(SP)
#define OBJ_ps  ps-40(SP)
#define OBJ_pn  ps-32(SP)
#define OBJ_pi  pi-24(SP)
#define OBJ_vp  vp-16(SP)
#define OBJ_mp  mp-8(SP)

TEXT ·decodeObject(SB), NOSPLIT, $104 - 8
	NO_LOCAL_POINTERS
	lspace(_check_object_empty)
	check_empty(_make_map, $'}')

	// create the result map
	MOVQ ST, OBJ_st
	MOVQ FL, OBJ_fl
	MOVQ PS, OBJ_ps
	MOVQ PN, OBJ_pn
	MOVQ PI, OBJ_pi
	CALL runtime·makemap_small(SB)
	MOVQ (SP), AX
	MOVQ AX, OBJ_mp
	MOVQ AX, ret+0(FP)
	MOVQ OBJ_st, ST
	MOVQ OBJ_fl, FL
	MOVQ OBJ_ps, PS
	MOVQ OBJ_pn, PN
	MOVQ OBJ_pi, PI
	GO_RESULTS_INITIALIZED

_parse_loop:
	CALL  decodeObjectKey(SB)
	MOVQ  DI, OBJ_ss_Sp
	MOVQ  SI, OBJ_ss_Sn
	TESTQ AX, AX
	JNS   _value_delim
	RET

_value_delim:
	lspace(_check_value_delim)
	check_char(_parse_value, $':')

	// allocate a new slot in the map
	MOVQ  ST, OBJ_st
	MOVQ  FL, OBJ_fl
	MOVQ  PS, OBJ_ps
	MOVQ  PN, OBJ_pn
	MOVQ  PI, OBJ_pi
	MOVQ  ·_type_strmap(SB), AX
	MOVQ  OBJ_mp, CX
	MOVOU OBJ_ss, X0
	MOVQ  AX, (SP)
	MOVQ  CX, 8(SP)
	MOVOU X0, 16(SP)
	CALL  runtime·mapassign_faststr(SB)
	MOVQ  32(SP), AX
	MOVQ  AX, OBJ_vp
	MOVQ  OBJ_st, ST
	MOVQ  OBJ_fl, FL
	MOVQ  OBJ_ps, PS
	MOVQ  OBJ_pn, PN
	MOVQ  OBJ_pi, PI

	// decode the value
	MOVQ  ·_subr_decode_value(SB), AX
	CALL  AX
	TESTQ RE, RE
	JNZ   _value_error

	// set the map value
	MOVQ OBJ_vp, AX
	MOVQ RT, (AX)
	MOVQ RV, 8(AX)

	// check for the delimiter
	lspace(_check_object_delim)
	check_delim(_done, _parse_loop, $'}')

	// set the result register and clear the errors
	XORQ AX, AX
	MOVQ OBJ_mp, RV
	RET

_value_error:
	MOVQ RE, AX
	NEGQ AX
	RET

// Object Key Parsing Helper

#define V_STRING            $7
#define F_DISABLE_URC       $2
#define B_UNICODE_REPLACE   $1

#define VAR_in      VAR_in_PS
#define VAR_in_ST   in_PS-88(SP)
#define VAR_in_FL   in_PS-80(SP)
#define VAR_in_PS   in_PS-72(SP)
#define VAR_in_PN   in_PN-64(SP)
#define VAR_in_PI   in_PI-56(SP)

#define VAR_ss      VAR_ss_Sp
#define VAR_ss_Sp   ss_Sp-48(SP)
#define VAR_ss_Sn   ss_Sp-40(SP)

#define VAR_vv      VAR_vv_Vt
#define VAR_vv_Vt   vv_Vt-32(SP)
#define VAR_vv_Dv   vv_Dv-24(SP)
#define VAR_vv_Iv   vv_Iv-16(SP)
#define VAR_vv_Ep   vv_Ep-8(SP)

TEXT decodeObjectKey(SB), NOSPLIT, $120 - 0
	NO_LOCAL_POINTERS
	lspace(_parse_key_begin)
	check_char(_parse_key_body, $'"')

	// parse the string
	MOVQ PS, VAR_in_PS
	MOVQ PN, VAR_in_PN
	MOVQ PI, VAR_in_PI
	LEAQ VAR_in, DI
	LEAQ VAR_in_PI, SI
	LEAQ VAR_vv, DX
	MOVQ github·com∕bytedance∕sonic∕internal∕native·_subr__vstring(SB), AX
	CALL AX
	MOVQ VAR_in_PI, PI

	// check for errors
	MOVQ  VAR_vv_Vt, AX
	TESTQ AX, AX
	JNS   _check_quote
	RET

_check_quote:
	CMPQ AX, V_STRING
	JNE  _invalid_type

	// extract the string
	MOVQ VAR_vv_Iv, CX
	MOVQ PS, DI
	MOVQ PI, SI
	ADDQ CX, DI
	SUBQ CX, SI
	SUBQ $1, SI

	// check for quotes
	CMPQ VAR_vv_Ep, $-1
	JNE  _unquote
	MOVQ SI, AX
	RET

_unquote:
	XORQ AX, AX
	MOVQ ST, VAR_in_ST
	MOVQ FL, VAR_in_FL
	MOVQ PS, VAR_in_PS
	MOVQ PN, VAR_in_PN
	MOVQ PI, VAR_in_PI
	MOVQ ·_type_byte(SB), CX
	MOVQ DI, VAR_ss_Sp
	MOVQ SI, VAR_ss_Sn

	// allocate space for unquoted string
	MOVQ SI, (SP)
	MOVQ CX, 8(SP)
	MOVQ AX, 16(SP)
	CALL runtime·mallocgc(SB)
	MOVQ 24(SP), DX
	MOVQ VAR_in_ST, ST
	MOVQ VAR_in_FL, FL
	MOVQ VAR_in_PS, PS
	MOVQ VAR_in_PN, PN
	MOVQ VAR_in_PI, PI

	// unquote the string
	MOVQ  VAR_ss_Sp, DI
	MOVQ  VAR_ss_Sn, SI
	MOVQ  DX, VAR_ss_Sp
	LEAQ  VAR_vv_Ep, CX
	XORQ  R8, R8
	BTQ   F_DISABLE_URC, FL
	SETCC R8
	SHLQ  B_UNICODE_REPLACE, R8
	MOVQ  github·com∕bytedance∕sonic∕internal∕native·_subr__unquote(SB), AX
	CALL  AX
    TESTQ AX, AX
    JS    _escape_error
	MOVQ  AX, SI
	MOVQ  VAR_ss_Sp, DI
	RET

_escape_error:
    MOVQ VAR_ss_Sn, CX
    SUBQ VAR_vv_Ep, CX
    SUBQ CX, PI
    SUBQ $1, PI
    RET

_invalid_type:
	MOVQ AX, (SP)
	CALL ·throw_invalid_type(SB)
	BYTE $0xcc
	WORD $0xfdeb
