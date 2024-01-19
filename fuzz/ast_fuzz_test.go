//go:build go1.18
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
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

// data is random, check whether is panic
func fuzzAst(t *testing.T, data []byte) {
	sonic.GetValueFromString(string(data))
}


func fuzzASTGetFromObject(t *testing.T, v []byte, m map[string]interface{}) {
	data := string(v)
	for k, expv := range(m) {
		msg := fmt.Sprintf("Data:\n%s\nKey:\n%s\n", spew.Sdump(&data), spew.Sdump(&k))
		node, err := sonic.GetValueFromString(data, k)
		require.NoErrorf(t, err, "error in ast get key\n%s", msg)
		v, err := node.Interface()
		require.NoErrorf(t, err, "error in node convert\n%s", msg)
		require.Equalf(t, v, expv, "error in node equal\n%sGot:\n%s\nExp:\n%s\n",msg, spew.Sdump(v), spew.Sdump(expv))
		nj, err := sonic.DeleteFromString(data, k)
		require.NoErrorf(t, err, "error in node delete\n%s", msg)
		nj, err = sonic.SetFromString(nj, expv, k)
		require.NoErrorf(t, err, "error in node set\n%s", msg)
		nn, _ := sonic.GetValueFromString(nj)
		nv, err := nn.Interface()
		require.NoErrorf(t, err, "error in node set\n%s", msg)
		require.Equalf(t, m, nv, msg)
	}
}

func fuzzASTGetFromArray(t *testing.T, v []byte, a []interface{}) {
	i := 0
	data := string(v)
	for ; i < len(a); i++ {
		msg := fmt.Sprintf("Data:\n%s\nIndex:\n%d\n", spew.Sdump(data), i)
		node, err := sonic.GetValueFromString(data, i)
		require.NoErrorf(t, err, "error in ast get index\n%s", msg)
		v, err := node.Interface()
		require.NoErrorf(t, err, "error in node convert\n%s", msg)
		require.Equalf(t, v, a[i], "error in node equal\n%sGot:\n%s\nExp:\n%s\n", msg, spew.Sdump(v), spew.Sdump(a[i]))
		next, err := sonic.DeleteFromString(data, i)
		require.NoErrorf(t, err, "error in node delete\n%s", msg)
		nj := ast.NewValueJSON(next)
		err = nj.AddAny(i, v)
		require.NoErrorf(t, err, "error in node add\n%s", msg)
		nv, err := nj.Interface()
		require.NoErrorf(t, err, "error in node set\n%s", msg)
		require.Equalf(t, a, nv, msg)
	}
	_, err := sonic.GetValueFromString(data, i)
	require.Errorf(t, err, "no error in ast get out of range\nData:\n%s\n", spew.Sdump(data))
}