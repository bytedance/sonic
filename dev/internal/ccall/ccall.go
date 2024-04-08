package ccall

import "unsafe"

//go:linkname Cgo_runtime_cgocall runtime.cgocall
//go:noescape
func Cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname Cgo_always_false runtime.cgoAlwaysFalse
var Cgo_always_false bool

//go:linkname Cgo_use runtime.cgoUse
func Cgo_use(interface{})
