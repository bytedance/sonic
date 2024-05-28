/*
 * Copyright 2021 ByteDance Inc.
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

package decoder

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

type errTest struct {
	in	string
	ptr interface{}
	pos int
}


func TestErrors_ParseError(t *testing.T) {
	testCases := []errTest {
		{
			in: `{123}`,
			pos: 1,
		},
		{
			in: `tru`,
			pos: 0,
		},
		{
			in: ` fx`,
			pos: 1,
		},
		{
			in: `{"12" 12}`,
			pos: 6,
		},
	}

	for _, tt := range testCases {
		var v1, v2 interface{}
		got := NewDecoder(tt.in).Decode(&v1)
		exp := json.Unmarshal([]byte(tt.in), &v2)
		assert.Error(t, exp)
		e := got.(SyntaxError); 
		assert.Equal(t, tt.pos, e.Pos, tt)
		println(e.Description())
	}
}

type A struct {
	A string
}

type B struct {
	A int `json:"a,string"`
}

func TestErrors_MismatchType(t *testing.T) {
	testCases := []errTest {
		{
			in: `{"a": 123}`,
			ptr: &A{},
			pos: 6,
		},
		{
			in: ` {"a": true}`,
			ptr: &A{},
			pos: 7,
		},
		{
			in: ` {"a": true}`,
			ptr: &B{},
			pos: 7,
		},
		// {
		// 	in: ` {"a": "true"}`,
		// 	ptr: &B{},
		// 	pos: 7,
		// },
		{
			in: ` [1, 2, "3", 4]`,
			ptr: &[4]int{},
			pos: 9,
		},
		{
			in: ` [1, 2, "3", 4]`,
			ptr: &[]int{},
			pos: 9,
		},
		{
			in: ` [1, 256, "3", 4]`,
			ptr: &[]int8{},
			pos: 5,
		},
		{
			in: ` [1, 256, "3", 4]`,
			ptr: &[]byte{}, // []byte is special
			pos: 1,
		},
		{
			in: ` {"key": 123}`,
			ptr: &map[string]string{},
			pos: 9,
		},
		{
			in: ` {"key": 123}`,
			ptr: &map[int64]interface{}{},
			pos: 3,
		},
		{
			in: ` "key"`,
			ptr: new(json.Number),
			pos: 2,
		},
	}

	for _, tt := range testCases {
		spew.Dump(tt)
		got := NewDecoder(tt.in).Decode(tt.ptr)
		e := got.(*MismatchTypeError); 
		assert.Equal(t, tt.pos, e.Pos)
		println(e.Description())

		exp := json.Unmarshal([]byte(tt.in), tt.ptr)
		assert.Error(t, exp)
	}
}

func TestErrors_ParseMultiJsonError(t *testing.T) {
	testCases := []errTest {
		{
			in: ` {"a":"b"} {"1":"2"}  true    false   null 1.23 0 -1 1e123 456 "hello" "" "\\" "\"" fx`,
			pos: 84,
		},
	}

	for _, tt := range testCases {
		dec := NewDecoder(tt.in)
		var val interface{}
		var err error
		for err == nil {
			err = dec.Decode(&val)
			spew.Dump(val)
		}

		e := err.(SyntaxError); 
		assert.Equal(t, tt.pos, e.Pos)
		println(e.Description())

		jdec := json.NewDecoder(strings.NewReader(tt.in))
		var jval interface{}
		var jerr error
		for jerr == nil {
			jerr = jdec.Decode(&jval)
		}
		println(jerr.Error())
	}
}
