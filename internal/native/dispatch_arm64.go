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

	`github.com/bytedance/sonic/internal/native/neon`
	`github.com/bytedance/sonic/internal/native/types`
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
func SkipOneFast(s *string, p *int) int

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func GetByPath(s *string, p *int, path *[]interface{}, m *types.StateMachine) int

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

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func F64toa(out *byte, val float64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func ValidateUTF8(s *string, p *int, m *types.StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func ValidateUTF8Fast(s *string) (ret int)

func useNeon() {
	S_f64toa = neon.S_f64toa
	S_f32toa = neon.S_f32toa
	S_i64toa = neon.S_i64toa
	S_u64toa = neon.S_u64toa
	S_lspace = neon.S_lspace
	S_quote = neon.S_quote
	S_unquote = neon.S_unquote
	S_value = neon.S_value
	S_vstring = neon.S_vstring
	S_vnumber = neon.S_vnumber
	S_vsigned = neon.S_vsigned
	S_vunsigned = neon.S_vunsigned
	S_skip_one = neon.S_skip_one
	S_skip_one_fast = neon.S_skip_one_fast
	S_skip_array = neon.S_skip_array
	S_skip_object = neon.S_skip_object
	S_skip_number = neon.S_skip_number
	S_get_by_path = neon.S_get_by_path
}

func init() {
	useNeon()
}
