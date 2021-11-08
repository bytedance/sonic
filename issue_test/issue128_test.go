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
    `runtime/debug`
    `testing`

    `github.com/bytedance/sonic`
)

//NOTICE: only DEBUG mode can reproduce problem
func TestDebug(t *testing.T) {
    debug.SetGCPercent(-1)
    data := &Data{
        Details: []*Detail{
            {
                Info: Info{},
            },
        },
    }

    body, err := sonic.Marshal(data)
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(string(body))
}

type Detail struct {
    Info Info
}

type Info struct {
    A int 
}

type Data struct {
    Details []*Detail
}

