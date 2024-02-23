/**
 * Copyright 2024 ByteDance Inc.
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

package issue_test

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/require"
)


func TestIssue600(t *testing.T) {
    // object
    obj := ast.NewRaw("{\"x\":\"a\",\"y\":\"b\"}")
    if ok, err := obj.Unset("x"); !ok || err != nil {
        panic(fmt.Errorf("unset x fail, ok=%v, err=%v", ok, err))
    }
    if ok, err := obj.Unset("y"); !ok || err != nil {
        panic(fmt.Errorf("unset y fail, ok=%v, err=%v", ok, err))
    }
    result, err := obj.MarshalJSON()
    if err != nil {
        panic(fmt.Errorf("MarshalJSON fail: err=%v", err))
    }
    require.Equal(t, `{}`, string(result))

    obj = ast.NewRaw("{\"x\":\"a\",\"y\":\"b\"}")
    if ok, err := obj.Unset("y"); !ok || err != nil {
        panic(fmt.Errorf("unset x fail, ok=%v, err=%v", ok, err))
    }
    if ok, err := obj.Unset("x"); !ok || err != nil {
        panic(fmt.Errorf("unset y fail, ok=%v, err=%v", ok, err))
    }
    result, err = obj.MarshalJSON()
    if err != nil {
        panic(fmt.Errorf("MarshalJSON fail: err=%v", err))
    }
    require.Equal(t, `{}`, string(result))

    // array
    obj = ast.NewRaw("[1,2]")
    if ok, err := obj.UnsetByIndex(0); !ok || err != nil {
        panic(fmt.Errorf("unset x fail, ok=%v, err=%v", ok, err))
    }
    if ok, err := obj.UnsetByIndex(0); !ok || err != nil {
        panic(fmt.Errorf("unset y fail, ok=%v, err=%v", ok, err))
    }
    result, err = obj.MarshalJSON()
    if err != nil {
        panic(fmt.Errorf("MarshalJSON fail: err=%v", err))
    }
    require.Equal(t, `[]`, string(result))

    obj = ast.NewRaw("[1,2]")
    if ok, err := obj.UnsetByIndex(1); !ok || err != nil {
        panic(fmt.Errorf("unset x fail, ok=%v, err=%v", ok, err))
    }
    if ok, err := obj.UnsetByIndex(0); !ok || err != nil {
        panic(fmt.Errorf("unset y fail, ok=%v, err=%v", ok, err))
    }
    result, err = obj.MarshalJSON()
    if err != nil {
        panic(fmt.Errorf("MarshalJSON fail: err=%v", err))
    }
    require.Equal(t, `[]`, string(result))
}
