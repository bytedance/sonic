package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"encoding/json"
)

var _ = json.Marshal

var _ = sonic.Marshal

func TestIssue758_UnmarshalIntoAnyPointer(t *testing.T) {
	for _, cas := range []unmTestCase {
		{
			name: "non-nil typed pointer",
			data: []byte(`["one","2"]`),
			newfn: func() interface{} {
				a := []string{}
				var aPtr interface{} = &a
				b := interface{}(&aPtr)
				return &b
			},
		},
		{
			name: "nil typed pointer",
			data: []byte(`["one","2"]`),
			newfn: func() interface{} {
				var aPtr interface{} = (*[]string)(nil)
				b := interface{}(&aPtr)
				return &b
			},
		},
		{
			name: "non-nil eface pointer recursive1",
			data: []byte(`{"a": "b"}`),
			newfn: func() interface{} {
				var v interface{}
				v = &v
				return v
			},
		},
		// TODO: the case is also failed for encoding/json
		// {
		// 	name: "non-nil eface pointer recursive2",
		// 	data: []byte(`{"a": "b"}`),
		// 	newfn: func() interface{} {
		// 		var v interface{}
		// 		var v1 = &v
		// 		v = &v1
		// 		return v
		// 	},
		// },
		{
			name: "non-nil eface pointer",
			data: []byte(`{"a": "b"}`),
			newfn: func() interface{} {
				var v1 interface{} = & struct {
					A string `json:"a"`
					B string `json:"b"`
				} {
					A: "c",
					B: "d",
				}
				var v = (*interface{})(&v1)
				return v
			},
		},
		{
			name: "nil eface pointer",
			data: []byte(`{"a": "b"}`),
			newfn: func() interface{} {
				var v interface{}
				v = (*interface{})(nil)
				return v
			},
		},
		{
			name: "non-nil iface pointer",
			data: []byte(`{"id": "2"}`),
			newfn: func() interface{} {
				var a MockEface = &fooEface{}
				var aPtr interface{} = &a
				b := interface{}(&aPtr)
				return &b
			},
		},
		{
			name: "root nil iface pointer shoule be error",
			data: []byte(`{"id": "2"}`),
			newfn: func() interface{} {
				var aPtr interface{} = (*MockEface)(nil)
				return aPtr
			},
		},
		{
			name: "nil iface pointer to be eface",
			data: []byte(`{"id": "2"}`),
			newfn: func() interface{} {
				var aPtr interface{} = (*MockEface)(nil)
				var a interface{} = &aPtr
				return a
			},
		},
		{
			name: "iface type",
			data: []byte(`{"id": "2"}`),
			newfn: func() interface{} {
				var a MockEface = fooEface3{}
				var aPtr interface{} = &a
				b := interface{}(&aPtr)
				return &b
			},
		},
	} {
		t.Run(cas.name, func(t *testing.T) {
			assertUnmarshal(t, sonic.ConfigDefault, cas)
		})
	}
}


type MockEface interface {
	MyMock()
}

type fooEface struct {
	Id *string `json:"id"`
}

func (self *fooEface) MyMock() {

}

type fooEface3 struct {
	Id *string `json:"id"`
}

func (self fooEface3) MyMock() {

}


