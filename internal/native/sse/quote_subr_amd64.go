// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
`github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __quote_entry() uintptr

var (
    _subr__quote uintptr = rt.GetFuncPC(__quote_entry) + 78
)

const (
    _stack__quote = 64
)

var (
    _ = _subr__quote
)

const (
    _ = _stack__quote
)
