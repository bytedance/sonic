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
