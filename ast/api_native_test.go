//go:build (amd64 && go1.16 && !go1.23) || (arm64 && go1.20 && !go1.23)
// +build amd64,go1.16,!go1.23 arm64,go1.20,!go1.23

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
    `encoding/json`
    `fmt`
    `reflect`
    `runtime`
    `runtime/debug`
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/require`
)

func TestSortNodeTwitter(t *testing.T) {
    if encoder.EnableFallback {
        return
    }
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    obj, err := root.MapUseNumber()
    if err != nil {
        t.Fatal(err)
    }
    exp, err := encoder.Encode(obj, encoder.SortMapKeys|encoder.NoEncoderNewline)
    if err != nil {
        t.Fatal(err)
    }
    var expObj interface{}
    require.NoError(t, json.Unmarshal(exp, &expObj))

    if err := root.SortKeys(true); err != nil {
        t.Fatal(err)
    }
    act, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    var actObj interface{}
    require.NoError(t, json.Unmarshal(act, &actObj))
    require.Equal(t, expObj, actObj)
    require.Equal(t, len(exp), len(act))
    require.Equal(t, string(exp), string(act))
}

func TestNodeAny(t *testing.T) {
    empty := Node{}
    _,err := empty.SetAny("any", map[string]interface{}{"a": []int{0}})
    if err != nil {
        t.Fatal(err)
    }
    if m, err := empty.Get("any").Interface(); err != nil {
        t.Fatal(err)
    } else if v, ok := m.(map[string]interface{}); !ok {
        t.Fatal(v)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]}}` {
        t.Fatal(string(buf))
    }
    if _, err := empty.Set("any2", Node{}); err != nil {
        t.Fatal(err)
    }
    if err := empty.Get("any2").AddAny(nil); err != nil {
        t.Fatal(err)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]},"any2":[null]}` {
        t.Fatal(string(buf))
    }
    if _, err := empty.Get("any2").SetAnyByIndex(0, NewNumber("-0.0")); err != nil {
        t.Fatal(err)
    }
    if buf, err := empty.MarshalJSON(); err != nil {
        t.Fatal(err)
    } else if string(buf) != `{"any":{"a":[0]},"any2":[-0.0]}` {
        t.Fatal(string(buf))
    }
}


func TestTypeCast2(t *testing.T) {
    type tcase struct {
        method string
        node Node
        exp interface{}
        err error
    }
    var cases = []tcase{
        {"Raw", NewAny(""), "\"\"", nil},
    }

    for i, c := range cases {
        fmt.Println(i, c)
        rt := reflect.ValueOf(&c.node)
        m := rt.MethodByName(c.method)
        rets := m.Call([]reflect.Value{})
        if len(rets) != 2 {
            t.Fatal(i, rets)
        }
        require.Equal(t, c.exp, rets[0].Interface())
        v := rets[1].Interface();
        if  v != c.err {
            t.Fatal(i, v)
        }
    }
}

func TestStackAny(t *testing.T) {
    var obj = stackObj()
    any := NewAny(obj)
    fmt.Printf("any: %#v\n", any)
    runtime.GC()
    debug.FreeOSMemory()
    println("finish GC")
    buf, err := any.MarshalJSON()
    println("finish marshal")
    if err != nil {
        t.Fatal(err)
    }
    if string(buf) != "1" {
        t.Fatal(string(buf))
    }
}


func Test_Export(t *testing.T) {
	type args struct {
		src  string
		path []interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantStart int
		wantEnd   int
		wantTyp   int
		wantErr   bool
		wantValid bool
	}{
		{"bool", args{`[true  ,2]`, []interface{}{0}}, 1, 5, V_TRUE, false, true},
		{"bool", args{`[t2ue  ,2]`, []interface{}{0}}, 1, 5, V_TRUE, false, false},
		{"number", args{`[1  ,2]`, []interface{}{0}}, 1, 2, V_NUMBER, false, true},
		{"number", args{`[1w ,2]`, []interface{}{0}}, 1, 3, V_NUMBER, false, false},
		{"string", args{`[" "  ,2]`, []interface{}{0}}, 1, 4, V_STRING, false, true},
		{"string", args{`[" "]  ,2]`, []interface{}{0}}, 1, 4, V_STRING, false, true},
		{"object", args{`[{"":""}  ,2]`, []interface{}{0}}, 1, 8, V_OBJECT, false, true},
		{"object", args{`[{x}  ,2]`, []interface{}{0}}, 1, 4, V_OBJECT, false, false},
		{"arrauy", args{`[[{}]  ,2]`, []interface{}{0}}, 1, 5, V_ARRAY, false, true},
		{"arrauy", args{`[[xx]  ,2]`, []interface{}{0}}, 1, 5, V_ARRAY, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, gotTyp, err := _GetByPath(tt.args.src, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("_GetByPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStart != tt.wantStart {
				t.Errorf("_GetByPath() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("_GetByPath() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
			if gotTyp != tt.wantTyp {
				t.Errorf("_GetByPath() gotTyp = %v, want %v", gotTyp, tt.wantTyp)
			}
			gotStart, gotEnd, err = _SkipFast(tt.args.src, tt.wantStart)
			if (err != nil) != tt.wantErr {
				t.Errorf("_SkipFast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStart != tt.wantStart {
				t.Errorf("_SkipFast() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("_SkipFast() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
			valid := _ValidSyntax(tt.args.src[tt.wantStart:tt.wantEnd])
			if valid != tt.wantValid {
				t.Errorf("_ValidSyntax() gotValid = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}
