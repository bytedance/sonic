// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)


type A1 struct {
	A *A1 `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
	D D `json:"d"`
}

type D int

func (d *D) UnmarshalJSON(data []byte) error {
	return nil
}

func TestMismatchErrorInRecusive(t *testing.T) {
	data := `{"a":{"a": null, "b": "123"}, "c": 123, "d": {}}`
	var obj1, obj2 A1
	es := sonic.Unmarshal([]byte(data), &obj1)
	ee := json.Unmarshal([]byte(data), &obj2)
	assert.Equal(t, ee ==nil, es == nil, es)
	assert.Equal(t, obj1, obj2)
	println(es.Error())
}