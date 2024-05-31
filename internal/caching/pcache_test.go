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
    `sync`
    `testing`

    `github.com/bytedance/sonic/internal/rt`
)

func TestPcacheRace(t *testing.T) {
    t.Parallel()

    pc := CreateProgramCache()
    wg := sync.WaitGroup{}
    wg.Add(2)
    start := make(chan struct{}, 2)

    go func(){
        defer wg.Done()
        var k = map[string]interface{}{}
        <- start
        for i:=0; i<100; i++ {
            _, _ = pc.Compute(rt.UnpackEface(k).Type, func(*rt.GoType, ... interface{}) (interface{}, error) {
                return map[string]interface{}{}, nil
            })
        }
    }()

    go func(){
        defer wg.Done()
        var k = map[string]interface{}{}
        <- start
        for i:=0; i<100; i++ {
            pc.Get(rt.UnpackEface(k).Type)
        }
    }()

    start <- struct{}{}
    start <- struct{}{}
    wg.Wait()
}