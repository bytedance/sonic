package sonic_fuzz

import (
	`bytes`
	`testing`
	`encoding/json`
	`unicode/utf8`

	`github.com/bytedance/sonic/encoder`
	`github.com/stretchr/testify/require`
	`github.com/davecgh/go-spew/spew`
)

func fuzzValidate(t *testing.T, data []byte){
	jok1 := json.Valid(data)
	jok2 := utf8.Valid(data)
	jok  := jok1 && jok2
	sok, _ := encoder.Valid(data)
	spew.Dump(data, jok1, jok2, sok)
	require.Equalf(t, jok, sok, "different validate results")
}

func fuzzHtmlEscape(t *testing.T, data []byte){
	var jdst bytes.Buffer
	var sdst []byte
	json.HTMLEscape(&jdst, data)
	sdst = encoder.HTMLEscape(sdst, data)
	require.Equalf(t, string(jdst.Bytes()), string(sdst), "different htmlescape results")
}