package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_vsigned func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

var S_vsigned uintptr

//go:nosplit
func vsigned(s *string, p *int, v *types.JsonState) {
    F_vsigned(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}
