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
    `testing`
    `encoding/json`
    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

type Foo struct {
    Name string
}

func (f *Foo) UnmarshalJSON(data []byte) error {
    f.Name = "Unmarshaler"
    return nil
}

type MyPtr *Foo

func TestIssue379(t *testing.T) {
    tests := []struct{
        data  string
        newf  func() interface{} 
    } {
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { return &Foo{} },
        },
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { return MyPtr(&Foo{}) },
        },
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
        },
        {
            data: `null`,
            newf:  func() interface{} { return MyPtr(&Foo{}) },
        },
        {
            data: `null`,
            newf:  func() interface{} { return &Foo{} },
        },
        {
            data: `null`,
            newf:  func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]MyPtr) },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]*Foo) },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]*MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]*MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]*Foo) },
        },
    }

    for _, tt := range tests {
        jv, sv := tt.newf(), tt.newf()
        jerr := json.Unmarshal([]byte(tt.data), jv)
        serr := sonic.Unmarshal([]byte(tt.data), sv)
        require.Equal(t, jv, sv)
        require.Equal(t, jerr, serr)
    }
}
