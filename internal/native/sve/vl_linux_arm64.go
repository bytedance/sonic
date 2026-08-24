//go:build linux && arm64
// +build linux,arm64

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

import (
	"golang.org/x/sys/unix"
)

// VectorLength returns this thread's SVE vector length in bytes, or 0 if it cannot be
// determined.
//
// prctl(PR_SVE_GET_VL) is the per-thread value the kernel will actually give
// the natives, which is what matters. /proc/sys/abi/sve_default_vector_length is
// only the default new processes inherit and can be overridden.
func VectorLength() int {
	vl, err := unix.PrctlRetInt(unix.PR_SVE_GET_VL, 0, 0, 0, 0)
	if err != nil {
		return 0
	}
	return vl & unix.PR_SVE_VL_LEN_MASK
}
