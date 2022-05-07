package decoder

import (
    `bytes`
    `encoding/json`
    `io`
    `io/ioutil`
    `strings`
    `testing`

    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
    `github.com/stretchr/testify/require`
)

var _Single_JSON = `{"aaaaa":"` + strings.Repeat("b",1024) + `"} b {}`
var _Double_JSON = `{"aaaaa":"bbbbb"}          {"11111":"22222"} b {}`     
var _Triple_JSON = `{"aaaaa":"` + strings.Repeat("b",1024) + `"}{ } {"11111":"` + 
    strings.Repeat("2",1024)+`"} b {}`

func TestDecodeMulti(t *testing.T) {
    var str = _Triple_JSON

    var r1 = strings.NewReader(str)
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    es1 := d1.Decode(&v1)
    var r2 = strings.NewReader(str)
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)
    ee1 := d2.Decode(&v2)
    assert.Equal(t, es1, ee1)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())
    // fmt.Printf("v:%#v\n", v1)

    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)

    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
}

type HaltReader struct {
    halts map[int]bool
    buf string
    p int
}

func NewHaltReader(buf string, halts map[int]bool) *HaltReader {
    return &HaltReader{
        halts: halts,
        buf: buf,
        p: 0,
    }
}

func (self *HaltReader) Read(p []byte) (int, error) {
    t := 0
    for ; t < len(p); {
        if self.p >= len(self.buf) {
            return t, io.EOF
        }
        if b, ok := self.halts[self.p]; b {
            self.halts[self.p] = false
            return t, nil
        } else if ok {
            delete(self.halts, self.p)
            return 0, nil
        }
        p[t] = self.buf[self.p]
        self.p++
        t++
    }
    return t, nil
}

func (self *HaltReader) Reset(buf string) {
    self.p = 0
    self.buf = buf
}

var testHalts = func () map[int]bool {
    return map[int]bool{
        1: true,
        10:true,
        20: true}
}

func TestDecodeHalt(t *testing.T) {
    var str = _Triple_JSON
    var r1 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    err1 := d1.Decode(&v1)
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)
    err2 := d2.Decode(&v2)
    assert.Equal(t, err1, err2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    // assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)

    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
}

func TestBuffered(t *testing.T) {
    var str = _Triple_JSON
    var r1 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var d1 = json.NewDecoder(r1)
    require.Nil(t, d1.Decode(&v1))
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)
    require.Nil(t, d2.Decode(&v2))
    left1, err1 := ioutil.ReadAll(d1.Buffered())
    require.Nil(t, err1)
    left2, err2 := ioutil.ReadAll(d2.Buffered())
    require.Nil(t, err2)
    require.Equal(t, d1.InputOffset(), d2.InputOffset())
    if !bytes.Contains(left2, left1) {
        t.Fatal(string(left2), string(left1))
    }

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, d1.InputOffset(), d2.InputOffset())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, d1.InputOffset(), d2.InputOffset())
}

func TestMore(t *testing.T) {
    var str = _Triple_JSON
    var r2 = NewHaltReader(str, testHalts())
    var v2 map[string]interface{}
    var d2 = NewStreamDecoder(r2)
    var r1 = NewHaltReader(str, testHalts())
    var v1 map[string]interface{}
    var d1 = jsoniter.NewDecoder(r1)
    require.Nil(t, d1.Decode(&v1))
    require.Nil(t, d2.Decode(&v2))
    require.Equal(t, d1.More(), d2.More())

    es4 := d1.Decode(&v1)
    ee4 := d2.Decode(&v2)
    assert.Equal(t, es4, ee4)
    assert.Equal(t, v1, v2)
    require.Equal(t, d1.More(), d2.More())

    es2 := d1.Decode(&v1)
    ee2 := d2.Decode(&v2)
    assert.Equal(t, es2, ee2)
    assert.Equal(t, v1, v2)
    require.Equal(t, d1.More(), d2.More())

    es3 := d1.Decode(&v1)
    assert.NotNil(t, es3)
    ee3 := d2.Decode(&v2)
    assert.NotNil(t, ee3)
    require.Equal(t, d1.More(), d2.More())

    es5 := d1.Decode(&v1)
    assert.NotNil(t, es5)
    ee5 := d2.Decode(&v2)
    assert.NotNil(t, ee5)
    require.Equal(t, d1.More(), d2.More())
}

func BenchmarkDecode_Std(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })
    
    b.Run("halt", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := json.NewDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}

func BenchmarkDecode_Jsoniter(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("halt", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := jsoniter.NewDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}

func BenchmarkDecodeError_Sonic(b *testing.B) {
    var str = `\b测试1234`
    for i:=0; i<b.N; i++ {
        var v1 map[string]interface{}
        _ = NewDecoder(str).Decode(&v1)
    }
}

func BenchmarkDecode_Compat(b *testing.B) {
    b.Run("single", func (b *testing.B) {
        var str = _Single_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("double", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = strings.NewReader(str)
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
            _ = dc.Decode(&v1)
        }
    })

    b.Run("halt", func (b *testing.B) {
        var str = _Double_JSON
        for i:=0; i<b.N; i++ {
            var r1 = NewHaltReader(str, testHalts())
            var v1 map[string]interface{}
            dc := NewStreamDecoder(r1)
            _ = dc.Decode(&v1)
        }
    })
}