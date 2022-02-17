package issue_test

import (
    `bytes`
    `fmt`
    `os/exec`
    `plugin`
    `runtime`
    `strings`
    `testing`
    `reflect`

    _ `github.com/bytedance/sonic`
)

func init() {
    if !strings.Contains(runtime.Version(), "go1.16") {
        return
    }
    bin, err := exec.LookPath("go")
    if err != nil {
        panic(err)
    }
    out := bytes.NewBuffer(nil)
    cmd := exec.Cmd{
        Path: bin,
        Args: []string{"go", "build", "-buildmode", "plugin", "-o", "plugin/plugin."+runtime.Version()+".so", "plugin/main.go"},
        Stdout: out,
        Stderr: out,
    }
    if err := cmd.Run(); err != nil {
        panic(string(out.Bytes()))
    }
}

func TestPlugin(t *testing.T) {
    if !strings.Contains(runtime.Version(), "go1.16") {
        return
    }
    p, err := plugin.Open("plugin/plugin."+runtime.Version()+".so")
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
