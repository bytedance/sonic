//go:build go1.24
//+build go1.24

/**
 * Copyright 2025 ByteDance Inc.
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 *     https://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package encoder

import (
    "testing"
    "time"
    "fmt"
    "strings"
)

type NonZeroStruct struct{}

func (nzs NonZeroStruct) IsZero() bool {
    return false
}

type NoPanicStruct struct {
    Int int `json:"int,omitzero"`
}

func (nps *NoPanicStruct) IsZero() bool {
    return nps.Int != 0
}

type isZeroer interface {
    IsZero() bool
}

type OptionalsZero struct {
    Sr string `json:"sr"`
    So string `json:"so,omitzero"`
    Sw string `json:"-"`

    Ir int `json:"omitzero"` // actually named omitzero, not an option
    Io int `json:"io,omitzero"`

    Slr       []string `json:"slr,random"`
    Slo       []string `json:"slo,omitzero"`
    SloNonNil []string `json:"slononnil,omitzero"`

    Mr  map[string]interface{} `json:"mr"`
    Mo  map[string]interface{} `json:",omitzero"`
    Moo map[string]interface{} `json:"moo,omitzero"`

    Fr   float64    `json:"fr"`
    Fo   float64    `json:"fo,omitzero"`
    Foo  float64    `json:"foo,omitzero"`
    Foo2 [2]float64 `json:"foo2,omitzero"`

    Br bool `json:"br"`
    Bo bool `json:"bo,omitzero"`

    Ur uint `json:"ur"`
    Uo uint `json:"uo,omitzero"`

    Str struct{} `json:"str"`
    Sto struct{} `json:"sto,omitzero"`

    Time      time.Time     `json:"time,omitzero"`
    TimeLocal time.Time     `json:"timelocal,omitzero"`
    Nzs       NonZeroStruct `json:"nzs,omitzero"`

    NilIsZeroer    isZeroer       `json:"niliszeroer,omitzero"`    // nil interface
    NonNilIsZeroer isZeroer       `json:"nonniliszeroer,omitzero"` // non-nil interface
    NoPanicStruct0 isZeroer       `json:"nps0,omitzero"`           // non-nil interface with nil pointer
    NoPanicStruct1 isZeroer       `json:"nps1,omitzero"`           // non-nil interface with non-nil pointer
    NoPanicStruct2 *NoPanicStruct `json:"nps2,omitzero"`           // nil pointer
    NoPanicStruct3 *NoPanicStruct `json:"nps3,omitzero"`           // non-nil pointer
    NoPanicStruct4 NoPanicStruct  `json:"nps4,omitzero"`           // concrete type
}

func TestOmitZero(t *testing.T) {
    // ForceUseVM()
    const want = `{
 "sr": "",
 "omitzero": 0,
 "slr": null,
 "slononnil": [],
 "mr": {},
 "Mo": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "nzs": {},
 "nps1": {},
 "nps3": {},
 "nps4": {}
}`
    var o OptionalsZero
    o.Sw = "something"
    o.SloNonNil = make([]string, 0)
    o.Mr = map[string]interface{}{}
    o.Mo = map[string]interface{}{}

    o.Foo = -0
    o.Foo2 = [2]float64{+0, -0}

    o.TimeLocal = time.Time{}.Local()

    o.NonNilIsZeroer = time.Time{}
    o.NoPanicStruct0 = (*NoPanicStruct)(nil)
    o.NoPanicStruct1 = &NoPanicStruct{}
    o.NoPanicStruct3 = &NoPanicStruct{}

    got, err := EncodeIndented(&o, "", " ", 0)
    if err != nil {
        t.Fatalf("MarshalIndent error: %v", err)
    }
    if got := string(got); got != want {
        t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
    }
}

func TestOmitZeroMap(t *testing.T) {
    const want = `{
 "foo": {
  "sr": "",
  "omitzero": 0,
  "slr": null,
  "mr": null,
  "fr": 0,
  "br": false,
  "ur": 0,
  "str": {},
  "nzs": {},
  "nps4": {}
 }
}`
    m := map[string]OptionalsZero{"foo": {}}
    got, err := EncodeIndented(m, "", " ", 0)
    if err != nil {
        t.Fatalf("MarshalIndent error: %v", err)
    }
    if got := string(got); got != want {
        fmt.Println(got)
        t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
    }
}

type OptionalsEmptyZero struct {
    Sr string `json:"sr"`
    So string `json:"so,omitempty,omitzero"`
    Sw string `json:"-"`

    Io int `json:"io,omitempty,omitzero"`

    Slr       []string `json:"slr,random"`
    Slo       []string `json:"slo,omitempty,omitzero"`
    SloNonNil []string `json:"slononnil,omitempty,omitzero"`

    Mr map[string]interface{} `json:"mr"`
    Mo map[string]interface{} `json:",omitempty,omitzero"`

    Fr float64 `json:"fr"`
    Fo float64 `json:"fo,omitempty,omitzero"`

    Br bool `json:"br"`
    Bo bool `json:"bo,omitempty,omitzero"`

    Ur uint `json:"ur"`
    Uo uint `json:"uo,omitempty,omitzero"`

    Str struct{} `json:"str"`
    Sto struct{} `json:"sto,omitempty,omitzero"`

    Time time.Time     `json:"time,omitempty,omitzero"`
    Nzs  NonZeroStruct `json:"nzs,omitempty,omitzero"`
}

func TestOmitEmptyZero(t *testing.T) {
    const want = `{
 "sr": "",
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "nzs": {}
}`
    var o OptionalsEmptyZero
    o.Sw = "something"
    o.SloNonNil = make([]string, 0)
    o.Mr = map[string]interface{}{}
    o.Mo = map[string]interface{}{}

    got, err := EncodeIndented(&o, "", " ", 0)
    if err != nil {
        t.Fatalf("MarshalIndent error: %v", err)
    }
    if got := string(got); got != want {
        t.Errorf("MarshalIndent:\n\tgot:  %s\n\twant: %s\n", indentNewlines(got), indentNewlines(want))
    }
}

func indentNewlines(s string) string {
    return strings.Join(strings.Split(s, "\n"), "\n\t")
}
