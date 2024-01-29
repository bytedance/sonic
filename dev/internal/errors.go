package internal

import (
	"fmt"
	"runtime"
)

type ParseError struct {
	Offset    int64
	Msg       string
	Backtrace string
}

func newError(msg string, offset int64) error {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	return &ParseError{
		Offset:    offset,
		Msg:       msg,
		Backtrace: string(buf),
	}
}

type UnmatchedError struct {
	Offset    int
	Msg       string
	Backtrace string
}

func newUnmatched(msg string) error {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	return &UnmatchedError{
		Offset:    -1,
		Msg:       msg,
		Backtrace: string(buf),
	}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Internal Parse Error: %v\n%v\n", e.Msg, e.Backtrace)
}

func (e *UnmatchedError) Error() string {
	return fmt.Sprintf("Internal Unmatched Error: %v\n%v\n", e.Msg, e.Backtrace)
}
