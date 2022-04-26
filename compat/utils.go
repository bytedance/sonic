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

package compat

import (
	`unsafe`
)

type goSlice struct {
    Ptr unsafe.Pointer
    Len int
    Cap int
}

type goString struct {
    Ptr unsafe.Pointer
    Len int
}

//go:nosplit
func mem2Str(v []byte) (s string) {
    (*goString)(unsafe.Pointer(&s)).Len = (*goSlice)(unsafe.Pointer(&v)).Len
    (*goString)(unsafe.Pointer(&s)).Ptr = (*goSlice)(unsafe.Pointer(&v)).Ptr
    return
}

//go:nosplit
func str2Mem(s string) (v []byte) {
    (*goSlice)(unsafe.Pointer(&v)).Cap = (*goString)(unsafe.Pointer(&s)).Len
    (*goSlice)(unsafe.Pointer(&v)).Len = (*goString)(unsafe.Pointer(&s)).Len
    (*goSlice)(unsafe.Pointer(&v)).Ptr = (*goString)(unsafe.Pointer(&s)).Ptr
    return
}