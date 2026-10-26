//go:build (amd64 && go1.17 && !go1.28) || (arm64 && go1.20 && !go1.28)
// +build amd64,go1.17,!go1.28 arm64,go1.20,!go1.28

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

package decoder

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"
)

func TestFloat32PrecisionPaths(t *testing.T) {
	for _, src := range []string{"7.328900098800659", "-7.328900098800659", "3.4028235e38", "1.401298464324817e-45", "1e-100", "-0", "3.4028236e38"} {
		t.Run(src, func(t *testing.T) {
			var want float32
			stdErr := json.Unmarshal([]byte(src), &want)
			check := func(t *testing.T, got float32, err error) {
				t.Helper()
				if (err != nil) != (stdErr != nil) || (stdErr == nil && math.Float32bits(got) != math.Float32bits(want)) {
					t.Fatalf("got %g (%08x), %v; want %g (%08x), error=%v", got, math.Float32bits(got), err, want, math.Float32bits(want), stdErr)
				}
			}
			t.Run("value", func(t *testing.T) {
				var got float32
				_, err := decode(src, &got, false)
				check(t, got, err)
			})
			t.Run("string-tag", func(t *testing.T) {
				var got struct {
					F float32 `json:"f,string"`
				}
				_, err := decode(`{"f":`+strconv.Quote(src)+`}`, &got, false)
				check(t, got.F, err)
			})
			t.Run("map-key", func(t *testing.T) {
				var got map[float32]int
				_, err := decode(`{`+strconv.Quote(src)+`:1}`, &got, false)
				var key float32
				for k := range got {
					key = k
				}
				check(t, key, err)
			})
		})
	}
}
