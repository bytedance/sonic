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

package rt

import (
    `unsafe`
)

//go:noescape
//go:linkname Memmove runtime.memmove
func Memmove(to unsafe.Pointer, from unsafe.Pointer, n uintptr)

//go:linkname Mapiternext runtime.mapiternext
func Mapiternext(it *GoMapIterator)

//go:linkname Mapiterinit runtime.mapiterinit
func Mapiterinit(t *GoMapType, m *GoMap, it *GoMapIterator)

//go:linkname IsValidNumber encoding/json.isValidNumber
func IsValidNumber(s string) bool

//go:noescape
//go:linkname MemclrNoHeapPointers runtime.memclrNoHeapPointers
func MemclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)


