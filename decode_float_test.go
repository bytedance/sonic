// +build amd64,go1.15,!go1.21

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

package sonic

import (
    `encoding/json`
    `reflect`
    `strings`
    `testing`

    `github.com/bytedance/sonic/decoder`
)

type atofTest struct {
    in  string
    out string
    err error
}

// Tests from Go strconv package, https://github.com/golang/go/blob/master/src/strconv/atof_test.go
// All tests are passed in Go encoding/json.
var atoftests = []atofTest{
    {"1.234e", "", nil}, // error
    {"1i", "1", nil},    // pass
    {"1", "1", nil},
    {"1e23", "1e+23", nil},
    {"1E23", "1e+23", nil},
    {"100000000000000000000000", "1e+23", nil},
    {"1e-100", "1e-100", nil},
    {"123456700", "1.234567e+08", nil},
    {"99999999999999974834176", "9.999999999999997e+22", nil},
    {"100000000000000000000001", "1.0000000000000001e+23", nil},
    {"100000000000000008388608", "1.0000000000000001e+23", nil},
    {"100000000000000016777215", "1.0000000000000001e+23", nil},
    {"100000000000000016777216", "1.0000000000000003e+23", nil},
    {"-1", "-1", nil},
    {"-0.1", "-0.1", nil},
    {"-0", "-0", nil},
    {"1e-20", "1e-20", nil},
    {"625e-3", "0.625", nil},

    // zeros
    {"0", "0", nil},
    {"0e0", "0", nil},
    {"-0e0", "-0", nil},
    {"0e-0", "0", nil},
    {"-0e-0", "-0", nil},
    {"0e+0", "0", nil},
    {"-0e+0", "-0", nil},
    {"0e+01234567890123456789", "0", nil},
    {"0.00e-01234567890123456789", "0", nil},
    {"-0e+01234567890123456789", "-0", nil},
    {"-0.00e-01234567890123456789", "-0", nil},

    {"0e291", "0", nil}, // issue 15364
    {"0e292", "0", nil}, // issue 15364
    {"0e347", "0", nil}, // issue 15364
    {"0e348", "0", nil}, // issue 15364
    {"-0e291", "-0", nil},
    {"-0e292", "-0", nil},
    {"-0e347", "-0", nil},
    {"-0e348", "-0", nil},

    // largest float64
    {"1.7976931348623157e308", "1.7976931348623157e+308", nil},
    {"-1.7976931348623157e308", "-1.7976931348623157e+308", nil},

    // the border is ...158079
    // borderline - okay
    {"1.7976931348623158e308", "1.7976931348623157e+308", nil},
    {"-1.7976931348623158e308", "-1.7976931348623157e+308", nil},

    // a little too large
    {"1e308", "1e+308", nil},

    // denormalized
    {"1e-305", "1e-305", nil},
    {"1e-306", "1e-306", nil},
    {"1e-307", "1e-307", nil},
    {"1e-308", "1e-308", nil},
    {"1e-309", "1e-309", nil},
    {"1e-310", "1e-310", nil},
    {"1e-322", "1e-322", nil},
    // smallest denormal
    {"5e-324", "5e-324", nil},
    {"4e-324", "5e-324", nil},
    {"3e-324", "5e-324", nil},
    // too small
    {"2e-324", "0", nil},
    // way too small
    {"1e-350", "0", nil},
    {"1e-400000", "0", nil},

    // try to overflow exponent
    {"1e-4294967296", "0", nil},
    {"1e-18446744073709551616", "0", nil},

    // https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
    {"2.2250738585072012e-308", "2.2250738585072014e-308", nil},
    // https://www.exploringbinary.com/php-hangs-on-numeric-value-2-2250738585072011e-308/
    {"2.2250738585072011e-308", "2.225073858507201e-308", nil},

    // A very large number (initially wrongly parsed by the fast algorithm).
    {"4.630813248087435e+307", "4.630813248087435e+307", nil},

    // A different kind of very large number.
    {"22.222222222222222", "22.22222222222222", nil},
    {"2." + strings.Repeat("2", 800) + "e+1", "22.22222222222222", nil},

    // Exactly halfway between 1 and math.Nextafter(1, 2).
    // Round to even (down).
    {"1.00000000000000011102230246251565404236316680908203125", "1", nil},
    // Slightly lower; still round down.
    {"1.00000000000000011102230246251565404236316680908203124", "1", nil},
    // Slightly higher; round up.
    {"1.00000000000000011102230246251565404236316680908203126", "1.0000000000000002", nil},
    // Slightly higher, but you have to read all the way to the end.
    {"1.00000000000000011102230246251565404236316680908203125" + strings.Repeat("0", 10000) + "1", "1.0000000000000002", nil},

    // Halfway between x := math.Nextafter(1, 2) and math.Nextafter(x, 2)
    // Round to even (up).
    {"1.00000000000000033306690738754696212708950042724609375", "1.0000000000000004", nil},

    // Halfway between 1090544144181609278303144771584 and 1090544144181609419040633126912
    // (15497564393479157p+46, should round to even 15497564393479156p+46, issue 36657)
    {"1090544144181609348671888949248", "1.0905441441816093e+30", nil},

    // Corner case between int64 and float64 for the input
    {"9223372036854775807", "9223372036854775807", nil}, // max int64: (1 << 63) - 1
    {"9223372036854775808", "9223372036854775808", nil},
    {"-9223372036854775808", "-9223372036854775808", nil}, // min int64: 1 << 63
    {"-9223372036854775809", "-9223372036854775809", nil},
}

func TestDecodeFloat(t *testing.T) {
    for i, tt := range atoftests {
        // default float64
        var sonicout, stdout float64
        sonicerr := decoder.NewDecoder(tt.in).Decode(&sonicout)
        stderr := json.NewDecoder(strings.NewReader(tt.in)).Decode(&stdout)
        if !reflect.DeepEqual(sonicout, stdout) {
            t.Fatalf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sonicout, stdout)
        }
        if !reflect.DeepEqual(sonicerr == nil, stderr == nil) {
            t.Fatalf("Test %d, %#v\ngot:\n   %#v\nexp:\n   %#v\n", i, tt.in, sonicerr, stderr)
        }
    }
}
