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
    `unsafe`

    _ `github.com/chenzhuoyu/base64x`

    `github.com/bytedance/sonic/internal/rt`
)

//go:linkname _subr__b64encode github.com/chenzhuoyu/base64x._subr__b64encode
var _subr__b64encode uintptr

//go:noescape
//go:linkname memmove runtime.memmove
//goland:noinspection ALL
func memmove(to unsafe.Pointer, from unsafe.Pointer, n uintptr)

//go:noescape
//go:linkname growslice runtime.growslice
//goland:noinspection ALL
func growslice(et *rt.GoType, old rt.GoSlice, cap int) rt.GoSlice

//go:noescape
//go:linkname assertI2I runtime.assertI2I
//goland:noinspection ALL
func assertI2I(inter *rt.GoType, i rt.GoIface) rt.GoIface

//go:noescape
//go:linkname mapiternext runtime.mapiternext
//goland:noinspection ALL
func mapiternext(it unsafe.Pointer)

//go:noescape
//go:linkname mapiterinit reflect.mapiterinit
//goland:noinspection ALL
func mapiterinit(t *rt.GoType, m unsafe.Pointer) unsafe.Pointer

//go:noescape
//go:linkname isValidNumber encoding/json.isValidNumber
//goland:noinspection ALL
func isValidNumber(s string) bool
