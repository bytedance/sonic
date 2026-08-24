package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_vnumber func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

var S_vnumber uintptr

//go:nosplit
func vnumber(s *string, p *int, v *types.JsonState) {
    F_vnumber(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}
