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

package caching

import (
    `fmt`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

const (
    _H_basis uint64 = 0xcbf29ce484222325
    _H_prime uint64 = 0x00000100000001b3
)

func fnv1a(s string) uint64 {
    v := _H_basis
    m := (*rt.GoString)(unsafe.Pointer(&s))

    /* hash each byte */
    for i := 0; i < m.Len; i++ {
        v ^= uint64(*(*uint8)(unsafe.Pointer(uintptr(m.Ptr) + uintptr(i))))
        v *= _H_prime
    }

    /* never returns 0 for hash */
    if v == 0 {
        return 1
    } else {
        return v
    }
}

func TestHashing_Fnv1a(t *testing.T) {
    fmt.Printf("%#x\n", fnv1a("hello, world"))
}

func TestHashing_StrHash(t *testing.T) {
    s := "hello, world"
    fmt.Printf("%#x\n", StrHash(s))
}

var fn_fnv1a = fnv1a
func BenchmarkHashing_Fnv1a(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fn_fnv1a("accountid_interval_aweme_second")
    }
}

func BenchmarkHashing_StrHash(b *testing.B) {
    for i := 0; i < b.N; i++ {
        StrHash("accountid_interval_aweme_second")
    }
}
