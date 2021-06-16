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

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/bytedance/sonic/unquote`
)

const (
    _CAP_BITS          = 32
    _LEN_MASK          = 1 << _CAP_BITS - 1
    _APPEND_EXTRA_SIZE = 5

    _NODE_SIZE = unsafe.Sizeof(Node{})
    _PAIR_SIZE = unsafe.Sizeof(Pair{})
)

const (
    _V_NONE       types.ValueType = 0

    _V_NODE_BASE  types.ValueType = 1<<5
    _V_LAZY        types.ValueType = 1 << 7
    _V_RAW         types.ValueType = 1 << 8
    _V_NUMBER                     = _V_NODE_BASE + 1
    _V_ARRAY_LAZY                  = _V_LAZY | types.V_ARRAY
    _V_OBJECT_LAZY                 = _V_LAZY | types.V_OBJECT
    _MASK_LAZY                     = _V_LAZY - 1
    _MASK_RAW = _V_RAW - 1
)

const (
    V_NONE   = 0
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
//    V_NULL   = 2
//    V_TRUE   = 3
//    V_FALSE  = 4
//    V_ARRAY  = 5
//    V_OBJECT = 6
//    V_STRING = 7
//    V_NUMBER = 33
func (self *Node) Type() int {
    return int(self.t & _MASK_LAZY & _MASK_RAW)
}

func (self *Node) itype() types.ValueType {
    return self.t & _MASK_LAZY & _MASK_RAW
}

// Exists returns false only if the node is nil or got by invalid path
func (self *Node) Exists() bool {
    return self != nil && self.t != _V_NONE
}

// IsRaw returns true if the node is type of below three:
//
// 1. V_RAW (never parsed)
// 2. types.V_Object_RAW (partially parsed)
// 3. types.V_Array_RAW (partially parsed)
func (self *Node) IsRaw() bool {
    return self.t&_V_RAW != 0
}

func (self *Node) isLazy() bool {
    return self.t&_V_LAZY != 0
}

/** Simple Value Methods **/

// Raw returns underlying json string of an raw node,
// which usually created by Search() api
func (self *Node) Raw() string {
    if !self.IsRaw() {
        panic("value cannot be represented as raw json")
    }
    return addr2str(self.p, self.v)
}

func (self *Node) checkRaw() {
    if !self.IsRaw() {
        return
    }
    *self = self.parseRaw()
}

// Bool returns bool value represented by this node
//
// If node type is not types.V_TRUE or types.V_FALSE, or V_RAW (must be a bool json value)
// it will panic
func (self *Node) Bool() bool {
    self.checkRaw()
    switch self.t {
        case types.V_TRUE  : return true
        case types.V_FALSE : return false
        default : panic("value cannot be represented as a boolean")
    }
}

// Int64 as above.
func (self *Node) Int64() int64 {
    self.checkRaw()
    switch self.t {
        case _V_NUMBER         : return numberToInt64(self)
        case types.V_TRUE     : return 1
        case types.V_FALSE    : return 0
        default               : panic("value cannot be represented as an integer")
    }
}

// Number as above.
func (self *Node) Number() json.Number {
    self.checkRaw()
    switch self.t {
        case _V_NUMBER         : return toNumber(self)
        default               : panic("value cannot be represented as a json.Number")
    }
}

// String as above.
func (self *Node) String() string {
    self.checkRaw()
    switch self.t {
        case _V_NUMBER        : return toNumber(self).String()
        case types.V_NULL    : return "null"
        case types.V_TRUE    : return "true"
        case types.V_FALSE   : return "false"
        case types.V_STRING  : return addr2str(self.p, self.v)
        default              : panic("value cannot be represented as a simple string")
    }
}

// Float64 as above.
func (self *Node) Float64() float64 {
    self.checkRaw()
    switch self.t {
        case _V_NUMBER        : return numberToFloat64(self)
        case types.V_TRUE    : return 1.0
        case types.V_FALSE   : return 0.0
        default              : panic("value cannot be represented as an integer")
    }
}

/** Sequencial Value Methods **/

// Len returns children count of a array|object|string node
// For partially loaded node, it also works but only counts the parsed children
func (self *Node) Len() int {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        return int(self.v & _LEN_MASK)
    } else if self.t == types.V_STRING {
        return int(self.v)
    } else {
        panic("value does not have a length")
    }
}

// Cap returns malloc capacity of a array|object node for children
func (self *Node) Cap() int {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        return int(self.v >> _CAP_BITS)
    } else {
        panic("value does not have a capacity")
    }
}

// Set sets the given node for the key under object node
func (self *Node) Set(key string, node Node) {
    p := self.Get(key)
    if !p.Exists() {
        l := self.Len()
        c := self.Cap()
        if l == c {
            // TODO: maybe change append_extra_size in future
            c += _APPEND_EXTRA_SIZE
            mem := unsafe_NewArray(_PAIR_TYPE, c)
            memmove(mem, self.p, _PAIR_SIZE * uintptr(l))
            self.p = mem
        }
        v := self.pairAt(l)
        v.Key = key
        v.Value = node
        self.setCapAndLen(c, l+1)
    } else {
        *p = node
    }
}

// SetByIndex sets the given node for the index under array node
//
// The index must within parent array's range
func (self *Node) SetByIndex(index int, node Node) {
    p := self.Index(index)
    if !p.Exists() {
        panic("index to nil node")
    } else {
        *p = node
    }
}

// Add appends the given node under array node
func (self *Node) Add(node Node) {
    self.must(types.V_ARRAY, "an array")
    self.loadAllIndex()
    l := self.Len()
    c := self.Cap()
    if l == c {
        // TODO: maybe change append_extra_size in future
        c += _APPEND_EXTRA_SIZE
        mem := unsafe_NewArray(_NODE_TYPE, c)
        memmove(mem, self.p, _NODE_SIZE * uintptr(l))
        self.p = mem
    }
    v := self.nodeAt(l)
    *v = node
    self.setCapAndLen(c, l+1)
}

// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed
func (self *Node) GetByPath(path ...interface{}) *Node {
    var s = self
    for _, p := range path {
        switch p.(type) {
        case int:
            s = s.Index(p.(int))
            if !s.Exists() {
                return s
            }
        case string:
            s = s.Get(p.(string))
            if !s.Exists() {
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
    self.must(types.V_OBJECT, "an object")
    return self.loadKey(key)
}

// Index loads given index of an array node on demands
func (self *Node) Index(idx int) *Node {
    self.must(types.V_ARRAY, "an array")
    return self.loadIndex(idx)
}

// Values returns iterator for array's children traversal.
// non-parsed children will be represented as V_RAW nodes
func (self *Node) Values() ListIterator {
    self.must(types.V_ARRAY, "an array")
    self.skipAllIndex()
    return ListIterator{Iterator{p: self}}
}

// Properties returns iterator for object's children traversal
// non-parsed children will be represented as V_RAW nodes.
func (self *Node) Properties() ObjectIterator {
    self.must(types.V_OBJECT, "an object")
    self.skipAllKey()
    return ObjectIterator{Iterator{p: self}}
}

/** Generic Value Converters **/

// Map loads all keys of an object node
func (self *Node) Map() map[string]interface{} {
    self.must(types.V_OBJECT, "an object")
    self.loadAllKey()
    return self.toGenericObject()
}

// MapUseNumber loads all keys of an object node, with numeric nodes casted to json.Number
func (self *Node) MapUseNumber() map[string]interface{} {
    self.must(types.V_OBJECT, "an object")
    self.loadAllKey()
    return self.toGenericObjectUseNumber()
}

// MapUseNode scans both parsed and non-parsed chidren nodes, 
// and indexes them on map by their keys
func (self *Node) MapUseNode() map[string]Node {
    self.must(types.V_OBJECT, "an object")
    self.skipAllKey()
    return self.toGenericObjectUseNode()
}

// MapUnsafe exports the internal pointer to its children map
// WARN: don't use it unless you know what you are doing
func (self *Node) UnsafeMap() []Pair {
    self.must(types.V_OBJECT, "an object")
    self.skipAllKey()
    s := ptr2slice(self.p, self.Len(), self.Cap())
    return *(*[]Pair)(s)
}

// Array loads all indexes of an array node
func (self *Node) Array() []interface{} {
    self.must(types.V_ARRAY, "an array")
    self.loadAllIndex()
    return self.toGenericArray()
}

// ArrayUseNumber loads all indexes of an array node, with numeric nodes casted to json.Number
func (self *Node) ArrayUseNumber() []interface{} {
    self.must(types.V_ARRAY, "an array")
    self.loadAllIndex()
    return self.toGenericArrayUseNumber()
}

// ArrayUseNode copys both parsed and non-parsed chidren nodes, 
// and indexes them by original order
func (self *Node) ArrayUseNode() []Node {
    self.must(types.V_ARRAY, "an array")
    self.skipAllIndex()
    return self.toGenericArrayUseNode()
}

// ArrayUnsafe exports the internal pointer to its children array
// WARN: don't use it unless you know what you are doing
func (self *Node) UnsafeArray() []Node {
    self.must(types.V_ARRAY, "an array")
    self.skipAllIndex()
    s := ptr2slice(self.p, self.Len(), self.Cap())
    return *(*[]Node)(s)
}

// Interface loads all children under all pathes from this node,
// and converts itself as generic go type
// all numberic nodes are casted to float64
func (self *Node) Interface() interface{} {
    self.checkRaw()
    switch self.t {
        case types.V_EOF     : panic("invalid value")
        case types.V_NULL    : return nil
        case types.V_TRUE    : return true
        case types.V_FALSE   : return false
        case types.V_ARRAY   : return self.toGenericArray()
        case types.V_OBJECT  : return self.toGenericObject()
        case types.V_STRING  : return addr2str(self.p, self.v)
        case _V_NUMBER        : return numberToFloat64(self)
        case _V_ARRAY_LAZY:
            self.loadAllIndex()
            return self.toGenericArray()
        case _V_OBJECT_LAZY:
            self.loadAllKey()
            return self.toGenericObject()
        default              : panic("not gonna happen")
    }
}

// InterfaceUseNumber works same with Interface()
// except numberic nodes  are casted to json.Number
func (self *Node) InterfaceUseNumber() interface{} {
    self.checkRaw()
    switch self.t {
        case types.V_EOF     : panic("invalid value")
        case types.V_NULL    : return nil
        case types.V_TRUE    : return true
        case types.V_FALSE   : return false
        case types.V_ARRAY   : return self.toGenericArrayUseNumber()
        case types.V_OBJECT  : return self.toGenericObjectUseNumber()
        case types.V_STRING  : return addr2str(self.p, self.v)
        case _V_NUMBER        : return toNumber(self)
        case _V_ARRAY_LAZY:
            self.loadAllIndex()
            return self.toGenericArrayUseNumber()
        case _V_OBJECT_LAZY:
            self.loadAllKey()
            return self.toGenericObjectUseNumber()
        default              : panic("not gonna happen")
    }
}

// InterfaceUseNode clone itself as a new node, 
// or its children as map[string]Node (or []Node)
func (self *Node) InterfaceUseNode() interface{} {
    self.checkRaw()
    switch self.t {
        case types.V_ARRAY   : return self.toGenericArrayUseNode()
        case types.V_OBJECT  : return self.toGenericObjectUseNode()
        case _V_ARRAY_LAZY:
            self.skipAllIndex()
            return self.toGenericArrayUseNode()
        case _V_OBJECT_LAZY:
            self.skipAllKey()
            return self.toGenericObjectUseNode()
        default              : return *self
    }
}

/** Internal Helper Methods **/

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
    self.checkRaw()
    if self.itype() != t {
        panic("value cannot be represented as " + s)
    }
}

func (self *Node) bound(i int) {
    if i < 0 || i >= self.Len() {
        panic("list index out of range")
    }
}

func (self *Node) nodeAt(i int) *Node {
    return (*Node)(unsafe.Pointer(uintptr(self.p) + uintptr(i)*_NODE_SIZE))
}

func (self *Node) pairAt(i int) *Pair {
    return (*Pair)(unsafe.Pointer(uintptr(self.p) + uintptr(i)*_PAIR_SIZE))
}

func (self *Node) findKey(key string) *Node {
    nb := self.Len()
    if nb <= 0 {
        return nil
    }

    var p *Pair
    if !self.isLazy() {
        p = (*Pair)(self.p)
    } else {
        s := (*parseObjectStack)(self.p)
        p = &s.v[0]
    }

    if p.Key == key {
        return &p.Value
    }
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        if p.Key == key {
            return &p.Value
        }
    }

    /* not found */
    return nil
}

func (self *Node) getParserAndArrayStack() (*Parser, *parseArrayStack) {
    stack := (*parseArrayStack)(self.p)
    return &stack.parser, stack
}

func (self *Node) getParserAndObjectStack() (*Parser, *parseObjectStack) {
    stack := (*parseObjectStack)(self.p)
    return &stack.parser, stack
}

func (self *Node) loadAllIndex() {
    if !self.isLazy() {
        return
    }
    var err types.ParsingError
    parser, stack := self.getParserAndArrayStack()
    old := parser.noLazy
    parser.noLazy = true
    start := parser.p
    *self, err = parser.decodeArray(stack.v)
    if err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.noLazy = old
}

func (self *Node) skipAllIndex() {
    if !self.isLazy() {
        return
    }
    var err types.ParsingError
    parser, stack := self.getParserAndArrayStack()
    olds := parser.skipValue
    parser.skipValue = true
    oldl := parser.noLazy
    parser.noLazy = true
    start := parser.p
    *self, err = parser.decodeArray(stack.v)
    if err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.skipValue = olds
    parser.noLazy = oldl
}

func (self *Node) loadAllKey() {
    if !self.isLazy() {
        return
    }
    var err types.ParsingError
    parser, stack := self.getParserAndObjectStack()
    old := parser.noLazy
    parser.noLazy = true
    start := parser.p
    *self, err = parser.decodeObject(stack.v)
    if err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.noLazy = old
}

func (self *Node) skipAllKey() {
    if !self.isLazy() {
        return
    }
    var err types.ParsingError
    parser, stack := self.getParserAndObjectStack()
    olds := parser.skipValue
    parser.skipValue = true
    oldl := parser.noLazy
    parser.noLazy = true
    start := parser.p
    *self, err = parser.decodeObject(stack.v)
    if err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.skipValue = olds
    parser.noLazy = oldl
}

func (self *Node) loadIndex(index int) *Node {
    nb := self.Len()
    if nb > index {
        return self.nodeAt(index)
    }
    if !self.isLazy() {
        return &Node{}
    }

    // lazy load
    for last := self.loadNextNode(); last != nil; last = self.loadNextNode(){
        if self.Len() > index {
            return last
        }
        if !self.isLazy() {
            break
        }
    }

    return &Node{}
}

func (self *Node) loadNextNode() *Node {
    stack := (*parseArrayStack)(self.p)
    ret := stack.v
    parser := &stack.parser
    sp := parser.p
    ns := len(parser.s)
    start := sp

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        panic(parser.ExportError(types.ERR_EOF, start))
    }

    /* check for empty array */
    if parser.s[parser.p] == ']' {
        parser.p++
        self.setArray(ret)
        return nil
    }

    var val Node
    var err types.ParsingError

    /* decode the value */
    old := parser.noLazy
    parser.noLazy = true
    start = parser.p
    if val, err = parser.Parse(); err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.noLazy = old

    /* add the value to result */
    ret = append(ret, val)
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        panic(parser.ExportError(types.ERR_EOF, start))
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
        panic(parser.ExportError(types.ERR_INVALID_CHAR, start))
    }
}

func (self *Node) loadKey(key string) *Node {
    node := self.findKey(key)
    if node != nil {
        return node
    }
    if !self.isLazy() {
        return &Node{}
    }

    // lazy load
    for last := self.loadNextPair(); last != nil; last = self.loadNextPair() {
        if last.Key == key {
            return &last.Value
        }
        if !self.isLazy() {
            break
        }
    }

    return &Node{}
}

func (self *Node) loadNextPair() (*Pair) {
    stack := (*parseObjectStack)(self.p)
    ret := stack.v
    parser := &stack.parser
    sp := parser.p
    ns := len(parser.s)
    start := sp

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        panic(parser.ExportError(types.ERR_EOF, start))
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
    start = parser.p
    if njs = parser.decodeValue(); njs.Vt != types.V_STRING {
        panic(parser.ExportError(types.ERR_INVALID_CHAR, start))
    }

    /* extract the key */
    idx := parser.p - 1
    key := parser.s[njs.Iv:idx]

    /* check for escape sequence */
    if njs.Ep != -1 {
        if key, err = unquote.String(key); err != 0 {
            panic(parser.ExportError(err, start))
        }
    }

    /* expect a ':' delimiter */
    start = parser.p
    if err = parser.delim(); err != 0 {
        panic(parser.ExportError(err, start))
    }

    /* decode the value */
    old := parser.noLazy
    parser.noLazy = true
    start = parser.p
    if val, err = parser.Parse(); err != 0 {
        panic(parser.ExportError(err, start))
    }
    parser.noLazy = old

    /* add the value to result */
    ret = append(ret, Pair{Key: key, Value: val})
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        panic(parser.ExportError(types.ERR_EOF, start))
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
        panic(parser.ExportError(types.ERR_INVALID_CHAR, start))
    }
}

func (self *Node) toGenericArray() []interface{} {
    nb := self.Len()
    ret := make([]interface{}, nb)
    if nb == 0 {
        return ret
    }

    /* convert each item */
    var p = (*Node)(self.p)
    ret[0] = p.Interface()
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        ret[i] = p.Interface()
    }

    /* all done */
    return ret
}

func (self *Node) toGenericArrayUseNumber() []interface{} {
    nb := self.Len()
    ret := make([]interface{}, nb)
    if nb == 0 {
        return ret
    }

    /* convert each item */
    var p = (*Node)(self.p)
    ret[0] = p.InterfaceUseNumber()
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        ret[i] = p.InterfaceUseNumber()
    }

    /* all done */
    return ret
}

func (self *Node) toGenericArrayUseNode() []Node {
    var nb = self.Len()
    var out = make([]Node, nb)
    if nb == 0 {
        return out
    }

    var p = (*Node)(self.p)
    out[0] = *p
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        out[i] = *p
    }

    return out
}

func (self *Node) toGenericObject() map[string]interface{} {
    nb := self.Len()
    ret := make(map[string]interface{}, nb)
    if nb == 0 {
        return ret
    }

    /* convert each item */
    var p = (*Pair)(self.p)
    ret[p.Key] = p.Value.Interface()
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        ret[p.Key] = p.Value.Interface()
    }

    /* all done */
    return ret
}


func (self *Node) toGenericObjectUseNumber() map[string]interface{} {
    nb := self.Len()
    ret := make(map[string]interface{}, nb)
    if nb == 0 {
        return ret
    }

    /* convert each item */
    var p = (*Pair)(self.p)
    ret[p.Key] = p.Value.InterfaceUseNumber()
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        ret[p.Key] = p.Value.InterfaceUseNumber()
    }

    /* all done */
    return ret
}

func (self *Node) toGenericObjectUseNode() map[string]Node {
    var nb = self.Len()
    var out = make(map[string]Node, nb)
    if nb == 0 {
        return out
    }

    var p = (*Pair)(self.p)
    out[p.Key] = p.Value
    for i := 1; i < nb; i++ {
        p = p.unsafe_next()
        out[p.Key] = p.Value
    }

    /* all done */
    return out
}

/** Internal Factory Methods **/

var (
    nullNode  = Node{t: types.V_NULL}
    trueNode  = Node{t: types.V_TRUE}
    falseNode = Node{t: types.V_FALSE}

    emptyArrayNode  = Node{t: types.V_ARRAY}
    emptyObjectNode = Node{t: types.V_OBJECT}
)

func newNumber(v string) Node {
    return Node{
        v: int64(len(v)),
        p: str2ptr(v),
        t: _V_NUMBER,
    }
}

func toNumber(node *Node) json.Number {
    return json.Number(addr2str(node.p, node.v))
}

func numberToFloat64(node *Node) float64 {
    ret,err := toNumber(node).Float64()
    if err != nil {
        panic(err)
    }
    return ret
}

func numberToInt64(node *Node) int64 {
    ret,err := toNumber(node).Int64()
    if err != nil {
        panic(err)
    }
    return ret
}

func newBytes(v []byte) Node {
    return Node{
        t: types.V_STRING,
        p: mem2ptr(v),
        v: int64(len(v)),
    }
}

func newString(v string) Node {
    return Node{
        t: types.V_STRING,
        p: str2ptr(v),
        v: int64(len(v)),
    }
}

func newArray(v []Node) Node {
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

func newObject(v []Pair) Node {
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
        v: int64(len(str)),
    }
}

func (self *Node) parseRaw() Node {
    raw := addr2str(self.p, self.v)
    parser := NewParser(raw)
    n, e := parser.Parse()
    if e != 0 {
        panic(parser.ExportError(e, 0))
    }
    return n
}

func switchRawType(c byte) types.ValueType {
    var t types.ValueType = _V_NONE
    switch c {
        case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' :
            t = _V_NUMBER
        case '"' :
            t = types.V_STRING
        case 'n' :
            t = types.V_NULL
        case 't' :
            t = types.V_TRUE
        case 'f' :
            t = types.V_FALSE
        case '[' :
            t = types.V_ARRAY
        case '{' :
            t = types.V_OBJECT
    }
    return t
}