//go:build go1.18
// +build go1.18

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

package sonic_fuzz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func generateNullType() reflect.Type {
	tab := []reflect.Type {
		reflect.TypeOf(int(0)),
		reflect.TypeOf(uint(0)),
		reflect.TypeOf("string"),
		reflect.TypeOf(struct{}{}),
		reflect.TypeOf(json.Number("0")),
		reflect.TypeOf([]interface{}{}),
		reflect.TypeOf(map[string]interface{}{}),
	}
	return tab[int(rand.Int() % len(tab))]
}

func generateNumberType() reflect.Type {
	tab := []reflect.Type {
		reflect.TypeOf(float64(0)),
		reflect.TypeOf(int64(0)),
		reflect.TypeOf(uintptr(0)),
		reflect.TypeOf(uint64(0)),
		reflect.TypeOf(json.Number("0")),
	}
	return tab[int(rand.Int() % len(tab))]
}

func generatePointerType(ft reflect.Type) reflect.Type {
	if ft == nil {
		ft = generateNullType()
	}
	ftp := reflect.PtrTo(ft)
	ftpp := reflect.PtrTo(ftp)
	ftppp := reflect.PtrTo(ftpp)
	tab := []reflect.Type { ft, ftp, ftpp, ftppp }
	return tab[int(rand.Int() % len(tab))]
}

func generateJSONTag(name string) reflect.StructTag {
	var opt string
	name = strings.Split(name, ",")[0] // remove origin "," in tag name
	switch int(rand.Int() % 5) {
		case 0: return reflect.StructTag(`json:"-"`) // always omitted
		case 1: opt = "" // empty opt
		case 2: opt = "omitempty"
		// case 3: opt = "string"
		default: return reflect.StructTag("") // empty tag
	}
	return reflect.StructTag(fmt.Sprintf(`json:"%s,%s"`, name, opt))
}

// Map2StructType generate random dynamic Golang Struct by Golang Map
func Map2StructType(m map[string]interface{}, maxDepth int) reflect.Type {
	if maxDepth <= 0 {
		return reflect.TypeOf(map[string]interface{}{})
	}
	fields := make([]reflect.StructField, 0)
	i := 0
	for k, v := range m {
		/* skip the empty key */
		if k == "" {
			continue
		}

		/* set exported field name */
		fn := "F" + strconv.Itoa(i)

		/* set field type */
		ft := reflect.TypeOf(v)
		if ft == nil {
			ft = generateNullType()

		} else {
			switch ft.Kind() {
				case reflect.Map: ft = Map2StructType(v.(map[string]interface{}), maxDepth-1)
				case reflect.Float64: ft = generateNumberType()
			}
		}
		ft = generatePointerType(ft)
		ef := reflect.StructField {
			Name: fn,
			Type: ft,
			Tag : generateJSONTag(k),
		}
		fields = append(fields, ef)
		i += 1

		/* insert some private field randomly */
		if int(rand.Int() % 3) != 0 {
			continue
		}
		fn = "p" + strconv.Itoa(i)
		pf := reflect.StructField {
			Name: fn,
			Type: ft,
			PkgPath: "sonic_fuzz",
		}
		fields = append(fields, pf)

	}
	rt := reflect.StructOf(fields)
	return rt
}

const _MAX_STRUCT_DEPTH = 30


func isAscii(s string) bool {
	for i :=0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

var letters = []byte("_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genRandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return rt.Mem2Str(b)
}

func removeNonAscii(m map[string]interface{}) ([]byte, map[string]interface{}) {
	m2 := make(map[string]interface{}, len(m))
	for k, v := range m {
		if !isAscii(k) {
			// filled with random ascii
			m[genRandString(len(k))] = v
		}
		m2[k] = v
	}

	// marshal to json
	data2, err := json.Marshal(m2)
	if err != nil {
		panic("remashal failed")
	}
	return data2, m2
}

// fuzzDynamicStruct is schema-based fuzz testing, 
// a struct type is a JSON schema.
func fuzzDynamicStruct(t *testing.T, data []byte, v map[string]interface{}) {
	// for most case, tag is always ascii
	if rand.Intn(1000) % 3  != 0 {
		data, v = removeNonAscii(v)
	}
	typ  := Map2StructType(v, _MAX_STRUCT_DEPTH)
	sv  := reflect.New(typ).Interface()
	jv  := reflect.New(typ).Interface()

	// Pretouch fuzz
	err := sonic.Pretouch(typ)
	require.NoErrorf(t, err, "error in sonic pretouch struct %v", typ)

	// Unmarshal fuzz
	serr := target.Unmarshal(data, &sv)
	jerr := json.Unmarshal(data, &jv)
	require.Equalf(t, serr != nil, jerr != nil, "different error in sonic unmarshal %v", typ, spew.Sdump(serr), spew.Sdump(jerr), string(data))
	if serr != nil {
		return
	}
	require.Equal(t, jv, sv, "different result in sonic unmarshal %v", typ, string(data))

	// Marshal fuzz
	sout, serr := target.Marshal(sv)
	jout, jerr := json.Marshal(jv)
	require.NoError(t, serr, "error in sonic marshal %v", typ)
	require.NoError(t, jerr, "error in json marshal %v", typ)
	// not comparing here because sonic marshal is different from encoding/json, as:
	// require.Equalf(t, sout, jout, "different in sonic marshal %#v", typ)
	var _, _ = sout, jout
}