// +build amd64,go1.15,!go1.21

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

package encoder

import (
    `github.com/bytedance/sonic/internal/encoder`
)

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