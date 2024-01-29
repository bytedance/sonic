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
	"runtime"
	"unsafe"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/chenzhuoyu/base64x"
)

// Valid validates json and returns first non-blank character position,
// if it is only one valid json value.
// Otherwise returns invalid character position using start.
//
// Note: it does not check for the invalid UTF-8 characters.
func Valid(data []byte) (ok bool, start int) {
    n := len(data)
    if n == 0 {
        return false, -1
    }
    s := rt.Mem2Str(data)
    p := 0
    m := types.NewStateMachine()
    ret := native.ValidateOne(&s, &p, m)
    types.FreeStateMachine(m)

    if ret < 0 {
        return false, p-1
    }

    /* check for trailing spaces */
    for ;p < n; p++ {
        if (types.SPACE_MASK & (1 << data[p])) == 0 {
            return false, p
        }
    }

    return true, ret
}

func EncodeBase64(buf []byte, src []byte) []byte {
	if len(src) == 0 {
		return append(buf, '"', '"')
	}
	buf = append(buf, '"')
	need := base64x.StdEncoding.EncodedLen(len(src))
	if cap(buf) - len(buf) < need {
		tmp := make([]byte, len(buf), len(buf) + need*2)
		copy(tmp, buf)
		buf = tmp
	}
	base64x.StdEncoding.Encode(buf[len(buf):cap(buf)], src)
	buf = buf[:len(buf) + need]
	buf = append(buf, '"')
	return buf
}


var typeByte = rt.UnpackEface(byte(0)).Type

//go:nocheckptr
func Quote(buf []byte, val string, double bool) []byte {
	if len(val) == 0 {
		if double {
			return append(buf, `"\"\""`...)
		}
		return append(buf, `""`...)
	}

	if double {
		buf = append(buf, `"\"`...)
	} else {
		buf = append(buf, `"`...)
	}
	sp := rt.IndexChar(val, 0)
	nb := len(val)
	b := (*rt.GoSlice)(unsafe.Pointer(&buf))

	// input buffer
	for nb > 0 {
		// output buffer
		dp := unsafe.Pointer(uintptr(b.Ptr) + uintptr(b.Len))
		dn := b.Cap - b.Len
		// call native.Quote, dn is byte count it outputs
		opts := uint64(0)
		if double {
			opts = types.F_DOUBLE_UNQUOTE
		}
		ret := native.Quote(sp, nb, dp, &dn, opts)
		// update *buf length
		b.Len += dn

		// no need more output
		if ret >= 0 {
			break
		}

		// double buf size
		*b = rt.Growslice(typeByte, *b, b.Cap*2)
		// ret is the complement of consumed input
		ret = ^ret
		// update input buffer
		nb -= ret
		sp = unsafe.Pointer(uintptr(sp) + uintptr(ret))
	}

	runtime.KeepAlive(buf)
	runtime.KeepAlive(sp)
	if double {
		buf = append(buf, `\""`...)
	} else {
		buf = append(buf, `"`...)
	}

	return buf
}

func HtmlEscape(dst []byte, src []byte) []byte {
	var sidx int

	dst = append(dst, src[:0]...) // avoid check nil dst
	sbuf := (*rt.GoSlice)(unsafe.Pointer(&src))
	dbuf := (*rt.GoSlice)(unsafe.Pointer(&dst))

	/* grow dst if it is shorter */
	if cap(dst)-len(dst) < len(src)+native.BufPaddingSize {
		cap := len(src)*3/2 + native.BufPaddingSize
		*dbuf = rt.Growslice(typeByte, *dbuf, cap)
	}

	for sidx < sbuf.Len {
		sp := Padd(sbuf.Ptr, sidx)
		dp := Padd(dbuf.Ptr, dbuf.Len)

		sn := sbuf.Len - sidx
		dn := dbuf.Cap - dbuf.Len
		nb := native.HTMLEscape(sp, sn, dp, &dn)

		/* check for errors */
		if dbuf.Len += dn; nb >= 0 {
			break
		}

		/* not enough space, grow the slice and try again */
		sidx += ^nb
		*dbuf = rt.Growslice(typeByte, *dbuf, dbuf.Cap*2)
	}
	return dst
}
