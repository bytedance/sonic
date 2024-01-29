package dev

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/dev/decoder"
)

// func UnmarshalString_Bak(json string, val interface{}) error {
// 	ctx, err := internal.NewContext(json)
// 	if err != nil {
// 		return err
// 	}

// 	node := ctx.Dom.Root()
// 	ret, err := node.AsIface(&ctx)
// 	if err != nil {
// 		return err
// 	}

// 	rv := reflect.ValueOf(val)
// 	rv.Elem().Set(reflect.ValueOf(ret))
// 	ctx.Delete()
// 	return nil
// }

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
	return sonic.Marshal(val)
}

// MarshalString returns the JSON encoding string of v.
func MarshalString(val interface{}) (string, error) {
	out, err := sonic.Marshal(val)
	return string(out), err
}
