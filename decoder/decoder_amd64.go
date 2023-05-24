// +build amd64,go1.15,!go1.21

package decoder

import (
    `github.com/bytedance/sonic/internal/decoder`
)

type Decoder = decoder.Decoder

type MismatchTypeError = decoder.MismatchTypeError

type Options = decoder.Options

const (
	OptionUseInt64         Options = decoder.OptionUseInt64
	OptionUseNumber        Options = decoder.OptionUseNumber
	OptionUseUnicodeErrors Options = decoder.OptionUseUnicodeErrors
	OptionDisableUnknown   Options = decoder.OptionDisableUnknown
	OptionCopyString       Options = decoder.OptionCopyString
	OptionValidateString   Options = decoder.OptionValidateString
)

type StreamDecoder = decoder.StreamDecoder

type SyntaxError = decoder.SyntaxError

var Pretouch = decoder.Pretouch

var Skip = decoder.Skip

var NewDecoder = decoder.NewDecoder

var NewStreamDecoder = decoder.NewStreamDecoder
