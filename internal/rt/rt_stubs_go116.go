//go:build go1.16 && !go1.20
// +build go1.16,!go1.20

package rt

import (
	"unsafe"
)

//go:linkname roundupsize  runtime.roundupsize
func roundupsize(size uintptr) uintptr

//go:linkname makeslice runtime.makeslice
//goland:noinspection GoUnusedParameter
func makeslice(et *GoType, len int, cap int) unsafe.Pointer

func MakeSlice(oldPtr unsafe.Pointer, et *GoType, newLen int) *GoSlice {
	if newLen == 0 {
		return &GoSlice{
			Ptr: ZSTPtr,
			Len: 0,
			Cap: 0,
		}
	}

	if *(*unsafe.Pointer)(oldPtr) == nil {
		return &GoSlice{
			Ptr: makeslice(et, newLen, newLen),
			Len: newLen,
			Cap: newLen,
		}
	}

	old := (*GoSlice)(oldPtr)
	if old.Cap >= newLen {
		old.Len = newLen
		return old
	}

	new := GrowSlice(et, *old, newLen)

	// we sould clear the memory from [oldLen:newLen]
	if et.PtrData == 0 {
		oldlenmem := uintptr(old.Len) * et.Size
		newlenmem := uintptr(newLen) * et.Size
		MemclrNoHeapPointers(add(new.Ptr, oldlenmem), newlenmem-oldlenmem)
	}

	new.Len = newLen
	return &new
}
