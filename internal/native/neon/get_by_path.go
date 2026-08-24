package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_get_by_path func(s unsafe.Pointer, p unsafe.Pointer, path unsafe.Pointer, m unsafe.Pointer) (ret int)

var S_get_by_path uintptr

//go:nosplit
func get_by_path(s *string, p *int, path *[]interface{}, m *types.StateMachine) (ret int) {
    return F_get_by_path(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(path)), rt.NoEscape(unsafe.Pointer(m)))
}
