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

package jit

import (
    `testing`

    `github.com/davecgh/go-spew/spew`
    `github.com/twitchyliquid64/golang-asm/obj`
    `github.com/twitchyliquid64/golang-asm/obj/x86`
)

func TestBackend(t *testing.T) {
    e := newBackend("amd64")
    p := e.New()
    p.As = x86.AVPTEST
    (*BaseAssembler)(nil).assignOperands(p, []obj.Addr{Reg("Y2"), Reg("Y1")})
    e.Append(p)
    spew.Dump(e.Assemble())
}
