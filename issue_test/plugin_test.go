//go:build (linux && !race) || (unix && !race)
// +build linux,!race unix,!race

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
    `bytes`
    `fmt`
    `os/exec`
    `plugin`
    `reflect`
    `runtime`
    `testing`

    _ `github.com/bytedance/sonic`
)

func buildPlugin() {
    out := bytes.NewBuffer(nil)
    bin0, err := exec.LookPath("rm")
    if err != nil {
        panic(err)
    }
    cmd0 := exec.Cmd{
        Path:   bin0,
        Args:   []string{"rm", "-f", "plugin/plugin." + runtime.Version() + ".so"},
        Stdout: out,
        Stderr: out,
    }
    if err := cmd0.Run(); err != nil {
        panic(string(out.Bytes()))
    }
    out.Reset()
    bin, err := exec.LookPath("go")
    if err != nil {
        panic(err)
    }
    cmd := exec.Cmd{
        Path:   bin,
        Args:   []string{"go", "build", "-buildmode", "plugin", "-o", "plugin/plugin." + runtime.Version() + ".so", "plugin/main.go"},
        Stdout: out,
        Stderr: out,
    }
    if err := cmd.Run(); err != nil {
        panic(string(out.Bytes()))
    }
}

func TestPlugin(t *testing.T) {
    buildPlugin()
    p, err := plugin.Open("plugin/plugin." + runtime.Version() + ".so")
    if err != nil {
        t.Fatal(err)
    }
    v, err := p.Lookup("V")
    if err != nil {
        t.Fatal(err)
    }
    f, err := p.Lookup("F")
    if err != nil {
        t.Fatal(err)
    }
    *v.(*int) = 7
    f.(func())() // prints "Hello, number 7"
    obj, err := p.Lookup("Obj")
    m := *(obj.(*map[string]string))
    fmt.Printf("%#v\n", m)
    d, err := p.Lookup("Unmarshal")
    if err != nil {
        t.Fatal(err)
    }
    dec := d.(func(json string, val interface{}) error)
    var exp map[string]string
    if err := dec(`{"a":"b"}`, &exp); err != nil {
        t.Fatal(err)
    }
    if !reflect.DeepEqual(m, exp) {
        t.Fatal(m, exp)
    }
}
