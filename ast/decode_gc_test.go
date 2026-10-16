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

package ast

import (
	"runtime"
	"testing"

	"github.com/bytedance/sonic/internal/native/types"
)

// Keep the calls in one function so that inlined scan addresses may remain
// live across the GC safe point.
//
//go:noinline
func skipBlankAcrossGC(src string) (int, int, int) {
	first := skipBlank(src, 0)
	runtime.GC()
	return first, skipBlank(src, 1), skipBlank(src, len(src))
}

func TestScannerStartAtEnd(t *testing.T) {
	src := "1"
	pos := len(src)
	want := -int(types.ERR_EOF)

	if got := skipBlank(src, pos); got != want {
		t.Fatalf("skipBlank returned %d, want %d", got, want)
	}
	if got := skipNumber(src, pos); got != want {
		t.Fatalf("skipNumber returned %d, want %d", got, want)
	}
	if got, _, _ := decodeInt64(src, pos); got != want {
		t.Fatalf("decodeInt64 returned %d, want %d", got, want)
	}
	if got, _, _ := decodeFloat64(src, pos); got != want {
		t.Fatalf("decodeFloat64 returned %d, want %d", got, want)
	}
	runtime.KeepAlive(src)
}

func TestSkipBlankEndPointerGC(t *testing.T) {
	// Use a large allocation whose logical end is recorded in span.limit.
	// A one-past-the-end unsafe.Pointer must not survive across runtime.GC.
	const size = 400001
	buf := make([]byte, size)
	for i := range buf {
		buf[i] = ' '
	}
	buf[size-1] = 'x'
	src := string(buf)
	first, second, atEnd := skipBlankAcrossGC(src)
	if first != size-1 || second != size-1 {
		t.Fatalf("skipBlank returned (%d, %d), want (%d, %d)", first, second, size-1, size-1)
	}
	if atEnd != -int(types.ERR_EOF) {
		t.Fatalf("skipBlank at end returned %d, want %d", atEnd, -int(types.ERR_EOF))
	}
	runtime.KeepAlive(src)
}
