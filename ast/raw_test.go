package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestRawNode(t *testing.T) {
	_, err := NewSearcher(` { ] `).GetRawByPath()
	require.Error(t, err)
	d1 := ` {"a":1,"b":[1,1,1],"c":{"d":1,"e":1,"f":1}} `
	root, err := NewSearcher(d1).GetRawByPath()
	require.NoError(t, err)
	require.Equal(t, len(d1)-2, len(root.js))
	v1, err := root.GetByPath("a").Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v1)
	v2, err := root.GetByPath("b", 1).Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v2)
	v3, err := root.GetByPath("c", "f").Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v3)
}