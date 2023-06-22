/**
 * Copyright 2023 ByteDance Inc.
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

package ast

type nodeChunk [_DEFAULT_NODE_CAP]Node

type linkedNodes struct {
	head   nodeChunk
	tail   []*nodeChunk
	size   int
}

func (self *linkedNodes) Cap() int {
	return (len(self.tail)+1)*_DEFAULT_NODE_CAP 
}

func (self *linkedNodes) Len() int {
	return self.size 
}

func (self *linkedNodes) At(i int) (*Node) {
	if i >= 0 && i<self.size && i < _DEFAULT_NODE_CAP {
		return &self.head[i]
	} else if i >= _DEFAULT_NODE_CAP && i<self.size  {
		a, b := i/_DEFAULT_NODE_CAP-1, i%_DEFAULT_NODE_CAP
		if a < len(self.tail) {
			return &self.tail[a][b]
		}
	}
	return nil
}

func (self *linkedNodes) Add(v Node) {
	if self.size < _DEFAULT_NODE_CAP {
		self.head[self.size] = v
		self.size++
		return
	}

	a, b, c := self.size/_DEFAULT_NODE_CAP-1 , self.size%_DEFAULT_NODE_CAP, cap(self.tail)
	if a - c >= 0 {
		c += 1 + c>>_APPEND_GROW_SHIFT
		tmp := make([]*nodeChunk, a + 1, c)
		copy(tmp, self.tail)
		self.tail = tmp
	} else if a >= len(self.tail) {
		self.tail = self.tail[:a+1]
	}
	
	var n = &self.tail[a]
	if *n == nil {
		*n = new(nodeChunk)
	}
	(*n)[b] = v
	self.size++
}

func (self *linkedNodes) ToSlice(con []Node) {
	if len(con) < self.size {
		return
	}
	a, b := self.size/_DEFAULT_NODE_CAP-1, self.size%_DEFAULT_NODE_CAP
	if a < 0 {
		copy(con, self.head[:self.size])
		return
	} else {
		copy(con, self.head[:])
		con = con[_DEFAULT_NODE_CAP:]
	}

	for i:=0; i<a; i++ {
		copy(con, self.tail[i][:])
		con = con[_DEFAULT_NODE_CAP:]
	}
	copy(con, self.tail[a][:b])
}

func (self *linkedNodes) FromSlice(con []Node) {
	self.size = len(con)
	a, b := self.size/_DEFAULT_NODE_CAP-1, self.size%_DEFAULT_NODE_CAP
	if a < 0 {
		copy(self.head[:self.size], con)
		return
	} else {
		copy(self.head[:], con)
		con = con[_DEFAULT_NODE_CAP:]
	}

	if cap(self.tail) <= a {
		c := (a+1) + (a+1)>>_APPEND_GROW_SHIFT
		self.tail = make([]*nodeChunk, a+1, c)
	}
	self.tail = self.tail[:a+1]

	for i:=0; i<a; i++ {
		self.tail[i] = new(nodeChunk)
		copy(self.tail[i][:], con)
		con = con[_DEFAULT_NODE_CAP:]
	}

	self.tail[a] = new(nodeChunk)
	copy(self.tail[a][:b], con)
}

type pairChunk [_DEFAULT_NODE_CAP]Pair

type linkedPairs struct {
	head pairChunk
	tail []*pairChunk
	size int
}

func (self *linkedPairs) Cap() int {
	return (len(self.tail)+1)*_DEFAULT_NODE_CAP 
}

func (self *linkedPairs) Len() int {
	return self.size 
}

func (self *linkedPairs) At(i int) *Pair {
	if i >= 0 && i < _DEFAULT_NODE_CAP && i<self.size {
		return &self.head[i]
	} else if i >= _DEFAULT_NODE_CAP && i<self.size {
		a, b := i/_DEFAULT_NODE_CAP-1, i%_DEFAULT_NODE_CAP
		if a < len(self.tail) {
			return &self.tail[a][b]
		}
	}
	return nil
}

func (self *linkedPairs) Add(v Pair) {
	if self.size < _DEFAULT_NODE_CAP {
		self.head[self.size] = v
		self.size++
		return
	}

	a, b, c := self.size/_DEFAULT_NODE_CAP-1 , self.size%_DEFAULT_NODE_CAP, cap(self.tail)
	if a - c >= 0 {
		c += 1 + c>>_APPEND_GROW_SHIFT
		tmp := make([]*pairChunk, a + 1, c)
		copy(tmp, self.tail)
		self.tail = tmp
	} else if a >= len(self.tail) {
		self.tail = self.tail[:a+1]
	}

	var n = &self.tail[a]
	if *n == nil {
		*n = new(pairChunk)
	}
	(*n)[b] = v
	self.size++
}

// linear search
func (self *linkedPairs) Get(key string) (*Pair, int) {
	for i:=0; i<self.size; i++ {
		if n := self.At(i); n.Key == key {
			return n, i
		}
	}
	return nil, -1
}

func (self *linkedPairs) ToSlice(con []Pair) {
	if len(con) < self.size {
		return
	}
	a, b := self.size/_DEFAULT_NODE_CAP-1, self.size%_DEFAULT_NODE_CAP
	if a < 0 {
		copy(con, self.head[:self.size])
		return
	} else {
		copy(con, self.head[:])
		con = con[_DEFAULT_NODE_CAP:]
	}

	for i:=0; i<a; i++ {
		copy(con, self.tail[i][:])
		con = con[_DEFAULT_NODE_CAP:]
	}
	copy(con, self.tail[a][:b])
}

func (self *linkedPairs) ToMap(con map[string]Node) {
	if len(con) < self.size {
		return
	}
	for i:=0; i<self.size; i++ {
		n := self.At(i)
		con[n.Key] = n.Value
	}
}

func (self *linkedPairs) FromSlice(con []Pair) {
	self.size = len(con)
	a, b := self.size/_DEFAULT_NODE_CAP-1, self.size%_DEFAULT_NODE_CAP
	if a < 0 {
		copy(self.head[:self.size], con)
		return
	} else {
		copy(self.head[:], con)
		con = con[_DEFAULT_NODE_CAP:]
	}

	if cap(self.tail) <= a {
		c := (a+1) + (a+1)>>_APPEND_GROW_SHIFT
		self.tail = make([]*pairChunk, a+1, c)
	}
	self.tail = self.tail[:a+1]

	for i:=0; i<a; i++ {
		self.tail[i] = new(pairChunk)
		copy(self.tail[i][:], con)
		con = con[_DEFAULT_NODE_CAP:]
	}

	self.tail[a] = new(pairChunk)
	copy(self.tail[a][:b], con)
}