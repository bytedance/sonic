package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)
var F_quote func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer, flags uint64) (ret int)

var S_quote uintptr

//go:nosplit
func quote(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) (ret int) {
    return F_quote(rt.NoEscape(sp), nb, rt.NoEscape(dp), rt.NoEscape(unsafe.Pointer(dn)), flags)
}
