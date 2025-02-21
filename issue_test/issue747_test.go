// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
	"testing"

	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestIssue747(t *testing.T) {
	tests := []struct {
		name           string
		input         string
		expected interface{}
		newfn func() interface{}
	}{
		{
			name:   "unmarshal map key float64",
			input: `{"1.2":1.8}`,
			expected: &map[float64]float64{
				1.2:  1.8,
			},
			newfn: func() interface{} { return new(map[float64]float64) },
		},
		{
			name:    "unmarshal map key float32",
			input: `{"1.2":1.8}`,
			expected: &map[float32]float32{
				1.2:  1.8,
			},
			newfn: func() interface{} { return new(map[float32]float32) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv, jv := tt.newfn(), tt.newfn()
			serr := sonic.Unmarshal([]byte(tt.input), &sv)
			assert.Equal(t, tt.expected, sv)
			assert.NoError(t, serr)

			// Note: it is different from encoding/json 
			jerr := json.Unmarshal([]byte(tt.input), &jv)
			assert.NotEqual(t, jerr == nil, serr == nil)
			assert.NotEqual(t, jv, sv)
		})
	}
}
