//go:build goexperiment.swisstable || (!go1.16 && go1.23)
// +build goexperiment.swisstable !go1.16,go1.23

// 	Most Codes are copied from Go1.20.4. Modified parts pls see comments with #MODIFIED

package rt

import (
	"sync/atomic"
	"unsafe"
)

var EnbaleFastMap bool = false

// TODO: support for swisstable
type MapPool struct {
}

func NewMapPool(t *GoMapType, mapHint int, kvHint int) MapPool {
	return MapPool {}
}

func (self *MapPool) GetMap(hint int) *Hmap {
	panic("not support")
	return nil
}

func (self *MapPool) Remain() (hdrs int, btks int) {
	panic("not support")
	return
}
