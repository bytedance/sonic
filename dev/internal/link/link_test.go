package link


import (
	"testing"
	"unsafe"
	"github.com/bytedance/sonic/dev/ccall"
	"github.com/bytedance/sonic/dev/internal"
)


var cfunc unsafe.Pointer


func init() {
	cfuncs := Link(sonic_rs_blob, []string{"func_1args"});
	ret := ccall.C_systemstack_1arg(123, cfuncs[0])
	cfunc = cfuncs[0]
	println("ret is ", ret);
}

func Testx(t *testing.T) {

}


func BenchmarkFfi(b *testing.B) {
	r1 :=  internal.Cgo_func(123)
	r2 :=  ccall.C_systemstack_1arg(123, cfunc)
	if r1 != r2  || r1 != 123 {
		panic("invalid tests")
	}
	b.Run("Cgo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = internal.Cgo_func(123)
		}
	})

	b.Run("Ccall", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ccall.C_systemstack_1arg(123, cfunc)
		}
	})

	b.Run("Asm2Asm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ccall.C_systemstack_1arg(123, cfunc)
		}
	})
}