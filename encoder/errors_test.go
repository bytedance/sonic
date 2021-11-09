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
	`encoding/json`
	`fmt`
	`runtime`
	`runtime/debug`
	`testing`
	`unsafe`

	`github.com/bytedance/sonic/internal/rt`
)

func Test_error_number(t *testing.T) {
	src := "0xffff"
	if isValidNumber(src) {
		t.Fatal()
	}
	gs := (*rt.GoString)(unsafe.Pointer(&src))
	println(gs.Len, gs.Ptr)
	jn := json.Number(src)
	runtime.SetFinalizer(&jn, func(jn *json.Number){
		fmt.Printf("json.Number(%v) got dropped!\n", unsafe.Pointer(jn))
	})

	println("before encode")
	debugGC = true
	_, err := Encode(jn, 0)
	debugGC = false
	println("after encode")
	runtime.GC()
	debug.FreeOSMemory()

	println(err, err.Error())
}
