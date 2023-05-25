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
    `bytes`
    `encoding/json`
    `strings`
    `testing`

    `github.com/stretchr/testify/require`
)

func TestEncodeStream(t *testing.T) {
    var o = map[string]interface{}{
        "a": "<>",
        "b": json.RawMessage(" [ ] "),
    }
    var w1 = bytes.NewBuffer(nil)
    var w2 = bytes.NewBuffer(nil)
    var enc1 = json.NewEncoder(w1)
    var enc2 = NewStreamEncoder(w2)
    enc2.SetEscapeHTML(true)
    enc2.SortKeys()
    enc2.SetCompactMarshaler(true)

    require.Nil(t, enc1.Encode(o))
    require.Nil(t, enc2.Encode(o))
    require.Equal(t, w1.String(), w2.String())

    enc1.SetEscapeHTML(true)
    enc2.SetEscapeHTML(true)
    enc1.SetIndent("你好", "\b")
    enc2.SetIndent("你好", "\b")
    require.Nil(t, enc1.Encode(o))
    require.Nil(t, enc2.Encode(o))
    require.Equal(t, w1.String(), w2.String())

    enc1.SetEscapeHTML(false)
    enc2.SetEscapeHTML(false)
    enc1.SetIndent("", "")
    enc2.SetIndent("", "")
    require.Nil(t, enc1.Encode(o))
    require.Nil(t, enc2.Encode(o))
    require.Equal(t, w1.String(), w2.String())
}

func BenchmarkEncodeStream_Sonic(b *testing.B) {
    var o = map[string]interface{}{
        "a": `<`+strings.Repeat("1", 1024)+`>`,
        "b": json.RawMessage(` [ `+strings.Repeat(" ", 1024)+` ] `),
    }

    b.Run("single", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = NewStreamEncoder(w)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            w.Reset()
        }
    })

    b.Run("double", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = NewStreamEncoder(w)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            _ = enc.Encode(o)
            w.Reset()
        }
    })

    b.Run("compatible", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = NewStreamEncoder(w)
        enc.SetEscapeHTML(true)
        enc.SortKeys()
        enc.SetCompactMarshaler(true)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            w.Reset()
        }
    })
}

func BenchmarkEncodeStream_Std(b *testing.B) {
    var o = map[string]interface{}{
        "a": `<`+strings.Repeat("1", 1024)+`>`,
        "b": json.RawMessage(` [ `+strings.Repeat(" ", 1024)+` ] `),
    }

    b.Run("single", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = json.NewEncoder(w)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            w.Reset()
        }
    })

    b.Run("double", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = json.NewEncoder(w)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            _ = enc.Encode(o)
            w.Reset()
        }
    })

    b.Run("compatible", func(b *testing.B){
        var w = bytes.NewBuffer(nil)
        var enc = json.NewEncoder(w)
        enc.SetEscapeHTML(true)
        b.ResetTimer()
        for i:=0; i<b.N; i++ {
            _ = enc.Encode(o)
            w.Reset()
        }
    })
}