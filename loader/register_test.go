/*
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

package loader

import (
    "fmt"
    "testing"
    "sync"
)

// Test_registerModuleChainReachable checks the invariant the runtime relies on
// while a module is being appended: every module that is already published as
// the tail must be reachable from the first module, otherwise runtime.findfunc
// cannot resolve frames of that module and stack growth aborts the process
// with "unknown caller pc".
func Test_registerModuleChainReachable(t *testing.T) {
    const (
        perGoroutine = 100
        parallel     = 8
        total        = perGoroutine * parallel
    )

    head := &moduledata{}
    registerModule(head)

    done := make(chan struct{})
    var once sync.Once
    var failure error
    fail := func(err error) {
        once.Do(func() { failure = err })
    }

    watcher := func() {
        for {
            select {
            case <-done:
                return
            default:
            }
            tail := loadModule(&lastmoduledatap)
            reachable := tail == head
            n := 0
            for p := head; !reachable; p = loadModule(&p.next) {
                n++
                if p == nil || n > total+1 {
                    fail(fmt.Errorf("the last module %p is not reachable from the first module "+
                        "after %d modules: it was published before it was linked", tail, n))
                    return
                }
                reachable = p == tail
            }
        }
    }

    go watcher()

    wg := sync.WaitGroup{}
    wg.Add(parallel)
    for i := 0; i < parallel; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < perGoroutine; j++ {
                registerModule(&moduledata{})
            }
        }()
    }
    wg.Wait()
    close(done)

    if failure != nil {
        t.Fatal(failure)
    }
}

func Test_registerModuleRace(t *testing.T) {
    n, parallel := 100, 8
    s := &moduledata{}
    registerModule(s)
    wg := sync.WaitGroup{}
    wg.Add(parallel)
    filler := func(n int) {
        defer wg.Done()
        for i := 0; i < n; i++ {
            registerModule(&moduledata{})
        }
    }
    for i := 0; i < parallel; i++ {
        go filler(n)
    }
    wg.Wait()
    i := 0
    for p := s.next; p != nil; p = p.next {
        i += 1
    }
    if i != parallel * n {
        t.Errorf("got %v, expected %v", i, parallel * n)
    }
}
