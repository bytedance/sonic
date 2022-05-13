// +build go1.15,!go1.19,amd64,linux go1.15,!go1.19,amd64,darwin

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

package compat

import (
	"io"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
)

// Marshal returns the JSON encoding string of v, with faster config.
func Marshal(val interface{}) ([]byte, error) {
    return ConfigDefault.Marshal(val)
}

// MarshalString is like Marshal, except its output is string.
func MarshalString(val interface{}) (string, error) {
    return ConfigDefault.MarshalToString(val)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v, with faster config.
func Unmarshal(buf []byte, val interface{}) error {
    return ConfigDefault.Unmarshal(buf, val)
}

// UnmarshalString is like Unmarshal, except buf is a string.
func UnmarshalString(buf string, val interface{}) error {
    return ConfigDefault.UnmarshalFromString(buf, val)
}

// MarshalStd returns the JSON encoding string of v, compatibly with encoding/json.
func MarshalStd(val interface{}) ([]byte, error) {
    return ConfigStd.Marshal(val)
}

// MarshalStringStd is like MarshalStd, except its output is string.
func MarshalStringStd(val interface{}) (string, error) {
    return ConfigStd.MarshalToString(val)
}

// UnmarshalStd parses the JSON-encoded data and stores the result in the value
// pointed to by v, compatibly with encoding/json.
func UnmarshalStd(buf []byte, val interface{}) error {
    return ConfigStd.Unmarshal(buf, val)
}

// UnmarshalStringStd is like UnmarshalStd, except buf is a string.
func UnmarshalStringStd(buf string, val interface{}) error {
    return ConfigStd.UnmarshalFromString(buf, val)
}

func checkTrailings(buf string, pos int) error {
    /* skip all the trailing spaces */
    if pos != len(buf) {
        for pos < len(buf) && (types.SPACE_MASK & (1 << buf[pos])) != 0 {
            pos++
        }
    }

    /* then it must be at EOF */
    if pos == len(buf) {
        return nil
    }

    /* junk after JSON value */
    return decoder.SyntaxError {
        Src  : buf,
        Pos  : pos,
        Code : types.ERR_INVALID_CHAR,
    }
}

type frozenConfig struct {
    Config
    encoderOpts encoder.Options
    decoderOpts decoder.Options
}

func (cfg Config) Froze() API {
    api := &frozenConfig{Config: cfg}
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
    return api
}

func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
    return encoder.Encode(val, cfg.encoderOpts)
}

func (cfg *frozenConfig) MarshalToString(val interface{}) (string, error) {
    buf, err := encoder.Encode(val, cfg.encoderOpts)
    return mem2Str(buf), err
}

func (cfg *frozenConfig) MarshalIndent(val interface{}, prefix, indent string) ([]byte, error) {
    return encoder.EncodeIndented(val, prefix, indent, cfg.encoderOpts)
}

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

func (cfg *frozenConfig) Unmarshal(buf []byte, val interface{}) error {
    return cfg.UnmarshalFromString(string(buf), val)
}

func (cfg *frozenConfig) NewEncoder(writer io.Writer) Encoder {
    enc := encoder.NewStreamEncoder(writer)
    enc.Opts = cfg.encoderOpts
    return enc
}

func (cfg *frozenConfig) NewDecoder(reader io.Reader) Decoder {
    dec := decoder.NewStreamDecoder(reader)
    dec.SetOptions(cfg.decoderOpts)
    return dec
}

func (cfg *frozenConfig) Valid(data []byte) bool {
    ok, _ := encoder.Valid(data)
    return ok
}