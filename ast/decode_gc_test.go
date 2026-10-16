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
)

// Keep both calls in one function so that an inlined IndexChar result can
// be shared across the GC safe point.
//
//go:noinline
func skipBlankAcrossGC(src string) (int, int) {
	first := skipBlank(src, 0)
	runtime.GC()
	return first, skipBlank(src, 1)
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
	first, second := skipBlankAcrossGC(src)
	if first != size-1 || second != size-1 {
		t.Fatalf("skipBlank returned (%d, %d), want (%d, %d)", first, second, size-1, size-1)
	}
	runtime.KeepAlive(src)
}
