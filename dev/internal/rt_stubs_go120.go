//go:build go1.20
// +build go1.20

package internal

import (
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
)

//go:linkname makeslice runtime.makeslice
//goland:noinspection GoUnusedParameter
func makeslice(et *rt.GoType, len int, cap int) unsafe.Pointer

//go:noescape
//go:linkname growslice reflect.growslice
//goland:noinspection GoUnusedParameter
func growslice(et *rt.GoType, old rt.GoSlice, num int) rt.GoSlice

func MakeSlice(oldPtr unsafe.Pointer, et *rt.GoType, newLen int) *rt.GoSlice {
	if newLen == 0 {
		return &rt.GoSlice{
			Ptr: _ZSTPtr,
			Len: 0,
			Cap: 0,
		}
	}

	if *(*unsafe.Pointer)(oldPtr) == nil {
		return &rt.GoSlice{
			Ptr: makeslice(et, newLen, newLen),
			Len: newLen,
			Cap: newLen,
		}
	}

	old := (*rt.GoSlice)(oldPtr)
	if old.Cap >= newLen {
		old.Len = newLen
		return old
	}

	new := growslice(et, *old, newLen-old.Len)
	new.Len = newLen
	return &new
}
