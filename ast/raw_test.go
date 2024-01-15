package ast

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var concurrency = 1000

func TestValAPI(t *testing.T) {
    var cases = []struct {
        method string
        js     string
        exp    []interface{}
    }{
		{"Len", `""`, []interface{}{0}},
		{"Len", `"a"`, []interface{}{1}},
		{"Len", `1`, []interface{}{-1}},
		{"Len", `true`, []interface{}{-1}},
		{"Len", `null`, []interface{}{-1}},
		{"Len", `[]`, []interface{}{0}},
		{"Len", `[1]`, []interface{}{1}},
		{"Len", `[ 1 , 2 ]`, []interface{}{2}},
		{"Len", `{}`, []interface{}{0}},
		{"Len", `{"a":1}`, []interface{}{1}},
		{"Len", `{ "a" : 1, "b" : [1] }`, []interface{}{2}},
	}
	for i, c := range cases {
        fmt.Println(i, c)
		node := NewValue(c.js)
		rt := reflect.ValueOf(&node)
		m := rt.MethodByName(c.method)
        rets := m.Call([]reflect.Value{})
		for i, v := range c.exp {
			require.Equal(t, v, rets[i].Interface())
		}
	}
}

// IndexVal
type IndexVal struct {
	Index int
	Val Value
}

func TestGetMany(t *testing.T) {
	var cases = []struct {
		name string
        js   string
		kvs  interface{} // []KeyVal or []IndexVal
		err  error
    }{
		{"Get fail", `{}`, []keyVal{{"b", Value{}}}, nil},
		{"Get 0", `{"a":1, "b":2, "c":3}`, []keyVal{}, nil},
		{"Get 1", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", rawNode(`2`)}}, nil},
		{"Get 2", `{"a":1, "b":2, "c":3}`, []keyVal{{"a", rawNode(`1`)}, {"c", rawNode(`3`)}}, nil},
		{"Get 2", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", rawNode(`2`)}, {"a", rawNode(`1`)}}, nil},
		{"Get 3", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", rawNode(`2`)}, {"c", rawNode(`3`)}, {"a", rawNode(`1`)}}, nil},
		{"Get 3", `{"a":1, "b":2, "c":3}`, []keyVal{{"b", rawNode(`2`)}, {"c", rawNode(`3`)}, {"d", Value{}}, {"a", rawNode(`1`)}}, nil},
		{"Index fail", `[]`, []IndexVal{{1, Value{}}}, nil},
		{"Index 0", `[1, 2, 3]`, []IndexVal{{1, rawNode(`2`)}}, nil},
		{"Index 1", `[1, 2, 3]`, []IndexVal{{1, rawNode(`2`)}}, nil},
		{"Index 2", `[1, 2, 3]`, []IndexVal{{0, rawNode(`1`)}, {2, rawNode(`3`)}}, nil},
		{"Index 2", `[1, 2, 3]`, []IndexVal{{1, rawNode(`2`)}, {0, rawNode(`1`)}}, nil},
		{"Index 3", `[1, 2, 3]`, []IndexVal{{1, rawNode(`2`)}, {2, rawNode(`3`)}, {0, rawNode(`1`)}}, nil},
		{"Index 3", `[1, 2, 3]`, []IndexVal{{1, rawNode(`2`)}, {2, rawNode(`3`)}, {3, Value{}}, {0, rawNode(`1`)}}, nil},
	}
	for i, c := range cases {
        fmt.Println(i, c)
		node := NewValue(c.js)
		var err error
		if kvs, ok := c.kvs.([]keyVal); ok {
			keys := []string{}
			vals := []Value{}
			for _, kv := range kvs {
				keys = append(keys, kv.Key)
				vals = append(vals, kv.Val)
			}
			act := make([]Value, len(keys))
			err = node.GetMany(keys, act)
			require.Equal(t, vals, act)
		} else  if ids, ok := c.kvs.([]IndexVal); ok {
			keys := []int{}
			vals := []Value{}
			for _, kv := range ids {
				keys = append(keys, kv.Index)
				vals = append(vals, kv.Val)
			}
			act := make([]Value, len(keys))
			err = node.IndexMany(keys, act)
			require.Equal(t, vals, act)
		}
		if err != nil && c.err.Error() != err.Error() {
			t.Fatal(err)
		}
	}
}

func TestSetMany(t *testing.T) {
	var cases = []struct{
		name string
		js string
		kvs interface{}
		exp string
		err string
	}{
		{"replace 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", rawNode(`11`)}}, `{ "a" : 11, "b" : 2}`, ""},
		{"replace 2", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", rawNode(`11`)}, {"b", rawNode(`22`)}}, `{ "a" : 11, "b" : 22}`, ""},
		{"replace repeated", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", rawNode(`11`)}, {"a", rawNode(`22`)}}, `{ "a" : 11, "b" : 2,"a":22}`, ""},
		{"insert empty", `{ } `, []keyVal{{"c", rawNode(`33`)}}, `{ "c":33}`, ""},
		{"insert repeated", `{ } `, []keyVal{{"c", rawNode(`33`)}, {"c", rawNode(`33`)}}, `{ "c":33,"c":33}`, ""},
		{"insert 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"c", rawNode(`33`)}}, `{ "a" : 1, "b" : 2,"c":33}`, ""},
		{"insert 2", `{ "a" : 1, "b" : 2} `, []keyVal{{"c", rawNode(`33`)},{"d", rawNode(`44`)}}, `{ "a" : 1, "b" : 2,"c":33,"d":44}`, ""},
		{"replace 1, insert 1", `{ "a" : 1, "b" : 2} `, []keyVal{{"a", rawNode(`11`)}, {"c", rawNode(`33`)}}, `{ "a" : 11, "b" : 2,"c":33}`, ""},
		
		{"replace 1", `[ 1, 2] `, []IndexVal{{0, rawNode(`11`)}}, `[ 11, 2]`, ""},
		{"replace 2", `[ 1, 2] `, []IndexVal{{0, rawNode(`11`)}, {0, rawNode(`22`)}}, `[ 11, 2,22]`, ""},
		{"replace repeated", `[ 1, 2] `, []IndexVal{{0, rawNode(`11`)}, {1, rawNode(`22`)}}, `[ 11, 22]`, ""},
		{"insert empty", `[ ] `, []IndexVal{{2, rawNode(`33`)}}, `[ 33]`, ""},
		{"insert 1", `[ 1, 2] `, []IndexVal{{2, rawNode(`33`)}}, `[ 1, 2,33]`, ""},
		{"insert 2", `[ 1, 2] `, []IndexVal{{2, rawNode(`33`)},{3, rawNode(`44`)}}, `[ 1, 2,33,44]`, ""},
		{"insert repeated", `[ 1, 2] `, []IndexVal{{2, rawNode(`33`)},{2, rawNode(`44`)}}, `[ 1, 2,33,44]`, ""},
		{"replace 1, insert 1", `[ 1, 2] `, []IndexVal{{0, rawNode(`11`)}, {2, rawNode(`33`)}}, `[ 11, 2,33]`, ""},
	}

	for i, c := range cases {
		println(i, c.name)
		node := NewValue(c.js)
		var err error
		if kvs, ok := c.kvs.([]keyVal); ok {
			keys := []string{}
			vals := []Value{}
			for _, kv := range kvs {
				keys = append(keys, kv.Key)
				vals = append(vals, kv.Val)
			}
			err = node.SetMany(keys, vals)
			require.Equal(t, c.exp, node.Raw())
		} else  if ids, ok := c.kvs.([]IndexVal); ok {
			keys := []int{}
			vals := []Value{}
			for _, kv := range ids {
				keys = append(keys, kv.Index)
				vals = append(vals, kv.Val)
			}
			err = node.SetManyByIndex(keys, vals)
			require.Equal(t, c.exp, node.Raw())
		}
		if err != nil && c.err != err.Error() {
			t.Fatal(err)
		}
	}
}

func TestForEachRaw(t *testing.T) {
    val := _TwitterJson
    node, err := NewSearcher(val).GetValueByPath()
    require.Nil(t, err)
    nodes := []Value{}

	var dfs func(key string, node Value) bool
	var dfs2 func(i int, node Value) bool
	dfs = func(key string, node Value) bool {
        if node.Type() == V_OBJECT  {
        	if err := node.ForEachKV(dfs); err != nil {
				panic(err)
			}
		}
		if node.Type() == V_ARRAY  {
        	if err := node.ForEachElem(dfs2); err != nil {
				panic(err)
			}
		}
		nodes = append(nodes, node)
		return true
    }
	dfs2 = func(i int, node Value) bool {
		if node.Type() == V_OBJECT  {
        	if err := node.ForEachKV(dfs); err != nil {
				panic(err)
			}
		}
		if node.Type() == V_ARRAY  {
        	if err := node.ForEachElem(dfs2); err != nil {
				panic(err)
			}
		}
		nodes = append(nodes, node)
		return true
	}
	
    require.NoError(t, node.ForEachKV(dfs))
    require.NotEmpty(t, nodes)
}

func TestRawNode(t *testing.T) {
	_, err := NewSearcher(` { ] `).GetValueByPath()
	require.Error(t, err)
	d1 := ` {"a":1,"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"d":"{\"你好\":\"hello\"}"} `
	root, err := NewSearcher(d1).GetValueByPath()
	require.NoError(t, err)
	v1, err := root.GetByPath("a").Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v1)
	v2, err := root.GetByPath("b", 1).Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v2)
	v3, err := root.GetByPath("c", "f").Int64()
	require.NoError(t, err)
	require.Equal(t, int64(1), v3)
	v4, err := root.GetByPath("a").Interface()
	require.NoError(t, err)
	require.Equal(t, float64(1), v4)
	v5, err := root.GetByPath("b").Interface()
	require.NoError(t, err)
	require.Equal(t, []interface{}{float64(1), float64(1), float64(1)}, v5)
	v6, err := root.GetByPath("c").Interface()
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"d": float64(1), "e": float64(1), "f": float64(1)}, v6)
	v7, err := root.GetByPath("d").String()
	require.NoError(t, err)
	require.Equal(t, `{"你好":"hello"}`, v7)
}

func TestConcurrentGetByPath(t *testing.T) {
	cont, err := NewSearcher(`{"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"a":1}`).GetValueByPath()
	if err != nil {
		t.Fatal(err)
	}
	c := make(chan struct{}, 7)
	wg := sync.WaitGroup{}

	var helper = func(ps ...interface{}){
		wg.Add(1)
		defer wg.Done()
		<-c
		v := cont.GetByPath(ps...)
		require.NoError(t, v.Check())
		vv, _ := v.Int64()
		require.Equal(t, int64(1), vv)
	}
	for i := 0; i < concurrency; i++ {
		go func() {
			helper("b", 1)
		}()
		go func() {
			helper("b", 0)
		}()
		go func() {
			helper("b", 2)
		}()
		go func() {
			helper("c", "d")
		}()
		go func() {
			helper("c", "f")
		}()
		go func() {
			helper("c", "f")
		}()
		go func() {
			helper("a")
		}()
	}

	for i := 0; i < 7*concurrency; i++ {
		c <- struct{}{}
	}
	
	wg.Wait()
}

func TestRawNode_Set(t *testing.T) {
	tests := []struct{
		name string
		js string
		key interface{}
		val Value
		exist bool
		err string
		out string
	}{
		{
			name: "exist object",
			js: `{"a":1}`,
			key: "a",
			val: NewValue(`2`),
			exist: true,
			err: "",
			out: `{"a":2}`,
		},
		{
			name: "not-exist object space",
			js: `{"b":1 }`,
			key: "a",
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `{"b":1,"a":2}`,
		},
		{
			name: "not-exist object",
			js: `{"b":1}`,
			key: "a",
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `{"b":1,"a":2}`,
		},
		{
			name: "empty object",
			js: `{}`,
			key: "a",
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `{"a":2}`,
		},
		{
			name: "empty object space",
			js: `{ }`,
			key: "a",
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `{"a":2}`,
		},
		{
			name: "exist array",
			js: `[1]`,
			key: 0,
			val: NewValue(`2`),
			exist: true,
			err: "",
			out: `[2]`,
		},
		{
			name: "not exist array",
			js: `[1]`,
			key: 1,
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `[1,2]`,
		},
		{
			name: "not exist array over",
			js: `[1 ]`,
			key: 99,
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `[1,2]`,
		},
		{
			name: "empty array",
			js: `[]`,
			key: 1,
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `[2]`,
		},
		{
			name: "empty array space",
			js: `[ ]`,
			key: 1,
			val: NewValue(`2`),
			exist: false,
			err: "",
			out: `[2]`,
		},
	}
	for _, c := range tests {
		println(c.name)
		root := NewValue(c.js)
		var exist bool
		var err error
		if key, ok:= c.key.(string); ok{
			exist, err = root.Set(key, c.val)
		} else if id, ok := c.key.(int); ok {
			exist, err = root.SetByIndex(id, c.val)
		}
		if exist != c.exist {
			t.Fatal()
		}
		if err != nil && err.Error() != c.err {
			t.Fatal()
		}
		if out := root.Raw(); err != nil {
			t.Fatal()
		} else if out != c.out {
			t.Fatal()
		}
	}
}

func TestRawNode_Unset(t *testing.T) {
	tests := []struct{
		name string
		js string
		key interface{}
		exist bool
		err string
		out string
	}{
		{
			name: "exist object",
			js: `{"a":1}`,
			key: "a",
			exist: true,
			err: "",
			out: `{}`,
		},
		{
			name: "exist object space",
			js: `{ "a":1 }`,
			key: "a",
			exist: true,
			err: "",
			out: `{ }`,
		},
		{
			name: "exist object comma",
			js: `{ "a":1 , "b":2 }`,
			key: "a",
			exist: true,
			err: "",
			out: `{  "b":2 }`,
		},
		{
			name: "empty object",
			js: `{ }`,
			key: "a",
			exist: false,
			err: "",
			out: `{ }`,
		},
		{
			name: "not-exist object",
			js: `{"b":1}`,
			key: "a",
			exist: false,
			err: "",
			out: `{"b":1}`,
		},
		{
			name: "exist array",
			js: `[1]`,
			key: 0,
			exist: true,
			err: "",
			out: `[]`,
		},
		{
			name: "exist array space",
			js: `[ 1 ]`,
			key: 0,
			exist: true,
			err: "",
			out: `[ ]`,
		},
		{
			name: "exist array comma",
			js: `[ 1, 2 ]`,
			key: 0,
			exist: true,
			err: "",
			out: `[  2 ]`,
		},
		{
			name: "not exist array",
			js: `[1]`,
			key: 1,
			exist: false,
			err: "",
			out: `[1]`,
		},
		{
			name: "empty array",
			js: `[ ]`,
			key: 0,
			exist: false,
			err: "",
			out: `[ ]`,
		},
	}

	for _, c := range tests {
		println(c.name)
		root := NewValue(c.js)
		var exist bool
		var err error
		if key, ok:= c.key.(string); ok{
			exist, err = root.Unset(key)
		} else if id, ok := c.key.(int); ok {
			exist, err = root.UnsetByIndex(id)
		}
		if exist != c.exist {
			t.Fatal()
		}
		if err != nil && err.Error() != c.err {
			t.Fatal(err.Error())
		}
		if out := root.Raw(); err != nil {
			t.Fatal()
		} else {
			require.Equal(t, c.out, out)
		}
	}
}

func BenchmarkGetByPath_ReuseNode(b *testing.B) {
	b.Run("Node", func(b *testing.B) {
		root := NewRaw(_TwitterJson)
		_, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			_, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
    b.Run("Value", func(b *testing.B) {
		cont := NewValue(_TwitterJson)
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			_, _ = cont.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
}

func BenchmarkNodesGetByPath_NewNode(b *testing.B) {
	b.Run("Node", func(b *testing.B) {
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			root := NewRaw(_TwitterJson)
			_, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
    b.Run("Value", func(b *testing.B) {
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			cont := NewValue(_TwitterJson)
			_, _ = cont.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
}
