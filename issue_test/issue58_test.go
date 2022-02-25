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

package issue_test

import (
    . `github.com/bytedance/sonic`
    `testing`
    `encoding/json`

    `github.com/stretchr/testify/require`
)

type (
    Issue58ValueReceiver   struct {}
    Issue58PointerReceiver struct {}
)

func (_ Issue58ValueReceiver) MarshalJSON() ([]byte, error)    { return []byte(`"value"`), nil }
func (_ *Issue58PointerReceiver) MarshalJSON() ([]byte, error) { return []byte(`"pointer"`), nil }

func TestIssue58_NilPointerOnValueMethod(t *testing.T) {
    v := struct {
        X *Issue58ValueReceiver
        Y *Issue58PointerReceiver
    }{}
    buf, err := Marshal(v)
    require.NoError(t, err)
    require.Equal(t, []byte(`{"X":null,"Y":null}`), buf)
    buf, err = json.Marshal(v)
    require.NoError(t, err)
    require.Equal(t, []byte(`{"X":null,"Y":null}`), buf)

}
