package optdec

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestParseNative(t *testing.T) {
	data := `{"a":1, "b": true,    "":    false, "d":                             null,"\"f":[], "\\": "\\", "hi": {}, "":[[[]], "", [{}]]}`
	p := newParser(data, 0, 0)
	ecode := p.parse()
	defer p.free()
	assert.Equal(t, int(ecode), 0)
	assert.Equal(t, p.Pos(), len(data))
	spew.Dump(p.nbuf.stat)
}

func TestParseNativeRetryLargeJson(t *testing.T) {
	t.Run("Object", func (t *testing.T)  {
		data := "{" + strings.Repeat("\"a\":1,", 1 << 20) + "\"a\":1}"
		p := newParser(data, 0, 0)
		ecode := p.parse()
		defer p.free()
		assert.Equal(t, int(ecode), 0)
		assert.Equal(t, int(p.Pos()), len(data))
		assert.Equal(t, int(p.nbuf.stat.object),  1)
		assert.Equal(t, int(p.nbuf.stat.object_keys), 1 << 20 + 1)
		assert.Equal(t, int(p.nbuf.stat.max_depth), 1)
		assert.Equal(t, int(p.nbuf.stat.number), 1 << 20 + 1)
	})

	t.Run("ObjectNull", func (t *testing.T)  {
		data := "{" + strings.Repeat("\"a\":null,", 1 << 20) + "\"a\":null}"
		p := newParser(data, 0, 0)
		ecode := p.parse()
		defer p.free()
		assert.Equal(t, int(ecode), 0)
		assert.Equal(t, int(p.Pos()), len(data))
		assert.Equal(t, int(p.nbuf.stat.object),  1)
		assert.Equal(t, int(p.nbuf.stat.object_keys), 1 << 20 + 1)
		assert.Equal(t, int(p.nbuf.stat.max_depth), 1)
	})

	t.Run("Object2", func (t *testing.T)  {
		data := "{\"top\": {" + strings.Repeat("\"a\":1,", 1 << 20) + "\"a\":1}, \"final\": true}"
		p := newParser(data, 0, 0)
		ecode := p.parse()
		defer p.free()
		assert.Equal(t, int(ecode), 0)
		assert.Equal(t, int(p.Pos()), len(data))
		assert.Equal(t, int(p.nbuf.stat.object),  2)
		assert.Equal(t, int(p.nbuf.stat.object_keys), 1 << 20 + 3)
		assert.Equal(t, int(p.nbuf.stat.max_depth), 2)
		assert.Equal(t, int(p.nbuf.stat.number), 1 << 20 + 1)
	})

	t.Run("Array", func (t *testing.T)  {
		data := "[" + strings.Repeat("1,", 1 << 20) + "1]"
		p := newParser(data, 0, 0)
		ecode := p.parse()
		defer p.free()
		assert.Equal(t, int(ecode), 0)
		assert.Equal(t, p.Pos(), len(data))
		assert.Equal(t, int(p.nbuf.stat.array),  1)
		assert.Equal(t, int(p.nbuf.stat.array_elems), 1 << 20 + 1)
		assert.Equal(t, int(p.nbuf.stat.number), 1 << 20 + 1)
		assert.Equal(t, int(p.nbuf.stat.max_depth), 1)
	})
}
