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

package types

import (
    `fmt`
)

type ValueType int
type ParsingError uint
type SearchingError uint

const (
    V_EOF        ValueType = 1
    V_NULL       ValueType = 2
    V_TRUE       ValueType = 3
    V_FALSE      ValueType = 4
    V_ARRAY      ValueType = 5
    V_OBJECT     ValueType = 6
    V_STRING     ValueType = 7
    V_DOUBLE     ValueType = 8
    V_INTEGER    ValueType = 9
    V_KEY_SEP    ValueType = 10
    V_ELEM_SEP   ValueType = 11
    V_ARRAY_END  ValueType = 12
    V_OBJECT_END ValueType = 13
    V_MAX
)

const (
    B_DOUBLE_UNQUOTE  = 0
    B_UNICODE_REPLACE = 1
)

const (
    F_DOUBLE_UNQUOTE  = 1 << B_DOUBLE_UNQUOTE
    F_UNICODE_REPLACE = 1 << B_UNICODE_REPLACE
)

const (
    MAX_RECURSE = 65536
)

const (
    ERR_EOF                ParsingError = 1
    ERR_INVALID_CHAR       ParsingError = 2
    ERR_INVALID_ESCAPE     ParsingError = 3
    ERR_INVALID_UNICODE    ParsingError = 4
    ERR_INTEGER_OVERFLOW   ParsingError = 5
    ERR_INVALID_NUMBER_FMT ParsingError = 6
    ERR_RECURSE_EXCEED_MAX ParsingError = 7
)

var _ParsingErrors = []string{
    0                      : "ok",
    ERR_EOF                : "eof",
    ERR_INVALID_CHAR       : "invalid char",
    ERR_INVALID_ESCAPE     : "invalid escape char",
    ERR_INVALID_UNICODE    : "invalid unicode escape",
    ERR_INTEGER_OVERFLOW   : "integer overflow",
    ERR_INVALID_NUMBER_FMT : "invalid number format",
    ERR_RECURSE_EXCEED_MAX : "recursion exceeded max depth",
}

func (self ParsingError) Error() string {
    return "json: error when parsing input: " + self.Message()
}

func (self ParsingError) Message() string {
    if int(self) < len(_ParsingErrors) {
        return _ParsingErrors[self]
    } else {
        return fmt.Sprintf("unknown error %d", self)
    }
}

type JsonState struct {
    Vt ValueType
    Dv float64
    Iv int64
    Ep int
}

type StateMachine struct {
    Sp int
    Vt [MAX_RECURSE]int
}
