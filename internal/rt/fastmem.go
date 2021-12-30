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

//go:nosplit
func Get16(v []byte) int16 {
    return *(*int16)((*GoSlice)(unsafe.Pointer(&v)).Ptr)
}

//go:nosplit
func Get32(v []byte) int32 {
    return *(*int32)((*GoSlice)(unsafe.Pointer(&v)).Ptr)
}

//go:nosplit
func Get64(v []byte) int64 {
    return *(*int64)((*GoSlice)(unsafe.Pointer(&v)).Ptr)
}

//go:nosplit
func Mem2Str(v []byte) (s string) {
    (*GoString)(unsafe.Pointer(&s)).Len = (*GoSlice)(unsafe.Pointer(&v)).Len
    (*GoString)(unsafe.Pointer(&s)).Ptr = (*GoSlice)(unsafe.Pointer(&v)).Ptr
    return
}

//go:nosplit
func Str2Mem(s string) (v []byte) {
    (*GoSlice)(unsafe.Pointer(&v)).Cap = (*GoString)(unsafe.Pointer(&s)).Len
    (*GoSlice)(unsafe.Pointer(&v)).Len = (*GoString)(unsafe.Pointer(&s)).Len
    (*GoSlice)(unsafe.Pointer(&v)).Ptr = (*GoString)(unsafe.Pointer(&s)).Ptr
    return
}

//go:nosplit
func MoreStack(size uintptr)