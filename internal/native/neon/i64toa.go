package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_i64toa func(out unsafe.Pointer, val int64) (ret int)

var S_i64toa uintptr

//go:nosplit
func i64toa(out *byte, val int64) (ret int) {
    return F_i64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

