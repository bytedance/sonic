//go:build darwin && arm64
// +build darwin,arm64

package link

import (
	_ "embed"
)

//go:embed libsonic_rs_aarch64-apple-darwin.dylib
var sonic_rs_blob []byte