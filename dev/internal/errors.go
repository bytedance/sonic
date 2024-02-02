package internal

import (
	"fmt"
)

type ParseError struct {
	Offset    int64
	Msg       string
}

func newError(msg string, offset int64) error {
	return &ParseError{
		Offset:    offset,
		Msg:       msg,
	}
}

type UnmatchedError struct {
	Offset    int
	Msg       string
	Backtrace string
}

func newUnmatched(msg string) error {
	return &UnmatchedError{
		Offset:    0,
		Msg:       msg,
	}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Internal Parse Error: %v\n", e.Msg)
}

func (e *UnmatchedError) Error() string {
	return fmt.Sprintf("Internal Unmatched Error: %v\n", e.Msg)
}
