package internal

import (
	"fmt"

	"github.com/bytedance/sonic/internal/rt"
)

type ParseError struct {
	Offset    int
	Msg       string
}

func newError(msg string, offset int) error {
	return &ParseError{
		Offset:    offset,
		Msg:       msg,
	}
}

type UnmatchedError struct {
	Offset    int
	Type	  *rt.GoType
}

func newUnmatched(pos int, typ *rt.GoType) error {
	return &UnmatchedError{
		Offset:	pos,
		Type: 	typ,
	}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Internal Parse Error: %v\n", e.Msg)
}

func (e *UnmatchedError) Error() string {
	return fmt.Sprintf("Internal Unmatched Error for Type %v\n", e.Type)
}
