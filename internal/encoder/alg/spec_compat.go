// +build !amd64

/**
 * Copyright 2024 ByteDance Inc.
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

package alg

import (
	_ "unsafe"
	"unicode/utf8"
	"strconv"
	"bytes"
	"encoding/json"
	"encoding/base64"

	"github.com/bytedance/sonic/internal/rt"
)

// Valid validates json and returns first non-blank character position,
// if it is only one valid json value.
// Otherwise returns invalid character position using start.
//
// Note: it does not check for the invalid UTF-8 characters.
func Valid(data []byte) (ok bool, start int) {
    ok = json.Valid(data)
	return ok, 0
}

func EncodeBase64(buf []byte, src []byte) []byte {
	if len(src) == 0 {
		return append(buf, '"', '"')
	}
	buf = append(buf, '"')
	need := base64.StdEncoding.EncodedLen(len(src))
	if cap(buf) - len(buf) < need {
		tmp := make([]byte, len(buf), len(buf) + need*2)
		copy(tmp, buf)
		buf = tmp
	}
	base64.StdEncoding.Encode(buf[len(buf):cap(buf)], src)
	buf = buf[:len(buf) + need]
	buf = append(buf, '"')
	return buf
}


var typeByte = rt.UnpackEface(byte(0)).Type

func Quote(e []byte, s string, double bool, correct bool) []byte {
	if len(s) == 0 {
		if double {
			return append(e, `"\"\""`...)
		}
		return append(e, `""`...)
	}
	b := e
	
	e = append(e, '"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if safeSet[b] {
				i++
				continue
			}
			if start < i {
				e = append(e, s[start:i]...)
			}
			e = append(e, '\\')
			switch b {
			case '\\', '"':
				e = append(e, b)
			case '\n':
				e = append(e, 'n')
			case '\r':
				e = append(e, 'r')
			case '\t':
				e = append(e, 't')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				e = append(e, `u00`...)
				e = append(e, hex[b>>4])
				e = append(e, hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if correct && c == utf8.RuneError && size == 1 {
		    if start < i {
		        e = append(e, s[start:i]...)
		    }
		    e = append(e, `\ufffd`...)
		    i += size
		    start = i
		    continue
		}
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				e = append(e, s[start:i]...)
			}
			e = append(e, `\u202`...)
			e = append(e, hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		e = append(e, s[start:]...)
	}
	e = append(e, '"')

	if double {
		return strconv.AppendQuote(b, rt.Mem2Str(e))
	} else {
		return e
	}
}

var (
	//go:linkname safeSet encoding/json.safeSet
	safeSet [utf8.RuneSelf]bool

	//go:linkname hex encoding/json.hex
	hex string
)

func HtmlEscape(dst []byte, src []byte) []byte {
	buf := bytes.NewBuffer(dst)
	json.HTMLEscape(buf, src)
	return buf.Bytes()
}

func F64toa(buf []byte, v float64) ([]byte) {
	return strconv.AppendFloat(buf, float64(v), 'g', -1, 64)
}

func F32toa(buf []byte, v float64) ([]byte) {
	return strconv.AppendFloat(buf, float64(v), 'g', -1, 32)
}

func I64toa(buf []byte, v int64) ([]byte) {
	return buf = strconv.AppendInt(buf, int64(v), 10)
}

func U64toa(buf []byte, v uint64) ([]byte) {
	return buf = strconv.AppendUint(buf, v, 10)
}


