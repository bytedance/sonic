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

package jitdec

import (
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/loader"
)

type x15StoreFunc func(*[16]byte)
type x15DirtyFunc func()

func loadDirtyX15Func(t *testing.T, name string) x15DirtyFunc {
	t.Helper()
	// PCMPEQD X15, X15; RET.
	text := []byte{0x66, 0x45, 0x0f, 0x76, 0xff, 0xc3}
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

func loadGenericCallCStoreX15(t *testing.T, dirty x15DirtyFunc) x15StoreFunc {
	t.Helper()
	var a _ValueDecoder
	a.BaseAssembler.Init(func() {
		a.Emit("SUBQ", jit.Imm(_VD_size), _SP)
		a.Emit("MOVQ", _BP, jit.Ptr(_SP, _VD_offs))
		a.Emit("LEAQ", jit.Ptr(_SP, _VD_offs), _BP)
		a.callc(jit.Func(dirty))
		a.Emit("MOVUPS", _X15, jit.Ptr(_AX, 0))
		a.Emit("MOVQ", jit.Ptr(_SP, _VD_offs), _BP)
		a.Emit("ADDQ", jit.Imm(_VD_size), _SP)
		a.Emit("RET")
	})
	p := a.BaseAssembler.Load("test_generic_callc_x15", _VD_size, 8, []bool{true}, nil)
	return *(*x15StoreFunc)(unsafe.Pointer(&p))
}

func TestGenericCallCClearsX15(t *testing.T) {
	dirty := loadDirtyX15Func(t, "dirty_x15_generic")
	call := loadGenericCallCStoreX15(t, dirty)
	var got [16]byte
	call(&got)
	for i, b := range got {
		if b != 0 {
			t.Fatalf("X15 byte %d = %#x, want 0; full=% x", i, b, got)
		}
	}
}
