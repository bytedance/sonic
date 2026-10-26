//go:build amd64
// +build amd64

package encoder

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bytedance/sonic/option"
	"github.com/stretchr/testify/assert"
)

type testMutualA struct {
	Name string       `json:"name"`
	B    *testMutualB `json:"b,omitempty"`
}

type testMutualB struct {
	ID int          `json:"id"`
	A  *testMutualA `json:"a,omitempty"`
}

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

func TestPretouchEncodeMultiLevelStructWithRec(t *testing.T) {
	type level3 struct {
		Score int      `json:"score"`
		Tags  []string `json:"tags"`
	}
	type level2 struct {
		Label string  `json:"label"`
		Next  *level3 `json:"next"`
	}
	type level1 struct {
		ID   int    `json:"id"`
		Node level2 `json:"node"`
	}
	type root struct {
		Name string `json:"name"`
		Data level1 `json:"data"`
	}

	v := root{
		Name: "pretouch-rec",
		Data: level1{
			ID: 7,
			Node: level2{
				Label: "L2",
				Next: &level3{
					Score: 99,
					Tags:  []string{"a", "b"},
				},
			},
		},
	}

	err := Pretouch(
		reflect.TypeOf(root{}),
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	got, err := Encode(v, 0)
	assert.NoError(t, err)

	want, err := json.Marshal(v)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPretouchManyEncodeMultiTypes(t *testing.T) {
	type leafA struct {
		V int `json:"v"`
	}
	type rootA struct {
		Name string `json:"name"`
		Sub  leafA  `json:"sub"`
	}
	type leafB struct {
		S string `json:"s"`
	}
	type rootB struct {
		ID   int    `json:"id"`
		Next *leafB `json:"next"`
	}

	vA := rootA{Name: "A", Sub: leafA{V: 1}}
	vB := rootB{ID: 2, Next: &leafB{S: "B"}}

	err := PretouchMany(
		[]reflect.Type{reflect.TypeOf(rootA{}), reflect.TypeOf(rootB{})},
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	gotA, err := Encode(vA, 0)
	assert.NoError(t, err)
	wantA, err := json.Marshal(vA)
	assert.NoError(t, err)
	assert.Equal(t, wantA, gotA)

	gotB, err := Encode(vB, 0)
	assert.NoError(t, err)
	wantB, err := json.Marshal(vB)
	assert.NoError(t, err)
	assert.Equal(t, wantB, gotB)
}

func TestPretouchManyMutualRefTypesEncodeOK(t *testing.T) {
	err := PretouchMany(
		[]reflect.Type{reflect.TypeOf(testMutualA{}), reflect.TypeOf(testMutualB{})},
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	v := testMutualA{Name: "a", B: &testMutualB{ID: 7}}
	got, err := Encode(v, 0)
	assert.NoError(t, err)

	want, err := json.Marshal(v)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPretouchManyMutualRefValueCycleEncodeError(t *testing.T) {
	err := PretouchMany(
		[]reflect.Type{reflect.TypeOf(testMutualA{}), reflect.TypeOf(testMutualB{})},
		option.WithCompileMaxInlineDepth(1),
		option.WithCompileRecursiveDepth(8),
	)
	assert.NoError(t, err)

	a := &testMutualA{Name: "cycle"}
	b := &testMutualB{ID: 1}
	a.B = b
	b.A = a

	_, err = Encode(a, 0)
	assert.Error(t, err)
	_, ok := err.(*json.UnsupportedValueError)
	assert.True(t, ok)
}
