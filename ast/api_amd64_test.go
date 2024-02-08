// +build amd64,go1.16,!go1.23

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

package ast

import (
    `encoding/json`
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/require`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/stretchr/testify/assert`
)

func TestEncodeValue(t *testing.T) {
    obj := new(_TwitterStruct)
    if err := json.Unmarshal([]byte(_TwitterJson), obj); err != nil {
        t.Fatal(err)
    }
    // buf, err := encoder.Encode(obj, encoder.EscapeHTML|encoder.SortMapKeys)
    buf, err := json.Marshal(obj)
    if err != nil {
        t.Fatal(err)
    }
    quote, err := json.Marshal(_TwitterJson)
    if err != nil {
        t.Fatal(err)
    }
    type Case struct {
        node Node
        exp string
        err bool
    }
    input := []Case{
        {NewNull(), "null", false},
        {NewBool(true), "true", false},
        {NewBool(false), "false", false},
        {NewNumber("0.0"), "0.0", false},
        {NewString(""), `""`, false},
        {NewString(`\"\"`), `"\\\"\\\""`, false},
        {NewString(_TwitterJson), string(quote), false},
        {NewArray([]Node{}), "[]", false},
        {NewArray([]Node{NewString(""), NewNull()}), `["",null]`, false},
        {NewArray([]Node{NewBool(true), NewString("true"), NewString("\t")}), `[true,"true","\t"]`, false},
        {NewObject([]Pair{Pair{"a", NewNull()}, Pair{"b", NewNumber("0")}}), `{"a":null,"b":0}`, false},
        {NewObject([]Pair{Pair{"\ta", NewString("\t")}, Pair{"\bb", NewString("\b")}, Pair{"\nb", NewString("\n")}, Pair{"\ra", NewString("\r")}}),`{"\ta":"\t","\u0008b":"\u0008","\nb":"\n","\ra":"\r"}`, false},
        {NewObject([]Pair{}), `{}`, false},
        {NewObject([]Pair{Pair{Key: "", Value: NewNull()}}), `{"":null}`, false},
        {NewBytes([]byte("hello, world")), `"aGVsbG8sIHdvcmxk"`, false},
        {NewAny(obj), string(buf), false},
        {NewRaw(`[{ }]`), "[{}]", false},
        {Node{}, "", true},
        {Node{t: types.ValueType(1)}, "", true},
    }
    for i, c := range input {
        t.Log(i)
        buf, err := json.Marshal(&c.node)
        if c.err {
            if err == nil {
                t.Fatal(i)
            }
            continue
        }
        if err != nil {
            t.Fatal(i, err)
        }
        assert.Equal(t, c.exp, string(buf))
    }
}

func TestSortNodeTwitter(t *testing.T) {
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    obj, err := root.MapUseNumber()
    if err != nil {
        t.Fatal(err)
    }
    exp, err := encoder.Encode(obj, encoder.SortMapKeys)
    if err != nil {
        t.Fatal(err)
    }
    var expObj interface{}
    require.NoError(t, json.Unmarshal(exp, &expObj))

    if err := root.SortKeys(true); err != nil {
        t.Fatal(err)
    }
    act, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    var actObj interface{}
    require.NoError(t, json.Unmarshal(act, &actObj))
    require.Equal(t, expObj, actObj)
    require.Equal(t, len(exp), len(act))
    require.Equal(t, string(exp), string(act))
}