package decoder

import (
	"fmt"
	"reflect"

	"github.com/bytedance/sonic/internal/resolver"
	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/bytedance/sonic/dev/internal/caching"
)

func (c *compiler) compileIntStringOption(vt reflect.Type) decFunc {
	switch vt.Size() {
	case 4:
		switch vt.Kind() {
		case reflect.Uint:
			fallthrough
		case reflect.Uintptr:
			return &u32StringDecoder{}
		case reflect.Int:
			return &i32StringDecoder{}
		}
	case 8:
		switch vt.Kind() {
		case reflect.Uint:
			fallthrough
		case reflect.Uintptr:
			return &u64StringDecoder{}
		case reflect.Int:
			return &i64StringDecoder{}
		}
	default:
		panic("not supported pointer size: " + fmt.Sprint(vt.Size()))
	}
	panic("unreachable")
}

func (c *compiler) compileFieldStringOption(vt reflect.Type) decFunc {
	switch vt.Kind() {
	case reflect.String:
		if vt == jsonNumberType {
			return &numberStringDecoder{}
		}
		return &strStringDecoder{}
	case reflect.Bool:
		return &boolStringDecoder{}
	case reflect.Int8:
		return &i8StringDecoder{}
	case reflect.Int16:
		return &i16StringDecoder{}
	case reflect.Int32:
		return &i32StringDecoder{}
	case reflect.Int64:
		return &i64StringDecoder{}
	case reflect.Uint8:
		return &u8StringDecoder{}
	case reflect.Uint16:
		return &u16StringDecoder{}
	case reflect.Uint32:
		return &u32StringDecoder{}
	case reflect.Uint64:
		return &u64StringDecoder{}
	case reflect.Float32:
		return &f32StringDecoder{}
	case reflect.Float64:
		return &f64StringDecoder{}
	case reflect.Uint:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Int:
		return c.compileIntStringOption(vt)
	case reflect.Pointer:
		return &ptrStrDecoder{
			typ:   rt.UnpackType(vt.Elem()),
			deref: c.compileFieldStringOption(vt.Elem()),
		}
	default:
		panic("string options should appliy only to fields of string, floating point, integer, or boolean types.")
	}
}

func (c *compiler) compileStruct(vt reflect.Type) decFunc {
	c.enter(vt)
	defer c.exit(vt)
	fv := resolver.ResolveStruct(vt)
	fm := caching.CreateFieldMap(len(fv))

	entries := make([]fieldEntry, 0, len(fv))
	for i, f := range fv {
		fm.Set(f.Name, i)

		var dec decFunc
		/* dealt with field tag options */
		if f.Opts&resolver.F_stringize != 0 {
			dec = c.compileFieldStringOption(f.Type)
		} else {
			dec = c.compile(f.Type)
		}

		/* deal with embedded pointer fields */
		var derefTyps []*rt.GoType
		var edecOffset uintptr = 0
		for _, off := range f.Path {
			if off.Kind == resolver.F_deref {
				derefTyps = append(derefTyps, rt.UnpackType(off.Type))
			} else {
				edecOffset = off.Size
			}
		}
		if derefTyps != nil {
			dec = &embeddedFieldPtrDecoder{
				derefTypes: derefTyps,
				offset:     edecOffset,
				fieldDec:   dec,
				fieldName:  f.Name,
			}
		}

		entries = append(entries, fieldEntry{
			FieldMeta: f,
			fieldDec:  dec,
		})
	}
	return &structDecoder{
		fieldMap:   fm,
		fields:     entries,
		structName: vt.Name(),
	}
}
