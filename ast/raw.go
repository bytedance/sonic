package ast

import (
	"encoding/json"
	"sort"
	"strconv"
	"sync"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// Value represents a raw json value or error
// It's safe to concurrently read Value
//
// Notice: passing a Value equals to pass-by-value,
// which measn any moditification on passed Value won't affect the origin one
type Value struct {
	t  int
	js string
}

// NewValueJSON converts a json string to Value
func NewValueJSON(json string) Value {
	p := NewParser(json)
	s, e := p.skip()
	if e != 0 {
		return errValue(p.ExportError(e))
	}
	return value(json[s:p.p])
}

// NewValueJSON converts a json bytes to Value
func NewValueJSONBytes(json []byte) Value {
	js := rt.Mem2Str(json)
	p := NewParser(js)
	s, e := p.skip()
	if e != 0 {
		return errValue(p.ExportError(e))
	}
	return value(js[s:p.p])
}

// NewValue converts a go primitive object to Value
func NewValue(val interface{}) Value {
	if v, ok := val.(Value); ok {
		return v
	}
	if v, ok := val.(*Value); ok {
		return *v
	}
	js, err := encoder.Encode(val, 0)
	if err != nil {
		return errValue(err)
	}
	return value(rt.Mem2Str(js))
}

func value(js string) Value {
	return Value{
		t:  int(switchRawType(js[0])),
		js: js,
	}
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
//	V_NUMBER = 33 (json value number )
func (self Value) Type() int {
	return self.t
}

// Exists tells if the nodes exists (including V_NULL)
func (self Value) Exists() bool {
	return self.t != 0 && self.t != V_ERROR
}

// MarshalJSON for encoding/json.Marshaler
func (self Value) MarshalJSON() ([]byte, error) {
	js, err := self.Raw()
	if err != nil {
		return nil, err
	}
	return rt.Str2Mem(js), nil
}

// UnmarshalJSON for encoding/json.Unmarshaler
func (self *Value) UnmarshalJSON(in []byte) error {
	n := NewValueJSON(rt.Mem2Str(in))
	if err := n.Check(); err != nil {
		return err
	}
	*self = n
	return nil
}

func (self Value) itype() types.ValueType {
	return types.ValueType(self.t)
}

// Error returns error message if the node is invalid
func (self Value) Error() string {
	if self.t == V_ERROR {
		return self.js
	} else if self.t == 0 {
		return ErrNotExist.Error()
	}
	return ""
}

// Check checks if the node itself is valid, and return
func (self Value) Check() error {
	if self.t == V_ERROR {
		return self
	} else if self.t == 0 {
		return ErrNotExist
	}
	return nil
}

func errValue(err error) Value {
	return Value{t: V_ERROR, js: err.Error()}
}

// Len returns children count of a array|object|string node
//
// WARN: this calculation consumes much CPU time
func (self Value) Len() (int, error) {
	switch self.t {
	case V_STRING:
		str, e := self.toString()
		if e != nil {
			return -1, e
		}
		return len(str), nil
	case V_ARRAY:
		c, _ := self.count_elems()
		return c, nil
	case V_OBJECT:
		c, _ := self.count_kvs()
		return c, nil
	default:
		return -1, ErrUnsupportType
	}
}

func (self Value) count_elems() (int, types.ParsingError) {
	p := NewParserObj(self.js)
	if empty, e := p.arrayBegin(); e != 0 {
		return -1, e
	} else if empty {
		return 0, 0
	}

	i := 0
	for {
		if _, e := p.skipFast(); e != 0 {
			return -1, e
		}
		i++
		if end, e := p.arrayEnd(); e != 0 {
			return -1, e
		} else if end {
			return i, 0
		}
	}
}

func (self Value) count_kvs() (int, types.ParsingError) {
	p := NewParserObj(self.js)
	if empty, e := p.objectBegin(); e != 0 {
		return -1, e
	} else if empty {
		return 0, 0
	}

	i := 0
	for {
		if _, e := p.key(); e != 0 {
			return -1, e
		}
		if _, e := p.skipFast(); e != 0 {
			return -1, e
		}
		i++
		if end, e := p.objectEnd(); e != 0 {
			return -1, e
		} else if end {
			return i, 0
		}
	}
}

// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self Value) GetByPath(path ...interface{}) Value {
	if self.Check() != nil {
		return self
	}
	p := NewParserObj(self.js)
	s, e := p.getByPathNoValidate(path...)
	if e != 0 {
		return errValue(p.ExportError(e))
	}
	return value(self.js[s:p.p])
}

// SetAnyByPath set value on given path and create nodes on the json if not exist
func (self *Value) SetAnyByPath(val interface{}, allowInsert bool, path ...interface{}) (bool, error) {
	return self.SetByPath(NewValue(val), allowInsert, path...)
}

// SetByPath set value on given path and create nodes on the json if not exist
// allowAppend controls if append val to V_ARRAY node if the path index is out of its range
func (self *Value) SetByPath(val Value, allowAppend bool, path ...interface{}) (bool, error) {
	exist := false
	if self.Check() != nil {
		return exist, self
	}
	if val.Check() != nil {
		return exist, val
	}

	p := NewParserObj(self.js)
	var err types.ParsingError
	var idx int

	for i, k := range path {
		if id, ok := k.(int); ok && id >= 0 {
			if _, err = p.searchIndex(id); err != 0 {
				if !allowAppend {
					return false, p.ExportError(err)
				}
				if err != _ERR_NOT_FOUND {
					return exist, p.ExportError(err)
				}
				idx = i
				break
			}
		} else if key, ok := k.(string); ok {
			if _, err = p.searchKey(key); err != 0 {
				if err != _ERR_NOT_FOUND {
					return exist, p.ExportError(err)
				}
				idx = i
				break
			}
		} else {
			return false, ErrInvalidPath
		}
	}

	var b []byte
	if err == 0 {
		exist = true
		// found, just skip and replace
		s, err := p.skipFast()
		if err != 0 {
			return exist, p.ExportError(err)
		}
		size := len(self.js) + len(val.js) - (p.p - s)
		b = make([]byte, 0, size)
		b = append(b, self.js[:s]...)
		b = append(b, val.js...)
	} else {
		// not found, stop at end of idx's parent
		s := p.p - 1
		for ; s >= 0 && isSpace(self.js[s]); s-- {
		}
		empty := (self.js[s] == '[' || self.js[s] == '{')
		size := len(self.js) + len(val.js) + 8*(len(path)-idx)
		b = make([]byte, 0, size)
		s = s + 1
		b = append(b, self.js[:s]...)
		if !empty {
			b = append(b, ","...)
		}
		// creat new nodes on path
		var err error
		b, err = appendPathValue(b, path[idx:], val)
		if err != nil {
			return exist, err
		}
	}

	b = append(b, self.js[p.p:]...)
	self.js = rt.Mem2Str(b)
	return exist, nil
}

// [2,"a"],1 => {"a":1}
// ["a",2],1  => "a":[1]
func appendPathValue(b []byte, path []interface{}, val Value) ([]byte, error) {
	for i, k := range path {
		if key, ok := k.(string); ok {
			quote(&b, key)
			b = append(b, ":"...)
		}
		if i == len(path)-1 {
			b = append(b, val.js...)
			break
		}
		n := path[i+1]
		if _, ok := n.(int); ok {
			b = append(b, "["...)
		} else if _, ok := n.(string); ok {
			b = append(b, `{`...)
		} else {
			return nil, ErrInvalidPath
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

// UnsetByPath delete value on given path.
func (self *Value) UnsetByPath(path ...interface{}) (bool, error) {
	if self.Check() != nil {
		return false, self
	}

	p := NewParserObj(self.js)
	var err types.ParsingError
	var comma = -1

	for _, k := range path {
		if id, ok := k.(int); ok && id >= 0 {
			if comma, err = p.searchIndex(id); err != 0 {
				if err == _ERR_NOT_FOUND {
					return false, nil
				}
				return false, p.ExportError(err)
			}
		} else if key, ok := k.(string); ok {
			if comma, err = p.searchKey(key); err != 0 {
				if err == _ERR_NOT_FOUND {
					return false, nil
				}
				return false, p.ExportError(err)
			}
		} else {
			return false, ErrInvalidPath
		}
	}

	var b []byte
	s, err := p.skipFast()
	if err != 0 {
		return true, p.ExportError(err)
	}
	if comma != -1 {
		s = comma
	}

	e := p.p
	if self.js[s] != ',' { // first elem
		// check if trailling ','
		p.p = p.lspace(p.p)
		if p.p < len(self.js) && self.js[p.p] == ',' {
			e = p.p + 1
		}
	}

	size := len(self.js) - (e - s)
	b = make([]byte, 0, size)
	b = append(b, self.js[:s]...)
	b = append(b, self.js[e:]...)
	self.js = rt.Mem2Str(b)
	return true, nil
}

// Get loads given key of an object node on demands
func (self Value) Get(key string) Value {
	if self.Check() != nil {
		return self
	}
	p := NewParserObj(self.js)
	s, e := p.getByPathNoValidate(key)
	if e != 0 {
		return errValue(p.ExportError(e))
	}
	return value(self.js[s:p.p])
}

// GetMany retrieves all the keys in kvs and set found Value at correpsonding index
//
// WARN: kvs shouldn't contains any repeated key, otherwise only first-occured key will be given value
func (self Value) GetMany(keys []string, vals []Value) error {
	if self.Check() != nil {
		return self
	}
	if self.t != V_OBJECT {
		return ErrUnsupportType
	}
	if e := self.getMany(keys, false, func(i, s, e int) {
		vals[i] = value(self.js[s:e])
	}); e != 0 {
		return NewParserObj(self.js).ExportError(e)
	}
	return nil
}

func (self Value) getMany(kvs []string, delete bool, hook func(i, s, e int)) types.ParsingError {
	p := NewParserObj(self.js)
	if empty, e := p.objectBegin(); e != 0 {
		return e
	} else if empty {
		return 0
	}

	count := len(kvs)
	for count > 0 {
		ks := p.p
		key, err := p.key()
		if err != 0 {
			return err
		}
		s, err := p.skipFast()
		if err != 0 {
			return err
		}

		var found = -1
		for i, kv := range kvs {
			if kv == key {
				if !delete {
					hook(i, s, p.p)
				} else {
					found = i
				}
				count--
				break
			}
		}

		if end, e := p.objectEnd(); e != 0 {
			return e
		} else if end {
			if found != -1 {
				hook(found, ks, p.p-1)
			}
			break
		}

		if found != -1 {
			hook(found, ks, p.p)
		}
	}

	return 0
}

// Index indexies node at given idx
func (self Value) Index(idx int) Value {
	if self.Check() != nil {
		return self
	}
	p := NewParserObj(self.js)
	s, e := p.getByPathNoValidate(idx)
	if e != 0 {
		return errValue(p.ExportError(e))
	}
	return value(self.js[s:p.p])
}

// GetMany retrieves all the indexes in ids and set found Value at correpsonding index of vals
//
// WARN: ids shouldn't contains any repeated index, otherwise only first-occured id will be given value
func (self Value) IndexMany(ids []int, vals []Value) error {
	if self.Check() != nil {
		return self
	}
	if self.t != V_ARRAY {
		return ErrUnsupportType
	}
	if e := self.indexMany(ids, false, func(i, s, e int) {
		vals[i] = value(self.js[s:e])
	}); e != 0 {
		return NewParserObj(self.js).ExportError(e)
	}
	return nil
}

func (self Value) indexMany(ids []int, delete bool, hook func(i, s, e int)) types.ParsingError {
	p := NewParserObj(self.js)
	if empty, e := p.arrayBegin(); e != 0 {
		return e
	} else if empty {
		return 0
	}

	count := len(ids)
	i := 0
	for count > 0 {
		s, e := p.skipFast()
		if e != 0 {
			return e
		}

		var found = -1
		for j, id := range ids {
			if id == i {
				if !delete {
					hook(j, s, p.p)
				} else {
					found = j
				}
				count--
				break
			}
		}

		if end, e := p.arrayEnd(); e != 0 {
			return e
		} else if end {
			if found != -1 {
				hook(found, s, p.p-1)
			}
			break
		}
		if found != -1 {
			hook(found, s, p.p)
		}
		i++
	}

	return 0
}

type point struct {
	i int
	s int
	e int
}

type points []point

func (ps points) Less(i, j int) bool {
	return ps[i].s < ps[j].s
}

func (ps points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps points) Len() int {
	return len(ps)
}

// Set sets the node of given key under self, and insert new value if not exist.
// It reports if the key has existed.
func (self *Value) SetAny(key string, val interface{}, allowInsert bool) (bool, error) {
	return self.SetMany([]string{key}, []Value{NewValue(val)})
}

// Set sets the node of given key under self, and insert new value if not exist.
// It reports if the key has existed.
func (self *Value) Set(key string, val Value, allowInsert bool) (bool, error) {
	return self.SetMany([]string{key}, []Value{val})
}

// SetMany retries kvs in the V_OBJECT value,
// and replace (exist) or insert (not-exist) key with correpsonding value
//
// WARN: kvs shouldn't contains any repeated key, otherwise the repeated key will be regarded as new key and insert
func (self *Value) SetMany(keys []string, vals []Value) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if len(keys) == 0 {
		return false, nil
	}
	if self.t == V_NULL {
		*self = value(`[]`)
	}
	if self.t != V_OBJECT {
		return false, ErrUnsupportType
	}

	var points = make(points, len(keys))
	var size = len(self.js)
	var exist = false

	// collect replace points
	if e := self.getMany(keys, false, func(i, s, e int) {
		size += len(vals[i].js) - (e - s)
		points[i] = point{i, s, e}
		exist = true
	}); e != 0 {
		return exist, NewParserObj(self.js).ExportError(e)
	}

	// collect insert ponts
	for i, r := range points {
		if r.s == 0 {
			points[i].i = i
			size += len(vals[i].js) + 4 + len(keys[i])
		}
	}

	b := make([]byte, 0, size)

	// write replace points
	sort.Stable(points)
	s := 0
	for i, r := range points {
		if r.s == 0 {
			continue
		}
		// write left
		b = append(b, self.js[s:r.s]...)
		// write new val
		b = append(b, vals[r.i].js...)
		if i < len(points)-1 {
			s = points[i+1].s
		} else {
			// not write '}'
			s = len(self.js) - 1
		}
		// write right
		b = append(b, self.js[r.e:s]...)
	}

	s = len(self.js) - 2
	if !exist {
		b = append(b, self.js[:s+1]...)
	}
	for ; s >= 0 && isSpace(self.js[s]); s-- {
	}
	empty := self.js[s] == '{'

	// write insert points
	for _, r := range points {
		if r.s != 0 {
			continue
		}
		// write last ','
		if !empty {
			b = append(b, ","...)
		}
		empty = false
		// write key
		quote(&b, keys[r.i])
		b = append(b, ":"...)
		// write new val
		b = append(b, vals[r.i].js...)
	}

	b = append(b, "}"...)
	self.js = rt.Mem2Str(b)
	return exist, nil
}

// SetByIndex replaces the node of given index.
func (self *Value) SetAnyByIndex(id int, val interface{}) (bool, error) {
	return self.SetByIndex(id, NewValue(val))
}

// SetByIndex replaces the node of given index.
func (self *Value) SetByIndex(id int, val Value) (bool, error) {
	return self.SetManyByIndex([]int{id}, []Value{val})
}

// SetManyByIndex retries ids in the V_ARRAY value,
// and replace existed ids with correpsonding value of vals
func (self *Value) SetManyByIndex(ids []int, vals []Value) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if len(ids) == 0 {
		return false, nil
	}
	if self.t == V_NULL {
		*self = value(`[]`)
	}
	if self.t != V_ARRAY {
		return false, ErrUnsupportType
	}

	var points = make(points, len(ids))
	var size = len(self.js)
	var exist = false

	// collect replace points
	if e := self.indexMany(ids, false, func(i, s, e int) {
		size += len(vals[i].js) - (e - s)
		points[i] = point{i, s, e}
		exist = true
	}); e != 0 {
		return exist, NewParserObj(self.js).ExportError(e)
	}
	if !exist {
		return false, ErrNotExist
	}

	// collect insert ponts
	for i, r := range points {
		if r.s == 0 {
			points[i].i = i
			size += len(vals[i].js) + 1
		}
	}

	b := make([]byte, 0, size)
	// write replace points
	sort.Stable(points)
	s := 0
	notExist := false
	for i, r := range points {
		if r.s == 0 {
			notExist = true
			continue
		}
		// write left
		b = append(b, self.js[s:r.s]...)
		// write new val
		b = append(b, vals[r.i].js...)
		if i < len(points)-1 {
			s = points[i+1].s
		} else {
			s = len(self.js) 
		}
		// write right
		b = append(b, self.js[r.e:s]...)
	}

	// if allowAppend {
	// 	s = len(self.js) - 2
	// 	if !exist {
	// 		b = append(b, self.js[:s+1]...)
	// 	}
	// 	for ; s >= 0 && isSpace(self.js[s]); s-- {
	// 	}
	// 	empty := self.js[s] == '['

	// 	// write insert points
	// 	for _, r := range points {
	// 		if r.s != 0 {
	// 			continue
	// 		}
	// 		// write last ','
	// 		if !empty {
	// 			b = append(b, ","...)
	// 		}
	// 		empty = false
	// 		// write new val
	// 		b = append(b, vals[r.i].js...)
	// 	}
	// }
	// b = append(b, "]"...)

	self.js = rt.Mem2Str(b)
	if notExist {
		return exist, ErrNotExist
	} else {
		return exist, nil
	}
}

// Add inserts single val into the left of the node at `from` (-1 means last)
func (self *Value) AddAny(from int, val interface{}) error {
	return self.AddMany(from, []Value{NewValue(val)})
}

// Add inserts single val into the left of the node at `from` (-1 means last)
func (self *Value) Add(from int, val Value) error {
	return self.AddMany(from, []Value{val})
}

// Add inserts multiple vals into the left of the node at `from` (-1 means last)
func (self *Value) AddMany(from int, vals []Value) error {
	if self.Check() != nil {
		return self
	}
	if len(vals) == 0 {
		return nil
	}
	if self.t == V_NULL {
		*self = value(`[]`)
	}
	if self.t != V_ARRAY {
		return ErrUnsupportType
	}

	e := 1 // from == 0
	if from > 0 {
		p := NewParserObj(self.js)
		_, err := p.searchIndex(from)
		if err != 0 {
			if err == _ERR_NOT_FOUND {
				e = len(self.js) - 1 //  end before ']'
			} else {
				return p.ExportError(err)
			}
		} else {
			e = p.p
		}
	} else if from < 0 {
		e = len(self.js) - 1
	}

	n := len(self.js)
	for _, v := range vals {
		n += len(v.js) + 1
	}
	b := make([]byte, 0, n)

	lastComma := self.js[e] != ']'
	s := e - 1
	for ;s>=0 && isSpace(self.js[s]); s++ {}
	firstComma := self.js[s] != ',' && self.js[s] != '['

	b = append(b, self.js[:e]...)
	for i, val := range vals {
		if i == 0 && firstComma {
			b = append(b, ","...)
		}
		b = append(b, val.js...)
		if i < len(vals)-1 || lastComma {
			b = append(b, ","...)
		}
	}
	b = append(b, self.js[e:]...)

	self.js = rt.Mem2Str(b)
	return nil
}

// Pop pops at most 1 trailling elements in the array.
func (self *Value) Pop() error {
	return self.PopMany(1)
}

var intsPool = sync.Pool{
	New: func() interface{} {
		return []int{}
	},
}

// PopMany pops at most n trailling elements in the array. -1 means clears all
func (self *Value) PopMany(n int) error {
	if self.Check() != nil {
		return self
	}
	if n == 0 {
		return nil
	}
	if self.t == V_NULL {
		return nil
	}
	if self.t != V_ARRAY {
		return ErrUnsupportType
	}
	if n < 0 {
		*self = value(`[]`)
		return nil
	}

	p := NewParserObj(self.js)
	if empty, e := p.arrayBegin(); e != 0 {
		return p.ExportError(e)
	} else if empty {
		return nil
	}

	ends := intsPool.Get().([]int)
	ends = ends[:0]
	for {
		_, e := p.skipFast()
		if e != 0 {
			return p.ExportError(e)
		}
		ends = append(ends, p.p)
		if end, e := p.arrayEnd(); e != 0 {
			return p.ExportError(e)
		} else if end {
			break
		}
	}

	i := len(ends) - n
	s := 1
	if i > 0 {
		s = ends[i-1]
	}
	intsPool.Put(ends)

	b := make([]byte, 0, s+1)
	b = append(b, self.js[:s]...)
	b = append(b, "]"...)
	self.js = rt.Mem2Str(b)
	return nil
}

// Unset REMOVE the node of given key under object parent, and reports if the key has existed.
func (self *Value) Unset(key string) (bool, error) {
	return self.UnsetMany([]string{key})
}

// UnsetMany REMOVE existing key and corresponding value of given keys.
func (self *Value) UnsetMany(keys []string) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if self.t == V_NULL {
		return false, nil
	}
	if self.t != V_OBJECT {
		return false, ErrUnsupportType
	}

	points := make(points, len(keys))
	size := len(self.js)
	replaced := false
	if err := self.getMany(keys, true, func(i, s, e int) {
		points[i] = point{i, s, e}
		size -= (e - s)
		replaced = true
	}); err != 0 {
		return replaced, NewParserObj(self.js).ExportError(err)
	}
	if !replaced {
		return replaced, nil
	}

	b := make([]byte, 0, size)
	// write replace points
	sort.Stable(points)
	s := 0
	for i, r := range points {
		if r.s == 0 {
			continue
		}

		// write left
		b = append(b, self.js[s:r.s]...)

		// write right
		if i < len(points)-1 {
			s = points[i+1].s
		} else {
			s = len(self.js)
		}
		b = append(b, self.js[r.e:s]...)
	}

	// check if any redundant ','
	for s = len(b) - 2; s >= 0 && isSpace(b[s]); s-- {
	}
	if b[s] == ',' {
		copy(b[s:], b[s+1:])
		b = b[:len(b)-1]
	}

	self.js = rt.Mem2Str(b)
	return replaced, nil
}

// UnsetByIndex REOMVE the node of given index.
func (self *Value) UnsetByIndex(id int) (bool, error) {
	return self.UnsetManyByIndex([]int{id})
}

// UnsetMany REMOVE existing id and corresponding value of given keys.
func (self *Value) UnsetManyByIndex(ids []int) (bool, error) {
	if self.Check() != nil {
		return false, self
	}
	if self.t == V_NULL {
		return false, nil
	}
	if self.t != V_ARRAY {
		return false, ErrUnsupportType
	}

	points := make(points, len(ids))
	size := len(self.js)
	replaced := false

	if err := self.indexMany(ids, true, func(i, s, e int) {
		points[i] = point{i, s, e}
		size -= (e - s)
		replaced = true
	}); err != 0 {
		return replaced, NewParserObj(self.js).ExportError(err)
	}

	if !replaced {
		return replaced, nil
	}

	b := make([]byte, 0, size)
	// write replace points
	sort.Stable(points)
	s := 0
	for i, r := range points {
		if r.s == 0 {
			continue
		}
		// write left
		b = append(b, self.js[s:r.s]...)
		if i < len(points)-1 {
			s = points[i+1].s
		} else {
			s = len(self.js)
		}
		// write right
		b = append(b, self.js[r.e:s]...)
	}

	// check if any redundant ','
	for s = len(b) - 2; s >= 0 && isSpace(b[s]); s-- {
	}
	if b[s] == ',' {
		copy(b[s:], b[s+1:])
		b = b[:len(b)-1]
	}

	self.js = rt.Mem2Str(b)
	return replaced, nil
}

func (self Value) str() string {
	return self.js[1 : len(self.js)-1]
}

func (self Value) raw() string {
	return self.js
}

// Raw returns json representation of the node
func (self Value) Raw() (string, error) {
	if e := self.Check(); e != nil {
		return "", e
	}
	return self.js, nil
}

func (self Value) toInt64() (int64, error) {
	return json.Number(self.js).Int64()
}

func (self Value) toFloat64() (float64, error) {
	return json.Number(self.js).Float64()
}

func (self Value) toString() (string, error) {
	ret, str := DecodeString(self.js, 0)
	if ret < 0 {
		return "", NewParserObj(self.js).ExportError(types.ParsingError(-ret))
	}
	return str, nil
}

// Bool returns bool value represented by this node,
// including types.V_TRUE|V_FALSE|V_NUMBER|V_STRING|V_ANY|V_NULL
func (self Value) Bool() (bool, error) {
	if e := self.Check(); e != nil {
		return false, e
	}
	switch self.t {
	case V_NULL:
		return false, nil
	case V_TRUE:
		return true, nil
	case V_FALSE:
		return false, nil
	case V_STRING:
		return strconv.ParseBool(self.str())
	case V_NUMBER:
		if i, err := self.toInt64(); err == nil {
			return i != 0, nil
		} else if f, err := self.toFloat64(); err == nil {
			return f != 0, nil
		} else {
			return false, err
		}
	default:
		return false, ErrUnsupportType
	}
}

// Int64 casts the node to int64 value,
// including V_NUMBER|V_TRUE|V_FALSE|V_STRING
func (self Value) Int64() (int64, error) {
	if e := self.Check(); e != nil {
		return 0, e
	}
	switch self.t {
	case V_NULL:
		return 0, nil
	case V_TRUE:
		return 1, nil
	case V_FALSE:
		return 0, nil
	case V_STRING:
		return json.Number(self.str()).Int64()
	case V_NUMBER:
		return self.toInt64()
	default:
		return 0, ErrUnsupportType
	}
}

// Float64 cast node to float64,
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL
func (self Value) Float64() (float64, error) {
	if e := self.Check(); e != nil {
		return 0, e
	}
	switch self.t {
	case V_NULL:
		return 0.0, nil
	case V_TRUE:
		return 1.0, nil
	case V_FALSE:
		return 0.0, nil
	case V_STRING:
		return json.Number(self.str()).Float64()
	case V_NUMBER:
		return self.toFloat64()
	default:
		return 0, ErrUnsupportType
	}
}

// Number casts node to float64,
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL,
func (self Value) Number() (json.Number, error) {
	if e := self.Check(); e != nil {
		return "", e
	}
	switch self.t {
	case V_NULL:
		return json.Number("0"), nil
	case V_TRUE:
		return json.Number("1"), nil
	case V_FALSE:
		return json.Number("0"), nil
	case V_STRING:
		num := json.Number(self.str())
		if _, err := num.Float64(); err == nil {
			return num, nil
		} else {
			return json.Number(""), err
		}
	case V_NUMBER:
		return json.Number(self.js), nil
	default:
		return "", ErrUnsupportType
	}
}

// String cast node to string,
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL
func (self Value) String() (string, error) {
	if e := self.Check(); e != nil {
		return "", e
	}
	switch self.t {
	case V_NULL:
		return "", nil
	case V_TRUE:
		return "true", nil
	case V_FALSE:
		return "false", nil
	case V_STRING:
		return self.toString()
	case V_NUMBER:
		return self.js, nil
	default:
		return "", ErrUnsupportType
	}
}

// Array appends children of the V_ARRAY to buf, in original order
func (self Value) Array(buf *[]Value) (err error) {
	if self.t != V_ARRAY {
		return ErrUnsupportType
	}
	if *buf == nil {
		*buf = make([]Value, 0, _DEFAULT_NODE_CAP)
	}
	return self.ForEachElem(func(i int, node Value) bool {
		*buf = append(*buf, node)
		return true
	})
}

// Map appends children of the V_OBJECT to buf
func (self Value) Map(buf *map[string]Value) (err error) {
	if self.t != V_OBJECT {
		return ErrUnsupportType
	}
	if *buf == nil {
		*buf = make(map[string]Value, _DEFAULT_NODE_CAP)
	}
	return self.ForEachKV(func(key string, node Value) bool {
		(*buf)[key] = node
		return true
	})
}

// Map appends children of the V_OBJECT to buf, in original array
func (self Value) MapAsSlice(keys *[]string, vals *[]Value) (err error) {
	if self.t != V_OBJECT {
		return ErrUnsupportType
	}
	if *keys == nil {
		*keys = make([]string, 0, _DEFAULT_NODE_CAP)
	}
	if *vals == nil {
		*vals = make([]Value, 0, _DEFAULT_NODE_CAP)
	}
	return self.ForEachKV(func(key string, node Value) bool {
		*keys = append(*keys, key)
		*vals = append(*vals, node)
		return true
	})
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic type (like encoding/json).
// It's behavior is same with default behavior of sonic/decoder
func (self Value) Interface() (interface{}, error) {
	if e := self.Check(); e != nil {
		return nil, e
	}
	var ret interface{}
	err := decoder.NewDecoder(self.js).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic type.
// It's behavior depends on decoder.Options
func (self Value) InterfaceUseOption(opts decoder.Options) (interface{}, error) {
	if e := self.Check(); e != nil {
		return nil, e
	}
	var ret interface{}
	dc := decoder.NewDecoder(self.js)
	dc.SetOptions(opts)
	err := dc.Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Interface loads all children under all pathes from this node,
// and converts itself as scalar type (like encoding/json), or nested value types:
//   - V_ARRAY: []Value
//   - V_OBJECT: map[string]Value
// It's behavior depends on decoder.Options
func (self Value) InterfaceUseNode() (interface{}, error) {
	if e := self.Check(); e != nil {
		return nil, e
	}
	switch self.t {
	case V_NULL:
		return nil, nil
	case V_TRUE:
		return true, nil
	case V_FALSE:
		return false, nil
	case V_NUMBER:
		return self.toFloat64()
	case V_STRING:
		return self.toString()
	case V_ARRAY:
		ret := make([]Value, 0, _DEFAULT_NODE_CAP)
		err := self.Array(&ret)
		return ret, err
	case V_OBJECT:
		ret := make(map[string]Value, _DEFAULT_NODE_CAP)
		err := self.Map(&ret)
		return ret, err
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
			return p.ExportError(err)
		} else if empty {
			return nil
		}

		for {
			k, e := p.key()
			if e != 0 {
				return p.ExportError(e)
			}
			s, e := p.skipFast()
			if e != 0 {
				return p.ExportError(e)
			}
			n := value(self.js[s:p.p])
			if !sc(k, n) {
				return nil
			}
			if end, e := p.objectEnd(); e != 0 {
				return p.ExportError(e)
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
			return p.ExportError(err)
		} else if empty {
			return nil
		}

		i := 0
		for {
			s, e := p.skipFast()
			if e != 0 {
				return p.ExportError(e)
			}
			n := value(self.js[s:p.p])
			if !sc(i, n) {
				return nil
			}
			i++
			if end, e := p.arrayEnd(); e != 0 {
				return p.ExportError(e)
			} else if end {
				return nil
			}
		}
	default:
		return ErrUnsupportType
	}
}
