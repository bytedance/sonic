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

package alg

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/bytedance/sonic/testdata"
)

func BenchmarkU64toa(b *testing.B) {
	b.ReportAllocs()
	buf := make([]byte, 0, 64)
	for x :=0 ;x <= 62; x+=4  {
		d := 1<<x
		b.Run("sonic-"+strconv.Itoa(d), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = U64toa(buf, uint64(d))
			}
		})
		b.Run("std-"+strconv.Itoa(d), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strconv.AppendUint(buf, uint64(d), 10)
			}
		})
	}
}

func BenchmarkI64toa(b *testing.B) {
	b.ReportAllocs()
	buf := make([]byte, 0, 64)
	for x :=0 ;x <= 62; x+=4  {
		d := 1<<x
		b.Run("sonic-"+strconv.Itoa(d), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = I64toa(buf, int64(d))
			}
		})
		b.Run("std-"+strconv.Itoa(d), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strconv.AppendInt(buf, int64(d), 10)
			}
		})
	}
}

func BenchmarkF64toa(b *testing.B) {
	b.ReportAllocs()
	buf := make([]byte, 0, 64)
	for x :=0 ;x <= 62; x+=4  {
		d := 1<<x
		f := float64(d)+rand.Float64()
		b.Run("sonic-"+strconv.FormatFloat(f, 'g', -1, 64), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = F64toa(buf, f)
			}
		})
		b.Run("std-"+strconv.FormatFloat(f, 'g', -1, 64), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strconv.AppendFloat(buf, f, 'g', -1, 64)
			}
		})
	}
}

func BenchmarkF32toa(b *testing.B) {
	b.ReportAllocs()
	buf := make([]byte, 0, 64)
	for x :=0 ;x <= 30; x+=2  {
		d := 1<<x
		f := float32(d)+rand.Float32()
		b.Run("sonic-"+strconv.FormatFloat(float64(f), 'g', -1, 32), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = F32toa(buf, f)
			}
		})
		b.Run("std-"+strconv.FormatFloat(float64(f), 'g', -1, 32), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strconv.AppendFloat(buf, float64(f), 'g', -1, 32)
			}
		})
	}
}

func BenchmarkQuote(b *testing.B) {
	b.ReportAllocs()
	var runner = func(seed string) func(b *testing.B) {
		return func(b *testing.B) {
			buf := make([]byte, 0, len(seed)*1024*1024)
			for l := 1; l< cap(buf)*10; l*=10 {
				src := strings.Repeat(seed, l)
				b.Run("sonic-"+strconv.Itoa(len(src)), func(b *testing.B) {
					b.ResetTimer()
					for i:=0 ; i < b.N; i++ {
						_ = Quote(buf, src, false)
					}
				})
				b.Run("std-"+strconv.Itoa(len(src)), func(b *testing.B) {
					b.ResetTimer()
					for i:=0 ; i < b.N; i++ {
						_ = strconv.AppendQuote(buf, src)
					}
				})
			}
		}
	}
	
	b.Run("no quote", runner("abcdefghij"))
	b.Run("1/10 quote", runner("abcdefghi\n"))
	b.Run("1/5 quote", runner("abcd\nfghi\n"))
}

func BenchmarkValid(b *testing.B) {
	b.ReportAllocs()
	var runner = func(seed []byte) func(b *testing.B) {
		return func(b *testing.B) {
			b.Run("sonic", func(b *testing.B) {
				b.ResetTimer()
				for i:=0 ; i < b.N; i++ {
					_, _ = Valid(seed)
				}
			})
			b.Run("std", func(b *testing.B) {
				b.ResetTimer()
				for i:=0 ; i < b.N; i++ {
					_ = json.Valid(seed)
				}
			})
		}
	}
	b.Run("valid-small", runner([]byte(`{"a":1}`)))
	b.Run("invalid-small", runner([]byte(`{"a":1>`)))
	b.Run("valid-large", runner([]byte(testdata.TwitterJson)))
	b.Run("invalid-large", runner([]byte(strings.ReplaceAll(testdata.TwitterJson, "}", ">"))))
}

func BenchmarkEscapeHTML(b *testing.B) {
	b.ReportAllocs()
	var runner = func(seed []byte) func(b *testing.B) {
		return func(b *testing.B) {
			buf := make([]byte, 0, len(seed)*10)
				b.Run("sonic", func(b *testing.B) {
					b.ResetTimer()
					for i:=0 ; i < b.N; i++ {
						_ = HtmlEscape(buf, seed)
					}
				})
				b.Run("std", func(b *testing.B) {
					b.ResetTimer()
					for i:=0 ; i < b.N; i++ {
						bf := bytes.NewBuffer(buf)
						json.HTMLEscape(bf, seed)
					}
				})
		}
	}
	
	b.Run("small", runner([]byte(`{"a":"<>"}`)))
	b.Run("large", runner([]byte(testdata.TwitterJson)))
}