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
    `testing`

    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/require`
    `github.com/tidwall/gjson`
)

func TestNotFoud(t *testing.T) {
    data := `{}`

    ia := jsoniter.Get([]byte(data), "b")
    require.Error(t, ia.LastError())
    require.Equal(t, false, ia.ToBool())

    ga := gjson.GetBytes([]byte(data), "b")
    require.True(t, ga.Type == gjson.Null)
    require.Equal(t, false, ga.Bool())

    sa, err := NewSearcher(data).GetByPath("b")
    require.True(t, sa.Type() == V_NONE)
    require.Error(t, err)
    sv, err := sa.Bool()
    require.Error(t, err)
    require.Equal(t, false, sv)
}

func TestNull(t *testing.T) {
    data := `{"b": null}`

    ia := jsoniter.Get([]byte(data), "b")
    require.NoError(t, ia.LastError())
    require.Equal(t, false, ia.ToBool())

    ga := gjson.GetBytes([]byte(data), "b")
    require.True(t, ga.Type == gjson.Null)
    require.Equal(t, false, ga.Bool())

    sa, err := NewSearcher(data).GetByPath("b")
    require.True(t, sa.Type() == V_NULL)
    require.NoError(t, err)
    sv, err := sa.Bool()
    require.NoError(t, err)
    require.Equal(t, false, sv)
}