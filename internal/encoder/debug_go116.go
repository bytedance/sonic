//go:build go1.16 && !go1.17
// +build go1.16,!go1.17

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
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/twitchyliquid64/golang-asm/obj"
)

var (
    debugSyncGC  = os.Getenv("SONIC_SYNC_GC") != ""
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

var (
    _Instr_End _Instr = newInsOp(_OP_null)

    _F_gc       = jit.Func(runtime.GC)
    _F_force_gc = jit.Func(debug.FreeOSMemory)
    _F_println  = jit.Func(println_wrapper)
)

func println_wrapper(i int, op1 int, op2 int){
    println(i, " Intrs ", op1, _OpNames[op1], "next: ", op2, _OpNames[op2])
}

func (self *_Assembler) force_gc() {
    self.call_go(_F_gc)
    self.call_go(_F_force_gc)
}

func (self *_Assembler) debug_instr(i int, v *_Instr) {
    if debugSyncGC {
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
}


var (
    debug_OOM int
    _F_debug_string = jit.Func(print_string)
)

func init() {
    val := os.Getenv("SONIC_DEBUG_OOM")
    iv, err := strconv.ParseInt(val, 10, 64)
    if err == nil {
        debug_OOM = int(iv)
    }
}

func (self *_Assembler) debug_string(str obj.Addr) {
    if debug_OOM <= 0 {
        return
    }
    self.Emit("MOVQ", jit.Ptr(str, 8), _AX)
    self.Emit("CMPQ", _AX, jit.Imm(int64(debug_OOM)))
    self.Sjmp("JBE", "_debug_string_ret_{n}")
    self.Emit("MOVQ", _RP, jit.Ptr(_SP, 0))        
    self.Emit("MOVQ", _RL, jit.Ptr(_SP, 8))        
    self.Emit("MOVQ", str, jit.Ptr(_SP, 16))     
    self.call_go(_F_debug_string)
    self.Link("_debug_string_ret_{n}")
}

func print_string(curjson string, val *string) {
    println("[SONIC_DEBUG] Current output json: ", curjson)
    if val != nil {
        s := (*rt.GoString)(unsafe.Pointer(val))
        print("[SONIC_DEBUG] GoString Ptr:", s.Ptr, ", Len:", s.Len, ", first bytes: ")
        if s.Len <= 100 {
            println(*val)
        } else {
            println((*val)[:100])
        }
    } else {
        println("[SONIC_DEBUG] nil pointer")
    }
}
