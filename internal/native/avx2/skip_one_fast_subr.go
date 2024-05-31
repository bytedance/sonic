// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_one_fast = 128
)

const (
    _stack__skip_one_fast = 136
)

const (
    _size__skip_one_fast = 3404
)

var (
    _pcsp__skip_one_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {658, 136},
        {662, 48},
        {663, 40},
        {665, 32},
        {667, 24},
        {669, 16},
        {671, 8},
        {672, 0},
        {3404, 136},
    }
)

var _cfunc_skip_one_fast = []loader.CFunc{
    {"_skip_one_fast_entry", 0,  _entry__skip_one_fast, 0, nil},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
}
