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

package caching

import (
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

const (
	charPoint = '.'

	fieldTrieTreeSize = unsafe.Sizeof(TrieNode{})

	defaultTrieLeavesSize = 4

	byteTypeSize = unsafe.Sizeof(byte(0))
)

func ascii2Int(c byte) uint8 {
	if c < charPoint {
		return uint8(255) + uint8(c) - uint8(charPoint)
	}
	return uint8(c) - uint8(charPoint)
}

type Pair struct {
	Val unsafe.Pointer
	Key string
}

type TrieTree struct {
	Count     int
	Positions []int
	Empty     unsafe.Pointer
	TrieNode
}

type TrieNode struct {
	Leaves *[]Pair
	Index  []TrieNode
}

func (fn *TrieTree) Set(k string, field unsafe.Pointer) bool {
	if k == "" {
		exist := fn.Empty != nil
		if !exist {
			fn.Count++
		}
		fn.Empty = field
		return exist
	}

	var l = len(k) - 1
	var ks = *(*unsafe.Pointer)(unsafe.Pointer(&k))
	var fs = fn.Index
	if fn.Leaves == nil {
		x := make([]Pair, 0, defaultTrieLeavesSize)
		fn.Leaves = &x
	}
	var ls = *fn.Leaves
	var fp = &fn.TrieNode
	var count = 1

	for _, i := range fn.Positions {
		if i > l {
			i = l
		}
		c := *(*byte)(rt.IndexPtr(ks, byteTypeSize, i))
		j := ascii2Int(c)
		if int(j) >= len(fs) {
			tmp := make([]TrieNode, j+1)
			copy(tmp, fs)
			fs = tmp
			fp.Index = tmp
		}
		fp = (*TrieNode)(rt.IndexPtr(unsafe.Pointer(&fs[0]), fieldTrieTreeSize, int(j)))
		fs = fp.Index
		if fp.Leaves == nil {
			x := make([]Pair, 0, defaultTrieLeavesSize)
			fp.Leaves = &x
		}
		ls = *fp.Leaves
		count++
	}

	for j, v := range ls {
		if k == v.Key {
			ls[j].Val = field
			return true
		}
	}
	ls = append(ls, Pair{field, k})
	fp.Leaves = &ls
	fn.Count++
	return false
}

func (fn *TrieTree) Size() int {
	return fn.Count
}

func (fn *TrieTree) Get(k string) unsafe.Pointer {
	if k == "" {
		return fn.Empty
	}

	var l = len(k) - 1
	var ks = *(*unsafe.Pointer)(unsafe.Pointer(&k))
	var fs = fn.Index
	var ls = *fn.Leaves

	for _, i := range fn.Positions {
		if i > l {
			i = l
		}
		c := *(*byte)(rt.IndexPtr(ks, byteTypeSize, i))
		j := ascii2Int(c)
		if int(j) >= len(fs) {
			return nil
		}
		fp := (*TrieNode)(rt.IndexPtr(unsafe.Pointer(&fs[0]), fieldTrieTreeSize, int(j)))
		// fp.Leaves is nil will not gonna happen!
		if fp.Leaves == nil {
			return nil
		}
		fs = fp.Index
		ls = *fp.Leaves
	}

	for _, v := range ls {
		if k == v.Key {
			return v.Val
		}
	}
	return nil
}
