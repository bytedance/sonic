/*
 * Copyright 2022 ByteDance Inc.
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

package abi

import (
	"fmt"
	"reflect"
	"unsafe"

	. "github.com/chenzhuoyu/iasm/x86_64"
)

const (
    PtrSize  = 8    // pointer size
    PtrAlign = 8    // pointer alignment
)

var iregOrderC = []Register{
    RDI, 
    RSI, 
    RDX, 
    RCX,
    R8, 
    R9,
}

var xregOrderC = []Register{
    XMM0,
    XMM1,
    XMM2,
    XMM3,
    XMM4,
    XMM5,
    XMM6,
    XMM7,
}

var (
    intType = reflect.TypeOf(0)
    ptrType = reflect.TypeOf(unsafe.Pointer(nil))
)

/** Frame Structure of the Generated Function
    FP  +------------------------------+
        |             . . .            |
        | 2nd reg argument spill space |
        + 1st reg argument spill space |
        | <pointer-sized alignment>    |
        |             . . .            |
        | 2nd stack-assigned result    |
        + 1st stack-assigned result    |
        | <pointer-sized alignment>    |
        |             . . .            |
        | 2nd stack-assigned argument  |
        | 1st stack-assigned argument  |
        | stack-assigned receiver      |
prev()  +------------------------------+ (Previous Frame)
                Return PC              |
size()  -------------------------------|
               Saved RBP               |
offs()  -------------------------------|
           1th Reserved Registers      |
        -------------------------------|
           2th Reserved Registers      |
        -------------------------------|
           Local Variables             |
    RSP -------------------------------|↓ lower addresses
*/

func (self *Frame) Argv(i int) *MemoryOperand {
    return Ptr(RSP, int32(self.Prev() + self.desc.Args[i].Mem))
}

// Spillv is used for growstack spill registers
func (self *Frame) Spillv(i int) *MemoryOperand {
    // remain one slot for caller return pc
    return Ptr(RSP, PtrSize + int32(self.desc.Args[i].Mem))
}

func (self *Frame) Retv(i int) *MemoryOperand {
    return Ptr(RSP, int32(self.Prev() + self.desc.Rets[i].Mem))
}

func (self *Frame) Resv(i int) *MemoryOperand {
    return Ptr(RSP, int32(self.Offs() - uint32((i+1) * PtrSize)))
}

func (self *Frame) emitGrowStack(p *Program, entry *Label) {
    // spill all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            if v.IsFloat {
                p.MOVSD(v.Reg, self.Spillv(i))
            } else {
                p.MOVQ(v.Reg, self.Spillv(i))
            }
        }
    }

    // call runtime.morestack_noctxt
    p.MOVQ(F_morestack_noctxt, R12)
    p.CALLQ(R12)
    // load all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            if v.IsFloat {
                p.MOVSD(self.Spillv(i), v.Reg)
            } else {
                p.MOVQ(self.Spillv(i), v.Reg)
            }
        }
    }

    // jump back to the function entry
    p.JMP(entry)
}

func (self *Frame) GrowStackTextSize() uint32 {
    p := DefaultArch.CreateProgram()
    // spill all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            if v.IsFloat {
                p.MOVSD(v.Reg, self.Spillv(i))
            } else {
                p.MOVQ(v.Reg, self.Spillv(i))
            }
        }
    }

    // call runtime.morestack_noctxt
    p.MOVQ(F_morestack_noctxt, R12)
    p.CALLQ(R12)
    // load all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            if v.IsFloat {
                p.MOVSD(self.Spillv(i), v.Reg)
            } else {
                p.MOVQ(self.Spillv(i), v.Reg)
            }
        }
    }

    // jump back to the function entry
    l := CreateLabel("")
    p.Link(l)
    p.JMP(l)

    return uint32(len(p.Assemble(0)))
}

func (self *Frame) emitPrologue(p *Program) {
    p.SUBQ(self.Size(), RSP)
    p.MOVQ(RBP, Ptr(RSP, int32(self.Offs())))
    p.LEAQ(Ptr(RSP, int32(self.Offs())), RBP)
}

func (self *Frame) emitEpilogue(p *Program) {
    p.MOVQ(Ptr(RSP, int32(self.Offs())), RBP)
    p.ADDQ(self.Size(), RSP)
    p.RET()
}

func (self *Frame) emitSpillRegs(p *Program) {
    // spill reserved registers
    for i, r := range ReservedRegs {
        p.MOVQ(r, self.Resv(i))
    }
    // spill pointer argument registers
    for i, r := range self.desc.Args {
        if r.InRegister && r.IsPointer {
            p.MOVQ(r.Reg, self.Argv(i))
        }
    }
}

func (self *Frame) emitLoadRegs(p *Program) {
    // load reserved registers
    for i, r := range ReservedRegs {
        p.MOVQ(self.Resv(i), r)
    }
}

func (self *Frame) emitCallC(p *Program, addr uintptr) {
    p.MOVQ(addr, RAX)
    p.CALLQ(RAX)
}

func (self *Frame) emitDebug(p *Program) {
    p.INT(3)
}

type Parameter struct {
    InRegister bool
    IsPointer  bool
    IsFloat    bool
    Reg        Register
    Mem        uint32
    Type       reflect.Type
}

func mkIReg(vt reflect.Type, reg Register64) (p Parameter) {
    p.Reg = reg
    p.Type = vt
    p.InRegister = true
    p.IsPointer = isPointer(vt)
    return
}

func mkXReg(vt reflect.Type, reg XMMRegister) (p Parameter) {
    p.Reg = reg
    p.Type = vt
    p.InRegister = true
    p.IsFloat = true
    return
}

func mkStack(vt reflect.Type, mem uint32) (p Parameter) {
    p.Mem = mem
    p.Type = vt
    p.InRegister = false
    p.IsPointer = isPointer(vt)
    return
}

func (self Parameter) String() string {
    if self.InRegister {
        return fmt.Sprintf("[%%%s, Pointer(%v), Float(%v)]", self.Reg, self.IsPointer, self.IsFloat)
    } else {
        return fmt.Sprintf("[%d(FP), Pointer(%v), Float(%v)]", self.Mem, self.IsPointer, self.IsFloat)
    }
}

func CallC(addr uintptr, fr Frame, maxStack uintptr) []byte {
    p := DefaultArch.CreateProgram()

    stack := CreateLabel("_stack_grow")
    entry := CreateLabel("_entry")
    p.Link(entry)
    fr.emitStackCheck(p, stack, maxStack)
    fr.emitPrologue(p)
    fr.emitSpillRegs(p)
    fr.emitExchangeArgs(p)
    fr.emitCallC(p, addr)
    fr.emitExchangeRets(p)
    fr.emitLoadRegs(p)
    fr.emitEpilogue(p)
    p.Link(stack)
    fr.emitGrowStack(p, entry)

    return p.Assemble(0)
}