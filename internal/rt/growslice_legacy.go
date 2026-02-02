// +build go1.16,!go1.20

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
    _ `unsafe`

    "github.com/bytedance/sonic/option"
)

// Growslice to newCap, not append length
// Note: the [old, newCap) will not be zeroed if et does not have any ptr data.
func GrowSlice(et *GoType, old GoSlice, newCap int) GoSlice {
	if newCap < old.Len {
		panic("growslice's newCap is smaller than old length")
	}
	if old.Cap < int(option.FastGrowSliceThreshold) {
		newCap *= int(option.FastGrowSliceFactor)
	}
	s := growslice(et, old, newCap)
	s.Len = old.Len
	return s
}

//go:linkname growslice runtime.growslice
//goland:noinspection GoUnusedParameter
func growslice(et *GoType, old GoSlice, cap int) GoSlice
