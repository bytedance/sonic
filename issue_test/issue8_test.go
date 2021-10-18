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
    `reflect`
    `testing`

    . `github.com/bytedance/sonic`
)

func TestFloat(t *testing.T) {
    data := `{"test":0.6666}`
    var stdobj map[string]interface{}
    if err := json.Unmarshal([]byte(data), &stdobj); err != nil {
        t.Fatal(err)
    }
    var sonicobj map[string]interface{}
    if err := Unmarshal([]byte(data), &sonicobj); err != nil {
        t.Fatal(err)
    }
    if !reflect.DeepEqual(stdobj, sonicobj) {
        t.Fatalf("exp: \n%#v, \ngot: \n%#v\n", stdobj, sonicobj)
    }
}
