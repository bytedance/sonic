// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

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
    _size__lookup_small_key = 892
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
        {877, 88},
        {881, 48},
        {882, 40},
        {884, 32},
        {886, 24},
        {888, 16},
        {890, 8},
        {892, 0},
    }
)

var _cfunc_lookup_small_key = []loader.CFunc{
    {"_lookup_small_key_entry", 0,  _entry__lookup_small_key, 0, nil},
    {"_lookup_small_key", _entry__lookup_small_key, _size__lookup_small_key, _stack__lookup_small_key, _pcsp__lookup_small_key},
}
