/**
* Copyright 2023 ByteDance Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express orimplied.
* See the License for the specific language governing permissionsand
* limitations under the License.
*/
package issue_test
import (
    `encoding/json`
    `testing`
    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func TestIssue460_UnmarshalMaxFloat32(t *testing.T) {
    tests := []string {
        // max.MaxFloat32
        "3.40282346638528859811704183484516925440e+38",

        // round up to max.MaxFloat32
        "3.402823469e+38",
        "3.40282346e+38",
        "3.40282347e+38",
        "3.40282348e+38",
        "3.4028235e+38",
        // TODO: fix this case
        // "3.4028235677973366e+38",  // Bits: 1000111111011111111111111111111111_10000000000000000000000000000

        // overflow for float32, round up to max.MaxFloat32 + 1
        "3.402823567797337e+38", // Bits: 1000111111011111111111111111111111_10000000000000000000000000001
        "3.402823567797338e+38",
        "3.4028236e+38",
    }

    t.Run("max float32", func(t *testing.T) {
        for _, data := range(tests) {
            var f1, f2  float32
            se := sonic.UnmarshalString(data, &f1)
            je := json.Unmarshal([]byte(data), &f2)
            require.Equal(t, se != nil, je != nil, data, se, je)
            require.Equal(t, f2, f1, data)
        }
    })

    t.Run("min float32", func(t *testing.T) {
        for _, data := range(tests) {
            data = string("-") + data
            var f1, f2  float32
            se := sonic.UnmarshalString(data, &f1)
            je := json.Unmarshal([]byte(data), &f2)
            require.Equal(t, se != nil, je != nil, data, se, je)
            require.Equal(t, f2, f1, data)
        }
    })
}
