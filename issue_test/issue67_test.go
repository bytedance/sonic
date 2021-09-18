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
    `testing`
)

func TestIssue67_JunkAfterJSON(t *testing.T) {
    data := `1e2e3`
    var stdobj, sonicobj interface{}
    stderr := json.Unmarshal([]byte(data), &stdobj)
    sonicerr := Unmarshal([]byte(data), &sonicobj)
    if (stderr == nil) != (sonicerr == nil) {
        t.Fatalf("exp err: \n%#v, \ngot err: \n%#v\n", stderr, sonicerr)
    }
}
