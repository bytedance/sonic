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

//go:noescape
//go:linkname growslice runtime.growslice
//goland:noinspection GoUnusedParameter
func growslice(et *GoType, old GoSlice, cap int) GoSlice

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

	new := growslice(et, *old, newLen)
	// growslice does not zero out new[old.cap:new.len] since it assumes that
	// the memory will be overwritten by an append() that called growslice.
	// Since the caller of reflect_growslice is not append(),
	// zero out this region before returning the slice to the reflect package.
	if et.PtrData == 0 {
		oldcapmem := uintptr(old.Cap) * et.Size
		newlenmem := uintptr(new.Len) * et.Size
		MemclrNoHeapPointers(add(new.Ptr, oldcapmem), newlenmem-oldcapmem)
	}

	new.Len = newLen
	return &new
}
