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
    `github.com/tidwall/gjson`
)

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
    if e != nil || node.Array()[0] != "h" {
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

func BenchmarkSearchOne_Gjson(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    for i := 0; i < b.N; i++ {
        ast := gjson.Get(_TwitterJson, "statuses.2.id")
        node := ast.Int()
        if node != 249289491129438208 {
            b.Fail()
        }
    }
}

func BenchmarkSearchOne_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    data := []byte(_TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 2, "id")
        node := ast.ToInt()
        if node != 249289491129438208 {
            b.Fail()
        }
    }
}

func BenchmarkSearchOne_Sonic(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    ast := NewSearcher(_TwitterJson)
    for i := 0; i < b.N; i++ {
        node, err := ast.GetByPath("statuses", 2, "id")
        if err != nil {
            b.Fatal(err)
        }
        if node.Int64() != 249289491129438208 {
            b.Fatal(node.Interface())
        }
    }
}

func BenchmarkSearchOne_Parallel_Gjson(b *testing.B) {
    b.SetBytes(int64(len(_TwitterJson)))
    b.SetParallelism(parallelism)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := gjson.Get(_TwitterJson, "statuses.2.id")
            node := ast.Int()
            if node != 249289491129438208 {
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
            ast := jsoniter.Get(data, "statuses", 2, "id")
            node := ast.ToInt()
            if node != 249289491129438208 {
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
            node, err := ast.GetByPath("statuses", 2, "id")
            if err != nil {
                b.Fatal(err)
            }
            if node.Int64() != 249289491129438208 {
                b.Fatal(node.Interface())
            }
        }
    })
}
