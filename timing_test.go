// Copyright 2023 CloudWeGo Authors
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

package sonic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestWork(t *testing.T) {
	N, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		return
	}
	p, err := strconv.Atoi(os.Args[len(os.Args)-2])
	if err != nil {
		return
	}
	j := os.Args[len(os.Args)-3]
	var obj = map[string]interface{}{}
	for i:=0; i<100; i++ {
		obj[strconv.Itoa(i)] = i
	}
	var work func()
	if j == "sonic" {
		work = func() {
			for i:=0; i<N; i++ {
				out, _ := ConfigStd.Marshal(obj)
				var v interface{}
				_ = ConfigStd.Unmarshal(out, &v)
			}
		}
	} else {
		work = func() {
			for i:=0; i<N; i++ {
				out, _ := json.Marshal(obj)
				var v interface{}
				_ = json.Unmarshal(out, &v)
			}
		}
	}
	s := time.Now()
	work()
	fmt.Printf("process %d cost: %vus\n", p, time.Since(s).Microseconds())
}

//go:generate go test -c -o ptest.bin .
func TestScalability_Process(t *testing.T) {
	var T = runtime.NumCPU()
	var N = 300
	var tests = func(t* testing.T, name string) {
		var test = func(t *testing.T, p, n int) {
			var ch = make(chan error, T)
			s := time.Now()
			for i:=0; i<p; i++ {
				var buf = bytes.NewBuffer(nil)
				cmd := exec.Command("./ptest.bin", "-test.run", "TestWork", name, strconv.Itoa(i), strconv.Itoa(n))
				cmd.Stdout = buf
				// cmd.Stderr = buf
				if err := cmd.Start(); err != nil {
					t.Fatal(err)
				}
				go func(i int) {
					e := cmd.Wait()
					// println(buf.String())
					ch <- e
				}(i)
			}
			for i:=0; i<p; i++ {
				e := <- ch
				if e != nil {
					t.Fatal(e)
				}
			}
			fmt.Printf("%d processes total cost: %dus\n", p, time.Since(s).Microseconds())
		}
	
		println("single-process")
		test(t, 1, T*N)
		println("1/4-processes")
		test(t, T/4, N*4)
		println("half-processes")
		test(t, T/2, N*2)
		println("3/4-processes")
		test(t, T/4*3, N/3*4)
		println("full-processes")
		test(t, T, N)
		println("full-processes-1")
		test(t, T-1, N+N/(T-1))
	}
	
	t.Run("sonic", func(t *testing.T) {
		tests(t, "sonic")
	})
	t.Run("std", func(t *testing.T) {
		tests(t, "std")
	})
}

func TestScalability_Thread(t *testing.T) {
	const N = 300
	var T = runtime.NumCPU()
	println("num cpu:", T)

	var obj = map[string]interface{}{}
	for i:=0; i<100; i++ {
		obj[strconv.Itoa(i)] = i
	}
	out, err := ConfigStd.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	var x interface{}
	err = ConfigStd.Unmarshal(out, &x)
	if err != nil {
		t.Fatal(err)
	}

	var tests = func(t *testing.T, work func()) {
		var test = func(t *testing.T, p int, n int) {
			wg := sync.WaitGroup{}
			s := time.Now()
			for i:=0; i<p; i++ {
				wg.Add(1)
				go func ()  {
					for j:=0; j<n; j++ {
						work()
					}
					wg.Done()
				}()
			}
			wg.Wait()
			fmt.Printf("%d threads total cost: %vus\n", p, time.Since(s).Microseconds())
		}
		println("single-thread")
		test(t, 1, T*N)
		println("1/4-threads")
		test(t, T/4, N*4)
		println("half-threads")
		test(t, T/2, N*2)
		println("3/4-threads")
		test(t, T/4*3, N/3*4)
		println("full-threads")
		test(t, T, N)
		println("full-threads-1")
		test(t, T-1, N+N/(T-1))
	}

	t.Run("sonic", func(t *testing.T) {
		work := func() {
			out, _ := ConfigStd.Marshal(obj)
			var x interface{}
			_ = ConfigStd.Unmarshal(out, &x)
		}
		tests(t, work)
	})

	t.Run("std", func(t *testing.T) {
		work := func() {
			out, _ := json.Marshal(obj)
			var x interface{}
			_ = json.Unmarshal(out, &x)
		}
		tests(t, work)
	})
}


