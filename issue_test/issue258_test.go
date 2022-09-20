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

package issue_test

import (
    `encoding/json`
    `fmt`
    `testing`

    `github.com/bytedance/sonic`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`
)


type M1 map[string]int

func (m *M1) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`{"m":%q}`, spew.Sprintf("%#+v", m))), nil
}

type M2 map[string]int

func (m M2) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`{"m":%q}`, spew.Sprintf("%#+v", m))), nil
}

func TestIssue258(t *testing.T) {
    m1 := M1{}
    oe,ee := json.Marshal(m1)
    os,es := sonic.Marshal(m1)
    require.Equal(t, ee, es)
    require.Equal(t, oe, os)

    m1p := &M1{}
    oe,ee = json.Marshal(m1p)
    os,es = sonic.Marshal(m1p)
    require.Equal(t, ee, es)
    require.Equal(t, oe, os)

    m2 := M2{}
    oe,ee = json.Marshal(m2)
    os,es = sonic.Marshal(m2)
    require.Equal(t, ee, es)
    require.Equal(t, oe, os)

    m2p := &M2{}
    oe,ee = json.Marshal(m2p)
    os,es = sonic.Marshal(m2p)
    require.Equal(t, ee, es)
    require.Equal(t, oe, os)
}