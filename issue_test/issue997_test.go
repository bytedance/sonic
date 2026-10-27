package issue_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue997StreamNumberBoundaries(t *testing.T) {
	for _, input := range []string{"42 43", "-7 1.5", "1e3 1e-3"} {
		t.Run(input, func(t *testing.T) {
			want := json.NewDecoder(iotest.OneByteReader(strings.NewReader(input)))
			want.UseNumber()
			got := sonic.Config{UseNumber: true}.Froze().NewDecoder(iotest.OneByteReader(strings.NewReader(input)))
			for i := 0; i < 2; i++ {
				var expected, actual interface{}
				require.NoError(t, want.Decode(&expected))
				require.NoError(t, got.Decode(&actual))
				require.Equal(t, expected, actual)
			}
			var extra interface{}
			require.Equal(t, io.EOF, got.Decode(&extra))
		})
	}
}
