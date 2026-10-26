package issue_test

import (
	"testing"

	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/require"
)

func TestIssue1004EmptyKeyAfterUnset(t *testing.T) {
	root := ast.NewRaw(`{"a":1,"":2}`)
	existed, err := root.Unset("a")
	require.NoError(t, err)
	require.True(t, existed)

	value, err := root.Get("").Int64()
	require.NoError(t, err)
	require.Equal(t, int64(2), value)

	existed, err = root.Set("", ast.NewNumber("3"))
	require.NoError(t, err)
	require.True(t, existed)
	data, err := root.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, `{"":3}`, string(data))

	existed, err = root.Unset("")
	require.NoError(t, err)
	require.True(t, existed)
	require.Nil(t, root.Get(""))
	existed, err = root.Set("", ast.NewNull())
	require.NoError(t, err)
	require.False(t, existed)
	data, err = root.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, `{"":null}`, string(data))
}
