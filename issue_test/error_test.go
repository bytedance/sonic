package issue_test

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

type S struct {
	A int    `json:",omitempty"`
	B string `json:"B1,omitempty"`
	C float64
	D bool
	E uint8
	// F []byte // unmarshal []byte is different with encoding/json
	G interface{}
	H map[string]interface{}
	I map[string]string
	J []interface{}
	K []string
	L S1
	M *S1
	N *int
	O **int
    P int `json:",string"`
    Q float64 `json:",string"`
	R int `json:"-"`
	T struct {}
	U [2]int
	V uintptr
    W json.Number
    X json.RawMessage
    Y Marshaller 
	Z TextMarshaller
}


type S1 struct {
	A int
	B string
}

type Marshaller struct {
	v string
}

func (m *Marshaller) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.v)
}

func (m *Marshaller) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.v)
}

type TextMarshaller struct {
    v int
}

func (k *TextMarshaller) MarshalText() ([]byte, error) {
    return json.Marshal(k.v)
}

func (k *TextMarshaller)  UnmarshalText(data []byte) error {
    return json.Unmarshal(data, &k.v)
}

const data = `{"A":1,"B1":"2","C":3,"D":true,"E":4,"G":5,"H":{"a":6},"I":{"a":"b"},"J":[7,8],"K":["a","b"],"L":{"A":9,"B":"10"},"M":{"A":11,"B":"12"},"N":13,"O":14,"P":"15","Q":"16","R":17,"T":{},"U":[19,20],"V":21,"W":"22","X":"23","Y":"24","Z":"25"}`

func TestErrorUnmarshalInvalidJSON(t *testing.T) {
	var objj, objs, obji S
	errj := json.Unmarshal([]byte(data), &objj)
	erri := jsoniter.Unmarshal([]byte(data), &obji)
	errs := sonic.Unmarshal([]byte(data), &objs)
	require.Equal(t, errj, errs)
	require.Equal(t, errj, erri)
	require.Equal(t, objj, objs)
	require.Equal(t, objj, obji)

	fields := []string{"B"}
	for _, field := range fields {
		root, err := sonic.Get([]byte(data))
		require.Nil(t, err)
		_, err = root.Set(field, ast.NewRaw(`{]`))
		require.Nil(t, err)
		buf, err := root.MarshalJSON()
		require.Nil(t, err)

		var objj, objs, obji S
		errj := json.Unmarshal(buf, &objj)
		errs := sonic.Unmarshal(buf, &objs)
		erri := jsoniter.Unmarshal(buf, &obji)
		require.Equal(t, errj!=nil, erri!=nil, "jsoniter:%s", string(buf))
		require.Equal(t, errj!=nil, errs!=nil, "json:%s", string(buf))
		require.Equal(t, objj, obji, "jsoniter:%s", string(buf))
		require.Equal(t, objj, objs, "json:%s", string(buf))
	}
}

func TestErrorUnmarshalDismatchedJSON(t *testing.T) {
	var objj, objs, obji S
	errj := json.Unmarshal([]byte(data), &objj)
	erri := jsoniter.Unmarshal([]byte(data), &obji)
	errs := sonic.Unmarshal([]byte(data), &objs)
	require.Equal(t, errj, errs)
	require.Equal(t, errj, erri)
	require.Equal(t, objj, objs)
	require.Equal(t, objj, obji)

	fields := []string{"L"}
	for _, field := range fields {
		root, err := sonic.Get([]byte(data))
		require.Nil(t, err)
		_, err = root.Set(field, ast.NewRaw(`[1,2]`))
		require.Nil(t, err)
		buf, err := root.MarshalJSON()
		require.Nil(t, err)

		var objj, objs, obji S
		errj := json.Unmarshal(buf, &objj)
		errs := sonic.Unmarshal(buf, &objs)
		erri := jsoniter.Unmarshal(buf, &obji)
		require.Equal(t, errj!=nil, erri!=nil, "jsoniter:%s", string(buf))
		require.Equal(t, errj!=nil, errs!=nil, "json:%s", string(buf))
		require.Equal(t, objj, obji, "jsoniter:%s", string(buf))
		require.Equal(t, objj, objs, "json:%s", string(buf))
	}
}