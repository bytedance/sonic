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
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
)

// singleK implements encoding.TextUnmarshaler for map key (pointer receiver).
// Used to reproduce issue #923: ConfigStd.Unmarshal into map[*K]V can segfault when SONIC_USE_OPTDEC=0.
type singleK923 struct{ S string }

func (k *singleK923) UnmarshalText(b []byte) error {
	k.S = string(b)
	return nil
}

func runInSubprocess(t *testing.T, testName, helperEnv string) {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^"+testName+"$")
	cmd.Env = append(os.Environ(), helperEnv+"=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("subprocess failed: %v\n%s", err, string(out))
	}
}

// TestIssue923_MapPointerKeyTextUnmarshalerNoSegfault verifies that unmarshaling
// into map[*K]V where K implements encoding.TextUnmarshaler does not segfault
// when using jitdec (SONIC_USE_OPTDEC=0). See https://github.com/bytedance/sonic/issues/923
func TestIssue923_MapPointerKeyTextUnmarshalerNoSegfault(t *testing.T) {
	if os.Getenv("SONIC_923_HELPER") != "1" {
		// Keep this one isolated: historically it can crash the process.
		runInSubprocess(t, "TestIssue923_MapPointerKeyTextUnmarshalerNoSegfault", "SONIC_923_HELPER")
		return
	}

	vals := make([]string, 256)
	for i := range vals {
		vals[i] = strconv.Itoa(i % 10)
	}
	jsonText := "{\"alpha\":[" + strings.Join(vals, ",") + "]}"

	var m map[*singleK923][256]byte
	if err := sonic.ConfigStd.Unmarshal([]byte(jsonText), &m); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if len(m) != 1 {
		t.Fatalf("unexpected map size: %d", len(m))
	}
	for k, v := range m {
		if k == nil {
			t.Fatalf("decoded nil key")
		}
		if k.S != "alpha" || v[0] != 0 || v[255] != 5 {
			t.Fatalf("unexpected decode result: key=%q first=%d last=%d", k.S, v[0], v[255])
		}
	}
}

type largeElem923 struct {
	N   int    `json:"n"`
	S   string `json:"s"`
	Pad [256]byte
}

func TestIssue923_MapIndirectElemStringKeyNoSegfault(t *testing.T) {
	var m map[string]largeElem923
	if err := sonic.ConfigStd.Unmarshal([]byte(`{"k":{"n":7,"s":"ok"}}`), &m); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	v, ok := m["k"]
	if !ok || v.N != 7 || v.S != "ok" {
		t.Fatalf("unexpected decode result: ok=%v n=%d s=%q len=%d", ok, v.N, v.S, len(m))
	}
}

func TestIssue923_MapIndirectElemUint64KeyNoSegfault(t *testing.T) {
	var m map[uint64]largeElem923
	if err := sonic.ConfigStd.Unmarshal([]byte(`{"9":{"n":9,"s":"u64"}}`), &m); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	v, ok := m[9]
	if !ok || v.N != 9 || v.S != "u64" {
		t.Fatalf("unexpected decode result: ok=%v n=%d s=%q len=%d", ok, v.N, v.S, len(m))
	}
}
