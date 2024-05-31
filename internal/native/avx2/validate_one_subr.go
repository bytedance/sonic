// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__validate_one = 144
)

const (
    _stack__validate_one = 160
)

const (
    _size__validate_one = 10216
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
        {9556, 160},
        {9560, 48},
        {9561, 40},
        {9563, 32},
        {9565, 24},
        {9567, 16},
        {9569, 8},
        {9570, 0},
        {10216, 160},
    }
)

var _cfunc_validate_one = []loader.CFunc{
    {"_validate_one_entry", 0,  _entry__validate_one, 0, nil},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
}
