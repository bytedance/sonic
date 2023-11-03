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
	"bytes"
	"io"
	"sync"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/option"
)

var (
    minLeftBufferShift   uint = 1
)

// StreamDecoder is the decoder context object for streaming input.
type StreamDecoder struct {
    r       io.Reader
    buf     []byte
    scanp   int
    scanned int64
    err     error
    Decoder
}

var bufPool = sync.Pool{
    New: func () interface{} {
        return make([]byte, 0, option.DefaultDecoderBufferSize)
    },
}

// NewStreamDecoder adapts to encoding/json.NewDecoder API.
//
// NewStreamDecoder returns a new decoder that reads from r.
func NewStreamDecoder(r io.Reader) *StreamDecoder {
    return &StreamDecoder{r : r}
}

// Decode decodes input stream into val with corresponding data. 
// Redundantly bytes may be read and left in its buffer, and can be used at next call.
// Either io error from underlying io.Reader (except io.EOF) 
// or syntax error from data will be recorded and stop subsequently decoding.
func (self *StreamDecoder) Decode(val interface{}) (err error) {
    // read more data into buf
    for self.More() {
        var src = rt.Mem2Str(self.buf[self.scanp:])
        var x int
        var s int
        if s = native.SkipOneFast(&src, &x); s < 0 {
            // mandatorily set buf full, trigger refill 
            self.scanp = len(self.buf)
            if self.More()  {
                continue
            } else {
                err = SyntaxError{x, self.s, types.ParsingError(-s), ""}
                self.err = err
                return
            }
        }
        s += self.scanp
        x += self.scanp

        // must copy string here for safety
        self.Decoder.Reset(string(self.buf[s:x]))
        err = self.Decoder.Decode(val)
        if err != nil {
            self.err = err
            return 
        }

        self.scanned += int64(x)
        self.scanp = 0

        if len(self.buf) == x  {
            // fully scan, thus we just recycle buffer
            mem := self.buf
            self.buf = nil
            bufPool.Put(mem[:0])
        } else {
            n := copy(self.buf, self.buf[x:])
            self.buf = self.buf[:n]
        }   

        break
    }    

    return self.err
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

func realloc(buf *[]byte) bool {
    l := uint(len(*buf))
    c := uint(cap(*buf))
    if c == 0 {
        println("use pool!")
       *buf = bufPool.Get().([]byte)
       return true
    }
    if c - l <= c >> minLeftBufferShift {
        println("realloc!")
        e := l+(l>>minLeftBufferShift)
        if e < option.DefaultDecoderBufferSize {
            e = option.DefaultDecoderBufferSize
        }
        tmp := make([]byte, l, e)
        copy(tmp, *buf)
        *buf = tmp
        return true
    }
    return false
}

