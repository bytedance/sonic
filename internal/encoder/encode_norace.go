//go:build !race
// +build !race

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

package encoder

import (
    `runtime`

    `github.com/bytedance/sonic/internal/rt`
)


func encodeInto(buf *[]byte, val interface{}, opts Options) error {
    stk := newStack()
    efv := rt.UnpackEface(val)
    err := encodeTypedPointer(buf, efv.Type, &efv.Value, stk, uint64(opts))

    /* return the stack into pool */
    if err != nil {
        resetStack(stk)
    }
    freeStack(stk)

    /* avoid GC ahead */
    runtime.KeepAlive(buf)
    runtime.KeepAlive(efv)
    return err
}
