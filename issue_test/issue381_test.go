/**
 * Copyright 2023 ByteDance Inc.
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
    `bytes`
    `encoding/json`
    `strings`
    `testing`

    `github.com/bytedance/sonic`
    `github.com/bytedance/sonic/option`
    _ `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`

    _ `github.com/stretchr/testify/assert`
)

var decoderBufferSize = option.DefaultDecoderBufferSize

func testSonicUnmarshal(t *testing.T) {
    val := strings.Repeat(" ", int(decoderBufferSize-3)) + `{"123":{}}`

    res := make(map[int64]map[string]interface{})
    res2 := make(map[int64]map[string]interface{})

    dec := sonic.ConfigDefault.NewDecoder(bytes.NewBufferString(val))
    err := dec.Decode(&res)
    require.NoError(t, err)

    err2 := json.Unmarshal([]byte(val), &res2)
    require.NoError(t, err2)
    require.Equal(t, res2, res)
}

func TestStreamxx(t *testing.T) {
    testSonicUnmarshal(t)
}