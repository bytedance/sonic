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

TEXT ·decodeInterface(SB), NOSPLIT, $0 - 56
	NO_LOCAL_POINTERS
	MOVQ    s+0(FP), R12
	MOVQ    n+8(FP), R13
	MOVQ    n+16(FP), R15
	XORQ    R14, R14
	MOVQ    ·_subr_decode_value(SB), AX
	CALL    AX
	XORQ    AX, AX
	TESTQ   R11, R11
	CMOVQNE AX, R8
	CMOVQNE AX, R9
	MOVQ    R14, i+24(FP)
	MOVQ    R8, t+32(FP)
	MOVQ    R9, v+40(FP)
	MOVQ    R11, e+48(FP)
	RET

TEXT ·decodeObjectKeyString(SB), NOSPLIT, $0 - 40
    NO_LOCAL_POINTERS
	MOVQ    s+0(FP), R12
	MOVQ    n+8(FP), R13
	MOVQ    i+16(FP), R14
	CALL    decodeObjectKey(SB)
	XORQ    CX, CX
	TESTQ   AX, AX
	CMOVQMI CX, DI
    MOVQ    DI, rp+24(FP)
    MOVQ    AX, rl+32(FP)
    RET
