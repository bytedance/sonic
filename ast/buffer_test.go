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
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
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
		r[i] = NewPair(strconv.Itoa(i), NewBool(true))
	}
	return r
}

func Test_linkedPairs_Push(t *testing.T) {
	type args struct {
		in  []Pair
		v   Pair
		exp []Pair
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "add empty",
			args: args{
				in:  []Pair{},
				v:   NewPair("a", NewBool(true)),
				exp: []Pair{NewPair("a", NewBool(true))},
			},
		},
		{
			name: "add one",
			args: args{
				in:  []Pair{NewPair("a", NewBool(false))},
				v:   NewPair("b", NewBool(true)),
				exp: []Pair{NewPair("a", NewBool(false)), NewPair("b", NewBool(true))},
			},
		},
		{
			name: "add _DEFAULT_NODE_CAP",
			args: args{
				in:  makePairs(_DEFAULT_NODE_CAP),
				v:   NewPair(strconv.Itoa(_DEFAULT_NODE_CAP), NewBool(true)),
				exp: makePairs(_DEFAULT_NODE_CAP + 1),
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
		v   Node
		exp []Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "add empty",
			args: args{
				in:  []Node{},
				v:   NewBool(true),
				exp: []Node{NewBool(true)},
			},
		},
		{
			name: "add one",
			args: args{
				in:  []Node{NewBool(false)},
				v:   NewBool(true),
				exp: []Node{NewBool(false), NewBool(true)},
			},
		},
		{
			name: "add _DEFAULT_NODE_CAP",
			args: args{
				in:  makeNodes(_DEFAULT_NODE_CAP),
				v:   NewBool(true),
				exp: makeNodes(_DEFAULT_NODE_CAP + 1),
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
		name string
		args args
	}{
		{
			name: "remove empty",
			args: args{
				in:  []Node{},
				exp: []Node{},
			},
		},
		{
			name: "remove one",
			args: args{
				in:  []Node{NewBool(false)},
				exp: []Node{},
			},
		},
		{
			name: "add _DEFAULT_NODE_CAP",
			args: args{
				in:  makeNodes(_DEFAULT_NODE_CAP),
				exp: makeNodes(_DEFAULT_NODE_CAP - 1),
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
		in     []Node
		source int
		target int
		exp    []Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "over index",
			args: args{
				in:     []Node{NewBool(true)},
				source: 1,
				target: 0,
				exp:    []Node{NewBool(true)},
			},
		},
		{
			name: "equal index",
			args: args{
				in:     []Node{NewBool(true)},
				source: 0,
				target: 0,
				exp:    []Node{NewBool(true)},
			},
		},
		{
			name: "forward index",
			args: args{
				in:     []Node{NewString("a"), NewString("b"), NewString("c")},
				source: 0,
				target: 2,
				exp:    []Node{NewString("b"), NewString("c"), NewString("a")},
			},
		},
		{
			name: "backward index",
			args: args{
				in:     []Node{NewString("a"), NewString("b"), NewString("c")},
				source: 2,
				target: 1,
				exp:    []Node{NewString("a"), NewString("c"), NewString("b")},
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

func TestLinkedPairsEmptyKeySlots(t *testing.T) {
	for _, indexed := range []bool{false, true} {
		t.Run(strconv.FormatBool(indexed), func(t *testing.T) {
			var pairs linkedPairs
			pairs.Push(NewPair("removed", NewNumber("1")))
			pairs.Push(NewPair("", NewNumber("2")))
			if indexed {
				pairs.BuildIndex()
			}
			pairs.Unset(0)
			p, i := pairs.Get("")
			require.Same(t, pairs.At(1), p)
			require.Equal(t, 1, i)
			pairs.Unset(1)
			p, i = pairs.Get("")
			require.Nil(t, p)
			require.Equal(t, -1, i)
		})
	}
}

func TestEmptyKeyAfterUnsetByIndex(t *testing.T) {
	for _, size := range []int{2, _Threshold_Index + 1} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			pairs := makePairs(size)
			pairs[size-1] = NewPair("", NewNumber("2"))
			root := NewObject(pairs)
			existed, err := root.UnsetByIndex(0)
			require.NoError(t, err)
			require.True(t, existed)
			value, err := root.Get("").Int64()
			require.NoError(t, err)
			require.Equal(t, int64(2), value)
			existed, err = root.Unset("")
			require.NoError(t, err)
			require.True(t, existed)
			require.Nil(t, root.Get(""))
		})
	}
}

func TestLinkedPairsEmptyKeyHashConflict(t *testing.T) {
	var pairs linkedPairs
	pairs.Push(NewPair("removed", NewNull()))
	pairs.Push(NewPair("", NewNumber("2")))
	collision := NewPair("collision", NewNull())
	collision.hash = pairs.At(1).hash
	pairs.Push(collision)
	pairs.Unset(0)
	pairs.BuildIndex()
	// The index points to the colliding key; fallback must skip the deleted slot.
	p, i := pairs.Get("")
	require.Same(t, pairs.At(1), p)
	require.Equal(t, 1, i)
	// A literal replacement must also remain visible through the fallback.
	*pairs.At(1) = Pair{Key: "", Value: NewNumber("3")}
	p, i = pairs.Get("")
	require.Same(t, pairs.At(1), p)
	require.Equal(t, 1, i)
}

func TestNewObjectPairLiterals(t *testing.T) {
	for _, size := range []int{2, _DEFAULT_NODE_CAP + 1, _Threshold_Index + 1, 3*_DEFAULT_NODE_CAP + 1} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			input := make([]Pair, size)
			for i := range input {
				input[i] = Pair{Key: strconv.Itoa(i), Value: NewNumber(strconv.Itoa(i))}
			}
			// Even a zero Pair literal is a writable, empty-key placeholder.
			input[size-1] = Pair{}
			root := NewObject(input)
			for i := 0; i < size-1; i++ {
				value, err := root.Get(strconv.Itoa(i)).Int64()
				require.NoError(t, err)
				require.Equal(t, int64(i), value)
			}
			require.NotNil(t, root.Get(""))
			require.NoError(t, root.Get("").Add(NewNull()))
			_, err := root.Unset("0")
			require.NoError(t, err)
			require.NotNil(t, root.Get(""))
			require.Zero(t, input[size-1].hash, "NewObject must not modify its input")
		})
	}
}
