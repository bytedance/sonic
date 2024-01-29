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
	"testing"
)

func TestErrors_ParseError(t *testing.T) {
	var got interface{}
	err := NewDecoder(`{123}`).Decode(&got)
	println(err.(SyntaxError).Description())
}

type A struct {
	A string
}

func TestErrors_MismatchType(t *testing.T) {
	var got A
	err := NewDecoder(`{"a": 123}`).Decode(&got)
	println(err.(MismatchTypeError).Description())
}
