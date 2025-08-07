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
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestIssue825_IntegerTypeBoundaries(t *testing.T) {
	// Test different integer types with boundary values
	type testCase struct {
		name     string
		input    string
		value    interface{}
		expected interface{}
		shouldError bool
	}
	
	testCases := []testCase{
		// uint8 tests
		{"uint8_valid", "255", new(uint8), uint8(255), false},
		{"uint8_overflow", "256", new(uint8), uint8(0), true},
		{"uint8_large", "9223372036854775807", new(uint8), uint8(0), true},
		
		// uint16 tests
		{"uint16_valid", "65535", new(uint16), uint16(65535), false},
		{"uint16_overflow", "65536", new(uint16), uint16(0), true},
		{"uint16_large", "9223372036854775807", new(uint16), uint16(0), true},
		
		// uint32 tests (the main issue)
		{"uint32_valid", "4294967295", new(uint32), uint32(4294967295), false},
		{"uint32_overflow", "4294967296", new(uint32), uint32(0), true},
		{"uint32_large", "9223372036854775807", new(uint32), uint32(0), true},
		
		// uint64 tests
		{"uint64_valid", "9223372036854775807", new(uint64), uint64(9223372036854775807), false},
		{"uint64_max", "18446744073709551615", new(uint64), uint64(18446744073709551615), false},
		
		// int8 tests
		{"int8_valid", "127", new(int8), int8(127), false},
		{"int8_overflow", "128", new(int8), int8(0), true},
		{"int8_negative", "-128", new(int8), int8(-128), false},
		{"int8_negative_overflow", "-129", new(int8), int8(0), true},
		
		// int16 tests
		{"int16_valid", "32767", new(int16), int16(32767), false},
		{"int16_overflow", "32768", new(int16), int16(0), true},
		{"int16_negative", "-32768", new(int16), int16(-32768), false},
		{"int16_negative_overflow", "-32769", new(int16), int16(0), true},
		
		// int32 tests
		{"int32_valid", "2147483647", new(int32), int32(2147483647), false},
		{"int32_overflow", "2147483648", new(int32), int32(0), true},
		{"int32_negative", "-2147483648", new(int32), int32(-2147483648), false},
		{"int32_negative_overflow", "-2147483649", new(int32), int32(0), true},
		
		// int64 tests
		{"int64_valid", "9223372036854775807", new(int64), int64(9223372036854775807), false},
		{"int64_negative", "-9223372036854775808", new(int64), int64(-9223372036854775808), false},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test Sonic
			err := sonic.Unmarshal([]byte(tc.input), tc.value)
			sonicHasError := err != nil
			
			// Test standard library
			errStd := json.Unmarshal([]byte(tc.input), tc.value)
			stdHasError := errStd != nil
			
			fmt.Printf("Test: %s\n", tc.name)
			fmt.Printf("  Input: %s\n", tc.input)
			fmt.Printf("  Sonic error: %v\n", err)
			fmt.Printf("  Std error: %v\n", errStd)
			fmt.Printf("  Error match: %v == %v\n", sonicHasError, stdHasError)
			fmt.Println()
			
			// Check error behavior consistency
			require.Equal(t, stdHasError, sonicHasError,
				"Error behavior mismatch for %s: Sonic error=%v, Std error=%v", 
				tc.input, sonicHasError, stdHasError)
		})
	}
} 