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
	"testing"
	"time"

	"github.com/bytedance/sonic/internal/rt"
)

func TestPcacheRace(t *testing.T) {
	t.Parallel()

	pc := CreateProgramCache()
	wg := sync.WaitGroup{}
	wg.Add(2)
	start := make(chan struct{}, 2)

	go func() {
		defer wg.Done()
		var k = map[string]interface{}{}
		<-start
		for i := 0; i < 100; i++ {
			_, _ = pc.Compute(rt.UnpackEface(k).Type, func(*rt.GoType, ...interface{}) (interface{}, error) {
				return map[string]interface{}{}, nil
			})
		}
	}()

	go func() {
		defer wg.Done()
		var k = map[string]interface{}{}
		<-start
		for i := 0; i < 100; i++ {
			pc.Get(rt.UnpackEface(k).Type)
		}
	}()

	start <- struct{}{}
	start <- struct{}{}
	wg.Wait()
}

func TestPcacheComputesDifferentTypesConcurrently(t *testing.T) {
	t.Parallel()

	pc := CreateProgramCache()
	var a struct{ A int }
	var b struct{ B string }
	vtA := rt.UnpackEface(a).Type
	vtB := rt.UnpackEface(b).Type

	enteredA := make(chan struct{})
	unblockA := make(chan struct{})
	doneA := make(chan struct{})

	go func() {
		defer close(doneA)
		_, _ = pc.Compute(vtA, func(*rt.GoType, ...interface{}) (interface{}, error) {
			close(enteredA)
			<-unblockA
			return "a", nil
		})
	}()

	<-enteredA

	enteredB := make(chan struct{})
	doneB := make(chan struct{})
	go func() {
		defer close(doneB)
		_, _ = pc.Compute(vtB, func(*rt.GoType, ...interface{}) (interface{}, error) {
			close(enteredB)
			return "b", nil
		})
	}()

	select {
	case <-enteredB:
	case <-time.After(time.Second):
		t.Fatal("different type compute was blocked by unrelated in-flight compute")
	}

	close(unblockA)
	<-doneA
	<-doneB
}

func TestPcacheComputesSameTypeOnce(t *testing.T) {
	t.Parallel()

	pc := CreateProgramCache()
	var k map[string]interface{}
	vt := rt.UnpackEface(k).Type

	entered := make(chan struct{})
	unblock := make(chan struct{})
	done := make(chan interface{}, 2)
	var calls int32

	compute := func(*rt.GoType, ...interface{}) (interface{}, error) {
		if atomic.AddInt32(&calls, 1) == 1 {
			close(entered)
		}
		<-unblock
		return "program", nil
	}

	for i := 0; i < 2; i++ {
		go func() {
			val, _ := pc.Compute(vt, compute)
			done <- val
		}()
	}

	<-entered
	close(unblock)

	for i := 0; i < 2; i++ {
		if val := <-done; val != "program" {
			t.Fatalf("unexpected compute result: %v", val)
		}
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("compute called %d times, want 1", got)
	}
}

func TestPcacheResetDoesNotCacheInFlightCompute(t *testing.T) {
	t.Parallel()

	pc := CreateProgramCache()
	var k map[string]interface{}
	vt := rt.UnpackEface(k).Type

	entered := make(chan struct{})
	unblock := make(chan struct{})
	done := make(chan interface{})

	go func() {
		val, _ := pc.Compute(vt, func(*rt.GoType, ...interface{}) (interface{}, error) {
			close(entered)
			<-unblock
			return "program", nil
		})
		done <- val
	}()

	<-entered
	pc.Reset()
	close(unblock)

	if val := <-done; val != "program" {
		t.Fatalf("unexpected compute result: %v", val)
	}
	if val := pc.Get(vt); val != nil {
		t.Fatalf("in-flight value was cached after Reset: %v", val)
	}
}

func TestPcacheComputePanicReleasesPending(t *testing.T) {
	t.Parallel()

	pc := CreateProgramCache()
	var k map[string]interface{}
	vt := rt.UnpackEface(k).Type

	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected compute panic")
			}
		}()

		_, _ = pc.Compute(vt, func(*rt.GoType, ...interface{}) (interface{}, error) {
			panic("compile failed")
		})
	}()

	val, err := pc.Compute(vt, func(*rt.GoType, ...interface{}) (interface{}, error) {
		return "program", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if val != "program" {
		t.Fatalf("unexpected compute result: %v", val)
	}
}
