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

// corpus returns the simple and basic JSON corpus for fuzzing test.
func corpus() [][]byte {
	data := []string {
		`[]`, `{}`, `[{`, `}`, `{{}}`, `,`, `:`,// structural chars
		`null`, `true`, `false`, `truf`,  // primitive types
		`1.234567890e-123`, `01`, `00`, "+1", // numbers
		`e`, `-`, `+`, `.`, // signs
		" ", "\n", "\t", "\r",  // space
		"\b",  "\f",  "\\", "/", "\"", "\u2028", "\x00", // unescaped chars
		"\\b", "\\n", "\\f", "\\\\", "\\/", "\\\"", "\\r", "\\t", "\\u2028", // escaped chars
		"<", ">", "&", "\u2028", "\u2029", // html escape
		`😁`, "\xff", "\xf0", "\x80", // utf-8
		"\xed\xa0\x80" /* \ud800 */, "\xef\xbf\xbf", /* \uffff */ "\xed\xbf\xbf", /* \udfff */
		`"haha"`, `"你好"`, `"😁"`, `"\\uD800\\udc00"`,  `""`, // json strings
		"\"\u2028\u2029\"", `<>&`, 
		"\"\xff\"", "\"\x00\"", `"\\uDFFF"`,
		`[2, 3, null, true, false, "hi"]`, // short json
		`{
			"object": {
				"slice": [
					1,
					2.0,
					"3",
					[4],
					{5: {}}
				]
			},
			"slice": [[]],
			"string": ":)",
			"int": 1e5,
			"float": 3e-9"
		}`,
		`{"a":{"a":1,"b":[1,1,1],"c":{"d":1,"e":1,"f":1},"d":"{\"你好\":\"hi\"}"}}`,
	}
	var corpus [][]byte
	for _, t := range(data) {
		corpus = append(corpus, []byte(t))
	}
	return corpus
}