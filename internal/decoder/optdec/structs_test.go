package optdec

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/bytedance/sonic/internal/decoder/consts"
	"github.com/stretchr/testify/require"
)

func TestStructFieldOrder(t *testing.T) {
	type record struct {
		A int
		B string
		C bool
		D []int
	}
	fields := []string{`"A":1`, `"B":"two"`, `"C":true`, `"D":[3,4]`}
	// Exercise all orders, including transitions from a fallback lookup to a hint.
	for a := range fields {
		for b := range fields {
			for c := range fields {
				for d := range fields {
					if a == b || a == c || a == d || b == c || b == d || c == d {
						continue
					}
					input := "{" + strings.Join([]string{fields[a], fields[b], fields[c], fields[d]}, ",") + "}"
					t.Run(fmt.Sprintf("%d%d%d%d", a, b, c, d), func(t *testing.T) {
						var got, want record
						require.NoError(t, json.Unmarshal([]byte(input), &want))
						pos := 0
						require.NoError(t, Decode(&input, &pos, 0, &got))
						require.Equal(t, want, got)
					})
				}
			}
		}
	}
}

func TestStructFieldHintFallback(t *testing.T) {
	type embedded struct {
		Promoted int
	}
	type record struct {
		embedded
		A      int `json:"name"`
		B      int `json:"NAME"`
		C      int
		Nested *struct {
			A int
			B int
		}
	}
	for _, input := range []string{
		`{}`,
		`{"C":3}`,
		`{"name":1,"name":2,"NAME":3,"C":4,"C":5}`,
		`{"missing":{"C":99},"name":1,"unknown":null,"NAME":2,"C":3}`,
		`{"na\u006de":1,"NAME":2,"C":3}`,
		`{"nAmE":1,"NAME":2,"c":3}`,
		`{"Promoted":1,"name":2,"NAME":3,"C":4}`,
		`{"Nested":{"B":2,"A":1},"C":3,"name":4}`,
		`{"name":"wrong","NAME":2,"C":3}`,
	} {
		t.Run(input, func(t *testing.T) {
			var got, want record
			wantErr := json.Unmarshal([]byte(input), &want)
			pos := 0
			gotErr := Decode(&input, &pos, 0, &got)
			require.Equal(t, wantErr == nil, gotErr == nil)
			require.Equal(t, want, got)
		})
	}
	t.Run("case sensitive", func(t *testing.T) {
		input := `{"nAmE":99,"name":1,"NAME":2,"c":99,"C":3}`
		var got record
		pos := 0
		require.NoError(t, Decode(&input, &pos, uint64(consts.OptionCaseSensitive), &got))
		require.Equal(t, record{A: 1, B: 2, C: 3}, got)
	})
	t.Run("reject unknown", func(t *testing.T) {
		input := `{"name":1,"unknown":2,"NAME":3}`
		var got record
		pos := 0
		err := Decode(&input, &pos, uint64(OptionDisableUnknown), &got)
		require.EqualError(t, err, `json: unknown field "unknown"`)
		require.Equal(t, record{A: 1}, got)
	})
}

func TestStructFieldHintConcurrent(t *testing.T) {
	type record struct{ A, B, C int }
	inputs := []string{`{"A":1,"B":2,"C":3}`, `{"C":3,"A":1,"B":2}`}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(input string) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				var got record
				pos := 0
				if err := Decode(&input, &pos, 0, &got); err != nil || got != (record{1, 2, 3}) {
					t.Errorf("decode %s: got %+v, error %v", input, got, err)
					return
				}
			}
		}(inputs[i%len(inputs)])
	}
	wg.Wait()
}
