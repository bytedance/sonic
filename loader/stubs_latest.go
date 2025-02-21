// go:build go1.24 && !go1.25
// +build go1.24,!go1.25

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
    `unsafe`
    `runtime`
)

func hacklastmoduledatap() *moduledata {
    memprofile_addr := uintptr(unsafe.Pointer(&runtime.MemProfileRate))
    /*
        TODO: fixme, it is very tricky to get the address of lastmoduledatap
        objdump the main binary, we can find the address of runtime.MemProfileRate and the offset of lastmoduledatap
        ➜  .vscode git:(fix/float64) ✗ go tool nm -size ./main  | awk '$3 == "D"' | sort 
        567000        416 D go:buildinfo
        5671a0          1 D strconv.optimize
        5671a1          1 D runtime.oneptrmask
        5671a4          4 D runtime.adviseUnused
        5671a8          4 D runtime.epfd
        5671ac          4 D runtime.worldsema
        5671b0          4 D runtime.gcsema
        5671b4          4 D runtime.traceback_cache
        5671b8          4 D runtime.startingStackSize
        5671bc          5 D runtime.finalizer1
        5671e0          8 D runtime.MemProfileRate
        5671e8          8 D runtime.minOffAddr
        5671f0          8 D runtime.maxOffAddr
        5671f8          8 D runtime.sigset_all
        567200          8 D runtime.asyncPreemptStack
        567208          8 D runtime.forcegcperiod
        567210          8 D runtime.maxstacksize
        567218          8 D runtime.maxstackceiling
        567220          8 D runtime.intArgRegs
        567228          8 D runtime.lastmoduledatap
        567230          8 D reflect.intArgRegs
    */
    lastmoduledatap_addr := memprofile_addr + 0x48
    return *(**moduledata)(unsafe.Pointer(lastmoduledatap_addr))
}

var lastmoduledatap = hacklastmoduledatap()

func moduledataverify1(_ *moduledata) {
    // do nothing
}
