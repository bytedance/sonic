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
    `sync`
    `testing`

    `github.com/bytedance/sonic`
)

type ByteStruct struct {
    Bytes    []byte
}

type ObjStruct struct {
    Obj ByteStruct
}

func TestIssue213(t *testing.T) {
    // bytes := []byte("{\"Obj\":{\"Bytes\":\"eyJUZXN0Q29kZSI6MjIyMiwiVGVzdFN0cmluZyI6InRlc3Rfc3RyaW5n\"}}") // this is OK
    bytes := []byte("{\"Obj\":{\"Bytes\":\"eyJUZXN0Q29kZSI6MjIyMiwiVGVzdFN0cmluZyI6InRlc3Rfc3RyaW5n\", \"x\":0}}")
    wg := sync.WaitGroup{}
    for i:=0;i<1000;i++{
        wg.Add(1)
        go func(){
            defer wg.Done()
            var o *ObjStruct
            if err := sonic.Unmarshal(bytes, &o); err != nil {
                t.Fatal(err)
            }
        }()
    }
    wg.Wait()
}

func BenchmarkIssue213(b *testing.B) {
    // bytes := []byte("{\"Obj\":{\"Bytes\":\"eyJUZXN0Q29kZSI6MjIyMiwiVGVzdFN0cmluZyI6InRlc3Rfc3RyaW5n\"}}") // this is OK
    js := "{\"Obj\":{\"Bytes\":\"eyJUZXN0Q29kZSI6MjIyMiwiVGVzdFN0cmluZyI6InRlc3Rfc3RyaW5n\", \"x\":0}}"
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        var o *ObjStruct
        _ = sonic.UnmarshalString(js, &o)
    }
}
