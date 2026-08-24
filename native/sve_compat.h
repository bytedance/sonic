/*
 * Copyright 2026 ByteDance Inc.
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

/*
 * Vector-length agnostic helpers for the SVE natives.
 *
 * SVE is a vector-length agnostic ISA: one encoding runs on a 128-bit and a
 * 512-bit implementation alike. This port was not written that way. It modelled
 * SVE as "AVX2 with different spelling", assuming a vector is exactly 32 bytes,
 * and extracted lane masks by reading a predicate register's spill as a fixed
 * uint32:
 *
 *	svbool_t pg = svcmpeq_n_u8(svptrue_b8(), v, c);
 *	uint32_t *bits = (uint32_t *)&pg;
 *
 * A byte predicate holds one bit per byte lane, so it occupies VL/8 bytes in
 * memory: 4 bytes at VL=32, but only 2 at VL=16. Reading a fixed 4 therefore
 * yields 16 real bits and 16 bits of neighbouring stack on a 128-bit machine,
 * while the surrounding loops still advance 32 bytes per iteration. Nothing
 * faults -- every instruction is legal -- so it presents as silent corruption.
 * Measured on Graviton4 (Neoverse V2, 128-bit): skip_one_fast returned 31 where
 * 42 was expected.
 *
 * The helpers here keep each algorithm's 32- and 64-byte block structure, which
 * the escape and odd/even bit twiddling depends on, and fix only how masks are
 * built: a block is covered by predicated chunks of min(VL, 32) bytes. At VL=32
 * that is one chunk per 32-byte block, matching what this port always emitted.
 *
 * Note these natives link -nostdlib against a custom linker script, so nothing
 * here may emit a libc call. That rules out variable-length __builtin_memcpy,
 * which lowers to a call to memcpy; every read below is a fixed size so it
 * lowers to a single load.
 */

#pragma once

#include "native.h"

#if defined(__SVE__)

#include <arm_sve.h>
#include <stdint.h>

/* Bytes per SVE vector on this machine, known only at run time. */
#define SVE_VL_BYTES ((uint64_t)svcntb())

/*
 * Bytes of input handled per chunk. Capped at 32 so a chunk never needs more
 * than 32 lane bits, which keeps the predicate read to a single fixed-size
 * load. A wider machine simply leaves its upper lanes predicated off.
 */
static always_inline uint64_t sve_chunk_bytes(void) {
    uint64_t vl = SVE_VL_BYTES;
    return vl > 32 ? 32 : vl;
}

/*
 * Lane bits of a byte predicate: bit i is lane i, one bit per byte.
 *
 * The predicate occupies VL/8 bytes, so read exactly as many as this machine
 * has: 2 bytes at VL=16, 4 at VL>=32. Both are fixed-size reads. Reading 4 on a
 * 128-bit machine is what the original code did, and is where the garbage came
 * from.
 *
 * Only the low sve_chunk_bytes() lanes are meaningful to callers; a wider
 * machine's remaining lanes were predicated off by the caller and read as zero.
 */
static always_inline uint64_t sve_pred_bits(const svbool_t *pg) {
    if (SVE_VL_BYTES >= 32) {
        uint32_t bits;
        __builtin_memcpy(&bits, pg, sizeof(bits));
        return (uint64_t)bits;
    } else {
        uint16_t bits;
        __builtin_memcpy(&bits, pg, sizeof(bits));
        return (uint64_t)bits;
    }
}

/*
 * One bit per byte over the n-byte window at s, set where the byte equals c.
 * n must be <= 64. Loads are predicated, so no byte outside the window is read
 * even when the vector is wider than the window.
 */
static always_inline uint64_t sve_mask_eq(const char *s, uint64_t n, uint8_t c) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mask = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t eq = svcmpeq_n_u8(pg, v, c);
        mask |= sve_pred_bits(&eq) << off;
    }
    return mask;
}

/*
 * One bit per byte over the n-byte window at s, set where the byte is less than
 * c compared as unsigned. n must be <= 64.
 */
static always_inline uint64_t sve_mask_lt(const char *s, uint64_t n, uint8_t c) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mask = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t lt = svcmplt_n_u8(pg, v, c);
        mask |= sve_pred_bits(&lt) << off;
    }
    return mask;
}

/*
 * The three masks the string scanner needs over an n-byte window, computed in a
 * single pass so each chunk is loaded once rather than once per predicate:
 *
 *	quote - byte == '"'
 *	bs    - byte == '\\'
 *	ctrl  - byte <  0x20, i.e. a control character
 *
 * ctrl matches the AVX2 helper it replaces: that computed (v > -1) & ~(v > 31)
 * on signed bytes, which admits 0x00..0x1F and excludes 0x80..0xFF because they
 * are negative. Comparing unsigned against 0x20 is the same set.
 *
 * n must be <= 64.
 */
static always_inline void sve_string_masks(const char *s, uint64_t n,
                                           uint64_t *quote, uint64_t *bs,
                                           uint64_t *ctrl) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mq = 0, mb = 0, mc = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t eq_q = svcmpeq_n_u8(pg, v, (uint8_t)'"');
        svbool_t eq_b = svcmpeq_n_u8(pg, v, (uint8_t)'\\');
        svbool_t lt_c = svcmplt_n_u8(pg, v, (uint8_t)0x20);
        mq |= sve_pred_bits(&eq_q) << off;
        mb |= sve_pred_bits(&eq_b) << off;
        mc |= sve_pred_bits(&lt_c) << off;
    }
    *quote = mq;
    *bs = mb;
    *ctrl = mc;
}

/*
 * One bit per byte over the n-byte window at s, set where the byte equals any
 * of a, b or c. Computed in a single pass. n must be <= 64.
 */
static always_inline uint64_t sve_mask_eq3(const char *s, uint64_t n,
                                           uint8_t a, uint8_t b, uint8_t c) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mask = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t ea = svcmpeq_n_u8(pg, v, a);
        svbool_t eb = svcmpeq_n_u8(pg, v, b);
        svbool_t ec = svcmpeq_n_u8(pg, v, c);
        mask |= (sve_pred_bits(&ea) | sve_pred_bits(&eb) | sve_pred_bits(&ec)) << off;
    }
    return mask;
}

/*
 * One bit per byte over the n-byte window, set where s1 and s2 differ.
 * n must be <= 64.
 */
static always_inline uint64_t sve_mask_ne2(const char *s1, const char *s2,
                                           uint64_t n) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mask = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t a = svld1_u8(pg, (const uint8_t *)s1 + off);
        svuint8_t b = svld1_u8(pg, (const uint8_t *)s2 + off);
        svbool_t ne = svcmpne_u8(pg, a, b);
        mask |= sve_pred_bits(&ne) << off;
    }
    return mask;
}

/*
 * The four masks the number scanner needs over an n-byte window, in one pass:
 *
 *	digit - '0' <= byte <= '9'
 *	dot   - byte == '.'
 *	exp   - byte == 'e' or 'E'
 *	sign  - byte == '+' or '-'
 *
 * digit matches the signed comparisons it replaces -- (v > '/') & ~(v > '9') --
 * which admit exactly '0'..'9' and reject 0x80..0xFF because those are negative
 * as signed bytes. An unsigned range check is the same set.
 *
 * n must be <= 64.
 */
static always_inline void sve_number_masks(const char *s, uint64_t n,
                                           uint64_t *digit, uint64_t *dot,
                                           uint64_t *exp, uint64_t *sign) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mv = 0, md = 0, me = 0, ms = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t ge0 = svcmpge_n_u8(pg, v, (uint8_t)'0');
        svbool_t le9 = svcmple_n_u8(pg, v, (uint8_t)'9');
        svbool_t edot = svcmpeq_n_u8(pg, v, (uint8_t)'.');
        svbool_t ee = svcmpeq_n_u8(pg, v, (uint8_t)'e');
        svbool_t eE = svcmpeq_n_u8(pg, v, (uint8_t)'E');
        svbool_t ep = svcmpeq_n_u8(pg, v, (uint8_t)'+');
        svbool_t em = svcmpeq_n_u8(pg, v, (uint8_t)'-');
        mv |= (sve_pred_bits(&ge0) & sve_pred_bits(&le9)) << off;
        md |= sve_pred_bits(&edot) << off;
        me |= (sve_pred_bits(&ee) | sve_pred_bits(&eE)) << off;
        ms |= (sve_pred_bits(&ep) | sve_pred_bits(&em)) << off;
    }
    *digit = mv;
    *dot = md;
    *exp = me;
    *sign = ms;
}

/*
 * One bit per byte over the n-byte window at s, set where the byte is NOT
 * one of the four JSON whitespace characters (space, tab, LF, CR). n must be
 * <= 64.
 *
 * The AVX2/SSE originals this replaces answer the same question with a
 * shuffle-table trick: index a 32-byte table by byte&0x1F and compare against
 * the original byte. That trick does not survive SVE unchanged -- the table
 * itself is loaded as an svuint8_t, so at VL<32 the register holds only the
 * table's first VL bytes, and svtbl_u8 zeroes any index the register doesn't
 * have a lane for. A byte whose low 5 bits pick an index past that truncated
 * width (e.g. 0x39, low bits 25) then looks up as 0 instead of matching,
 * misclassifying it. Since there are only four whitespace bytes, comparing
 * against each directly sidesteps the table (and the truncation) entirely.
 */
static always_inline uint64_t sve_nonspace_mask(const char *s, uint64_t n) {
    uint64_t step = sve_chunk_bytes();
    uint64_t mask = 0;
    for (uint64_t off = 0; off < n; off += step) {
        svbool_t pg = svwhilelt_b8_u64(off, n);
        svuint8_t v = svld1_u8(pg, (const uint8_t *)s + off);
        svbool_t sp = svcmpeq_n_u8(pg, v, (uint8_t)' ');
        svbool_t tab = svcmpeq_n_u8(pg, v, (uint8_t)'\t');
        svbool_t lf = svcmpeq_n_u8(pg, v, (uint8_t)'\n');
        svbool_t cr = svcmpeq_n_u8(pg, v, (uint8_t)'\r');
        uint64_t space = sve_pred_bits(&sp) | sve_pred_bits(&tab) |
                         sve_pred_bits(&lf) | sve_pred_bits(&cr);
        mask |= space << off;
    }
    return ~mask;
}

#endif /* __SVE__ */
