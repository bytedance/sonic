//go:build !go1.26 && !goexperiment.swissmap
// +build !go1.26,!goexperiment.swissmap

package rt

func (self *GoMapType) IndirectElem() bool {
	return self.Flags&2 != 0
}
