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

import "testing"


func TestLoad(t *testing.T) {
    var f = func(i *int) {
        *i = 1234
    }
    bc := []byte {
        0x48, 0x8b, 0x44, 0x24, 0x08,               // 0x0000 MOVQ  8(%rsp), %rax
        0x48, 0xc7, 0x00, 0xd2, 0x04, 0x00, 0x00,   // 0x0005 MOVQ  $1234, (%rax)
        0xc3,                                       // 0x000c RET
    }

    fn := Func{
        ID: 0,
        Flag: 0,
        EntryOff: 0,
        TextSize: uint32(len(bc)),
        StartLine: 1,
        DeferReturn: 0,
        FileName: "dummy.go",
        Name: "dummy",
    }

    fn.Pcsp = &Pcdata{
		{PC: 5, Val: 8},
	}
}
