//go:build amd64
// +build amd64

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate make
package sonic

import (
	"io"
	"reflect"

	"github.com/bytedance/sonic/ast"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/option"
)

func checkTrailings(buf string, pos int) error {
	/* skip all the trailing spaces */
	if pos != len(buf) {
		for pos < len(buf) && (types.SPACE_MASK&(1<<buf[pos])) != 0 {
			pos++
		}
	}

	/* then it must be at EOF */
	if pos == len(buf) {
		return nil
	}

	/* junk after JSON value */
	return decoder.SyntaxError{
		Src:  buf,
		Pos:  pos,
		Code: types.ERR_INVALID_CHAR,
	}
}

type frozenConfig struct {
	Config
	encoderOpts encoder.Options
	decoderOpts decoder.Options
}

// Froze convert the Config to API
func (cfg Config) Froze() API {
	api := &frozenConfig{Config: cfg}

	// configure encoder options:
	if cfg.EscapeHTML {
		api.encoderOpts |= encoder.EscapeHTML
	}
	if cfg.SortMapKeys {
		api.encoderOpts |= encoder.SortMapKeys
	}
	if cfg.CompactMarshaler {
		api.encoderOpts |= encoder.CompactMarshaler
	}
	if cfg.NoQuoteTextMarshaler {
		api.encoderOpts |= encoder.NoQuoteTextMarshaler
	}
	if cfg.NoNullSliceOrMap {
		api.encoderOpts |= encoder.NoNullSliceOrMap
	}

	// configure decoder options:
	if cfg.UseInt64 {
		api.decoderOpts |= decoder.OptionUseInt64
	}
	if cfg.UseNumber {
		api.decoderOpts |= decoder.OptionUseNumber
	}
	if cfg.DisallowUnknownFields {
		api.decoderOpts |= decoder.OptionDisableUnknown
	}
	if cfg.CopyString {
		api.decoderOpts |= decoder.OptionCopyString
	}
	if cfg.ValidateString {
		api.decoderOpts |= decoder.OptionValidateString
	}
	return api
}

// Marshal is implemented by sonic
func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
	return encoder.Encode(val, cfg.encoderOpts)
}

// MarshalToString is implemented by sonic
func (cfg *frozenConfig) MarshalToString(val interface{}) (string, error) {
	buf, err := encoder.Encode(val, cfg.encoderOpts)
	return rt.Mem2Str(buf), err
}

// MarshalIndent is implemented by sonic
func (cfg *frozenConfig) MarshalIndent(val interface{}, prefix, indent string) ([]byte, error) {
	return encoder.EncodeIndented(val, prefix, indent, cfg.encoderOpts)
}

// UnmarshalFromString is implemented by sonic
func (cfg *frozenConfig) UnmarshalFromString(buf string, val interface{}) error {
	dec := decoder.NewDecoder(buf)
	dec.SetOptions(cfg.decoderOpts)
	err := dec.Decode(val)
	pos := dec.Pos()

	/* check for errors */
	if err != nil {
		return err
	}
	return checkTrailings(buf, pos)
}

// Unmarshal is implemented by sonic
func (cfg *frozenConfig) Unmarshal(buf []byte, val interface{}) error {
	return cfg.UnmarshalFromString(string(buf), val)
}

// NewEncoder is implemented by sonic
func (cfg *frozenConfig) NewEncoder(writer io.Writer) Encoder {
	enc := encoder.NewStreamEncoder(writer)
	enc.Opts = cfg.encoderOpts
	return enc
}

// NewDecoder is implemented by sonic
func (cfg *frozenConfig) NewDecoder(reader io.Reader) Decoder {
	dec := decoder.NewStreamDecoder(reader)
	dec.SetOptions(cfg.decoderOpts)
	return dec
}

// Valid is implemented by sonic
func (cfg *frozenConfig) Valid(data []byte) bool {
	ok, _ := encoder.Valid(data)
	return ok
}

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
	if err := encoder.Pretouch(vt, opts...); err != nil {
	    return err
	} else if err := decoder.Pretouch(vt, opts...); err != nil {
		return err
	} else {
		return nil
	}
}

// Get searches the given path json,
// and returns its representing ast.Node.
//
// Each path arg must be integer or string:
//     - Integer means searching current node as array
//     - String means searching current node as object
func Get(src []byte, path ...interface{}) (ast.Node, error) {
	return GetFromString(string(src), path...)
}

// GetFromString is same with Get except src is string,
// which can reduce unnecessary memory copy.
func GetFromString(src string, path ...interface{}) (ast.Node, error) {
	return ast.NewSearcher(src).GetByPath(path...)
}
