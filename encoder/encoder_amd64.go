// +build amd64,go1.15,!go1.21

package encoder

import (
    `github.com/bytedance/sonic/internal/encoder`
)

// func Encode(val interface{}, opts Options) ([]byte, error)
// func EncodeIndented(val interface{}, prefix string, indent string, opts Options) ([]byte, error)
// func EncodeInto(buf *[]byte, val interface{}, opts Options) error
// func HTMLEscape(dst []byte, src []byte) []byte
// func Pretouch(vt reflect.Type, opts ...option.CompileOption) error
// func Quote(s string) string
// func Valid(data []byte) (ok bool, start int)
// type Encoder
// func (self *Encoder) Encode(v interface{}) ([]byte, error)
// func (self *Encoder) SetCompactMarshaler(f bool)
// func (self *Encoder) SetEscapeHTML(f bool)
// func (enc *Encoder) SetIndent(prefix, indent string)
// func (self *Encoder) SetNoQuoteTextMarshaler(f bool)
// func (self *Encoder) SetValidateString(f bool)
// func (self *Encoder) SortKeys() *Encoder
// type Options
// type StreamEncoder
// func NewStreamEncoder(w io.Writer) *StreamEncoder
// func (enc *StreamEncoder) Encode(val interface{}) (err error)

var (
    Encode = encoder.Encode
    EncodeIndented = encoder.EncodeIndented
    EncodeInto = encoder.EncodeInto
    HTMLEscape = encoder.HTMLEscape
    Pretouch = encoder.Pretouch
    Quote = encoder.Quote
    Valid = encoder.Valid
)

type Encoder = encoder.Encoder

type Options = encoder.Options

const (
	// SortMapKeys indicates that the keys of a map needs to be sorted
	// before serializing into JSON.
	// WARNING: This hurts performance A LOT, USE WITH CARE.
	SortMapKeys Options = encoder.SortMapKeys

	// EscapeHTML indicates encoder to escape all HTML characters
	// after serializing into JSON (see https://pkg.go.dev/encoding/json#HTMLEscape).
	// WARNING: This hurts performance A LOT, USE WITH CARE.
	EscapeHTML Options = encoder.EscapeHTML

	// CompactMarshaler indicates that the output JSON from json.Marshaler
	// is always compact and needs no validation
	CompactMarshaler Options = encoder.CompactMarshaler

	// NoQuoteTextMarshaler indicates that the output text from encoding.TextMarshaler
	// is always escaped string and needs no quoting
	NoQuoteTextMarshaler Options = encoder.NoQuoteTextMarshaler

	// NoNullSliceOrMap indicates all empty Array or Object are encoded as '[]' or '{}',
	// instead of 'null'
	NoNullSliceOrMap Options = encoder.NoNullSliceOrMap

	// ValidateString indicates that encoder should validate the input string
	// before encoding it into JSON.
	ValidateString Options = encoder.ValidateString

	// CompatibleWithStd is used to be compatible with std encoder.
	CompatibleWithStd Options = encoder.CompatibleWithStd
)

type StreamEncoder = encoder.StreamEncoder

var NewStreamEncoder = encoder.NewStreamEncoder