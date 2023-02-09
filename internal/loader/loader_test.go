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
    `testing`
    `unsafe`

    `github.com/stretchr/testify/assert`
)

func TestLoader_Load(t *testing.T) {
    bc := []byte {
        0x48, 0x8b, 0x44, 0x24, 0x08,               // MOVQ  8(%rsp), %rax
        0x48, 0xc7, 0x00, 0xd2, 0x04, 0x00, 0x00,   // MOVQ  $1234, (%rax)
        0xc3,                                       // RET
    }
    v0 := 0
    fn := Loader(bc).Load("test", 0, 8, nil, nil)
    (*(*func(*int))(unsafe.Pointer(&fn)))(&v0)
    assert.Equal(t, 1234, v0)
    println(runtime.FuncForPC(*(*uintptr)(fn)).Name())
}

func faker1(in string) error {
    runtime.KeepAlive(in)
    return errors.New("1")
}  

func faker2(in string) error {
    runtime.KeepAlive(in)
    return errors.New("2")
}

func faker4(in string) error {
    return nil
}

func TestStackMap(t *testing.T) {
    args1, locals1 := stackMap(faker1)
    fi1 := findfunc(reflect.ValueOf(faker1).Pointer())
    fmt.Printf("func1: %#v, args: %x, locals: %x\n", fi1, args1, locals1)

    args2, locals2 := stackMap(faker2)
    fi2 := findfunc(reflect.ValueOf(faker2).Pointer())
    fmt.Printf("func2: %#v, args: %x, locals: %x\n", fi2, args2, locals2)

    args4, locals4 := stackMap(faker4)
    fi4 := findfunc(reflect.ValueOf(faker4).Pointer())
    fmt.Printf("func4: %#v, args: %x, locals: %x\n", fi4, args4, locals4)

    if reflect.DeepEqual(fi1, fi2) || reflect.DeepEqual(fi1, fi2) {
        t.Fatal()
    }
    if args1 != args2 || locals1 != locals2 {
        t.Fatal()
    }
    if args1 == args4 || locals1 == locals4 || args2 == args4 || locals2 == locals4 {
        t.Fatal()
    }
}

func funcWrap(f func(i *int)) int {
    var ret int
    var x int = 0
    runtime.SetFinalizer(&x, func(xp *int){
        fmt.Printf("x got dropped: %x\n", unsafe.Pointer(xp))
    })
    f(&x)
    ret = x
    return ret
}
