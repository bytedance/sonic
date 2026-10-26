package issue_test

import (
	"fmt"
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

func TestIssue1004PairLiteralReplacement(t *testing.T) {
	for _, size := range []int{1, 17} {
		for _, key := range []string{"a", ""} {
			for _, placeholder := range []bool{false, true} {
				t.Run(fmt.Sprintf("size=%d/key=%q/placeholder=%t", size, key, placeholder), func(t *testing.T) {
					pairs := []ast.Pair{ast.NewPair(key, ast.NewNull())}
					for i := 1; i < size; i++ {
						pairs = append(pairs, ast.NewPair(fmt.Sprintf("key%d", i), ast.NewNull()))
					}
					root := ast.NewObject(pairs)
					value := ast.NewNumber("2")
					if placeholder {
						value = ast.Node{}
					}
					*root.IndexPair(0) = ast.Pair{Key: key, Value: value}
					got := root.Get(key)
					require.NotNil(t, got)
					if placeholder {
						require.NoError(t, got.Add(ast.NewNull()))
					} else {
						n, err := got.Int64()
						require.NoError(t, err)
						require.Equal(t, int64(2), n)
					}
					existed, err := root.Set(key, ast.NewNumber("3"))
					require.NoError(t, err)
					require.True(t, existed)
					pairs[0] = ast.NewPair(key, ast.NewNumber("3"))
					expected := ast.NewObject(pairs)
					want, err := expected.MarshalJSON()
					require.NoError(t, err)
					data, err := root.MarshalJSON()
					require.NoError(t, err)
					require.Equal(t, string(want), string(data))
				})
			}
		}
	}
}
