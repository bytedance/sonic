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
	"errors"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// Searcher used to skip and search json values
type Searcher struct {
    parser Parser
}

// NewSearcher
func NewSearcher(str string) *Searcher {
    return &Searcher{
        parser: Parser{
            s:      str,
            noLazy: false,
        },
    }
}

// GetByPath searches in the json and located a Node at path
func (self *Searcher) GetByPath(path ...interface{}) (Node, error) {
    var err types.ParsingError
    var start int

    self.parser.p = 0
    start, err = self.parser.getByPath(path...)
    if err != 0 {
        // for compatibility with old version
        if err == types.ERR_NOT_FOUND {
            return Node{}, ErrNotExist
        }
        if err == types.ERR_UNSUPPORT_TYPE {
            panic("path must be either int(>=0) or string")
        }
        return Node{}, self.parser.syntaxError(err)
    }

    t := switchRawType(self.parser.s[start])
    if t == _V_NONE {
        return Node{}, self.parser.ExportError(err)
    }
    return newRawNode(self.parser.s[start:self.parser.p], t), nil
}

// GetValueByPath searches in the json and located a Value at path
func (self *Searcher) GetValueByPath(path ...interface{}) (Value, error) {
	if self.parser.s == "" {
		err := errors.New("empty input")
		return errValue(err), err
	}

    self.parser.p = 0
    s, err := self.parser.getByPath(path...)
    if err != 0 {
		e := self.parser.ExportError(err)
        return errValue(e), e
    }

    t := switchRawType(self.parser.s[s])
    if t == _V_NONE {
		e := self.parser.ExportError(err)
        return errValue(e), e 
    }
    return Value{int(t), self.parser.s[s:self.parser.p]}, nil
}

// SetValueByPath searches and relpace a value at path.
// if path not exist, it will insert new json value
func (self *Searcher) SetValueByPath(val Value, path ...interface{}) (string, error) {
	if self.parser.s == "" {
		err := errors.New("empty input")
		return self.parser.s, err
	}

    self.parser.p = 0
    s, err := self.parser.getByPath(path...)

    if err != 0 {
        if err != _ERR_NOT_FOUND {
            e := self.parser.ExportError(err)
            return self.parser.s, e
        } else {
            // not exist, slow path
            n := value(self.parser.s)
            if  _, err := n.SetByPath(val, true, path...); err != nil {
                return self.parser.s, err
            }
            return n.js, nil
        }
    }

    // exist, fast-path replace
    e := self.parser.p
    b := make([]byte, 0, len(self.parser.s)+len(val.js)-(e-s))
    b = append(b, self.parser.s[:s]...)
    b = append(b, val.js...)
    b = append(b, self.parser.s[e:]...)
    return rt.Mem2Str(b), nil
}


// DeleteByPath searches and remove a value at path.
func (self *Searcher) DeleteByPath(path ...interface{}) (string, error) {
    if self.parser.s == "" {
		err := errors.New("empty input")
		return self.parser.s, err
	}
    
	// not exist, slow path
    n := value(self.parser.s)
    if  _, err := n.UnsetByPath(path...); err != nil {
        return self.parser.s, err
    }
    return n.js, nil
}
