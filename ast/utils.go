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

    `github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
func mem2ptr(s []byte) unsafe.Pointer {
    return (*rt.GoSlice)(unsafe.Pointer(&s)).Ptr
}

//go:nosplit
func ptr2slice(s unsafe.Pointer, l int, c int) unsafe.Pointer {
    slice := &rt.GoSlice{
        Ptr: s,
        Len: l,
        Cap: c,
    }
    return unsafe.Pointer(slice)
}

//go:nosplit
func str2ptr(s string) unsafe.Pointer {
    return (*rt.GoString)(unsafe.Pointer(&s)).Ptr
}

//go:nosplit
func addr2str(p unsafe.Pointer, n int64) (s string) {
    (*rt.GoString)(unsafe.Pointer(&s)).Ptr = p
    (*rt.GoString)(unsafe.Pointer(&s)).Len = int(n)
    return
}

const _SPACE_CHAR_MASK = (1<<' ')|(1<<'\t')|(1<<'\r')|(1<<'\n')

func isSpace(c byte) bool {
    return (int(1<<c) & _SPACE_CHAR_MASK) != 0
}