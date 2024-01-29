// +build !amd64 !go1.16 go1.22

/*
* Copyright 2023 ByteDance Inc.
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

package decoder

import (
    `encoding/json`
    `io`
    `reflect`

    `github.com/bytedance/sonic/dev/decoder`
    `github.com/bytedance/sonic/option`
)

func init() {
     println("WARNING(decoder): you are using the test branch in aarch64")
}

// Decoder is the decoder context object
type Decoder = decoder.Decoder

// SyntaxError represents json syntax error
type SyntaxError = decoder.SyntaxError

// MismatchTypeError represents dismatching between json and object
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

var (
    // NewDecoder creates a new decoder instance.
    NewDecoder = decoder.NewDecoder
)

// Pretouch compiles vt ahead-of-time to avoid JIT compilation on-the-fly, in
// order to reduce the first-hit latency.
//
// Opts are the compile options, for example, "option.WithCompileRecursiveDepth" is
// a compile option to set the depth of recursive compile for the nested struct type.
func Pretouch(vt reflect.Type, opts ...option.CompileOption) error {
     return nil
}

// TODO: replace stream decoder with dev/decoder
type StreamDecoder = json.Decoder

// NewStreamDecoder adapts to encoding/json.NewDecoder API.
//
// NewStreamDecoder returns a new decoder that reads from r.
func NewStreamDecoder(r io.Reader) *StreamDecoder {
   return json.NewDecoder(r)
}