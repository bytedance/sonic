/**
 * Copyright 2026 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package issue_test

import (
	"testing"

	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/require"
)

func TestIssue732MoveAndCopyAtAnyDepth(t *testing.T) {
	root := ast.NewRaw(`{"foo":{"w":123,"h":345},"items":["a","b","c"]}`)

	require.NoError(t, root.MovePath(ast.Path{"foo"}, ast.Path{"bar"}))
	require.NoError(t, root.MovePath(ast.Path{"bar", "w"}, ast.Path{"dimensions", 2, "width"}))
	require.NoError(t, root.CopyPath(ast.Path{"bar"}, ast.Path{"snapshot"}))
	require.NoError(t, root.MovePath(ast.Path{"items", 0}, ast.Path{"items", 2}))

	// The copy must remain detached when the source subtree is changed.
	require.NoError(t, root.MovePath(ast.Path{"bar", "h"}, ast.Path{"height"}))

	raw, err := root.Raw()
	require.NoError(t, err)
	require.JSONEq(t, `{
		"bar": {},
		"dimensions": [null, null, {"width": 123}],
		"height": 345,
		"items": ["b", "c", "a"],
		"snapshot": {"h": 345}
	}`, raw)
}

func TestIssue732MoveToRootAndRejectDescendant(t *testing.T) {
	root := ast.NewRaw(`{"foo":{"w":123,"h":345}}`)
	require.NoError(t, root.MovePath(ast.Path{"foo"}, nil))

	raw, err := root.Raw()
	require.NoError(t, err)
	require.JSONEq(t, `{"w":123,"h":345}`, raw)

	require.Error(t, root.MovePath(nil, ast.Path{"child"}))
	require.Error(t, root.CopyPath(ast.Path{"missing"}, ast.Path{"copy"}))
	require.Error(t, root.MovePath(ast.Path{"w"}, ast.Path{-1}))
}
