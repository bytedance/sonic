package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

type myfoo struct{}

func TestIssue739(t *testing.T) {
	var bar myfoo
	s := `{"a":"b
c"}`
	assert.NoError(t, sonic.ConfigDefault.UnmarshalFromString(s, &bar))
	assert.Error(t, sonic.ConfigStd.UnmarshalFromString(s, &bar))
}