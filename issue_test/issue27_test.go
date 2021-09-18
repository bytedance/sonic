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
    `reflect`
    `sync`
    `testing`

    `github.com/stretchr/testify/require`
)

func TestIssue27_EmptySlice(t *testing.T) {
    _ = Pretouch(reflect.TypeOf([]*int{}))
    wg := sync.WaitGroup{}
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            v := make([]*int, 16)
            err := UnmarshalString(`[null, null]`, &v)
            require.NoError(t, err)
            require.Equal(t, []*int{nil, nil}, v)
        }()
    }
    wg.Wait()
}
