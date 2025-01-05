//
// Copyright 2024 CloudWeGo Authors
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

package x86_64

import (
	"bytes"
	"testing"

	"github.com/bytedance/sonic/loader/internal/iasm/expr"
	"github.com/davecgh/go-spew/spew"
)

func TestProgram_Assemble(t *testing.T) {
	a := CreateArch()
	b := CreateLabel("bak")
	s := CreateLabel("tab")
	j := CreateLabel("jmp")
	p := a.CreateProgram()
	p.JMP(j)
	p.JMP(j)
	p.Link(b)
	p.Data(bytes.Repeat([]byte{0x0f, 0x1f, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00}, 15))
	p.Data([]byte{0x0f, 0x1f, 0x00})
	p.JMP(b)
	p.Link(j)
	p.LEAQ(Ref(s), RDI)
	p.MOVSLQ(Sib(RDI, RAX, 4, -4), RAX)
	p.ADDQ(RDI, RAX)
	p.JMPQ(RAX)
	p.Align(32, expr.Int(0xcc))
	p.Link(s)
	p.Long(expr.Ref(s.Retain()).Sub(expr.Ref(j.Retain())))
	p.Long(expr.Ref(s.Retain()).Sub(expr.Ref(b.Retain())))
	spew.Dump(p.AssembleAndFree(0))
}
