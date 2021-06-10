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

    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
func i64tof(v int64) float64 {
    return *(*float64)(unsafe.Pointer(&v))
}

//go:nosplit
func f64toi(v float64) int64 {
    return *(*int64)(unsafe.Pointer(&v))
}

//go:nosplit
func mem2ptr(s []byte) unsafe.Pointer {
    return (*rt.GoSlice)(unsafe.Pointer(&s)).Ptr
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

func unquoteBytes(s string, m *[]byte) types.ParsingError {
    pos := -1
    slv := (*rt.GoSlice)(unsafe.Pointer(m))
    str := (*rt.GoString)(unsafe.Pointer(&s))
    ret := native.Unquote(str.Ptr, str.Len, slv.Ptr, &pos, 0)

    /* check for errors */
    if ret < 0 {
        return types.ParsingError(-ret)
    }

    /* update the length */
    slv.Len = ret
    return 0
}

func UnquoteString(s string) (ret string, err types.ParsingError) {
    mm := make([]byte, 0, len(s))
    err = unquoteBytes(s, &mm)
    ret = rt.Mem2Str(mm)
    return
}
