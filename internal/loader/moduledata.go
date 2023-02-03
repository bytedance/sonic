/**
 * Copyright 2023 ByteDance Inc.
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
	"os"
	"sync"
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

// moduledata used to cache the funcdata and findfuncbucket of one module
var moduleCache = struct {
    m map[*moduledata][]byte
    sync.Mutex
}{
    m: make(map[*moduledata][]byte),
}

func makeModuledata(name string, filenames []string, funcs []Func, text []byte) (mod *moduledata) {
    mod = new(moduledata)
    mod.modulename = name

    // make filename table
    cu := make([]string, 0, len(filenames))
    for _, f := range filenames {
        cu = append(cu, f)
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
    // NOTICE: _func use mod.gofunc+offset to directly point funcdata, thus need cache funcdata
    // TODO: estimate accurate capacity
    cache := make([]byte, 0, len(funcs)*int(_PtrSize)) 
    fstart, funcdataOffs := writeFuncdata(&cache, funcs)

    // make pc->func (binary search) func table
    lastFuncsize := funcs[len(funcs)-1].TextSize
    ftab := makeFtab(_funcs, lastFuncsize)
    mod.ftab = ftab

    // write pc->func (modmap) findfunc table
    ffstart := writeFindfunctab(&cache, ftab)

    // make pclnt table
    pclntab := makePclntable(_funcs, lastFuncsize, pcdataOffs, funcdataOffs)
    mod.pclntable = pclntab

    // mmap() text and funcdata segements
    p := os.Getpagesize()
    size := int(rnd(int64(len(text)), int64(p)))
    addr := mmap(size)
    // copy the machine code
    s := rt.BytesFrom(unsafe.Pointer(addr), len(text), size)
    copy(s, text)
    // make it executable
    mprotect(addr, size)

    // assign addresses
    mod.text = addr
    mod.etext = addr + uintptr(size)
    mod.minpc = addr
    mod.maxpc = addr + uintptr(len(text))

    // cache funcdata and findfuncbucket
    moduleCache.Lock()
    moduleCache.m[mod] = cache
    moduleCache.Unlock()
    mod.gofunc = uintptr(unsafe.Pointer(&cache[fstart]))
    mod.findfunctab = uintptr(unsafe.Pointer(&cache[ffstart]))

    // make pc header
    mod.pcHeader = &pcHeader {
        magic   : _Magic,
        minLC   : _MinLC,
        ptrSize : _PtrSize,
        nfunc   : len(funcs),
        nfiles: uint(len(cu)),
        textStart: mod.text,
        funcnameOffset: getOffsetOf(moduledata{}, "funcnametab"),
        cuOffset: getOffsetOf(moduledata{}, "cutab"),
        filetabOffset: getOffsetOf(moduledata{}, "filetab"),
        pctabOffset: getOffsetOf(moduledata{}, "pctab"),
        pclnOffset: getOffsetOf(moduledata{}, "pclntable"),
    }

    // sepecial case: gcdata and gcbss must by non-empty
    mod.gcdata = uintptr(unsafe.Pointer(&emptyByte))
    mod.gcbss = uintptr(unsafe.Pointer(&emptyByte))

    return
}
