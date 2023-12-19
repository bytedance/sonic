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
    "os"
    "runtime"
    "runtime/debug"
    "strings"
    "unsafe"

    "github.com/bytedance/sonic/internal/jit"
    "github.com/twitchyliquid64/golang-asm/obj"
)

var (
    debugSyncGC  = os.Getenv("SONIC_SYNC_GC") != ""
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
    debugCheckPtr = os.Getenv("SONIC_CHECK_POINTER") == ""
)

const _FP_debug  = 0   //NOTICE: for debug = 8*len(_REG_debug)

var  _REG_debug = []obj.Addr { _ST, _VP, _IP, _IL, _IC, _ET, _EP, _AX, _BX, _CX, _DX, _DI, _SI, _R8, _R9, jit.Reg("R10")}

var (
    _Instr_End _Instr = newInsOp(_OP_nil_1)

    _F_gc       = jit.Func(gc)
    _F_println  = jit.Func(println_wrapper)
    _F_print    = jit.Func(print)
)

func (self *_Assembler) dsave(r ...obj.Addr) {
    for i, v := range r {
        if i > _FP_debug / 8 - 1 {
            panic("too many registers to save")
        } else {
            self.Emit("MOVQ", v, jit.Ptr(_SP, _FP_fargs + _FP_saves + _FP_locals + int64(i) * 8))
        }
    }
}

func (self *_Assembler) dload(r ...obj.Addr) {
    for i, v := range r {
        if i > _FP_debug / 8 - 1 {
            panic("too many registers to load")
        } else {
            self.Emit("MOVQ", jit.Ptr(_SP, _FP_fargs + _FP_saves + _FP_locals + int64(i) * 8), v)
        }
    }
}

func (self *_ValueDecoder) dsave(r ...obj.Addr) {
    for i, v := range r {
        if i > _FP_debug / 8 - 1 {
            panic("too many registers to save")
        } else {
            self.Emit("MOVQ", v, jit.Ptr(_SP, _VD_fargs  + _VD_fargs + _VD_locals + int64(i) * 8))
        }
    }
}

func (self *_ValueDecoder) dload(r ...obj.Addr) {
    for i, v := range r {
        if i > _FP_debug / 8 - 1 {
            panic("too many registers to load")
        } else {
            self.Emit("MOVQ", jit.Ptr(_SP, _VD_fargs + _VD_fargs + _VD_locals + int64(i) * 8), v)
        }
    }
}

func println_wrapper(i int, op1 int, op2 int){
    println(i, " Intrs ", op1, _OpNames[op1], "next: ", op2, _OpNames[op2])
}

func print(i int){
    println(i)
}

func gc() {
    if !debugSyncGC {
        return
    }
    runtime.GC()
    debug.FreeOSMemory()
}

func (self *_Assembler) dcall(fn obj.Addr) {
    self.Emit("MOVQ", fn, jit.Reg("R10"))  // MOVQ ${fn}, R10
    self.Rjmp("CALL", jit.Reg("R10"))       // CALL R10
}

func (self *_Assembler) debug_gc() {
    if !debugSyncGC {
        return
    }
    self.dsave(_REG_debug...)
    self.dcall(_F_gc)
    self.dload(_REG_debug...)
}

func (self *_ValueDecoder) dcall(fn obj.Addr) {
    self.Emit("MOVQ", fn, jit.Reg("R10"))  // MOVQ ${fn}, R10
    self.Rjmp("CALL", jit.Reg("R10"))       // CALL R10
}

func (self *_ValueDecoder) debug_gc() {
    if !debugSyncGC {
        return
    }
    self.dsave(_REG_debug...)
    self.dcall(_F_gc)
    self.dload(_REG_debug...)
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
        // self.debug_gc()
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
    findobj(ptr)
}

func findobj(ptr uintptr) {
    _, s, objIndex := findObject(ptr, 0, 0)
    fmt.Printf("objIndex: %d\n", objIndex)
    if s == nil {
        fmt.Printf("! invalid pointer: %x\n", ptr)
    }
}

func (self *_Assembler) check_ptr(ptr obj.Addr, lea bool) {
    if !debugCheckPtr {
        return
    }

    self.dsave(_REG_debug...)
    if lea {
        self.Emit("LEAQ", ptr, _AX)
    } else {
        self.Emit("MOVQ", ptr, _AX)
    }
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))
    self.dcall(_F_checkptr)  
    self.dload(_REG_debug...)
}

func (self *_ValueDecoder) check_ptr(ptr obj.Addr, lea bool) {
    if !debugCheckPtr {
        return
    }

    self.dsave(_REG_debug...)
    if lea {
        self.Emit("LEAQ", ptr, _AX)
    } else {
        self.Emit("MOVQ", ptr, _AX)
    }
    self.Emit("MOVQ", _AX, jit.Ptr(_SP, 0))
    self.dcall(_F_checkptr)  
    self.dload(_REG_debug...)
}

func printptr(i int, ptr uintptr) {
    fmt.Printf("[%d] ptr: %x\n", i, ptr)
}

func (self *_Assembler) print_ptr(i int, ptr obj.Addr, lea bool) {
    self.dsave(_REG_debug...)
    if lea {
        self.Emit("LEAQ", ptr, _BX)
    } else {
        self.Emit("MOVQ", ptr, _BX)
    }

    self.Emit("MOVQ", jit.Imm(int64(i)), _AX)
    self.Emit("MOVQ", _BX, jit.Ptr(_SP, 8))
    self.dcall(_F_printptr)  
    self.dload(_REG_debug...)
}

func (self *_ValueDecoder) print_ptr(i int, ptr obj.Addr, lea bool) {
    self.dsave(_REG_debug...)
    if lea {
        self.Emit("LEAQ", ptr, _BX)
    } else {
        self.Emit("MOVQ", ptr, _BX)
    }

    self.Emit("MOVQ", jit.Imm(int64(i)), _AX)
    self.Emit("MOVQ", _BX, jit.Ptr(_SP, 8))
    self.dcall(_F_printptr)  
    self.dload(_REG_debug...)
}