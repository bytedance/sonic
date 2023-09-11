// Code generated by Makefile, DO NOT EDIT.

// Code generated by Makefile, DO NOT EDIT.

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

package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
)


//go:nosplit
func i64toa(out *byte, val int64) (ret int) {
    return __i64toa(out, val)
}

//go:nosplit
func u64toa(out *byte, val uint64) (ret int) {
    return __u64toa(out, val)
}

//go:nosplit
func f64toa(out *byte, val float64) (ret int) {
    return __f64toa(out, val)
}

//go:nosplit
func f32toa(out *byte, val float32) (ret int) {
    return __f32toa(out, val)
}

//go:nosplit
func lspace(sp unsafe.Pointer, nb int, off int) (ret int) {
    return __lspace(sp, nb, off)
}

//go:nosplit
func quote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) (ret int) {
    return __quote(sp, nb, dp, dn, flags)
}

//go:nosplit
func html_escape(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) (ret int) {
    return __html_escape(sp, nb, dp, dn)
}

//go:nosplit
func unquote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) (ret int) {
    return __unquote(sp, nb, dp, ep, flags)
}

//go:nosplit
func value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) (ret int) {
    return __value(s, n, p, v, flags)
}

//go:nosplit
func vstring(s *string, p *int, v *types.JsonState, flags uint64) {
    __vstring(s, p, v, flags)
}

//go:nosplit
func vnumber(s *string, p *int, v *types.JsonState) {
    __vnumber(s, p, v)
}

//go:nosplit
func vsigned(s *string, p *int, v *types.JsonState) {
    __vsigned(s, p, v)
}

//go:nosplit
func vunsigned(s *string, p *int, v *types.JsonState) {
    __vunsigned(s, p, v)
}

//go:nosplit
func skip_one(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
    return __skip_one(s, p, m, flags)
}

//go:nosplit
func skip_one_fast(s *string, p *int) (ret int) {
    return __skip_one_fast(s, p)
}

//go:nosplit
func skip_array(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
    return __skip_array(s, p, m, flags)
}

//go:nosplit
func skip_object(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
    return __skip_object(s, p, m, flags)
}

//go:nosplit
func skip_number(s *string, p *int) (ret int) {
    return __skip_number(s, p)
}

//go:nosplit
func validate_one(s *string, p *int, m *types.StateMachine) (ret int) {
    return __validate_one(s, p, m)
}

//go:nosplit
func get_by_path(s *string, p *int, path *[]interface{}, m *types.StateMachine) (ret int) {
    return __get_by_path(s, p, path, m)
}

//go:nosplit
func validate_utf8(s *string, p *int, m *types.StateMachine) (ret int) {
    return __validate_utf8(s, p, m)
}

//go:nosplit
func validate_utf8_fast(s *string)  (ret int) {
    return __validate_utf8_fast(s)
}

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __i64toa(out *byte, val int64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __u64toa(out *byte, val uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __f64toa(out *byte, val float64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __f32toa(out *byte, val float32) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __lspace(sp unsafe.Pointer, nb int, off int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __quote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __html_escape(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __unquote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vstring(s *string, p *int, v *types.JsonState, flags uint64)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vnumber(s *string, p *int, v *types.JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vsigned(s *string, p *int, v *types.JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vunsigned(s *string, p *int, v *types.JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_one(s *string, p *int, m *types.StateMachine, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_one_fast(s *string, p *int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_array(s *string, p *int, m *types.StateMachine, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_object(s *string, p *int, m *types.StateMachine, flags uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_number(s *string, p *int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __validate_one(s *string, p *int, m *types.StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __get_by_path(s *string, p *int, path *[]interface{}, m *types.StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __validate_utf8(s *string, p *int, m *types.StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __validate_utf8_fast(s *string)  (ret int)
