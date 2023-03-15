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
	"unsafe"

	"github.com/bytedance/sonic/internal/cpu"
	"github.com/bytedance/sonic/internal/native/avx"
	"github.com/bytedance/sonic/internal/native/avx2"
	"github.com/bytedance/sonic/internal/native/sse"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/loader"
)

const (
    MaxFrameSize   uintptr = 400
    BufPaddingSize int     = 64
)

var (
    S_f64toa uintptr
    S_f32toa uintptr
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
    S_skip_one_fast    uintptr
    S_get_by_path    uintptr
    S_skip_array  uintptr
    S_skip_object uintptr
    S_skip_number uintptr
)

var (
    Quote func(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int

    Unquote func(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) int

    HTMLEscape func(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int

    Value func(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) int

    SkipOne func(s *string, p *int, m *types.StateMachine, flags uint64) int

    SkipOneFast func(s *string, p *int) int

    GetByPath func(s *string, p *int, path *[]interface{}) int

    ValidateOne func(s *string, p *int, m *types.StateMachine) int

    I64toa func(out *byte, val int64) (ret int)

    U64toa func(out *byte, val uint64) (ret int)

    F64toa func(out *byte, val float64) (ret int)

    ValidateUTF8 func(s *string, p *int, m *types.StateMachine) (ret int)

    ValidateUTF8Fast func(s *string) (ret int)
)

var Stubs = []loader.GoFunc{
    {"_f64toa", &F64toa},
    {"_get_by_path", &GetByPath},
    {"_html_escape", &HTMLEscape},
    {"_i64toa", &I64toa},
    {"_quote", &Quote},
    {"_skip_one", &SkipOne},
    {"_skip_one_fast", &SkipOneFast},
    {"_u64toa", &U64toa},
    {"_unquote", &Unquote},
    {"_validate_one", &ValidateOne},
    {"_validate_utf8", &ValidateUTF8},
    {"_validate_utf8_fast", &ValidateUTF8Fast},
    {"_value", &Value},
}


func useAVX() {
    S_f64toa      = avx.S_f64toa
    S_f32toa      = avx.S_f32toa
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
    S_skip_one_fast = avx.S_skip_one_fast
    S_skip_array  = avx.S_skip_array
    S_skip_object = avx.S_skip_object
    S_skip_number = avx.S_skip_number
    S_get_by_path = avx.S_get_by_path

    loader.WrapC(avx.Text___native_entry__, avx.Funcs, Stubs, "avx", "avx/native.c")
}

func useAVX2() {
    S_f64toa      = avx2.S_f64toa
    S_f32toa      = avx2.S_f32toa
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
    S_skip_one_fast = avx2.S_skip_one_fast
    S_skip_array  = avx2.S_skip_array
    S_skip_object = avx2.S_skip_object
    S_skip_number = avx2.S_skip_number
    S_get_by_path = avx2.S_get_by_path

    loader.WrapC(avx2.Text___native_entry__, avx2.Funcs, Stubs, "avx2", "avx2/native.c")
}

func useSSE() {
    S_f64toa = sse.S_f64toa
    S_f32toa = sse.S_f32toa
    S_i64toa = sse.S_i64toa
    S_u64toa = sse.S_u64toa
    S_lspace = sse.S_lspace
    S_quote = sse.S_quote
    S_unquote = sse.S_unquote
    S_value = sse.S_value
    S_vstring = sse.S_vstring
    S_vnumber = sse.S_vnumber
    S_vsigned = sse.S_vsigned
    S_vunsigned = sse.S_vunsigned
    S_skip_one = sse.S_skip_one
    S_skip_one_fast = sse.S_skip_one_fast
    S_skip_array = sse.S_skip_array
    S_skip_object = sse.S_skip_object
    S_skip_number = sse.S_skip_number
    S_get_by_path = sse.S_get_by_path

    loader.WrapC(sse.Text___native_entry__, sse.Funcs, Stubs, "sse", "sse/native.c")
}

func init() {
    if cpu.HasAVX2 {
        useAVX2()
    } else if cpu.HasAVX {
        useAVX()
    } else if cpu.HasSSE {
        useSSE()
    } else {
        panic("Unsupported CPU, maybe it's too old to run Sonic.")
    }
}
