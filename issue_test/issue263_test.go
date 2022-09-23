/*
 * Copyright 2022 ByteDance Inc.
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
    `strings`
    `testing`

    `github.com/bytedance/sonic/decoder`
    `github.com/stretchr/testify/require`
)

type Response struct {
    Menu Menu `json:"menu"`
}

type Menu struct {
    Items []*Item `json:"items"`
}

type Item struct {
    ID string `json:"id"`
}

func (i *Item) UnmarshalJSON(buf []byte) error {    
	return nil
}

func TestIssue263(t *testing.T) {
    q := `{
		"menu": {
			"items": [
				{`+strings.Repeat(" ", 1024)+`}
			]
		}
	}`

    var response Response
    require.Nil(t, decoder.NewStreamDecoder(bytes.NewReader([]byte(q))).Decode(&response))

    q = `{
        "menu": {
            "items": [
                {"a":"`+strings.Repeat("b", 2048)+`"}
            ]
        }
    }`
    
    require.Nil(t, decoder.NewStreamDecoder(bytes.NewReader([]byte(q))).Decode(&response))
}