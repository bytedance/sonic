/**
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
	`os`
	`runtime`
	`runtime/debug`
	`testing`
	`time`

	`github.com/stretchr/testify/require`
)

var (
	debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

func TestMain(m *testing.M) {
	go func ()  {
		if !debugAsyncGC {
			return
		}
		println("Begin GC looping...")
		for {
		runtime.GC()
		debug.FreeOSMemory() 
		}
		println("stop GC looping!")
	}()
	time.Sleep(time.Millisecond*100)
	m.Run()
}

func TestWrapC(t *testing.T) {
	var stub func(a int64, val *int64) (ret int64) 
	ct := []byte{
		0x55,                    // pushq   %rbp
		0x48, 0x89, 0xe5,        // movq    %rsp, %rbp
		0x48, 0x89, 0x7d, 0xf8,  // movq    %rdi, -8(%rbp)
		0x48, 0x89, 0x75, 0xf0,  // movq    %rsi, -16(%rbp)
		0x48, 0x8b, 0x75, 0xf8,  // movq    -8(%rbp), %rsi
		0x48, 0x8b, 0x7d, 0xf0,  // movq    -16(%rbp), %rdi
		0x48, 0x03, 0x37,        // addq    (%rdi), %rsi
		0x48, 0x89, 0xf0,        // movq    %rsi, %rax
		0x5d,                    // popq    %rbp
		0xc3,                    // ret
	}
	
	WrapGoC(ct, []CFunc{{
		Name:     "add",
		EntryOff: 0,
		TextSize: uint32(len(ct)),
		MaxStack: uintptr(16),
		Pcsp:     [][2]uint32{
			{1, 8},
		},
	}}, []GoC{{
		CName:     "add",
		GoFunc:   &stub,
	} }, "dummy/native", "dummy/native.c")
	
	// defer func(){
	//     if err := recover(); err!= nil {
	//         println("panic:", err)
	//     } else {
	//         t.Fatal("not panic")
	//     }
	// }()

	f := stub
	b := int64(2)    
	println("b : ", &b)
	var c *int64 = &b
	runtime.SetFinalizer(c, func(x *int64){
		println("c got GC: ", x)
	})
	runtime.GC()
	println("before")
	act := f(1, c)
	println("after")
	runtime.GC()
	require.Equal(t, int64(3), act)
}
