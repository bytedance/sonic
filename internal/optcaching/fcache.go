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

	"github.com/bytedance/sonic/internal/resolver"
)

type FieldCache struct {
	caseSensitive   map[string]int
	caseInsensitive map[string]int
}

func NewFieldCache(fields []resolver.FieldMeta) *FieldCache {
	f := &FieldCache{
		caseSensitive:   make(map[string]int, len(fields)),
		caseInsensitive: make(map[string]int, len(fields)),
	}
	for i, field := range fields {
		name := field.Name
		f.caseSensitive[name] = i
		lowerName := strings.ToLower(name)
		if existingIdx, ok := f.caseInsensitive[lowerName]; !ok || i < existingIdx {
			f.caseInsensitive[lowerName] = i
		}
	}
	return f
}

func (f *FieldCache) Get(name string, caseSensitive bool) int {
	if idx, ok := f.caseSensitive[name]; ok {
		return idx
	}
	if !caseSensitive {
		if idx, ok := f.caseInsensitive[strings.ToLower(name)]; ok {
			return idx
		}
	}
	return -1
}
