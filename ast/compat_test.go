package ast

import (
    `testing`

    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/require`
    `github.com/tidwall/gjson`
)

func TestNotFoud(t *testing.T) {
    data := `{}`

    ia := jsoniter.Get([]byte(data), "b")
    require.Error(t, ia.LastError())
    require.Equal(t, false, ia.ToBool())

    ga := gjson.GetBytes([]byte(data), "b")
    require.True(t, ga.Type == gjson.Null)
    require.Equal(t, false, ga.Bool())

    sa, err := NewSearcher(data).GetByPath("b")
    require.True(t, sa.Type() == V_NONE)
    require.Error(t, err)
    sv, err := sa.Bool()
    require.Error(t, err)
    require.Equal(t, false, sv)
}

func TestNull(t *testing.T) {
    data := `{"b": null}`

    ia := jsoniter.Get([]byte(data), "b")
    require.NoError(t, ia.LastError())
    require.Equal(t, false, ia.ToBool())

    ga := gjson.GetBytes([]byte(data), "b")
    require.True(t, ga.Type == gjson.Null)
    require.Equal(t, false, ga.Bool())

    sa, err := NewSearcher(data).GetByPath("b")
    require.True(t, sa.Type() == V_NULL)
    require.NoError(t, err)
    sv, err := sa.Bool()
    require.NoError(t, err)
    require.Equal(t, false, sv)
}