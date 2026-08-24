/*
 * Copyright 2026 ByteDance Inc.
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

package native

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// sve_linkname is produced by asm2arm_tool's SL mode, which rewrites clang's
// assembly as Go assembly. Most instructions pass through as raw WORDs, but
// branches are re-expressed in Go mnemonics so the Go assembler can resolve the
// labels. That rewrite is where semantics can leak.
//
// cbz/cbnz/tbz/tbnz must become the assembler's own CBZ/CBNZ/CBZW/CBNZW/TBZ/TBNZ.
// An earlier version of the tool lowered them to CMP/TST + B.cond instead. The
// originals do not touch NZCV, and the compiler schedules them freely between a
// flag-setting instruction and the B.cond or CSEL that consumes those flags, so
// a synthesized compare in that window clobbers flags that are still live.
// skip_one hit exactly this:
//
//	subs x17, x8, x5        // nb = len - p; Z set when len == p
//	tbnz w3, #5, ...        // dispatch on a flags bit
//	b.eq ERR_EOF            // the len == p check
//
// The TST turned the b.eq into "flags bit 5 clear", and every top-level string
// scan returned ERR_EOF. The 64-bit CMP against ZR also widened every 32-bit
// cbz/cbnz, which happened to be benign only because the upper halves were zero.
//
// This pins the invariant to the checked-in artifacts: every compare-and-branch
// the compiler emitted must survive as a single native compare-and-branch, with
// the operand width preserved. It reads the generated sources, so it runs on
// every platform, not just arm64.
func TestLinknameNativesKeepCompareAndBranch(t *testing.T) {
	files, err := filepath.Glob("sve_linkname/*_arm64.s")
	if err != nil {
		t.Fatalf("glob: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("no sve_linkname/*_arm64.s files; the natives should be checked in alongside this test")
	}

	// The original instruction, recorded by the tool in a trailing comment.
	orig := regexp.MustCompile(`//\s*(cbz|cbnz|tbz|tbnz)\s+([wx])(\w+),`)
	// What the Go line in front of that comment must be.
	want := map[string]map[string]string{
		"cbz":  {"w": "CBZW", "x": "CBZ"},
		"cbnz": {"w": "CBNZW", "x": "CBNZ"},
		"tbz":  {"w": "TBZ", "x": "TBZ"},
		"tbnz": {"w": "TBNZ", "x": "TBNZ"},
	}

	for _, file := range files {
		body, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("read %s: %v", file, err)
		}
		lines := strings.Split(string(body), "\n")
		sites := 0
		for i, line := range lines {
			m := orig.FindStringSubmatch(line)
			if m == nil {
				continue
			}
			sites++
			op, width := m[1], m[2]
			code := strings.TrimSpace(strings.SplitN(line, "//", 2)[0])
			mnemonic := strings.Fields(code)[0]
			if got, ok := want[op][width]; !ok || mnemonic != got {
				t.Errorf("%s:%d: %s %s… lowered to %q, want %s: %s",
					file, i+1, op, width, mnemonic, want[op][width], strings.TrimSpace(line))
			}
			// A synthesized compare in front of the branch is the exact bug.
			if i > 0 {
				prev := strings.TrimSpace(lines[i-1])
				if strings.HasPrefix(prev, "CMP ") || strings.HasPrefix(prev, "TST ") {
					t.Errorf("%s:%d: %s preceded by a synthesized %q, which clobbers NZCV",
						file, i+1, op, prev)
				}
			}
		}
		if sites == 0 {
			t.Errorf("%s: no cbz/cbnz/tbz/tbnz sites found; the tool's disassembly comments may have changed shape", file)
		}
	}
}
