package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_u64toa func(out unsafe.Pointer, val uint64) (ret int)

var S_u64toa uintptr

//go:nosplit
func u64toa(out *byte, val uint64) (ret int) {
    return F_u64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}
