/*
 * Copyright 2026 ByteDance Inc.
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

// Catalog of Sonic vs Go 1.27 encoding/json compatibility cases.
// See docs/sonic-go127-compatibility.md for the full table.
package compatibility_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/internal/envs"
	"github.com/stretchr/testify/require"
)

// ---- from case_float_overflow_test.go ----
// Case F32-OVF / F64-OVF: on Go 1.27 jsonv2 backend, std saturates overflow
// destinations to ±Inf while still returning an error. Sonic keeps dest at 0.

func TestCase_F32_OVF(t *testing.T) {
	const data = "3.402823567797337e+38"

	var s, j float32
	se := sonic.ConfigStd.Unmarshal([]byte(data), &s)
	je := json.Unmarshal([]byte(data), &j)

	require.Error(t, se)
	require.Error(t, je)
	require.Equal(t, float32(0), s, "sonic leaves dest at 0")

	if stdUsesJSONV2 {
		require.True(t, math.IsInf(float64(j), 1), "std jsonv2 saturates to +Inf")
	} else {
		require.Equal(t, float32(0), j, "legacy std leaves dest at 0")
	}
}

func TestCase_F64_OVF(t *testing.T) {
	const data = "1e1000"

	var s, j float64
	se := sonic.ConfigStd.Unmarshal([]byte(data), &s)
	je := json.Unmarshal([]byte(data), &j)

	require.Error(t, se)
	require.Error(t, je)
	require.Equal(t, float64(0), s, "sonic leaves dest at 0")

	if stdUsesJSONV2 {
		require.True(t, math.IsInf(j, 1), "std jsonv2 saturates to +Inf")
	} else {
		require.Equal(t, float64(0), j, "legacy std leaves dest at 0")
	}
}

// ---- from case_map_keys_test.go ----
// Case MAP-F64-ENC: Go 1.27 jsonv2 allows marshaling maps with float keys;
// sonic (and legacy std) reject them as unsupported.

func TestCase_MAP_F64_ENC(t *testing.T) {
	in := map[float64]string{1: ""}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)

	require.Error(t, serr)
	require.Nil(t, sout)

	if stdUsesJSONV2 {
		require.NoError(t, jerr)
		require.JSONEq(t, `{"1":""}`, string(jout))
	} else {
		require.Error(t, jerr)
	}
}

// Case MAP-PTR-ENC: pointer-keyed maps become marshalable under jsonv2 backend.

func TestCase_MAP_PTR_ENC(t *testing.T) {
	k, v := 1, 2
	in := map[*int]*int{&k: &v}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)

	require.Error(t, serr)
	require.Nil(t, sout)

	if stdUsesJSONV2 {
		require.NoError(t, jerr)
		require.JSONEq(t, `{"1":2}`, string(jout))
	} else {
		require.Error(t, jerr)
	}
}

// Case MAP-F64-DEC: sonic accepts float map keys. Legacy std rejects them;
// Go 1.27 jsonv2 accepts them (matching sonic).

func TestCase_MAP_F64_DEC(t *testing.T) {
	const data = `{"1.2":1.8}`
	var s, j map[float64]float64
	se := sonic.ConfigStd.Unmarshal([]byte(data), &s)
	je := json.Unmarshal([]byte(data), &j)

	require.NoError(t, se)
	require.Equal(t, map[float64]float64{1.2: 1.8}, s)

	if stdUsesJSONV2 {
		require.NoError(t, je)
		require.Equal(t, map[float64]float64{1.2: 1.8}, j)
	} else {
		require.Error(t, je, "legacy std cannot unmarshal into map[float64]float64")
	}
}

// ---- from case_string_option_test.go ----
type stringOptStruct struct {
	S *string `json:"s,string"`
	I *int    `json:"i,string"`
}

// Case STR-NULL: `,string` option with JSON string "null".
// Sonic clears pointer fields without error. Go 1.27 jsonv2 differs on *string.

func TestCase_STR_NULL(t *testing.T) {
	data := []byte(`{"s":"null","i":"null"}`)

	var s, j stringOptStruct
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	require.NoError(t, se)
	require.Nil(t, s.S)
	require.Nil(t, s.I)

	if stdUsesJSONV2 {
		var typeErr *json.UnmarshalTypeError
		require.ErrorAs(t, je, &typeErr)
		require.NotNil(t, j.S)
		require.Empty(t, *j.S)
		require.Nil(t, j.I)
	} else {
		require.NoError(t, je)
		require.Nil(t, j.S)
		require.Nil(t, j.I)
	}
}

// ---- from case_utf8_test.go ----
// Case UTF8-ENC: invalid UTF-8 in Go strings.
// Sonic ConfigStd escapes the replacement as \ufffd.
// Go 1.27 jsonv2 embeds the raw U+FFFD rune in the JSON string bytes.

func TestCase_UTF8_ENC(t *testing.T) {
	in := "a\xffb"
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)

	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.Equal(t, `"a\ufffdb"`, string(sout))

	if stdUsesJSONV2 {
		require.Equal(t, "\"a\ufffdb\"", string(jout))
	} else {
		require.Equal(t, string(sout), string(jout))
	}
}

// Case UTF8-DEC: both replace invalid UTF-8 on unmarshal (aligned under DefaultOptionsV1).

func TestCase_UTF8_DEC(t *testing.T) {
	data := []byte("\"\xff\"")
	var s, j string
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)
	require.NoError(t, se)
	require.NoError(t, je)
	require.Equal(t, "\ufffd", s)
	require.Equal(t, s, j)
}

// ---- from case_unmarshaler_test.go ----
type unjFoo struct {
	Name string
}

func (f *unjFoo) UnmarshalJSON(data []byte) error {
	f.Name = "Unmarshaler"
	return nil
}

// named pointer type whose element implements UnmarshalJSON
type unjMyPtr *unjFoo

// Case TMPTR-UNJ: named pointer + UnmarshalJSON dispatch.
// Sonic follows classic v1 (does not call UnmarshalJSON through named pointer).
// Go 1.27 jsonv2 calls UnmarshalJSON.

func TestCase_TMPTR_UNJ(t *testing.T) {
	data := []byte(`{"Name":"MyPtr"}`)

	s := unjMyPtr(&unjFoo{})
	j := unjMyPtr(&unjFoo{})
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)
	require.NoError(t, se)
	require.NoError(t, je)
	require.NotNil(t, s)
	require.NotNil(t, j)

	require.Equal(t, "MyPtr", s.Name, "sonic uses default decode into named pointer")

	if stdUsesJSONV2 {
		require.Equal(t, "Unmarshaler", j.Name, "std jsonv2 invokes UnmarshalJSON")
	} else {
		require.Equal(t, "MyPtr", j.Name, "legacy std matches sonic")
	}
}

type unjDate int64

func (d *unjDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = unjDate(t.Unix())
	return nil
}

type unjPartial struct {
	D unjDate `json:"D"`
	E int     `json:"E"`
}

// Case UNJ-PARTIAL: error-time partial writes currently depend on Sonic's
// decoder. JIT leaves later fields untouched; OptDec records the first error
// and continues populating fields. Go 1.27 jsonv2 also populates later fields.

func TestCase_UNJ_PARTIAL(t *testing.T) {
	data := []byte(`{"D":11,"E":1}`)
	var s, j unjPartial
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	require.Error(t, se)
	require.Error(t, je)
	if envs.UseOptDec {
		require.Equal(t, 1, s.E, "Sonic OptDec applies later fields after UnmarshalJSON error")
	} else {
		require.Equal(t, 0, s.E, "Sonic JIT does not apply later fields after UnmarshalJSON error")
	}

	if stdUsesJSONV2 {
		require.Equal(t, 1, j.E, "std jsonv2 still applies later fields")
	} else {
		require.Equal(t, 0, j.E, "legacy std matches sonic")
	}
}

// ---- from case_text_marshaler_test.go ----
// TextMarshaler that returns already-quoted text (historical Issue827 string3 case).
type quotedTextKey string

func (s quotedTextKey) MarshalText() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

type plainTextKey string

func (s plainTextKey) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

// Case TM-KEY-QUOTE: when MarshalText returns bytes that look like a JSON string
// (including quotes), sonic uses them as the map key content without re-escaping
// the surrounding quotes into the key; Go 1.27 jsonv2 re-escapes them.

func TestCase_TM_KEY_QUOTE(t *testing.T) {
	in := map[quotedTextKey]int{"one": 1, "two": 2}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)
	require.NoError(t, serr)
	require.NoError(t, jerr)

	require.JSONEq(t, `{"one":1,"two":2}`, string(sout))

	if stdUsesJSONV2 {
		require.JSONEq(t, `{"\"one\"":1,"\"two\"":2}`, string(jout))
	} else {
		require.JSONEq(t, string(sout), string(jout))
	}
}

func TestCase_TM_KEY_PLAIN(t *testing.T) {
	in := map[plainTextKey]int{"one": 1, "two": 2}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)
	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.JSONEq(t, `{"one":1,"two":2}`, string(sout))
	require.JSONEq(t, string(sout), string(jout), "plain TextMarshaler keys stay aligned")
}

// ---- from case_iface_skip_test.go ----
type mockEface interface {
	myMock()
}

type fooEfacePtr struct {
	Id *string `json:"id"`
}

func (f *fooEfacePtr) myMock() {}

type fooEfaceValue struct {
	Id *string `json:"id"`
}

func (f fooEfaceValue) myMock() {}

type wrapperEface struct {
	Name string    `json:"name"`
	Id   mockEface `json:"id"`
}

// Case IFACE-SKIP: unmarshaling into iface-typed fields that hold a non-pointer
// concrete value or a typed nil. Under Go 1.27 jsonv2 the error/value outcome
// diverges from classic v1 / sonic.

func TestCase_IFACE_SKIP_nonPtr(t *testing.T) {
	data := []byte(`{"id":{"id":"2"},"name":"name"}`)
	objS := wrapperEface{Id: fooEfaceValue{}}
	objJ := wrapperEface{Id: fooEfaceValue{}}

	se := sonic.ConfigStd.Unmarshal(data, &objS)
	je := json.Unmarshal(data, &objJ)

	if stdUsesJSONV2 {
		require.Error(t, se)
		var typeErr *json.UnmarshalTypeError
		require.ErrorAs(t, je, &typeErr)
		require.Equal(t, "name", objS.Name)
		require.IsType(t, fooEfaceValue{}, objS.Id)
		require.Equal(t, "name", objJ.Name)
		require.Nil(t, objJ.Id)
		return
	}

	require.Equal(t, se == nil, je == nil, "sonic=%v std=%v", se, je)
	require.Equal(t, objS.Name, objJ.Name)
	require.Equal(t, objS.Id, objJ.Id)
}

func TestCase_IFACE_SKIP_nilPtr(t *testing.T) {
	data := []byte(`{"id":{"id":"2"},"name":"name"}`)
	objS := wrapperEface{Id: (*fooEfacePtr)(nil)}
	objJ := wrapperEface{Id: (*fooEfacePtr)(nil)}

	se := sonic.ConfigStd.Unmarshal(data, &objS)
	je := json.Unmarshal(data, &objJ)

	if stdUsesJSONV2 {
		require.Error(t, se)
		var typeErr *json.UnmarshalTypeError
		require.ErrorAs(t, je, &typeErr)
		require.Equal(t, "name", objS.Name)
		typedNil, ok := objS.Id.(*fooEfacePtr)
		require.True(t, ok)
		require.Nil(t, typedNil)
		require.Equal(t, "name", objJ.Name)
		require.Nil(t, objJ.Id)
		return
	}

	require.Equal(t, se == nil, je == nil, "sonic=%v std=%v", se, je)
	require.Equal(t, objS.Name, objJ.Name)
	require.Equal(t, objS.Id, objJ.Id)
}

// ---- from case_self_ref_test.go ----
// Case SELF-REF: self-referential interface{} (`v = &v`).
// Sonic unmarshals successfully. Go 1.27 jsonv2 recurses until stack overflow,
// so this test must NOT call encoding/json under stdUsesJSONV2.

func TestCase_SELF_REF(t *testing.T) {
	var v interface{}
	v = &v
	err := sonic.ConfigStd.Unmarshal([]byte(`{"a":"b"}`), v)
	require.NoError(t, err)

	// Decoder writes through the self pointer and replaces the cycle with a map.
	m, ok := v.(map[string]interface{})
	require.True(t, ok, "got %#v", v)
	require.Equal(t, "b", m["a"])

	if stdUsesJSONV2 {
		t.Log("std encoding/json jsonv2 backend stack-overflows on this input; not invoked")
	}
}

// ---- from case_config_preset_test.go ----
// Case HTML-ESC: ConfigDefault does not escape HTML; ConfigStd / encoding/json do.

func TestCase_HTML_ESC(t *testing.T) {
	in := "<&>"
	def, err := sonic.ConfigDefault.Marshal(in)
	require.NoError(t, err)
	require.Equal(t, `"<&>"`, string(def))

	stdCfg, err := sonic.ConfigStd.Marshal(in)
	require.NoError(t, err)
	jout, jerr := json.Marshal(in)
	require.NoError(t, jerr)
	require.Equal(t, string(jout), string(stdCfg))
	require.Equal(t, `"\u003c\u0026\u003e"`, string(stdCfg))
}

// Case NIL-SLICE: default nil slice encodes as null (same as std).

func TestCase_NIL_SLICE(t *testing.T) {
	var sl []int
	def, err := sonic.ConfigDefault.Marshal(sl)
	require.NoError(t, err)
	require.Equal(t, "null", string(def))

	stdCfg, err := sonic.ConfigStd.Marshal(sl)
	require.NoError(t, err)
	jout, jerr := json.Marshal(sl)
	require.NoError(t, jerr)
	require.Equal(t, "null", string(stdCfg))
	require.Equal(t, string(jout), string(stdCfg))
}

// Case CASE-FOLD: default unmarshal is case-insensitive, matching std v1 /
// Go 1.27 DefaultOptionsV1.

func TestCase_CASE_FOLD(t *testing.T) {
	type S struct{ Foo int }
	data := []byte(`{"foo":1}`)
	var s, j S
	require.NoError(t, sonic.ConfigStd.Unmarshal(data, &s))
	require.NoError(t, json.Unmarshal(data, &j))
	require.Equal(t, 1, s.Foo)
	require.Equal(t, j, s)

	api := sonic.Config{CaseSensitive: true}.Froze()
	var strict S
	require.NoError(t, api.Unmarshal(data, &strict))
	require.Equal(t, 0, strict.Foo, "CaseSensitive skips folded name")
}

// ---- from case_aligned_test.go ----
// Cases that remain aligned between sonic and Go 1.27 encoding/json
// (DefaultOptionsV1). Kept so accidental std/sonic drift is noticed.

func TestCase_DUP_KEY(t *testing.T) {
	data := []byte(`{"a":1,"a":2}`)
	var s, j map[string]int
	require.NoError(t, sonic.ConfigStd.Unmarshal(data, &s))
	require.NoError(t, json.Unmarshal(data, &j))
	require.Equal(t, map[string]int{"a": 2}, s)
	require.Equal(t, s, j)
}

func TestCase_ARR_LEN(t *testing.T) {
	data := []byte(`[1,2,3]`)
	var s, j [2]int
	require.NoError(t, sonic.ConfigStd.Unmarshal(data, &s))
	require.NoError(t, json.Unmarshal(data, &j))
	require.Equal(t, [2]int{1, 2}, s)
	require.Equal(t, s, j)
}

func TestCase_BYTE_ARR(t *testing.T) {
	in := [3]byte{1, 2, 3}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)
	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.Equal(t, `[1,2,3]`, string(sout))
	require.Equal(t, string(sout), string(jout))
}

func TestCase_OMIT_FALSE_PTR(t *testing.T) {
	type S struct {
		B *bool `json:",omitempty"`
	}
	f := false
	in := S{B: &f}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)
	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.JSONEq(t, `{"B":false}`, string(sout))
	require.JSONEq(t, string(sout), string(jout))
}

func TestCase_TRAIL(t *testing.T) {
	data := []byte(`{"a":1}x`)
	var s, j map[string]int
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)
	require.Error(t, se)
	require.Error(t, je)
}

// ---- additional gaps found by auditing encoding/json.DefaultOptionsV1 ----

// Case SYN-PREMUT: encoding/json validates the complete JSON syntax before
// applying semantic mutations. Sonic JIT has already written earlier fields
// when it discovers a later syntax error; OptDec matches encoding/json.
func TestCase_SYN_PREMUT(t *testing.T) {
	data := []byte(`{"A":1,"B":`)
	type S struct {
		A int
		B int
	}
	s := S{A: 9, B: 9}
	j := S{A: 9, B: 9}

	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	require.Error(t, se)
	require.Error(t, je)
	if envs.UseOptDec {
		require.Equal(t, S{A: 9, B: 9}, s, "Sonic OptDec validates syntax before mutation")
	} else {
		require.Equal(t, S{A: 1, B: 9}, s, "Sonic JIT applies fields before the syntax error")
	}
	require.Equal(t, S{A: 9, B: 9}, j, "std validates syntax before mutation")
}

// Case MAP-VALUE-REPLACE: v1 replaces an existing map value before decoding
// an object into it. Sonic merges the object into the old map value instead.
func TestCase_MAP_VALUE_REPLACE(t *testing.T) {
	type V struct {
		Old bool
		New bool
	}
	data := []byte(`{"k":{"New":true}}`)
	s := map[string]V{"k": {Old: true}}
	j := map[string]V{"k": {Old: true}}

	require.NoError(t, sonic.ConfigStd.Unmarshal(data, &s))
	require.NoError(t, json.Unmarshal(data, &j))
	require.Equal(t, V{Old: true, New: true}, s["k"], "sonic merges into the old map value")
	require.Equal(t, V{New: true}, j["k"], "std replaces the old map value")
}

// Case STRING-GONUM: v1's ,string mode delegates numeric parsing to strconv.
// Sonic JIT rejects some Go numeric literals; OptDec matches encoding/json.
func TestCase_STRING_GONUM(t *testing.T) {
	type S struct {
		Int   int     `json:",string"`
		Float float64 `json:",string"`
	}

	tests := []struct {
		name string
		data string
		want S
	}{
		{name: "int-zero-padded", data: `{"Int":"00012"}`, want: S{Int: 12}},
		{name: "float-hex", data: `{"Float":"0x1_4p-2"}`, want: S{Float: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s, j S
			se := sonic.ConfigStd.Unmarshal([]byte(tt.data), &s)
			je := json.Unmarshal([]byte(tt.data), &j)

			require.NoError(t, je)
			require.Equal(t, tt.want, j)
			if envs.UseOptDec {
				require.NoError(t, se)
				require.Equal(t, j, s)
			} else {
				require.Error(t, se)
				require.Equal(t, S{}, s)
			}
		})
	}
}

var errMarshalSentinel = errors.New("marshal sentinel")

type marshalErrorValue struct{}

func (marshalErrorValue) MarshalJSON() ([]byte, error) {
	return nil, errMarshalSentinel
}

// Case MARSHAL-ERR-WRAP: encoding/json wraps errors returned by MarshalJSON
// with *json.MarshalerError. Sonic returns the original error directly.
func TestCase_MARSHAL_ERR_WRAP(t *testing.T) {
	_, se := sonic.ConfigStd.Marshal(marshalErrorValue{})
	_, je := json.Marshal(marshalErrorValue{})

	require.ErrorIs(t, se, errMarshalSentinel)
	require.Same(t, errMarshalSentinel, se)
	require.ErrorIs(t, je, errMarshalSentinel)
	var marshalerErr *json.MarshalerError
	require.ErrorAs(t, je, &marshalerErr)
	if stdUsesJSONV2 {
		require.Equal(t, reflect.TypeOf((*marshalErrorValue)(nil)), marshalerErr.Type)
	} else {
		require.Equal(t, reflect.TypeOf(marshalErrorValue{}), marshalerErr.Type)
	}
}

// Case STRING-GONUM-PLUS: the jsonv2-backed v1 facade accepts leading plus
// signs and decimal points in quoted numbers. Legacy encoding/json rejects them.
func TestCase_STRING_GONUM_PLUS(t *testing.T) {
	type S struct {
		Int   int     `json:",string"`
		Float float64 `json:",string"`
	}

	tests := []struct {
		name string
		data string
		want S
	}{
		{name: "int-leading-plus", data: `{"Int":"+1"}`, want: S{Int: 1}},
		{name: "float-leading-dot", data: `{"Float":".5"}`, want: S{Float: 0.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s, j S
			se := sonic.ConfigStd.Unmarshal([]byte(tt.data), &s)
			je := json.Unmarshal([]byte(tt.data), &j)

			if envs.UseOptDec {
				require.NoError(t, se)
				require.Equal(t, tt.want, s)
			} else {
				require.Error(t, se)
				require.Equal(t, S{}, s)
			}
			if stdUsesJSONV2 {
				require.NoError(t, je)
				require.Equal(t, tt.want, j)
			} else {
				require.Error(t, je)
				require.Equal(t, S{}, j)
			}
		})
	}
}

type rewrittenTextKey string

func (rewrittenTextKey) MarshalText() ([]byte, error) {
	return []byte("TEXTKEY"), nil
}

// Case MAP-STRING-TEXT: the jsonv2-backed v1 facade calls MarshalText for a
// named string map key. Legacy encoding/json and Sonic use the underlying key.
func TestCase_MAP_STRING_TEXT(t *testing.T) {
	in := map[rewrittenTextKey]int{"raw": 1}
	sout, serr := sonic.ConfigStd.Marshal(in)
	jout, jerr := json.Marshal(in)

	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.JSONEq(t, `{"raw":1}`, string(sout))
	if stdUsesJSONV2 {
		require.JSONEq(t, `{"TEXTKEY":1}`, string(jout))
	} else {
		require.JSONEq(t, string(sout), string(jout))
	}
}

// Case BASE64-ERROR: all implementations reject malformed Base64, but
// GoStd127 wraps the corrupt-input error and Sonic also clears the destination.
func TestCase_BASE64_ERROR(t *testing.T) {
	data := []byte(`"YWJjZA=!"`)
	s := []byte("old")
	j := []byte("old")

	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	require.Error(t, se)
	require.Error(t, je)
	require.Empty(t, s)
	require.Equal(t, []byte("old"), j)
	var corrupt base64.CorruptInputError
	require.Equal(t, !envs.UseOptDec, errors.As(se, &corrupt))
	require.ErrorAs(t, je, &corrupt)
	var typeErr *json.UnmarshalTypeError
	require.Equal(t, stdUsesJSONV2, errors.As(je, &typeErr))
}

// Case UNSUPPORTED-VALUE-PAYLOAD: GoStd127 no longer populates the Value field
// of UnsupportedValueError. Legacy encoding/json and Sonic retain a payload.
func TestCase_UNSUPPORTED_VALUE_PAYLOAD(t *testing.T) {
	_, se := sonic.ConfigStd.Marshal(math.NaN())
	_, je := json.Marshal(math.NaN())

	var sonicErr, jsonErr *json.UnsupportedValueError
	require.ErrorAs(t, se, &sonicErr)
	require.ErrorAs(t, je, &jsonErr)
	require.True(t, sonicErr.Value.IsValid())
	require.Equal(t, !stdUsesJSONV2, jsonErr.Value.IsValid())
}

type nestedMiddle struct {
	N int
}

type nestedOuter struct {
	M nestedMiddle
}

// Case UNMARSHAL-TYPE-ROOT: GoStd127 reports the root destination type in
// UnmarshalTypeError.Struct; legacy encoding/json reports the immediate owner.
func TestCase_UNMARSHAL_TYPE_ROOT(t *testing.T) {
	data := []byte(`{"M":{"N":"bad"}}`)
	var s, j nestedOuter
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	var sonicErr *decoder.MismatchTypeError
	require.ErrorAs(t, se, &sonicErr)
	var typeErr *json.UnmarshalTypeError
	require.ErrorAs(t, je, &typeErr)
	require.Equal(t, "M.N", typeErr.Field)
	if stdUsesJSONV2 {
		require.Equal(t, "nestedOuter", typeErr.Struct)
	} else {
		require.Equal(t, "nestedMiddle", typeErr.Struct)
	}
}

type appendAndMarshalText int

func (appendAndMarshalText) AppendText(dst []byte) ([]byte, error) {
	return append(dst, "APPEND"...), nil
}

func (appendAndMarshalText) MarshalText() ([]byte, error) {
	return []byte("MARSHAL"), nil
}

// Case TEXT-APPEND: Go 1.27's jsonv2-backed v1 facade prefers AppendText over
// MarshalText. Legacy encoding/json and Sonic call MarshalText.
func TestCase_TEXT_APPEND(t *testing.T) {
	sout, serr := sonic.ConfigStd.Marshal(appendAndMarshalText(1))
	jout, jerr := json.Marshal(appendAndMarshalText(1))

	require.NoError(t, serr)
	require.NoError(t, jerr)
	require.Equal(t, `"MARSHAL"`, string(sout))
	if stdUsesJSONV2 {
		require.Equal(t, `"APPEND"`, string(jout))
	} else {
		require.Equal(t, string(sout), string(jout))
	}
}

// Case F64-LONG-OVF: a mantissa longer than the 800-digit fallback buffer
// exposes a historical range-detection bug. Go 1.27 reports the mathematical
// overflow; Sonic still truncates the mantissa and returns a finite value.
func TestCase_F64_LONG_OVF(t *testing.T) {
	number := "5" + strings.Repeat("3", 383) + "5" + strings.Repeat("3", 1327) + "e-913"
	data := []byte("[" + number + "]")

	var s, j []interface{}
	se := sonic.ConfigStd.Unmarshal(data, &s)
	je := json.Unmarshal(data, &j)

	require.NoError(t, se)
	require.Len(t, s, 1)
	require.Equal(t, 5.3333333333333336e-114, s[0])

	_, parseErr := strconv.ParseFloat(number, 64)
	if errors.Is(parseErr, strconv.ErrRange) {
		require.Error(t, je)
		require.Len(t, j, 1)
		if stdUsesJSONV2 {
			require.True(t, math.IsInf(j[0].(float64), 1))
		} else {
			require.Nil(t, j[0])
		}
	} else {
		// Go 1.26 and earlier share Sonic's historical finite-value behavior.
		require.NoError(t, je)
		require.Equal(t, s, j)
	}
}
