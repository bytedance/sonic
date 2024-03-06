package link

import (
	"fmt"
	"unsafe"
	"os"
)

func Link(blob []byte, names []string) []unsafe.Pointer {
	tmpFile, err := os.CreateTemp("/tmp", "libsonic-blob*.so")
	defer os.Remove(tmpFile.Name())
	if err != nil {
		fmt.Println("unable to create temp files:", err)
		return nil
	}

	if _, err := tmpFile.Write(blob); err != nil {
		fmt.Println("unable to write temp file:", err)
		return nil
	}
	handle := DlOpen(tmpFile.Name(), 2 /* RTLD_NOW */)

	cfuncs := make([]unsafe.Pointer, len(names))
	for i, sym := range(names) {
		cfuncs[i] =  DlSym(handle, sym)
	}
	return cfuncs
}

func AddPcSp()