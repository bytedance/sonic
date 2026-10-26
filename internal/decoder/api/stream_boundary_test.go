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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"
)

type streamReadStep struct {
	data string
	err  error
}

type streamStepReader struct {
	steps []streamReadStep
	reads int
	tail  error
}

func (r *streamStepReader) Read(p []byte) (int, error) {
	r.reads++
	if len(r.steps) == 0 {
		return 0, r.tail
	}
	s := &r.steps[0]
	n := copy(p, s.data)
	s.data = s.data[n:]
	if s.data != "" {
		return n, nil
	}
	err := s.err
	r.steps = r.steps[1:]
	return n, err
}

func TestStreamDecodeNumberAcrossReads(t *testing.T) {
	for _, input := range []string{"42", "42 ", "42\n", "-7", "1234567890", "1.5", "1e3", "0", "1e-3", " 42"} {
		for _, useNumber := range []bool{false, true} {
			t.Run(fmt.Sprintf("%q/number=%v", input, useNumber), func(t *testing.T) {
				want := json.NewDecoder(iotest.OneByteReader(strings.NewReader(input)))
				got := NewStreamDecoder(iotest.OneByteReader(strings.NewReader(input)))
				if useNumber {
					want.UseNumber()
					got.UseNumber()
				}
				var expected, actual interface{}
				require.NoError(t, want.Decode(&expected))
				require.NoError(t, got.Decode(&actual))
				require.Equal(t, expected, actual)
				require.Equal(t, want.InputOffset(), got.InputOffset())
				require.NotPanics(t, func() { _, err := io.ReadAll(got.Buffered()); require.NoError(t, err) })
				require.Equal(t, io.EOF, got.Decode(new(interface{})))
			})
		}
	}
}

func TestStreamNumberStopsAtDelimiter(t *testing.T) {
	for _, space := range []string{" ", "\t", "\r", "\n"} {
		t.Run(fmt.Sprintf("%q", space), func(t *testing.T) {
			readErr := errors.New("read after delimiter")
			r := &streamStepReader{steps: []streamReadStep{{data: "42"}, {data: space}}, tail: readErr}
			d := NewStreamDecoder(r)
			var got int
			require.NoError(t, d.Decode(&got))
			require.Equal(t, 42, got)
			require.Equal(t, 2, r.reads, "Decode must not wait for data after the delimiter")
			require.Equal(t, int64(2), d.InputOffset())
			require.NotPanics(t, func() { _ = d.Buffered() })
		})
	}
}

func TestStreamNumberLookaheadError(t *testing.T) {
	for _, prefix := range []string{"4", " 4"} {
		for _, withData := range []bool{false, true} {
			t.Run(fmt.Sprintf("%q/data=%v", prefix, withData), func(t *testing.T) {
				readErr := errors.New("number read failed")
				steps := []streamReadStep{{data: prefix}, {data: "2"}}
				reads := 3
				if withData {
					steps[1].err = readErr
					reads = 2
				}
				r := &streamStepReader{steps: steps, tail: readErr}
				d := NewStreamDecoder(r)
				got := 99
				require.Equal(t, readErr, d.Decode(&got))
				require.Equal(t, 99, got)
				require.Equal(t, readErr, d.Decode(&got))
				require.Equal(t, reads, r.reads, "do not read again after the saved error")
				require.NotPanics(t, func() { _ = d.Buffered() })
			})
		}
	}
}

func TestStreamNumberLookaheadDataAndError(t *testing.T) {
	readErr := errors.New("lookahead failed")
	for _, tc := range []struct {
		name   string
		suffix string
		err    error
	}{
		{"eof", "2", io.EOF},
		{"delimiter", "2 ", readErr},
		{"buffered-value", "2 {}", io.EOF},
		{"incomplete-buffered-value", "2 {", readErr},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := &streamStepReader{steps: []streamReadStep{{data: "4"}, {data: tc.suffix, err: tc.err}}, tail: errors.New("unexpected read")}
			d := NewStreamDecoder(r)
			var got interface{}
			require.NoError(t, d.Decode(&got))
			require.Equal(t, float64(42), got)
			require.NotPanics(t, func() { _ = d.Buffered() })
			if tc.name == "buffered-value" {
				require.NoError(t, d.Decode(&got))
				require.Equal(t, map[string]interface{}{}, got)
			}
			require.Equal(t, tc.err, d.Decode(&got))
			require.Equal(t, 2, r.reads)
		})
	}
}

func TestStreamIncompleteNumber(t *testing.T) {
	for _, input := range []string{"-", "1.", "1e", "1e+", "1e-"} {
		t.Run(input, func(t *testing.T) {
			d := NewStreamDecoder(iotest.OneByteReader(strings.NewReader(input)))
			got := 99
			require.Error(t, d.Decode(&got))
			require.Equal(t, 99, got)
		})
	}
}

func TestStreamNumberEverySplit(t *testing.T) {
	for _, input := range []string{"42 43", "01", "1-2", "42true", "42{}", "-1.5e-3 2", strings.Repeat("1", 64) + " 2", "42" + strings.Repeat(" ", 64) + "43"} {
		for split := 1; split < len(input); split++ {
			t.Run(fmt.Sprintf("%q/%d", input, split), func(t *testing.T) {
				reader := func() io.Reader {
					return io.MultiReader(strings.NewReader(input[:split]), strings.NewReader(input[split:]))
				}
				want := json.NewDecoder(reader())
				want.UseNumber()
				got := NewStreamDecoder(reader())
				got.UseNumber()
				for count := 0; count <= len(input); count++ {
					var expected, actual interface{}
					err := want.Decode(&expected)
					if err == io.EOF {
						require.Equal(t, io.EOF, got.Decode(&actual))
						return
					}
					require.NoError(t, err)
					require.NoError(t, got.Decode(&actual))
					require.Equal(t, expected, actual)
				}
				t.Fatal("decoder did not reach EOF")
			})
		}
	}
}

func TestStreamNumberFirstReadDataAndError(t *testing.T) {
	readErr := errors.New("first read failed")
	for _, tc := range []struct {
		input   string
		err     error
		success bool
	}{
		{"42", io.EOF, true}, {"42", readErr, false}, {"42 ", readErr, true},
	} {
		t.Run(fmt.Sprintf("%q/%v", tc.input, tc.err), func(t *testing.T) {
			r := &streamStepReader{steps: []streamReadStep{{data: tc.input, err: tc.err}}, tail: errors.New("unexpected read")}
			d := NewStreamDecoder(r)
			got := 99
			err := d.Decode(&got)
			if tc.success {
				require.NoError(t, err)
				require.Equal(t, 42, got)
			} else {
				require.Equal(t, readErr, err)
				require.Equal(t, 99, got)
			}
			require.Equal(t, tc.err, d.Decode(&got))
			require.Equal(t, 1, r.reads)
			require.NotPanics(t, func() { _ = d.Buffered() })
		})
	}
}
