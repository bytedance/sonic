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

package issue_test

import (
	`encoding/json`
	`fmt`
	`sync`
	`testing`

	`github.com/bytedance/sonic`
)

func TestGcWriteBarrier(t *testing.T) {
	// debug.SetGCPercent(-1)
	size := 10_0000
	data := make([]int, size)
	date, _ := json.Marshal(data)	

	// debug.SetGCPercent(-1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		cd := make([]byte, len(date))
		copy(cd, date)
		go func() {
			var w []int
			defer wg.Done()
			// Decode
			if err := sonic.Unmarshal(cd, &w); err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkGcGuard(b *testing.B) {
	size := 10_0000
	data := make([]int, size)
	date, _ := json.Marshal(data)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var w []int
			_ = sonic.Unmarshal(date, &w)
		}
	})
}