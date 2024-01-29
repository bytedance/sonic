package decoder

import (
	"encoding"
	"encoding/json"
	"math"
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/bytedance/sonic/dev/internal"
)

/** Decoder for most common map types: map[string]interface{}, map[string]string **/

type mapIfaceDecoder struct {
}

func (d *mapIfaceDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*map[string]interface{})(vp) = nil
		return nil
	}

	if ctx.options&OptionUseNumber == 0 {
		return node.AsMapIface(&ctx.Context, vp)

	}

	return node.AsMapIfaceUseNumber(&ctx.Context, vp)
}

type mapStringDecoder struct {
}

func (d *mapStringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*map[string]string)(vp) = nil
		return nil
	}

	return node.AsMapString(&ctx.Context, vp)
}

/** Decoder for map with string key **/

type mapStrKeyFastDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapStrKeyFastDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		key, err := keyn.AsStr(&ctx.Context)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign_faststr(d.mapType, m, key)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapStrKeyStdDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapStrKeyStdDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		key, err := keyn.AsStr(&ctx.Context)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign(d.mapType, m, unsafe.Pointer(&key))
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			*(*unsafe.Pointer)(vp) = nil
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

/** Decoder for map with int32 or int64 key **/

type mapI32KeyFastDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapI32KeyFastDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		k, err := keyn.AsKeyI64(&ctx.Context)
		if k > math.MaxInt32 || k < math.MinInt32 {
			return error_value(keyn.AsRaw(&ctx.Context), d.mapType.Key.Pack())
		}

		key := int32(k)
		ku32 := *(*uint32)(unsafe.Pointer(&key))
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign_fast32(d.mapType, m, ku32)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}

		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapI32KeyStdDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapI32KeyStdDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		k, err := keyn.AsKeyI64(&ctx.Context)
		if k > math.MaxInt32 || k < math.MinInt32 {
			return error_value(keyn.AsRaw(&ctx.Context), d.mapType.Key.Pack())
		}

		key := int32(k)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign(d.mapType, m, unsafe.Pointer(&key))
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapI64KeyFastDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapI64KeyFastDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		key, err := keyn.AsKeyI64(&ctx.Context)
		ku64 := *(*uint64)(unsafe.Pointer(&key))
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign_fast64(d.mapType, m, ku64)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapI64KeyStdDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapI64KeyStdDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		key, err := keyn.AsKeyI64(&ctx.Context)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign(d.mapType, m, unsafe.Pointer(&key))
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

/** Decoder for map with unt32 or uint64 key **/

type mapU32KeyFastDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapU32KeyFastDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		k, err := keyn.AsKeyU64(&ctx.Context)
		if k > math.MaxUint32 {
			return error_value(keyn.AsRaw(&ctx.Context), d.mapType.Key.Pack())
		}
		key := uint32(k)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign_fast32(d.mapType, m, key)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapU32KeyStdDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapU32KeyStdDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		k, err := keyn.AsKeyU64(&ctx.Context)
		if k > math.MaxUint32 {
			return error_value(keyn.AsRaw(&ctx.Context), d.mapType.Key.Pack())
		}

		key := uint32(k)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign(d.mapType, m, unsafe.Pointer(&key))
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}

		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapU64KeyFastDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapU64KeyFastDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		key, err := internal.NewNode(next).AsKeyU64(&ctx.Context)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign_fast64(d.mapType, m, key)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}
		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

type mapU64KeyStdDecoder struct {
	mapType *rt.GoMapType
	elemDec decFunc
}

func (d *mapU64KeyStdDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		keyn := internal.NewNode(next)
		key, err := keyn.AsKeyU64(&ctx.Context)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		valp := mapassign(d.mapType, m, unsafe.Pointer(&key))
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}

		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}

/** Decoder for generic cases */

type decKey func(dec *mapDecoder, raw string, ctx *context) (interface{}, error)

func decodeKeyU8(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	key, err := internal.Unquote(raw)
	if err != nil {
		return nil, err
	}
	ret, err := internal.ParseU64(key)
	if err != nil {
		return nil, err
	}
	if ret > math.MaxUint8 {
		return nil, error_value(key, dec.mapType.Key.Pack())
	}
	return uint8(ret), nil
}

func decodeKeyU16(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	key, err := internal.Unquote(raw)
	if err != nil {
		return nil, err
	}
	ret, err := internal.ParseU64(key)
	if err != nil {
		return nil, err
	}
	if ret > math.MaxUint16 {
		return nil, error_value(key, dec.mapType.Key.Pack())
	}
	return uint16(ret), nil
}

func decodeKeyI8(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	key, err := internal.Unquote(raw)
	if err != nil {
		return nil, err
	}
	ret, err := internal.ParseI64(key)
	if err != nil {
		return nil, err
	}
	if ret > math.MaxInt8 || ret < math.MinInt8 {
		return nil, error_value(key, dec.mapType.Key.Pack())
	}
	return int8(ret), nil
}

func decodeKeyI16(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	key, err := internal.Unquote(raw)
	if err != nil {
		return nil, err
	}
	ret, err := internal.ParseI64(key)
	if err != nil {
		return nil, err
	}
	if ret > math.MaxInt16 || ret < math.MinInt16 {
		return nil, error_value(key, dec.mapType.Key.Pack())
	}
	return int16(ret), nil
}

func decodeKeyJSONUnmarshaler(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	ret := reflect.New(dec.mapType.Key.Pack()).Interface()
	err := ret.(json.Unmarshaler).UnmarshalJSON([]byte(raw))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func decodeKeyTextUnmarshaler(dec *mapDecoder, raw string, ctx *context) (interface{}, error) {
	key, err := internal.Unquote(raw)
	if err != nil {
		return nil, err
	}
	ret := reflect.New(dec.mapType.Key.Pack()).Interface()
	err = ret.(encoding.TextUnmarshaler).UnmarshalText([]byte(key))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type mapDecoder struct {
	mapType *rt.GoMapType
	keyDec  decKey
	elemDec decFunc
}

func (d *mapDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	obj, err := node.AsObj()
	if err != nil {
		return err
	}

	// allocate map
	m := *(*unsafe.Pointer)(vp)
	if m == nil {
		m = makemap(&d.mapType.GoType, obj.Len())
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		raw := internal.NewNode(next).AsRaw(&ctx.Context)
		key, err := d.keyDec(d, raw, ctx)
		if err != nil {
			return err
		}

		valn := internal.NewNode(internal.PtrOffset(next, 1))
		keyp := rt.UnpackEface(key).Value
		valp := mapassign(d.mapType, m, keyp)
		err = d.elemDec.FromDom(valp, valn, ctx)
		if err != nil {
			return err
		}

		next = valn.Next()
	}

	*(*unsafe.Pointer)(vp) = m
	return nil
}
