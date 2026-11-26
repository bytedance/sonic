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

package sonic

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic/internal/envs"
	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	require.False(t, Valid(nil))

	testCase := []struct {
		data     string
		expected bool
	}{
		{``, false},
		{`s`, false},
		{`{`, false},
		{`[`, false},
		{`[1,2`, false},
		{`{"so":nic"}`, false},

		{`null`, true},
		{`""`, true},
		{`1`, true},
		{`"sonic"`, true},
		{`{}`, true},
		{`[]`, true},
		{`[1,2]`, true},
		{`{"so":"nic"}`, true},
	}
	for _, tc := range testCase {
		require.Equal(t, tc.expected, Valid([]byte(tc.data)), tc.data)
	}
}

func TestConfigUseUnicodeErrors(t *testing.T) {
	if APIKind == UseStdJSON {
		t.Skip("UseUnicodeErrors is not supported by encoding/json fallback")
	}
	if envs.UseOptDec {
		t.Skip("UseUnicodeErrors is not supported by OptDec")
	}

	var got string
	err := Config{UseUnicodeErrors: true}.Froze().UnmarshalFromString(`"\ud800"`, &got)
	require.Error(t, err)
}

func TestIdent(t *testing.T) {
	foo := struct {
		Name string
		Age  int
	}{
		Name: "sonic",
		Age:  20,
	}

	out, err := MarshalIndent(&foo, "", "  ")
	require.Nil(t, err)
	require.Equal(t, `{
  "Name": "sonic",
  "Age": 20
}`, string(out))
}

func TestUnmarshalCaseFoldedFieldName(t *testing.T) {
	type dotted struct{ Id int }

	// U+0130 lowercases to ASCII 'i' but is not in its fold set, so
	// encoding/json leaves the field untouched.
	data := []byte("{\"\u0130d\":1}")
	var sv, jv dotted
	require.NoError(t, json.Unmarshal(data, &jv))
	require.NoError(t, ConfigStd.Unmarshal(data, &sv))
	require.Equal(t, 0, sv.Id)
	require.Equal(t, jv, sv)

	type sigma struct {
		V int `json:"σ"`
	}

	// U+03C2 and U+03C3 share a fold set, so both decoders match them.
	data = []byte("{\"\u03c2\":1}")
	var sv2, jv2 sigma
	require.NoError(t, json.Unmarshal(data, &jv2))
	require.NoError(t, ConfigStd.Unmarshal(data, &sv2))
	require.Equal(t, 1, sv2.V)
	require.Equal(t, jv2, sv2)
}
