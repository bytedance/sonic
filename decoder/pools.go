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

package decoder

import (
    `errors`
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/caching`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _MinSlice = 16
    _MaxStack = 65536 // 64k slots
    _MaxStackBytes = _MaxStack * _PtrBytes
    _MaxDigitNums = 800  // used in atof fallback algorithm
)

const (
    _PtrBytes   = _PTR_SIZE / 8
    _FsmOffset  = (_MaxStack + 1) * _PtrBytes
    _DbufOffset = _FsmOffset + int64(unsafe.Sizeof(types.StateMachine{})) + types.MAX_RECURSE * _PtrBytes
    _StackSize  = unsafe.Sizeof(_Stack{})
)

var (
    stackPool     = sync.Pool{}
    valueCache    = []unsafe.Pointer(nil)
    fieldCache    = []*caching.FieldMap(nil)
    fieldCacheMux = sync.Mutex{}
    programCache  = caching.CreateProgramCache()
)

type _Stack struct {
    sp uintptr
    sb [_MaxStack]unsafe.Pointer
    mm types.StateMachine
    vp [types.MAX_RECURSE]unsafe.Pointer
    dp [_MaxDigitNums]byte
}

type _Decoder func(
    s  string,
    i  int,
    vp unsafe.Pointer,
    sb *_Stack,
    fv uint64,
    sv string, // DO NOT pass value to this arguement, since it is only used for local _VAR_sv
    vk unsafe.Pointer, // DO NOT pass value to this arguement, since it is only used for local _VAR_vk
) (int, error)

var _KeepAlive struct {
    s string
    i int
    vp unsafe.Pointer
    sb *_Stack
    fv uint64
    sv string
    vk unsafe.Pointer

    ret int
    err error

    frame_decoder [_FP_offs]byte
    frame_generic [_VD_offs]byte
}

var errCallShadow = errors.New("DON'T CALL THIS!")

// Faker func of _Decoder, used to export its stackmap as _Decoder's
func _Decoder_Shadow(s string, i int, vp unsafe.Pointer, sb *_Stack, fv uint64, sv string, vk unsafe.Pointer) (ret int, err error) {
    // align to assembler_amd64.go: _FP_offs
    var frame [_FP_offs]byte

    // keep all args and stacks alive
    _KeepAlive.s = s
    _KeepAlive.i = i
    _KeepAlive.vp = vp
    _KeepAlive.sb = sb
    _KeepAlive.fv = fv
    _KeepAlive.ret = ret
    _KeepAlive.err = err
    _KeepAlive.sv = sv
    _KeepAlive.vk = vk
    _KeepAlive.frame_decoder = frame
    
    return 0, errCallShadow
}

// Faker func of _Decoder_Generic, used to export its stackmap
func _Decoder_Generic_Shadow(sb *_Stack) {
    // align to generic_amd64.go: _VD_offs
    var frame [_VD_offs]byte

    // must keep sb noticeable to GC
    _KeepAlive.sb = sb
    _KeepAlive.frame_generic = frame
}

func newStack() *_Stack {
    if ret := stackPool.Get(); ret == nil {
        return new(_Stack)
    } else {
        return ret.(*_Stack)
    }
}

func resetStack(p *_Stack) {
    memclrNoHeapPointers(unsafe.Pointer(p), _StackSize)
}

func freeStack(p *_Stack) {
    p.sp = 0
    stackPool.Put(p)
}

func freezeValue(v unsafe.Pointer) uintptr {
    valueCache = append(valueCache, v)
    return uintptr(v)
}

func freezeFields(v *caching.FieldMap) int64 {
    fieldCacheMux.Lock()
    fieldCache = append(fieldCache, v)
    fieldCacheMux.Unlock()
    return referenceFields(v)
}

func referenceFields(v *caching.FieldMap) int64 {
    return int64(uintptr(unsafe.Pointer(v)))
}

func makeDecoder(vt *rt.GoType) (interface{}, error) {
    if pp, err := newCompiler().compile(vt.Pack()); err != nil {
        return nil, err
    } else {
        return newAssembler(pp).Load(), nil
    }
}

func findOrCompile(vt *rt.GoType) (_Decoder, error) {
    if val := programCache.Get(vt); val != nil {
        return val.(_Decoder), nil
    } else if ret, err := programCache.Compute(vt, makeDecoder); err == nil {
        return ret.(_Decoder), nil
    } else {
        return nil, err
    }
}