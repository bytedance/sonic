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

// Register represents a hardware register.
type Register interface {
	fmt.Stringer
	implRegister()
}

// General Purpose Registers (64-bit)
type Register64 byte

// General Purpose Registers (32-bit)
type Register32 byte

// NEON/FP Registers (128-bit vector)
type VRegister byte

// NEON/FP Registers (64-bit)
type DRegister byte

// NEON/FP Registers (32-bit)
type SRegister byte

// NEON/FP Registers (16-bit half precision)
type HRegister byte

// NEON/FP Registers (8-bit)
type BRegister byte

// 64-bit General Purpose Registers
const (
	X0 Register64 = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	XZR // Zero Register
	SP  // Stack Pointer (special encoding)
)

// 32-bit General Purpose Registers
const (
	W0 Register32 = iota
	W1
	W2
	W3
	W4
	W5
	W6
	W7
	W8
	W9
	W10
	W11
	W12
	W13
	W14
	W15
	W16
	W17
	W18
	W19
	W20
	W21
	W22
	W23
	W24
	W25
	W26
	W27
	W28
	W29
	W30
	WZR // Zero Register
	WSP // Stack Pointer (special encoding)
)

// NEON/FP Vector Registers (128-bit)
const (
	V0 VRegister = iota
	V1
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	V9
	V10
	V11
	V12
	V13
	V14
	V15
	V16
	V17
	V18
	V19
	V20
	V21
	V22
	V23
	V24
	V25
	V26
	V27
	V28
	V29
	V30
	V31
)

// NEON/FP Registers (64-bit double precision)
const (
	D0 DRegister = iota
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	D9
	D10
	D11
	D12
	D13
	D14
	D15
	D16
	D17
	D18
	D19
	D20
	D21
	D22
	D23
	D24
	D25
	D26
	D27
	D28
	D29
	D30
	D31
)

// NEON/FP Registers (32-bit single precision)
const (
	S0 SRegister = iota
	S1
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
	S10
	S11
	S12
	S13
	S14
	S15
	S16
	S17
	S18
	S19
	S20
	S21
	S22
	S23
	S24
	S25
	S26
	S27
	S28
	S29
	S30
	S31
)

// Register interface implementations
func (Register64) implRegister() {}
func (Register32) implRegister() {}
func (VRegister) implRegister()  {}
func (DRegister) implRegister()  {}
func (SRegister) implRegister()  {}
func (HRegister) implRegister()  {}
func (BRegister) implRegister()  {}

// String implementations for Register64
func (r Register64) String() string {
	switch r {
	case X0:
		return "x0"
	case X1:
		return "x1"
	case X2:
		return "x2"
	case X3:
		return "x3"
	case X4:
		return "x4"
	case X5:
		return "x5"
	case X6:
		return "x6"
	case X7:
		return "x7"
	case X8:
		return "x8"
	case X9:
		return "x9"
	case X10:
		return "x10"
	case X11:
		return "x11"
	case X12:
		return "x12"
	case X13:
		return "x13"
	case X14:
		return "x14"
	case X15:
		return "x15"
	case X16:
		return "x16"
	case X17:
		return "x17"
	case X18:
		return "x18"
	case X19:
		return "x19"
	case X20:
		return "x20"
	case X21:
		return "x21"
	case X22:
		return "x22"
	case X23:
		return "x23"
	case X24:
		return "x24"
	case X25:
		return "x25"
	case X26:
		return "x26"
	case X27:
		return "x27"
	case X28:
		return "x28"
	case X29:
		return "x29"
	case X30:
		return "x30"
	case XZR:
		return "xzr"
	case SP:
		return "sp"
	default:
		return fmt.Sprintf("x?%d", r)
	}
}

// String implementations for Register32
func (r Register32) String() string {
	switch r {
	case WZR:
		return "wzr"
	case WSP:
		return "wsp"
	default:
		if r <= W30 {
			return fmt.Sprintf("w%d", r)
		}
		return fmt.Sprintf("w?%d", r)
	}
}

// String implementations for VRegister
func (r VRegister) String() string {
	if r <= V31 {
		return fmt.Sprintf("v%d", r)
	}
	return fmt.Sprintf("v?%d", r)
}

// String implementations for DRegister
func (r DRegister) String() string {
	if r <= D31 {
		return fmt.Sprintf("d%d", r)
	}
	return fmt.Sprintf("d?%d", r)
}

// String implementations for SRegister
func (r SRegister) String() string {
	if r <= S31 {
		return fmt.Sprintf("s%d", r)
	}
	return fmt.Sprintf("s?%d", r)
}

// RegIndex returns the register index (0-31)
func (r Register64) RegIndex() byte {
	if r == SP {
		return 31
	}
	if r == XZR {
		return 31
	}
	return byte(r)
}

// RegIndex returns the register index (0-31)
func (r Register32) RegIndex() byte {
	if r == WSP || r == WZR {
		return 31
	}
	return byte(r)
}

// RegIndex returns the register index (0-31)
func (r VRegister) RegIndex() byte {
	return byte(r)
}

// RegIndex returns the register index (0-31)
func (r DRegister) RegIndex() byte {
	return byte(r)
}

// RegIndex returns the register index (0-31)
func (r SRegister) RegIndex() byte {
	return byte(r)
}

// Aliases for common register names
const (
	// X29 is the Frame Pointer
	FP = X29
	// X30 is the Link Register
	LR = X30
)
