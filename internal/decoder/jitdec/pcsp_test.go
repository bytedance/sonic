/*
* Copyright 2025 ByteDance Inc.
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type _MockDecoder struct {
    jit.BaseAssembler
}

func (self *_MockDecoder) compile() {
    self.Emit("SUBQ", jit.Imm(_VD_size), _SP)       // SUBQ $_VD_size, SP
    self.Byte(0xcc) // INT3
    self.Emit("MOVQ", _BP, jit.Ptr(_SP, _VD_offs))  // MOVQ BP, _VD_offs(SP)
    self.Byte(0xcc) // INT3

    self.Emit("MOVQ", jit.Ptr(_SP, _VD_offs), _BP)  // MOVQ _VD_offs(SP), BP
    self.Emit("MOVQ", jit.Imm(199), _AX)
    self.Byte(0xcc) // INT3
    self.Emit("ADDQ", jit.Imm(_VD_size), _SP)       // ADDQ $_VD_size, SP
    self.Byte(0xcc) // INT3
    self.Emit("RET")
}

func (self *_MockDecoder) Load() _Decoder {
    self.Init(self.compile)
    return ptodec(self.BaseAssembler.Load("decode_mock", _FP_size, _FP_args, argPtrs, localPtrs))
}

// copied from g01.20.4 signal_linux_amd64.go
type sigctxt struct {
	info unsafe.Pointer
	ctxt unsafe.Pointer
}

type stackt struct {
	ss_sp     *byte
	ss_flags  int32
	pad_cgo_0 [4]byte
	ss_size   uintptr
}

type mcontext struct {
	gregs       [23]uint64
	fpregs      unsafe.Pointer
	__reserved1 [8]uint64
}

type sigcontext struct {
	r8          uint64
	r9          uint64
	r10         uint64
	r11         uint64
	r12         uint64
	r13         uint64
	r14         uint64
	r15         uint64
	rdi         uint64
	rsi         uint64
	rbp         uint64
	rbx         uint64
	rdx         uint64
	rax         uint64
	rcx         uint64
	rsp         uint64
	rip         uint64
	eflags      uint64
	cs          uint16
	gs          uint16
	fs          uint16
	__pad0      uint16
	err         uint64
	trapno      uint64
	oldmask     uint64
	cr2         uint64
	fpstate     unsafe.Pointer
	__reserved1 [8]uint64
}
type ucontext struct {
	uc_flags     uint64
	uc_link      *ucontext
	uc_stack     stackt
	uc_mcontext  mcontext
}

//go:nosplit
func (c *sigctxt) regs() *sigcontext {
	return (*sigcontext)(unsafe.Pointer(&(*ucontext)(c.ctxt).uc_mcontext))
}

func (c *sigctxt) rsp() uint64 { return c.regs().rsp }

//go:nosplit
func (c *sigctxt) sigpc() uintptr { return uintptr(c.rip()) }

//go:nosplit
func (c *sigctxt) rip() uint64 { return c.regs().rip }
func (c *sigctxt) sigsp() uintptr    { return uintptr(c.rsp()) }
func (c *sigctxt) siglr() uintptr    { return 0 }

// only used for test sonic trace
//go:linkname testSigtrap runtime.testSigtrap
var testSigtrap func(info unsafe.Pointer, c *sigctxt, gp unsafe.Pointer) bool 

//go:linkname traceback1 runtime.traceback1
func traceback1(pc, sp, lr uintptr, gp unsafe.Pointer, flags uint);

func sonicSigTrap(info unsafe.Pointer, c *sigctxt, gp unsafe.Pointer) bool {
	pc := c.sigpc()
	sp := c.sigsp()
	lr := c.siglr()
	traceback1(pc, sp, lr, gp, 0);
	return true
}

func TestAssembler_PCSP(t *testing.T) {
    testSigtrap = sonicSigTrap
    f := new(_MockDecoder).Load()
    pos, err := f("", 0, nil, nil, 0, "", nil)
    require.NoError(t, err)
    assert.Equal(t, 199, pos)
    testSigtrap = nil
}