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

#if USE_AVX2
static const uintptr_t ALIGN_MASK = 31;
#else
static const uintptr_t ALIGN_MASK = 15;
#endif

size_t lspace(const char *sp, size_t nb, size_t p) {
    int32_t      ms;
    const char * ss = sp;

    /* seek to `p` */
    sp += p;
    nb -= p;

    /* likely to run into non-spaces within a few characters, try scalar code first */
    while (nb > 0 && ((uintptr_t)sp & ALIGN_MASK)) {
        switch ((nb--, *sp++)) {
            case ' '  : break;
            case '\r' : break;
            case '\n' : break;
            case '\t' : break;
            default   : return sp - ss - 1;
        }
    }

#if USE_AVX2
    /* 32-byte loop */
    while (likely(nb >= 32)) {
        __m256i x = _mm256_load_si256 ((const void *)sp);
        __m256i a = _mm256_cmpeq_epi8 (x, _mm256_set1_epi8(' '));
        __m256i b = _mm256_cmpeq_epi8 (x, _mm256_set1_epi8('\t'));
        __m256i c = _mm256_cmpeq_epi8 (x, _mm256_set1_epi8('\n'));
        __m256i d = _mm256_cmpeq_epi8 (x, _mm256_set1_epi8('\r'));
        __m256i u = _mm256_or_si256   (a, b);
        __m256i v = _mm256_or_si256   (c, d);
        __m256i w = _mm256_or_si256   (u, v);

        /* check for matches */
        if ((ms = _mm256_movemask_epi8(w)) != -1) {
            _mm256_zeroupper();
            return sp - ss + __builtin_ctzll(~(uint64_t)ms);
        }

        /* move to next block */
        sp += 32;
        nb -= 32;
    }

    /* clear upper half to avoid AVX-SSE transition penalty */
    _mm256_zeroupper();
#endif

    /* 16-byte loop */
    while (likely(nb >= 16)) {
        __m128i x = _mm_load_si128 ((const void *)sp);
        __m128i a = _mm_cmpeq_epi8 (x, _mm_set1_epi8(' '));
        __m128i b = _mm_cmpeq_epi8 (x, _mm_set1_epi8('\t'));
        __m128i c = _mm_cmpeq_epi8 (x, _mm_set1_epi8('\n'));
        __m128i d = _mm_cmpeq_epi8 (x, _mm_set1_epi8('\r'));
        __m128i u = _mm_or_si128   (a, b);
        __m128i v = _mm_or_si128   (c, d);
        __m128i w = _mm_or_si128   (u, v);

        /* check for matches */
        if ((ms = _mm_movemask_epi8(w)) != 0xffff) {
            return sp - ss + __builtin_ctz(~ms);
        }

        /* move to next block */
        sp += 16;
        nb -= 16;
    }

    /* remaining bytes, do with scalar code */
    while (nb-- > 0) {
        switch (*sp++) {
            case ' '  : break;
            case '\r' : break;
            case '\n' : break;
            case '\t' : break;
            default   : return sp - ss - 1;
        }
    }

    /* all the characters are spaces */
    return sp - ss;
}