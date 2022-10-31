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
    `runtime`
    `github.com/bytedance/sonic/internal/native`
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
    var err types.ParsingError
    self.parser.p = 0
    start := native.GetByPath(&self.parser.s, &self.parser.p, &path)
    // prohibit path collceted when executing native code.
    runtime.KeepAlive(path)
    if start < 0 {
        return Node{}, self.parser.syntaxError(types.ParsingError(-start))
    }
    t := switchRawType(self.parser.s[start])
    if t == _V_NONE {
        return Node{}, self.parser.ExportError(err)
    }
    return newRawNode(self.parser.s[start:self.parser.p], t), nil
}