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

package issue_test

import (
    `encoding/json`
    `reflect`
    `runtime`
    `runtime/debug`
    `sync`
    `testing`
    `time`

    `fmt`
    `unsafe`
    `github.com/bytedance/sonic/internal/rt`

    . `github.com/bytedance/sonic`
)

type jsonNumberWrapper struct {
    Number json.Number `json:"number"`
}

func wrapJsonNumber(t *testing.T) interface{} {
    var data = []byte(`{"number":123456789123456789}`)
    var src = rt.Mem2Str(data)
    ptr := (*rt.GoString)(unsafe.Pointer(&src))
    fmt.Printf("src:%#v\n", *ptr)
    var bs = (*byte)(ptr.Ptr)
    runtime.SetFinalizer(bs, func (bs *byte)  {
        fmt.Printf("bs(%v) got dropped\n", bs)
    })

    var obj jsonNumberWrapper
    if err := UnmarshalString(src, &obj); err != nil {
        t.Fatal(err)
    }

    time.Sleep(time.Millisecond)
    runtime.GC()
    debug.FreeOSMemory()
    time.Sleep(time.Millisecond)
    println("out")
    return obj
}

func TestIssueJsonNumber(t *testing.T) {
    var obj interface{}
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        defer wg.Done()
        obj = wrapJsonNumber(t)
        runtime.GC()
        debug.FreeOSMemory()
        x := obj.(jsonNumberWrapper)
        fmt.Printf("src:%#v\n", *(*rt.GoString)(unsafe.Pointer(&x.Number)))
    }()

    wg.Wait()
    exp := jsonNumberWrapper{json.Number("123456789123456789")}
    if !reflect.DeepEqual(obj, exp) {
        t.Fatal(obj, exp)
    }
    
}