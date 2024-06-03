//go:build !goexperiment.swisstable
// +build !goexperiment.swisstable

package rt

import (
	"testing"
	"unsafe"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestB(t *testing.T) {
	for i := 0; i < 64; i++ {
		println(calcaulateB(i))
	}
}

func TestFastMap(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		pool := NewMapPool(MapType(MapStringType), 2, 2)
		h := pool.GetMap(4)
		m := *((*map[string]interface{})((unsafe.Pointer)((&h))))
		assert.Equal(t, m, map[string]interface{}{})

		hdr, btk := pool.Remain()
		assert.Equal(t, hdr, 1)
		assert.Equal(t, btk, 1)

		h = pool.GetMap(4)
		m = *((*map[string]interface{})((unsafe.Pointer)((&h))))
		assert.Equal(t, m, map[string]interface{}{})

		hdr, btk = pool.Remain()
		assert.Equal(t, hdr, 0)
		assert.Equal(t, btk, 0)
	})

	t.Run("Normal", func(t *testing.T) {
		pool := NewMapPool(MapType(MapStringType), 2, 2)
		h := pool.GetMap(4)
		m := *((*map[string]interface{})((unsafe.Pointer)((&h))))
		exp := map[string]interface{}{}

		for i := 0; i < 8; i++ {
			m[strconv.Itoa(i)] = i
		}
		for i := 0; i < 8; i++ {
			exp[strconv.Itoa(i)] = i
		}
		assert.Equal(t, m, exp)
		hdr, btk := pool.Remain()
		assert.Equal(t, hdr, 1)
		assert.Equal(t, btk, 1)

		for i := 0; i < 100; i++ {
			m[strconv.Itoa(i)] = i
		}
		for i := 0; i < 100; i++ {
			exp[strconv.Itoa(i)] = i
		}
		assert.Equal(t, m, exp)
		hdr, btk = pool.Remain()
		assert.Equal(t, hdr, 1)
		assert.Equal(t, btk, 1)
	})

	t.Run("Large", func(t *testing.T) {
		pool := NewMapPool(MapType(MapStringType), 2, 164)
		h := pool.GetMap(127)
		m := *((*map[string]interface{})((unsafe.Pointer)((&h))))
		exp := map[string]interface{}{}

		for i := 0; i < 8; i++ {
			m[strconv.Itoa(i)] = i
		}
		for i := 0; i < 8; i++ {
			exp[strconv.Itoa(i)] = i
		}
		assert.Equal(t, m, exp)

		hdr, btk := pool.Remain()
		isOvf := btk < 164 - 128/8
		assert.Equal(t, hdr, 1)
		assert.True(t, isOvf) 
	})

	t.Run("Many", func(t *testing.T) {
		totalMap := 1000
		pool := NewMapPool(MapType(MapStringType), 1000, 2000 * 8)

		for i := 0; i < 1000; i++ {
			h := pool.GetMap(10) // B is 2
			m := *((*map[string]interface{})((unsafe.Pointer)((&h))))
			exp := map[string]interface{}{}
	
			for i := 0; i < 8; i++ {
				m[strconv.Itoa(i)] = i
			}
			for i := 0; i < 8; i++ {
				exp[strconv.Itoa(i)] = i
			}
			assert.Equal(t, m, exp)
			hdr, _ := pool.Remain()
			assert.Equal(t, hdr, totalMap - (i+1))
			// assert.Equal(t, btk, totalBtk - 2* (i+1))
		}
	})
}