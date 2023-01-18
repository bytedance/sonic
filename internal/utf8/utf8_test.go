package utf8


import (
	`testing`
	`github.com/davecgh/go-spew/spew`
	`github.com/stretchr/testify/assert`
	`unicode/utf8`
)


var (
    _Header_2Bytes  = string([]byte{0xC0})
    _Header_3Bytes  = string([]byte{0xE0})
    _Header_4Bytes  = string([]byte{0xF0})
    _Low_Surrogate  = string([]byte{0xED, 0xA0, 0x80}) // \ud800
    _High_Surrogate = string([]byte{0xED, 0xB0, 0x80}) // \udc00
    _Cont           = "\xb0"
)

func TestDecoder_InvalidUtf8(t *testing.T) {
    var quote = func(s string) string {
        return `"` + s + `"`
    };
    var tests = []struct {
        name string
        input string
        expect string
        hasError bool
        opt     int
    } {
        {"basic", `"abc"`, "abc", false, 0},

        // invalid utf8 - single byte
        {"single_Cont", quote(_Cont), "\ufffd", false, 0},
        {"single_Header_2Bytes", quote(_Header_2Bytes), "\ufffd", false, 0},
        {"single_Header_3Bytes", quote(_Header_3Bytes), "\ufffd", false, 0},
        {"single_Header_4Bytes", quote(_Header_4Bytes), "\ufffd", false, 0},

        // invalid utf8 - two bytes
        {"two_Header_2Bytes + _Cont", quote(_Header_2Bytes + _Cont), "\ufffd\ufffd", false, 0},
        {`two_Header_4Bytes + _Cont+ "xx"`, quote(_Header_4Bytes + _Cont + "xx"),  "\ufffd\ufffdxx", false, 0},
        {`three_Header_4Bytes + _Cont + _Cont + "xx"`, quote(_Header_4Bytes + _Cont + _Cont + "xx"), "\ufffd\ufffd\ufffdxx", false, 0},

        // invalid utf8 - three bytes
        {`three_Low_Surrogate`, quote(_Low_Surrogate), "\ufffd\ufffd\ufffd", false, 0},
        {`three__High_Surrogate`, quote(_High_Surrogate), "\ufffd\ufffd\ufffd", false, 0},

        // invalid utf8 - multi bytes
        {`_High_Surrogate + _Low_Surrogate`, quote(_High_Surrogate + _Low_Surrogate), "\ufffd\ufffd\ufffd\ufffd\ufffd\ufffd", false, 0},
        {`"\x80\x80\x80\x80"`, quote("\x80\x80\x80\x80"), "\ufffd\ufffd\ufffd\ufffd", false, 0},

    }
    s := "\xf0\x80\x80\x80"
    var _ = s
    for _, test := range tests {
        spew.Dump(test)
        // check validutf8
        {
            got, err := CorrectWith(nil, []byte(test.input), "\ufffd")
            assert.Equal(t, quote(test.expect), string(got), test.name)
            assert.Equal(t, err == nil, utf8.ValidString(test.input), test.name)
        }
    }
}
