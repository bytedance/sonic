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

package sse

import (
	"runtime"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
	"github.com/stretchr/testify/assert"
)

func TestA_Native_Value(t *testing.T) {
    runtime.GC()
    var v types.JsonState
    s := `   -12345`
    p := (*rt.GoString)(unsafe.Pointer(&s))
    x := value(p.Ptr, p.Len, 0, &v, 0)
    assert.Equal(t, 9, x)
    assert.Equal(t, types.V_INTEGER, v.Vt)
    assert.Equal(t, int64(-12345), v.Iv)
    assert.Equal(t, 3, v.Ep)
}

