//
// Copyright 2025 Huawei Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package arm64

// Arch represents the ARM64 architecture
type Arch struct{}

// DefaultArch is the default ARM64 architecture instance
var DefaultArch = &Arch{}

// CreateProgram creates a new ARM64 program
func (a *Arch) CreateProgram() *Program {
	return NewProgram()
}

// Name returns the architecture name
func (a *Arch) Name() string {
	return "arm64"
}

// PointerSize returns the pointer size in bytes
func (a *Arch) PointerSize() int {
	return 8
}

// InstructionAlignment returns the instruction alignment requirement
func (a *Arch) InstructionAlignment() int {
	return 4
}
