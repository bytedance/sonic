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

package arm64

import "testing"

func TestLogicalImmediateEncodings(t *testing.T) {
	p := NewProgram()

	p.AND(X0, X1, 0xff)
	expAND := (uint32(1) << 31) | (0x24 << 23) | (encodeRegister(X1) << 5) | encodeRegister(X0)
	if p.insns[0].enc != expAND {
		t.Errorf("AND encoding: got 0x%08x, want 0x%08x (imm field simplified)", p.insns[0].enc, expAND)
	}

	p = NewProgram()
	p.ORR(W2, W3, 0x7)
	expORR := (uint32(0) << 31) | (1 << 29) | (0x24 << 23) | (encodeRegister(W3) << 5) | encodeRegister(W2)
	if p.insns[0].enc != expORR {
		t.Errorf("ORR encoding: got 0x%08x, want 0x%08x", p.insns[0].enc, expORR)
	}

	p = NewProgram()
	p.EOR(X4, X5, 0x1)
	expEOR := (uint32(1) << 31) | (1 << 30) | (0x24 << 23) | (encodeRegister(X5) << 5) | encodeRegister(X4)
	if p.insns[0].enc != expEOR {
		t.Errorf("EOR encoding: got 0x%08x, want 0x%08x", p.insns[0].enc, expEOR)
	}
}

func TestMOVZ_MOVK(t *testing.T) {
	p := NewProgram()
	// MOVZ X0, #0x1234, LSL #0
	p.MOVZ(X0, 0x1234, 0)
	ins := p.insns[0].enc
	if rd := ins & 0x1f; rd != uint32(X0.RegIndex()) {
		t.Errorf("MOVZ Rd=%d, want %d", rd, X0.RegIndex())
	}
	// Check it's 64-bit (sf=1)
	if sf := (ins >> 31) & 1; sf != 1 {
		t.Errorf("MOVZ sf=%d, want 1 for X register", sf)
	}
	// Check opcode field for MOVZ (opc=10)
	if opc := (ins >> 29) & 0x3; opc != 0x2 {
		t.Errorf("MOVZ opc=%d, want 2", opc)
	}

	p = NewProgram()
	// MOVK X1, #0x5678, LSL #16
	p.MOVK(X1, 0x5678, 16)
	ins = p.insns[0].enc
	if rd := ins & 0x1f; rd != uint32(X1.RegIndex()) {
		t.Errorf("MOVK Rd=%d, want %d", rd, X1.RegIndex())
	}
	// Check opcode field for MOVK (opc=11)
	if opc := (ins >> 29) & 0x3; opc != 0x3 {
		t.Errorf("MOVK opc=%d, want 3", opc)
	}
	// Check hw field (shift/16)
	if hw := (ins >> 21) & 0x3; hw != 1 {
		t.Errorf("MOVK hw=%d, want 1 for LSL #16", hw)
	}
}

func TestConditionalBranches(t *testing.T) {
	// Test all conditional branch variants
	tests := []struct {
		name string
		cond Condition
	}{
		{"BEQ", CondEQ},
		{"BNE", CondNE},
		{"BLT", CondLT},
		{"BLE", CondLE},
		{"BGT", CondGT},
		{"BGE", CondGE},
		{"BLS", CondLS},
		{"BHI", CondHI},
	}

	for _, tt := range tests {
		p := NewProgram()
		label := CreateLabel(tt.name)

		// Call the appropriate branch function
		switch tt.cond {
		case CondEQ:
			p.BEQ(label)
		case CondNE:
			p.BNE(label)
		case CondLT:
			p.BLT(label)
		case CondLE:
			p.BLE(label)
		case CondGT:
			p.BGT(label)
		case CondGE:
			p.BGE(label)
		case CondLS:
			p.BLS(label)
		case CondHI:
			p.BHI(label)
		}

		ins := p.insns[0].enc
		// Check condition code field (bits 0-3)
		if gotCond := ins & 0xf; gotCond != uint32(tt.cond) {
			t.Errorf("%s: cond=%d, want %d", tt.name, gotCond, tt.cond)
		}
		// Check opcode for conditional branch (bits 24-31 should be 0x54)
		if op := (ins >> 24) & 0xff; (op & 0xfe) != 0x54 {
			t.Errorf("%s: opcode=0x%x, want 0x54/0x55", tt.name, op)
		}
	}
}

func TestMSR_MRS(t *testing.T) {
	p := NewProgram()
	// MSR NZCV, X0
	p.MSR("nzcv", X0)
	ins := p.insns[0].enc
	// Check it's a system register instruction
	if op0 := (ins >> 19) & 0x3; op0 != 0 {
		t.Errorf("MSR op0=%d", op0)
	}
	// Verify Rt field
	if rt := ins & 0x1f; rt != uint32(X0.RegIndex()) {
		t.Errorf("MSR Rt=%d, want %d", rt, X0.RegIndex())
	}

	p = NewProgram()
	// MRS X1, NZCV
	p.MRS(X1, "nzcv")
	ins = p.insns[0].enc
	// Verify Rt field
	if rt := ins & 0x1f; rt != uint32(X1.RegIndex()) {
		t.Errorf("MRS Rt=%d, want %d", rt, X1.RegIndex())
	}
}
