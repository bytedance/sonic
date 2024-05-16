package ast

import (
	"fmt"

	"github.com/bytedance/sonic/internal/encoder/alg"
	"github.com/bytedance/sonic/internal/native/types"
	uq "github.com/bytedance/sonic/unquote"
	"github.com/cloudwego/base64x"
)

// NOTICE: must assume pos is at ']' or '}'
func backward(json string, pos int) (stop int, empty bool) {
    if json[pos] != ']' && json[pos] != '}' {
        panic(fmt.Sprintf("char '%c' at stop %d is not close dilimeter of JSON", json[pos], pos))
    }
    stop = pos - 1
    for ; stop >= 0 && isSpace(json[stop]); stop-- {
    }
    empty = (json[stop] == '[' || json[stop] == '{')
    stop += 1
    return
}

func decodeBase64(src string) ([]byte, error) {
    return base64x.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64x.StdEncoding.EncodeToString(src)
}

func unquote(src string) (string, types.ParsingError) {
    return uq.String(src)
}

// [2,"a"],1 => {"a":1}
// ["a",2],1  => "a":[1]
func makePathAndValue(b []byte, path []interface{}, allowAppend bool, val string) ([]byte, error) {
	for i, k := range path {
		if key, ok := k.(string); ok {
			b = alg.Quote(b, key, false)
			b = append(b, ":"...)
		}
		if i == len(path)-1 {
			b = append(b, val...)
			break
		}
		n := path[i+1]
		if _, ok := n.(int); ok {
			if !allowAppend {
				return nil, ErrNotExist
			}
			b = append(b, "["...)
		} else if _, ok := n.(string); ok {
			b = append(b, `{`...)
		} else {
			panic("path must be either int or string")
		}
	}
	for i := len(path) - 1; i >= 1; i-- {
		k := path[i]
		if _, ok := k.(int); ok {
			b = append(b, "]"...)
		} else if _, ok := k.(string); ok {
			b = append(b, `}`...)
		}
	}
	return b, nil
}

// func castToInt64(v interface{}) (int64, bool) {
// 	switch vv := v.(type) {
// 	case json.Number:
// 		iv, err := vv.Int64()
// 		if err != nil {
// 			return 0, false
// 		}
// 		return iv, true
// 	case int:
// 		return int64(vv), true
// 	case uint:
// 		return int64(vv), true
// 	case int64:
// 		return vv, true
// 	case uint64:
// 		return int64(vv), true
// 	case int32:
// 		return int64(vv), true
// 	case uint16:
// 		return int64(vv), true
// 	case int16:
// 		return int64(vv), true
// 	case uint8:
// 		return int64(vv), true
// 	case int8:
// 		return int64(vv), true
// 	default:
// 		return 0, false
// 	}
// }

// func castToFloat64(v interface{}) (float64, bool) {
// 	switch vv := v.(type) {
// 	case json.Number:
// 		iv, err := vv.Float64()
// 		if err != nil {
// 			return 0, false
// 		}
// 		return iv, true
// 	case float32:
// 		return float64(vv), true
// 	case float64:
// 		return float64(vv), true
// 	case int:
// 		return float64(vv), true
// 	case uint:
// 		return float64(vv), true
// 	case int64:
// 		return float64(vv), true
// 	case uint64:
// 		return float64(vv), true
// 	case int32:
// 		return float64(vv), true
// 	case uint16:
// 		return float64(vv), true
// 	case int16:
// 		return float64(vv), true
// 	case uint8:
// 		return float64(vv), true
// 	case int8:
// 		return float64(vv), true
// 	default:
// 		return 0, false
// 	}
// }

// func castToNumber(v interface{}) (json.Number, bool) {
// 	switch vv := v.(type) {
// 	case json.Number:
// 		return vv, true
// 	case float32:
// 		return json.Number(strconv.FormatFloat(float64(vv), 'g', -1, 32)), true
// 	case float64:
// 		return json.Number(strconv.FormatFloat(float64(vv), 'g', -1, 32)), true
// 	case int:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case uint:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case int64:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case uint64:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case int32:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case uint16:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case int16:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case uint8:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	case int8:
// 		return json.Number(strconv.FormatInt(int64(vv), 10)), true
// 	default:
// 		return "", false
// 	}
// }

// func castToString(v interface{}) (string, bool) {
// 	switch vv := v.(type) {
// 	case string:
// 		return vv, true
//     case []byte:
//         return string(vv), false
// 	default:
// 		return "", false
// 	}
// }

// func castToMap(v interface{}, buf map[string]Node) bool {
// 	switch vv := v.(type) {
// 	case map[string]Node:
//         for k, v := range vv {
//             buf[k] = v
//         }
// 		return true
//     case map[string]interface{}:
//         for k, v := range vv {
//             buf[k] = NewAny(v)
//         }
//         return true
// 	default:
// 		return false
// 	}
// }

// func castToArray(v interface{}, buf *[]Node) bool {
// 	switch vv := v.(type) {
// 	case []Node:
// 		*buf = append(*buf, vv...)
//         return true
//     case []interface{}:
//         for _, v := range vv {
//             *buf = append(*buf, NewAny(v))
//         }
//         return true
// 	default:
// 		return false
// 	}
// }