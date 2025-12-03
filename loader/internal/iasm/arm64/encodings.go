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
	"encoding/binary"
	"fmt"
)

// encodeRegister encodes a register into its 5-bit field
func encodeRegister(r Register) uint32 {
	switch reg := r.(type) {
	case Register64:
		return uint32(reg.RegIndex())
	case Register32:
		return uint32(reg.RegIndex())
	case VRegister:
		return uint32(reg.RegIndex())
	case DRegister:
		return uint32(reg.RegIndex())
	case SRegister:
		return uint32(reg.RegIndex())
	default:
		panic(fmt.Sprintf("unsupported register type: %T", r))
	}
}

// encodeImm12 encodes a 12-bit unsigned immediate
func encodeImm12(imm int64) uint32 {
	if imm < 0 || imm > 0xfff {
		panic(fmt.Sprintf("immediate value out of range for 12-bit encoding: %d", imm))
	}
	return uint32(imm)
}

// encodeImm9 encodes a 9-bit signed immediate
func encodeImm9(imm int32) uint32 {
	if imm < -256 || imm > 255 {
		panic(fmt.Sprintf("immediate value out of range for 9-bit encoding: %d", imm))
	}
	return uint32(imm) & 0x1ff
}

// encodeImm7 encodes a 7-bit signed immediate (for load/store pair)
func encodeImm7(imm int32, scale uint8) uint32 {
	if imm%(1<<scale) != 0 {
		panic(fmt.Sprintf("immediate value must be aligned to %d bytes: %d", 1<<scale, imm))
	}
	scaled := imm >> scale
	if scaled < -64 || scaled > 63 {
		panic(fmt.Sprintf("scaled immediate out of range for 7-bit encoding: %d (scaled: %d)", imm, scaled))
	}
	return uint32(scaled) & 0x7f
}

// encodeImm19 encodes a 19-bit signed immediate for conditional branches (word offset)
func encodeImm19(offset int32) uint32 {
	if offset&3 != 0 {
		panic(fmt.Sprintf("branch offset must be 4-byte aligned: %d", offset))
	}
	wordOffset := offset >> 2
	if wordOffset < -(1<<18) || wordOffset >= (1<<18) {
		panic(fmt.Sprintf("branch offset out of range: %d", offset))
	}
	return uint32(wordOffset) & 0x7ffff
}

// encodeImm26 encodes a 26-bit signed immediate for unconditional branches (word offset)
func encodeImm26(offset int32) uint32 {
	if offset&3 != 0 {
		panic(fmt.Sprintf("branch offset must be 4-byte aligned: %d", offset))
	}
	wordOffset := offset >> 2
	if wordOffset < -(1<<25) || wordOffset >= (1<<25) {
		panic(fmt.Sprintf("branch offset out of range: %d", offset))
	}
	return uint32(wordOffset) & 0x3ffffff
}

// encodeLogicalImmediate encodes a bitmask immediate for logical instructions
// This is a simplified version - full implementation requires complex algorithm
func encodeLogicalImmediate(imm uint64, is64bit bool) uint32 {
	// TODO: Implement full bitmask immediate encoding
	// For now, just return 0 as placeholder
	// See ARM Architecture Reference Manual for the full algorithm
	return 0
}

// encodeMovImmediate encodes a 16-bit immediate for MOV instructions
func encodeMovImmediate(imm uint16, shift uint8) uint32 {
	if shift > 3 {
		panic(fmt.Sprintf("MOV immediate shift must be 0-3: %d", shift))
	}
	return (uint32(shift) << 21) | (uint32(imm) << 5)
}

// encodeShift encodes shift type and amount
func encodeShift(shiftType ShiftType, amount uint8) uint32 {
	if amount > 63 {
		panic(fmt.Sprintf("shift amount out of range: %d", amount))
	}
	return (uint32(shiftType) << 22) | (uint32(amount) << 10)
}

// encodeExtend encodes extend type and amount
func encodeExtend(extendType ExtendType, amount uint8) uint32 {
	if amount > 4 {
		panic(fmt.Sprintf("extend amount out of range: %d", amount))
	}
	return (uint32(extendType) << 13) | (uint32(amount) << 10)
}

// encodeCondition encodes condition code
func encodeCondition(cond Condition) uint32 {
	return uint32(cond)
}

// Instruction encoding helper functions

// encodeDataProcessingImmediate encodes data processing (immediate) instructions
func encodeDataProcessingImmediate(op uint32, sf uint32, rd, rn Register, imm12 uint32) uint32 {
	return (sf << 31) | (op << 23) | (imm12 << 10) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// encodeDataProcessingRegister encodes data processing (register) instructions
func encodeDataProcessingRegister(op uint32, sf uint32, rd, rn, rm Register) uint32 {
	return (sf << 31) | (op << 21) | (encodeRegister(rm) << 16) | (encodeRegister(rn) << 5) | encodeRegister(rd)
}

// encodeLoadStore encodes load/store instructions
func encodeLoadStore(op uint32, size uint32, rt Register, mem *MemoryOperand) uint32 {
	rn := encodeRegister(mem.Base)
	rtEnc := encodeRegister(rt)

	// Determine V (vector) bit: 1 for FP/SIMD registers (D/S/V), 0 for GPR
	var vBit uint32 = 0
	switch rt.(type) {
	case DRegister, SRegister, VRegister:
		vBit = 1
	default:
		vBit = 0
	}

	switch mem.Mode {
	case AddrModeOffset:
		// Support both positive and negative offsets
		// Positive: use unsigned offset encoding (imm12)
		// Negative: use STUR/LDUR encoding (imm9)
		if mem.Offset < 0 {
			// Use unscaled immediate offset (STUR/LDUR): -256 to +255
			imm9 := encodeImm9(mem.Offset)
			// Format: size 111 V 00 opc 0 imm9 00 rn rt
			fixedBits := uint32(0x38<<24) | (op << 22)
			return (size << 30) | fixedBits | (imm9 << 12) | (rn << 5) | rtEnc
		}
		// Unsigned offset: [base, #offset]
		// Format: size 111 V 00 opc imm12 rn rt
		// V=0 for GPR, V=1 for FP/SIMD
		// opc: 00=STR, 01=LDR
		offset := uint32(mem.Offset)
		scale := size
		if offset%(1<<scale) != 0 {
			panic(fmt.Sprintf("offset must be aligned to %d bytes: %d", 1<<scale, offset))
		}
		scaledOffset := offset >> scale
		if scaledOffset > 0xfff {
			panic(fmt.Sprintf("offset out of range: %d", offset))
		}
		// Bits 29-22: 111 V 00 opc
		// Use base 0x39 for unsigned offset encoding, then set V bit and opc
		fixedBits := (0x39 << 24) | (vBit << 26) | (op << 22)
		return (size << 30) | fixedBits | (scaledOffset << 10) | (rn << 5) | rtEnc

	case AddrModePreIndex:
		// Pre-indexed: [base, #offset]!
		imm9 := encodeImm9(mem.Offset)
		// Use unscaled immediate encoding base bits (same as STUR/LDUR format)
		fixedBits := (0x38 << 24) | (vBit << 26) | (op << 22)
		return (size << 30) | fixedBits | (imm9 << 12) | (0x3 << 10) | (rn << 5) | rtEnc

	case AddrModePostIndex:
		// Post-indexed: [base], #offset
		imm9 := encodeImm9(mem.Offset)
		fixedBits := (0x38 << 24) | (vBit << 26) | (op << 22)
		return (size << 30) | fixedBits | (imm9 << 12) | (0x1 << 10) | (rn << 5) | rtEnc

	case AddrModeRegister, AddrModeExtended:
		// Register offset: [base, index]
		rm := encodeRegister(mem.Index)
		option := uint32(0x6) // LSL
		s := uint32(0)
		if mem.Scale > 0 {
			s = 1
		}
		// Include V bit for FP/SIMD
		return (size << 30) | (0x1c << 24) | (vBit << 26) | (op << 22) | (1 << 21) | (rm << 16) | (option << 13) | (s << 12) | (rn << 5) | rtEnc

	default:
		panic(fmt.Sprintf("unsupported addressing mode: %v", mem.Mode))
	}
}

// encodeBranch encodes branch instructions
func encodeBranch(op uint32, offset int32) uint32 {
	imm26 := encodeImm26(offset)
	return (op << 26) | imm26
}

// encodeConditionalBranch encodes conditional branch instructions
func encodeConditionalBranch(cond Condition, offset int32) uint32 {
	imm19 := encodeImm19(offset)
	return (0x54 << 24) | (imm19 << 5) | encodeCondition(cond)
}

// append32 appends a 32-bit instruction to the buffer
func append32(buf *[]byte, instr uint32) {
	var tmp [4]byte
	binary.LittleEndian.PutUint32(tmp[:], instr)
	*buf = append(*buf, tmp[:]...)
}

// Helper functions for common instruction patterns

// isValidImm12 checks if a value can be encoded as 12-bit immediate
func isValidImm12(imm int64) bool {
	return imm >= 0 && imm <= 0xfff
}

// isValidImm9 checks if a value can be encoded as 9-bit signed immediate
func isValidImm9(imm int32) bool {
	return imm >= -256 && imm <= 255
}

// alignmentScale returns the log2 of alignment required for the size
func alignmentScale(size uint32) uint8 {
	switch size {
	case 0:
		return 0 // byte
	case 1:
		return 1 // halfword
	case 2:
		return 2 // word
	case 3:
		return 3 // doubleword
	default:
		panic(fmt.Sprintf("invalid size: %d", size))
	}
}
