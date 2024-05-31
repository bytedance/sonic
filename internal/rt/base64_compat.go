// +build !amd64 !go1.16 go1.23

package rt

import (
	"encoding/base64"
)

func DecodeBase64(raw []byte) ([]byte, error) {
	ret := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	n, err := base64.StdEncoding.Decode(ret, raw)
	if err != nil {
		return nil, err
	}
	return ret[:n], nil
}
