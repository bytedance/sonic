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
    `encoding/json`
    `testing`
    `runtime`
    `runtime/debug`
    `sync`

    `github.com/bytedance/sonic/internal/native/types`
)

func TestGC_Encode(t *testing.T) {
	if debugSyncGC {
        return
    }
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    root.LoadAll()
    _, err = root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
			defer wg.Done()
			root, err := NewSearcher(_TwitterJson).GetByPath()
			if err != nil {
				t.Fatal(err)
			}
			root.Load()
            _, err = root.MarshalJSON()
            if err != nil {
                t.Fatal(err)
            }
            runtime.GC()
            debug.FreeOSMemory()
        }(wg)
    }
    wg.Wait()
}

func TestEncodeValue(t *testing.T) {
	type Case struct {
		node Node
		exp string
		err bool
	}
	input := []Case{
		{NewNull(), "null", false},
		{NewBool(true), "true", false},
		{NewBool(false), "false", false},
		{NewNumber("0.0"), "0.0", false},
		{NewString(""), `""`, false},
		{NewArray([]Node{}), "[]", false},
		{NewArray([]Node{NewBool(true), NewString("true")}), `[true,"true"]`, false},
		{NewObject([]Pair{Pair{"a", NewNull()}, Pair{"b", NewNumber("0")}}), `{"a":null,"b":0}`, false},
		{NewObject([]Pair{}), `{}`, false},
		{newRawNode(`[{ }]`, types.V_ARRAY), "[{}]", false},
		{Node{}, "", true},
		{Node{t: types.V_EOF}, "", true},
	}
	for i, c := range input {
		buf, err := json.Marshal(&c.node)
		if c.err {
			if err == nil {
				t.Fatal(i)
			}
			continue
		}
		if err != nil {
			t.Fatal(i, err)
		}
		if string(buf) != c.exp {
			t.Fatal(i, string(buf))
		}
	}
}

func TestEncodeNode(t *testing.T) {
	data := `{"a":[{},[],-0.1,true,false,null,""],"b":0,"c":true,"d":false,"e":null,"g":""}`
	root, e := NewSearcher(data).GetByPath()
	if e != nil {
		t.Fatal(root)
	}
	ret, err := root.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(ret) != data {
		t.Fatal(string(ret))
	}
	root.skipAllKey()
	ret, err = root.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(ret) != data {
		t.Fatal(string(ret))
	}
	root.loadAllKey()
	ret, err = root.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(ret) != data {
		t.Fatal(string(ret))
	}
}

func BenchmarkEncodeRaw(b *testing.B) {
	data := _TwitterJson
	root, e := NewSearcher(data).GetByPath()
	if e != nil {
		b.Fatal(root)
	}
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, err := root.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeSkip(b *testing.B) {
	data := _TwitterJson
	root, e := NewParser(data).Parse()
	if e != 0 {
		b.Fatal(root)
	}
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, err := root.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeLoad(b *testing.B) {
	data := _TwitterJson
	root, e := NewParser(data).Parse()
	if e != 0 {
		b.Fatal(root)
	}
	root.loadAllKey()
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, err := root.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}