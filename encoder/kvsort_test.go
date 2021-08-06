package encoder

import (
    "bytes"
    "math/rand"
    "reflect"
    "sort"
    "testing"
    "unsafe"
)

// Make kvSlice meet sort.Interface.
func (x kvSlice) Less(i, j int) bool { return bytes.Compare(x[i].k, x[j].k) < 0 }
func (x kvSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x kvSlice) Len() int           { return len(x) }

var keyLen = 15

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
