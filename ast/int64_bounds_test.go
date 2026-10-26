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

package ast

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestInt64TextBounds(t *testing.T) {
	for _, tc := range []struct {
		text     string
		want     int64
		overflow bool
	}{
		{"-9223372036854775809", 0, true},
		{"-9223372036854775809.0", 0, true},
		{"-9.223372036854775809e18", 0, true},
		{"9223372036854775808", 0, true},
		{"9223372036854775808.0", 0, true},
		{"9.223372036854775808e18", 0, true},
		{"-9223372036854775808", math.MinInt64, false},
		{"-9223372036854775808.9", math.MinInt64, false},
		{"9223372036854775807", math.MaxInt64, false},
		{"9223372036854775807.9", math.MaxInt64, false},
		{"9.223372036854775807e18", math.MaxInt64, false},
		{"-9223372036854775807.0", math.MinInt64 + 1, false},
		{"1.9", 1, false},
		{"-1.9", -1, false},
		{"1e3", 1000, false},
	} {
		t.Run(tc.text, func(t *testing.T) {
			for name, n := range map[string]Node{
				"raw": NewRaw(tc.text), "number": NewNumber(tc.text),
				"string": NewString(tc.text), "any-string": NewAny(tc.text),
				"any-number": NewAny(json.Number(tc.text)),
			} {
				t.Run(name, func(t *testing.T) {
					got, err := n.Int64()
					if tc.overflow {
						if !errors.Is(err, strconv.ErrRange) || got != 0 {
							t.Fatalf("Int64() = %d, %v; want 0, ErrRange", got, err)
						}
					} else if err != nil || got != tc.want {
						t.Fatalf("Int64() = %d, %v; want %d, nil", got, err, tc.want)
					}
				})
			}
		})
	}
}
