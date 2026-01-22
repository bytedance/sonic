// +build !amd64,!arm64 go1.27 !go1.17 arm64,!go1.20

package compat

import (
    "fmt"
    "os"
)

func Warn(prefix string) {
    fmt.Fprintf(os.Stderr, "WARNING: %s only supports (go1.17~1.26 and amd64 CPU) or (go1.20~1.26 and arm64 CPU), but your environment is not suitable and will fallback to encoding/json\n", prefix)
}
