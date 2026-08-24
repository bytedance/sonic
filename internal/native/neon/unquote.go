package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_unquote func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep unsafe.Pointer, flags uint64) (ret int)

var S_unquote uintptr

//go:nosplit
func unquote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) (ret int) {
    return F_unquote(rt.NoEscape(sp), nb, dp, rt.NoEscape(unsafe.Pointer(ep)), flags)
}
