package rt

import (
	"sync"
	"testing"
)

var poolTest = sync.Pool{
	New:  newBytes,
}

func newBytes() interface{}  {
	return make([]byte, 1024)
}

func BenchmarkSyncPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		poolTest.Get()
	}
}

func BenchmarkDirectAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newBytes()
	}
}