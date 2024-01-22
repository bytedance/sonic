//go:build !amd64 || !go1.16 || go1.22
// +build !amd64 !go1.16 go1.22

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
	"encoding/base64"
	"encoding/json"
	"runtime"
	"unsafe"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

func init() {
    println("WARNING: sonic only supports Go1.16~1.21 && CPU amd64, but your environment is not suitable")
}

func Quote(buf *[]byte, val string) {
    quoteString(buf, val)
}

func Unquote(src string) (string, types.ParsingError) {
    sp := rt.IndexChar(src, -1)
    out, ok := unquoteBytes(rt.BytesFrom(sp, len(src)+2, len(src)+2))
    if !ok {
        return "", types.ERR_INVALID_ESCAPE
    }
    return rt.Mem2Str(out), 0
}

func decodeBase64(src string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64.StdEncoding.EncodeToString(src)
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

func (self *Parser) getByPath(path ...interface{}) (int, types.ValueType, types.ParsingError) {
    for _, p := range path {
        if idx, ok := p.(int); ok && idx >= 0 {
            if _, err := self.searchIndex(idx); err != 0 {
                return self.p, 0, err
            }
        } else if key, ok := p.(string); ok {
            if _, err := self.searchKey(key); err != 0 {
                return self.p, 0, err
            }
        } else {
            panic("path must be either int(>=0) or string")
        }
    }
    start, e := self.skip()
    if e != 0 {
        return self.p, 0, e
    }
    t := switchRawType(self.s[start])
    if t == _V_NUMBER {
        self.p = 1 + backward(self.s, self.p-1)
    }
    return start, t, 0
}

func (self *Parser) getByPathNoValidate(path ...interface{}) (int, types.ValueType, types.ParsingError) {
    return self.getByPath(path...)
}

//go:nocheckptr
func DecodeString(src string, pos int, needEsc bool) (v string, ret int, hasEsc bool) {
    ret, ep := skipString(src, pos)
    if ep == -1 {
        (*rt.GoString)(unsafe.Pointer(&v)).Ptr = rt.IndexChar(src, pos+1)
        (*rt.GoString)(unsafe.Pointer(&v)).Len = ret - pos - 2
        return v, ret, false
    } else if needEsc {
        return src[pos+1:ret-1], ret, true
    }

    vv, ok := unquoteBytes(rt.Str2Mem(src[pos:ret]))
    if !ok {
        return "", -int(types.ERR_INVALID_CHAR), true
    }

    runtime.KeepAlive(src)
    return rt.Mem2Str(vv), ret, true
}

// ValidSyntax check if a json has a valid JSON syntax,
// while not validate UTF-8 charset
func ValidSyntax(json string) bool {
	p, _ := skipValue(json, 0)
    if p < 0 {
        return false
    }
    /* check for trailing spaces */
    for ;p < len(json); p++ {
        if !isSpace(json[p]) {
            return false
        }
    }

    return true
}
