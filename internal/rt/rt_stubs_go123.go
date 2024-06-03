//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package rt

import (
	"unsafe"
)

//go:linkname rt_roundupsize  runtime.roundupsize
func rt_roundupsize(size uintptr, noscan bool) uintptr

// a wrapper for fastmap
func roundupsize(size uintptr) uintptr {
	return rt_roundupsize(size, MapType(MapEfaceType).Bucket.PtrData == 0)
}

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

	new := GrowSlice(et, *old, newLen-old.Len)
	new.Len = newLen
	return &new
}
