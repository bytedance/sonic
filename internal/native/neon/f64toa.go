package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_f64toa func(out unsafe.Pointer, val float64) (ret int) 

var S_f64toa uintptr

//go:nosplit
func f64toa(out *byte, val float64) (ret int) {
	return F_f64toa((rt.NoEscape(unsafe.Pointer(out))), val)
}

