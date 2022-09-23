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

#include "native.h"
#include "tab.h"

static const char Vec16xA0[16] __attribute__((aligned(16))) = {
    '0', '0', '0', '0', '0', '0', '0', '0',
    '0', '0', '0', '0', '0', '0', '0', '0',
};

static const uint16_t Vec8x10[8] __attribute__((aligned(16))) = {
    10, 10, 10, 10,
    10, 10, 10, 10,
};

static const uint32_t Vec4x10k[4] __attribute__((aligned(16))) = {
    10000,
    10000,
    10000,
    10000,
};

static const uint32_t Vec4xDiv10k[4] __attribute__((aligned(16))) = {
    0xd1b71759,
    0xd1b71759,
    0xd1b71759,
    0xd1b71759,
};

static const uint16_t VecDivPowers[8] __attribute__((aligned(16))) = {
    0x20c5, 0x147b,
    0x3334, 0x8000,
    0x20c5, 0x147b,
    0x3334, 0x8000,
};

static const uint16_t VecShiftPowers[8] __attribute__((aligned(16))) = {
    0x0080, 0x0800,
    0x2000, 0x8000,
    0x0080, 0x0800,
    0x2000, 0x8000,
};

static const uint8_t VecShiftShuffles[144] __attribute__((aligned(16))) = {
    0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
    0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff,
    0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff,
    0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff,
    0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff,
    0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff,
    0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
    0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
    0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
};

static inline int itoa1(char *out, int n, uint32_t v) {
    out[n++] = (char)v + '0';
    return n;
}

static inline int itoa2(char *out, int n, uint32_t v) {
    out[n++] = Digits[v];
    out[n++] = Digits[v + 1];
    return n;
}

static inline __m128i itoa8_sse2(uint32_t v) {
    __m128i v00 = _mm_cvtsi32_si128  (v);
    __m128i v01 = _mm_mul_epu32      (v00, as_m128v(Vec4xDiv10k));
    __m128i v02 = _mm_srli_epi64     (v01, 45);
    __m128i v03 = _mm_mul_epu32      (v02, as_m128v(Vec4x10k));
    __m128i v04 = _mm_sub_epi32      (v00, v03);
    __m128i v05 = _mm_unpacklo_epi16 (v02, v04);
    __m128i v06 = _mm_slli_epi64     (v05, 2);
    __m128i v07 = _mm_unpacklo_epi16 (v06, v06);
    __m128i v08 = _mm_unpacklo_epi32 (v07, v07);
    __m128i v09 = _mm_mulhi_epu16    (v08, as_m128v(VecDivPowers));
    __m128i v10 = _mm_mulhi_epu16    (v09, as_m128v(VecShiftPowers));
    __m128i v11 = _mm_mullo_epi16    (v10, as_m128v(Vec8x10));
    __m128i v12 = _mm_slli_epi64     (v11, 16);
    __m128i v13 = _mm_sub_epi16      (v10, v12);
    return v13;
}

static inline int u32toa_small(char *out, uint32_t val) {
    int      n  = 0;
    uint32_t d1 = (val / 100) << 1;
    uint32_t d2 = (val % 100) << 1;

    /* 1000-th digit */
    if (val >= 1000) {
        out[n++] = Digits[d1];
    }

    /* 100-th digit */
    if (val >= 100) {
        out[n++] = Digits[d1 + 1];
    }

    /* 10-th digit */
    if (val >= 10) {
        out[n++] = Digits[d2];
    }

    /* last digit */
    out[n++] = Digits[d2 + 1];
    return n;
}

static inline int u32toa_medium(char *out, uint32_t val) {
    int      n  = 0;
    uint32_t b  = val / 10000;
    uint32_t c  = val % 10000;
    uint32_t d1 = (b / 100) << 1;
    uint32_t d2 = (b % 100) << 1;
    uint32_t d3 = (c / 100) << 1;
    uint32_t d4 = (c % 100) << 1;

    /* 10000000-th digit */
    if (val >= 10000000) {
        out[n++] = Digits[d1];
    }

    /* 1000000-th digit */
    if (val >= 1000000) {
        out[n++] = Digits[d1 + 1];
    }

    /* 100000-th digit */
    if (val >= 100000) {
        out[n++] = Digits[d2];
    }

    /* remaining digits */
    out[n++] = Digits[d2 + 1];
    out[n++] = Digits[d3];
    out[n++] = Digits[d3 + 1];
    out[n++] = Digits[d4];
    out[n++] = Digits[d4 + 1];
    return n;
}

static inline int u64toa_large_sse2(char *out, uint64_t val) {
    uint32_t a = (uint32_t)(val / 100000000);
    uint32_t b = (uint32_t)(val % 100000000);

    /* convert to digits */
    __m128i v0 = itoa8_sse2(a);
    __m128i v1 = itoa8_sse2(b);

    /* convert to bytes, add '0' */
    __m128i v2 = _mm_packus_epi16 (v0, v1);
    __m128i v3 = _mm_add_epi8     (v2, as_m128v(Vec16xA0));

    /* count number of digit */
    __m128i  v4 = _mm_cmpeq_epi8    (v3, as_m128v(Vec16xA0));
    uint32_t bm = _mm_movemask_epi8 (v4);
    uint32_t nd = __builtin_ctz     (~bm | 0x8000);

    /* shift digits to the beginning */
    __m128i p = _mm_loadu_si128  (as_m128c(&VecShiftShuffles[nd * 16]));
    __m128i r = _mm_shuffle_epi8 (v3, p);

    /* store the result */
    _mm_storeu_si128(as_m128p(out), r);
    return 16 - nd;
}

static inline int u64toa_xlarge_sse2(char *out, uint64_t val) {
    int      n = 0;
    uint64_t b = val % 10000000000000000;
    uint32_t a = (uint32_t)(val / 10000000000000000);

    /* the highest 4 digits */
    if (a < 10) {
        n = itoa1(out, n, a);
    } else if (a < 100) {
        n = itoa2(out, n, a << 1);
    } else if (a < 1000) {
        n = itoa1(out, n, a / 100);
        n = itoa2(out, n, (a % 100) << 1);
    } else {
        n = itoa2(out, n, (a / 100) << 1);
        n = itoa2(out, n, (a % 100) << 1);
    }

    /* remaining digits */
    __m128i v0 = itoa8_sse2       ((uint32_t)(b / 100000000));
    __m128i v1 = itoa8_sse2       ((uint32_t)(b % 100000000));
    __m128i v2 = _mm_packus_epi16 (v0, v1);
    __m128i v3 = _mm_add_epi8     (v2, as_m128v(Vec16xA0));

    /* convert to bytes, add '0' */
    _mm_storeu_si128(as_m128p(&out[n]), v3);
    return n + 16;
}

int i64toa(char *out, int64_t val) {
    if (likely(val >= 0)) {
        return u64toa(out, (uint64_t)val);
    } else {
        *out = '-';
        return u64toa(out + 1, (uint64_t)(-val)) + 1;
    }
}

int u64toa(char *out, uint64_t val) {
    if (likely(val < 10000)) {
        return u32toa_small(out, (uint32_t)val);
    } else if (likely(val < 100000000)) {
        return u32toa_medium(out, (uint32_t)val);
    } else if (likely(val < 10000000000000000)) {
        return u64toa_large_sse2(out, val);
    } else {
        return u64toa_xlarge_sse2(out, val);
    }
}
