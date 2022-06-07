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
	`os`
	`testing`
	`reflect`
	`fmt`
	`github.com/bytedance/sonic`
	`encoding/json`
	gojson `github.com/goccy/go-json`
	jsoniter `github.com/json-iterator/go`
	jsonv2 `github.com/go-json-experiment/json`
)

var (
	validFlag = os.Getenv("SONIC_VALID_GENERIC_BENCH")  != ""
	pretouchFlag = os.Getenv("SONIC_NO_PRETOUCH_BENCH") == ""
)

type jsonLibEntry struct {
	name      string
	marshal   func(any) ([]byte, error)
	unmarshal func([]byte, any) error
}

func sonicMarshalStd(v any) ([]byte, error) {
	return sonic.ConfigStd.Marshal(v)
}

func sonicUnmarshalStd(data []byte, v any) error {
	return sonic.ConfigStd.Unmarshal(data, v)
}

var jsonLibs = []jsonLibEntry {
	{"Std", json.Marshal, json.Unmarshal},
	{"StdV2", jsonv2.Marshal, jsonv2.Unmarshal},
	{"Sonic", sonic.Marshal, sonic.Unmarshal},
	{"SonicStd", sonicMarshalStd, sonicUnmarshalStd},
	{"GoJson", gojson.Marshal, gojson.Unmarshal},
	{"JsonIter", jsoniter.Marshal, jsoniter.Unmarshal},
	{"JsonIterStd", jsoniter.ConfigCompatibleWithStandardLibrary.Marshal, jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal},
}

func BenchmarkUnmarshalConcrete(b *testing.B)     { 
	runUnmarshalC(b)
}

func BenchmarkUnmarshalInterface(b *testing.B)     { 
	runUnmarshalI(b)
}

func BenchmarkMarshalConcrete(b *testing.B)     { 
	runMarshalC(b) 
}

func BenchmarkMarshalInterface(b *testing.B)     { 
	runMarshalI(b) 
}

func runUnmarshalC(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			var val any = tt.new()
			pretouch := func() {
				_ = lib.unmarshal(tt.data, val)
			}
	
			run := func(b *testing.B) {
				val = tt.new()
				if err := lib.unmarshal(tt.data, val); err != nil {
					b.Fatalf("%s Unmarshal error: %v", lib.name, err)
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
					b.Fatalf("%s Unmarshal output mismatch:\ngot  %v\nwant %v", lib.name, val2, val1)
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				if pretouchFlag {
					pretouch()
				}
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

func runUnmarshalI(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			var val any = tt.newI()
			pretouch := func() {
				_ = lib.unmarshal(tt.data, val)
			}
	
			run := func(b *testing.B) {
				val = tt.newI()
				if err := lib.unmarshal(tt.data, val); err != nil {
					b.Fatalf("%s Unmarshal error: %v", lib.name, err)
				}
			}

			valid := func(b *testing.B) {
				val1, val2 := tt.newI(), tt.newI()
				if err := json.Unmarshal(tt.data, val1); err != nil {
					panic(err)
				}
				if err := lib.unmarshal(tt.data, val2); err != nil {
					panic(err)
				}
				if !reflect.DeepEqual(val1, val2) {
					b.Fatalf("%s Unmarshal output mismatch:\ngot  %v\nwant %v", lib.name, val2, val1)
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				if pretouchFlag {
					pretouch()
				}
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

func runMarshalC(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			pretouch := func() {
				_, _ = lib.marshal(tt.val)
			}

			run := func(b *testing.B) {
				if _, err := lib.marshal(tt.val); err != nil {
					b.Fatalf("%s Marshal error: %v", lib.name, err)
				}
			}

			valid := func(b *testing.B) {
				// some details are different with encoding/json, so we compare the unmarshal results from marshaled buffer
				var buf1, buf2 []byte
				buf1, err := json.Marshal(tt.val)
				if err != nil {
					b.Fatalf("encoding/json Marshal error: %v", err)
				}
				buf2, err = lib.marshal(tt.val); 
				if err != nil {
					b.Fatalf("%s Marshal error: %v", lib.name, err)
				}
				var val1, val2 = tt.new(), tt.new()
				if err := json.Unmarshal(buf1, val1); err != nil {
					b.Fatalf("encoding/json Unmarshal again error: %v", err)
				}
				if err := lib.unmarshal(buf2, val2); err != nil {
					b.Fatalf("%s Unmarshal again error: %v", lib.name, err)
				}
				if !reflect.DeepEqual(val1, val2) {
					b.Fatalf("Unmarshal again output mismatch\n")
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				if pretouchFlag {
					pretouch()
				}
				if (validFlag) {
					valid(b)
				}
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

func runMarshalI(b *testing.B) {
	for _, tt := range jsonTestdata() {
		for _, lib := range jsonLibs {
			pretouch := func() {
				_, _ = lib.marshal(tt.valI)
			}

			run := func(b *testing.B) {
				if _, err := lib.marshal(tt.valI); err != nil {
					b.Fatalf("%s Marshal error: %v", lib.name, err)
				}
			}

			valid := func(b *testing.B) {
				// some details are different with encoding/json, so we compare the unmarshal results from marshaled buffer
				var buf1, buf2 []byte
				buf1, err := json.Marshal(tt.valI)
				if err != nil {
					b.Fatalf("encoding/json Marshal error: %v", err)
				}
				buf2, err = lib.marshal(tt.valI); 
				if err != nil {
					b.Fatalf("%s Marshal error: %v", lib.name, err)
				}
				var val1, val2 = tt.new(), tt.new()
				if err := json.Unmarshal(buf1, val1); err != nil {
					b.Fatalf("encoding/json Unmarshal again error: %v", err)
				}
				if err := lib.unmarshal(buf2, val2); err != nil {
					b.Fatalf("%s Unmarshal again error: %v", lib.name, err)
				}
				if !reflect.DeepEqual(val1, val2) {
					b.Fatalf("Unmarshal again output mismatch\n")
				}
			}

			name := fmt.Sprintf("%s_%s", tt.name, lib.name)
			b.Run(name, func(b *testing.B) {
				if pretouchFlag {
					pretouch()
				}
				if (validFlag) {
					valid(b)
				}
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