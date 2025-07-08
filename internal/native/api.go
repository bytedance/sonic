/**
 * Copyright 2025 ByteDance Inc.
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 *     https://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package native

import "github.com/bytedance/sonic/internal/utils"

func SkipOneFast(json *string, pos *int) int {
	start := SkipOneFastTrailing(json, pos)
	if start < 0 {
		return start
	}

	// because the `SkipOneFast` ignored the trailing white spaces, should trim them
	for utils.IsSpace((*json)[*pos - 1]) && *pos > start {
		*pos -= 1;
	}
	return start
}
