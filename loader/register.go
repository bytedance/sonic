// +build !bytedance_tango

/**
 * Copyright 2024 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package loader

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// registerModuleMu serializes appends to the runtime module list.
//
// Appending has two steps: link mod from the previous tail, then publish mod
// as the new tail. Both steps must stay in that order, and no other append may
// interleave between them. Publishing first leaves a tail that is still
// unreachable from runtime.firstmoduledata, so a second registration that
// samples that tail chains itself behind an unlinked module. Frames of such a
// module are then invisible to runtime.findfunc, and stack growth or a GC scan
// aborts the process with "unknown caller pc".
var registerModuleMu sync.Mutex

func registerModule(mod *moduledata) {
    registerModuleMu.Lock()
    appendModule(&lastmoduledatap, mod)
    registerModuleMu.Unlock()
}

func appendModule(tail **moduledata, mod *moduledata) {
    oldTail := loadModule(tail)
    storeModule(&oldTail.next, mod)
    storeModule(tail, mod)
}

func loadModule(p **moduledata) *moduledata {
    return (*moduledata)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
}

func storeModule(p **moduledata, value *moduledata) {
    atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(value))
}

