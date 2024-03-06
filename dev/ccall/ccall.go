package ccall

import (
	"unsafe"
)

//go:nosplit
func C_systemstack_1arg(arg0 uintptr, callback unsafe.Pointer) uintptr {
	return __c_systemstack_1arg(arg0, callback)

	// TODO: keepalive
}

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __c_systemstack_1arg(arg0 uintptr, callback unsafe.Pointer) uintptr