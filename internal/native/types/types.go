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

package types

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

type ValueType int
type ParsingError uint
type SearchingError uint

// NOTE: !NOT MODIFIED ONLY.
// This definitions are followed in native/types.h.

const (
    V_EOF     ValueType = 1
    V_NULL    ValueType = 2
    V_TRUE    ValueType = 3
    V_FALSE   ValueType = 4
    V_ARRAY   ValueType = 5
    V_OBJECT  ValueType = 6
    V_STRING  ValueType = 7
    V_DOUBLE  ValueType = 8
    V_INTEGER ValueType = 9
    _         ValueType = 10    // V_KEY_SEP
    _         ValueType = 11    // V_ELEM_SEP
    _         ValueType = 12    // V_ARRAY_END
    _         ValueType = 13    // V_OBJECT_END
    V_MAX
    V_NUMBER ValueType = 33
)

const (
    // for native.Unquote() flags
    B_DOUBLE_UNQUOTE  = 0
    B_UNICODE_REPLACE = 1

    // for native.Value() flags
    B_USE_NUMBER      = 1
    B_VALIDATE_STRING = 5
    B_ALLOW_CONTROL   = 31
)

const (
    F_DOUBLE_UNQUOTE  = 1 << B_DOUBLE_UNQUOTE
    F_UNICODE_REPLACE = 1 << B_UNICODE_REPLACE

    F_USE_NUMBER      = 1 << B_USE_NUMBER
    F_VALIDATE_STRING = 1 << B_VALIDATE_STRING
    F_ALLOW_CONTROL   = 1 << B_ALLOW_CONTROL
)

const (
    MAX_RECURSE = 4096
)

const (
    SPACE_MASK = (1 << ' ') | (1 << '\t') | (1 << '\r') | (1 << '\n')
)

const (
    ERR_EOF                ParsingError = 1
    ERR_INVALID_CHAR       ParsingError = 2
    ERR_INVALID_ESCAPE     ParsingError = 3
    ERR_INVALID_UNICODE    ParsingError = 4
    ERR_INTEGER_OVERFLOW   ParsingError = 5
    ERR_INVALID_NUMBER_FMT ParsingError = 6
    ERR_RECURSE_EXCEED_MAX ParsingError = 7
    ERR_FLOAT_INFINITY     ParsingError = 8
    ERR_MISMATCH           ParsingError = 9
    ERR_INVALID_UTF8       ParsingError = 10

    // error code used in ast
    ERR_NOT_FOUND          ParsingError = 33
    ERR_UNSUPPORT_TYPE     ParsingError = 34
)

var _ParsingErrors = []string{
    0                      : "ok",
    ERR_EOF                : "eof",
    ERR_INVALID_CHAR       : "invalid char",
    ERR_INVALID_ESCAPE     : "invalid escape char",
    ERR_INVALID_UNICODE    : "invalid unicode escape",
    ERR_INTEGER_OVERFLOW   : "integer overflow",
    ERR_INVALID_NUMBER_FMT : "invalid number format",
    ERR_RECURSE_EXCEED_MAX : "recursion exceeded max depth",
    ERR_FLOAT_INFINITY     : "float number is infinity",
    ERR_MISMATCH           : "mismatched type with value",
    ERR_INVALID_UTF8       : "invalid UTF8",

	ERR_NOT_FOUND          : "not found",
	ERR_UNSUPPORT_TYPE     : "unsupported type in path",
}

func (self ParsingError) Error() string {
    return "json: error when parsing input: " + self.Message()
}

func (self ParsingError) Message() string {
    if int(self) < len(_ParsingErrors) {
        return _ParsingErrors[self]
    } else {
        return fmt.Sprintf("unknown error %d", self)
    }
}

type JsonState struct {
    Vt ValueType
    Dv   float64
    Iv   int64
    Ep   int
    Dbuf *byte
    Dcap int
}

type StateMachine struct {
    Sp int
    Vt [MAX_RECURSE]int
}

var stackPool = sync.Pool{
    New: func()interface{}{
        return &StateMachine{}
    },
}

func NewStateMachine() *StateMachine {
    return stackPool.Get().(*StateMachine)
}

func FreeStateMachine(fsm *StateMachine) {
    stackPool.Put(fsm)
}

const MaxDigitNums = 800

var digitPool = sync.Pool{
    New: func() interface{} {
        return (*byte)(unsafe.Pointer(&[MaxDigitNums]byte{}))
    },
}

func NewDbuf() *byte {
    return digitPool.Get().(*byte)
}

func FreeDbuf(p *byte) {
    digitPool.Put(p)
}

type Flag uint16

const (
	F_ESC	= Flag(1<<0)
)

type Type uint8 

const (
    T_NULL    Type = 2
    T_TRUE    Type = 3
    T_FALSE   Type = 4
    T_ARRAY   Type = 5
    T_OBJECT  Type = 6
    T_STRING  Type = 7
    T_NUMBER  Type = 8
)

type Token struct {
	Kind Type
	Flag Flag
	Off uint32
	Len uint32
}

func (t Type) IsComplex() bool {
    return t == T_ARRAY || t == T_OBJECT
}

func (t Type) String() string {
    switch t {
    case T_NULL:
        return "null"
    case T_TRUE:
        return "true"
    case T_FALSE:
        return "false"
    case T_ARRAY:
        return "array"
    case T_OBJECT:
        return "object"
    case T_STRING:
        return "string"
    case T_NUMBER:
        return "number"
    default:
        return strconv.Itoa(int(t))
    }
}

type Node struct {
	Kind Type
	Flag Flag
	JSON string
	Kids []Token
}

type stats struct {
    min int64
    max int64
    last int64
    over int64
}

const _DefaultTokenSize = 8 

var tokenSizeStats = stats{
    min: _DefaultTokenSize,
    max: _DefaultTokenSize,
    last: _DefaultTokenSize,
}

func avg(min, max, last int64) int64 {
    return (last + (max + min)/2)
}

// TODO
func PredictTokenSize() int64 {
    return avg((atomic.LoadInt64(&tokenSizeStats.min)), (atomic.LoadInt64(&tokenSizeStats.max)), (atomic.LoadInt64(&tokenSizeStats.last)))
}

// TODO
func RecordTokenSize(last int64) {
    min, max, lastSize, over := (atomic.LoadInt64(&tokenSizeStats.min)), (atomic.LoadInt64(&tokenSizeStats.max)), (atomic.LoadInt64(&tokenSizeStats.last)), (atomic.LoadInt64(&tokenSizeStats.last))

    avg := avg(min, max, lastSize)

    if last > max || last > avg {
        // fast enlarge 2x
        max = last
        over = 0
    } else if last < avg {
        over++
        if over > 5 && last < max {
            // slow shrink 1/4
            max -= (max - last) >> 1
            over = 0
        }
    }

    if last < min {
        min = last
    }
    
    lastSize = last
	atomic.StoreInt64(&tokenSizeStats.min, min)
	atomic.StoreInt64(&tokenSizeStats.max, max)
	atomic.StoreInt64(&tokenSizeStats.last, lastSize)
	atomic.StoreInt64(&tokenSizeStats.over, over)
}

var tokenPool = sync.Pool{
    New: func() interface{} {
        return make([]Token, 0, _DefaultTokenSize)
    },
}

func NewToken() []Token {
    return tokenPool.Get().([]Token)
}

func FreeToken(t []Token) {
    t = (t)[:0]
    tokenPool.Put(t)
}

func (n *Node) Grow()  {
    n.Kids = make([]Token, 0, 2 * cap(n.Kids))
}

// encoding 64-bit Token as follows:
const (
	/* specific error code */
	MUST_RETRY = 0x12345
)

// func (t Flag) IsRaw() bool {
// 	return t & _F_RAW != 0
// }

func (t Flag) IsEsc() bool {
	return t & F_ESC != 0
}

func (t Token) Peek(json string) byte {
	return json[t.Off]
}

func (t Token) Raw(json string) string {
	return json[t.Off:t.Off + t.Len]
}

// for T_OBJECT | T_ARRAY, must remember to handle Kids
func NewNode(json string, start int, flag Flag) Node {
    kind := typeJumpTable[json[start]]
	// if kind == T_OBJECT || kind == T_ARRAY {
    //     flag |= _F_RAW
	// } 
    return Node{
        Kind: kind,
        Flag: flag,
        JSON: json,
    }
}

var typeJumpTable = [256]Type{
    '"' : T_STRING,
    '-' : T_NUMBER,
    '0' : T_NUMBER,
    '1' : T_NUMBER,
    '2' : T_NUMBER,
    '3' : T_NUMBER,
    '4' : T_NUMBER,
    '5' : T_NUMBER,
    '6' : T_NUMBER,
    '7' : T_NUMBER,
    '8' : T_NUMBER,
    '9' : T_NUMBER,
    '[' : T_ARRAY,
    'f' : T_FALSE,
    'n' : T_NULL,
    't' : T_TRUE,
    '{' : T_OBJECT,
}

const (
    F_GetByPath_StartAtKey uint64 = 1 << 1
)