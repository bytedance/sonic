//go:build go1.18
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
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"testing"
	"time"
	_ "unicode/utf8"

	"github.com/bytedance/gopkg/util/gctuner"
	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func FuzzMain(f *testing.F) {
    for _, corp := range(corpus()) {
        f.Add(corp)
    }
    f.Fuzz(fuzzMain)
}

// Used for debug falied fuzz corpus
func TestCorpus(t *testing.T) {
    fuzzMain(t, []byte("[1\x00"))
    fuzzMain(t, []byte("\"\\uDE1D\\uDE1D\\uDEDD\\uDE1D\\uDE1D\\uDE1D\\uDE1D\\uDEDD\\uDE1D\""))
    // fuzzMain(t, []byte(`{"":null}`))
}

var sonicConfigStdUseNumber = sonic.Config{
	EscapeHTML : true,
	SortMapKeys: true,
	CompactMarshaler: true,
	CopyString : true,
	ValidateString : true,
	UseNumber: true,
}.Froze()

func jsonUnmarshal(data []byte, v any) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.UseNumber()
	return dec.Decode(v)
}

type NewType = func() interface{}

type Api struct {
	Unmarshal func(data []byte, v any) error
	Marshal func(v any) ([]byte, error)
}

var target = sonic.ConfigStd
var sonicStdDefault Api = Api {
	Unmarshal: sonic.ConfigStd.Unmarshal,
	Marshal: sonic.ConfigStd.Marshal,
}

var stdjsonDefault Api = Api {
	Unmarshal: json.Unmarshal,
	Marshal: json.Marshal,
}

var sonicDefaultUseNumber Api = Api {
	Unmarshal: sonicConfigStdUseNumber.Unmarshal,
	Marshal: sonicConfigStdUseNumber.Marshal,
}

var stdjsonUseNumber Api = Api {
	Unmarshal: func (data []byte, v any) error {
		dec := json.NewDecoder(bytes.NewBuffer(data))
		dec.UseNumber()
		return dec.Decode(v)
	},
	Marshal: json.Marshal,
}

func fuzzUnmarshal(t *testing.T, data []byte, typ NewType, jstd, sour Api) ([]byte, any) {
	var sv = typ()
	var jv = typ()
	serr := sour.Unmarshal(data, sv)
	jerr := jstd.Unmarshal(data, jv)
	require.Equal(t, serr != nil, jerr != nil, 
			dump(data, jv, jerr, sv, serr))
	if jerr != nil {
		return nil, nil
	}
	require.Equal(t, sv, jv, dump(data, jv, jerr, sv, serr))

	v := jv
	sout, serr := sour.Marshal(v)
	jout, jerr := jstd.Marshal(v)
	require.NoError(t, serr, dump(v, jout, jerr, sout, serr))
	require.NoError(t, jerr, dump(v, jout, jerr, sout, serr))

	// compare the marshal result
	{
		sv, jv = typ(), typ()
		serr := sour.Unmarshal(sout, sv)
		jerr := jstd.Unmarshal(jout, jv)
		require.Equalf(t, serr != nil, jerr != nil, dump(data, jv, jerr, sv, serr))
		if jerr != nil {
			return nil, nil
		}
		require.Equal(t, sv, jv, dump(data, jv, jerr, sv, serr))
	}
	return jout, sv
}

// compare the unmarshaled result to compare two jsons. useNumber to make it not
// not return error when very large numbers
func assertJsonEqual(t *testing.T, json1, json2 []byte, msg string) {
	var v1, v2 interface{}
	err1 := stdjsonUseNumber.Unmarshal(json1, &v1)
	err2 := stdjsonUseNumber.Unmarshal(json2, &v2)
	require.Equal(t, err1, nil, msg)
	require.Equal(t, err2, nil, msg)
	require.Equal(t, v1, v2, msg)
}

func sonicAstMarshal(t *testing.T, data []byte) {
	root, aerr := sonic.Get(data)
	require.Equal(t, aerr, nil)
	aerr = root.LoadAll()
	require.Equal(t, aerr, nil, dump(data, root, aerr))
	aout, aerr := root.MarshalJSON()
	require.Equal(t, aerr, nil)
	assertJsonEqual(t, data, aout, dump(data, aout, aerr))
}

func fuzzMain(t *testing.T, data []byte) {
    fuzzValidate(t, data)
    fuzzHtmlEscape(t, data)
    // fuzz ast get api, should not panic here.
    fuzzAst(t, data)
    // Only fuzz the validate json here.
    if !json.Valid(data) {
        return
    }
	sonicAstMarshal(t, data)
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
		jout, sv := fuzzUnmarshal(t, data, typ, stdjsonDefault, sonicStdDefault)
		if i < 3 {
			fuzzUnmarshal(t, data, typ, stdjsonUseNumber, sonicDefaultUseNumber)
		}
        // if m, ok := sv.(*map[string]interface{}); ok {
        //     fuzzDynamicStruct(t, data, *m)
        //     fuzzASTGetFromObject(t, jout, *m)
        // }
        // if a, ok := sv.(*[]interface{}); ok {
        //     fuzzASTGetFromArray(t, jout, *a)
        // }
		var _, _ = jout, sv
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
