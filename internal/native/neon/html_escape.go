package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_html_escape func(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer) (ret int)

var S_html_escape uintptr

//go:nosplit
func html_escape(sp unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) (ret int) {
    return F_html_escape(rt.NoEscape(sp), nb, dp, rt.NoEscape(unsafe.Pointer(dn)))
}
