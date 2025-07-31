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
		A *int `xx:"yy"`
		B map[int]struct {
			C *int `json:"c" xx:"yy"`
		}
		D []struct {
			E *int `json:"e"`
		}
		F *struct {
			G *int `json:"g"`
		}
		H struct {
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
		name string
		args args
		want string
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
			tr := NewTagChanger(ReplacerAddOmitempty)
			oldObj := caseStruct{}
			oldJs, err := json.Marshal(oldObj)
			if err != nil {
				t.Errorf("TagChanger.ReplaceTag() error = %v", err)
				return
			}
			fmt.Println(string(oldJs))
			got := tr.Replace(reflect.TypeOf(tt.args.old))
			gotObj := reflect.NewAt(got, unsafe.Pointer(&oldObj)).Interface()
			js, err := json.Marshal(gotObj)
			println(string(js))
			if err != nil {
				t.Errorf("TagChanger.ReplaceTag() error = %v", err)
				return
			} else if string(js) != tt.want {
				t.Errorf("TagChanger.ReplaceTag() = %v, want %v", string(js), tt.want)
			}
		})
	}
}

func TestTagChanger_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				v: &struct {
					A *int `json:"a"`
					B *string 
				}{},
			},
			want:    []byte(`{}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTagChanger(ReplacerAddOmitempty)
			got, err := tr.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagChanger.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagChanger.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagChanger_Unmarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	type testStruct struct {
		A int64 `json:"a"`
		B uint64 
		C *int64
	}
	vc := int64(3)
	tests := []struct {
		name    string
		args    args
		js      []byte
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				v: &testStruct{},
			},
			js:      []byte(`{"a":"1","b":"2","c":"3"}`),
			want:    &testStruct{
				A: 1,
				B: 2,
				C: &vc,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTagChanger(ReplaceerAddStringInt64)
			err := tr.Unmarshal(tt.js, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagChanger.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("TagChanger.Marshal() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
