// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f64toa       = __native_entry__() + 814
    _subr__html_escape  = __native_entry__() + 10717
    _subr__i64toa       = __native_entry__() + 3449
    _subr__lspace       = __native_entry__() + 379
    _subr__lzero        = __native_entry__() + 13
    _subr__quote        = __native_entry__() + 4842
    _subr__skip_array   = __native_entry__() + 25356
    _subr__skip_number  = __native_entry__() + 27886
    _subr__skip_object  = __native_entry__() + 25393
    _subr__skip_one     = __native_entry__() + 23264
    _subr__u64toa       = __native_entry__() + 3544
    _subr__unquote      = __native_entry__() + 7467
    _subr__validate_one = __native_entry__() + 28003
    _subr__value        = __native_entry__() + 15533
    _subr__vnumber      = __native_entry__() + 21377
    _subr__vsigned      = __native_entry__() + 22682
    _subr__vstring      = __native_entry__() + 17698
    _subr__vunsigned    = __native_entry__() + 22962
)

const (
    _stack__f64toa = 136
    _stack__html_escape = 72
    _stack__i64toa = 24
    _stack__lspace = 8
    _stack__lzero = 8
    _stack__quote = 72
    _stack__skip_array = 168
    _stack__skip_number = 88
    _stack__skip_object = 168
    _stack__skip_one = 168
    _stack__u64toa = 8
    _stack__unquote = 72
    _stack__validate_one = 168
    _stack__value = 416
    _stack__vnumber = 312
    _stack__vsigned = 16
    _stack__vstring = 1136
    _stack__vunsigned = 24
)

var (
    _ = _subr__f64toa
    _ = _subr__html_escape
    _ = _subr__i64toa
    _ = _subr__lspace
    _ = _subr__lzero
    _ = _subr__quote
    _ = _subr__skip_array
    _ = _subr__skip_number
    _ = _subr__skip_object
    _ = _subr__skip_one
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
    _ = _stack__f64toa
    _ = _stack__html_escape
    _ = _stack__i64toa
    _ = _stack__lspace
    _ = _stack__lzero
    _ = _stack__quote
    _ = _stack__skip_array
    _ = _stack__skip_number
    _ = _stack__skip_object
    _ = _stack__skip_one
    _ = _stack__u64toa
    _ = _stack__unquote
    _ = _stack__validate_one
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
