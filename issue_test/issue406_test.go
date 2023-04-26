/*
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
    `testing`
    `encoding/json`
    `github.com/stretchr/testify/require`
    `github.com/bytedance/sonic`
)

type FooId struct {
    Id   int    `json:"id"`
}

func TestUnmarshalErrorInMapSlice(t *testing.T) {
    var a, b map[string][]FooId
    mapdata := `{"ptrslice": [{"id": "1"}, {"id": "2"}, {"id": "3"}, {"id": "4"}]}`
    se := json.Unmarshal([]byte(mapdata), &a)
    je := sonic.Unmarshal([]byte(mapdata), &b)
    require.Equal(t, se == nil,  je == nil);
    require.Equal(t, a, b);
}

func TestUnmarshalErrorInSlice(t *testing.T) {
   var a, b []*FooId
   slicedata := `[{"id": "1"}, {"id": "2"}, {"id": "3"}, {"id": 4}]`
   je := json.Unmarshal([]byte(slicedata), &a)
   se := sonic.Unmarshal([]byte(slicedata), &b)
   require.Equal(t, se == nil,  je == nil);
   require.Equal(t, a, b);
}
