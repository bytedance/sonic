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
	"strings"
)

type FieldMap struct {
	oders  []string
	inner  map[string]int
	backup map[string]int
}

func CreateFieldMap(n int) *FieldMap {
	return &FieldMap{
		oders:  make([]string, n, n),
		inner:  make(map[string]int, n*2),
		backup: make(map[string]int, n*2),
	}
}

func (self *FieldMap) TryGet(name string, idx int) int {
	if idx < len(self.oders) && self.oders[idx] == name {
		return idx
	}
	return -1
}

func (self *FieldMap) Get(name string) int {
	if i, ok := self.inner[name]; ok {
		return i
	} else {
		return -1
	}
}

func (self *FieldMap) Set(name string, i int) {
	self.oders[i] = name
	self.inner[name] = i

	/* add the case-insensitive version, prefer the one with smaller field ID */
	key := strings.ToLower(name)
	if v, ok := self.backup[key]; !ok || i < v {
		self.backup[key] = i
	}
}

func (self *FieldMap) GetCaseInsensitive(name string) int {
	if i, ok := self.backup[strings.ToLower(name)]; ok {
		return i
	} else {
		return -1
	}
}
