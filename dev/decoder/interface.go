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

	// not pointer type, or nil pointer, or *interface{}
	if eface.Value == nil || eface.Type.Kind() != reflect.Ptr || rt.PtrElem(eface.Type) == anyType {
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


	dec, err := findOrCompile(rt.PtrElem(eface.Type))
	if err != nil {
		return err
	}

	return dec.FromDom(eface.Value, node, ctx)
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

	return v.(encoding.TextUnmarshaler).UnmarshalText(rt.Str2Mem(s))
}

type unmarshalJSONDecoder struct {
	typ *rt.GoType
}

func (d *unmarshalJSONDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	// not deal with null here
	v := *(*interface{})(unsafe.Pointer(&rt.GoEface{
		Type:  d.typ,
		Value: vp,
	}))

	return v.(json.Unmarshaler).UnmarshalJSON([]byte(node.AsRaw(&ctx.Context)))
}