// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f64toa      = __native_entry__() + 825
    _subr__html_escape = __native_entry__() + 10636
    _subr__i64toa      = __native_entry__() + 3908
    _subr__lspace      = __native_entry__() + 379
    _subr__lzero       = __native_entry__() + 13
    _subr__quote       = __native_entry__() + 6308
    _subr__skip_array  = __native_entry__() + 21086
    _subr__skip_object = __native_entry__() + 21121
    _subr__skip_one    = __native_entry__() + 19344
    _subr__u64toa      = __native_entry__() + 4003
    _subr__unquote     = __native_entry__() + 8141
    _subr__value       = __native_entry__() + 14442
    _subr__vnumber     = __native_entry__() + 17358
    _subr__vsigned     = __native_entry__() + 18788
    _subr__vstring     = __native_entry__() + 16451
    _subr__vunsigned   = __native_entry__() + 19068
)

const (
    _stack__f64toa = 120
    _stack__html_escape = 80
    _stack__i64toa = 24
    _stack__lspace = 8
    _stack__lzero = 8
    _stack__quote = 80
    _stack__skip_array = 144
    _stack__skip_object = 144
    _stack__skip_one = 144
    _stack__u64toa = 8
    _stack__unquote = 72
    _stack__value = 408
    _stack__vnumber = 320
    _stack__vsigned = 16
    _stack__vstring = 112
    _stack__vunsigned = 16
)

var (
    _ = _subr__f64toa
    _ = _subr__html_escape
    _ = _subr__i64toa
    _ = _subr__lspace
    _ = _subr__lzero
    _ = _subr__quote
    _ = _subr__skip_array
    _ = _subr__skip_object
    _ = _subr__skip_one
    _ = _subr__u64toa
    _ = _subr__unquote
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
    _ = _stack__skip_object
    _ = _stack__skip_one
    _ = _stack__u64toa
    _ = _stack__unquote
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
