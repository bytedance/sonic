package utf8


import (
	`testing`
    `strings`
	`github.com/stretchr/testify/assert`
	`unicode/utf8`
    `bytes`
    `math/rand`
)

var (
    _Header_2Bytes  = string([]byte{0xC0})
    _Header_3Bytes  = string([]byte{0xE0})
    _Header_4Bytes  = string([]byte{0xF0})
    _Low_Surrogate  = string([]byte{0xED, 0xA0, 0x80}) // \ud800
    _High_Surrogate = string([]byte{0xED, 0xB0, 0x80}) // \udc00
    _Cont           = "\xb0"
)

func TestCorrectWith_InvalidUtf8(t *testing.T) {
    var tests = []struct {
        name   string
        input  string
        expect string
        errpos int
    } {
        {"basic", `abc`, "abc", -1},
        {"long", strings.Repeat("helloα，景😊", 1000), strings.Repeat("helloα，景😊", 1000), -1},

        // invalid utf8 - single byte
        {"single_Cont", _Cont, "\ufffd", 0},
        {"single_Header_2Bytes", _Header_2Bytes, "\ufffd", 0},
        {"single_Header_3Bytes", _Header_3Bytes, "\ufffd", 0},
        {"single_Header_4Bytes", _Header_4Bytes, "\ufffd", 0},

        // invalid utf8 - two bytes
        {"two_Header_2Bytes + _Cont", _Header_2Bytes + _Cont, "\ufffd\ufffd", 0},
        {`two_Header_4Bytes + _Cont+ "xx"`, _Header_4Bytes + _Cont + "xx",  "\ufffd\ufffdxx", 0},
        { `"xx" + three_Header_4Bytes + _Cont + _Cont`, "xx" + _Header_4Bytes + _Cont + _Cont, "xx\ufffd\ufffd\ufffd", 2},

        // invalid utf8 - three bytes
        {`three_Low_Surrogate`, _Low_Surrogate, "\ufffd\ufffd\ufffd", 0},
        {`three__High_Surrogate`, _High_Surrogate, "\ufffd\ufffd\ufffd", 0},

        // invalid utf8 - multi bytes
        {`_High_Surrogate + _Low_Surrogate`, _High_Surrogate + _Low_Surrogate, "\ufffd\ufffd\ufffd\ufffd\ufffd\ufffd", 0},
        {`"\x80\x80\x80\x80"`, "\x80\x80\x80\x80", "\ufffd\ufffd\ufffd\ufffd", 0},
    }
    for _, test := range tests {
        got, err := CorrectWith(nil, []byte(test.input), "\ufffd")
        assert.Equal(t, []byte(test.expect), got, test.name)
        assert.Equal(t, err == nil, utf8.ValidString(test.input), test.name)
        if err != nil {
            assert.Equal(t, err.(*InvalidUTF8Error).Pos, test.errpos, test.name)
        }
    }
}

func genRandBytes(maxLen int) []byte {
    var buf bytes.Buffer
    length := rand.Intn(maxLen)
    for j := 0; j < length; j++ {
        buf.WriteByte(byte(rand.Intn(256)))
    }
    return buf.Bytes()
}

func genRandRune(maxLen int) []byte {
    var buf bytes.Buffer
    length := rand.Intn(maxLen)
    for j := 0; j < length; j++ {
        buf.WriteRune(rune(rand.Intn(0x10FFFF)))
    }
    return buf.Bytes()
}

func TestValidate_Random(t *testing.T) {
    // compare with stdlib
    compare := func(t *testing.T, data []byte) {
        assert.Equal(t, utf8.Valid(data), Validate(data), string(data))
    }

    // random testing
    nums   := 1000
    maxLen := 1000
    for i := 0; i < nums; i++ {
        compare(t, genRandBytes(maxLen))
        compare(t, genRandRune(maxLen))
    }
}
