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

package decoder

import (
    `runtime`
    `testing`
    `time`
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var referred = false

func TestStringReferring(t *testing.T) {
    str := []byte(`{"A":"0","B":"1"}`)
    sp := *(**byte)(unsafe.Pointer(&str))
    println("malloc *byte ", sp)
    runtime.SetFinalizer(sp, func(sp *byte){
        referred = false
        println("*byte ", sp, " got free 1")
    })
    runtime.GC()
    println("first GC")
    var obj struct{
        A string
        B string
    }
    dc := NewDecoder(rt.Mem2Str(str))
    dc.CopyString()
    referred = true
    if err := dc.Decode(&obj); err != nil {
        t.Fatal(err)
    }
    runtime.GC()
    println("second GC")
    time.Sleep(time.Millisecond)
    if referred {
        t.Fatal("*byte is being referred")
    }

    str2 := []byte(`{"A":"0","B":"1"}`)
    sp2 := *(**byte)(unsafe.Pointer(&str2))
    println("malloc *byte ", sp2)
    runtime.SetFinalizer(sp2, func(sp *byte){
        referred = false
        println("*byte ", sp, " got free")
    })
    runtime.GC()
    println("first GC")
    var obj2 interface{}
    dc2 := NewDecoder(rt.Mem2Str(str2))
    dc2.UseNumber()
    dc2.CopyString()
    referred = true
    if err := dc2.Decode(&obj2); err != nil {
        t.Fatal(err)
    }
    runtime.GC()
    println("second GC")
    time.Sleep(time.Millisecond)
    if referred {
        t.Fatal("*byte is being referred")
    }
    
    runtime.KeepAlive(&obj)
    runtime.KeepAlive(&obj2)
}

func TestDecoderErrorStackOverflower(t *testing.T) {
    src := `{"a":[]}`
    N := _MaxStack
    for i:=0; i<N; i++ {
        var obj map[string]string
        err := NewDecoder(src).Decode(&obj)
        if err == nil {
            t.Fatal(err)
        }
    }
}