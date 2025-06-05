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

//go:build go1.21
// +build go1.21

package rt

import (
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/option"
)

//go:nosplit
func Get16(v []byte) int16 {
	return *(*int16)(unsafe.Pointer(unsafe.SliceData(v)))
}

//go:nosplit
func Get32(v []byte) int32 {
	return *(*int32)(unsafe.Pointer(unsafe.SliceData(v)))
}

//go:nosplit
func Get64(v []byte) int64 {
	return *(*int64)(unsafe.Pointer(unsafe.SliceData(v)))
}

// Mem2Str uses Go 1.20+ unsafe.String and unsafe.SliceData for cross-platform, zero-copy string to []byte.
//go:nosplit
func Mem2Str(v []byte) (s string) {
	return unsafe.String(unsafe.SliceData(v), len(v))
}

// Str2Mem uses Go 1.20+ unsafe.Slice for cross-platform, zero-copy, zero-copy []byte to string
//go:nosplit
func Str2Mem(s string) (v []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesFrom(p unsafe.Pointer, n int, c int) (r []byte) {
	return unsafe.Slice((*byte)(p), n)
}

func FuncAddr(f interface{}) unsafe.Pointer {
	if vv := UnpackEface(f); vv.Type.Kind() != reflect.Func {
		panic("f is not a function")
	} else {
		return *(*unsafe.Pointer)(vv.Value)
	}
}

//go:nocheckptr
func IndexChar(src string, index int) unsafe.Pointer {
	return unsafe.Pointer(uintptr((*GoString)(unsafe.Pointer(&src)).Ptr) + uintptr(index))
}

//go:nocheckptr
func IndexByte(ptr []byte, index int) unsafe.Pointer {
	return unsafe.Pointer(uintptr((*GoSlice)(unsafe.Pointer(&ptr)).Ptr) + uintptr(index))
}

func GuardSlice(buf *[]byte, n int) {
	c := cap(*buf)
	l := len(*buf)
	if c-l < n {
		c = c>>1 + n + l
		if c < 32 {
			c = 32
		}
		tmp := make([]byte, l, c)
		copy(tmp, *buf)
		*buf = tmp
	}
}

func GuardSlice2(buf []byte, n int) []byte {
	c := cap(buf)
	l := len(buf)
	if c-l < n {
		c = c>>1 + n + l
		if c < 32 {
			c = 32
		}
		tmp := make([]byte, l, c)
		copy(tmp, buf)
		buf = tmp
	}
	return buf
}

//go:nosplit
func Ptr2SlicePtr(s unsafe.Pointer, l int, c int) unsafe.Pointer {
	slice := &GoSlice{
		Ptr: s,
		Len: l,
		Cap: c,
	}
	return unsafe.Pointer(slice)
}

// StrPtr uses Go 1.20+ unsafe.StringData to get the pointer to the string's underlying data.
//go:nosplit
func StrPtr(s string) unsafe.Pointer {
	return unsafe.Pointer(unsafe.StringData(s))
}

// StrFrom uses Go 1.20+ unsafe.String for cross-platform, zero-copy string construction.
//go:nosplit
func StrFrom(p unsafe.Pointer, n int64) (s string) {
	return unsafe.String((*byte)(p), int(n))
}

// NoEscape hides a pointer from escape analysis. NoEscape is
// the identity function but escape analysis doesn't think the
// output depends on the input. NoEscape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//
//go:nosplit
//goland:noinspection GoVetUnsafePointer
func NoEscape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

//go:nosplit
func MoreStack(size uintptr)

//go:nosplit
func Add(ptr unsafe.Pointer, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + off)
}

// CanSizeResue
func CanSizeResue(cap int) bool {
	return cap <= int(option.LimitBufferSize)
}
