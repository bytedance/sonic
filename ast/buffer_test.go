/**
 * Copyright 2023 ByteDance Inc.
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

package ast

import (
    `strconv`
    `testing`

    `github.com/stretchr/testify/require`
)

func makeNodes(l int) []Node {
    r := make([]Node, l)
    for i := 0; i < l; i++ {
        r[i] = NewBool(true)
    }
    return r
}

func makePairs(l int) []Pair {
    r := make([]Pair, l)
    for i := 0; i < l; i++ {
        r[i] = Pair{strconv.Itoa(i), NewBool(true)}
    }
    return r
}

func Test_linkedPairs_Push(t *testing.T) {
    type args struct {
        in  []Pair
        v Pair
        exp []Pair
    }
    tests := []struct {
        name   string
        args   args
    }{
        {
            name: "add empty",
            args: args{
                in: []Pair{},
                v: Pair{"a", NewBool(true)},
                exp: []Pair{Pair{"a", NewBool(true)}},
            },
        },
        {
            name: "add one",
            args: args{
                in: []Pair{{"a", NewBool(false)}},
                v: Pair{"b", NewBool(true)},
                exp: []Pair{{"a", NewBool(false)}, {"b", NewBool(true)}},
            },
        },
        {
            name: "add _DEFAULT_NODE_CAP",
            args: args{
                in: makePairs(_DEFAULT_NODE_CAP),
                v: Pair{strconv.Itoa(_DEFAULT_NODE_CAP), NewBool(true)},
                exp: makePairs(_DEFAULT_NODE_CAP+1),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            self := &linkedPairs{}
            self.FromSlice(tt.args.in)
            self.Push(tt.args.v)
            act := make([]Pair, self.Len())
            self.ToSlice(act)
            require.Equal(t, tt.args.exp, act)
        })
    }
}


func Test_linkedNodes_Push(t *testing.T) {
    type args struct {
        in  []Node
        v Node
        exp []Node
    }
    tests := []struct {
        name   string
        args   args
    }{
        {
            name: "add empty",
            args: args{
                in: []Node{},
                v: NewBool(true),
                exp: []Node{NewBool(true)},
            },
        },
        {
            name: "add one",
            args: args{
                in: []Node{NewBool(false)},
                v: NewBool(true),
                exp: []Node{NewBool(false), NewBool(true)},
            },
        },
        {
            name: "add _DEFAULT_NODE_CAP",
            args: args{
                in: makeNodes(_DEFAULT_NODE_CAP),
                v: NewBool(true),
                exp: makeNodes(_DEFAULT_NODE_CAP+1),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            self := &linkedNodes{}
            self.FromSlice(tt.args.in)
            self.Push(tt.args.v)
            act := make([]Node, self.Len())
            self.ToSlice(act)
            require.Equal(t, tt.args.exp, act)
        })
    }
}

func Test_linkedNodes_Pop(t *testing.T) {
    type args struct {
        in  []Node
        exp []Node
    }
    tests := []struct {
        name   string
        args   args
    }{
        {
            name: "remove empty",
            args: args{
                in: []Node{},
                exp: []Node{},
            },
        },
        {
            name: "remove one",
            args: args{
                in: []Node{NewBool(false)},
                exp: []Node{},
            },
        },
        {
            name: "add _DEFAULT_NODE_CAP",
            args: args{
                in: makeNodes(_DEFAULT_NODE_CAP),
                exp: makeNodes(_DEFAULT_NODE_CAP-1),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            self := &linkedNodes{}
            self.FromSlice(tt.args.in)
            self.Pop()
            act := make([]Node, self.Len())
            self.ToSlice(act)
            require.Equal(t, tt.args.exp, act)
        })
    }
}

func Test_linkedNodes_MoveOne(t *testing.T) {
    type args struct {
        in  []Node
        source int
        target int
        exp []Node
    }
    tests := []struct {
        name   string
        args   args
    }{
        {
            name: "over index",
            args: args{
                in: []Node{NewBool(true)},
                source: 1,
                target: 0,
                exp: []Node{NewBool(true)},
            },
        },
        {
            name: "equal index",
            args: args{
                in: []Node{NewBool(true)},
                source: 0,
                target: 0,
                exp: []Node{NewBool(true)},
            },
        },
        {
            name: "forward index",
            args: args{
                in: []Node{NewString("a"), NewString("b"), NewString("c")},
                source: 0,
                target: 2,
                exp: []Node{NewString("b"), NewString("c"), NewString("a")},
            },
        },
        {
            name: "backward index",
            args: args{
                in: []Node{NewString("a"), NewString("b"), NewString("c")},
                source: 2,
                target: 1,
                exp: []Node{NewString("a"), NewString("c"), NewString("b")},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            self := &linkedNodes{}
            self.FromSlice(tt.args.in)
            self.MoveOne(tt.args.source, tt.args.target)
            act := make([]Node, self.Len())
            self.ToSlice(act)
            require.Equal(t, tt.args.exp, act)
        })
    }
}


