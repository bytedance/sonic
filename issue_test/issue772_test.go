package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue772_SkipIfaceType(t *testing.T) {
	for _, tt := range []struct {
		cas   unmTestCase
		check func(*testing.T, *WrapperEface)
	}{
		{
			cas: unmTestCase{
				name: "should skip non-ptr iface type",
				data: []byte(`{"id": {"id": "2"},"name": "name"}`),
				newfn: func() interface{} {
					obj := WrapperEface{}
					obj.Id = fooEface3{}
					return &obj
				},
			},
			check: func(t *testing.T, got *WrapperEface) {
				require.Equal(t, "name", got.Name)
				require.IsType(t, fooEface3{}, got.Id)
				require.Nil(t, got.Id.(fooEface3).Id)
			},
		},
		{
			cas: unmTestCase{
				name: "should skip nil iface type",
				data: []byte(`{"id": {"id": "2"},"name": "name"}`),
				newfn: func() interface{} {
					obj := WrapperEface{}
					obj.Id = (*fooEface)(nil)
					return &obj
				},
			},
			check: func(t *testing.T, got *WrapperEface) {
				require.Equal(t, "name", got.Name)
				require.IsType(t, (*fooEface)(nil), got.Id)
				require.Nil(t, got.Id.(*fooEface))
			},
		},
	} {
		t.Run(tt.cas.name, func(t *testing.T) {
			if stdUsesJSONV2 {
				value := requireSonicUnmarshalSnapshot(
					t,
					sonic.ConfigDefault,
					tt.cas,
					true,
					false,
					`{"name":"name","id":null}`,
				)
				tt.check(t, value.(*WrapperEface))
				return
			}
			assertUnmarshal(t, sonic.ConfigDefault, tt.cas, true)
		})
	}
}
