// +build go1.18

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

package sonic_fuzz

import (
	`encoding/json`
    `testing`
    `unicode/utf8`
    `reflect`
    `os`
    `runtime`
    `runtime/debug`
    `time`

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

func FuzzMain(f *testing.F) {
    for _, corp := range(corpus()) {
        f.Add(corp)
    }
    f.Fuzz(fuzzMain)
}

func fuzzMain(t *testing.T, data []byte) {
    fuzzValidate(t, data)
    fuzzHtmlEscape(t, data)
    // Only fuzz the validate json here, because the default configuration does not have validation in SONIC.
    if !utf8.Valid(data) || !json.Valid(data) {
        return
    }
    for _, typ := range []func() interface{}{
        func() interface{} { return new(interface{}) },
        func() interface{} { return new(map[string]interface{}) },
        func() interface{} { return new([]interface{}) },
        func() interface{} { return new(string) },
        func() interface{} { return new(int64) },
        func() interface{} { return new(uint64) },
        func() interface{} { return new(float64) },
        func() interface{} { return new(json.Number) },
        func() interface{} { return new(S) },
    } {
        sv, jv := typ(), typ()
        serr := sonic.Unmarshal([]byte(data), sv)
        jerr := json.Unmarshal([]byte(data), jv)
        require.Equalf(t, serr != nil, jerr != nil, "different error in sonic unmarshal %v", reflect.TypeOf(jv))
        if jerr != nil {
            continue
        }
        require.Equal(t, sv, jv, "different result in sonic unmarshal %v", reflect.TypeOf(jv))
        sout, serr := sonic.Marshal(sv)
        jout, jerr := json.Marshal(jv)
        require.NoError(t, serr, "error in sonic marshal %v", reflect.TypeOf(jv))
        require.NoError(t, jerr, "error in json marshal %v", reflect.TypeOf(jv))

        {
            sv, jv := typ(), typ()
            serr := sonic.Unmarshal(sout, sv)
            jerr := json.Unmarshal(jout, jv)
            require.Equalf(t, serr != nil, jerr != nil, "different error in sonic unmarshal again %v", reflect.TypeOf(jv))
            if jerr != nil {
                continue
            }
            require.Equal(t, sv, jv, "different result in sonic unmarshal again %v", reflect.TypeOf(jv))
        }

        if m, ok := sv.(*map[string]interface{}); ok {
            fuzzDynamicStruct(t, jout, *m)
            fuzzASTGetFromObject(t, jout, *m)
        }
        if a, ok := sv.(*[]interface{}); ok {
            fuzzASTGetFromArray(t, jout, *a)
        }
    }
}

type S struct {
	A int    `json:",omitempty"`
	B string `json:"B1,omitempty"`
	C float64
	D bool
	E uint8
	// F []byte // unmarshal []byte is different with encoding/json
	G interface{}
	H map[string]interface{}
	I map[string]string
	J []interface{}
	K []string
	L S1
	M *S1
	N *int
	O **int
    P int `json:",string"`
    Q float64 `json:",string"`
	R int `json:"-"`
	T struct {}
	U [2]int
	V uintptr
    W json.Number
    // X json.RawMessage
    Y Marshaller 
	Z TextMarshaller
}


type S1 struct {
	A int
	B string
}

type Marshaller struct {
	v string
}

func (m *Marshaller) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.v)
}

func (m *Marshaller) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.v)
}

type TextMarshaller struct {
    v int
}

func (k *TextMarshaller) MarshalText() ([]byte, error) {
    return json.Marshal(k.v)
}

func (k *TextMarshaller)  UnmarshalText(data []byte) error {
    return json.Unmarshal(data, &k.v)
}

var debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""

func TestMain(m *testing.M) {
    go func ()  {
        if !debugAsyncGC {
            return 
        }
        println("Begin GC looping...")
        for {
            runtime.GC()
            debug.FreeOSMemory() 
        }
    }()
    time.Sleep(time.Millisecond)
    m.Run()
}