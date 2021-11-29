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
    `encoding/json`
    `testing`
    `reflect`

    `github.com/stretchr/testify/require`
    .`github.com/bytedance/sonic`
)

type Issue141_Case_Insentive1 struct {
    Field0 int `json:"foo"`
    Field1 int `json:"FOO"`
    Field2 int `json:"FoO"`
    FOo    int
}

type Issue141_Case_Insentive2 struct {
    Field0 int `json:"FOO"`
    Field1 int `json:"foo"`
    Field2 int `json:"FoO"`
    FOo    int
}

type Issue141_Case_Insentive3 struct {
    FOo    int
    Field0 int `json:"FoO"`
    Field1 int `json:"foo"`
    Field2 int `json:"FOO"`
}

type Issue141_Case_Insentive4 struct {
    foo    int
    Field0 int `json:"FoO"`
    Field1 int `json:"foo"`
    Field2 int `json:"FOO"`
}

type  Issue141_Matched1 struct {
    Field0 int `json:"FOO"`
    Field1 int `json:"foo"`
    Field2 int `json:"FoO"`
    Foo    int
}

type  Issue141_Matched2 struct {
    Field0 int `json:"FOO"`
    Field1 int `json:"foo"`
    Field2 int `json:"FoO"`
    Foo    int
    Field3 int `json:"Foo"`
}

// Struct field priority in unmarshal, see https://go.dev/blog/json
func TestIssue141_StructFieldPriority(t *testing.T) {
    data := []byte("{\"Foo\":1}")
    for _, factory := range []func() interface{}{
        func() interface{} { return new(Issue141_Case_Insentive1) },
        func() interface{} { return new(Issue141_Case_Insentive2) },
        func() interface{} { return new(Issue141_Case_Insentive3) },
        func() interface{} { return new(Issue141_Case_Insentive4) },
        func() interface{} { return new(Issue141_Matched1) },
        func() interface{} { return new(Issue141_Matched2) },
    }{
        v1, v2 := factory(), factory()
        err1 := json.Unmarshal(data, &v1)
        err2 := Unmarshal(data, &v2)
        require.NoError(t, err1)
        require.NoError(t, err2)
        require.Equal(t, v1, v2)

        switch reflect.TypeOf(v2).Elem() {
        case reflect.TypeOf(Issue141_Case_Insentive1{}):
            println("Issue141_Case_Insentive1.Field0(tag foo) is ", v2.(*Issue141_Case_Insentive1).Field0)
        case reflect.TypeOf(Issue141_Case_Insentive2{}):
            println("Issue141_Case_Insentive2.Field0(tag FOO) is ", v2.(*Issue141_Case_Insentive2).Field0)
        case reflect.TypeOf(Issue141_Case_Insentive3{}):
            println("Issue141_Case_Insentive3.FOo is ", v2.(*Issue141_Case_Insentive3).FOo)
        case reflect.TypeOf(Issue141_Case_Insentive4{}):
            println("Issue141_Case_Insentive4.Field0(tag FoO) is ", v2.(*Issue141_Case_Insentive4).Field0)
        case reflect.TypeOf(Issue141_Matched1{}):
            println("Issue141_Matched1.Foo is ", v2.(*Issue141_Matched1).Foo)
        case reflect.TypeOf(Issue141_Matched2{}):
            println("Issue141_Matched2.Field3(tag Foo) is ", v2.(*Issue141_Matched2).Field3)
            println("Issue141_Matched2.Foo is ", v2.(*Issue141_Matched2).Foo)
        }
    }
}
