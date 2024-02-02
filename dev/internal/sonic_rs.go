package internal

//#cgo linux,arm64 LDFLAGS:  -L ../rs_wrapper/lib/linux -lsonic_rs_aarch64-unknown-linux-gnu
//#cgo linux,amd64 LDFLAGS:  -L ../rs_wrapper/lib/linux -lsonic_rs_x86_64-unknown-linux-gnu
//#cgo darwin,arm64 LDFLAGS: -L ../rs_wrapper/lib/darwin -lsonic_rs_aarch64-apple-darwin
//#cgo darwin,amd64 LDFLAGS: -L ../rs_wrapper/lib/darwin -lsonic_rs_x86_64-apple-darwin
//
//#include "../rs_wrapper/include/sonic_rs.h"
import "C"
import (
	"encoding/json"
	"math"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
	_ "github.com/davecgh/go-spew/spew"
)

type Context struct {
	Options   uint64
	Json  string
	Start uintptr
	Dom   Dom
}

func NewContext(json string, opts uint64) (Context, error) {
	dom, err := Parse(json, opts)
	if err != nil {
		return Context{}, err
	}
	return Context{
		Options:   opts,
		Json:  json,
		Start: dom.StrStart(),
		Dom:   dom,
	}, nil
}

func (ctx *Context) Delete() {
	ctx.Dom.Delete()
}

type Node struct {
	cptr *C.Node
}

func NewNode(cptr *C.Node) Node {
	return Node{cptr: cptr}
}

type Dom struct {
	cdom C.Document
}

func (dom *Dom) Root() Node {
	return Node{cptr: dom.cdom.node}
}

func (dom *Dom) HasUtf8Lossy() bool {
	return bool(dom.cdom.has_utf8_lossy)
}

func (dom *Dom) CopyJsonString() string {
	cptr := (*C.char)(unsafe.Pointer(dom.cdom.str_buf))
	len := C.int(dom.cdom.str_len);
	return C.GoStringN(cptr, len);
}

func (dom *Dom) StrStart() uintptr {
	return uintptr(unsafe.Pointer(dom.cdom.str_buf))
}

func (dom *Dom) Delete() {
	C.sonic_rs_ffi_free(dom.cdom.dom, dom.cdom.str_buf, dom.cdom.error_msg_cap)
}

type Array struct {
	cptr *C.Node
}

type Object struct {
	cptr *C.Node
}

func (obj Object) Len() int {
	cobj := (*C.Object)(unsafe.Pointer(obj.cptr))
	return int(uint64(cobj.len) & ConLenMask)
}

func (arr Array) Len() int {
	carr := (*C.Array)(unsafe.Pointer(arr.cptr))
	return int(uint64(carr.len) & ConLenMask)
}

func Parse(data string, opt uint64) (Dom, error) {
	var s = (*reflect.StringHeader)((unsafe.Pointer)(&data))
	cdom := C.sonic_rs_ffi_parse((*C.char)(unsafe.Pointer((s.Data))), C.size_t(s.Len), C.uint64_t(opt))
	runtime.KeepAlive(data)
	ret := Dom{
		cdom: cdom,
	}

	// parse error
	if offset := int64(cdom.error_offset); offset != -1 {
		msg := C.GoStringN(cdom.error_msg, C.int(cdom.error_msg_len))
		err := newError(msg, offset)
		ret.Delete()
		return ret, err
	}

	return ret, nil
}

// / Helper functions to eliminate CGO calls
func (val Node) Type() uint8 {
	ctype := (*C.Type)(unsafe.Pointer(val.cptr))
	return uint8(ctype.t)
}

func (val Node) Next() *C.Node {
	if val.Type() != KObject && val.Type() != KArray {
		return PtrOffset(val.cptr, 1)
	}
	cobj := (*C.Object)(unsafe.Pointer(val.cptr))
	offset := int(uint64(cobj.len) >> ConLenBits)
	return PtrOffset(val.cptr, uintptr(offset))
}

func (val Node) U64() uint64 {
	cnum := (*C.Number)(unsafe.Pointer(val.cptr))
	return *(*uint64)((unsafe.Pointer)(&(cnum.num)))
}

func (val Node) I64() int64 {
	cnum := (*C.Number)(unsafe.Pointer(val.cptr))
	return *(*int64)((unsafe.Pointer)(&(cnum.num)))
}

func (val Node) IsNull() bool {
	ctype := (*C.Type)(unsafe.Pointer(val.cptr))
	return ctype.t == KNull
}

func (val Node) IsNumber() bool {
	ctype := (*C.Type)(unsafe.Pointer(val.cptr))
	return ctype.t&KNumber != 0
}

func (val Node) F64() float64 {
	cnum := (*C.Number)(unsafe.Pointer(val.cptr))
	return *(*float64)((unsafe.Pointer)(&(cnum.num)))
}

func (val Node) Bool() bool {
	ctype := (*C.Type)(unsafe.Pointer(val.cptr))
	return ctype.t == KTrue
}

func (val Node) AsU64() (uint64, error) {
	if val.Type() != KUint {
		return 0, newUnmatched("expect uint64")
	}

	cnum := (*C.Number)(unsafe.Pointer(val.cptr))
	return *(*uint64)((unsafe.Pointer)(&(cnum.num))), nil
}

func (val *Node) AsObj() (Object, error) {
	var ret Object
	if val.Type() != KObject {
		return ret, newUnmatched("expect object")
	}
	return Object{
		cptr: val.cptr,
	}, nil
}

func (val Node) Obj() Object {
	return Object{cptr: val.cptr}
}

func (val Node) Arr() Array {
	return Array{cptr: val.cptr}
}

func (val *Node) AsArr() (Array, error) {
	var ret Array
	if val.Type() != KArray {
		return ret, newUnmatched("expect array")
	}
	return Array{
		cptr: val.cptr,
	}, nil
}

func (val Node) AsI64() (int64, error) {
	if val.Type() == KUint && val.U64() <= math.MaxInt64 {
		return int64(val.U64()), nil
	}

	if val.Type() == KSint {
		return val.I64(), nil
	}

	return 0, newUnmatched("expect int64")
}

func (val Node) ParseI64(ctx *Context) (int64, error) {
	s, err := val.AsStrRef(ctx)
	if err != nil {
		return 0, err
	}

	i, err := ParseI64(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (val Node) ParseBool(ctx *Context) (bool, error) {
	s, err := val.AsStrRef(ctx)
	if err != nil {
		return false, err
	}

	b, err := ParseBool(s)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (val Node) ParseU64(ctx *Context) (uint64, error) {
	s, err := val.AsStrRef(ctx)
	if err != nil {
		return 0, err
	}

	i, err := ParseU64(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (val Node) ParseF64(ctx *Context) (float64, error) {
	s, err := val.AsStrRef(ctx)
	if err != nil {
		return 0, err
	}

	i, err := ParseF64(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (val Node) AsF64() (float64, error) {
	if val.Type() == KUint {
		return float64(val.U64()), nil
	}
	if val.Type() == KSint {
		return float64(val.I64()), nil
	}
	if val.Type() == KReal {
		return float64(val.F64()), nil
	}
	return 0, newUnmatched("expect float64")
}

func (val Node) AsBool() (bool, error) {
	if val.Type() == KTrue {
		return true, nil
	}
	if val.Type() == KFalse {
		return false, nil
	}
	return false, newUnmatched("expect bool")
}

func (val Node) AsStr(ctx *Context) (string, error) {
	if !val.IsStr() {
		return "", newUnmatched("expect string")
	}

	if (ctx.Options & (1 << _F_copy_string) == 0) && val.Type() == KStringCommon {
		return val.String(ctx), nil
	}

	return val.StringCopy(), nil
}


func (val Node) AsStrRef(ctx *Context) (string, error) {
	switch val.Type() {
	case KStringHasEscaped:
		return val.StringCopy(), nil
	case KStringCommon:
		return val.String(ctx), nil
	default:
		return "", newUnmatched("expect string")
	}
}

func (val Node) AsStringCopy() (string, error) {
	if !val.IsStr() {
		return "", newUnmatched("expect string")
	}
	cstr := (*C.String)(unsafe.Pointer(val.cptr))
	len := cstr.len >> PosBits
	ret := C.GoStringN(cstr.p, C.int(len))
	return ret, nil
}

func (val Node) IsStr() bool {
	return (val.Type() == KStringCommon) || (val.Type() == KStringHasEscaped)
}

func (val Node) IsRawNumber() bool {
	return val.Type() == KRawNumber
}

func (val Node) Number(ctx *Context) json.Number {
	cnum := (*C.RawNumber)(unsafe.Pointer(val.cptr))
	len := int(cnum.len >> PosBits)
	offset := int(uintptr(unsafe.Pointer(cnum.p)) - ctx.Start)
	ref := rt.Str2Mem(ctx.Json)[offset:int(offset+len)]
	return json.Number(rt.Mem2Str(ref))
}

func (val Node) Position(ctx *Context) int {
	if !val.IsStr() && !val.IsRawNumber() {
		cstr := (*C.String)(unsafe.Pointer(val.cptr))
		return int(cstr.len >> PosBits)
	} else {
		cstr := (*C.String)(unsafe.Pointer(val.cptr))
		ret := int(uintptr(unsafe.Pointer(cstr.p)) - ctx.Start - 1)
		return ret
	}
}

func (val Node) AsNumber(ctx *Context) (json.Number, error) {
	if val.IsStr() {
		s, _ := val.AsStr(ctx)
		err := ValidNumberFast(s)
		if err != nil {
			return "", err
		}
		return json.Number(s), nil
	}

	if !val.IsNumber() {
		return json.Number(""), newUnmatched("expect number")
	}

	if val.IsRawNumber() {
		return val.Number(ctx), nil
	}

	start := val.Position(ctx)
	end, err := SkipNumberFast(ctx.Json, start)
	if err != nil {
		return "", err
	}
	return json.Number(ctx.Json[start:end]), nil
}

func (val Node) AsRaw(ctx *Context) string {
	// fast path for unescaped strings
	switch val.Type() {
	case KNull:
		return "null"
	case KTrue:
		return "true"
	case KFalse:
		return "false"
	case KStringCommon:
		cstr := (*C.String)(unsafe.Pointer(val.cptr))
		len := int(cstr.len >> PosBits)
		offset := int(uintptr(unsafe.Pointer(cstr.p)) - ctx.Start)
		// add start abd end quote
		ref := rt.Str2Mem(ctx.Json)[offset-1 : int(offset+len)+1]
		return rt.Mem2Str(ref)
	default:
		raw, err := SkipOneFast(ctx.Json, val.Position(ctx))
		if err != nil {
			panic("should always be valid json here")
		}
		return raw
	}
}

// reference from the input JSON as possible
func (val Node) String(ctx *Context) string {
	cstr := (*C.String)(unsafe.Pointer(val.cptr))
	len := int(cstr.len >> PosBits)
	offset := int(uintptr(unsafe.Pointer(cstr.p)) - ctx.Start)
	ref := rt.Str2Mem(ctx.Json)[offset:int(offset+len)]
	return rt.Mem2Str(ref)
}

func (val Node) StringCopy() string {
	cstr := (*C.String)(unsafe.Pointer(val.cptr))
	len := cstr.len >> PosBits
	ret := C.GoStringN(cstr.p, C.int(len))
	return ret
}

func (val Node) Object() Object {
	return Object{cptr: val.cptr}
}

func (val Node) Array() Array {
	return Array{cptr: val.cptr}
}

func (val *Array) Children() *C.Node {
	return PtrOffset(val.cptr, 1)
}

func (val *Object) Children() *C.Node {
	return PtrOffset(val.cptr, 1)
}

func (val *Node) Equal(lhs string) bool {
	// check whether escaped
	cstr := (*C.String)(unsafe.Pointer(val.cptr))
	len := int(cstr.len >> PosBits)

	gos := rt.GoString{
		Ptr: unsafe.Pointer(cstr.p),
		Len: len,
	}

	// TODO: FIXME: maybe bad pointer here
	s := *((*string)(unsafe.Pointer(&gos)))
	return lhs == s
}

func (node *Node) AsMapEface(ctx *Context, vp unsafe.Pointer) error {
	if node.IsNull() {
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	size := obj.Len()
	var m map[string]interface{}
	if *(*unsafe.Pointer)(vp) == nil {
		m = make(map[string]interface{}, size)
	} else {
		m = *(*map[string]interface{})(vp)
	}

	var gerr error
	next := obj.Children()
	for i := 0; i < size; i++ {
		knode := NewNode(next)
		key, err := knode.AsStr(ctx)
		if err != nil {
			return err
		}

		val := NewNode(PtrOffset(next, 1))
		m[key], err = val.AsEface(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
		next = val.cptr
	}

	*(*map[string]interface{})(vp) = m
	return gerr
}

func (node *Node) AsMapEfaceUseNumber(ctx *Context, vp unsafe.Pointer) error {
	obj, err := node.AsObj()
	if err != nil {
		return nil
	}

	size := obj.Len()

	var m map[string]interface{}
	if *(*unsafe.Pointer)(vp) == nil {
		m = make(map[string]interface{}, size)
	} else {
		m = *(*map[string]interface{})(vp)
	}

	var gerr error
	*node = NewNode(obj.Children())
	for i := 0; i < size; i++ {
		key, err := node.AsStr(ctx)
		if err != nil {
			return err
		}

		*node = NewNode(PtrOffset(node.cptr, 1))
		m[key], err = node.AsEfaceUseNumber(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
	}

	*(*map[string]interface{})(vp) = m
	return gerr
}

func (node *Node) AsMapEfaceUseInt64(ctx *Context, vp unsafe.Pointer) error {
	obj, err := node.AsObj()
	if err != nil {
		return nil
	}

	size := obj.Len()

	var m map[string]interface{}
	var gerr error
	if *(*unsafe.Pointer)(vp) == nil {
		m = make(map[string]interface{}, size)
	} else {
		m = *(*map[string]interface{})(vp)
	}

	*node = NewNode(obj.Children())
	for i := 0; i < size; i++ {
		key, err := node.AsStr(ctx)
		if err != nil {
			return err
		}

		*node = NewNode(PtrOffset(node.cptr, 1))
		m[key], err = node.AsEfaceUseInt64(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
	}

	*(*map[string]interface{})(vp) = m
	return gerr
}

func (node *Node) AsMapString(ctx *Context, vp unsafe.Pointer) error {
	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	size := obj.Len()

	var m map[string]string
	if *(*unsafe.Pointer)(vp) == nil {
		m = make(map[string]string, size)
	} else {
		m = *(*map[string]string)(vp)
	}

	next := obj.Children()
	var gerr error
	for i := 0; i < size; i++ {
		knode := NewNode(next)
		key, err := knode.AsStr(ctx)
		if err != nil {
			return err
		}

		val := NewNode(PtrOffset(next, 1))
		m[key], err = val.AsStr(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
		next = val.Next()
	}

	*(*map[string]string)(vp) = m
	return gerr
}

func (node *Node) AsSliceEface(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()

	s := *(*[]interface{})((unsafe.Pointer)(MakeSlice(vp, anyType, size)))
	next := arr.Children()

	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)
		s[i], err = val.AsEface(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
		next = val.cptr
	}

	*(*[]interface{})(vp) = s
	return gerr
}

func (node *Node) AsSliceEfaceUseNumber(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()

	s := *(*[]interface{})((unsafe.Pointer)(MakeSlice(vp, anyType, size)))
	*node = NewNode(arr.Children())

	var gerr error
	for i := 0; i < size; i++ {
		s[i], err = node.AsEfaceUseNumber(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
	}

	*(*[]interface{})(vp) = s
	return gerr
}

func (node *Node) AsSliceEfaceUseInt64(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()

	s := *(*[]interface{})((unsafe.Pointer)(MakeSlice(vp, anyType, size)))
	*node = NewNode(arr.Children())

	var gerr error
	for i := 0; i < size; i++ {
		s[i], err = node.AsEfaceUseInt64(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
	}

	*(*[]interface{})(vp) = s
	return gerr
}

func (node *Node) AsSliceI32(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()

	s := *(*[]int32)((unsafe.Pointer)(MakeSlice(vp, int32Type, size)))
	next := arr.Children()

	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)
		ret, err := val.AsI64()
		if gerr == nil && err != nil {
			gerr = err
		}

		if ret > math.MaxInt32 || ret < math.MinInt32 {
			if gerr == nil {
				gerr = newUnmatched("expect int32")
			}
			ret = 0
		}

		next = val.Next()
		s[i] = int32(ret)
	}

	*(*[]int32)(vp) = s
	return gerr
}

func (node *Node) AsSliceI64(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()
	next := arr.Children()

	s := *(*[]int64)((unsafe.Pointer)(MakeSlice(vp, int64Type, size)))

	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)

		ret, err := val.AsI64()
		if gerr == nil && err != nil {
			gerr = err
		}

		s[i] = ret
		next = val.Next()
	}

	*(*[]int64)(vp) = s
	return gerr
}

func (node *Node) AsSliceU32(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()
	next := arr.Children()

	s := *(*[]uint32)((unsafe.Pointer)(MakeSlice(vp, uint32Type, size)))
	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)
		ret, err := val.AsU64()
		if gerr == nil && err != nil {
			gerr = err
		}

		if ret > math.MaxUint32 {
			if gerr == nil {
				gerr =  newUnmatched("expect uint32")
			}
			ret = 0
		}

		s[i] = uint32(ret)
		next = val.Next()
	}

	*(*[]uint32)(vp) = s
	return gerr
}

func (node *Node) AsSliceU64(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	size := arr.Len()
	next := arr.Children()

	s := *(*[]uint64)((unsafe.Pointer)(MakeSlice(vp, uint64Type, size)))
	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)

		ret, err := val.AsU64()
		if gerr == nil && err != nil {
			gerr = err
		}

		s[i] = ret
		next = val.Next()
	}

	*(*[]uint64)(vp) = s
	return gerr
}

func (node *Node) AsSliceString(ctx *Context, vp unsafe.Pointer) error {
	arr, err := node.AsArr()
	if err != nil {
		return err
	}
	size := arr.Len()
	next := arr.Children()

	s := *(*[]string)((unsafe.Pointer)(MakeSlice(vp, strType, size)))
	var gerr error
	for i := 0; i < size; i++ {
		val := NewNode(next)

		ret, err := val.AsStr(ctx)
		if gerr == nil && err != nil {
			gerr = err
		}

		s[i] = ret
		next = val.Next()
	}

	*(*[]string)(vp) = s
	return gerr
}

func (node *Node) AsSliceBytes(ctx *Context) ([]byte, error) {
	s, err := node.AsStrRef(ctx)
	if err != nil {
		return nil, err
	}

	b64, err := decodeBase64(s)
	if err != nil {
		return nil, err
	}

	return b64, nil
}

func (node *Node) AsEface(ctx *Context) (interface{}, error) {
	switch node.Type() {
	case KObject:
		obj := node.Object()
		size := obj.Len()
		m := make(map[string]interface{}, size)

		*node = NewNode(obj.Children())
		var key string
		var err error

		for i := 0; i < size; i++ {
			switch node.Type() {
			case KStringHasEscaped:
				key = node.StringCopy()
			case KStringCommon:
				key = node.String(ctx)
			default:
				return nil, newUnmatched("expect string")
			}
			*node = NewNode(PtrOffset(node.cptr, 1))
			m[key], err = node.AsEface(ctx)
			if err != nil {
				return nil, err
			}
		}
		return m, nil
	case KArray:
		arr := node.Array()
		size := arr.Len()
		garr := make([]interface{}, size)
		*node = NewNode(arr.Children())
		var err error
		for i := 0; i < size; i++ {
			garr[i], err = node.AsEface(ctx)
			if err != nil {
				return nil, err
			}
		}
		return garr, nil
	case KStringCommon:
		str := node.String(ctx)
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KStringHasEscaped:
		str := node.StringCopy()
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KUint:
		f := (float64)(node.U64())
		*node = NewNode(PtrOffset(node.cptr, 1))
		return f, nil
	case KReal:
		f := (float64)(node.F64())
		*node = NewNode(PtrOffset(node.cptr, 1))
		return f, nil
	case KSint:
		f := (float64)(node.I64())
		*node = NewNode(PtrOffset(node.cptr, 1))
		return f, nil
	case KTrue:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return true, nil
	case KFalse:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return false, nil
	case KNull:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return nil, nil
	default:
		return nil, newUnmatched("invalid JSON type")
	}
}

func (node *Node) AsEfaceUseNumber(ctx *Context) (interface{}, error) {
	switch node.Type() {
	case KObject:
		obj := node.Object()
		size := obj.Len()
		m := make(map[string]interface{}, size)
		*node = NewNode(obj.Children())

		var key string
		var err error
		for i := 0; i < size; i++ {
			switch node.Type() {
			case KStringHasEscaped:
				key = node.StringCopy()
			case KStringCommon:
				key = node.String(ctx)
			default:
				return nil, newUnmatched("expect string")
			}

			*node = NewNode(PtrOffset(node.cptr, 1))
			m[key], err = node.AsEfaceUseNumber(ctx)
			if err != nil {
				return nil, err
			}
		}
		return m, nil
	case KArray:
		arr := node.Array()
		size := arr.Len()
		garr := make([]interface{}, size)
		*node = NewNode(arr.Children())

		var err error
		for i := 0; i < size; i++ {
			garr[i], err = node.AsEfaceUseNumber(ctx)
			if err != nil {
				return nil, err
			}
		}
		return garr, nil
	case KStringCommon:
		str := node.String(ctx)
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KStringHasEscaped:
		str := node.StringCopy()
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KRawNumber:
		num := node.Number(ctx)
		*node = NewNode(PtrOffset(node.cptr, 1))
		return num, nil
	case KTrue:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return true, nil
	case KFalse:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return false, nil
	case KNull:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return nil, nil
	default:
		num, err := node.AsNumber(ctx)
		if err != nil {
			return nil, err
		}

		*node = NewNode(PtrOffset(node.cptr, 1))
		return num, nil
	}
}

func (node *Node) AsEfaceUseInt64(ctx *Context) (interface{}, error) {
	switch node.Type() {
	case KObject:
		{
			obj := node.Object()
			size := obj.Len()
			m := make(map[string]interface{}, size)
			*node = NewNode(obj.Children())

			var key string
			var err error
			for i := 0; i < size; i++ {
				switch node.Type() {
				case KStringHasEscaped:
					key = node.StringCopy()
				case KStringCommon:
					key = node.String(ctx)
				default:
					return nil, newUnmatched("expect string")
				}

				*node = NewNode(PtrOffset(node.cptr, 1))
				m[key], err = node.AsEfaceUseInt64(ctx)
				if err != nil {
					return nil, err
				}
			}
			return m, nil
		}
	case KArray:
		arr := node.Array()
		size := arr.Len()
		garr := make([]interface{}, size)
		*node = NewNode(arr.Children())
		var err error
		for i := 0; i < size; i++ {
			garr[i], err = node.AsEfaceUseInt64(ctx)
			if err != nil {
				return nil, err
			}
		}
		return garr, nil
	case KStringCommon:
		str := node.String(ctx)
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KStringHasEscaped:
		str := node.StringCopy()
		*node = NewNode(PtrOffset(node.cptr, 1))
		return str, nil
	case KTrue:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return true, nil
	case KFalse:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return false, nil
	case KNull:
		*node = NewNode(PtrOffset(node.cptr, 1))
		return nil, nil
	default:
		val, err := node.AsI64()
		*node = NewNode(PtrOffset(node.cptr, 1))
		return val, err
	}
}

func PtrOffset(ptr *C.Node, off uintptr) *C.Node {
	uptr := uintptr(unsafe.Pointer(ptr))
	uptr += off * unsafe.Sizeof(C.Node{})
	return (*C.Node)(unsafe.Pointer(uptr))
}
