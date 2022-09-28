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
    `strings`
    `unsafe`
)

const (
    FieldMap_C = int64(unsafe.Offsetof(FieldMap{}.Count))
)

type FieldMap struct {
    TrieTree
    M map[string]FieldEntry
    All []FieldEntry
    MaxKeyLength int
}

type FieldEntry struct {
    ID   int
    Name string
}

func CreateFieldMap(n int) *FieldMap {
    return &FieldMap {
        TrieTree: TrieTree{},
        M: make(map[string]FieldEntry, n),
    }
}

// Get searches FieldMap by name. JIT generated assembly does NOT call this
// function, rather it implements its own version directly in assembly. So
// we must ensure this function stays in sync with the JIT generated one.
func (self *FieldMap) Get(name string) int {
    v := self.TrieTree.Get(name)
    if v == nil {
        return -1
    }
    return (*FieldEntry)(v).ID
}

func (self *FieldMap) Set(name string, i int) {
    if len(name) > self.MaxKeyLength {
        self.MaxKeyLength = len(name)
    }

    fi := FieldEntry{
        Name: name,
        ID:   i,
    }
    self.TrieTree.Set(name, unsafe.Pointer(&fi))

    /* add the case-insensitive version, prefer the one with smaller field ID */
    key := strings.ToLower(name)
    if v, ok := self.M[key]; !ok || i < v.ID {
        self.M[key] = fi
    }

    self.All = append(self.All, fi)
}

func (self *FieldMap) GetCaseInsensitive(name string) int {
    if i, ok := self.M[strings.ToLower(name)]; ok {
        return i.ID
    } else {
        return -1
    }
}

// Build calcaulates the best index for radix searching and reconstruct the trie tree
func (self *FieldMap) Build() {
	var empty unsafe.Pointer

    // map every fields under the same char for either index j (backward)
	var charCount = make([]map[byte][]int, self.MaxKeyLength)
	for i, v := range self.All {
		for j := self.MaxKeyLength - 1; j >= 0; j-- {
			if v.Name == "" {
				empty = unsafe.Pointer(&v)
			}
			var c = byte(0)
			if j < len(v.Name) {
				c = v.Name[j]
			}
			if charCount[j] == nil {
				charCount[j] = make(map[byte][]int, 16)
			}
			charCount[j][c] = append(charCount[j][c], i)
		}
	}

    if empty != nil {
        self.Empty = empty
    }

	var idealPos = 0
	var minF = float64(len(self.All))

    // find the best position to split the trie tree (fieldCount/charCount closest to 1)
	for i := self.MaxKeyLength - 1; i >= 0; i-- {
		cd := charCount[i]
		charCount := len(cd)
		fieldCount := 0
		for _, v := range cd {
			fieldCount += len(v)
		}
		f := float64(fieldCount) / float64(charCount)
		if f < minF {
			minF = f
			idealPos = i
		}
		if minF == 1 {
			break
		}
	}

    // record the index position for the tree.Get()
    self.Positions = append(self.Positions, idealPos)

    // build the trie tree again
    for _, v := range self.All {
        self.Set(v.Name, v.ID)
    }

	return
}