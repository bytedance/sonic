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
	`fmt`
	`os`
	`syscall`
	`unsafe`

	`github.com/bytedance/sonic/internal/rt`
)

const (
    _AP = syscall.MAP_ANON  | syscall.MAP_PRIVATE
    _RX = syscall.PROT_READ | syscall.PROT_EXEC
    _RW = syscall.PROT_READ | syscall.PROT_WRITE
)

type Loader   []byte
type Function unsafe.Pointer

// func (self Loader) LoadWithFaker(fn string, fp int, args int, faker interface{}) (f Function) {
//     p := os.Getpagesize()
//     n := (((len(self) - 1) / p) + 1) * p

//     /* register the function */
//     m := mmap(n)
//     v := fmt.Sprintf("runtime.__%s_%x", fn, m)
//     argsptr, localsptr := stackMap(faker)
//     registerFunction(v, m, fp, args, uintptr(len(self)), argsptr, localsptr)

//     /* reference as a slice */
//     s := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader {
//         Data : m,
//         Cap  : n,
//         Len  : len(self),
//     }))

//     /* copy the machine code, and make it executable */
//     copy(s, self)
//     mprotect(m, n)
//     return Function(&m)
// }

func mkptr(m uintptr) unsafe.Pointer {
    return *(*unsafe.Pointer)(unsafe.Pointer(&m))
}

func alignUp(n uintptr, a int) uintptr {
    return (n + uintptr(a) - 1) &^ (uintptr(a) - 1)
}

func (self Loader) LoadWithFaker(fn string, frame rt.Frame, faker interface{}) Function {
    frame.ArgPtrs, frame.LocalPtrs = stackMap(faker)
    return self.Load(fn, frame)
}

func (self Loader) Load(fn string, frame rt.Frame) (f Function) {
    var mm uintptr
    var er syscall.Errno

    /* align the size to pages */
    nf := uintptr(len(self))
    nb := alignUp(nf, os.Getpagesize())

    /* allocate a block of memory */
    if mm, _, er = syscall.Syscall6(syscall.SYS_MMAP, 0, nb, _RW, _AP, 0, 0); er != 0 {
        panic(er)
    }

    /* copy code into the memory, and register the function */
    copy(*(*[]byte)(unsafe.Pointer(&rt.GoSlice{mkptr(mm), len(self), int(nb)})), self)
    registerFunction(fmt.Sprintf("runtime.%s_%x", fn, mm), mm, nf, frame)

    /* make it executable */
    if _, _, err := syscall.Syscall(syscall.SYS_MPROTECT, mm, nb, _RX); err != 0 {
        panic(err)
    } else {
        return Function(&mm)
    }
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
