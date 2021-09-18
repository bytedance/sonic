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
    `bytes`
    `math/rand`
    `reflect`
    `sort`
    `strconv`
    `testing`
    `unsafe`
)

var keyLen = 15

type encodedKeyValues []encodedKV
type encodedKV struct {
	key      string
	_MapPair []byte
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
                key: "\"test_" + strconv.Itoa(i) + "\"",
            }
        }
        return kvs
    } else {
        kvs := make([]_MapPair, map_size)
        for i:=map_size-1; i>=0; i-- {
            kvs[i] = _MapPair{
                k: "\"test_" + strconv.Itoa(i) + "\"",
            }
        }
        return kvs
    }
}

func BenchmarkSort_Sonic(b *testing.B) {
    ori := getKvs(false).([]_MapPair)
    kvs := make([]_MapPair, len(ori))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        copy(kvs, ori)
        radixQsort(kvs, 0, maxDepth(len(kvs)))
    }
}

func BenchmarkSort_Std(b *testing.B) {
    ori := getKvs(true).(encodedKeyValues)
    kvs := make(encodedKeyValues, len(ori))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        copy(kvs, ori)
        sort.Sort(kvs)
    }
}

func BenchmarkSort_Parallel_Sonic(b *testing.B) {
    ori := getKvs(false).([]_MapPair)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        kvs := make([]_MapPair, len(ori))
        for p.Next() {
            copy(kvs, ori)
            radixQsort(kvs, 0, maxDepth(len(kvs)))
        }
    })
}

func BenchmarkSort_Parallel_Std(b *testing.B) {
    ori := getKvs(true).(encodedKeyValues)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        kvs := make(encodedKeyValues, len(ori))
        for p.Next() {
            copy(kvs, ori)
            sort.Sort(kvs)
        }
    })
}

type kvSlice []_MapPair

// Make kvSlice meet sort.Interface.
func (self kvSlice) Less(i, j int) bool { return self[i].k < self[j].k }
func (self kvSlice) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self kvSlice) Len() int           { return len(self) }

//go:nosplit
func (self kvSlice) Sort() {
    radixQsort(self, 0, maxDepth(len(self)))
}

func (self kvSlice) String() string {
    buf := bytes.NewBuffer(nil)
    for i, kv := range self {
        if i > 0 {
            buf.WriteByte(',')
        }
        buf.WriteString(kv.k)
    }
    return buf.String()
}

func TestSort_SortRandomKeys(t *testing.T) {
    kvs := getRandKvs(100, keyLen)
    sorted := make([]_MapPair, len(kvs))

    copy(sorted, kvs)
    sort.Sort(kvSlice(sorted))
    kvs.Sort()

    got := kvs.String()
    want := kvSlice(sorted).String()
    if !reflect.DeepEqual(got, want) {
        t.Errorf(" got: %v\nwant: %v\n", got, want)
    }
}

func genKey(kl int) []byte {
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
    kvs := make(kvSlice, 0)
    for len(keys) < kn {
        k := genKey(kl)
        keys[string(k)] = true
    }
    for k := range keys {
        var kv _MapPair
        kv.k = k
        kv.v = unsafe.Pointer(&k)
        kvs = append(kvs, kv)
    }
    return kvs
}
