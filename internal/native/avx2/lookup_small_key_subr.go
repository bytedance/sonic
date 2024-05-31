// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__lookup_small_key = 48
)

const (
    _stack__lookup_small_key = 88
)

const (
    _size__lookup_small_key = 876
)

var (
    _pcsp__lookup_small_key = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {707, 88},
        {711, 48},
        {712, 40},
        {714, 32},
        {716, 24},
        {718, 16},
        {720, 8},
        {721, 0},
        {876, 88},
    }
)

var _cfunc_lookup_small_key = []loader.CFunc{
    {"_lookup_small_key_entry", 0,  _entry__lookup_small_key, 0, nil},
    {"_lookup_small_key", _entry__lookup_small_key, _size__lookup_small_key, _stack__lookup_small_key, _pcsp__lookup_small_key},
}
