//go:build amd64 || arm64
// +build amd64 arm64

/**
 * Copyright 2025 ByteDance Inc.
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

package utils


// Hack: this is used for both checking space and cause friendly compile errors in 32-bit arch.
const _Sonic_Not_Support_32Bit_Arch__Checking_32Bit_Arch_Here = (1 << ' ') | (1 << '\t') | (1 << '\r') | (1 << '\n')

func IsSpace(c byte) bool {
    return (int(1<<c) & _Sonic_Not_Support_32Bit_Arch__Checking_32Bit_Arch_Here) != 0
}