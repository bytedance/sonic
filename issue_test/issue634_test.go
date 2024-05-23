/*
 * Copyright 2024 ByteDance Inc.
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
    "strings"
    "testing"

    "sync"

    "github.com/bytedance/sonic"
    "github.com/bytedance/sonic/option"
    "github.com/stretchr/testify/assert"
)

func marshalSingle() {
    var m = map[string]interface{}{
        "1": map[string]interface{} {
            `"`+strings.Repeat("a", int(option.DefaultEncoderBufferSize) - 38)+`"`: "b",
            "1": map[string]int32{
                "b": 1658219785,
            },
        },
    }
    _, err := sonic.Marshal(&m)
    if err != nil {
        panic("err")
    }
}

type zoo foo

func (z *zoo) MarshalJSON() ([]byte, error) {
    marshalSingle()
    return sonic.Marshal((*foo)(z))
}

type foo bar

func (f *foo) MarshalJSON() ([]byte, error) {
    marshalSingle()
    return sonic.Marshal((*bar)(f))
}

type bar int

func (b *bar) MarshalJSON() ([]byte, error) {
    marshalSingle()
    return sonic.Marshal(int(*b))
}

 func TestEncodeOOM(t *testing.T) {
    wg := &sync.WaitGroup{}
    N := 10000
    for i:=0; i<N; i++ {
        wg.Add(1)
        go func (wg *sync.WaitGroup)  {
            defer wg.Done()
            var z zoo
            _, err := sonic.Marshal(&z)
            assert.NoError(t, err)
        }(wg)
    }
    wg.Wait()
 }