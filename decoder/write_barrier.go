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
	`fmt`
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

//go:linkname gcWriteBarrierAX runtime.gcWriteBarrier
func gcWriteBarrierAX()

//go:linkname gcWriteBarrierCX runtime.gcWriteBarrierCX
func gcWriteBarrierCX()

//go:linkname gcWriteBarrierDX runtime.gcWriteBarrierDX
func gcWriteBarrierDX()

 var (
    _V_writeBarrier = jit.Imm(int64(uintptr(unsafe.Pointer(&_runtime_writeBarrier))))

    // _VAR_gc_DI = jit.Ptr(_SP, _FP_fargs + _FP_saves + 96)
    // _VD_gc_DI = jit.Ptr(_SP, _VD_fargs + _VD_saves + 40)

    _R10 = jit.Reg("R10")

    _F_gcWriteBarrierAX = jit.Func(gcWriteBarrierAX)
    _F_gcWriteBarrierCX = jit.Func(gcWriteBarrierCX)
    _F_gcWriteBarrierDX = jit.Func(gcWriteBarrierDX)
    _F_print_ptr  = jit.Func(printPtr)
)

func printPtr(i int, ptrs [3]uintptr) {
    fmt.Printf("%d: [", i)
    for _, ptr := range ptrs {
        fmt.Printf("%x, ", ptr)
    }
    fmt.Println("]")
}

func (self *_Assembler) print_ptr(i int, ptrs ...obj.Addr) {
    // self.Emit("MOVQ", _AX, _VAR_gc_AX)
    self.Emit("MOVQ", jit.Imm(int64(i)), _AX)
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))
    for i:=0; i<3; i++ {
        if i<len(ptrs) {
            self.Emit("MOVQ", ptrs[i],  _AX)      
            self.Emit("MOVQ", _AX,  jit.Ptr(_SP, int64(i*8+8)))  
        }else{
            self.Emit("MOVQ", jit.Imm(0),  jit.Ptr(_SP, int64(i*8+8)))  
        }
    }
    self.call_go(_F_print_ptr)
    // self.Emit("MOVQ", _VAR_gc_AX, _AX)
}

// This uses AX to pass ptr, thus AX's value will change
func (self *_Assembler) WritePtrAX(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, _AX)
    self.save(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierAX, _R10)  // MOVQ ${fn}, AX
    self.Rjmp("CALL", _R10)      
    self.load(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Link("_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
}

// This uses CX to pass ptr, thus CX's value will change
func (self *_Assembler) WritePtrCX(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, _CX)
    self.save(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierCX, _R10)  // MOVQ ${fn}, CX
    self.Rjmp("CALL", _R10)      
    self.load(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Link("_no_writeBarrier" + strconv.Itoa(i) + "_{n}")
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i) + "_{n}")
}

func (self *_ValueDecoder) print_ptr(i int, ptrs ...obj.Addr) {
    // self.Emit("MOVQ", _AX, _VAR_gc_AX)
    self.Emit("MOVQ", jit.Imm(int64(i)), _AX)
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))
    for i:=0; i<3; i++ {
        if i<len(ptrs) {
            self.Emit("MOVQ", ptrs[i],  _AX)      
            self.Emit("MOVQ", _AX,  jit.Ptr(_SP, int64(i*8+8)))  
        }else{
            self.Emit("MOVQ", jit.Imm(0),  jit.Ptr(_SP, int64(i*8+8)))  
        }
    }
    self.call_go(_F_print_ptr)
    // self.Emit("MOVQ", _VAR_gc_AX, _AX)
}

// This uses AX and R10 to pass ptr, thus their value will change
func (self *_ValueDecoder) WritePtrAx(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, _AX)
    self.load(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierAX, _R10)  
    self.Rjmp("CALL", _R10)  
    self.load(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i)) 
    self.Link("_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i))
}

// This uses CX and R10 to pass ptr, thus their value will change
func (self *_ValueDecoder) WritePtrCx(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, _CX)
    self.load(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierCX, _R10)  
    self.Rjmp("CALL", _R10)  
    self.load(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i))
    self.Link("_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i))
}

// This uses DX and R10 to pass ptr, thus their value will change
func (self *_ValueDecoder) WritePtrDx(i int, ptr obj.Addr, rec obj.Addr) {
    self.Emit("MOVQ", _V_writeBarrier, _R10)
    self.Emit("CMPL", jit.Ptr(_R10, 0), jit.Imm(0))
    self.Sjmp("JE", "_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, _DX)
    self.load(_DI)
    self.Emit("LEAQ", rec, _DI)
    self.Emit("MOVQ", _F_gcWriteBarrierDX, _R10)  // MOVQ ${fn}, CX
    self.Rjmp("CALL", _R10)  
    self.load(_DI)
    self.Sjmp("JMP", "_end_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", _DX, rec)
    self.Link("_no_writeBarrier" + strconv.Itoa(i))
    self.Emit("MOVQ", ptr, rec)
    self.Link("_end_writeBarrier" + strconv.Itoa(i))
}