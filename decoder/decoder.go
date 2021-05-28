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
    `encoding/json`
    `reflect`

    `github.com/bytedance/sonic/internal/rt`
)

const (
    _F_use_int64 = iota
    _F_use_number
    _F_disable_urc
    _F_disable_unknown
)

type Decoder struct {
    i int
    f uint64
    s string
}

func NewDecoder(s string) *Decoder {
    return &Decoder{s: s}
}

func (self *Decoder) Pos() int {
    return self.i
}

func (self *Decoder) Decode(val interface{}) error {
    vv := rt.UnpackEface(val)
    vp := vv.Value

    /* check for nil type */
    if vv.Type == nil {
        return &json.InvalidUnmarshalError{}
    }

    /* must be a non-nil pointer */
    if vp == nil || vv.Type.Kind() != reflect.Ptr {
        return &json.InvalidUnmarshalError{Type: vv.Type.Pack()}
    }

    /* create a new stack, and call the decoder */
    sb, etp := newStack(), rt.PtrElem(vv.Type)
    nb, err := decodeTypedPointer(self.s, self.i, etp, vp, sb, self.f)

    /* return the stack back */
    self.i = nb
    freeStack(sb)
    return err
}

func (self *Decoder) UseInt64() {
    self.f  |= 1 << _F_use_int64
    self.f &^= 1 << _F_use_number
}

func (self *Decoder) UseNumber() {
    self.f &^= 1 << _F_use_int64
    self.f  |= 1 << _F_use_number
}

func (self *Decoder) UseUnicodeErrors() {
    self.f |= 1 << _F_disable_urc
}

func (self *Decoder) DisallowUnknownFields() {
    self.f |= 1 << _F_disable_unknown
}

func Pretouch(vt reflect.Type) (err error) {
    _, err = findOrCompile(rt.UnpackType(vt))
    return
}
