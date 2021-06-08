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
	`sync`
	`unsafe`

	`github.com/bytedance/sonic/internal/native`
	`github.com/bytedance/sonic/internal/rt`
)

const _DEFAULT_NODE_CAP = 16

type Parser struct {
    p      int
    s      string
    noLazy bool
}

var stackPool = sync.Pool{
    New: func()interface{}{
        return &native.StateMachine{}
    },
}

/** Parser Private Methods **/

func (self *Parser) delim() native.ParsingError {
    n := len(self.s)
    p := self.lspace(self.p)

    /* check for EOF */
    if p >= n {
        return native.ERR_EOF
    }

    /* check for the delimtier */
    if self.s[p] != ':' {
        return native.ERR_INVALID_CHAR
    }

    /* update the read pointer */
    self.p = p + 1
    return 0
}

func (self *Parser) object() native.ParsingError {
    n := len(self.s)
    p := self.lspace(self.p)

    /* check for EOF */
    if p >= n {
        return native.ERR_EOF
    }

    /* check for the delimtier */
    if self.s[p] != '{' {
        return native.ERR_INVALID_CHAR
    }

    /* update the read pointer */
    self.p = p + 1
    return 0
}

func (self *Parser) array() native.ParsingError {
    n := len(self.s)
    p := self.lspace(self.p)

    /* check for EOF */
    if p >= n {
        return native.ERR_EOF
    }

    /* check for the delimtier */
    if self.s[p] != '[' {
        return native.ERR_INVALID_CHAR
    }

    /* update the read pointer */
    self.p = p + 1
    return 0
}

func (self *Parser) lspace(sp int) int {
    sv := (*rt.GoString)(unsafe.Pointer(&self.s))
    return native.Lspace(sv.Ptr, sv.Len, sp)
}

func (self *Parser) decodeValue() (val native.JsonState) {
    sv := (*rt.GoString)(unsafe.Pointer(&self.s))
    self.p = native.Value(sv.Ptr, sv.Len, self.p, &val)
    return
}

func (self *Parser) decodeArray(ret []Node) (Node, native.ParsingError) {
    sp := self.p
    ns := len(self.s)

    /* check for EOF */
    if self.p = self.lspace(sp); self.p >= ns {
        return Node{}, native.ERR_EOF
    }

    /* check for empty array */
    if self.s[self.p] == ']' {
        self.p++
        return emptyArrayNode, 0
    }

    /* allocate array space and parse every element */
    for {
        var val Node
        var err native.ParsingError

        /* decode the value */
        if val, err = self.Parse(); err != 0 {
            return Node{}, err
        }

        /* add the value to result */
        ret = append(ret, val)
        self.p = self.lspace(self.p)

        /* check for EOF */
        if self.p >= ns {
            return Node{}, native.ERR_EOF
        }

        /* check for the next character */
        switch self.s[self.p] {
            case ',' : self.p++
            case ']' : self.p++; return newArray(ret), 0
        default:
            if val.IsRaw() {
                return newRawArray(self, ret), 0
            }
            return Node{}, native.ERR_INVALID_CHAR
        }
    }
}

func (self *Parser) decodeObject(ret []Pair) (Node, native.ParsingError) {
    sp := self.p
    ns := len(self.s)

    /* check for EOF */
    if self.p = self.lspace(sp); self.p >= ns {
        return Node{}, native.ERR_EOF
    }

    /* check for empty object */
    if self.s[self.p] == '}' {
        self.p++
        return emptyObjectNode, 0
    }

    /* decode each pair */
    for {
        var val Node
        var njs native.JsonState
        var err native.ParsingError

        /* decode the key */
        if njs = self.decodeValue(); njs.Vt != native.V_STRING {
            return Node{}, native.ERR_INVALID_CHAR
        }

        /* extract the key */
        idx := self.p - 1
        key := self.s[njs.Iv:idx]

        /* check for escape sequence */
        if njs.Ep != -1 {
            if key, err = UnquoteString(key); err != 0 {
                return Node{}, err
            }
        }

        /* expect a ':' delimiter */
        if err = self.delim(); err != 0 {
            return Node{}, err
        }

        /* decode the value */
        if val, err = self.Parse(); err != 0 {
            return Node{}, err
        }

        /* add the value to result */
        ret = append(ret, Pair{Key: key, Value: val})
        self.p = self.lspace(self.p)

        /* check for EOF */
        if self.p >= ns {
            return Node{}, native.ERR_EOF
        }

        /* check for the next character */
        switch self.s[self.p] {
            case ',' : self.p++
            case '}' : self.p++; return newObject(ret), 0
        default:
            if val.IsRaw() {
                return newRawObject(self, ret), 0
            }
            return Node{}, native.ERR_INVALID_CHAR
        }
    }
}

func (self *Parser) decodeString(iv int64, ep int) (Node, native.ParsingError) {
    p := self.p - 1
    s := self.s[iv:p]

    /* fast path: no escape sequence */
    if ep == -1 {
        return newString(s), 0
    }

    /* unquote the string */
    buf := make([]byte, 0, len(s))
    err := unquoteBytes(s, &buf)

    /* check for errors */
    if err != 0 {
        return Node{}, err
    } else {
        return newBytes(buf), 0
    }
}

/** Parser Interface **/
// Pos returns current position at parsed json
func (self *Parser) Pos() int {
    return self.p
}

// Parse loads the inner json with lazy-load mode
func (self *Parser) Parse() (Node, native.ParsingError) {
    switch val := self.decodeValue(); val.Vt {
        case native.V_EOF     : return Node{}, native.ERR_EOF
        case native.V_NULL    : return nullNode, 0
        case native.V_TRUE    : return trueNode, 0
        case native.V_FALSE   : return falseNode, 0
    case native.V_ARRAY:
        if self.noLazy {
            return self.decodeArray(make([]Node, 0, _DEFAULT_NODE_CAP))
        }
        return newRawArray(self, make([]Node, 0, _DEFAULT_NODE_CAP)), 0
    case native.V_OBJECT:
        if self.noLazy {
            return self.decodeObject(make([]Pair, 0, _DEFAULT_NODE_CAP))
        }
        return newRawObject(self, make([]Pair, 0, _DEFAULT_NODE_CAP)), 0
        case native.V_STRING  : return self.decodeString(val.Iv, val.Ep)
        case native.V_DOUBLE  : return newFloat64(val.Dv), 0
        case native.V_INTEGER : return newInt64(val.Iv), 0
        default               : return Node{}, native.ParsingError(-val.Vt)
    }
}

func (self *Parser) skip() (int, native.ParsingError) {
    fsm := stackPool.Get().(*native.StateMachine)
    start := native.SkipOne(&self.s, &self.p, fsm)
    stackPool.Put(fsm)
    
    if start < 0 {
        return self.p, native.ParsingError(-start)
    }
    return start, 0
}

// Loads fully parses given json and returns:
//  the finally parsed position,
//  the interface{} representing results,
//  or ParsingError if any mistake happends
func Loads(src string) (int, interface{}, native.ParsingError) {
    ps := &Parser{s: src}
    np, err := ps.Parse()

    /* check for errors */
    if err != 0 {
        return 0, nil, err
    } 
    return ps.Pos(), np.Interface(), 0
}

// LoadUseNumber is almost same with Loads,
// except returned numbers are json.Number
func LoadUseNuumber(src string) (int, interface{}, native.ParsingError) {
    ps := &Parser{s: src}
    np, err := ps.Parse()

    /* check for errors */
    if err != 0 {
        return 0, nil, err
    } 
    return ps.Pos(), np.InterfaceUseNumber(), 0
}

// LoadUseNumber is almost same with Loads,
// except returned numbers are int64
func LoadUseInt64(src string) (int, interface{}, native.ParsingError) {
    ps := &Parser{s: src}
    np, err := ps.Parse()

    /* check for errors */
    if err != 0 {
        return 0, nil, err
    } 
    return ps.Pos(), np.InterfaceUseInt64(), 0
}



// NewParser.
func NewParser(src string) *Parser {
    return &Parser{s: src}
}