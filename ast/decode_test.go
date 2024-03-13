/*
 * Copyright 2022 ByteDance Inc.
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

package ast

import (
	"testing"
	"unicode/utf8"

	"github.com/bytedance/sonic/internal/rt"
)

func Test_DecodeString(t *testing.T) {
	type args struct {
		src      string
		pos      int
		needEsc  bool
		validStr bool
	}
	invalidstr := rt.Mem2Str([]byte{'"',193,255,'"'})
	println(utf8.ValidString(invalidstr))

	tests := []struct {
		name       string
		args       args
		wantV      string
		wantRet    int
		wantHasEsc bool
	}{
		{"empty", args{`""`, 0, false, false}, "", 2, false},
		{"one", args{`"1"`, 0, false, false}, "1", 3, false},
		{"escape", args{`"\\"`, 0, false, false}, `\\`, 4, true},
		{"escape", args{`"\\"`, 0, true, true}, `\`, 4, true},
		{"uft8", args{`"\u263a"`, 0, false, false}, `\u263a`, 8, true},
		{"uft8", args{`"\u263a"`, 0, true, true}, `☺`, 8, true},
		{"invalid uft8", args{`"\xx"`, 0, false, false}, `\xx`, 5, true},
		{"invalid escape", args{`"\xx"`, 0, false, true}, `\xx`, 5, true},
		{"invalid escape", args{`"\xx"`, 0, true, true}, ``, -3, true},
		{"invalid string", args{invalidstr, 0, false, false}, invalidstr[1:3], 4, false},
		{"invalid string", args{invalidstr, 0, true, true}, "", -10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotRet, gotHasEsc := _DecodeString(tt.args.src, tt.args.pos, tt.args.needEsc, tt.args.validStr)
			if gotV != tt.wantV {
				t.Errorf("_DecodeString() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotRet != tt.wantRet {
				t.Errorf("_DecodeString() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
			if gotHasEsc != tt.wantHasEsc {
				t.Errorf("_DecodeString() gotHasEsc = %v, want %v", gotHasEsc, tt.wantHasEsc)
			}
		})
	}
}
