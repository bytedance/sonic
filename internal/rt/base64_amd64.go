// +build amd64,go1.17,!go1.24

package rt

import (
	"github.com/cloudwego/base64x"
)

func DecodeBase64(raw []byte) ([]byte, error) {
	ret := make([]byte, base64x.StdEncoding.DecodedLen(len(raw)))
	n, err := base64x.StdEncoding.Decode(ret, raw)
	if err != nil {
		return nil, err
	}
	return ret[:n], nil
}

func EncodeBase64(src []byte) string {
    return base64x.StdEncoding.EncodeToString(src)
}
