//go:build darwin && amd64
// +build darwin,amd64

package link

import (
	_ "embed"
)

// TODO: fixme
//go:embed libsonic_rs_aarch64-apple-darwin.dylib
var sonic_rs_blob []byte