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
    S_quote       uintptr
    S_unquote     uintptr
    S_html_escape uintptr
)

var (
    S_value     uintptr
    S_vstring   uintptr
    S_vnumber   uintptr
    S_vsigned   uintptr
    S_vunsigned uintptr
)

var (
    S_skip_one      uintptr
    S_skip_one_fast uintptr
    S_get_by_path   uintptr
    S_skip_array    uintptr
    S_skip_object   uintptr
    S_skip_number   uintptr
)

var (
    S_validate_one       uintptr
    S_validate_utf8      uintptr
    S_validate_utf8_fast uintptr
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

var Stubs = []loader.GoC{
    {"_f64toa", &S_f64toa, &F64toa},
    {"_f32toa", &S_f32toa, nil},
    {"_i64toa", &S_i64toa, &I64toa},
    {"_u64toa", &S_u64toa, &U64toa},
    {"_lspace", &S_lspace, nil},
    {"_quote", &S_quote, &Quote},
    {"_unquote", &S_unquote, &Unquote},
    {"_html_escape", &S_html_escape, &HTMLEscape},
    {"_value", &S_value, &Value},
    {"_vstring", &S_vstring, nil},
    {"_vnumber", &S_vnumber, nil},
    {"_vsigned", &S_vsigned, nil},
    {"_vunsigned", &S_vunsigned, nil},
    {"_skip_one", &S_skip_one, &SkipOne},
    {"_skip_one_fast", &S_skip_one_fast, &SkipOneFast},
    {"_get_by_path", &S_get_by_path, &GetByPath},
    {"_skip_array", &S_skip_array, nil},
    {"_skip_object", &S_skip_object, nil},
    {"_skip_number", &S_skip_number, nil},
    {"_validate_one", &S_validate_one, &ValidateOne},
    {"_validate_utf8", &S_validate_utf8, &ValidateUTF8},
    {"_validate_utf8_fast", &S_validate_utf8_fast, &ValidateUTF8Fast},
}


func useAVX() {
    loader.WrapGoC(avx.Text__native_entry__, avx.Funcs, Stubs, "avx", "avx/native.c")
}

func useAVX2() {
    loader.WrapGoC(avx2.Text__native_entry__, avx2.Funcs, Stubs, "avx2", "avx2/native.c")
}

func useSSE() {
    loader.WrapGoC(sse.Text__native_entry__, sse.Funcs, Stubs, "sse", "sse/native.c")
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
