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
    `encoding/json`
    `strings`
    `testing`

    `github.com/stretchr/testify/require`

    `github.com/bytedance/sonic`
    `github.com/bytedance/sonic/decoder`
)

var issue_19x_idata = "\"" + strings.Repeat("9", 1000) + "\""
var issue_19x_fdata = "\"" + strings.Repeat("9", 100) + "." + strings.Repeat("9", 1000) + "\""
var issue_19x_ndata = strings.Repeat("9", 1000)
var issue_19x_invalid = strings.Repeat("9", 100) + "abc99"

func TestDecodeLongStringToJsonNumber(t *testing.T) {
    var objs, obje json.Number 
    errs := sonic.UnmarshalString(issue_19x_idata, &objs)
    erre := json.Unmarshal([]byte(issue_19x_idata), &obje)
    require.Equal(t, erre, errs)
    require.Equal(t, obje, objs)

    var fobjs, fobje json.Number 
    errs = sonic.UnmarshalString(issue_19x_fdata, &fobjs)
    erre = json.Unmarshal([]byte(issue_19x_fdata), &fobje)
    require.Equal(t, erre, errs)
    require.Equal(t, fobje, fobjs)

    var objs2, obje2 json.Number 
    errs = sonic.UnmarshalString(issue_19x_invalid, &objs2)
    erre = json.Unmarshal([]byte(issue_19x_invalid), &obje2)
    require.NotNil(t, erre)
    require.NotNil(t, errs)


    var iobjs, iobje interface{}
    dc := decoder.NewDecoder(issue_19x_ndata)
    dc.UseNumber()
    errs = dc.Decode(&iobjs)
    r := json.NewDecoder(bytes.NewBufferString(issue_19x_ndata))
    r.UseNumber()
    erre = r.Decode(&iobje)
    require.Equal(t, erre, errs)
    require.Equal(t, iobje, iobjs)

    var iobjs2, iobje2 interface{}
    dc = decoder.NewDecoder(issue_19x_invalid)
    dc.UseNumber()
    errs = dc.Decode(&iobjs2)
    r = json.NewDecoder(bytes.NewBufferString(issue_19x_invalid))
    r.UseNumber()
    erre = r.Decode(&iobje2)
    require.Equal(t, erre, errs)
    require.Equal(t, iobje2, iobjs2)
    // spew.Dump(iobje2)
}

var jsonNumberBig = "\"" + strings.Repeat("9", 10) + "." + strings.Repeat("9", 100) + "\""

func BenchmarkDecodeJsonNumber_Sonic(b *testing.B) {
    b.SetBytes(int64(len(jsonNumberBig)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        var obj json.Number 
        _ = sonic.UnmarshalString(jsonNumberBig, &obj)
    }
}

func BenchmarkDecodeUseNumber_Sonic(b *testing.B) {
    b.SetBytes(int64(len(jsonNumberBig)))
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        var obj interface{} 
        dc := decoder.NewDecoder(jsonNumberBig)
        dc.UseNumber()
        _ = dc.Decode(&obj)
    }
}