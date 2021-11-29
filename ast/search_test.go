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
    `runtime`
    `runtime/debug`
    `sync`
    `fmt`
	`math`
    `strconv`

    `github.com/tidwall/sjson`
    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
    `github.com/tidwall/gjson`
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
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
            defer wg.Done()
            _, err := NewSearcher(_TwitterJson).GetByPath("statuses", 0, "id")
            if err != nil {
                t.Fatal(err)
            }
            runtime.GC()
            debug.FreeOSMemory()
        }(wg)
    }
    wg.Wait()
}

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
        x, _ := node.Int64()
        if x != 249279667666817024 {
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
            x, _ := node.Int64()
            if x != 249279667666817024 {
                b.Fatal(node.Interface())
            }
        }
    })
}

func BenchmarkSetOne_Sjson(b *testing.B) {
	path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
	_, err := sjson.Set(_TwitterJson, path, math.MaxInt32)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(_TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sjson.Set(_TwitterJson, path, math.MaxInt32)
	}
}

func BenchmarkSetOne_Jsoniter(b *testing.B) {
    data := []byte(_TwitterJson)
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

func BenchmarkSetOne_Parallel_Sjson(b *testing.B) {
	path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
	_, err := sjson.Set(_TwitterJson, path, math.MaxInt32)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(_TwitterJson)))
	b.ReportAllocs()
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            sjson.Set(_TwitterJson, path, math.MaxInt32)
        }
    })
}

func BenchmarkSetOne_Parallel_Jsoniter(b *testing.B) {
    data := []byte(_TwitterJson)
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