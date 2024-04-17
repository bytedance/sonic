// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unit_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	var obj = new(Context)
	out, err := sonic.Marshal(obj)
	out1, err1 := jsoniter.Marshal(obj)
	out2, err2 := json.Marshal(obj)
	println(string(out2))
	assert.Equal(t, err2 ==nil, err1 ==nil)
	assert.Equal(t, err2 ==nil, err ==nil)
	assert.Equal(t, out2, out1)
	assert.Equal(t, out2, out)
	// err = sonic.Unmarshal(out, obj)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

type Context struct {
	*http.Request
}

