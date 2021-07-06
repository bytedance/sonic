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
    `reflect`
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

func encodeTypedPointer(buf *[]byte, vt *rt.GoType, vp *unsafe.Pointer, sb *_Stack) error {
    if vt == nil {
        return encodeNil(buf)
    } else if fn, err := findOrCompile(vt); err != nil {
        return err
    } else if (vt.KindFlags & rt.F_direct) == 0 {
        return fn(buf, *vp, sb)
    } else {
        return fn(buf, unsafe.Pointer(vp), sb)
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

func isZeroSafe(p unsafe.Pointer, vt *rt.GoType) bool {
    if native.Lzero(p, int(vt.Size)) == 0 {
        return true
    } else {
        return isZeroTyped(p, vt)
    }
}

func isZeroTyped(p unsafe.Pointer, vt *rt.GoType) bool {
    switch vt.Kind() {
        case reflect.Map       : return (*(**rt.GoMap)(p)).Count == 0
        case reflect.Slice     : return (*rt.GoSlice)(p).Len == 0
        case reflect.String    : return (*rt.GoString)(p).Len == 0
        case reflect.Struct    : return isZeroStruct(p, vt)
        case reflect.Interface : return (*rt.GoEface)(p).Value == nil
        default                : return false
    }
}

func isZeroStruct(p unsafe.Pointer, vt *rt.GoType) bool {
    var dp uintptr
    var fp unsafe.Pointer

    /* check for each field */
    for _, fv := range (*rt.GoStructType)(unsafe.Pointer(vt)).Fields {
        dp = fv.OffEmbed >> 1
        fp = unsafe.Pointer(uintptr(p) + dp)

        /* check for the field */
        if !isZeroSafe(fp, fv.Type) {
            return false
        }
    }

    /* all tests are passed */
    return true
}

func isTrivialZeroable(vt reflect.Type) bool {
    switch vt.Kind() {
        case reflect.Bool      : return true
        case reflect.Int       : return true
        case reflect.Int8      : return true
        case reflect.Int16     : return true
        case reflect.Int32     : return true
        case reflect.Int64     : return true
        case reflect.Uint      : return true
        case reflect.Uint8     : return true
        case reflect.Uint16    : return true
        case reflect.Uint32    : return true
        case reflect.Uint64    : return true
        case reflect.Uintptr   : return true
        case reflect.Float32   : return true
        case reflect.Float64   : return true
        case reflect.String    : return false
        case reflect.Array     : return true
        case reflect.Interface : return false
        case reflect.Map       : return false
        case reflect.Ptr       : return true
        case reflect.Slice     : return false
        case reflect.Struct    : return isStructTrivialZeroable(vt)
        default                : return false
    }
}

func isStructTrivialZeroable(vt reflect.Type) bool {
    for i := 0; i < vt.NumField(); i++ { if !isTrivialZeroable(vt.Field(i).Type) { return false } }
    return true
}
