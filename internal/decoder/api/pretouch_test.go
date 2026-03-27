package api

import (
	"reflect"
	"testing"

	"github.com/bytedance/sonic/option"
	"github.com/stretchr/testify/assert"
)

type testAPIPretouchManyA struct {
	Name string                `json:"name"`
	B    *testAPIPretouchManyB `json:"b,omitempty"`
}

type testAPIPretouchManyB struct {
	ID int                   `json:"id"`
	A  *testAPIPretouchManyA `json:"a,omitempty"`
}

func TestAPIPretouchMany(t *testing.T) {
	err := PretouchMany(
		[]reflect.Type{reflect.TypeOf(testAPIPretouchManyA{}), reflect.TypeOf(testAPIPretouchManyB{})},
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	dec := NewDecoder(`{"name":"api","b":{"id":3}}`)
	var out testAPIPretouchManyA
	err = dec.Decode(&out)
	assert.NoError(t, err)
	assert.Equal(t, "api", out.Name)
	if assert.NotNil(t, out.B) {
		assert.Equal(t, 3, out.B.ID)
	}
}
