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

var _QuoteTab = [256]string {
    '\x00' : `\u0000`,
    '\x01' : `\u0001`,
    '\x02' : `\u0002`,
    '\x03' : `\u0003`,
    '\x04' : `\u0004`,
    '\x05' : `\u0005`,
    '\x06' : `\u0006`,
    '\x07' : `\u0007`,
    '\b'   : `\b`,
    '\t'   : `\t`,
    '\n'   : `\n`,
    '\x0b' : `\u000b`,
    '\f'   : `\f`,
    '\r'   : `\r`,
    '\x0e' : `\u000e`,
    '\x0f' : `\u000f`,
    '\x10' : `\u0010`,
    '\x11' : `\u0011`,
    '\x12' : `\u0012`,
    '\x13' : `\u0013`,
    '\x14' : `\u0014`,
    '\x15' : `\u0015`,
    '\x16' : `\u0016`,
    '\x17' : `\u0017`,
    '\x18' : `\u0018`,
    '\x19' : `\u0019`,
    '\x1a' : `\u001a`,
    '\x1b' : `\u001b`,
    '\x1c' : `\u001c`,
    '\x1d' : `\u001d`,
    '\x1e' : `\u001e`,
    '\x1f' : `\u001f`,
    '"'    : `\"`,
    '\\'   : `\\`,
}

var _DoubleQuoteTab = [256]string {
    '\x00' : `\\u0000`,
    '\x01' : `\\u0001`,
    '\x02' : `\\u0002`,
    '\x03' : `\\u0003`,
    '\x04' : `\\u0004`,
    '\x05' : `\\u0005`,
    '\x06' : `\\u0006`,
    '\x07' : `\\u0007`,
    '\b'   : `\\b`,
    '\t'   : `\\t`,
    '\n'   : `\\n`,
    '\x0b' : `\\u000b`,
    '\f'   : `\\f`,
    '\r'   : `\\r`,
    '\x0e' : `\\u000e`,
    '\x0f' : `\\u000f`,
    '\x10' : `\\u0010`,
    '\x11' : `\\u0011`,
    '\x12' : `\\u0012`,
    '\x13' : `\\u0013`,
    '\x14' : `\\u0014`,
    '\x15' : `\\u0015`,
    '\x16' : `\\u0016`,
    '\x17' : `\\u0017`,
    '\x18' : `\\u0018`,
    '\x19' : `\\u0019`,
    '\x1a' : `\\u001a`,
    '\x1b' : `\\u001b`,
    '\x1c' : `\\u001c`,
    '\x1d' : `\\u001d`,
    '\x1e' : `\\u001e`,
    '\x1f' : `\\u001f`,
    '"'    : `\\\"`,
    '\\'   : `\\\\`,
}

func encodeNil(rb *[]byte) error {
    *rb = append(*rb, 'n', 'u', 'l', 'l')
    return nil
}

func encodeStr(buf *[]byte, val string) error {
    *buf = append(*buf, '"')
    encodeQuote(buf, native.Lquote(&val, 0), val)
    *buf = append(*buf, '"')
    return nil
}

func encodeQuote(buf *[]byte, i int, val string) {
    p := 0
    n := len(val)

    /* quote all the characters, if any */
    for i < n {
        *buf = append(*buf, rt.Str2Mem(val[p:i])...)
        *buf = append(*buf, rt.Str2Mem(_QuoteTab[val[i]])...)
        p, i = i + 1, native.Lquote(&val, i + 1)
    }

    /* add the remaining characters */
    if p < n {
        *buf = append(*buf, rt.Str2Mem(val[p:])...)
    }
}

func encodeDoubleQuote(buf *[]byte, i int, val string) {
    p := 0
    n := len(val)

    /* quote all the characters, if any */
    for i < n {
        *buf = append(*buf, rt.Str2Mem(val[p:i])...)
        *buf = append(*buf, rt.Str2Mem(_DoubleQuoteTab[val[i]])...)
        p, i = i + 1, native.Lquote(&val, i + 1)
    }

    /* add the remaining characters */
    if p < n {
        *buf = append(*buf, rt.Str2Mem(val[p:])...)
    }
}

func encodeTypedPointer(buf *[]byte, vt *rt.GoType, vp *unsafe.Pointer, sb *_Stack) error {
    if vt == nil {
        return encodeNil(buf)
    } else if fn, err := findOrCompile(vt); err != nil {
        return err
    } else if vt.Indir() {
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
        return encodeStr(buf, rt.Mem2Str(ret))
    }
}

func isZeroSafe(p unsafe.Pointer, vt *rt.GoType) bool {
    if native.Lzero(p, vt.Size()) == 0 {
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
