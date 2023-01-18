// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f32toa        = __native_entry__() + 28032
    _subr__f64toa        = __native_entry__() + 464
    _subr__get_by_path   = __native_entry__() + 26992
    _subr__html_escape   = __native_entry__() + 9584
    _subr__i64toa        = __native_entry__() + 3744
    _subr__lspace        = __native_entry__() + 80
    _subr__quote         = __native_entry__() + 5472
    _subr__skip_array    = __native_entry__() + 19760
    _subr__skip_number   = __native_entry__() + 23120
    _subr__skip_object   = __native_entry__() + 21664
    _subr__skip_one      = __native_entry__() + 23264
    _subr__skip_one_fast = __native_entry__() + 23472
    _subr__u64toa        = __native_entry__() + 4016
    _subr__unquote       = __native_entry__() + 7184
    _subr__validate_one  = __native_entry__() + 23296
    _subr__value         = __native_entry__() + 13216
    _subr__vnumber       = __native_entry__() + 17520
    _subr__vsigned       = __native_entry__() + 19056
    _subr__vstring       = __native_entry__() + 15408
    _subr__vunsigned     = __native_entry__() + 19408
)

const (
    _stack__f32toa = 56
    _stack__f64toa = 80
    _stack__get_by_path = 248
    _stack__html_escape = 64
    _stack__i64toa = 16
    _stack__lspace = 8
    _stack__quote = 80
    _stack__skip_array = 112
    _stack__skip_number = 72
    _stack__skip_object = 112
    _stack__skip_one = 112
    _stack__skip_one_fast = 144
    _stack__u64toa = 8
    _stack__unquote = 128
    _stack__validate_one = 112
    _stack__value = 368
    _stack__vnumber = 280
    _stack__vsigned = 16
    _stack__vstring = 128
    _stack__vunsigned = 24
)

var (
    _ = _subr__f32toa
    _ = _subr__f64toa
    _ = _subr__get_by_path
    _ = _subr__html_escape
    _ = _subr__i64toa
    _ = _subr__lspace
    _ = _subr__quote
    _ = _subr__skip_array
    _ = _subr__skip_number
    _ = _subr__skip_object
    _ = _subr__skip_one
    _ = _subr__skip_one_fast
    _ = _subr__u64toa
    _ = _subr__unquote
    _ = _subr__validate_one
    _ = _subr__value
    _ = _subr__vnumber
    _ = _subr__vsigned
    _ = _subr__vstring
    _ = _subr__vunsigned
)

const (
    _ = _stack__f32toa
    _ = _stack__f64toa
    _ = _stack__get_by_path
    _ = _stack__html_escape
    _ = _stack__i64toa
    _ = _stack__lspace
    _ = _stack__quote
    _ = _stack__skip_array
    _ = _stack__skip_number
    _ = _stack__skip_object
    _ = _stack__skip_one
    _ = _stack__skip_one_fast
    _ = _stack__u64toa
    _ = _stack__unquote
    _ = _stack__validate_one
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
