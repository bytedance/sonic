// +build !go1.15 go1.19 !amd64 !linux,!darwin

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
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var compatJsoniter = jsoniter.Config{
    ValidateJsonRawMessage: true,
}.Froze()

// Marshal returns the JSON encoding string of v with faster config.
func Marshal(val interface{}) ([]byte, error) {
    return compatJsoniter.Marshal(val)
}

// MarshalString is like MarshalStd, except its output is string.
func MarshalString(val interface{}) (string, error) {
   return compatJsoniter.MarshalToString(val)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v, with faster config.
func Unmarshal(buf []byte, val interface{}) error {
    return compatJsoniter.Unmarshal(buf, val)
}

// UnmarshalStringStd is like Unmarshal, except buf is a string.
func UnmarshalString(buf string, val interface{}) error {
    return compatJsoniter.UnmarshalFromString(buf, val)
}

// MarshalStd returns the JSON encoding string of v, compatibly with encoding/json.
func MarshalStd(val interface{}) ([]byte, error) {
    return json.Marshal(val)
}

// MarshalStringStd is like MarshalStd, except its output is string.
func MarshalStringStd(val interface{}) (string, error) {
    buf, err := json.Marshal(val)
    return mem2Str(buf), err
}

// UnmarshalStd parses the JSON-encoded data and stores the result in the value
// pointed to by v, compatibly with encoding/json.
func UnmarshalStd(buf []byte, val interface{}) error {
    return json.Unmarshal(buf, val)
}

// UnmarshalStringStd is like Unmarshal, except buf is a string.
func UnmarshalStringStd(buf string, val interface{}) error {
    return json.Unmarshal(str2Mem(buf), val)
}