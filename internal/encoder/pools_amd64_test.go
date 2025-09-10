//go:build amd64
// +build amd64

package encoder

import (
	"reflect"
	"testing"

	"github.com/bytedance/sonic/option"
	"github.com/stretchr/testify/assert"
)

func TestPretouchTypeX86(t *testing.T) {
	type subA struct{}
	type subB struct{}
	type subC struct{}
	type data struct {
		SubA subA
		SubB subB
		SubC subC
	}

	sub, err := pretouchTypeX86(
		reflect.TypeOf(data{}),
		option.CompileOptions{
			MaxInlineDepth: 1,
			RecursiveDepth: 1000,
		},
		0,
	)
	assert.NoError(t, err)
	assert.Contains(t, sub, reflect.TypeOf(subA{}))
	assert.Contains(t, sub, reflect.TypeOf(subB{}))
	assert.Contains(t, sub, reflect.TypeOf(subC{}))
}
