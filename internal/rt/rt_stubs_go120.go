//go:build go1.20 && !go1.22
// +build go1.20,!go1.22

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
	new.Len = newLen
	return &new
}
