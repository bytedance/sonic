/*
 * Copyright 2022 ByteDance Inc.
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

package sonic

import (
    `reflect`
    `testing`

    `github.com/bytedance/sonic/option`
)

func TestPretouch(t *testing.T) {
    var v map[string]interface{}
    if err := Pretouch(reflect.TypeOf(v)); err != nil {
        t.Errorf("err:%v", err)
    }

    if err := Pretouch(reflect.TypeOf(v),
       option.WithCompileRecursiveDepth(1),
       option.WithCompileMaxInlineDepth(2),
    ); err != nil {
        t.Errorf("err:%v", err)
    }
}

func TestGet(t *testing.T) {
    var data = `{"a":"b"}`
    r, err := GetFromString(data, "a")
    if err != nil {
        t.Fatal(err)
    }
    v, err := r.String()
    if err != nil {
        t.Fatal(err)
    }
    if v != "b" {
        t.Fatal(v)
    }
}

