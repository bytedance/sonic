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
	"testing"

	"github.com/bytedance/sonic/ast"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func BenchmarkParseOne_Gjson(b *testing.B) {
    ast := gjson.Parse(TwitterJson)
    node := ast.Get("statuses.2.id")
    v := node.Int()
    if v != 249289491129438208 {
        b.Fatal(node)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast := gjson.Parse(TwitterJson)
        node := ast.Get("statuses.2.id")
        _ = node.Int()
    }
}

func BenchmarkParseOne_Jsoniter(b *testing.B) {
    data := []byte(TwitterJson)
    ast := jsoniter.Get(data, "statuses", 2, "id")
    node := ast.ToInt()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data, "statuses", 2, "id")
        _ = ast.ToInt()
    }
}

func BenchmarkParseOne_Sonic(b *testing.B) {
    r, _ := ast.NewParser(TwitterJson).Parse()
    node, _ := r.Get("statuses").Index(2).Get("id").Int64()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast, _ := ast.NewParser(TwitterJson).Parse()
        _, _ = ast.Get("statuses").Index(2).Get("id").Int64()
    }
}

func BenchmarkParseOne_Parallel_Gjson(b *testing.B) {
    ast := gjson.Parse(TwitterJson)
    node := ast.Get("statuses.2.id")
    v := node.Int()
    if v != 249289491129438208 {
        b.Fatal(node)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := gjson.Parse(TwitterJson)
            node := ast.Get("statuses.2.id")
            _ = node.Int()
        }
    })
}

func BenchmarkParseOne_Parallel_Jsoniter(b *testing.B) {
    data := []byte(TwitterJson)
    ast := jsoniter.Get(data, "statuses", 2, "id")
    node := ast.ToInt()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            data := []byte(TwitterJson)
            ast := jsoniter.Get(data, "statuses", 2, "id")
            _ = ast.ToInt()
        }
    })
}

func BenchmarkParseOne_Parallel_Sonic(b *testing.B) {
    r, _ := ast.NewParser(TwitterJson).Parse()
    node, _ := r.Get("statuses").Index(2).Get("id").Int64()
    if node != 249289491129438208 {
        b.Fail()
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast, _ := ast.NewParser(TwitterJson).Parse()
            _, _ = ast.Get("statuses").Index(2).Get("id").Int64()
        }
    })
}

func BenchmarkParseSeven_Gjson(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast := gjson.Parse(TwitterJson)
        node := ast.Get("statuses.3.id")
        node = ast.Get("statuses.3.user.entities.description")
        node = ast.Get("statuses.3.user.entities.url.urls")
        node = ast.Get("statuses.3.user.entities.url")
        node = ast.Get("statuses.3.user.created_at")
        node = ast.Get("statuses.3.user.name")
        node = ast.Get("statuses.3.text")
        if node.Value() == nil {
            b.Fail()
        }
    }
}

func BenchmarkParseSeven_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    data := []byte(TwitterJson)
    for i := 0; i < b.N; i++ {
        ast := jsoniter.Get(data)
        node := ast.Get("statuses", 3, "id")
        node = ast.Get("statuses", 3, "user", "entities", "description")
        node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
        node = ast.Get("statuses", 3, "user", "entities", "url")
        node = ast.Get("statuses", 3, "user", "created_at")
        node = ast.Get("statuses", 3, "user", "name")
        node = ast.Get("statuses", 3, "text")
        if node.LastError() != nil {
            b.Fail()
        }
    }
}

func BenchmarkParseSeven_Sonic(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ast, _ := ast.NewParser(TwitterJson).Parse()
        node := ast.GetByPath("statuses", 3, "id")
        node = ast.GetByPath("statuses",  3, "user", "entities","description")
        node = ast.GetByPath("statuses",  3, "user", "entities","url","urls")
        node = ast.GetByPath("statuses",  3, "user", "entities","url")
        node = ast.GetByPath("statuses",  3, "user", "created_at")
        node = ast.GetByPath("statuses",  3, "user", "name")
        node = ast.GetByPath("statuses",  3, "text")
        if node.Check() != nil {
            b.Fail()
        }
    }
}

func BenchmarkParseSeven_Parallel_Gjson(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast := gjson.Parse(TwitterJson)
            node := ast.Get("statuses.3.id")
            node = ast.Get("statuses.3.user.entities.description")
            node = ast.Get("statuses.3.user.entities.url.urls")
            node = ast.Get("statuses.3.user.entities.url")
            node = ast.Get("statuses.3.user.created_at")
            node = ast.Get("statuses.3.user.name")
            node = ast.Get("statuses.3.text")
            if node.Value() == nil {
                b.Fail()
            }
        }
    })
}

func BenchmarkParseSeven_Parallel_Jsoniter(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            data := []byte(TwitterJson)
            ast := jsoniter.Get(data)
            node := ast.Get("statuses", 3, "id")
            node = ast.Get("statuses", 3, "user", "entities", "description")
            node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
            node = ast.Get("statuses", 3, "user", "entities", "url")
            node = ast.Get("statuses", 3, "user", "created_at")
            node = ast.Get("statuses", 3, "user", "name")
            node = ast.Get("statuses", 3, "text")
            if node.LastError() != nil {
                b.Fail()
            }
        }
    })
}

func BenchmarkParseSeven_Parallel_Sonic(b *testing.B) {
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ast, _ := ast.NewParser(TwitterJson).Parse()
            node := ast.GetByPath("statuses", 3, "id")
            node = ast.GetByPath("statuses",  3, "user", "entities","description")
            node = ast.GetByPath("statuses",  3, "user", "entities","url","urls")
            node = ast.GetByPath("statuses",  3, "user", "entities","url")
            node = ast.GetByPath("statuses",  3, "user", "created_at")
            node = ast.GetByPath("statuses",  3, "user", "name")
            node = ast.GetByPath("statuses",  3, "text")
            if node.Check() != nil {
                b.Fail()
            }
        }
    })
}

