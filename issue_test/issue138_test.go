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
    `fmt`
    `time`
    `testing`
    `reflect`

    `github.com/stretchr/testify/require`
    `github.com/bytedance/sonic`
    `github.com/bytedance/sonic/option`
)

type Issue138_DeepStruct struct {
    L0 struct {
    L1 struct {
    L2 struct {
    L3 struct {
    L4 struct {
    L5 struct {
    L6 struct {
    L7 struct {
    L8 struct {
    L9 struct {
    L10 struct {
    L11 struct {
    L12 struct {
    L13 struct {
    L14 struct {
    L15 struct {
    L16 struct {
    L17 struct {
    L18 struct {
    L19 struct {
    L20 struct {
    L21 struct {
    L22 struct {
        A int
        B string
        C []float64
        E map[string]bool
        F *Issue138_DeepStruct
    }}}}}}}}}}}}}}}}}}}}}}}
}

func testPretouchTime(depth int) {
    start := time.Now()
    sonic.Pretouch(reflect.TypeOf(Issue138_DeepStruct{}), option.WithCompileRecursiveDepth(depth))
    elapsed := time.Since(start)
    fmt.Printf("Pretouch with recursive depth %d, time is %s\n", depth, elapsed)
}

func TestIssue138_PretouchTime(t *testing.T) {
    testPretouchTime(4)
    var obj Issue138_DeepStruct
    start := time.Now()
    data, err := sonic.Marshal(obj)
    err = sonic.Unmarshal([]byte(data), &obj)
    elapsed := time.Since(start)
    fmt.Printf("Marshal and unmarshal time is %s\n", elapsed)
    require.NoError(t, err)
}

