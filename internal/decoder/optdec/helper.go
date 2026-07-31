package optdec

import (
	"strconv"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/utils"
)

func SkipNumberFast(json string, start int) (int, bool) {
	// find the number ending, we parsed in native, it always valid
	pos := start
	for pos < len(json) && json[pos] != ']' && json[pos] != '}' && json[pos] != ',' {
		if json[pos] >= '0' && json[pos] <= '9' || json[pos] == '.' || json[pos] == '-' || json[pos] == '+' || json[pos] == 'e' || json[pos] == 'E' {
			pos += 1
		} else {
			break
		}
	}

	// if not found number, return false
	if pos == start {
		return pos, false
	}
	return pos, true
}

// pos is the start index of the raw
func ValidNumberFast(raw string) bool {
	ret := utils.SkipNumber(raw, 0)
	if ret < 0 {
		return false
	}

	// check trailing chars
	return ret >= len(raw)
}

func SkipOneFast(json string, pos int) (string, error) {
	start := native.SkipOneFast(&json, &pos)
	if start < 0 {
		return "", error_syntax(pos, json, types.ParsingError(-start).Error())
	}

	return json[start:pos], nil
}

func ParseI64(raw string) (int64, error) {
	return strconv.ParseInt(raw, 10, 64)
}

func ParseBool(raw string) (bool, error) {
	switch raw {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, &strconv.NumError{Func: "ParseBool", Num: raw, Err: strconv.ErrSyntax}
	}
}

func ParseU64(raw string) (uint64, error) {
	return strconv.ParseUint(raw, 10, 64)
}

func ParseF64(raw string) (float64, error) {
	return strconv.ParseFloat(raw, 64)
}

func Unquote(raw string) (string, error) {
	return strconv.Unquote(raw)
}
