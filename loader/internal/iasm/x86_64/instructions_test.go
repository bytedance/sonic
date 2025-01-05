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
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestInstr_Encode(t *testing.T) {
	m := []byte(nil)
	a := CreateArch()
	p := a.CreateProgram()
	p.VPERMIL2PD(7, Sib(R8, R9, 1, 12345), YMM1, YMM2, YMM3).encode(&m)
	spew.Dump(m)
}

func TestInstr_EncodeSegment(t *testing.T) {
	m := []byte(nil)
	a := CreateArch()
	p := a.CreateProgram()
	p.MOVQ(Abs(0x30), RCX).GS().encode(&m)
	spew.Dump(m)
}

func BenchmarkInstr_Encode(b *testing.B) {
	a := CreateArch()
	m := make([]byte, 0, 16)
	p := a.CreateProgram()
	p.VPERMIL2PD(7, Sib(R8, R9, 1, 12345), YMM1, YMM2, YMM3).encode(&m)
	p.Free()
	b.SetBytes(int64(len(m)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m = m[:0]
		p = a.CreateProgram()
		p.VPERMIL2PD(7, Sib(R8, R9, 1, 12345), YMM1, YMM2, YMM3).encode(&m)
		p.Free()
	}
}
