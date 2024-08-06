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
	"reflect"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestIssue670_JSONMarshaler(t *testing.T) {
	var obj = Issue670JSONMarshaler{ D: Date(time.Now().Unix()) }
	so, _ := sonic.MarshalString(obj)
	eo, _ := json.Marshal(obj)
	assert.Equal(t, string(eo), so)
	println(string(eo))
}

func TestIssue670_JSONUnmarshaler(t *testing.T) {
	// match
	eo := []byte(`{"D":"2021-08-26","E":1}`)
	et := reflect.TypeOf(Issue670JSONMarshaler{})
	testUnmarshal(t, eo, et, true)

	// mismatch
	eo = []byte(`{"D":11,"E":1}`)
	testUnmarshal(t, eo, et, true)

	// null
	eo = []byte(`{"D":null,"E":1}`)
	testUnmarshal(t, eo, et, true)
}

func testUnmarshal(t *testing.T, eo []byte, rt reflect.Type, checkobj bool) {
	obj := reflect.New(rt).Interface()
	println(string(eo))
	println("sonic")
	es := sonic.Unmarshal(eo, obj)
	obj2 := reflect.New(rt).Interface()
	println("std")
	ee := json.Unmarshal(eo, obj2)
	assert.Equal(t, ee ==nil, es == nil, es)
	if checkobj {
		assert.Equal(t, obj2, obj)
	}
	fmt.Printf("std: %v, obj: %#v", ee, obj2)
	fmt.Printf("sonic error: %v, obj: %#v", es, obj)
}

func TestIssue670_TextMarshaler(t *testing.T) {
	var obj = Issue670TextMarshaler{ D: int(time.Now().Unix()) }
	so, _ := sonic.MarshalString(obj)
	eo, _ := json.Marshal(obj)
	assert.Equal(t, string(eo), so)
	println(string(eo))
}

func TestIssue670_TextUnmarshaler(t *testing.T) {
	// match
	eo := []byte(`{"D":"2021-08-26","E":1}`)
	et := reflect.TypeOf(Issue670TextMarshaler{})
	testUnmarshal(t, eo, et, false)

	// mismatch
	eo = []byte(`{"D":11,"E":1}`)
	testUnmarshal(t, eo, et, false)

	// null
	eo = []byte(`{"D":null,"E":1}`)
	testUnmarshal(t, eo, et, true)
}

type Issue670JSONMarshaler struct {
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

type Issue670TextMarshaler struct {
	D      int                 `form:"D" json:"D,string" query:"D"`
	E      int
}


type Date2 int64

func (d Date2) MarshalText() ([]byte, error) {
	println("hook 1")
	if d == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Unix(int64(d), 0).Format("2006-01-02"))), nil
}

func (d *Date2) UnmarshalText(in []byte) error {
	println("hook 2", string(in))
	if string(in) == "null" {
		*d = 0
		return nil
	}
	t, err := time.Parse("2006-01-02", string(in))
	if err != nil {
		return err
	}
	*d = Date2(t.Unix())
	return nil
}