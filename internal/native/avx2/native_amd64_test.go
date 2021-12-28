// Code generated by Makefile, DO NOT EDIT.

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

package avx2

import (
    `encoding/hex`
    `fmt`
    `math`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

func TestNative_Value(t *testing.T) {
    var v types.JsonState
    s := `   -12345`
    p := (*rt.GoString)(unsafe.Pointer(&s))
    x := __value(p.Ptr, p.Len, 0, &v, 0)
    assert.Equal(t, 9, x)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    assert.Equal(t, int64(-12345), v.Iv)
    assert.Equal(t, 3, v.Ep)
}

func TestNative_Value_OutOfBound(t *testing.T) {
    var v types.JsonState
    mem := []byte{'"', '"'}
    s := rt.Mem2Str(mem[:1])
    p := (*rt.GoString)(unsafe.Pointer(&s))
    x := __value(p.Ptr, p.Len, 0, &v, 0)
    assert.Equal(t, 1, x)
    assert.Equal(t, -int(types.ERR_EOF), int(v.Vt))
}

func TestNative_Quote(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 256)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, 0)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    assert.Equal(t, len(s), rv)
    assert.Equal(t, 27, len(d))
    assert.Equal(t, `hello\b\f\n\r\t\\\"景world`, string(d))
}

func TestNative_QuoteNoMem(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 10)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, 0)
    assert.Equal(t, -8, rv)
    assert.Equal(t, 9, len(d))
    assert.Equal(t, `hello\b\f`, string(d))
}

func TestNative_DoubleQuote(t *testing.T) {
    s := "hello\b\f\n\r\t\\\"\u666fworld"
    d := make([]byte, 256)
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __quote(sp.Ptr, sp.Len, dp.Ptr, &dp.Len, types.F_DOUBLE_UNQUOTE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    assert.Equal(t, len(s), rv)
    assert.Equal(t, 36, len(d))
    assert.Equal(t, `hello\\b\\f\\n\\r\\t\\\\\\\"景world`, string(d))
}

func TestNative_Unquote(t *testing.T) {
    s := `hello\b\f\n\r\t\\\"\u2333world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\b\f\n\r\t\\\"\u2333world", string(d))
}

func TestNative_UnquoteError(t *testing.T) {
    s := `asdf\`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_EOF), rv)
    assert.Equal(t, 5, ep)
    s = `asdf\gqwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_ESCAPE), rv)
    assert.Equal(t, 5, ep)
    s = `asdf\u1gggqwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_CHAR), rv)
    assert.Equal(t, 7, ep)
    s = `asdf\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 6, ep)
    s = `asdf\\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 7, ep)
    s = `asdf\ud800\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, 0)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 12, ep)
    s = `asdf\\ud800\\ud800qwer`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    assert.Equal(t, -int(types.ERR_INVALID_UNICODE), rv)
    assert.Equal(t, 14, ep)
}

func TestNative_DoubleUnquote(t *testing.T) {
    s := `hello\\b\\f\\n\\r\\t\\\\\\\"\\u2333world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_DOUBLE_UNQUOTE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\b\f\n\r\t\\\"\u2333world", string(d))
}

func TestNative_UnquoteUnicodeReplacement(t *testing.T) {
    s := `hello\ud800world`
    d := make([]byte, 0, len(s))
    ep := -1
    dp := (*rt.GoSlice)(unsafe.Pointer(&d))
    sp := (*rt.GoString)(unsafe.Pointer(&s))
    rv := __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_UNICODE_REPLACE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\ufffdworld", string(d))
    s = `hello\ud800\ud800world`
    d = make([]byte, 0, len(s))
    ep = -1
    dp = (*rt.GoSlice)(unsafe.Pointer(&d))
    sp = (*rt.GoString)(unsafe.Pointer(&s))
    rv = __unquote(sp.Ptr, sp.Len, dp.Ptr, &ep, types.F_UNICODE_REPLACE)
    if rv < 0 {
        require.NoError(t, types.ParsingError(-rv))
    }
    dp.Len = rv
    assert.Equal(t, -1, ep)
    assert.Equal(t, "hello\ufffd\ufffdworld", string(d))
}

func TestNative_Vstring(t *testing.T) {
    var v types.JsonState
    i := 0
    s := `test"test\n2"`
    __vstring(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, -1, v.Ep)
    assert.Equal(t, int64(0), v.Iv)
    __vstring(&s, &i, &v)
    assert.Equal(t, 13, i)
    assert.Equal(t, 9, v.Ep)
    assert.Equal(t, int64(5), v.Iv)
}

func TestNative_VstringEscapeEOF(t *testing.T) {
    var v types.JsonState
    i := 0
    s := `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\"xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"x`
    __vstring(&s, &i, &v)
    assert.Equal(t, 95, i)
    assert.Equal(t, 63, v.Ep)
    assert.Equal(t, int64(0), v.Iv)
}

func TestNative_VstringHangUpOnRandomData(t *testing.T) {
    v, e := hex.DecodeString(
        "228dc61efd54ef80a908fb6026b7f2d5f92a257ba8b347c995f259eb8685376a" +
        "8c4500262d9c308b3f3ec2577689cf345d9f86f9b5d18d3e463bec5c22df2d2e" +
        "4506010eba1dae7278",
    )
    assert.Nil(t, e)
    p := 1
    s := rt.Mem2Str(v)
    var js types.JsonState
    __vstring(&s, &p, &js)
    fmt.Printf("js: %s\n", spew.Sdump(js))
}

func TestNative_Vnumber(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "1.234"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 1.234, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "1.234e5"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 7, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 1.234e5, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "0.0125"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 6, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 0.0125, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "100000000000000000000"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 21, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 100000000000000000000.0, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "999999999999999900000"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 21, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, 999999999999999900000.0, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
    i = 0
    s = "-1.234"
    __vnumber(&s, &i, &v)
    assert.Equal(t, 6, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, -1.234, v.Dv)
    assert.Equal(t, types.V_DOUBLE, v.Vt)
}

func TestNative_Vsigned(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "-1234"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(-1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "9223372036854775807"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(math.MaxInt64), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "-9223372036854775808"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 20, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(math.MinInt64), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "9223372036854775808"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 18, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "-9223372036854775809"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "1.234"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "0.0125"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e5"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e-5"
    __vsigned(&s, &i, &v)
    assert.Equal(t, 5, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
}

func TestNative_Vunsigned(t *testing.T) {
    var v types.JsonState
    i := 0
    s := "1234"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, int64(1234), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "18446744073709551615"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 20, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, ^int64(0), v.Iv)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    i = 0
    s = "18446744073709551616"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 19, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INTEGER_OVERFLOW)), v.Vt)
    i = 0
    s = "-1234"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "1.234"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "0.0125"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 1, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "1234e5"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 4, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1234e5"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1.234e5"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
    i = 0
    s = "-1.234e-5"
    __vunsigned(&s, &i, &v)
    assert.Equal(t, 0, i)
    assert.Equal(t, 0, v.Ep)
    assert.Equal(t, types.ValueType(-int(types.ERR_INVALID_NUMBER_FMT)), v.Vt)
}

func TestNative_SkipOne(t *testing.T) {
    p := 0
    s := ` {"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
    q := __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 42, p)
    assert.Equal(t, 1, q)
    p = 0
    s = `1 2.5 -3 "asdf\nqwer" true false null {} []`
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 1, p)
    assert.Equal(t, 0, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 5, p)
    assert.Equal(t, 2, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 8, p)
    assert.Equal(t, 6, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 21, p)
    assert.Equal(t, 9, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 26, p)
    assert.Equal(t, 22, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 32, p)
    assert.Equal(t, 27, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 37, p)
    assert.Equal(t, 33, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 40, p)
    assert.Equal(t, 38, q)
    q = __skip_one(&s, &p, &types.StateMachine{})
    assert.Equal(t, 43, p)
    assert.Equal(t, 41, q)
}

func TestNative_SkipArray(t *testing.T) {
    p := 0
    s := `null, true, false, 1, 2.0, -3, {"asdf": "wqer"}],`
    __skip_array(&s, &p, &types.StateMachine{})
    assert.Equal(t, p, 48)
}

func TestNative_SkipObject(t *testing.T) {
    p := 0
    s := `"asdf": "wqer"},`
    __skip_object(&s, &p, &types.StateMachine{})
    assert.Equal(t, p, 15)
}
