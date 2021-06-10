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
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/caching`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _MinSlice  = 16
    _MaxStack  = 262144  // 256k slots
    _FsmOffset = int64(unsafe.Offsetof(_Stack{}.mm))
)

var (
    stackPool    = sync.Pool{}
    fieldCache   = []*caching.FieldMap(nil)
    programCache = caching.CreateProgramCache()
)

type _Stack struct {
    sp uintptr
    sb [_MaxStack]unsafe.Pointer
    mm types.StateMachine
}

type _Decoder func(
    s  string,
    i  int,
    vp unsafe.Pointer,
    sb *_Stack,
    fv uint64,
) (int, error)

func newStack() *_Stack {
    if ret := stackPool.Get(); ret == nil {
        return new(_Stack)
    } else {
        return ret.(*_Stack)
    }
}

func freeStack(p *_Stack) {
    stackPool.Put(p)
}

func freezeFields(v *caching.FieldMap) int64 {
    fieldCache = append(fieldCache, v)
    return referenceFields(v)
}

func referenceFields(v *caching.FieldMap) int64 {
    return int64(uintptr(unsafe.Pointer(v)))
}

func findOrCompile(vt *rt.GoType) (_Decoder, error) {
    var ex error
    var fn _Decoder
    var pp *_Program
    var fv interface{}

    /* fast path: the program is in the cache */
    if fv = programCache.Get(vt); fv != nil {
        return fv.(_Decoder), nil
    }

    /* slow path: not found, compile the type on the fly */
    if pp, ex = newCompiler().compile(vt.Pack()); ex != nil {
        return nil, ex
    }

    /* link the program, and put it into cache */
    fn = newAssembler(pp).Load()
    programCache.Put(vt, fn)
    return fn, nil
}
