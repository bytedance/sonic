/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	`encoding/json`
	`unsafe`

	`github.com/bytedance/sonic/decoder`
	`github.com/bytedance/sonic/internal/native/types`
	`github.com/bytedance/sonic/internal/rt`
	`github.com/bytedance/sonic/unquote`
)

const (
    _CAP_BITS          = 32
    _LEN_MASK          = 1 << _CAP_BITS - 1

    _NODE_SIZE = unsafe.Sizeof(Node{})
    _PAIR_SIZE = unsafe.Sizeof(Pair{})
)

const (
    _V_NONE         types.ValueType = 0
    _V_NODE_BASE    types.ValueType = 1<<5
    _V_LAZY         types.ValueType = 1 << 7
    _V_RAW          types.ValueType = 1 << 8
    _V_NUMBER                       = _V_NODE_BASE + 1
    _V_ARRAY_LAZY                   = _V_LAZY | types.V_ARRAY
    _V_OBJECT_LAZY                  = _V_LAZY | types.V_OBJECT
    _MASK_LAZY                      = _V_LAZY - 1
    _MASK_RAW                       = _V_RAW - 1
)

const (
    V_NONE   = 0
    V_ERROR  = 1
    V_NULL   = 2
    V_TRUE   = 3
    V_FALSE  = 4
    V_ARRAY  = 5
    V_OBJECT = 6
    V_STRING = 7
    V_NUMBER = int(_V_NUMBER)
)

type Node struct {
    v int64
    t types.ValueType
    p unsafe.Pointer
}

/** Node Type Accessor **/

// Type returns json type represented by the node
// It will be one of belows:
//    V_NONE   = 0
//    V_ERROR  = 1
//    V_NULL   = 2
//    V_TRUE   = 3
//    V_FALSE  = 4
//    V_ARRAY  = 5
//    V_OBJECT = 6
//    V_STRING = 7
//    V_NUMBER = 33
func (self Node) Type() int {
    return int(self.t & _MASK_LAZY & _MASK_RAW)
}

func (self Node) itype() types.ValueType {
    return self.t & _MASK_LAZY & _MASK_RAW
}

// Exists returns false only if the node is nil or got by invalid path
func (self *Node) Exists() bool {
    return self != nil && self.t != _V_NONE
}

// Valid returns true when the node has any type of V_NULL ~ V_STRING, or V_NUMBER
func (self *Node) Valid() bool {
    if self == nil {
        return false
    }
    it := self.Type()
    return it >= V_NULL && it <= V_STRING || it == V_NUMBER
}

// Check check if the node itself is valid, and return:
//   - ErrNotFound If the node does not exist
//   - Its underlying error If the node is V_ERROR
func (self *Node)  Check() error {
    if self == nil || self.t == V_NONE {
        return ErrNotExist
    } else if self.t != V_ERROR {
        return nil
    } else {
        return self
    }
}

// Error returns error message if the node is invalid
func (self Node) Error() string {
    if self.t == V_NONE {
        return "unsupported type"
    } else if self.t != V_ERROR {
        return ""
    } else {
        return *(*string)(self.p)
    } 
}

// IsRaw returns true if node's underlying value is raw json
func (self Node) IsRaw() bool {
    return self.t&_V_RAW != 0
}

func (self Node) isLazy() bool {
    return self.t&_V_LAZY != 0
}

/** Simple Value Methods **/

// Raw returns underlying json string of an raw node,
// which usually created by Search() api
func (self *Node) Raw() (string, error) {
    if !self.IsRaw() {
        return "", ErrUnsupportType
    }
    return addr2str(self.p, self.v), nil
}

func (self *Node) checkRaw() error {
    if self.IsRaw() {
        *self = self.parseRaw()
    }
    if err := self.Check(); err != nil {
        return err
    }
    return nil
}

// Bool_E returns bool value represented by this node
//
// If node type is not types.V_TRUE or types.V_FALSE, or V_RAW (must be a bool json value),
// it will return error
func (self *Node) Bool() (bool, error) {
    if err := self.checkRaw(); err != nil {
        return false, err
    }
    switch self.t {
        case types.V_TRUE  : return true , nil
        case types.V_FALSE : return false, nil
        default            : return false, ErrUnsupportType
    }
}

// Int64 as above.
func (self *Node) Int64() (int64, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    switch self.t {
        case _V_NUMBER        : return numberToInt64(self)
        case types.V_TRUE     : return 1, nil
        case types.V_FALSE    : return 0, nil
        default               : return 0, ErrUnsupportType
    }
}

// Number as above.
func (self *Node) Number() (json.Number, error) {
    if err := self.checkRaw(); err != nil {
        return json.Number(""), err
    }
    switch self.t {
        case _V_NUMBER        : return toNumber(self)  , nil
        case types.V_TRUE     : return json.Number("1"), nil
        case types.V_FALSE    : return json.Number("0"), nil
        default               : return json.Number(""), ErrUnsupportType
    }
}

// String returns raw string value if node type is V_STRING.
// Or return the string representation of other types:
//  V_NULL => "null",
//  V_TRUE => "true",
//  V_FALSE => "false",
//  V_NUMBER => "[0-9\.]*"
func (self *Node) String() (string, error) {
    if err := self.checkRaw(); err != nil {
        return "", err
    }
    switch self.t {
        case _V_NUMBER       : return toNumber(self).String(), nil
        case types.V_NULL    : return "null" , nil
        case types.V_TRUE    : return "true" , nil
        case types.V_FALSE   : return "false", nil
        case types.V_STRING  : return addr2str(self.p, self.v), nil
        default              : return ""     , ErrUnsupportType
    }
}

// Float64 as above.
func (self *Node) Float64() (float64, error) {
    if err := self.checkRaw(); err != nil {
        return 0.0, err
    }
    switch self.t {
        case _V_NUMBER       : return numberToFloat64(self)
        case types.V_TRUE    : return 1.0, nil
        case types.V_FALSE   : return 0.0, nil
        default              : return 0.0, ErrUnsupportType
    }
}

/** Sequencial Value Methods **/

// Len returns children count of a array|object|string node
// For partially loaded node, it also works but only counts the parsed children
func (self *Node) Len() (int, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        return int(self.v & _LEN_MASK), nil
    } else if self.t == types.V_STRING {
        return int(self.v), nil
    } else {
        return 0, ErrUnsupportType
    }
}

func (self Node) len() int {
    return int(self.v & _LEN_MASK)
}

// Cap returns malloc capacity of a array|object node for children
func (self *Node) Cap() (int, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        return int(self.v >> _CAP_BITS), nil
    } else {
        return 0, ErrUnsupportType
    }
}

func (self Node) cap() int {
    return int(self.v >> _CAP_BITS)
}

// Set sets the node of given key under object parent
// If the key doesn't exist, it will be append to the last
func (self *Node) Set(key string, node Node) (bool, error) {
    p := self.Get(key)

    if !p.Exists() {
        l := self.len()
        c := self.cap()
        if l == c {
            // TODO: maybe change append size in future
            c += _DEFAULT_NODE_CAP
            mem := unsafe_NewArray(_PAIR_TYPE, c)
            memmove(mem, self.p, _PAIR_SIZE * uintptr(l))
            self.p = mem
        }
        v := self.pairAt(l)
        v.Key = key
        v.Value = node
        self.setCapAndLen(c, l+1)
        return false, nil

    } else if err := p.Check(); err != nil {
        return false, err
    } 

    *p = node
    return true, nil
}

// Unset remove the node of given key under object parent
func (self *Node) Unset(key string) (bool, error) {
    self.must(types.V_OBJECT, "an object")
    p, i := self.skipKey(key)
    if !p.Exists() {
        return false, nil
    } else if err := p.Check(); err != nil {
        return false, err
    }
    
    self.removePair(i)
    return true, nil
}

// SetByIndex sets the node of given index
//
// The index must within parent array's children
func (self *Node) SetByIndex(index int, node Node) (bool, error) {
    p := self.Index(index)
    if !p.Exists() {
        return false, ErrNotExist
    } else if err := p.Check(); err != nil {
        return false, err
    }

    *p = node
    return true, nil
}

// UnsetByIndex remove the node of given index
func (self *Node) UnsetByIndex(index int) (bool, error) {
    var p *Node
    it := self.itype()
    if it == types.V_ARRAY {
        p = self.Index(index)
    }else if it == types.V_OBJECT {
        pr := self.skipIndexPair(index)
        if pr == nil {
           return false, ErrNotExist
        }
        p = &pr.Value
    }else{
        return false, ErrUnsupportType
    }

    if !p.Exists() {
        return false, ErrNotExist
    }

    if it == types.V_ARRAY {
        self.removeNode(index)
    }else if it == types.V_OBJECT {
        self.removePair(index)
    }
    return true, nil
}

// Add appends the given node under array node
func (self *Node) Add(node Node) error {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return err
    }
    if err := self.skipAllIndex(); err != nil {
        return err
    }

    l := self.len()
    c := self.cap()
    if l == c {
        // TODO: maybe change append_extra_size in future
        c += _DEFAULT_NODE_CAP
        mem := unsafe_NewArray(_NODE_TYPE, c)
        memmove(mem, self.p, _NODE_SIZE * uintptr(l))
        self.p = mem
    }

    v := self.nodeAt(l)
    *v = node
    self.setCapAndLen(c, l+1)
    return nil
}

// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self *Node) GetByPath(path ...interface{}) *Node {
    if !self.Valid() {
        return self
    }
    var s = self
    for _, p := range path {
        switch p.(type) {
        case int:
            s = s.Index(p.(int))
            if !s.Valid() {
                return s
            }
        case string:
            s = s.Get(p.(string))
            if !s.Valid() {
                return s
            }
        default:
            panic("path must be either int or string")
        }
    }
    return s
}

// Get loads given key of an object node on demands
func (self *Node) Get(key string) *Node {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return unwrapError(err)
    }
    n, _ := self.skipKey(key)
    return n
}

// Index loads given index of an node on demands,
// node type can be either V_OBJECT or V_ARRAY
func (self *Node) Index(idx int) *Node {
    if err := self.checkRaw(); err != nil {
        return unwrapError(err)
    }

    it := self.itype()
    if it == types.V_ARRAY {
        return self.skipIndex(idx)

    }else if it == types.V_OBJECT {
        pr := self.skipIndexPair(idx)
        if pr == nil {
           return nodeNotExist
        }
        return &pr.Value

    }else{
        return nodeUnsupportType
    }
}

// Values returns iterator for array's children traversal
func (self *Node) Values() (ListIterator, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return ListIterator{}, err
    }
    if err := self.skipAllIndex(); err != nil {
        return ListIterator{}, err
    }
    return ListIterator{Iterator{p: self}}, nil
}

// Properties returns iterator for object's children traversal
func (self *Node) Properties() (ObjectIterator, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return ObjectIterator{}, err
    }
    if err := self.skipAllKey(); err != nil {
        return ObjectIterator{}, err
    }
    return ObjectIterator{Iterator{p: self}}, nil
}

/** Generic Value Converters **/

// Map loads all keys of an object node
func (self *Node) Map() (map[string]interface{}, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil, err
    }
    if err := self.loadAllKey(); err != nil {
        return nil, err
    }
    return self.toGenericObject()
}

// MapUseNumber loads all keys of an object node, with numeric nodes casted to json.Number
func (self *Node) MapUseNumber() (map[string]interface{}, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil, err
    }
    if err := self.loadAllKey(); err != nil {
        return nil, err
    }
    return self.toGenericObjectUseNumber()
}

// MapUseNode scans both parsed and non-parsed chidren nodes, 
// and map them by their keys
func (self *Node) MapUseNode() (map[string]Node, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil, err
    }
    if err := self.skipAllKey(); err != nil {
        return nil, err
    }
    return self.toGenericObjectUseNode()
}

// MapUnsafe exports the underlying pointer to its children map
// WARN: don't use it unless you know what you are doing
func (self *Node) UnsafeMap() ([]Pair, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil, err
    }
    if err := self.skipAllKey(); err != nil {
        return nil, err
    }
    s := ptr2slice(self.p, int(self.len()), self.cap())
    return *(*[]Pair)(s), nil
}

// Array loads all indexes of an array node
func (self *Node) Array() ([]interface{}, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.loadAllIndex(); err != nil {
        return nil, err
    }
    return self.toGenericArray()
}

// ArrayUseNumber loads all indexes of an array node, with numeric nodes casted to json.Number
func (self *Node) ArrayUseNumber() ([]interface{}, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.loadAllIndex(); err != nil {
        return nil, err
    }
    return self.toGenericArrayUseNumber()
}

// ArrayUseNode copys both parsed and non-parsed chidren nodes, 
// and indexes them by original order
func (self *Node) ArrayUseNode() ([]Node, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.loadAllIndex(); err != nil {
        return nil, err
    }
    return self.toGenericArrayUseNode()
}

// ArrayUnsafe exports the underlying pointer to its children array
// WARN: don't use it unless you know what you are doing
func (self *Node) UnsafeArray() ([]Node, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.loadAllIndex(); err != nil {
        return nil, err
    }
    s := ptr2slice(self.p, self.len(), self.cap())
    return *(*[]Node)(s), nil
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic type.
// WARN: all numberic nodes are casted to float64
func (self *Node) Interface() (interface{}, error) {
    if err := self.checkRaw(); err != nil {
        return nil, err
    }
    switch self.t {
        case V_ERROR         : return nil, self.Check()
        case types.V_NULL    : return nil, nil
        case types.V_TRUE    : return true, nil
        case types.V_FALSE   : return false, nil
        case types.V_ARRAY   : return self.toGenericArray()
        case types.V_OBJECT  : return self.toGenericObject()
        case types.V_STRING  : return addr2str(self.p, self.v), nil
        case _V_NUMBER       : 
            v, err := numberToFloat64(self)
            if err != nil {
                return nil, err
            }
            return v, nil
        case _V_ARRAY_LAZY   :
            if err := self.loadAllIndex(); err != nil {
                return nil, err
            }
            return self.toGenericArray()
        case _V_OBJECT_LAZY  :
            if err := self.loadAllKey(); err != nil {
                return nil, err
            }
            return self.toGenericObject()
        default              : return nil,  ErrUnsupportType
    }
}

// InterfaceUseNumber works same with Interface()
// except numberic nodes  are casted to json.Number
func (self *Node) InterfaceUseNumber() (interface{}, error) {
    if err := self.checkRaw(); err != nil {
        return nil, err
    }
    switch self.t {
        case V_ERROR         : return nil, self.Check()
        case types.V_NULL    : return nil, nil
        case types.V_TRUE    : return true, nil
        case types.V_FALSE   : return false, nil
        case types.V_ARRAY   : return self.toGenericArrayUseNumber()
        case types.V_OBJECT  : return self.toGenericObjectUseNumber()
        case types.V_STRING  : return addr2str(self.p, self.v), nil
        case _V_NUMBER       : return toNumber(self), nil
        case _V_ARRAY_LAZY   :
            if err := self.loadAllIndex(); err != nil {
                return nil, err
            }
            return self.toGenericArrayUseNumber()
        case _V_OBJECT_LAZY  :
            if err := self.loadAllKey(); err != nil {
                return nil, err
            }
            return self.toGenericObjectUseNumber()
        default              : return nil, ErrUnsupportType
    }
}

// InterfaceUseNode clone itself as a new node, 
// or its children as map[string]Node (or []Node)
func (self *Node) InterfaceUseNode() (interface{}, error) {
    if err := self.checkRaw(); err != nil {
        return nil, err
    }
    switch self.t {
        case types.V_ARRAY   : return self.toGenericArrayUseNode()
        case types.V_OBJECT  : return self.toGenericObjectUseNode()
        case _V_ARRAY_LAZY   :
            if err := self.skipAllIndex(); err != nil {
                return nil, err
            }
            return self.toGenericArrayUseNode()
        case _V_OBJECT_LAZY  :
            if err := self.loadAllKey(); err != nil {
                return nil, err
            }
            return self.toGenericObjectUseNode()
        default              : return *self, nil
    }
}

/**---------------------------------- Internal Helper Methods ----------------------------------**/

var (
    _NODE_TYPE = rt.UnpackEface(Node{}).Type
    _PAIR_TYPE = rt.UnpackEface(Pair{}).Type
)

func (self *Node) setCapAndLen(cap int, len int) {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        self.v = int64(len&_LEN_MASK | cap<<_CAP_BITS)
    } else {
        panic("value does not have a length")
    }
}

func (self *Node) unsafe_next() *Node {
    return (*Node)(unsafe.Pointer(uintptr(unsafe.Pointer(self)) + _NODE_SIZE))
}

func (self *Pair) unsafe_next() *Pair {
    return (*Pair)(unsafe.Pointer(uintptr(unsafe.Pointer(self)) + _PAIR_SIZE))
}

func (self *Node) must(t types.ValueType, s string) {
    if err := self.checkRaw(); err != nil {
        panic(err)
    }
    if err := self.Check(); err != nil {
        panic(err)
    }
    if  self.itype() != t {
        panic("value cannot be represented as " + s)
    }
}

func (self *Node) should(t types.ValueType, s string) error {
    if err := self.checkRaw(); err != nil {
        return err
    }
    if  self.itype() != t {
        return ErrUnsupportType
    }
    return nil
}

func (self *Node) nodeAt(i int) *Node {
    var p = self.p
    if self.isLazy() {
        _, stack := self.getParserAndArrayStack()
        p = *(*unsafe.Pointer)(unsafe.Pointer(&stack.v))
    }
    return (*Node)(unsafe.Pointer(uintptr(p) + uintptr(i)*_NODE_SIZE))
}

func (self *Node) pairAt(i int) *Pair {
    var p = self.p
    if self.isLazy() {
        _, stack := self.getParserAndObjectStack()
        p = *(*unsafe.Pointer)(unsafe.Pointer(&stack.v))
    }
    return (*Pair)(unsafe.Pointer(uintptr(p) + uintptr(i)*_PAIR_SIZE))
}

func (self *Node) findKey(key string) (*Node, int) {
    nb := self.len()
    if nb <= 0 {
        return nil, -1
    }

    var p *Pair
    if !self.isLazy() {
        p = (*Pair)(self.p)
    } else {
        s := (*parseObjectStack)(self.p)
        p = &s.v[0]
    }

    if p.Key == key {
        return &p.Value, 0
    }
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        if p.Key == key {
            return &p.Value, i
        }
    }

    /* not found */
    return nil, -1
}

func (self *Node) getParserAndArrayStack() (*Parser, *parseArrayStack) {
    stack := (*parseArrayStack)(self.p)
    ret := (*rt.GoSlice)(unsafe.Pointer(&stack.v))
    ret.Len = self.len()
    ret.Cap = self.cap()
    return &stack.parser, stack
}

func (self *Node) getParserAndObjectStack() (*Parser, *parseObjectStack) {
    stack := (*parseObjectStack)(self.p)
    ret := (*rt.GoSlice)(unsafe.Pointer(&stack.v))
    ret.Len = self.len()
    ret.Cap = self.cap()
    return &stack.parser, stack
}

func (self *Node) skipAllIndex() error {
    if !self.isLazy() {
        return nil
    }
    var err types.ParsingError
    parser, stack := self.getParserAndArrayStack()
    parser.skipValue = true
    parser.noLazy = true
    *self, err = parser.decodeArray(stack.v)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) skipAllKey() error {
    if !self.isLazy() {
        return nil
    }
    var err types.ParsingError
    parser, stack := self.getParserAndObjectStack()
    parser.skipValue = true
    parser.noLazy = true
    *self, err = parser.decodeObject(stack.v)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) skipNextNode() *Node {
    if !self.isLazy() {
        return nil
    }

    parser, stack := self.getParserAndArrayStack()
    ret := stack.v
    sp := parser.p
    ns := len(parser.s)

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        return newSyntaxError(parser.syntaxError(types.ERR_EOF))
    }

    /* check for empty array */
    if parser.s[parser.p] == ']' {
        parser.p++
        self.setArray(ret)
        return nil
    }

    var val Node
    /* skip the value */
    if start, err := parser.skip(); err != 0 {
        return newSyntaxError(parser.syntaxError(err))
    }else{
        t := switchRawType(parser.s[start])
        if t == _V_NONE {
            return newSyntaxError(parser.syntaxError(types.ERR_INVALID_CHAR))
        }
        val = newRawNode(parser.s[start:parser.p], t)
    }

    /* add the value to result */
    ret = append(ret, val)
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        return newSyntaxError(parser.syntaxError(types.ERR_EOF))
    }

    /* check for the next character */
    switch parser.s[parser.p] {
    case ',':
        parser.p++
        self.setLazyArray(parser, ret)
        return &ret[len(ret)-1]
    case ']':
        parser.p++
        self.setArray(ret)
        return &ret[len(ret)-1]
    default:
        return newSyntaxError(parser.syntaxError(types.ERR_INVALID_CHAR))
    }
}

func (self *Node) skipNextPair() (*Pair) {
    if !self.isLazy() {
        return nil
    }

    parser, stack := self.getParserAndObjectStack()
    ret := stack.v
    sp := parser.p
    ns := len(parser.s)

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        return &Pair{"", *newSyntaxError(parser.syntaxError(types.ERR_EOF))}
    }

    /* check for empty object */
    if parser.s[parser.p] == '}' {
        parser.p++
        self.setObject(ret)
        return nil
    }

    /* decode one pair */
    var val Node
    var njs types.JsonState
    var err types.ParsingError

    /* decode the key */
    if njs = parser.decodeValue(); njs.Vt != types.V_STRING {
        return &Pair{"", *newSyntaxError(parser.syntaxError(types.ERR_INVALID_CHAR))}
    }

    /* extract the key */
    idx := parser.p - 1
    key := parser.s[njs.Iv:idx]

    /* check for escape sequence */
    if njs.Ep != -1 {
        if key, err = unquote.String(key); err != 0 {
            return &Pair{key, *newSyntaxError(parser.syntaxError(err))}
        }
    }

    /* expect a ':' delimiter */
    if err = parser.delim(); err != 0 {
        return &Pair{key, *newSyntaxError(parser.syntaxError(err))}
    }

    /* skip the value */
    if start, err := parser.skip(); err != 0 {
        return &Pair{key, *newSyntaxError(parser.syntaxError(err))}
    }else{
        t := switchRawType(parser.s[start])
        if t == _V_NONE {
            return &Pair{key, *newSyntaxError(parser.syntaxError(types.ERR_INVALID_CHAR))}
        }
        val = newRawNode(parser.s[start:parser.p], t)
    }

    /* add the value to result */
    ret = append(ret, Pair{Key: key, Value: val})
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        return &Pair{key, *newSyntaxError(parser.syntaxError(types.ERR_EOF))}
    }

    /* check for the next character */
    switch parser.s[parser.p] {
    case ',':
        parser.p++
        self.setLazyObject(parser, ret)
        return &ret[len(ret)-1]
    case '}':
        parser.p++
        self.setObject(ret)
        return &ret[len(ret)-1]
    default:
        return &Pair{key, *newSyntaxError(parser.syntaxError(types.ERR_INVALID_CHAR))}
    }
}

func (self *Node) skipKey(key string) (*Node, int) {
    node, pos := self.findKey(key)
    if node != nil {
        return node, pos
    }
    if !self.isLazy() {
        return &Node{}, -1
    }

    // lazy load
    var i = self.len()
    for last := self.skipNextPair(); last != nil; last = self.skipNextPair() {
        if last.Value.Check() != nil {
            return &last.Value, -1
        }
        if last.Key == key {
            return &last.Value, i
        }
        i++
    }
    return &Node{}, -1
}

func (self *Node) skipIndex(index int) *Node {
    nb := self.len()
    if nb > index {
        v := self.nodeAt(index)
        return v
    }
    if !self.isLazy() {
        return &Node{}
    }

    // lazy load
    for last := self.skipNextNode(); last != nil; last = self.skipNextNode(){
        if last.Check() != nil {
            return last
        }
        if self.len() > index {
            return last
        }
    }

    return &Node{}
}

func (self *Node) skipIndexPair(index int) *Pair {
    nb := self.len()
    if nb > index {
        return self.pairAt(index)
    }
    if !self.isLazy() {
        return nil
    }

    // lazy load
    for last := self.skipNextPair(); last != nil; last = self.skipNextPair(){
        if last.Value.Check() != nil {
            return last
        }
        if self.len() > index {
            return last
        }
    }

    return nil
}

func (self *Node) loadAllIndex() error {
    if !self.isLazy() {
        return nil
    }
    var err types.ParsingError
    parser, stack := self.getParserAndArrayStack()
    parser.noLazy = true
    *self, err = parser.decodeArray(stack.v)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) loadAllKey() error {
    if !self.isLazy() {
        return nil
    }
    var err types.ParsingError
    parser, stack := self.getParserAndObjectStack()
    parser.noLazy = true
    *self, err = parser.decodeObject(stack.v)
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) removeNode(i int) {
    nb := self.len() - 1
    node := self.nodeAt(i)
    if i == nb {
        self.setCapAndLen(self.cap(), nb)
        *node = Node{}
        return
    }

    from := self.nodeAt(i + 1)
    memmove(unsafe.Pointer(node), unsafe.Pointer(from), _NODE_SIZE * uintptr(nb - i))

    last := self.nodeAt(nb)
    *last = Node{}
    
    self.setCapAndLen(self.cap(), nb)
}

func (self *Node) removePair(i int) {
    nb := self.len() - 1
    node := self.pairAt(i)
    if i == nb {
        self.setCapAndLen(self.cap(), nb)
        *node = Pair{}
        return
    }

    from := self.pairAt(i + 1)
    memmove(unsafe.Pointer(node), unsafe.Pointer(from), _PAIR_SIZE * uintptr(nb - i))

    last := self.pairAt(nb)
    *last = Pair{}
    
    self.setCapAndLen(self.cap(), nb)
}

func (self *Node) toGenericArray() ([]interface{}, error) {
    nb := self.len()
    ret := make([]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var p = (*Node)(self.p)
    x, err := p.Interface()
    if err != nil {
        return nil, err
    }
    ret[0] = x

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        x, err := p.Interface()
        if err != nil {
            return nil, err
        }
        ret[i] = x
    }

    /* all done */
    return ret, nil
}

func (self *Node) toGenericArrayUseNumber() ([]interface{}, error) {
    nb := self.len()
    ret := make([]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var p = (*Node)(self.p)
    x, err := p.InterfaceUseNumber()
    if err != nil {
        return nil, err
    }
    ret[0] = x

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        x, err := p.InterfaceUseNumber()
        if err != nil {
            return nil, err
        }
        ret[i] = x
    }

    /* all done */
    return ret, nil
}

func (self *Node) toGenericArrayUseNode() ([]Node, error) {
    var nb = self.len()
    var out = make([]Node, nb)
    if nb == 0 {
        return out, nil
    }

    var p = (*Node)(self.p)
    out[0] = *p
    if err := p.Check(); err != nil {
        return nil, err
    }

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        if err := p.Check(); err != nil {
            return nil, err
        }
        out[i] = *p
    }

    return out, nil
}

func (self *Node) toGenericObject() (map[string]interface{}, error) {
    nb := self.len()
    ret := make(map[string]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var p = (*Pair)(self.p)
    x, err := p.Value.Interface()
    if err != nil {
        return nil, err
    }
    ret[p.Key] = x

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        x, err := p.Value.Interface()
        if err != nil {
            return nil, err
        }
        ret[p.Key] = x
    }

    /* all done */
    return ret, nil
}


func (self *Node) toGenericObjectUseNumber() (map[string]interface{}, error) {
    nb := self.len()
    ret := make(map[string]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var p = (*Pair)(self.p)
    x, err := p.Value.InterfaceUseNumber()
    if err != nil {
        return nil, err
    }
    ret[p.Key] = x

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        x, err := p.Value.InterfaceUseNumber()
        if err != nil {
            return nil, err
        }
        ret[p.Key] = x
    }

    /* all done */
    return ret, nil
}

func (self *Node) toGenericObjectUseNode() (map[string]Node, error) {
    var nb = self.len()
    var out = make(map[string]Node, nb)
    if nb == 0 {
        return out, nil
    }

    var p = (*Pair)(self.p)
    out[p.Key] = p.Value
    if err := p.Value.Check(); err != nil {
        return nil, err
    }

    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        if err := p.Value.Check(); err != nil {
            return nil, err
        }
        out[p.Key] = p.Value
    }

    /* all done */
    return out, nil
}

/**------------------------------------ Factory Methods ------------------------------------**/

var (
    nullNode  = Node{t: types.V_NULL}
    trueNode  = Node{t: types.V_TRUE}
    falseNode = Node{t: types.V_FALSE}

    emptyArrayNode  = Node{t: types.V_ARRAY}
    emptyObjectNode = Node{t: types.V_OBJECT}
)

// NewNull creates a node of type V_NULL
func NewNull() Node {
    return Node{
        v: 0,
        p: nil,
        t: types.V_NULL,
    }
}

// NewBool creates a node of type bool:
//  If v is true, returns V_TRUE node
//  If v is false, returns V_FALSE node
func NewBool(v bool) Node {
    var t = types.V_FALSE
    if v {
        t = types.V_TRUE
    }
    return Node{
        v: 0,
        p: nil,
        t: t,
    }
}

// NewNumber creates a json.Number node
// v must be a decimal string complying with RFC8259
func NewNumber(v string) Node {
    return Node{
        v: int64(len(v) & _LEN_MASK),
        p: str2ptr(v),
        t: _V_NUMBER,
    }
}

func toNumber(node *Node) json.Number {
    return json.Number(addr2str(node.p, node.v))
}

func numberToFloat64(node *Node) (float64, error) {
    ret,err := toNumber(node).Float64()
    if err != nil {
        return 0, err
    }
    return ret, nil
}

func numberToInt64(node *Node) (int64, error) {
    ret,err := toNumber(node).Int64()
    if err != nil {
        return 0, err
    }
    return ret, nil
}

func newBytes(v []byte) Node {
    return Node{
        t: types.V_STRING,
        p: mem2ptr(v),
        v: int64(len(v) & _LEN_MASK),
    }
}

// NewString creates a node of type string
func NewString(v string) Node {
    return Node{
        t: types.V_STRING,
        p: str2ptr(v),
        v: int64(len(v) & _LEN_MASK),
    }
}

// NewArray creates a node of type V_ARRAY,
// using v as its underlying children
func NewArray(v []Node) Node {
    return Node{
        t: types.V_ARRAY,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: *(*unsafe.Pointer)(unsafe.Pointer(&v)),
    }
}

func (self *Node) setArray(v []Node) {
    self.t = types.V_ARRAY
    self.setCapAndLen(cap(v), len(v))
    self.p = *(*unsafe.Pointer)(unsafe.Pointer(&v))
}

// NewObject creates a node of type V_OBJECT,
// using v as its underlying children
func NewObject(v []Pair) Node {
    return Node{
        t: types.V_OBJECT,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: *(*unsafe.Pointer)(unsafe.Pointer(&v)),
    }
}

func (self *Node) setObject(v []Pair) {
    self.t = types.V_OBJECT
    self.setCapAndLen(cap(v), len(v))
    self.p = *(*unsafe.Pointer)(unsafe.Pointer(&v))
}

type parseObjectStack struct {
    parser Parser
    v      []Pair
}

type parseArrayStack struct {
    parser Parser
    v      []Node
}

func newLazyArray(p *Parser, v []Node) Node {
    s := new(parseArrayStack)
    s.parser = *p
    s.v = v
    return Node{
        t: _V_ARRAY_LAZY,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setLazyArray(p *Parser, v []Node) {
    s := new(parseArrayStack)
    s.parser = *p
    s.v = v
    self.t = _V_ARRAY_LAZY
    self.setCapAndLen(cap(v), len(v))
    self.p = (unsafe.Pointer)(s)
}

func newLazyObject(p *Parser, v []Pair) Node {
    s := new(parseObjectStack)
    s.parser = *p
    s.v = v
    return Node{
        t: _V_OBJECT_LAZY,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setLazyObject(p *Parser, v []Pair) {
    s := new(parseObjectStack)
    s.parser = *p
    s.v = v
    self.t = _V_OBJECT_LAZY
    self.setCapAndLen(cap(v), len(v))
    self.p = (unsafe.Pointer)(s)
}

func newRawNode(str string, typ types.ValueType) Node {
    return Node{
        t: _V_RAW | typ,
        p: str2ptr(str),
        v: int64(len(str) & _LEN_MASK),
    }
}

func (self *Node) parseRaw() Node {
    raw := addr2str(self.p, self.v)
    parser := NewParser(raw)
    n, e := parser.Parse()
    if e != 0 {
        return *newSyntaxError(parser.syntaxError(e))
    }
    return n
}

func newError(err types.ParsingError, msg string) *Node {
    return &Node{
        t: V_ERROR,
        v: int64(err),
        p: unsafe.Pointer(&msg),
    }
}

func newSyntaxError(err *decoder.SyntaxError) *Node {
    msg := err.Description()
    return &Node{
        t: V_ERROR,
        v: int64(err.Code),
        p: unsafe.Pointer(&msg),
    }
}

var typeJumpTable = [256]types.ValueType{
    '"' : types.V_STRING,
    '-' : _V_NUMBER,
    '0' : _V_NUMBER,
    '1' : _V_NUMBER,
    '2' : _V_NUMBER,
    '3' : _V_NUMBER,
    '4' : _V_NUMBER,
    '5' : _V_NUMBER,
    '6' : _V_NUMBER,
    '7' : _V_NUMBER,
    '8' : _V_NUMBER,
    '9' : _V_NUMBER,
    '[' : types.V_ARRAY,
    'f' : types.V_FALSE,
    'n' : types.V_NULL,
    't' : types.V_TRUE,
    '{' : types.V_OBJECT,
}

func switchRawType(c byte) types.ValueType {
    return typeJumpTable[c]
}

func unwrapError(err error) *Node {
    if se, ok := err.(*Node); ok {
        return se
    }else if sse, ok := err.(Node); ok {
        return &sse
    }else{
        msg := err.Error()
        return &Node{
            t: V_ERROR,
            v: 0,
            p: unsafe.Pointer(&msg),
        }
    }
}