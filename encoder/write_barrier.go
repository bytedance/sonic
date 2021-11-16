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
	`strconv`
	`unsafe`

	`github.com/bytedance/sonic/internal/jit`
	`github.com/twitchyliquid64/golang-asm/obj`
)

type writeBarrier struct {
	enabled bool    // compiler emits a check of this before calling write barrier
	pad     [3]byte // compiler uses 32-bit load for "enabled" field
	needed  bool    // whether we need a write barrier for current GC phase
	cgo     bool    // whether we need a write barrier for a cgo check
	alignme uint64  // guarantee alignment so that compiler can use a 32 or 64-bit load
}

//go:linkname _runtime_writeBarrier runtime.writeBarrier
var _runtime_writeBarrier writeBarrier

//go:linkname gcWriteBarrierCX runtime.gcWriteBarrierCX
func gcWriteBarrierCX()

 var (
    _V_writeBarrier = jit.Imm(int64(uintptr(unsafe.Pointer(&_runtime_writeBarrier))))

    _R10 = jit.Reg("R10")

    _F_gcWriteBarrierCX = jit.Func(gcWriteBarrierCX)
)

// This uses CX to pass ptr, thus CX's value will change
func (self *_Assembler) WritePtrCX(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, _CX)
    self.xsave(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierCX, _R10)  // MOVQ ${fn}, CX
    self.Rjmp("CALL", _R10)      
    self.xload(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Link("_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
}