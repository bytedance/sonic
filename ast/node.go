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
    `strconv`
    `unsafe`
    `reflect`
    
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _CAP_BITS          = 32
    _LEN_MASK          = 1 << _CAP_BITS - 1

    _NODE_SIZE = unsafe.Sizeof(Node{})
    _PAIR_SIZE = unsafe.Sizeof(Pair{})
)

const (
    _V_NONE         types.ValueType = 0
    _V_NODE_BASE    types.ValueType = 1 << 5
    _V_LAZY         types.ValueType = 1 << 7
    _V_RAW          types.ValueType = 1 << 8
    _V_NUMBER                       = _V_NODE_BASE + 1
    _V_ANY                          = _V_NODE_BASE + 2
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
    V_ANY    = int(_V_ANY)
)

var (
    byteType = rt.UnpackType(reflect.TypeOf(byte(0)))
)

type Node struct {
    l uint32
    c uint32
    t types.ValueType
    p unsafe.Pointer
}

// UnmarshalJSON is just an adapter to json.Unmarshaler.
// If you want better performance, use Searcher.GetByPath() directly
func (self *Node) UnmarshalJSON(data []byte) (err error) {
    *self, err = NewSearcher(string(data)).GetByPath()
    return 
}

/** Node Type Accessor **/

// Type returns json type represented by the node
// It will be one of belows:
//    V_NONE   = 0 (empty node)
//    V_ERROR  = 1 (error node)
//    V_NULL   = 2 (json value `null`)
//    V_TRUE   = 3 (json value `true`)
//    V_FALSE  = 4 (json value `false`)
//    V_ARRAY  = 5 (json value array)
//    V_OBJECT = 6 (json value object)
//    V_STRING = 7 (json value string)
//    V_NUMBER = 33 (json value number )
//    V_ANY    = 34 (golang interface{})
func (self Node) Type() int {
    return int(self.t & _MASK_LAZY & _MASK_RAW)
}

func (self Node) itype() types.ValueType {
    return self.t & _MASK_LAZY & _MASK_RAW
}

// Exists returns false only if the self is nil or empty node V_NONE
func (self *Node) Exists() bool {
    return self != nil && self.t != _V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
    if self == nil {
        return false
    }
    return self.t != V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotFound If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node)  Check() error {
    if self == nil {
        return ErrNotExist
    } else if self.t != V_ERROR {
        return nil
    } else {
        return self
    }
}

// IsRaw returns true if node's underlying value is raw json
func (self Node) IsRaw() bool {
    return self.t&_V_RAW != 0
}

func (self *Node) isLazy() bool {
    return self != nil && self.t&_V_LAZY != 0
}

func (self *Node) isAny() bool {
    return self != nil && self.t == _V_ANY
}

/** Simple Value Methods **/

// Raw returns json representation of the node,
func (self *Node) Raw() (string, error) {
    if !self.IsRaw() {
        buf, err := self.MarshalJSON()
        return rt.Mem2Str(buf), err
    }
    return self.toString(), nil
}

func (self *Node) checkRaw() error {
    if err := self.Check(); err != nil {
        return err
    }
    if self.IsRaw() {
        self.parseRaw(false)
    }
    return nil
}

// Bool returns bool value represented by this node, 
// including types.V_TRUE|V_FALSE|V_NUMBER|V_STRING|V_ANY|V_NULL, 
// V_NONE will return error
func (self *Node) Bool() (bool, error) {
    if err := self.checkRaw(); err != nil {
        return false, err
    }
    switch self.t {
        case types.V_TRUE  : return true , nil
        case types.V_FALSE : return false, nil
        case types.V_NULL  : return false, nil
        case _V_NUMBER     : 
            if i, err := self.toInt64(); err == nil {
                return i != 0, nil
            } else if f, err := self.toFloat64(); err == nil {
                return f != 0, nil
            } else {
                return false, err
            }
        case types.V_STRING: return strconv.ParseBool(self.toString())
        case _V_ANY        :   
            any := self.packAny()     
            switch v := any.(type) {
                case bool   : return v, nil
                case int    : return v != 0, nil
                case int8   : return v != 0, nil
                case int16  : return v != 0, nil
                case int32  : return v != 0, nil
                case int64  : return v != 0, nil
                case uint   : return v != 0, nil
                case uint8  : return v != 0, nil
                case uint16 : return v != 0, nil
                case uint32 : return v != 0, nil
                case uint64 : return v != 0, nil
                case float32: return v != 0, nil
                case float64: return v != 0, nil
                case string : return strconv.ParseBool(v)
                case json.Number: 
                    if i, err := v.Int64(); err == nil {
                        return i != 0, nil
                    } else if f, err := v.Float64(); err == nil {
                        return f != 0, nil
                    } else {
                        return false, err
                    }
                default: return false, ErrUnsupportType
            }
        default            : return false, ErrUnsupportType
    }
}

// Int64 casts the node to int64 value, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING
// V_NONE it will return error
func (self *Node) Int64() (int64, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    switch self.t {
        case _V_NUMBER, types.V_STRING :
            if i, err := self.toInt64(); err == nil {
                return i, nil
            } else if f, err := self.toFloat64(); err == nil {
                return int64(f), nil
            } else {
                return 0, err
            }
        case types.V_TRUE     : return 1, nil
        case types.V_FALSE    : return 0, nil
        case types.V_NULL     : return 0, nil
        case _V_ANY           :  
            any := self.packAny()
            switch v := any.(type) {
                case bool   : if v { return 1, nil } else { return 0, nil }
                case int    : return int64(v), nil
                case int8   : return int64(v), nil
                case int16  : return int64(v), nil
                case int32  : return int64(v), nil
                case int64  : return int64(v), nil
                case uint   : return int64(v), nil
                case uint8  : return int64(v), nil
                case uint16 : return int64(v), nil
                case uint32 : return int64(v), nil
                case uint64 : return int64(v), nil
                case float32: return int64(v), nil
                case float64: return int64(v), nil
                case string : 
                    if i, err := strconv.ParseInt(v, 10, 64); err == nil {
                        return i, nil
                    } else if f, err := strconv.ParseFloat(v, 64); err == nil {
                        return int64(f), nil
                    } else {
                        return 0, err
                    }
                case json.Number: 
                    if i, err := v.Int64(); err == nil {
                        return i, nil
                    } else if f, err := v.Float64(); err == nil {
                        return int64(f), nil
                    } else {
                        return 0, err
                    }
                default: return 0, ErrUnsupportType
            }
        default               : return 0, ErrUnsupportType
    }
}

// StrictInt64 exports underlying int64 value, including V_NUMBER, V_ANY
func (self *Node) StrictInt64() (int64, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    switch self.t {
        case _V_NUMBER        : return self.toInt64()
        case _V_ANY           :  
            any := self.packAny()
            switch v := any.(type) {
                case int   : return int64(v), nil
                case int8  : return int64(v), nil
                case int16 : return int64(v), nil
                case int32 : return int64(v), nil
                case int64 : return int64(v), nil
                case uint  : return int64(v), nil
                case uint8 : return int64(v), nil
                case uint16: return int64(v), nil
                case uint32: return int64(v), nil
                case uint64: return int64(v), nil
                case json.Number: 
                    if i, err := v.Int64(); err == nil {
                        return i, nil
                    } else {
                        return 0, err
                    }
                default: return 0, ErrUnsupportType
            }
        default               : return 0, ErrUnsupportType
    }
}

func castNumber(v bool) json.Number {
    if v {
        return json.Number("1")
    } else {
        return json.Number("0")
    }
}

// Number casts node to float64, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL,
// V_NONE it will return error
func (self *Node) Number() (json.Number, error) {
    if err := self.checkRaw(); err != nil {
        return json.Number(""), err
    }
    switch self.t {
        case _V_NUMBER        : return self.toNumber(), nil
        case types.V_STRING : 
            if _, err := self.toInt64(); err == nil {
                return self.toNumber(), nil
            } else if _, err := self.toFloat64(); err == nil {
                return self.toNumber(), nil
            } else {
                return json.Number(""), err
            }
        case types.V_TRUE     : return json.Number("1"), nil
        case types.V_FALSE    : return json.Number("0"), nil
        case types.V_NULL     : return json.Number("0"), nil
        case _V_ANY           :        
            any := self.packAny()
            switch v := any.(type) {
                case bool   : return castNumber(v), nil
                case int    : return castNumber(v != 0), nil
                case int8   : return castNumber(v != 0), nil
                case int16  : return castNumber(v != 0), nil
                case int32  : return castNumber(v != 0), nil
                case int64  : return castNumber(v != 0), nil
                case uint   : return castNumber(v != 0), nil
                case uint8  : return castNumber(v != 0), nil
                case uint16 : return castNumber(v != 0), nil
                case uint32 : return castNumber(v != 0), nil
                case uint64 : return castNumber(v != 0), nil
                case float32: return castNumber(v != 0), nil
                case float64: return castNumber(v != 0), nil
                case string : 
                    if _, err := strconv.ParseFloat(v, 64); err == nil {
                        return json.Number(v), nil
                    } else {
                        return json.Number(""), err
                    }
                case json.Number: return v, nil
                default: return json.Number(""), ErrUnsupportType
            }
        default               : return json.Number(""), ErrUnsupportType
    }
}

// Number exports underlying float64 value, including V_NUMBER, V_ANY of json.Number
func (self *Node) StrictNumber() (json.Number, error) {
    if err := self.checkRaw(); err != nil {
        return json.Number(""), err
    }
    switch self.t {
        case _V_NUMBER        : return self.toNumber()  , nil
        case _V_ANY        :        
            if v, ok := self.packAny().(json.Number); ok {
                return v, nil
            } else {
                return json.Number(""), ErrUnsupportType
            }
        default               : return json.Number(""), ErrUnsupportType
    }
}

// String cast node to string, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL,
// V_NONE it will return error
func (self *Node) String() (string, error) {
    if err := self.checkRaw(); err != nil {
        return "", err
    }
    switch self.t {
        case types.V_NULL    : return "" , nil
        case types.V_TRUE    : return "true" , nil
        case types.V_FALSE   : return "false", nil
        case types.V_STRING, _V_NUMBER  : return self.toString(), nil
        case _V_ANY          :        
        any := self.packAny()
        switch v := any.(type) {
            case bool   : return strconv.FormatBool(v), nil
            case int    : return strconv.Itoa(v), nil
            case int8   : return strconv.Itoa(int(v)), nil
            case int16  : return strconv.Itoa(int(v)), nil
            case int32  : return strconv.Itoa(int(v)), nil
            case int64  : return strconv.Itoa(int(v)), nil
            case uint   : return strconv.Itoa(int(v)), nil
            case uint8  : return strconv.Itoa(int(v)), nil
            case uint16 : return strconv.Itoa(int(v)), nil
            case uint32 : return strconv.Itoa(int(v)), nil
            case uint64 : return strconv.Itoa(int(v)), nil
            case float32: return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
            case float64: return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
            case string : return v, nil 
            case json.Number: return v.String(), nil
            default: return "", ErrUnsupportType
        }
        default              : return ""     , ErrUnsupportType
    }
}

// StrictString returns string value (unescaped), includeing V_STRING, V_ANY of string.
// In other cases, it will return empty string.
func (self *Node) StrictString() (string, error) {
    if err := self.checkRaw(); err != nil {
        return "", err
    }
    switch self.t {
        case types.V_STRING  : return self.toString(), nil
        case _V_ANY          :        
            if v, ok := self.packAny().(string); ok {
                return v, nil
            } else {
                return "", ErrUnsupportType
            }
        default              : return "", ErrUnsupportType
    }
}

// Float64 cast node to float64, 
// including V_NUMBER|V_TRUE|V_FALSE|V_ANY|V_STRING|V_NULL,
// V_NONE it will return error
func (self *Node) Float64() (float64, error) {
    if err := self.checkRaw(); err != nil {
        return 0.0, err
    }
    switch self.t {
        case _V_NUMBER, types.V_STRING : return self.toFloat64()
        case types.V_TRUE    : return 1.0, nil
        case types.V_FALSE   : return 0.0, nil
        case types.V_NULL    : return 0.0, nil
        case _V_ANY          :        
            any := self.packAny()
            switch v := any.(type) {
                case bool    : 
                    if v {
                        return 1.0, nil
                    } else {
                        return 0.0, nil
                    }
                case int    : return float64(v), nil
                case int8   : return float64(v), nil
                case int16  : return float64(v), nil
                case int32  : return float64(v), nil
                case int64  : return float64(v), nil
                case uint   : return float64(v), nil
                case uint8  : return float64(v), nil
                case uint16 : return float64(v), nil
                case uint32 : return float64(v), nil
                case uint64 : return float64(v), nil
                case float32: return float64(v), nil
                case float64: return float64(v), nil
                case string : 
                    if f, err := strconv.ParseFloat(v, 64); err == nil {
                        return float64(f), nil
                    } else {
                        return 0, err
                    }
                case json.Number: 
                    if f, err := v.Float64(); err == nil {
                        return float64(f), nil
                    } else {
                        return 0, err
                    }
                default     : return 0, ErrUnsupportType
            }
        default             : return 0.0, ErrUnsupportType
    }
}

// Float64 exports underlying float64 value, includeing V_NUMBER, V_ANY 
func (self *Node) StrictFloat64() (float64, error) {
    if err := self.checkRaw(); err != nil {
        return 0.0, err
    }
    switch self.t {
        case _V_NUMBER       : return self.toFloat64()
        case _V_ANY        :        
            any := self.packAny()
            switch v := any.(type) {
                case float32 : return float64(v), nil
                case float64 : return float64(v), nil
                default      : return 0, ErrUnsupportType
            }
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
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY || self.t == types.V_STRING {
        return int(self.l), nil
    } else if self.t == _V_NONE || self.t == types.V_NULL {
        return 0, nil
    } else {
        return 0, ErrUnsupportType
    }
}

func (self Node) len() int {
    return int(self.l)
}

// Cap returns malloc capacity of a array|object node for children
func (self *Node) Cap() (int, error) {
    if err := self.checkRaw(); err != nil {
        return 0, err
    }
    if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
        return int(self.c), nil
    } else if self.t == _V_NONE || self.t == types.V_NULL {
        return 0, nil
    } else {
        return 0, ErrUnsupportType
    }
}

func (self Node) cap() int {
    return int(self.c)
}

// Set sets the node of given key under self, and reports if the key has existed.
//
// If self is V_NONE or V_NULL, it becomes V_OBJECT and sets the node at the key.
func (self *Node) Set(key string, node Node) (bool, error) {
    if self != nil && (self.t == _V_NONE || self.t == types.V_NULL) {
        *self = NewObject([]Pair{{key, node}})
        return false, nil
    }

    if err := node.Check(); err != nil {
        return false, err 
    }

    p := self.Get(key)

    if !p.Exists() {
        // self must be fully-loaded here
        if self.len() == 0 {
            *self = newObject(new(linkedPairs), 0)
        }
        s := (*linkedPairs)(self.p)
        s.Add(Pair{key, node})
        self.l++
        return false, nil

    } else if err := p.Check(); err != nil {
        return false, err
    } 

    *p = node
    return true, nil
}

// SetAny wraps val with V_ANY node, and Set() the node.
func (self *Node) SetAny(key string, val interface{}) (bool, error) {
    return self.Set(key, NewAny(val))
}

// Unset remove the node of given key under object parent, and reports if the key has existed.
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

// SetByIndex sets the node of given index, and reports if the key has existed.
//
// The index must be within self's children.
func (self *Node) SetByIndex(index int, node Node) (bool, error) {
    if err := node.Check(); err != nil {
        return false, err 
    }

    p := self.Index(index)
    if !p.Exists() {
        return false, ErrNotExist
    } else if err := p.Check(); err != nil {
        return false, err
    }

    *p = node
    return true, nil
}

// SetAny wraps val with V_ANY node, and SetByIndex() the node.
func (self *Node) SetAnyByIndex(index int, val interface{}) (bool, error) {
    return self.SetByIndex(index, NewAny(val))
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
    } else {
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

// Add appends the given node under self.
//
// If self is V_NONE or V_NULL, it becomes V_ARRAY and sets the node at index 0.
func (self *Node) Add(node Node) error {
    if self != nil && (self.t == _V_NONE || self.t == types.V_NULL) {
        *self = NewArray([]Node{node})
        return nil
    }

    s, err := self.unsafeArray()
    if err != nil {
        return err
    }

    s.Add(node)
    self.l++
    return nil
}


// SetAny wraps val with V_ANY node, and Add() the node.
func (self *Node) AddAny(val interface{}) error {
    return self.Add(NewAny(val))
}

// GetByPath load given path on demands,
// which only ensure nodes before this path got parsed.
//
// Note, the api expects the json is well-formed at least,
// otherwise it may return unexpected result.
func (self *Node) GetByPath(path ...interface{}) *Node {
    if !self.Valid() {
        return self
    }
    var s = self
    for _, p := range path {
        switch p := p.(type) {
        case int:
            s = s.Index(p)
            if !s.Valid() {
                return s
            }
        case string:
            s = s.Get(p)
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

// Index indexies node at given idx,
// node type CAN be either V_OBJECT or V_ARRAY
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
           return newError(_ERR_NOT_FOUND, "value not exists")
        }
        return &pr.Value

    } else {
        return newError(_ERR_UNSUPPORT_TYPE, fmt.Sprintf("unsupported type: %v", self.itype()))
    }
}

// IndexPair indexies pair at given idx,
// node type MUST be either V_OBJECT
func (self *Node) IndexPair(idx int) *Pair {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil
    }
    return self.skipIndexPair(idx)
}

// IndexOrGet firstly use idx to index a value and check if its key matches
// If not, then use the key to search value
func (self *Node) IndexOrGet(idx int, key string) *Node {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return unwrapError(err)
    }

    pr := self.skipIndexPair(idx)
    if pr != nil && pr.Key == key {
        return &pr.Value
    }
    n, _ := self.skipKey(key)
    return n
}

/** Generic Value Converters **/

// Map loads all keys of an object node
func (self *Node) Map() (map[string]interface{}, error) {
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.(map[string]interface{}); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
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
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.(map[string]interface{}); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
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
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.(map[string]Node); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
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
    return self.toGenericObjectUsePair()
}

func (self *Node) unsafeMap() (*linkedPairs, error) {
    if err := self.should(types.V_OBJECT, "an object"); err != nil {
        return nil, err
    }
    if err := self.skipAllKey(); err != nil {
        return nil, err
    }
    if self.p == nil {
        *self = newObject(new(linkedPairs), 0)
    }
    return (*linkedPairs)(self.p), nil
}

// SortKeys sorts children of a V_OBJECT node in ascending key-order.
// If recurse is true, it recursively sorts children's children as long as a V_OBJECT node is found.
func (self *Node) SortKeys(recurse bool) (err error) {
    ps, err := self.unsafeMap()
    if err != nil {
        return err
    }
    ps.Sort()
    if recurse {
        var sc Scanner
        sc = func(path Sequence, node *Node) bool {
            if node.itype() == types.V_OBJECT {
                if err := node.SortKeys(recurse); err != nil {
                    return false
                }
            }
            if node.itype() == types.V_ARRAY {
                if err := node.ForEach(sc); err != nil {
                    return false
                }
            }
            return true
        }
        self.ForEach(sc)
    }
    return nil
}

// Array loads all indexes of an array node
func (self *Node) Array() ([]interface{}, error) {
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.([]interface{}); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
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
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.([]interface{}); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
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
    if self.isAny() {
        any := self.packAny()
        if v, ok := any.([]Node); ok {
            return v, nil
        } else {
            return nil, ErrUnsupportType
        }
    }
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.skipAllIndex(); err != nil {
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
    if err := self.skipAllIndex(); err != nil {
        return nil, err
    }
    return self.toGenericArrayUseNode()
}

func (self *Node) unsafeArray() (*linkedNodes, error) {
    if err := self.should(types.V_ARRAY, "an array"); err != nil {
        return nil, err
    }
    if err := self.skipAllIndex(); err != nil {
        return nil, err
    }
    if self.p == nil {
        *self = newArray(new(linkedNodes), 0)
    }
    return (*linkedNodes)(self.p), nil
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
        case types.V_STRING  : return self.toString(), nil
        case _V_NUMBER       : 
            v, err := self.toFloat64()
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
        case _V_ANY:
            switch v := self.packAny().(type) {
                case Node : return v.Interface()
                case *Node: return v.Interface()
                default   : return v, nil
            }
        default              : return nil,  ErrUnsupportType
    }
}

func (self *Node) packAny() interface{} {
    return *(*interface{})(self.p)
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
        case types.V_STRING  : return self.toString(), nil
        case _V_NUMBER       : return self.toNumber(), nil
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
        case _V_ANY          : return self.packAny(), nil
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
            if err := self.skipAllKey(); err != nil {
                return nil, err
            }
            return self.toGenericObjectUseNode()
        default              : return *self, self.Check()
    }
}

// LoadAll loads all the node's children and children's children as parsed.
// After calling it, the node can be safely used on concurrency
func (self *Node) LoadAll() error {
    if self.IsRaw() {
        self.parseRaw(true)
        return self.Check()
    }

    switch self.itype() {
    case types.V_ARRAY:
        e := self.len()
        if err := self.loadAllIndex(); err != nil {
            return err
        }
        for i := 0; i < e; i++ {
            n := self.nodeAt(i)
            if n.IsRaw() {
                n.parseRaw(true)
            }
            if err := n.Check(); err != nil {
                return err
            }
        }
        return nil
    case types.V_OBJECT:
        e := self.len()
        if err := self.loadAllKey(); err != nil {
            return err
        }
        for i := 0; i < e; i++ {
            n := self.pairAt(i)
            if n.Value.IsRaw() {
                n.Value.parseRaw(true)
            }
            if err := n.Value.Check(); err != nil {
                return err
            }
        }
        return nil
    default:
        return self.Check()
    }
}

// Load loads the node's children as parsed.
// After calling it, only the node itself can be used on concurrency (not include its children)
func (self *Node) Load() error {
    if self.IsRaw() {
        self.parseRaw(false)
        return self.Load()
    }

    switch self.t {
    case _V_ARRAY_LAZY:
        return self.skipAllIndex()
    case _V_OBJECT_LAZY:
        return self.skipAllKey()
    default:
        return self.Check()
    }
}

/**---------------------------------- Internal Helper Methods ----------------------------------**/

var (
    _NODE_TYPE = rt.UnpackEface(Node{}).Type
    _PAIR_TYPE = rt.UnpackEface(Pair{}).Type
)

// func (self *Node) setCapAndLen(cap int, len int) {
//     if self.t == types.V_ARRAY || self.t == types.V_OBJECT || self.t == _V_ARRAY_LAZY || self.t == _V_OBJECT_LAZY {
//         self.l = uint32(len)
//         self.c = uint32(cap)
//     } else {
//         panic("value does not have a length")
//     }
// }

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
    var p *linkedNodes
    if self.isLazy() {
        _, stack := self.getParserAndArrayStack()
        p = stack.v
    } else {
        p = (*linkedNodes)(self.p)
    }
    return p.At(i)
}

func (self *Node) pairAt(i int) *Pair {
    var p *linkedPairs
    if self.isLazy() {
        _, stack := self.getParserAndObjectStack()
        p = stack.v
    } else {
        p = (*linkedPairs)(self.p)
    }
    return p.At(i)
}

func (self *Node) getParserAndArrayStack() (*Parser, *parseArrayStack) {
    stack := (*parseArrayStack)(self.p)
    return &stack.parser, stack
}

func (self *Node) getParserAndObjectStack() (*Parser, *parseObjectStack) {
    stack := (*parseObjectStack)(self.p)
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
    *self, err = parser.decodeArray(stack.v, self.len())
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
    *self, err = parser.decodeObject(stack.v, self.len())
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) skipKey(key string) (*Node, int) {
    nb := self.len()
    lazy := self.isLazy()

    if nb > 0 {
        /* linear search */
        var p *Pair
        var i int
        if lazy {
            s := (*parseObjectStack)(self.p)
            p, i = s.v.Get(key)
        } else {
            p, i = (*linkedPairs)(self.p).Get(key)
        }

        if p != nil {
            return &p.Value, i
        }
    }

    /* not found */
    if !lazy {
        return nil, -1
    }

    // lazy load
    for last, i := self.skipNextPair(), nb; last != nil; last, i = self.skipNextPair(), i+1 {
        if last.Value.Check() != nil {
            return &last.Value, -1
        }
        if last.Key == key {
            return &last.Value, i
        }
    }

    return nil, -1
}

func (self *Node) skipIndex(index int) *Node {
    nb := self.len()
    if nb > index {
        v := self.nodeAt(index)
        return v
    }
    if !self.isLazy() {
        return nil
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

    return nil
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
    *self, err = parser.decodeArray(stack.v, self.len())
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
    *self, err = parser.decodeObject(stack.v, self.len())
    if err != 0 {
        return parser.ExportError(err)
    }
    return nil
}

func (self *Node) removeNode(i int) {
    node := self.nodeAt(i)
    if node == nil {
        return
    }
    *node = Node{}
    // NOTICE: for consistency with linkedNodes, we DOSEN'T reduce size here
    // self.l--
}

func (self *Node) removePair(i int) {
    last := self.pairAt(i)
    if last == nil {
        return
    }
    *last = Pair{}
    // NOTICE: for consistency with linkedNodes, we DOSEN'T reduce size here
    // self.l--
}

func (self *Node) toGenericArray() ([]interface{}, error) {
    nb := self.len()
    ret := make([]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var s = (*linkedNodes)(self.p)
    for i := 0; i < nb; i++ {
        p := s.At(i)
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
    var s = (*linkedNodes)(self.p)
    for i := 0; i < nb; i++ {
        p := s.At(i)
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
    if nb == 0 {
        return []Node{}, nil
    }

    var s = (*linkedNodes)(self.p)
    var out = make([]Node, s.Len())
    s.ToSlice(out)

    return out, nil
}

func (self *Node) toGenericObject() (map[string]interface{}, error) {
    nb := self.len()
    ret := make(map[string]interface{}, nb)
    if nb == 0 {
        return ret, nil
    }

    /* convert each item */
    var s = (*linkedPairs)(self.p)
    for i := 0; i < nb; i++ {
        p := s.At(i)
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
    var s = (*linkedPairs)(self.p)
    for i := 0; i < nb; i++ {
        p := s.At(i)
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
    if nb == 0 {
        return map[string]Node{}, nil
    }

    var s = (*linkedPairs)(self.p)
    var out = make(map[string]Node, s.Len())
    s.ToMap(out)

    /* all done */
    return out, nil
}

func (self *Node) toGenericObjectUsePair() ([]Pair, error) {
    var nb = self.len()
    if nb == 0 {
        return []Pair{}, nil
    }

    var s = (*linkedPairs)(self.p)
    var out = make([]Pair, s.Len())
    s.ToSlice(out)

    /* all done */
    return out, nil
}

/**------------------------------------ Factory Methods ------------------------------------**/

var (
    nullNode  = Node{t: types.V_NULL}
    trueNode  = Node{t: types.V_TRUE}
    falseNode = Node{t: types.V_FALSE}
)

// NewRaw creates a node of raw json.
// If the input json is invalid, NewRaw returns a error Node.
func NewRaw(json string) Node {
    parser := NewParser(json)
    start, err := parser.skip()
    if err != 0 {
        return *newError(err, err.Message()) 
    }
    it := switchRawType(parser.s[start])
    if it == _V_NONE {
        return Node{}
    }
    return newRawNode(parser.s[start:parser.p], it)
}

// NewAny creates a node of type V_ANY if any's type isn't Node or *Node, 
// which stores interface{} and can be only used for `.Interface()`\`.MarshalJSON()`.
func NewAny(any interface{}) Node {
    switch n := any.(type) {
    case Node:
        return n
    case *Node:
        return *n
    default:
        return Node{
            t: _V_ANY,
            p: unsafe.Pointer(&any),
        }
    }
}

// NewBytes encodes given src with Base64 (RFC 4648), and creates a node of type V_STRING.
func NewBytes(src []byte) Node {
    if len(src) == 0 {
        panic("empty src bytes")
    }
    out := encodeBase64(src)
    return NewString(out)
}

// NewNull creates a node of type V_NULL
func NewNull() Node {
    return Node{
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
        p: nil,
        t: t,
    }
}

// NewNumber creates a json.Number node
// v must be a decimal string complying with RFC8259
func NewNumber(v string) Node {
    return Node{
        l: uint32(len(v)),
        p: rt.StrPtr(v),
        t: _V_NUMBER,
    }
}

func (node Node) toNumber() json.Number {
    return json.Number(rt.StrFrom(node.p, int64(node.l)))
}

func (self Node) toString() string {
    return rt.StrFrom(self.p, int64(self.l))
}

func (node Node) toFloat64() (float64, error) {
    ret, err := node.toNumber().Float64()
    if err != nil {
        return 0, err
    }
    return ret, nil
}

func (node Node) toInt64() (int64, error) {
    ret,err := node.toNumber().Int64()
    if err != nil {
        return 0, err
    }
    return ret, nil
}

func newBytes(v []byte) Node {
    return Node{
        t: types.V_STRING,
        p: mem2ptr(v),
        l: uint32(len(v)),
    }
}

// NewString creates a node of type V_STRING. 
// v is considered to be a valid UTF-8 string,
// which means it won't be validated and unescaped.
// when the node is encoded to json, v will be escaped.
func NewString(v string) Node {
    return Node{
        t: types.V_STRING,
        p: rt.StrPtr(v),
        l: uint32(len(v)),
    }
}

// NewArray creates a node of type V_ARRAY,
// using v as its underlying children
func NewArray(v []Node) Node {
    s := new(linkedNodes)
    s.FromSlice(v)
    return newArray(s, s.Len())
}

func newArray(v *linkedNodes, l int) Node {
    return Node{
        t: types.V_ARRAY,
        l: uint32(l),
        c: uint32(v.Cap()),
        p: unsafe.Pointer(v),
    }
}

func (self *Node) setArray(v *linkedNodes) {
    self.t = types.V_ARRAY
    self.p = unsafe.Pointer(v)
}

// NewObject creates a node of type V_OBJECT,
// using v as its underlying children
func NewObject(v []Pair) Node {
    s := new(linkedPairs)
    s.FromSlice(v)
    return newObject(s, s.Len())
}

func newObject(v *linkedPairs, l int) Node {
    return Node{
        t: types.V_OBJECT,
        l: uint32(l),
        c: uint32(v.Cap()),
        p: unsafe.Pointer(v),
    }
}

func (self *Node) setObject(v *linkedPairs) {
    self.t = types.V_OBJECT
    self.p = unsafe.Pointer(v)
}

type parseObjectStack struct {
    parser Parser
    v      *linkedPairs
}

type parseArrayStack struct {
    parser Parser
    v      *linkedNodes
}

func newLazyArray(p *Parser, v *linkedNodes) Node {
    s := new(parseArrayStack)
    s.parser = *p
    s.v = v
    return Node{
        t: _V_ARRAY_LAZY,
        l: uint32(v.Len()),
        c: uint32(v.Cap()),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setLazyArray(p *Parser, v *linkedNodes, s *parseArrayStack) {
    s.parser = *p
    s.v = v
    self.t = _V_ARRAY_LAZY
    self.p = (unsafe.Pointer)(s)
}

func newLazyObject(p *Parser, v *linkedPairs) Node {
    s := new(parseObjectStack)
    s.parser = *p
    s.v = v
    return Node{
        t: _V_OBJECT_LAZY,
        l: uint32(v.Len()),
        c: uint32(v.Cap()),
        p: unsafe.Pointer(s),
    }
}

func (self *Node) setLazyObject(p *Parser, v *linkedPairs, s *parseObjectStack) {
    s.parser = *p
    s.v = v
    self.t = _V_OBJECT_LAZY
    self.p = (unsafe.Pointer)(s)
}

func newRawNode(str string, typ types.ValueType) Node {
    return Node{
        t: _V_RAW | typ,
        p: rt.StrPtr(str),
        l: uint32(len(str)),
    }
}

func (self *Node) parseRaw(full bool) {
    raw := self.toString()
    parser := NewParser(raw)
    if full {
        parser.noLazy = true
        parser.skipValue = false
    }
    var e types.ParsingError
    *self, e = parser.Parse()
    if e != 0 {
        *self = *newSyntaxError(parser.syntaxError(e))
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