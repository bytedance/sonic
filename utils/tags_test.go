/**
 * Copyright 2025 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestTagChanger_ReplaceTag(t *testing.T) {

	type caseStruct struct {
		A *int 
		B map[int]struct{
			C *int `json:"c"`
		} 
		D []struct{
			E *int `json:"e"`
		}
		F *struct{
			G *int `json:"g"`
		}
		H struct{
			I *int `json:"i"`
		}
	}


	type fields struct {
		types    map[reflect.Type]reflect.Type
		replacer func(field reflect.StructField, tag reflect.StructTag) reflect.StructTag
	}
	type args struct {
		old interface{}
	}
	tests := []struct {
		name   string
		args   args
		want   string
	}{
		{
			name: "test1",
			args: args{
				old: caseStruct{},
			},
			want: `{"H":{}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTagChanger(ReplacerAddOmitemptyfunc)
			oldObj := caseStruct{}
			oldJs, err := json.Marshal(oldObj)
			if err!= nil {
				t.Errorf("TagChanger.ReplaceTag() error = %v", err)
				return
			}
			fmt.Println(string(oldJs))
			got := tr.ReplaceTag(reflect.TypeOf(tt.args.old))
			gotObj := reflect.NewAt(got, unsafe.Pointer(&oldObj)).Interface()
			js, err := json.Marshal(gotObj)
			println(string(js))
			if err!= nil {
				t.Errorf("TagChanger.ReplaceTag() error = %v", err)
				return
			} else if string(js)!= tt.want {
				t.Errorf("TagChanger.ReplaceTag() = %v, want %v", string(js), tt.want)
			}
		})
	}
}
