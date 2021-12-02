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
    `encoding/json`
    `testing`

    `github.com/stretchr/testify/require`
    `github.com/davecgh/go-spew/spew`
    `github.com/bytedance/sonic`
)

type Issue144_StringOption struct {
    S1 *string      `json:"s1,string"`
    S2 *string      `json:"s2,string"`
    S3 string       `json:"s3,string"`
    J1 json.Number  `json:"j1,string"`
    J2 *json.Number `json:"j2,string"`
    J3 *json.Number `json:"j3,string"`
    I1 int          `json:"i1,string"`
    I2 *int         `json:"i2,string"`
    I3 *int         `json:"i3,string"`
}

func TestIssue144_StringOption(t *testing.T) {
    data := []byte(`{
        "s1":"\"null\"",
        "s2":"null",
        "s3":"null",
        "j1":"null",
        "j2":"null",
        "j3":"123.456",
        "i1":"null",
        "i2":"null",
        "i3":"-123"
    }`)

    var v1, v2 Issue144_StringOption
    e1 := json.Unmarshal(data, &v1)
    e2 := sonic.Unmarshal(data, &v2)
    require.NoError(t, e1)
    require.NoError(t, e2)
    require.Equal(t, v1, v2)
    spew.Dump(v1)

    i, j, s := int(1), json.Number("1"), "null"
    v1.I2, v2.I2 = &i, &i
    v1.J2, v2.J2 = &j, &j
    v1.S2, v2.S2 = &s, &s

    e1 = json.Unmarshal(data, &v1)
    e2 = sonic.Unmarshal(data, &v2)
    require.NoError(t, e1)
    require.NoError(t, e2)
    require.Equal(t, v1, v2)
    spew.Dump(v1)
}