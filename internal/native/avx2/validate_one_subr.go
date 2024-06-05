// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__validate_one = 480
)

const (
    _stack__validate_one = 112
)

const (
    _size__validate_one = 10664
)

var (
    _pcsp__validate_one = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {10225, 112},
        {10229, 48},
        {10230, 40},
        {10232, 32},
        {10234, 24},
        {10236, 16},
        {10238, 8},
        {10242, 0},
        {10664, 112},
    }
)

var _cfunc_validate_one = []loader.CFunc{
    {"_validate_one_entry", 0,  _entry__validate_one, 0, nil},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
}
