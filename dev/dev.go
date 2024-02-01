package dev

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/dev/decoder"
)

func UnmarshalString(json string, val interface{}) error {
	dec := decoder.NewDecoder(json)
	err := dec.Decode(val)
	return err
}

func Unmarshal(json []byte, val interface{}) error {
	return UnmarshalString(string(json), val)
}

// Marshal returns the JSON encoding bytes of v.
func Marshal(val interface{}) ([]byte, error) {
	return sonic.ConfigDefault.Marshal(val)
}

// MarshalString returns the JSON encoding string of v.
func MarshalString(val interface{}) (string, error) {
	out, err := sonic.ConfigDefault.Marshal(val)
	return string(out), err
}
