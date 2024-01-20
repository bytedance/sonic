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
	"strconv"
	"unsafe"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

func init() {
    println("WARNING: sonic only supports Go1.16~1.21 && CPU amd64, but your environment is not suitable")
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

func (self *Parser) getByPath(path ...interface{}) (int, types.ParsingError) {
    for _, p := range path {
        if idx, ok := p.(int); ok && idx >= 0 {
            if _, err := self.searchIndex(idx); err != 0 {
                return self.p, err
            }
        } else if key, ok := p.(string); ok {
            if _, err := self.searchKey(key); err != 0 {
                return self.p, err
            }
        } else {
            panic("path must be either int(>=0) or string")
        }
    }
    start, e := self.skip()
    if e != 0 {
        return self.p, e
    }
    return start, 0
}

func (self *Parser) getByPathNoValidate(path ...interface{}) (int, types.ParsingError) {
    return self.getByPath(path...)
}

//go:nocheckptr
func decodeString(src string, pos int) (ret int, v string) {
    ret, ep := skipString(src, pos)
    if ep == -1 {
        (*rt.GoString)(unsafe.Pointer(&v)).Ptr = rt.IndexChar(src, pos+1)
        (*rt.GoString)(unsafe.Pointer(&v)).Len = ret - pos - 2
        return ret, v
    }

    vv, ok := unquoteBytes(rt.Str2Mem(src[pos:ret]))
    if !ok {
        return -int(types.ERR_INVALID_CHAR), ""
    }

    runtime.KeepAlive(src)
    return ret, rt.Mem2Str(vv)
}

//go:nocheckptr
func decodeInt64(src string, pos int) (ret int, v int64, err error) {
    sp := uintptr(rt.IndexChar(src, pos))
    ss := uintptr(sp)
    se := uintptr(rt.IndexChar(src, len(src)))
    if uintptr(sp) >= se {
        return -int(types.ERR_EOF), 0, nil
    }

    if c := *(*byte)(unsafe.Pointer(sp)); c == '-' {
        sp += 1
    }
    if sp == se {
        return -int(types.ERR_EOF), 0, nil
    }

    for ; sp < se; sp += uintptr(1) {
        if !isDigit(*(*byte)(unsafe.Pointer(sp))) {
            break
        }
    }

    if sp < se {
        if c := *(*byte)(unsafe.Pointer(sp)); c == '.' || c == 'e' || c == 'E' {
            return -int(types.ERR_INVALID_NUMBER_FMT), 0, nil
        }
    }

    var vv string
    ret = int(uintptr(sp) - uintptr((*rt.GoString)(unsafe.Pointer(&src)).Ptr))
    (*rt.GoString)(unsafe.Pointer(&vv)).Ptr = unsafe.Pointer(ss)
    (*rt.GoString)(unsafe.Pointer(&vv)).Len = ret - pos

    v, err = strconv.ParseInt(vv, 10, 64)
    if err != nil {
        //NOTICE: allow overflow here
        if err.(*strconv.NumError).Err == strconv.ErrRange {
            return ret, 0, err
        }
        return -int(types.ERR_INVALID_CHAR), 0, err
    }

    runtime.KeepAlive(src)
    return ret, v, nil
}

//go:nocheckptr
func decodeFloat64(src string, pos int) (ret int, v float64, err error) {
    sp := uintptr(rt.IndexChar(src, pos))
    ss := uintptr(sp)
    se := uintptr(rt.IndexChar(src, len(src)))
    if uintptr(sp) >= se {
        return -int(types.ERR_EOF), 0, nil
    }

    if c := *(*byte)(unsafe.Pointer(sp)); c == '-' {
        sp += 1
    }
    if sp == se {
        return -int(types.ERR_EOF), 0, nil
    }

    for ; sp < se; sp += uintptr(1) {
        if !isNumberChars(*(*byte)(unsafe.Pointer(sp))) {
            break
        }
    }

    var vv string
    ret = int(uintptr(sp) - uintptr((*rt.GoString)(unsafe.Pointer(&src)).Ptr))
    (*rt.GoString)(unsafe.Pointer(&vv)).Ptr = unsafe.Pointer(ss)
    (*rt.GoString)(unsafe.Pointer(&vv)).Len = ret - pos

    v, err = strconv.ParseFloat(vv, 64)
    if err != nil {
        //NOTICE: allow overflow here
        if err.(*strconv.NumError).Err == strconv.ErrRange {
            return ret, 0, err
        }
        return -int(types.ERR_INVALID_CHAR), 0, err
    }

    runtime.KeepAlive(src)
    return ret, v, nil
}
