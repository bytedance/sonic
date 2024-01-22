// +build amd64,go1.16,!go1.22

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
    `runtime`
    `unsafe`

    `github.com/bytedance/sonic/encoder`
    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    uq `github.com/bytedance/sonic/unquote`
    `github.com/chenzhuoyu/base64x`
)

var typeByte = rt.UnpackEface(byte(0)).Type

//go:nocheckptr
func quote(buf *[]byte, val string) {
    *buf = append(*buf, '"')
    if len(val) == 0 {
        *buf = append(*buf, '"')
        return
    }

    sp := rt.IndexChar(val, 0)
    nb := len(val)
    b := (*rt.GoSlice)(unsafe.Pointer(buf))

    // input buffer
    for nb > 0 {
        // output buffer
        dp := unsafe.Pointer(uintptr(b.Ptr) + uintptr(b.Len))
        dn := b.Cap - b.Len
        // call native.Quote, dn is byte count it outputs
        ret := native.Quote(sp, nb, dp, &dn, 0)
        // update *buf length
        b.Len += dn

        // no need more output
        if ret >= 0 {
            break
        }

        // double buf size
        *b = growslice(typeByte, *b, b.Cap*2)
        // ret is the complement of consumed input
        ret = ^ret
        // update input buffer
        nb -= ret
        sp = unsafe.Pointer(uintptr(sp) + uintptr(ret))
    }

    runtime.KeepAlive(buf)
    runtime.KeepAlive(sp)
    *buf = append(*buf, '"')
}

func unquote(src string) (string, types.ParsingError) {
    return uq.String(src)
}

func decodeBase64(src string) ([]byte, error) {
    return base64x.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64x.StdEncoding.EncodeToString(src)
}

func (self *Parser) decodeValue() (val types.JsonState) {
    sv := (*rt.GoString)(unsafe.Pointer(&self.s))
    flag := types.F_USE_NUMBER
    if self.dbuf != nil {
        flag = 0
        val.Dbuf = self.dbuf
        val.Dcap = types.MaxDigitNums
    }
    self.p = native.Value(sv.Ptr, sv.Len, self.p, &val, uint64(flag))
    return
}

func (self *Parser) skip() (int, types.ParsingError) {
    fsm := types.NewStateMachine()
    start := native.SkipOne(&self.s, &self.p, fsm, 0)
    types.FreeStateMachine(fsm)

    if start < 0 {
        return self.p, types.ParsingError(-start)
    }
    return start, 0
}

func (self *Node) encodeInterface(buf *[]byte) error {
    //WARN: NOT compatible with json.Encoder
    return encoder.EncodeInto(buf, self.packAny(), 0)
}

func (self *Parser) skipFast() (int, types.ParsingError) {
    start := native.SkipOneFast(&self.s, &self.p)
    if start < 0 {
        return self.p, types.ParsingError(-start)
    }
    return start, 0
}

func (self *Parser) getByPath(path ...interface{}) (int, types.ValueType, types.ParsingError) {
    fsm := types.NewStateMachine()
    start := native.GetByPath(&self.s, &self.p, &path, fsm)
    types.FreeStateMachine(fsm)
    runtime.KeepAlive(path)
    if start < 0 {
        return self.p, 0, types.ParsingError(-start)
    }
    t := switchRawType(self.s[start])
    if t == _V_NUMBER {
        self.p = 1 + backward(self.s, self.p-1)
    }
    return start, t, 0
}

func (self *Parser) getByPathNoValidate(path ...interface{}) (int, types.ValueType, types.ParsingError) {
    start := native.GetByPath(&self.s, &self.p, &path, nil)
    runtime.KeepAlive(path)
    if start < 0 {
        return self.p, 0, types.ParsingError(-start)
    }
    t := switchRawType(self.s[start])
    if t == _V_NUMBER {
        self.p = 1 + backward(self.s, self.p-1)
    }
    return start, t, 0
}

func DecodeString(src string, pos int) (ret int, v string) {
	p := NewParserObj(src)
    p.p = pos
	switch val := p.decodeValue(); val.Vt {
	case types.V_STRING:
        str := p.s[val.Iv : p.p-1]
		/* fast path: no escape sequence */
		if val.Ep == -1 {
			return p.p, str
		}
		/* unquote the string */
		out, err := unquote(str)
		/* check for errors */
		if err != 0 {
			return -int(err), ""
		} else {
			return p.p, out
		}
	default:
		return -int(_ERR_UNSUPPORT_TYPE), ""
	}
}

// ValidSyntax check if a json has a valid JSON syntax,
// while not validate UTF-8 charset
func ValidSyntax(json string) bool {
	fsm := types.NewStateMachine()
    p := 0
    ret := native.ValidateOne(&json, &p, fsm, 0)
    types.FreeStateMachine(fsm)

     if ret < 0 {
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
