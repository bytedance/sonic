package internal

import (
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic/dev/internal/rt"
)

var (
	int32Type  = rt.UnpackType(reflect.TypeOf(int32(0)))
	int64Type  = rt.UnpackType(reflect.TypeOf(int64(0)))
	uint32Type = rt.UnpackType(reflect.TypeOf(uint32(0)))
	uint64Type = rt.UnpackType(reflect.TypeOf(uint64(0)))
	strType    = rt.UnpackType(reflect.TypeOf(""))
	anyType    = rt.UnpackType(reflect.TypeOf((*interface{})(nil)).Elem())
	_ZSTPtr    = unsafe.Pointer(&struct{}{})
)
