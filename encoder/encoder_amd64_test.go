// +build amd64,go1.16,!go1.23

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

package encoder

import (
    `encoding/json`
    `testing`

    `github.com/stretchr/testify/require`
)

func TestOptionSliceOrMapNoNull(t *testing.T) {
    obj := sample{}
    out, err := Encode(obj, NoNullSliceOrMap)
    if err != nil {
        t.Fatal(err)
    }
    require.Equal(t, `{"M":{},"S":[],"A":[],"MP":null,"SP":null,"AP":null}`, string(out))

    obj2 := sample{}
    out, err = Encode(obj2, 0)
    if err != nil {
        t.Fatal(err)
    }
    require.Equal(t, `{"M":null,"S":null,"A":[],"MP":null,"SP":null,"AP":null}`, string(out))
}

func TestEncoder_Marshaler(t *testing.T) {
    v := MarshalerStruct{V: MarshalerImpl{X: 12345}}
    ret, err := Encode(&v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":12345    }`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":{"X":12345}}`, string(ret))

    ret2, err2 := Encode(&v, 0)
    require.NoError(t, err2)
    require.Equal(t, `{"V":12345    }`, string(ret2))
    ret3, err3 := Encode(v, CompactMarshaler)
    require.NoError(t, err3)
    require.Equal(t, `{"V":{"X":12345}}`, string(ret3))
}

func TestMarshalerError(t *testing.T) {
    v := MarshalerErrorStruct{}
    ret, err := Encode(&v, 0)
    require.EqualError(t, err, `invalid Marshaler output json syntax at 5: "[\"\"] {"`)
    require.Equal(t, []byte(nil), ret)
}


func TestEncoder_RawMessage(t *testing.T) {
    rms := RawMessageStruct{
        X: json.RawMessage("123456    "),
    }
    ret, err := Encode(&rms, 0)
    require.NoError(t, err)
    require.Equal(t, `{"X":123456    }`, string(ret))

    ret, err = Encode(&rms, CompactMarshaler)
    require.NoError(t, err)
    require.Equal(t, `{"X":123456}`, string(ret))
}


func TestEncoder_TextMarshaler(t *testing.T) {
    v := TextMarshalerStruct{V: TextMarshalerImpl{X: (`{"a"}`)}}
    ret, err := Encode(&v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":"{\"a\"}"}`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret))

    ret2, err2 := Encode(&v, NoQuoteTextMarshaler)
    require.NoError(t, err2)
    require.Equal(t, `{"V":{"a"}}`, string(ret2))
    ret3, err3 := Encode(v, NoQuoteTextMarshaler)
    require.NoError(t, err3)
    require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret3))
}

func TestEncoder_Marshal_EscapeHTML(t *testing.T) {
    v := map[string]TextMarshalerImpl{"&&":{"<>"}}
    ret, err := Encode(v, EscapeHTML)
    require.NoError(t, err)
    require.Equal(t, `{"\u0026\u0026":{"X":"\u003c\u003e"}}`, string(ret))
    ret, err = Encode(v, 0)
    require.NoError(t, err)
    require.Equal(t, `{"&&":{"X":"<>"}}`, string(ret))

    // “ is \xe2\x80\x9c, and ” is \xe2\x80\x9d,
    // similar as HTML escaped chars \u2028(\xe2\x80\xa8) and \u2029(\xe2\x80\xa9)
    m := map[string]string{"test": "“123”"}
    ret, err = Encode(m, EscapeHTML)
    require.Equal(t, string(ret), `{"test":"“123”"}`)
    require.NoError(t, err)

    m = map[string]string{"K": "\u2028\u2028\xe2"}
    ret, err = Encode(m, EscapeHTML)
    require.Equal(t, string(ret), "{\"K\":\"\\u2028\\u2028\xe2\"}")
    require.NoError(t, err)
}
