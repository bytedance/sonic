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

package sonic

import (
    `bytes`
    `encoding/json`
    `reflect`
    `testing`

    `github.com/bytedance/sonic/option`
    `github.com/stretchr/testify/require`
)

func TestCompatUnmarshalStd(t *testing.T) {
    var sobj = map[string]interface{}{}
    var jobj = map[string]interface{}{}
    var data = []byte(`{"a":1.00000001E-10}`)
    var str = string(data)
    serr := ConfigStd.UnmarshalFromString(str, &sobj)
    jerr := json.Unmarshal(data, &jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)
    data[2] = '0'
    require.Equal(t, jobj, sobj)

    sobj = map[string]interface{}{}
    jobj = map[string]interface{}{}
    data = []byte(`{"a":1}`)
    cfg := Config{
        UseNumber: true,
    }.Froze()
    serr = cfg.Unmarshal(data, &sobj)
    dec := json.NewDecoder(bytes.NewBuffer(data))
    dec.UseNumber()
    jerr = dec.Decode(&jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)

    x := struct{
        A json.Number
        B json.Number
    }{}
    y := struct{
        A json.Number
        B json.Number
    }{}
    data = []byte(`{"A":"1", "C":-1, "B":1}`)
    cfg = Config{
        DisallowUnknownFields: true,
    }.Froze()
    serr = cfg.Unmarshal(data, &x)
    dec = json.NewDecoder(bytes.NewBuffer(data))
    dec.UseNumber()
    dec.DisallowUnknownFields()
    jerr = dec.Decode(&y)
    require.Equal(t, jerr, serr)
    // require.Equal(t, y, x)
}

func TestCompatMarshalStd(t *testing.T) {
    t.Parallel()
    var obj = map[string]interface{}{
        "c": json.RawMessage(" [ \"<&>\" ] "),
        "b": json.RawMessage(" [ ] "),
    }
    sout, serr := ConfigStd.Marshal(obj)
    jout, jerr := json.Marshal(obj)
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))

    obj = map[string]interface{}{
        "a": json.RawMessage(" [} "),
    }
    sout, serr = ConfigStd.Marshal(obj)
    jout, jerr = json.Marshal(obj)
    require.NotNil(t, jerr)
    require.NotNil(t, serr)
    require.Equal(t, string(jout), string(sout))

    obj = map[string]interface{}{
        "a": json.RawMessage("1"),
    }
    sout, serr = ConfigStd.MarshalIndent(obj, "xxxx", "  ")
    jout, jerr = json.MarshalIndent(obj, "xxxx", "  ")
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))
}

func TestCompatEncoderStd(t *testing.T) {
    var o = map[string]interface{}{
        "a": "<>",
        "b": json.RawMessage(" [ ] "),
    }
    var w1 = bytes.NewBuffer(nil)
    var w2 = bytes.NewBuffer(nil)
    var enc1 = json.NewEncoder(w1)
    var enc2 = ConfigStd.NewEncoder(w2)

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
    var enc2 = ConfigStd.NewDecoder(w2)

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

func TestPretouch(t *testing.T) {
    var v map[string]interface{}
    if err := Pretouch(reflect.TypeOf(v)); err != nil {
        t.Errorf("err:%v", err)
    }

    if err := Pretouch(reflect.TypeOf(v),
       option.WithCompileRecursiveDepth(1),
       option.WithCompileMaxInlineDepth(2),
    ); err != nil {
        t.Errorf("err:%v", err)
    }
}

func TestGet(t *testing.T) {
    var data = `{"a":"b"}`
    r, err := GetFromString(data, "a")
    if err != nil {
        t.Fatal(err)
    }
    v, err := r.String()
    if err != nil {
        t.Fatal(err)
    }
    if v != "b" {
        t.Fatal(v)
    }
}

