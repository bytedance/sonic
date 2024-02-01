package decoder

import (
	"encoding/json"
	"math"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/bytedance/sonic/dev/internal"
)

type ptrStrDecoder struct {
	typ   *rt.GoType
	deref decFunc
}

// Pointer Value is allocated in the Caller
func (d *ptrStrDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*unsafe.Pointer)(vp) = nil
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		*(*unsafe.Pointer)(vp) = nil
		return err
	}

	if *(*unsafe.Pointer)(vp) == nil {
		*(*unsafe.Pointer)(vp) = mallocgc(d.typ.Size, d.typ, true)
	}

	return d.deref.FromDom(*(*unsafe.Pointer)(vp), node, ctx)
}

type boolStringDecoder struct {
}

func (d *boolStringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseBool(s)
	if err != nil {
		return err
	}

	*(*bool)(vp) = ret
	return nil
}

type i8StringDecoder struct{}

func (d *i8StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseI64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxInt8 || ret < math.MinInt8 {
		return error_value(node.AsRaw(&ctx.Context), int8Type)
	}

	*(*int8)(vp) = int8(ret)
	return nil
}

type i16StringDecoder struct{}

func (d *i16StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseI64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxInt16 || ret < math.MinInt16 {
		return error_value(node.AsRaw(&ctx.Context), int16Type)
	}

	*(*int16)(vp) = int16(ret)
	return nil
}

type i32StringDecoder struct{}

func (d *i32StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseI64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return error_value(node.AsRaw(&ctx.Context), int32Type)
	}

	*(*int32)(vp) = int32(ret)
	return nil
}

type i64StringDecoder struct{}

func (d *i64StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseI64(s)
	if err != nil {
		return err
	}

	*(*int64)(vp) = int64(ret)
	return nil
}

type u8StringDecoder struct{}

func (d *u8StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseU64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxUint8 {
		return error_value(node.AsRaw(&ctx.Context), uint8Type)
	}

	*(*uint8)(vp) = uint8(ret)
	return nil
}

type u16StringDecoder struct{}

func (d *u16StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseU64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxUint16 {
		return error_value(node.AsRaw(&ctx.Context), uint16Type)
	}

	*(*uint16)(vp) = uint16(ret)
	return nil
}

type u32StringDecoder struct{}

func (d *u32StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseU64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxUint32 {
		return error_value(node.AsRaw(&ctx.Context), uint32Type)
	}

	*(*uint32)(vp) = uint32(ret)
	return nil
}

type u64StringDecoder struct{}

func (d *u64StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseU64(s)
	if err != nil {
		return err
	}

	*(*uint64)(vp) = uint64(ret)
	return nil
}

type f32StringDecoder struct{}

func (d *f32StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseF64(s)
	if err != nil {
		return err
	}

	if ret > math.MaxFloat32 || ret < -math.MaxFloat32 {
		return error_value(node.AsRaw(&ctx.Context), float32Type)
	}

	*(*float32)(vp) = float32(ret)
	return nil
}

type f64StringDecoder struct{}

func (d *f64StringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	// check null
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.ParseF64(s)
	if err != nil {
		return err
	}

	*(*float64)(vp) = float64(ret)
	return nil
}

/* parse string field with string options */
type strStringDecoder struct{}

func (d *strStringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStrRef(&ctx.Context)
	/* deal with empty string */
	if err != nil || s == "null" {
		return err
	}

	ret, err := internal.Unquote(s)
	if err != nil {
		return err
	}

	*(*string)(vp) = ret
	return nil
}

type numberStringDecoder struct{}

func (d *numberStringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	s, err := node.AsStr(&ctx.Context)
	if err != nil || s == "null" {
		return err
	}

	end, err := internal.SkipNumberFast(s, 0)
	if err != nil {
		return err
	}

	*(*json.Number)(vp) = json.Number(s[:end])
	return nil
}
