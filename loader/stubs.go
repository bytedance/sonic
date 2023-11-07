/**
 * Copyright 2023 ByteDance Inc.
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

package loader

import (
    `sync/atomic`
    `unsafe`
    _ `unsafe`
)

//go:linkname lastmoduledatap runtime.lastmoduledatap
//goland:noinspection GoUnusedGlobalVariable
var lastmoduledatap *moduledata

func registerModule(mod *moduledata) {
load:
    old := (*moduledata)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&lastmoduledatap))))
    next := (*moduledata)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&old.next))))
    // load and swap old.next
    if !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&old.next)), 
        unsafe.Pointer(next), unsafe.Pointer(mod)){
        goto load
    }
    // sucussefully exchange old.next, then load and swap lastmoduledatap
    if !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&lastmoduledatap)), 
        unsafe.Pointer(old), unsafe.Pointer(mod)) {
        // exchange fail, recover old.next...
        atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&old.next)), unsafe.Pointer(next))
        goto load
    }
}

//go:linkname moduledataverify1 runtime.moduledataverify1
func moduledataverify1(_ *moduledata)


