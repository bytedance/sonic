// +build go1.15,!go1.19,amd64,linux go1.15,!go1.19,amd64,darwin

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

package compat

import (
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// Marshal returns the JSON encoding string of v, with faster config.
func Marshal(val interface{}) ([]byte, error) {
    return encoder.Encode(val, 0)
}

// MarshalString is like Marshal, except its output is string.
func MarshalString(val interface{}) (string, error) {
    buf, err := encoder.Encode(val, 0)
    return rt.Mem2Str(buf), err
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v, with faster config.
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
        for pos < len(buf) && (types.SPACE_MASK & (1 << buf[pos])) != 0 {
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

// MarshalStd returns the JSON encoding string of v, compatibly with encoding/json.
func MarshalStd(val interface{}) ([]byte, error) {
    return encoder.Encode(val, encoder.EscapeHTML|encoder.SortMapKeys|encoder.CompactMarshaler)
}

// MarshalStringStd is like MarshalStd, except its output is string.
func MarshalStringStd(val interface{}) (string, error) {
    buf, err := encoder.Encode(val, encoder.EscapeHTML|encoder.SortMapKeys|encoder.CompactMarshaler)
    return rt.Mem2Str(buf), err
}

// UnmarshalStd parses the JSON-encoded data and stores the result in the value
// pointed to by v, compatibly with encoding/json.
func UnmarshalStd(buf []byte, val interface{}) error {
    return UnmarshalString(string(buf), val)
}

// UnmarshalStringStd is like UnmarshalStd, except buf is a string.
func UnmarshalStringStd(buf string, val interface{}) error {
    dec := decoder.NewDecoder(buf)
    dec.CopyString()
    err := dec.Decode(val)
    pos := dec.Pos()

    /* check for errors */
    if err != nil {
        return err
    }

    /* skip all the trailing spaces */
    if pos != len(buf) {
        for pos < len(buf) && (types.SPACE_MASK & (1 << buf[pos])) != 0 {
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