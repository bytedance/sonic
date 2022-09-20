
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
    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func TestMarshal_Float32To64(t *testing.T) {
	var f float32 = 0.1
	oe,ee := json.Marshal(f)
	os,es := sonic.Marshal(f)
	require.Equal(t, ee == nil, es == nil)
	require.Equal(t, string(oe), string(os))

	var f2,f3 float64
	require.Nil(t, json.Unmarshal(oe, &f2))
	require.Nil(t, sonic.Unmarshal(os, &f3))
	require.Equal(t, f2, f3)
}