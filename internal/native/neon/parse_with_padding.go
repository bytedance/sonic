package neon

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_parse_with_padding func(parser unsafe.Pointer) (ret int)

var S_parse_with_padding uintptr

//go:nosplit
func parse_with_padding(parser unsafe.Pointer) (ret int) {
    return F_parse_with_padding(rt.NoEscape(parser))
}
