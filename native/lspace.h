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

#pragma once

#include "native.h"
#include "test/xassert.h"

static always_inline size_t lspace_1(const char *sp, size_t nb, size_t p) {
    const char * ss = sp;

    /* seek to `p` */
    sp += p;
    nb -= p;

    /* likely to run into non-spaces within a few characters, try scalar code first */
#if USE_AVX2
    __m256i space_tab = _mm256_setr_epi8(
        '\x20', 0, 0, 0, 0, 0, 0, 0,
         0, '\x09', '\x0A', 0, 0, '\x0D', 0, 0,
        '\x20', 0, 0, 0, 0, 0, 0, 0,
         0, '\x09', '\x0A', 0, 0, '\x0D', 0, 0
    );

    /* 32-byte loop */
    while (likely(nb >= 32)) {
        __m256i input = _mm256_loadu_si256((__m256i*)sp);
        __m256i shuffle = _mm256_shuffle_epi8(space_tab, input);
        __m256i result = _mm256_cmpeq_epi8(input, shuffle);
        int32_t mask = _mm256_movemask_epi8(result);
        if (mask != -1) {
            return sp - ss + __builtin_ctzll(~(uint64_t)mask);
        }
        sp += 32;
        nb -= 32;
    }
#endif

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
