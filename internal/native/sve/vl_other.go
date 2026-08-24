//go:build !(linux && arm64)
// +build !linux !arm64

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

package sve

// VectorLength reports the SVE vector length in bytes, or 0 when it cannot be
// established.
//
// prctl(PR_SVE_GET_VL) is Linux-specific. Everywhere else there is no way to
// ask, so report unknown and let the caller refuse to enable the SVE natives.
// darwin in particular has no HWCAP and no prctl, so this keeps macOS on neon
// even if x/sys/cpu ever grows SVE detection there.
func VectorLength() int {
	return 0
}
