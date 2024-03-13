/**
 * Copyright 2024 ByteDance Inc.
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

// This package is only used for exporting internal API to thrid-party libs. DO NOT USE IT.
package export

import (
	_ "unsafe"
)

// DecoderString decodes a JSON string from pos and return golang string.
//   - needEsc indicates if to unescaped escaping chars
//   - hasEsc tells if the returned string has escaping chars
//   - validStr enables validating UTF8 charset
//
//go:linkname DecodeString github.com/bytedance/sonic/ast._DecodeString
func DecodeString(src string, pos int, needEsc bool, validStr bool) (v string, ret int, hasEsc bool)

// GetByPath searches a path and returns relaction and types of target
//
//go:linkname GetByPath github.com/bytedance/sonic/ast._GetByPath  
func GetByPath(src string, path ...interface{}) (start int, end int, typ int, err error)

// ValidSyntax check if a json has a valid JSON syntax,
// while not validate UTF-8 charset
//
//go:linkname ValidSyntax github.com/bytedance/sonic/ast._ValidSyntax  
func ValidSyntax(json string) bool 

// SkipFast skip a json value in fast-skip algs, 
// while not strictly validate JSON syntax and UTF-8 charset.
//
//go:linkname SkipFast github.com/bytedance/sonic/ast._SkipFast  
func SkipFast(src string, i int) (int, int, error)

// Unquote unescapes a escaped string (not including `"` at begining and end)
//   - replace enables replacing invalid utf8 escaped char with `\uffd`
//
//go:linkname Unquote github.com/bytedance/sonic/unquote._String  
func Unquote(s string, replace bool) (ret string, err error)
