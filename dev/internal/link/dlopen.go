// +build !cgo

package link

import "unsafe"

func DlOpen(filename string, flags int32) unsafe.Pointer {
	panic("must enable cgo")
}

func DlClose(handle unsafe.Pointer) int32 {
	panic("must enable cgo")
}

func DlSym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	panic("must enable cgo")
}

func DlError() (err string) {
	panic("must enable cgo")
}