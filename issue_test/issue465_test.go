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
    `testing`

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func TestIssue465(t *testing.T) {
    data := `
{
    "data": {
    "1": "",
    "2": true,
    "3": "",
    "4": "",
    "5": {
        "a": {},
        "b": "hello",
        "c": ""
    },
    "6": "",
    "7": false,
    "8": true,
    "9": 7,
    "10": "",
    "11": {},
    "12": 0,
    "13": {},
    "14": "",
    "15": "",
    "16": true,
    "17": {"d":{"e":1}}
    }
}`
    root, _ := sonic.GetFromString(data)
    sp := root.GetByPath("data", "5")
    node := sp.Get("a")
    require.True(t, node.Exists())
    root.GetByPath("data", "17")
    node2 := sp.Get("b")
    get2, err2 := node2.String()
    require.Nil(t, err2)
    require.Equal(t, get2, "hello")
    b, err := root.MarshalJSON()
    require.NoError(t, err)
    buf := bytes.NewBuffer(nil)
    json.Compact(buf, b)
    require.Equal(t, buf.String(), string(b))
}
