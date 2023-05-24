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
    `testing`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/stretchr/testify/assert`
)

func make_err(src string, pos int) SyntaxError {
    return SyntaxError {
        Src  : src,
        Pos  : pos,
        Code : types.ERR_INVALID_CHAR,
    }
}

func TestErrors_Normal(t *testing.T) {
    println(make_err("this is a very long message with 'hello, world' embedded in the string", 33).Description())
}

func TestErrors_LeftEdge(t *testing.T) {
    println(make_err("this is a very long message with 'hello, world' embedded in the string", 6).Description())
}

func TestErrors_RightEdge(t *testing.T) {
    println(make_err("this is a very long message with 'hello, world' embedded in the string", 65).Description())
}

func TestErrors_AfterRightEdge(t *testing.T) {
    println(make_err("this is a very long message with 'hello, world' embedded in the string", 70).Description())
}

func TestErrors_ShortDescription(t *testing.T) {
    e := make_err("hello, world", 5)
    println(e.Description())
    assert.Equal(t, "Syntax error at index 5: invalid char\n\n\thello, world\n\t.....^......\n", e.Description())
    assert.Equal(t, `"Syntax error at index 5: invalid char\n\n\thello, world\n\t.....^......\n"`, e.Error())
}

func TestErrors_EmptyDescription(t *testing.T) {
    println(make_err("", 0).Description())
}