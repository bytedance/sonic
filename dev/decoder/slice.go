package decoder

import (
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
	"github.com/bytedance/sonic/dev/internal"
)

type sliceDecoder struct {
	elemType *rt.GoType
	elemDec  decFunc
}

var (
	emptyPtr = &struct{}{}
)

func (d *sliceDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	slice := internal.MakeSlice(vp, d.elemType, arr.Len())
	elems := slice.Ptr

	next := arr.Children()

	var gerr error
	for i := 0; i < arr.Len(); i++ {
		val := internal.NewNode(next)
		elem := unsafe.Pointer(uintptr(elems) + uintptr(i)*d.elemType.Size)
		err = d.elemDec.FromDom(elem, val, ctx)
		if gerr == nil && err != nil {
			gerr = err
		}
		next = val.Next()
	}

	*(*rt.GoSlice)(vp) = *slice
	return gerr
}

type arrayDecoder struct {
	len      int
	elemType *rt.GoType
	elemDec  decFunc
}

func (d *arrayDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		return nil
	}

	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	next := arr.Children()
	i := 0
	for ; i < d.len && i < arr.Len(); i++ {
		elem := unsafe.Pointer(uintptr(vp) + uintptr(i)*d.elemType.Size)
		val := internal.NewNode(next)
		err = d.elemDec.FromDom(elem, val, ctx)
		if err != nil {
			return err
		}
		next = val.Next()
	}

	/* zero rest of array */
	ptr := unsafe.Pointer(uintptr(vp) + uintptr(i)*d.elemType.Size)
	n := uintptr(d.len-i) * d.elemType.Size
	internal.ClearMemory(d.elemType, ptr, n)
	return nil
}

type sliceEfaceDecoder struct {
}

func (d *sliceEfaceDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	if Options(ctx.Options)&OptionUseNumber == 0 && Options(ctx.Options)&OptionUseInt64 == 0 {
		return node.AsSliceEface(&ctx.Context, vp)
	}

	if Options(ctx.Options)&OptionUseNumber != 0 {
		return node.AsSliceEfaceUseNumber(&ctx.Context, vp)
	}

	return node.AsSliceEfaceUseInt64(&ctx.Context, vp)
}

type sliceI32Decoder struct {
}

func (d *sliceI32Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	return node.AsSliceI32(&ctx.Context, vp)
}

type sliceI64Decoder struct {
}

func (d *sliceI64Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	return node.AsSliceI64(&ctx.Context, vp)
}

type sliceU32Decoder struct {
}

func (d *sliceU32Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	return node.AsSliceU32(&ctx.Context, vp)
}

type sliceU64Decoder struct {
}

func (d *sliceU64Decoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	return node.AsSliceU64(&ctx.Context, vp)
}

type sliceStringDecoder struct {
}

func (d *sliceStringDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	return node.AsSliceString(&ctx.Context, vp)
}

type sliceBytesDecoder struct {
}

func (d *sliceBytesDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	s, err := node.AsSliceBytes(&ctx.Context)
	if err != nil {
		return err
	}

	*(*[]byte)(vp) = s
	return nil
}

type sliceBytesUnmarshalerDecoder struct {
	elemType *rt.GoType
	elemDec  decFunc
}

func (d *sliceBytesUnmarshalerDecoder) FromDom(vp unsafe.Pointer, node internal.Node, ctx *context) error {
	if node.IsNull() {
		*(*rt.GoSlice)(vp) = rt.GoSlice{}
		return nil
	}

	/* parse JSON string into `[]byte` */
	if node.IsStr() {
		slice, err := node.AsSliceBytes(&ctx.Context)
		if err != nil {
			return err
		}
		*(*[]byte)(vp) = slice
		return nil
	}

	/* parse JSON array into `[]byte` */
	arr, err := node.AsArr()
	if err != nil {
		return err
	}

	slice := internal.MakeSlice(vp, d.elemType, arr.Len())
	elems := slice.Ptr

	next := arr.Children()
	for i := 0; i < arr.Len(); i++ {
		child := internal.NewNode(next)
		elem := unsafe.Pointer(uintptr(elems) + uintptr(i)*d.elemType.Size)
		err = d.elemDec.FromDom(elem, child, ctx)
		if err != nil {
			return err
		}
		next = child.Next()
	}

	*(*rt.GoSlice)(vp) = *slice
	return nil
}
