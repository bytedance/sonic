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
    `strings`
    `testing`

    `github.com/bytedance/sonic/decoder`
)

func TestIssue293(t *testing.T) {
    left := `{"a":`
    var data = left+strings.Repeat(" ", 4096 - len(left)-3) + "33.0}"
    sd := decoder.NewStreamDecoder(strings.NewReader(data))
    var v = struct{
        A json.RawMessage
    }{}
    err := sd.Decode(&v)
    if err != nil {
        t.Fatal(err)
    }
}

func TestIssue293Sign(t *testing.T) {
    left := `{"a":`
    var data = left+strings.Repeat(" ", 4096 - len(left)-1) + "-33.0}"
    sd := decoder.NewStreamDecoder(strings.NewReader(data))
    var v = struct{
        A json.RawMessage
    }{}
    err := sd.Decode(&v)
    if err != nil {
	if e, ok := err.(decoder.SyntaxError); ok {
	    t.Fatal(e.Description())
	}
        t.Fatal(err)
    }
}