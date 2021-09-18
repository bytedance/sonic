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

package issue_test

import (
    . `github.com/bytedance/sonic`
	`fmt`
	`reflect`
	_ `sync`
	`testing`
	`unsafe`

	stdjson `encoding/json`
)

func TestLargeMapValue(t *testing.T) {
    var jsonStr = `{
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {}
	}`
    type Case struct {
        std interface{}
        sonic interface{}
    }
    cases := []Case{
        {&map[string]TestIssue100_LargeMapValue{}                         , &map[string]TestIssue100_LargeMapValue{}},
        {&map[int32]TestIssue100_LargeMapValue{}                          , &map[int32]TestIssue100_LargeMapValue{}},
        {&map[int64]TestIssue100_LargeMapValue{}                          , &map[int64]TestIssue100_LargeMapValue{}},
        {&map[uint32]TestIssue100_LargeMapValue{}                         , &map[uint32]TestIssue100_LargeMapValue{}},
        {&map[uint64]TestIssue100_LargeMapValue{}                         , &map[uint64]TestIssue100_LargeMapValue{}},
        {&map[TestIssue100_textMarshalKey]TestIssue100_LargeMapValue{}    , &map[TestIssue100_textMarshalKey]TestIssue100_LargeMapValue{}},
        {&map[TestIssue100_textMarshalKeyPtr]TestIssue100_LargeMapValue{} , &map[TestIssue100_textMarshalKeyPtr]TestIssue100_LargeMapValue{}},
    }
    for i, c := range cases {
        var stdw, sonicw = c.std, c.sonic
        if err := stdjson.Unmarshal([]byte(jsonStr), stdw); err != nil {
            t.Fatal(i, err)
        }
        fmt.Printf("[%d]struct size: %d\tmap length: %d\n", i, unsafe.Sizeof(TestIssue100_LargeMapValue{}), reflect.ValueOf(stdw).Elem().Len())
        if err := Unmarshal([]byte(jsonStr), sonicw); err != nil {
            t.Fatal(err)
        }
        if !reflect.DeepEqual(stdw, sonicw) {
            fmt.Printf("have:\n\t%#v\nwant:\n\t%#v\n", sonicw, stdw)
            t.Fatal(i)
        }
    }
}

type TestIssue100_textMarshalKey string

func(self TestIssue100_textMarshalKey) UnmarshalText(text []byte) error {
    _ = TestIssue100_textMarshalKey(text)
    return nil
}

type TestIssue100_textMarshalKeyPtr string

func(self *TestIssue100_textMarshalKeyPtr) UnmarshalText(text []byte) error {
    *self = TestIssue100_textMarshalKeyPtr(text)
    return nil
}

type TestIssue100_LargeMapValue struct {
    Id [129]byte
}