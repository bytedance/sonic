// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package issue_test

import (
    "fmt"
    "testing"

    "github.com/bytedance/sonic"
    "github.com/stretchr/testify/assert"
)

type UnmFoo struct {
    Name string
    Age  int
}

func (p *UnmFoo) UnmarshalJSON(data []byte) error {
    var aux struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }

    if err := sonic.Unmarshal(data, &aux); err != nil {
        return err
    }

    p.Name = aux.Name
    p.Age = aux.Age
    return nil
}

func TestIssue716(t *testing.T) {
    jsonData := `{"name": "Alice", "age": "30"}`
    var obj UnmFoo
    err := sonic.Unmarshal([]byte(jsonData), &obj)
    assert.Error(t, err)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
