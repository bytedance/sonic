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
    `bytes`
    `encoding/json`
    `reflect`

    `github.com/bytedance/sonic/internal/rt`
)

func Quote(s string) string {
    var n int
    var p []byte

    /* check for empty string */
    if s == "" {
        return `""`
    }

    /* allocate space for result */
    n = len(s) + 2
    p = make([]byte, 0, n)

    /* call the encoder */
    _ = encodeString(&p, s)
    return rt.Mem2Str(p)
}

func Encode(val interface{}) ([]byte, error) {
    buf := newBytes()
    err := EncodeInto(&buf, val)

    /* check for errors */
    if err != nil {
        freeBytes(buf)
        return nil, err
    }

    /* make a copy of the result */
    ret := make([]byte, len(buf))
    copy(ret, buf)

    /* return the buffer into pool */
    freeBytes(buf)
    return ret, nil
}

func EncodeInto(buf *[]byte, val interface{}) error {
    stk := newStack()
    efv := rt.UnpackEface(val)
    err := encodeTypedPointer(buf, efv.Type, &efv.Value, stk)

    /* return the stack into pool */
    freeStack(stk)
    return err
}

func EncodeIndented(val interface{}, prefix string, indent string) ([]byte, error) {
    var err error
    var out []byte
    var buf *bytes.Buffer

    /* encode into the buffer */
    out = newBytes()
    err = EncodeInto(&out, val)

    /* check for errors */
    if err != nil {
        freeBytes(out)
        return nil, err
    }

    /* indent the JSON */
    buf = newBuffer()
    err = json.Indent(buf, out, prefix, indent)

    /* check for errors */
    if err != nil {
        freeBytes(out)
        freeBuffer(buf)
        return nil, err
    }

    /* copy to the result buffer */
    ret := make([]byte, buf.Len())
    copy(ret, buf.Bytes())

    /* return the buffers into pool */
    freeBytes(out)
    freeBuffer(buf)
    return ret, nil
}

func Pretouch(vt reflect.Type) (err error) {
    _, err = findOrCompile(rt.UnpackType(vt))
    return
}
