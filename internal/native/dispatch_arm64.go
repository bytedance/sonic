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

package native

import (
	"os"
	"unsafe"
	"bufio"
	"strings"
	"github.com/shirou/gopsutil/cpu"
	neon "github.com/bytedance/sonic/internal/native/neon"
	sve_linkname "github.com/bytedance/sonic/internal/native/sve_linkname"
	"github.com/bytedance/sonic/internal/native/sve_wrapgoc"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)



var (
	__GetByPath func(s unsafe.Pointer, p unsafe.Pointer, path unsafe.Pointer, m unsafe.Pointer) int

	__ParseWithPadding func(parser unsafe.Pointer) (ret int)

	__I64toa func(out unsafe.Pointer, val int64) (ret int)

	__U64toa func(out unsafe.Pointer, val uint64) (ret int)

	__F64toa func(out unsafe.Pointer, val float64) (ret int)

	__F32toa func(out unsafe.Pointer, val float32) (ret int)

	__Quote func(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn unsafe.Pointer, flags uint64) int

	__LookupSmallKey func(key unsafe.Pointer, table unsafe.Pointer, lowerOff int) (index int)

	__SkipOne func(s unsafe.Pointer, p unsafe.Pointer, m unsafe.Pointer, flags uint64) int

	__SkipOneFast func(s unsafe.Pointer, p unsafe.Pointer) int
)

const (
	MaxFrameSize   uintptr = 200
	BufPaddingSize int     = 64
)

var (
	S_f64toa uintptr
	S_f32toa uintptr
	S_i64toa uintptr
	S_u64toa uintptr
	S_lspace uintptr
)

var (
	S_quote   uintptr
	S_unquote uintptr
)

var (
	S_value     uintptr
	S_vstring   uintptr
	S_vnumber   uintptr
	S_vsigned   uintptr
	S_vunsigned uintptr
)

var (
	S_skip_one    uintptr
	S_skip_one_fast    uintptr
	S_get_by_path    uintptr
	S_skip_array  uintptr
	S_skip_object uintptr
	S_skip_number uintptr
	S_parse_with_padding uintptr
	S_lookup_small_key uintptr
)
var UseSveWrapgoc bool
var UseSveLinkname bool

//go:nosplit
func GetByPathSveWrapgoc(s *string, p *int, path *[]interface{}, m *types.StateMachine) int {
	return __GetByPath(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(path)), rt.NoEscape(unsafe.Pointer(m)))
}

//go:nosplit
//go:noescape
//go:linkname GetByPathNeonLinkname github.com/bytedance/sonic/internal/native/neon.__get_by_path
func GetByPathNeonLinkname(s *string, p *int, path *[]interface{}, m *types.StateMachine) int

//go:nosplit
//go:noescape
//go:linkname GetByPathSveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__get_by_path
func GetByPathSveLinkname(s *string, p *int, path *[]interface{}, m *types.StateMachine) int

//go:nosplit
func GetByPath(s *string, p *int, path *[]interface{}, m *types.StateMachine) int {
	if UseSveWrapgoc {
		return GetByPathSveWrapgoc(s, p, path, m)
	} else if UseSveLinkname {
		return GetByPathSveLinkname(s, p, path, m)
	} else {
		return GetByPathNeonLinkname(s, p, path, m)
	}
}

//go:nosplit
func ParseWithPaddingSveWrapgoc(parser unsafe.Pointer) (ret int) {
	return __ParseWithPadding(rt.NoEscape(unsafe.Pointer(parser)))
}

//go:nosplit
//go:noescape
//go:linkname ParseWithPaddingNeonLinkname github.com/bytedance/sonic/internal/native/neon.__parse_with_padding
func ParseWithPaddingNeonLinkname(parser unsafe.Pointer) (ret int)

//go:nosplit
//go:noescape
//go:linkname ParseWithPaddingSveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__parse_with_padding
func ParseWithPaddingSveLinkname(parser unsafe.Pointer) (ret int)

//go:nosplit
func ParseWithPadding(parser unsafe.Pointer) (ret int) {
	if UseSveWrapgoc {
		return ParseWithPaddingSveWrapgoc(parser)
	} else if UseSveLinkname {
		return ParseWithPaddingSveLinkname(parser)
	} else {
		return ParseWithPaddingNeonLinkname(parser)
	}
}

//go:nosplit
//go:noescape
//go:linkname F64toaNeonLinkname github.com/bytedance/sonic/internal/native/neon.__f64toa
func F64toaNeonLinkname(out *byte, val float64) (ret int)

//go:nosplit
func F64toaSveWrapgoc(out *byte, val float64) (ret int) {
	return __F64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func F64toa(out *byte, val float64) (ret int) {
	if UseSveWrapgoc {
		return F64toaSveWrapgoc(out, val)
	} else {
		return F64toaNeonLinkname(out, val)
	}
}

//go:nosplit
//go:noescape
//go:linkname F32toaNeonLinkname github.com/bytedance/sonic/internal/native/neon.__f32toa
func F32toaNeonLinkname(out *byte, val float32) (ret int)

//go:nosplit
func F32toaSveWrapgoc(out *byte, val float32) (ret int) {
	return __F32toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
func F32toa(out *byte, val float32) (ret int) {
	if UseSveWrapgoc {
		return F32toaSveWrapgoc(out, val)
	} else {
		return F32toaNeonLinkname(out, val)
	}
}

//go:nosplit
func I64toaSveWrapgoc(out *byte, val int64) (ret int) {
	return __I64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
//go:noescape
//go:linkname I64toaNeonLinkname github.com/bytedance/sonic/internal/native/neon.__i64toa
func I64toaNeonLinkname(out *byte, val int64) (ret int)

//go:nosplit
func I64toa(out *byte, val int64) (ret int) {
	if UseSveWrapgoc {
		return I64toaSveWrapgoc(out, val)
	} else {
		return I64toaNeonLinkname(out, val)
	}
}

//go:nosplit
func U64toaSveWrapgoc(out *byte, val uint64) (ret int) {
	return __U64toa(rt.NoEscape(unsafe.Pointer(out)), val)
}

//go:nosplit
//go:noescape
//go:linkname U64toaNeonLinkname github.com/bytedance/sonic/internal/native/neon.__u64toa
func U64toaNeonLinkname(out *byte, val uint64) (ret int)

//go:nosplit
func U64toa(out *byte, val uint64) (ret int) {
	if UseSveWrapgoc {
		return U64toaSveWrapgoc(out, val)
	} else {
		return U64toaNeonLinkname(out, val)
	}
}

//go:nosplit
func QuoteSveWrapgoc(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int {
	return __Quote(rt.NoEscape(unsafe.Pointer(s)), nb, rt.NoEscape(unsafe.Pointer(dp)), rt.NoEscape(unsafe.Pointer(dn)), flags)
}

//go:nosplit
//go:noescape
//go:linkname QuoteNeonLinkname github.com/bytedance/sonic/internal/native/neon.__quote
func QuoteNeonLinkname(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int

//go:nosplit
func Quote(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int, flags uint64) int {
	if UseSveWrapgoc {
		return QuoteSveWrapgoc(s, nb, dp, dn, flags)
	} else {
		return QuoteNeonLinkname(s, nb, dp, dn, flags)
	}
}

//go:nosplit
func LookupSmallKeySveWrapgoc(key *string, table *[]byte, lowerOff int) (index int) {
	return __LookupSmallKey(rt.NoEscape(unsafe.Pointer(key)), rt.NoEscape(unsafe.Pointer(table)), lowerOff)
}

//go:nosplit
//go:noescape
//go:linkname LookupSmallKeySveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__lookup_small_key
func LookupSmallKeySveLinkname(key *string, table *[]byte, lowerOff int) (index int)

//go:nosplit
//go:noescape
//go:linkname LookupSmallKeyNeonLinkname github.com/bytedance/sonic/internal/native/neon.__lookup_small_key
func LookupSmallKeyNeonLinkname(key *string, table *[]byte, lowerOff int) (index int)

//go:nosplit
func LookupSmallKey(key *string, table *[]byte, lowerOff int) (index int) {
	if UseSveWrapgoc {
		return LookupSmallKeySveWrapgoc(key, table, lowerOff)
	} else if UseSveLinkname {
		return LookupSmallKeySveLinkname(key, table, lowerOff)
	} else {
		return LookupSmallKeyNeonLinkname(key, table, lowerOff)
	}
}

//go:nosplit
func SkipOneSveWrapgoc(s *string, p *int, m *types.StateMachine, flags uint64) int {
	return __SkipOne(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(m)), flags)
}

//go:nosplit
//go:noescape
//go:linkname SkipOneNeonLinkname github.com/bytedance/sonic/internal/native/neon.__skip_one
func SkipOneNeonLinkname(s *string, p *int, m *types.StateMachine, flags uint64) int

//go:nosplit
//go:noescape
//go:linkname SkipOneSveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__skip_one
func SkipOneSveLinkname(s *string, p *int, m *types.StateMachine, flags uint64) int

//go:nosplit
func SkipOne(s *string, p *int, m *types.StateMachine, flags uint64) int {
	if UseSveWrapgoc {
		return SkipOneSveWrapgoc(s, p, m, flags)
	} else if UseSveLinkname {
		return SkipOneSveLinkname(s, p, m, flags)
	} else {
		return SkipOneNeonLinkname(s, p, m, flags)
	}
}

//go:nosplit
func SkipOneFastSveWrapgoc(s *string, p *int) int {
	return __SkipOneFast(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)))
}

//go:nosplit
//go:noescape
//go:linkname SkipOneFastNeonLinkname github.com/bytedance/sonic/internal/native/neon.__skip_one_fast
func SkipOneFastNeonLinkname(s *string, p *int) int

//go:nosplit
//go:noescape
//go:linkname SkipOneFastSveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__skip_one_fast
func SkipOneFastSveLinkname(s *string, p *int) int

//go:nosplit
func SkipOneFast(s *string, p *int) int {
	if UseSveWrapgoc {
		return SkipOneFastSveWrapgoc(s, p)
	} else if UseSveLinkname {
		return SkipOneFastSveLinkname(s, p)
	} else {
		return SkipOneFastNeonLinkname(s, p)
	}
}

//go:nosplit
//go:noescape
//go:linkname HTMLEscapeSveLinkname github.com/bytedance/sonic/internal/native/sve_linkname.__html_escape
func HTMLEscapeSveLinkname(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int

//go:nosplit
//go:noescape
//go:linkname HTMLEscapeNeonLinkname github.com/bytedance/sonic/internal/native/neon.__html_escape
func HTMLEscapeNeonLinkname(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int

//go:nosplit
func HTMLEscape(s unsafe.Pointer, nb int, dp unsafe.Pointer, dn *int) int {
	if UseSveLinkname {
		return HTMLEscapeSveLinkname(s, nb, dp, dn)
	} else {
		return HTMLEscapeNeonLinkname(s, nb, dp, dn)
	}
}

//go:nosplit
//go:noescape
//go:linkname Unquote github.com/bytedance/sonic/internal/native/neon.__unquote
func Unquote(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) int

//go:nosplit
//go:noescape
//go:linkname Value github.com/bytedance/sonic/internal/native/neon.__value
func Value(s unsafe.Pointer, n int, p int, v *types.JsonState, flags uint64) int

//go:nosplit
//go:noescape
//go:linkname ValidateOne github.com/bytedance/sonic/internal/native/neon.__validate_one
func ValidateOne(s *string, p *int, m *types.StateMachine, flags uint64) int

//go:nosplit
//go:noescape
//go:linkname ValidateUTF8 github.com/bytedance/sonic/internal/native/neon.__validate_utf8
func ValidateUTF8(s *string, p *int, m *types.StateMachine) (ret int)

//go:nosplit
//go:noescape
//go:linkname ValidateUTF8Fast github.com/bytedance/sonic/internal/native/neon.__validate_utf8_fast
func ValidateUTF8Fast(s *string) (ret int)

func useNeon() {
	S_f64toa = neon.S_f64toa
	S_f32toa = neon.S_f32toa
	S_i64toa = neon.S_i64toa
	S_u64toa = neon.S_u64toa
	S_lspace = neon.S_lspace
	S_quote = neon.S_quote
	S_unquote = neon.S_unquote
	S_value = neon.S_value
	S_vstring = neon.S_vstring
	S_vnumber = neon.S_vnumber
	S_vsigned = neon.S_vsigned
	S_vunsigned = neon.S_vunsigned
	S_skip_one = neon.S_skip_one
	S_skip_one_fast = neon.S_skip_one_fast
	S_skip_array = neon.S_skip_array
	S_skip_object = neon.S_skip_object
	S_skip_number = neon.S_skip_number
	S_get_by_path = neon.S_get_by_path
	S_parse_with_padding = neon.S_parse_with_padding
	S_lookup_small_key = neon.S_lookup_small_key
}

func useSveLinkname() {
	S_f64toa = neon.S_f64toa
	S_f32toa = neon.S_f32toa
	S_i64toa = neon.S_i64toa
	S_u64toa = neon.S_u64toa
	S_lspace = neon.S_lspace
	S_quote = neon.S_quote
	S_unquote = neon.S_unquote
	S_value = neon.S_value
	S_vstring = neon.S_vstring
	S_vnumber = neon.S_vnumber
	S_vsigned = neon.S_vsigned
	S_vunsigned = neon.S_vunsigned
	S_skip_one = sve_linkname.S_skip_one
	S_skip_one_fast = sve_linkname.S_skip_one_fast
	S_skip_array = neon.S_skip_array
	S_skip_object = neon.S_skip_object
	S_skip_number = neon.S_skip_number
	S_get_by_path = sve_linkname.S_get_by_path
	S_parse_with_padding = sve_linkname.S_parse_with_padding
	S_lookup_small_key = sve_linkname.S_lookup_small_key
}

func useSveWrapgoc() {
	sve_wrapgoc.Use()
	__GetByPath = sve_wrapgoc.F_get_by_path
	__ParseWithPadding = sve_wrapgoc.F_parse_with_padding
	__F32toa = sve_wrapgoc.F_f32toa
	__F64toa = sve_wrapgoc.F_f64toa
	__I64toa = sve_wrapgoc.F_i64toa
	__U64toa = sve_wrapgoc.F_u64toa
	__Quote = sve_wrapgoc.F_quote
	__SkipOne = sve_wrapgoc.F_skip_one
	__SkipOneFast = sve_wrapgoc.F_skip_one_fast
	__LookupSmallKey = sve_wrapgoc.F_lookup_small_key

	S_f64toa = sve_wrapgoc.S_f64toa
	S_f32toa = sve_wrapgoc.S_f32toa
	S_i64toa = sve_wrapgoc.S_i64toa
	S_u64toa = sve_wrapgoc.S_u64toa
	S_lspace = neon.S_lspace
	S_quote  = sve_wrapgoc.S_quote
	S_unquote = neon.S_unquote
	S_value = neon.S_value
	S_vstring = neon.S_vstring
	S_vnumber = neon.S_vnumber
	S_vsigned = neon.S_vsigned
	S_vunsigned = neon.S_vunsigned
	S_skip_one = sve_wrapgoc.S_skip_one
	S_skip_one_fast = sve_wrapgoc.S_skip_one_fast
	S_skip_array = neon.S_skip_array
	S_skip_object = neon.S_skip_object
	S_skip_number = neon.S_skip_number
	S_get_by_path = sve_wrapgoc.S_get_by_path
	S_parse_with_padding = sve_wrapgoc.S_parse_with_padding
	S_lookup_small_key = sve_wrapgoc.S_lookup_small_key
}

func CpuDetect() bool {
	cpuinfo, err := cpu.Info()
	if err != nil {
		return false
	}

	if cpuinfo[0].Model == "0xd02" || cpuinfo[0].Model == "0xd06" {
		return true
	}

    file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return false
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	line := scanner.Text()
    	if strings.HasPrefix(line, "CPU part") {
    		parts := strings.SplitN(line, ":", 2)
			model := strings.TrimSpace(parts[1])
		if model == "0xd02" || model == "0xd06" {
				return true
			}
    	}
    }
	return false
}

func init() {
	if CpuDetect() {
		UseSveWrapgoc = os.Getenv("SONIC_USE_SVE_WRAPGOC") == "1"
		UseSveLinkname = os.Getenv("SONIC_USE_SVE_LINKNAME") == "1"	
	}
	if UseSveWrapgoc {
		useSveWrapgoc()
	} else if UseSveLinkname {
		useSveLinkname()
	} else {
		useNeon()
	}
}
