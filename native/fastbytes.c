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

#define loop_decl()         \
    size_t       v;         \
    size_t       n = 0;     \
    const char * p = s;     \

#define loop_m128(func, ...) {                                                  \
    if (nb >= 16) {                                                             \
        if ((v = func(_mm_loadu_si128(as_m128c(p)), ## __VA_ARGS__)) < 16) {    \
            return n + v;                                                       \
        } else {                                                                \
            n += v;                                                             \
            p += 16;                                                            \
            nb -= 16;                                                           \
        }                                                                       \
    }                                                                           \
}

#define loop_m256(func, ...) {                                                      \
    while (nb >= 32) {                                                              \
        if ((v = func(_mm256_loadu_si256(as_m256c(p)), ## __VA_ARGS__)) < 32) {     \
            return n + v;                                                           \
        } else {                                                                    \
            n += v;                                                                 \
            p += 32;                                                                \
            nb -= 32;                                                               \
        }                                                                           \
    }                                                                               \
}

#define loop_last(func, ...) {                                                          \
    return func(_mm_loadu_si128(as_m128c(p + nb - 16)), ## __VA_ARGS__) + n + nb - 16;  \
}

#define loop_simd(func, ...) {                  \
    loop_decl()                                 \
    loop_m256(func ## _avx2, ## __VA_ARGS__)    \
    _mm256_zeroupper();                         \
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

static inline size_t strchr2_sse2(__m128i v0, uint64_t c0, uint64_t c1) {
    __m128i  v1 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8((char)c0));
    __m128i  v2 = _mm_cmpeq_epi8    (v0, _mm_set1_epi8((char)c1));
    __m128i  v3 = _mm_or_si128      (v1, v2);
    uint32_t v4 = _mm_movemask_epi8 (v3);
    uint32_t v5 = __builtin_ctz     (v4 | 0xffff0000);
    return v5;
}

static inline size_t strchr2_avx2(__m256i v0, uint64_t c0, uint64_t c1) {
    __m256i  v1 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8((char)c0));
    __m256i  v2 = _mm256_cmpeq_epi8    (v0, _mm256_set1_epi8((char)c1));
    __m256i  v3 = _mm256_or_si256      (v1, v2);
    uint32_t v4 = _mm256_movemask_epi8 (v3);
    uint64_t v5 = __builtin_ctzll      ((uint64_t)v4 | 0xffffffff00000000);
    return v5;
}

static inline size_t lspace_p(const char *s, size_t nb) {
    if (nb == 0) {
        return 0;
    } else if (nb < 16) {
        loop_duff(lspace)
    } else {
        loop_simd(lspace)
    }
}

static inline size_t lquote_p(const char *s, size_t nb) {
    if (nb == 0) {
        return 0;
    } else if (nb < 16) {
        loop_duff(lquote)
    } else {
        loop_simd(lquote)
    }
}

static inline size_t strchr1_p(const char *p, size_t nb, uint64_t ch) {
    __m256i  a;
    __m256i  b;
    __m256i  c;
    __m256i  d;
    __m256i  u;
    __m256i  v;
    __m256i  w;
    int32_t  r;
    uint32_t t;

    /* prepare the vector */
    __m256i      x = _mm256_set1_epi8(ch);
    ssize_t      n = nb;
    uintptr_t    m = (uintptr_t)p;
    const char * q = p;

    /* check for pointer alignment */
    if (m & 31) {
        v = _mm256_load_si256    ((const void *)(m & -32));
        v = _mm256_cmpeq_epi8    (v, x);
        r = _mm256_movemask_epi8 (v);

        /* check for match in the first characters */
        if ((r = r >> (t = m & 31)) != 0) {
            if ((r = _mm_tzcnt_64(r)) < n) {
                return r;
            } else {
                return -1;
            }
        }

        /* make the pointer aligned */
        p += 32 - t;
        n -= 32 - t;
    }

    /* attempt to compare 128-bytes at a time */
    while (n >= 128) {
        a = _mm256_cmpeq_epi8 (_mm256_load_si256((const void *)(p +  0)), x);
        b = _mm256_cmpeq_epi8 (_mm256_load_si256((const void *)(p + 32)), x);
        c = _mm256_cmpeq_epi8 (_mm256_load_si256((const void *)(p + 64)), x);
        d = _mm256_cmpeq_epi8 (_mm256_load_si256((const void *)(p + 96)), x);
        u = _mm256_or_si256   (a, b);
        v = _mm256_or_si256   (c, d);
        w = _mm256_or_si256   (u, v);

        /* check if anything matches */
        if (_mm256_testz_si256(w, w)) {
            p += 128;
            n -= 128;
            continue;
        }

        /* match something in the 128-byte region */
        if ((r = _mm256_movemask_epi8(a)) != 0) {
            return p - q + _mm_tzcnt_64(r);
        } else if ((r = _mm256_movemask_epi8(b)) != 0) {
            return p - q + _mm_tzcnt_64(r) + 32;
        } else if ((r = _mm256_movemask_epi8(c)) != 0) {
            return p - q + _mm_tzcnt_64(r) + 64;
        } else {
            return p - q + _mm_tzcnt_64(_mm256_movemask_epi8(d)) + 96;
        }
    }

    /* check every 32 bytes, at most 4 times */
    for (int i = 0; i < 4 && n >= 0; i++) {
        v = _mm256_cmpeq_epi8    (_mm256_load_si256((const void *)p), x);
        r = _mm256_movemask_epi8 (v);

        /* found something */
        if (r != 0) {
            if ((r = _mm_tzcnt_64(r)) >= n) {
                return -1;
            } else {
                return p - q + r;
            }
        }

        /* otherwise advance to next block */
        p += 32;
        n -= 32;
    }

    /* not found */
    return nb;
}

static inline size_t strchr2_p(const char *s, size_t nb, uint64_t c0, uint64_t c1) {
    if (nb == 0) {
        return 0;
    } else if (nb < 16) {
        loop_duff(strchr2, c0, c1)
    } else {
        loop_simd(strchr2, c0, c1)
    }
}

size_t lzero(const char *p, size_t n) {
    __m256i a;
    __m256i b;
    __m256i c;
    __m256i d;
    __m256i u;
    __m256i v;
    __m256i w;

    /* zero vector */
    size_t  r = 0;
    __m256i y = _mm256_set1_epi8(0xff);
    __m256i z = _mm256_setzero_si256();

    /* 128 bytes loop */
    while (n >= 128) {
        a = _mm256_cmpeq_epi8 (_mm256_loadu_si256(as_m256c(p +  0)), z);
        b = _mm256_cmpeq_epi8 (_mm256_loadu_si256(as_m256c(p + 32)), z);
        c = _mm256_cmpeq_epi8 (_mm256_loadu_si256(as_m256c(p + 64)), z);
        d = _mm256_cmpeq_epi8 (_mm256_loadu_si256(as_m256c(p + 96)), z);
        u = _mm256_and_si256  (a, b);
        v = _mm256_and_si256  (c, d);
        w = _mm256_xor_si256  (v, y);

        /* test for zeros */
        if (!_mm256_testc_si256(u, w)) {
            return 1;
        }

        /* move to next block */
        p += 128;
        n -= 128;
    }

    /* 32 bytes loop */
    while (n >= 32) {
        a = _mm256_loadu_si256 (as_m256c(p));
        b = _mm256_cmpeq_epi8  (a, z);

        /* test for zeros */
        if (!_mm256_testc_si256(b, y)) {
            return 1;
        }

        /* move to next block */
        p += 32;
        n -= 32;
    }

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
