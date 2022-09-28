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
}

func (self *FieldMap) GetCaseInsensitive(name string) int {
    if i, ok := self.M[strings.ToLower(name)]; ok {
        return i.ID
    } else {
        return -1
    }
}