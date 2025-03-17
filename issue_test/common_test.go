// Copyright 2025 CloudWeGo Authors
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
	"testing"
	"encoding/json"

	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

type unmTestCase struct {
	name string
	data []byte
	newfn func() interface{}
}

func assertUnmarshal(t *testing.T, api sonic.API, cas unmTestCase, args ...interface{}) {
	sv, jv := cas.newfn(), cas.newfn()
	serr := api.Unmarshal(cas.data, sv)
	jerr := json.Unmarshal(cas.data, jv)
	assert.Equal(t, jv, sv, spew.Sdump(jv, sv))
	assert.Equal(t, serr == nil, jerr == nil, spew.Sdump(jerr, serr))
	if len(args) > 0 && args[0].(bool) {
		spew.Dump(jv, jerr, sv, serr)
	}
}

func assertMarshal(t *testing.T, api sonic.API, obj  interface{}, args ...interface{}) {
	sout, serr := api.Marshal(&obj)
	jout, jerr := json.Marshal(&obj)
	assert.Equal(t, jerr == nil, serr == nil, spew.Sdump(jerr, serr))
	assert.Equal(t, jout, sout, spew.Sdump(jout, sout))
	if len(args) > 0 && args[0].(bool) {
		spew.Dump(string(jout), jerr)
	}
}
