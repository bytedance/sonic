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
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/caching`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _MaxStack  = 65536      // 64k states
    _MaxBuffer = 1048576    // 1MB buffer size
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
    sb [_MaxStack]_State
}

type _Encoder func(
    rb *[]byte,
    vp unsafe.Pointer,
    sb *_Stack,
    fv uint64,
) error

func newBytes() []byte {
    if ret := bytesPool.Get(); ret != nil {
        return ret.([]byte)
    } else {
        return make([]byte, 0, _MaxBuffer)
    }
}

func newStack() *_Stack {
    if ret := stackPool.Get(); ret == nil {
        return new(_Stack)
    } else {
        return ret.(*_Stack)
    }
}

func newBuffer() *bytes.Buffer {
    if ret := bufferPool.Get(); ret != nil {
        return ret.(*bytes.Buffer)
    } else {
        return bytes.NewBuffer(make([]byte, 0, _MaxBuffer))
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

func findOrCompile(vt *rt.GoType) (_Encoder, error) {
    var ex error
    var fn _Encoder
    var pp *_Program
    var fv interface{}

    /* fast path: the program is in the cache */
    if fv = programCache.Get(vt); fv != nil {
        return fv.(_Encoder), nil
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
