package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_skip_object func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) (ret int)

var S_skip_object uintptr

//go:nosplit
func skip_object(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
    return F_skip_object(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}
