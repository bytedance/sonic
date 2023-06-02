//go:build !go1.16
// +build !go1.16

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

package loader

import (
    `github.com/bytedance/sonic/internal/loader`
)

// LoadFuncs loads only one function as module, and returns the function pointer
//   - text: machine code
//   - funcName: function name
//   - frameSize: stack frame size. 
//   - argSize: argument total size (in bytes)
//   - argPtrs: indicates if a slot (8 Bytes) of arguments memory stores pointer, from low to high
//   - localPtrs: indicates if a slot (8 Bytes) of local variants memory stores pointer, from low to high
// 
// WARN: 
//   - the function MUST has fixed SP offset equaling to this, otherwise it go.gentraceback will fail
//   - the function MUST has only one stack map for all arguments and local variants
func (self Loader) LoadOne(text []byte, funcName string, frameSize int, argSize int, argPtrs []bool, localPtrs []bool) Function {
    return Function(loader.Loader(text).Load(funcName, frameSize, argSize, argPtrs, localPtrs))
}

// Load loads given machine codes and corresponding function information into go moduledata
// and returns runnable function pointer
// WARN: this API is experimental, use it carefully
func Load(text []byte, funcs []Func, modulename string, filenames []string) (out []Function) {
    panic("not implemented")
}
