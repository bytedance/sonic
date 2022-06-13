// +build go1.18

/*
 * Copyright 2021 ByteDance Inc.
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

package sonic_fuzz

import (
    `testing`
    `github.com/bytedance/sonic`
    `github.com/bytedance/sonic/ast`
    `github.com/stretchr/testify/require`
)

func fuzzASTGetFromObject(t *testing.T, data []byte, m map[string]interface{}) {
    for k, expv := range(m) {
        node, err := sonic.Get(data, k)
        require.NoErrorf(t, err, "error in ast get key -> %s", k)
        assertAstNode(t, node, expv)
    }
}

func fuzzASTGetFromArray(t *testing.T, data []byte, a []interface{}) {
    var i = 0
    for ; i < len(a); i++ {
        node, err := sonic.Get(data, i)
        require.NoErrorf(t, err, "error in ast get index -> %s", i)
        assertAstNode(t, node, a[i])
    }
    _, err := sonic.Get(data, i)
    require.Errorf(t, err, "error in ast get index -> %s", i)
}

func assertAstNode(t *testing.T, node ast.Node, expv interface{}) {
    switch node.Type() {
        case ast.V_NULL: require.Nilf(t, expv, "wrong in ast null")
        case ast.V_TRUE: fallthrough
        case ast.V_FALSE: 
            gotv, err := node.Bool()
            require.NoErrorf(t, err, "error in ast get bool")
            require.Equalf(t, gotv, expv, "wrong in get bool")
        case ast.V_STRING:
            gotv, err := node.String()
            require.NoErrorf(t, err, "error in ast get string")
            require.Equalf(t, gotv, expv, "wrong in get string")
        case ast.V_ARRAY:
            gotv, err := node.Array()
            require.NoErrorf(t, err, "error in ast get array")
            require.Equalf(t, gotv, expv, "wrong in get array")
        case ast.V_OBJECT:
            gotv, err := node.Map()
            require.NoErrorf(t, err, "error in ast get object")
            require.Equalf(t, gotv, expv, "wrong in get object")
        case ast.V_NUMBER:
            gotv, err := node.Float64()
            require.NoErrorf(t, err, "error in ast get number")
            require.Equalf(t, gotv, expv, "wrong in get number")
        case ast.V_ANY:
            gotv, err := node.Interface()
            require.NoErrorf(t, err, "error in ast get any")
            require.Equalf(t, gotv, expv, "wrong in get any")
    }
}