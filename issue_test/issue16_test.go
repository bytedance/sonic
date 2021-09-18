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
    `io/ioutil`
    `testing`

    `github.com/stretchr/testify/require`
)

func benchmarkEncodeSonic(b *testing.B, data []byte) {
    var xbook = map[string]interface{}{}
    if err := Unmarshal(data, &xbook); err != nil {
        b.Fatal(err)
    }
    if _, err := Marshal(&xbook); err != nil {
        b.Fatal(err)
    }
    b.SetBytes(int64(len(data)))
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Marshal(&xbook)
    }
}

func BenchmarkIssue16(b *testing.B) {
    data, err := ioutil.ReadFile("../testdata/twitterescaped.json")
    require.Nil(b, err)
    benchmarkEncodeSonic(b, data)
}