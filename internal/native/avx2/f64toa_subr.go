// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f64toa = 48
)

const (
    _stack__f64toa = 72
)

const (
    _size__f64toa = 5088
)

var (
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {4985, 72},
        {4989, 48},
        {4990, 40},
        {4992, 32},
        {4994, 24},
        {4996, 16},
        {4998, 8},
        {5002, 0},
        {5088, 72},
    }
)

var _cfunc_f64toa = []loader.CFunc{
    {"_f64toa_entry", 0,  _entry__f64toa, 0, nil},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
}
