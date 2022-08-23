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

package issue_test

import (
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/require`
)

type Issue113_OmitemptyOpt struct {
    Bool      bool                `json:"bool,omitempty"`
    Int       int                 `json:"int,omitempty"`
    Int8      int8                `json:"int8,omitempty"`
    Int16     int16               `json:"int16,omitempty"`
    Int32     int32               `json:"int32,omitempty"`
    Int64     int64               `json:"int64,omitempty"`
    Uint      uint                `json:"uint,omitempty"`
    Uint8     uint8               `json:"uint8,omitempty"`
    Uint16    uint16              `json:"uint16,omitempty"`
    Uint32    uint32              `json:"uint32,omitempty"`
    Uint64    uint64              `json:"uint64,omitempty"`
    Float32   float32             `json:"float32,omitempty"`
    Float64   float64             `json:"float64,omitempty"`
    Uintptr   uintptr             `json:"uintptr,omitempty"`
    String    string              `json:"string,omitempty"`
    Array0    [0]uint             `json:"array0,omitempty"`
    Array     [2]int              `json:"array,omitempty"`
    Interface interface{}         `json:"interface,omitempty"`
    Map0      map[int]interface{} `json:"map0,omitempty"`
    Map       map[string]float64  `json:"map,omitempty"`
    Slice0    []int               `json:"slice0,omitempty"`
    Slice     []byte              `json:"slice,omitempty"`
    Ptr       * Issue113_Inner    `json:"ptr,omitempty"`
    Struct1   Issue113_Inner      `json:"struct1,omitempty"`
    Struct2   struct{}            `json:"struct2,omitempty"`
}

type Issue113_Inner struct {
    S string `json:"s,"`
    So string `json:"so,omitempty"`
}

var issue13ExpectedEmptyOpt = `{
 "array": [
  0,
  0
 ],
 "struct1": {
  "s": ""
 },
 "struct2": {}
}`

var issue13ExpectedNonemptyOpt = `{
 "bool": true,
 "int": 1,
 "int8": -1,
 "int16": 1,
 "int32": 2,
 "int64": 64,
 "uint": 1,
 "uint8": 8,
 "uint16": 16,
 "uint32": 32,
 "uint64": 64,
 "float32": 1,
 "float64": -2.34e+64,
 "uintptr": 1,
 "string": "string",
 "array": [
  0,
  -1
 ],
 "interface": {
  "s": "not omit"
 },
 "map0": {
  "0": "zero"
 },
 "map": {
  "key": 0
 },
 "slice0": [
  0
 ],
 "slice": "Yg==",
 "ptr": {
  "s": "not omit"
 },
 "struct1": {
  "s": "not omit"
 },
 "struct2": {}
}`

func TestIssue113_MarshalEmptyFieldsWithOmitemptyOpt(t *testing.T) {
    var obj Issue113_OmitemptyOpt
    obj.Slice0 = make([]int, 0, 100)     // empty slice
    obj.Map0 = make(map[int]interface{}) // empty map

    got, err := encoder.EncodeIndented(&obj, "", " ", 0)

    require.NoError(t, err)
    require.Equal(t, issue13ExpectedEmptyOpt, string(got))
}

func TestIssue113_MarshalNonemptyFieldsWithOmitemptyOpt(t *testing.T) {
    var inner = & Issue113_Inner {
        S : "not omit",
    }

    var obj = & Issue113_OmitemptyOpt{
        Bool      : true,
        Int       : 1,
        Int8      : -1,
        Int16     : 1,
        Int32     : 2,
        Int64     : 64,
        Uint      : 1,
        Uint8     : 8,
        Uint16    : 16,
        Uint32    : 32,
        Uint64    : 64,
        Float32   : 1.0,
        Float64   : -2.34e+64,
        Uintptr   : uintptr(0x1),
        String    : "string", 
        Array0    : [0]uint{},
        Array     : [2]int{0, -1},
        Interface : *inner,
        Map0      : map[int]interface{}{0 : "zero"},
        Map       : map[string]float64{"key" : 0.0},
        Slice0    : make([]int, 1, 1),
        Slice     : []byte("b"),
        Ptr       : inner,
        Struct1   : *inner,
        Struct2   : struct{}{},
    }

    got, err := encoder.EncodeIndented(&obj, "", " ", 0)

    require.NoError(t, err)
    require.Equal(t, issue13ExpectedNonemptyOpt, string(got))
}