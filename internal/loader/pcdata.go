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

const (
    _N_PCDATA   = 4

    _PCDATA_UnsafePoint   = 0
    _PCDATA_StackMapIndex = 1
    _PCDATA_InlTreeIndex  = 2
    _PCDATA_ArgLiveIndex  = 3

    _PCDATA_INVALID_OFFSET = 0
)

const (
    // PCDATA_UnsafePoint values.
    PCDATA_UnsafePointSafe   = -1 // Safe for async preemption
    PCDATA_UnsafePointUnsafe = -2 // Unsafe for async preemption

    // PCDATA_Restart1(2) apply on a sequence of instructions, within
    // which if an async preemption happens, we should back off the PC
    // to the start of the sequence when resume.
    // We need two so we can distinguish the start/end of the sequence
    // in case that two sequences are next to each other.
    PCDATA_Restart1 = -3
    PCDATA_Restart2 = -4

    // Like PCDATA_RestartAtEntry, but back to function entry if async
    // preempted.
    PCDATA_RestartAtEntry = -5

    _PCDATA_START_VAL = -1
)

type Pcvalue struct {
    PC  uint32 // PC offset from func entry
    Val int32
}

type Pcdata []Pcvalue

// see https://docs.google.com/document/d/1lyPIbmsYbXnpNj57a261hgOYVpNRcgydurVQIyZOz_o/pub
func (self Pcdata) MarshalBinary() (data []byte, err error) {
    // delta value always starts from -1
    sv := int32(_PCDATA_START_VAL)
    sp := uint32(0)
    for _, v := range self {
        data = append(data, encodeVariant(toZigzag(int(v.Val - sv)))...)
        data = append(data, encodeVariant(toZigzag(int(v.PC - sp)))...)
        sp = v.PC
        sv = v.Val
    }
    return
}

// makePctab generates pcdelta->valuedelta tables for functions,
// and returns the table and the entry offset of every kind pcdata in the table.
func makePctab(funcs []Func, cuOffset []uint32, nameOffset []int32) (pctab []byte, pcdataOffs [][]uint32, _funcs []_func) {
    _funcs = make([]_func, len(funcs))

    // Pctab offsets of 0 are considered invalid in the runtime. We respect
    // that by just padding a single byte at the beginning of runtime.pctab,
    // that way no real offsets can be zero.
    pctab = make([]byte, 1, 12*len(funcs)+1)
    pcdataOffs = make([][]uint32, len(funcs))

    for i, f := range funcs {
        _f := &_funcs[i]

        var writer = func(pc *Pcdata) {
            var ab []byte
            var err error
            if pc != nil {
                ab, err = pc.MarshalBinary()
                if err != nil {
                    panic(err)
                }
                pcdataOffs[i] = append(pcdataOffs[i], uint32(len(pctab)))
            } else {
                ab = []byte{0}
                pcdataOffs[i] = append(pcdataOffs[i], _PCDATA_INVALID_OFFSET)
            }
            pctab = append(pctab, ab...)
        }

        if f.Pcsp != nil {
            _f.pcsp = uint32(len(pctab))
        }
        writer(f.Pcsp)
        if f.Pcfile != nil {
            _f.pcfile = uint32(len(pctab))
        }
        writer(f.Pcfile)
        if f.Pcline != nil {
            _f.pcln = uint32(len(pctab))
        }
        writer(f.Pcline)
        writer(f.PcUnsafePoint)
        writer(f.PcStackMapIndex)
        writer(f.PcInlTreeIndex)
        writer(f.PcArgLiveIndex)
        
        _f.entryOff = f.EntryOff
        _f.nameOff = nameOffset[i]
        _f.args = f.ArgsSize
        _f.deferreturn = f.DeferReturn
        // NOTICE: _func.pcdata is always as [PCDATA_UnsafePoint(0) : PCDATA_ArgLiveIndex(3)]
        _f.npcdata = uint32(_N_PCDATA)
        _f.cuOffset = cuOffset[i]
        _f.funcID = f.ID
        _f.flag = f.Flag
        _f.nfuncdata = uint8(_N_FUNCDATA)
    }

    return
}