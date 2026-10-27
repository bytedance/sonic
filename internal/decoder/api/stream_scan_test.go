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
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStreamValueScanner(t *testing.T) {
	for _, tc := range []struct {
		src   string
		final bool
		state streamScanStatus
		end   int
	}{
		{"4", false, streamNeedMore, 1}, {"4", true, streamDone, 1},
		{"1e", false, streamNeedMore, 2}, {"1e", true, streamInvalid, 2},
		{"1e-", false, streamNeedMore, 3}, {"1e-3 ", false, streamDone, 4},
		{"42 ", false, streamDone, 2}, {"42true", false, streamDone, 2},
		{"01", false, streamDone, 1}, {"-01", false, streamDone, 2},
		{"-", false, streamNeedMore, 1}, {"-", true, streamInvalid, 1},
		{"1ee", false, streamInvalid, 2}, {"1.x", false, streamInvalid, 2},
		{"{", false, streamNeedMore, 1}, {"{", true, streamInvalid, 1},
		{`{"x":1}`, false, streamDone, 7}, {"?", false, streamInvalid, 0},
		{strings.Repeat("1", 64) + " 2", false, streamDone, 64},
	} {
		t.Run(tc.src, func(t *testing.T) {
			var s streamValueScanner
			end, state, err := s.scan(tc.src, tc.final)
			require.Equal(t, tc.state, state)
			require.Equal(t, tc.end, end)
			if state == streamInvalid {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStreamValueScannerContinuesNumber(t *testing.T) {
	var s streamValueScanner
	for _, prefix := range []string{"-", "-1", "-1.", "-1.5", "-1.5e", "-1.5e-", "-1.5e-3"} {
		end, state, err := s.scan(prefix, false)
		require.NoError(t, err)
		require.Equal(t, streamNeedMore, state)
		require.Equal(t, len(prefix), end)
	}
	end, state, err := s.scan("-1.5e-3 2", false)
	require.NoError(t, err)
	require.Equal(t, streamDone, state)
	require.Equal(t, 7, end)

	var unfinished streamValueScanner
	_, state, err = unfinished.scan("1e-", true)
	require.Equal(t, streamInvalid, state)
	require.Equal(t, io.ErrUnexpectedEOF, err)
}

func FuzzStreamValueScanner(f *testing.F) {
	for _, src := range []string{"0", "-0", "42", "-1.5", "1e+3", "1e-3", strings.Repeat("1", 64)} {
		f.Add(src, uint8(1), false)
	}
	f.Fuzz(func(t *testing.T, number string, chunk uint8, eof bool) {
		if len(number) == 0 || len(number) > 128 || (number[0] != '-' && (number[0] < '0' || number[0] > '9')) {
			t.Skip()
		}
		var value json.Number
		if err := json.Unmarshal([]byte(number), &value); err != nil || value.String() != number {
			t.Skip()
		}
		input := number
		if !eof {
			input += " {}"
		}
		var s streamValueScanner
		step := 1 + int(chunk)%32
		for end := step; ; end += step {
			if end > len(input) {
				end = len(input)
			}
			n, state, err := s.scan(input[:end], end == len(input))
			require.NoError(t, err)
			if state == streamDone {
				require.Equal(t, len(number), n)
				return
			}
			require.Equal(t, streamNeedMore, state)
			require.Less(t, end, len(input))
		}
	})
}
