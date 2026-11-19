package optdec

import (
	"encoding"
	"testing"

	"github.com/stretchr/testify/require"
)

type textSlice []string
type textMap map[string]string

func (*textSlice) UnmarshalText([]byte) error { panic("unexpected UnmarshalText") }
func (*textMap) UnmarshalText([]byte) error   { panic("unexpected UnmarshalText") }

func TestTextUnmarshalerNullContainers(t *testing.T) {
	s := textSlice{"old"}
	dst := struct {
		Slice textSlice
		Map   textMap
		Iface encoding.TextUnmarshaler
	}{s, textMap{"key": "old"}, &s}

	input := `{"Slice":null,"Map":null,"Iface":null}`
	pos := 0
	require.NoError(t, Decode(&input, &pos, 0, &dst))
	require.Nil(t, dst.Slice)
	require.Zero(t, len(dst.Slice))
	require.Zero(t, cap(dst.Slice))
	require.Nil(t, dst.Map)
	require.Nil(t, dst.Iface)
}
