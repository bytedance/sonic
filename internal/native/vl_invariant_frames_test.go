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

// The SVE natives run with one pcsp table per function on every supported
// vector length. That is only true because their scalable stack allocations
// are sized for sve.MaxVectorLength at generation time (asm2arm_tool
// --max-vl): `addvl sp, sp, #-N` becomes `sub sp, sp, #N*32`, so the frame is
// the 256-bit frame on a 128-bit machine too. The rewrite is sound because
// LLVM anchors every scalable object on the frame base (x9), which is computed
// before the allocation, and never on sp.
//
// Both halves of that are invariants of the generated artifacts, not of any
// Go code, so pin them here where every CI arch can check them:
//
//   - no instruction may adjust sp by a multiple of the vector length
//     (addvl/addpl on sp, or sp += xN);
//   - no VL-scaled memory access may be anchored on sp.
//
// A regeneration that violates either would produce natives whose frames
// differ by machine and whose pcsp tables are wrong everywhere but VL=32 --
// invisible to every functional test, visible only as an unrecoverable fault
// or a corrupted GC stack scan on 128-bit hardware.
func TestSveNativesHaveVectorLengthInvariantFrames(t *testing.T) {
	forbidden := []struct {
		name string
		re   *regexp.Regexp
	}{
		{"addvl/addpl writing sp", regexp.MustCompile(`\b(addvl|addpl)\s+sp\b`)},
		{"addvl/addpl reading sp", regexp.MustCompile(`\b(addvl|addpl)\s+\w+,\s*sp\b`)},
		{"sp adjusted by a register", regexp.MustCompile(`\b(add|sub)\s+sp,\s*sp,\s*x\d+`)},
		{"VL-scaled access anchored on sp", regexp.MustCompile(`\[\s*sp\s*,[^\]]*mul vl`)},
	}
	// SL mode records the original instruction next to the rewrite; JIT mode
	// copies the linked ELF, so its comments only show the result. A native
	// generated without --max-vl would fail the addvl check above regardless;
	// this is a second, positive signal for the flavour that can give one.
	rewritten := regexp.MustCompile(`frame sized for VL=`)

	for _, flavour := range generatedNatives {
		if flavour.pkg == "neon" {
			continue // no SVE, no scalable stack
		}
		flavour := flavour
		t.Run(flavour.pkg, func(t *testing.T) {
			files, err := filepath.Glob(flavour.glob)
			if err != nil {
				t.Fatalf("glob %s: %v", flavour.glob, err)
			}
			if len(files) == 0 {
				t.Fatalf("no files matched %s; the natives should be checked in alongside this test", flavour.glob)
			}
			seenRewrite := 0
			for _, file := range files {
				for _, line := range disassemblyComments(t, file) {
					if rewritten.MatchString(line) {
						seenRewrite++
					}
					// Judge the instruction, not the note recording what it
					// replaced.
					insn, _, _ := strings.Cut(line, "(was")
					for _, f := range forbidden {
						if f.re.MatchString(insn) {
							t.Errorf("%s: %s: %s", file, f.name, line)
						}
					}
				}
			}
			if flavour.pkg == "sve_linkname" && seenRewrite == 0 {
				t.Errorf("%s: no native carries a 'frame sized for VL=' rewrite; was this generated without --max-vl?", flavour.pkg)
			}
		})
	}
}

// disassemblyComments returns every disassembly comment in a generated file,
// with the leading hex address (present in *_text_arm64.go, absent in .s)
// dropped and whitespace normalised. Only the text after "//" is returned, so
// encoded instruction bytes are never examined.
func disassemblyComments(t *testing.T, file string) []string {
	t.Helper()

	body, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read %s: %v", file, err)
	}

	var out []string
	for _, line := range strings.Split(string(body), "\n") {
		_, comment, ok := strings.Cut(line, "//")
		if !ok {
			continue
		}
		asm := strings.TrimSpace(comment)
		if head, rest, ok := strings.Cut(asm, " "); ok && strings.HasPrefix(head, "0x") {
			asm = rest
		}
		out = append(out, strings.Join(strings.Fields(asm), " "))
	}
	return out
}
