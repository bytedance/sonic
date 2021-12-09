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
    `errors`
    `fmt`
    `reflect`
    `strconv`
    `strings`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

type SyntaxError struct {
    Pos  int
    Src  string
    Code types.ParsingError
}

func (self SyntaxError) Error() string {
    return fmt.Sprintf("%q", self.Description())
}

func (self SyntaxError) Description() string {
    i := 16
    p := self.Pos - i
    q := self.Pos + i

    /* check for empty source */
    if self.Src == "" {
        return fmt.Sprintf("no sources available: %#v", self)
    }

    /* prevent slicing before the beginning */
    if p < 0 {
        p, q, i = 0, q - p, i + p
    }

    /* prevent slicing beyond the end */
    if n := len(self.Src); q > n {
        n = q - n
        q = len(self.Src)

        /* move the left bound if possible */
        if p > n {
            i += n
            p -= n
        }
    }

    /* left and right length */
    x := clamp_zero(i)
    y := clamp_zero(q - p - i - 1)

    /* compose the error description */
    return fmt.Sprintf(
        "Syntax error at index %d: %s\n\n\t%s\n\t%s^%s\n",
        self.Pos,
        self.Code.Message(),
        self.Src[p:q],
        strings.Repeat(".", x),
        strings.Repeat(".", y),
    )
}

func clamp_zero(v int) int {
    if v < 0 {
        return 0
    } else {
        return v
    }
}

/** JIT Error Helpers **/

var stackOverflow = &json.UnsupportedValueError {
    Str   : "Value nesting too deep",
    Value : reflect.ValueOf("..."),
}

//go:nosplit
func error_wrap(src string, pos int, code types.ParsingError) error {
    return SyntaxError {
        Pos  : pos,
        Src  : src,
        Code : code,
    }
}

//go:nosplit
func error_type(vt *rt.GoType) error {
    return &json.UnmarshalTypeError{Type: vt.Pack()}
}

//go:nosplit
func error_field(name string) error {
    return errors.New("json: unknown field " + strconv.Quote(name))
}

//go:nosplit
func error_value(value string, vtype reflect.Type) error {
    return &json.UnmarshalTypeError {
        Type  : vtype,
        Value : value,
    }
}
