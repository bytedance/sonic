package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

type Function = func()

type Unsupported struct {
	Functions []Function
}

type StructWithUnsupported struct {
	Foo *Unsupported `json:"foo"`
	Bar *Unsupported `json:"bar,omitempty"`
}

type Foo2 struct {
	A int
	B *chan int
}

type MockContext struct {
	*Foo2
}

func TestIssue491_MarshalUnsupportedType(t *testing.T) {
	// Wrapper a unbale serde type
	tests := []struct {
		value    interface{}
		wantErr  bool
		wantJSON string
	}{
		{value: map[string]*Function{}, wantJSON: `{}`},
		{value: map[*Function]*Function{}, wantErr: true},
		{value: map[string]Function{}, wantJSON: `{}`},
		{value: []Function{}, wantJSON: `[]`},
		{value: StructWithUnsupported{}, wantJSON: `{"foo":null}`},
		{value: struct {
			Foo *int
		}{}, wantJSON: `{"Foo":null}`},
		{value: struct {
			Foo Function
		}{}, wantErr: true},
		{value: chan int(nil), wantErr: true},
		{value: new(MockContext), wantJSON: `{}`},
	}
	for _, tt := range tests {
		if stdUsesJSONV2 {
			requireSonicMarshalSnapshot(t, sonic.ConfigDefault, tt.value, tt.wantErr, tt.wantJSON)
			continue
		}
		assertMarshal(t, sonic.ConfigDefault, tt.value)
	}
}

func TestIssue491_UnmarshalUnsupported(t *testing.T) {
	type Test struct {
		data  string
		value interface{}
	}

	tests := []struct {
		cas              unmTestCase
		wantUnmarshalErr bool
		wantMarshalErr   bool
		wantJSON         string
		check            func(*testing.T, interface{})
	}{
		{
			cas: unmTestCase{
				name:  "unsupported type slice",
				data:  []byte("null"),
				newfn: func() interface{} { return new([]Function) },
			},
			wantJSON: `null`,
			check: func(t *testing.T, value interface{}) {
				require.Nil(t, *value.(*[]Function))
			},
		},
		{
			cas: unmTestCase{
				name:  "unsupported type",
				data:  []byte("[null, null]"),
				newfn: func() interface{} { return new([]chan int) },
			},
			wantMarshalErr: true,
			check: func(t *testing.T, value interface{}) {
				got := *value.(*[]chan int)
				require.Len(t, got, 2)
				require.Nil(t, got[0])
				require.Nil(t, got[1])
			},
		},
		{
			cas: unmTestCase{
				name: "unsupported type in struct",
				data: []byte("{\"foo\": null}"),
				newfn: func() interface{} {
					return new(struct {
						Foo Function
					})
				},
			},
			wantMarshalErr: true,
			check: func(t *testing.T, value interface{}) {
				got := value.(*struct {
					Foo Function
				})
				require.Nil(t, got.Foo)
			},
		},
		{
			cas: unmTestCase{
				name: "unsupported type in map key should be error",
				data: []byte("null"),
				newfn: func() interface{} {
					return map[chan int]Function{}
				},
			},
			wantUnmarshalErr: true,
			wantMarshalErr:   true,
			check: func(t *testing.T, value interface{}) {
				require.Empty(t, value.(map[chan int]Function))
			},
		},
	}
	for _, tt := range tests {
		if stdUsesJSONV2 {
			value := requireSonicUnmarshalSnapshot(
				t,
				sonic.ConfigDefault,
				tt.cas,
				tt.wantUnmarshalErr,
				tt.wantMarshalErr,
				tt.wantJSON,
			)
			tt.check(t, value)
			continue
		}
		assertUnmarshal(t, sonic.ConfigDefault, tt.cas)
	}
}
