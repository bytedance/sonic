package rt

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

const POOL_SIZE int = 128

func TestFastConvTSlice(t *testing.T) {
	pool := NewTslicePool(POOL_SIZE)
	typ := UnpackType(reflect.TypeOf([]string(nil)))

	t.Run("Empty",  func(t *testing.T) {
		buf := []string{ }

		var got, exp interface{}
		exp = interface{}(buf)

		b := *((*GoSlice)(unsafe.Pointer(&buf)))
		spew.Dump(b)
		pool.Conv(b, typ, &got)
		assert.Equal(t, exp, got)

		spew.Dump(&b)
		spew.Dump(*((*GoEface)(unsafe.Pointer(&exp))))
	})

	t.Run("Nil",  func(t *testing.T) {
		var buf []string
		buf = nil

		var got, exp interface{}
		exp = interface{}(buf)

		b := *((*GoSlice)(unsafe.Pointer(&buf)))
		pool.Conv(b, typ, &got)
		assert.Equal(t, exp, got)
		spew.Dump(exp, got)
		spew.Dump(*((*GoEface)(unsafe.Pointer(&exp))))
	})

	t.Run("Normal",  func(t *testing.T) {
		buf := []string{ "hello" }

		var got, exp interface{}
		exp = interface{}(buf)

		b := *((*GoSlice)(unsafe.Pointer(&buf)))
		pool.Conv(b, typ, &got)
		assert.Equal(t, exp, got)
	})
}

func TestFastConvTString(t *testing.T) {
	pool := NewTstringPool(POOL_SIZE)
	t.Run("Empty", func(t *testing.T) {
		s := "" 
		var got, exp interface{}
		exp = interface{}(s)
		pool.Conv(s, &got)
		assert.Equal(t, exp, got)
	})

	t.Run("Normal", func(t *testing.T) {
		s := "hello" 
		var got, exp interface{}
		exp = interface{}(s)
		pool.Conv(s, &got)
		assert.Equal(t, exp, got)
	})
}



func TestFastConvT64(t *testing.T) {
	pool := NewT64Pool(POOL_SIZE)
	t.Run("Small", func(t *testing.T) {
		v := true
		typ := UnpackType(reflect.TypeOf(v))

		var got, exp interface{}
		exp = interface{}(v)
		pool.Conv(uint64(1), typ, &got)
		assert.Equal(t, exp, got)
	})

	t.Run("Normal", func(t *testing.T) {
		v := 123456
		typ := UnpackType(reflect.TypeOf(v))

		var got, exp interface{}
		exp = interface{}(v)
		pool.Conv(uint64(v), typ, &got)
		assert.Equal(t, exp, got)
	})
}

func BenchmarkFastConvT64(b *testing.B) {
	typ := UnpackType(reflect.TypeOf(int(0)))

	var v1, v2 interface{}
	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pool := NewT64Pool(POOL_SIZE)
			for j := 0; j < POOL_SIZE; j++ {
				v := 123456
				pool.Conv(uint64(v), typ, &v1)
			}
		}
	})


	b.Run("Naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < POOL_SIZE; j++ {
				v := 123456
				v2 = interface{}(v)
			}
		}
	})

	assert.Equal(b, v1, v2)
}


func BenchmarkFastConvTString(b *testing.B) {
	var v1, v2 interface{}
	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pool := NewTstringPool(POOL_SIZE)
			for j := 0; j < POOL_SIZE; j++ {
				v := "123456"
				pool.Conv(v, &v1)
			}
		}
	})


	b.Run("Naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < POOL_SIZE; j++ {
				v := "123456"
				v2 = interface{}(v)
			}
		}
	})

	assert.Equal(b, v1, v2)
}

func BenchmarkFastConvTSlice(b *testing.B) {
	typ := UnpackType(reflect.TypeOf([]byte(nil)))
	v := []byte("123456")
	var v1, v2 interface{}

	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pool := NewTslicePool(POOL_SIZE)
			for j := 0; j < POOL_SIZE; j++ {
				pool.Conv(*((*GoSlice)(unsafe.Pointer(&v))), typ, &v1)
			}
		}
	})


	b.Run("Naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < POOL_SIZE; j++ {
				v2 = interface{}(v)
			}
		}
	})

	assert.Equal(b, v1, v2)
}