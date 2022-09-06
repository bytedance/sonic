package ast

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestNotFoud(t *testing.T) {
	data := `{"a":null}`
	ia := jsoniter.Get([]byte(data), "b")
	require.Error(t, ia.LastError())
	require.Equal(t, false, ia.ToBool())

	ga := gjson.GetBytes([]byte(data), "b")
	require.True(t, ga.Type == gjson.Null)
	require.Equal(t, false, ga.Bool())
}