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
	"bytes"
	"io"
	"sync"

	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/internal/utils"
	"github.com/bytedance/sonic/option"
)

var (
	minLeftBufferShift uint = 1
)

// StreamDecoder is the decoder context object for streaming input.
type StreamDecoder struct {
	r       io.Reader
	buf     []byte
	scanp   int
	scanned int64
	err     error
	readErr error // pending Reader error; buffered data is scanned first
	Decoder
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, option.DefaultDecoderBufferSize)
	},
}

func freeBytes(buf []byte) {
	if rt.CanSizeResue(cap(buf)) {
		bufPool.Put(buf[:0])
	}
}

// NewStreamDecoder adapts to encoding/json.NewDecoder API.
//
// NewStreamDecoder returns a new decoder that reads from r.
func NewStreamDecoder(r io.Reader) *StreamDecoder {
	return &StreamDecoder{r: r}
}

// Decode decodes input stream into val with corresponding data.
// Redundantly bytes may be read and left in its buffer, and can be used at next call.
// Either io error from underlying io.Reader (except io.EOF)
// or syntax error from data will be recorded and stop subsequently decoding.
func (self *StreamDecoder) Decode(val interface{}) (err error) {
	if self.err != nil {
		return self.err
	}
	// Decode must inspect closing delimiters as values so that a stray top-level
	// ']' or '}' becomes a syntax error. More intentionally reports false for
	// those bytes and is only suitable for array/object iteration.
	if _, err = self.peek(); err != nil {
		return err
	}

	s := self.scanp
	e, scanErr := self.readValue()
	if scanErr != nil {
		return scanErr
	}
	// must copy string here for safety
	self.Decoder.Reset(string(self.buf[s:e]))
	err = self.Decoder.Decode(val)
	if err != nil {
		self.setErr(err)
		return
	}

	self.scanp = e
	_, empty := self.scan()
	if empty {
		// no remain valid bytes, thus we just recycle buffer
		mem := self.buf
		self.buf = nil
		freeBytes(mem)
		self.err = self.readErr
	} else {
		// remain undecoded bytes, move them onto head
		n := copy(self.buf, self.buf[self.scanp:])
		self.buf = self.buf[:n]
	}

	self.scanned += int64(self.scanp)
	self.scanp = 0
	return nil
}

// InputOffset returns the input stream byte offset of the current decoder position.
// The offset gives the location of the end of the most recently returned token and the beginning of the next token.
func (self *StreamDecoder) InputOffset() int64 {
	return self.scanned + int64(self.scanp)
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (self *StreamDecoder) Buffered() io.Reader {
	return bytes.NewReader(self.buf[self.scanp:])
}

// More reports whether there is another element in the
// current array or object being parsed.
func (self *StreamDecoder) More() bool {
	if self.err != nil {
		return false
	}
	c, err := self.peek()
	return err == nil && c != ']' && c != '}'
}

// readValue owns the scan/read loop without consuming or releasing the token.
func (self *StreamDecoder) readValue() (int, error) {
	s := self.scanp
	var scanner streamValueScanner
	for {
		src := rt.Mem2Str(self.buf[s:])
		end, status, err := scanner.scan(src, self.readErr == io.EOF)
		switch status {
		case streamDone:
			return s + end, nil
		case streamInvalid:
			self.setErr(err)
			return 0, err
		case streamNeedMore:
			if self.readErr != nil {
				err = self.readErr
				self.setErr(err)
				return 0, err
			}
			self.readMore()
		}
	}
}

// readMore appends one read. Cursor movement and error handling belong to callers.
func (self *StreamDecoder) readMore() error {
	if self.readErr != nil {
		return self.readErr
	}
	l := len(self.buf)
	realloc(&self.buf)
	n, err := self.r.Read(self.buf[l:cap(self.buf)])
	self.buf = self.buf[:l+n]
	self.readErr = err
	return err
}

func (self *StreamDecoder) setErr(err error) {
	self.err = err
	self.scanned += int64(self.scanp)
	self.scanp = 0
	if self.buf != nil {
		mem := self.buf[:0]
		self.buf = nil
		freeBytes(mem)
	}
}

func (self *StreamDecoder) peek() (byte, error) {
	var err error
	for {
		c, empty := self.scan()
		if !empty {
			return byte(c), nil
		}
		// buffer has been scanned, now report any error
		if err != nil {
			self.setErr(err)
			return 0, err
		}
		err = self.refill()
	}
}

func (self *StreamDecoder) scan() (byte, bool) {
	for i := self.scanp; i < len(self.buf); i++ {
		c := self.buf[i]
		if utils.IsSpace(c) {
			continue
		}
		self.scanp = i
		return c, false
	}
	return 0, true
}

func (self *StreamDecoder) refill() error {
	// Make room to read more into the buffer.
	// First slide down data already consumed.
	if self.scanp > 0 {
		self.scanned += int64(self.scanp)
		n := copy(self.buf, self.buf[self.scanp:])
		self.buf = self.buf[:n]
		self.scanp = 0
	}

	return self.readMore()
}

func realloc(buf *[]byte) bool {
	l := uint(len(*buf))
	c := uint(cap(*buf))
	if c == 0 {
		*buf = bufPool.Get().([]byte)
		return true
	}
	if c-l <= c>>minLeftBufferShift {
		e := l + (l >> minLeftBufferShift)
		if e <= c {
			e = c * 2
		}
		tmp := make([]byte, l, e)
		copy(tmp, *buf)
		*buf = tmp
		return true
	}
	return false
}
