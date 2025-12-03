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

package abi_test

import (
	"testing"

	"github.com/bytedance/sonic/loader/internal/iasm/arm64"
	"github.com/stretchr/testify/require"
)

func TestSimpleFunctionCodeGeneration(t *testing.T) {
	// Generate a simple function using arm64 iasm: long long simple() { return 42; }
	// Compare with clang -O0 compiled result

	p := arm64.DefaultArch.CreateProgram()
	// Prologue: SUB SP, SP, #16
	p.SUB(arm64.SP, arm64.SP, 16)
	// MOV X8, #42
	p.MOVZ(arm64.X8, 42, 0)
	// STR X8, [SP, #8]
	p.STR(arm64.X8, arm64.Ptr(arm64.SP, 8))
	// LDR X0, [SP, #8]
	p.LDR(arm64.X0, arm64.Ptr(arm64.SP, 8))
	// Epilogue: ADD SP, SP, #16
	p.ADD(arm64.SP, arm64.SP, 16)
	// RET
	p.RET()
	code := p.Assemble(0)

	// Expected bytecode from clang -O0: ff4300d1480580d2e80700f9e00740f9ff430091c0035fd6
	expected := []byte{0xff, 0x43, 0x00, 0xd1, 0x48, 0x05, 0x80, 0xd2, 0xe8, 0x07, 0x00, 0xf9, 0xe0, 0x07, 0x40, 0xf9, 0xff, 0x43, 0x00, 0x91, 0xc0, 0x03, 0x5f, 0xd6}

	require.Equal(t, expected, code, "Generated code should match clang -O0 output")
}
