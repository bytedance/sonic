/*
 * Copyright 2025 Huawei Technologies Co., Ltd.
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

	arm64 "github.com/bytedance/sonic/loader/internal/iasm/arm64"
)

// ARM64 Register definitions
// Based on ARM64 calling convention (AAPCS64)
type Register = arm64.Register
type Register64 = arm64.Register64
type VRegister = arm64.VRegister
type Program = arm64.Program
type MemoryOperand = arm64.MemoryOperand
type Label = arm64.Label

var (
	Ptr         = arm64.Ptr
	DefaultArch = arm64.DefaultArch
	CreateLabel = arm64.CreateLabel

	intType = reflect.TypeOf(0)
	ptrType = reflect.TypeOf(unsafe.Pointer(nil))
)

// Address modes
const (
	AddrModePreIndex  = arm64.AddrModePreIndex
	AddrModePostIndex = arm64.AddrModePostIndex
)

// ARM64 general purpose registers
const (
	X0  = arm64.X0
	X1  = arm64.X1
	X2  = arm64.X2
	X3  = arm64.X3
	X4  = arm64.X4
	X5  = arm64.X5
	X6  = arm64.X6
	X7  = arm64.X7
	X8  = arm64.X8
	X9  = arm64.X9
	X10 = arm64.X10
	X11 = arm64.X11
	X12 = arm64.X12
	X13 = arm64.X13
	X14 = arm64.X14
	X15 = arm64.X15
	X16 = arm64.X16
	X17 = arm64.X17
	X18 = arm64.X18
	X19 = arm64.X19
	X20 = arm64.X20
	X21 = arm64.X21
	X22 = arm64.X22
	X23 = arm64.X23
	X24 = arm64.X24
	X25 = arm64.X25
	X26 = arm64.X26
	X27 = arm64.X27
	X28 = arm64.X28
	X29 = arm64.X29
	X30 = arm64.X30
	SP  = arm64.SP
	XZR = arm64.XZR
)

// ARM64 NEON/FP registers
const (
	V0  = arm64.V0
	V1  = arm64.V1
	V2  = arm64.V2
	V3  = arm64.V3
	V4  = arm64.V4
	V5  = arm64.V5
	V6  = arm64.V6
	V7  = arm64.V7
	V8  = arm64.V8
	V9  = arm64.V9
	V10 = arm64.V10
	V11 = arm64.V11
	V12 = arm64.V12
	V13 = arm64.V13
	V14 = arm64.V14
	V15 = arm64.V15
	V16 = arm64.V16
	V17 = arm64.V17
	V18 = arm64.V18
	V19 = arm64.V19
	V20 = arm64.V20
	V21 = arm64.V21
	V22 = arm64.V22
	V23 = arm64.V23
	V24 = arm64.V24
	V25 = arm64.V25
	V26 = arm64.V26
	V27 = arm64.V27
	V28 = arm64.V28
	V29 = arm64.V29
	V30 = arm64.V30
	V31 = arm64.V31
)

// Aliases
const (
	FP = arm64.FP
	LR = arm64.LR
)

const (
	PtrSize  = 8 // pointer size on ARM64
	PtrAlign = 8 // pointer alignment on ARM64
)

// ARM64 C calling convention (AAPCS64)
// Arguments are passed in X0-X7 for integers/pointers
// V0-V7 for floating point
var iregOrderC = []Register64{
	X0, X1, X2, X3, X4, X5, X6, X7,
}

var vregOrderC = []VRegister{
	V0, V1, V2, V3, V4, V5, V6, V7,
}

// ARM64 Go calling convention (ABIInternal)
// According to Go internal ABI spec, ARM64 uses R0-R15 (X0-X15) for integer arguments
// and F0-F15 (V0-V15) for floating-point arguments
// We support up to 16 integer and 16 float registers for Go ABI
var iregOrderGo = []Register64{
	X0, X1, X2, X3, X4, X5, X6, X7,
	X8, X9, X10, X11, X12, X13, X14, X15,
}

var vregOrderGo = []VRegister{
	V0, V1, V2, V3, V4, V5, V6, V7,
	V8, V9, V10, V11, V12, V13, V14, V15,
}

type stackAlloc struct {
	s uint32
	i int
	x int
}

func (self *stackAlloc) reset() {
	self.i, self.x = 0, 0
}

func (self *stackAlloc) ireg(vt reflect.Type) (p Parameter) {
	p = mkIReg(vt, iregOrderGo[self.i])
	self.i++
	return
}

func (self *stackAlloc) xreg(vt reflect.Type) (p Parameter) {
	p = mkVReg(vt, vregOrderGo[self.x])
	self.x++
	return
}

func (self *stackAlloc) stack(vt reflect.Type) (p Parameter) {
	p = mkStack(vt, self.s)
	self.s += uint32(vt.Size())
	return
}

func (self *stackAlloc) spill(n uint32, a int) uint32 {
	self.s = alignUp(self.s, a) + n
	return self.s
}

func (self *stackAlloc) alloc(p []Parameter, vt reflect.Type) []Parameter {
	nb := vt.Size()
	vk := vt.Kind()

	/* zero-sized objects are allocated on stack */
	if nb == 0 {
		return append(p, mkStack(intType, self.s))
	}

	/* check for value type */
	switch vk {
	case reflect.Bool:
		return self.valloc(p, reflect.TypeOf(false))
	case reflect.Int:
		return self.valloc(p, intType)
	case reflect.Int8:
		return self.valloc(p, reflect.TypeOf(int8(0)))
	case reflect.Int16:
		return self.valloc(p, reflect.TypeOf(int16(0)))
	case reflect.Int32:
		return self.valloc(p, reflect.TypeOf(int32(0)))
	case reflect.Int64:
		return self.valloc(p, reflect.TypeOf(int64(0)))
	case reflect.Uint:
		return self.valloc(p, reflect.TypeOf(uint(0)))
	case reflect.Uint8:
		return self.valloc(p, reflect.TypeOf(uint8(0)))
	case reflect.Uint16:
		return self.valloc(p, reflect.TypeOf(uint16(0)))
	case reflect.Uint32:
		return self.valloc(p, reflect.TypeOf(uint32(0)))
	case reflect.Uint64:
		return self.valloc(p, reflect.TypeOf(uint64(0)))
	case reflect.Uintptr:
		return self.valloc(p, reflect.TypeOf(uintptr(0)))
	case reflect.Float32:
		return self.valloc(p, reflect.TypeOf(float32(0)))
	case reflect.Float64:
		return self.valloc(p, reflect.TypeOf(float64(0)))
	case reflect.Complex64:
		panic("abi: arm64: not implemented: complex64")
	case reflect.Complex128:
		panic("abi: arm64: not implemented: complex128")
	case reflect.Array:
		panic("abi: arm64: not implemented: arrays")
	case reflect.Chan:
		return self.valloc(p, reflect.TypeOf((chan int)(nil)))
	case reflect.Func:
		return self.valloc(p, reflect.TypeOf((func())(nil)))
	case reflect.Map:
		return self.valloc(p, reflect.TypeOf((map[int]int)(nil)))
	case reflect.Ptr:
		return self.valloc(p, reflect.TypeOf((*int)(nil)))
	case reflect.UnsafePointer:
		return self.valloc(p, ptrType)
	case reflect.Interface:
		return self.valloc(p, ptrType, ptrType)
	case reflect.Slice:
		return self.valloc(p, ptrType, intType, intType)
	case reflect.String:
		return self.valloc(p, ptrType, intType)
	case reflect.Struct:
		panic("abi: arm64: not implemented: structs")
	default:
		panic("abi: invalid value type")
	}
}

func (self *stackAlloc) valloc(p []Parameter, vts ...reflect.Type) []Parameter {
	for _, vt := range vts {
		enum := isFloat(vt)
		if enum != notFloatKind && self.x < len(vregOrderGo) {
			p = append(p, self.xreg(vt))
		} else if enum == notFloatKind && self.i < len(iregOrderGo) {
			p = append(p, self.ireg(vt))
		} else {
			p = append(p, self.stack(vt))
		}
	}
	return p
}

func NewFunctionLayout(ft reflect.Type) FunctionLayout {
	var sa stackAlloc
	var fn FunctionLayout

	/* assign every arguments */
	for i := 0; i < ft.NumIn(); i++ {
		fn.Args = sa.alloc(fn.Args, ft.In(i))
	}

	/* reset the register counter, and add a pointer alignment field */
	sa.reset()

	/* assign every return value */
	for i := 0; i < ft.NumOut(); i++ {
		fn.Rets = sa.alloc(fn.Rets, ft.Out(i))
	}

	sa.spill(0, PtrAlign)

	/* assign spill slots */
	for i := 0; i < len(fn.Args); i++ {
		if fn.Args[i].InRegister {
			fn.Args[i].Mem = sa.spill(PtrSize, PtrAlign) - PtrSize
		}
	}

	/* add the final pointer alignment field */
	fn.FP = sa.spill(0, PtrAlign)
	return fn
}

func (self *Frame) emitGrowStack(p *Program, entry *Label) {
	// ARM64 morestack calling convention:
	// - morestack expects the caller's return address in R3 register (not on stack!)
	// - We must spill register arguments to their spill slots in caller's frame
	// - Spill slots are at POSITIVE offsets from SP (in caller's frame)
	//
	// Stack layout when we arrive here (BEFORE prologue):
	// +------------------+
	// | args from caller |  <- Stack arguments (if any)
	// +------------------+ <- SP points here
	// | spill slots      |  <- At positive offsets: SP+0, SP+8, SP+16...
	// +------------------+
	//
	// LR (X30) contains our return address - morestack needs it in R3!

	// Step 1: Spill register arguments to their spill slots in caller's frame FIRST
	// These slots are at POSITIVE offsets from current SP
	// IMPORTANT: Must save args BEFORE moving LR to R3, because R3 might contain arg data!
	for _, v := range self.desc.Args {
		if v.InRegister {
			// v.Mem contains the spill slot offset in caller's frame
			offset := int32(v.Mem) + 8
			if v.IsFloat == floatKind64 {
				p.STR(arm64.DRegister(v.Reg.(VRegister)), Ptr(SP, offset))
			} else if v.IsFloat == floatKind32 {
				p.STR(arm64.SRegister(v.Reg.(VRegister)), Ptr(SP, offset))
			} else {
				p.STR(v.Reg.(Register64), Ptr(SP, offset))
			}
		}
	}

	// Step 2: Move LR to R3 (morestack calling convention)
	// morestack will save R3 to g->sched->gobuf_lr
	// NOW it's safe to move X30 to X3 since all args (including X3 if it was an arg) are saved
	p.MOV(X3, X30)

	// Step 3: Call runtime.morestack_noctxt
	// It will:
	// - Read return address from R3
	// - Save current goroutine state (SP, FP, PC from LR, LR from R3)
	// - Allocate new stack and copy old stack contents
	// - Switch to new stack
	// - Return to this point (on the new stack)
	p.MOVZ(X16, uint16(F_morestack_noctxt&0xffff), 0)
	p.MOVK(X16, uint16((F_morestack_noctxt>>16)&0xffff), 16)
	p.MOVK(X16, uint16((F_morestack_noctxt>>32)&0xffff), 32)
	p.MOVK(X16, uint16((F_morestack_noctxt>>48)&0xffff), 48)
	p.BLR(X16)

	// Step 4: After morestack returns, we're on the new stack
	// Reload register arguments from their spill slots
	for _, v := range self.desc.Args {
		if v.InRegister {
			offset := int32(v.Mem) + 8
			if v.IsFloat == floatKind64 {
				p.LDR(arm64.DRegister(v.Reg.(VRegister)), Ptr(SP, offset))
			} else if v.IsFloat == floatKind32 {
				p.LDR(arm64.SRegister(v.Reg.(VRegister)), Ptr(SP, offset))
			} else {
				p.LDR(v.Reg.(Register64), Ptr(SP, offset))
			}
		}
	}

	// Step 5: Jump back to function entry to retry with the grown stack
	p.B(entry)
}

func (self *Frame) GrowStackTextSize() uint32 {
	// Generate actual grow stack code and measure its size
	// This matches what emitGrowStack produces
	p := DefaultArch.CreateProgram()
	entry := CreateLabel("_entry")
	p.Link(entry)
	self.emitGrowStack(p, entry)
	return uint32(len(p.Assemble(0)))
}

func (self *Frame) emitPrologue(p *Program, maxStack uintptr) {
	// ARM64 standard prologue (following Go compiler conventions):
	// 1. STR X30, [SP, #-frameSize]!   (pre-index: allocate stack and save LR)
	// 2. STUR X29, [SP, #-8]           (save old FP)
	// 3. SUB X29, SP, #8               (set new FP)

	// Step 1: Allocate stack and save X30, X29 together
	// Standard Go uses: STR.W X30, [SP, #-frameSize]! then STUR X29, [SP, #-8]
	// But we save both inside the frame for simplicity
	mem := &arm64.MemoryOperand{
		Base:   SP,
		Offset: -int32(self.Size()),
		Mode:   arm64.AddrModePreIndex,
	}
	p.STR(X30, mem)
	p.STR(X29, Ptr(SP, -8)) // Save X29 (FP) at SP-8

	// Step 3: Set new FP = SP - 8 (pointing to saved old FP)
	// This allows FP chain: [FP] points to previous FP
	p.SUB(X29, SP, 8)
}

func (self *Frame) emitEpilogue(p *Program, maxStack uintptr) {
	// ARM64 standard epilogue (following Go compiler conventions):
	// 1. LDP (X29, X30), [SP, #...]    (restore FP and LR)
	// 2. ADD SP, SP, #frameSize        (deallocate stack)
	// 3. RET

	// Restore X29 (FP) and X30 (LR) from stack
	p.LDR(X30, Ptr(SP, 0))  // Restore X30 from SP-0
	p.LDR(X29, Ptr(SP, -8)) // Restore X29 from SP-8
	// Deallocate stack space
	p.ADD(SP, SP, int(self.Size()))
	// RET
	p.RET()
}

// ReservedRegs returns the list of registers that need to be preserved
// For ARM64, we need to preserve:
// - X29 (FP): Frame Pointer - required for stack unwinding and debugging
// - X30 (LR): Link Register - stores return address, clobbered by BL/BLR
// - X28 (g): Current goroutine pointer - critical for Go runtime
func ReservedRegs(callc bool) []Register {
	return []Register{
		X28, // current goroutine (g)
	}
}

func (self *Frame) emitReserveRegs(p *Program) {
	// Save X28 (g pointer) to stack
	// X30 and X29 are already saved in emitPrologue
	// X30 is at [SP-0], X29 is at [SP+8], X29 points to [SP+8]
	// Save X28 at [SP-16]
	//p.STR(X28, Ptr(SP, 8)) // asm2asm will used fixed-x28
}

func (self *Frame) emitRestoreRegs(p *Program) {
	// Restore X28 (g pointer) from stack
	// X30 and X29 will be restored in emitEpilogue
	// p.LDR(X28, Ptr(SP, 8))
}

func (self *Frame) emitStackCheck(p *Program, to *Label, maxStack uintptr) {
	// ARM64 stack check: compare SP - frameSize with g.stackguard0
	// g is in X28, stackguard0 is at offset 16 in g struct
	// Note: frameSize must include space for reserved registers and maxStack
	baseFrameSize := self.Size()
	// Add maxStack for C function call stack space
	// frameSize must match what prologue allocates
	frameSize := alignUp(baseFrameSize+uint32(maxStack), 0x10)
	p.LDR(X16, Ptr(X28, 0x10)) // load g.stackguard0
	p.SUB(X17, SP, int(frameSize))
	p.CMP(X17, X16) // compare calculated SP with stackguard0 (X9 - X10, sets flags)
	p.BLS(to)       // branch if calculated SP < stackguard0 (unsigned less than)
}

func (self *Frame) emitExchangeArgs(p *Program) {

}

func (self *Frame) emitExchangeRets(p *Program) {
	// ARM64 Go ABI and C ABI (AAPCS64) return value conventions:
	// Both use:
	// - Integer/pointer returns in X0-X7 (first return in X0, second in X1, etc.)
	// - Float returns in V0-V7 (first return in V0, second in V1, etc.)
	//
	// Since both ABIs use the same registers for return values, no exchange is needed.
	// Return values are already in the correct registers after the C function returns.
	//
	// Note: This is different from AMD64 where some register mapping may be needed.
	// For ARM64, the register-based return convention is compatible between Go and C.

	// No operation needed - return values in X0, V0, etc. are already correct
}

func (self *Frame) emitCallC(p *Program, addr uintptr) {
	// Load C function address into X16 and call it
	// Using MOVZ + MOVK sequence to load 64-bit address
	p.MOVZ(X16, uint16(addr&0xffff), 0)
	p.MOVK(X16, uint16((addr>>16)&0xffff), 16)
	p.MOVK(X16, uint16((addr>>32)&0xffff), 32)
	p.MOVK(X16, uint16((addr>>48)&0xffff), 48)
	// BLR X16 (branch with link to register)
	p.BLR(X16)
}

type floatKind uint8

const (
	notFloatKind floatKind = iota
	floatKind32
	floatKind64
)

type Parameter struct {
	InRegister bool
	IsPointer  bool
	IsFloat    floatKind
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

func isFloat(vt reflect.Type) floatKind {
	switch vt.Kind() {
	case reflect.Float32:
		return floatKind32
	case reflect.Float64:
		return floatKind64
	default:
		return notFloatKind
	}
}

func mkVReg(vt reflect.Type, reg VRegister) (p Parameter) {
	p.Reg = reg
	p.Type = vt
	p.InRegister = true
	p.IsFloat = isFloat(vt)
	return
}

func mkStack(vt reflect.Type, mem uint32) (p Parameter) {
	p.Mem = mem
	p.Type = vt
	p.InRegister = false
	p.IsPointer = isPointer(vt)
	p.IsFloat = isFloat(vt)
	return
}

func (self Parameter) String() string {
	if self.InRegister {
		return fmt.Sprintf("[%s, Pointer(%v), Float(%v)]", self.Reg, self.IsPointer, self.IsFloat)
	} else {
		return fmt.Sprintf("[%d(SP), Pointer(%v), Float(%v)]", self.Mem, self.IsPointer, self.IsFloat)
	}
}

func CallC(addr uintptr, fr Frame, maxStack uintptr) []byte {
	// Generate ARM64 machine code that wraps a C function call
	// Following AMD64 implementation pattern for consistency
	p := DefaultArch.CreateProgram()

	stack := CreateLabel("_stack_grow")
	entry := CreateLabel("_entry")

	// Entry point
	p.Link(entry)
	fr.emitStackCheck(p, stack, maxStack)
	fr.emitPrologue(p, maxStack)

	fr.emitReserveRegs(p)
	fr.emitCallC(p, addr)
	fr.emitExchangeRets(p)
	fr.emitRestoreRegs(p)
	fr.emitEpilogue(p, maxStack)

	// Always link stack growth path like AMD64 version does
	p.Link(stack)
	fr.emitGrowStack(p, entry)

	return p.Assemble(0)
}

func (self *Frame) StackCheckTextSize() uint32 {
	// Generate actual stack check code and measure its size
	// This is used to build correct PCSP (PC to Stack Pointer) mapping table
	// Note: This must match the exact code generated by emitStackCheck
	// Including maxStack in calculation to match runtime behavior
	p := DefaultArch.CreateProgram()
	to := CreateLabel("")
	p.Link(to)
	self.emitStackCheck(p, to, 0)
	return uint32(len(p.Assemble(0)))
}
