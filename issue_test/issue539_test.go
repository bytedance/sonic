// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
   `encoding/json`
   `testing`

   `github.com/bytedance/sonic`
   `github.com/stretchr/testify/require`
)

type OmitEmptyInterface struct {
    ErrCode int32       `json:"code"`
    Data    interface{} `json:"data,omitempty"`
}

func TestOmitEmptyInterface(t *testing.T) {
    // non-enmpty type
    var data *string
        resp := &OmitEmptyInterface{
        ErrCode: 123,
        Data:    data,
        }
    eout, eerr := json.Marshal(resp)
    sout, serr := sonic.Marshal(resp)
    require.Equal(t, eerr == nil, serr == nil)
    require.Equal(t, string(eout), string(sout))

    // empty type and value
    resp = &OmitEmptyInterface{
        ErrCode: 123,
        Data:    nil,
        }
        eout, eerr = json.Marshal(resp)
        sout, serr = sonic.Marshal(resp)
        require.Equal(t, eerr == nil, serr == nil)
        require.Equal(t, string(eout), string(sout))
}