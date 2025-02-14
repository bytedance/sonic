//go:build go1.24 && !go1.25 && goexperiment.swissmap

package rt

import (
	"unsafe"
)

type GoMap struct {
	Count int// used uint64
	seed uintptr
	dirPtr unsafe.Pointer
	dirLen int
	globalDepth uint8
	globalShift uint8
	writing uint8
	clearSeq uint64
}

type GoMapIterator struct {
	K     unsafe.Pointer
	V     unsafe.Pointer
	T     *GoMapType
	It    unsafe.Pointer
}
