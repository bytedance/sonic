//go:build !race
// +build !race

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

package encoder

import (
    `runtime`
    `sync`
    `testing`

	`github.com/bytedance/sonic/internal/encoder/vars`
)

func TestGC(t *testing.T) {
    if !vars.DebugAsyncGC {
        return
    }
    out, err := Encode(_GenericValue, 0)
    if err != nil {
        t.Fatal(err)
    }
    n := len(out)
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup, size int)  {
            defer wg.Done()
            out, err := Encode(_GenericValue, 0)
            if err != nil {
                t.Fatal(err)
            }
            if len(out) != size {
                t.Fatal(len(out), size)
            }
            runtime.GC()
        }(wg, n)
    }
    wg.Wait()
}