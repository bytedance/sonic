//go:build amd64

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
	"os/exec"
	"testing"
)

func TestNoPCLMULQDQ_SubProcess(t *testing.T) {
	// This test verifies that Sonic works correctly when PCLMULQDQ is unavailable.
	// It runs a subprocess with SONIC_MODE=noavx to force the SSE path (scalar prefix-XOR).
	if os.Getenv("SONIC_TEST_NOPCLMUL") == "1" {
		// We are in the subprocess — verify flags are disabled
		if HasAVX2 {
			t.Fatal("HasAVX2 should be false with SONIC_MODE=noavx")
		}
		if HasPCLMULQDQ {
			t.Fatal("HasPCLMULQDQ should be false with SONIC_MODE=noavx")
		}
		return
	}

	// Parent process: launch subprocess with SONIC_MODE=noavx
	cmd := exec.Command(os.Args[0], "-test.run=^TestNoPCLMULQDQ_SubProcess$", "-test.v")
	cmd.Env = append(os.Environ(), "SONIC_MODE=noavx", "SONIC_TEST_NOPCLMUL=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("subprocess failed: %v\n%s", err, out)
	}
	t.Logf("subprocess output:\n%s", out)
}
