// +build amd64,go1.15,!go1.21

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

package sonic

import (
    `encoding/json`
    `reflect`
    `testing`

    `github.com/bytedance/sonic/decoder`
)

type useInt64Test struct {
    in  string
    out int64
}

type useFloatTest struct {
    in  string
    out float64
}

var useinttest = []useInt64Test{
    // int64
    {"0", 0},
    {"1", 1},
    {"-1", -1},
    {"100", 100},

    {"-9223372036854775807", -9223372036854775807},
    {"-9223372036854775808", -9223372036854775808}, //min int64
    {"9223372036854775807", 9223372036854775807},   //max int64
    {"9223372036854775806", 9223372036854775806},
}

var usefloattest = []useFloatTest{
    // float64
    {"-9223372036854775809", -9223372036854775809}, // int64 overflow
    {"9223372036854775808", 9223372036854775808},   // int64 overflow
    {"1e2", 1e2},
    {"1e-20", 1e-20},
    {"1.0", 1},
}

func TestUseInt64(t *testing.T) {
    for i, tt := range useinttest {
        var sout interface{}
        dc := decoder.NewDecoder(tt.in)
        dc.UseInt64()
        serr := dc.Decode(&sout)
        if !reflect.DeepEqual(sout, tt.out) {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sout, tt.in)
        }
        if serr != nil {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   nil\n", i, tt, serr)
        }
    }

    for i, tt := range usefloattest {
        var sout interface{}
        dc := decoder.NewDecoder(tt.in)
        dc.UseInt64()
        //the input string is not int64, still return float64
        serr := dc.Decode(&sout)
        if !reflect.DeepEqual(sout, tt.out) {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sout, tt.in)
        }
        if serr != nil {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   nil\n", i, tt, serr)
        }
    }
}

func TestUseNumber(t *testing.T) {
    for i, tt := range useinttest {
        var sout interface{}
        dc := decoder.NewDecoder(tt.in)
        dc.UseNumber()
        serr := dc.Decode(&sout)
        if !reflect.DeepEqual(sout, json.Number(tt.in)) {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sout, tt.out)
        }
        if serr != nil {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   nil\n", i, tt, serr)
        }
    }

    for i, tt := range usefloattest {
        var sout interface{}
        dc := decoder.NewDecoder(tt.in)
        dc.UseNumber()
        serr := dc.Decode(&sout)
        if !reflect.DeepEqual(sout, json.Number(tt.in)) {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sout, tt.out)
        }
        if serr != nil {
            t.Errorf("Test %d, %#v\ngot:\n   %#v\nexp:\n   nil\n", i, tt, serr)
        }
    }
}
