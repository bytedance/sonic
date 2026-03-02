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
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type unknownKeyMap struct {
	X []json.Number        `json:"x"`
	Y map[*int]json.Number `json:"y"`
	Z map[int]func()       `json:"z"`
}

func TestIssue777_NumberSlice(t *testing.T) {
	cas := unmTestCase{
		name: "number slice",
		data: []byte(`["1", "2", 123, 456]`),
		newfn: func() interface{} {
			var a []json.Number
			return &a
		},
	}
	assertUnmarshal(t, sonic.ConfigStd, cas)
}

func TestIssue777_NumberKey(t *testing.T) {
	cas := unmTestCase{
		name: "number key",
		data: []byte(`{"1": "2", "123": 123}`),
		newfn: func() interface{} {
			var a map[json.Number]json.Number
			return &a
		},
	}
	assertUnmarshal(t, sonic.ConfigStd, cas)

	// TODO: encoding/json has a bug
	// cas = unmTestCase {
	// 	name: "number key",
	// 	data: []byte(`{"1": "2", "123x": 123}`),
	// 	newfn: func() interface{} {
	// 		var a map[json.Number]json.Number
	// 		return &a
	// 	},
	// }
	// assertUnmarshal(t, sonic.ConfigStd, cas, true)

}

func TestIssue777_SkipUnsupportedFieldStillDecodeLaterFields(t *testing.T) {
	cases := []unmTestCase{
		{
			name: "skip unsupported field and continue",
			data: []byte(`{"z":{}, "y": {"1": "2", "123": 123}, "x": [1, 2, 3]}`),
			newfn: func() interface{} {
				var a unknownKeyMap
				return &a
			},
		},
		{
			name: "skip unsupported field with null value and continue",
			data: []byte(`{"z": null, "y": {"1": "2", "123": 123}, "x": [1, 2, 3]}`),
			newfn: func() interface{} {
				var a unknownKeyMap
				return &a
			},
		},
		{
			name: "skip unsupported field with nested value and continue",
			data: []byte(`{"z":{"k":[{"m":1},{"n":[2,3,{"p":"q"}]}]}, "y": {"1": "2", "123": 123}, "x": [1, 2, 3]}`),
			newfn: func() interface{} {
				var a unknownKeyMap
				return &a
			},
		},
	}

	for _, cas := range cases {
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}
