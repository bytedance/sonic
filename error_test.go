package sonic

import (
    `encoding`
    `encoding/json`
    `errors`
    `strconv`
    `testing`
    `unicode/utf8`

    `github.com/bytedance/sonic/ast`
    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
)

type S struct {
    A int    `json:",omitempty"`
    B string `json:"B1,omitempty"`
    C float64
    D bool
    E uint8
    // F []byte // unmarshal []byte is different with encoding/json
    G interface{}
    H map[string]interface{}
    I map[string]string
    J []interface{}
    K []string
    L S1
    M *S1
    N *int
    O **int
    P int `json:",string"`
    Q float64 `json:",string"`
    R int `json:"-"`
    T struct {}
    U [2]int
    V uintptr
    W json.Number
    X json.RawMessage
    Y Marshaller 
    Z TextMarshaller
}


type S1 struct {
    A int
    B string
}

type Marshaller struct {
    v string
}

func (m *Marshaller) MarshalJSON() ([]byte, error) {
    return []byte(m.v), nil
}

func (m *Marshaller) UnmarshalJSON(data []byte) error {
    return json.Unmarshal(data, &m.v)
}

type TextMarshaller struct {
    v int
}

func (k *TextMarshaller) MarshalText() ([]byte, error) {
    return []byte(strconv.FormatInt(int64(k.v), 10)), nil
}

func (k *TextMarshaller)  UnmarshalText(data []byte) error {
    return json.Unmarshal(data, &k.v)
}

type TextMarshaller2 struct {
    v string
}

func (k *TextMarshaller2) MarshalText() ([]byte, error) {
    return []byte(k.v), errors.New("error")
}

func (k *TextMarshaller2)  UnmarshalText(data []byte) error {
    k.v = string(data)
    return errors.New("error")
}

const data = `{"A":1,"B1":"2","C":3,"D":true,"E":4,"G":5,"H":{"a":6},"I":{"a":"b"},"J":[7,8],"K":["a","b"],"L":{"A":9,"B":"10"},"M":{"A":11,"B":"12"},"N":13,"O":14,"P":"15","Q":"16","R":17,"T":{},"U":[19,20],"V":21,"W":"22","X":"23","Y":"24","Z":"25"}`


func TestErrorUnmarshalInvalidJSON(t *testing.T) {
    fields := []string{"B", "W", "X", "Y", "Z"}
    for _, field := range fields {
        root, err := Get([]byte(data))
        assert.Nil(t, err)
        _, err = root.Set(field, ast.NewRaw(`{]`))
        assert.Nil(t, err)
        buf, err := root.MarshalJSON()
        assert.Nil(t, err)

        var objj, objs S
        errj := json.Unmarshal(buf, &objj)
        errs := ConfigStd.Unmarshal(buf, &objs)
        // erri := jsoniter.Unmarshal(buf, &obji)
        // assert.Equal(t, errj!=nil, erri!=nil, "jsoniter:%s", string(buf))
        assert.Equal(t, errj!=nil, errs!=nil, "json:%s", string(buf))
        // assert.Equal(t, objj, obji, "jsoniter:%s", string(buf))
        assert.Equal(t, objj, objs, "json:%s", string(buf))
    }
}

func TestErrorUnmarshalDismatchedJSON(t *testing.T) {
    fields := []string{"L", "W", "X", "Y", "Z"}
    for _, field := range fields {
        root, err := Get([]byte(data))
        assert.Nil(t, err)
        _, err = root.Set(field, ast.NewRaw(`[1,2]`))
        assert.Nil(t, err)
        buf, err := root.MarshalJSON()
        assert.Nil(t, err)

        var objj, objs S
        errj := json.Unmarshal(buf, &objj)
        errs := ConfigStd.Unmarshal(buf, &objs)
        // erri := jsoniter.Unmarshal(buf, &obji)
        // assert.Equal(t, errj!=nil, erri!=nil, "jsoniter:%s", string(buf))
        assert.Equal(t, errj!=nil, errs!=nil, "json:%s", string(buf))
        // assert.Equal(t, objj, obji, "jsoniter:%s", string(buf))
        assert.Equal(t, objj, objs, "json:%s", string(buf))
    }
}

func TestErrorUnmarshalErrorUnmarshaler(t *testing.T) {
    var obje, objs, obji encoding.TextUnmarshaler = &TextMarshaller2{}, &TextMarshaller2{}, &TextMarshaller2{}

    data := []byte("{}")
    errj := json.Unmarshal(data, &obje)
    erri := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &obji)
    errs := ConfigStd.Unmarshal(data, &objs)
    assert.Equal(t, errj==nil, errs==nil)
    assert.Equal(t, errj==nil, erri==nil)
    assert.Equal(t, obje, objs)
    assert.Equal(t, obje, obji)
}

func TestErrorUnmarshalInvalidUTF8(t *testing.T) {
    var obje, objs, obji string

    data := []byte{'"', 0xFF, 0xFF, '"'}
    if utf8.Valid(data) {
        t.Errorf("%s is valid UTF-8", data)
    }
    errj := json.Unmarshal(data, &obje)
    erri := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &obji)
    errs := ConfigStd.Unmarshal(data, &objs)
    assert.Equal(t, errj==nil, errs==nil)
    assert.Equal(t, errj==nil, erri==nil)
    // assert.Equal(t, obje, objs)
    // assert.Equal(t, obje, obji)
}

func TestErrorMarshalInvalidUTF8(t *testing.T) {
    var obj string

    obj = string([]byte{0xFF, 0xFF})
    if utf8.Valid([]byte(obj)) {
        t.Errorf("%s is valid UTF-8", obj)
    }
    outj, errj := json.Marshal(&obj)
    outi, erri := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
    outs, errs := ConfigStd.Marshal(&obj)
    assert.Equal(t, errj==nil, errs==nil)
    assert.Equal(t, errj==nil, erri==nil)
    t.Log(outj, outs)
    t.Log(outj, outi)
}

func TestErrorMarshalInvalidJSON(t *testing.T) {
    t.Run("json.Number", func(t *testing.T) {
        var obj json.Number

        obj = json.Number("1.797693134862315708145274237317e+309")
        outj, errj := json.Marshal(&obj)
        outi, erri := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
        outs, errs := ConfigStd.Marshal(&obj)
        assert.Equal(t, errj==nil, errs==nil)
        assert.Equal(t, errj==nil, erri==nil)
        assert.Equal(t, outj, outs)
        assert.Equal(t, outj, outi)
    
        obj = json.Number(" [ invalid } ")
        outj, errj = json.Marshal(&obj)
        outi, erri = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
        outs, errs = ConfigStd.Marshal(&obj)
        assert.Equal(t, errj==nil, errs==nil)
        // assert.Equal(t, errj==nil, erri==nil)
        assert.Equal(t, outj, outs)
        // assert.Equal(t, outj, outi)
    })
    
    t.Run("json.RawMessage", func(t *testing.T) {
        var obj json.RawMessage

        obj = json.RawMessage(" [ invalid } ")
        outj, errj := json.Marshal(&obj)
        // outi, erri := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
        outs, errs := ConfigStd.Marshal(&obj)
        assert.Equal(t, errj==nil, errs==nil)
        // assert.Equal(t, errj==nil, erri==nil)
        assert.Equal(t, outj, outs)
        // assert.Equal(t, outj, outi)
    })

    t.Run("json.Marshaler", func(t *testing.T) {
        var obj json.Marshaler

        obj = &Marshaller{v: " [ invalid } "}
        outj, errj := json.Marshal(&obj)
        // outi, erri := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
        outs, errs := ConfigStd.Marshal(&obj)
        assert.Equal(t, errj==nil, errs==nil)
        // assert.Equal(t, errj==nil, erri==nil)
        assert.Equal(t, outj, outs)
        // assert.Equal(t, outj, outi)
    })
}

func TestErrorMarshalErrorMarshaler(t *testing.T) {
    var obj encoding.TextMarshaler

    obj = &TextMarshaller2{v: " [ invalid } "}
    outj, errj := json.Marshal(&obj)
    outi, erri := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&obj)
    outs, errs := ConfigStd.Marshal(&obj)
    assert.Equal(t, errj==nil, errs==nil)
    assert.Equal(t, errj==nil, erri==nil)
    assert.Equal(t, outj, outs)
    assert.Equal(t, outj, outi)
}