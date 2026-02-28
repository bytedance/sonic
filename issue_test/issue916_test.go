package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue916_StringTagTypeMismatchShouldContinue(t *testing.T) {
	type A struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	type B struct {
		ID   int64  `json:"id,string"`
		Name string `json:"name"`
	}

	data, err := sonic.Marshal(A{
		ID:   1,
		Name: "test1",
	})
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	for _, cas := range []unmTestCase{
		{
			name: "issue916",
			data: data,
			newfn: func() interface{} {
				return new(B)
			},
		},
	} {
		assertUnmarshal(t, sonic.ConfigDefault, cas)
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}
