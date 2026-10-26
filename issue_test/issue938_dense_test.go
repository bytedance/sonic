// Copyright 2026 CloudWeGo Authors
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
	stdjson "encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
)

type issue938DenseUnknownTarget struct {
	Known string `json:"known"`
}

type issue938DenseKnownTarget struct {
	Payload string `json:"payload"`
	Known   string `json:"known"`
}

type issue938DenseWorkload struct {
	name       string
	encoded    string
	decoded    string
	encodedLen int
}

type preparedIssue938DenseWorkload struct {
	issue938DenseWorkload
	payload []byte
}

var (
	issue938DenseUnknownSink issue938DenseUnknownTarget
	issue938DenseKnownSink   issue938DenseKnownTarget
)

func TestIssue938_DenseEscapeValidParity(t *testing.T) {
	simpleEncoded, simpleDecoded := issue938DenseSimpleEscapes()
	fixtures := []struct {
		name    string
		encoded string
		decoded string
	}{
		{
			name:    "AllEightSimpleEscapes",
			encoded: strings.Repeat(simpleEncoded, 128),
			decoded: strings.Repeat(simpleDecoded, 128),
		},
		{
			name:    "MixedUnicode",
			encoded: strings.Repeat(`\n\u0061`, 128),
			decoded: strings.Repeat("\na", 128),
		},
		{
			name:    "SurrogatePairBetweenDenseBlocks",
			encoded: strings.Repeat(`\n`, 40) + `\uD800\uDC00` + strings.Repeat(`\t`, 40),
			decoded: strings.Repeat("\n", 40) + "\U00010000" + strings.Repeat("\t", 40),
		},
	}

	for _, fixture := range fixtures {
		fixture := fixture
		t.Run(fixture.name, func(t *testing.T) {
			assertIssue938DenseEscapeValidParity(t, fixture.encoded, fixture.decoded)
		})
	}
}

func TestIssue938_DenseEscapeBoundaryParity(t *testing.T) {
	simpleEncoded, simpleDecoded := issue938DenseSimpleEscapes()
	for pad := 0; pad <= 65; pad++ {
		pad := pad
		prefix := strings.Repeat("a", pad)
		t.Run(fmt.Sprintf("Pad%02d/AllEightSimpleEscapes", pad), func(t *testing.T) {
			assertIssue938DenseEscapeValidParity(t, prefix+simpleEncoded, prefix+simpleDecoded)
		})
		t.Run(fmt.Sprintf("Pad%02d/BMPUnicode", pad), func(t *testing.T) {
			assertIssue938DenseEscapeValidParity(t, prefix+`\u0061`, prefix+"a")
		})
		t.Run(fmt.Sprintf("Pad%02d/SurrogatePair", pad), func(t *testing.T) {
			encoded := strings.Repeat(`\n`, 40) + prefix + `\uD800\uDC00` + strings.Repeat(`\t`, 40)
			decoded := strings.Repeat("\n", 40) + prefix + "\U00010000" + strings.Repeat("\t", 40)
			assertIssue938DenseEscapeValidParity(t, encoded, decoded)
		})
	}
}

func TestIssue938_DenseEscapeBackslashThenLiteralXParity(t *testing.T) {
	for _, pad := range []int{31, 32, 33, 63, 64, 65} {
		pad := pad
		t.Run(fmt.Sprintf("Pad%02d", pad), func(t *testing.T) {
			prefix := strings.Repeat("a", pad)
			assertIssue938DenseEscapeValidParity(t, prefix+`\\x`, prefix+`\x`)
		})
	}
}

func TestIssue938_DenseEscapeMalformedParity(t *testing.T) {
	assertIssue938DenseEscapeInvalidParity(t,
		strings.Repeat(`\n`, 40)+string([]byte{0x01})+`\q`)
	assertIssue938DenseEscapeInvalidParity(t, strings.Repeat(`\n`, 512)+`\q`)
	assertIssue938DenseEscapeInvalidParity(t, strings.Repeat(`\n`, 40)+`\u12G4`)

	for _, pad := range []int{28, 29, 30, 31, 32, 33, 34, 35, 36, 60, 61, 62, 63, 64, 65} {
		pad := pad
		t.Run(fmt.Sprintf("MalformedUnicode/Pad%02d", pad), func(t *testing.T) {
			assertIssue938DenseEscapeInvalidParity(t, strings.Repeat("a", pad)+`\u12G4`)
		})
	}
}

func BenchmarkIssue938_DenseEscape(b *testing.B) {
	b.StopTimer()
	workloads := prepareIssue938DenseWorkloads(b)
	b.Run("ConfigStd", func(b *testing.B) {
		for _, workload := range workloads {
			workload := workload
			b.Run("UnknownField/"+workload.name, func(b *testing.B) {
				benchmarkIssue938DenseUnknownField(b, workload)
			})
			b.Run("KnownFieldControl/"+workload.name, func(b *testing.B) {
				benchmarkIssue938DenseKnownField(b, workload)
			})
		}
	})
}

func assertIssue938DenseEscapeValidParity(t *testing.T, encoded, wantPayload string) {
	t.Helper()
	payload := issue938DensePayload(encoded)

	var standardKnown issue938DenseKnownTarget
	if err := stdjson.Unmarshal(payload, &standardKnown); err != nil {
		t.Fatalf("encoding/json rejected valid fixture: %v", err)
	}
	wantKnown := issue938DenseKnownTarget{Payload: wantPayload, Known: "ok"}
	if standardKnown != wantKnown {
		t.Fatalf("encoding/json decoded known fields (%q, %q), want (%q, %q)",
			standardKnown.Payload, standardKnown.Known, wantKnown.Payload, wantKnown.Known)
	}
	var sonicKnown issue938DenseKnownTarget
	if err := sonic.ConfigStd.Unmarshal(payload, &sonicKnown); err != nil {
		t.Fatalf("Sonic ConfigStd rejected valid fixture: %v", err)
	}
	if sonicKnown != standardKnown {
		t.Fatalf("Sonic decoded known fields (%q, %q), encoding/json decoded (%q, %q)",
			sonicKnown.Payload, sonicKnown.Known, standardKnown.Payload, standardKnown.Known)
	}

	var standardUnknown issue938DenseUnknownTarget
	if err := stdjson.Unmarshal(payload, &standardUnknown); err != nil {
		t.Fatalf("encoding/json rejected unknown-field fixture: %v", err)
	}
	var sonicUnknown issue938DenseUnknownTarget
	if err := sonic.ConfigStd.Unmarshal(payload, &sonicUnknown); err != nil {
		t.Fatalf("Sonic ConfigStd rejected unknown-field fixture: %v", err)
	}
	if sonicUnknown != standardUnknown || sonicUnknown.Known != "ok" {
		t.Fatalf("Sonic unknown-field result = %+v, encoding/json = %+v", sonicUnknown, standardUnknown)
	}
}

func assertIssue938DenseEscapeInvalidParity(t *testing.T, encoded string) {
	t.Helper()
	payload := issue938DensePayload(encoded)

	var standardKnown issue938DenseKnownTarget
	if err := stdjson.Unmarshal(payload, &standardKnown); err == nil {
		t.Fatal("encoding/json accepted malformed known-field fixture")
	}
	var sonicKnown issue938DenseKnownTarget
	if err := sonic.ConfigStd.Unmarshal(payload, &sonicKnown); err == nil {
		t.Fatal("Sonic ConfigStd accepted malformed known-field fixture")
	}
	var standardUnknown issue938DenseUnknownTarget
	if err := stdjson.Unmarshal(payload, &standardUnknown); err == nil {
		t.Fatal("encoding/json accepted malformed unknown-field fixture")
	}
	var sonicUnknown issue938DenseUnknownTarget
	if err := sonic.ConfigStd.Unmarshal(payload, &sonicUnknown); err == nil {
		t.Fatal("Sonic ConfigStd accepted malformed unknown-field fixture")
	}
}

func issue938DenseSimpleEscapes() (encoded, decoded string) {
	encodedParts := []string{`\"`, `\\`, `\/`, `\b`, `\f`, `\n`, `\r`, `\t`}
	decodedParts := []string{`"`, `\`, "/", "\b", "\f", "\n", "\r", "\t"}
	return strings.Join(encodedParts, ""), strings.Join(decodedParts, "")
}

func issue938DensePayload(encoded string) []byte {
	return []byte(`{"payload":"` + encoded + `","known":"ok"}`)
}

func prepareIssue938DenseWorkloads(b *testing.B) []preparedIssue938DenseWorkload {
	workloads := issue938DenseWorkloads()
	prepared := make([]preparedIssue938DenseWorkload, 0, len(workloads))
	for _, workload := range workloads {
		if got := len(workload.encoded); got != workload.encodedLen {
			b.Fatalf("fixture %s encoded length = %d, want %d", workload.name, got, workload.encodedLen)
		}
		payload := issue938DensePayload(workload.encoded)
		if !stdjson.Valid(payload) {
			b.Fatalf("encoding/json rejected generated fixture %s", workload.name)
		}

		var standardKnown issue938DenseKnownTarget
		if err := stdjson.Unmarshal(payload, &standardKnown); err != nil {
			b.Fatalf("encoding/json Unmarshal known fixture %s: %v", workload.name, err)
		}
		wantKnown := issue938DenseKnownTarget{Payload: workload.decoded, Known: "ok"}
		if standardKnown != wantKnown {
			b.Fatalf("encoding/json decoded known fixture %s as (%q, %q), want (%q, %q)",
				workload.name, standardKnown.Payload, standardKnown.Known, wantKnown.Payload, wantKnown.Known)
		}
		var standardUnknown issue938DenseUnknownTarget
		if err := stdjson.Unmarshal(payload, &standardUnknown); err != nil {
			b.Fatalf("encoding/json Unmarshal unknown fixture %s: %v", workload.name, err)
		}
		if standardUnknown != (issue938DenseUnknownTarget{Known: "ok"}) {
			b.Fatalf("encoding/json decoded unknown fixture %s as %+v", workload.name, standardUnknown)
		}

		prepared = append(prepared, preparedIssue938DenseWorkload{
			issue938DenseWorkload: workload,
			payload:               payload,
		})
	}
	return prepared
}

func issue938DenseWorkloads() []issue938DenseWorkload {
	plain := func(n int) string { return strings.Repeat("a", n) }
	workloads := make([]issue938DenseWorkload, 0, 20)
	for _, n := range []int{16, 32, 64, 256, 4096} {
		text := plain(n)
		workloads = append(workloads, issue938DenseWorkload{
			name: fmt.Sprintf("Plain%d", n), encoded: text, decoded: text, encodedLen: n,
		})
	}
	for _, n := range []int{32, 64, 256, 4096} {
		workloads = append(workloads, issue938DenseWorkload{
			name:    fmt.Sprintf("SparseEscapeStart%d", n),
			encoded: `\"` + plain(n-2), decoded: `"` + plain(n-2), encodedLen: n,
		})
	}

	simpleEncoded, simpleDecoded := issue938DenseSimpleEscapes()
	workloads = append(workloads,
		issue938DenseWorkload{
			name: "DenseNewline4096", encoded: strings.Repeat(`\n`, 2048),
			decoded: strings.Repeat("\n", 2048), encodedLen: 4096,
		},
		issue938DenseWorkload{
			name: "DenseQuote4096", encoded: strings.Repeat(`\"`, 2048),
			decoded: strings.Repeat(`"`, 2048), encodedLen: 4096,
		},
		issue938DenseWorkload{
			name: "DenseBackslash4096", encoded: strings.Repeat(`\\`, 2048),
			decoded: strings.Repeat(`\`, 2048), encodedLen: 4096,
		},
		issue938DenseWorkload{
			name: "DenseMixedSimple4096", encoded: strings.Repeat(simpleEncoded, 256),
			decoded: strings.Repeat(simpleDecoded, 256), encodedLen: 4096,
		},
		issue938DenseWorkload{
			name: "DenseUnicode4092", encoded: strings.Repeat(`\u0061`, 682),
			decoded: strings.Repeat("a", 682), encodedLen: 4092,
		},
		issue938DenseWorkload{
			name: "DenseSurrogate4092", encoded: strings.Repeat(`\uD800\uDC00`, 341),
			decoded: strings.Repeat("\U00010000", 341), encodedLen: 4092,
		},
		issue938DenseWorkload{
			name: "MixedUnicode4096", encoded: strings.Repeat(`\n\u0061`, 512),
			decoded: strings.Repeat("\na", 512), encodedLen: 4096,
		},
	)
	for _, n := range []int{16, 32, 64, 128} {
		repeats := n / len(simpleEncoded)
		workloads = append(workloads, issue938DenseWorkload{
			name:    fmt.Sprintf("ShortMixed%d", n),
			encoded: strings.Repeat(simpleEncoded, repeats),
			decoded: strings.Repeat(simpleDecoded, repeats), encodedLen: n,
		})
	}
	return workloads
}

func benchmarkIssue938DenseUnknownField(b *testing.B, workload preparedIssue938DenseWorkload) {
	var target issue938DenseUnknownTarget
	b.SetBytes(int64(len(workload.payload)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = issue938DenseUnknownTarget{}
		if err := sonic.ConfigStd.Unmarshal(workload.payload, &target); err != nil {
			b.Fatalf("Unmarshal fixture %s: %v", workload.name, err)
		}
	}
	b.StopTimer()
	want := issue938DenseUnknownTarget{Known: "ok"}
	if target != want {
		b.Fatalf("fixture %s decoded unknown-field target as %+v, want %+v", workload.name, target, want)
	}
	issue938DenseUnknownSink = target
}

func benchmarkIssue938DenseKnownField(b *testing.B, workload preparedIssue938DenseWorkload) {
	var target issue938DenseKnownTarget
	b.SetBytes(int64(len(workload.payload)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = issue938DenseKnownTarget{}
		if err := sonic.ConfigStd.Unmarshal(workload.payload, &target); err != nil {
			b.Fatalf("Unmarshal fixture %s: %v", workload.name, err)
		}
	}
	b.StopTimer()
	want := issue938DenseKnownTarget{Payload: workload.decoded, Known: "ok"}
	if target != want {
		b.Fatalf("fixture %s decoded known-field target as (%q, %q), want (%q, %q)",
			workload.name, target.Payload, target.Known, want.Payload, want.Known)
	}
	issue938DenseKnownSink = target
}
