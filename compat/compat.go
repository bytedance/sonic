// +build !go1.15 go1.19 !amd64 !linux,!darwin

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
	"encoding/json"
	"io"

	jsoniter "github.com/json-iterator/go"
)

// Marshal returns the JSON encoding string of v with faster config.
func Marshal(val interface{}) ([]byte, error) {
    return ConfigDefault.Marshal(val)
}

// MarshalString is like MarshalStd, except its output is string.
func MarshalString(val interface{}) (string, error) {
   return ConfigDefault.MarshalToString(val)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value
// pointed to by v, with faster config.
func Unmarshal(buf []byte, val interface{}) error {
    return ConfigDefault.Unmarshal(buf, val)
}

// UnmarshalStringStd is like Unmarshal, except buf is a string.
func UnmarshalString(buf string, val interface{}) error {
    return ConfigDefault.UnmarshalFromString(buf, val)
}

// MarshalStd returns the JSON encoding string of v, compatibly with encoding/json.
func MarshalStd(val interface{}) ([]byte, error) {
    return json.Marshal(val)
}

// MarshalStringStd is like MarshalStd, except its output is string.
func MarshalStringStd(val interface{}) (string, error) {
    buf, err := json.Marshal(val)
    return mem2Str(buf), err
}

// UnmarshalStd parses the JSON-encoded data and stores the result in the value
// pointed to by v, compatibly with encoding/json.
func UnmarshalStd(buf []byte, val interface{}) error {
    return json.Unmarshal(buf, val)
}

// UnmarshalStringStd is like Unmarshal, except buf is a string.
func UnmarshalStringStd(buf string, val interface{}) error {
    return json.Unmarshal(str2Mem(buf), val)
}

type frozenConfig struct {
    Config
    ja jsoniter.API
}

func (cfg Config) Froze() API {
    api := &frozenConfig{Config:cfg}
    jcfg := jsoniter.Config{}
    if cfg.EscapeHTML {
        jcfg.EscapeHTML = cfg.EscapeHTML
    }
    if cfg.SortMapKeys {
        jcfg.SortMapKeys = cfg.SortMapKeys
    }
    if cfg.UseNumber {
        jcfg.UseNumber = cfg.UseNumber
    }
    if cfg.DisallowUnknownFields {
        jcfg.DisallowUnknownFields = cfg.DisallowUnknownFields
    }
    jcfg.ValidateJsonRawMessage = true
    api.ja = jcfg.Froze()
    return api
}

func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
    return cfg.ja.Marshal(val)
}

func (cfg *frozenConfig) MarshalToString(val interface{}) (string, error) {
    return cfg.ja.MarshalToString(val)
}

func (cfg *frozenConfig) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
    return cfg.ja.MarshalIndent(v, prefix, indent)
}

func (cfg *frozenConfig) UnmarshalFromString(str string, v interface{}) error {
    return cfg.ja.UnmarshalFromString(str, v)
}

func (cfg *frozenConfig) Unmarshal(data []byte, v interface{}) error {
    return cfg.ja.Unmarshal(data, v)
}

func (cfg *frozenConfig) NewEncoder(writer io.Writer) Encoder {
    return cfg.ja.NewEncoder(writer)
}

func (cfg *frozenConfig) NewDecoder(reader io.Reader) Decoder {
    return cfg.ja.NewDecoder(reader)
}

func (cfg *frozenConfig) Valid(data []byte) bool {
    return cfg.ja.Valid(data)
}