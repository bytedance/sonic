package ast

import (
	"errors"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/internal/native/types"
)

var (
	// ErrNotExist means both key and value doesn't exist
	ErrNotExist = newError(errors.New("not exist"))

	// ErrInvalidPath means path is invalid
	ErrInvalidPath = newError(errors.New("ErrInvalidPath"))

	// ErrUnsupportType means API on the node is unsupported
	ErrUnsupportType = newError(errors.New("not supported type"))
)

func newError(err error) Node {
	return Node{
		node: types.Node{
			Kind: V_ERROR,
			JSON: err.Error(),
		},
	}
}

// Error returns error message if the node is invalid
func (self Node) Error() string {
	if self.node.Kind == V_ERROR {
		return self.node.JSON
	} else if self.node.Kind == V_NONE {
		return "not exist"
	} else {
		return ""
	}
}

func makeSyntaxError(json string, p int, msg string) decoder.SyntaxError {
	return decoder.SyntaxError{
		Pos: p,
		Src: json,
		Msg: msg,
	}
}

// speciall used for `SetByPath()`, tells which path is array index of range
type ErrIndexOutOfRange struct {
	Index int // index of path
	Err   error
}

func (e ErrIndexOutOfRange) Error() string {
	return e.Err.Error()
}
