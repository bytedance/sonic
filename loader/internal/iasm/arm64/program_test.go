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
	"encoding/hex"
	"testing"
)

func TestRegisterString(t *testing.T) {
	tests := []struct {
		reg  Register
		want string
	}{
		{X0, "x0"},
		{X15, "x15"},
		{X30, "x30"},
		{SP, "sp"},
		{XZR, "xzr"},
		{W0, "w0"},
		{W15, "w15"},
		{WZR, "wzr"},
		{V0, "v0"},
		{V31, "v31"},
		{D0, "d0"},
		{D31, "d31"},
		{S0, "s0"},
		{S31, "s31"},
	}

	for _, tt := range tests {
		if got := tt.reg.String(); got != tt.want {
			t.Errorf("%T.String() = %v, want %v", tt.reg, got, tt.want)
		}
	}
}

func TestMemoryOperand(t *testing.T) {
	tests := []struct {
		mem  *MemoryOperand
		want string
	}{
		{Ptr(SP, 0), "[sp]"},
		{Ptr(SP, 16), "[sp, #16]"},
		{Ptr(X0, -8), "[x0, #-8]"},
		{PtrIndex(SP, X1), "[sp, x1]"},
		{PtrIndexScale(X0, X1, 3), "[x0, x1, lsl #3]"},
	}

	for _, tt := range tests {
		if got := tt.mem.String(); got != tt.want {
			t.Errorf("MemoryOperand.String() = %v, want %v", got, tt.want)
		}
	}
}

func TestCondition(t *testing.T) {
	tests := []struct {
		cond Condition
		want string
	}{
		{CondEQ, "eq"},
		{CondNE, "ne"},
		{CondLT, "lt"},
		{CondGE, "ge"},
		{CondAL, "al"},
	}

	for _, tt := range tests {
		if got := tt.cond.String(); got != tt.want {
			t.Errorf("Condition.String() = %v, want %v", got, tt.want)
		}
	}
}

func TestSimpleInstructions(t *testing.T) {
	p := NewProgram()

	// Test NOP
	p.NOP()
	if len(p.insns) != 1 || p.insns[0].enc != 0xd503201f {
		t.Errorf("NOP encoding failed: got 0x%08x", p.insns[0].enc)
	}

	// Test RET
	p = NewProgram()
	p.RET()
	if len(p.insns) != 1 || p.insns[0].enc != 0xd65f03c0 {
		t.Errorf("RET encoding failed: got 0x%08x", p.insns[0].enc)
	}
}

func TestADD(t *testing.T) {
	p := NewProgram()
	p.ADD(X0, X1, 42)

	// ADD X0, X1, #42 should encode to: 0x91 0x0a 0x80 0x20
	// Format: sf=1, op=00, S=0, shift=00, imm12=0x02a (42), Rn=1, Rd=0
	expected := uint32(0x9100a820)
	if p.insns[0].enc != expected {
		t.Errorf("ADD encoding failed: got 0x%08x, want 0x%08x", p.insns[0].enc, expected)
	}
}

func TestSUB(t *testing.T) {
	p := NewProgram()
	p.SUB(X0, X1, 16)

	// SUB X0, X1, #16
	// Format: sf=1, op=10, S=0, shift=00, imm12=0x010 (16), Rn=1, Rd=0
	expected := uint32(0xd1004020)
	if p.insns[0].enc != expected {
		t.Errorf("SUB encoding failed: got 0x%08x, want 0x%08x", p.insns[0].enc, expected)
	}
}

func TestMOV(t *testing.T) {
	p := NewProgram()
	p.MOV(X0, X1)

	// MOV X0, X1 (encoded as ORR X0, XZR, X1)
	// Should have X1 in Rm field and XZR in Rn field
	if p.insns[0].enc&0x1f != 0 { // Check Rd is X0
		t.Errorf("MOV encoding failed: Rd field incorrect")
	}
}

func TestLDR_STR(t *testing.T) {
	p := NewProgram()

	// LDR X0, [SP, #16]
	p.LDR(X0, Ptr(SP, 16))

	// STR X1, [SP, #8]
	p.STR(X1, Ptr(SP, 8))

	if len(p.insns) != 2 {
		t.Errorf("Expected 2 instructions, got %d", len(p.insns))
	}
}

func TestSTRPreIndex(t *testing.T) {
	p := NewProgram()

	// STR X30, [SP, #-0x70]!
	// Expected machine code (little-endian bytes): FE 0F 19 F8
	// As uint32 little-endian value: 0xF8190FFE
	mem := Ptr(SP, -0x70)
	mem.Mode = AddrModePreIndex
	p.STR(X30, mem)

	if len(p.insns) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(p.insns))
	}

	ins := p.insns[0].enc

	// idx bits (bits 11-10) should be 0x3 for pre-index
	if ((ins >> 10) & 0x3) != 0x3 {
		t.Errorf("Expected pre-index idx=3, got %d", (ins>>10)&0x3)
	}

	// Rn should be SP (31)
	if ((ins >> 5) & 0x1f) != uint32(SP.RegIndex()) {
		t.Errorf("Expected Rn=SP(31), got %d", (ins>>5)&0x1f)
	}

	// Rt should be X30 (30)
	if (ins & 0x1f) != uint32(X30.RegIndex()) {
		t.Errorf("Expected Rt=X30(30), got %d", ins&0x1f)
	}

	// Check full encoding equals expected machine code
	expected := uint32(0xF8190FFE)
	if ins != expected {
		t.Errorf("Expected encoding 0x%08x (FE 0F 19 F8), got 0x%08x", expected, ins)
	}
}

func TestSTP_LDP(t *testing.T) {
	p := NewProgram()

	// STP X29, X30, [SP, #-16]!
	mem := Ptr(SP, -16)
	mem.Mode = AddrModePreIndex
	p.STP(X29, X30, mem)

	// LDP X29, X30, [SP], #16
	mem2 := Ptr(SP, 16)
	mem2.Mode = AddrModePostIndex
	p.LDP(X29, X30, mem2)

	if len(p.insns) != 2 {
		t.Errorf("Expected 2 instructions, got %d", len(p.insns))
	}
}

func TestBranches(t *testing.T) {
	p := NewProgram()

	label := CreateLabel("test")

	// B label
	p.B(label)

	// BL label
	p.BL(label)

	// B.EQ label
	p.BEQ(label)

	// Link the label
	p.Link(label)
	p.NOP()

	if len(p.insns) != 4 {
		t.Errorf("Expected 4 instructions, got %d", len(p.insns))
	}
}

func TestBR_BLR(t *testing.T) {
	p := NewProgram()

	// BR X16
	p.BR(X16)
	expected := uint32(0xd61f0200) // BR X16
	if p.insns[0].enc != expected {
		t.Errorf("BR encoding failed: got 0x%08x, want 0x%08x", p.insns[0].enc, expected)
	}

	// BLR X16
	p.BLR(X16)
	expected = uint32(0xd63f0200) // BLR X16
	if p.insns[1].enc != expected {
		t.Errorf("BLR encoding failed: got 0x%08x, want 0x%08x", p.insns[1].enc, expected)
	}
}

func TestCMP(t *testing.T) {
	p := NewProgram()

	// CMP X0, #10
	p.CMP(X0, 10)

	// CMP is encoded as SUBS XZR, X0, #10
	// Check that Rd field is XZR (31)
	if p.insns[0].enc&0x1f != 31 {
		t.Errorf("CMP encoding failed: Rd should be XZR")
	}
}

func TestAssemble(t *testing.T) {
	p := NewProgram()

	// Simple program: add, sub, ret
	p.ADD(X0, X0, 1)
	p.SUB(X0, X0, 1)
	p.RET()

	// Assemble
	code := p.Assemble(0)

	// Should be 3 instructions × 4 bytes = 12 bytes
	if len(code) != 12 {
		t.Errorf("Expected 12 bytes, got %d", len(code))
	}

	// Print hex for debugging
	t.Logf("Assembled code: %s", hex.EncodeToString(code))
}

// Test LDR/STR for floating point registers (D/S) to ensure V bit and size are encoded
func TestLDR_STR_Float(t *testing.T) {
	p := NewProgram()

	// STR D0, [SP, #16]
	p.STR(D0, Ptr(SP, 16))
	// LDR D1, [SP, #24]
	p.LDR(D1, Ptr(SP, 24))

	// STR S0, [SP, #8]
	p.STR(S0, Ptr(SP, 8))
	// LDR S1, [SP, #12]
	p.LDR(S1, Ptr(SP, 12))

	if len(p.insns) != 4 {
		t.Fatalf("Expected 4 instructions, got %d", len(p.insns))
	}

	// Compare exact expected machine code (bytes given as little-endian uint32)
	// Expected bytes (little-endian):
	// STR D0, [SP, #16] -> E0 0B 00 FD  => uint32 0xFD000BE0
	// LDR D1, [SP, #24] -> E1 0F 40 FD  => uint32 0xFD400FE1
	// STR S0, [SP, #8]  -> E0 0B 00 BD  => uint32 0xBD000BE0
	// LDR S1, [SP, #12] -> E1 0F 40 BD  => uint32 0xBD400FE1
	expected := []uint32{
		0xFD000BE0,
		0xFD400FE1,
		0xBD000BE0,
		0xBD400FE1,
	}

	for i, ins := range p.insns {
		if ins.enc != expected[i] {
			t.Errorf("Instruction %d: encoding mismatch, got 0x%08x, want 0x%08x", i, ins.enc, expected[i])
		}
	}
}

func TestProgramWithLabels(t *testing.T) {
	p := NewProgram()

	entry := CreateLabel("entry")
	loop := CreateLabel("loop")

	p.Link(entry)
	p.MOV(X0, XZR)

	p.Link(loop)
	p.ADD(X0, X0, 1)
	p.CMP(X0, 10)
	p.BNE(loop)
	p.RET()

	// Should have 5 instructions
	if len(p.insns) != 5 {
		t.Errorf("Expected 5 instructions, got %d", len(p.insns))
	}

	// Check labels are linked
	if entry.Pos != 0 {
		t.Errorf("Entry label pos = %d, want 0", entry.Pos)
	}
	if loop.Pos != 1 {
		t.Errorf("Loop label pos = %d, want 1", loop.Pos)
	}
}
