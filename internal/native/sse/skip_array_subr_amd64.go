// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
`github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __skip_array_entry() uintptr

var (
    _subr__skip_array uintptr = rt.GetFuncPC(__skip_array_entry) + 190
)

const (
    _stack__skip_array = 160
)

var (
    _ = _subr__skip_array
)

const (
    _ = _stack__skip_array
)
