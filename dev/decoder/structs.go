package decoder

import (
	"unsafe"

	"github.com/bytedance/sonic/internal/resolver"
	"github.com/bytedance/sonic/dev/internal"
	"github.com/bytedance/sonic/dev/internal/caching"
)

type fieldEntry struct {
	resolver.FieldMeta
	fieldDec decFunc
}

type structDecoder struct {
	fieldMap   *caching.FieldMap
	fields     []fieldEntry
	structName string
}

func (d *structDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	var gerr error
	obj, err := node.AsObj()
	if err != nil {
		return nil
	}

	next := obj.Children()
	for i := 0; i < obj.Len(); i++ {
		key, err := internal.NewNode(next).AsStr(&ctx.Context)
		val := internal.NewNode(internal.PtrOffset(next, 1))
		next = val.Next()
		if err != nil {
			return nil
		}

		idx := d.fieldMap.TryGet(key, i)
		if idx == -1 {
			idx = d.fieldMap.Get(key)
		}
		if idx == -1 {
			idx = d.fieldMap.GetCaseInsensitive(key)
		}
        if idx == -1 {
            if ctx.options&OptionDisableUnknown != 0 {
                return error_field(key)
            }
            continue
        }

		offset := d.fields[idx].Path[0].Size
		elem := unsafe.Pointer(uintptr(vp) + offset)
		err = d.fields[idx].fieldDec.FromDom(elem, val, ctx)

		// deal with mismatch type errors
		if err != nil {
			gerr = error_mismatch_internal(err, d.fields[idx].Type, ctx.Json)
			continue
		}
	}
	return gerr
}
