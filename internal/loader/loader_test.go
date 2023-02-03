// +build go1.15,!go1.17

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

package loader

import (
    `errors`
    `fmt`
    `reflect`
    `runtime`
    `runtime/debug`
    `testing`
    `unsafe`

    `github.com/stretchr/testify/assert`
)

func TestLoader_LoadFunc(t *testing.T) {
    bc := []byte {
        0x48, 0x8b, 0x44, 0x24, 0x08,               // MOVQ  8(%rsp), %rax
        0x48, 0xc7, 0x00, 0xd2, 0x04, 0x00, 0x00,   // MOVQ  $1234, (%rax)
        0xc3,                                       // RET
    }
    v0 := 0
    l := Loader{
        Name: "test",
        File: "test.go",
    }
    fn := l.LoadFunc(bc, "test", 0, 8, []bool{true}, []bool{})
    assert.Equal(t, 1234, v0)
    println(runtime.FuncForPC(*(*uintptr)(fn)).Name())
}