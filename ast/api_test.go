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

package ast

import (
    `testing`
    `strings`

    `github.com/stretchr/testify/assert`
)

type Path = []interface{}

type testGetApi struct {
    json      string
    path      Path
}

type checkError func(error) bool

func isSyntaxError(err error) bool {
    if err == nil {
        return false
    }
    return strings.HasPrefix(err.Error(), `"Syntax error at index`)
}

func isEmptySource(err error) bool {
    if err == nil {
        return false
    }
    return strings.Contains(err.Error(), "no sources available")
}

func isErrNotExist(err error) bool {
    return err == ErrNotExist
}

func isErrUnsupportType(err error) bool {
    return err == ErrUnsupportType
}

func testSyntaxJson(t *testing.T, json string, path ...interface{}) {
    search := NewSearcher(json)
    _, err := search.GetByPath(path...)
    assert.True(t, isSyntaxError(err))
}

func TestGetFromEmptyJson(t *testing.T) {
    tests := []testGetApi {
        { "", nil },
        { "", Path{}},
        { "", Path{""}},
        { "", Path{0}},
        { "", Path{"", ""}},
    }
    for _, test := range tests {
        f := func(t *testing.T) {
            search := NewSearcher(test.json)
            _, err := search.GetByPath(test.path...)
            assert.True(t, isEmptySource(err))
        }
        t.Run(test.json, f)
    }
}

func TestGetFromSyntaxError(t *testing.T) {
    tests := []testGetApi {
        { " \r\n\f\t", Path{} },
        { "123.", Path{} },
        { "+124", Path{} },
        { "-", Path{} },
        { "-e123", Path{} },
        { "-1.e123", Path{} },
        { "-12e456.1", Path{} },
        { "-12e.1", Path{} },
        { "[", Path{} },
        { "{", Path{} },
        { "[}", Path{} },
        { "{]", Path{} },
        { "{,}", Path{} },
        { "[,]", Path{} },
        { "tru", Path{} },
        { "fals", Path{} },
        { "nul", Path{} },
        { `{"a":"`, Path{"a"} },
        { `{"`, Path{} },
        { `"`, Path{} },
        { `"\"`, Path{} },
        { `"\\\"`, Path{} },
        { `"hello`, Path{} },
        { `{{}}`, Path{} },
        { `{[]}`, Path{} },
        { `{:,}`, Path{} },
        { `{test:error}`, Path{} },
        { `{":true}`, Path{} },
        { `{"" false}`, Path{} },
        { `{ "" : "false }`, Path{} },
        { `{"":"",}`, Path{} },
        { `{ " test : true}`, Path{} },
        { `{ "test" : tru }`, Path{} },
        { `{ "test" : true , }`, Path{} },
        { `{ {"test" : true , } }`, Path{} },
        { `{"test":1. }`, Path{} },
        { `{"\\\""`, Path{} },
        { `{"\\\"":`, Path{} },
        { `{"\\\":",""}`, Path{} },
        { `[{]`, Path{} },
        { `[tru]`, Path{} },
        { `[-1.]`, Path{} },
        { `[[]`, Path{} },
        { `[[],`, Path{} },
        { `[ true , false , [ ]`, Path{} },
        { `[true, false, [],`, Path{} },
        { `[true, false, [],]`, Path{} },
        { `{"key": [true, false, []], "key2": {{}}`, Path{} },
    }

    for _, test := range tests {
        f := func(t *testing.T) {
            testSyntaxJson(t, test.json, test.path...)
            path := append(Path{"key"}, test.path...)
            testSyntaxJson(t, `{"key":` + test.json, path...)
            path  = append(Path{""}, test.path...)
            testSyntaxJson(t, `{"":` + test.json, path...)
            path  = append(Path{1}, test.path...)
            testSyntaxJson(t, `["",` + test.json, path...)
        }
        t.Run(test.json, f)
    }
}

// NOTE: GetByPath API not validate the undemanded fields for performance.
func TestGetWithInvalidUndemandedField(t *testing.T) {
    type Any = interface{}
    tests := []struct {
        json string
        path Path
        exp  Any
    } {
        { "-0xyz", Path{}, Any(float64(-0))},
        { "-12e4xyz", Path{}, Any(float64(-12e4))},
        { "truex",  Path{}, Any(true)},
        { "false,", Path{}, Any(false)},
        { `{"a":{,xxx},"b":true}`, Path{"b"}, Any(true)},
        { `{"a":[,xxx],"b":true}`, Path{"b"}, Any(true)},
    }

    for _, test := range tests {
        f := func(t *testing.T) {
            search := NewSearcher(test.json)
            node, err := search.GetByPath(test.path...)
            assert.NoError(t, err)
            v, err := node.Interface()
            assert.NoError(t, err)
            assert.Equal(t, v, test.exp)
        }
        t.Run(test.json, f)
    }
}

func TestGet_InvalidPathType(t *testing.T) {
    assert.Panics(t, assert.PanicTestFunc(func() {
        data := `{"a":[{"b":true}]}`
        s := NewSearcher(data)
        s.GetByPath("a", true)

        s = NewSearcher(data)
        s.GetByPath("a", nil)

        s = NewSearcher(data)
        s.GetByPath("a", -1)
    }))
}
