package sonic

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
    require.False(t, Valid(nil))

    testCase := []struct {
        data     string
        expected bool
    }{
        {``, false},
        {`s`, false},
        {`{`, false},
        {`[`, false},
        {`[1,2`, false},
        {`{"so":nic"}`, false},

        {`null`, true},
        {`""`, true},
        {`1`, true},
        {`"sonic"`, true},
        {`{}`, true},
        {`[]`, true},
        {`[1,2]`, true},
        {`{"so":"nic"}`, true},
    }
    for _, tc := range testCase {
        require.Equal(t, tc.expected, Valid([]byte(tc.data)), tc.data)
    }
}
