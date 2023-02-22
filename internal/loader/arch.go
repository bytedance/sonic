//go:build darwin || linux
// +build darwin linux

/*
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
    `fmt`
)

// Arch describes an architecture.
type Arch int

const (
        // AMD64 is the x86-64 architecture.
        AMD64 Arch = iota
        // ARM64 is the aarch64 architecture.
        ARM64
)

// String implements fmt.Stringer.
func (a Arch) String() string {
        switch a {
        case AMD64:
                return "amd64"
        case ARM64:
                return "arm64"
        default:
                return fmt.Sprintf("Arch(%d)", a)
        }
}
