// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue_test

import (
	"testing"

	"github.com/bytedance/sonic"
)

func genSlice() interface{} {
	a := []string{}
	return &[]interface{}{&a}
}

func TestSlicePointer_Issue750(t *testing.T) {
	assertUnmarshal(t, sonic.ConfigStd, unmTestCase{
		name: "non-empty eface slice",
		newfn: genSlice,
		data: []byte(`["one","2"]`),
	})
}