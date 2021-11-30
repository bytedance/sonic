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

//go:generate make
package sonic

import (
    `reflect`

    `github.com/bytedance/sonic/ast`
    `github.com/bytedance/sonic/decoder`
    `github.com/bytedance/sonic/encoder`
    `github.com/bytedance/sonic/option`
    `github.com/bytedance/sonic/internal/native/types`
)

const (
    _SpaceMask = (1 << ' ') | (1 << '\t') | (1 << '\r') | (1 << '\n')
)

// Marshal returns the JSON encoding of v.
func Marshal(val interface{}) ([]byte, error) {
    return encoder.Encode(val, 0)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v.
func Unmarshal(buf []byte, val interface{}) error {
    return UnmarshalString(string(buf), val)
}

// UnmarshalString is like Unmarshal, except buf is a string.
func UnmarshalString(buf string, val interface{}) error {
    dec := decoder.NewDecoder(buf)
    err := dec.Decode(val)
    pos := dec.Pos()

    /* check for errors */
    if err != nil {
        return err
    }

    /* skip all the trailing spaces */
    if pos != len(buf) {
        for pos < len(buf) && (_SpaceMask & (1 << buf[pos])) != 0 {
            pos++
        }
    }

    /* then it must be at EOF */
    if pos == len(buf) {
        return nil
    }

    /* junk after JSON value */
    return decoder.SyntaxError {
        Src  : buf,
        Pos  : dec.Pos(),
        Code : types.ERR_INVALID_CHAR,
    }
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
    if err := encoder.Pretouch(vt, opts...); err != nil {
        return err
    } else if err = decoder.Pretouch(vt, opts...); err != nil {
        return err
    } else {
        return nil
    }
}

// Get searches the given path json,
// and returns its representing ast.Node.
//
// Each path arg must be integer or string:
//     - Integer means searching current node as array
//     - String means searching current node as object
func Get(src []byte, path ...interface{}) (ast.Node, error) {
    return GetFromString(string(src), path...)
}

// GetFromString is same with Get except src is string,
// which can reduce unnecessary memory copy.
func GetFromString(src string, path ...interface{}) (ast.Node, error) {
    return ast.NewSearcher(src).GetByPath(path...)
}