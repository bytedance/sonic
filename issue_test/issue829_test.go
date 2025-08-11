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

package issue_test

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue829_NullCharacterInVM(t *testing.T) {
	var v interface{}
	err := sonic.Unmarshal([]byte("[\u0000]"), &v)
	require.Error(t, err, "Null character should cause syntax error")
	
	var vStd interface{}
	errStd := json.Unmarshal([]byte("[\u0000]"), &vStd)
	require.Error(t, errStd, "Standard library should return error for null character")
	require.Equal(t, err != nil, errStd != nil)
}


