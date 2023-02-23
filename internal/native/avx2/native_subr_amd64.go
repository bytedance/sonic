// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f32toa             = __native_entry__() + 59936
    _subr__f64toa             = __native_entry__() + 736
    _subr__get_by_path        = __native_entry__() + 49664
    _subr__html_escape        = __native_entry__() + 12384
    _subr__i64toa             = __native_entry__() + 5456
    _subr__lspace             = __native_entry__() + 224
    _subr__quote              = __native_entry__() + 6960
    _subr__skip_array         = __native_entry__() + 27264
    _subr__skip_number        = __native_entry__() + 38448
    _subr__skip_object        = __native_entry__() + 29920
    _subr__skip_one           = __native_entry__() + 40288
    _subr__skip_one_fast      = __native_entry__() + 45872
    _subr__u64toa             = __native_entry__() + 5568
    _subr__unquote            = __native_entry__() + 9776
    _subr__validate_one       = __native_entry__() + 42944
    _subr__validate_utf8      = __native_entry__() + 56384
    _subr__validate_utf8_fast = __native_entry__() + 57328
    _subr__value              = __native_entry__() + 14592
    _subr__vnumber            = __native_entry__() + 18560
    _subr__vsigned            = __native_entry__() + 26352
    _subr__vstring            = __native_entry__() + 16672
    _subr__vunsigned          = __native_entry__() + 26720
)

const (
    _stack__f32toa = 48
    _stack__f64toa = 56
    _stack__get_by_path = 304
    _stack__html_escape = 72
    _stack__i64toa = 16
    _stack__lspace = 8
    _stack__quote = 56
    _stack__skip_array = 120
    _stack__skip_number = 56
    _stack__skip_object = 120
    _stack__skip_one = 120
    _stack__skip_one_fast = 200
    _stack__u64toa = 8
    _stack__unquote = 72
    _stack__validate_one = 104
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 176
    _stack__value = 184
    _stack__vnumber = 128
    _stack__vsigned = 16
    _stack__vstring = 72
    _stack__vunsigned = 16
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
    _ = _subr__validate_utf8
    _ = _subr__validate_utf8_fast
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
    _ = _stack__validate_utf8
    _ = _stack__validate_utf8_fast
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
