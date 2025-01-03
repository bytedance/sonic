//go:build darwin
// +build darwin

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

package obj

import (
	"os"
	"os/exec"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestMachO_Create(t *testing.T) {
	fp, err := os.CreateTemp("", "macho_out-")
	require.NoError(t, err)
	code := []byte{
		0x48, 0xc7, 0xc7, 0x01, 0x00, 0x00, 0x00, // MOVQ    $1, %rdi
		0x48, 0x8d, 0x35, 0x1b, 0x00, 0x00, 0x00, // LEAQ    0x1b(%rip), %rsi
		0x48, 0xc7, 0xc2, 0x0e, 0x00, 0x00, 0x00, // MOVQ    $14, %rdx
		0x48, 0xc7, 0xc0, 0x04, 0x00, 0x00, 0x02, // MOVQ    $0x02000004, %rax
		0x0f, 0x05, // SYSCALL
		0x31, 0xff, // XORL    %edi, %edi
		0x48, 0xc7, 0xc0, 0x01, 0x00, 0x00, 0x02, // MOVQ    $0x02000001, %rax
		0x0f, 0x05, // SYSCALL
		'h', 'e', 'l', 'l', 'o', ',', ' ',
		'w', 'o', 'r', 'l', 'd', '\r', '\n',
	}
	err = assembleMachO(fp, code, 0, 0)
	require.NoError(t, err)
	err = fp.Close()
	require.NoError(t, err)
	err = os.Chmod(fp.Name(), 0755)
	require.NoError(t, err)
	println("Saved to", fp.Name())
	out, err := exec.Command(fp.Name()).Output()
	require.NoError(t, err)
	spew.Dump(out)
	require.Equal(t, []byte("hello, world\r\n"), out)
	err = os.Remove(fp.Name())
	require.NoError(t, err)
}
