//
// Copyright 2025 Huawei Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package arm64

import (
	"fmt"
)

// Label represents a code label
type Label struct {
	Name string
	Pos  int    // Position in instruction list (-1 if not yet linked)
	PC   uint32 // Actual PC when assembled
}

// CreateLabel creates a new label
func CreateLabel(name string) *Label {
	return &Label{
		Name: name,
		Pos:  -1,
	}
}

// Instruction represents a single ARM64 instruction
type Instruction struct {
	op     string        // Mnemonic
	enc    uint32        // Encoded instruction (0 if not yet encoded)
	args   []interface{} // Arguments
	label  *Label        // Associated label (if any)
	pc     uint32        // PC of this instruction
	size   int           // Size in bytes (usually 4)
	linked bool          // Whether label references are resolved
}

// Program represents an ARM64 assembly program
type Program struct {
	insns     []*Instruction
	labels    map[string]*Label
	pc        uint32
	assembled bool
}

// NewProgram creates a new ARM64 program
func NewProgram() *Program {
	return &Program{
		insns:  make([]*Instruction, 0, 64),
		labels: make(map[string]*Label),
	}
}

// Link associates a label with the current position
func (p *Program) Link(label *Label) {
	if label.Pos != -1 {
		panic(fmt.Sprintf("label %s already linked at position %d", label.Name, label.Pos))
	}
	label.Pos = len(p.insns)
	if label.Name != "" {
		if _, exists := p.labels[label.Name]; exists {
			panic(fmt.Sprintf("duplicate label: %s", label.Name))
		}
		p.labels[label.Name] = label
	}
}

// add adds an instruction to the program
func (p *Program) add(op string, args ...interface{}) *Instruction {
	insn := &Instruction{
		op:   op,
		args: args,
		size: 4, // ARM64 instructions are always 4 bytes
	}
	p.insns = append(p.insns, insn)
	return insn
}

// Assemble assembles the program into machine code
func (p *Program) Assemble(base uintptr) []byte {
	if p.assembled {
		panic("program already assembled")
	}

	// First pass: calculate positions and resolve label addresses
	pc := uint32(base)
	for i, insn := range p.insns {
		insn.pc = pc
		// Check if this position has a linked label
		for _, lbl := range p.labels {
			if lbl.Pos == i {
				lbl.PC = pc
			}
		}
		pc += uint32(insn.size)
	}

	// Second pass: encode instructions with label resolution
	buf := make([]byte, 0, len(p.insns)*4)
	for _, insn := range p.insns {
		if insn.enc == 0 {
			// Instruction not yet encoded - this shouldn't happen
			// if all emit functions properly set enc field
			panic(fmt.Sprintf("instruction %s not encoded", insn.op))
		}

		// Update instructions with label references
		finalEnc := insn.enc
		if insn.label != nil {
			offset := p.calculateOffset(insn.pc, insn.label)

			// Check instruction type and update encoding
			if (insn.enc & 0xff000000) == 0x54000000 { // B.cond
				// Conditional branch: 01010100 imm19 0 cond
				// Offset is in instructions (4-byte aligned), so divide by 4
				if offset%4 != 0 {
					panic(fmt.Sprintf("unaligned branch offset %d for %s", offset, insn.op))
				}
				immOffset := offset / 4
				if immOffset < -262144 || immOffset > 262143 { // 19-bit signed range
					panic(fmt.Sprintf("branch offset %d out of range for %s", immOffset, insn.op))
				}
				// Clear old imm19 field and set new one
				finalEnc = (insn.enc & 0xff00001f) | (uint32(immOffset&0x7ffff) << 5)

			} else if (insn.enc & 0xfc000000) == 0x14000000 { // B (unconditional)
				// Unconditional branch: 000101 imm26
				if offset%4 != 0 {
					panic(fmt.Sprintf("unaligned branch offset %d for %s", offset, insn.op))
				}
				immOffset := offset / 4
				if immOffset < -33554432 || immOffset > 33554431 { // 26-bit signed range
					panic(fmt.Sprintf("branch offset %d out of range for %s", immOffset, insn.op))
				}
				// Clear old imm26 field and set new one
				finalEnc = (insn.enc & 0xfc000000) | uint32(immOffset&0x3ffffff)
			}
		}

		append32(&buf, finalEnc)
	}

	p.assembled = true
	return buf
}

// calculateOffset calculates the offset from current PC to target label
func (p *Program) calculateOffset(fromPC uint32, to *Label) int32 {
	if to.Pos == -1 {
		panic(fmt.Sprintf("label %s not yet linked", to.Name))
	}
	return int32(to.PC) - int32(fromPC)
}

// currentPC returns the PC of the next instruction to be added
func (p *Program) currentPC() uint32 {
	pc := p.pc
	for _, insn := range p.insns {
		pc += uint32(insn.size)
	}
	return pc
}

// Byte emits a raw byte
func (p *Program) Byte(val byte) {
	insn := p.add(".byte", val)
	insn.size = 1
	insn.enc = uint32(val)
}

// Word emits a 16-bit word
func (p *Program) Word(val uint16) {
	insn := p.add(".word", val)
	insn.size = 2
	insn.enc = uint32(val)
}

// Long emits a 32-bit long
func (p *Program) Long(val uint32) {
	insn := p.add(".long", val)
	insn.size = 4
	insn.enc = val
}

// Quad emits a 64-bit quad
func (p *Program) Quad(val uint64) {
	insn := p.add(".quad", val)
	insn.size = 8
	// Store both halves
	insn.enc = uint32(val) // Lower 32 bits
	// We'll need special handling in Assemble for 64-bit values
}

// NOP emits a NOP instruction
func (p *Program) NOP() {
	insn := p.add("nop")
	insn.enc = 0xd503201f // NOP encoding
}

// String returns a string representation of the program
func (p *Program) String() string {
	s := "ARM64 Program:\n"
	for i, insn := range p.insns {
		s += fmt.Sprintf("  %04d: %s\n", i, insn.op)
	}
	return s
}
