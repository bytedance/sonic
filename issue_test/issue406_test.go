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
	"testing"
    "encoding/json"
    `github.com/davecgh/go-spew/spew`
	`github.com/stretchr/testify/require`

	"github.com/bytedance/sonic"
)
var mapdata = `{
      "ptrslice": [{"id": "1"}, {"id": "2"}, {"id": "3"}, {"id": "4"}]
  }`;

  
type FooId struct {
	Id   int    `json:"id"`
}

func TestUnmarshalErrorInMapSlice(t *testing.T) {
 	var a, b map[string][]FooId
 	se := json.Unmarshal([]byte(mapdata), &a)
	je := sonic.Unmarshal([]byte(mapdata), &b)
    spew.Dump(se, a) // len(a) = 4
    spew.Dump(je, b) // len(b) = 1
	require.Equal(t, a, b);
}

var slicedata = `[{"id": "1"}, {"id": "2"}, {"id": "3"}, {"id": 4}]`;

func TestUnmarshalErrorInSlice(t *testing.T) {
   var a, b []*FooId
   je := json.Unmarshal([]byte(slicedata), &a)
   se := sonic.Unmarshal([]byte(slicedata), &b)
   spew.Dump("sonic ", se, b)
   spew.Dump("json ", je, a)
   require.Equal(t, a, b);
}
