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
    `reflect`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/ast`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`
)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func decodeInterface(s string, f uint64) (int, interface{}, types.ParsingError)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func decodeObjectKeyString(s string, i int) string

func TestGeneric_DecodeInterface(t *testing.T) {
    s := `[null, true, false, 1234, -1.25e-8, "hello\nworld", [], {"asdf": [1, 2.5, "qwer", null, true, false, [], {"zxcv": "fghj"}]}]`
    i, v, err := decodeInterface(s, 0)
    if err < 0 {
        require.NoError(t, err)
    }
    require.Equal(t, len(s), i)
    fmt.Print("v: ")
    spew.Dump(v)
    fmt.Printf("type: %s\n", reflect.TypeOf(v))
}

func TestGeneric_DecodeObjectKeyString(t *testing.T) {
    r := decodeObjectKeyString(`    "hello\u2333world"`, 4)
    if v := (*rt.GoString)(unsafe.Pointer(&r)).Len; v < 0 {
        require.NoError(t, types.ParsingError(-v))
    }
    spew.Dump(r)
}

func BenchmarkGeneric_DecodeAST(b *testing.B) {
    _, _, _ = ast.Loads(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _, _ = ast.Loads(TwitterJson)
    }
}

func BenchmarkGeneric_DecodeInterface(b *testing.B) {
    _, _, _ = decodeInterface(TwitterJson, 0)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _, _ = decodeInterface(TwitterJson, 0)
    }
}

func BenchmarkGeneric_Parallel_DecodeAST(b *testing.B) {
    _, _, _ = ast.Loads(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _, _ = ast.Loads(TwitterJson)
        }
    })
}

func BenchmarkGeneric_Parallel_DecodeInterface(b *testing.B) {
    _, _, _ = decodeInterface(TwitterJson, 0)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _, _ = decodeInterface(TwitterJson, 0)
        }
    })
}
