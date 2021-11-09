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
	"errors"
	"runtime"
	"sync"
	"unsafe"

	"github.com/bytedance/sonic/internal/caching"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

const (
    _MinSlice = 16
    _MaxStack = 65536 // 64k slots
    _StackError = 8 + _MaxStack * 8 + 8 + types.MAX_RECURSE * 8 * 2
)

const (
    _PtrBytes  = _PTR_SIZE / 8
    _FsmOffset = (_MaxStack + 1) * _PtrBytes
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
    vp [types.MAX_RECURSE]*interface{}
    err error
}

type _Decoder func(
    s  string,
    i  int,
    vp unsafe.Pointer,
    sb *_Stack,
    fv uint64,
) (int, error)

var errCallShadow = errors.New("DON'T CALL THIS!")

//go:nosplit
// Faker func of _Decoder, used to export its stackmap as _Decoder's
func _Decoder_Shadow(rb *[]byte, vp unsafe.Pointer, sb *_Stack, fv uint64) error {
    // align to assembler_amd64.go: _FP_offs
    var stacks [_FP_offs]byte
    runtime.KeepAlive(stacks)

    // must keep rb, vp and sb noticeable to GC
    runtime.KeepAlive(sb)
    runtime.KeepAlive(rb)
    runtime.KeepAlive(vp)

    return errCallShadow
}

func newStack() *_Stack {
    if ret := stackPool.Get(); ret == nil {
        return new(_Stack)
    } else {
        return ret.(*_Stack)
    }
}

func freeStack(p *_Stack) {
    p.err = nil
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
    if pp, err := make(_Compiler).compile(vt.Pack()); err != nil {
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