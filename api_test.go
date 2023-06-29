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
    `testing`

    `github.com/stretchr/testify/require`
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
