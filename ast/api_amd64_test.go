// +build amd64,go1.16,!go1.22

/*
 * Copyright 2022 ByteDance Inc.
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
    `encoding/json`
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/require`
)

func TestSortNodeTwitter(t *testing.T) {
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    obj, err := root.MapUseNumber()
    if err != nil {
        t.Fatal(err)
    }
    exp, err := encoder.Encode(obj, encoder.SortMapKeys)
    if err != nil {
        t.Fatal(err)
    }
    var expObj interface{}
    require.NoError(t, json.Unmarshal(exp, &expObj))

    if err := root.SortKeys(true); err != nil {
        t.Fatal(err)
    }
    act, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    var actObj interface{}
    require.NoError(t, json.Unmarshal(act, &actObj))
    require.Equal(t, expObj, actObj)
    require.Equal(t, len(exp), len(act))
    require.Equal(t, string(exp), string(act))
}