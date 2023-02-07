//go:build go1.18
// +build go1.18

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
    `runtime`
    `runtime/debug`
    `strconv`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
    `github.com/stretchr/testify/require`
)

func TestLoad(t *testing.T) {
    // defer func() {
    //     if r := recover(); r != nil {
    //         runtime.GC()
    //         if r != "hook1" {
    //             t.Fatal("not right panic:" + r.(string))
    //         }
    //     } else {
    //         t.Fatal("not panic")
    //     }
    // }()

    var hstr string

    type TestFunc func(i *int, hook func(i *int)) int
    var hook = func(i *int) {
        runtime.GC()
        debug.FreeOSMemory()
        hstr = ("hook" + strconv.Itoa(*i))
        runtime.GC()
        debug.FreeOSMemory()
    }
    // var f TestFunc = func(i *int, hook func(i *int)) int {
    //     var t = *i
    //     hook(i)
    //     return t + *i
    // }
    bc := []byte {
        0x48, 0x83, 0xec, 0x18,         // (0x00) subq $24, %rsp
        0x48, 0x89, 0x6c, 0x24, 0x10,   // (0x04) movq %rbp, 16(%rsp)
        0x48, 0x8d, 0x6c, 0x24, 0x10,   // (0x09) leaq 16(%rsp), %rbp
        0x48, 0x89, 0x44, 0x24, 0x20,   // (0x0e) movq %rax, 32(%rsp)
        0x48, 0x8b, 0x08,               // (0x13) movq (%rax), %rcx
        0x48, 0x89, 0x4c, 0x24, 0x08,   // (0x16) movq %rcx, 8(%rsp)
        0x48, 0x8b, 0x33,               // (0x1b) movq (%rbx), %rsi
        0x48, 0x89, 0xda,               // (0x1e) movq %rbx, %rdx
        0xff, 0xd6,                     // (0x21) callq %rsi
        0x48, 0x8b, 0x44, 0x24, 0x08,   // (0x23) movq 8(%rsp), %rax
        0x48, 0x8b, 0x4c, 0x24, 0x20,   // (0x28) movq 32(%rsp), %rcx
        0x48, 0x03, 0x01,               // (0x2d) addq (%rcx), %rax
        0x48, 0x8b, 0x6c, 0x24, 0x10,   // (0x30) movq 16(%rsp), %rbp
        0x48, 0x83, 0xc4, 0x18,         // (0x35) addq $24, %rsp
        0xc3,                           // (0x39) ret
    }
    size := uint32(len(bc))
    fn := Func{
        ID: 0,
        Flag: 0,
        ArgsSize: 16,
        EntryOff: 0,
        TextSize: size,
        DeferReturn: 0,
        FileIndex: 0,
        Name: "dummy",
    }

    fn.Pcsp = &Pcdata{
        {PC: 0x04, Val: 0},
        {PC: size, Val: 24},
    }

    fn.Pcline = &Pcdata{
        {PC: 0x00, Val: 0},
        {PC: 0x0e, Val: 1},
        {PC: 0x1d, Val: 2},
        {PC: size, Val: 3},
    }

    fn.Pcfile = &Pcdata{
        {PC: size, Val: 0},
    }

    fn.PcUnsafePoint = &Pcdata{
        {PC: size, Val: PCDATA_UnsafePointUnsafe},
    }

    fn.PcStackMapIndex = &Pcdata{
        {PC: size, Val: 0},
    }

    args := rt.StackMapBuilder{}
    args.AddField(true)
    args.AddField(true)
    fn.ArgsPointerMaps = args.Build()

    locals := rt.StackMapBuilder{}
    locals.AddField(false)
    locals.AddField(false)
    fn.LocalsPointerMaps = locals.Build()

    rets := Load(bc, []Func{fn}, "dummy_module", []string{"github.com/bytedance/sonic/dummy.go"})
    println("func address ", *(*unsafe.Pointer)(rets[0]))
    // for k, _ := range moduleCache.m {
    //     spew.Dump(k)
    // }

    f := *(*TestFunc)(unsafe.Pointer(&rets[0]))
    i := 1
    j := f(&i, hook)
    require.Equal(t, 2, j)
    require.Equal(t, "hook1", hstr)
}