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

package compat

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

var jt = jsoniter.Config{
    ValidateJsonRawMessage: true,
}.Froze()

func TestMarshalMarshal(t *testing.T){
    var obj = map[string]interface{}{
        "c": "<&> [ ] ",
    }
    sout, serr := Marshal(obj)
    jout, jerr := jt.Marshal(obj)
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))
}

func TestUnmarshal(t *testing.T){
    var sobj = map[string]interface{}{}
    var jobj = map[string]interface{}{}
    var data = []byte(`{"a":1.00000001E-10}`)
    var str = mem2Str(data)
    serr := UnmarshalString(str, &sobj)
    jerr := jt.UnmarshalFromString(str, &jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)
    // data[2] = '0'
    // require.Equal(t, jobj, sobj)
}

func TestMarshalStd(t *testing.T){
    var obj = map[string]interface{}{
        "c": "<&>",
        "b": json.RawMessage(" [ ] "),
    }
    sout, serr := MarshalStd(obj)
    jout, jerr := json.Marshal(obj)
    require.Equal(t, jerr, serr)
    require.Equal(t, string(jout), string(sout))
}

func TestUnmarshalStd(t *testing.T){
    var sobj = map[string]interface{}{}
    var jobj = map[string]interface{}{}
    var data = []byte(`{"a":1.00000001E-10}`)
    var str = mem2Str(data)
    serr := UnmarshalStringStd(str, &sobj)
    jerr := json.Unmarshal(data, &jobj)
    require.Equal(t, jerr, serr)
    require.Equal(t, jobj, sobj)
    data[2] = '0'
    require.Equal(t, jobj, sobj)
}