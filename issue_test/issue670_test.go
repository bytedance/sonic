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
	"fmt"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestIssue670_encode(t *testing.T) {
	var obj = Issue670Case{ D: Date(time.Now().Unix()) }
	so, _ := sonic.MarshalString(obj)
	eo, _ := json.Marshal(obj)
	assert.Equal(t, string(eo), so)
	println(string(eo))
}

func TestIssue670_decode(t *testing.T) {
	// match
	eo := []byte(`{"D":"2021-08-26","E":1}`)
	testUnmarshal(t, eo)

	// mismatch
	eo = []byte(`{"D":11,"E":1}`)
	testUnmarshal(t, eo)

	// null
	eo = []byte(`{"D":null,"E":1}`)
	testUnmarshal(t, eo)
}

func testUnmarshal(t *testing.T, eo []byte) {
	obj := Issue670Case{}
	println(string(eo))
	println("sonic")
	es := sonic.Unmarshal(eo, &obj)
	obj2 := Issue670Case{}
	println("std")
	ee := json.Unmarshal(eo, &obj2)
	assert.Equal(t, ee ==nil, es == nil, es)
	assert.Equal(t, obj2, obj)
	fmt.Printf("std: %v, obj: %#v", ee, obj2)
	fmt.Printf("sonic error: %v, obj: %#v", es, obj)
}

type Issue670Case struct {
	D      Date                 `form:"D" json:"D,string" query:"D"`
	E      int
}

type Date int64

func (d Date) MarshalJSON() ([]byte, error) {
	if d == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Unix(int64(d), 0).Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(in []byte) error {
	if string(in) == "null" {
		*d = 0
		return nil
	}
	println("hook ", string(in))
	t, err := time.Parse("2006-01-02", string(in))
	if err != nil {
		return err
	}
	*d = Date(t.Unix())
	return nil
}
