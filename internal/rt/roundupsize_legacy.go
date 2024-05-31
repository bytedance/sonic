//go:build go1.16 && !go1.22
// +build go1.16,!go1.22

package rt

import (
	_ "unsafe"
)

//go:linkname roundupsize  runtime.roundupsize
func roundupsize(size uintptr) uintptr
