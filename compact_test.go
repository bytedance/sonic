/*
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

package sonic

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCompactAPI(t *testing.T) {
	input := []byte(`{  "name" : "sonic"  ,  "fast" : true  }`)
	want := `{"name":"sonic","fast":true}`

	var dst []byte
	err := Compact(&dst, input)
	if err != nil {
		t.Fatalf("Compact error: %v", err)
	}
	if string(dst) != want {
		t.Errorf("got %q, want %q", string(dst), want)
	}

	// Verify against standard library
	var stdDst bytes.Buffer
	if err := json.Compact(&stdDst, input); err != nil {
		t.Fatalf("json.Compact error: %v", err)
	}
	if string(dst) != stdDst.String() {
		t.Errorf("mismatch with std: got %q, std %q", string(dst), stdDst.String())
	}
}

func BenchmarkCompactAPI_Sonic(b *testing.B) {
	src := []byte(`{  "name" : "sonic"  ,  "version" : 1.0  ,  "features" : [ "fast" , "safe" ]  }`)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst []byte
		_ = Compact(&dst, src)
	}
}

func BenchmarkCompactAPI_Std(b *testing.B) {
	src := []byte(`{  "name" : "sonic"  ,  "version" : 1.0  ,  "features" : [ "fast" , "safe" ]  }`)
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst bytes.Buffer
		_ = json.Compact(&dst, src)
	}
}
