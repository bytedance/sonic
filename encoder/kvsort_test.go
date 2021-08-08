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

// Make kvSlice meet sort.Interface.
func (x kvSlice) Less(i, j int) bool { return bytes.Compare(x[i].k, x[j].k) < 0 }
func (x kvSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x kvSlice) Len() int           { return len(x) }

var keyLen = 15

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
                key: "test_" + strconv.Itoa(i),
            }
        }
        return kvs
    }else{
        kvs := make(kvSlice, map_size)
        for i:=map_size-1; i>=0; i-- {
            kvs[i] = keyValue{
                k: []byte("test_" + strconv.Itoa(i)),
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

func BenchmarkSort_Parallel_Std(b *testing.B) {
    kvs := getKvs(true).(encodedKeyValues)
    b.ResetTimer()
    b.RunParallel(func(p *testing.PB) {
        for p.Next() {
            sort.Sort(kvs)
        }
    })
}

func TestSortRandKvs(t *testing.T) {
    kvs := getRandKvs(100, keyLen)
    sorted := make([]keyValue, len(kvs))

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
    l := int(rand.Uint32()%uint32(kl) + 1)
    k := make([]byte, l)
    for i := 0; i < l; i++ {
        k[i] = byte('a' + int(rand.Uint32()%26))
    }
    return k
}

func getRandKvs(kn int, kl int) kvSlice {
    keys := make(map[string]bool)
    kvs := make([]keyValue, 0)
    for len(keys) < kn {
        k := genKey(kl)
        keys[string(k)] = true
    }
    for k := range keys {
        var kv keyValue
        kv.k = []byte(k)
        kv.v = unsafe.Pointer(&k)
        kvs = append(kvs, kv)
    }
    return kvs[:]
}
