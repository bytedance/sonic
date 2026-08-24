//go:build arm64
// +build arm64

/*
 * Copyright 2026 The sonic Authors.
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

package neon

import (
	"io"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// TestProfileAndGCStress is a regression check for
// https://github.com/bytedance/sonic/issues/970 /
// https://github.com/golang/go/issues/80771: before this package moved to
// asm2arm_tool's JIT mode (real Pcsp via internal/loader, matching how
// internal/native/avx2 already works on amd64), a CPU-profiling signal
// landing inside one of these native calls could hit the Go runtime's
// "traceback stuck" or "unknown caller pc" fatal errors, because the
// static-link .s files this package used to compile through carried no
// real PC->SP unwind data for their C-compiler-generated bodies.
//
// This does not deterministically reproduce that crash -- the underlying
// race requires a profiling sample to land within a few-instruction window
// of a function's prologue, which in production surfaced only "roughly
// monthly across a production arm64 fleet" per the original report. What
// this test verifies is that heavy concurrent CPU profiling and GC (both
// of which unwind through these native frames -- sigprof with
// unwindSilentErrors, GC stack scans with the strict flags that require
// real unwind info to succeed) run cleanly through every native in this
// package without the process going down. Combined with the
// TestRecover_* tests below (which force a real hardware fault, and
// therefore a real signal-driven unwind, inside each native call and
// require a clean recoverable panic), this exercises the same unwind
// machinery the original bug broke.
func TestProfileAndGCStress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	if err := pprof.StartCPUProfile(io.Discard); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	var workers sync.WaitGroup
	stop := make(chan struct{})

	call := func() {
		var v types.JsonState
		s := `   -12345`
		sp := (*rt.GoString)(unsafe.Pointer(&s))
		_ = value(sp.Ptr, sp.Len, 0, &v, 0)

		vs := "asdf"
		vp := 0
		_ = validate_one(&vs, &vp, &types.StateMachine{}, 0)

		s2 := `{"asdf": [null, true, false, 1, 2.0, -3]}, 1234.5`
		p2 := 0
		path := []interface{}{"asdf", 4}
		_ = get_by_path(&s2, &p2, &path, types.NewStateMachine())
	}

	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for {
				select {
				case <-stop:
					return
				default:
					call()
				}
			}
		}()
	}

	// Drive GC synchronously on the main goroutine so the test's duration
	// is bounded by how long these GC cycles actually take, not a fixed
	// sleep -- the workers above keep hammering native calls the whole
	// time, maximizing the odds a profiling signal lands mid-call.
	for i := 0; i < 50; i++ {
		runtime.GC()
	}

	close(stop)
	workers.Wait()
}
