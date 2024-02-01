// +build !amd64,!arm64 go1.22 !go1.16 arm64,!go1.20

/*
* Copyright 2022 ByteDance Inc.
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

	`github.com/bytedance/sonic/internal/native/types`
	`github.com/bytedance/sonic/internal/rt`
)

func init() {
	println("WARNING:(ast) sonic only supports Go1.16~1.21, but your environment is not suitable")
}

func quote(buf *[]byte, val string) {
	quoteString(buf, val)
}

func unquote(src string) (string, types.ParsingError) {
	sp := rt.IndexChar(src, -1)
	out, ok := unquoteBytes(rt.BytesFrom(sp, len(src)+2, len(src)+2))
	if !ok {
		return "", types.ERR_INVALID_ESCAPE
	}
	return rt.Mem2Str(out), 0
}


func (self *Parser) decodeValue() (val types.JsonState) {
	e, v := decodeValue(self.s, self.p, self.dbuf == nil)
	if e < 0 {
		return v
	}
	self.p = e
	return v
}

func (self *Parser) skip() (int, types.ParsingError) {
	e, s := skipValue(self.s, self.p)
	if e < 0 {
		return self.p, types.ParsingError(-e)
	}
	self.p = e
	return s, 0
}

func (self *Parser) skipFast() (int, types.ParsingError) {
	e, s := skipValueFast(self.s, self.p)
	if e < 0 {
		return self.p, types.ParsingError(-e)
	}
	self.p = e
	return s, 0
}

func (self *Node) encodeInterface(buf *[]byte) error {
	out, err := json.Marshal(self.packAny())
	if err != nil {
		return err
	}
	*buf = append(*buf, out...)
	return nil
}
