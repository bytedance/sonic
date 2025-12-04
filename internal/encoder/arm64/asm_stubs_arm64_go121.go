//go:build arm64 && go1.21 && !go1.26
// +build arm64,go1.21,!go1.26

// Copyright 2025 Huawei Technologies Co., Ltd.
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

package arm64

import (
	"strconv"
	"unsafe"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/twitchyliquid64/golang-asm/obj"
	"github.com/twitchyliquid64/golang-asm/obj/arm64"
)

var (
	_V_writeBarrier = jit.Imm(int64(uintptr(unsafe.Pointer(&rt.RuntimeWriteBarrier))))

	_F_gcWriteBarrier2 = jit.Func(rt.GcWriteBarrier2)
)

func (self *Assembler) WritePtr(i int, ptr obj.Addr, base obj.Addr, index obj.Addr, offset int64, script_reg obj.Addr) {
	// index - _T0, script_reg - _T1
	if base.Reg == arm64.REG_R0 || index.Reg == arm64.REG_R0 {
		panic("rec contains R0!")
	}
	self.Emit("MOVD", _V_writeBarrier, _T2)
	self.EmitCmpqLdrPtr(jit.Ptr(_T2, 0), _ZR, script_reg)
	self.Sjmp("BEQ", "_no_writeBarrier"+strconv.Itoa(i)+"_{n}")

	self.xsave(_R25, _R30, _T3)
	self.Emit("MOVD", _SP_q, _R25)
	self.Emit("MOVD", _F_gcWriteBarrier2, _T3)
	self.Rjmp("CALL", _T3)
	self.Emit("MOVD", ptr, jit.Ptr(_R25, 0))
	self.EmitLdrri(jit.Ptr, base, index, offset, _T3, script_reg)
	self.Emit("MOVD", _T3, jit.Ptr(_R25, 8))
	self.xload(_R25, _R30, _T3)
	self.Link("_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
	self.EmitStrri(jit.Ptr, ptr, base, index, offset, script_reg)
}
