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