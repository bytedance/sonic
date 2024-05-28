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
    _ `unicode/utf8`
    `os`
    `runtime`
    `runtime/debug`
    `time`
    `io`
    `log`
    `strconv`

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
    `github.com/davecgh/go-spew/spew`
    `github.com/bytedance/gopkg/util/gctuner`
)

func FuzzMain(f *testing.F) {
    for _, corp := range(corpus()) {
        f.Add(corp)
    }
    f.Fuzz(fuzzMain)
}

type testFuzzCase struct {
	data []byte
	newf func () interface{}
}

func testJson(t *testing.T, data []byte, newf func() interface{}) {
	jv := newf()
	jerr := json.Unmarshal(data, jv)
	sv := newf()
	serr := sonic.Unmarshal(data, sv)
	require.Equal(t, jerr == nil, serr == nil)
	require.Equal(t, jv, sv)
}

var testFuzzCases = []testFuzzCase{
	{
		data: []byte(`{"x":"","":"","$$$$$ſ":"","RRRRRſ":"","ppppſ":"","ŝ":"","Ţ":"","ţ":"","Ť":"","Ũ":"","Ŭ":"","Ű":"","ų":"","Ŷ":"","Ÿ":"","ź":"","Ż":"","ſ":"","ſſ":"","ǿ":"","ɿ":"","տ":"","ٿſ":"","ڵ":""}`),
		newf: func() interface{} {
			return new(struct { F0 ***string; F1 string "json:\"ڵ,omitempty\""; F2 *string; F3 string; p4 string; F4 **string; F5 string; F6 *string "json:\"-\""; F7 ***string; F8 string; p9 string; F9 string; p10 string; F10 **string "json:\"Ŷ,\""; F11 **string "json:\"Ż,omitempty\""; F12 **string "json:\"ſ,\""; F13 ***string; F14 *string; p15 *string; F15 string "json:\"-\""; p16 string; F16 **string "json:\"ſſ,omitempty\""; F17 **string "json:\"ɿ,omitempty\""; p18 **string; F18 *string "json:\"-\""; F19 **string "json:\"RRRRRſ,omitempty\""; F20 ***string; p21 ***string; F21 string "json:\"ź,omitempty\""; p22 string })
		},
	},
	{
		data: []byte(`{"":"","$$$$$ſ":"","RRRRRſ":"","ppppſ":"","ŝ":"","Ţ":"","ţ":"","Ť":"","Ũ":"","Ŭ":"","Ű":"","ų":"","Ŷ":"","Ÿ":"","ź":"","Ż":"","ſ":"","ſſ":"","ǿ":"","ɿ":"","տ":"","ٿſ":"","ڵ":""}`),
		newf:  func() interface{} {
				return new(struct { F6 **string "json:\"x,\""; F7 string; p8 string; F8 *string; F9 string "json:\"ٿſ,omitempty\""; p10 string; F10 **string "json:\"-\""; F11 string "json:\"$$$$$ſ,\""; p12 string; F12 *string; p13 *string; F13 *string "json:\"Ű,\""; p14 *string; F14 **string; F15 string "json:\"Ż,omitempty\""; F16 string "json:\"-\""; p17 string; F17 **string "json:\"-\""; F18 string "json:\"ppppſ,\""; F19 ***string "json:\"Ţ,omitempty\""; p20 ***string; F20 ***string; F21 *string })
			},
	},
	{
		// FIXME: encoding/json has bugs because the limited dbuf capcaity is 800?
		data: []byte("[53333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333353333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333e-913]"),
		newf: func() interface{} { return new([]interface{}) },
	},
}

// Used for debug falied fuzz corpus
func TestFuzzCases(t *testing.T) {
	for _, c := range testFuzzCases {
		testJson(t, c.data, c.newf)
	}
    fuzzMain(t, []byte("[1\x00"))
    fuzzMain(t, []byte("\"\\uDE1D\\uDE1D\\uDEDD\\uDE1D\\uDE1D\\uDE1D\\uDE1D\\uDEDD\\uDE1D\""))
    // fuzzMain(t, []byte(`{"":null}`))
}

var target = sonic.ConfigStd

func fuzzMain(t *testing.T, data []byte) {
    fuzzValidate(t, data)
    fuzzHtmlEscape(t, data)
    // fuzz ast get api, should not panic here.
    fuzzAst(t, data)
    // Only fuzz the validate json here.
    if !json.Valid(data) {
        return
    }
    fuzzStream(t, data)
    for i, typ := range []func() interface{}{
        func() interface{} { return new(interface{}) },
        func() interface{} { return new(map[string]interface{}) },
        func() interface{} { return new([]interface{}) },
        func() interface{} { return new(string) },
        func() interface{} { return new(int64) },
        func() interface{} { return new(uint64) },
        func() interface{} { return new(float64) },
        // func() interface{} { return new(json.Number) },
        // func() interface{} { return new(S) },
    } {
        var sv = typ()
        var jv = typ()
        serr := target.Unmarshal(data, sv)
        jerr := json.Unmarshal(data, jv)
        require.Equal(t, serr != nil, jerr != nil, 
                dump(data, jv, jerr, sv, serr))
        if jerr != nil {
            continue
        }
        require.Equal(t, sv, jv, dump(string(data), jv, jerr, sv, serr))
    
        v := jv
        sout, serr := target.Marshal(v)
        jout, jerr := json.Marshal(v)
        require.NoError(t, serr, dump(v, jout, jerr, sout, serr))
        require.NoError(t, jerr, dump(v, jout, jerr, sout, serr))

        {
            sv, jv = typ(), typ()
            serr := target.Unmarshal(sout, sv)
            jerr := json.Unmarshal(jout, jv)
            require.Equalf(t, serr != nil, jerr != nil, dump(data, jv, jerr, sv, serr))
            if jerr != nil {
                continue
            }
            require.Equal(t, sv, jv, dump(data, jv, jerr, sv, serr))
        }

        // fuzz ast MarshalJSON API
        if i == 0 {
            root, aerr := sonic.Get(data)
            require.Equal(t, aerr, nil)
            aerr = root.LoadAll()
            require.Equal(t, aerr, nil, dump(data, jv, jerr, root, aerr))
            aout, aerr := root.MarshalJSON()
            require.Equal(t, aerr, nil)
            sv = typ()
            serr := json.Unmarshal(aout, sv)
            require.Equal(t, serr, nil)
            require.Equal(t, sv, jv, dump(data, jv, jerr, sv, serr))
        }

        if m, ok := sv.(*map[string]interface{}); ok {
            fuzzASTGetFromObject(t, jout, *m)
            fuzzDynamicStruct(t, jout, *m)
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


func dump(args ...interface{}) string {
    return spew.Sdump(args)
}

func fdump(w io.Writer, args ...interface{}) {
    spew.Fdump(w, args)
}

const (
    MemoryLimitEnv = "SONIC_FUZZ_MEM_LIMIT"
    AsynyncGCEnv   = "SONIC_NO_ASYNC_GC"
    KB      uint64 = 1024
    MB      uint64 = 1024 * KB
    GB      uint64 = 1024 * MB
)

func setMemLimit(limit uint64) {
    threshold := uint64(float64(limit) * 0.7)
    numWorker := uint64(runtime.GOMAXPROCS(0))
    if os.Getenv(MemoryLimitEnv) != "" {
        if memGB, err := strconv.ParseUint(os.Getenv(MemoryLimitEnv), 10, 64); err == nil {
            limit = memGB * GB
        }
    }
    gctuner.Tuning(threshold / numWorker)
    log.Printf("[%d] Memory Limit: %d GB, Memory Threshold: %d MB\n", os.Getpid(), limit/GB, threshold/MB)
    log.Printf("[%d] Memory Threshold Per Worker: %d MB\n", os.Getpid(), threshold/numWorker/MB)
}

func enableSyncGC() {
    var debugAsyncGC = os.Getenv("AsynyncGCEnv") == ""
    go func ()  {
        if !debugAsyncGC {
            return 
        }
        log.Printf("Begin GC looping...")
        for {
            runtime.GC()
            debug.FreeOSMemory() 
        }
    }()
}

func TestMain(m *testing.M) {
    // Avoid OOM
    setMemLimit(12 * GB)
    enableSyncGC()
    time.Sleep(time.Millisecond)
    m.Run()
}
