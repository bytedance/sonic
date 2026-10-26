//go:build go1.17 && !go1.28
// +build go1.17,!go1.28

/*
 * Copyright 2021 ByteDance Inc.
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

package x86

import (
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/loader"
)

type x15StoreFunc func(*[16]byte)
type x15DirtyFunc func()

var x15TestReg = jit.Reg("X15")

func loadDirtyX15Func(t *testing.T, name string) x15DirtyFunc {
	t.Helper()
	// PCMPEQD X15, X15; XORL AX, AX; RET.
	// Native calls may clobber AX, which initially holds the output pointer.
	text := []byte{0x66, 0x45, 0x0f, 0x76, 0xff, 0x31, 0xc0, 0xc3}
	l := loader.Loader{
		Name: "sonic.test.",
		File: "x15_regabi_amd64_test.go",
		Options: loader.Options{
			NoPreempt: true,
		},
	}
	p := l.LoadOne(text, name, 0, 0, nil, nil, loader.Pcdata{})
	return *(*x15DirtyFunc)(unsafe.Pointer(&p))
}

func loadEncoderCallStoreX15(t *testing.T, name string, dirty x15DirtyFunc, callKind string) x15StoreFunc {
	t.Helper()
	var a Assembler
	a.BaseAssembler.Init(func() {
		a.Emit("SUBQ", jit.Imm(_FP_size), _SP)
		a.Emit("MOVQ", _BP, jit.Ptr(_SP, FP_offs))
		a.Emit("LEAQ", jit.Ptr(_SP, FP_offs), _BP)
		// Preserve the output pointer in its ABI argument spill slot.
		a.Emit("MOVQ", _AX, jit.Ptr(_SP, _FP_size+8))
		a.save_c()
		switch callKind {
		case "call_c":
			a.call_c(jit.Func(dirty))
		case "call_b64":
			a.call_b64(jit.Func(dirty))
		default:
			t.Fatalf("unknown native call wrapper %q", callKind)
		}
		a.Emit("MOVQ", jit.Ptr(_SP, _FP_size+8), _AX)
		a.Emit("MOVUPS", x15TestReg, jit.Ptr(_AX, 0))
		// Return to Go with a clean X15 even when the wrapper regresses.
		a.Emit("XORPS", x15TestReg, x15TestReg)
		a.Emit("MOVQ", jit.Ptr(_SP, FP_offs), _BP)
		a.Emit("ADDQ", jit.Imm(_FP_size), _SP)
		a.Emit("RET")
	})
	p := a.BaseAssembler.Load(name, _FP_size, 8, []bool{true}, nil)
	return *(*x15StoreFunc)(unsafe.Pointer(&p))
}

func TestEncoderCallCClearsX15(t *testing.T) {
	dirty := loadDirtyX15Func(t, "dirty_x15_encoder")
	call := loadEncoderCallStoreX15(t, "test_encoder_callc_x15", dirty, "call_c")
	var got [16]byte
	for i := range got {
		got[i] = 0xa5
	}
	call(&got)
	for i, b := range got {
		if b != 0 {
			t.Fatalf("X15 byte %d = %#x, want 0; full=% x", i, b, got)
		}
	}
}

func TestEncoderCallB64ClearsX15(t *testing.T) {
	dirty := loadDirtyX15Func(t, "dirty_x15_encoder_b64")
	call := loadEncoderCallStoreX15(t, "test_encoder_callb64_x15", dirty, "call_b64")
	var got [16]byte
	for i := range got {
		got[i] = 0xa5
	}
	call(&got)
	for i, b := range got {
		if b != 0 {
			t.Fatalf("X15 byte %d = %#x, want 0; full=% x", i, b, got)
		}
	}
}
