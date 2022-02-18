/*
 * Copyright (c) 2009 The Go Authors. All rights reserved.
 * Modifications Copyright 2021 ByteDance Inc.
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

// ascii: 0x00 ~ 0x7F
static inline int _mm_ascii_mask(__m128i vv) {
    return _mm_movemask_epi8(vv);
}

#if USE_AVX2

// ascii: 0x00 ~ 0x7F
static inline int _mm256_ascii_mask(__m256i vv) {
    return _mm256_movemask_epi8(vv);
}

#endif

static inline bool is_ascii(uint8_t ch) {
    return ch < 0x80;
}

// The default lowest and highest continuation byte.
const static uint8_t locb = 0x80;
const static uint8_t hicb = 0xBF;
const static uint8_t xx = 0xF1; // invalid: size 1
const static uint8_t as = 0xF0; // ASCII: size 1
const static uint8_t s1 = 0x02; // accept 0, size 2
const static uint8_t s2 = 0x13; // accept 1, size 3
const static uint8_t s3 = 0x03; // accept 0, size 3
const static uint8_t s4 = 0x23; // accept 2, size 3
const static uint8_t s5 = 0x34; // accept 3, size 4
const static uint8_t s6 = 0x04; // accept 0, size 4
const static uint8_t s7 = 0x44; // accept 4, size 4

// first is information about the first byte in a UTF-8 sequence.
static const uint8_t first[256] = {
	//   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x00-0x0F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x10-0x1F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x20-0x2F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x30-0x3F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x40-0x4F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x50-0x5F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x60-0x6F
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x70-0x7F
	//   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x80-0x8F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x90-0x9F
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xA0-0xAF
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xB0-0xBF
	xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xC0-0xCF
	s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xD0-0xDF
	s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3, // 0xE0-0xEF
	s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xF0-0xFF
};

// AcceptRange gives the range of valid values for the second byte in a UTF-8
// sequence.
struct AcceptRange {
	uint8_t lo; // lowest value for second byte.
	uint8_t hi; // highest value for second byte.
};

// ranges has size 16 to avoid bounds checks in the code that uses it.
const static struct AcceptRange ranges[5] = {
    {locb, hicb}, // 0
    {0xA0, hicb}, // 1
    {locb, 0x9F}, // 2
    {0x90, hicb}, // 3
    {locb, 0x8F}, // 4
};

//  UTF-8 code point  | first byte | second byte | thrid byte | fourth byte
//  U+0000  -  U+007F | 0___ ____
//  U+0080  -  U+07FF | 110_ ____  | 10__ ____
//  U+0800  -  U+D7FF | 1110 ____  | 10__ ____   | 10__ ____
//  U+D800  -  U+DFFF | reserved for UTF-16 surrogate pairs
//  U+E000  -  U+FFFF | 1110 ____  | 10__ ____   | 10__ ____
// U+10000 - U+10FFFF | 1111 0___  | 10__ ____   | 10__ ____  | 10__ ____
// checks non-ascii characters, and returns the utf-8 length
static inline ssize_t nonascii_is_utf8(const uint8_t* sp, size_t n) {
    uint8_t mask = first[sp[0]];
    uint8_t size = mask & 7;
    if (n < size) {
        return 0;
    }
    struct AcceptRange accept = ranges[mask >> 4];
    switch (size) {
        case 4 : if (sp[3] < locb || hicb < sp[3]) return 0;
        case 3 : if (sp[2] < locb || hicb < sp[2]) return 0;
        case 2 : if (sp[1] < accept.lo || accept.hi < sp[1]) return 0;
        case 1 : // only validate non-ascii chars here
        default: return 0;
    }
    return size;
}

ssize_t find_non_ascii(const uint8_t*sp, ssize_t nb, ssize_t rb) {
    const uint8_t* ss = sp;
    int64_t m;

#if USE_AVX2
    while (rb >= 32 && nb > 0) {
        __m256i v = _mm256_loadu_si256 ((const void *)(sp));
        if (unlikely((m = _mm256_ascii_mask(v)) != 0)) {
            return sp - ss + __builtin_ctzll(m);
        }
        rb -= 32;
        nb -= 32;
        sp += 32;
    }

    /* clear spper half to avoid AVX-SSE transition penalty */
     _mm256_zeroupper();
#endif
    while (rb >= 16 && nb > 0) {
        __m128i v = _mm_loadu_si128 ((const void *)(sp));
        if (unlikely((m = _mm_ascii_mask(v)) != 0)) {
            return sp - ss + __builtin_ctzll(m);
        }
        rb -= 16;
        nb -= 16;
        sp += 16;
    }

    /* remaining bytes, do with scalar code */
    while (nb--) {
        if (is_ascii(*sp)) {
            sp++;
        } else {
            return sp - ss;
        }
    }

    /* nothing found */
    return -1;
}

// utf8_validate validates whether the JSON string is valid UTF-8.
// nb is string length, rb is the remain JSON length
// return -1 if validate, otherwise, return the error postion.
ssize_t utf8_validate(const char *sp, ssize_t nb, ssize_t rb) {
    const uint8_t* p = (const uint8_t*)sp;
    const uint8_t* s = (const uint8_t*)sp;
    ssize_t n;
    ssize_t b;

    // Optimize for the continuous non-ascii chars */
    while (nb > 0 && (n = (!is_ascii(*p) ? 0 : find_non_ascii(p, nb, rb))) != -1) {
        /* not found non-ascii in string */
        if (n >= nb) {
            return -1;
        }

        nb -= n;
        rb -= n;
        p  += n;

        /* validate the non-ascii */
        if (unlikely((b = nonascii_is_utf8(p, nb)) == 0)) {
            return p - s;
        }

        nb -= b;
        rb -= b;
        p  += b;
    }

    return -1;
}