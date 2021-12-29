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

    `github.com/bytedance/sonic/internal/native/types`
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
    self.parser.p = 0

    var err types.ParsingError
    for _, p := range path {
        switch p.(type) {
        case int:
            if err = self.parser.searchIndex(p.(int)); err != 0 {
                return Node{}, self.parser.ExportError(err)
            }
        case string:
            if err = self.parser.searchKey(p.(string)); err != 0 {
                return Node{}, self.parser.ExportError(err)
            }
        default:
            panic("path must be either int or string")
        }
    }

    var start = self.parser.p
    if start, err = self.parser.skip(); err != 0 {
        return Node{}, self.parser.ExportError(err)
    }
    ns := len(self.parser.s)
    if self.parser.p > ns || start >= ns || start>=self.parser.p {
        return Node{}, fmt.Errorf("skip %d char out of json boundary", start)
    }

    t := switchRawType(self.parser.s[start])
    if t == _V_NONE {
        return Node{}, self.parser.ExportError(err)
    }

    return newRawNode(self.parser.s[start:self.parser.p], t), nil
}