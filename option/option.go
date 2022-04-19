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

package option

import (
    `sync`
)

// CompileOptions includes all options for encoder or decoder compiler.
type CompileOptions struct {
    // the depth for recursive compile
    RecursiveDepth int
    // SliceOrMapNoNull indicates encoder to output "[]" or "{}" chars
    // while encountering a nil value of slice or map type, insetead of "null"
    SliceOrMapNoNull bool
}

func DefaultCompileOptions() CompileOptions {
    return CompileOptions{
        RecursiveDepth: 0,
        SliceOrMapNoNull: false,
    }
}

type CompileOption func(o *CompileOptions)

// WithCompileRecursiveDepth sets the depth of recursive compile 
// in decoder or encoder.
//
// Default value(0) is suitable for basic types and small nested struct types.
// 
// For large or deep nested struct, try to set larger depth to reduce compile 
// time in the first Marshal or Unmarshal.
func WithCompileRecursiveDepth(depth int) CompileOption {
    return func(o *CompileOptions) {
            o.RecursiveDepth = depth
        }
}

// SliceOrMapNoNull indicates encoder to output "[]" or "{}" chars
// while encountering a nil value of slice or map type, insetead of "null"
func WithSliceOrMapNoNull(no bool) CompileOption {
    return func(o *CompileOptions) {
            o.SliceOrMapNoNull = no
        }
}

var (
	global CompileOptions = DefaultCompileOptions()
    // Settled bool
    globalMux sync.RWMutex
)
 
func SetCompileOptions(opts ...CompileOption) CompileOptions {
    // if Settled {
    //     panic("compile options can only be set once!")
    // }
    globalMux.Lock()
    defer globalMux.Unlock()
    for _, opt := range opts {
        opt(&global)
    }
    return global
}

func GetCompileOptions() (ret CompileOptions) {
    globalMux.RLock() 
    ret = global
    globalMux.RUnlock()
    return
}