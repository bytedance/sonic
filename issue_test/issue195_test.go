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

package issue_test

import (
    `testing`
    `encoding/json`
    `github.com/stretchr/testify/require`

    `github.com/bytedance/sonic`
)

func TestDecodeStringToJsonNumber(t *testing.T) {
	var objs json.Number 
	errs := sonic.UnmarshalString(`"1234"`, &objs)
	var obje json.Number 
	erre := json.Unmarshal([]byte(`"1234"`), &obje)
	require.Equal(t, erre, errs)
	require.Equal(t, obje, objs)

	errs = sonic.UnmarshalString(`"12x4"`, &objs)
	erre = json.Unmarshal([]byte(`"12x4"`), &obje)
	require.Error(t, errs)
	require.Error(t, erre)

	errs = sonic.UnmarshalString(`"1234`, &objs)
	erre = json.Unmarshal([]byte(`"1234`), &obje)
	require.Error(t, errs)
	require.Error(t, erre)

	errs = sonic.UnmarshalString(`1234"`, &objs)
	erre = json.Unmarshal([]byte(`1234"`), &obje)
	require.Error(t, errs)
	require.Error(t, erre)
}