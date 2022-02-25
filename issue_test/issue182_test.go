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
    `encoding/json`
    `strconv`
    `testing`

    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

type MarshalerWrap struct {
    A *ValueMarshaler
    B ValueMarshaler
    C *PointerMarshaler
    D PointerMarshaler
    
    E *ValueTextMarshaler
    F ValueTextMarshaler
    G *PointerTextMarshaler
    H PointerTextMarshaler
}

type ValueMarshaler struct {
    X int 
}

func (v ValueMarshaler) MarshalJSON() ([]byte, error) {
    return []byte(strconv.Itoa(v.X)), nil
}

type PointerMarshaler struct {
    X int 
}

func (v *PointerMarshaler) MarshalJSON() ([]byte, error) {
    return []byte(strconv.Itoa(v.X)), nil
}

type ValueTextMarshaler struct {
    X int 
}

func (v ValueTextMarshaler) MarshalText() ([]byte, error) {
    return []byte(strconv.Itoa(v.X)), nil
}

type PointerTextMarshaler struct {
    X int 
}

func (v *PointerTextMarshaler) MarshalText() ([]byte, error) {
    return []byte(strconv.Itoa(v.X)), nil
}

func TestIssue182(t *testing.T) {
    v0 := MarshalerWrap{}
    ret, err := json.Marshal(v0)
    rets, errs := sonic.Marshal(v0)
    require.Equal(t, err, errs)
    require.Equal(t, string(ret), string(rets))
    ret, err = json.Marshal(&v0)
    rets, errs = sonic.Marshal(&v0)
    require.Equal(t, err, errs)
    require.Equal(t, string(ret), string(rets))
    
    v1 := MarshalerWrap{A:&ValueMarshaler{}, C:&PointerMarshaler{}, E: &ValueTextMarshaler{}, G:&PointerTextMarshaler{}}
    ret, err = json.Marshal(v1)
    rets, errs = sonic.Marshal(v1)
    require.Equal(t, err, errs)
    require.Equal(t, string(ret), string(rets))
    ret, err = json.Marshal(&v1)
    rets, errs = sonic.Marshal(&v1)
    require.Equal(t, err, errs)
    require.Equal(t, string(ret), string(rets))
}


