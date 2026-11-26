/*
 * Copyright 2025 ByteDance Inc.
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
	"unicode"
	"unicode/utf8"
)

// FoldName returns a folded string such that FoldName(x) == FoldName(y)
// is identical to bytes.EqualFold(x, y). It is kept in sync with
// encoding/json, so that case-insensitive field-name matching agrees
// with the standard library.
func FoldName(in []byte) []byte {
	// This is inlinable to take advantage of "function outlining".
	var arr [32]byte // large enough for most JSON names
	return appendFoldedName(arr[:0], in)
}

// FoldNameString is FoldName for a string key.
func FoldNameString(s string) string {
	return string(FoldName(Str2Mem(s)))
}

func appendFoldedName(out, in []byte) []byte {
	for i := 0; i < len(in); {
		// Handle single-byte ASCII.
		if c := in[i]; c < utf8.RuneSelf {
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
			out = append(out, c)
			i++
			continue
		}
		// Handle multi-byte Unicode.
		r, n := utf8.DecodeRune(in[i:])
		out = utf8.AppendRune(out, foldRune(r))
		i += n
	}
	return out
}

// foldRune returns the smallest rune for all runes in the same fold set.
func foldRune(r rune) rune {
	for {
		r2 := unicode.SimpleFold(r)
		if r2 <= r {
			return r2
		}
		r = r2
	}
}
