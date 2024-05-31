// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vsigned = 0
)

const (
    _stack__vsigned = 16
)

const (
    _size__vsigned = 336
)

var (
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {119, 16},
        {120, 8},
        {121, 0},
        {132, 16},
        {133, 8},
        {134, 0},
        {276, 16},
        {277, 8},
        {278, 0},
        {282, 16},
        {283, 8},
        {284, 0},
        {322, 16},
        {323, 8},
        {324, 0},
        {332, 16},
        {333, 8},
        {336, 0},
    }
)

var _cfunc_vsigned = []loader.CFunc{
    {"_vsigned_entry", 0,  _entry__vsigned, 0, nil},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
}
