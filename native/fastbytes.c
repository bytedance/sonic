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

#if USE_SSE
#define loop_decl()         \
    size_t       v;         \
    size_t       n = 0;     \
    const char * p = s;     \

#define loop_simd(size, load, func, ...) {                                      \
    while (nb >= size) {                                                        \
        if ((v = func(load((const void *)(p)), ## __VA_ARGS__)) < size) {       \
            return n + v;                                                       \
        } else {                                                                \
            n += v;                                                             \
            p += size;                                                          \
            nb -= size;                                                         \
        }                                                                       \
    }                                                                           \
}

#if !USE_AVX2
#define loop_zero()
#define loop_m256(func, ...)
#else
#define loop_zero()             _mm256_zeroupper();
#define loop_m256(func, ...)    loop_simd(32, _mm256_loadu_si256, func, ## __VA_ARGS__)
#endif

#define loop_m128(func, ...)    loop_simd(16, _mm_loadu_si128, func, ## __VA_ARGS__)
#define loop_last(func, ...)    return func(_mm_loadu_si128(as_m128c(p + nb - 16)), ## __VA_ARGS__) + n + nb - 16;

#define loop_bulk(func, ...) {                  \
    loop_decl()                                 \
    loop_m256(func ## _avx2, ## __VA_ARGS__)    \
    loop_zero();                                \
    loop_m128(func ## _sse2, ## __VA_ARGS__)    \
    loop_last(func ## _sse2, ## __VA_ARGS__)    \
}

#define loop_duff(func, ...) {                              \
    size_t  r = nb;                                         \
    __m128i m = _mm_set1_epi8(0xff);                        \
                                                            \
    /* remaining bytes */                                   \
    switch (r) {                                            \
        case 15 : m = _mm_insert_epi8(m, s[14], 14);        \
        case 14 : m = _mm_insert_epi8(m, s[13], 13);        \
        case 13 : m = _mm_insert_epi8(m, s[12], 12);        \
        case 12 : m = _mm_insert_epi8(m, s[11], 11);        \
        case 11 : m = _mm_insert_epi8(m, s[10], 10);        \
        case 10 : m = _mm_insert_epi8(m, s[ 9],  9);        \
        case  9 : m = _mm_insert_epi8(m, s[ 8],  8);        \
        case  8 : m = _mm_insert_epi8(m, s[ 7],  7);        \
        case  7 : m = _mm_insert_epi8(m, s[ 6],  6);        \
        case  6 : m = _mm_insert_epi8(m, s[ 5],  5);        \
        case  5 : m = _mm_insert_epi8(m, s[ 4],  4);        \
        case  4 : m = _mm_insert_epi8(m, s[ 3],  3);        \
        case  3 : m = _mm_insert_epi8(m, s[ 2],  2);        \
        case  2 : m = _mm_insert_epi8(m, s[ 1],  1);        \
        case  1 : m = _mm_insert_epi8(m, s[ 0],  0);        \
        default : return func ## _sse2(m, ## __VA_ARGS__);  \
    }                                                       \
}

static inline size_t lspace_sse2(__m128i v0) {
    __m128i  v1 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8(' '));
    __m128i  v2 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8('\t'));
    __m128i  v3 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8('\n'));
    __m128i  v4 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8('\r'));
    __m128i  v5 = _mm_or_si128      (v1, v2);
    __m128i  v6 = _mm_or_si128      (v3, v4);
    __m128i  v7 = _mm_or_si128      (v5, v6);
    uint32_t v8 = _mm_movemask_epi8 (v7);
    uint32_t v9 = __builtin_ctz     (~v8);
    return v9;
}

#if USE_AVX2
static inline size_t lspace_avx2(__m256i v0) {
    __m256i  v1 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8(' '));
    __m256i  v2 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8('\t'));
    __m256i  v3 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8('\n'));
    __m256i  v4 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8('\r'));
    __m256i  v5 = _mm256_or_si256      (v1, v2);
    __m256i  v6 = _mm256_or_si256      (v3, v4);
    __m256i  v7 = _mm256_or_si256      (v5, v6);
    uint32_t v8 = _mm256_movemask_epi8 (v7);
    uint64_t v9 = __builtin_ctzll      (~(uint64_t)(v8));
    return v9;
}
#endif

static inline size_t lquote_sse2(__m128i v0) {
    __m128i  v1 = _mm_cmpgt_epi8    (v0, _mm_set1_epi8(-1));
    __m128i  v2 = _mm_cmplt_epi8    (v0, _mm_set1_epi8(' '));
    __m128i  v3 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8('"'));
    __m128i  v4 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8('\\'));
    __m128i  v5 = _mm_and_si128     (v1, v2);
    __m128i  v6 = _mm_or_si128      (v3, v4);
    __m128i  v7 = _mm_or_si128      (v5, v6);
    uint32_t v8 = _mm_movemask_epi8 (v7);
    uint32_t v9 = __builtin_ctz     (v8 | 0xffff0000);
    return v9;
}

#if USE_AVX2
static inline size_t lquote_avx2(__m256i v0) {
    __m256i  v1 = _mm256_cmpgt_epi8    (v0, _mm256_set1_epi8(-1));
    __m256i  v2 = _mm256_cmpgt_epi8    (v0, _mm256_set1_epi8(31));
    __m256i  v3 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8('"'));
    __m256i  v4 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8('\\'));
    __m256i  v5 = _mm256_andnot_si256  (v2, v1);
    __m256i  v6 = _mm256_or_si256      (v3, v4);
    __m256i  v7 = _mm256_or_si256      (v5, v6);
    uint32_t v8 = _mm256_movemask_epi8 (v7);
    uint64_t v9 = __builtin_ctzll      ((uint64_t)v8 | 0xffffffff00000000);
    return v9;
}
#endif

static inline size_t strchr2_sse2(__m128i v0, uint64_t c0, uint64_t c1) {
    __m128i  v1 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8((char)c0));
    __m128i  v2 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8((char)c1));
    __m128i  v3 = _mm_or_si128      (v1, v2);
    uint32_t v4 = _mm_movemask_epi8 (v3);
    uint32_t v5 = __builtin_ctz     (v4 | 0xffff0000);
    return v5;
}

#if USE_AVX2
static inline size_t strchr2_avx2(__m256i v0, uint64_t c0, uint64_t c1) {
    __m256i  v1 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8((char)c0));
    __m256i  v2 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8((char)c1));
    __m256i  v3 = _mm256_or_si256      (v1, v2);
    uint32_t v4 = _mm256_movemask_epi8 (v3);
    uint64_t v5 = __builtin_ctzll      ((uint64_t)v4 | 0xffffffff00000000);
    return v5;
}
#endif

#define do_simd(func, ...) {                \
    if (nb == 0) {                          \
        return 0;                           \
    } if (nb < 16) {                        \
        loop_duff(func, ## __VA_ARGS__)     \
    } else {                                \
        loop_bulk(func, ## __VA_ARGS__)     \
    }                                       \
}
#endif

#define is_quote(c) ((c) == '"' || (c) == '\\' || ((c) >= 0 && (c) <= 31))
#define is_space(c) ((c) == ' ' || (c) == '\t' || (c) == '\n' || (c) == '\r')

static inline size_t lspace_p(const char *s, size_t nb) {
#if USE_SSE
    do_simd(lspace)
#else
    size_t i = 0;
    while (i < nb && !is_space(s[i])) i++;
    return i;
#endif
}

static inline size_t lquote_p(const char *s, size_t nb) {
#if USE_SSE
    do_simd(lquote)
#else
    size_t i = 0;
    while (i < nb && !is_quote(s[i])) i++;
    return i;
#endif
}

static inline size_t strchr1_p(const char *p, size_t nb, uint64_t ch) {
#if USE_SSE
    int64_t  r;
    uint32_t t;

    /* prepare the vector */
    ssize_t      n = nb;
    uintptr_t    m = (uintptr_t)p;
    const char * q = p;

#if USE_AVX2
#define ALIGN_VAL           31
#define _mm_or              _mm256_or_si256
#define _mm_load            _mm256_load_si256
#define _mm_cmpeq(a, b)     _mm256_cmpeq_epi8(a, b)
#define _mm_testz(v)        _mm256_testz_si256(v, v)
#define _mm_movemask(v)     _mm256_movemask_epi8(v)
    __m256i a;
    __m256i b;
    __m256i c;
    __m256i d;
    __m256i u;
    __m256i v;
    __m256i w;
    __m256i x = _mm256_set1_epi8(ch);
#else
#define ALIGN_VAL           15
#define _mm_or              _mm_or_si128
#define _mm_load            _mm_load_si128
#define _mm_cmpeq(a, b)     _mm_cmpeq_epi8(a, b)
#define _mm_testz(v)        (_mm_movemask_epi8(v) == 0)
#define _mm_movemask(v)     _mm_movemask_epi8(v)
    __m128i a;
    __m128i b;
    __m128i c;
    __m128i d;
    __m128i u;
    __m128i v;
    __m128i w;
    __m128i x = _mm_set1_epi8(ch);
#endif

#define BLOCK_SIZE      (ALIGN_VAL + 1)
#define BLOCK_MASK      (1ull << BLOCK_SIZE)
#define BLOCK_LARGE     (BLOCK_SIZE * 4)

    /* check for pointer alignment */
    if (m & ALIGN_VAL) {
        v = _mm_load     ((const void *)(m & -BLOCK_SIZE));
        v = _mm_cmpeq    (v, x);
        r = _mm_movemask (v);

        /* check for match in the first characters */
        if ((r >>= (t = m & ALIGN_VAL)) != 0) {
            if ((r = __builtin_ctzll(r | BLOCK_MASK)) < n) {
                return r;
            } else {
                return -1;
            }
        }

        /* make the pointer aligned */
        p += BLOCK_SIZE - t;
        n -= BLOCK_SIZE - t;
    }

    /* attempt to compare 4 blocks at a time */
    while (n >= BLOCK_LARGE) {
        a = _mm_load  ((const void *)(p + BLOCK_SIZE * 0));
        b = _mm_load  ((const void *)(p + BLOCK_SIZE * 1));
        c = _mm_load  ((const void *)(p + BLOCK_SIZE * 2));
        d = _mm_load  ((const void *)(p + BLOCK_SIZE * 3));
        a = _mm_cmpeq (a, x);
        b = _mm_cmpeq (b, x);
        c = _mm_cmpeq (c, x);
        d = _mm_cmpeq (d, x);
        u = _mm_or    (a, b);
        v = _mm_or    (c, d);
        w = _mm_or    (u, v);

        /* check if anything matches */
        if (_mm_testz(w)) {
            p += BLOCK_LARGE;
            n -= BLOCK_LARGE;
            continue;
        }

        /* match something in the 4-blocks region */
        if ((r = _mm_movemask(a)) != 0) {
            return p - q + __builtin_ctzll(r | BLOCK_MASK);
        } else if ((r = _mm_movemask(b)) != 0) {
            return p - q + __builtin_ctzll(r | BLOCK_MASK) + BLOCK_SIZE;
        } else if ((r = _mm_movemask(c)) != 0) {
            return p - q + __builtin_ctzll(r | BLOCK_MASK) + BLOCK_SIZE * 2;
        } else {
            return p - q + __builtin_ctzll(_mm_movemask(d) | BLOCK_MASK) + BLOCK_SIZE * 3;
        }
    }

    /* check every block, at most 4 times */
    for (int i = 0; i < 4 && n >= 0; i++) {
        v = _mm_load     ((const void *)p);
        v = _mm_cmpeq    (v, x);
        r = _mm_movemask (v);

        /* found something */
        if (r != 0) {
            if ((r = __builtin_ctzll(r | BLOCK_MASK)) >= n) {
                return -1;
            } else {
                return p - q + r;
            }
        }

        /* otherwise advance to next block */
        p += BLOCK_SIZE;
        n -= BLOCK_SIZE;
    }

#undef _mm_load
#undef _mm_bitor
#undef _mm_cmpeq
#undef _mm_testz
#undef _mm_movemask
#undef ALIGN_VAL
#undef BLOCK_SIZE
#undef BLOCK_LARGE
#else
    for (size_t i = 0; i < nb; i++) {
        if (p[i] == ch) {
            return i;
        }
    }
#endif

    /* not found */
    return nb;
}

static inline size_t strchr2_p(const char *s, size_t nb, uint64_t c0, uint64_t c1) {
#if USE_SSE
    do_simd(strchr2, c0, c1)
#else
    size_t i = 0;
    while (i < nb && s[i] != c0 && s[i] != c1) i++;
    return i;
#endif
}

size_t lzero(const char *p, size_t n) {
#if USE_SSE
#if USE_AVX
    __m256i a;
    __m256i b;
    __m256i c;
    __m256i d;
    __m256i u;
    __m256i v;
    __m256i w;
    __m256i y = _mm256_set1_epi8(0xff);
    __m256i z = _mm256_setzero_si256();
    #define BLOCK_SIZE 32
#else
    __m128i a;
    __m128i b;
    __m128i c;
    __m128i d;
    __m128i u;
    __m128i v;
    __m128i w;
    __m128i z = _mm_setzero_si128();
    #define BLOCK_SIZE 16
#endif

#if USE_AVX2
#define _mm_load            _mm256_load_si256
#define _mm_and(a, b)       _mm256_and_si256(a, b)
#define _mm_cmpeq(a, b)     _mm256_cmpeq_epi8(a, b)
#define _mm_testinz(v)      (!_mm256_testc_si256(v, y))
#elif USE_AVX
#define _mm_load            _mm256_load_si256
#define _mm_and(a, b)       _mm256_and_ps((__m256)a, (__m256)b)
#define _mm_cmpeq(a, b)     _mm256_cmp_ps(a, b, _CMP_EQ_OQ)
#define _mm_testinz(v)      (!_mm256_testc_si256(v, y))
#else
#define _mm_load            _mm_load_si128
#define _mm_and(a, b)       _mm_and_si128(a, b)
#define _mm_cmpeq(a, b)     _mm_cmpeq_epi8(a, b)
#define _mm_testinz(v)      (_mm_movemask_epi8(v) != 0xffff)
#endif

    /* multi-block loop */
    while (n >= BLOCK_SIZE * 4) {
        a = _mm_load  ((const void *)(p + BLOCK_SIZE * 0));
        b = _mm_load  ((const void *)(p + BLOCK_SIZE * 1));
        c = _mm_load  ((const void *)(p + BLOCK_SIZE * 2));
        d = _mm_load  ((const void *)(p + BLOCK_SIZE * 3));
        a = _mm_cmpeq (a, z);
        b = _mm_cmpeq (b, z);
        c = _mm_cmpeq (c, z);
        d = _mm_cmpeq (d, z);
        u = _mm_and   (a, b);
        v = _mm_and   (c, d);
        w = _mm_and   (u, v);

        /* test for zeros */
        if (_mm_testinz(w)) {
            return 1;
        }

        /* move to next block */
        p += BLOCK_SIZE * 4;
        n -= BLOCK_SIZE * 4;
    }

    /* single block loop */
    while (n >= BLOCK_SIZE) {
        a = _mm_load  ((const void *)(p));
        b = _mm_cmpeq (a, z);

        /* test for zeros */
        if (_mm_testinz(b)) {
            return 1;
        }

        /* move to next block */
        p += BLOCK_SIZE;
        n -= BLOCK_SIZE;
    }

#undef _mm_load
#undef _mm_cmpeq
#undef _mm_bitand
#undef _mm_testinz
#undef BLOCK_SIZE
#endif

    /* 8 bytes loop */
    while (n >= 8) {
        if (*(uint64_t *)p) {
            return 1;
        } else {
            p += 8;
            n -= 8;
        }
    }

    /* 4 bytes test */
    if (n >= 4) {
        if (*(uint32_t *)p) {
            return 1;
        } else {
            p += 4;
            n -= 4;
        }
    }

    /* 2 bytes test */
    if (n >= 2) {
        if (*(uint16_t *)p) {
            return 1;
        } else {
            p += 2;
            n -= 2;
        }
    }

    /* the final byte */
    if (n == 0) {
        return 0;
    } else {
        return *p != 0;
    }
}

size_t lquote(const GoString *s, size_t p) {
    return lquote_p(s->buf + p, s->len - p) + p;
}

size_t lspace(const char *sp, size_t nb, size_t p) {
    return lspace_p(sp + p, nb - p) + p;
}

ssize_t strchr1(const GoString *s, size_t p, char ch) {
    size_t n = s->len - p;
    size_t v = strchr1_p(s->buf + p, n, ch);
    return v >= n ? -1 : v + p;
}

ssize_t strchr2(const GoString *s, size_t p, char c0, char c1) {
    size_t n = s->len - p;
    size_t v = strchr2_p(s->buf + p, n, c0, c1);
    return v >= n ? -1 : v + p;
}
