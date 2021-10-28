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
    `runtime`

    `github.com/bytedance/sonic/internal/rt`
)

// Options is a set of encoding options.
type Options uint64

const (
    bitSortMapKeys = iota
)

const (
    // SortMapKeys indicate that the keys of a map needs to be sorted before
    // serializing into JSON.
    // WARNING: This hurts performance A LOT, USE WITH CARE.
    SortMapKeys Options = 1 << bitSortMapKeys
)

// Encoder represents a specific set of encoder configurations.
type Encoder struct {
    Opts Options
}

// Encode returns the JSON encoding of v.
func (self *Encoder) Encode(v interface{}) ([]byte, error) {
    return Encode(v, self.Opts)
}

// SortKeys enables the SortMapKeys option.
func (self *Encoder) SortKeys() *Encoder {
    self.Opts |= SortMapKeys
    return self
}

// Quote returns the JSON-quoted version of s.
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

// Encode returns the JSON encoding of val, encoded with opts.
func Encode(val interface{}, opts Options) ([]byte, error) {
    buf := newBytes()
    err := EncodeInto(&buf, val, opts)

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

// EncodeInto is like Encode but uses a user-supplied buffer instead of allocating
// a new one.
func EncodeInto(buf *[]byte, val interface{}, opts Options) error {
    stk := newStack()
    efv := rt.UnpackEface(val)
    err := encodeTypedPointer(buf, efv.Type, &efv.Value, stk, uint64(opts))

    /* return the stack into pool */
    freeStack(stk)

    /* avoid GC ahead */
    runtime.KeepAlive(efv)
    return err
}

// EncodeIndented is like Encode but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
func EncodeIndented(val interface{}, prefix string, indent string, opts Options) ([]byte, error) {
    var err error
    var out []byte
    var buf *bytes.Buffer

    /* encode into the buffer */
    out = newBytes()
    err = EncodeInto(&out, val, opts)

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

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
func Pretouch(vt reflect.Type) (err error) {
    _, err = findOrCompile(rt.UnpackType(vt))
    return
}
