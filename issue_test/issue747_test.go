package issue_test

import (
	"testing"

	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestIssue747(t *testing.T) {
	tests := []struct {
		name           string
		input         string
		expected interface{}
		newfn func() interface{}
	}{
		{
			name:   "unmarshal map key float64",
			input: `{"1.2":1.8}`,
			expected: &map[float64]float64{
				1.2:  1.8,
			},
			newfn: func() interface{} { return new(map[float64]float64) },
		},
		{
			name:    "unmarshal map key float32",
			input: `{"1.2":1.8}`,
			expected: &map[float32]float32{
				1.2:  1.8,
			},
			newfn: func() interface{} { return new(map[float32]float32) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv, jv := tt.newfn(), tt.newfn()
			serr := sonic.Unmarshal([]byte(tt.input), &sv)
			assert.Equal(t, tt.expected, sv)
			assert.NoError(t, serr)

			// Note: it is different from encoding/json 
			jerr := json.Unmarshal([]byte(tt.input), &jv)
			assert.NotEqual(t, jerr == nil, serr == nil)
			assert.NotEqual(t, jv, sv)
		})
	}
}
