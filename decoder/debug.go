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
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/twitchyliquid64/golang-asm/obj"
)

//WARN: MUST set false after release
var debugGC = false

var (
    _Instr_End _Instr = newInsOp(_OP_nil_1)

    _F_gc       = jit.Func(runtime.GC)
    _F_force_gc = jit.Func(debug.FreeOSMemory)
    _F_print_int  = jit.Func(println_int)
    _F_print_gc  = jit.Func(println_wrapper)
    _F_print_ptr  = jit.Func(printPtr)

    _REG_print_ptr = []obj.Addr{ _ST,_VP,_IP,_IL,_IC,_ET,_EP,_DF}
)

var printCount = 0

func println_int(){
    println("gc: ", printCount)
    printCount++
}

func (self *_Assembler) print_int() {
    self.call_go(_F_print_int)
}

func println_wrapper(i int, op1 int, op2 int){
    println(i, " Intrs ", op1, _OpNames[op1], "next: ", op2, _OpNames[op2])
}

func (self *_Assembler) print_gc(i int, p1 *_Instr, p2 *_Instr) {
    self.Emit("MOVQ", jit.Imm(int64(p2.op())),  jit.Ptr(_SP, 16))// MOVQ $(p2.op()), 16(SP)
    self.Emit("MOVQ", jit.Imm(int64(p1.op())),  jit.Ptr(_SP, 8)) // MOVQ $(p1.op()), 8(SP)
    self.Emit("MOVQ", jit.Imm(int64(i)),  jit.Ptr(_SP, 0))       // MOVQ $(i), (SP)
    self.call_go(_F_print_gc)
}

func (self *_Assembler) force_gc() {
    if debugGC {
        self.print_int()
        self.call_go(_F_gc)
        self.call_go(_F_force_gc)
    }
}

func (self *_Assembler) append_gc(i int, v *_Instr) {
    if (i+1 == len(self.p)) {
        self.print_gc(i, v, &_Instr_End) 
    } else {
        next := &(self.p[i+1])
        self.print_gc(i, v, next)
        name := _OpNames[next.op()]
        if strings.Contains(name, "save") {
            return
        }
    }
    self.force_gc()
}

func (self *_Assembler) debug_instr(i int, v *_Instr) {
    if debugGC {
        self.append_gc(i, v)
    }
}

func (self *_ValueDecoder) force_gc() {
    if debugGC {
        self.call_go(_F_gc)
        self.call_go(_F_force_gc)
    }
}

func (self *_ValueDecoder) debug_gc(i int) {
    if debugGC {
        self.force_gc()
    }
}

func printPtr(i int, ptrs [3]uintptr) {
    fmt.Printf("%d: [", i)
    for _, ptr := range ptrs {
        fmt.Printf("%d, ", ptr)
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
    self.save(_REG_print_ptr...)   // SAVE $REG_go
    self.call(_F_print_ptr)
    self.load(_REG_print_ptr...)   // SAVE $REG_go
    // self.Emit("MOVQ", _VAR_gc_AX, _AX)
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
    self.save(_REG_print_ptr...)   // SAVE $REG_go
    self.call(_F_print_ptr)
    self.load(_REG_print_ptr...)   // SAVE $REG_go
    // self.Emit("MOVQ", _VAR_gc_AX, _AX)
}