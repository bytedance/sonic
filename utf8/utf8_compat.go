// +build !amd64,!arm64 go1.22 !go1.16 arm64,!go1.20

/*
 * Copyright 2022 ByteDance Inc.
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

package utf8

import (
    `unicode/utf8`
    `github.com/bytedance/sonic/internal/rt`
)

// CorrectWith corrects the invalid utf8 byte with repl string.
func CorrectWith(dst []byte, src []byte, repl string) []byte {
    for len(src) > 0 {
        r, size := utf8.DecodeRune(src)
        if r == utf8.RuneError && (size == 1 || size == 0) {
            dst = append(dst, repl...)
        } else {
            dst = append(dst, string(r)...)
        }
        src = src[size:]
    }
    return dst
}

// Validate is a simd-accelereated drop-in replacement for the standard library's utf8.Valid.
func Validate(src []byte) bool {
    return ValidateString(rt.Mem2Str(src))
}

// ValidateString as Validate, but for string.
func ValidateString(src string) bool {
    return utf8.ValidString(src)
}

