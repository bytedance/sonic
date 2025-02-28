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
    `testing`

    `github.com/bytedance/sonic/ast`
    jsoniter `github.com/json-iterator/go`
    `github.com/tidwall/gjson`
    fastjson `github.com/valyala/fastjson`
)

func BenchmarkParser_Gjson(b *testing.B) {
    gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
        if !value.Exists() {
            b.Fatal(value.Index)
        }
        _ = value.Value()
        return true
    })
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
            if !value.Exists() {
                b.Fatal(value.Index)
            }
            _ = value.Value()
            return true
        })
    }
}

func BenchmarkParser_Jsoniter(b *testing.B) {
    v := jsoniter.Get([]byte(TwitterJson)).GetInterface()
    if v == nil {
        b.Fatal(v)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = jsoniter.Get([]byte(TwitterJson)).GetInterface()
    }
}

func BenchmarkParser_Parallel_Gjson(b *testing.B) {
    gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
        if !value.Exists() {
            b.Fatal(value.Index)
        }
        return true
    })
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
                if !value.Exists() {
                    b.Fatal(value.Index)
                }
                _ = value.Value()
                return true
            })
        }
    })
}

func BenchmarkParser_Parallel_Jsoniter(b *testing.B) {
    var bv = []byte(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var out interface{}
            _ = jsoniter.Unmarshal(bv, &out)
        }
    })
}

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

func BenchmarkParseOne_Fastjson(b *testing.B) {
    data := []byte(TwitterJson)
    var p fastjson.Parser
    v, err := p.ParseBytes(data)
    if err != nil {
        b.Fatal(err)
    }
    id := v.Get("statuses").GetArray()[2].GetInt64("id")
    if id != 249289491129438208 {
        b.Fatal("value mismatch")
    }
    
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var p fastjson.Parser
        v, _ := p.ParseBytes(data)
        _ = v.Get("statuses").GetArray()[2].GetInt64("id")
    }
}


func BenchmarkParseOne_Parallel_Fastjson(b *testing.B) {
    data := []byte(TwitterJson)
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var p fastjson.Parser
            v, _ := p.ParseBytes(data)
            _ = v.Get("statuses").GetArray()[2].GetInt64("id")
        }
    })
}

func BenchmarkParseSeven_Fastjson(b *testing.B) {
    data := []byte(TwitterJson)
    var p fastjson.Parser
    v, err := p.ParseBytes(data)
    if err != nil {
        b.Fatal(err)
    }
    
    statuses := v.GetArray("statuses")
    if len(statuses) < 4 {
        b.Fatal("insufficient statuses")
    }
    status := statuses[3]
    
    checks := []func(*fastjson.Value){
        func(v *fastjson.Value) { v.GetInt64("id") },
        func(v *fastjson.Value) { v.Get("user").Get("entities").GetStringBytes("description") },
        func(v *fastjson.Value) { v.Get("user").Get("entities").Get("url").GetArray("urls") },
        func(v *fastjson.Value) { v.Get("user").Get("entities").Get("url") },
        func(v *fastjson.Value) { v.Get("user").GetStringBytes("created_at") },
        func(v *fastjson.Value) { v.Get("user").GetStringBytes("name") },
        func(v *fastjson.Value) { v.GetStringBytes("text") },
    }
    
    for _, check := range checks {
        check(status)
    }
    
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var p fastjson.Parser
        v, _ := p.ParseBytes(data)
        status := v.GetArray("statuses")[3]
        for _, check := range checks {
            check(status)
        }
    }
}

func BenchmarkParseSeven_Parallel_Fastjson(b *testing.B) {
    data := []byte(TwitterJson)
    b.SetBytes(int64(len(data)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            var p fastjson.Parser
            v, _ := p.ParseBytes(data)
            status := v.GetArray("statuses")[3]
            status.GetInt64("id")
            status.Get("user").Get("entities").GetStringBytes("description")
            status.Get("user").Get("entities").Get("url").GetArray("urls")
            status.Get("user").Get("entities").Get("url")
            status.Get("user").GetStringBytes("created_at")
            status.Get("user").GetStringBytes("name")
            status.GetStringBytes("text")
        }
    })
}
