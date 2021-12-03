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

import (
    `github.com/bytedance/sonic/internal/native/types`
)

type Pair struct {
    Key   string
    Value Node
}

// type Scanner = func (key interface{}, val Node) bool

// func (self *Node) ForEach(f Scanner) error {
//     if err := self.checkRaw(); err != nil {
//         return unwrapError(err)
//     }

//     it := self.itype()
//     if it == types.V_ARRAY {

//         if err := self.skipAllIndex(); err != nil {
//             return err
//         }
//         var p = (*Node)(self.p)
//         for i := 0; i < self.len(); i++ {
//             p = p.unsafe_next()
//             if err := p.Check(); err != nil {
//                 return err
//             }
//             if !f(i, *p) {
//                 return nil
//             }
//         }

//     }else if it == types.V_OBJECT {
        
//         if err := self.skipAllKey(); err != nil {
//             return err
//         }
//         var p = (*Pair)(self.p)
//         for i := 0; i < self.len(); i++ {
//             p = p.unsafe_next()
//             if err := p.Value.Check(); err != nil {
//                 return err
//             }
//             if !f(p.Key, p.Value) {
//                 return nil
//             }
//         }

//     }else{

//         f(nil, *self)
//     }

//     return nil
// }

// Values returns iterator for array's children traversal
func (self *Node) Values() (ListIterator, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return ListIterator{}, err
    }
    return ListIterator{Iterator{p: self}}, nil
}

// Properties returns iterator for object's children traversal
func (self *Node) Properties() (ObjectIterator, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return ObjectIterator{}, err
    }
    return ObjectIterator{Iterator{p: self}}, nil
}

type Iterator struct {
    i int
    p *Node
}

func (self *Iterator) Pos() int {
    return self.i
}

func (self *Iterator) Len() int {
    return self.p.len()
}

func (self *Iterator) HasNext() bool {
    if !self.p.isLazy() {
        return self.i < self.p.len() && self.p.Valid()
    } else if self.p.t == _V_ARRAY_LAZY {
        return self.p.skipNextNode().Valid()
    } else if self.p.t == _V_OBJECT_LAZY {
        return self.p.skipNextPair().Value.Valid()
    }
    return false
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
        *v, self.i = *self.p.nodeAt(self.i), self.i + 1
        return true
    }
}

func (self *ObjectIterator) Next(p *Pair) bool {
    if !self.HasNext() {
        return false
    } else {
        *p, self.i = *self.p.pairAt(self.i), self.i + 1
        return true
    }
}
