package ast

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bytedance/sonic/internal/native/types"
)

// RawNode represents a raw json value or error
type RawNode struct {
	t  int
	js string
}

func NewRawNode(js string) RawNode {
	s := NewParser(js).lspace(0)
	if s > len(js) {
		return errRawNode(types.ERR_EOF)
	}
	return rawNode(js[s:])
}

func rawNode(js string) RawNode {
	return RawNode{
		t: int(switchRawType(js[0])),
		js: js,
	}
}

// Type returns json type represented by the node
// It will be one of belows:
//    V_NONE   = 0 (empty node)
//    V_NULL   = 2 (json value `null`)
//    V_TRUE   = 3 (json value `true`)
//    V_FALSE  = 4 (json value `false`)
//    V_ARRAY  = 5 (json value array)
//    V_OBJECT = 6 (json value object)
//    V_STRING = 7 (json value string)
//    V_NUMBER = 33 (json value number )
func (self RawNode) Type() int {
	return self.t
}

func (self RawNode) Exists() bool {
	return self.t != 0 && self.t != V_ERROR
}

func (self RawNode) itype() types.ValueType {
	return types.ValueType(self.t)
}

// Error returns error message if the node is invalid
func (self RawNode) Error() string {
	if self.t == V_ERROR {
		return self.js
	}
	return ""
}

// Check checks if the node itself is valid, and return
func (self RawNode) Check() error {
	if self.t == V_ERROR {
		return errors.New(self.js)
	}
	return nil
}


// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self RawNode) GetByPath(path ...interface{}) RawNode {
	if self.Check() != nil {
		return self
	}
	p := NewParserObj(self.js)
	s, e := p.getByPath(path...)
	if e != 0 {
		return errRawNode(p.ExportError(e))
	}
    return rawNode(self.js[s:p.p])
}

func errRawNode(err error) RawNode {
	return RawNode{t: V_ERROR, js: err.Error()}
}


// Get loads given key of an object node on demands
func (self RawNode) Get(key string) RawNode {
	if self.Check() != nil {
		return self
	}
    p := NewParserObj(self.js)
	s, e := p.getByPath(key)
	if e != 0 {
		return errRawNode(p.ExportError(e))
	}
    return rawNode(self.js[s:p.p])
}

// Index indexies node at given idx
func (self RawNode) Index(idx int) RawNode {
	if self.Check() != nil {
		return self
	}
	p := NewParserObj(self.js)
	s, e := p.getByPath(idx)
	if e != 0 {
		return errRawNode(p.ExportError(e))
	}
    return rawNode(self.js[s:p.p])
}

func (self RawNode) str() string {
	return self.js[1:len(self.js)-1]
}

// Raw returns json representation of the node
func (self RawNode) Raw() (string, error) {
    if e := self.Check(); e != nil {
        return "", e
    }
    return self.js, nil
}

// Bool returns bool value represented by this node, 
// including types.V_TRUE|V_FALSE|V_NUMBER|V_STRING|V_ANY|V_NULL
func (self RawNode) Bool() (bool, error) {
	if e := self.Check(); e != nil {
		return false, e
	}
	p := NewParserObj(self.js)
	p.decodeNumber(true)
	val := p.decodeValue()
	p.decodeNumber(false)
	switch val.Vt {
        case types.V_NULL    : return false, nil
        case types.V_TRUE    : return true, nil
        case types.V_FALSE   : return false, nil
        case types.V_STRING  : return strconv.ParseBool(self.str())
        case types.V_DOUBLE  : return val.Dv == 0, nil
        case types.V_INTEGER : return val.Iv == 0, nil
        default              : return false, types.ParsingError(-val.Vt)
    } 
}

// Int64 casts the node to int64 value, 
// including V_NUMBER|V_TRUE|V_FALSE|V_STRING
func (self RawNode) Int64() (int64, error) {
	if e := self.Check(); e != nil {
		return 0, e
	}
	p := NewParserObj(self.js)
	p.decodeNumber(true)
	val := p.decodeValue()
	p.decodeNumber(false)
	switch val.Vt {
        case types.V_NULL    : return 0, nil
        case types.V_TRUE    : return 1, nil
        case types.V_FALSE   : return 0, nil
        case types.V_STRING  : return json.Number(self.str()).Int64()
        case types.V_DOUBLE  : return int64(val.Dv), nil
        case types.V_INTEGER : return int64(val.Iv), nil
        default              : return 0, types.ParsingError(-val.Vt)
    } 
}

// Float64 cast node to float64, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL
func (self RawNode) Float64() (float64, error) {
	if e := self.Check(); e != nil {
		return 0, e
	}
	p := NewParserObj(self.js)
	p.decodeNumber(true)
	val := p.decodeValue()
	p.decodeNumber(false)
	switch val.Vt {
        case types.V_NULL    : return 0, nil
        case types.V_TRUE    : return 1, nil
        case types.V_FALSE   : return 0, nil
        case types.V_STRING  : return json.Number(self.str()).Float64()
        case types.V_DOUBLE  : return float64(val.Dv), nil
        case types.V_INTEGER : return float64(val.Iv), nil
        default              : return 0, types.ParsingError(-val.Vt)
    } 
}

// Number casts node to float64, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL,
func (self RawNode) Number() (json.Number, error) {
	if e := self.Check(); e != nil {
		return "", e
	}
	p := NewParserObj(self.js)
	p.decodeNumber(true)
	val := p.decodeValue()
	p.decodeNumber(false)
	switch val.Vt {
        case types.V_NULL    : return json.Number("0"), nil
        case types.V_TRUE    : return json.Number("1"), nil
        case types.V_FALSE   : return json.Number("0"), nil
        case types.V_STRING  : return json.Number(self.str()), nil
        case types.V_DOUBLE  : return json.Number(self.js), nil
        case types.V_INTEGER : return json.Number(self.js), nil
        default              : return "", types.ParsingError(-val.Vt)
    } 
}

// String cast node to string, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL
func (self RawNode) String() (string, error) {
	if e := self.Check(); e != nil {
		return "", e
	}
	p := NewParserObj(self.js)
	p.decodeNumber(true)
	val := p.decodeValue()
	p.decodeNumber(false)
	switch val.Vt {
		case types.V_NULL    : return "", nil
		case types.V_TRUE    : return "true", nil
		case types.V_FALSE   : return "false", nil
		case types.V_STRING  : 
			n, e := p.decodeString(val.Iv, val.Ep)
			if e != 0 {
				return "", p.ExportError(e)
			}
			return n.toString(), nil
		case types.V_DOUBLE  : return strconv.FormatFloat(val.Dv, 'g', -1, 64), nil
		case types.V_INTEGER : return strconv.FormatInt(val.Iv, 10), nil
		default              : return "", types.ParsingError(-val.Vt)
	} 
}

// ArrayUseNode copys both parsed and non-parsed chidren nodes, 
// and indexes them by original order
func (self RawNode) ArrayUseNode() (ret []RawNode, err error) {
	ret = make([]RawNode, 0, _DEFAULT_NODE_CAP)
	err = self.ForEachElem(func(i int, node RawNode) bool {
		ret = append(ret, node)
		return true
	})
    return ret, err
}

// Array loads all indexes of an array node
func (self RawNode) Array() (ret []interface{}, err error) {
	node := NewRaw(self.js)
	return node.Array()
}

// ObjectUseNode scans both parsed and non-parsed chidren nodes, 
// and map them by their keys
func (self RawNode) MapUseNode() (ret []RawPair, err error) {
	ret = make([]RawPair, 0, _DEFAULT_NODE_CAP)
	err = self.ForEachKV(func(key string, node RawNode) bool {
		ret = append(ret, RawPair{key, node})
		return true
	})
    return ret, err
}

// Map loads all keys of an object node
func (self RawNode) Map() (ret map[string]interface{}, err error) {
	node := NewRaw(self.js)
	return node.Map()
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic type.
// WARN: all numberic nodes are casted to float64
func (self RawNode) Interface() (interface{}, error) {
	if e := self.Check(); e != nil {
		return nil, e
	}
	switch self.itype() {
	case types.V_OBJECT:
		return self.Map()
	case types.V_ARRAY:
		return self.Array()
	case types.V_STRING:
		return self.str(), nil
	case _V_NUMBER:
		return self.Float64()
	case types.V_TRUE:
		return true, nil
	case types.V_FALSE:
		return false, nil
	case types.V_NULL:
		return nil, nil
	default:
		return nil, ErrUnsupportType
	}
}


// ForEach scans one V_OBJECT node's children from JSON head to tail
func (self RawNode) ForEachKV(sc func(key string, node RawNode) bool) error {
	if e := self.Check(); e != nil {
		return e
	}
    switch self.itype() {
	case types.V_OBJECT:

		p := NewParser(self.js)
		if empty, err := p.objectBegin(); err != 0 {
			return err
		} else if empty {
			return nil
		} 

		for {
			k, e := p.key()
			if e != 0 {
				return e
			}
			s, e := p.skipFast()
			if e != 0 {
				return e
			}
			n := rawNode(self.js[s:p.p])
			if !sc(k, n) {
				return nil
			}
			if end, e := p.objectEnd(); e != 0 {
				return e
			} else if end {
				return nil
			}
		}
		
    default:
        return ErrUnsupportType
    }
}

// ForEach scans one V_OBJECT node's children from JSON head to tail
func (self RawNode) ForEachElem(sc func(i int, node RawNode) bool) error {
	if e := self.Check(); e != nil {
		return e
	}
    switch self.itype() {
    case types.V_ARRAY:
		p := NewParser(self.js)
		if empty, err := p.arrayBegin(); err != 0 {
			return err
		} else if empty {
			return nil
		} 

		i := 0
		for {
			s, e := p.skipFast()
			if e != 0 {
				return e
			}
			n := rawNode(self.js[s:p.p])
			if !sc(i, n) {
				return nil
			}
			i++
			if end, e := p.arrayEnd(); e != 0 {
				return e
			} else if end {
				return nil
			}
		}
    default:
        return ErrUnsupportType
    }
}

// RawPair is a pair of key and value (RawNode)
type RawPair struct {
    Key   string
    Value RawNode
}

// GetRawByPath
func (self Searcher) GetRawByPath(path ...interface{}) (RawNode, error) {
	if self.parser.s == "" {
		err := errors.New("empty input")
		return errRawNode(err), err
	}

    self.parser.p = 0
    s, err := self.parser.getByPath(path...)
    if err != 0 {
		e := self.parser.ExportError(err)
        return errRawNode(e), e
    }

    t := switchRawType(self.parser.s[s])
    if t == _V_NONE {
		e := self.parser.ExportError(err)
        return errRawNode(e), e 
    }
    return RawNode{int(t), self.parser.s[s:self.parser.p]}, nil
}