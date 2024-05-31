// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vm_test

import (
	"encoding/json"
	"runtime"
	"runtime/debug"
	"testing"
	"time"

	"github.com/bytedance/sonic/internal/encoder"
	"github.com/bytedance/sonic/internal/encoder/vars"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
    go func ()  {
        if !vars.DebugAsyncGC {
            return
        }
        println("Begin GC looping...")
        for {
            runtime.GC()
            debug.FreeOSMemory()
        }
        println("stop GC looping!")
    }()
    time.Sleep(time.Millisecond)
    encoder.ForceUseVM()
    m.Run()
}

type StringStruct struct {
    X *int        `json:"x,string,omitempty"`
    Y []int       `json:"y"`
    Z json.Number `json:"z,string"`
    W string      `json:"w,string"`
}

func TestEncoder_FieldStringize(t *testing.T) {
    x := 12345
    v := StringStruct{X: &x, Y: []int{1, 2, 3}, Z: "4567456", W: "asdf"}
    r, e := encoder.Encode(v, 0)
    require.NoError(t, e)
    println(string(r))
}

func TestCorpusMarshal(t *testing.T) {
    var v = struct { F0 int "json:\"C,omitempty\""; F1 **int "json:\"a,\""; p2 **int }{}
    sout, serr := encoder.Encode(v, 0)
    jout, jerr := json.Marshal(v)
    require.Equal(t, jerr == nil, serr == nil)
    require.Equal(t, string(jout), string(sout))
}
