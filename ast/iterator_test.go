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
	"fmt"
	"strconv"
	"testing"
)

func getTestIteratorSample() (string, int) {
	var data []int
	var v1 = ""
	var v2 = ""
	loop := _DEFAULT_NODE_CAP + 1
	for i := 0; i < loop; i++ {
		data = append(data, i*i)
		v1 += strconv.Itoa(i)
		v2 += `"k` + strconv.Itoa(i) + `":` + strconv.Itoa(i)
		if i != loop-1 {
			v1 += `,`
			v2 += `,`
		}
	}
	return `{"array":[` + v1 + `], "object":{` + v2 + `}}`, loop
}

func TestRawIterator(t *testing.T) {
	str, loop := getTestIteratorSample()
	fmt.Println(str)

	root, err := NewSearcher(str).GetByPath("array")
	if err != nil {
		t.Fatal(err)
	}
	ai := root.Values()
	i := int64(0)
	for ai.HasNext() {
		v := &Node{}
		if !ai.Next(v) {
			t.Fatalf("no next")
		}
		if i < int64(loop) && v.Int64() != i {
			t.Fatalf("exp:%v, got:%v", i, v)
		}
		if i != int64(ai.Pos())-1 || i >= int64(ai.Len()) {
			t.Fatal(i)
		}
		i++
	}

	root, err = NewSearcher(str).GetByPath("object")
	if err != nil {
		t.Fatal(err)
	}
	mi := root.Properties()
	i = int64(0)
	for mi.HasNext() {
		v := &Pair{}
		if !mi.Next(v) {
			t.Fatalf("no next")
		}
		if i < int64(loop) && (v.Value.Int64() != i || v.Key != fmt.Sprintf("k%d", i)) {
			t.Fatalf("exp:%v, got:%v", i, v.Value.Interface())
		}
		if i != int64(mi.Pos())-1 || i >= int64(mi.Len()) {
			t.Fatal(i)
		}
		i++
	}
}

func TestIterator(t *testing.T) {
	str, loop := getTestIteratorSample()
	fmt.Println(str)

	root, err := NewParser(str).Parse()
	if err != 0 {
		t.Fatal(err)
	}
	ai := root.Get("array").Values()
	i := int64(0)
	for ai.HasNext() {
		v := &Node{}
		if !ai.Next(v) {
			t.Fatalf("no next")
		}
		if i < int64(loop) && v.Int64() != i {
			t.Fatalf("exp:%v, got:%v", i, v)
		}
		if i != int64(ai.Pos())-1 || i >= int64(ai.Len()) {
			t.Fatal(i)
		}
		i++
	}

	root, err = NewParser(str).Parse()
	if err != 0 {
		t.Fatal(err)
	}
	mi := root.Get("object").Properties()
	i = int64(0)
	for mi.HasNext() {
		v := &Pair{}
		if !mi.Next(v) {
			t.Fatalf("no next")
		}
		if i < int64(loop) && (v.Value.Int64() != i || v.Key != fmt.Sprintf("k%d", i)) {
			t.Fatalf("exp:%v, got:%v", i, v.Value.Interface())
		}
		if i != int64(mi.Pos())-1 || i >= int64(mi.Len()) {
			t.Fatal(i)
		}
		i++
	}
}

func BenchmarkArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 1, "entities", "hashtags")
		if err != nil {
			b.Fatal(err)
		}
		a := root.Array()
		for _, v := range a {
			_ = v
		}
	}
}

func BenchmarkListIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 1, "entities", "hashtags")
		if err != nil {
			b.Fatal(err)
		}
		it := root.Values()
		for it.HasNext() {
			v := &Node{}
			if !it.Next(v) {
				b.Fatalf("no value")
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 1, "user")
		if err != nil {
			b.Fatal(err)
		}
		m := root.Map()
		for k, v := range m {
			_ = v
			_ = k
		}
	}
}

func BenchmarkObjectIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root, err := NewSearcher(_TwitterJson).GetByPath("statuses", 1, "user")
		if err != nil {
			b.Fatal(err)
		}
		it := root.Properties()
		for it.HasNext() {
			v := &Pair{}
			if !it.Next(v) {
				b.Fatalf("no value")
			}
		}
	}
}
