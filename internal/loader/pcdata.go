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

package loader

import (
	// "encoding"
)

const (
	PCDATA_UnsafePoint   = 0
	PCDATA_StackMapIndex = 1
	PCDATA_InlTreeIndex  = 2
	PCDATA_ArgLiveIndex  = 3
)

const (
	// PCDATA_UnsafePoint values.
	_PCDATA_UnsafePointSafe   = -1 // Safe for async preemption
	_PCDATA_UnsafePointUnsafe = -2 // Unsafe for async preemption

	// _PCDATA_Restart1(2) apply on a sequence of instructions, within
	// which if an async preemption happens, we should back off the PC
	// to the start of the sequence when resume.
	// We need two so we can distinguish the start/end of the sequence
	// in case that two sequences are next to each other.
	_PCDATA_Restart1 = -3
	_PCDATA_Restart2 = -4

	// Like _PCDATA_RestartAtEntry, but back to function entry if async
	// preempted.
	_PCDATA_RestartAtEntry = -5

	_PCDATA_START_VAL = -1
)

type Pcvalue struct {
	PC  uint32 // PC offset from text section
	Val int32
}

type Pcdata []Pcvalue

func (self Pcdata) Marshal(startPC uintptr) (data []byte, err error) {
	// delta value always starts from -1
	// see https://docs.google.com/document/d/1lyPIbmsYbXnpNj57a261hgOYVpNRcgydurVQIyZOz_o/pub
	sv := int32(_PCDATA_START_VAL)
	sp := uint32(startPC)
	for _, v := range self {
		data = append(data, encodeVariant(toZigzag(int(v.Val - sv)))...)
		data = append(data, encodeVariant(toZigzag(int(v.PC - sp)))...)
		sp = v.PC
		sv = v.Val
	}
	return
}