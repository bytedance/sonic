package ast

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/bytedance/sonic/internal/native/types"
)

// RawNode represents a raw json value or error
type RawNode struct {
	err error
	js string
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
	t := switchRawType(self.js[0])
	return int(t)
}

func (self RawNode) Exists() bool {
	return self.err == nil
}

func (self RawNode) itype() types.ValueType {
	return switchRawType(self.js[0])
}

// Error returns error message if the node is invalid
func (self RawNode) Error() string {
	if self.err != nil {
		return self.err.Error()
	}
	return ""
}

// Check checks if the node itself is valid, and return
func (self RawNode) Check() error {
	return self.err
}


// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self RawNode) GetByPath(path ...interface{}) RawNode {
	if self.err != nil {
		return self
	}
	p := Parser{s: self.js}
	s, e := p.getByPath(path...)
	if e != 0 {
		err := p.ExportError(e)
		return RawNode{err: err}
	}
    return RawNode{js: self.js[s:p.p]}
}


// Get loads given key of an object node on demands
func (self RawNode) Get(key string) RawNode {
	if self.err != nil {
		return self
	}
    p := Parser{s: self.js}
	s, e := p.getByPath(key)
	if e != 0 {
		err := p.ExportError(e)
		return RawNode{err: err}
	}
    return RawNode{js: self.js[s:p.p]}
}

// Index indexies node at given idx
func (self RawNode) Index(idx int) RawNode {
	if self.err != nil {
		return self
	}
	p := Parser{s: self.js}
	s, e := p.getByPath(idx)
	if e != 0 {
		err := p.ExportError(e)
		return RawNode{err: err}
	}
    return RawNode{js: self.js[s:p.p]}
}

func (self RawNode) str() string {
	return self.js[1:len(self.js)-1]
}

// Raw returns json representation of the node
func (self RawNode) Raw() (string, error) {
    if self.err != nil {
        return "", self.err
    }
    return self.js, nil
}

// Bool returns bool value represented by this node, 
// including types.V_TRUE|V_FALSE|V_NUMBER|V_STRING|V_ANY|V_NULL
func (self RawNode) Bool() (bool, error) {
	if self.err != nil {
		return false, self.err
	}
	p := Parser{s: self.js}
	switch val := p.decodeValue(); val.Vt {
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
	if self.err != nil {
		return 0, self.err
	}
	p := Parser{s: self.js}
	switch val := p.decodeValue(); val.Vt {
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
	if self.err != nil {
		return 0, self.err
	}
	p := Parser{s: self.js}
	switch val := p.decodeValue(); val.Vt {
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
	if self.err != nil {
		return "", self.err
	}
	p := Parser{s: self.js}
	switch val := p.decodeValue(); val.Vt {
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
	if self.err != nil {
		return "", self.err
	}
	p := Parser{s: self.js}
	switch val := p.decodeValue(); val.Vt {
		case types.V_NULL    : return "", nil
		case types.V_TRUE    : return "true", nil
		case types.V_FALSE   : return "false", nil
		case types.V_STRING  : 
			n, e := p.decodeString(val.Iv, val.Ep)
			if e != 0 {
				return "", p.ExportError(e)
			}
			return addr2str(n.p, n.v), nil
		case types.V_DOUBLE  : return strconv.FormatFloat(val.Dv, 'g', -1, 64), nil
		case types.V_INTEGER : return strconv.FormatInt(val.Iv, 10), nil
		default              : return "", types.ParsingError(-val.Vt)
	} 
}

// ArrayUseNode copys both parsed and non-parsed chidren nodes, 
// and indexes them by original order
func (self RawNode) ArrayUseNode() (ret []RawNode, err error) {
	ret = make([]RawNode, 0, _DEFAULT_NODE_CAP)
	err = self.skipAllIndex(&ret)
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
	err = self.skipAllKey(&ret)
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
	if self.err != nil {
		return 0, self.err
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
func (self RawNode) ForEach(sc func(index int, key string, node RawNode) bool) error {
	if self.err != nil {
		return self.err
	}
    switch self.itype() {
    case types.V_ARRAY:
        ns := rawNodesPool.Get().(*[]RawNode) 
        if err := self.skipAllIndex(ns); err != nil {
			return err
		}
		for i, n := range *ns {
			if !sc(i, "", n) {
				break
			}
		}
		*ns = (*ns)[:0]
		rawNodesPool.Put(ns)
		return nil
	case types.V_OBJECT:
		ps := rawPairsPool.Get().(*[]RawPair)
		if err := self.skipAllKey(ps); err != nil {
			return err
		}
		for i, p := range *ps {
			if !sc(i, p.Key, p.Value) {
				break
			}
		}
		*ps = (*ps)[:0]
		rawPairsPool.Put(ps)
		return nil
    default:
        return ErrUnsupportType
    }
}


func (self RawNode) skipAllIndex(ret *[]RawNode) error {
    parser := Parser{s: self.js}
    parser.skipValue = true
    parser.noLazy = true
    _, err := parser.decodeArray(nil, ret)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

// RawPair is a pair of key and value (RawNode)
type RawPair struct {
    Key   string
    Value RawNode
}

func (self RawNode) skipAllKey(ret *[]RawPair) error {
    parser := Parser{s: self.js}
    parser.skipValue = true
    parser.noLazy = true
    _, err := parser.decodeObject(nil, ret)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

var rawNodesPool = sync.Pool{
	New: func() interface{} {
		ret := make([]RawNode, 0, _DEFAULT_NODE_CAP)
		return &ret
	},
}

var rawPairsPool = sync.Pool{
	New: func() interface{} {
		ret := make([]RawPair, 0, _DEFAULT_NODE_CAP)
		return &ret
	},
}

func (self Searcher) GetRawByPath(path ...interface{}) (RawNode, error) {
    var err types.ParsingError
    var start int

    self.parser.p = 0
    start, err = self.parser.getByPath(path...)
    if err != 0 {
        return RawNode{}, self.parser.syntaxError(err)
    }

    t := switchRawType(self.parser.s[start])
    if t == _V_NONE {
        return RawNode{}, self.parser.ExportError(err)
    }
    return RawNode{js: self.parser.s[start:self.parser.p]}, nil
}