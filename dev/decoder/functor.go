package decoder

import (
	"encoding/json"
	"math"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/bytedance/sonic/dev/internal"
)

type decFunc interface {
	FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error
}

type ptrDecoder struct {
	typ   *rt.GoType
	deref decFunc
}

// Pointer Value is allocated in the Caller
func (d *ptrDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	if *(*unsafe.Pointer)(vp) == nil {
		*(*unsafe.Pointer)(vp) = mallocgc(d.typ.Size, d.typ, true)
	}

	err := d.deref.FromDom(*(*unsafe.Pointer)(vp), node, ctx)
	if err != nil {
		*(*unsafe.Pointer)(vp) = nil
		return err
	}
	return nil
}

type embeddedFieldPtrDecoder struct {
	derefTypes []*rt.GoType
	offset     uintptr
	fieldDec   decFunc
	fieldName  string
}

// Pointer Value is allocated in the Caller
func (d *embeddedFieldPtrDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	for _, tp := range d.derefTypes {
		if *(*unsafe.Pointer)(vp) == nil {
			*(*unsafe.Pointer)(vp) = mallocgc(tp.Size, tp, true)
		}
		vp = *(*unsafe.Pointer)(vp)
	}

	vp = unsafe.Pointer(uintptr(vp) + d.offset)
	err := d.fieldDec.FromDom(vp, node, ctx)
	if err != nil {
		return err
	}
	return nil
}

type i8Decoder struct{}

func (d *i8Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsI64()
	if err != nil {
		return err
	}

	if ret > math.MaxInt8 || ret < math.MinInt8 {
		return error_value(node.AsRaw(&ctx.Context), int8Type)
	}

	*(*int8)(vp) = int8(ret)
	return nil
}

type i16Decoder struct{}

func (d *i16Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsI64()
	if err != nil {
		return err
	}

	if ret > math.MaxInt16 || ret < math.MinInt16 {
		return error_value(node.AsRaw(&ctx.Context), int16Type)
	}

	*(*int16)(vp) = int16(ret)
	return nil
}

type i32Decoder struct{}

func (d *i32Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsI64()
	if err != nil {
		return err
	}

	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return error_value(node.AsRaw(&ctx.Context), int32Type)
	}

	*(*int32)(vp) = int32(ret)
	return nil
}

type i64Decoder struct{}

func (d *i64Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsI64()
	if err != nil {
		return err
	}

	*(*int64)(vp) = int64(ret)
	return nil
}

type u8Decoder struct{}

func (d *u8Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsU64()
	if err != nil {
		return err
	}

	if ret > math.MaxUint8 {
		return error_value(node.AsRaw(&ctx.Context), uint8Type)
	}
	*(*uint8)(vp) = uint8(ret)
	return nil
}

type u16Decoder struct{}

func (d *u16Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsU64()
	if err != nil {
		return err
	}
	if ret > math.MaxUint16 {
		return error_value(node.AsRaw(&ctx.Context), uint16Type)
	}
	*(*uint16)(vp) = uint16(ret)
	return nil
}

type u32Decoder struct{}

func (d *u32Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsU64()
	if err != nil {
		return err
	}
	if ret > math.MaxUint32 {
		return error_value(node.AsRaw(&ctx.Context), uint32Type)
	}
	*(*uint32)(vp) = uint32(ret)
	return nil
}

type u64Decoder struct{}

func (d *u64Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsU64()
	if err != nil {
		return err
	}
	*(*uint64)(vp) = uint64(ret)
	return nil
}

type f32Decoder struct{}

func (d *f32Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsF64()
	if err != nil {
		return err
	}

	if ret > math.MaxFloat32 || ret < -math.MaxFloat32 {
		return error_value(node.AsRaw(&ctx.Context), float32Type)
	}

	*(*float32)(vp) = float32(ret)
	return nil
}

type f64Decoder struct{}

func (d *f64Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsF64()
	if err != nil {
		return err
	}
	*(*float64)(vp) = float64(ret)
	return nil
}

type boolDecoder struct {
}

func (d *boolDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsBool()
	if err != nil {
		return err
	}
	*(*bool)(vp) = bool(ret)
	return nil
}

type stringDecoder struct {
}

func (d *stringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	ret, err := node.AsStr(&ctx.Context)
	if err != nil {
		return err
	}
	*(*string)(vp) = ret
	return nil
}

type numberDecoder struct {
}

func (d *numberDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	num, err := node.AsNumber(&ctx.Context)
	if err != nil {
		return err
	}
	*(*json.Number)(vp) = num
	return nil
}

type recuriveDecoder struct {
	typ *rt.GoType
}

func (d *recuriveDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	dec, err := findOrCompile(d.typ)
	if err != nil {
		return err
	}
	return dec.FromDom(vp, node, ctx)
}
