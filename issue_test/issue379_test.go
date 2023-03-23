package issue_test

import (
    `testing`
    `encoding/json`
    `github.com/bytedance/sonic`
    `github.com/stretchr/testify/require`
)

type Foo struct {
    Name string
}

func (f *Foo) UnmarshalJSON(data []byte) error {
    f.Name = "Unmarshaler"
    return nil
}

type MyPtr *Foo

func TestIssue379(t *testing.T) {
    tests := []struct{
        data  string
        newf  func() interface{} 
    } {
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { return &Foo{} },
        },
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { return MyPtr(&Foo{}) },
        },
        {
            data: `{"Name":"MyPtr"}`,
            newf:  func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
        },
        {
            data: `null`,
            newf:  func() interface{} { return MyPtr(&Foo{}) },
        },
        {
            data: `null`,
            newf:  func() interface{} { return &Foo{} },
        },
        {
            data: `null`,
            newf:  func() interface{} { ptr := MyPtr(&Foo{}); return &ptr },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]MyPtr) },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]*Foo) },
        },
        {
            data: `{"map":{"Name":"MyPtr"}}`,
            newf:  func() interface{} { return new(map[string]*MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]*MyPtr) },
        },
        {
            data: `[{"Name":"MyPtr"}]`,
            newf:  func() interface{} { return new([]*Foo) },
        },
    }

    for _, tt := range tests {
        jv, sv := tt.newf(), tt.newf()
        jerr := json.Unmarshal([]byte(tt.data), jv)
        serr := sonic.Unmarshal([]byte(tt.data), sv)
        require.Equal(t, jv, sv)
        require.Equal(t, jerr, serr)
    }
}
// *Foo TypeInfo:
//reflect.rtype {size: 8, ptrdata: 8, hash: 3423087479, tflag: tflagUncommon|tflagRegularMemory (9), align: 8, fieldAlign: 8, kind: 54, equal: runtime.memequal64, gcdata: *1, str: 48583, ptrToThis: 0}

// MyPtrType TypeInfo:
// reflect.rtype {size: 8, ptrdata: 8, hash: 3863140068, tflag: tflagUncommon|tflagExtraStar|tflagNamed|tflagRegularMemory (15), align: 8, fieldAlign: 8, kind: 54, equal: runtime.memequal64, gcdata: *1, str: 56864, ptrToThis: 135104}