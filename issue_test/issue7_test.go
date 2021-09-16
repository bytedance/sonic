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

    `github.com/stretchr/testify/require`
)

func TestIssue7(t *testing.T) {
    v := &[...]int{1, 2, 3, 4, 5, 6, 7}
    err := Unmarshal([]byte(`[3]`), v)
    require.Nil(t, err)
    require.Equal(t, &[...]int{3, 0, 0, 0, 0, 0, 0}, v)
}
