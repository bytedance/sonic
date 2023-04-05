/*
 * Copyright 2023 ByteDance Inc.
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
    `testing`
    `github.com/bytedance/sonic`
    `encoding/json`
    `github.com/stretchr/testify/require`
)

type DirectStruct struct {
    Ptr *int
}

func (d DirectStruct) MarshalJSON() ([]byte, error) {
    return json.Marshal((d.Ptr))
}

type DirectArray [1]*int

func (d DirectArray) MarshalJSON() ([]byte, error) {
    return json.Marshal(d[0])
}

type DirectNested [1]DirectStruct

func (d DirectNested) MarshalJSON() ([]byte, error) {
    return json.Marshal(d[0])
}


type DirectStruct2 struct {
    Ptr *int
}

func (d DirectStruct2) MarshalText() ([]byte, error) {
    return json.Marshal((d.Ptr))
}

type DirectArray2 [1]*int

func (d DirectArray2) MarshalText() ([]byte, error) {
    return json.Marshal(d[0])
}

type DirectNested2 [1]DirectStruct2

func (d DirectNested2) MarshalText() ([]byte, error) {
    return json.Marshal(d[0])
}

func TestDirectStructType(t *testing.T) {
    val  := 123
    real := &val
    realds := DirectStruct{real}
    nullds := DirectStruct{}
    realda := DirectArray{real}
    nullda := DirectArray{}
    nested := DirectNested{realds}

    realds2 := DirectStruct2{real}
    nullds2 := DirectStruct2{}
    realda2 := DirectArray2{real}
    nullda2 := DirectArray2{}
    nested2 := DirectNested2{realds2}

    tests := []interface{} {
        // test direct iface type implemented encoding.JSONMarshaler
        &realds, realds, &nullds, nullds,
        map[string]DirectStruct{ "a": realds, "b": nullds},
        map[string]*DirectStruct{ "a": &realds, "b": &nullds},
        []DirectStruct{realds, nullds},
        []*DirectStruct{&realds, &nullds},
        &realda, realda, &nullda, nullda,
        nested, &nested,

        // test direct iface implemented encoding.TextMarshaler
        &realds2, realds2, &nullds2, nullds2,
        map[string]DirectStruct2{ "a": realds2, "b": nullds2},
        map[string]*DirectStruct2{ "a": &realds2, "b": &nullds2},
        []DirectStruct2{realds2, nullds2},
        []*DirectStruct2{&realds2, &nullds2},
        &realda2, realda2, &nullda2, nullda2,
        nested2, &nested2,

        // test map key implement encoding.TextMarshaler
        map[DirectStruct2]DirectArray{
            realds2 : realda,
            nullds2 : nullda,
        },
    }
    for _, tt := range tests {
        jout, jerr := json.Marshal(tt)
        sout, serr := sonic.ConfigStd.Marshal(tt)
        require.Equal(t, string(jout), string(sout))
        require.NoError(t, jerr)
        require.Equal(t, jerr, serr)
    }
}


type recurePtr struct {
    Name string
    Recur *recurePtr
}

func TestRecursiveIssue(t *testing.T) {
    data := `{
        "Name": "",
        "Recur": null
    }`
    jv, sv := recurePtr{}, recurePtr{}
    jerr := json.Unmarshal([]byte(data), &jv)
    serr := sonic.Unmarshal([]byte(data), &sv)
    require.Equal(t, jv, sv)
    require.Equal(t, jerr, serr)
}

