package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_value func(s unsafe.Pointer, n int, p int, v unsafe.Pointer, flags uint64) (ret int)

var S_value uintptr

//go:nosplit
func value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) (ret int) {
    return F_value(rt.NoEscape(unsafe.Pointer(s)), n, p, rt.NoEscape(unsafe.Pointer(v)), flags)
}
