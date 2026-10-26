// Copyright 2026 ByteDance Inc.
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

package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func requireSonicMarshalSnapshot(t *testing.T, api sonic.API, value interface{}, wantErr bool, wantJSON string) {
	t.Helper()
	out, err := api.Marshal(&value)
	require.Equal(t, wantErr, err != nil, err)
	if wantErr {
		require.Nil(t, out)
		return
	}
	require.JSONEq(t, wantJSON, string(out))
}

func requireSonicUnmarshalSnapshot(
	t *testing.T,
	api sonic.API,
	cas unmTestCase,
	wantUnmarshalErr bool,
	wantMarshalErr bool,
	wantJSON string,
) interface{} {
	t.Helper()
	value := cas.newfn()
	err := api.Unmarshal(cas.data, value)
	require.Equal(t, wantUnmarshalErr, err != nil, err)
	requireSonicValueSnapshot(t, api, value, wantMarshalErr, wantJSON)
	return value
}

func requireSonicValueSnapshot(t *testing.T, api sonic.API, value interface{}, wantErr bool, wantJSON string) {
	t.Helper()
	out, err := api.Marshal(value)
	require.Equal(t, wantErr, err != nil, err)
	if wantErr {
		require.Nil(t, out)
		return
	}
	require.JSONEq(t, wantJSON, string(out))
}
