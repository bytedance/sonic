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
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/internal/abi"
	"github.com/bytedance/sonic/internal/rt"
)

type CFunc struct {
	Name     string
	EntryOff uint32
	TextSize uint32
	MaxStack uintptr
	Pcsp     [][2]uint32
}

type GoFunc struct {
	Name     string
	GoFunc   interface{}
}

func WrapC(text []byte, natives []CFunc, stubs []GoFunc, modulename string, filename string) {
	funcs := make([]Func, len(natives))
	
	// register C funcs
	for i, f := range natives {
		fn := Func{
			EntryOff: f.EntryOff,
			TextSize: f.TextSize,
			Name: f.Name,
		}
		if len(f.Pcsp) != 0 {
			fn.Pcsp = (*Pcdata)(unsafe.Pointer(&natives[i].Pcsp))
		}
		fn.PcUnsafePoint = &Pcdata{
			{PC: f.TextSize, Val: PCDATA_UnsafePointUnsafe},
		}
		funcs[i] = fn
	}
	rets := Load(text, funcs, modulename, []string{filename})

	// got absolute entry address
	native_entry := **(**uintptr)(unsafe.Pointer(&rets[0]))

	wraps := make([]Func, len(stubs))
	code := make([]byte, 0, len(wraps)*256)
	entryOff := uint32(0)

	// register wrapper go funcs
	for i, w := range stubs {
		for _, f := range natives {
			if w.Name != f.Name {
				continue
			}

			// assemble wrapper codes
			layout := abi.NewFunctionLayout(reflect.TypeOf(w.GoFunc).Elem())
			frame := abi.NewFrame(&layout, []bool{false, false}) 
			tcode := abi.CallC(uintptr(native_entry + uintptr(f.EntryOff)), frame, f.MaxStack)
			code = append(code, tcode...)

			size := uint32(len(tcode))
		
			fn := Func{
				ArgsSize: int32(layout.ArgSize()),
				EntryOff: entryOff,
				TextSize: size,
				Name: reflect.ValueOf(w.GoFunc).String(),
			}
			// NOTICE: need add check-stack and grow-stack pcsp
			fn.Pcsp = &Pcdata{
				{PC: uint32(frame.StackCheckTextSize()), Val: 0},
				{PC: size - uint32(frame.GrowStackTextSize()), Val: int32(frame.Size())},
				{PC: size, Val: 0},
			}
			fn.PcUnsafePoint = &Pcdata{
				{PC: size, Val: PCDATA_UnsafePointUnsafe},
			}
			fn.PcStackMapIndex = &Pcdata{
				{PC: size, Val: 0},
			}

			fn.ArgsPointerMaps = frame.ArgPtrs()
			fn.LocalsPointerMaps = frame.LocalPtrs()

			entryOff += size
			wraps[i] = fn
		}
	}
    gofuncs := Load(code, wraps, modulename+"/go", []string{filename+".go"})
	
	// set go func value 
	for i, f := range gofuncs {
		println(stubs[i].Name, " pc:", **(**unsafe.Pointer)(unsafe.Pointer(&f)))
		println(stubs[i].GoFunc, " type:", reflect.ValueOf(stubs[i].GoFunc).String())
        w := rt.UnpackEface(stubs[i].GoFunc)
		*(*Function)(w.Value) = f
	}
}