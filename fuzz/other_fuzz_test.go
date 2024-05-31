// +build go1.18

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sonic_fuzz

import (
	`bytes`
	`testing`
	`encoding/json`
	`unicode/utf8`

	`github.com/bytedance/sonic/encoder`
	`github.com/bytedance/sonic/decoder`
	`github.com/stretchr/testify/require`
	// `github.com/davecgh/go-spew/spew`
)

func fuzzValidate(t *testing.T, data []byte){
	jok1 := json.Valid(data)
	jok2 := utf8.Valid(data)
	_  = jok1 && jok2
	_, _ = encoder.Valid(data)
	// spew.Dump(data, jok1, jok2, sok)
	// require.Equalf(t, jok, sok, "different validate results")
}

func fuzzHtmlEscape(t *testing.T, data []byte){
	var jdst bytes.Buffer
	var sdst []byte
	json.HTMLEscape(&jdst, data)
	sdst = encoder.HTMLEscape(sdst, data)
	require.Equalf(t, string(jdst.Bytes()), string(sdst), "different htmlescape results")
}

// data is random, check whether is panic
func fuzzStream(t *testing.T, data []byte) {
	r := bytes.NewBuffer(data)
	dc := decoder.NewStreamDecoder(r)
	decoderEnableValidateString(dc)
	r1 := bytes.NewBuffer(data)
	dc1 := decoder.NewStreamDecoder(r1)

	w := bytes.NewBuffer(nil)
	ec := encoder.NewStreamEncoder(w)
	ec.SetCompactMarshaler(true)
	ec.SetValidateString(true)
	ec.SetEscapeHTML(true)
	ec.SortKeys()
	w1 := bytes.NewBuffer(nil)
	ec1 := encoder.NewStreamEncoder(w1)

	for dc1.More() {
		if !dc.More() {
			t.Fatal()
		}
		var obj interface{}
		err := dc.Decode(&obj)
		var obj1 interface{}
		err1 := dc1.Decode(&obj1)
		require.Equal(t, err1 == nil, err == nil)
		// require.Equal(t, obj, obj1)
		if err1 != nil {
			return
		}

		ee := ec.Encode(obj)
		ee1 := ec1.Encode(obj1)
		require.Equal(t, ee == nil, ee1 == nil)
	}
}