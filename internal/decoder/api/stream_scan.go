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

package api

import (
	"io"
	"strings"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
)

type streamScanStatus uint8

const (
	streamDone streamScanStatus = iota
	streamNeedMore
	streamInvalid
)

type streamNumberState uint8

const (
	streamNumStart streamNumberState = iota
	streamNumSign
	streamNumZero
	streamNumInt
	streamNumDot
	streamNumFrac
	streamNumExp
	streamNumExpSign
	streamNumExpDigits
)

// streamValueScanner frames one value. src starts at a non-space byte and may
// only be extended between calls. Numeric state is retained across reads;
// other values use the existing native scanner. Only a real EOF sets final.
type streamValueScanner struct {
	pos    int
	number streamNumberState
}

func (s *streamValueScanner) scan(src string, final bool) (int, streamScanStatus, error) {
	if len(src) != 0 && src[0] != '-' && (src[0] < '0' || src[0] > '9') {
		end := 0
		start := native.SkipOneFast(&src, &end)
		if start >= 0 {
			return end, streamDone, nil
		}
		code := types.ParsingError(-start)
		if code != types.ERR_EOF {
			return end, streamInvalid, SyntaxError{end, strings.Clone(src), code, ""}
		}
		if !final {
			return end, streamNeedMore, nil
		}
		return end, streamInvalid, io.ErrUnexpectedEOF
	}

	for s.pos < len(src) {
		c := src[s.pos]
		digit := c >= '0' && c <= '9'
		switch s.number {
		case streamNumStart, streamNumSign:
			switch {
			case c == '-' && s.number == streamNumStart:
				s.number = streamNumSign
			case c == '0':
				s.number = streamNumZero
			case c >= '1' && c <= '9':
				s.number = streamNumInt
			default:
				return s.invalidNumber(src)
			}
		case streamNumZero, streamNumInt, streamNumFrac, streamNumExpDigits:
			switch {
			case digit && s.number != streamNumZero:
			case c == '.' && (s.number == streamNumZero || s.number == streamNumInt):
				s.number = streamNumDot
			case (c == 'e' || c == 'E') && s.number != streamNumExpDigits:
				s.number = streamNumExp
			default:
				return s.pos, streamDone, nil
			}
		case streamNumDot, streamNumExpSign:
			if !digit {
				return s.invalidNumber(src)
			}
			if s.number == streamNumDot {
				s.number = streamNumFrac
			} else {
				s.number = streamNumExpDigits
			}
		case streamNumExp:
			if c == '+' || c == '-' {
				s.number = streamNumExpSign
			} else if digit {
				s.number = streamNumExpDigits
			} else {
				return s.invalidNumber(src)
			}
		}
		s.pos++
	}
	if !final {
		return s.pos, streamNeedMore, nil
	}
	switch s.number {
	case streamNumZero, streamNumInt, streamNumFrac, streamNumExpDigits:
		return s.pos, streamDone, nil
	default:
		return s.pos, streamInvalid, io.ErrUnexpectedEOF
	}
}

func (s *streamValueScanner) invalidNumber(src string) (int, streamScanStatus, error) {
	return s.pos, streamInvalid, SyntaxError{s.pos, strings.Clone(src), types.ERR_INVALID_NUMBER_FMT, ""}
}
