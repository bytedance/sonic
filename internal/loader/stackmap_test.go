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
    `fmt`
    `testing`
    `unsafe`
)

type bitvector struct {
    n        int32 // # of bits
    bytedata *uint8
}

//go:nosplit
func addb(p *byte, n uintptr) *byte {
    return (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + n))
}

func (bv *bitvector) ptrbit(i uintptr) uint8 {
    b := *(addb(bv.bytedata, i/8))
    return (b >> (i % 8)) & 1
}

//go:linkname stackmapdata runtime.stackmapdata
func stackmapdata(_ *StackMap, _ int32) bitvector

func dumpstackmap(m *StackMap) {
    for i := int32(0); i < m.N; i++ {
        fmt.Printf("bitmap #%d/%d: ", i, m.L)
        bv := stackmapdata(m, i)
        for j := int32(0); j < bv.n; j++ {
            fmt.Printf("%d ", bv.ptrbit(uintptr(j)))
        }
        fmt.Printf("\n")
    }
}

var keepalive struct {
    s  string
    i  int
    vp unsafe.Pointer
    sb interface{}
    fv uint64
}

func stackmaptestfunc(s string, i int, vp unsafe.Pointer, sb interface{}, fv uint64) (x *uint64, y string, z *int) {
    z = new(int)
    x = new(uint64)
    y = s + "asdf"
    keepalive.s = s
    keepalive.i = i
    keepalive.vp = vp
    keepalive.sb = sb
    keepalive.fv = fv
    return
}

func TestStackMap_Dump(t *testing.T) {
    args, locals := stackMap(stackmaptestfunc)
    println("--- args ---")
    dumpstackmap((*StackMap)(unsafe.Pointer(args)))
    println("--- locals ---")
    dumpstackmap((*StackMap)(unsafe.Pointer(locals)))
}