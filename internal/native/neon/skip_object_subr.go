// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package neon

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __skip_object_entry__() uintptr

var (
    _subr__skip_object uintptr = __skip_object_entry__() + 64
)

const (
    _stack__skip_object = 160
)

var (
    _ = _subr__skip_object
)

const (
    _ = _stack__skip_object
)