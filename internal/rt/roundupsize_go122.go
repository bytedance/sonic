//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package rt

import (
	_ "unsafe"
)

//go:linkname rt_roundupsize  runtime.roundupsize
func rt_roundupsize(size uintptr, noscan bool) uintptr

// a wrapper for fastmap
func roundupsize(size uintptr) uintptr {
	return rt_roundupsize(size, MapType(MapEfaceType).Bucket.PtrData == 0)
}
