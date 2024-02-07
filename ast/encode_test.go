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
    `runtime`
    `sync`
    `testing`

    `github.com/stretchr/testify/require`
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
                t.Error(err)
                return
            }
            root.Load()
            _, err = root.MarshalJSON()
            if err != nil {
                t.Error(err)
                return
            }
            runtime.GC()
        }(wg)
    }
    wg.Wait()
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

type SortableNode struct {
    sorted bool
	*Node
}

func (j *SortableNode) UnmarshalJSON(data []byte) (error) {
    j.Node = new(Node)
	return j.Node.UnmarshalJSON(data)
}

func (j *SortableNode) MarshalJSON() ([]byte, error) {
    if !j.sorted {
        j.Node.SortKeys(true)
        j.sorted = true
    }
	return j.Node.MarshalJSON()
}

func TestMarshalSort(t *testing.T) {
    var data = `{"d":3,"a":{"c":1,"b":2},"e":null}`
    var obj map[string]*SortableNode
    require.NoError(t, json.Unmarshal([]byte(data), &obj))
    out, err := json.Marshal(obj)
    require.NoError(t, err)
    require.Equal(t, `{"a":{"b":2,"c":1},"d":3,"e":null}`, string(out))
    out, err = json.Marshal(obj)
    require.NoError(t, err)
    require.Equal(t, `{"a":{"b":2,"c":1},"d":3,"e":null}`, string(out))
}

func BenchmarkEncodeRaw_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewSearcher(data).GetByPath()
    if e != nil {
        b.Fatal(root)
    }
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
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

func BenchmarkEncodeSkip_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewParser(data).Parse()
    if e != 0 {
        b.Fatal(root)
    }
    root.skipAllKey()
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
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

func BenchmarkEncodeLoad_Sonic(b *testing.B) {
    data := _TwitterJson
    root, e := NewParser(data).Parse()
    if e != 0 {
        b.Fatal(root)
    }
    root.loadAllKey()
    _, err := root.MarshalJSON()
    if err != nil {
        b.Fatal(err)
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

func TestEncodeNone(t *testing.T) {
    n := NewObject([]Pair{{Key:"a", Value:Node{}}})
    out, err := n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, "{}", string(out))
    n = NewObject([]Pair{{Key:"a", Value:NewNull()}, {Key:"b", Value:Node{}}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, `{"a":null}`, string(out))

    n = NewArray([]Node{Node{}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, "[]", string(out))
    n = NewArray([]Node{NewNull(), Node{}})
    out, err = n.MarshalJSON()
    require.NoError(t, err)
    require.Equal(t, `[null]`, string(out))
}
