package example

import (
	"bytes"
	"fmt"
	"strings"
	"github.com/bytedance/sonic"
)

// This example uses a Decoder to decode a stream of distinct JSON values.
func ExampleStreamDecoder() {
	var o =  map[string]interface{}{}
	var r = strings.NewReader(`{"a":"b"}{"1":"2"}`)
	var dec = sonic.ConfigDefault.NewDecoder(r)
	dec.Decode(&o)
	dec.Decode(&o)
	fmt.Printf("%+v", o)
	// Output:
	// map[1:2 a:b]
}


// This example uses a Encoder to encode streamingly.
func ExampleStreamEncoder() {
	var o1 = map[string]interface{}{
		"a": "b",
	}
	var o2 = 1
	var w = bytes.NewBuffer(nil)
	var enc = sonic.ConfigDefault.NewEncoder(w)
	enc.Encode(o1)
	enc.Encode(o2)
	fmt.Println(w.String())
	// Output:
	// {"a":"b"}
	// 1
}