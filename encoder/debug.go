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
	`fmt`
	`os`
	`runtime`
	`runtime/debug`
	`strings`
	`unsafe`

	`github.com/bytedance/sonic/internal/jit`
	`github.com/twitchyliquid64/golang-asm/obj`
)

var (
    debugSyncGC  = os.Getenv("SONIC_SYNC_GC") != ""
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
    debugCheckPtr = os.Getenv("SONIC_CHECK_POINTER") != ""
)

var (
    _Instr_End _Instr = newInsOp(_OP_null)

    _F_gc       = jit.Func(gc)
    _F_println  = jit.Func(println_wrapper)
)

func println_wrapper(i int, op1 int, op2 int){
    println(i, " Intrs ", op1, _OpNames[op1], "next: ", op2, _OpNames[op2])
}

func gc() {
	if !debugSyncGC {
        return
    }
	runtime.GC()
	debug.FreeOSMemory()
}

func (self *_Assembler) debug_gc() {
	if !debugSyncGC {
		return
	}
    self.xsave(_REG_checkptr...)
    self.call(_F_gc)
    self.xload(_REG_checkptr...)
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
        self.debug_gc()
    }
}

//go:noescape
//go:linkname checkptrBase runtime.checkptrBase
func checkptrBase(p unsafe.Pointer) uintptr

//go:noescape
//go:linkname findObject runtime.findObject
func findObject(p, refBase, refOff uintptr) (base uintptr, s unsafe.Pointer, objIndex uintptr)

var (
    _F_checkptr = jit.Func(checkptr)
    _F_printptr = jit.Func(printptr)
)

var _REG_checkptr = []obj.Addr {
    _ST, _SP_x, _SP_f, _SP_p, _SP_q, _RP, _RL, _RC, _EP, _ET,
    _AX,
    _CX,
    _R8,
    jit.Reg("R9"),
}

//go:nosplit
func checkptr(ptr uintptr) {
    if ptr == 0 {
        return
    }
    fmt.Printf("pointer: %x\n", ptr)
    f := checkptrBase(unsafe.Pointer(uintptr(ptr)))
    if f == 0 {
        fmt.Printf("! unknown-based pointer: %x\n", ptr)
    } else if f == 1 {
        fmt.Printf("! stack pointer: %x\n", ptr)
    } else {
        fmt.Printf("base: %x\n", f)
    }
    // findobj(ptr)
}

func findobj(ptr uintptr) {
    base, s, objIndex := findObject(ptr, 0, 0)
    if s != nil && base == 0 {
        fmt.Printf("! not in-heap pointer: %x\n", ptr)
    } else {
		fmt.Printf("objIndex: %d\n", objIndex)
		// if s != nil && s.isFree(objIndex) {
		// 	fmt.Printf("! obj(%x) is marked free!\n", ptr)
		// }
	}
}

func (self *_Assembler) check_ptr(ptr obj.Addr, lea bool) {
    if !debugCheckPtr {
        return
    }
    self.xsave(_REG_checkptr...)
    if lea {
        self.Emit("LEAQ", ptr, _R10)
    } else {
        self.Emit("MOVQ", ptr, _R10)
    }
    self.Emit("MOVQ", _R10, jit.Ptr(_SP, 0))
    self.Emit("MOVQ", _F_checkptr, _R10)
    self.Rjmp("CALL", _R10)  
    self.xload(_REG_checkptr...)
}

func printptr(i int, ptr uintptr) {
    fmt.Printf("[%d] ptr: %x\n", i, ptr)
}

func (self *_Assembler) print_ptr(i int, ptr obj.Addr, lea bool) {
    if !debugCheckPtr {
        return
    }
    self.xsave(_REG_checkptr...)
    if lea {
        self.Emit("LEAQ", ptr, _R10)
    } else {
        self.Emit("MOVQ", ptr, _R10)
    }
    self.Emit("MOVQ", jit.Imm(int64(i)), jit.Ptr(_SP, 0))
    self.Emit("MOVQ", _R10, jit.Ptr(_SP, 8))
    self.Emit("MOVQ", _F_printptr, _R10)
    self.Rjmp("CALL", _R10)  
    self.xload(_REG_checkptr...)
}