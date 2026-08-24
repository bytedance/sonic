package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_lspace func(sp unsafe.Pointer, nb int, off int) (ret int)

var S_lspace uintptr

//go:nosplit
func lspace(sp *byte, nb int, off int) (ret int) {
    return F_lspace(rt.NoEscape(unsafe.Pointer(sp)), nb, off)
}

