package ast

import (
	"encoding/json"
	"strconv"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

const (
	V_NONE   = 0
	V_ERROR  = 1
	V_NULL   = int(types.T_NULL)
	V_TRUE   = int(types.T_TRUE)
	V_FALSE  = int(types.T_FALSE)
	V_ARRAY  = int(types.T_ARRAY)
	V_OBJECT = int(types.T_OBJECT)
	V_STRING = int(types.T_STRING)
	V_NUMBER = int(types.T_NUMBER)
	// V_ANY    = V_NUMBER + 1 // go interface{}, only appears when using `SetXX()` API
)

// JSON Node
type Node struct {
	node types.Node
	// mut []interface{}
}

// Type returns json type represented by the node
// It will be one of belows:
//
//	V_NONE   = 0 (empty node)
//	V_ERROR  = 1 (something wrong)
//	V_NULL   = 2 (json value `null`)
//	V_TRUE   = 3 (json value `true`)
//	V_FALSE  = 4 (json value `false`)
//	V_ARRAY  = 5 (json value array)
//	V_OBJECT = 6 (json value object)
//	V_STRING = 7 (json value string)
//	V_NUMBER = 8 (json value number )
//	V_ANY 	 = 9 (go interface{})
func (self Node) Type() int {
	return int(self.node.Kind)
}

/***************** Check APIs ***********************/

// Exists returns false only if the self is nil or empty node V_NONE
func (self *Node) Exists() bool {
	return self.Valid() && self.node.Kind != V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
	if self == nil {
		return false
	}
	return self.node.Kind != V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotExist If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node) Check() error {
	if self == nil {
		return ErrNotExist
	} else if self.node.Kind == V_ERROR || self.node.Kind == V_NONE {
		return self
	} else {
		return nil
	}
}

// func (self *Node) isMut() bool{
// 	return self != nil && len(self.mut) != 0
// }

var (
	nullNode  = types.NewNode("null", 0, 0)
	trueNode  = types.NewNode("true", 0, 0)
	falseNode = types.NewNode("false", 0, 0)
)

// NewRaw creates a node of raw json.
// If the input json is invalid, NewRaw returns a error Node.
func NewRaw(json string) Node {
	n, e := NewParser(json).Parse()
	if e != nil {
		return newError(e)
	}
	return n
}

// NewAny creates a node of type:
//   - V_ANY: if any's type isn't Node or *Node
//     val.Type(): if any's type is Node or *Node
func NewAny(val interface{}, opts encoder.Options) Node {
	// if n, isNode := val.(Node); isNode {
	// 	return n
	// } else if nn, isNode := val.(*Node); isNode {
	// 	return *nn
	// }
	// return Node{
	// 	node: types.Node{
	// 		Kind: types.Type(V_ANY),
	// 	},
	// }
	js, err := encoder.Encode(val, opts)
	if err != nil {
		return newError(err)
	}
	return NewRaw(rt.Mem2Str(js))
}

// func (n *Node) any() interface{} {
// 	return n.mut[0]
// }

// NewNull creates a node of type V_NULL
func NewNull() Node {
	return Node{nullNode}
}

// NewBool creates a node of type bool:
//
//	If v is true, returns V_TRUE node
//	If v is false, returns V_FALSE node
func NewBool(v bool) Node {
	if v {
		return Node{trueNode}
	} else {
		return Node{falseNode}
	}
}

// NewNumber creates a json.Number node
// v must be a decimal string complying with RFC8259
func NewNumber(v json.Number) Node {
	return newRawNodeLoad(string(v), 0)
}

// NewString creates a node of type V_STRING.
// v is considered to be a valid UTF-8 string, which means it won't be validated and unescaped.
// when the node is encoded to json, v will be escaped.
func NewString(v string) Node {
	s := encoder.Quote(v)
	esc := len(s) > len(v)+2
	if esc {
		return newRawNodeLoad(s, types.F_ESC)
	}
	return newRawNodeLoad(s, 0)
}

func NewInt64(v int64) Node {
	s, err := encoder.Encode(v, 0)
	if err != nil {
		return newError(err)
	}
	return newRawNodeLoad(rt.Mem2Str(s), 0)
}

func NewFloat(v float64) Node {
	s, err := encoder.Encode(v, 0)
	if err != nil {
		return newError(err)
	}
	return newRawNodeLoad(rt.Mem2Str(s), 0)
}

func (n *Node) Raw() (string, error) {
	if n.Error() != "" {
		return "", n
	}
	// if n.isMut() {
	// 	println("marshal json")
	// 	js, err := n.MarshalJSON()
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return rt.Mem2Str(js), nil
	// }
	return n.node.JSON, nil
}

func (n *Node) Len() (int, error) {
	switch n.node.Kind {
	case types.T_ARRAY, types.T_OBJECT:
		return len(n.node.Kids), nil
	default:
		return 0, ErrUnsupportType
	}
}

func (n *Node) Bool() (bool, error) {
	switch n.node.Kind {
	case types.T_FALSE:
		return false, nil
	case types.T_TRUE:
		return true, nil
	// case types.Type(V_ANY):
	// 	b, ok := n.any().(bool)
	// 	if !ok {
	// 		return false, ErrUnsupportType
	// 	}
	// 	return b, nil
	case V_ERROR:
		return false, n
	case V_NONE:
		return false, ErrNotExist
	default:
		return false, ErrUnsupportType
	}
}

func (n *Node) Int64() (int64, error) {
	if err := n.Check(); err != nil {
		return 0, err
	}
	switch n.node.Kind {
	case types.T_NUMBER:
		iv, err := strconv.ParseInt(n.node.JSON, 10, 64)
		if err != nil {
			return 0, makeSyntaxError(n.node.JSON, 0, err.Error())
		}
		return iv, nil
	// case types.Type(V_ANY):
	// 	v, ok := castToInt64(n.any())
	// 	if !ok {
	// 		return 0, ErrUnsupportType
	// 	}
	// 	return v, nil
	default:
		return 0, ErrUnsupportType
	}
}

func (n *Node) Float64() (float64, error) {
	if err := n.Check(); err != nil {
		return 0, err
	}
	switch n.node.Kind {
	case types.T_NUMBER:
		fv, err := strconv.ParseFloat(n.node.JSON, 64)
		if err != nil {
			return 0, makeSyntaxError(n.node.JSON, 0, err.Error())
		}
		return fv, nil
	// case types.Type(V_ANY):
	// 	v, ok := castToFloat64(n.any())
	// 	if !ok {
	// 		return 0, ErrUnsupportType
	// 	}
	// 	return v, nil
	default:
		return 0, ErrUnsupportType
	}
}

func (n *Node) Number() (json.Number, error) {
	if err := n.Check(); err != nil {
		return "", err
	}
	switch n.node.Kind {
	case types.T_NUMBER:
		return json.Number(n.node.JSON), nil
	// case types.Type(V_ANY):
	// 	v, ok := castToNumber(n.any())
	// 	if !ok {
	// 		return "", ErrUnsupportType
	// 	}
	// 	return v, nil
	default:
		return "", ErrUnsupportType
	}
}

func (n *Node) String() (string, error) {
	if err := n.Check(); err != nil {
		return "", err
	}
	switch n.node.Kind {
	case types.T_STRING:
		return raw2str(n.node.JSON, n.node.Flag.IsEsc(), 0)
	// case types.Type(V_ANY):
	// 	v, ok := castToString(n.any())
	// 	if !ok {
	// 		return "", ErrUnsupportType
	// 	}
	// 	return v, nil
	default:
		return "", ErrUnsupportType
	}
}

func (n *Node) Array(buf *[]Node) error {
	if err := n.Check(); err != nil {
		return err
	}
	switch n.node.Kind {
	case types.T_ARRAY:
		l := len(n.node.Kids)
		ol := len(*buf)
		if cap(*buf)-ol < l {
			tmp := make([]Node, ol+l)
			copy(tmp, *buf)
			*buf = tmp[:ol]
		}
		for _, t := range n.node.Kids {
			*buf = append(*buf, n.getKidLoad(t))
		}
		return nil
	// case types.Type(V_ANY):
	// 	ok := castToArray(n.any(), buf)
	// 	if !ok {
	// 		return ErrUnsupportType
	// 	}
	// 	return nil
	default:
		return ErrUnsupportType
	}
}

func (n *Node) Map(buf map[string]Node) error {
	if err := n.should(types.T_OBJECT); err != nil {
		return err
	}
	if err := n.Check(); err != nil {
		return err
	}
	switch n.node.Kind {
	case types.T_OBJECT:
		for i := 0; i < len(n.node.Kids); {
			t := n.node.Kids[i]
			key, err := n.str(t)
			if err != nil {
				return err
			}
			tt := n.node.Kids[i+1]
			val := n.getKidLoad(tt)
			buf[key] = val
			i += 2
		}
		return nil
	// case types.Type(V_ANY):
	// 	ok := castToMap(n.any(), buf)
	// 	if !ok {
	// 		return ErrUnsupportType
	// 	}
	// 	return nil
	default:
		return ErrUnsupportType
	}

}

func (n *Node) Interface(opts decoder.Options) (interface{}, error) {
	switch n.node.Kind {
	// case types.Type(V_ANY):
	// 	return n.any(), nil
	case types.T_NULL:
		return nil, nil
	case types.T_FALSE:
		return false, nil
	case types.T_TRUE:
		return true, nil
	case types.T_STRING:
		return n.String()
	case types.T_NUMBER:
		if opts&decoder.OptionUseNumber != 0 {
			return json.Number(n.node.JSON), nil
		} else if opts&decoder.OptionUseInt64 != 0 {
			if iv, err := n.Int64(); err == nil {
				return iv, nil
			}
		}
		return n.Float64()
	case types.T_ARRAY:
		buf := make([]interface{}, len(n.node.Kids))
		for i, tok := range n.node.Kids {
			js := tok.Raw(n.node.JSON)
			dc := decoder.NewDecoder(js)
			dc.SetOptions(opts)
			if err := dc.Decode(&buf[i]); err != nil {
				return nil, err
			}
		}
		return buf, nil
	case types.T_OBJECT:
		buf := make(map[string]interface{}, len(n.node.Kids)/2)
		for i := 0; i < len(n.node.Kids); i += 2 {
			key, err := n.str(n.node.Kids[i])
			if err != nil {
				return nil, err
			}
			var val interface{}
			js := n.node.Kids[i+1].Raw(n.node.JSON)
			dc := decoder.NewDecoder(js)
			dc.SetOptions(opts)
			if err := dc.Decode(&val); err != nil {
				return nil, err
			}
			buf[key] = val
		}
		return buf, nil
	case V_ERROR:
		return nil, n
	case V_NONE:
		return nil, ErrNotExist
	default:
		return nil, ErrUnsupportType
	}
}

func (n *Node) ForEachKV(scanner func(key string, elem Node) bool) error {
	if err := n.should(types.T_OBJECT); err != nil {
		return err
	}
	for i := 0; i < len(n.node.Kids); i += 2 {
		key, err := n.str(n.node.Kids[i])
		if err != nil {
			return err
		}
		val := n.getKidLoad(n.node.Kids[i+1])
		if !scanner(key, val) {
			return nil
		}
	}
	return nil
}

func (n *Node) ForEachElem(scanner func(index int, elem Node) bool) error {
	if err := n.should(types.T_ARRAY); err != nil {
		return err
	}
	for i, t := range n.node.Kids {
		elem := n.getKidLoad(t)
		if !scanner(i, elem) {
			return nil
		}
	}
	return nil
}

func (n *Node) Get(key string) Node {
	if err := n.should(types.T_OBJECT); err != nil {
		return newError(err)
	}
	_, t, err := n.objAt(key)
	if err != nil {
		return newError(err)
	}
	return n.getKidLoad(*t)
}

func (n *Node) get(key string) (string, error) {
	if err := n.should(types.T_OBJECT); err != nil {
		return "", err
	}
	_, t, err := n.objAt(key)
	if err != nil {
		return "", err
	}
	return n.json(*t), nil
}

func (n *Node) Index(key int) Node {
	if err := n.should(types.T_ARRAY); err != nil {
		return newError(err)
	}
	t := n.arrAt(key)
	if t == nil {
		return Node{}
	}
	return n.getKidLoad(*t)
}

func (n *Node) index(key int) (string, error) {
	if err := n.should(types.T_ARRAY); err != nil {
		return "", err
	}
	t := n.arrAt(key)
	if t == nil {
		return "", ErrNotExist
	}
	return n.json(*t), nil
}

func (self *Node) GetByPath(paths ...interface{}) Node {
	if l := len(paths); l == 0 {
		return *self
	} else {

		// first path, using node.index
		var js = self.node.JSON
		var err error
		switch p := paths[0].(type) {
		case int:
			if l == 1 {
				return self.Index(p)
			}
			js, err = self.index(p)
			if err != nil {
				return newError(err)
			}
		case string:
			if l == 1 {
				return self.Get(p)
			}
			js, err = self.get(p)
			if err != nil {
				return newError(err)
			}
		default:
			return ErrInvalidPath
		}

		// more paths, using parser..
		n, err := NewParser(js).GetByPath(paths[1:]...)
		if err != nil {
			return newError(err)
		}
		return n
	}
}

var EstimatedInsertedPathCharSize = 8

func (self *Node) SetByPath(allowArrayAppend bool, json string, path ...interface{}) (bool, error) {
	// TODO, pass option
	// js, e := encoder.Encode(val, 0)
	// if e != nil {
	// 	return false, e
	// }
	valjs := json
	if l := len(path); l == 0 {
		*self = NewRaw(valjs)
		return true, nil
		// } else if l == 1 {
		// 	// for one layer set
		// 	switch p := path[0].(type) {
		// 	case int:
		// 		e := self.arrSet(p, val)
		// 		if e == ErrNotExist && allowArrayAppend {
		// 			ee := self.arrAdd(val)
		// 			return false, ee
		// 		}
		// 		return e == nil, e
		// 	case string:
		// 		e := self.objSet(p, val)
		// 		if e == ErrNotExist {
		// 			ee := self.objAdd(p, val)
		// 			return false, ee
		// 		}
		// 		return e == nil, e
		// 	default:
		// 		panic("path must be either int or string")
		// 	}
	} else {
		// multi layers set
		p := NewParser(self.node.JSON)
		var err types.ParsingError
		var missing int
		var start int
		var exist bool
		for i, k := range path {
			if id, ok := k.(int); ok {
				if start, err = p.locate(0, id); err != 0 {
					if err != types.ERR_NOT_FOUND {
						return exist, makeSyntaxError(self.node.JSON, p.pos, err.Message())
					} else if !allowArrayAppend {
						return false, ErrIndexOutOfRange{Index: i, Err: ErrNotExist}
					} else {
						missing = i
						break
					}
				}
			} else if key, ok := k.(string); ok {
				if start, err = p.locate(0, key); err != 0 {
					// for object, we allow insert non-exist key
					if err != types.ERR_NOT_FOUND {
						return exist, makeSyntaxError(self.node.JSON, p.pos, err.Message())
					} else {
						missing = i
						break
					}
				}
			} else {
				return false, ErrInvalidPath
			}
			// next layer, reset pos
			if i != l-1 {
				p.pos = start
			}
		}
		var b []byte
		if err == types.ERR_NOT_FOUND {
			// NOTICE: pos must stop at '}' or ']'
			s, c := backward(self.node.JSON, p.pos)
			path := path[missing:]
			size := len(self.node.JSON) + len(valjs) + EstimatedInsertedPathCharSize*len(path)
			b = make([]byte, 0, size)
			b = append(b, self.node.JSON[:s]...)
			if c != '[' && c != '{' {
				// non-empty, insert sep
				b = append(b, ","...)
			}
			// creat new nodes on path
			var err error
			b, err = makePathAndValue(b, path, allowArrayAppend, valjs)
			if err != nil {
				return exist, err
			}
			b = append(b, self.node.JSON[s:]...)
		} else {
			b = make([]byte, 0, start+len(valjs)+(len(self.node.JSON)-p.pos))
			b = append(b, self.node.JSON[:start]...)
			b = append(b, valjs...)
			b = append(b, self.node.JSON[p.pos:]...)
			exist = true
		}
		// refrest the node
		p.src = rt.Mem2Str(b)
		p.pos = 0
		node, ee := p.Parse()
		if ee != nil {
			return true, ee
		}
		*self = node
		return exist, nil
	}
}

func (self *Node) UnsetByPath(path ...interface{}) (bool, error) {
	if l := len(path); l == 0 {
		*self = Node{}
		return true, nil
		// } else if l == 1 {
		// 	// for one layer set
		// 	switch p := path[0].(type) {
		// 	case int:
		// 		e := self.arrDel(p)
		// 		return e != ErrNotExist, e
		// 	case string:
		// 		e := self.objDel(p)
		// 		return e != ErrNotExist, e
		// 	default:
		// 		panic("path must be either int or string")
		// 	}
	} else {
		// multi layers set
		p := NewParser(self.node.JSON)
		var err types.ParsingError
		var start int
		if start, err = p.locate(types.F_GetByPath_StartAtKey, path...); err != 0 {
			if err == types.ERR_NOT_FOUND {
				return false, nil
			} else if err == types.ERR_UNSUPPORT_TYPE {
				return false, ErrUnsupportType
			} else {
				return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
			}
		}

		end := p.pos
		start, c := backward(self.node.JSON, start)
		if c == ',' {
			// not first, remove sep
			start -= 1
		} else {
			// first, remove subfix ',' (if any)
			end, c = forward(self.node.JSON, end)
			if c == ',' {
				end += 1
			}
		}

		var b []byte
		b = make([]byte, 0, len(self.node.JSON)-(end-start))
		b = append(b, self.node.JSON[:start]...)
		b = append(b, self.node.JSON[end:]...)
		// refrest the node
		p.src = rt.Mem2Str(b)
		p.pos = 0
		node, ee := p.Parse()
		if ee != nil {
			return true, ee
		}
		*self = node
		return true, nil
	}
}
