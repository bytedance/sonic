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
    `testing`

    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
    `github.com/tidwall/gjson`
)

func TestExportError(t *testing.T) {
    data := `{"a":]`
    p := NewSearcher(data)
    _, err := p.GetByPath("a")
    if err == nil {
        t.Fatal()
    }
    if err.Error() != `"Syntax error at index 6: invalid char\n\n\t{\"a\":]\n\t......^\n"` {
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
}

func TestSearcher_GetByPath(t *testing.T) {
    s := NewSearcher(` { "xx" : [] ,"yy" :{ }, "test" : [ true , 0.1 , "abc", ["h"], {"a":"bc"} ] } `)

    node, e := s.GetByPath("test", 0)
    if e != nil || node.Bool() != true {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = s.GetByPath("test", 1)
    if e != nil || node.Float64() != 0.1 {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = s.GetByPath("test", 2)
    if e != nil || node.String() != "abc" {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = s.GetByPath("test", 3)
    arr, _ := node.Array()
    if e != nil || arr[0] != "h" {
        t.Fatalf("node: %v, err: %v", node, e)
    }

    node, e = s.GetByPath("test", 4, "a")
    if e != nil || node.String() != "bc" {
        t.Fatalf("node: %v, err: %v", node, e)
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
    a := node.Index(3).Float64()
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

func BenchmarkSearchOne_Gjson(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    for i := 0; i < b.N; i++ {
        ast := gjson.Get(_TwitterJson, "statuses.3.id")
        node := ast.Int()
        if node != 249279667666817024 {
            b.Fail()
        }
    }
}

func BenchmarkSearchOne_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    data := []byte(_TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 3, "id")
        node := ast.ToInt()
        if node != 249279667666817024 {
            b.Fail()
        }
    }
}

func BenchmarkSearchOne_Sonic(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    ast := NewSearcher(_TwitterJson)
    for i := 0; i < b.N; i++ {
        node, err := ast.GetByPath("statuses", 3, "id")
        if err != nil {
            b.Fatal(err)
        }
        if node.Int64() != 249279667666817024 {
            b.Fatal(node.Interface())
        }
    }
}

func BenchmarkSearchOne_Parallel_Gjson(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := gjson.Get(_TwitterJson, "statuses.3.id")
            node := ast.Int()
            if node != 249279667666817024 {
                b.Fail()
            }
        }
    })
}

func BenchmarkSearchOne_Parallel_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    data := []byte(_TwitterJson)
    b.SetParallelism(parallelism)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := jsoniter.Get(data, "statuses", 3, "id")
            node := ast.ToInt()
            if node != 249279667666817024 {
                b.Fail()
            }
        }
    })
}

func BenchmarkSearchOne_Parallel_Sonic(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.RunParallel(func(pb *testing.PB) {
        ast := NewSearcher(_TwitterJson)
        for pb.Next() {
            node, err := ast.GetByPath("statuses", 3, "id")
            if err != nil {
                b.Fatal(err)
            }
            if node.Int64() != 249279667666817024 {
                b.Fatal(node.Interface())
            }
        }
    })
}
