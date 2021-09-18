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
    `fmt`
    `testing`

    `github.com/davecgh/go-spew/spew`
)

type SingleMapField struct {
    Z *int
}

type SingleMapFieldOuter struct {
    Y *SingleMapField
}

type SingleMapFieldOuterContainer struct {
    X *SingleMapFieldOuter
}

func TestIssue76_MarshalSingleMapField(t *testing.T) {
    data := `{"X": {"Y": {"Z": 1}}}`
    obj := new(SingleMapFieldOuterContainer)
    if err := json.Unmarshal([]byte(data), obj); err != nil {
        t.Fatal(err)
    }
    spew.Dump(obj)
    buf, err := Marshal(obj)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(string(buf))
}