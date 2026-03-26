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

package encoder

import (
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/bytedance/sonic/internal/encoder/x86"
	"github.com/bytedance/sonic/internal/jit"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/loader"
	"github.com/bytedance/sonic/option"
)

func ForceUseJit() {
	x86.SetCompiler(makeEncoderX86)
	pretouchType = pretouchTypeX86
	encodeTypedPointer = x86.EncodeTypedPointer
	vars.UseVM = false
}

func init() {
	if vars.UseVM {
		ForceUseVM()
	} else {
		ForceUseJit()
	}
}

var _KeepAlive struct {
	rb    *[]byte
	vp    unsafe.Pointer
	sb    *vars.Stack
	fv    uint64
	err   error
	frame [x86.FP_offs]byte
}

func makeEncoderX86(vt *rt.GoType, ex ...interface{}) (interface{}, error) {
	pp, err := NewCompiler().Compile(vt.Pack(), ex[0].(bool))
	if err != nil {
		return nil, err
	}
	as := x86.NewAssembler(pp)
	as.Name = vt.String()
	return as.Load(), nil
}

func pretouchTypeX86(_vt reflect.Type, opts option.CompileOptions, v uint8) (map[reflect.Type]uint8, error) {
	/* compile function */
	compiler := NewCompiler().apply(opts)
	encoder := func(vt *rt.GoType, ex ...interface{}) (interface{}, error) {
		pp, err := compiler.Compile(vt.Pack(), ex[0].(bool))
		if err != nil {
			return nil, err
		}
		as := x86.NewAssembler(pp)
		as.Name = vt.String()
		return as.Load(), nil
	}

	/* find or compile */
	vt := rt.UnpackType(_vt)
	if val := vars.GetProgram(vt); val != nil {
		return nil, nil
	} else if _, err := vars.ComputeProgram(vt, encoder, v == 1); err == nil {
		return compiler.rec, nil
	} else {
		return nil, err
	}
}

type x86PretouchProgram struct {
	vt   *rt.GoType
	pv   bool
	item loader.LoadOneItem
}

func pretouchRecX86(vtm map[reflect.Type]uint8, opts option.CompileOptions) error {
	pendings := make(map[*rt.GoType]x86PretouchProgram)

	for opts.RecursiveDepth >= 0 && len(vtm) > 0 {
		next := make(map[reflect.Type]uint8)
		for vt, v := range vtm {
			gvt := rt.UnpackType(vt)
			if vars.GetProgram(gvt) != nil {
				continue
			}
			if _, ok := pendings[gvt]; ok {
				continue
			}

			compiler := NewCompiler().apply(opts)
			pp, err := compiler.Compile(vt, v == 1)
			if err != nil {
				return err
			}

			as := x86.NewAssembler(pp)
			as.Name = vt.String()
			text, pcdata := as.Export()

			pendings[gvt] = x86PretouchProgram{
				vt: gvt,
				pv: v == 1,
				item: loader.LoadOneItem{
					Text:      text,
					FuncName:  "encode_" + as.Name,
					ArgSize:   x86.FP_args,
					ArgPtrs:   vars.ArgPtrs,
					LocalPtrs: vars.LocalPtrs,
					Pcdata:    pcdata,
				},
			}

			for svt, pv := range compiler.rec {
				next[svt] = pv
			}
		}

		opts.RecursiveDepth--
		vtm = next
	}

	if len(pendings) == 0 {
		return nil
	}

	entries := make([]x86PretouchProgram, 0, len(pendings))
	items := make([]loader.LoadOneItem, 0, len(pendings))
	for _, p := range pendings {
		entries = append(entries, p)
		items = append(items, p.item)
	}

	loaded := jit.LoadMany(items)
	for i, p := range entries {
		enc := x86.ToEncoder(loaded[i])
		_, err := vars.ComputeProgram(p.vt, func(*rt.GoType, ...interface{}) (interface{}, error) {
			return enc, nil
		}, p.pv)
		if err != nil {
			return err
		}
	}

	return nil
}
