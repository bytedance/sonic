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

// The arm64 natives are compiled with `-ffixed-x28 -ffixed-x18 -ffixed-x9` (see
// tools/asm2arm_tool/scripts/generate_native_go.sh). Two of those must never
// appear in the generated code:
//
//   - x28 is Go's goroutine (g) register on arm64. Clobbering it corrupts the
//     runtime's view of the current goroutine.
//   - x18 is the platform register. On Apple platforms the kernel does not
//     preserve it across exception returns, so any value the natives park there
//     is asynchronously zeroed. This is not theoretical: a build that lost
//     -ffixed-x18 segfaulted on macOS with a fault address exactly equal to the
//     displacement of the first x18-relative load.
//
// x9 is deliberately *not* asserted here. It is reserved from the register
// allocator, but LLVM's AArch64 frame lowering still uses it as a fixed scratch
// register when addressing scalable (SVE) stack slots -- `add x9, sp, #N`
// feeding `addpl`. That is unavoidable and harmless: x9 is call-clobbered under
// AAPCS and Go never keeps anything there. The sve natives have always
// contained it.
//
// These registers are enforced only by compiler flags, so nothing in the normal
// build checks them. Dropping a flag, or regenerating with a compiler that
// defaults to a triple where the register is allocatable, silently produces
// natives that pass every functional test on Linux and then crash elsewhere.
// This test pins the invariant to the checked-in artifacts instead.
//
// It reads the generated sources rather than the compiled package, so it is
// arch-independent and runs on every CI platform, not just arm64.
var reservedRegs = []struct {
	name   string
	reason string
}{
	{"x28", "Go's goroutine (g) register on arm64"},
	{"x18", "the platform register; not preserved by the darwin/arm64 kernel"},
}

// Every flavour of generated arm64 native, and where its machine code lives.
// neon and sve_wrapgoc are loaded through internal/loader, so their code sits in
// byte slices; sve_linkname is statically linked, so its code sits in Go asm.
var generatedNatives = []struct {
	pkg  string
	glob string
}{
	{"neon", "neon/*_text_arm64.go"},
	{"sve_wrapgoc", "sve_wrapgoc/*_text_arm64.go"},
	{"sve_linkname", "sve_linkname/*_arm64.s"},
}

func TestGeneratedNativesAvoidReservedRegisters(t *testing.T) {
	for _, flavour := range generatedNatives {
		flavour := flavour
		t.Run(flavour.pkg, func(t *testing.T) {
			files, err := filepath.Glob(flavour.glob)
			if err != nil {
				t.Fatalf("glob %s: %v", flavour.glob, err)
			}
			if len(files) == 0 {
				t.Fatalf("no files matched %s; the natives should be checked in alongside this test", flavour.glob)
			}

			for _, reg := range reservedRegs {
				reg := reg
				t.Run(reg.name, func(t *testing.T) {
					// Match the 64- and 32-bit views of the same register, as
					// whole words so hex byte literals never match.
					re := regexp.MustCompile(`\b[wx]` + strings.TrimPrefix(reg.name, "x") + `\b`)
					for _, file := range files {
						for _, hit := range scanDisassembly(t, file, re) {
							t.Errorf("%s uses reserved register %s (%s): %s",
								file, reg.name, reg.reason, hit)
						}
					}
				})
			}
		})
	}
}

// scanDisassembly reports the disassembly comments in a generated file that
// match re. Only the text after "//" is considered, so the encoded instruction
// bytes are never searched. A leading hex address, which the *_text_arm64.go
// files carry but the .s files do not, is dropped so it cannot look like an
// operand.
func scanDisassembly(t *testing.T, file string, re *regexp.Regexp) []string {
	t.Helper()

	body, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read %s: %v", file, err)
	}

	var hits []string
	for _, line := range strings.Split(string(body), "\n") {
		_, comment, ok := strings.Cut(line, "//")
		if !ok {
			continue
		}
		asm := strings.TrimSpace(comment)
		if head, rest, ok := strings.Cut(asm, " "); ok && strings.HasPrefix(head, "0x") {
			asm = rest
		}
		if re.MatchString(asm) {
			hits = append(hits, strings.Join(strings.Fields(asm), " "))
		}
	}
	return hits
}
