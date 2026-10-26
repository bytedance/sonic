package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue758_UnmarshalIntoAnyPointer(t *testing.T) {
	snapshots := map[string]struct {
		wantErr  bool
		wantJSON string
	}{
		"non-nil typed pointer":                  {wantJSON: `["one","2"]`},
		"nil typed pointer":                      {wantJSON: `["one","2"]`},
		"non-nil eface pointer recursive1":       {wantJSON: `{"a":"b"}`},
		"non-nil eface pointer":                  {wantJSON: `{"a":"b","b":"d"}`},
		"nil eface pointer":                      {wantErr: true, wantJSON: `null`},
		"non-nil iface pointer":                  {wantJSON: `{"id":"2"}`},
		"root nil iface pointer shoule be error": {wantErr: true, wantJSON: `null`},
		"nil iface pointer to be eface":          {wantJSON: `{"id":"2"}`},
		"iface type should be error":             {wantErr: true, wantJSON: `null`},
		"non-nil iface type as struct field":     {wantJSON: `{"name":"name","id":{"id":"2"}}`},
		"null iface type as struct field":        {wantJSON: `{"name":"name","id":null}`},
	}
	for _, cas := range []unmTestCase{
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
				var v1 interface{} = &struct {
					A string `json:"a"`
					B string `json:"b"`
				}{
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
			name: "iface type should be error",
			data: []byte(`{"id": "2"}`),
			newfn: func() interface{} {
				var a MockEface = fooEface3{}
				var aPtr interface{} = &a
				b := interface{}(&aPtr)
				return &b
			},
		},
		{
			name: "non-nil iface type as struct field",
			data: []byte(`{"id": {"id": "2"},"name": "name"}`),
			newfn: func() interface{} {
				obj := WrapperEface{}
				foo := fooEface{}
				obj.Id = &foo
				return &obj
			},
		},
		{
			name: "null iface type as struct field",
			data: []byte(`{"name": "name", "id": null}`),
			newfn: func() interface{} {
				obj := WrapperEface{}
				a := "123"
				foo := fooEface{
					Id: &a,
				}
				obj.Id = &foo
				return &obj
			},
		},
	} {
		t.Run(cas.name, func(t *testing.T) {
			if stdUsesJSONV2 {
				value := cas.newfn()
				err := sonic.ConfigDefault.Unmarshal(cas.data, value)
				want, ok := snapshots[cas.name]
				if !ok {
					t.Fatalf("missing Sonic snapshot for %q", cas.name)
				}
				require.Equal(t, want.wantErr, err != nil, err)
				requireIssue758Target(t, cas.name, value)
				if cas.name == "null iface type as struct field" {
					return
				}
				requireSonicValueSnapshot(t, sonic.ConfigDefault, value, false, want.wantJSON)
				return
			}
			assertUnmarshal(t, sonic.ConfigDefault, cas, true)
		})
	}
}

func requireIssue758Target(t *testing.T, name string, value interface{}) {
	t.Helper()
	switch name {
	case "non-nil typed pointer":
		require.IsType(t, (*interface{})(nil), value)
		outer := value.(*interface{})
		require.IsType(t, (*interface{})(nil), *outer)
		middle := (*outer).(*interface{})
		require.IsType(t, (*[]string)(nil), *middle)
		got := (*middle).(*[]string)
		require.Equal(t, []string{"one", "2"}, *got)
	case "nil typed pointer":
		require.IsType(t, (*interface{})(nil), value)
		outer := value.(*interface{})
		require.IsType(t, (*interface{})(nil), *outer)
		middle := (*outer).(*interface{})
		require.IsType(t, []interface{}{}, *middle)
		got := (*middle).([]interface{})
		require.Equal(t, []interface{}{"one", "2"}, got)
	case "non-nil eface pointer recursive1":
		require.IsType(t, (*interface{})(nil), value)
		got := value.(*interface{})
		require.IsType(t, map[string]interface{}{}, *got)
		require.Equal(t, map[string]interface{}{"a": "b"}, (*got).(map[string]interface{}))
	case "non-nil eface pointer":
		require.IsType(t, (*interface{})(nil), value)
		got := value.(*interface{})
		wantType := (*struct {
			A string `json:"a"`
			B string `json:"b"`
		})(nil)
		require.IsType(t, wantType, *got)
		target := (*got).(*struct {
			A string `json:"a"`
			B string `json:"b"`
		})
		require.Equal(t, "b", target.A)
		require.Equal(t, "d", target.B)
	case "nil eface pointer":
		require.IsType(t, (*interface{})(nil), value)
		require.Nil(t, value.(*interface{}))
	case "non-nil iface pointer":
		require.IsType(t, (*interface{})(nil), value)
		outer := value.(*interface{})
		require.IsType(t, (*interface{})(nil), *outer)
		middle := (*outer).(*interface{})
		require.IsType(t, (*MockEface)(nil), *middle)
		iface := (*middle).(*MockEface)
		require.IsType(t, (*fooEface)(nil), *iface)
		target := (*iface).(*fooEface)
		require.NotNil(t, target.Id)
		require.Equal(t, "2", *target.Id)
	case "root nil iface pointer shoule be error":
		require.IsType(t, (*MockEface)(nil), value)
		require.Nil(t, value.(*MockEface))
	case "nil iface pointer to be eface":
		require.IsType(t, (*interface{})(nil), value)
		got := value.(*interface{})
		require.IsType(t, map[string]interface{}{}, *got)
		require.Equal(t, map[string]interface{}{"id": "2"}, (*got).(map[string]interface{}))
	case "iface type should be error":
		require.IsType(t, (*interface{})(nil), value)
		outer := value.(*interface{})
		require.IsType(t, (*interface{})(nil), *outer)
		middle := (*outer).(*interface{})
		require.IsType(t, (*MockEface)(nil), *middle)
		iface := (*middle).(*MockEface)
		require.IsType(t, fooEface3{}, *iface)
		target := (*iface).(fooEface3)
		require.Nil(t, target.Id)
	case "non-nil iface type as struct field":
		require.IsType(t, (*WrapperEface)(nil), value)
		got := value.(*WrapperEface)
		require.Equal(t, "name", got.Name)
		require.IsType(t, (*fooEface)(nil), got.Id)
		target := got.Id.(*fooEface)
		require.NotNil(t, target.Id)
		require.Equal(t, "2", *target.Id)
	case "null iface type as struct field":
		require.IsType(t, (*WrapperEface)(nil), value)
		got := value.(*WrapperEface)
		require.Equal(t, "name", got.Name)
		require.Nil(t, got.Id)
	default:
		t.Fatalf("missing direct target assertion for %q", name)
	}
}

type WrapperEface struct {
	Name string    `json:"name"`
	Id   MockEface `json:"id"`
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
