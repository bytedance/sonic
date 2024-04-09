//go:build linux && amd64
// +build linux,amd64

package link

import (
	_ "embed"
)

//go:embed libsonic_rs_x86_64-unknown-linux-gnu.so
var sonic_rs_blob []byte