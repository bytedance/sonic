//go:build linux || darwin
// +build linux darwin

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
    `encoding/binary`
    `fmt`
    `os`
    `reflect`
    `syscall`
    `unsafe`
)

const (
    _AP = syscall.MAP_ANON  | syscall.MAP_PRIVATE
    _RX = syscall.PROT_READ | syscall.PROT_EXEC
    _RW = syscall.PROT_READ | syscall.PROT_WRITE
)

const (
    _MinLC uint8 = 1
    _PtrSize uint8 = 8
)

var (
    byteOrder binary.ByteOrder = binary.LittleEndian
)

type Func struct {
    ID          uint8  // see runtime/symtab.go
    Flag        uint8  // see runtime/symtab.go
    ArgsSize    int32  // args byte size
    EntryOff    uint32 // start pc, offset to moduledata.text
    TextSize    uint32 // size of func text
    StartLine   int32  // line number of first line in file of function
    DeferReturn uint32 // offset of start of a deferreturn call instruction from entry, if any.
    FileIndex   uint32 // index into filetab 
    Name        string // name of function

    // PC data
    Pcsp            *Pcdata
    Pcfile          *Pcdata
    Pcline          *Pcdata
    PcUnsafePoint   *Pcdata
    PcStackMapIndex *Pcdata
    PcInlTreeIndex  *Pcdata
    PcArgLiveIndex  *Pcdata
    
    // Func data
    ArgsPointerMaps    interface{}
    LocalsPointerMaps  interface{}
    StackObjects       interface{}
    InlTree            interface{}
    OpenCodedDeferInfo interface{}
    ArgInfo            interface{}
    ArgLiveInfo        interface{}
    WrapInfo           interface{}
}

func getOffsetOf(data interface{}, field string) uintptr {
    t := reflect.TypeOf(data)
    fv, ok := t.FieldByName(field)
    if !ok {
        panic(fmt.Sprintf("field %s not found in struct %s", field, t.Name()))
    }
    return fv.Offset
}


type Loader []byte

type Function unsafe.Pointer

func (self Loader) LoadWithFaker(fn string, fp int, args int, faker interface{}) (f Function) {
    p := os.Getpagesize()
    n := int(rnd(int64(len(self)), int64(p)))

    /* register the function */
    m := mmap(n)
    v := fmt.Sprintf("runtime.__%s_%x", fn, m)
    argsptr, localsptr := GetStackMap(faker)
    registerFunction(v, m, uintptr(n), fp, args, uintptr(len(self)), argsptr, localsptr)

    /* reference as a slice */
    s := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader {
        Data : m,
        Cap  : n,
        Len  : len(self),
    }))

    /* copy the machine code, and make it executable */
    copy(s, self)
    mprotect(m, n)
    return Function(&m)
}

func (self Loader) Load(fn string, fp int, args int) (f Function) {
    return self.LoadWithFaker(fn, fp, args, func(){})
}

func mmap(nb int) uintptr {
    if m, _, e := syscall.RawSyscall6(syscall.SYS_MMAP, 0, uintptr(nb), _RW, _AP, 0, 0); e != 0 {
        panic(e)
    } else {
        return m
    }
}

func mprotect(p uintptr, nb int) {
    if _, _, err := syscall.RawSyscall(syscall.SYS_MPROTECT, p, uintptr(nb), _RX); err != 0 {
        panic(err)
    }
}
