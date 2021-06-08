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
    `fmt`

    `github.com/bytedance/sonic/internal/native`
)

// Searcher is used for path searches.
// It looks for specific key and skip unnecessary parsing
type Searcher struct {
    parser Parser
    // cache map[string]*Node
}

// NewSearch creates a Searcher with lazy-load parser by default
func NewSearcher(str string) *Searcher {
    return &Searcher{
        parser: Parser{
            s:      str,
            noLazy: false,
        },
    }
}


func (self *Parser) printNear(start int) string {
    start -= 10
    if start < 0 {
        start = 0
    }
    end := self.p + 10
    if end > len(self.s) {
        end = len(self.s)
    }
    return self.s[start:end]
}

// GetByPath searches the given path json,
// and returns its representing ast.Node
//
// Each path arg must be integer or string:
//     - Integer means searching current node as array,
//     - String means searching current node as object
func (self *Searcher) GetByPath(path ...interface{}) (Node, error) {
    self.parser.p = 0

    var err native.ParsingError
    for _, p := range path {
        switch p.(type) {
        case int:
            start := self.parser.p
            if err = self.parser.searchIndex(p.(int)); err != 0 {
                return Node{}, fmt.Errorf("%v at %d, near '%s'", err, start, self.parser.printNear(start))
            }
        case string:
            start := self.parser.p
            if err = self.parser.searchKey(p.(string)); err != 0 {
                return Node{}, fmt.Errorf("%v at %d, near '%s'", err, start, self.parser.printNear(start))
            }
        default:
            panic("path must be either int or string")
        }
    }

    var start int
    if start, err = self.parser.skip(); err != 0 {
        return Node{}, fmt.Errorf("%v at %d, near '%s'", err, self.parser.p, self.parser.printNear(start))
    }
    ns := len(self.parser.s)
    if self.parser.p > ns || start > ns {
        return Node{}, fmt.Errorf("skip %d char out of json boundary", start)
    }

    return newRawNode(self.parser.s[start:self.parser.p]), nil
}

func (self *Parser) searchKey(match string) native.ParsingError {
    ns := len(self.s)
    if err := self.object(); err != 0 {
        return err
    }

    /* check for EOF */
    if self.p = self.lspace(self.p); self.p >= ns {
        return native.ERR_EOF
    }

    /* check for empty object */
    if self.s[self.p] == '}' {
        self.p++
        return native.ERR_EOF
    }

    var njs native.JsonState
    var err native.ParsingError
    /* decode each pair */
    for {

        /* decode the key */
        if njs = self.decodeValue(); njs.Vt != native.V_STRING {
            return native.ERR_INVALID_CHAR
        }

        /* extract the key */
        idx := self.p - 1
        key := self.s[njs.Iv:idx]

        /* check for escape sequence */
        if njs.Ep != -1 {
            if key, err = UnquoteString(key); err != 0 {
                return err
            }
        }

        /* expect a ':' delimiter */
        if err = self.delim(); err != 0 {
            return err
        }

        /* skip value */
        if key != match {
            if _, err = self.skip(); err != 0 {
                return err
            }
        } else {
            return 0
        }

        /* check for EOF */
        self.p = self.lspace(self.p)
        if self.p >= ns {
            return native.ERR_EOF
        }

        /* check for the next character */
        switch self.s[self.p] {
        case ',':
            self.p++
        case '}':
            self.p++
            return native.ERR_EOF
        default:
            return native.ERR_INVALID_CHAR
        }
    }
}

func (self *Parser) searchIndex(idx int) native.ParsingError {
    ns := len(self.s)
    if err := self.array(); err != 0 {
        return err
    }

    /* check for EOF */
    if self.p = self.lspace(self.p); self.p >= ns {
        return native.ERR_EOF
    }

    /* check for empty array */
    if self.s[self.p] == ']' {
        self.p++
        return native.ERR_EOF
    }

    var err native.ParsingError
    /* allocate array space and parse every element */
    for i := 0; i < idx; i++ {

        /* decode the value */
        if _, err = self.skip(); err != 0 {
            return err
        }

        /* check for EOF */
        self.p = self.lspace(self.p)
        if self.p >= ns {
            return native.ERR_EOF
        }

        /* check for the next character */
        switch self.s[self.p] {
        case ',':
            self.p++
        case ']':
            self.p++
            return native.ERR_EOF
        default:
            return native.ERR_INVALID_CHAR
        }
    }

    return 0
}
