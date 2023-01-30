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
	"os"
	"reflect"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

const (
    _Magic uint32 = 0xFFFFFFF1

    _N_PC_DATA = 4
    _N_FUNC_DATA = 4
    _FUNC_SIZE = 11 * 4

	MINFUNC = 16 // minimum size for a function
    BUCKETSIZE    = 256 * MINFUNC
	SUBBUCKETS    = 16
	SUBBUCKETSIZE = BUCKETSIZE / SUBBUCKETS
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

	typemap map[int32]*rt.GoType // offset to *_rtype in previous module

	bad bool // module failed to load and should be ignored

	next *moduledata
}

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
    name int32
    typ  int32
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

type FuncInfo struct {
	ID          uint8
	Flag        uint8
    Args        int32 // args count
    EntryOff    uint32 // start pc, offset to moduledata.text
    TextSize    uint32
    StartLine   int32 
    DeferReturn uint32 // offset of start of a deferreturn call instruction from entry, if any.
    FileName    string
    Pcsp        uint32
    Pcfile      uint32
    Pcline      uint32
    Name        string
    Pcdata      []Pcvalue
    Funcdata    Funcdata
}

type Pcvalue struct {
    Val   int32
    Delta int32
}

func (Pcvalue) Marshal() ([]byte, error)

type Funcdata struct {
    ArgsPointerMaps   StackMap  
    LocalsPointerMaps StackMap
}

type StackMap struct {
	n        int32   // number of bitmaps
	nbit     int32   // number of bits in each bitmap
	bytedata [1]byte // bitmaps, each starting on a byte boundary
}

func (self StackMap) Marshal() ([]byte, error)

func getOffsetOf(data interface{}, field string) uintptr {
    t := reflect.TypeOf(data)
    fv, ok := t.FieldByName(field)
    if !ok {
        panic(fmt.Sprintf("field %s not found in struct %s", field, t.Name()))
    }
    return fv.Offset
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
func makeFuncnameTab(funcs []FuncInfo) (tab []byte, offs []int32) {
    offs = make([]int32, len(funcs))
    offset := 0

    for i, f := range funcs {
        offs[i] = int32(offset)

        a, b, c := funcNameParts(f.Name)
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
func makeFilenametab(cus []compilationUnit) (cutab []uint32, filetab []byte, cuOffsets []uint32) {
    cuOffsets = make([]uint32, len(cus))
    cuOffset := 0
    fileOffset := 0

    for i, cu := range cus {
        cuOffsets[i] = uint32(cuOffset)

        for _, name := range cu.fileNames {
            cutab = append(cutab, uint32(fileOffset))

            fileOffset += len(name) + 1
            filetab = append(filetab, name...)
            filetab = append(filetab, 0)
        }

        cuOffset += len(cu.fileNames)
    }

    return
}

// PCTab format:
//   pcDatas[0].pcsp pcDatas[0].pcfile pcDatas[0].pcln pcdatas[0].others[...] 
//   pcDatas[1].pcsp pcDatas[1].pcfile pcDatas[1].pcln pcdatas[1].others[...]
//   ...
func makePctab(funcs []FuncInfo, cuOffset []uint32, nameOffset []int32) (pctab []byte, pcdataOffs [][]uint32, _funcs []_Func) {
    _funcs = make([]_Func, len(funcs))

    // Pctab offsets of 0 are considered invalid in the runtime. We respect
	// that by just padding a single byte at the beginning of runtime.pctab,
	// that way no real offsets can be zero.
    pctab = make([]byte, 1, 12*len(funcs)+1)
    pcdataOffs = make([][]uint32, len(funcs))

    for i, f := range funcs {
        _f := &_funcs[i]

        rt.GuardSlice(&pctab, 12)
        l := len(pctab)
        _f.pcsp = uint32(l)
        byteOrder.PutUint32(pctab[l:l+4], uint32(f.Pcsp))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l))

        _f.pcfile = uint32(l + 4)
        byteOrder.PutUint32(pctab[l+4:l+8], uint32(f.Pcfile))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l+4))

        _f.pcln = uint32(l + 8)
        byteOrder.PutUint32(pctab[l+8:l+12], uint32(f.Pcline))
        pcdataOffs[i] = append(pcdataOffs[i], uint32(l+8))

        pctab = pctab[:len(pctab)+12]

        for _, pc := range f.Pcdata {
            pcdataOffs[i] = append(pcdataOffs[i], uint32(len(pctab)))
            pb, err := pc.Marshal()
            if err != nil {
                panic(err)
            }
            pctab = append(pctab, pb...)
        }
        
        _f.entryOff = f.EntryOff
        _f.nameOff = nameOffset[i]
        _f.args = f.Args
        _f.deferreturn = f.DeferReturn
        _f.npcdata = uint32(len(f.Pcdata)+3)
        _f.cuOffset = cuOffset[i]
        _f.startLine = f.StartLine
        _f.funcID = f.ID
        _f.flag = f.Flag
        _f.nfuncdata = 2 // TODO: how to get this value dynamically?
    }

    return
}

func writeFuncdata(out *[]byte, funcs []FuncInfo) (funcdataOffs [][]uint32) {
    offs := uint32(0)
    for i, f := range funcs {

        // write ArgsPointerMaps
        funcdataOffs[i] = append(funcdataOffs[i], offs)
        ab, err := f.Funcdata.ArgsPointerMaps.Marshal()
        if err != nil {
            panic(err)
        }
        *out = append(*out, ab...)
        offs += uint32(len(ab))

        // write LocalsPointerMaps
        funcdataOffs[i] = append(funcdataOffs[i], offs)
        lb, err := f.Funcdata.LocalsPointerMaps.Marshal()
        if err != nil {
            panic(err)
        }
        *out = append(*out, lb...)
        offs += uint32(len(lb))
    }
    return 
}

func makeFtab(funcs []_Func, lastFuncSize uint32) (ftab []funcTab) {
    // Allocate space for the pc->func table. This structure consists of a pc offset
	// and an offset to the func structure. After that, we have a single pc
	// value that marks the end of the last function in the binary.
    var size int64 = int64(len(funcs)*2*4 + 4)
    var startLocations = make([]uint32, len(funcs))
    for i, f := range funcs {
        size = rnd(size, int64(_PtrSize))
        //writePCToFunc
        startLocations[i] = uint32(size)
        size += int64(_FUNC_SIZE+f.nfuncdata*4+uint8(f.npcdata)*4)
    }

    ftab = make([]funcTab, 0, len(funcs)+1)

    // write a map of pc->func info offsets 
    for i, f := range funcs {
        ftab = append(ftab, funcTab{uint32(f.entryOff), uint32(startLocations[i])})
    }

    // Final entry of table is just end pc offset.
    lastFunc := funcs[len(funcs)-1]
    ftab = append(ftab, funcTab{uint32(lastFunc.entryOff + lastFuncSize), 0})

    return
}

// Pcln table format: [...]funcTab + [...]_Func
func makePclntable(funcs []_Func, lastFuncSize uint32, pcdataOffs [][]uint32, funcdataOffs [][]uint32) (pclntab []byte) {
    // Allocate space for the pc->func table. This structure consists of a pc offset
	// and an offset to the func structure. After that, we have a single pc
	// value that marks the end of the last function in the binary.
    var size int64 = int64(len(funcs)*2*4 + 4)
    var startLocations = make([]uint32, len(funcs))
    for i := range funcs {
        size = rnd(size, int64(_PtrSize))
        //writePCToFunc
        startLocations[i] = uint32(size)
        size += int64(_FUNC_SIZE+len(funcdataOffs[i])*4+len(pcdataOffs[i])*4)
    }

    pclntab = make([]byte, 0, size)

    // write a map of pc->func info offsets 
    for i, f := range funcs {

        byteOrder.PutUint32(pclntab[len(pclntab):len(pclntab)+4], uint32(f.entryOff))
        byteOrder.PutUint32(pclntab[len(pclntab)+4:len(pclntab)+8], uint32(startLocations[i]))
        pclntab = pclntab[:len(pclntab)+8]
    }

    // Final entry of table is just end pc offset.
    lastFunc := funcs[len(funcs)-1]
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
            byteOrder.PutUint32(pclntab[off:off+4], uint32(funcdata))
            off += 4
        }

    }

    return
}

// findfunc table used to find the funcinfo by pc
// see findfunc() in runtime/symtab.go
func writeFindfunctab(out *[]byte, ftab []funcTab) {
    
    max := ftab[len(ftab)-1].entry
    min := ftab[0].entry
    nbuckets := (max - min + BUCKETSIZE - 1) / BUCKETSIZE
    n := (max - min + SUBBUCKETSIZE - 1) / SUBBUCKETSIZE

    tab := make([]findfuncbucket, 0, nbuckets)

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

func makeModuledata(name string, funcs []FuncInfo, text []byte) (mod *moduledata) {
    mod = new(moduledata)
    mod.modulename = name

    // TODO estimate accurate capacity
    out := make([]byte, len(text), len(text)+256) 
    copy(out, text)

    // make filename table
    cu := make([]string, 0, len(funcs))
    for _, f := range funcs {
        cu = append(cu, f.FileName)
    }
    cutab, filetab, cuOffs := makeFilenametab([]compilationUnit{{cu}})
    mod.cutab = cutab
    mod.filetab = filetab

    // make funcname table
    funcnametab, nameOffs := makeFuncnameTab(funcs)
    mod.funcnametab = funcnametab

    // make pcdata table
    // NOTICE: _func only use offset to index pcdata, thus no need mmap() pcdata 
    pctab, pcdataOffs, _funcs := makePctab(funcs, cuOffs, nameOffs)
    mod.pctab = pctab

    // write func data
    // NOTICE: _func use mod.gofunc+offset to directly point funcdata, thus need mmap() funcdata
    fstart := len(out)
    funcdataOffs := writeFuncdata(&out, funcs)

    // make pc->func (binary search) func table
    lastFuncsize := funcs[len(funcs)-1].TextSize
    ftab := makeFtab(_funcs, lastFuncsize)
    mod.ftab = ftab

    // write pc->func (modmap) findfunc table
    ffstart := len(out)
    writeFindfunctab(&out, ftab)

    // make pclnt table
    pclntab := makePclntable(_funcs, lastFuncsize, pcdataOffs, funcdataOffs)
    mod.pclntable = pclntab

    // MMAP() text and funcdata segements
    p := os.Getpagesize()
    size := (((len(text) - 1) / p) + 1) * p
    addr := mmap(size)
    mod.text = addr
    mod.etext = addr + uintptr(size)
    mod.minpc = addr
    mod.maxpc = addr + uintptr(len(text))
    mod.gofunc = addr + uintptr(fstart)
    mod.findfunctab = addr + uintptr(ffstart)

    // make pc header
    mod.pcHeader = &pcHeader {
        magic   : _Magic,
        minLC   : _MinLC,
        nfunc   : len(funcs),
        ptrSize : _PtrSize,
        textStart: mod.text,
        funcnameOffset: getOffsetOf(moduledata{}, "funcnametab"),
        cuOffset: getOffsetOf(moduledata{}, "cutab"),
        filetabOffset: getOffsetOf(moduledata{}, "filetab"),
        pctabOffset: getOffsetOf(moduledata{}, "pctab"),
        pclnOffset: getOffsetOf(moduledata{}, "pclntab"),
    }

    // sepecial case: gcdata and gcbss must by non-empty
    mod.gcdata = uintptr(unsafe.Pointer(&emptyByte))
    mod.gcbss = uintptr(unsafe.Pointer(&emptyByte))

    // copy the machine code
    s := rt.BytesFrom(unsafe.Pointer(addr), size, size)
    copy(s, out)

    // make it executable
    mprotect(addr, size)
    return
}

func Load(modulename string, funcs []FuncInfo, text []byte) (out []Function) {
    // generate module data and allocate memory address
    mod := makeModuledata(modulename, funcs, text)

    // verify and register the new module
    moduledataverify1(mod)
    registerModule(mod)

    // encapsulate function address
    out = make([]Function, len(funcs))
    for i, f := range funcs {
        out[i] = Function(mod.text + uintptr(f.EntryOff))
    }

    return 
}