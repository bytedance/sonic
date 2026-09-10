//go:build (amd64 && go1.17 && !go1.27) || (arm64 && go1.20 && !go1.27)
// +build amd64,go1.17,!go1.27 arm64,go1.20,!go1.27

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

package sonic

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

// TestNoPCLMULQDQ_JSON runs a subprocess with SONIC_MODE=noavx to simulate
// a CPU without PCLMULQDQ (e.g. Hygon Dhyana under QEMU). The SSE path
// uses scalar prefix-XOR instead of PCLMULQDQ for the inquote bitmask.
func TestNoPCLMULQDQ_JSON(t *testing.T) {
	if os.Getenv("SONIC_TEST_NOPCLMUL") == "1" {
		// Subprocess: run JSON operations using SSE path (no PCLMULQDQ)
		runJSONTests(t)
		return
	}

	// Parent: launch subprocess with SONIC_MODE=noavx
	cmd := exec.Command(os.Args[0], "-test.run=^TestNoPCLMULQDQ_JSON$", "-test.v")
	cmd.Env = append(os.Environ(), "SONIC_MODE=noavx", "SONIC_TEST_NOPCLMUL=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("subprocess failed: %v\n%s", err, out)
	}
	t.Logf("subprocess output:\n%s", out)
}

func runJSONTests(t *testing.T) {
	cfg := ConfigDefault

	// Test 1: basic marshal/unmarshal
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		type Obj struct {
			Name    string   `json:"name"`
			Age     int      `json:"age"`
			Tags    []string `json:"tags"`
			Escaped string   `json:"escaped"`
		}
		orig := Obj{
			Name:    "hello",
			Age:     42,
			Tags:    []string{"a", "b", "c"},
			Escaped: `quote"slash\tab	newline` + "\n",
		}
		data, err := cfg.Marshal(orig)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		var got Obj
		if err := cfg.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if !reflect.DeepEqual(got, orig) {
			t.Fatalf("roundtrip mismatch: got %+v, want %+v", got, orig)
		}
	})

	// Test 2: strings with many quotes (exercises the inquote bitmask heavily)
	t.Run("ManyQuotes", func(t *testing.T) {
		input := `{"a":"he\"ll\"o","b":"wo\"r\"ld","c":"\"\"\"","d":"no quotes"}`
		var m map[string]string
		if err := cfg.UnmarshalFromString(input, &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m["a"] != `he"ll"o` {
			t.Fatalf("a = %q, want %q", m["a"], `he"ll"o`)
		}
		if m["b"] != `wo"r"ld` {
			t.Fatalf("b = %q, want %q", m["b"], `wo"r"ld`)
		}
		if m["c"] != `"""` {
			t.Fatalf("c = %q, want %q", m["c"], `"""`)
		}
		if m["d"] != "no quotes" {
			t.Fatalf("d = %q, want %q", m["d"], "no quotes")
		}
	})

	// Test 3: deeply nested JSON (exercises skip_one, skip_array, skip_object)
	t.Run("DeepNested", func(t *testing.T) {
		input := `{"a":{"b":{"c":{"d":[1,2,{"e":"val"}]}}}}`
		var m map[string]interface{}
		if err := cfg.UnmarshalFromString(input, &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		// Traverse to the leaf
		a := m["a"].(map[string]interface{})
		b := a["b"].(map[string]interface{})
		c := b["c"].(map[string]interface{})
		d := c["d"].([]interface{})
		inner := d[2].(map[string]interface{})
		if inner["e"] != "val" {
			t.Fatalf("nested value = %v, want %q", inner["e"], "val")
		}
	})

	// Test 4: long string > 64 bytes (exercises the full 64-byte scanning loop)
	t.Run("LongString", func(t *testing.T) {
		// 200-char string with embedded quotes and backslashes
		longVal := `abcdefghij\"klmnopqrst\\uvwxyz0123456789ABCDEFGHIJ\"KLMNOPQRST\\UVWXYZ` +
			`abcdefghij\"klmnopqrst\\uvwxyz0123456789ABCDEFGHIJ\"KLMNOPQRST\\UVWXYZ` +
			`final\"end`
		input := `{"key":"` + longVal + `"}`
		var m map[string]string
		if err := cfg.UnmarshalFromString(input, &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if len(m["key"]) == 0 {
			t.Fatal("expected non-empty string for long value")
		}
	})

	// Test 5: GetByPath (this was the exact crash point in #867)
	t.Run("GetByPath", func(t *testing.T) {
		input := `{"users":[{"name":"alice","age":30},{"name":"bob","age":25}]}`
		node, err := GetFromString(input, "users", 1, "name")
		if err != nil {
			t.Fatalf("GetFromString: %v", err)
		}
		val, err := node.String()
		if err != nil {
			t.Fatalf("String: %v", err)
		}
		if val != "bob" {
			t.Fatalf("got %q, want %q", val, "bob")
		}
	})

	// Test 6: Valid (exercises validate_one)
	t.Run("Valid", func(t *testing.T) {
		if !cfg.Valid([]byte(`{"a":1,"b":[true,false,null]}`)) {
			t.Fatal("expected valid JSON")
		}
		if cfg.Valid([]byte(`{"a":1,"b":[true,false,null}`)) {
			t.Fatal("expected invalid JSON")
		}
	})
}
