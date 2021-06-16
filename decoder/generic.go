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
    `reflect`
    `sync`
    `unsafe`

    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
    `github.com/bytedance/sonic/unquote`
)

const (
    _S_val = iota
    _S_arr
    _S_arr_0
    _S_obj
    _S_obj_x
    _S_obj_delim
)

var (
    mapVal  = map[string]interface{}(nil)
	mapType = rt.UnpackType(reflect.TypeOf(mapVal))
)

type _GenericDecoder struct {
    types.StateMachine
    Vp [types.MAX_RECURSE]*interface{}
}

func (self *_GenericDecoder) val(v interface{}) types.ParsingError {
    sp := self.Sp
    vt := self.Vt[sp]

    /* must be a value or the first element of an array */
    if vt != _S_val && vt != _S_arr_0 {
        return types.ERR_INVALID_CHAR
    }

    /* set the value */
    self.Sp--
    *self.Vp[sp] = v
    return 0
}

func (self *_GenericDecoder) exec(s string, i int, f uint64) (int, interface{}, types.ParsingError) {
    var rv interface{}
    var ss types.JsonState
    var ex types.ParsingError

    /* initialize the state machine */
    self.Sp = 0
    self.Vp[0] = &rv
    self.Vt[0] = _S_val

    /* string length and pointer */
    slen := len(s)
    sbuf := (*rt.GoString)(unsafe.Pointer(&s)).Ptr

    /* run until the state goes empty */
    for self.Sp >= 0 {
        switch i = native.Value(sbuf, slen, i, &ss, 1); ss.Vt {
            default: {
                return i, nil, types.ParsingError(-ss.Vt)
            }

            /* EOF */
            case types.V_EOF: {
                return i, nil, types.ERR_EOF
            }

            /* null */
            case types.V_NULL: {
                if ex = self.val(nil); ex != 0 {
                    return i - 4, nil, ex
                }
            }

            /* boolean true */
            case types.V_TRUE: {
                if ex = self.val(true); ex != 0 {
                    return i - 4, nil, ex
                }
            }

            /* boolean false */
            case types.V_FALSE: {
                if ex = self.val(false); ex != 0 {
                    return i - 5, nil, ex
                }
            }

            /* strings */
            case types.V_STRING: {
                p := i - 1
                v := s[ss.Iv:p]

                /* check for escape sequence */
                if ss.Ep != -1 {
                    if v, ex = unquote.String(v); ex != 0 {
                        return int(ss.Iv) - 1, nil, ex
                    }
                }

                /* check for object key */
                if vt := self.Vt[self.Sp]; vt != _S_obj && vt != _S_obj_x {
                    if ex = self.val(v); ex != 0 {
                        return int(ss.Iv) - 1, nil, ex
                    } else {
                        continue
                    }
                }

                /* get the map */
                ef := rt.UnpackEface(*self.Vp[self.Sp])
                mp := ef.Value

                /* add the delimiter */
                self.Sp++
                self.Vt[self.Sp] = _S_obj_delim
                self.Vp[self.Sp] = (*interface{})(mapassign_faststr(mapType, mp, v))
            }

            /* nested arrays */
            case types.V_ARRAY: {
                vt := self.Vt[self.Sp]
                vv := make([]interface{}, 1, 16)

                /* must be a value */
                if vt != _S_val && vt != _S_arr_0 {
                    return i - 1, nil, types.ERR_INVALID_CHAR
                }

                /* set the value */
                self.Vt[self.Sp] = _S_arr
                *self.Vp[self.Sp] = vv

                /* add the first element */
                self.Sp++
                self.Vt[self.Sp] = _S_arr_0
                self.Vp[self.Sp] = &vv[0]
            }

            /* nested objects */
            case types.V_OBJECT: {
                vt := self.Vt[self.Sp]
                vv := map[string]interface{}{}

                /* must be a value */
                if vt != _S_val && vt != _S_arr_0 {
                    return i - 1, nil, types.ERR_INVALID_CHAR
                }

                /* set the value */
                self.Vt[self.Sp] = _S_obj
                *self.Vp[self.Sp] = vv
            }

            /* floating point numbers */
            case types.V_DOUBLE: {
                if (f & (1 << _F_use_number)) == 0 {
                    if ex = self.val(ss.Dv); ex != 0 {
                        return ss.Ep, nil, ex
                    }
                } else {
                    if ex = self.val(json.Number(s[ss.Ep:i])); ex != 0 {
                        return ss.Ep, nil, ex
                    }
                }
            }

            /* integers */
            case types.V_INTEGER: {
                if (f & (1 << _F_use_number)) != 0 {
                    if ex = self.val(json.Number(s[ss.Ep:i])); ex != 0 {
                        return ss.Ep, nil, ex
                    }
                } else if (f & (1 << _F_use_int64)) == 0 {
                    if ex = self.val(float64(ss.Iv)); ex != 0 {
                        return ss.Ep, nil, ex
                    }
                } else {
                    if ex = self.val(ss.Iv); ex != 0 {
                        return ss.Ep, nil, ex
                    }
                }
            }

            /* key separator ':' */
            case types.V_KEY_SEP: {
                if self.Vt[self.Sp] == _S_obj_delim {
                    self.object_elem()
                } else {
                    return i - 1, nil, types.ERR_INVALID_CHAR
                }
            }

            /* element separator ',' */
            case types.V_ELEM_SEP: {
                switch self.Vt[self.Sp] {
                    case _S_obj : self.Vt[self.Sp] = _S_obj_x
                    case _S_arr : self.array_push(self.Vp[self.Sp])
                    default     : return i - 1, nil, types.ERR_INVALID_CHAR
                }
            }

            /* array end ']' */
            case types.V_ARRAY_END: {
                switch self.Vt[self.Sp] {
                    case _S_arr   : self.Sp--
                    case _S_arr_0 : self.array_end()
                    default       : return i - 1, nil, types.ERR_INVALID_CHAR
                }
            }

            /* object end '}' */
            case types.V_OBJECT_END: {
                if self.Vt[self.Sp] == _S_obj {
                    self.Sp--
                } else {
                    return i - 1, nil, types.ERR_INVALID_CHAR
                }
            }
        }
    }

    /* all done */
    return i, rv, 0
}

func (self *_GenericDecoder) array_end() {
    (*rt.GoSlice)((*rt.GoEface)(unsafe.Pointer(self.Vp[self.Sp - 1])).Value).Len = 0
    self.Sp -= 2
}

func (self *_GenericDecoder) array_add(v *interface{}) *interface{} {
    vv := (*[]interface{})((*rt.GoEface)(unsafe.Pointer(v)).Value)
    nb := len(*vv)
    *vv = append(*vv, nil)
    return &(*vv)[nb]
}

func (self *_GenericDecoder) array_push(v *interface{}) {
    self.Sp++
    self.Vt[self.Sp] = _S_val
    self.Vp[self.Sp] = self.array_add(v)
}

func (self *_GenericDecoder) object_elem() {
    self.Vt[self.Sp] = _S_val
    self.Vt[self.Sp - 1] = _S_obj
}

var decoderPool = sync.Pool {
    New: func() interface{} {
        return new(_GenericDecoder)
    },
}

func decodeGeneric(s string, i int, f uint64) (p int, v interface{}, e types.ParsingError) {
    dec := decoderPool.Get().(*_GenericDecoder)
    p, v, e = dec.exec(s, i, f)
    decoderPool.Put(dec)
    return
}
