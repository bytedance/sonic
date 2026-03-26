package jitdec

import (
	"reflect"
	"testing"

	"github.com/bytedance/sonic/option"
	"github.com/stretchr/testify/assert"
)

type testPretouchManyA struct {
	Name string             `json:"name"`
	B    *testPretouchManyB `json:"b,omitempty"`
}

type testPretouchManyB struct {
	ID int                `json:"id"`
	A  *testPretouchManyA `json:"a,omitempty"`
}

func TestPretouchMany(t *testing.T) {
	err := PretouchMany(
		[]reflect.Type{reflect.TypeOf(testPretouchManyA{}), reflect.TypeOf(testPretouchManyB{})},
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	s := `{"name":"x","b":{"id":7}}`
	i := 0
	var out testPretouchManyA
	err = Decode(&s, &i, 0, &out)
	assert.NoError(t, err)
	assert.Equal(t, "x", out.Name)
	if assert.NotNil(t, out.B) {
		assert.Equal(t, 7, out.B.ID)
	}
}
