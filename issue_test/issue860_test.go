package issue_test

/**
 * Copyright 2025 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"encoding"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue860(t *testing.T) {
	var m1 map[encoding.TextUnmarshaler]interface{}
	err := sonic.Unmarshal([]byte(`{"a":1}`), &m1)
	t.Logf("sonic err: %#v", err)
	var m2 map[encoding.TextUnmarshaler]interface{}
	err1 := json.Unmarshal([]byte(`{"a":1}`), &m2)
	require.Equal(t, err1 == nil, err == nil)
}
