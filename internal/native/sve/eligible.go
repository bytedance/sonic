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

// Package sve decides whether this machine may run sonic's SVE natives.
//
// It lives apart from internal/native so that the sve_linkname and sve_wrapgoc
// packages can use it too. They cannot import internal/native, which imports
// them, and before this each carried its own private copy of the check in its
// test files -- copies that were never updated alongside the real one, so the
// tests gated on different hardware than the dispatcher did and skipped
// themselves on machines where the natives were live.
package sve

import (
	"bufio"
	"os"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	xcpu "golang.org/x/sys/cpu"
)

// Two independent things have to line up for the SVE natives to work at a
// given vector length.
//
// The kernels must compute the right answer at that width. They originally did
// not: native/scanning.h read predicate registers as fixed 32-bit lane masks,
// so it was correct only at 32 bytes. That is now expressed vector-length
// agnostically in native/sve_compat.h, so the results are right at any width.
//
// And the frame metadata handed to Go must match, because a native's scalable
// spill area is VL-sized. Get that wrong and results are still correct but
// unwinding is not: Go computes the wrong caller SP, so a fault inside a
// native cannot be recovered and precise GC stack scanning walks the wrong
// frame. Rather than carry one table per width, the natives are generated with
// every scalable allocation sized for MaxVectorLength (asm2arm_tool --max-vl),
// so the frame is identical on every machine with VL <= MaxVectorLength and a
// single table serves both flavours -- including sve_linkname, whose frame
// sizes are baked into the TEXT directives of generated Go assembly.

// MaxVectorLength is the SVE vector length, in bytes, that the natives' stack
// frames are sized for. Must match SVE_MAX_VL in
// tools/asm2arm_tool/scripts/generate_native_go.sh. A machine with a wider
// vector would overrun those frames, so it is not eligible.
const MaxVectorLength = 32

// SupportedVectorLengths lists the widths the natives have been verified on:
// 16 bytes (Neoverse V2/N2: Graviton4, Cobalt, Axion, Grace) and 32 bytes
// (Neoverse V1: Graviton3; Kunpeng). These are also the only SVE widths
// shipping in general-purpose parts.
var SupportedVectorLengths = [...]int{16, 32}

// SupportsVectorLength reports whether the natives may run at vl bytes.
func SupportsVectorLength(vl int) bool {
	if vl > MaxVectorLength {
		return false
	}
	for _, v := range SupportedVectorLengths {
		if vl == v {
			return true
		}
	}
	return false
}

// Eligible reports whether this CPU may run the SVE natives.
//
// The CPU must implement SVE, and its vector length must be one the natives
// support (see SupportedVectorLengths). Checking for SVE alone is not enough:
// a wider vector than the frames were sized for would overrun them.
//
// Ask the kernel rather than keeping a vendor list: x/sys/cpu decodes AT_HWCAP,
// which the kernel derives from ID_AA64PFR0_EL1 at boot. That picks up
// Graviton3 and Graviton4, which the original Kunpeng part-id check excluded
// despite them working. That check is kept as a fallback for when the vector
// length cannot be established at all, so existing deployments cannot regress.
//
// Eligibility is not selection: sonic still requires SONIC_USE_SVE_WRAPGOC or
// SONIC_USE_SVE_LINKNAME before anything but neon is used.
func Eligible() bool {
	if xcpu.ARM64.HasSVE {
		if vl := VectorLength(); vl != 0 {
			return SupportsVectorLength(vl)
		}
		// Vector length unknown; fall through to the legacy check.
	}
	return isKunpeng()
}

// Describe returns a short human-readable account of the decision, for tests
// and CI logs. A path that quietly falls back to neon looks identical to one
// that ran the natives, which is how the SVE breakage stayed hidden.
func Describe() string {
	var b strings.Builder
	b.WriteString("sve: ")
	if !xcpu.ARM64.HasSVE {
		b.WriteString("not implemented by this CPU")
	} else if vl := VectorLength(); vl == 0 {
		b.WriteString("present, vector length unknown")
	} else {
		b.WriteString("present, vector length ")
		b.WriteString(itoa(vl))
		b.WriteString(" bytes (natives support")
		for _, v := range SupportedVectorLengths {
			b.WriteString(" ")
			b.WriteString(itoa(v))
		}
		b.WriteString(")")
	}
	if Eligible() {
		b.WriteString("; eligible")
	} else {
		b.WriteString("; NOT eligible, neon will be used")
	}
	return b.String()
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	return string(buf[i:])
}

func isKunpeng() bool {
	cpuinfo, err := cpu.Info()
	if err == nil && len(cpuinfo) != 0 {
		if cpuinfo[0].Model == "0xd02" || cpuinfo[0].Model == "0xd06" {
			return true
		}
	}

	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "CPU part") {
			parts := strings.SplitN(line, ":", 2)
			model := strings.TrimSpace(parts[1])
			if model == "0xd02" || model == "0xd06" {
				return true
			}
		}
	}
	return false
}
