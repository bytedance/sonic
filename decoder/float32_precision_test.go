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
