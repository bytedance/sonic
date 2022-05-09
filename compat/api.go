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

import "io"

 type Config struct {
	EscapeHTML                    bool
	SortMapKeys                   bool
	CompactMarshaler              bool
	NoQuoteTextMarshaler          bool
	UseInt64                      bool
	UseNumber                     bool
	UseUnicodeErrors              bool
	DisallowUnknownFields         bool
	CopyString                    bool
}

var (
	ConfigDefault = Config{}.Froze()

    ConfigStd = Config{
		EscapeHTML : true,
		SortMapKeys: true,
		CompactMarshaler: true,
		CopyString : true,
	}.Froze()

    ConfigFastest = Config{
		NoQuoteTextMarshaler: false,
	}.Froze()
)


type API interface {
	MarshalToString(v interface{}) (string, error)
	Marshal(v interface{}) ([]byte, error)
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	UnmarshalFromString(str string, v interface{}) error
	Unmarshal(data []byte, v interface{}) error
	// Get(data []byte, path ...interface{}) Value
	NewEncoder(writer io.Writer) Encoder
	NewDecoder(reader io.Reader) Decoder
	Valid(data []byte) bool
}

type Encoder interface {
	Encode(val interface{}) ([]byte, error)
}

type Decoder interface {
	Decode(buf []byte, val interface{}) (error)
}