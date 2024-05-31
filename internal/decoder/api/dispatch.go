package api

import (
	"os"
)

func useDecOpt() bool {
	return os.Getenv("SONIC_USE_OPTDEC") == "1" 
}
