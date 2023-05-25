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
	"reflect"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
	"github.com/stretchr/testify/assert"
)

func TestCompiler_Compile(t *testing.T) {
    p, err := newCompiler().compile(reflect.TypeOf(_BindingValue), false)
    assert.Nil(t, err)
    p.disassemble()
}

func TestReflectDirect(t *testing.T) {
    type A struct {
        A int
        B int
    }
    var a A
    var b = &a
    println("b:", unsafe.Pointer(b))
    v := rt.UnpackEface(a)
    vv := reflect.ValueOf(a)
    _ = vv
    println("v:", v.Type.KindFlags, v.Value)
    pv := rt.UnpackEface(&a)
    pvv := reflect.ValueOf(&a)
    _ = pvv
    println("pv:", pv.Type.KindFlags, pv.Value)
}