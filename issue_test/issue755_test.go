package issue_test

import (
	"testing"
	"github.com/bytedance/sonic"
)

var _emptyFunc func()

func TestIssue755_NilEfaceWithDirectValue(t *testing.T) {
	tests := []interface{} {
		struct {
			Foo *int
		}{},
		struct {
			Foo func()
		}{},
		chan int(nil),
		_emptyFunc,
	}
	for _, v := range(tests) {
		assertMarshal(t, sonic.ConfigDefault, v)
	}
}

type NilMarshaler struct {}

func (n *NilMarshaler) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte(`"my null value"`), nil
	}
	return []byte(`{}`), nil
}

func TestIssue755_MarshalIface(t *testing.T) {
	tests := []interface{} {
		&NilMarshaler{},
		(*NilMarshaler)(nil),
	}
	for _, v := range(tests) {
		assertMarshal(t, sonic.ConfigDefault, v)
	}
}

