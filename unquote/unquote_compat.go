// +build !amd64,!arm64 go1.22 !go1.16 arm64,!go1.20

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

package unquote

import (
	"runtime"
	"strconv"
	"unsafe"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

func String(s string) (ret string, err types.ParsingError) {
    mm := make([]byte, 0, len(s))
    err = intoBytesUnsafe(s, &mm)
    ret = rt.Mem2Str(mm)
    return
}

func IntoBytes(s string, m *[]byte) types.ParsingError {
    if cap(*m) < len(s) {
        return types.ERR_EOF
    } else {
        return intoBytesUnsafe(s, m)
    }
}

func intoBytesUnsafe(s string, m *[]byte) types.ParsingError {
    o, e := strconv.Unquote(s)
    if e != nil {
        return types.ERR_INVALID_ESCAPE
    }
    *m = append(*m, o...)
    return 0
}
