// Copyright 2026 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
	stdjson "encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue938_InvalidEscapesInSkippedString(t *testing.T) {
	type target struct{}

	for _, escape := range []string{`\0`, `\1`, `\v`, `\x`, `\u12xz`} {
		for _, prefixLen := range []int{0, 31, 63, 64} {
			t.Run(fmt.Sprintf("%s/offset-%d", escape, prefixLen), func(t *testing.T) {
				input := []byte(`{"ignored":"` + strings.Repeat("p", prefixLen) + escape + strings.Repeat("p", 64) + `"}`)

				var stdValue, sonicValue target
				stdErr := stdjson.Unmarshal(input, &stdValue)
				sonicErr := sonic.ConfigStd.Unmarshal(input, &sonicValue)

				require.Error(t, stdErr)
				require.Error(t, sonicErr)
			})
		}
	}

	for _, value := range []string{
		`\"`, `\\`, `\/`, `\b`, `\f`, `\n`, `\r`, `\t`, `\u0061`,
		strings.Repeat("p", 63) + `\\x`,
	} {
		t.Run("valid/"+value, func(t *testing.T) {
			input := []byte(`{"ignored":"` + value + strings.Repeat("p", 64) + `"}`)

			var stdValue, sonicValue target
			stdErr := stdjson.Unmarshal(input, &stdValue)
			sonicErr := sonic.ConfigStd.Unmarshal(input, &sonicValue)

			require.NoError(t, stdErr)
			require.NoError(t, sonicErr)
		})
	}
}
