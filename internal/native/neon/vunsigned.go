package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_vunsigned func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer)

var S_vunsigned uintptr

//go:nosplit
func vunsigned(s *string, p *int, v *types.JsonState) {
    F_vunsigned(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)))
}
