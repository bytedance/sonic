//
// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Eval(t *testing.T) {
	p := new(Parser).SetSource(`3 - 2 * (5 + 6) ** 4 / 7 + (1 << (1234 % 23)) & 0x5436 ^ 0x5a5a - 2 | 1`)
	v, err := p.Parse(nil)
	require.NoError(t, err)
	r, err := v.Evaluate()
	require.NoError(t, err)
	assert.Equal(t, int64(7805), r)
}
