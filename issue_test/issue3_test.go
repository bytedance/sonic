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
    `encoding/json`
    `reflect`
    `testing`

    `github.com/stretchr/testify/assert`
)

func TestIssue3_Encode(t *testing.T) {
    var v HugeStruct6
    ret, err := Marshal(v)
    assert.Nil(t, err)
    assert.Equal(t, []byte(`{}`), ret)
}

func TestIssue3_Decode(t *testing.T) {
    var v HugeStruct6
    err := Unmarshal([]byte(`{}`), &v)
    assert.Nil(t, err)
    assert.Equal(t, HugeStruct6{}, v)
}

func BenchmarkIssue3_Encode_Sonic(b *testing.B) {
    var v HugeStruct6
    err := Pretouch(reflect.TypeOf(v))
    assert.Nil(b, err)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Marshal(v)
    }
}

func BenchmarkIssue3_Encode_StdLib(b *testing.B) {
    var v HugeStruct6
    ret, err := json.Marshal(v)
    assert.Nil(b, err)
    assert.Equal(b, []byte(`{}`), ret)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = json.Marshal(v)
    }
}

func BenchmarkIssue3_Decode_Sonic(b *testing.B) {
    var v HugeStruct6
    buf := []byte(`{}`)
    err := Pretouch(reflect.TypeOf(v))
    assert.Nil(b, err)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = Unmarshal(buf, &v)
    }
}

func BenchmarkIssue3_Decode_StdLib(b *testing.B) {
    var v HugeStruct6
    buf := []byte(`{}`)
    err := json.Unmarshal(buf, &v)
    assert.Nil(b, err)
    assert.Equal(b, HugeStruct6{}, v)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = json.Unmarshal(buf, &v)
    }
}
