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

package benchmark_test

import (
    `fmt`
    `math`
    `testing`

    `github.com/buger/jsonparser`
    `github.com/bytedance/sonic/ast`
    jsoniter `github.com/json-iterator/go`
    `github.com/tidwall/gjson`
    `github.com/tidwall/sjson`
)

func BenchmarkGetOne_Gjson(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    for i := 0; i < b.N; i++ {
        ast := gjson.Get(TwitterJson, "statuses.3.id")
        node := ast.Int()
        if node != 249279667666817024 {
            b.Fail()
        }
    }
}

func BenchmarkGetOne_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    data := []byte(TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 3, "id")
        node := ast.ToInt()
        if node != 249279667666817024 {
            b.Fail()
        }
    }
}

func BenchmarkGetOne_Sonic(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    ast := ast.NewSearcher(TwitterJson)
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

func BenchmarkGet(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("gjson", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            node := gjson.Get(TwitterJson, "statuses.3.id").Int()
            if node != 249279667666817024 {
                b.Fail()
            }
        }
    })
    b.Run("Node", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses", 3, "id")
            x, _ := n.Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses", 3, "id")
            x, _ := n.Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
}

func BenchmarkNodeGet(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("gjson", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            node := gjson.Get(TwitterJson, "statuses")
            v := node.Get("3").Get("id").Int()
            if v != 249279667666817024 {
                b.Fail()
            }
        }
    })
    b.Run("Node", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses")
            x, _ := n.Index(3).Get("id").Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetValueByPath("statuses")
            x, _ := n.Index(3).Get("id").Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
}

func BenchmarkSubGet(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("gjson", func(b *testing.B) {
        node := gjson.Get(TwitterJson, "statuses")
        for i := 0; i < b.N; i++ {
            v := node.Get("3").Get("id").Int()
            if v != 249279667666817024 {
                b.Fail()
            }
        }
    })
    b.Run("Node", func(b *testing.B) {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses")
        for i := 0; i < b.N; i++ {
            x, _ := n.Index(3).Get("id").Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetValueByPath("statuses")
        for i := 0; i < b.N; i++ {
            x, _ := n.Index(3).Get("id").Int64()
            if x != 249279667666817024 {
                b.Fatal(n.Interface())
            }
        }
    })
}

func BenchmarkNodeGetMany(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("gjson", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            nodes := gjson.GetMany(TwitterJson, "statuses.0", "statuses.3", "statuses.100")
            if !nodes[0].Exists() {
                b.Fatal()
            }
            if !nodes[1].Exists() {
                b.Fatal()
            }
            if nodes[2].Exists() {
                b.Fatal()
            }
        }
    })
    b.Run("Node", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses")
            nodes := []*ast.Node{}
            x := n.Index(0)
            nodes = append(nodes, x)
            x = n.Index(3)
            nodes = append(nodes, x)
            x = n.Index(100)
            nodes = append(nodes, x)
            if !nodes[0].Exists() {
                b.Fatal()
            }
            if !nodes[1].Exists() {
                b.Fatal()
            }
            if nodes[2].Exists() {
                b.Fatal()
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetValueByPath("statuses")
            nodes := [3]ast.Value{}
            err := n.IndexMany([]int{0,3,100}, nodes[:])
            if err != nil {
                b.Fatal(err)
            }
            if !nodes[0].Exists() {
                b.Fatal()
            }
            if !nodes[1].Exists() {
                b.Fatal()
            }
            if nodes[2].Exists() {
                b.Fatal()
            }
        }
    })
}

func TestSetNotExist(t *testing.T) {
    js, err := sjson.Set(`[]`, "0", 1)
    if err != nil {
        t.Fatal(err)
    }
    println(js)
    js, err = sjson.Set(js, "10", 1)
    if err != nil {
        t.Fatal(err)
    }
    println(js)
}

func BenchmarkNodeSetMany(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("sjson", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            js, err := sjson.Set(TwitterJson, "statuses.0", 1)
            if err != nil {
                b.Fatal(err)
            }
            _, err = sjson.Set(js, "statuses.3", 1)
            if err != nil {
                b.Fatal(err)
            }
            // _, err = sjson.Set(js, "statuses.100", 1)
            // if err != nil {
            //     b.Fatal(err)
            // }
        }
    })
    b.Run("Node", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses")
            exist, err := n.SetByIndex(0, ast.NewAny(1))
            if !exist || err != nil {
                b.Fatal()
            }
            exist, err = n.SetByIndex(3, ast.NewAny(1))
            if !exist || err != nil {
                b.Fatal()
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetValueByPath("statuses")
            nodes := []ast.Value{ast.NewValue(1), ast.NewValue(2)}
            exist, err := n.SetManyByIndex([]int{0,3}, nodes)
            if !exist || err != nil {
                b.Fatal()
            }
        }
    })
}

func BenchmarkSet(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.Run("sjson", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _, err := sjson.Set(TwitterJson, "statuses.3.id", 1)
            if err != nil {
                b.Fatal(err)
            }
        }
    })
    b.Run("Node", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            n, _ := s.GetByPath("statuses", 3)
            _, err := n.SetAny("id", 1)
            if err != nil {
                b.Fatal(err)
            }
        }
    })
    b.Run("Value", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s := ast.NewSearcher(TwitterJson)
            _, err := s.SetValueByPath(ast.NewValue(1), "statuses", 3, "id")
            if err != nil {
                b.Fatal(err)
            }
        }
    })
}

func BenchmarkGetOne_Parallel_Sonic(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.RunParallel(func(pb *testing.PB) {
        ast := ast.NewSearcher(TwitterJson)
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

func BenchmarkGetOne_Parallel_Gjson(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := gjson.Get(TwitterJson, "statuses.3.id")
            node := ast.Int()
            if node != 249279667666817024 {
                b.Fail()
            }
        }
    })
}

func BenchmarkGetOne_Parallel_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    data := []byte(TwitterJson)
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

func BenchmarkSetOne_Sjson(b *testing.B) {
    path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
    _, err := sjson.Set(TwitterJson, path, math.MaxInt32)
    if err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sjson.Set(TwitterJson, path, math.MaxInt32)
    }
}

func BenchmarkSetOne_Jsoniter(b *testing.B) {
    data := []byte(TwitterJson)
    node, ok := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
    if !ok {
        b.Fatal(node)
    }

    b.SetBytes(int64(len(data)))
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        node, _ := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
        node["id"] = math.MaxInt32
    }
}

func BenchmarkSetOne_Parallel_Sjson(b *testing.B) {
    path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
    _, err := sjson.Set(TwitterJson, path, math.MaxInt32)
    if err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ReportAllocs()
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            sjson.Set(TwitterJson, path, math.MaxInt32)
        }
    })
}

func BenchmarkSetOne_Parallel_Jsoniter(b *testing.B) {
    data := []byte(TwitterJson)
    node, ok := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
    if !ok {
        b.Fatal(node)
    }

    b.SetBytes(int64(len(data)))
    b.ReportAllocs()
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            node, _ := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
            node["id"] = math.MaxInt32
        }
    })
}


func BenchmarkGetByKeys_Sonic(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    ast := ast.NewSearcher(TwitterJson)
    const _count = 4
    for i := 0; i < b.N; i++ {
        node, err := ast.GetByPath("search_metadata", "count")
        if err != nil {
            b.Fatal(err)
        }
        x, _ := node.Int64()
        if x != _count {
            b.Fatal(node.Interface())
        }
    }
}


func BenchmarkGetByKeys_JsonParser(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    data := []byte(TwitterJson)
    const _count = 4
    for i := 0; i < b.N; i++ {
        value, err := jsonparser.GetInt(data, "search_metadata", "count")
        if err != nil {
            b.Fatal(err)
        }
        if value != _count {
            b.Fatal(value)
        }
    }
}