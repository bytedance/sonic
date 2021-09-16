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
    . `github.com/bytedance/sonic`
    `testing`

    `github.com/bytedance/sonic/decoder`
    `github.com/stretchr/testify/require`
)

type Issue82String string

func (s *Issue82String) UnmarshalJSON(b []byte) error {
    *s = Issue82String(b)
    return nil
}

func TestIssue82_MapValueIsStringUnmarshaler(t *testing.T) {
    var v map[string]Issue82String
    err := Unmarshal([]byte(`{"a":123}`), &v)
    if err != nil {
        println(err.(decoder.SyntaxError).Description())
        require.NoError(t, err)
    }
    require.Equal(t, map[string]Issue82String{"a": "123"}, v)
}
