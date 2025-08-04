//go:build (amd64 && go1.17 && !go1.26) || (arm64 && go1.20 && !go1.26)
// +build amd64,go1.17,!go1.26 arm64,go1.20,!go1.26

/*
 * Copyright 2022 ByteDance Inc.
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

package utf8

import (
    `testing`
    `strings`
    `github.com/stretchr/testify/assert`
    `unicode/utf8`
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
        got := CorrectWith(nil, []byte(test.input), "\ufffd")
        assert.Equal(t, []byte(test.expect), got, test.name)
        assert.Equal(t,test.errpos == -1, utf8.ValidString(test.input), test.name)
    }
}
