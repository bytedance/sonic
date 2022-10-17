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
    `fmt`
    `strconv`
    `testing`

    `github.com/stretchr/testify/assert`
)

func getTestIteratorSample(loop int) (string, int) {
    var data []int
    var v1 = ""
    var v2 = ""
    for i:=0;i<loop;i++{
        data = append(data, i*i)
        v1 += strconv.Itoa(i)
        v2 += `"k`+strconv.Itoa(i)+`":`+strconv.Itoa(i)
        if i!=loop-1{
            v1+=`,`
            v2+=`,`
        }
    }
    return `{"array":[`+v1+`], "object":{`+v2+`}}`, loop
}

func TestForEach(t *testing.T) {
    pathes := []Sequence{}
    values := []*Node{}
    sc := func(path Sequence, node *Node) bool {
        pathes = append(pathes, path)
        values = append(values, node)
        if path.Key != nil && *path.Key == "array" {
            node.ForEach(func(path Sequence, node *Node)bool{
                pathes = append(pathes, path)
                values = append(values, node)
                return true
            })
        }
        return true
    }

    str, _ := getTestIteratorSample(3)
    fmt.Println(str)
    root, err := NewSearcher(str).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    err = root.ForEach(sc)
    if err != nil {
        t.Fatal(err)
    }
    eObjKey := "object"
    eArrKey := "array"
    expPath := []Sequence{
        {0, &eArrKey},
        {0, nil},
        {1, nil},
        {2, nil},
        {1, &eObjKey},
    }
    expValue := []*Node{
        root.Get("array"),
        root.GetByPath("array", 0),
        root.GetByPath("array", 1),
        root.GetByPath("array", 2),
        root.Get("object"),
    }
    // fmt.Printf("pathes:%+v\n", pathes)
    // fmt.Printf("values:%+v\n", values)
    assert.Equal(t, expPath, pathes)
    assert.Equal(t, expValue, values)
    
}

func TestRawIterator(t *testing.T) {
    str, loop := getTestIteratorSample(_DEFAULT_NODE_CAP)
    fmt.Println(str)
    
    root, err := NewSearcher(str).GetByPath("array")
    if err != nil {
        t.Fatal(err)
    }
    ai, _ := root.Values()
    i := int64(0)
    for ai.HasNext() {
        v := &Node{}
        if !ai.Next(v) {
            t.Fatalf("no next")
        }
        x, _ := v.Int64()
        if i < int64(loop) && x != i {
            t.Fatalf("exp:%v, got:%v", i, v)
        }
        if i != int64(ai.Pos())-1 || i >= int64(ai.Len()) {
            t.Fatal(i)
        }
        i++
    }
    if i != int64(loop) {
        t.Fatal(i)
    }
    
    root, err = NewSearcher(str).GetByPath("object")
    if err != nil {
        t.Fatal(err)
    }
    mi, _ := root.Properties()
    i = int64(0)
    for mi.HasNext() {
        v := &Pair{}
        if !mi.Next(v) {
            t.Fatalf("no next")
        }
        x, _ := v.Value.Int64()
        if i < int64(loop) &&( x != i ||v.Key != fmt.Sprintf("k%d", i)) {
            vv, _ := v.Value.Interface()
            t.Fatalf("exp:%v, got:%v", i, vv)
        }
        if i != int64(mi.Pos())-1 || i >= int64(mi.Len()) {
            t.Fatal(i)
        }
        i++
    }
    if i != int64(loop) {
        t.Fatal(i)
    }
}

func TestIterator(t *testing.T) {
    str, loop := getTestIteratorSample(_DEFAULT_NODE_CAP)
    fmt.Println(str)

    root, err := NewParser(str).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    ai, _ := root.Get("array").Values()
    i := int64(0)
    for ai.HasNext() {
        v := &Node{}
        if !ai.Next(v) {
            t.Fatalf("no next")
        }
        x, _ := v.Int64()
        if i < int64(loop) && x != i {
            t.Fatalf("exp:%v, got:%v", i, v)
        }
        if i != int64(ai.Pos())-1 || i >= int64(ai.Len()) {
            t.Fatal(i)
        }
        i++
    }
    if i != int64(loop) {
        t.Fatal(i)
    }

    root, err = NewParser(str).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    mi, _ := root.Get("object").Properties()
    i = int64(0)
    for mi.HasNext() {
        v := &Pair{}
        if !mi.Next(v) {
            t.Fatalf("no next")
        }
        x, _ := v.Value.Int64()
        if i < int64(loop) &&( x != i ||v.Key != fmt.Sprintf("k%d", i)) {
            vv, _ := v.Value.Interface()
            t.Fatalf("exp:%v, got:%v", i, vv)
        }
        if i != int64(mi.Pos())-1 || i >= int64(mi.Len()) {
            t.Fatal(i)
        }
        i++
    }
    if i != int64(loop) {
        t.Fatal(i)
    }

    str, _ = getTestIteratorSample(0)
    root, err = NewParser(str).Parse()
    if err != 0 {
        t.Fatal(err)
    }
    mi, _ = root.Get("object").Properties()
    if mi.HasNext() {
        t.Fatalf("should not have next")
    }
}

func BenchmarkArrays(b *testing.B) {
    for i:=0;i<b.N;i++{
        root,err := NewSearcher(_TwitterJson).GetByPath("statuses",1,"entities","hashtags")
        if err != nil {
            b.Fatal(err)
        }
        a, _ := root.Array()
        for _,v := range a {
            _ = v
        }
    }
}

func BenchmarkListIterator(b *testing.B) {
    for i:=0;i<b.N;i++{
        root,err := NewSearcher(_TwitterJson).GetByPath("statuses",1,"entities","hashtags")
        if err != nil {
            b.Fatal(err)
        }
        it, _ := root.Values()
        for it.HasNext() {
            v := &Node{}
            if !it.Next(v) {
                b.Fatalf("no value")
            }
        }
    }
}

func BenchmarkMap(b *testing.B) {
    for i:=0;i<b.N;i++{
        root,err := NewSearcher(_TwitterJson).GetByPath("statuses",1, "user")
        if err != nil {
            b.Fatal(err)
        }
        m, _ := root.Map()
        for k,v := range m {
            _ = v
            _ = k
        }
    }
}

func BenchmarkObjectIterator(b *testing.B) {
    for i:=0;i<b.N;i++{
        root,err := NewSearcher(_TwitterJson).GetByPath("statuses",1, "user")
        if err != nil {
            b.Fatal(err)
        }
        it, _ := root.Properties()
        for it.HasNext() {
            v := &Pair{}
            if !it.Next(v)  {
                b.Fatalf("no value")
            }
        }
    }
}