/*
 * Copyright 2025 Huawei Technologies Co., Ltd.
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

package native

import (
	"os"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/testdata"
)

/* how to use it
SONIC_NO_ASYNC_GC=1 SONIC_USE_SVE_LINKNAME=1  taskset -c 32-47 go test ./internal/native/ -bench=. -benchtime=500000x
*/

type node struct {
	typ uint64
	val uint64
}

type jsonStat struct {
	object      uint32
	array       uint32
	str         uint32
	number      uint32
	array_elems uint32
	object_keys uint32
	max_depth   uint32
}

type nodeBuf struct {
	ncur   uintptr
	parent int64
	depth  uint64
	nstart uintptr
	nend   uintptr
	iskey  bool
	stat   jsonStat
}

type Parser struct {
	Json    string
	padded  []byte
	nodes   []node
	dbuf    []byte
	backup  []node
	options uint64
	start   uintptr
	cur     uintptr
	end     uintptr
	nbk     [16]byte
	nbuf    nodeBuf
	Utf8Inv bool
	isEface bool
}

func loadTwitterJson() string {
	data, err := os.ReadFile("../../testdata/twitterescaped.json")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func BenchmarkParseWithPadding(b *testing.B) {
	b.Run("Complex", func(b *testing.B) {
		s := loadTwitterJson()
		benchmarkParseWithPadding(b, s, 1024*1024)
	})
	b.Run("Medium", func(b *testing.B) {
		s := testdata.TwitterJson
		benchmarkParseWithPadding(b, s, 1000)
	})
}

func benchmarkParseWithPadding(b *testing.B, s string, nodeCap int) {
	// Allocate fresh buffers for each iteration
	padded := make([]byte, len(s)+64)
	copy(padded, s)
	for j := len(s); j < len(padded); j++ {
		padded[j] = 'x'
	}
	nodes := make([]node, nodeCap)
	dbuf := make([]byte, 800)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create parser
		parser := &Parser{
			Json:   s,
			nodes:  nodes,
			dbuf:   dbuf,
		}

		parser.padded = append(parser.padded, padded...)
		parser.start = uintptr(unsafe.Pointer(&parser.padded[0]))
		parser.cur= parser.start
		parser.end =  uintptr(unsafe.Pointer(&parser.padded[len(s)]))

		// Initialize node buffer
		parser.nbuf.ncur = uintptr(unsafe.Pointer(&nodes[0]))
		parser.nbuf.nstart = parser.nbuf.ncur
		parser.nbuf.nend = parser.nbuf.ncur + uintptr(cap(nodes))*unsafe.Sizeof(node{})
		parser.nbuf.parent = -1
		parser.nbuf.depth = 0
		parser.nbuf.iskey = false
		parser.nbuf.stat = jsonStat{}

		ret := ParseWithPadding(unsafe.Pointer(parser))
		if ret != 0 {
			b.Fatalf("ParseWithPadding failed with code %d", ret)
		}
	}
}

func BenchmarkGetByPath(b *testing.B) {
	s := loadTwitterJson()
	p := 0
	path := []interface{}{"statuses", 3, "entities", "user_mentions", 0, "screen_name"} // get the id of the 4th tweet
	m := types.NewStateMachine()
	defer types.FreeStateMachine(m)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = 0 // reset position
		ret := GetByPath(&s, &p, &path, m)
		if ret < 0 {
			b.Fatalf("GetByPath failed with code %d", ret)
		}
	}
}
