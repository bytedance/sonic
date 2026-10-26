package rt

import (
	"unsafe"
)

// GoMapIterator mirrors the pre-Go 1.24 hiter layout. Go 1.24+ runtime
// provides a linkname compatibility shim whose real iterator pointer lands in H.
type GoMapIterator struct {
	K           unsafe.Pointer
	V           unsafe.Pointer
	T           *GoMapType
	H           unsafe.Pointer
	Buckets     unsafe.Pointer
	Bptr        *unsafe.Pointer
	Overflow    *[]unsafe.Pointer
	OldOverflow *[]unsafe.Pointer
	StartBucket uintptr
	Offset      uint8
	Wrapped     bool
	B           uint8
	I           uint8
	Bucket      uintptr
	CheckBucket uintptr
	// Required by the Go 1.24-1.25 GOEXPERIMENT=noswissmap hiter layout.
	// Newer legacy linkname shims ignore this trailing storage.
	ClearSeq uint64
}
