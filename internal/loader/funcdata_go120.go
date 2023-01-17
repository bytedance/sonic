//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

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

package loader

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytedance/sonic/internal/rt"
)

type moduledata struct {
	pcHeader     *pcHeader
	funcnametab  []byte
	cutab        []uint32
	filetab      []byte
	pctab        []byte
	pclntable    []byte
	ftab         []funcTab
	findfunctab  uintptr
	minpc, maxpc uintptr // first func address, last func address + last func size

	text, etext           uintptr // start/end of text, (etext-text) must be greater than MIN_FUNC
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	covctrs, ecovctrs     uintptr
	end, gcdata, gcbss    uintptr
	types, etypes         uintptr
	rodata                uintptr

    // TODO: generate funcinfo object to memory
	gofunc                uintptr // go.func.* is actual funcinfo object in image

	textsectmap []textSection // see runtime/symtab.go: textAddr()
	typelinks   []int32 // offsets from types
	itablinks   []*rt.GoItab

	ptab []ptabEntry

	pluginpath string
	pkghashes  []modulehash

	modulename   string
	modulehashes []modulehash

	hasmain uint8 // 1 if module contains the main function, 0 otherwise

	gcdatamask, gcbssmask bitVector

	typemap map[typeOff]*rt.GoType // offset to *_rtype in previous module

	bad bool // module failed to load and should be ignored

	next *moduledata
}

// A FuncFlag holds bits about a function.
// This list must match the list in cmd/internal/objabi/funcid.go.
type funcFlag uint8

const (
    _N_PC_DATA = 4
    _N_FUNC_DATA = 4
    _FUNC_SIZE = 11 * 4

	MINFUNC = 16 // minimum size for a function
    BUCKETSIZE    = 256 * MINFUNC
	SUBBUCKETS    = 16
	SUBBUCKETSIZE = BUCKETSIZE / SUBBUCKETS
)

type _Func struct {
    entryOff uint32 // start pc, as offset from moduledata.text/pcHeader.textStart
	nameOff  int32  // function name, as index into moduledata.funcnametab.

	args        int32  // in/out args size
	deferreturn uint32 // offset of start of a deferreturn call instruction from entry, if any.

	pcsp      uint32
	pcfile    uint32
	pcln      uint32
	npcdata   uint32
	cuOffset  uint32 // runtime.cutab offset of this function's CU
	startLine int32  // line number of start of function (func keyword/TEXT directive)
	funcID    uint8 // set for certain special runtime functions
	flag      uint8
	_         [1]byte // pad
	nfuncdata uint8   // 
    
    // The end of the struct is followed immediately by two variable-length
	// arrays that reference the pcdata and funcdata locations for this
	// function.

	// pcdata contains the offset into moduledata.pctab for the start of
	// that index's table. e.g.,
	// &moduledata.pctab[_func.pcdata[_PCDATA_UnsafePoint]] is the start of
	// the unsafe point table.
	//
	// An offset of 0 indicates that there is no table.
	//
	// pcdata [npcdata]uint32

	// funcdata contains the offset past moduledata.gofunc which contains a
	// pointer to that index's funcdata. e.g.,
	// *(moduledata.gofunc +  _func.funcdata[_FUNCDATA_ArgsPointerMaps]) is
	// the argument pointer map.
	//
	// An offset of ^uint32(0) indicates that there is no entry.
	//
	// funcdata [nfuncdata]uint32
}

type funcTab struct {
    entry   uint32
    funcoff uint32
}

type pcHeader struct {
    magic          uint32  // 0xFFFFFFF0
    pad1, pad2     uint8   // 0,0
    minLC          uint8   // min instruction size
    ptrSize        uint8   // size of a ptr in bytes
    nfunc          int     // number of functions in the module
    nfiles         uint    // number of entries in the file tab
    textStart      uintptr // base for function entry PC offsets in this module, equal to moduledata.text
    funcnameOffset uintptr // offset to the funcnametab variable from pcHeader
    cuOffset       uintptr // offset to the cutab variable from pcHeader
    filetabOffset  uintptr // offset to the filetab variable from pcHeader
    pctabOffset    uintptr // offset to the pctab variable from pcHeader
    pclnOffset     uintptr // offset to the pclntab variable from pcHeader
}

type bitVector struct {
    n        int32 // # of bits
    bytedata *uint8
}

type ptabEntry struct {
    name nameOff
    typ  typeOff
}

type textSection struct {
    vaddr    uintptr // prelinked section vaddr
	end      uintptr // vaddr + section length
	baseaddr uintptr // relocated section address
}

type modulehash struct {
	modulename   string
	linktimehash string
	runtimehash  *string
}

type nameOff int32
type typeOff int32
type textOff int32

// findfuncbucket is an array of these structures.
// Each bucket represents 4096 bytes of the text segment.
// Each subbucket represents 256 bytes of the text segment.
// To find a function given a pc, locate the bucket and subbucket for
// that pc. Add together the idx and subbucket value to obtain a
// function index. Then scan the functab array starting at that
// index to find the target function.
// This table uses 20 bytes for every 4096 bytes of code, or ~0.5% overhead.
type findfuncbucket struct {
	idx        uint32
	subbuckets [16]byte
}

func getOffsetOf(data interface{}, field string) uintptr {
    t := reflect.TypeOf(data)
    fv, ok := t.FieldByName(field)
    if !ok {
        panic(fmt.Sprintf("field %s not found in struct %s", field, t.Name()))
    }
    return fv.Offset
}

const _Magic uint32 = 0xFFFFFFF1

func newPCHeader(text uintptr, nfunc int) *pcHeader {
    return &pcHeader{
		magic:          _Magic,
		pad1:           0,
		pad2:           0,
		minLC:          _MinLC,
		ptrSize:        _PtrSize,
		nfunc:          nfunc,
		nfiles:         0,
		textStart:      text,
		funcnameOffset: getOffsetOf(moduledata{}, "funcnametab"),
		cuOffset:       getOffsetOf(moduledata{}, "cutab"),
		filetabOffset:  getOffsetOf(moduledata{}, "filetab"),
		pctabOffset:    getOffsetOf(moduledata{}, "pctab"),
		pclnOffset:     getOffsetOf(moduledata{}, "pclntable"),
	}
}

func funcNameParts(name string) (string, string, string) {
    i := strings.IndexByte(name, '[')
    if i < 0 {
        return name, "", ""
    }
    // TODO: use LastIndexByte once the bootstrap compiler is >= Go 1.5.
    j := len(name) - 1
    for j > i && name[j] != ']' {
        j--
    }
    if j <= i {
        return name, "", ""
    }
    return name[:i], "[...]", name[j+1:]
}

// func name table format: 
//   nameOff[0] -> namePartA namePartB namePartC \x00 
//   nameOff[1] -> namePartA namePartB namePartC \x00
//  ...
func writeFuncNameTab(names []string) (tab []byte, offs []nameOff) {
    offs = make([]nameOff, len(names))
    offset := 0

    for i, name := range names {
        offs[i] = nameOff(offset)

        a, b, c := funcNameParts(name)
        tab = append(tab, a...)
        tab = append(tab, b...)
        tab = append(tab, c...)
        tab = append(tab, 0)
        offset += len(a) + len(b) + len(c) + 1
    }

    return
}

type compilationUnit struct {
    fileNames []string
}

// CU table format:
//  cuOffsets[0] -> filetabOffset[0] filetabOffset[1] ... filetabOffset[len(CUs[0].fileNames)-1]
//  cuOffsets[1] -> filetabOffset[len(CUs[0].fileNames)] ... filetabOffset[len(CUs[0].fileNames) + len(CUs[1].fileNames)-1]
//  ...
//
// file name table format:
//  filetabOffset[0] -> CUs[0].fileNames[0] \x00
//  ...
//  filetabOffset[len(CUs[0]-1)] -> CUs[0].fileNames[len(CUs[0].fileNames)-1] \x00
//  ...
//  filetabOffset[SUM(CUs,fileNames)-1] -> CUs[len(CU)-1].fileNames[len(CUs[len(CU)-1].fileNames)-1] \x00
func writeFilenameTabs(CUs []compilationUnit) (cutab []byte, filetab []byte, cuOffsets []uint32) {
    cuOffsets = make([]uint32, len(CUs))
    cuOffset := 0
    fileOffset := 0

    for i, cu := range CUs {
        cuOffsets[i] = uint32(cuOffset)
        cuOffset += len(cu.fileNames)

        for _, name := range cu.fileNames {
            rt.GuardSlice(&cutab, 4)
            byteOrder.PutUint32(cutab[len(cutab):len(cutab)+4], uint32(fileOffset))
            cutab = cutab[:len(cutab)+4]

            fileOffset += len(name) + 1
            filetab = append(filetab, name...)
            filetab = append(filetab, 0)
        }
    }

    return
}


type _pcdata int32

type pcData struct {
    pcsp   uint32
    pcfile uint32
    pcln   uint32
    others []_pcdata
}

// FuncInfo is serialized as a symbol (aux symbol). The symbol data is
// the binary encoding of the struct below.
type FuncInfo struct {
	Args      uint32
	Locals    uint32
	FuncID    uint8
	FuncFlag  uint8
	StartLine int32
	File      []CUFileIndex
	InlTree   []InlTreeNode
}

// InlTreeNode is the serialized form of FileInfo.InlTree.
type InlTreeNode struct {
	Parent   int32
	File     CUFileIndex
	Line     int32
	Func     SymRef
	ParentPC int32
}

// Symbol reference.
type SymRef struct {
	PkgIdx uint32
	SymIdx uint32
}

// CUFileIndex is used to index the filenames that are stored in the
// per-package/per-CU FileList.
type CUFileIndex uint32

// PCTab format:
//   pcDatas[0].pcsp pcDatas[0].pcfile pcDatas[0].pcln pcdatas[0].others[...] 
//   pcDatas[1].pcsp pcDatas[1].pcfile pcDatas[1].pcln pcdatas[1].others[...]
//   ...
func writePctab(funcs []_Func, pcdatas []pcData) (pctab []byte, pcdataOffs [][]uint32) {
    if len(funcs) != len(pcdatas) {
        panic("len(funcs) != len(pcdatas)")
    }

    // Pctab offsets of 0 are considered invalid in the runtime. We respect
	// that by just padding a single byte at the beginning of runtime.pctab,
	// that way no real offsets can be zero.
    pctab = make([]byte, 1, 12*len(funcs)+1)
    pcdataOffs = make([][]uint32, len(funcs))

    for i, f := range funcs {
        if int(f.npcdata) != len(pcdatas[i].others) + 3 {
            panic("f.npcdata  !=  len(pcdatas)")
        }
        rt.GuardSlice(&pctab, 12)

        l := len(pctab)
        f.pcsp = uint32(l)
        byteOrder.PutUint32(pctab[l:l+4], uint32(pcdatas[i].pcsp))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l))

        f.pcfile = uint32(l + 4)
        byteOrder.PutUint32(pctab[l+4:l+8], uint32(pcdatas[i].pcfile))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l+4))

        f.pcln = uint32(l + 8)
        byteOrder.PutUint32(pctab[l+8:l+12], uint32(pcdatas[i].pcln))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l+8))

        pctab = pctab[:len(pctab)+12]

        for _, pcdata := range pcdatas[i].others {
            rt.GuardSlice(&pctab, 4)
            byteOrder.PutUint32(pctab[len(pctab):len(pctab)+4], uint32(pcdata))
            pcdataOffs[i] = append(pcdataOffs[i], uint32(len(pctab)))
            pctab = pctab[:len(pctab)+4]
        }
    }

    return
}

// Pcln table format: [...]funcTab + [...]_Func
func writeFunctable(funcs []_Func, lastFuncSize uint32, gofuncBase uintptr, pcdataOffs [][]uint32, funcdataOffs [][]uint32) (pclntab []byte, ftab []funcTab) {
    // Allocate space for the pc->func table. This structure consists of a pc offset
	// and an offset to the func structure. After that, we have a single pc
	// value that marks the end of the last function in the binary.
    var size int64 = int64(len(funcs)*2*4 + 4)
    var startLocations = make([]uint32, len(funcs))
    for i := range funcs {
        size = Rnd(size, int64(_PtrSize))
        //writePCToFunc
        startLocations[i] = uint32(size)
        size += int64(_FUNC_SIZE+len(funcdataOffs[i])*4+len(pcdataOffs[i])*4)
    }

    pclntab = make([]byte, 0, size)
    ftab = make([]funcTab, 0, len(funcs)+1)

    // write a map of pc->func info offsets 
    for i, f := range funcs {
        ftab = append(ftab, funcTab{uint32(f.entryOff), uint32(startLocations[i])})

        byteOrder.PutUint32(pclntab[len(pclntab):len(pclntab)+4], uint32(f.entryOff))
        byteOrder.PutUint32(pclntab[len(pclntab)+4:len(pclntab)+8], uint32(startLocations[i]))
        pclntab = pclntab[:len(pclntab)+8]
    }

    // Final entry of table is just end pc offset.
    lastFunc := funcs[len(funcs)-1]
    ftab = append(ftab, funcTab{uint32(lastFunc.entryOff + lastFuncSize), 0})
    byteOrder.PutUint32(pclntab[len(pclntab):len(pclntab)+4], uint32(lastFunc.entryOff+lastFuncSize))

    // write func info table
    for i, f := range funcs {
        // functiono structure
        off := startLocations[i]
        byteOrder.PutUint32(pclntab[off:off+4], uint32(f.entryOff))
        byteOrder.PutUint32(pclntab[off+4:off+8], uint32(f.nameOff))
        byteOrder.PutUint32(pclntab[off+8:off+12], uint32(f.args))
        byteOrder.PutUint32(pclntab[off+12:off+16], uint32(f.deferreturn))
        byteOrder.PutUint32(pclntab[off+16:off+20], uint32(f.pcsp))
        byteOrder.PutUint32(pclntab[off+20:off+24], uint32(f.pcfile))
        byteOrder.PutUint32(pclntab[off+24:off+28], uint32(f.pcln))
        byteOrder.PutUint32(pclntab[off+28:off+32], uint32(f.npcdata))
        byteOrder.PutUint32(pclntab[off+32:off+36], uint32(f.cuOffset))
        byteOrder.PutUint32(pclntab[off+36:off+40], uint32(f.startLine))
        pclntab[off+40] = f.funcID
        pclntab[off+41] = f.flag
        pclntab[off+43] = f.nfuncdata

        // pcdata
        off += _FUNC_SIZE
        for _, pcdata := range pcdataOffs[i] {
            byteOrder.PutUint32(pclntab[off:off+4], uint32(pcdata))
            off += 4
        }

        // Write funcdata refs as offsets from go:func.* and go:funcrel.*.
        for _, funcdata := range funcdataOffs[i] {
            byteOrder.PutUint32(pclntab[off:off+4], uint32(funcdata-uint32(gofuncBase)))
            off += 4
        }
    }

    return
}

// findfunc table used to find the funcinfo by pc
// see findfunc() in runtime/symtab.go
func makeFindfunctab(ftab []funcTab) (tab []findfuncbucket) {
    max := ftab[len(ftab)-1].entry
    min := ftab[0].entry
    nbuckets := (max - min + BUCKETSIZE - 1) / BUCKETSIZE
    n := (max - min + SUBBUCKETSIZE - 1) / SUBBUCKETSIZE

    tab = make([]findfuncbucket, 0, nbuckets)

    var s, e = 0, 0
    for i := 0; i<int(nbuckets); i++ {
        var pc = min + uint32((i+1)*BUCKETSIZE)
        // find the end func of the bucket
        for ; e < len(ftab)-1 && ftab[e+1].entry <= pc; e++ {}
        // store the start func of the bucket
        var fb = findfuncbucket{idx: uint32(s)}

        for j := 0; j<SUBBUCKETS && (i*SUBBUCKETS+j)<int(n); j++ {
            pc = min + uint32(i*BUCKETSIZE) + uint32((j+1)*SUBBUCKETSIZE)
            var ss = s
            // find the end func of the subbucket
            for ; ss < len(ftab)-1 && ftab[ss+1].entry <= pc; ss++ {}
            // store the start func of the subbucket
            fb.subbuckets[j] = byte(uint32(s) - fb.idx)
            s = ss
        }
        s = e
        tab = append(tab, fb)
    }

    return 
}

func registerFunction(name string, pc uintptr, textSize uintptr, fp int, args int, size uintptr, argptrs uintptr, localptrs uintptr) {
    
}