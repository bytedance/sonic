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
)

type Searcher struct {
    parser Parser
}

func NewSearcher(str string) *Searcher {
    return &Searcher{
        parser: Parser{
            s:      str,
            noLazy: false,
        },
    }
}

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

// GetValueByPath
func (self Searcher) GetValueByPath(path ...interface{}) (Value, error) {
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