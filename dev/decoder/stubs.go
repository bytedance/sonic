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

package decoder

import (
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
)

// runtime.maxElementSize
const _max_map_element_size uintptr = 128

func mapfast(vt reflect.Type) bool {
	return vt.Elem().Size() <= _max_map_element_size
}

//go:linkname mallocgc runtime.mallocgc
//goland:noinspection GoUnusedParameter
func mallocgc(size uintptr, typ *rt.GoType, needzero bool) unsafe.Pointer

//go:linkname makemap reflect.makemap
func makemap(*rt.GoType, int) unsafe.Pointer

//go:linkname mapassign runtime.mapassign
//goland:noinspection GoUnusedParameter
func mapassign(t *rt.GoMapType, h unsafe.Pointer, k unsafe.Pointer) unsafe.Pointer

//go:linkname mapassign_fast32 runtime.mapassign_fast32
//goland:noinspection GoUnusedParameter
func mapassign_fast32(t *rt.GoMapType, h unsafe.Pointer, k uint32) unsafe.Pointer

//go:linkname mapassign_fast64 runtime.mapassign_fast64
//goland:noinspection GoUnusedParameter
func mapassign_fast64(t *rt.GoMapType, h unsafe.Pointer, k uint64) unsafe.Pointer

//go:linkname mapassign_fast64ptr runtime.mapassign_fast64ptr
//goland:noinspection GoUnusedParameter
func mapassign_fast64ptr(t *rt.GoMapType, h unsafe.Pointer, k unsafe.Pointer) unsafe.Pointer

//go:linkname mapassign_faststr runtime.mapassign_faststr
//goland:noinspection GoUnusedParameter
func mapassign_faststr(t *rt.GoMapType, h unsafe.Pointer, s string) unsafe.Pointer
