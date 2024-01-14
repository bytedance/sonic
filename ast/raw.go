package ast

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// Value represents a raw json value or error
type Value struct {
	t  int
	js string
}

func NewValue(js string) Value {
	p := NewParser(js)
	s, e := p.skip()
	if e != 0 {
		return errRawNode(p.ExportError(e))
	}
	return rawNode(js[s:p.p])
}

func rawNode(js string) Value {
	return Value{
		t: int(switchRawType(js[0])),
		js: js,
	}
}

// Type returns json type represented by the node
// It will be one of belows:
//    V_NONE   = 0 (empty node)
//    V_ERROR  = 1 (something wrong)
//    V_NULL   = 2 (json value `null`)
//    V_TRUE   = 3 (json value `true`)
//    V_FALSE  = 4 (json value `false`)
//    V_ARRAY  = 5 (json value array)
//    V_OBJECT = 6 (json value object)
//    V_STRING = 7 (json value string)
//    V_NUMBER = 33 (json value number )
func (self Value) Type() int {
	return self.t
}

func (self Value) Exists() bool {
	return self.t != 0 && self.t != V_ERROR
}

func (self Value) itype() types.ValueType {
	return types.ValueType(self.t)
}

// Error returns error message if the node is invalid
func (self Value) Error() string {
	if self.t == V_ERROR {
		return self.js
	}
	return ""
}

// Check checks if the node itself is valid, and return
func (self Value) Check() error {
	if self.t == V_ERROR {
		return errors.New(self.js)
	}
	return nil
}


// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self Value) GetByPath(path ...interface{}) Value {
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

func errRawNode(err error) Value {
	return Value{t: V_ERROR, js: err.Error()}
}


// Get loads given key of an object node on demands
func (self Value) Get(key string) Value {
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

// Set sets the node of given key under self, and reports if the key has existed.
func (self *Value) Set(key string, val Value) (bool, error) {
	if val.Check() != nil {
		return false, val
	}
	if self.Check() != nil {
		return false, self
	}
	if self.t != V_OBJECT {
		return false, ErrUnsupportType
	}

	in := val.js
	exist := true

    p := NewParserObj(self.js)
	s, e := p.getByPath(key)
	if e != 0 {
		if e == _ERR_NOT_FOUND {
			exist = false
		} else {
			return false, p.ExportError(e)
		}
	}
	
	var b []byte
	// at least size is left + val + right
	n := len(in)+s+(len(self.js)-p.p)

	if !exist {
		// not exist, need write "key":
		var s = p.p-1 // _ERR_NOT_FOUND stop at ']'
		for ; s>=0 && isSpace(self.js[s]); s-- {}
		if self.js[s] == '{' {
			b = make([]byte, 0, 3+len(key)+n)
			b = append(b, self.js[:s+1]...)
		} else {
			// not empty, no need ','
			b = make([]byte, 0, 4+len(key)+n)
			b = append(b, self.js[:s+1]...)
			b = append(b, ","...)
		}
		quote(&b, key)
		b = append(b, ":"...)
	} else {
		// exist
		// canInplace := len(in) <= (p.p - s)
		// if inplace == 1 && canInplace {
		// 	// join and shrink old string
		// 	b := rt.Str2Mem(self.js)
		// 	copy(b[s:], in)
		// 	copy(b[s+len(in):], b[p.p:])
		// 	self.js = rt.Mem2Str(b[:n])
		// 	return exist, nil
		// } else if inplace == 2 && canInplace {
		// 	// join and trucate old string
		// 	b := rt.Str2Mem(self.js)
		// 	copy(b[s:], in)
		// 	rt.WriteChar(&b[s+len(in)], (p.p - s) - len(in), ' ')
		// 	return exist, nil
		// }
		// slow path: allocate new string
		b = make([]byte, 0, n)
		b = append(b, self.js[:s]...)
	}

	// write val
	b = append(b, in...)
	b = append(b, self.js[p.p:]...)
	self.js = rt.Mem2Str(b)

	return exist, nil
}

// Unset REMOVE the node of given key under object parent, and reports if the key has existed.
func (self *Value) Unset(key string) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if self.t != V_OBJECT {
		return false, ErrUnsupportType
	}

    p := NewParserObj(self.js)
	// start pos of "key"
	s, e := p.searchKey(key)
	if e != 0 {
		if e == _ERR_NOT_FOUND {
			return false, nil
		} else {
			return false, p.ExportError(e)
		}
	}

	// end pos of val
	_, e = p.skipFast()
	if e != 0 {
		return true, p.ExportError(e)
	}

	// trailling ',' or '}'
	end, e := p.objectEnd()
	if e != 0 {
		return true, p.ExportError(e)
	}
	if end {
		p.p--
	}
	
	var b []byte
	d := (p.p-s)
	// if inplace == 0 { 
		// allocate new string
		b = make([]byte, len(self.js)-d)
		copy(b, self.js[:s])
		copy(b[s:], self.js[p.p:])
	// } else if inplace == 1 { 
	// 	// shrink string
	// 	b = rt.Str2Mem(self.js)
	// 	copy(b[s:], b[p.p:])
	// 	b = b[:len(b)-d]
	// } else {
	// 	// trucate string
	// 	b = rt.Str2Mem(self.js)
	// 	rt.WriteChar(&b[s], d, ' ')
	// }

	self.js = rt.Mem2Str(b)
	return true, nil
}

// SetByIndex sets the node of given index, and reports if the key has existed.
func (self *Value) SetByIndex(id int, val Value) (bool, error) {
	if val.Check() != nil {
		return false, val
	}
	if self.Check() != nil {
		return false, self
	}
	if self.t != V_ARRAY {
		return false, ErrUnsupportType
	}

	in := val.js
	exist := true

	// try search from raw
    p := NewParserObj(self.js)
	s, e := p.getByPath(id)
	if e != 0 {
		if e == _ERR_NOT_FOUND {
			exist = false
		} else {
			return false, p.ExportError(e)
		}
	}
	
	var b []byte
	// at least size is left + val + right
	n := len(in)+s+(len(self.js)-p.p)
	if !exist {
		// not exist, need write "key":
		var s = p.p-1 // _ERR_NOT_FOUND stop at ']'
		for ; s>=0 && isSpace(self.js[s]); s-- {}
		if self.js[s] == '[' {
			b = make([]byte, 0, n)
			b = append(b, self.js[:s+1]...)
		} else {
			// the container is not empty, need ','
			b = make([]byte, 0, 1+n)
			b = append(b, self.js[:s+1]...)
			b = append(b, ","...)
		}
	} else {
		b = make([]byte, 0, n)
		b = append(b, self.js[:s]...)
	}

	// write val
	b = append(b, in...)
	b = append(b, self.js[p.p:]...)
	self.js = rt.Mem2Str(b)

	return exist, nil
}

// UnsetByIndex REOMVE the node of given index.
func (self *Value) UnsetByIndex(id int) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if self.t != V_ARRAY {
		return false, ErrUnsupportType
	}

	// try search from raw
    p := NewParserObj(self.js)
	s, e := p.getByPath(id)
	if e != 0 {
		if e == _ERR_NOT_FOUND {
			return false, nil
		} else {
			return false, p.ExportError(e)
		}
	}
	
	// trailling ',' or ']'
	end, e := p.arrayEnd()
	if e != 0 {
		return true, p.ExportError(e)
	}
	if end {
		p.p--
	}
	
	b := make([]byte, len(self.js)-(p.p-s))
	copy(b, self.js[:s])
	copy(b[s:], self.js[p.p:])
	self.js = rt.Mem2Str(b)
	return true, nil
}

// Index indexies node at given idx
func (self Value) Index(idx int) Value {
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

func (self Value) str() string {
	return self.js[1:len(self.js)-1]
}

// Raw returns json representation of the node
func (self Value) Raw() (string, error) {
    if e := self.Check(); e != nil {
        return "", e
    }
    return self.js, nil
}

// Bool returns bool value represented by this node, 
// including types.V_TRUE|V_FALSE|V_NUMBER|V_STRING|V_ANY|V_NULL
func (self Value) Bool() (bool, error) {
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
func (self Value) Int64() (int64, error) {
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
func (self Value) Float64() (float64, error) {
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
func (self Value) Number() (json.Number, error) {
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
func (self Value) String() (string, error) {
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
func (self Value) ArrayUseNode() (ret []Value, err error) {
	if self.t != V_ARRAY {
		return nil, ErrUnsupportType
	}
	ret = make([]Value, 0, _DEFAULT_NODE_CAP)
	err = self.ForEachElem(func(i int, node Value) bool {
		ret = append(ret, node)
		return true
	})
    return ret, err
}

// Array loads all indexes of an array node
func (self Value) Array() (ret []interface{}, err error) {
	node := NewRaw(self.js)
	return node.Array()
}

// ObjectUseNode scans both parsed and non-parsed chidren nodes, 
// and map them by their keys
func (self Value) MapUseNode() (ret []KeyVal, err error) {
	if self.t != V_OBJECT {
		return nil, ErrUnsupportType
	}
	ret = make([]KeyVal, 0, _DEFAULT_NODE_CAP)
	err = self.ForEachKV(func(key string, node Value) bool {
		ret = append(ret, KeyVal{key, node})
		return true
	})
    return ret, err
}

// Map loads all keys of an object node
func (self Value) Map() (ret map[string]interface{}, err error) {
	if self.t != V_OBJECT {
		return nil, ErrUnsupportType
	}
	node := NewRaw(self.js)
	return node.Map()
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic type.
// WARN: all numberic nodes are casted to float64
func (self Value) Interface() (interface{}, error) {
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
func (self Value) ForEachKV(sc func(key string, node Value) bool) error {
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
func (self Value) ForEachElem(sc func(i int, node Value) bool) error {
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

// KeyVal is a pair of key and value (RawNode)
type KeyVal struct {
    Key   string
    Value Value
}

// GetRawByPath
func (self Searcher) GetRawByPath(path ...interface{}) (Value, error) {
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
    return Value{int(t), self.parser.s[s:self.parser.p]}, nil
}