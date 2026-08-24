package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_validate_utf8 func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer) (ret int)

var S_validate_utf8 uintptr

//go:nosplit
func validate_utf8(s *string, p *int, m *types.StateMachine) (ret int) {
    return F_validate_utf8(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)))
}


