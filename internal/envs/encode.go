package envs

import (
	"os"
)

var UseEncodeVM = os.Getenv("SONIC_USE_ENCODE_VM") == "1" 