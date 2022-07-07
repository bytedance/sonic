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

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func TestIssue248(t *testing.T) {
    var d = `{}`
    n, err := sonic.GetFromString(d)
    require.Nil(t, err)
    v := n.GetByPath("a", 0)
    require.Nil(t, v)
    s, err := v.Get("a").Index(0).String()
    require.NotNil(t, err)
    require.Equal(t, "", s)
}