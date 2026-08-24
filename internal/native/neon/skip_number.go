package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_skip_number func(s unsafe.Pointer, p unsafe.Pointer) (ret int)

var S_skip_number uintptr

//go:nosplit
func skip_number(s *string, p *int) (ret int) {
    return F_skip_number(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}
