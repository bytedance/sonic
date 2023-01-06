package issue_test

import (
    `encoding/json`
    `fmt`
    `reflect`
    `strconv`
    `testing`
    `time`

    `github.com/bytedance/sonic`
    `github.com/davecgh/go-spew/spew`
    `github.com/stretchr/testify/require`
)

func TestPointerValueRecurseMarshal(t *testing.T) {
    info := &TestStruct1{
        StartTime: JSONTime(time.Now()),
    }
    infos := &[]*TestStruct1{info}

    bytes, err1 := json.Marshal(infos)
    fmt.Printf("%+v\n", string(bytes))
    spew.Dump(bytes, err1)

    jbytes, err2 := sonic.Marshal(infos)
    fmt.Printf("%+v\n", string(jbytes))
    spew.Dump(jbytes, err2)
    require.Equal(t, bytes, jbytes)
}

func TestPointerValueRecursePretouch(t *testing.T) {
    info := &TestStruct2{
        StartTime: JSONTime(time.Now()),
    }
    infos := &[]*TestStruct2{info}

    bytes, err1 := json.Marshal(infos)
    fmt.Printf("%+v\n", string(bytes))
    spew.Dump(bytes, err1)

    sonic.Pretouch(reflect.TypeOf(infos))
    jbytes, err2 := sonic.Marshal(infos)
    fmt.Printf("%+v\n", string(jbytes))
    spew.Dump(jbytes, err2)
    require.Equal(t, bytes, jbytes)
}

type TestStruct1 struct {
    StartTime JSONTime
}

type TestStruct2 struct {
    StartTime JSONTime
}

type JSONTime time.Time

func (t *JSONTime) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatInt(time.Time(*t).Unix(), 10)), nil
}