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

package issue_test

import (
    `encoding`
    `encoding/json`
    `strconv`
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/require`
)

type ptrTextMarshaler string

func (k *ptrTextMarshaler) MarshalText() ([]byte, error) {
    if k == nil {
        return []byte("ptrisnil"), nil
    }
    return []byte("ptrto" + string(*k)), nil
}

type textMarshaler string

func (k textMarshaler) MarshalText() ([]byte, error) {
    return []byte(string(k)), nil
}

func TestIssue115_MarshalMapWithSort(t *testing.T) {
    nptext := (*ptrTextMarshaler)(nil)
    ptext  := ptrTextMarshaler("key")
    text0  := textMarshaler("")
    text1  := textMarshaler("1")
    text2  := textMarshaler("2")
    testCases := []struct {
        v    interface{}
        want string
    }{
        { v: map[string]int{"b":2, "a":1, "c":3}, want: `{"a":1,"b":2,"c":3}`},
        { v: map[int64]int{1:-1, -2:2, 0:0}, want: `{"-2":2,"0":0,"1":-1}`},
        { v: map[uint]int{1:-1, ^uint(0):2, 0:0}, want: `{"0":0,"1":-1,"18446744073709551615":2}`},
        { v: map[uintptr]int{uintptr(0xf):0xf, uintptr(0x0):0}, want: `{"0":0,"15":15}`}, 
        { v: map[encoding.TextMarshaler]interface{}{
               nptext : nil,
               &ptext : struct{}{},
               text0  : "", 
               &text1 : 1, 
               text2  : text2,
            },
          want: `{"":"","1":1,"2":"2","ptrisnil":null,"ptrtokey":{}}`,
        },
    }

    for _, tt := range testCases {
        out, err := encoder.Encode(tt.v, encoder.SortMapKeys)
        require.NoError(t, err)
        require.Equal(t, tt.want, string(out))
    }
}

func TestIssue115_MarshalLargeIntKeyMapWitSort(t *testing.T) {
    N := 10000
    m := map[int]string{}
    for i := 0; i < N; i++ {
        a := strconv.Itoa(i)
        m[i] = a
    }

    exp, err := json.Marshal(&m)
    require.NoError(t, err)
    got, err := encoder.Encode(&m, encoder.SortMapKeys)
    require.NoError(t, err)
    require.Equal(t, string(exp), string(got))
}

