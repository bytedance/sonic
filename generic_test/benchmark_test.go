// +build go1.18

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

package generic_test

import(
	`testing`
	`reflect`
	`fmt`
	`github.com/bytedance/sonic`
	`github.com/bytedance/sonic/encoder`
	`github.com/bytedance/sonic/decoder`
	`github.com/bytedance/sonic/internal/rt`
	`encoding/json`
	gojson `github.com/goccy/go-json`
	jsoniter `github.com/json-iterator/go`
	jsonv2 `github.com/go-json-experiment/json`
)

type jsonLibEntry struct {
	name string
	marshal  func(any) ([]byte, error)
	unmarshal func([]byte, any) error
}

func sonicMarshalStd(v any) ([]byte, error) {
	return encoder.Encode(v, encoder.CompatibleWithStd)
}

func sonicUnmarshalCopy(data []byte, v any) error {
	dc := decoder.NewDecoder(rt.Mem2Str(data))
	dc.CopyString()
	return dc.Decode(&v)
}

var jsonLibs = []jsonLibEntry {
	{"Std", json.Marshal, json.Unmarshal},
	{"JsonV2", jsonv2.Marshal, jsonv2.Unmarshal},
	{"Sonic", sonic.Marshal, sonic.Unmarshal},
	{"SonicStd", sonicMarshalStd, sonicUnmarshalCopy},
	{"GoJson", gojson.Marshal, gojson.Unmarshal},
	{"JsonIter", jsoniter.Marshal, jsoniter.Unmarshal},
}

func BenchmarkUnmarshal(b *testing.B)     { 
	runUnmarshal(b)
}

func BenchmarkMarshal(b *testing.B)     { 
	runMarshal(b) 
}

func runUnmarshal(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			var val any = tt.new()
			pretouch := func() {
				_ = lib.unmarshal(tt.data, val)
			}
	
			run := func(b *testing.B) {
				val = tt.new()
				if err := lib.unmarshal(tt.data, val); err != nil {
					b.Fatalf("Unmarshal error: %v", err)
				}
			}

			valid := func(b *testing.B) {
				val1, val2 := tt.new(), tt.new()
				if err := json.Unmarshal(tt.data, val1); err != nil {
					panic(err)
				}
				if err := lib.unmarshal(tt.data, val2); err != nil {
					panic(err)
				}
				if !reflect.DeepEqual(val1, val2) {
					b.Fatalf("Unmarshal output mismatch:\ngot  %v\nwant %v", val2, val1)
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				pretouch()
				valid(b)
				b.ResetTimer()
				b.ReportAllocs()
				b.SetBytes(int64(len(tt.data)))
				for i := 0; i < b.N; i++ {
					run(b)
				}
			})
		}
	}
}

func runMarshal(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			pretouch := func() {
				_, _ = lib.marshal(tt.val)
			}

			run := func(b *testing.B) {
				if _, err := lib.marshal(tt.val); err != nil {
					b.Fatalf("Marshal error: %v", err)
				}
			}

			valid := func(b *testing.B) {
				// some details are different with encoding/json, so we compare the unmarshal results from marshaled buffer
				var stdbuf, sonicbuf []byte
				stdbuf, err := lib.marshal(tt.val)
				if err != nil {
					b.Fatalf("encoding/json Marshal error: %v", err)
				}
				sonicbuf, err = json.Marshal(tt.val); 
				if err != nil {
					b.Fatalf("sonic Marshal error: %v", err)
				}
				var stdv, sonicv any = tt.new(), tt.new()
				if err := json.Unmarshal(stdbuf, stdv); err != nil {
					b.Fatalf("encoding/json Unmarshal again error: %v", err)
				}
				if err := lib.unmarshal(sonicbuf, sonicv); err != nil {
					b.Fatalf("sonic Unmarshal again error: %v", err)
				}
				if !reflect.DeepEqual(stdv, sonicv) {
					b.Fatalf("Unmarshal again output mismatch\n")
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				pretouch()
				valid(b)
				b.ResetTimer()
				b.ReportAllocs()
				b.SetBytes(int64(len(tt.data)))
				for i := 0; i < b.N; i++ {
					run(b)
				}
			})
		}
	}
}