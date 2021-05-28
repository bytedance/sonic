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

package decoder

import (
    `encoding/json`
    `math`
    `reflect`
    `sync`
    `testing`
    `unsafe`

    `github.com/bytedance/sonic/ast`
    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/chenzhuoyu/base64x`
    `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
)

var (
    ref_stk_pool  = sync.Pool{New: new_stk}
    ref_prg_cache = map[reflect.Type]*_Program{}
)

func new_stk() interface{} {
    return make([]unsafe.Pointer, 0, _MaxStack)
}

func ref_eval(prg []_Instr, s string, p unsafe.Pointer) error {
    k := ref_stk_pool.Get().([]unsafe.Pointer)
    _, e := ref_eval_impl(prg, s, 0, p, k)
    ref_stk_pool.Put(k[:0])
    return e
}

func ref_unquote(s string, m *[]byte, du bool) int {
    var flags uint64
    if du {
        flags |= native.F_DOUBLE_UNQUOTE
    }

    /* unquote the string */
    pos := -1
    slv := (*rt.GoSlice)(unsafe.Pointer(m))
    str := (*rt.GoString)(unsafe.Pointer(&s))
    ret := native.Unquote(str.Ptr, str.Len, slv.Ptr, &pos, flags)

    /* check for errors */
    if ret < 0 {
        return -ret
    }

    /* update the length */
    slv.Len = ret
    return 0
}

func ref_eval_impl(prg []_Instr, s string, i int, p unsafe.Pointer, st []unsafe.Pointer) (int, error) {
    pc := 0
    lr := 0
    vv := native.JsonState{}
    mm := native.StateMachine{}

    for pc < len(prg) {
        ins := prg[pc]
        pc++

        switch ins.op() {
        case _OP_any:
            j, v, e := ast.Loads(s[i:])
            if e != 0 {
                return 0, e
            }
            i += j
            *(*interface{})(p) = v

        case _OP_str:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v := s[vv.Iv:i - 1]
            if vv.Ep == -1 {
                *(*string)(p) = v
                continue
            }
            m := make([]byte, 0, len(v))
            e := ref_unquote(v, &m, false)
            if e != 0 {
                return 0, native.ParsingError(e)
            }
            *(*string)(p) = rt.Mem2Str(m)

        case _OP_bin:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v, e := base64x.StdEncoding.DecodeString(s[vv.Iv:i - 1])
            if e != nil {
                return 0, e
            }
            *(*[]byte)(p) = v

        case _OP_bool:
            if i + 4 <= len(s) && s[i:i + 4] == "true" {
                i += 4
                *(*bool)(p) = true
            } else if i + 4 <= len(s) && s[i:i + 4] == "null" {
                i += 4
                *(*bool)(p) = false
            } else if i + 5 <= len(s) && s[i:i + 5] == "false" {
                i += 5
                *(*bool)(p) = false
            } else {
                return 0, native.ERR_INVALID_CHAR
            }

        case _OP_num,
            _OP_f32,
            _OP_f64:
            native.Vnumber(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            switch ins.op() {
            case _OP_num:
                *(*json.Number)(p) = json.Number(s[vv.Ep:i])

            case _OP_f32:
                if vv.Dv < -math.MaxFloat32 || vv.Dv > math.MaxFloat32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(float32(0.0)))
                }
                *(*float32)(p) = float32(vv.Dv)

            case _OP_f64:
                *(*float64)(p) = vv.Dv
            }

        case _OP_i8,
            _OP_i16,
            _OP_i32,
            _OP_i64:
            native.Vsigned(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            switch ins.op() {
            case _OP_i8:
                if vv.Iv < math.MinInt8 || vv.Iv > math.MaxInt8 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int8(0)))
                }
                *(*int8)(p) = int8(vv.Iv)

            case _OP_i16:
                if vv.Iv < math.MinInt16 || vv.Iv > math.MaxInt16 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int16(0)))
                }
                *(*int16)(p) = int16(vv.Iv)

            case _OP_i32:
                if vv.Iv < math.MinInt32 || vv.Iv > math.MaxInt32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int32(0)))
                }
                *(*int32)(p) = int32(vv.Iv)

            case _OP_i64:
                *(*int64)(p) = vv.Iv
            }

        case _OP_u8,
            _OP_u16,
            _OP_u32,
            _OP_u64:
            native.Vunsigned(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            switch ins.op() {
            case _OP_u8:
                if vv.Iv < 0 || vv.Iv > math.MaxUint8 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint8(0)))
                }
                *(*uint8)(p) = uint8(vv.Iv)

            case _OP_u16:
                if vv.Iv < 0 || vv.Iv > math.MaxUint16 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint16(0)))
                }
                *(*uint16)(p) = uint16(vv.Iv)

            case _OP_u32:
                if vv.Iv < 0 || vv.Iv > math.MaxUint32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint32(0)))
                }
                *(*uint32)(p) = uint32(vv.Iv)

            case _OP_u64:
                *(*uint64)(p) = uint64(vv.Iv)
            }

        case _OP_unquote:
            if i + 2 > len(s) {
                return 0, native.ERR_EOF
            }
            if s[i] != '\\' || s[i + 1] != '"' {
                return 0, native.ERR_INVALID_CHAR
            }
            i += 2
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            if vv.Ep == -1 {
                return 0, native.ERR_EOF
            }
            v := s[vv.Iv:i - 3]
            if vv.Ep == i - 3 {
                *(*string)(p) = v
                continue
            }
            m := make([]byte, 0, len(v))
            e := ref_unquote(v, &m, true)
            if e != 0 {
                return 0, native.ParsingError(e)
            }
            *(*string)(p) = rt.Mem2Str(m)

        case _OP_nil_1:
            *(*[1]uintptr)(p) = [1]uintptr{}

        case _OP_nil_2:
            *(*[2]uintptr)(p) = [2]uintptr{}

        case _OP_nil_3:
            *(*[3]uintptr)(p) = [3]uintptr{}

        case _OP_deref:
            v := (*unsafe.Pointer)(p)
            if *v == nil {
                t := rt.UnpackType(ins.vt())
                *v = mallocgc(uintptr(t.Size()), t, true)
            }
            p = *v

        case _OP_index:
            p = unsafe.Pointer(uintptr(p) + uintptr(ins.vi()))

        case _OP_is_null:
            if i + 4 <= len(s) && s[i:i + 4] == "null" {
                i += 4
                pc = ins.vi()
            }

        case _OP_map_init:
            v := (*unsafe.Pointer)(p)
            if *v == nil {
                *v = makemap_small()
            }
            p = *v

        case _OP_map_key_i8,
            _OP_map_key_i16,
            _OP_map_key_i32,
            _OP_map_key_i64:
            native.Vsigned(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            mt := rt.UnpackType(ins.vt())
            switch ins.op() {
            case _OP_map_key_i8:
                if vv.Iv < math.MinInt8 || vv.Iv > math.MaxInt8 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int8(0)))
                }
                p = mapassign(mt, p, unsafe.Pointer(&vv.Iv))

            case _OP_map_key_i16:
                if vv.Iv < math.MinInt16 || vv.Iv > math.MaxInt16 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int16(0)))
                }
                p = mapassign(mt, p, unsafe.Pointer(&vv.Iv))

            case _OP_map_key_i32:
                if vv.Iv < math.MinInt32 || vv.Iv > math.MaxInt32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(int32(0)))
                }
                p = mapassign_fast32(mt, p, uint32(vv.Iv))

            case _OP_map_key_i64:
                p = mapassign_fast64(mt, p, uint64(vv.Iv))
            }

        case _OP_map_key_u8,
            _OP_map_key_u16,
            _OP_map_key_u32,
            _OP_map_key_u64:
            native.Vunsigned(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            mt := rt.UnpackType(ins.vt())
            switch ins.op() {
            case _OP_map_key_u8:
                if vv.Iv < 0 || vv.Iv > math.MaxUint8 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint8(0)))
                }
                p = mapassign(mt, p, unsafe.Pointer(&vv.Iv))

            case _OP_map_key_u16:
                if vv.Iv < 0 || vv.Iv > math.MaxUint16 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint16(0)))
                }
                p = mapassign(mt, p, unsafe.Pointer(&vv.Iv))

            case _OP_map_key_u32:
                if vv.Iv < 0 || vv.Iv > math.MaxUint32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(uint32(0)))
                }
                p = mapassign_fast32(mt, p, uint32(vv.Iv))

            case _OP_map_key_u64:
                p = mapassign_fast64(mt, p, uint64(vv.Iv))
            }

        case _OP_map_key_f32,
            _OP_map_key_f64:
            native.Vnumber(&s, &i, &vv)
            if vv.Vt < 0 {
                return 0, native.ParsingError(-vv.Vt)
            }
            mt := rt.UnpackType(ins.vt())
            switch ins.op() {
            case _OP_map_key_f32:
                if vv.Dv < -math.MaxFloat32 || vv.Dv > math.MaxFloat32 {
                    return 0, error_value(s[vv.Ep:i], reflect.TypeOf(float32(0)))
                }
                x := float32(vv.Dv)
                p = mapassign(mt, p, unsafe.Pointer(&x))

            case _OP_map_key_f64:
                p = mapassign(mt, p, unsafe.Pointer(&vv.Dv))
            }

        case _OP_map_key_str:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v := s[vv.Iv:i - 1]
            if vv.Ep != -1 {
                m := make([]byte, 0, len(v))
                e := ref_unquote(v, &m, false)
                if e != 0 {
                    return 0, native.ParsingError(e)
                }
                v = rt.Mem2Str(m)
            }
            p = mapassign_faststr(rt.UnpackType(ins.vt()), p, v)

        case _OP_map_key_utext,
            _OP_map_key_utext_p:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v := s[vv.Iv:i - 1]
            if vv.Ep != -1 {
                m := make([]byte, 0, len(v))
                e := ref_unquote(v, &m, false)
                if e != 0 {
                    return 0, native.ParsingError(e)
                }
                v = rt.Mem2Str(m)
            }
            kk := ins.vt().Key()
            pk := rt.UnpackType(kk)
            kt := pk
            fn := mapassign
            if ins.op() == _OP_map_key_utext_p {
                pk = rt.UnpackType(reflect.PtrTo(kk))
            }
            if kk.Kind() == reflect.Ptr {
                kt = rt.UnpackType(kk.Elem())
                fn = mapassign_fast64ptr
            }
            kp := mallocgc(uintptr(kt.Size()), kt, true)
            if err := decodeTextUnmarshaler(rt.GoEface{Type: pk, Value: kp}.Pack(), v); err != nil {
                return 0, err
            }
            p = fn(rt.UnpackType(ins.vt()), p, kp)

        case _OP_array_skip:
            native.SkipArray(&s, &i, &mm)
            if i < 0 {
                return 0, native.ParsingError(-i)
            }

        case _OP_slice_init:
            v := (*rt.GoSlice)(p)
            v.Len = 0
            if v.Ptr == nil {
                v.Cap = 16
                v.Ptr = makeslice(rt.UnpackType(ins.vt()), 0, v.Cap)
            }

        case _OP_slice_append:
            sl := (*rt.GoSlice)(p)
            if sl.Len >= sl.Cap {
                *sl = growslice(rt.UnpackType(ins.vt()), *sl, sl.Cap * 2)
            }
            p = unsafe.Pointer(uintptr(sl.Ptr) + uintptr(sl.Len) * ins.vt().Size())
            sl.Len++

        case _OP_object_skip:
            native.SkipObject(&s, &i, &mm)
            if i < 0 {
                return 0, native.ParsingError(-i)
            }

        case _OP_object_next:
            q := native.SkipOne(&s, &i, &mm)
            if q < 0 {
                return 0, native.ParsingError(-q)
            }

        case _OP_struct_field:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v := s[vv.Iv:i - 1]
            if vv.Ep != -1 {
                m := make([]byte, 0, len(v))
                e := ref_unquote(v, &m, false)
                if e != 0 {
                    return 0, native.ParsingError(e)
                }
                v = rt.Mem2Str(m)
            }
            lr = ins.vf().Get(v)
            if lr == -1 {
                lr = ins.vf().GetCaseInsensitive(v)
            }

        case _OP_unmarshal,
            _OP_unmarshal_p:
            q := native.SkipOne(&s, &i, &mm)
            if q < 0 {
                return 0, native.ParsingError(-q)
            }
            v := s[q:i]
            kk := ins.vt()
            vp := p
            if kk.Kind() == reflect.Ptr {
                kp := (*unsafe.Pointer)(p)
                if *kp == nil {
                    kt := rt.UnpackType(kk.Elem())
                    *kp = mallocgc(uintptr(kt.Size()), kt, true)
                }
                vp = *kp
            }
            if ins.op() == _OP_unmarshal_p {
                kk = reflect.PtrTo(kk)
            }
            if err := decodeJsonUnmarshaler(rt.GoEface{Type: rt.UnpackType(kk), Value: vp}.Pack(), v); err != nil {
                return 0, err
            }

        case _OP_unmarshal_text,
            _OP_unmarshal_text_p:
            native.Vstring(&s, &i, &vv)
            if vv.Vt != native.V_STRING {
                return 0, native.ParsingError(-vv.Vt)
            }
            v := s[vv.Iv:i - 1]
            if vv.Ep != -1 {
                m := make([]byte, 0, len(v))
                e := ref_unquote(v, &m, false)
                if e != 0 {
                    return 0, native.ParsingError(e)
                }
                v = rt.Mem2Str(m)
            }
            kk := ins.vt()
            vp := p
            if kk.Kind() == reflect.Ptr {
                kp := (*unsafe.Pointer)(p)
                if *kp == nil {
                    kt := rt.UnpackType(kk.Elem())
                    *kp = mallocgc(uintptr(kt.Size()), kt, true)
                }
                vp = *kp
            }
            if ins.op() == _OP_unmarshal_text_p {
                kk = reflect.PtrTo(kk)
            }
            if err := decodeTextUnmarshaler(rt.GoEface{Type: rt.UnpackType(kk), Value: vp}.Pack(), v); err != nil {
                return 0, err
            }

        case _OP_lspace:
            sv := (*rt.GoString)(unsafe.Pointer(&s))
            if i = native.Lspace(sv.Ptr, sv.Len, i); i >= len(s) {
                return 0, native.ERR_EOF
            }
            if i < 0 {
                return 0, native.ParsingError(-i)
            }

        case _OP_match_char:
            if i == len(s) {
                return 0, native.ERR_EOF
            }
            if s[i] != ins.vb() {
                return 0, native.ERR_INVALID_CHAR
            }
            i++

        case _OP_check_char:
            if i == len(s) {
                return 0, native.ERR_EOF
            }
            if s[i] == ins.vb() {
                i++
                pc = ins.vi()
            }

        case _OP_load:
            p = st[len(st) - 1]

        case _OP_save:
            st = append(st, p)

        case _OP_drop:
            p = st[len(st) - 1]
            st = st[:len(st) - 1]

        case _OP_drop_2:
            p = st[len(st) - 2]
            st = st[:len(st) - 2]

        case _OP_recurse:
            var err error
            np, ok := ref_prg_cache[ins.vt()]
            if !ok {
                np, err = newCompiler().compile(ins.vt())
                if err != nil {
                    return 0, err
                }
                ref_prg_cache[ins.vt()] = np
            }
            if i, err = ref_eval_impl(np.ins, s, i, p, st); err != nil {
                return 0, err
            }

        case _OP_goto:
            pc = ins.vi()

        case _OP_switch:
            if lr >= 0 && lr < len(ins.vs()) {
                pc = ins.vs()[lr]
            }

        default:
            panic("invalid opcode: " + ins.op().String())
        }
    }

    return i, nil
}

func TestInterpreter_OpCodes_any(t *testing.T) {
    var v interface{}
    e := ref_eval([]_Instr{newInsOp(_OP_any)}, `{"a": [1.0, 2, -3]}`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, map[string]interface{}{"a": []interface{}{1.0, int64(2), int64(-3)}}, v)
}

func TestInterpreter_OpCodes_str(t *testing.T) {
    s := ""
    e := ref_eval([]_Instr{newInsOp(_OP_str)}, `hello, world"`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, "hello, world", s)
    s = ""
    e = ref_eval([]_Instr{newInsOp(_OP_str)}, `hello, world \\ \/ \b \f \n \r \t \u666f 测试中文 \ud83d\ude00"`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, "hello, world \\ / \b \f \n \r \t 景 测试中文 😀", s)
}

func TestInterpreter_OpCodes_bin(t *testing.T) {
    s := []byte(nil)
    e := ref_eval([]_Instr{newInsOp(_OP_bin)}, `aGVsbG8sIHdvcmxk"`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, []byte("hello, world"), s)
}

func TestInterpreter_OpCodes_bool_s(t *testing.T) {
    v := false
    e := ref_eval([]_Instr{newInsOp(_OP_bool)}, `true`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
    assert.True(t, v)
}

func TestInterpreter_OpCodes_num_s(t *testing.T) {{
    v := json.Number("")
    e := ref_eval([]_Instr{newInsOp(_OP_num)}, `12345`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, json.Number("12345"), v)
}; {
    v := int8(0)
    e := ref_eval([]_Instr{newInsOp(_OP_i8)}, `123`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, int8(123), v)
    v = 0
    e = ref_eval([]_Instr{newInsOp(_OP_i8)}, `-123`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
}; {
    v := uint64(0)
    e := ref_eval([]_Instr{newInsOp(_OP_u64)}, `1234567890123`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
}}

func TestInterpreter_OpCodes_unquote(t *testing.T) {
    v := ""
    e := ref_eval([]_Instr{newInsOp(_OP_unquote)}, `\"hello\\b\\f\\n\\r\\tworld\""`, unsafe.Pointer(&v))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, "hello\b\f\n\r\tworld", v)
}

func TestInterpreter_OpCodes_map(t *testing.T) {
    s := (map[string]string)(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"asdf":"qwer","zxcv":"fdgh"}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, map[string]string{"asdf": "qwer", "zxcv": "fdgh"}, s)
}

func TestInterpreter_OpCodes_map_i64(t *testing.T) {
    s := (map[int64]string)(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"1234":"qwer","-2345":"fdgh"}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, map[int64]string{1234: "qwer", -2345: "fdgh"}, s)
}

type UtextStruct struct {
    V string
}

func (self *UtextStruct) UnmarshalText(text []byte) error {
    self.V = string(text)
    return nil
}

func TestInterpreter_OpCodes_map_utext(t *testing.T) {
    s := (map[*UtextStruct]string)(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"asdf":"qwer","zxcv":"fdgh"}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, 2, len(s))
    m := map[string]string{}
    for k, v := range s {
        m[k.V] = v
    }
    assert.Equal(t, map[string]string{"asdf": "qwer", "zxcv": "fdgh"}, m)
}

func TestInterpreter_OpCodes_map_utext_p(t *testing.T) {
    s := map[UtextStruct]string{}
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"asdf":"qwer","zxcv":"fdgh"}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, 2, len(s))
    m := map[string]string{}
    for k, v := range s {
        m[k.V] = v
    }
    assert.Equal(t, map[string]string{"asdf": "qwer", "zxcv": "fdgh"}, m)
}

func TestInterpreter_OpCodes_array(t *testing.T) {
    s := [3]uint64{}
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `[1, 2, 3, 4, 5]`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, [3]uint64{1, 2, 3}, s)
}

func TestInterpreter_OpCodes_slice(t *testing.T) {
    s := []uint64(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `[1, 2, 3, 4, 5]`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, []uint64{1, 2, 3, 4, 5}, s)
}

type JsonStruct struct {
    A int
    B string
    C map[string]int
    D []int
}

func TestInterpreter_OpCodes_struct(t *testing.T) {
    s := JsonStruct{}
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"A": 123, "B": "asdf", "C": {"qwer": 4567}, "D": [1, 2, 3, 4, 5]}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, JsonStruct{
        A: 123,
        B: "asdf",
        C: map[string]int{"qwer": 4567},
        D: []int{1, 2, 3, 4, 5},
    }, s)
}

type UjsonStruct struct {
    V string
}

func (self *UjsonStruct) UnmarshalJSON(v []byte) error {
    self.V = string(v)
    return nil
}

func TestInterpreter_OpCodes_ujson(t *testing.T) {
    s := (*UjsonStruct)(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `{"test": "foo"}`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, "{\"test\": \"foo\"}", s.V)
}

func TestInterpreter_OpCodes_utext(t *testing.T) {
    s := (*UtextStruct)(nil)
    p, e := newCompiler().compile(reflect.TypeOf(s))
    if e != nil {
        panic(e)
    }
    e = ref_eval(p.ins, `"hello, world"`, unsafe.Pointer(&s))
    if e != nil {
        panic(e)
    }
    assert.Equal(t, "hello, world", s.V)
}

var _BindingValue TwitterStruct

type StringTag struct {
    BoolStr    bool        `json:",string"`
    IntStr     int64       `json:",string"`
    UintptrStr uintptr     `json:",string"`
    StrStr     string      `json:",string"`
    NumberStr  json.Number `json:",string"`
}

func init() {
    _ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func TestInterpreter_ParseJson(t *testing.T) {
    var v TwitterStruct
    prg, err := newCompiler().compile(reflect.TypeOf(v))
    if err != nil {
        panic(err)
    }
    err = ref_eval(prg.ins, TwitterJson, unsafe.Pointer(&v))
    if err != nil {
        panic(err)
    }
    assert.Equal(t, _BindingValue, v)
}

func TestInterpreter_ParseStringize(t *testing.T) {
    var v StringTag
    prg, err := newCompiler().compile(reflect.TypeOf(v))
    if err != nil {
        panic(err)
    }
    s := `{
        "BoolStr": "true",
        "IntStr": "42",
        "NumberStr": "46",
        "StrStr": "\"xzbit\"",
        "UintptrStr": "44"
    }`
    err = ref_eval(prg.ins, s, unsafe.Pointer(&v))
    if err != nil {
        panic(err)
    }
    assert.Equal(t, StringTag{
        BoolStr:    true,
        IntStr:     42,
        UintptrStr: 44,
        StrStr:     "xzbit",
        NumberStr:  "46",
    }, v)
}

func BenchmarkInterpreter_ParseJson_Sonic(b *testing.B) {
    var v TwitterStruct
    prg, err := newCompiler().compile(reflect.TypeOf(v))
    if err != nil {
        panic(err)
    }
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = ref_eval(prg.ins, TwitterJson, unsafe.Pointer(&v))
        }
    })
}

func BenchmarkInterpreter_ParseJson_JsonIter(b *testing.B) {
    var v TwitterStruct
    s := []byte(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = jsoniter.Unmarshal(s, &v)
        }
    })
}

func BenchmarkInterpreter_ParseJson_StdLib(b *testing.B) {
    var v TwitterStruct
    s := []byte(TwitterJson)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = json.Unmarshal(s, &v)
        }
    })
}
