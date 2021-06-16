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
    `fmt`
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/bytedance/sonic/unquote`
)

const (
    _MAP_THRESHOLD     = 5
    _CAP_BITS          = 32
    _LEN_MASK          = 1 << _CAP_BITS - 1
    _APPEND_EXTRA_SIZE = 5

    _NODE_SIZE = unsafe.Sizeof(Node{})
    _PAIR_SIZE = unsafe.Sizeof(Pair{})
)

const (
    V_NODE_BASE  types.ValueType = 32768
    V_RAW        types.ValueType = 1 << 16
    V_NUMBER                     = V_NODE_BASE + 0
    V_ARRAY_RAW                  = V_RAW | types.V_ARRAY
    V_OBJECT_RAW                 = V_RAW | types.V_OBJECT
    MASK_RAW                     = V_RAW - 1
)

type Node struct {
    v int64
    t types.ValueType
    p unsafe.Pointer
    m map[string]unsafe.Pointer
}

/** Node Type Accessor **/

func (self *Node) Type() types.ValueType {
    return self.t & MASK_RAW
}

/** Simple Value Methods **/

// Raw returns underlying json string of an raw node,
// which usually created by Search() api
func (self *Node) Raw() string {
    if self.t != V_RAW {
        panic("value cannot be represented as raw json")
    }
    return addr2str(self.p, self.v)
}

// Bool returns bool value represented by this node
//
// If node type is not types.V_TRUE or types.V_FALSE, or V_RAW (must be a bool json value)
// it will panic
func (self *Node) Bool() bool {
    switch self.t {
        case types.V_TRUE  : return true
        case types.V_FALSE : return false
        case V_RAW : 
            n := self.parseRaw()
            return n.Bool()
        default : panic("value cannot be represented as a boolean")
    }
}

// Int64 as above.
func (self *Node) Int64() int64 {
    switch self.t {
        case V_NUMBER         : return numberToInt64(self)
        case types.V_TRUE     : return 1
        case types.V_FALSE    : return 0
        case V_RAW :
            n := self.parseRaw()
            return n.Int64()
        default               : panic("value cannot be represented as an integer")
    }
}

// Number as above.
func (self *Node) Number() json.Number {
    switch self.t {
        case V_RAW : 
            n := self.parseRaw()
            return n.Number()
        case V_NUMBER         : return toNumber(self)
        default               : panic("value cannot be represented as a json.Number")
    }
}

// String as above.
func (self *Node) String() string {
    switch self.t {
        case V_NUMBER        : return toNumber(self).String()
        case types.V_NULL    : return "null"
        case types.V_TRUE    : return "true"
        case types.V_FALSE   : return "false"
        case types.V_STRING  : return addr2str(self.p, self.v)
        case V_RAW : 
            n := self.parseRaw()
            return n.String()
        default              : panic("value cannot be represented as a simple string")
    }
}

// Float64 as above.
func (self *Node) Float64() float64 {
    switch self.t {
        case V_NUMBER        : return numberToFloat64(self)
        case types.V_TRUE    : return 1.0
        case types.V_FALSE   : return 0.0
        case V_RAW : 
            n := self.parseRaw()
            return n.Float64()
        default              : panic("value cannot be represented as an integer")
    }
}

// IsRaw returns true if the node is type of below three:
//
// 1. V_RAW (never parsed)
// 2. types.V_Object_RAW (partially parsed)
// 3. types.V_Array_RAW (partially parsed)
func (self *Node) IsRaw() bool {
    return self.t&V_RAW != 0
}

/** Sequencial Value Methods **/

// Len returns children count of a array|object|string node
// For partially loaded node, it also works but only counts the parsed children
func (self *Node) Len() int {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == V_ARRAY_RAW || self.t == V_OBJECT_RAW {
        return int(self.v & _LEN_MASK)
    } else if self.t == types.V_STRING {
        return int(self.v)
    } else {
        panic("value does not have a length")
    }
}

// Cap returns malloc capacity of a array|object node for children
func (self *Node) Cap() int {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == V_ARRAY_RAW || self.t == V_OBJECT_RAW {
        return int(self.v >> _CAP_BITS)
    } else {
        panic("value does not have a capacity")
    }
}

// Set sets the given node for the key under object node
func (self *Node) Set(key string, node Node) {
    p := self.Get(key)
    if p == nil {
        l := self.Len()
        c := self.Cap()
        if l == c {
            // TODO: maybe change append_extra_size in future
            c += _APPEND_EXTRA_SIZE
            mem := unsafe_NewArray(_PAIR_TYPE, c)
            memmove(mem, self.p, _PAIR_SIZE*uintptr(l))
            self.p = mem
        }
        v := self.pairAt(l)
        v.Key = key
        v.Value = node
        if self.m != nil {
            self.m[key] = unsafe.Pointer(&v.Value)
        }
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
    if p == nil {
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
        memmove(mem, self.p, _NODE_SIZE*uintptr(l))
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
        case string:
            s = s.Get(p.(string))
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

// Values returns iterator for array's children traversal
func (self *Node) Values() ListIterator {
    self.must(types.V_ARRAY, "an array")
    self.loadAllIndex()
    return ListIterator{Iterator{p: self}}
}

// Properties returns iterator for object's children traversal
func (self *Node) Properties() ObjectIterator {
    self.must(types.V_OBJECT, "an object")
    self.loadAllKey()
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

// Interface loads all children under all pathes from this node,
// and converts itself as generic go type
// all numberic nodes are casted to float64
func (self *Node) Interface() interface{} {
    switch self.t {
        case types.V_EOF     : panic("invalid value")
        case types.V_NULL    : return nil
        case types.V_TRUE    : return true
        case types.V_FALSE   : return false
        case types.V_ARRAY   : return self.toGenericArray()
        case types.V_OBJECT  : return self.toGenericObject()
        case types.V_STRING  : return addr2str(self.p, self.v)
        case V_NUMBER        :
            return numberToFloat64(self)
        case V_ARRAY_RAW:
            self.loadAllIndex()
            return self.toGenericArray()
        case V_OBJECT_RAW:
            self.loadAllKey()
            return self.toGenericObject()
        case V_RAW : 
            n := self.parseRaw()
            return n.Interface()
        default              : panic("not gonna happen")
    }
}

// InterfaceUseNumber works same with Interface()
// except numberic nodes  are casted to json.Number
func (self *Node) InterfaceUseNumber() interface{} {
    switch self.t {
        case types.V_EOF     : panic("invalid value")
        case types.V_NULL    : return nil
        case types.V_TRUE    : return true
        case types.V_FALSE   : return false
        case types.V_ARRAY   : return self.toGenericArrayUseNumber()
        case types.V_OBJECT  : return self.toGenericObjectUseNumber()
        case types.V_STRING  : return addr2str(self.p, self.v)
        case V_NUMBER        :
            return toNumber(self)
        case V_ARRAY_RAW:
            self.loadAllIndex()
            return self.toGenericArrayUseNumber()
        case V_OBJECT_RAW:
            self.loadAllKey()
            return self.toGenericObjectUseNumber()
        case V_RAW :
            n := self.parseRaw()
            return n.InterfaceUseNumber()
        default              : panic("not gonna happen")
    }
}

/** Internal Helper Methods **/

var (
    _NODE_TYPE = rt.UnpackEface(Node{}).Type
    _PAIR_TYPE = rt.UnpackEface(Pair{}).Type
)

func (self *Node) setCapAndLen(cap int, len int) {
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == V_ARRAY_RAW || self.t == V_OBJECT_RAW {
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
    if self.t == V_RAW {
        *self = self.parseRaw()
    }
    if self.t&MASK_RAW != t {
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
    if self.mapped() {
        if n, ok := (self.m)[key]; !ok {
            return nil
        } else {
            return (*Node)(n)
        }
    }

    nb := self.Len()
    if nb <= 0 {
        return nil
    }

    var p *Pair
    if !self.IsRaw() {
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

func (self *Node) mapped() bool {
    if self.m != nil && len(self.m) == self.Len() {
        return true
    }
    if self.Len() > _MAP_THRESHOLD {
        self.appendMap()
        return true
    }
    return false
}

func (self *Node) appendMap() {
    nb := self.Len()
    nm := len(self.m)
    if nb <= nm {
        return
    }

    if self.m == nil {
        self.m = make(map[string]unsafe.Pointer, nb)
    }

    var p *Pair
    if !self.IsRaw() {
        p = self.pairAt(nm)
    } else {
        s := (*parseObjectStack)(self.p)
        p = &s.v[nm]
    }

    (self.m)[p.Key] = unsafe.Pointer(&p.Value)
    for i := nm + 1; i < nb; i++ {
        p = p.unsafe_next()
        (self.m)[p.Key] = unsafe.Pointer(&p.Value)
    }
}

func (self *Node) loadAllIndex() {
    if !self.IsRaw() {
        return
    }
    var err types.ParsingError
    stack := (*parseArrayStack)(self.p)
    parser := &stack.parser
    old := parser.noLazy
    parser.noLazy = true
    *self, err = parser.decodeArray(stack.v)
    if err != 0 {
        panic(fmt.Sprintf("%s at position %d", err, stack.parser.Pos()))
    }
    parser.noLazy = old
}

func (self *Node) loadAllKey() {
    if !self.IsRaw() {
        return
    }
    var err types.ParsingError
    stack := (*parseObjectStack)(self.p)
    parser := &stack.parser
    old := parser.noLazy
    parser.noLazy = true
    *self, err = parser.decodeObject(stack.v)
    if err != 0 {
        panic(fmt.Sprintf("%s at position %d", err, stack.parser.Pos()))
    }
    parser.noLazy = old
}

func (self *Node) loadIndex(index int) *Node {
    if !self.IsRaw() {
        self.bound(index)
        return self.nodeAt(index)
    }
    nb := self.Len()
    if nb > index {
        return self.nodeAt(index)
    }

    // lazy load

    for last, err := self.loadNextNode(); last != nil; {
        if err != 0 {
            panic(fmt.Sprintf("%s at index %d", err, nb-1))
        }
        if self.Len() > index {
            return last
        }
        last, err = self.loadNextNode()
    }

    return nil
}

func (self *Node) loadNextNode() (*Node, types.ParsingError) {
    stack := (*parseArrayStack)(self.p)
    ret := stack.v
    parser := &stack.parser
    sp := parser.p
    ns := len(parser.s)

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        return nil, types.ERR_EOF
    }

    /* check for empty array */
    if parser.s[parser.p] == ']' {
        parser.p++
        self.setArray(ret)
        return nil, 0
    }

    var val Node
    var err types.ParsingError

    /* decode the value */
    parser.noLazy = true
    if val, err = parser.Parse(); err != 0 {
        return nil, err
    }
    parser.noLazy = false

    /* add the value to result */
    ret = append(ret, val)
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        return &ret[len(ret)-1], types.ERR_EOF
    }

    /* check for the next character */
    switch parser.s[parser.p] {
    case ',':
        parser.p++
        self.setRawArray(parser, ret)
        return &ret[len(ret)-1], 0
    case ']':
        parser.p++
        self.setArray(ret)
        return &ret[len(ret)-1], 0
    default:
        return &ret[len(ret)-1], types.ERR_INVALID_CHAR
    }
}

func (self *Node) loadKey(key string) *Node {
    node := self.findKey(key)
    if node != nil || !self.IsRaw() {
        return node
    }

    // lazy load
    for last, err := self.loadNextPair(); last != nil; {
        if err != 0 {
            panic(fmt.Sprintf("%s at key %s", err, last.Key))
        }
        if last.Key == key {
            return &last.Value
        }
        last, err = self.loadNextPair()
    }

    return nil
}

func (self *Node) loadNextPair() (*Pair, types.ParsingError) {
    stack := (*parseObjectStack)(self.p)
    ret := stack.v
    parser := &stack.parser
    sp := parser.p
    ns := len(parser.s)

    /* check for EOF */
    if parser.p = parser.lspace(sp); parser.p >= ns {
        return nil, types.ERR_EOF
    }

    /* check for empty object */
    if parser.s[parser.p] == '}' {
        parser.p++
        self.setObject(ret)
        return nil, 0
    }

    /* decode one pair */
    var val Node
    var njs types.JsonState
    var err types.ParsingError

    /* decode the key */
    if njs = parser.decodeValue(); njs.Vt != types.V_STRING {
        return nil, types.ERR_INVALID_CHAR
    }

    /* extract the key */
    idx := parser.p - 1
    key := parser.s[njs.Iv:idx]

    /* check for escape sequence */
    if njs.Ep != -1 {
        if key, err = unquote.String(key); err != 0 {
            return nil, err
        }
    }

    /* expect a ':' delimiter */
    if err = parser.delim(); err != 0 {
        return nil, err
    }

    /* decode the value */
    parser.noLazy = true
    if val, err = parser.Parse(); err != 0 {
        return nil, err
    }
    parser.noLazy = false

    /* add the value to result */
    ret = append(ret, Pair{Key: key, Value: val})
    parser.p = parser.lspace(parser.p)

    /* check for EOF */
    if parser.p >= ns {
        return &ret[len(ret)-1], types.ERR_EOF
    }

    /* check for the next character */
    switch parser.s[parser.p] {
    case ',':
        parser.p++
        self.setRawObject(parser, ret)
        return &ret[len(ret)-1], 0
    case '}':
        parser.p++
        self.setObject(ret)
        return &ret[len(ret)-1], 0
    default:
        return &ret[len(ret)-1], types.ERR_INVALID_CHAR
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
        t: V_NUMBER,
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

func newRawArray(p *Parser, v []Node) Node {
    s := new(parseArrayStack)
    s.parser = *p
    s.v = v
    return Node{
        t: V_ARRAY_RAW,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setRawArray(p *Parser, v []Node) {
    s := new(parseArrayStack)
    s.parser = *p
    s.v = v
    self.t = V_ARRAY_RAW
    self.setCapAndLen(cap(v), len(v))
    self.p = (unsafe.Pointer)(s)
}

func newRawObject(p *Parser, v []Pair) Node {
    s := new(parseObjectStack)
    s.parser = *p
    s.v = v
    return Node{
        t: V_OBJECT_RAW,
        v: int64(len(v)&_LEN_MASK | cap(v)<<_CAP_BITS),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setRawObject(p *Parser, v []Pair) {
    s := new(parseObjectStack)
    s.parser = *p
    s.v = v
    self.t = V_OBJECT_RAW
    self.setCapAndLen(cap(v), len(v))
    self.p = (unsafe.Pointer)(s)
}

func newRawNode(str string) Node {
    return Node{
        t: V_RAW,
        p: str2ptr(str),
        v: int64(len(str)),
    }
}

func (self *Node) parseRaw() Node {
    raw := addr2str(self.p, self.v)
    n, e := NewParser(raw).Parse()
    if e != 0 {
        panic(fmt.Sprintf("%v, raw json: %s", e.Error(), raw))
    }
    return n
}
