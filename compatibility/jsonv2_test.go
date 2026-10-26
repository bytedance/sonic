//go:build go1.27 && goexperiment.jsonv2
// +build go1.27,goexperiment.jsonv2

/*
 * Copyright 2026 ByteDance Inc.
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

package compatibility_test

import (
	stdjson "encoding/json"
	jsonv2 "encoding/json/v2"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

// TestJSONV2Semantics records the intentional differences between the
// encoding/json v1 facade and the experimental encoding/json/v2 defaults.
func TestJSONV2Semantics(t *testing.T) {
	t.Run("DUP-KEY", func(t *testing.T) {
		data := []byte(`{"a":1,"a":2}`)
		var sonicV, stdV, v2V map[string]int

		require.NoError(t, sonic.ConfigStd.Unmarshal(data, &sonicV))
		require.NoError(t, stdjson.Unmarshal(data, &stdV))
		require.Error(t, jsonv2.Unmarshal(data, &v2V))
		require.Equal(t, map[string]int{"a": 2}, sonicV)
		require.Equal(t, sonicV, stdV)
	})

	t.Run("CASE-FOLD", func(t *testing.T) {
		type value struct{ Foo int }
		data := []byte(`{"foo":1}`)
		var sonicV, stdV, v2V value

		require.NoError(t, sonic.ConfigStd.Unmarshal(data, &sonicV))
		require.NoError(t, stdjson.Unmarshal(data, &stdV))
		require.NoError(t, jsonv2.Unmarshal(data, &v2V))
		require.Equal(t, 1, sonicV.Foo)
		require.Equal(t, sonicV, stdV)
		require.Equal(t, 0, v2V.Foo)
	})

	t.Run("NIL-SLICE", func(t *testing.T) {
		var value []int
		sonicOut, sonicErr := sonic.ConfigStd.Marshal(value)
		stdOut, stdErr := stdjson.Marshal(value)
		v2Out, v2Err := jsonv2.Marshal(value)

		require.NoError(t, sonicErr)
		require.NoError(t, stdErr)
		require.NoError(t, v2Err)
		require.Equal(t, "null", string(sonicOut))
		require.Equal(t, string(sonicOut), string(stdOut))
		require.Equal(t, "[]", string(v2Out))
	})

	t.Run("ARR-LEN", func(t *testing.T) {
		data := []byte(`[1,2,3]`)
		var sonicV, stdV, v2V [2]int

		require.NoError(t, sonic.ConfigStd.Unmarshal(data, &sonicV))
		require.NoError(t, stdjson.Unmarshal(data, &stdV))
		require.Error(t, jsonv2.Unmarshal(data, &v2V))
		require.Equal(t, [2]int{1, 2}, sonicV)
		require.Equal(t, sonicV, stdV)
	})

	t.Run("BYTE-ARR", func(t *testing.T) {
		value := [3]byte{1, 2, 3}
		sonicOut, sonicErr := sonic.ConfigStd.Marshal(value)
		stdOut, stdErr := stdjson.Marshal(value)
		v2Out, v2Err := jsonv2.Marshal(value)

		require.NoError(t, sonicErr)
		require.NoError(t, stdErr)
		require.NoError(t, v2Err)
		require.Equal(t, `[1,2,3]`, string(sonicOut))
		require.Equal(t, string(sonicOut), string(stdOut))
		require.Equal(t, `"AQID"`, string(v2Out))
	})

	t.Run("UTF8-DEC", func(t *testing.T) {
		data := []byte{'"', 0xff, '"'}
		var sonicV, stdV, v2V string

		require.NoError(t, sonic.ConfigStd.Unmarshal(data, &sonicV))
		require.NoError(t, stdjson.Unmarshal(data, &stdV))
		require.Error(t, jsonv2.Unmarshal(data, &v2V))
		require.Equal(t, "\ufffd", sonicV)
		require.Equal(t, sonicV, stdV)
	})
}
