package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_f32toa func(out unsafe.Pointer, val float32) (ret int) 

var S_f32toa uintptr

//go:nosplit
func f32toa(out *byte, val float32) (ret int) {
    return F_f32toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

