package optdec

import (
	"encoding/json"
	"strconv"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
)


func SkipNumberFast(json string, start int) (int, error) {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	pos := start
	for pos < len(json) && json[pos] != ']' && json[pos] != '}' && json[pos] != ',' {
		if json[pos] >= '0' && json[pos] <= '9' || json[pos] == '.' || json[pos] == '-' || json[pos] == '+' || json[pos] == 'e' || json[pos] == 'E' {
			pos += 1
		} else {
			return pos, error_syntax(pos, json, "invalid number")
		}
	}
	return pos, nil
}

func ValidNumberFast(json string) error {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	pos := 0
	for pos < len(json) && json[pos] != ']' && json[pos] != '}' && json[pos] != ',' {
		if json[pos] >= '0' && json[pos] <= '9' || json[pos] == '.' || json[pos] == '-' || json[pos] == '+' || json[pos] == 'e' || json[pos] == 'E' {
			pos += 1
		} else {
			return error_syntax(pos, json, "invalid number")
		}
	}

	if pos == 0 {
		return error_syntax(pos, json, "invalid number")
	}
	return nil
}

func SkipOneFast2(json string, pos *int) (int, error) {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	start := native.SkipOneFast(&json, pos)
	if start < 0 {
		return -1, error_syntax(*pos, json, types.ParsingError(-start).Error())
	}
	return start, nil
}

func SkipOneFast(json string, pos int) (string, error) {
	// find the number ending, we pasred in sonic-cpp, it alway valid
	start := native.SkipOneFast(&json, &pos)
	if start < 0 {
		// TODO: details error code
		return "", error_syntax(pos, json, types.ParsingError(-start).Error())
	}
	return json[start:pos], nil
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
