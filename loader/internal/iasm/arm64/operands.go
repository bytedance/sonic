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

// MemoryOperand represents a memory addressing mode
type MemoryOperand struct {
	Base   Register64
	Offset int32
	Index  Register64 // For indexed addressing
	Scale  uint8      // For scaled indexed addressing (LSL amount)
	Mode   AddrMode   // Addressing mode
}

// AddrMode represents the addressing mode
type AddrMode uint8

const (
	AddrModeOffset    AddrMode = iota // [base, #offset]
	AddrModePreIndex                  // [base, #offset]! (pre-indexed)
	AddrModePostIndex                 // [base], #offset (post-indexed)
	AddrModeRegister                  // [base, index]
	AddrModeExtended                  // [base, index, extend]
)

// String implements fmt.Stringer
func (m *MemoryOperand) String() string {
	switch m.Mode {
	case AddrModeOffset:
		if m.Offset == 0 {
			return fmt.Sprintf("[%s]", m.Base)
		}
		return fmt.Sprintf("[%s, #%d]", m.Base, m.Offset)
	case AddrModePreIndex:
		return fmt.Sprintf("[%s, #%d]!", m.Base, m.Offset)
	case AddrModePostIndex:
		return fmt.Sprintf("[%s], #%d", m.Base, m.Offset)
	case AddrModeRegister:
		return fmt.Sprintf("[%s, %s]", m.Base, m.Index)
	case AddrModeExtended:
		if m.Scale > 0 {
			return fmt.Sprintf("[%s, %s, lsl #%d]", m.Base, m.Index, m.Scale)
		}
		return fmt.Sprintf("[%s, %s]", m.Base, m.Index)
	default:
		return fmt.Sprintf("[%s, ???]", m.Base)
	}
}

// Ptr creates a memory operand with base and offset
func Ptr(base Register64, offset int32) *MemoryOperand {
	return &MemoryOperand{
		Base:   base,
		Offset: offset,
		Mode:   AddrModeOffset,
	}
}

// PtrIndex creates a memory operand with base and index register
func PtrIndex(base Register64, index Register64) *MemoryOperand {
	return &MemoryOperand{
		Base:  base,
		Index: index,
		Mode:  AddrModeRegister,
	}
}

// PtrIndexScale creates a memory operand with base, index and scale
func PtrIndexScale(base Register64, index Register64, scale uint8) *MemoryOperand {
	return &MemoryOperand{
		Base:  base,
		Index: index,
		Scale: scale,
		Mode:  AddrModeExtended,
	}
}

// Condition represents ARM64 condition codes
type Condition uint8

const (
	CondEQ Condition = iota // Equal (Z == 1)
	CondNE                  // Not equal (Z == 0)
	CondCS                  // Carry set / unsigned higher or same (C == 1)
	CondCC                  // Carry clear / unsigned lower (C == 0)
	CondMI                  // Minus / negative (N == 1)
	CondPL                  // Plus / positive or zero (N == 0)
	CondVS                  // Overflow (V == 1)
	CondVC                  // No overflow (V == 0)
	CondHI                  // Unsigned higher (C == 1 && Z == 0)
	CondLS                  // Unsigned lower or same (C == 0 || Z == 1)
	CondGE                  // Signed greater than or equal (N == V)
	CondLT                  // Signed less than (N != V)
	CondGT                  // Signed greater than (Z == 0 && N == V)
	CondLE                  // Signed less than or equal (Z == 1 || N != V)
	CondAL                  // Always (unconditional)
	CondNV                  // Never (reserved)
)

var condNames = [16]string{
	"eq", "ne", "cs", "cc", "mi", "pl", "vs", "vc",
	"hi", "ls", "ge", "lt", "gt", "le", "al", "nv",
}

func (c Condition) String() string {
	if c <= CondNV {
		return condNames[c]
	}
	return fmt.Sprintf("cond?%d", c)
}

// Aliases for condition codes
const (
	CondHS = CondCS // Unsigned higher or same (same as CS)
	CondLO = CondCC // Unsigned lower (same as CC)
)

// Shift types for ARM64
type ShiftType uint8

const (
	ShiftLSL ShiftType = iota // Logical shift left
	ShiftLSR                  // Logical shift right
	ShiftASR                  // Arithmetic shift right
	ShiftROR                  // Rotate right
)

func (s ShiftType) String() string {
	switch s {
	case ShiftLSL:
		return "lsl"
	case ShiftLSR:
		return "lsr"
	case ShiftASR:
		return "asr"
	case ShiftROR:
		return "ror"
	default:
		return fmt.Sprintf("shift?%d", s)
	}
}

// Extend types for ARM64
type ExtendType uint8

const (
	ExtendUXTB ExtendType = iota // Unsigned extend byte
	ExtendUXTH                   // Unsigned extend halfword
	ExtendUXTW                   // Unsigned extend word
	ExtendUXTX                   // Unsigned extend doubleword (64-bit, essentially a copy)
	ExtendSXTB                   // Signed extend byte
	ExtendSXTH                   // Signed extend halfword
	ExtendSXTW                   // Signed extend word
	ExtendSXTX                   // Signed extend doubleword (64-bit, essentially a copy)
)

func (e ExtendType) String() string {
	switch e {
	case ExtendUXTB:
		return "uxtb"
	case ExtendUXTH:
		return "uxth"
	case ExtendUXTW:
		return "uxtw"
	case ExtendUXTX:
		return "uxtx"
	case ExtendSXTB:
		return "sxtb"
	case ExtendSXTH:
		return "sxth"
	case ExtendSXTW:
		return "sxtw"
	case ExtendSXTX:
		return "sxtx"
	default:
		return fmt.Sprintf("extend?%d", e)
	}
}

// Immediate represents an immediate value
type Immediate struct {
	Value int64
}

func (i Immediate) String() string {
	return fmt.Sprintf("#%d", i.Value)
}

// Imm creates an immediate operand
func Imm(value int64) Immediate {
	return Immediate{Value: value}
}
