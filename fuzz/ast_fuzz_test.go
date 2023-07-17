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
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/utf8"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

// data is random, check whether is panic
func fuzzAst(t *testing.T, data []byte) {
	n, err := sonic.Get(data)
	if err == nil && utf8.Validate(data) {
		var x interface{}
		if err := json.Unmarshal(data, &x); err != nil {
			return
		}
		y, err := n.Interface()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(x, y) {
			t.Fatalf("exp:%#v, got:%#v", x, y)
		}
		_, err = n.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func fuzzASTGetFromObject(t *testing.T, data []byte, m map[string]interface{}) {
	for k, expv := range(m) {
		msg := fmt.Sprintf("Data:\n%s\nKey:\n%s\n", spew.Sdump(&data), spew.Sdump(&k))
		node, err := sonic.Get(data, k)
		require.NoErrorf(t, err, "error in ast get key\n%s", msg)
		v, err := node.Interface()
		require.NoErrorf(t, err, "error in node convert\n%s", msg)
		require.Equalf(t, v, expv, "error in node equal\n%sGot:\n%s\nExp:\n%s\n", 
			msg, spew.Sdump(v), spew.Sdump(expv))
	}
}

func fuzzASTGetFromArray(t *testing.T, data []byte, a []interface{}) {
	i := 0
	for ; i < len(a); i++ {
		msg := fmt.Sprintf("Data:\n%s\nIndex:\n%d\n", spew.Sdump(data), i)
		node, err := sonic.Get(data, i)
		require.NoErrorf(t, err, "error in ast get index\n%s", msg)
		v, err := node.Interface()
		require.NoErrorf(t, err, "error in node convert\n%s", msg)
		require.Equalf(t, v, a[i], "error in node equal\n%sGot:\n%s\nExp:\n%s\n", 
			msg, spew.Sdump(v), spew.Sdump(a[i]))
	}
	_, err := sonic.Get(data, i)
	require.Errorf(t, err, "no error in ast get out of range\nData:\n%s\n", spew.Sdump(data))
}