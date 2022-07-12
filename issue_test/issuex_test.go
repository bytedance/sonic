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

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func TestValidStringObject(t *testing.T) { // Not OK
    data := `{
        }`
        var js1, js2 interface{}
        jerr1 := json.Unmarshal([]byte(data), &js1)
        jerr2 := sonic.Unmarshal([]byte(data), &js2)
        require.Equal(t, jerr1 == nil, jerr2 == nil)
}

func TestValidStringObject2(t *testing.T) { // OK
    data := `{   }`
        var js1, js2 interface{}
        jerr1 := json.Unmarshal([]byte(data), &js1)
        jerr2 := sonic.Unmarshal([]byte(data), &js2)
        require.Equal(t, jerr1 == nil, jerr2 == nil)
}