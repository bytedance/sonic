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

package ast

import (
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGC_Search(t *testing.T) {
	if debugSyncGC {
		return
	}
	_, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0, "id")
	if err != nil {
		t.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	// A limitation of the race detecting is 8128.
	// See https://github.com/golang/go/issues/43898
	N := 5000
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			_, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0, "id")
			if err != nil {
				t.Error(err)
				return
			}
			runtime.GC()
		}(wg)
	}
	wg.Wait()
}

func TestExportErrorInvalidChar(t *testing.T) {
	data := `{"a":]`
	p := NewSearcher(data)
	_, err := p.GetByPath("a")
	if err == nil {
		t.Fatal()
	}
	if strings.Index(err.Error(), `"Syntax error at `) != 0 {
		t.Fatal(err)
	}

	data = `:"b"]`
	p = NewSearcher(data)
	_, err = p.GetByPath("a")
	if err == nil {
		t.Fatal()
	}
	if err.Error() != `"Syntax error at index 0: invalid char\n\n\t:\"b\"]\n\t^....\n"` {
		t.Fatal(err)
	}

	data = `{:"b"]`
	p = NewSearcher(data)
	_, err = p.GetByPath("a")
	if err == nil {
		t.Fatal()
	}
	if err.Error() != `"Syntax error at index 1: invalid char\n\n\t{:\"b\"]\n\t.^....\n"` {
		t.Fatal(err)
	}

	data = `{`
	p = NewSearcher(data)
	_, err = p.GetByPath("he")
	if err == nil {
		t.Fatal()
	}
	if err == ErrNotExist {
		t.Fatal(err)
	}

	data = `[`
	p = NewSearcher(data)
	_, err = p.GetByPath(0)
	if err == nil {
		t.Fatal()
	}
	if err == ErrNotExist {
		t.Fatal(err)
	}
}

type testExportError struct {
	data string
	path []interface{}
	err  error
}

func TestExportErrNotExist(t *testing.T) {
	tests := []testExportError{
		// object
		{`{}`, []interface{}{"b"}, ErrNotExist},
		{` {  } `, []interface{}{"b"}, ErrNotExist},
		{`{"a":null}`, []interface{}{"b"}, ErrNotExist},
		// This should be invalid char errors.
		// {`{"a":null}`, []interface{}{"a", "b"}, ErrNotExist},
		// {`{"a":null}`, []interface{}{"a", 0}, ErrNotExist},
		// {`{"a":null}`, []interface{}{"a", "b", 0}, ErrNotExist},
		{`{"":{"b":123}}`, []interface{}{"b"}, ErrNotExist},
		{`{"":{"b":123}}`, []interface{}{"", ""}, ErrNotExist},
		{`{"a":{"b":123}}`, []interface{}{"b"}, ErrNotExist},
		{`{"a":{"b":123}}`, []interface{}{"a", "c"}, ErrNotExist},
		{`{"a":{"c": null, "b":{}}}`, []interface{}{"a", "b", "c"}, ErrNotExist},
		{`{"a":{"b":123}}`, []interface{}{"b", "b"}, ErrNotExist},
		{`{"\"\\":{"b":123}}`, []interface{}{"\"", "b"}, ErrNotExist},
		{`{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\"\\":{"b":123}}`,
			[]interface{}{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\"", "b"}, ErrNotExist},

		// array
		{`[]`, []interface{}{0}, ErrNotExist},
		{`[]`, []interface{}{1}, ErrNotExist},
		{` [ ] `, []interface{}{0}, ErrNotExist},
		{`[null]`, []interface{}{1}, ErrNotExist},
		{`[null, ["null", 123]]`, []interface{}{2}, ErrNotExist},
		{`[null, true , false, 14, 2.35, -46, "hello7", "\"8"]`, []interface{}{8}, ErrNotExist},
		{`[{}]`, []interface{}{1}, ErrNotExist},
		{`[[]]`, []interface{}{1}, ErrNotExist},
		{`[[],[{},{}, []],{}]`, []interface{}{3}, ErrNotExist},
	}

	for _, test := range tests {
		f := func(t *testing.T) {
			p := NewSearcher(test.data)
			node, err := p.GetByPath(test.path...)
			if err != test.err || node.Exists() {
				t.Fatal(err)
			}
		}
		t.Run(test.data, f)
	}
}

func TestSearcher_GetByPath(t *testing.T) {
	s := NewSearcher(` { "xx" : [] ,"yy" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ] } `)

	node, e := s.GetByPath("test", 0)
	a, _ := node.Bool()
	if e != nil || a != true {
		t.Fatalf("node: %v, err: %v", node, e)
	}

	node, e = s.GetByPath("test", 1)
	b, _ := node.Float64()
	if e != nil || b != 0.1 {
		t.Fatalf("node: %v, err: %v", node, e)
	}

	node, e = s.GetByPath("test", 2)
	c, _ := node.String()
	if e != nil || c != "abc" {
		t.Fatalf("node: %v, err: %v", node, e)
	}

	node, e = s.GetByPath("test", 3)
	arr, _ := node.Array()
	if e != nil || arr[0] != "h" {
		t.Fatalf("node: %v, err: %v", node, e)
	}

	node, e = s.GetByPath("test", 4, "a")
	d, _ := node.String()
	if e != nil || d != "bc" {
		t.Fatalf("node: %v, err: %v", node, e)
	}
}

type testGetByPath struct {
	json  string
	path  []interface{}
	value interface{}
	ok    bool
}

func TestSearcher_GetByPathSingle(t *testing.T) {
	type Path = []interface{}
	const Ok = true
	const Error = false
	tests := []testGetByPath{
		{`true`, Path{}, true, Ok},
		{`false`, Path{}, false, Ok},
		{`null`, Path{}, nil, Ok},
		{`12345`, Path{}, 12345.0, Ok},
		{`12345.6789`, Path{}, 12345.6789, Ok},
		{`"abc"`, Path{}, "abc", Ok},
		{`"a\"\\bc"`, Path{}, "a\"\\bc", Ok},
		{`{"a":1}`, Path{"a"}, 1.0, Ok},
		{`{"":1}`, Path{""}, 1.0, Ok},
		{`{"":{"":1}}`, Path{"", ""}, 1.0, Ok},
		{`[1,2,3]`, Path{0}, 1.0, Ok},
		{`[1,2,3]`, Path{1}, 2.0, Ok},
		{`[1,2,3]`, Path{2}, 3.0, Ok},

		{`tru`, Path{}, nil, Error},
		{`fal`, Path{}, nil, Error},
		{`nul`, Path{}, nil, Error},
		{`{"a":1`, Path{}, nil, Error},
		{`x12345.6789`, Path{}, nil, Error},
		{`"abc`, Path{}, nil, Error},
		{`"a\"\\bc`, Path{}, nil, Error},
		{`"a\"\`, Path{}, nil, Error},
		{`{"a":`, Path{"a"}, nil, Error},
		{`[1,2,3]`, Path{4}, nil, Error},
		{`[1,2,3]`, Path{"a"}, nil, Error},
	}
	for _, test := range tests {
		t.Run(test.json, func(t *testing.T) {
			s := NewSearcher(test.json)
			node, err1 := s.GetByPath(test.path...)
			assert.Equal(t, test.ok, err1 == nil)

			value, err2 := node.Interface()
			assert.Equal(t, test.value, value)
			assert.Equal(t, test.ok, err2 == nil)
		})
	}
}

func TestSearcher_GetByPathErr(t *testing.T) {
	s := NewSearcher(` { "xx" : [] ,"yy" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ], "err1":[a, ] , "err2":{ ,"x":"xx"} } `)
	node, e := s.GetByPath("zz")
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	s.parser.p = 0
	node, e = s.GetByPath("xx", 4)
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	s.parser.p = 0
	node, e = s.GetByPath("yy", "a")
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	s.parser.p = 0
	node, e = s.GetByPath("test", 2, "x")
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	s.parser.p = 0
	node, e = s.GetByPath("err1", 0)
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	s.parser.p = 0
	node, e = s.GetByPath("err2", "x")
	if e == nil {
		t.Fatalf("node: %v, err: %v", node, e)
	}
}

func TestLoadIndex(t *testing.T) {
	node, err := NewSearcher(`{"a":[-0, 1, -1.2, -1.2e-10]}`).GetByPath("a")
	if err != nil {
		t.Fatal(err)
	}
	a, _ := node.Index(3).Float64()
	assert.Equal(t, -1.2e-10, a)
	m, _ := node.Array()
	assert.Equal(t, m, []interface{}{
		float64(0),
		float64(1),
		-1.2,
		-1.2e-10,
	})
}

func TestSearchNotExist(t *testing.T) {
	s := NewSearcher(` { "xx" : [ 0, "" ] ,"yy" :{ "2": "" } } `)
	node, e := s.GetByPath("xx", 2)
	if node.Exists() {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	node, e = s.GetByPath("xx", 1)
	if e != nil || !node.Exists() {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	node, e = s.GetByPath("yy", "3")
	if node.Exists() {
		t.Fatalf("node: %v, err: %v", node, e)
	}
	node, e = s.GetByPath("yy", "2")
	if e != nil || !node.Exists() {
		t.Fatalf("node: %v, err: %v", node, e)
	}
}

func BenchmarkGetOne_Sonic(b *testing.B) {
	b.SetBytes(int64(len(_TwitterJson)))
	ast := NewSearcher(_TwitterJson)
	for i := 0; i < b.N; i++ {
		node, err := ast.GetByPath("statuses", 3, "id")
		if err != nil {
			b.Fatal(err)
		}
		x, _ := node.Int64()
		if x != 249279667666817024 {
			b.Fatal(node.Interface())
		}
	}
}

func BenchmarkGetFull_Sonic(b *testing.B) {
	ast := NewSearcher(_TwitterJson)
	b.SetBytes(int64(len(_TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node, err := ast.GetByPath()
		if err != nil || node.Type() != V_OBJECT {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetWithManyCompare_Sonic(b *testing.B) {
	b.SetBytes(int64(len(_LotsCompare)))
	ast := NewSearcher(_LotsCompare)
	for i := 0; i < b.N; i++ {
		node, err := ast.GetByPath("is")
		if err != nil {
			b.Fatal(err)
		}
		x, _ := node.Int64()
		if x != 1 {
			b.Fatal(node.Interface())
		}
	}
}

func BenchmarkGetOne_Parallel_Sonic(b *testing.B) {
	b.SetBytes(int64(len(_TwitterJson)))
	b.RunParallel(func(pb *testing.PB) {
		ast := NewSearcher(_TwitterJson)
		for pb.Next() {
			node, err := ast.GetByPath("statuses", 3, "id")
			if err != nil {
				b.Fatal(err)
			}
			x, _ := node.Int64()
			if x != 249279667666817024 {
				b.Fatal(node.Interface())
			}
		}
	})
}

func BenchmarkSetOne_Sonic(b *testing.B) {
	node, err := NewSearcher(_TwitterJson).GetByPath("statuses", 3)
	if err != nil {
		b.Fatal(err)
	}
	n := NewNumber(strconv.Itoa(math.MaxInt32))
	_, err = node.Set("id", n)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(_TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node, _ := NewSearcher(_TwitterJson).GetByPath("statuses", 3)
		_, _ = node.Set("id", n)
	}
}

func BenchmarkSetOne_Parallel_Sonic(b *testing.B) {
	node, err := NewSearcher(_TwitterJson).GetByPath("statuses", 3)
	if err != nil {
		b.Fatal(err)
	}
	n := NewNumber(strconv.Itoa(math.MaxInt32))
	_, err = node.Set("id", n)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(_TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			node, _ := NewSearcher(_TwitterJson).GetByPath("statuses", 3)
			_, _ = node.Set("id", n)
		}
	})
}
