// +build !amd64,!arm64 go1.26 !go1.17 arm64,!go1.20

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

package utf8

import (
	"unicode/utf8"
)

// ValidateFallback validates UTF-8 encoded bytes using standard library.
// This is used when native UTF-8 validation is not available.
func Validate(src []byte) bool {
	return utf8.Valid(src)
}

// ValidateStringFallback validates UTF-8 encoded string using standard library.
// This is used when native UTF-8 validation is not available.
func ValidateString(src string) bool {
	return utf8.ValidString(src)
}

// CorrectWith corrects the invalid utf8 byte with repl string.
// This is the fallback implementation using standard library.
func CorrectWith(dst []byte, src []byte, repl string) []byte {
	for len(src) > 0 {
		r, size := utf8.DecodeRune(src)
		if r == utf8.RuneError && size == 1 {
			// Invalid UTF-8 byte, replace with repl
			dst = append(dst, repl...)
			src = src[1:]
		} else {
			// Valid UTF-8 sequence, copy it
			dst = append(dst, src[:size]...)
			src = src[size:]
		}
	}
	return dst
}
