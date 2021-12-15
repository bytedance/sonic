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

package ast

import (
    `reflect`
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/encoder`
    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _MaxBuffer = 1024    // 1KB buffer size
)

const (
    bytesNull   = "null"
    bytesTrue   = "true"
    bytesFalse  = "false"
    bytesObject = "{}"
    bytesArray  = "[]"
)

var bytesPool   = sync.Pool{}

func (self *Node) MarshalJSON() ([]byte, error) {
    buf := newBuffer()
    err := self.encode(buf)
    if err != nil {
        freeBuffer(buf)
        return nil, err
    }

    ret := make([]byte, len(*buf))
    copy(ret, *buf)
    freeBuffer(buf)
    return ret, err
}

func newBuffer() *[]byte {
    if ret := bytesPool.Get(); ret != nil {
        return ret.(*[]byte)
    } else {
        buf := make([]byte, 0, _MaxBuffer)
        return &buf
    }
}

func freeBuffer(buf *[]byte) {
    *buf = (*buf)[:0]
    bytesPool.Put(buf)
}

func (self *Node) encode(buf *[]byte) error {
    if self.IsRaw() {
        return self.encodeRaw(buf)
    }
    switch self.Type() {
        case V_NONE  : return ErrNotExist
        case V_ERROR : return self.Check()
        case V_NULL  : return self.encodeNull(buf)
        case V_TRUE  : return self.encodeTrue(buf)
        case V_FALSE : return self.encodeFalse(buf)
        case V_ARRAY : return self.encodeArray(buf)
        case V_OBJECT: return self.encodeObject(buf)
        case V_STRING: return self.encodeString(buf)
        case V_NUMBER: return self.encodeNumber(buf)
        case V_ANY   : return self.encodeInterface(buf)
        default      : return ErrUnsupportType 
    }
}

func (self *Node) encodeRaw(buf *[]byte) error {
    raw, err := self.Raw()
    if err != nil {
        return err
    }
    *buf = append(*buf, raw...)
    return nil
}

func (self *Node) encodeNull(buf *[]byte) error {
    *buf = append(*buf, bytesNull...)
    return nil
}

func (self *Node) encodeTrue(buf *[]byte) error {
    *buf = append(*buf, bytesTrue...)
    return nil
}

func (self *Node) encodeFalse(buf *[]byte) error {
    *buf = append(*buf, bytesFalse...)
    return nil
}

func (self *Node) encodeNumber(buf *[]byte) error {
    str := addr2str(self.p, self.v)
    *buf = append(*buf, str...)
    return nil
}

var typeByte = rt.UnpackType(reflect.TypeOf(byte(0)))

func quote(buf *[]byte, sp unsafe.Pointer, nb int) {
    b := (*rt.GoSlice)(unsafe.Pointer(buf))
    // input buffer
    for nb > 0 {
        // output buffer
        dp := unsafe.Pointer(uintptr(b.Ptr) + uintptr(b.Len))
        dn := b.Cap - b.Len
        // call native.Quote, dn is byte count it outputs
        ret := native.Quote(sp, nb, dp, &dn, 0)
        // update *buf length
        b.Len += dn

        // no need more output
        if ret >= 0 {
            break
        }

        // double buf size
        *b = growslice(typeByte, *b, b.Cap * 2)
        // ret is the complement of consumed input
        ret = ^ret
        // update input buffer
        nb -= ret
        sp = unsafe.Pointer(uintptr(sp) + uintptr(ret))
    }
}

func (self *Node) encodeString(buf *[]byte) error {
    *buf = append(*buf, '"')
    nb := int(self.v)
    if nb == 0 {
        *buf = append(*buf, '"')
        return nil
    }

    quote(buf, self.p, nb)

    *buf = append(*buf, '"')
    return nil
}

func (self *Node) encodeArray(buf *[]byte) error {
    if self.isLazy() {
        if err := self.skipAllIndex(); err != nil {
            return err
        }
    }

    nb := self.len()
    if nb == 0 {
        *buf = append(*buf, bytesArray...)
        return nil
    }
    
    *buf = append(*buf, '[')

    var p = (*Node)(self.p)
    err := p.encode(buf)
    if err != nil {
        return err
    }
    for i := 1; i < nb; i++ {
        *buf = append(*buf, ',')
        p = p.unsafe_next()
        err := p.encode(buf)
        if err != nil {
            return err
        }
    }

    *buf = append(*buf, ']')
    return nil
}

func (self *Pair) encode(buf *[]byte) error {
    *buf = append(*buf, '"')
    sptr := (*rt.GoString)(unsafe.Pointer(&self.Key))
    if sptr.Len == 0 {
        *buf = append(*buf, '"', ':')
        return self.Value.encode(buf)
    }

    quote(buf, sptr.Ptr, sptr.Len)

    *buf = append(*buf, '"', ':')
    return self.Value.encode(buf)
}

func (self *Node) encodeObject(buf *[]byte) error {
    if self.isLazy() {
        if err := self.skipAllKey(); err != nil {
            return err
        }
    }
    
    nb := self.len()
    if nb == 0 {
        *buf = append(*buf, bytesObject...)
        return nil
    }
    
    *buf = append(*buf, '{')

    var p = (*Pair)(self.p)
    err := p.encode(buf)
    if err != nil {
        return err
    }
    for i := 1; i < nb; i++ {
        *buf = append(*buf, ',')
        p = p.unsafe_next()
        err := p.encode(buf)
        if err != nil {
            return err
        }
    }

    *buf = append(*buf, '}')
    return nil
}

func (self *Node) encodeInterface(buf *[]byte) error {
    return encoder.EncodeInto(buf, self.packAny(), 0)
}