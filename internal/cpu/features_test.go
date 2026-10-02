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

package cpu

import (
	"os"
	"testing"

	"github.com/klauspost/cpuid/v2"
	"github.com/stretchr/testify/assert"
)

func TestHasSSE(t *testing.T) {
	expected := cpuid.CPU.Has(cpuid.SSE) && cpuid.CPU.Has(cpuid.CLMUL)
	assert.Equal(t, expected, HasSSE)

	if !cpuid.CPU.Has(cpuid.CLMUL) {
		assert.False(t, HasSSE, "HasSSE must be false if CLMUL is not supported")
	}
	if !cpuid.CPU.Has(cpuid.SSE) {
		assert.False(t, HasSSE, "HasSSE must be false if SSE is not supported")
	}
}

func TestHasAVX2(t *testing.T) {
	mode := os.Getenv("SONIC_MODE")
	if mode == "noavx" || mode == "noavx2" {
		assert.False(t, HasAVX2)
	} else {
		assert.Equal(t, cpuid.CPU.Has(cpuid.AVX2), HasAVX2)
	}
}
