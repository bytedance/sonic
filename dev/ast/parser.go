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
	"runtime"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/davecgh/go-spew/spew"
)

type Parser struct {
    pos	int
    src	string
}

/** Parser Interface **/

func (self *Parser) Pos() int {
    return self.pos
}


// parse one layer but no validate sub layers
func (self *Parser) Parse() (Node, error) {
	node, err := parseLazy(self.src, nil)
    spew.Dump(node.node)
	if err != nil {
		return Node{}, err
	}
	return node, nil
}

// NewParser returns pointer of new allocated parser
func NewParser(src string) *Parser {
    return &Parser{src: src}
}

// get Node and  no validate
func (self *Parser) GetByPath(path ...interface{}) (Node, error) {
    return parseLazy(self.src[self.pos:], &path)
}

// get location no validate
func (self *Parser) locate(flags uint64, path ...interface{}) (start int, err types.ParsingError) {
    start = native.GetByPath(&self.src, &self.pos, &path, nil, flags)
    println("s", start, "p", self.pos)
    if start < 0 {
        return -1, types.ParsingError(-start)
    }
    runtime.KeepAlive(path)
    return start, 0
}

// get location and validate
func (self *Parser) locateAndValidate(flags uint64, path ...interface{}) (start int, err types.ParsingError) {
    var fsm = types.NewStateMachine()
    start = native.GetByPath(&self.src, &self.pos, &path, fsm, flags)
    types.FreeStateMachine(fsm)
    if start < 0 {
        return -1, types.ParsingError(-start)
    }
    runtime.KeepAlive(path)
    return start, 0
}

// skip and validate
func (self *Parser)  Skip() (start int, err error ) {
    fsm := types.NewStateMachine()
    start = native.SkipOne(&self.src, &self.pos, fsm, 0)
    types.FreeStateMachine(fsm)
    if start < 0 {
        if -start == int(types.ERR_NOT_FOUND) {
            return -1, ErrNotExist
        }
        return -1, makeSyntaxError(self.src, self.pos, types.ParsingError(-start).Message())
    }
    return start, nil
}

// skip no validate
func (self *Parser) skip() (int, types.ParsingError) {
    start := native.SkipOneFast(&self.src, &self.pos)
    if start < 0 {
        return -1, types.ParsingError(-start)
    }
    return start, 0
}

func Get(src []byte, path ...interface{}) (Node, error) {
    node, err := parseLazy(rt.Mem2Str(src), &path)
    // NOTICE: always copy JSON string to avoid memory leak
    r := node
    r.node.JSON = rt.Mem2Str([]byte(node.node.JSON))
    return node, err
}

func GetFromString(src string, path ...interface{}) (Node, error) {
    return parseLazy(src, &path)
}
