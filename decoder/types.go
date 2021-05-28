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

package decoder

import (
    `encoding`
    `encoding/base64`
    `encoding/json`
    `reflect`
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var (
    byteType                = reflect.TypeOf(byte(0))
    jsonNumberType          = reflect.TypeOf(json.Number(""))
    base64CorruptInputError = reflect.TypeOf(base64.CorruptInputError(0))
)

var (
    errorType                   = reflect.TypeOf((*error)(nil)).Elem()
    jsonUnmarshalerType         = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
    encodingTextUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

func rtype(t reflect.Type) (*rt.GoItab, *rt.GoType) {
    p := (*rt.GoIface)(unsafe.Pointer(&t))
    return p.Itab, (*rt.GoType)(p.Value)
}
