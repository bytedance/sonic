package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_vstring func(s unsafe.Pointer, p unsafe.Pointer, v unsafe.Pointer, flags uint64)

var S_vstring uintptr

//go:nosplit
func vstring(s *string, p *int, v *types.JsonState, flags uint64) {
    F_vstring(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(v)), flags)
}
