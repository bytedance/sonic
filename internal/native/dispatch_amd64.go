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

package native

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/cpu`
    `github.com/bytedance/sonic/internal/native/avx`
    `github.com/bytedance/sonic/internal/native/avx2`
    `github.com/bytedance/sonic/internal/native/sse4`
    `github.com/bytedance/sonic/internal/native/types`
)

const MaxFrameSize uintptr = 400

var (
    S_f64toa uintptr
    S_i64toa uintptr
    S_u64toa uintptr
    S_lspace uintptr
)

var (
    S_quote   uintptr
    S_unquote uintptr
)

var (
    S_value     uintptr
    S_vstring   uintptr
    S_vnumber   uintptr
    S_vsigned   uintptr
    S_vunsigned uintptr
)

var (
    S_skip_one    uintptr
    S_skip_array  uintptr
    S_skip_object uintptr
    S_skip_number uintptr
)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func Quote(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func Unquote(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func HTMLEscape(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func Value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func SkipOne(s *string, p *int, m *types.StateMachine, flags uint64) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func ValidateOne(s *string, p *int, m *types.StateMachine) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func I64toa(out *byte, val int64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func U64toa(out *byte, val uint64) (ret int)

func useAVX() {
    S_f64toa      = avx.S_f64toa
    S_i64toa      = avx.S_i64toa
    S_u64toa      = avx.S_u64toa
    S_lspace      = avx.S_lspace
    S_quote       = avx.S_quote
    S_unquote     = avx.S_unquote
    S_value       = avx.S_value
    S_vstring     = avx.S_vstring
    S_vnumber     = avx.S_vnumber
    S_vsigned     = avx.S_vsigned
    S_vunsigned   = avx.S_vunsigned
    S_skip_one    = avx.S_skip_one
    S_skip_array  = avx.S_skip_array
    S_skip_object = avx.S_skip_object
    S_skip_number = avx.S_skip_number
}

func useAVX2() {
    S_f64toa      = avx2.S_f64toa
    S_i64toa      = avx2.S_i64toa
    S_u64toa      = avx2.S_u64toa
    S_lspace      = avx2.S_lspace
    S_quote       = avx2.S_quote
    S_unquote     = avx2.S_unquote
    S_value       = avx2.S_value
    S_vstring     = avx2.S_vstring
    S_vnumber     = avx2.S_vnumber
    S_vsigned     = avx2.S_vsigned
    S_vunsigned   = avx2.S_vunsigned
    S_skip_one    = avx2.S_skip_one
    S_skip_array  = avx2.S_skip_array
    S_skip_object = avx2.S_skip_object
    S_skip_number = avx2.S_skip_number
}

func useSSE4() {
    S_f64toa = sse4.S_f64toa
    S_i64toa = sse4.S_i64toa
    S_u64toa = sse4.S_u64toa
    S_lspace = sse4.S_lspace
    S_quote = sse4.S_quote
    S_unquote = sse4.S_unquote
    S_value = sse4.S_value
    S_vstring = sse4.S_vstring
    S_vnumber = sse4.S_vnumber
    S_vsigned = sse4.S_vsigned
    S_vunsigned = sse4.S_vunsigned
    S_skip_one = sse4.S_skip_one
    S_skip_array = sse4.S_skip_array
    S_skip_object = sse4.S_skip_object
    S_skip_number = sse4.S_skip_number
}

func init() {
    if cpu.HasAVX2 {
        useAVX2()
    } else if cpu.HasAVX {
        useAVX()
    } else if cpu.HasSSE4 {
        useSSE4()
    } else {
        panic("Unsupported CPU, maybe it's too old to run Sonic.")
    }
}
