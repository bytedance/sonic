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

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// streamFuzzReader varies read boundaries independently of the JSON bytes.
// A finite read budget turns accidental extra reads into a reproducible failure.
type streamFuzzReader struct {
	input          string
	plan           []byte
	pos, reads     int
	dataReads      int
	emptyEvery     int
	emptyDelivered bool
	errWithData    bool
	tail           error
}

func (r *streamFuzzReader) Read(p []byte) (int, error) {
	if r.reads > 4*len(r.input)+64 {
		panic("stream decoder exceeded its read budget")
	}
	r.reads++
	if r.emptyEvery != 0 && !r.emptyDelivered && r.dataReads%r.emptyEvery == 0 {
		r.emptyDelivered = true
		return 0, nil
	}
	r.emptyDelivered = false
	if r.pos == len(r.input) {
		return 0, r.tail
	}
	n := len(r.input) - r.pos
	if len(r.plan) != 0 {
		limit := 1 + int(r.plan[(r.reads-1)%len(r.plan)])%64
		if n > limit {
			n = limit
		}
	}
	if n > len(p) {
		n = len(p)
	}
	copy(p, r.input[r.pos:r.pos+n])
	r.pos += n
	r.dataReads++
	if r.pos == len(r.input) && r.errWithData {
		return n, r.tail
	}
	return n, nil
}

func assertRawMessagesEqual(t *testing.T, got, want []json.RawMessage) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("decoded %d values, want %d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Fatalf("value %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func decodeFuzzRawStream(t *testing.T, input string, plan []byte, ending, emptyEvery uint8) ([]json.RawMessage, *StreamDecoder, *streamFuzzReader) {
	t.Helper()
	r := &streamFuzzReader{
		input:       input,
		plan:        plan,
		tail:        io.EOF,
		errWithData: ending&1 != 0,
		emptyEvery:  int(emptyEvery % 8),
	}
	d := NewStreamDecoder(r)
	var values []json.RawMessage
	for len(values) <= 32 {
		var value json.RawMessage
		err := d.Decode(&value)
		if err == io.EOF {
			return values, d, r
		}
		if err != nil {
			t.Fatalf("Decode value %d: %v; input=%q plan=%x ending=%d emptyEvery=%d", len(values), err, input, plan, ending, emptyEvery)
		}
		values = append(values, value)
	}
	t.Fatal("too many decoded stream values")
	return nil, nil, nil
}

func streamFuzzValues(input string) ([]json.RawMessage, error) {
	d := json.NewDecoder(strings.NewReader(input))
	var values []json.RawMessage
	for len(values) <= 32 {
		var value json.RawMessage
		err := d.Decode(&value)
		if err == io.EOF {
			return values, nil
		}
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return nil, errors.New("too many stream values")
}

func streamFuzzSonicValues(input string) ([]json.RawMessage, error) {
	d := NewStreamDecoder(strings.NewReader(input))
	var values []json.RawMessage
	for len(values) <= 32 {
		var value json.RawMessage
		err := d.Decode(&value)
		if err == io.EOF {
			return values, nil
		}
		if err != nil {
			return nil, err
		}
		values = append(values, append(json.RawMessage(nil), value...))
	}
	return nil, errors.New("too many stream values")
}

// This exercises the public lifecycle rather than just the numeric scanner:
// Reader -> Decode -> Buffered/InputOffset -> another Decode -> terminal EOF.
// RawMessage compares framing without depending on unrelated conversion rules.
func FuzzStreamDecoderChunking(f *testing.F) {
	// These all pass without fragmentation on the pre-fix implementation.
	// In particular, "42" only exposes truncation when fuzzing mutates plan.
	for _, input := range []string{"42", "0", "-1.5", "1e+3", "1e700", "true", `"stream"`, `{"n":42}`, "1 2"} {
		f.Add(input, []byte{}, uint8(0))
	}
	f.Fuzz(func(t *testing.T, input string, plan []byte, ending uint8) {
		if len(input) > 2048 || len(plan) > 32 {
			t.Skip()
		}
		standardValues, err := streamFuzzValues(input)
		if err != nil || len(standardValues) == 0 {
			t.Skip()
		}
		// This fuzzer checks read-boundary invariance. Use Sonic with a whole
		// Reader as the semantic oracle so architecture-specific number behavior
		// (for example, overflow handling) is not mistaken for framing failure.
		want, err := streamFuzzSonicValues(input)
		if err != nil || len(want) != len(standardValues) {
			t.Skip()
		}
		r := &streamFuzzReader{input: input, plan: plan, tail: io.EOF, errWithData: ending&1 != 0}
		d := NewStreamDecoder(r)
		var offset int64
		for i, expected := range want {
			var got json.RawMessage
			if err := d.Decode(&got); err != nil {
				t.Fatalf("value %d: Decode: %v; input=%q plan=%x ending=%d", i, err, input, plan, ending)
			}
			if !bytes.Equal(got, expected) {
				t.Fatalf("value %d: got %q, want %q; input=%q plan=%x ending=%d", i, got, expected, input, plan, ending)
			}
			if next := d.InputOffset(); next < offset || next > int64(len(input)) {
				t.Fatalf("invalid offset: previous=%d current=%d input=%q", offset, next, input)
			} else {
				offset = next
			}
			buffered, err := io.ReadAll(d.Buffered())
			if err != nil {
				t.Fatal(err)
			}
			// Ignore insignificant inter-value whitespace, which the decoder may
			// already have skipped, but ensure no future value was consumed.
			remaining := string(buffered) + input[r.pos:]
			replayed, err := streamFuzzValues(remaining)
			if err != nil || len(replayed) != len(want)-i-1 {
				t.Fatalf("lost buffered values after %d: remaining=%q err=%v got=%d want=%d", i, remaining, err, len(replayed), len(want)-i-1)
			}
			for j := range replayed {
				if !bytes.Equal(replayed[j], want[i+j+1]) {
					t.Fatalf("buffered value %d: got %q, want %q", j, replayed[j], want[i+j+1])
				}
			}
		}
		for i := 0; i < 2; i++ {
			if err := d.Decode(new(json.RawMessage)); err != io.EOF {
				t.Fatalf("Decode after last value: got %v, want EOF", err)
			}
			if _, err := io.ReadAll(d.Buffered()); err != nil {
				t.Fatal(err)
			}
		}
	})
}

// This is a metamorphic check: changing short-read boundaries, injecting
// occasional permitted (0, nil) reads, and returning EOF with the final bytes
// must not change value framing. It deliberately crosses the 4 KiB initial
// buffer boundary and keeps earlier RawMessages alive while the buffer is reused.
func FuzzStreamDecoderBufferLifecycle(f *testing.F) {
	f.Add([]byte("small"), []byte{}, uint8(0), uint8(0))
	f.Add(bytes.Repeat([]byte("x"), 4095), []byte{0, 7, 31}, uint8(1), uint8(3))
	f.Add(bytes.Repeat([]byte("y"), 8192), []byte{1}, uint8(0), uint8(1))
	f.Fuzz(func(t *testing.T, payload, plan []byte, ending, emptyEvery uint8) {
		if len(payload) > 64<<10 || len(plan) > 64 {
			t.Skip()
		}
		encoded, err := json.Marshal(string(payload))
		if err != nil {
			t.Fatal(err)
		}
		input := "0 \n" + string(encoded) + "\t-1.25e+3 true null {} []"
		want, err := streamFuzzValues(input)
		if err != nil {
			t.Fatal(err)
		}
		got, d, _ := decodeFuzzRawStream(t, input, plan, ending, emptyEvery)
		assertRawMessagesEqual(t, got, want)
		if d.InputOffset() < 0 || d.InputOffset() > int64(len(input)) {
			t.Fatalf("terminal offset %d outside input length %d", d.InputOffset(), len(input))
		}
		if buffered, err := io.ReadAll(d.Buffered()); err != nil || len(buffered) != 0 {
			t.Fatalf("terminal buffer=%q err=%v", buffered, err)
		}
		// The decoder has recycled and moved its internal buffer repeatedly.
		// Previously returned values must still own stable bytes.
		assertRawMessagesEqual(t, got, want)
	})
}

// Valid numeric streams also exercise the public conversion options. The
// stdlib UseNumber path is the oracle for generic values; UseInt64 is checked
// only for integer literals that strconv accepts without overflow.
func FuzzStreamDecoderNumericOptions(f *testing.F) {
	for _, input := range []string{"0", "42", "-1", "1.5", "1e+3", "0 -1 42"} {
		f.Add(input, []byte{}, uint8(0))
	}
	f.Fuzz(func(t *testing.T, input string, plan []byte, mode uint8) {
		if len(input) == 0 || len(input) > 2048 || len(plan) > 32 {
			t.Skip()
		}
		rawValues, err := streamFuzzValues(input)
		if err != nil || len(rawValues) == 0 {
			t.Skip()
		}
		var want []json.Number
		for _, raw := range rawValues {
			if len(raw) == 0 || (raw[0] != '-' && (raw[0] < '0' || raw[0] > '9')) {
				t.Skip()
			}
			var value json.Number
			if err := json.Unmarshal(raw, &value); err != nil || string(raw) != string(value) {
				t.Skip()
			}
			want = append(want, value)
		}
		r := &streamFuzzReader{input: input, plan: plan, tail: io.EOF, errWithData: mode&1 != 0, emptyEvery: int(mode>>1) % 4}
		d := NewStreamDecoder(r)
		if mode&8 == 0 {
			d.UseNumber()
			for i, expected := range want {
				var got interface{}
				if err := d.Decode(&got); err != nil || got != expected {
					t.Fatalf("UseNumber value %d: got=%#v err=%v want=%q input=%q plan=%x", i, got, err, expected, input, plan)
				}
			}
		} else {
			d.UseInt64()
			for i, expected := range want {
				integer, err := strconv.ParseInt(string(expected), 10, 64)
				if err != nil {
					t.Skip()
				}
				var got interface{}
				if err := d.Decode(&got); err != nil || !reflect.DeepEqual(got, integer) {
					t.Fatalf("UseInt64 value %d: got=%#v err=%v want=%d input=%q plan=%x", i, got, err, integer, input, plan)
				}
			}
		}
		if err := d.Decode(new(interface{})); err != io.EOF {
			t.Fatalf("terminal Decode: got %v, want EOF", err)
		}
	})
}

// Only EOF can terminate an undelimited numeric token. A non-EOF Reader error
// must survive Decode, including when Reader returns data and that error together.
func FuzzStreamDecoderReaderErrors(f *testing.F) {
	for _, number := range []string{"0", "42", "-1.5", "1e+3"} {
		f.Add(number, []byte{0}, uint8(0))
	}
	f.Fuzz(func(t *testing.T, number string, plan []byte, mode uint8) {
		if len(number) == 0 || len(number) > 256 || len(plan) > 32 || (number[0] != '-' && (number[0] < '0' || number[0] > '9')) {
			t.Skip()
		}
		var parsed json.Number
		if json.Unmarshal([]byte(number), &parsed) != nil || string(parsed) != number {
			t.Skip()
		}
		input := number
		delimited := mode&1 != 0
		if delimited {
			input += " "
		}
		readErr := errors.New("fuzz reader error")
		r := &streamFuzzReader{input: input, plan: plan, tail: readErr, errWithData: mode&2 != 0}
		d := NewStreamDecoder(r)
		d.UseNumber()
		var got interface{}
		err := d.Decode(&got)
		if delimited {
			if err != nil || got != parsed {
				t.Fatalf("delimited number: got=%v err=%v want=%q plan=%x", got, err, number, plan)
			}
			if !r.errWithData {
				// The space establishes the boundary before the Reader error.
				// A read past it could block on an actual open connection.
				if r.reads > len(input) || r.pos != len(input) {
					t.Fatalf("read after delimiter: reads=%d pos=%d input=%q", r.reads, r.pos, input)
				}
			}
		} else if err != readErr || got != nil {
			t.Fatalf("unterminated number committed: got=%v err=%v wantErr=%v input=%q plan=%x", got, err, readErr, input, plan)
		}
		for i := 0; i < 2; i++ {
			if err := d.Decode(new(interface{})); err != readErr {
				t.Fatalf("Reader error lost: got=%v want=%v", err, readErr)
			}
			if _, err := io.ReadAll(d.Buffered()); err != nil {
				t.Fatal(err)
			}
		}
	})
}

// Arbitrary bytes are useful even when there is no stable cross-library error
// oracle. This property checks liveness, monotonic bounded offsets, safe buffer
// inspection, and sticky terminal errors under adversarial Reader behavior.
func FuzzStreamDecoderRobustness(f *testing.F) {
	for _, input := range [][]byte{
		{}, []byte("-"), []byte("1e"), []byte("01"), []byte("[1,"),
		[]byte(`{"a":1}{"b":2}`), {0xff, 0x00, '1'},
	} {
		f.Add(input, []byte{0}, uint8(0))
	}
	f.Fuzz(func(t *testing.T, raw, plan []byte, behavior uint8) {
		if len(raw) > 4<<10 || len(plan) > 64 {
			t.Skip()
		}
		tail := error(io.EOF)
		if behavior&1 != 0 {
			tail = errors.New("fuzz terminal error")
		}
		r := &streamFuzzReader{
			input:       string(raw),
			plan:        plan,
			tail:        tail,
			errWithData: behavior&2 != 0,
			emptyEvery:  int(behavior>>2) % 8,
		}
		d := NewStreamDecoder(r)
		var previous int64
		var terminal error
		// Every successful Decode must consume at least one input byte, so one
		// more attempt than len(raw) is sufficient to reach a terminal error.
		for i := 0; i <= len(raw); i++ {
			var value json.RawMessage
			err := d.Decode(&value)
			offset := d.InputOffset()
			if offset < previous || offset > int64(len(raw)) {
				t.Fatalf("offset moved outside input: previous=%d current=%d len=%d", previous, offset, len(raw))
			}
			previous = offset
			if _, bufferedErr := io.ReadAll(d.Buffered()); bufferedErr != nil {
				t.Fatal(bufferedErr)
			}
			if err != nil {
				terminal = err
				break
			}
		}
		if terminal == nil {
			t.Fatal("decoder returned too many values without a terminal error")
		}
		for i := 0; i < 2; i++ {
			err := d.Decode(new(json.RawMessage))
			if reflect.TypeOf(err) != reflect.TypeOf(terminal) || (err != nil && err.Error() != terminal.Error()) {
				t.Fatalf("terminal error changed: first=%T %v later=%T %v", terminal, terminal, err, err)
			}
			if _, bufferedErr := io.ReadAll(d.Buffered()); bufferedErr != nil {
				t.Fatal(bufferedErr)
			}
		}
	})
}
