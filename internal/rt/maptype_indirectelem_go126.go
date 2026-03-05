//go:build go1.26
// +build go1.26

package rt

import "unsafe"

const mapIndirectElemGo126 = 1 << 3

type goMapTypeGo126 struct {
	GoType
	Key       *GoType
	Elem      *GoType
	Group     *GoType
	Hasher    func(unsafe.Pointer, uintptr) uintptr
	GroupSize uintptr
	SlotSize  uintptr
	ElemOff   uintptr
	Flags     uint32
}

func (self *GoMapType) IndirectElem() bool {
	return (*goMapTypeGo126)(unsafe.Pointer(self)).Flags&mapIndirectElemGo126 != 0
}
