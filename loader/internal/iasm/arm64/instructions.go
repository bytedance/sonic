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

import "fmt"

// Data Processing Instructions

// ADD - Add (immediate or register)
func (p *Program) ADD(rd, rn Register, op interface{}) {
	insn := p.add("add", rd, rn, op)

	var sf uint32 = 1 // 64-bit by default
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// Check if op is a register or immediate
	switch v := op.(type) {
	case Register64, Register32:
		// ADD (shifted register): sf 0 01011 shift 0 Rm imm6 Rn Rd
		// For simple ADD without shift: shift=00, imm6=000000
		rm := op.(Register)
		insn.enc = (sf << 31) | (0x0b << 24) | (encodeRegister(rm) << 16) | (encodeRegister(rn) << 5) | encodeRegister(rd)
	case int, int32, int64, uint32:
		// ADD (immediate): sf 0 010001 sh imm12 Rn Rd
		var immVal uint32
		switch iv := v.(type) {
		case int:
			immVal = encodeImm12(int64(iv))
		case int32:
			immVal = encodeImm12(int64(iv))
		case int64:
			immVal = encodeImm12(iv)
		case uint32:
			immVal = encodeImm12(int64(iv))
		}
		insn.enc = (sf << 31) | (0x11 << 24) | (immVal << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
	default:
		panic(fmt.Sprintf("unsupported operand type for ADD: %T", op))
	}
}

// SUB - Subtract (immediate)
func (p *Program) SUB(rd, rn Register, imm interface{}) {
	insn := p.add("sub", rd, rn, imm)

	var immVal uint32
	switch v := imm.(type) {
	case int:
		immVal = encodeImm12(int64(v))
	case int32:
		immVal = encodeImm12(int64(v))
	case int64:
		immVal = encodeImm12(v)
	case uint32:
		immVal = encodeImm12(int64(v))
	default:
		panic(fmt.Sprintf("unsupported immediate type: %T", imm))
	}

	var sf uint32 = 1
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// SUB (immediate): sf 1 010001 sh imm12 Rn Rd
	insn.enc = (sf << 31) | (1 << 30) | (0x11 << 24) | (immVal << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// MOV - Move (register)
func (p *Program) MOV(rd, rn Register) {
	// MOV is an alias for ORR Rd, XZR, Rn
	insn := p.add("mov", rd, rn)

	var sf uint32 = 1
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// ORR (shifted register): sf 01 01010 shift 0 Rm imm6 Rn Rd
	// With Rn = XZR and shift = 0, imm6 = 0, this becomes MOV
	zr := XZR
	if sf == 0 {
		zr = Register64(WZR)
	}
	insn.enc = (sf << 31) | (0x2a << 24) | (encodeRegister(rn) << 16) | (encodeRegister(zr) << 5) | encodeRegister(rd)
}

// MOVZ - Move wide with zero
func (p *Program) MOVZ(rd Register64, imm uint16, shift uint8) {
	insn := p.add("movz", rd, imm, shift)
	// MOVZ: sf 10 100101 hw imm16 Rd
	// For 64-bit: bits 31-23 = 110100101 = 0x1A5 << 1 | 1 << 8
	// Simpler: 0xD2800000 base + adjustments
	hw := uint32(shift / 16)
	if hw > 3 {
		panic(fmt.Sprintf("invalid shift for MOVZ: %d", shift))
	}
	// Base encoding for MOVZ X-reg: 0xD2800000
	// Add: hw field (bits 22-21) + imm16 (bits 20-5) + Rd (bits 4-0)
	insn.enc = 0xD2800000 | (hw << 21) | (uint32(imm) << 5) | encodeRegister(rd)
}

// MOVK - Move wide with keep
func (p *Program) MOVK(rd Register64, imm uint16, shift uint8) {
	insn := p.add("movk", rd, imm, shift)
	// MOVK: sf 11 100101 hw imm16 Rd
	// For 64-bit: 0xF2800000 base
	hw := uint32(shift / 16)
	if hw > 3 {
		panic(fmt.Sprintf("invalid shift for MOVK: %d", shift))
	}
	insn.enc = 0xF2800000 | (hw << 21) | (uint32(imm) << 5) | encodeRegister(rd)
}

// CMP - Compare (immediate or register)
func (p *Program) CMP(rn Register, operand interface{}) {
	// Check if operand is a register or immediate
	if rm, ok := operand.(Register); ok {
		// CMP with register is an alias for SUBS XZR, Rn, Rm
		insn := p.add("cmp", rn, rm)

		var sf uint32 = 1
		if _, ok := rn.(Register32); ok {
			sf = 0
		}

		zr := XZR
		if sf == 0 {
			zr = Register64(WZR)
		}

		// SUBS (shifted register): sf 11 01011 shift(00) 0 Rm imm6(000000) Rn Rd
		insn.enc = (sf << 31) | (0x6B << 24) | (encodeRegister(rm) << 16) | (encodeRegister(rn) << 5) | encodeRegister(zr)
		return
	}

	// CMP with immediate is an alias for SUBS XZR, Rn, #imm
	insn := p.add("cmp", rn, operand)

	var immVal uint32
	switch v := operand.(type) {
	case int:
		immVal = encodeImm12(int64(v))
	case int32:
		immVal = encodeImm12(int64(v))
	case int64:
		immVal = encodeImm12(v)
	case uint32:
		immVal = encodeImm12(int64(v))
	default:
		panic(fmt.Sprintf("unsupported immediate type: %T", operand))
	}

	var sf uint32 = 1
	if _, ok := rn.(Register32); ok {
		sf = 0
	}

	zr := XZR
	if sf == 0 {
		zr = Register64(WZR)
	}

	// SUBS (immediate): sf 1 110001 sh imm12 Rn Rd
	insn.enc = (sf << 31) | (0x71 << 24) | (immVal << 10) | (encodeRegister(rn) << 5) | encodeRegister(zr)
}

// Load/Store Instructions

// LDR - Load register
func (p *Program) LDR(rt Register, mem *MemoryOperand) {
	insn := p.add("ldr", rt, mem)

	var size uint32
	switch rt.(type) {
	case Register64:
		size = 3 // 64-bit
	case Register32:
		size = 2 // 32-bit
	case DRegister:
		size = 3 // 64-bit FP
	case SRegister:
		size = 2 // 32-bit FP
	default:
		panic(fmt.Sprintf("unsupported register type for LDR: %T", rt))
	}

	insn.enc = encodeLoadStore(0x1, size, rt, mem)
}

// STR - Store register
func (p *Program) STR(rt Register, mem *MemoryOperand) {
	insn := p.add("str", rt, mem)

	var size uint32
	switch rt.(type) {
	case Register64:
		size = 3 // 64-bit
	case Register32:
		size = 2 // 32-bit
	case DRegister:
		size = 3 // 64-bit FP
	case SRegister:
		size = 2 // 32-bit FP
	default:
		panic(fmt.Sprintf("unsupported register type for STR: %T", rt))
	}

	insn.enc = encodeLoadStore(0x0, size, rt, mem)
}

// STP - Store pair
func (p *Program) STP(rt1, rt2 Register, mem *MemoryOperand) {
	insn := p.add("stp", rt1, rt2, mem)

	var opc uint32 // size field in ARM docs
	switch rt1.(type) {
	case Register64:
		opc = 2 // 10 = 64-bit pair
	case Register32:
		opc = 0 // 00 = 32-bit pair
	default:
		panic(fmt.Sprintf("unsupported register type for STP: %T", rt1))
	}

	// STP format: opc(31-30) 101(29-27) V(26) idx(24-23) L(22) imm7(21-15) Rt2(14-10) Rn(9-5) Rt1(4-0)
	// V=0 for GPR, V=1 for FP/SIMD
	// L=0 for store (STP), L=1 for load (LDP)
	// idx: 10=signed offset, 11=pre-index, 01=post-index
	var idx uint32
	if mem.Mode == AddrModePreIndex {
		idx = 0x3 // Pre-indexed: 11
	} else if mem.Mode == AddrModePostIndex {
		idx = 0x1 // Post-indexed: 01
	} else {
		idx = 0x2 // Signed offset: 10
	}

	// Scale: 32-bit pairs aligned to 4, 64-bit pairs aligned to 8
	scale := uint32(2)
	if opc == 2 {
		scale = 3 // 64-bit
	}

	offset := mem.Offset
	if offset%(1<<scale) != 0 {
		panic(fmt.Sprintf("STP offset must be %d-byte aligned", 1<<scale))
	}
	imm7 := uint32((offset >> scale) & 0x7F)

	// Encoding: opc(31-30) 101(29-27) V(26) idx(24-23) L(22) imm7(21-15) Rt2(14-10) Rn(9-5) Rt1(4-0)
	vBit := uint32(0) // GPR
	lBit := uint32(0) // Store

	rt1Enc := encodeRegister(rt1)
	rt2Enc := encodeRegister(rt2)
	rnEnc := encodeRegister(mem.Base)

	insn.enc = (opc << 30) | (0x5 << 27) | (vBit << 26) | (idx << 23) | (lBit << 22) |
		(imm7 << 15) | (rt2Enc << 10) |
		(rnEnc << 5) | rt1Enc
}

// LDP - Load pair
func (p *Program) LDP(rt1, rt2 Register, mem *MemoryOperand) {
	insn := p.add("ldp", rt1, rt2, mem)

	var opc uint32 // size field
	switch rt1.(type) {
	case Register64:
		opc = 2 // 10 = 64-bit pair
	case Register32:
		opc = 0 // 00 = 32-bit pair
	default:
		panic(fmt.Sprintf("unsupported register type for LDP: %T", rt1))
	}

	// LDP format: opc(31-30) 101(29-27) V(26) idx(24-23) L(22) imm7(21-15) Rt2(14-10) Rn(9-5) Rt1(4-0)
	// V=0 for GPR, V=1 for FP/SIMD
	// L=0 for store (STP), L=1 for load (LDP)
	// idx: 10=signed offset, 11=pre-index, 01=post-index
	var idx uint32
	if mem.Mode == AddrModePreIndex {
		idx = 0x3 // Pre-indexed: 11
	} else if mem.Mode == AddrModePostIndex {
		idx = 0x1 // Post-indexed: 01
	} else {
		idx = 0x2 // Signed offset: 10
	}

	scale := uint32(2)
	if opc == 2 {
		scale = 3 // 64-bit
	}

	offset := mem.Offset
	if offset%(1<<scale) != 0 {
		panic(fmt.Sprintf("LDP offset must be %d-byte aligned", 1<<scale))
	}
	imm7 := uint32((offset >> scale) & 0x7F)

	// Encoding: opc(31-30) 101(29-27) V(26) idx(24-23) L(22) imm7(21-15) Rt2(14-10) Rn(9-5) Rt1(4-0)
	vBit := uint32(0) // GPR
	lBit := uint32(1) // Load
	insn.enc = (opc << 30) | (0x5 << 27) | (vBit << 26) | (idx << 23) | (lBit << 22) |
		(imm7 << 15) | (encodeRegister(rt2) << 10) |
		(encodeRegister(mem.Base) << 5) | encodeRegister(rt1)
}

// Branch Instructions

// B - Unconditional branch
func (p *Program) B(label *Label) {
	insn := p.add("b", label)
	insn.label = label
	// Will be encoded during assembly when label position is known
	// B: 0 00101 imm26
	// For now, use a placeholder that will be fixed up
	insn.enc = 0x14000000 // B with offset 0
}

// BL - Branch with link
func (p *Program) BL(label *Label) {
	insn := p.add("bl", label)
	insn.label = label
	// BL: 1 00101 imm26
	insn.enc = 0x94000000 // BL with offset 0
}

// BR - Branch to register
func (p *Program) BR(rn Register64) {
	insn := p.add("br", rn)
	// BR: 1101011 0000 11111 000000 Rn 00000
	insn.enc = (0xd61f << 16) | (encodeRegister(rn) << 5)
}

// BLR - Branch with link to register
func (p *Program) BLR(rn Register64) {
	insn := p.add("blr", rn)
	// BLR: 1101011 0001 11111 000000 Rn 00000
	insn.enc = (0xd63f << 16) | (encodeRegister(rn) << 5)
}

// RET - Return
func (p *Program) RET() {
	insn := p.add("ret")
	// RET: 1101011 0010 11111 000000 11110 00000 (RET X30)
	insn.enc = 0xd65f03c0 // arm64 RET encoding
}

// Conditional Branch

// Bcond - Conditional branch
func (p *Program) Bcond(cond Condition, label *Label) {
	insn := p.add(fmt.Sprintf("b.%s", cond), label)
	insn.label = label
	// B.cond: 01010100 imm19 0 cond
	insn.enc = 0x54000000 | encodeCondition(cond)
}

// Convenience methods for common conditions
func (p *Program) BEQ(label *Label) { p.Bcond(CondEQ, label) }
func (p *Program) BNE(label *Label) { p.Bcond(CondNE, label) }
func (p *Program) BLT(label *Label) { p.Bcond(CondLT, label) }
func (p *Program) BLE(label *Label) { p.Bcond(CondLE, label) }
func (p *Program) BGT(label *Label) { p.Bcond(CondGT, label) }
func (p *Program) BGE(label *Label) { p.Bcond(CondGE, label) }
func (p *Program) BLS(label *Label) { p.Bcond(CondLS, label) }
func (p *Program) BHI(label *Label) { p.Bcond(CondHI, label) }
func (p *Program) BLO(label *Label) { p.Bcond(CondLO, label) }

// Additional common instructions (merged from instructions_extra.go)

// AND - Bitwise AND (immediate)
func (p *Program) AND(rd, rn Register, imm uint64) {
	insn := p.add("and", rd, rn, imm)

	var sf uint32 = 1
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// AND (immediate): sf 00 100100 N immr imms Rn Rd
	// Using simplified encoding - full implementation requires bitmask immediate encoding
	immEnc := encodeLogicalImmediate(imm, sf == 1)
	insn.enc = (sf << 31) | (0x24 << 23) | (immEnc << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// ORR - Bitwise OR (immediate)
func (p *Program) ORR(rd, rn Register, imm uint64) {
	insn := p.add("orr", rd, rn, imm)

	var sf uint32 = 1
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// ORR (immediate): sf 01 100100 N immr imms Rn Rd
	immEnc := encodeLogicalImmediate(imm, sf == 1)
	insn.enc = (sf << 31) | (1 << 29) | (0x24 << 23) | (immEnc << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// EOR - Bitwise XOR (immediate)
func (p *Program) EOR(rd, rn Register, imm uint64) {
	insn := p.add("eor", rd, rn, imm)

	var sf uint32 = 1
	if _, ok := rd.(Register32); ok {
		sf = 0
	}

	// EOR (immediate): sf 10 100100 N immr imms Rn Rd
	immEnc := encodeLogicalImmediate(imm, sf == 1)
	insn.enc = (sf << 31) | (1 << 30) | (0x24 << 23) | (immEnc << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// System Instructions
// MSR - Move to system register (simplified)
func (p *Program) MSR(sysreg string, rt Register64) {
	insn := p.add("msr", sysreg, rt)
	// This is simplified - full implementation needs system register encoding
	insn.enc = 0xd5000000 | encodeRegister(rt)
}

// MRS - Move from system register (simplified)
func (p *Program) MRS(rt Register64, sysreg string) {
	insn := p.add("mrs", rt, sysreg)
	// This is simplified - full implementation needs system register encoding
	insn.enc = 0xd5200000 | encodeRegister(rt)
}
