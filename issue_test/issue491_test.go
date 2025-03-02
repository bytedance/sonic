package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

type Function = func()

type Unsupported struct {
	Functions []Function
}

type StructWithUnsupported struct {
	Foo *Unsupported	`json:"foo"`
	Bar *Unsupported	`json:"bar,omitempty"`
}

type Foo2 struct {
	A int
	B *chan int
}

type MockContext struct {
	*Foo2
}

func TestIssue491_MarshalUnsupportedType(t *testing.T) {
	// Wrapper a unbale serde type
	tests := []interface{} {
		map[string]*Function{},
		map[*Function]*Function{},
		map[string]Function{},
		[]Function{},
		StructWithUnsupported{},
		struct {
			Foo *int
		}{},
		struct {
			Foo Function
		}{},
		chan int(nil),
		new(MockContext),
	}
	for _, v := range(tests) {
		assertMarshal(t, sonic.ConfigDefault, v)
	}
}

 func TestIssue491_UnmarshalUnsupported(t *testing.T) {
	type Test struct {
		data string
		value interface{}
	}

	tests := []unmTestCase{
		{
			name: "unsupported type slice",
			data: []byte("null"),
			newfn: func() interface{} { return new([]Function)},
		},
		{ 
			name: "unsupported type",
			data: []byte("[null, null]"),
			newfn: func() interface{} { return new([]chan int) },
		},
		{
			name: "unsupported type in struct",
			data: []byte("{\"foo\": null}"),
			newfn: func() interface{} {  
				return new(struct {
					Foo Function
				})
			},
		},
		{
			name: "unsupported type in map key should be error",
			data: []byte("null"),
			newfn: func() interface{} {  
				return map[chan int]Function{}
			},
		},
	}
	for _, v := range(tests) {
		assertUnmarshal(t, sonic.ConfigDefault, v)
	}
 }
 