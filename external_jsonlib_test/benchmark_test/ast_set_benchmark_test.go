package benchmark_test

import (
	"strconv"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bytedance/sonic/ast"
)

func BenchmarkBuildObject_SonicAST_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root := ast.NewNull()
		for j := 0; j < 100; j++ {
			if _, err := root.Set(strconv.Itoa(j), ast.NewString("123321")); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func BenchmarkBuildObject_Simplejson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root := simplejson.New()
		for j := 0; j < 100; j++ {
			root.Set(strconv.Itoa(j), "123321")
		}
	}
}
