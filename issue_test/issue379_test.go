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
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

type Foo struct {
	Name string
}

func (f *Foo) UnmarshalJSON(data []byte) error {
	println("UnmarshalJSON called!!!")
	f.Name = "Unmarshaler"
	return nil
}

type MyPtr *Foo

func TestIssue379(t *testing.T) {
	tests := []struct {
		data     string
		newf     func() interface{}
		wantf    func() interface{}
		wantJSON string
	}{
		{
			data:     `{"Name":"MyPtr"}`,
			newf:     func() interface{} { return &Foo{} },
			wantf:    func() interface{} { return &Foo{Name: "Unmarshaler"} },
			wantJSON: `{"Name":"Unmarshaler"}`,
		},
		{
			data:     `{"Name":"MyPtr"}`,
			newf:     func() interface{} { ptr := &Foo{}; return &ptr },
			wantf:    func() interface{} { ptr := &Foo{Name: "Unmarshaler"}; return &ptr },
			wantJSON: `{"Name":"Unmarshaler"}`,
		},
		{
			data:     `{"Name":"MyPtr"}`,
			newf:     func() interface{} { return MyPtr(&Foo{}) },
			wantf:    func() interface{} { return MyPtr(&Foo{Name: "MyPtr"}) },
			wantJSON: `{"Name":"MyPtr"}`,
		},
		{
			data:     `{"Name":"MyPtr"}`,
			newf:     func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
			wantf:    func() interface{} { ptr := MyPtr(&Foo{Name: "MyPtr"}); return &ptr },
			wantJSON: `{"Name":"MyPtr"}`,
		},

		// TODO: fix jit tests
		// {
		//     data: `null`,
		//     newf:  func() interface{} { return MyPtr(&Foo{}) },
		// },
		// {
		//     data: `null`,
		//     newf:  func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
		// },
		// {
		//     data: `null`,
		//     newf:  func() interface{} {
		//         x :=  &Foo{Name: "mock"}
		//         return x },
		// },
		// {
		//     data: `null`,
		//     newf:  func() interface{} { ptr := &Foo{}; return &ptr },
		// },
		{
			data: `{"map":{"Name":"MyPtr"}}`,
			newf: func() interface{} { return new(map[string]MyPtr) },
			wantf: func() interface{} {
				value := MyPtr(&Foo{Name: "MyPtr"})
				return &map[string]MyPtr{"map": value}
			},
			wantJSON: `{"map":{"Name":"MyPtr"}}`,
		},
		{
			data: `{"map":{"Name":"MyPtr"}}`,
			newf: func() interface{} { return new(map[string]*Foo) },
			wantf: func() interface{} {
				return &map[string]*Foo{"map": {Name: "Unmarshaler"}}
			},
			wantJSON: `{"map":{"Name":"Unmarshaler"}}`,
		},
		{
			data: `{"map":{"Name":"MyPtr"}}`,
			newf: func() interface{} { return new(map[string]*MyPtr) },
			wantf: func() interface{} {
				value := MyPtr(&Foo{Name: "MyPtr"})
				return &map[string]*MyPtr{"map": &value}
			},
			wantJSON: `{"map":{"Name":"MyPtr"}}`,
		},
		{
			data: `[{"Name":"MyPtr"}]`,
			newf: func() interface{} { return new([]MyPtr) },
			wantf: func() interface{} {
				value := MyPtr(&Foo{Name: "MyPtr"})
				return &[]MyPtr{value}
			},
			wantJSON: `[{"Name":"MyPtr"}]`,
		},
		{
			data: `[{"Name":"MyPtr"}]`,
			newf: func() interface{} { return new([]*MyPtr) },
			wantf: func() interface{} {
				value := MyPtr(&Foo{Name: "MyPtr"})
				return &[]*MyPtr{&value}
			},
			wantJSON: `[{"Name":"MyPtr"}]`,
		},
		{
			data: `[{"Name":"MyPtr"}]`,
			newf: func() interface{} { return new([]*Foo) },
			wantf: func() interface{} {
				return &[]*Foo{{Name: "Unmarshaler"}}
			},
			wantJSON: `[{"Name":"Unmarshaler"}]`,
		},
	}

	for i, tt := range tests {
		println(i)
		sv := tt.newf()
		serr := sonic.Unmarshal([]byte(tt.data), sv)
		require.NoError(t, serr)
		if stdUsesJSONV2 {
			require.Equal(t, tt.wantf(), sv)
			requireSonicValueSnapshot(t, sonic.ConfigDefault, sv, false, tt.wantJSON)
		}
		if !stdUsesJSONV2 {
			jv := tt.newf()
			jerr := json.Unmarshal([]byte(tt.data), jv)
			require.Equal(t, jv, sv)
			require.Equal(t, jerr, serr)
		}

		sv = tt.newf()
		serr = sonic.Unmarshal([]byte(tt.data), &sv)
		require.NoError(t, serr)

		if stdUsesJSONV2 {
			require.Equal(t, tt.wantf(), sv)
			requireSonicValueSnapshot(t, sonic.ConfigDefault, sv, false, tt.wantJSON)
			continue
		}

		jv := tt.newf()
		jerr := json.Unmarshal([]byte(tt.data), &jv)
		require.Equal(t, jv, sv, spew.Sdump(jv, sv))
		require.Equal(t, jerr, serr)
	}
}
