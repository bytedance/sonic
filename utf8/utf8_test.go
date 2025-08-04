/*
 * Copyright 2022 ByteDance Inc.
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

package utf8

import (
    `testing`
    `github.com/stretchr/testify/assert`
    `unicode/utf8`
    `bytes`
    `math/rand`
)

var (
    _Header_2Bytes  = string([]byte{0xC0})
    _Header_3Bytes  = string([]byte{0xE0})
    _Header_4Bytes  = string([]byte{0xF0})
    _Low_Surrogate  = string([]byte{0xED, 0xA0, 0x80}) // \ud800
    _High_Surrogate = string([]byte{0xED, 0xB0, 0x80}) // \udc00
    _Cont           = "\xb0"
)

func genRandBytes(length int) []byte {
    var buf bytes.Buffer
    for j := 0; j < length; j++ {
        buf.WriteByte(byte(rand.Intn(0xFF + 1)))
    }
    return buf.Bytes()
}

func genRandAscii(length int) []byte {
    var buf bytes.Buffer
    for j := 0; j < length; j++ {
        buf.WriteByte(byte(rand.Intn(0x7F + 1)))
    }
    return buf.Bytes()
}

func genRandRune(length int) []byte {
    var buf bytes.Buffer
    for j := 0; j < length; j++ {
        buf.WriteRune(rune(rand.Intn(0x10FFFF + 1)))
    }
    return buf.Bytes()
}

func TestValidate_Random(t *testing.T) {
    // compare with stdlib
    compare := func(t *testing.T, data []byte) {
        assert.Equal(t, utf8.Valid(data), Validate(data), string(data))
    }

	compare(t, []byte{})
	compare(t, nil)

    // random testing
    nums   := 1000
    maxLen := 1000
    for i := 0; i < nums; i++ {
        length := rand.Intn(maxLen)
        compare(t, genRandBytes(length))
        compare(t, genRandRune(length))
    }
}

func BenchmarkValidate(b *testing.B) {
    bench := []struct {
        name string
        data []byte
    } {
        {"ValidAscii", genRandAscii(1000)},
        {"ValidUTF8",  genRandRune(1000)},
        {"RandomBytes", genRandBytes(1000)},
    }

    for _, test := range bench {
        if utf8.Valid(test.data) != Validate(test.data) {
            b.Fatalf("sonic utf8 validate wrong for %s string: %v", test.name, test.data)
        }
        b.Run("Sonic_" + test.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                Validate(test.data)
            }
        })
        b.Run("StdLib_" + test.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                utf8.Valid(test.data)
            }
        })
    }
}
