
/*
 * Copyright 2022 ByteDance Inc.
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
	 `testing`
 
	 `encoding/json`
	 . `github.com/bytedance/sonic`
	 `github.com/stretchr/testify/require`
 )

 func TestIssue_UnmarshalBase64(t *testing.T) {
	var obj, stdobj []byte
	tests := []string {
		`"xy\r\nzu"`,
		`"xy\/\/"`,
		`"\/\/=="`,
		`"\/\/\u003d\u003d"`,
		`"\u0030\u0030\u0030\u003d"`,
	}
	for _, data := range(tests) {
		stderr := json.Unmarshal([]byte(data), &stdobj)
		err := Unmarshal([]byte(data), &obj)
		require.NoError(t, stderr, data)
		require.NoError(t, err, data)
		require.Equal(t, stdobj, obj, data)
	}
}

func TestIssue_UnmarshalBase64Error(t *testing.T) {
	var obj, stdobj []byte
	tests := []string {
		`"xy\r\nzu0==="`,
		`"xy\/\/`,
		`"\/\/==`,
		`"\/\/\u003d0\u003d"`,
		`"\u0030\u0030\u0030\u003d`,
	}
	for _, data := range(tests) {
		stderr := json.Unmarshal([]byte(data), &stdobj)
		err := Unmarshal([]byte(data), &obj)
		require.Equal(t, stderr != nil, err != nil)
	}
	var _, _ = obj, stdobj
}