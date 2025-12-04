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

package jit

import (
	"encoding/binary"
	"strconv"
	"strings"
	"sync"

	"fmt"

	"github.com/bytedance/sonic/loader"
	"github.com/twitchyliquid64/golang-asm/obj"
	"github.com/twitchyliquid64/golang-asm/obj/arm64"
)

const (
	_LB_jump_pc = "_jump_pc_"
)

var (
	_DBG_ASM_OUTPUT = false
)

// condition codes for arm64
const (
	EQ = 0
	NE = 1
	HS = 2
	CS = 2
	LO = 3
	CC = 3
	MI = 4
	PL = 5
	VS = 6
	VC = 7
	HI = 8
	LS = 9
	GE = 10
	LT = 11
	GT = 12
	LE = 13
	AL = 14
	NV = 15
)

type BaseAssembler struct {
	i        int
	f        func()
	c        []byte
	Pcdata   loader.Pcdata
	o        sync.Once
	pb       *Backend
	xrefs    map[string][]*obj.Prog
	labels   map[string]*obj.Prog
	pendings map[string][]*obj.Prog
}

/** Instruction Encoders **/

func (self *BaseAssembler) NOP() *obj.Prog {
	p := self.pb.New()
	p.As = obj.ANOP
	self.pb.Append(p)
	return p
}

func (self *BaseAssembler) Mark(pc int) {
	self.i++
	self.Link(_LB_jump_pc + strconv.Itoa(pc))
}

func (self *BaseAssembler) Link(to string) {
	var p *obj.Prog
	var v []*obj.Prog

	/* placeholder substitution */
	if strings.Contains(to, "{n}") {
		to = strings.ReplaceAll(to, "{n}", strconv.Itoa(self.i))
	}

	/* check for duplications */
	if _, ok := self.labels[to]; ok {
		panic("label " + to + " has already been linked")
	}

	/* get the pending links */
	p = self.NOP()
	v = self.pendings[to]

	/* patch all the pending jumps */
	for _, q := range v {
		q.To.Val = p
	}

	/* mark the label as resolved */
	self.labels[to] = p
	delete(self.pendings, to)
}

func adr(rm int64, imm int64) int64 {
	var other int64 = 0x10000000 | rm
	var immhi int64 = imm << 3
	return adrBuild(other, immhi)
}

func adrBuild(other int64, immhi int64) int64 {
	return (other & 0xff00001f) | (immhi & 0x00ffffe0)
}

func adrInv(src int64) (other int64, immhi int64) {
	other = src & 0xff00001f
	immhi = src & 0x00ffffe0
	return
}

func (self *BaseAssembler) SrefRm(to string, d int64, rm int64) {
	p := self.pb.New()
	p.As = arm64.AWORD
	imm := adr(rm, d)
	p.To = Imm(imm)

	/* placeholder substitution */
	if strings.Contains(to, "{n}") {
		to = strings.ReplaceAll(to, "{n}", strconv.Itoa(self.i))
	}

	/* record the patch point */
	self.pb.Append(p)
	self.xrefs[to] = append(self.xrefs[to], p)
}

func (self *BaseAssembler) Xjmp(op string, to int) {
	self.Sjmp(op, _LB_jump_pc+strconv.Itoa(to))
}

func (self *BaseAssembler) Sjmp(op string, to string) {
	p := self.pb.New()
	p.As = As(op)

	/* placeholder substitution */
	if strings.Contains(to, "{n}") {
		to = strings.ReplaceAll(to, "{n}", strconv.Itoa(self.i))
	}

	/* check for backward jumps */
	if v, ok := self.labels[to]; ok {
		p.To.Val = v
	} else {
		self.pendings[to] = append(self.pendings[to], p)
	}

	/* mark as a branch, and add to instruction buffer */
	p.To.Type = obj.TYPE_BRANCH
	self.pb.Append(p)
}

func (self *BaseAssembler) Rjmp(op string, to obj.Addr) {
	p := self.pb.New()
	p.To = to
	p.To.Type = obj.TYPE_MEM
	p.As = As(op)
	self.pb.Append(p)
}

func (self *BaseAssembler) From(op string, val obj.Addr) {
	p := self.pb.New()
	p.As = As(op)
	p.From = val
	self.pb.Append(p)
}

func (self *BaseAssembler) EmitWithTwoSrcOps(op string, args ...obj.Addr) {
	p := self.pb.New()
	p.As = As(op)
	self.assignWithTwoSrcOps(p, args)
	self.pb.Append(p)
}

func (self *BaseAssembler) assignWithTwoSrcOps(p *obj.Prog, args []obj.Addr) {
	switch len(args) {
	case 0:
	case 1:
		p.From = args[0]
	case 2:
		p.Reg, p.From = args[1].Reg, args[0]
	case 3:
		p.To, p.From, p.Reg = args[2], args[0], args[1].Reg
	case 4:
		p.To, p.RegTo2, p.From, p.Reg = args[2], args[3].Reg, args[0], args[1].Reg
	default:
		panic("invalid operands")
	}
}

func (self *BaseAssembler) EmitRet() {
	p := self.pb.New()
	p.As = As("RET")
	p.To = obj.Addr{Type: obj.TYPE_REG, Reg: arm64.REG_R30}
	self.pb.Append(p)
}

func (self *BaseAssembler) EmitNOP() {
	p := self.pb.New()
	p.As = As("NOP")
	self.pb.Append(p)
}

func (self *BaseAssembler) EmitAdd(dest obj.Addr, op1 obj.Addr, op2 obj.Addr) {
	self.EmitWithThreeOps("ADD", dest, op1, op2)
}

func (self *BaseAssembler) EmitWithThreeOps(op string, dest obj.Addr, op1 obj.Addr, op2 obj.Addr) {
	p := self.pb.New()
	p.As = As(op)
	p.From = op2
	p.Reg = op1.Reg
	p.To = dest
	self.pb.Append(p)
}

func (self *BaseAssembler) EmitCmpq(plan9_op0 obj.Addr, plan9_op1 obj.Addr) {
	self.EmitWithTwoSrcOps("CMP", plan9_op0, plan9_op1)
}

func (self *BaseAssembler) EmitCmpqLiteral(literal obj.Addr, op1 obj.Addr) {
	self.EmitWithTwoSrcOps("CMP", literal, op1)
}

func (self *BaseAssembler) EmitCmpqLiteralLdr(literal obj.Addr, op1 obj.Addr, script_reg obj.Addr) {
	self.Emit("MOVD", op1, script_reg)
	self.EmitCmpqLiteral(literal, script_reg)
}

type JitPtrTt func(reg obj.Addr, offs int64) obj.Addr

func (self *BaseAssembler) EmitStur(ptr JitPtrTt, v obj.Addr, base obj.Addr, unscaled_offset int64) {
	self.Emit("MOVD", v, ptr(base, unscaled_offset))
}

func (self *BaseAssembler) EmitLdur(ptr JitPtrTt, base obj.Addr, unscaled_offset int64, v obj.Addr) {
	self.Emit("MOVD", ptr(base, unscaled_offset), v)
}

func (self *BaseAssembler) EmitCmpqLdr(plan9_op0 obj.Addr, plan9_op1 obj.Addr, script_reg obj.Addr) {
	self.EmitCmpqLdrPtr(plan9_op0, plan9_op1, script_reg)
}

func (self *BaseAssembler) EmitCmpqLdrPtr(plan9_op0 obj.Addr, plan9_op1 obj.Addr, script_reg obj.Addr) {
	self.Emit("MOVD", plan9_op0, script_reg)
	self.EmitCmpq(script_reg, plan9_op1)
}

func (self *BaseAssembler) EmitBrk(args ...obj.Addr) {
	self.EmitWithTwoSrcOps("BRK", args...)
}

func (self *BaseAssembler) EmitTbzSLdr(ptr JitPtrTt, op string, imm obj.Addr, from obj.Addr, to string, script_reg obj.Addr) {
	self.EmitLdur(ptr, from, 0, script_reg)
	self.EmitTbzS(op, imm, script_reg, to)
}

func (self *BaseAssembler) EmitTbzSLdriPtr(ptr JitPtrTt, op string, imm obj.Addr, from obj.Addr, to string, script_reg obj.Addr) {
	self.Emit("MOVD", from, script_reg)
	self.EmitTbzS(op, imm, script_reg, to)
}

func (self *BaseAssembler) EmitLdrri(ptr JitPtrTt, base obj.Addr, offset_r obj.Addr, offset_i int64, rst_r obj.Addr, script_reg obj.Addr) {
	self.EmitAdd(script_reg, base, offset_r)
	self.EmitLdur(ptr, script_reg, offset_i, rst_r)
}

func (self *BaseAssembler) EmitStrri(ptr JitPtrTt, rst_r obj.Addr, base obj.Addr, offset_r obj.Addr, offset_i int64, script_reg obj.Addr) {
	self.EmitAdd(script_reg, base, offset_r)
	self.EmitStur(ptr, rst_r, script_reg, offset_i)
}

func (self *BaseAssembler) EmitStrriWithInsn(ptr JitPtrTt, op string, rst_r obj.Addr, base obj.Addr, offset_r obj.Addr, offset_i int64, script_reg obj.Addr) {
	self.EmitAdd(script_reg, base, offset_r)
	self.Emit(op, rst_r, ptr(script_reg, offset_i))
}

// for CBZ/CBNZ from:reg to:label(Sjmp)
// (TESTQ R0, R0/CMP R0, 0+)JZ=CBZ(TESTQ R0, R0/CMP R0, 0+)+JNE=CBNZ
func (self *BaseAssembler) EmitCbzS(op string, from obj.Addr, to string) {
	p := self.pb.New()
	p.As = As(op)
	p.From = from

	/* placeholder substitution */
	if strings.Contains(to, "{n}") {
		to = strings.ReplaceAll(to, "{n}", strconv.Itoa(self.i))
	}

	/* check for backward jumps */
	if v, ok := self.labels[to]; ok {
		p.To.Val = v
	} else {
		self.pendings[to] = append(self.pendings[to], p)
	}

	/* mark as a branch, and add to instruction buffer */
	p.To.Type = obj.TYPE_BRANCH
	self.pb.Append(p)
}

// for CBZ/CBNZ from:reg to:label(Xjmp)
func (self *BaseAssembler) EmitCbzX(op string, from obj.Addr, to int) {
	self.EmitCbzS(op, from, _LB_jump_pc+strconv.Itoa(to))
}

// for TBZ/TBNZ from:reg to:label(Sjmp) imm:测试位 to:label
func (self *BaseAssembler) EmitTbzS(op string, imm obj.Addr, from obj.Addr, to string) {
	p := self.pb.New()
	p.As = As(op)
	p.From = imm
	p.Reg = from.Reg

	/* placeholder substitution */
	if strings.Contains(to, "{n}") {
		to = strings.ReplaceAll(to, "{n}", strconv.Itoa(self.i))
	}

	/* check for backward jumps */
	if v, ok := self.labels[to]; ok {
		p.To.Val = v
	} else {
		self.pendings[to] = append(self.pendings[to], p)
	}

	/* mark as a branch, and add to instruction buffer */
	p.To.Type = obj.TYPE_BRANCH
	self.pb.Append(p)
}

// for TBZ/TBNZ from:reg to:label(Sjmp) imm:测试位
// BTQ + JNC = TBZ, BTQ + JC = TBNZ
func (self *BaseAssembler) EmitTbzX(op string, imm obj.Addr, from obj.Addr, to int) {
	self.EmitTbzS(op, imm, from, _LB_jump_pc+strconv.Itoa(to))
}

func (self *BaseAssembler) Emit(op string, args ...obj.Addr) {
	p := self.pb.New()
	p.As = As(op)
	self.assignOperands(p, args)
	self.pb.Append(p)
}

// for TST from:reg
func (self *BaseAssembler) EmitTst(from obj.Addr, reg obj.Addr) {
	p := self.pb.New()
	p.As = As("TST")
	p.From = from
	p.Reg = reg.Reg
	self.pb.Append(p)
}

func (self *BaseAssembler) assignOperands(p *obj.Prog, args []obj.Addr) {
	switch len(args) {
	case 0:
	case 1:
		p.To = args[0]
	case 2:
		p.To, p.From = args[1], args[0]
	case 3:
		p.To, p.From, p.RestArgs = args[2], args[0], args[1:2]
	case 4:
		p.To, p.From, p.RestArgs = args[2], args[3], args[:2]
	default:
		panic("invalid operands")
	}
}

/** Assembler Helpers **/

func (self *BaseAssembler) Size() int {
	self.build()
	return len(self.c)
}

func (self *BaseAssembler) Init(f func()) {
	self.i = 0
	self.f = f
	self.c = nil
	self.o = sync.Once{}
}

var jitLoader = loader.Loader{
	Name: "sonic.jit.",
	File: "github.com/bytedance/sonic/jit.go",
	Options: loader.Options{
		NoPreempt: true,
	},
}

func (self *BaseAssembler) Load(name string, frameSize int, argSize int, argStackmap []bool, localStackmap []bool) loader.Function {
	self.build()
	return jitLoader.LoadOne(self.c, name, frameSize, argSize, argStackmap, localStackmap, self.Pcdata)
}

/** Assembler Stages **/

func (self *BaseAssembler) init() {
	self.pb = newBackend("arm64")
	self.xrefs = map[string][]*obj.Prog{}
	self.labels = map[string]*obj.Prog{}
	self.pendings = map[string][]*obj.Prog{}
}

func (self *BaseAssembler) build() {
	self.o.Do(func() {
		self.init()
		self.f()
		self.validate()
		self.assemble()
		self.resolve()
		self.debugAsmOutput()
		self.release()
	})
}

func (self *BaseAssembler) release() {
	self.pb.Release()
	self.pb = nil
	self.xrefs = nil
	self.labels = nil
	self.pendings = nil
}

func (self *BaseAssembler) resolve() {
	for s, v := range self.xrefs {
		for _, prog := range v {
			if prog.As != arm64.AWORD {
				panic("invalid RIP relative reference")
			} else if p, ok := self.labels[s]; !ok {
				panic("links are not fully resolved: " + s)
			} else {
				a := prog.To.Offset
				other, immho_a := adrInv(a)
				immhi := immho_a + ((p.Pc - prog.Pc) << 3)
				off := adrBuild(other, immhi)
				binary.LittleEndian.PutUint32(self.c[prog.Pc:], uint32(off))
			}
		}
	}
}

func (self *BaseAssembler) validate() {
	for key := range self.pendings {
		panic("links are not fully resolved: " + key)
	}
}

func (self *BaseAssembler) assemble() {
	self.c, self.Pcdata = self.pb.Assemble()
}

func (self *BaseAssembler) debugAsmOutput() {
	if !_DBG_ASM_OUTPUT {
		return
	}
	fmt.Print("========Asm Output BEGIN========\n")
	for i, b := range self.c {
		fmt.Printf("%02x", b)
		if i > 0 {
			// fmt.Print(" ")
		}
	}
	fmt.Print("\n")
	fmt.Print("========Asm Output END========\n")
}
