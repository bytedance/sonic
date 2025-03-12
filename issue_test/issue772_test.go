
package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue772_SkipIfaceType(t *testing.T) {
	for _, cas := range []unmTestCase {
		{
			name: "should skip non-ptr iface type",
			data: []byte(`{"id": {"id": "2"},"name": "name"}`),
			newfn: func() interface{} {
				obj := WrapperEface {
				}
				obj.Id = fooEface3{}
				return &obj
			},
		},
		{
			name: "should skip nil iface type",
			data: []byte(`{"id": {"id": "2"},"name": "name"}`),
			newfn: func() interface{} {
				obj := WrapperEface {
				}
				obj.Id = (*fooEface)(nil)
				return &obj
			},
		},
	} {
		t.Run(cas.name, func(t *testing.T) {
			assertUnmarshal(t, sonic.ConfigDefault, cas, true)
		})
	}
}



