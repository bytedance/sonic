// +build cgo

package link

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"
import "unsafe"

func DlOpen(filename string, flags int32) unsafe.Pointer {
	handle := C.dlopen(C.CString(filename), C.int(flags))
	if handle == nil {
		panic("dlopen failed")
	}
	return handle
}

func DlClose(handle unsafe.Pointer) int32 {
	ret := int32(C.dlclose(handle))
	if ret != 0 {
		panic("dlcolse failed")
	}
	return ret
}

func DlSym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	addr := C.dlsym(handle, C.CString(symbol))
	if addr == nil {
		panic(DlError())
	}
	return addr
}

func DlError() (err string) {
	cerr := C.dlerror()
	if cerr != nil {
		err = C.GoString(cerr)
	}
	return
}