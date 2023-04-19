/**
 * Copyright 2023 ByteDance Inc.
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
    `testing`

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

type Issue403 struct {
    Card   PtrAlias `json:"card"`
}

type PtrAlias *MarshalerImpl

type MarshalerImpl struct{
    A int
}

func (self *MarshalerImpl) MarshalJSON() ([]byte, error) {
    return []byte("1"), nil
}

func TestIssue403(t *testing.T) {
    obj := MarshalerImpl{0}
    messageInfo := &Issue403{
        Card:    &obj,
    }

    jsonData, err := json.Marshal(messageInfo)
    require.NoError(t, err)
    sonicData, err := sonic.Marshal(messageInfo)
    require.NoError(t, err)
    require.Equal(t, jsonData, sonicData)
}