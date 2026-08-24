package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var  F_skip_one func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) (ret int)

var S_skip_one uintptr

//go:nosplit
func skip_one(s *string, p *int, m *types.StateMachine, flags uint64) (ret int) {
    return F_skip_one(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}
