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

var (
    // Default value(3) means the compiler only inline 3 layers of nested struct. 
    // when the depth exceeds, the compiler will recurse 
    // and compile subsequent structs when they are decoded 
    DefaultMaxInlineDepth = 3

    // Default value(1) means `Pretouch()` will be recursively executed once,
    // if any nested struct is left (depth exceeds MaxInlineDepth)
    DefaultRecursiveDepth = 1
)

// CompileOptions includes all options for encoder or decoder compiler.
type CompileOptions struct {
    // the maximum depth for compilation inline
    MaxInlineDepth    int

    // the loop times for recursive pretouch
    RecursiveDepth int
}

// DefaultCompileOptions set default compile options.
func DefaultCompileOptions() CompileOptions {
    return CompileOptions{
        RecursiveDepth: DefaultRecursiveDepth,
        MaxInlineDepth: DefaultMaxInlineDepth,
    }
}

// CompileOption is a function used to change DefaultCompileOptions.
type CompileOption func(o *CompileOptions)

// WithCompileRecursiveDepth sets the loops of recursive pretouch 
// in decoder and encoder.
//
// For deep nested struct (depth exceeds MaxInlineDepth), 
// try to set larger depth to reduce compile time 
// in the first Marshal or Unmarshal.
func WithCompileRecursiveDepth(loop int) CompileOption {
    return func(o *CompileOptions) {
            if loop < 0 {
                panic("loop must be >= 0")
            }
            o.RecursiveDepth = loop
        }
}

// WithCompileMaxInlineDepth sets the max depth of inline compile 
// in decoder and encoder.
//
// The larger the depth is, the compilation time of one pretouch loop may be longer
func WithCompileMaxInlineDepth(depth int) CompileOption {
    return func(o *CompileOptions) {
            if depth <= 0 {
                panic("depth must be > 0")
            }
            o.MaxInlineDepth = depth
        }
}
 