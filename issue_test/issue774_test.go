package issue_test

import (
	"testing"

	"encoding/json"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue774_EmptyStringForByteSlice(t *testing.T) {
	var jv, sv struct {
		ByteArr []byte
	}

	err := sonic.Unmarshal([]byte(`{"ByteArr": ""}`), &sv)
	err2 := json.Unmarshal([]byte(`{"ByteArr": ""}`), &jv)
	require.Equal(t, err2 == nil, err == nil)
	require.Equal(t, jv, sv) 
}