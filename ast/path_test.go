/*
 * Copyright 2026 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeMovePath(t *testing.T) {
	tests := []struct {
		name string
		json string
		src  Path
		dst  Path
		want string
	}{
		{
			name: "rename object",
			json: `{"foo":{"w":123,"h":345}}`,
			src:  Path{"foo"},
			dst:  Path{"bar"},
			want: `{"bar":{"w":123,"h":345}}`,
		},
		{
			name: "replace root",
			json: `{"foo":{"w":123,"h":345}}`,
			src:  Path{"foo"},
			dst:  nil,
			want: `{"w":123,"h":345}`,
		},
		{
			name: "create destination parents",
			json: `{"foo":{"w":123}}`,
			src:  Path{"foo", "w"},
			dst:  Path{"dimensions", 2, "width"},
			want: `{"foo":{},"dimensions":[null,null,{"width":123}]}`,
		},
		{
			name: "reorder lazy array",
			json: `{"items":["a","b","c"]}`,
			src:  Path{"items", 0},
			dst:  Path{"items", 2},
			want: `{"items":["b","c","a"]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := NewRaw(test.json)
			require.NoError(t, root.MovePath(test.src, test.dst))
			raw, err := root.Raw()
			require.NoError(t, err)
			require.JSONEq(t, test.want, raw)
		})
	}
}

func TestNodeCopyPathIsDetached(t *testing.T) {
	root := NewRaw(`{"foo":{"nested":{"value":1}}}`)
	require.NoError(t, root.CopyPath(Path{"foo"}, Path{"copy"}))

	require.NoError(t, root.MovePath(Path{"foo", "nested", "value"}, Path{"moved"}))

	raw, err := root.Raw()
	require.NoError(t, err)
	require.JSONEq(t, `{
		"foo":{"nested":{}},
		"copy":{"nested":{"value":1}},
		"moved":1
	}`, raw)
}

func TestNodePathErrorsAndNoOp(t *testing.T) {
	root := NewRaw(`{"foo":{"value":1}}`)

	require.NoError(t, root.MovePath(Path{"foo"}, Path{"foo"}))
	require.Error(t, root.MovePath(Path{"foo"}, Path{"foo", "child"}))
	require.Error(t, root.MovePath(Path{"missing"}, Path{"value"}))
	require.Error(t, root.CopyPath(Path{"foo"}, Path{-1}))
	require.Error(t, root.CopyPath(Path{"foo"}, Path{1.5}))

	raw, err := root.Raw()
	require.NoError(t, err)
	require.JSONEq(t, `{"foo":{"value":1}}`, raw)
}
