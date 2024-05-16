package ast

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/bytedance/sonic/internal/decoder"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func getSample(width int, depth int) string {
	obj := map[string]interface{}{}
	for i := 0; i < width; i++ {
		var v interface{}
		if depth > 0 {
			v = json.RawMessage(getSample(width/2+1, depth-1))
		} else {
			v = 1
		}
		obj[strconv.Itoa(i)] = v
	}
	js, _ := json.Marshal(obj)
	return string(js)
}

func TestNodeParse(t *testing.T) {
	n1, err := NewParser(`[1,"1",true,null]`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.node.Kids), 4)

	n1, err = NewParser(`[]`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.node.Kids), 0)

	n1, err = NewParser(`{}`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.node.Kids), 0)

	n1, err = NewParser(`{"key": null, "k2": {}}`).Parse()
	require.NoError(t, err)
	spew.Dump(n1.node.Kids, len(n1.node.Kids))
	require.Equal(t, len(n1.node.Kids), 4)

	src := getSample(100, 0)
	n, err := NewParser(src).Parse()
	require.NoError(t, err)
	n50 := n.GetByPath("50")
	require.Empty(t, n50.Error())
	v, _ := n50.Int64()
	require.Equal(t, int64(1), v)
	js, err := n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
	src = getSample(100, 1)
	n, err = NewParser(src).Parse()
	require.NoError(t, err)
	js, err = n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
}

func BenchmarkNode_GetByPath(b *testing.B) {
	b.Run("10/2", func(b *testing.B) {
		src := getSample(10, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			_ = n.GetByPath("5")
		}
	})
	b.Run("10/2/2", func(b *testing.B) {
		src := getSample(10, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			x := n.GetByPath("5", "5")
			if x.Check() != nil {
				b.Fatal(x.Error())
			}
		}
	})
	b.Run("100/2", func(b *testing.B) {
		src := getSample(100, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			_ = n.GetByPath("50")
		}
	})
	b.Run("100/2/2", func(b *testing.B) {
		src := getSample(100, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			x := n.GetByPath("50", "50")
			if x.Check() != nil {
				b.Fatal(x.Error())
			}
		}
	})
	b.Run("1000/2", func(b *testing.B) {
		src := getSample(1000, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			_ = n.GetByPath("500")
		}
	})
	b.Run("1000/2/2", func(b *testing.B) {
		src := getSample(1000, 1)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i := 0; i < b.N; i++ {
			x := n.GetByPath("500", "500")
			if x.Check() != nil {
				b.Fatal(x.Error())
			}
		}
	})
}

func BenchmarkParse(b *testing.B) {
	b.Run("10-0", func(b *testing.B) {
		src := getSample(10, 0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("10-1", func(b *testing.B) {
		src := getSample(10, 1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("100-0", func(b *testing.B) {
		src := getSample(100, 0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("100-1", func(b *testing.B) {
		src := getSample(100, 1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("1000-0", func(b *testing.B) {
		src := getSample(1000, 0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("1000-1", func(b *testing.B) {
		src := getSample(1000, 1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
}

func TestNode_UnsetByPath(t *testing.T) {
	type args struct {
		path []interface{}
	}
	tests := []struct {
		name   string
		src    string
		args   args
		want   string
		val    interface{}
		exist  bool
		err    error
	}{
		{
			name: "self",
			src:  `[1]`,
			args: args{path: []interface{}{}},
			want: ``,
			exist: true,
			val: nil,
		},
		{
			name: "unsupporetd",
			src:  `1`,
			args: args{path: []interface{}{1}},
			want: `1`,
			exist: false,
			val: int64(1),
			err: ErrUnsupportType,
		},
		{
			name: "unsupporetd",
			src:  `1`,
			args: args{path: []interface{}{"1"}},
			want: `1`,
			exist: false,
			val: int64(1),
			err: ErrUnsupportType,
		},
		{
			name: "empty object",
			src:  ` { } `,
			args: args{path: []interface{}{"1"}},
			want: ` { } `,
			exist: false,
			val: map[string]interface {}{},
		},
		{
			name: "one object not-exist",
			src:  ` { "1" : 1 } `,
			args: args{path: []interface{}{"2"}},
			want: ` { "1" : 1 } `,
			exist: false,
			val: map[string]interface{}{"1":int64(1)},
		},
		{
			name: "one object exist",
			src:  ` { "1" : true } `,
			args: args{path: []interface{}{"1"}},
			want: ` {} `,
			exist: true,
			val: map[string]interface{}{},
		},
		{
			name: "two object exist",
			src:  ` { "1" : 1 , "2" : 2 } `,
			args: args{path: []interface{}{"1"}},
			want: ` { "2" : 2 } `,
			exist: true,
			val: map[string]interface{}{"2":int64(2)},
		},
		{
			name: "two object exist 2",
			src:  ` { "1" : 1 , "2" : 2 } `,
			args: args{path: []interface{}{"2"}},
			want: ` { "1" : 1 } `,
			exist: true,
			val: map[string]interface{}{"1":int64(1)},
		},
		{
			name: "empty array",
			src:  ` [ ] `,
			args: args{path: []interface{}{0}},
			want: ` [ ] `,
			exist: false,
			val: []interface{}{},
		},
		{
			name: "one array not-exist",
			src:  ` [ 1 ] `,
			args: args{path: []interface{}{1}},
			want: ` [ 1 ] `,
			exist: false,
			val: []interface{}{int64(1)},
		},
		{
			name: "one array exist",
			src:  ` [ 1 ] `,
			args: args{path: []interface{}{0}},
			want: ` [] `,
			exist: true,
			val: []interface{}{},
		},
		{
			name: "two array exist",
			src:  ` [ 1 , 2 ] `,
			args: args{path: []interface{}{0}},
			want: ` [ 2 ] `,
			exist: true,
			val: []interface{}{int64(2)},
		},
		{
			name: "two array exist 2",
			src:  ` [ 1 , 2 ] `,
			args: args{path: []interface{}{1}},
			want: ` [ 1 ] `,
			exist: true,
			val: []interface{}{int64(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println(tt.name)
			self := NewRaw(tt.src)
			exist, err := self.UnsetByPath(tt.args.path...)
			spew.Dump(self.node)
			if err != nil && tt.err == nil || err == nil && tt.err != nil {
				t.Errorf("err = %v, want %v", err, tt.err)
			}
			if exist != tt.exist {
				t.Errorf("exist = %v, want %v", exist, tt.exist)
			}
			if js, _ := self.Raw(); js != tt.want {
				t.Errorf("raw = `%v`, want `%v`", js, tt.want)
			}
			if val, e := self.Interface(decoder.OptionUseInt64); !reflect.DeepEqual(val, tt.val) {
				t.Errorf("val = %#v, val %#v, err = %v", val, tt.val, e)
			}
		})
	}
}

func TestNode_SetByPath(t *testing.T) {
	type args struct {
		path []interface{}
		val string
		allowArrayAppend bool
	}
	tests := []struct {
		name   string
		src    string
		args   args
		want   string
		val    interface{}
		exist  bool
		err    error
	}{
		{
			name: "self",
			src:  `{}`,
			args: args{path: []interface{}{}, val: "1"},
			want: `1`,
			exist: true,
			val: int64(1),
		},
		{
			name: "unsupported",
			src:  `1`,
			args: args{path: []interface{}{1}, val: "1"},
			want: `1`,
			exist: false,
			val: int64(1),
			err: ErrUnsupportType,
		},
		{
			name: "unsupported",
			src:  `1`,
			args: args{path: []interface{}{"1"}, val: "1"},
			want: `1`,
			exist: false,
			val: int64(1),
			err: ErrUnsupportType,
		},
		{
			name: "empty object",
			src:  ` { } `,
			args: args{path: []interface{}{"1"}, val: "1"},
			want: ` {"1":1 } `,
			exist: false,
			val: map[string]interface{}{"1":int64(1)},
		},
		{
			name: "one object not-exist",
			src:  ` { "1" : 1 } `,
			args: args{path: []interface{}{"2"}, val: "2"},
			want: ` { "1" : 1,"2":2 } `,
			exist: false,
			val: map[string]interface{}{"1":int64(1),"2":int64(2)},
		},
		{
			name: "one object exist",
			src:  ` { "1" : true } `,
			args: args{path: []interface{}{"1"}, val: "-1"},
			want: ` { "1" : -1 } `,
			exist: true,
			val: map[string]interface{}{"1":int64(-1)},
		},
		{
			name: "empty array",
			src:  ` [ ] `,
			args: args{path: []interface{}{0}, val: "1", allowArrayAppend: true},
			want: ` [1 ] `,
			exist: false,
			val: []interface{}{int64(1)},
		},
		{
			name: "empty array not allow insert",
			src:  ` [ ] `,
			args: args{path: []interface{}{0}, val: "1"},
			err: ErrNotExist,
			want: ` [ ] `,
			exist: false,
			val: []interface{}{},
		},
		{
			name: "one array not-exist",
			src:  ` [ 1 ] `,
			args: args{path: []interface{}{1}, val: "2", allowArrayAppend: true},
			want: ` [ 1,2 ] `,
			exist: false,
			val: []interface{}{int64(1), int64(2)},
		},
		{
			name: "one array exist",
			src:  ` [ "a" ] `,
			args: args{path: []interface{}{0}, val: "-1"},
			want: ` [ -1 ] `,
			exist: true,
			val: []interface{}{int64(-1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println(tt.name)
			self := NewRaw(tt.src)
			exist, err := self.SetByPath(tt.args.allowArrayAppend, tt.args.val, tt.args.path...)
			spew.Dump(self.node)
			if err != nil && tt.err == nil || err == nil && tt.err != nil {
				t.Errorf("err = %v, want %v", err, tt.err)
			}
			if exist != tt.exist {
				t.Errorf("exist = %v, want %v", exist, tt.exist)
			}
			if js, _ := self.Raw(); js != tt.want {
				t.Errorf("raw = `%v`, want `%v`", js, tt.want)
			}
			if val, e := self.Interface(decoder.OptionUseInt64); !reflect.DeepEqual(val, tt.val) {
				t.Errorf("val = %#v, val %#v, err = %v", val, tt.val, e)
			}
		})
	}
}

func TestNode_GetByPath(t *testing.T) {
	type args struct {
		path []interface{}
	}
	var srcArray = ` [ 1 , null , true , false , 1.0 , "\"" , [ ] , [ 1 ] , { } , { "1" : 1 } ] `
	var srcObject = ` { "1" : 1 , "2" : null , "3" : true , "4" : false , "5" : 1.0 , "6" : "\"" , "7" : [ ] , "8" : [ 1 ] , "9" : { } , "10" : { "1" : 1 } } `
	tests := []struct {
		name   string
		src    string
		args   args
		want   string
		val    interface{}
		err    error
	}{
		{
			name: "self",
			src: ` [ 1 ] `,
			args: args{path: []interface{}{}},
			want: ` [ 1 ] `,
			val:  []interface{}{int64(1)},
		},
		{
			name: "unsupported type",
			src: ` 1 `,
			args: args{path: []interface{}{0}},
			err: ErrUnsupportType,
		},
		{
			name: "unsupported type",
			src: ` 1 `,
			args: args{path: []interface{}{"0"}},
			err: ErrUnsupportType,
		},
		{
			name: "array not-exist",
			src: ` [ ] `,
			args: args{path: []interface{}{0}},
			err: ErrNotExist,
		},
		{
			name: "array int",
			src:  srcArray,
			args: args{path: []interface{}{0}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "array null",
			src:  srcArray,
			args: args{path: []interface{}{1}},
			want: `null`,
			val:  nil,
		},
		{
			name: "array true",
			src:  srcArray,
			args: args{path: []interface{}{2}},
			want: `true`,
			val:  true,
		},
		{
			name: "array false",
			src:  srcArray,
			args: args{path: []interface{}{3}},
			want: `false`,
			val:  false,
		},
		{
			name: "array float",
			src:  srcArray,
			args: args{path: []interface{}{4}},
			want: `1.0`,
			val:  1.0,
		},
		{
			name: "array string",
			src:  srcArray,
			args: args{path: []interface{}{5}},
			want: `"\""`,
			val:  `"`,
		},
		{
			name: "array empty array",
			src:  srcArray,
			args: args{path: []interface{}{6}},
			want: `[ ]`,
			val:  []interface{}{},
		},
		{
			name: "array array",
			src:  srcArray,
			args: args{path: []interface{}{7}},
			want: `[ 1 ]`,
			val:  []interface{}{int64(1)},
		},
		{
			name: "array array int",
			src:  srcArray,
			args: args{path: []interface{}{7, 0}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "array array not-exst",
			src:  srcArray,
			args: args{path: []interface{}{7, 1}},
			err: ErrNotExist,
		},
		{
			name: "array empty object",
			src:  srcArray,
			args: args{path: []interface{}{8}},
			want: `{ }`,
			val:  map[string]interface{}{},
		},
		{
			name: "array object",
			src:  srcArray,
			args: args{path: []interface{}{9}},
			want: `{ "1" : 1 }`,
			val:  map[string]interface{}{"1": int64(1)},
		},
		{
			name: "array object int",
			src:  srcArray,
			args: args{path: []interface{}{9, "1"}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "array object not-exist",
			src:  srcArray,
			args: args{path: []interface{}{9, "2"}},
			err: ErrNotExist,
		},
		{
			name: "object int",
			src:  srcObject,
			args: args{path: []interface{}{"1"}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "object null",
			src:  srcObject,
			args: args{path: []interface{}{"2"}},
			want: `null`,
			val:  nil,
		},
		{
			name: "object true",
			src:  srcObject,
			args: args{path: []interface{}{"3"}},
			want: `true`,
			val:  true,
		},
		{
			name: "object false",
			src:  srcObject,
			args: args{path: []interface{}{"4"}},
			want: `false`,
			val:  false,
		},
		{
			name: "object float",
			src:  srcObject,
			args: args{path: []interface{}{"5"}},
			want: `1.0`,
			val:  1.0,
		},
		{
			name: "object string",
			src:  srcObject,
			args: args{path: []interface{}{"6"}},
			want: `"\""`,
			val:  `"`,
		},
		{
			name: "object empty array",
			src:  srcObject,
			args: args{path: []interface{}{"7"}},
			want: `[ ]`,
			val:  []interface{}{},
		},
		{
			name: "object array",
			src:  srcObject,
			args: args{path: []interface{}{"8"}},
			want: `[ 1 ]`,
			val:  []interface{}{int64(1)},
		},
		{
			name: "object empty object",
			src:  srcObject,
			args: args{path: []interface{}{"9"}},
			want: `{ }`,
			val:  map[string]interface{}{},
		},
		{
			name: "object object",
			src:  srcObject,
			args: args{path: []interface{}{"10"}},
			want: `{ "1" : 1 }`,
			val:  map[string]interface{}{"1": int64(1)},
		},
		{
			name: "object array int",
			src:  srcObject,
			args: args{path: []interface{}{"8", 0}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "object array not-exist",
			src:  srcObject,
			args: args{path: []interface{}{"8", 1}},
			err: ErrNotExist,
		},
		{
			name: "object object int",
			src:  srcObject,
			args: args{path: []interface{}{"10", "1"}},
			want: `1`,
			val:  int64(1),
		},
		{
			name: "object object not-exist",
			src:  srcObject,
			args: args{path: []interface{}{"10", "2"}},
			err: ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println("name:", tt.name)
			self := NewRaw(tt.src)
			got := self.GetByPath(tt.args.path...)
			err := got.Check()
			if err != nil && tt.err.Error() != err.Error() || err == nil && tt.err != nil {
				t.Errorf("Node.Interface() = %v, err %v", err, tt.err)
			}
			if err == nil {
				spew.Dump(got.node)
				if js, _ := got.Raw(); js != tt.want {
					t.Errorf("Node.GetByPath() = `%v`, want `%v`", js, tt.want)
				} else if val, _ := got.Interface(decoder.OptionUseInt64); !reflect.DeepEqual(val, tt.val) {
					t.Errorf("Node.Interface() = %v, val %v", val, tt.val)
				}
			}
		})
	}
}
