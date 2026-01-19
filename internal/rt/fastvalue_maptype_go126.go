//go:build go1.26
// +build go1.26

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

import "unsafe"

// GoMapType for Go 1.26+ Swiss maps.
// This matches internal/abi.MapType in Go 1.26.
type GoMapType struct {
	GoType
	Key       *GoType
	Elem      *GoType
	Group     *GoType // internal type representing a slot group
	Hasher    func(unsafe.Pointer, uintptr) uintptr
	GroupSize uintptr // == Group.Size_
	SlotSize  uintptr // size of key/elem slot
	ElemOff   uintptr // offset of elem in key/elem slot
	Flags     uint32
}

// Flag values for Go 1.26 Swiss maps
const (
	mapNeedKeyUpdate  = 1 << iota // 1
	mapHashMightPanic             // 2
	mapIndirectKey                // 4
	mapIndirectElem               // 8
)

func (self *GoMapType) IndirectElem() bool {
	return self.Flags&mapIndirectElem != 0
}
