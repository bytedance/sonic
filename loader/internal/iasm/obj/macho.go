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
	"encoding/binary"
	"io"
	"unsafe"
)

type MachOHeader struct {
	Magic      uint32
	CPUType    uint32
	CPUSubType uint32
	FileType   uint32
	CmdCount   uint32
	CmdSize    uint32
	Flags      uint32
	_          uint32
}

type SegmentCommand struct {
	Cmd          uint32
	Size         uint32
	Name         [16]byte
	VMAddr       uint64
	VMSize       uint64
	FileOffset   uint64
	FileSize     uint64
	MaxProtect   uint32
	InitProtect  uint32
	SectionCount uint32
	Flags        uint32
}

type SegmentSection struct {
	Name      [16]byte
	SegName   [16]byte
	Addr      uint64
	Size      uint64
	Offset    uint32
	Align     uint32
	RelOffset uint32
	RelCount  uint32
	Flags     uint32
	_         [3]uint32
}

type Registers struct {
	RAX    uint64
	RBX    uint64
	RCX    uint64
	RDX    uint64
	RDI    uint64
	RSI    uint64
	RBP    uint64
	RSP    uint64
	R8     uint64
	R9     uint64
	R10    uint64
	R11    uint64
	R12    uint64
	R13    uint64
	R14    uint64
	R15    uint64
	RIP    uint64
	RFLAGS uint64
	CS     uint64
	FS     uint64
	GS     uint64
}

type UnixThreadCommand struct {
	Cmd    uint32
	Size   uint32
	Flavor uint32
	Count  uint32
	Regs   Registers
}

const (
	_MH_MAGIC_64 = 0xfeedfacf
	_MH_EXECUTE  = 0x02
	_MH_NOUNDEFS = 0x01
)

const (
	_CPU_TYPE_I386  = 0x00000007
	_CPU_ARCH_ABI64 = 0x01000000
)

const (
	_CPU_SUBTYPE_LIB64    = 0x80000000
	_CPU_SUBTYPE_I386_ALL = 0x00000003
)

const (
	_LC_SEGMENT_64 = 0x19
	_LC_UNIXTHREAD = 0x05
)

const (
	_VM_PROT_READ    = 0x01
	_VM_PROT_WRITE   = 0x02
	_VM_PROT_EXECUTE = 0x04
)

const (
	_S_ATTR_SOME_INSTRUCTIONS = 0x00000400
	_S_ATTR_PURE_INSTRUCTIONS = 0x80000000
)

const (
	_x86_THREAD_STATE64          = 0x04
	_x86_EXCEPTION_STATE64_COUNT = 42
)

const (
	_MACHO_SIZE      = uint32(unsafe.Sizeof(MachOHeader{}))
	_SEGMENT_SIZE    = uint32(unsafe.Sizeof(SegmentCommand{}))
	_SECTION_SIZE    = uint32(unsafe.Sizeof(SegmentSection{}))
	_UNIXTHREAD_SIZE = uint32(unsafe.Sizeof(UnixThreadCommand{}))
)

const (
	_IMAGE_SIZE = 4096
	_IMAGE_BASE = 0x04000000
)

const (
	_HDR_SIZE  = _MACHO_SIZE + _SEGMENT_SIZE*2 + _SECTION_SIZE + _UNIXTHREAD_SIZE
	_ZERO_SIZE = (_IMAGE_SIZE - _HDR_SIZE%_IMAGE_SIZE) % _IMAGE_SIZE
)

var (
	zeroBytes = [_ZERO_SIZE]byte{}
)

func assembleMachO(w io.Writer, code []byte, base uint64, entry uint64) error {
	var p0name [16]byte
	var txname [16]byte
	var tsname [16]byte

	/* segment names */
	copy(tsname[:], "__text")
	copy(txname[:], "__TEXT")
	copy(p0name[:], "__PAGEZERO")

	/* calculate size of code */
	clen := uint64(len(code))
	hlen := uint64(_HDR_SIZE + _ZERO_SIZE)

	/* Mach-O does not allow image base at zero */
	if base == 0 {
		base = _IMAGE_BASE
		entry += _IMAGE_BASE
	}

	/* Page-0 Segment */
	p0 := SegmentCommand{
		Cmd:    _LC_SEGMENT_64,
		Size:   _SEGMENT_SIZE,
		Name:   p0name,
		VMSize: base,
	}

	/* TEXT Segment */
	text := SegmentCommand{
		Cmd:          _LC_SEGMENT_64,
		Size:         _SEGMENT_SIZE + _SECTION_SIZE,
		Name:         txname,
		VMAddr:       base,
		VMSize:       hlen + clen,
		FileSize:     hlen + clen,
		MaxProtect:   _VM_PROT_READ | _VM_PROT_WRITE | _VM_PROT_EXECUTE,
		InitProtect:  _VM_PROT_READ | _VM_PROT_EXECUTE,
		SectionCount: 1,
	}

	/* __TEXT.__text section */
	tsec := SegmentSection{
		Name:    tsname,
		SegName: txname,
		Addr:    base + hlen,
		Size:    clen,
		Offset:  uint32(hlen),
		Flags:   _S_ATTR_SOME_INSTRUCTIONS | _S_ATTR_PURE_INSTRUCTIONS,
	}

	/* UNIX Thread Metadata */
	unix := UnixThreadCommand{
		Cmd:    _LC_UNIXTHREAD,
		Size:   _UNIXTHREAD_SIZE,
		Flavor: _x86_THREAD_STATE64,
		Count:  _x86_EXCEPTION_STATE64_COUNT,
		Regs:   Registers{RIP: hlen + entry},
	}

	/* Mach-O Header */
	macho := MachOHeader{
		Magic:      _MH_MAGIC_64,
		CPUType:    _CPU_ARCH_ABI64 | _CPU_TYPE_I386,
		CPUSubType: _CPU_SUBTYPE_LIB64 | _CPU_SUBTYPE_I386_ALL,
		FileType:   _MH_EXECUTE,
		CmdCount:   3,
		CmdSize:    _SEGMENT_SIZE*2 + _SECTION_SIZE + _UNIXTHREAD_SIZE,
		Flags:      _MH_NOUNDEFS,
	}

	/* write the headers */
	if err := binary.Write(w, binary.LittleEndian, &macho); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &p0); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &text); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &tsec); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &unix); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &zeroBytes); err != nil {
		return err
	}

	/* write the code */
	if n, err := w.Write(code); err != nil {
		return err
	} else if n != len(code) {
		return io.ErrShortWrite
	} else {
		return nil
	}
}
