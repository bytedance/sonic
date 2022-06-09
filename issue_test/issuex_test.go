package issue_test

import (
    `bytes`
    `encoding/json`
    `testing`

    `github.com/bytedance/sonic`
    jsoniter `github.com/json-iterator/go`
    `github.com/stretchr/testify/assert`
)


var (
    testErrStr   = `{"a":[1,2,3,4,],"b":"1"}`
    testErrBytes = []byte(testErrStr)
)

func TestUnmarshal(t *testing.T) {
    var v interface{} 

    d := jsoniter.NewDecoder(bytes.NewBuffer(testErrBytes))
    d.UseNumber()
    err := d.Decode(&v)
    assert.NotNil(t,err)
    assert.NotNil(t,v)
    t.Logf("%#v",v)

    e := json.NewDecoder(bytes.NewBuffer(testErrBytes))
    d.UseNumber()
    err = e.Decode(&v)
    assert.NotNil(t,err)
    assert.NotNil(t,v)
    t.Logf("%#v",v)

    err = sonic.ConfigStd.Unmarshal(testErrBytes, &v)
    assert.NotNil(t,err)
    assert.NotNil(t,v)
    t.Logf("%#v",v)
}