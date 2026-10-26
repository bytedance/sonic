// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/internal/rt"
)

func collectIssue999Numbers(v reflect.Value, out *[]json.Number) {
	if !v.IsValid() {
		return
	}
	if v.Type() == reflect.TypeOf(json.Number("")) {
		*out = append(*out, v.Interface().(json.Number))
		return
	}
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		if !v.IsNil() {
			collectIssue999Numbers(v.Elem(), out)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			collectIssue999Numbers(k, out)
			collectIssue999Numbers(v.MapIndex(k), out)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			collectIssue999Numbers(v.Index(i), out)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			collectIssue999Numbers(v.Field(i), out)
		}
	}
}

func TestIssue999CopyNumberInputReuse(t *testing.T) {
	for _, number := range []string{"7", "12345", "-12345.625e+10", strings.Repeat("1", 128)} {
		for _, tc := range []struct {
			name, format string
			newValue     func() interface{}
		}{
			{"interface", `%s`, func() interface{} { return new(interface{}) }},
			{"generic-map", `{"n":%s}`, func() interface{} { return new(map[string]interface{}) }},
			{"generic-slice", `[%s]`, func() interface{} { return new([]interface{}) }},
			{"nested", `{"a":[{"n":%s}]}`, func() interface{} { return new(interface{}) }},
			{"struct-interface", `{"N":%s}`, func() interface{} { return new(struct{ N interface{} }) }},
			{"typed-number", `%s`, func() interface{} { return new(json.Number) }},
			{"typed-map", `{"n":%s}`, func() interface{} { return new(map[string]json.Number) }},
			{"typed-slice", `[%s]`, func() interface{} { return new([]json.Number) }},
			{"typed-pointer", `%s`, func() interface{} { return new(*json.Number) }},
			{"quoted-number", `"%s"`, func() interface{} { return new(json.Number) }},
			{"string-tag", `{"N":"%s"}`, func() interface{} {
				return new(struct {
					N json.Number `json:",string"`
				})
			}},
			{"number-map-key", `{"%s":0}`, func() interface{} { return new(map[json.Number]int) }},
		} {
			t.Run(tc.name+"/"+number, func(t *testing.T) {
				input := []byte(fmt.Sprintf(tc.format, number))
				value := tc.newValue()
				api := sonic.Config{CopyString: true, UseNumber: true}.Froze()
				if err := api.UnmarshalFromString(rt.Mem2Str(input), value); err != nil {
					t.Fatal(err)
				}
				var got []json.Number
				collectIssue999Numbers(reflect.ValueOf(value), &got)
				if len(got) != 1 || string(got[0]) != number {
					t.Fatalf("unexpected decoded numbers: %v", got)
				}
				// Reuse the same allocation after decoding. A copied number must
				// retain its value even when every byte of the source is replaced.
				for i := range input {
					input[i] = '9'
				}
				if string(got[0]) != number {
					t.Fatalf("decoded number changed after input reuse: got %q, want %q", got[0], number)
				}
				if m, ok := value.(*map[json.Number]int); ok {
					if _, found := (*m)[json.Number(number)]; !found {
						t.Fatal("original number key is no longer searchable after input reuse")
					}
				}
			})
		}
	}
}

// Both configurations reuse the cached decoder for the same map type. Its copy
// policy must be taken from the current call, not cached in the shared decoder.
func TestIssue999NumberKeyCopyOption(t *testing.T) {
	if sonic.APIKind == sonic.UseStdJSON {
		t.Skip("encoding/json fallback always copies input")
	}
	for _, copyString := range []bool{true, false, true, false} {
		t.Run(fmt.Sprint(copyString), func(t *testing.T) {
			input := []byte(`{"12345":1}`)
			dc := decoder.NewDecoder(rt.Mem2Str(input))
			dc.UseNumber()
			if copyString {
				dc.CopyString()
			}
			var value map[json.Number]int
			if err := dc.Decode(&value); err != nil {
				t.Fatal(err)
			}
			var key json.Number
			for k := range value {
				key = k
			}
			input[2] = '9'
			want := json.Number("92345")
			if copyString {
				want = "12345"
			}
			if key != want {
				t.Fatalf("got %q, want %q", key, want)
			}
		})
	}
}

func TestIssue999StreamNumberRetention(t *testing.T) {
	api := sonic.Config{CopyString: true, UseNumber: true}.Froze()
	dc := api.NewDecoder(strings.NewReader(`{"n":12345} {"n":67890}`))
	var first, second map[string]interface{}
	if err := dc.Decode(&first); err != nil {
		t.Fatal(err)
	}
	if err := dc.Decode(&second); err != nil {
		t.Fatal(err)
	}
	if first["n"] != json.Number("12345") || second["n"] != json.Number("67890") {
		t.Fatalf("stream results changed: first=%v second=%v", first, second)
	}
}
