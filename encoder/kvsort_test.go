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

package encoder

import (
	"bytes"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"unsafe"
)

type encodedKeyValues []encodedKV
type encodedKV struct {
	key      string
	keyValue []byte
}

func (sv encodedKeyValues) Len() int           { return len(sv) }
func (sv encodedKeyValues) Swap(i, j int)      { sv[i], sv[j] = sv[j], sv[i] }
func (sv encodedKeyValues) Less(i, j int) bool { return sv[i].key < sv[j].key }

func getKvs(std bool) interface{} {
    var map_size = 1000
    if std {
        kvs := make(encodedKeyValues, map_size)
        for i:=map_size-1; i>=0; i-- {
            kvs[i] = encodedKV{
                key:"\"test_" + strconv.Itoa(i) + "\"",
                // key: string(genRandKey(10)),
            }
        }
        return kvs
    }else{
        kvs := make(kvSlice, map_size)
        for i:=map_size-1; i>=0; i-- {
            kvs[i] = keyValue{
                k: []byte("\"test_" + strconv.Itoa(i) + "\""),
                // k: genRandKey(10),
            }
        }
        return kvs
    }
}

func BenchmarkSort_Sonic(b *testing.B) {
    kvs := getKvs(false).(kvSlice)
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        radixQsort(kvs, 0, maxDepth(len(kvs)))
    }
}

func BenchmarkSort_Insert(b *testing.B) {
    kvs := getKvs(false).(kvSlice)
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        insertRadixSort(kvs, 0)
    }
}

func BenchmarkSort_Std(b *testing.B) {
    kvs := getKvs(true).(encodedKeyValues)
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        sort.Sort(kvs)
    }
}

func BenchmarkSort_Parallel_Sonic(b *testing.B) {
    kvs := getKvs(false).(kvSlice)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        for p.Next() {
            radixQsort(kvs, 0, maxDepth(len(kvs)))
        }
    })
}

func BenchmarkSort_Parallel_Insert(b *testing.B) {
    kvs := getKvs(false).(kvSlice)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        for p.Next() {
            insertRadixSort(kvs, 0)
        }
    })
}

func BenchmarkSort_Parallel_Std(b *testing.B) {
    kvs := getKvs(true).(encodedKeyValues)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        for p.Next() {
            sort.Sort(kvs)
        }
    })
}

// Make kvSlice meet sort.Interface.
func (x kvSlice) Less(i, j int) bool { return bytes.Compare(x[i].k, x[j].k) < 0 }
func (x kvSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x kvSlice) Len() int           { return len(x) }

func TestSortRadixRandKvs(t *testing.T) {
    kvs := getRandKvs(1000, 15)
    sorted := make([]keyValue, len(kvs))
    for i := range(sorted) {
        k := make([]byte, len(kvs[i].k))
        copy(k, kvs[i].k)
        sorted[i].k = k[:]
        sorted[i].v = kvs[i].v
    }
    sort.Sort(kvSlice(sorted))
    kvs.Sort()

    if !reflect.DeepEqual(kvs, kvSlice(sorted)) {
         t.Errorf(" got: %v\nwant: %v\n", kvs, kvSlice(sorted))
    }
}

func TestSortInsertRandKvs(t *testing.T) {
    kvs := getRandKvs(10, 15)
    sorted := make([]keyValue, len(kvs))
    for i := range(sorted) {
        k := make([]byte, len(kvs[i].k))
        copy(k, kvs[i].k)
        sorted[i].k = k[:]
        sorted[i].v = kvs[i].v
    }

    sort.Sort(kvSlice(sorted))
    insertRadixSort(kvs, 0)

    if !reflect.DeepEqual(kvs, kvSlice(sorted)) {
         t.Errorf(" got: %v\nwant: %v\n", kvs, kvSlice(sorted))
    }
}

func genRandKey(kl int) []byte {
    l := int(rand.Uint32()%uint32(kl) + 2)
    k := make([]byte, l)
    k[0], k[l-1] = '"', '"'
    for i := 1; i < l-1; i++ {
        k[i] = byte('a' + int(rand.Uint32()%26))
    }
    return k
}

func getRandKvs(kn int, kl int) kvSlice {
    keys := make(map[string]bool)
    kvs := make([]keyValue, 0)
    for len(keys) < kn {
        k := genRandKey(kl)
        keys[string(k)] = true
    }
    for k := range keys {
        var kv keyValue
        kv.k = []byte(k)
        kv.v = unsafe.Pointer(&k)
        kvs = append(kvs, kv)
    }
    return kvs
}
