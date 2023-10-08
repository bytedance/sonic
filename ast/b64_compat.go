// +build !amd64

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
    `encoding/base64`
)

func init() {
    println("WARNING: sonic only supports Go1.16~1.20 && CPU amd64, but your environment is not suitable")
}

func decodeBase64(src string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64.StdEncoding.EncodeToString(src)
}
