/**
 * Copyright 2024 ByteDance Inc.
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

package prim

import (
	"bytes"
	"encoding/json"
	"testing"
)

var testCases = []struct {
	name  string
	input string
	want  string
}{
	{"empty_object", `{ }`, `{}`},
	{"empty_array", `[ ]`, `[]`},
	{"simple_object", `{ "a" : 1 }`, `{"a":1}`},
	{"nested_object", `{ "a" : { "b" : 2 } }`, `{"a":{"b":2}}`},
	{"array", `[ 1 , 2 , 3 ]`, `[1,2,3]`},
	{"string_with_spaces", `{ "msg" : "hello world" }`, `{"msg":"hello world"}`},
	{"string_with_escapes", `{ "msg" : "hello\tworld\n" }`, `{"msg":"hello\tworld\n"}`},
	{"string_with_quote", `{ "msg" : "say \"hello\"" }`, `{"msg":"say \"hello\""}`},
	{"multiline", "{\n\t\"a\": 1,\n\t\"b\": 2\n}", `{"a":1,"b":2}`},
	{"mixed", `{ "arr" : [ 1 , { "x" : "y" } ] }`, `{"arr":[1,{"x":"y"}]}`},
	{"unicode", `{ "emoji" : "😀" }`, `{"emoji":"😀"}`},
	{"escaped_backslash", `{ "path" : "c:\\dir\\file" }`, `{"path":"c:\\dir\\file"}`},
}

func TestCompact(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var dst []byte
			err := Compact(&dst, []byte(tc.input))
			if err != nil {
				t.Fatalf("Compact error: %v", err)
			}
			if string(dst) != tc.want {
				t.Errorf("got %q, want %q", string(dst), tc.want)
			}

			// Verify against standard library
			var stdDst bytes.Buffer
			if err := json.Compact(&stdDst, []byte(tc.input)); err != nil {
				t.Fatalf("json.Compact error: %v", err)
			}
			if string(dst) != stdDst.String() {
				t.Errorf("mismatch with std: got %q, std %q", string(dst), stdDst.String())
			}
		})
	}
}

func TestCompactEmpty(t *testing.T) {
	var dst []byte
	err := Compact(&dst, []byte{})
	if err != nil {
		t.Fatalf("Compact error: %v", err)
	}
	if len(dst) != 0 {
		t.Errorf("expected empty, got %q", string(dst))
	}
}

func TestCompactAppend(t *testing.T) {
	// Test that Compact properly appends to existing slice
	dst := []byte("prefix:")
	err := Compact(&dst, []byte(`{ "a" : 1 }`))
	if err != nil {
		t.Fatalf("Compact error: %v", err)
	}
	want := `prefix:{"a":1}`
	if string(dst) != want {
		t.Errorf("got %q, want %q", string(dst), want)
	}
}

func TestCompactInvalidJSON(t *testing.T) {
	invalidCases := []struct {
		name  string
		input string
	}{
		{"unclosed_brace", `{`},
		{"unclosed_bracket", `[`},
		{"trailing_comma_array", `[1,2,]`},
		{"trailing_comma_object", `{"a":1,}`},
		{"wrong_order", `}{`},
		{"missing_colon", `{"a" 1}`},
		{"missing_value", `{"a":}`},
		{"double_comma", `[1,,2]`},
	}

	for _, tc := range invalidCases {
		t.Run(tc.name, func(t *testing.T) {
			var dst []byte
			err := Compact(&dst, []byte(tc.input))
			if err == nil {
				t.Errorf("expected error for invalid JSON %q, got nil", tc.input)
			}

			// Verify behavior matches standard library
			var stdDst bytes.Buffer
			stdErr := json.Compact(&stdDst, []byte(tc.input))
			if stdErr == nil {
				t.Fatalf("std json.Compact should also error for %q", tc.input)
			}
		})
	}
}

// Large JSON for benchmarking
var largeJSON = func() []byte {
	obj := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		obj[string(rune('a'+i%26))+string(rune('0'+i/26))] = map[string]interface{}{
			"value": i,
			"name":  "test string with some content",
			"nested": map[string]interface{}{
				"x": 1,
				"y": 2,
			},
		}
	}
	data, _ := json.MarshalIndent(obj, "", "    ")
	return data
}()

func BenchmarkCompact_Sonic(b *testing.B) {
	src := largeJSON
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst []byte
		_ = Compact(&dst, src)
	}
}

func BenchmarkCompact_Std(b *testing.B) {
	src := largeJSON
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst bytes.Buffer
		_ = json.Compact(&dst, src)
	}
}

func BenchmarkCompact_Sonic_Small(b *testing.B) {
	src := []byte(`{ "name" : "test" , "value" : 123 }`)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst []byte
		_ = Compact(&dst, src)
	}
}

func BenchmarkCompact_Std_Small(b *testing.B) {
	src := []byte(`{ "name" : "test" , "value" : 123 }`)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst bytes.Buffer
		_ = json.Compact(&dst, src)
	}
}
