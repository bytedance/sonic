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

package issue_test

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type User struct {
	Name string
	Age  int
}

var user = User{Name: "test", Age: 18}
var data []byte

// Benchmark Sonic serialization
func BenchmarkSonicMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark Sonic deserialization
func BenchmarkSonicUnmarshal(b *testing.B) {
	data, _ = sonic.Marshal(user)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var newUser User
		err := sonic.Unmarshal(data, &newUser)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark standard JSON serialization
func BenchmarkStandardMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark standard JSON deserialization
func BenchmarkStandardUnmarshal(b *testing.B) {
	data, _ = json.Marshal(user)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var newUser User
		err := json.Unmarshal(data, &newUser)
		if err != nil {
			b.Fatal(err)
		}
	}
}