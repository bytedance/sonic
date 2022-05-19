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
    `encoding/json`
    `io`

    jsoniter `github.com/json-iterator/go`
)

type frozenConfig struct {
    Config
    ja jsoniter.API
}

// Froze convert a Config to API
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
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.Marshal(val)
    }
    return cfg.ja.Marshal(val)
}

func (cfg *frozenConfig) MarshalToString(val interface{}) (string, error) {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        buf, err := json.Marshal(val)
        return mem2Str(buf), err    
    }
    return cfg.ja.MarshalToString(val)
}

func (cfg *frozenConfig) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.MarshalIndent(v, prefix, indent)
    }
    return cfg.ja.MarshalIndent(v, prefix, indent)
}

func (cfg *frozenConfig) UnmarshalFromString(str string, val interface{}) error {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.Unmarshal(str2Mem(str), val)
    }
    return cfg.ja.UnmarshalFromString(str, val)
}

func (cfg *frozenConfig) Unmarshal(data []byte, v interface{}) error {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.Unmarshal(data, v)
    }
    return cfg.ja.Unmarshal(data, v)
}

func (cfg *frozenConfig) NewEncoder(writer io.Writer) Encoder {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.NewEncoder(writer)
    }
    return cfg.ja.NewEncoder(writer)
}

func (cfg *frozenConfig) NewDecoder(reader io.Reader) Decoder {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.NewDecoder(reader)
    }
    return cfg.ja.NewDecoder(reader)
}

func (cfg *frozenConfig) Valid(data []byte) bool {
    if v := ConfigStd.(*frozenConfig); cfg == v {
        return json.Valid(data)
    }
    return cfg.ja.Valid(data)
}