//go:build linux && arm64
// +build linux,arm64

package link

import (
	_ "embed"
)

//go:embed libsonic_rs_aarch64-unknown-linux-gnu.so
var sonic_rs_blob []byte