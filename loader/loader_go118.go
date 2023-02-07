//go:build go1.18 && !go1.21
// +build go1.18,!go1.21

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
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

type Options struct {
    NoPreempt bool
}

type Function unsafe.Pointer

type ModuleLoader struct {
    Name string
    File string
    Options
}

func (self ModuleLoader) LoadFunc(text []byte, funcName string, frameSize int, argSize int, argStackmap []bool, localStackmap []bool) Function {
    size := uint32(len(text))

    fn := Func{
        Name: funcName,
        TextSize: size,
        ArgsSize: int32(argSize),
    }

    fn.Pcsp = &Pcdata{
		{PC: size, Val: int32(frameSize)},
	}

    if self.NoPreempt {
        fn.PcUnsafePoint = &Pcdata{
            {PC: size, Val: PCDATA_UnsafePointUnsafe},
        }
    } else {
        fn.PcUnsafePoint = &Pcdata{
            {PC: size, Val: PCDATA_UnsafePointSafe},
        }
    }

    fn.PcStackMapIndex = &Pcdata{
        {PC: size, Val: 0},
    }

    if argStackmap != nil {
        args := rt.StackMapBuilder{}
        for _, b := range argStackmap {
            args.AddField(b)
        }
        fn.ArgsPointerMaps = args.Build()
    }
    
    if localStackmap != nil {
        locals := rt .StackMapBuilder{}
        for _, b := range localStackmap {
            locals.AddField(b)
        }
        fn.LocalsPointerMaps = locals.Build()
    }

    out := Load(self.Name + funcName, []string{self.File}, []Func{fn}, text)
    return out[0]
}

func Load(modulename string, filenames []string, funcs []Func, text []byte) (out []Function) {
    // generate module data and allocate memory address
    mod := makeModuledata(modulename, filenames, funcs, text)

    // verify and register the new module
    moduledataverify1(mod)
    registerModule(mod)

    // encapsulate function address
    out = make([]Function, len(funcs))
    for i, f := range funcs {
        m := uintptr(mod.text + uintptr(f.EntryOff))
        out[i] = Function(&m)
    }
    return 
}