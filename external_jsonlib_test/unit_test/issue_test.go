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

package unit_test

import (
	"github.com/gin-gonic/gin"
	"github.com/bytedance/sonic"
    "context"
	"testing"
)

func TestContext(t *testing.T) {
	var obj = new(Context)
	obj.Context = context.Background()
	obj.Context = context.WithValue(obj.Context, "a", "b")
	obj.GinCtx = new(gin.Context)
	obj.GinCtx.Accepted = []string{"1"}
	out, err := sonic.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	println(string(out))
	err = sonic.Unmarshal(out, obj)
	if err != nil {
		t.Fatal(err)
	}
}


type Context struct {
	context.Context
	GinCtx *gin.Context
}

