// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vm

import (
	"encoding"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"sync"
	"unsafe"

	"github.com/bytedance/sonic/internal/encoder/alg"
	"github.com/bytedance/sonic/internal/encoder/ir"
	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/bytedance/sonic/internal/rt"
)

// const (
// 	_Realloc_Threshold_0 = 512
// 	_Realloc_Threshold_1 = 4096
// 	_Realloc_Threshold_2 = 128 * 1024
// 	_Realloc_Threshold_3 = 1024 * 1024
// )

// func ralloc(w unsafe.Pointer, c int, l int, n int) (unsafe.Pointer, int) {
// 	if c < _Realloc_Threshold_0 {
// 		c = c<<2 + n + l
// 	} else if c < _Realloc_Threshold_1 {
// 		c = c<<1 + n + l
// 	} else if c < _Realloc_Threshold_2 {
// 		c = c + n + l
// 	} else if c < _Realloc_Threshold_3 {
// 		c = c>>1 + n + l
// 	} else {
// 		c = c>>2 + n + l
// 	}
// 	buf := rt.BytesFrom(unsafe.Pointer(uintptr(w)-uintptr(l)), l, l)
// 	tmp := make([]byte, l, c)
// 	copy(tmp, buf)
// 	w = rt.IndexByte(tmp, l)
// 	return w, c
// }

// func check_size(ow unsafe.Pointer, oc int, l int, n int) (unsafe.Pointer, int) {
// 	if oc-l <= n {
// 		ow, oc = ralloc(ow, oc, l, n)
// 	}
// 	return ow, oc
// }

// func write_char(ow unsafe.Pointer, ol int, oc int, char byte) (w unsafe.Pointer, l int, c int) {
// 	w, c = check_size(ow, oc, ol, 1)
// 	*(*byte)(w) = char
// 	w = unsafe.Add(w, 1)
// 	l = ol + 1
// 	return
// }

// func write_str(ow unsafe.Pointer, ol int, oc int, str string) (w unsafe.Pointer, l int, c int) {
// 	return write_bytes(ow, ol, oc, rt.IndexChar(str, 0), len(str))
// }

// func write_bytes(ow unsafe.Pointer, ol int, oc int, p unsafe.Pointer, n int) (w unsafe.Pointer, l int, c int) {
// 	w, c = check_size(ow, oc, ol, n)
// 	memmove(w, p, uintptr(n))
// 	w = unsafe.Add(w, n)
// 	l = ol + n
// 	return
// }

type stackHolds struct {
	p unsafe.Pointer
}

var shPools = sync.Pool{
	New: func() interface{} {
		return &stackHolds{}
	},
}

func newStackHolds() *stackHolds {
	return shPools.Get().(*stackHolds)
}

func freeStackHolds(sh *stackHolds) {
	sh.p = nil
	shPools.Put(sh)
}

const (
	_S_cond = iota
	_S_init
)

var (
	_T_json_Marshaler         = rt.UnpackType(vars.JsonMarshalerType)
	_T_encoding_TextMarshaler = rt.UnpackType(vars.EncodingTextMarshalerType)
)

func ExecVM(b *[]byte, p unsafe.Pointer, s *vars.Stack, flags uint64, prog ir.Program) (error) {
	pl := len(prog)
	if pl <= 0 {
		return nil
	}

	var buf = *b
	var x int
	var q unsafe.Pointer
	var f uint64

	var pro = &(prog)[0]
	for pc := 0; pc < pl; {
		ins := (*ir.Instr)(unsafe.Add(unsafe.Pointer(pro), ir.OpSize*uintptr(pc)))
		pc++
		op := ins.Op()

		// if len(buf) > 20 {
		// 	fmt.Println(string(buf[len(buf)-20:]))
		// } else {
		// 	fmt.Println(string(buf))
		// }
		// fmt.Printf("\npc %04d, op %v, ins %#v\n", pc, op, ins.Disassemble())
		
		switch op {
		case ir.OP_goto:
			pc = ins.Vi()
			continue
		case ir.OP_byte:
			v := ins.Byte()
			buf = append(buf, v)
		case ir.OP_text:
			v := ins.Vs()
			buf = append(buf, v...)
		case ir.OP_deref:
			p = *(*unsafe.Pointer)(p)
		case ir.OP_index:
			p = unsafe.Add(p, uintptr(ins.I64()))
		case ir.OP_load:
			// NOTICE: load CANNOT change f!
			x, _, p, q = s.Load() 
		case ir.OP_save:
			if !s.Save(x, f, p, q) {
				return vars.ERR_too_deep
			}
		case ir.OP_drop:
			x, f, p, q = s.Drop()
		case ir.OP_drop_2:
			s.Drop()
			x, f, p, q = s.Drop()
		case ir.OP_recurse:
			vt, pv := ins.Vp2()
			f := flags
			if pv {
				f |= alg.BitPointerValue
			}
			*b = buf
			if vt.Indirect() {
				sh := newStackHolds()
				sh.p = p
				if err := EncodeTypedPointer(b, vt, &sh.p, s, f); err != nil {
					freeStackHolds(sh)
					return err
				}
				freeStackHolds(sh)
			} else {
				vp := (*unsafe.Pointer)(p)
				if err := EncodeTypedPointer(b, vt, vp, s, f); err != nil {
					return err
				}
			}
			buf = *b
		case ir.OP_is_nil:
			if is_nil(p) {
				pc = ins.Vi()
				continue
			}
		case ir.OP_is_nil_p1:
			if (*rt.GoEface)(p).Value == nil {
				pc = ins.Vi()
				continue
			}
		case ir.OP_null:
			buf = append(buf, "null"...)
		case ir.OP_str:
			v := *(*string)(p)
			buf = alg.Quote(buf, v, false)
		case ir.OP_bool:
			if *(*bool)(p) {
				buf = append(buf, "true"...)
			} else {
				buf = append(buf, "false"...)
			}
		case ir.OP_i8:
			v := *(*int8)(p)
			strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_i16:
			v := *(*int16)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_i32:
			v := *(*int32)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_i64:
			v := *(*int64)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_u8:
			v := *(*uint8)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_u16:
			v := *(*uint16)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_u32:
			v := *(*uint32)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_u64:
			v := *(*uint64)(p)
			buf = strconv.AppendInt(buf, int64(v), 10)
		case ir.OP_f32:
			v := *(*float32)(p)
			if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
				return vars.ERR_nan_or_infinite
			}
			buf = strconv.AppendFloat(buf, float64(v), 'g', -1, 32)
		case ir.OP_f64:
			v := *(*float64)(p)
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return vars.ERR_nan_or_infinite
			}
			buf = strconv.AppendFloat(buf, float64(v), 'g', -1, 64)
		case ir.OP_bin:
			v := *(*[]byte)(p)
			buf = alg.EncodeBase64(buf, v)
		case ir.OP_quote:
			v := *(*string)(p)
			buf = alg.Quote(buf, v, true)
		case ir.OP_number:
			v := *(*json.Number)(p)
			if v == "" {
				buf = append(buf, '0')
			} else if !rt.IsValidNumber(string(v)) {
				return vars.Error_number(v)
			} else {
				buf = append(buf, v...)
			}
		case ir.OP_eface:
			*b = buf
			if err := EncodeTypedPointer(b, *(**rt.GoType)(p), (*unsafe.Pointer)(unsafe.Add(p, 8)), s, flags); err != nil {
				return err
			}
			buf = *b
		case ir.OP_iface:
			*b = buf
			if err := EncodeTypedPointer(b,  (*(**rt.GoItab)(p)).Vt, (*unsafe.Pointer)(unsafe.Add(p, 8)), s, flags); err != nil {
				return err
			}
			buf = *b
		case ir.OP_is_zero_map:
			v := *(**rt.GoMap)(p)
			if v == nil || v.Count == 0 {
				pc = ins.Vi()
				continue
			}
		case ir.OP_map_iter:
			v := *(**rt.GoMap)(p)
			vt := ins.Vr()
			it, err := alg.IteratorStart(rt.MapType(vt), v, flags)
			if err != nil {
				return err
			}
			q = unsafe.Pointer(it)
		case ir.OP_map_stop:
			it := (*alg.MapIterator)(q)
			alg.IteratorStop(it)
			q = nil
		case ir.OP_map_value_next:
			it := (*alg.MapIterator)(q)
			p = it.It.V
			alg.IteratorNext(it)
		case ir.OP_map_check_key:
			it := (*alg.MapIterator)(q)
			if it.It.K == nil {
				pc = ins.Vi()
				continue
			}
			p = it.It.K
		case ir.OP_map_write_key:
			if has_opts(flags, alg.BitSortMapKeys) {
				v := *(*string)(p)
				buf = alg.Quote(buf, v, false)
				pc = ins.Vi()
				continue
			}
		case ir.OP_slice_len:
			v := (*rt.GoSlice)(p)
			x = v.Len
			p = v.Ptr
			//TODO: why?
			f |= 1<<_S_init 
		case ir.OP_slice_next:
			if x == 0 {
				pc = ins.Vi()
				continue
			}
			x--
			if has_opts(f, _S_init) {
				f &= ^uint64(1 << _S_init)
			} else {
				p = unsafe.Add(p, uintptr(ins.Vlen()))
			}
		case ir.OP_cond_set:
			f |= 1<<_S_cond
		case ir.OP_cond_testc:
			fmt.Printf("%x\n", f)
			if has_opts(f, _S_cond) {
				f &= ^uint64(1 << _S_cond)
				pc = ins.Vi()
				continue
			}
		case ir.OP_is_zero_1:
			if *(*uint8)(p) == 0 {
				pc = ins.Vi()
				continue
			}
		case ir.OP_is_zero_2:
			if *(*uint16)(p) == 0 {
				pc = ins.Vi()
				continue
			}
		case ir.OP_is_zero_4:
			if *(*uint32)(p) == 0 {
				pc = ins.Vi()
				continue
			}
		case ir.OP_is_zero_8:
			if *(*uint64)(p) == 0 {
				pc = ins.Vi()
				continue
			}
		case ir.OP_empty_arr:
			if has_opts(flags, alg.BitNoNullSliceOrMap) {
				buf = append(buf, "[]"...)
			} else {
				buf = append(buf, "null"...)
			}
		case ir.OP_empty_obj:
			if has_opts(flags, alg.BitNoNullSliceOrMap) {
				buf = append(buf, "{}"...)
			} else {
				buf = append(buf, "null"...)
			}
		case ir.OP_marshal:
			vt := ins.Vr()
			var err error
			if buf, err = call_json_marshaler(buf, vt, p, flags, false); err != nil {
				return err
			}
		case ir.OP_marshal_p:
			vt := ins.Vr()
			var err error
			if buf, err = call_json_marshaler(buf, vt, p, flags, true); err != nil {
				return err
			}
		case ir.OP_marshal_text:
			vt := ins.Vr()
			var err error
			if buf, err = call_text_marshaler(buf, vt, p, flags, false); err != nil {
				return err
			}
		case ir.OP_marshal_text_p:
			vt := ins.Vr()
			var err error
			if buf, err = call_text_marshaler(buf, vt, p, flags, true); err != nil {
				return err
			}
		default:
			panic(fmt.Sprintf("not implement %s at %d", ins.Op().String(), pc))
		}
	}

	*b = buf
	return nil
}

// func to_buf(w unsafe.Pointer, l int, c int) []byte {
// 	return rt.BytesFrom(unsafe.Pointer(uintptr(w)-uintptr(l)), l, c)
// }

// func from_buf(buf []byte) (unsafe.Pointer, int, int) {
// 	return rt.IndexByte(buf, len(buf)), len(buf), cap(buf)
// }

func has_opts(opts uint64, bit int) bool {
	return opts & (1<<bit) != 0
}

func is_nil(p unsafe.Pointer) bool {
	return *(*unsafe.Pointer)(p) == nil
}

func call_text_marshaler(buf []byte, vt *rt.GoType, p unsafe.Pointer, flags uint64, pointer bool) ([]byte, error) {
	var it rt.GoIface
	if !pointer {
		switch vt.Kind() {
			case reflect.Interface        : 
			if is_nil(p) {
				buf = append(buf, "null"...)
				return buf, nil
			}
			it = rt.AssertI2I(_T_encoding_TextMarshaler, *(*rt.GoIface)(p))
			case reflect.Ptr, reflect.Map : it = rt.ConvT2I(vt, p, _T_encoding_TextMarshaler, true)
			default                       : it = rt.ConvT2I(vt, p, _T_encoding_TextMarshaler, !vt.Indirect())
		}

	} else {
		it = rt.ConvT2I(vt, p, _T_encoding_TextMarshaler, false)
	}
	if err := alg.EncodeTextMarshaler(&buf, *(*encoding.TextMarshaler)(unsafe.Pointer(&it)), (flags)); err != nil {
		return buf, err
	}
	return buf, nil
}

func call_json_marshaler(buf []byte, vt *rt.GoType, p unsafe.Pointer, flags uint64, pointer bool) ([]byte, error) {
	var it rt.GoIface
	if !pointer {
		switch vt.Kind() {
			case reflect.Interface        : 
			if is_nil(p) {
				buf = append(buf, "null"...)
				return buf, nil
			}
			it = rt.AssertI2I(_T_json_Marshaler, *(*rt.GoIface)(p))
			case reflect.Ptr, reflect.Map : it = rt.ConvT2I(vt, p, _T_json_Marshaler, true)
			default                       : it = rt.ConvT2I(vt, p, _T_json_Marshaler, !vt.Indirect())
		}

	} else {
		it = rt.ConvT2I(vt, p, _T_json_Marshaler, false)
	}
	if err := alg.EncodeJsonMarshaler(&buf, *(*json.Marshaler)(unsafe.Pointer(&it)), (flags)); err != nil {
		return buf, err
	}
	return buf, nil
}