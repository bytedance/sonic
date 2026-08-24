package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_validate_utf8_fast func(s unsafe.Pointer)  (ret int)

var S_validate_utf8_fast uintptr

//go:nosplit
func validate_utf8_fast(s *string)  (ret int) {
    return F_validate_utf8_fast(rt.NoEscape(unsafe.Pointer(s)))
}
