package issue_test

import (
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

func TestDecodeGenericError(t *testing.T) {
    var v1 = map[string]interface{}{
        "a":"b",
    }
    err := json.Unmarshal(testErrBytes,&v1)
    assert.NotNil(t ,err)
    assert.Equal(t, map[string]interface{}{
        "a":"b",
    }, v1)
    t.Logf("%#v", v1)

    var v1j = map[string]interface{}{
        "a":"b",
    }
    err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(testErrBytes,&v1j)
    assert.NotNil(t ,err)
    assert.Equal(t, map[string]interface{}{
        "a":"b",
    }, v1j)
    t.Logf("%#v", v1j)

    var v1s = map[string]interface{}{
        "a":"b",
    }
    err = sonic.ConfigStd.Unmarshal(testErrBytes,&v1s)
    assert.NotNil(t ,err)
    assert.Equal(t, map[string]interface{}{
        "a":"b",
    }, v1s)
    t.Logf("%#v", v1s)

    var v2 interface{} 
    err = json.Unmarshal(testErrBytes,&v2)
    assert.NotNil(t,err)
    assert.Nil(t,v2)
    t.Logf("%#v",v2)

    var v2j interface{} 
    err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(testErrBytes, &v2j)
    assert.NotNil(t,err)
    assert.Nil(t,v2j)
    t.Logf("%#v",v2j)

    var v2s interface{} 
    err = sonic.ConfigStd.Unmarshal(testErrBytes, &v2s)
    assert.NotNil(t,err)
    assert.Nil(t,v2s)
    t.Logf("%#v",v2s)
}