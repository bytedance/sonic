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

package ast

import (
    `unsafe`
    `reflect`

    `github.com/bytedance/sonic/internal/rt`
)

var (
    byteType = rt.UnpackType(reflect.TypeOf(byte(0)))
)

//go:noescape
//go:linkname memmove runtime.memmove
//goland:noinspection GoUnusedParameter
func memmove(to unsafe.Pointer, from unsafe.Pointer, n uintptr)

//go:linkname unsafe_NewArray reflect.unsafe_NewArray
//goland:noinspection GoUnusedParameter
func unsafe_NewArray(typ *rt.GoType, n int) unsafe.Pointer

//go:linkname growslice runtime.growslice
//goland:noinspection GoUnusedParameter
func growslice(et *rt.GoType, old rt.GoSlice, cap int) rt.GoSlice