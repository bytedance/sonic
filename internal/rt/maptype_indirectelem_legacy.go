//go:build !go1.26
// +build !go1.26

package rt

func (self *GoMapType) IndirectElem() bool {
	return self.Flags&2 != 0
}
