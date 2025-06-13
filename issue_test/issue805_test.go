/*
* Copyright 2025 ByteDance Inc.
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
	"errors"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

type Foo805 struct {}

func (d *Foo805) UnmarshalJSON(data []byte) error {
	end := data[len(data) - 1]
	if  end == '\r' || end == '\n' || end == ' ' || end == '\t' {
		return errors.New("has trailing space")
	}
	return nil
}

type Foo805Wrapper struct {
	D Foo805 `json:"d"`
}

func TestIssue805_TrailingCharsInUnmarshaler(t *testing.T) {
	data := []byte(`{
		"d": 5.56030000
	}`)

	var sv, jv Foo805Wrapper
	serr := sonic.Unmarshal(data, &sv)
	jerr := json.Unmarshal(data, &jv)
	require.Equal(t, serr, jerr)
	require.Equal(t, sv, jv)
}
