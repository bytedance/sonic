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

func assertUnmarshal(t *testing.T, api sonic.API, input []byte, newfn func() interface{}) {
	sv, jv := newfn(), newfn()
	serr := api.Unmarshal(input, sv)
	jerr := json.Unmarshal(input, jv)
	assert.Equal(t, jv, sv)
	assert.Equal(t, serr == nil, jerr == nil)
}