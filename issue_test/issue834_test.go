/**
 * Copyright 2025 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 
package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue834_UnmarshalSingleEscapedCharIntoStringOption(t *testing.T) {
	type SS struct {
		S string `json:",string"`
	}

	for _, cas := range []unmTestCase{
		 {
			name: "double escaped",
			data: []byte(`{"S":"\"\\u003c\\u003e\\u0026\\u2028\\u2029\""}`),
			newfn: func() interface{} {
				return new(SS)
			},
		},
		{
			name: "single escaped",
			data: []byte(`{"S":"\"\u003c\u003e\u0026\u2028\u2029\""}`),
			newfn: func() interface{} {
				return new(SS)
			},
		},
		{
			name: "without quotes",
			data: []byte(`{"S":"\u003c\u003e\u0026\u2028\u2029"}`),
			newfn: func() interface{} {
				return new(SS)
			},
		},
	} {
		assertUnmarshal(t, sonic.ConfigDefault, cas)
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}
