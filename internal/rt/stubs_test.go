package rt

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)


func TestStubsMake(t *testing.T) {
	t.Run("NonPtr", func(t *testing.T) {
		old := &[]int{}
		news := MakeSlice(unsafe.Pointer(old),  UnpackType(reflect.TypeOf(int(1))), 10000)
		new := *(*[]int)(unsafe.Pointer(news))
		for i := 0; i < 10000; i++ {
			assert.Equal(t, new[i], 0)
		}
	})

	t.Run("HasPtr", func(t *testing.T) {
		old := &[]*int{}
		news := MakeSlice(unsafe.Pointer(old),  UnpackType(reflect.TypeOf((*int)(nil))), 10000)
		new := *(*[]*int)(unsafe.Pointer(news))
		for i := 0; i < 10000; i++ {
			assert.Equal(t, new[i], (*int)(nil))
		}
	})
}