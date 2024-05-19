package ast

import (
	"encoding/json"
	"errors"
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
			name: "two object exist",
			src:  ` { "1" : { "2" : 2 } } `,
			args: args{path: []interface{}{"1", "2"}, val: "-1"},
			want: ` { "1" : { "2" : -1} } `,
			exist: true,
			val: map[string]interface{}{"1":map[string]interface{}{"2":int64(-1)}},
		},
		{
			name: "two object not exist",
			src:  ` { "1" : { "2" : 2 } } `,
			args: args{path: []interface{}{"1", "3"}, val: "3"},
			want: ` { "1" : { "2" : 2,"3":3 } } `,
			exist: false,
			val: map[string]interface{}{"1":map[string]interface{}{"2":int64(2),"3":int64(3)}},
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
		{
			name: "two array exist",
			src:  ` [ [ 1 ] ] `,
			args: args{path: []interface{}{0, 0}, val: "-1"},
			want: ` [ [ -1] ] `,
			exist: true,
			val: []interface{}{[]interface{}{int64(-1)}},
		},
		{
			name: "two array not exist, disallow append",
			src:  ` [ [ 1 ] ] `,
			args: args{path: []interface{}{0, 1}, val: "-1"},
			want: ` [ [ 1 ] ] `,
			exist: false,
			val: []interface{}{[]interface{}{int64(1)}},
			err: ErrNotExist,
		},
		{
			name: "two array not exist, allow append",
			src:  ` [ [ 1 ] ] `,
			args: args{path: []interface{}{0, 1}, val: "-1", allowArrayAppend: true},
			want: ` [ [ 1,-1 ] ] `,
			exist: false,
			val: []interface{}{[]interface{}{int64(1), int64(-1)}},
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


func TestNodeCast(t *testing.T) {
	type tcase struct {
        method string
        node   Node
		opts decoder.Options
        exp    interface{}
        err    error
    }
	var cases = []tcase{
		// none
        {"Interface", Node{}, 0, interface{}(nil), ErrNotExist},
        {"Bool", Node{}, 0, false, ErrNotExist},
        {"Int64", Node{}, 0, int64(0), ErrNotExist},
        {"Float64", Node{}, 0, float64(0), ErrNotExist},
        {"Number", Node{}, 0, json.Number(""), ErrNotExist},
		{"String", Node{}, 0, "", ErrNotExist},
		{"Map", Node{}, 0, map[string]Node{}, ErrNotExist},
		{"Array", Node{}, 0, []Node{}, ErrNotExist},

		// error
		{"Interface", newError(ErrUnsupportType), 0, interface{}(nil), ErrUnsupportType},
  
		// string
		{"Interface", NewRaw(` "a"`), 0, "a", nil},
        {"Bool", NewRaw(`"a"`), 0, false, ErrUnsupportType},
        {"Int64", NewRaw(`"a"`), 0, int64(0), ErrUnsupportType},
        {"Float64", NewRaw(`"a"`), 0, float64(0), ErrUnsupportType},
        {"Number", NewRaw(`"a"`), 0, json.Number(""), ErrUnsupportType},
		{"String", NewRaw(` "a"`), 0, "a", nil},
		{"String", NewRaw(` "\"a"`), 0, `"a`, nil},
		{"String", NewRaw(` "\u263a"`), 0, `☺`, nil},
		{"String", NewString(` \u263a`), 0, ` \u263a`, nil},
		{"String", NewAny(` \u263a`, 0), 0, ` \u263a`, nil},
		{"Map", NewRaw(`"a"`), 0, map[string]Node{}, ErrUnsupportType},
		{"Array", NewRaw(`"a"`), 0, []Node{}, ErrUnsupportType},

		// bool
		{"Interface", NewRaw(` true`), 0, true, nil},
        {"Bool", NewRaw(` true`), 0, true, nil},
        {"Bool", NewAny(true, 0), 0, true, nil},
        {"Bool", NewBool(true), 0, true, nil},
        {"Bool", NewBool(false), 0, false, nil},
        {"Int64", NewRaw(` true`), 0, int64(0), ErrUnsupportType},
        {"Float64", NewRaw(` true`), 0, float64(0), ErrUnsupportType},
        {"Number", NewRaw(` true`), 0, json.Number(""), ErrUnsupportType},
		{"String", NewRaw(` true`), 0, "", ErrUnsupportType},
		{"Map", NewRaw(` true`), 0, map[string]Node{}, ErrUnsupportType},
		{"Array", NewRaw(` true`), 0, []Node{}, ErrUnsupportType},

		// number
		{"Interface", NewRaw(` 1 `), decoder.OptionUseInt64, int64(1), nil},
		{"Interface", NewRaw(` 1 `), decoder.OptionUseInt64, int64(1), nil},
		{"Interface", NewRaw(` 1.1 `), 0, float64(1.1), nil},
        {"Bool", NewRaw(` 1 `), 0, false, ErrUnsupportType},
        {"Int64", NewRaw(` 1 `), 0, int64(1), nil},
        {"Float64", NewRaw(` 1 `), 0, float64(1), nil},
        {"Number", NewRaw(` 1 `), 0, json.Number("1"), nil},
		{"Int64", NewRaw(` 1.1 `), 0, int64(0), errors.New("\"Syntax error at index 0: strconv.ParseInt: parsing \\\"1.1\\\": invalid syntax\\n\\n\\t1.1\\n\\t^..\\n\"")},
        {"Float64", NewRaw(` 1.1 `), 0, float64(1.1), nil},
        {"Number", NewRaw(` 1.1 `), 0, json.Number("1.1"), nil},
		{"String", NewRaw(` 1 `), 0, "", ErrUnsupportType},
		{"Map", NewRaw(` 1 `), 0, map[string]Node{}, ErrUnsupportType},
		{"Array", NewRaw(` 1 `), 0, []Node{}, ErrUnsupportType},

		// array
		{"Interface", NewRaw(` [ 1 ] `), 0, []interface{}{float64(1)}, nil},
        {"Bool", NewRaw(` [ 1 ] `), 0, false, ErrUnsupportType},
        {"Int64", NewRaw(` [ 1 ] `), 0, int64(0), ErrUnsupportType},
        {"Float64", NewRaw(` [ 1 ] `), 0, float64(0), ErrUnsupportType},
        {"Number", NewRaw(` [ 1 ] `), 0, json.Number(""), ErrUnsupportType},
		{"String", NewRaw(` [ 1 ] `), 0, "", ErrUnsupportType},
		{"Map", NewRaw(` [ 1 ] `), 0, map[string]Node{}, ErrUnsupportType},
		{"Array", NewRaw(` [ 1 ] `), 0, []Node{NewRaw(`1`)}, nil},

		// map
		{"Interface", NewRaw(` { "1" : 1 } `), 0, map[string]interface{}{"1":float64(1)}, nil},
        {"Bool", NewRaw(` { "1" : 1 } `), 0, false, ErrUnsupportType},
        {"Int64", NewRaw(` { "1" : 1 } `), 0, int64(0), ErrUnsupportType},
        {"Float64", NewRaw(` { "1" : 1 } `), 0, float64(0), ErrUnsupportType},
        {"Number", NewRaw(` { "1" : 1 } `), 0, json.Number(""), ErrUnsupportType},
		{"String", NewRaw(` { "1" : 1 } `), 0, "", ErrUnsupportType},
		{"Map", NewRaw(` { "1" : 1 } `), 0, map[string]Node{"1":NewRaw("1")}, nil},
		{"Array", NewRaw(` { "1" : 1 } `), 0, []Node{}, ErrUnsupportType},
	}

	for i, c := range cases {
        println(i, c.node.node.JSON)
        rt := reflect.ValueOf(&c.node)
        m := rt.MethodByName(c.method)
		args := []reflect.Value{}
		maps := map[string]Node{}
		arrs := []Node{}
		if c.method == "Interface" {
			args = append(args, reflect.ValueOf(c.opts))
		} else if c.method == "Map" {
			args = append(args, reflect.ValueOf(maps))
		} else if c.method == "Array" {
			args = append(args, reflect.ValueOf(&arrs))
		}
        rets := m.Call(args)
		var v interface{}
		var e interface{}
		if c.method == "Map" {
			v = maps
			e = rets[0].Interface()
		} else if c.method == "Array" {
			v = arrs
			e = rets[0].Interface()
		} else if len(rets) == 2 {
			v = rets[0].Interface()
			e = rets[1].Interface()
        } else {
			t.Fatal(i, rets)
		}
        require.Equal(t, c.exp,  v)
		require.Equal(t, c.err == nil, e == nil)
		if e != nil {
			require.Equal(t, c.err.Error(), e.(error).Error())
		}
	}
}