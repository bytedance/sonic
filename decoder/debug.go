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
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/twitchyliquid64/golang-asm/obj"
)


var (
    debugSyncGC  = os.Getenv("SONIC_SYNC_GC") != ""
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

var (
    _Instr_End _Instr = newInsOp(_OP_nil_1)

    _F_gc       = jit.Func(runtime.GC)
    _F_force_gc = jit.Func(debug.FreeOSMemory)
    _F_println  = jit.Func(println_wrapper)
    _F_print    = jit.Func(print)
)

var (
    _REG_debug = []obj.Addr { 
        jit.Reg("AX"),
        jit.Reg("BX"),
        jit.Reg("CX"),
        jit.Reg("DX"),
        jit.Reg("DI"),
        jit.Reg("SI"),
        jit.Reg("BP"),
        jit.Reg("SP"),
        jit.Reg("R8"),
        jit.Reg("R9"),
        jit.Reg("R10"),
        jit.Reg("R11"),
        jit.Reg("R12"),
        jit.Reg("R14"),
        jit.Reg("R13"),
        jit.Reg("R15"),
     }
)

func println_wrapper(i int, op1 int, op2 int){
    println(i, " Intrs ", op1, _OpNames[op1], "next: ", op2, _OpNames[op2])
}

func print(i int, a int, b int){
    println(i, a, b)
}

func (self *_Assembler) print(i int) {
    self.call_go(_F_gc)
    self.call_go(_F_force_gc)
}

func (self *_Assembler) call_debug(fn obj.Addr) {
    self.save(_REG_debug...)   
    self.call(fn)
    self.load(_REG_debug...)   
}

func (self *_Assembler) force_gc() {
    self.call_go(_F_gc)
    self.call_go(_F_force_gc)
}

func (self *_Assembler) debug_instr(i int, v *_Instr) {
    if debugSyncGC {
        if (i+1 == len(self.p)) {
            self.print_op(i, v, &_Instr_End) 
        } else {
            next := &(self.p[i+1])
            self.print_op(i, v, next)
            name := _OpNames[next.op()]
            if strings.Contains(name, "save") {
                return
            }
        }
        self.force_gc()
    }
}

func (self *_Assembler) print_op(i int, p1 *_Instr, p2 *_Instr) {
    self.Emit("MOVQ", jit.Imm(int64(p2.op())),  _CX)// MOVQ $(p2.op()), 16(SP)
    self.Emit("MOVQ", jit.Imm(int64(p1.op())),  _BX) // MOVQ $(p1.op()), 8(SP)
    self.Emit("MOVQ", jit.Imm(int64(i)),  _AX)       // MOVQ $(i), (SP)
    self.call_go(_F_println)
}

func (self *_Assembler) print_reg(i int, b, c obj.Addr) {
    self.save(_REG_debug...)   
    self.Emit("MOVQ", c,  _CX)
    self.Emit("MOVQ", b,  _BX) 
    self.Emit("MOVQ", jit.Imm(int64(i)),  _AX)       
    self.call(_F_print)
    self.load(_REG_debug...)   
}