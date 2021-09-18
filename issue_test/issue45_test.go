/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
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
    `sync`
    `testing`
)

func ExtractJson(idx int, body interface{}) {
    j := []byte(exampleJson[idx])

    // REPLACE sonic WITH json
    err := Unmarshal(j, &body)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
}

var exampleJson1 = `{}`

var exampleJson2 = `{}`

var exampleJson = []string{exampleJson1, exampleJson2}

type raceTestStruct struct {
    f1 *[]raceTestStruct2 `json:"f1"`
    f2 *int               `json:"f2"`
    f3 *string            `json:"f3"`
    f4 *string            `json:"f4"`
}
type raceTestStruct2 struct {
    g1 *string           `json:"g1"`
    g2 *string           `json:"g2"`
    g3 *string           `json:"g3"`
    g4 []raceTestStruct3 `json:"g4"`
}
type raceTestStruct3 struct {
    e1 *string  `json:"e1"`
    e2 *string  `json:"e2"`
    e3 *float64 `json:"e3"`
    e4 *float64 `json:"e4"`
}

func TestExtracJson(t *testing.T) {

    wg := sync.WaitGroup{}

    resultChan := make(chan raceTestStruct, 2)

    wg.Add(1)
    go func() {
        defer wg.Done()
        var model raceTestStruct
        ExtractJson(0, &model)
        resultChan <- model
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        var model raceTestStruct
        ExtractJson(1, &model)
        resultChan <- model
    }()

    var results []raceTestStruct
    for i := 0; i < 2; i++ {
        results = append(results, <-resultChan)
    }

    wg.Wait()
    close(resultChan)
}