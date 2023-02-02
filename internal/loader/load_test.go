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
	"runtime"
	"testing"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)


func TestLoad(t *testing.T) {
    // defer func() {
    //     if r := recover(); r != nil {
    //         runtime.GC()
    //         if r != "hook" {
    //             t.Fatal("not right panic")
    //         }
    //     } else {
    //         t.Fatal("not panic")
    //     }
    // }()

    type TestFunc func(i *int, hook func()) int
    var hook = func() {
        runtime.GC()
        panic("hook")
    }
    // var f TestFunc = func(i *int, hook func()) int {
    //     x := *i
    //     hook()
    //     return x + 1234
    // }
    bc := []byte {
        0x48, 0x83, 0xec, 0x10,             //(0x00)      subq $16, %rsp
        0x48, 0x89, 0x6c, 0x24, 0x08,       //(0x04)      movq %rbp, 8(%rsp)
        0x48, 0x8d, 0x6c, 0x24, 0x08,       //(0x09)      leaq 8(%rsp), %rbp
        0x48, 0x8b, 0x00,                   //(0x0e)      movq (%rax), %rax
        0x48, 0x89, 0x04, 0x24,             //(0x11)      movq %rax, (%rsp)
        0x48, 0x8b, 0x0b,                   //(0x15)      movq (%rbx), %rcx
        0x48, 0x89, 0xda,                   //(0x18)      movq %rbx, %rdx
        0xff, 0xd1,                         //(0x1b)      callq %rcx
        0x48, 0x8b, 0x04, 0x24,             //(0x1d)      movq (%rsp), %rax
        0x48, 0x05, 0xd2, 0x04, 0x00, 0x00, //(0x21)      addq $1234, %rax
        0x48, 0x8b, 0x6c, 0x24, 0x08,       //(0x27)      movq 8(%rsp), %rbp
        0x48, 0x83, 0xc4, 0x10,             //(0x2c)      addq $16, %rsp
        0xc3,                               //(0x30)      ret
    }

    fn := Func{
        ID: 0,
        Flag: 0,
        EntryOff: 0,
        TextSize: uint32(len(bc)),
        StartLine: 0,
        DeferReturn: 0,
        FileIndex: 0,
        Name: "dummy",
    }

    fn.Pcsp = &Pcdata{
		{PC: 0x4, Val: 0},
		{PC: 0x30, Val: 16},
		{PC: 0x31, Val: 0},
	}

    fn.Pcline = &Pcdata{
        {PC: 0x00, Val: 0},
        {PC: 0x0e, Val: 1},
        {PC: 0x1d, Val: 2},
        {PC: 0x31, Val: 3},
    }

    fn.Pcfile = &Pcdata{
        {PC: 0x31, Val: 0},
    }

    fn.PcUnsafePoint = &Pcdata{
        {PC: 0x00, Val: -2},
        {PC: 0x31, Val: -1},
    }

    fn.PcStackMapIndex = &Pcdata{
        {PC: 0x0, Val: 0},
        {PC: 0x31, Val: 0},
    }

    args := StackMapBuilder{}
    args.AddField(false)
    args.AddField(false)
    fn.ArgsPointerMaps = args.Build()

    locals := StackMapBuilder{}
    locals.AddField(false)
    locals.AddField(false)
    fn.LocalsPointerMaps = locals.Build()

    rets := Load("dummy_module", []string{"github.com/bytedance/sonic/dummy.go"}, []Func{fn}, bc)
    spew.Dump(rets)
    // for k, _ := range moduleCache.m {
    //     spew.Dump(k)
    // }

    f := *(*TestFunc)(unsafe.Pointer(&rets[0]))
    i := 1
    j := f(&i, hook)
    require.Equal(t, 1235, j)
}
