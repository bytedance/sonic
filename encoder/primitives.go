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
    `encoding`
    `encoding/json`
    `unsafe`

    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/rt`
)

/** Encoder Primitives **/

func encodeNil(rb *[]byte) error {
    *rb = append(*rb, 'n', 'u', 'l', 'l')
    return nil
}

func encodeString(buf *[]byte, val string) error {
    var sidx int
    var pbuf *rt.GoSlice
    var pstr *rt.GoString

    /* opening quote */
    *buf = append(*buf, '"')
    pbuf = (*rt.GoSlice)(unsafe.Pointer(buf))
    pstr = (*rt.GoString)(unsafe.Pointer(&val))

    /* encode with native library */
    for sidx < pstr.Len {
        sn := pstr.Len - sidx
        dn := pbuf.Cap - pbuf.Len
        sp := padd(pstr.Ptr, sidx)
        dp := padd(pbuf.Ptr, pbuf.Len)
        nb := native.Quote(sp, sn, dp, &dn, 0)

        /* check for errors */
        if pbuf.Len += dn; nb >= 0 {
            break
        }

        /* not enough space, grow the slice and try again */
        sidx += ^nb
        *pbuf = growslice(rt.UnpackType(byteType), *pbuf, pbuf.Cap * 2)
    }

    /* closing quote */
    *buf = append(*buf, '"')
    return nil
}

func encodeTypedPointer(buf *[]byte, vt *rt.GoType, vp *unsafe.Pointer, sb *_Stack, fv uint64) error {
    if vt == nil {
        return encodeNil(buf)
    } else if fn, err := findOrCompile(vt); err != nil {
        return err
    } else if (vt.KindFlags & rt.F_direct) == 0 {
        rt.MoreStack(_FP_size + native.MaxFrameSize)
        return fn(buf, *vp, sb, fv)
    } else {
        rt.MoreStack(_FP_size + native.MaxFrameSize)
        return fn(buf, unsafe.Pointer(vp), sb, fv)
    }
}

func encodeJsonMarshaler(buf *[]byte, val json.Marshaler) error {
    if ret, err := val.MarshalJSON(); err != nil {
        return err
    } else {
        return compact(buf, ret)
    }
}

func encodeTextMarshaler(buf *[]byte, val encoding.TextMarshaler) error {
    if ret, err := val.MarshalText(); err != nil {
        return err
    } else {
        return encodeString(buf, rt.Mem2Str(ret))
    }
}