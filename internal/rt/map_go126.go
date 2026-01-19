//go:build go1.26
// +build go1.26

package rt

import (
	"unsafe"
)

// GoMapIterator for Go 1.26+ where Swiss maps are the default.
// This matches the internal runtime iterator structure for Swiss maps.
type GoMapIterator struct {
	K  unsafe.Pointer
	V  unsafe.Pointer
	T  *GoMapType
	It unsafe.Pointer
}
