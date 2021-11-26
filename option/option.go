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

// CompileOptions includes all options for encoder or decoder compiler.
type CompileOptions struct {
    // the depth for recursive compile
    RecursiveDepth int
}

func DefaultCompileOptions() CompileOptions {
    return CompileOptions{
        RecursiveDepth: 0,
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
 