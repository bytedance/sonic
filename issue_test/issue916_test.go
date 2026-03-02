package issue_test

import (
	"reflect"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/option"
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

func TestIssue916_StringTagTypeMismatchShouldContinue_WithLowInlineDepth(t *testing.T) {
	type A struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	type B struct {
		ID   int64  `json:"id,string"`
		Name string `json:"name"`
	}

	if err := sonic.Pretouch(reflect.TypeOf(B{}), option.WithCompileMaxInlineDepth(2), option.WithCompileRecursiveDepth(8)); err != nil {
		t.Fatalf("pretouch failed: %v", err)
	}

	data, err := sonic.Marshal(A{
		ID:   1,
		Name: "test1",
	})
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	cas := unmTestCase{
		name: "issue916-inline-depth2",
		data: data,
		newfn: func() interface{} {
			return new(B)
		},
	}

	assertUnmarshal(t, sonic.ConfigDefault, cas)
	assertUnmarshal(t, sonic.ConfigStd, cas)
}
