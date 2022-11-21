package ast

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var concurrency = 1000

func TestRawNode(t *testing.T) {
	_, err := NewSearcher(` { ] `).GetRawByPath()
	require.Error(t, err)
	d1 := ` {"a":1,"b":[1,1,1],"c":{"d":1,"e":1,"f":1}} `
	root, err := NewSearcher(d1).GetRawByPath()
	require.NoError(t, err)
	require.Equal(t, len(d1)-2, len(root.js))
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
	require.Equal(t, int64(1), v4)
	v5, err := root.GetByPath("b").Interface()
	require.NoError(t, err)
	require.Equal(t, []interface{}{float64(1), float64(1), float64(1)}, v5)
	v6, err := root.GetByPath("c").Interface()
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"d": float64(1), "e": float64(1), "f": float64(1)}, v6)
}

func TestConcurrentGetByPath(t *testing.T) {
	cont, err := NewSearcher(`{"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"a":1}`).GetRawByPath()
	if err != nil {
		t.Fatal(err)
	}
	c := make(chan struct{}, 7)
	wg := sync.WaitGroup{}

	for i := 0; i < concurrency; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("b", 1)
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("b", 0)
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("b", 2)
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("c", "d")
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("c", "f")
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("c", "e")
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			<-c
			v := cont.GetByPath("a")
			require.NoError(t, v.Check())
			vv, _ := v.Int64()
			require.Equal(t, int64(1), vv)
		}()
	}

	for i := 0; i < 7*concurrency; i++ {
		c <- struct{}{}
	}
	
	wg.Wait()
}

func BenchmarkNodeGetContByPath(b *testing.B) {
    root, derr := NewParser(_TwitterJson).Parse()
    if derr != 0 {
        b.Fatalf("decode failed: %v", derr.Error())
    }
    _, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
	cont := RawNode{js: _TwitterJson}

	b.Run("Node", func(b *testing.B) {
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			_, _ = root.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
    b.Run("ContNode", func(b *testing.B) {
		b.ResetTimer()
        for i:=0; i<b.N; i++ {
			_, _ = cont.GetByPath("statuses", 3, "entities", "hashtags", 0, "text").String()
		}
    })
}