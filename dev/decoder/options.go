package decoder

import (
	"github.com/bytedance/sonic/internal/native/types"
)

const (
	_F_use_int64       = 0
	_F_disable_urc     = 2
	_F_disable_unknown = 3
	_F_copy_string     = 4

	_F_use_number      = types.B_USE_NUMBER
	_F_validate_string = types.B_VALIDATE_STRING
	_F_allow_control   = types.B_ALLOW_CONTROL
)

type Options uint64

const (
	OptionUseInt64         Options = 1 << _F_use_int64       // 1
	OptionUseNumber        Options = 1 << _F_use_number      // 2
	OptionUseUnicodeErrors Options = 1 << _F_disable_urc     // 4
	OptionDisableUnknown   Options = 1 << _F_disable_unknown // 8
	OptionCopyString       Options = 1 << _F_copy_string     // 16
	OptionValidateString   Options = 1 << _F_validate_string
)

func (self *Decoder) SetOptions(opts Options) {
	if (opts&OptionUseNumber != 0) && (opts&OptionUseInt64 != 0) {
		panic("can't set OptionUseInt64 and OptionUseNumber both!")
	}
	self.opts = opts
}

// UseInt64 indicates the Decoder to unmarshal an integer into an interface{} as an
// int64 instead of as a float64.
func (self *Decoder) UseInt64() {
	self.opts |= 1 << _F_use_int64
	self.opts &^= 1 << _F_use_number
}

// UseNumber indicates the Decoder to unmarshal a number into an interface{} as a
// json.Number instead of as a float64.
func (self *Decoder) UseNumber() {
	self.opts &^= 1 << _F_use_int64
	self.opts |= 1 << _F_use_number
}

// UseUnicodeErrors indicates the Decoder to return an error when encounter invalid
// UTF-8 escape sequences.
func (self *Decoder) UseUnicodeErrors() {
	self.opts |= 1 << _F_disable_urc
}

// DisallowUnknownFields indicates the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (self *Decoder) DisallowUnknownFields() {
	self.opts |= 1 << _F_disable_unknown
}

// CopyString indicates the Decoder to decode string values by copying instead of referring.
func (self *Decoder) CopyString() {
	self.opts |= 1 << _F_copy_string
}

// ValidateString causes the Decoder to validate string values when decoding string value
// in JSON. Validation is that, returning error when unescaped control chars(0x00-0x1f) or
// invalid UTF-8 chars in the string value of JSON.
func (self *Decoder) ValidateString() {
	self.opts |= 1 << _F_validate_string
}
