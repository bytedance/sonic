//go:build go1.17 && !go1.27
// +build go1.17,!go1.27

package x86

import "github.com/bytedance/sonic/loader"

const FP_args = _FP_args

func (self *Assembler) Export() ([]byte, loader.Pcdata) {
	return self.BaseAssembler.Export()
}
