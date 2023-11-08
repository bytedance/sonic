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
    "testing"
    "sync"
)

func Test_registerModuleLockFree(t *testing.T) {
    n, parallel := 1000, 8
    head := moduledata{}
    tail := &head
    wg := sync.WaitGroup{}
    wg.Add(parallel)
    filler := func(n int) {
        defer wg.Done()
        for i := 0; i < n; i++ {
            m := &moduledata{}
            registerModuleLockFree(&tail, m)
        }
    }
    for i := 0; i < parallel; i++ {
        go filler(n)
    }
    wg.Wait()
    i := 0
    for p := head.next; p != nil; p = p.next {
        i += 1
    }
    if i != parallel * n {
        t.Errorf("got %v, expected %v", i, parallel * n)
    }
}
