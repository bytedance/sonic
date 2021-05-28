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

static const char _UnquoteTab[256] = {
    ['/' ] = '/',
    ['"' ] = '"',
    ['b' ] = '\b',
    ['f' ] = '\f',
    ['n' ] = '\n',
    ['r' ] = '\r',
    ['t' ] = '\t',
    ['u' ] = -1,
    ['\\'] = '\\',
};

#define memcchr_p32_avx2()                          \
    while (n >= 32) {                               \
        u = _mm256_loadu_si256  ((const void *)s);  \
        v = _mm256_cmpeq_epi8   (u, b);             \
            _mm256_storeu_si256 ((void *)p, u);     \
                                                    \
        /* check for matches */                     \
        if ((r = _mm256_movemask_epi8(v)) != 0) {   \
            return s - q + _mm_tzcnt_64(r);         \
        }                                           \
                                                    \
        /* move to the next 32 bytes */             \
        s += 32;                                    \
        p += 32;                                    \
        n -= 32;                                    \
    }                                               \

#define memcchr_p32_sse2()                          \
    if (n >= 16) {                                  \
        x = _mm_loadu_si128  ((const void *)s);     \
        y = _mm_cmpeq_epi8   (x, a);                \
            _mm_storeu_si128 ((void *)p, x);        \
                                                    \
        /* check for matches */                     \
        if ((r = _mm_movemask_epi8(y)) != 0) {      \
            return s - q + _mm_tzcnt_64(r);         \
        }                                           \
                                                    \
        /* move to the next 16 bytes */             \
        s += 16;                                    \
        p += 16;                                    \
        n -= 16;                                    \
    }

static inline ssize_t memcchr_p32(const char *s, ssize_t nb, char *p) {
    int32_t      r;
    __m128i      x;
    __m128i      y;
    __m256i      u;
    __m256i      v;
    __m128i      a = _mm_set1_epi8('\\');
    __m256i      b = _mm256_set1_epi8('\\');
    ssize_t      n = nb;
    const char * q = s;

    /* scan & copy with SIMD */
    memcchr_p32_avx2();
    _mm256_zeroupper();
    memcchr_p32_sse2();

    /* remaining bytes, do with scalar code */
    while (n--) {
        if (*s != '\\') {
            *p++ = *s++;
        } else {
            return s - q;
        }
    }

    /* nothing found, but everything was copied */
    return -1;
}

#undef memcchr_p32_avx2
#undef memcchr_p32_sse2

#define ALL_01h     (~0ul / 255)
#define ALL_7fh     (ALL_01h * 127)
#define ALL_80h     (ALL_01h * 128)

static inline uint32_t hasless(uint32_t x, uint8_t n) {
    return (x - ALL_01h * n) & ~x & ALL_80h;
}

static inline uint32_t hasmore(uint32_t x, uint8_t n) {
    return (x + ALL_01h * (127 - n) | x) & ALL_80h;
}

static inline uint32_t hasbetween(uint32_t x, uint8_t m, uint8_t n) {
    return (ALL_01h * (127 + n) - (x & ALL_7fh) & ~x & (x & ALL_7fh) + ALL_01h * (127 - m)) & ALL_80h;
}

#undef ALL_01h
#undef ALL_7fh
#undef ALL_80h

static inline char ishex(char c) {
    return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F');
}

static inline void unirep(char **dp) {
    *(*dp)++ = 0xef;
    *(*dp)++ = 0xbf;
    *(*dp)++ = 0xbd;
}

static inline char unhex16_is(const char *s) {
    uint32_t v = *(uint32_t *)s;
    return !(hasless(v, '0') || hasmore(v, 'f') || hasbetween(v, '9', 'A') || hasbetween(v, 'F', 'a'));
}

static inline uint32_t unhex16_fast(const char *s) {
    uint32_t a = __builtin_bswap32(*(uint32_t *)s);
    uint32_t b = 9 * ((~a & 0x10101010) >> 4) + (a & 0x0f0f0f0f);
    uint32_t c = (b >> 4) | b;
    uint32_t d = ((c >> 8) & 0xff00) | (c & 0x00ff);
    return d;
}

ssize_t unquote(const char *sp, ssize_t nb, char *dp, ssize_t *ep, uint64_t flags) {
    ssize_t      n;
    ssize_t      x = nb;
    const char * s = sp;
    const char * p = dp;

    /* scan & copy all the non-escape characters */
    while (nb && (n = (*sp == '\\' ? 0 : memcchr_p32(sp, nb, dp))) != -1) {
        char     cc;
        uint32_t r0;
        uint32_t r1;

        /* skip the plain text */
        dp += n;
        sp += n + 2;
        nb -= n + 2;

        /* check for EOF */
        if (nb < 0) {
            *ep = x;
            return -ERR_EOF;
        }

        /* check for double unquote */
        if (unlikely(flags & F_DBLUNQ)) {
            int  nr = nb;
            char c1 = sp[-1];

            /* must have at least 1 character left */
            if (nr == 0) {
                *ep = x;
                return -ERR_EOF;
            }

            /* every quote must be a double quote */
            if (c1 != '\\') {
                *ep = sp - s - 1;
                return -ERR_INVAL;
            }

            /* special case of '\\\\' and '\\\"' */
            if (*sp == '\\') {
                if (nr < 2) {
                    *ep = x;
                    return -ERR_EOF;
                } else if (sp[1] != '"' && sp[1] != '\\') {
                    *ep = sp - s + 1;
                    return -ERR_INVAL;
                } else {
                    sp++;
                    nb--;
                }
            }

            /* skip the second escape */
            sp++;
            nb--;
        }

        /* check for escape sequence */
        if ((cc = _UnquoteTab[(uint8_t)sp[-1]]) == 0) {
            *ep = sp - s - 1;
            return -ERR_ESCAPE;
        }

        /* check for simple escape sequence */
        if (cc != -1) {
            *dp++ = cc;
            continue;
        }

        /* must have at least 4 characters */
        if (nb < 4) {
            *ep = x;
            return -ERR_EOF;
        }

        /* check for hexadecimal characters */
        if (!unhex16_is(sp)) {
            *ep = sp - s;
            for (int i = 0; i < 4 && ishex(*sp); i++, sp++) ++*ep;
            return -ERR_INVAL;
        }

        /* decode the code-point */
        r0 = unhex16_fast(sp);
        sp += 4;
        nb -= 4;

        /* ASCII characters, unlikely */
        if (unlikely(r0 <= 0x7f)) {
            *dp++ = (char)r0;
            continue;
        }

        /* latin-1 characters, unlikely */
        if (unlikely(r0 <= 0x07ff)) {
            *dp++ = (char)(0xc0 | (r0 >> 6));
            *dp++ = (char)(0x80 | (r0 & 0x3f));
            continue;
        }

        /* 3-byte characters, likely */
        if (likely(r0 < 0xd800 || r0 > 0xdfff)) {
            *dp++ = (char)(0xe0 | ((r0 >> 12)       ));
            *dp++ = (char)(0x80 | ((r0 >>  6) & 0x3f));
            *dp++ = (char)(0x80 | ((r0      ) & 0x3f));
            continue;
        }

        /* check for double unquote */
        if (unlikely(flags & F_DBLUNQ)) {
            if (nb < 1) {
                *ep = x;
                return -ERR_EOF;
            } else if (sp[0] != '\\') {
                *ep = sp - s - 4;
                return -ERR_UNICODE;
            } else {
                nb--;
                sp++;
            }
        }

        /* surrogate half, must follows by the other half */
        if (nb < 6 || r0 > 0xdbff || sp[0] != '\\' || sp[1] != 'u') {
            if (likely(flags & F_UNIREP)) {
                unirep(&dp);
                continue;
            } else {
                *ep = sp - s - ((flags & F_DBLUNQ) ? 5 : 4);
                return -ERR_UNICODE;
            }
        }

        /* check the hexadecimal escape */
        if (!unhex16_is(sp + 2)) {
            *ep = sp - s + 2;
            for (int i = 0; i < 4 && ishex(sp[2]); i++, sp++) ++*ep;
            return -ERR_INVAL;
        }

        /* decode the second code-point */
        r1 = unhex16_fast(sp + 2);
        sp += 6;
        nb -= 6;

        /* it must be the other half */
        if (r1 < 0xdc00 || r1 > 0xdfff) {
            if (likely(!(flags & F_UNIREP))) {
                *ep = sp - s - 4;
                return -ERR_UNICODE;
            } else {
                unirep(&dp);
                unirep(&dp);
                continue;
            }
        }

        /* merge two surrogates */
        r0 = (r0 - 0xd800) << 10;
        r1 = (r1 - 0xdc00) + 0x010000;
        r0 += r1;

        /* check the code point range */
        if (r0 > 0x10ffff) {
            if (likely(!(flags & F_UNIREP))) {
                *ep = sp - s - 4;
                return -ERR_UNICODE;
            } else {
                unirep(&dp);
                continue;
            }
        }

        /* encode the character */
        *dp++ = (char)(0xf0 | ((r0 >> 18)       ));
        *dp++ = (char)(0x80 | ((r0 >> 12) & 0x3f));
        *dp++ = (char)(0x80 | ((r0 >>  6) & 0x3f));
        *dp++ = (char)(0x80 | ((r0      ) & 0x3f));
    }

    /* calculate the result length */
    return dp + nb - p;
}
