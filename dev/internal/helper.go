package internal

import (
	"encoding/json"
	"strconv"

	"github.com/bytedance/sonic/ast"
)

func SkipNumberFast(json string, start int) (int, error) {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	pos := start
	for pos < len(json) && json[pos] != ']' && json[pos] != '}' && json[pos] != ',' {
		if json[pos] >= '0' && json[pos] <= '9' || json[pos] == '.' || json[pos] == '-' || json[pos] == '+' || json[pos] == 'e' || json[pos] == 'E' {
			pos += 1
		} else {
			return pos, newError("invalid number", int64(pos))
		}
	}
	return pos, nil
}

func skipOneFast(json string, start int) string {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	nast, err := ast.NewSearcher(json[start:]).GetByPath()
	if err != nil {
		println(json[start:])
		panic("json should always be valid here")
	}
	raw, err := nast.Raw()
	if err != nil {
		panic("json should always be valid here")
	}
	return raw
}

func decodeBase64(raw string) ([]byte, error) {
	var ret []byte
	err := json.Unmarshal([]byte("\""+raw+"\""), &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func ParseI64(raw string) (int64, error) {
	i64, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return i64, nil
}

func ParseBool(raw string) (bool, error) {
	var b bool
	err := json.Unmarshal([]byte(raw), &b)
	if err != nil {
		return false, err
	}
	return b, nil
}

func ParseU64(raw string) (uint64, error) {
	u64, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return u64, nil
}

func ParseF64(raw string) (float64, error) {
	f64, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, err
	}
	return f64, nil
}

func Unquote(raw string) (string, error) {
	var u string
	err := json.Unmarshal([]byte(raw), &u)
	if err != nil {
		return "", err
	}
	return u, nil
}
