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
    `bytes`
    `errors`
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/caching`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/bytedance/sonic/option`
)

const (
    _StackSize = unsafe.Sizeof(_Stack{})
    _StateSize  = int64(unsafe.Sizeof(_State{}))
)

var (
    bytesPool    = sync.Pool{}
    stackPool    = sync.Pool{}
    bufferPool   = sync.Pool{}
    programCache = caching.CreateProgramCache()
)

type _State struct {
    x int
    f uint64
    p unsafe.Pointer
    q unsafe.Pointer
}

type _Stack struct {
    sp uint64
    sb []_State
}

type _Encoder func(
    rb *[]byte,
    vp unsafe.Pointer,
    sb *_Stack,
    fv uint64,
) error

var _KeepAlive struct {
    rb *[]byte
    vp unsafe.Pointer
    sb *_Stack
    fv uint64
    err error
    frame [_FP_offs]byte
}

var errCallShadow = errors.New("DON'T CALL THIS!")

// Faker func of _Encoder, used to export its stackmap as _Encoder's
func _Encoder_Shadow(rb *[]byte, vp unsafe.Pointer, sb *_Stack, fv uint64) (err error) {
    // align to assembler_amd64.go: _FP_offs
    var frame [_FP_offs]byte

    // must keep all args and frames noticeable to GC
    _KeepAlive.rb = rb
    _KeepAlive.vp = vp
    _KeepAlive.sb = sb
    _KeepAlive.fv = fv
    _KeepAlive.err = err
    _KeepAlive.frame = frame

    return errCallShadow
}

func newBytes() []byte {
    if ret := bytesPool.Get(); ret != nil {
        return ret.([]byte)
    } else {
        return make([]byte, 0, option.DefaultEncodeBufferSize)
    }
}

func newStack() *_Stack {
    if ret := stackPool.Get(); ret == nil {
        st := new(_Stack)
        st.sb = make([]_State, 0, option.MaxEncodeStackSize)
        return st
    } else {
        return ret.(*_Stack)
    }
}

func resetStack(p *_Stack) {
    memclrNoHeapPointers(*(*unsafe.Pointer)(unsafe.Pointer(&p.sb)), uintptr(option.MaxEncodeStackSize * uint(_StateSize)))
}

func newBuffer() *bytes.Buffer {
    if ret := bufferPool.Get(); ret != nil {
        return ret.(*bytes.Buffer)
    } else {
        return bytes.NewBuffer(make([]byte, 0, option.DefaultEncodeBufferSize))
    }
}

func freeBytes(p []byte) {
    p = p[:0]
    bytesPool.Put(p)
}

func freeStack(p *_Stack) {
    p.sp = 0
    stackPool.Put(p)
}

func freeBuffer(p *bytes.Buffer) {
    p.Reset()
    bufferPool.Put(p)
}

func makeEncoder(vt *rt.GoType) (interface{}, error) {
    if pp, err := newCompiler().compile(vt.Pack()); err != nil {
        return nil, err
    } else {
        return newAssembler(pp).Load(), nil
    }
}

func findOrCompile(vt *rt.GoType) (_Encoder, error) {
    if val := programCache.Get(vt); val != nil {
        return val.(_Encoder), nil
    } else if ret, err := programCache.Compute(vt, makeEncoder); err == nil {
        return ret.(_Encoder), nil
    } else {
        return nil, err
    }
}
