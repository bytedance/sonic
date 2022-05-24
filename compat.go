// +build !amd64

package sonic

import (
	"bytes"
	"encoding/json"
	"io"
)

type frozenConfig struct {
    Config
}

// Froze convert the Config to API
func (cfg Config) Froze() API {
    api := &frozenConfig{Config: cfg}
    return api
}

func (cfg *frozenConfig) marshalOptions(val interface{}, prefix, indent string) ([]byte, error) {
	w := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(cfg.EscapeHTML)
	enc.SetIndent(prefix, indent)
	err := enc.Encode(val)
	return w.Bytes(), err
}

// Marshal is implemented by sonic
func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
	if !cfg.EscapeHTML {
		return cfg.marshalOptions(val, "", "")
	}
	return json.Marshal(val)
}

// MarshalToString is implemented by sonic
func (cfg *frozenConfig) MarshalToString(val interface{}) (string, error) {
	out, err := cfg.Marshal(val)
	return string(out), err
}

// MarshalIndent is implemented by sonic
func (cfg *frozenConfig) MarshalIndent(val interface{}, prefix, indent string) ([]byte, error) {
	if !cfg.EscapeHTML {
		return cfg.marshalOptions(val, prefix, indent)
	}
    return json.MarshalIndent(val, prefix, indent)
}

// UnmarshalFromString is implemented by sonic
func (cfg *frozenConfig) UnmarshalFromString(buf string, val interface{}) error {
	r := bytes.NewBufferString(buf)
	dec := json.NewDecoder(r)
    if cfg.UseNumber {
		dec.UseNumber()
	}
	if cfg.DisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
    return dec.Decode(val)
}

// Unmarshal is implemented by sonic
func (cfg *frozenConfig) Unmarshal(buf []byte, val interface{}) error {
    return cfg.UnmarshalFromString(string(buf), val)
}

// NewEncoder is implemented by sonic
func (cfg *frozenConfig) NewEncoder(writer io.Writer) Encoder {
	enc := json.NewEncoder(writer)
	if !cfg.EscapeHTML {
		enc.SetEscapeHTML(cfg.EscapeHTML)
	}
    return enc
}

// NewDecoder is implemented by sonic
func (cfg *frozenConfig) NewDecoder(reader io.Reader) Decoder {
    dec := json.NewDecoder(reader)
    if cfg.UseNumber {
		dec.UseNumber()
	}
	if cfg.DisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
    return dec
}

// Valid is implemented by sonic
func (cfg *frozenConfig) Valid(data []byte) bool {
    return json.Valid(data)
}