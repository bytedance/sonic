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

var _C_Redzone = []bool{false, false, false, false}

// CFunc is a function information for C func
type CFunc struct {
	// C function name
	Name     string

	// entry pc relative to entire text segment
	EntryOff uint32

	// function text size in bytes
	TextSize uint32

	// maximum stack depth of the function
	MaxStack uintptr

	// PC->SP delta lists of the function
	Pcsp     [][2]uint32
}

// GoC is the wrapper for Go calls to C
type GoC struct {
	// CName is the name of corresponding C function
	CName     string

	// CEntry points out where to store the entry address of corresponding C function.
	// It won't be set if nil
	CEntry   *uintptr

	// GoFunc is the POINTER of corresponding go stub function. 
	// It is used to generate Go-C ABI conversion wrapper and receive the wrapper's address 
	//   eg. &func(a int, b int) int 
	//     FOR 
	//     int add(int a, int b)
	// It won't be set if nil
	GoFunc   interface{} 
}

// WrapGoC wraps C functions and loader it into Go stubs
func WrapGoC(text []byte, natives []CFunc, stubs []GoC, modulename string, filename string) {
	panic("this is a fallback path, not implemented")
}
