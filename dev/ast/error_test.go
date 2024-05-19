package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrIndexOutOfRange_Error(t *testing.T) {
	tests := []struct {
		name   string
		src    string
		path   []interface{}
		want   int
		err    string
	}{
		{
			name: "top",
			src: ` [ 1 ] `,
			path: []interface{}{1},
			want: 0,
			err: ErrNotExist.Error(),
		},
		{
			name: "second",
			src: `{ "1" : [ 1 ] }`,
			path: []interface{}{"1", 2, 3},
			want: 1,
			err: ErrNotExist.Error(),
		},
		{
			name: "thrid",
			src: `{ "1" : [ 1, [ 1 ] ] }`,
			path: []interface{}{"1", 1, 3},
			want: 2,
			err: ErrNotExist.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewRaw(tt.src)
			exist, err := n.SetByPath(false, "1", tt.path...)
			require.False(t, exist)
			require.Error(t, err)
			require.Equal(t, tt.err, err.Error())
			require.Equal(t, tt.want, err.(ErrIndexOutOfRange).Index)
		})
	}
}
