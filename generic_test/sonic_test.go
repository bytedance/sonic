// +build go1.18

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

package generic_test

import (
	`testing`
    `reflect`
    
    `github.com/bytedance/sonic`
	`github.com/bytedance/sonic/option`
    `github.com/bytedance/sonic/ast`
)

type Str interface {
    string
}

type Bytes interface {
    []byte
}

type Any interface {
    any
}

type SliceAny interface {
    []any
}

func unmarshalAny[S Str, B Bytes, T Any](data S, val T) error {
    return sonic.Unmarshal(B(data), val)
}

func marshalAny[B Bytes, T Any](val T) (B, error) {
    return sonic.Marshal(val)
}

func getAny[S Str, B Bytes, T SliceAny](src S, path T) (ast.Node, error) {
    return sonic.Get(B(src), path...)
}

func pretouchAny[T Any](v T, opts ...option.CompileOption) error {
	rt := reflect.TypeOf(v)
	return sonic.Pretouch(rt, opts...)
}

type Basic interface {
    ~*int|~*float64|~float64|~*string
}

func unmarshalBasic[S Str, B Bytes, T Basic](data S, val T) error {
    return sonic.Unmarshal(B(data), val)
}

func marshalBasic[B Bytes, T Basic](val T) (B, error) {
    return sonic.Marshal(val)
}

func pretouchBasic[T Basic](v T, opts ...option.CompileOption) error {
	rt := reflect.TypeOf(v)
	return sonic.Pretouch(rt, opts...)
}

type Float64 float64

func TestGenericAPI(t *testing.T) {
    var x interface{}
    if err := unmarshalAny(`{"a":[true,0.5,"hello world"]}`, &x); err != nil {
        t.Fatal(t)
    }
    out, err := marshalAny(x)
    if err != nil {
        t.Fatal(err)
    }
	t.Logf("%s", out)
	
	var x0 = struct{
		A []Any `json:"a"`
	}{}
	if err := pretouchAny(x0); err != nil {
		t.Fatal(err)
	}
	if err := unmarshalAny(`{"a":[true,0.5,"hello world"]}`, &x0); err != nil {
        t.Fatal(t)
    }
    out0, err := marshalAny(x0)
    if err != nil {
        t.Fatal(err)
    }
	t.Logf("%s", out0)

    var x1 int
    if err := unmarshalBasic(`1`, &x1); err != nil {
        t.Fatal(t)
    }
    out1, err := marshalBasic(&x1)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("%s", out1)

	var x2 Float64 = 1
	// if err := unmarshalBasic(`1`, &x2); err != nil {
    //     t.Fatal(t)
    // }
    out2, err := marshalBasic(x2)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("%s", out2)

    var x3 string
    if err := unmarshalBasic(`"1"`, &x3); err != nil {
        t.Fatal(t)
    }
    out3, err := marshalBasic(&x3)
    if err != nil {
        t.Fatal(err)
    }
	t.Logf("%s", out3)

	var x4 Float64 = 1
	if err := pretouchBasic(x4); err != nil {
		t.Fatal(err)
	}
	// if err := unmarshalBasic(`0.5`, &x4); err != nil {
    //     t.Fatal(t)
    // }
    out4, err := marshalBasic(x4)
    if err != nil {
        t.Fatal(err)
    }
	t.Logf("%s", out4)

    root, err := getAny(`{"a":[true,1,"hello world"]}`, []interface{}{"a", 1})
    if err != nil {
        t.Fatal(err)
    }
    f, err := root.Float64()
    if err != nil {
        t.Fatal(err)
    }
    t.Log(f)
}