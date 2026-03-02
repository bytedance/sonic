package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue775_OutOfRangeIntegerInByteSlice(t *testing.T) {
	cas := unmTestCase{
		name: "byte slice with invalid integer items",
		data: []byte(`{"ByteArr": [1, 2999, -123, 3.14, 3]}`),
		newfn: func() interface{} {
			var v struct {
				ByteArr []uint8
			}
			return &v
		},
	}

	assertUnmarshal(t, sonic.ConfigDefault, cas)
	assertUnmarshal(t, sonic.ConfigStd, cas)
}

func TestIssue775_OutOfRangeIntegerInOtherSlices(t *testing.T) {
	cases := []unmTestCase{
		{
			name: "int8 slice with invalid items",
			data: []byte(`{"V": [1, 128, -129, 2]}`),
			newfn: func() interface{} {
				var v struct {
					V []int8
				}
				return &v
			},
		},
		{
			name: "int16 slice with invalid items",
			data: []byte(`{"V": [1, 32768, -32769, 2]}`),
			newfn: func() interface{} {
				var v struct {
					V []int16
				}
				return &v
			},
		},
		{
			name: "int32 slice with invalid items",
			data: []byte(`{"V": [1, 2147483648, -2147483649, 2]}`),
			newfn: func() interface{} {
				var v struct {
					V []int32
				}
				return &v
			},
		},
		{
			name: "uint16 slice with invalid items",
			data: []byte(`{"V": [1, 65536, -1, 2]}`),
			newfn: func() interface{} {
				var v struct {
					V []uint16
				}
				return &v
			},
		},
		{
			name: "uint32 slice with invalid items",
			data: []byte(`{"V": [1, 4294967296, -1, 2]}`),
			newfn: func() interface{} {
				var v struct {
					V []uint32
				}
				return &v
			},
		},
	}

	for _, cas := range cases {
		assertUnmarshal(t, sonic.ConfigDefault, cas)
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}

func TestIssue775_NonSliceIntegerMismatchShouldContinue(t *testing.T) {
	cases := []unmTestCase{
		{
			name: "int8 overflow should still decode sibling field",
			data: []byte(`{"A":128,"B":2}`),
			newfn: func() interface{} {
				var v struct {
					A int8
					B int
				}
				return &v
			},
		},
		{
			name: "uint8 underflow should still decode sibling field",
			data: []byte(`{"A":-1,"B":2}`),
			newfn: func() interface{} {
				var v struct {
					A uint8
					B int
				}
				return &v
			},
		},
		{
			name: "int16 overflow should still decode sibling field",
			data: []byte(`{"A":32768,"B":2}`),
			newfn: func() interface{} {
				var v struct {
					A int16
					B int
				}
				return &v
			},
		},
	}

	for _, cas := range cases {
		assertUnmarshal(t, sonic.ConfigDefault, cas)
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}

func TestIssue775_MapKeyIntegerMismatchShouldContinue(t *testing.T) {
	cases := []unmTestCase{
		{
			name: "map[int8]int key overflow should still decode sibling field",
			data: []byte(`{"M":{"1":1,"128":2,"2":3},"T":9}`),
			newfn: func() interface{} {
				var v struct {
					M map[int8]int
					T int
				}
				return &v
			},
		},
		{
			name: "map[uint8]int key underflow should still decode sibling field",
			data: []byte(`{"M":{"1":1,"-1":2,"2":3},"T":9}`),
			newfn: func() interface{} {
				var v struct {
					M map[uint8]int
					T int
				}
				return &v
			},
		},
	}

	for _, cas := range cases {
		assertUnmarshal(t, sonic.ConfigDefault, cas)
		assertUnmarshal(t, sonic.ConfigStd, cas)
	}
}
