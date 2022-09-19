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

package unit_test

import (
    `bytes`
    `encoding/json`
    `testing`

    `github.com/bytedance/sonic`
    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/require`
)

var jt = jsoniter.Config{
    ValidateJsonRawMessage: true,
}.Froze()

func TestCompatMarshalDefault(t *testing.T) {
    var obj = map[string]interface{}{
        "c": json.RawMessage("[\"<&>\"]"),
    }
    sout, serr := sonic.ConfigDefault.Marshal(obj)
    jout, jerr := jt.Marshal(obj)
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))

    // obj = map[string]interface{}{
    //     "a": json.RawMessage(" [} "),
    // }
    // sout, serr = ConfigDefault.Marshal(obj)
    // jout, jerr = json.Marshal(obj)
    // require.NotNil(t, jerr)
    // require.NotNil(t, serr)
    // require.Equal(t, string(jout), string(sout))

    obj = map[string]interface{}{
        "a": json.RawMessage("1"),
    }
    sout, serr = sonic.ConfigDefault.MarshalIndent(obj, "", "  ")
    jout, jerr = jt.MarshalIndent(obj, "", "  ")
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))
}

func TestCompatMarshalStd(t *testing.T) {
    t.Parallel()
    var obj = map[string]interface{}{
        "c": json.RawMessage(" [ \"<&>\" ] "),
        "b": json.RawMessage(" [ ] "),
    }
    sout, serr := sonic.ConfigStd.Marshal(obj)
    jout, jerr := json.Marshal(obj)
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))

    obj = map[string]interface{}{
        "a": json.RawMessage(" [} "),
    }
    sout, serr = sonic.ConfigStd.Marshal(obj)
    jout, jerr = json.Marshal(obj)
    require.NotNil(t, jerr)
    require.NotNil(t, serr)
    require.Equal(t, string(jout), string(sout))

    obj = map[string]interface{}{
        "a": json.RawMessage("1"),
    }
    sout, serr = sonic.ConfigStd.MarshalIndent(obj, "xxxx", "  ")
    jout, jerr = json.MarshalIndent(obj, "xxxx", "  ")
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))
}

func TestCompatUnmarshalDefault(t *testing.T) {
    var sobj = map[string]interface{}{}
    var jobj = map[string]interface{}{}
    var data = []byte(`{"a":-0}`)
    var str = string(data)
    serr := sonic.ConfigDefault.UnmarshalFromString(str, &sobj)
    jerr := jt.UnmarshalFromString(str, &jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)

    x := struct{ A json.Number }{}
    y := struct{ A json.Number }{}
    data = []byte(`{"A":"1", "B":-1}`)
    serr = sonic.ConfigDefault.Unmarshal(data, &x)
    jerr = jt.Unmarshal(data, &y)
    require.Equal(t, jerr, serr)
    require.Equal(t, y, x)
}

func TestCompatUnmarshalStd(t *testing.T) {
    var sobj = map[string]interface{}{}
    var jobj = map[string]interface{}{}
    var data = []byte(`{"a":1.00000001E-10}`)
    var str = string(data)
    serr := sonic.ConfigStd.UnmarshalFromString(str, &sobj)
    jerr := json.Unmarshal(data, &jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)
    data[2] = '0'
    require.Equal(t, jobj, sobj)

    sobj = map[string]interface{}{}
    jobj = map[string]interface{}{}
    data = []byte(`{"a":1}`)
    cfg := sonic.Config{
        UseNumber: true,
    }.Froze()
    serr = cfg.Unmarshal(data, &sobj)
    dec := json.NewDecoder(bytes.NewBuffer(data))
    dec.UseNumber()
    jerr = dec.Decode(&jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)

    x := struct{ A json.Number }{}
    y := struct{ A json.Number }{}
    data = []byte(`{"A":"1", "B":-1}`)
    cfg = sonic.Config{
        DisallowUnknownFields: true,
    }.Froze()
    serr = cfg.Unmarshal(data, &x)
    dec = json.NewDecoder(bytes.NewBuffer(data))
    dec.UseNumber()
    dec.DisallowUnknownFields()
    jerr = dec.Decode(&y)
    require.Equal(t, jerr, serr)
    require.Equal(t, y, x)
}

func TestCompatEncoderDefault(t *testing.T) {
    var o = map[string]interface{}{
        "a": "<>",
        // "b": json.RawMessage(" [ ] "),
    }
    var w1 = bytes.NewBuffer(nil)
    var w2 = bytes.NewBuffer(nil)
    var enc1 = jt.NewEncoder(w1)
    var enc2 = sonic.ConfigDefault.NewEncoder(w2)

    require.Nil(t, enc1.Encode(o))
    require.Nil(t, enc2.Encode(o))
    require.Equal(t, w1.String(), w2.String())

    enc1.SetEscapeHTML(true)
    enc2.SetEscapeHTML(true)
    enc1.SetIndent("", "  ")
    enc2.SetIndent("", "  ")
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

func TestCompatEncoderStd(t *testing.T) {
    var o = map[string]interface{}{
        "a": "<>",
        "b": json.RawMessage(" [ ] "),
    }
    var w1 = bytes.NewBuffer(nil)
    var w2 = bytes.NewBuffer(nil)
    var enc1 = json.NewEncoder(w1)
    var enc2 = sonic.ConfigStd.NewEncoder(w2)

    require.Nil(t, enc1.Encode(o))
    require.Nil(t, enc2.Encode(o))
    require.Equal(t, w1.String(), w2.String())

    enc1.SetEscapeHTML(true)
    enc2.SetEscapeHTML(true)
    enc1.SetIndent("\n", "  ")
    enc2.SetIndent("\n", "  ")
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

func TestCompatDecoderStd(t *testing.T) {
    var o1 = map[string]interface{}{}
    var o2 = map[string]interface{}{}
    var s = `{"a":"b"} {"1":"2"} a {}`
    var w1 = bytes.NewBuffer([]byte(s))
    var w2 = bytes.NewBuffer([]byte(s))
    var enc1 = json.NewDecoder(w1)
    var enc2 = sonic.ConfigStd.NewDecoder(w2)

    require.Equal(t, enc1.More(), enc2.More())
    require.Nil(t, enc1.Decode(&o1))
    require.Nil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())

    require.Equal(t, enc1.More(), enc2.More())
    require.Nil(t, enc1.Decode(&o1))
    require.Nil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())

    require.Equal(t, enc1.More(), enc2.More())
    require.NotNil(t, enc1.Decode(&o1))
    require.NotNil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())
}

func TestCompatDecoderDefault(t *testing.T) {
    var o1 = map[string]interface{}{}
    var o2 = map[string]interface{}{}
    var s = `{"a":"b"} {"1":"2"} a {}`
    var w1 = bytes.NewBuffer([]byte(s))
    var w2 = bytes.NewBuffer([]byte(s))
    var enc1 = jt.NewDecoder(w1)
    var enc2 = sonic.ConfigDefault.NewDecoder(w2)

    require.Equal(t, enc1.More(), enc2.More())
    require.Nil(t, enc1.Decode(&o1))
    require.Nil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())

    require.Equal(t, enc1.More(), enc2.More())
    require.Nil(t, enc1.Decode(&o1))
    require.Nil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())

    require.Equal(t, enc1.More(), enc2.More())
    require.NotNil(t, enc1.Decode(&o1))
    require.NotNil(t, enc2.Decode(&o2))
    require.Equal(t, w1.String(), w2.String())

    // require.Equal(t, enc1.More(), enc2.More())
    // require.NotNil(t, enc1.Decode(&o1))
    // require.NotNil(t, enc2.Decode(&o2))
    // require.Equal(t, w1.String(), w2.String())
}
