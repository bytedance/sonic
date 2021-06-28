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
	"reflect"

	"github.com/bytedance/sonic/ast"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
)

// Marshal returns the JSON encoding of v.
func Marshal(val interface{}) ([]byte, error) {
	return encoder.Encode(val)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v.
func Unmarshal(buf []byte, val interface{}) error {
	return UnmarshalString(string(buf), val)
}

// UnmarshalString is like Unmarshal, except buf is a string.
func UnmarshalString(buf string, val interface{}) error {
	return decoder.NewDecoder(buf).Decode(val)
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
func Pretouch(vt reflect.Type) error {
	if err := encoder.Pretouch(vt); err != nil {
		return err
	} else if err = decoder.Pretouch(vt); err != nil {
		return err
	} else {
		return nil
	}
}

// Get searches the given path json,
// and returns its representing ast.Node
//
// Each path arg must be integer or string:
//     - Integer means searching current node as array,
//     - String means searching current node as object
func Get(src []byte, path ...interface{}) (ast.Node, error) {
	return GetFromString(string(src), path...)
}

// GetFromString is same with Get except src is string,
// which can reduce unnecessary memory copy
func GetFromString(src string, path ...interface{}) (ast.Node, error) {
	return ast.NewSearcher(src).GetByPath(path...)
}
