package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_skip_one_fast func(s unsafe.Pointer, p unsafe.Pointer) (ret int)

var S_skip_one_fast uintptr

//go:nosplit
func skip_one_fast(s *string, p *int) (ret int) {
    return F_skip_one_fast(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}

