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
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

/** Program Map **/

const (
	_LoadFactor   = 0.5
	_InitCapacity = 4096 // must be a power of 2
)

type _ProgramMap struct {
	n uint64
	m uint32
	b []_ProgramEntry
}

type _ProgramEntry struct {
	vt *rt.GoType
	fn interface{}
}

func newProgramMap() *_ProgramMap {
	return &_ProgramMap{
		n: 0,
		m: _InitCapacity - 1,
		b: make([]_ProgramEntry, _InitCapacity),
	}
}

func (self *_ProgramMap) copy() *_ProgramMap {
	fork := &_ProgramMap{
		n: self.n,
		m: self.m,
		b: make([]_ProgramEntry, len(self.b)),
	}
	for i, f := range self.b {
		fork.b[i] = f
	}
	return fork
}

func (self *_ProgramMap) get(vt *rt.GoType) interface{} {
	i := self.m + 1
	p := vt.Hash & self.m

	/* linear probing */
	for ; i > 0; i-- {
		if b := self.b[p]; b.vt == vt {
			return b.fn
		} else if b.vt == nil {
			break
		} else {
			p = (p + 1) & self.m
		}
	}

	/* not found */
	return nil
}

func (self *_ProgramMap) add(vt *rt.GoType, fn interface{}) *_ProgramMap {
	p := self.copy()
	f := float64(atomic.LoadUint64(&p.n)+1) / float64(p.m+1)

	/* check for load factor */
	if f > _LoadFactor {
		p = p.rehash()
	}

	/* insert the value */
	p.insert(vt, fn)
	return p
}

func (self *_ProgramMap) rehash() *_ProgramMap {
	c := (self.m + 1) << 1
	r := &_ProgramMap{m: c - 1, b: make([]_ProgramEntry, int(c))}

	/* rehash every entry */
	for i := uint32(0); i <= self.m; i++ {
		if b := self.b[i]; b.vt != nil {
			r.insert(b.vt, b.fn)
		}
	}

	/* rebuild successful */
	return r
}

func (self *_ProgramMap) insert(vt *rt.GoType, fn interface{}) {
	h := vt.Hash
	p := h & self.m

	/* linear probing */
	for i := uint32(0); i <= self.m; i++ {
		if b := &self.b[p]; b.vt != nil {
			p += 1
			p &= self.m
		} else {
			b.vt = vt
			b.fn = fn
			atomic.AddUint64(&self.n, 1)
			return
		}
	}

	/* should never happens */
	panic("no available slots")
}

/** RCU Program Cache **/

type ProgramCache struct {
	m       sync.Mutex
	p       unsafe.Pointer
	pending map[*rt.GoType]*programCall
	epoch   uint64
}

type programCall struct {
	done       chan struct{}
	epoch      uint64
	val        interface{}
	err        error
	panicked   bool
	panicValue interface{}
}

func CreateProgramCache() *ProgramCache {
	return &ProgramCache{
		m:       sync.Mutex{},
		p:       unsafe.Pointer(newProgramMap()),
		pending: make(map[*rt.GoType]*programCall),
	}
}

func (self *ProgramCache) Reset() {
	self.m.Lock()
	defer self.m.Unlock()
	self.p = unsafe.Pointer(newProgramMap())
	self.pending = make(map[*rt.GoType]*programCall)
	self.epoch++
}

func (self *ProgramCache) Get(vt *rt.GoType) interface{} {
	return (*_ProgramMap)(atomic.LoadPointer(&self.p)).get(vt)
}

func (self *ProgramCache) Compute(vt *rt.GoType, compute func(*rt.GoType, ...interface{}) (interface{}, error), ex ...interface{}) (interface{}, error) {
	var err error
	var val interface{}

	if val = self.Get(vt); val != nil {
		return val, nil
	}

	self.m.Lock()

	/* double check with write lock held */
	if val = self.Get(vt); val != nil {
		self.m.Unlock()
		return val, nil
	}

	if call := self.pending[vt]; call != nil {
		self.m.Unlock()
		<-call.done
		if call.panicked {
			panic(call.panicValue)
		}
		return call.val, call.err
	}

	call := &programCall{
		done:  make(chan struct{}),
		epoch: self.epoch,
	}
	self.pending[vt] = call
	self.m.Unlock()

	/* compute the value */
	func() {
		defer func() {
			if r := recover(); r != nil {
				call.panicked = true
				call.panicValue = r
			}
		}()
		call.val, call.err = compute(vt, ex...)
	}()

	self.m.Lock()

	/* update the RCU cache */
	if !call.panicked && call.err == nil && call.epoch == self.epoch {
		atomic.StorePointer(&self.p, unsafe.Pointer((*_ProgramMap)(atomic.LoadPointer(&self.p)).add(vt, call.val)))
	}
	if self.pending[vt] == call {
		delete(self.pending, vt)
	}
	close(call.done)

	val = call.val
	err = call.err
	panicked := call.panicked
	panicValue := call.panicValue
	self.m.Unlock()

	if panicked {
		panic(panicValue)
	}
	return val, err
}
