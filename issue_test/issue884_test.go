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

package issue_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue884Float32Precision(t *testing.T) {
	input := []byte(`[7.328900098800659,-7.328900098800659,3.4028235e38,-0]`)
	var want, got []float32
	if err := json.Unmarshal(input, &want); err != nil {
		t.Fatal(err)
	}
	if err := sonic.Unmarshal(input, &got); err != nil {
		t.Fatal(err)
	}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if math.Float32bits(got[i]) != math.Float32bits(want[i]) {
			t.Errorf("element %d: got %g, want %g", i, got[i], want[i])
		}
	}
}
