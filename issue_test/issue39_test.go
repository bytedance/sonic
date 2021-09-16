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
    `testing`

    `github.com/bytedance/sonic/decoder`
    `github.com/stretchr/testify/require`
)

type normalIfaceIssue39 interface {
    Foo()
}

type normalWrapIssue39 struct {
    F normalIfaceIssue39
}

type normalImplIssue39 struct {
    X int
}

func (_ *normalImplIssue39) Foo() {}

type jsonIfaceIssue39 interface {
    UnmarshalJSON(b []byte) error
}

type jsonWrapIssue39 struct {
    F jsonIfaceIssue39
}

type jsonImplIssue39 struct {
    a string
}

func (self *jsonImplIssue39) UnmarshalJSON(b []byte) error{
    self.a = string(b)
    return nil
}

type textIfaceIssue39 interface {
    UnmarshalText(b []byte) error
}

type textWrapIssue39 struct {
    F textIfaceIssue39
}

type textImplIssue39 struct {
    a string
}

func (self *textImplIssue39) UnmarshalText(b []byte) error{
    self.a = string(b)
    return nil
}

func TestIssue39_Iface(t *testing.T) {
    p := new(normalImplIssue39)
    obj := normalWrapIssue39{F: p}
    err := Unmarshal([]byte(`{"F":{"X":123}}`), &obj)
    if err != nil {
        if v, ok := err.(decoder.SyntaxError); ok {
            println(v.Description())
        }
        require.NoError(t, err)
    }
    require.Equal(t, 123, p.X)
}

func TestIssue39_UnmarshalJSON(t *testing.T) {
    p := &jsonImplIssue39{}
    obj := jsonWrapIssue39{F: p}
    err := Unmarshal([]byte(`{"F":"xx"}`), &obj)
    if err != nil {
        if v, ok := err.(decoder.SyntaxError); ok {
            println(v.Description())
        }
        require.NoError(t, err)
    }
    require.Equal(t, `"xx"`, p.a)
}

func TestIssue39_UnmarshalText(t *testing.T) {
    p := &textImplIssue39{}
    obj := textWrapIssue39{F: p}
    err := Unmarshal([]byte(`{"F":"xx"}`), &obj)
    if err != nil {
        if v, ok := err.(decoder.SyntaxError); ok {
            println(v.Description())
        }
        require.NoError(t, err)
    }
    require.Equal(t, `xx`, p.a)
}
