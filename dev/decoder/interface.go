package decoder

import (
	"encoding"
	"encoding/json"
	"unsafe"
	"reflect"

	"github.com/bytedance/sonic/dev/internal"
	"github.com/bytedance/sonic/dev/internal/rt"
)

type efaceDecoder struct {
}

func (d *efaceDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	eface := *(*rt.GoEface)(vp)
	etp := rt.PtrElem(eface.Type)
	vp = eface.Value

	/* check the defined pointer type for issue 379 */
	if eface.Type.IsNamed() {
		newp := vp
		etp = eface.Type
		vp = unsafe.Pointer(&newp)
	}

	// not pointer type, or nil pointer, or *interface{}
	if eface.Value == nil || eface.Type.Kind() != reflect.Ptr ||  etp == anyType {
		var ret interface{}
		var err error
	
		if Options(ctx.Options)&OptionUseNumber == 0 &&  Options(ctx.Options)&OptionUseInt64 == 0 {
			ret, err = node.AsEface(&ctx.Context)
		} else if  Options(ctx.Options)&OptionUseNumber != 0 {
			ret, err = node.AsEfaceUseNumber(&ctx.Context)
		} else {
			ret, err = node.AsEfaceUseInt64(&ctx.Context)
		}

		if err != nil {
			return err
		}
	
		*(*interface{})(vp) = ret
		return nil
	}


	dec, err := findOrCompile(etp)
	if err != nil {
		return err
	}

	return dec.FromDom(vp, node, ctx)
}

type ifaceDecoder struct {
	typ *rt.GoType
}

func (d *ifaceDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	iface := *(*rt.GoIface)(vp)
	vt := iface.Itab.Vt
	etp := rt.PtrElem(iface.Itab.Vt)
	vp = iface.Value

	/* check the defined pointer type for issue 379 */
	if vt.IsNamed() {
		newp := vp
		etp = vt
		vp = unsafe.Pointer(&newp)
	}

	// not pointer type, or nil pointer, or *interface{}
	if vp == nil || vt.Kind() != reflect.Ptr ||  etp == anyType {
		var ret interface{}
		var err error
	
		if Options(ctx.Options)&OptionUseNumber == 0 &&  Options(ctx.Options)&OptionUseInt64 == 0 {
			ret, err = node.AsEface(&ctx.Context)
		} else if  Options(ctx.Options)&OptionUseNumber != 0 {
			ret, err = node.AsEfaceUseNumber(&ctx.Context)
		} else {
			ret, err = node.AsEfaceUseInt64(&ctx.Context)
		}

		if err != nil {
			return err
		}
	
		*(*interface{})(vp) = ret
		return nil
	}


	dec, err := findOrCompile(etp)
	if err != nil {
		return err
	}

	return dec.FromDom(vp, node, ctx)
}

type unmarshalTextDecoder struct {
	typ *rt.GoType
}

func (d *unmarshalTextDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	s, err := node.AsStringCopy()
	if err != nil {
		return err
	}

	v := *(*interface{})(unsafe.Pointer(&rt.GoEface{
		Type:  d.typ,
		Value: vp,
	}))

	// fast path
	if u, ok :=  v.(encoding.TextUnmarshaler); ok {
		return u.UnmarshalText(rt.Str2Mem(s))
	}

	// slow path
	rv := reflect.ValueOf(v)
	if u, ok := rv.Interface().(encoding.TextUnmarshaler); ok {
		return u.UnmarshalText(rt.Str2Mem(s))
	}

	return error_type(d.typ)
}

type unmarshalJSONDecoder struct {
	typ *rt.GoType
}

func (d *unmarshalJSONDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	v := *(*interface{})(unsafe.Pointer(&rt.GoEface{
		Type: d.typ,
		Value: vp,
	}))

	// fast path
	if u, ok :=  v.(json.Unmarshaler); ok {
		return u.UnmarshalJSON([]byte(node.AsRaw(&ctx.Context)))
	}

	// slow path
	rv := reflect.ValueOf(v)
	if u, ok := rv.Interface().(json.Unmarshaler); ok {
		return u.UnmarshalJSON([]byte(node.AsRaw(&ctx.Context)))
	}

	return error_type(d.typ)
}