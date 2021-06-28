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

package ast

type Pair struct {
	Key   string
	Value Node
}

type Iterator struct {
	i int
	p *Node
}

func (self *Iterator) Pos() int {
	return self.i
}

func (self *Iterator) Len() int {
	return self.p.Len()
}

func (self *Iterator) HasNext() bool {
	return self.i < self.p.Len()
}

type ListIterator struct {
	Iterator
}

type ObjectIterator struct {
	Iterator
}

func (self *ListIterator) Next(v *Node) bool {
	if !self.HasNext() {
		return false
	} else {
		*v, self.i = *self.p.nodeAt(self.i), self.i+1
		return true
	}
}

func (self *ObjectIterator) Next(p *Pair) bool {
	if !self.HasNext() {
		return false
	} else {
		*p, self.i = *self.p.pairAt(self.i), self.i+1
		return true
	}
}
