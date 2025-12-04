//go:build arm64 && go1.17 && !go1.21
// +build arm64,go1.17,!go1.21

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

	_F_gcWriteBarrier = jit.Func(rt.GcWriteBarrierAX)
)

func (self *Assembler) WritePtr(i int, ptr obj.Addr, base obj.Addr, index obj.Addr, offset int64, script_reg obj.Addr) {
	if base.Reg == arm64.REG_R0 || index.Reg == arm64.REG_R0 {
		panic("rec contains R0!")
	}
	self.Emit("MOVD", _V_writeBarrier, _T2)
	self.EmitCmpqLdrPtr(jit.Ptr(_T2, 0), _ZR, script_reg)
	self.Sjmp("BEQ", "_no_writeBarrier"+strconv.Itoa(i)+"_{n}")
	self.xsave(_R2, _R3, _R30, _T3)
	self.EmitAdd(_T2, base, index)
	self.EmitAdd(_R2, _T2, jit.Imm(offset))
	self.Emit("MOVD", ptr, _R3)
	self.Emit("MOVD", _F_gcWriteBarrier, _T3)
	self.Rjmp("CALL", _T3)
	self.xload(_R2, _R3, _R30, _T3)
	self.Sjmp("B", "_end_writeBarrier"+strconv.Itoa(i)+"_{n}")
	self.Link("_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
	self.EmitStrri(jit.Ptr, ptr, base, index, offset, script_reg)
	self.Link("_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
}
