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

package decoder

import (
    `bytes`
    `io`
    `sync`

    `github.com/bytedance/sonic/internal/native/types`
)

var (
    defaultBufferSize    uint = 4096
    growSliceFactorShift uint = 1
    minLeftBufferShift   uint = 2
)

type StreamDecoder struct {
    r       io.Reader
    buf     []byte
    scanp   int
    scanned int64
    Decoder
    err     error
}

var bufPool = sync.Pool{
    New: func () interface{} {
        return make([]byte, 0, defaultBufferSize)
    },
}

func NewStreamDecoder(r io.Reader) *StreamDecoder {
    return &StreamDecoder{r : r}
}

func (self *StreamDecoder) Decode(val interface{}) (error) {
	if self.err != nil {
		return self.err
	}

    var buf = self.buf
    var p = 0
    var recycle bool
    if cap(buf) == 0 {
        buf = bufPool.Get().([]byte)
        recycle = true
    }
	
	var first = true
read_more:
    for {
        l := len(buf)
        realloc(&buf)
        n, err := self.r.Read(buf[l:cap(buf)])
        buf = buf[:l+n]
        if err != nil {
            if err == io.EOF {
                break
            }
            self.err = err
            return err
        }
        if n > 0 || first {
            break
        }
	}
	first = false

    self.Decoder.Reset(string(buf))
    err := self.Decoder.Decode(val)
    if err != nil {
        if ee, ok := err.(SyntaxError); ok && ee.Code == types.ERR_EOF {
            goto read_more
        }
        self.err = err
    }

    // remain undecoded bytes, so copy them into self.buf
    p = self.Decoder.Pos()
    self.scanned += int64(p)
    self.scanp = 0
    if p < len(buf) {
        self.buf = append(self.buf[:0], buf[p:]...)
    }

    if recycle {
        // recycle buffer
        buf = buf[:0]
        bufPool.Put(buf)
    }
    return err
}

func (self *StreamDecoder) InputOffset() int64 {
    return self.scanned + int64(self.scanp)
}

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
    return err == nil && self.err == nil && c != ']' && c != '}'
}

func (self *StreamDecoder) peek() (byte, error) {
    var err error
    for {
        for i := self.scanp; i < len(self.buf); i++ {
            c := self.buf[i]
            if isSpace(c) {
                continue
            }
            self.scanp = i
            return c, nil
        }
        // buffer has been scanned, now report any error
        if err != nil {
            if err != io.EOF {
                self.err = err
            }
            return 0, err
        }
        err = self.refill()
    }
}

func isSpace(c byte) bool {
    return types.SPACE_MASK & (1 << c) != 0
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

    // Grow buffer if not large enough.
    realloc(&self.buf)

    // Read. Delay error for next iteration (after scan).
    n, err := self.r.Read(self.buf[len(self.buf):cap(self.buf)])
    self.buf = self.buf[0 : len(self.buf)+n]

    return err
}

func realloc(buf *[]byte) {
	l := uint(len(*buf))
	c := uint(cap(*buf))
    if c - l < c >> minLeftBufferShift {
        tmp := make([]byte, l, l+(l>>minLeftBufferShift))
        copy(tmp, *buf)
        *buf = tmp
    }
}

