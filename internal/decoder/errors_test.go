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
    `fmt`
    `strings`
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

func Fuzz_calcBounds(f *testing.F) {
    f.Add(33, 16)
    f.Fuzz(func(t *testing.T, a int, b int){
        if a < 0 || a > 1024*1024*10 {
            return
        }
        src := strings.Repeat("a", a)
        p, x, q, y := calcBounds(a, b)
        if x < 0 {
            t.Fatal("x < 0", x)
        }
        if y < 0 {
            t.Fatal("y < 0", y)
        }
        if x > a {
            t.Fatal("x > a", x)
        }
        if y > a {
            t.Fatal("y > a", y)
        }
        if p < 0 {
            t.Fatal("p < 0", p)
        }
        if q < 0 {
            t.Fatal("q < 0", 0)
        }
        if p > a {
            t.Fatal("p >= a", p)
        }
        if q > a {
            t.Fatal("q >= a", q)
        }
        if p > q {
            t.Fatal("p > q", q)
        }
        
        _ = fmt.Sprintf(
            "%s\n\t%s^%s\n",
            src[p:q],
            strings.Repeat(".", x),
            strings.Repeat(".", y),
        )
    })
}