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
#include "utils.h"
#include <stdint.h>

/** String Quoting **/
#define MAX_ESCAPED_BYTES 8
typedef struct {
    const long n;
    const char s[MAX_ESCAPED_BYTES];
} quoted_t;

static const quoted_t _SingleQuoteTab[256] = {
    ['\x00'] = { .n = 6, .s = "\\u0000" },
    ['\x01'] = { .n = 6, .s = "\\u0001" },
    ['\x02'] = { .n = 6, .s = "\\u0002" },
    ['\x03'] = { .n = 6, .s = "\\u0003" },
    ['\x04'] = { .n = 6, .s = "\\u0004" },
    ['\x05'] = { .n = 6, .s = "\\u0005" },
    ['\x06'] = { .n = 6, .s = "\\u0006" },
    ['\x07'] = { .n = 6, .s = "\\u0007" },
    ['\b'  ] = { .n = 6, .s = "\\u0008" },
    ['\t'  ] = { .n = 2, .s = "\\t"     },
    ['\n'  ] = { .n = 2, .s = "\\n"     },
    ['\x0b'] = { .n = 6, .s = "\\u000b" },
    ['\f'  ] = { .n = 6, .s = "\\u000c" },
    ['\r'  ] = { .n = 2, .s = "\\r"     },
    ['\x0e'] = { .n = 6, .s = "\\u000e" },
    ['\x0f'] = { .n = 6, .s = "\\u000f" },
    ['\x10'] = { .n = 6, .s = "\\u0010" },
    ['\x11'] = { .n = 6, .s = "\\u0011" },
    ['\x12'] = { .n = 6, .s = "\\u0012" },
    ['\x13'] = { .n = 6, .s = "\\u0013" },
    ['\x14'] = { .n = 6, .s = "\\u0014" },
    ['\x15'] = { .n = 6, .s = "\\u0015" },
    ['\x16'] = { .n = 6, .s = "\\u0016" },
    ['\x17'] = { .n = 6, .s = "\\u0017" },
    ['\x18'] = { .n = 6, .s = "\\u0018" },
    ['\x19'] = { .n = 6, .s = "\\u0019" },
    ['\x1a'] = { .n = 6, .s = "\\u001a" },
    ['\x1b'] = { .n = 6, .s = "\\u001b" },
    ['\x1c'] = { .n = 6, .s = "\\u001c" },
    ['\x1d'] = { .n = 6, .s = "\\u001d" },
    ['\x1e'] = { .n = 6, .s = "\\u001e" },
    ['\x1f'] = { .n = 6, .s = "\\u001f" },
    ['"'   ] = { .n = 2, .s = "\\\""    },
    ['\\'  ] = { .n = 2, .s = "\\\\"    },
};

static const quoted_t _DoubleQuoteTab[256] = {
    ['\x00'] = { .n = 7, .s = "\\\\u0000" },
    ['\x01'] = { .n = 7, .s = "\\\\u0001" },
    ['\x02'] = { .n = 7, .s = "\\\\u0002" },
    ['\x03'] = { .n = 7, .s = "\\\\u0003" },
    ['\x04'] = { .n = 7, .s = "\\\\u0004" },
    ['\x05'] = { .n = 7, .s = "\\\\u0005" },
    ['\x06'] = { .n = 7, .s = "\\\\u0006" },
    ['\x07'] = { .n = 7, .s = "\\\\u0007" },
    ['\b'  ] = { .n = 7, .s = "\\\\u0008" },
    ['\t'  ] = { .n = 3, .s = "\\\\t"     },
    ['\n'  ] = { .n = 3, .s = "\\\\n"     },
    ['\x0b'] = { .n = 7, .s = "\\\\u000b" },
    ['\f'  ] = { .n = 7, .s = "\\\\u000c" },
    ['\r'  ] = { .n = 3, .s = "\\\\r"     },
    ['\x0e'] = { .n = 7, .s = "\\\\u000e" },
    ['\x0f'] = { .n = 7, .s = "\\\\u000f" },
    ['\x10'] = { .n = 7, .s = "\\\\u0010" },
    ['\x11'] = { .n = 7, .s = "\\\\u0011" },
    ['\x12'] = { .n = 7, .s = "\\\\u0012" },
    ['\x13'] = { .n = 7, .s = "\\\\u0013" },
    ['\x14'] = { .n = 7, .s = "\\\\u0014" },
    ['\x15'] = { .n = 7, .s = "\\\\u0015" },
    ['\x16'] = { .n = 7, .s = "\\\\u0016" },
    ['\x17'] = { .n = 7, .s = "\\\\u0017" },
    ['\x18'] = { .n = 7, .s = "\\\\u0018" },
    ['\x19'] = { .n = 7, .s = "\\\\u0019" },
    ['\x1a'] = { .n = 7, .s = "\\\\u001a" },
    ['\x1b'] = { .n = 7, .s = "\\\\u001b" },
    ['\x1c'] = { .n = 7, .s = "\\\\u001c" },
    ['\x1d'] = { .n = 7, .s = "\\\\u001d" },
    ['\x1e'] = { .n = 7, .s = "\\\\u001e" },
    ['\x1f'] = { .n = 7, .s = "\\\\u001f" },
    ['"'   ] = { .n = 4, .s = "\\\\\\\""  },
    ['\\'  ] = { .n = 4, .s = "\\\\\\\\"  },
};

static const quoted_t _HtmlQuoteTab[256] = {
    ['<'] = { .n = 6, .s = "\\u003c" },
    ['>'] = { .n = 6, .s = "\\u003e" },
    ['&'] = { .n = 6, .s = "\\u0026" },
    // \u2028 and \u2029 is [E2 80 A8] and [E2 80 A9]
    [0xe2] = { .n = 0, .s = {0} },
    [0xa8] = { .n = 6, .s = "\\u2028" },
    [0xa9] = { .n = 6, .s = "\\u2029" },
};

static inline __m128i _mm_find_quote(__m128i vv) {
    __m128i e1 = _mm_cmpgt_epi8   (vv, _mm_set1_epi8(-1));
    __m128i e2 = _mm_cmpgt_epi8   (vv, _mm_set1_epi8(31));
    __m128i e3 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('"'));
    __m128i e4 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('\\'));
    __m128i r1 = _mm_andnot_si128 (e2, e1);
    __m128i r2 = _mm_or_si128     (e3, e4);
    __m128i rv = _mm_or_si128     (r1, r2);
    return rv;
}

#if USE_AVX2
static inline __m256i _mm256_find_quote(__m256i vv) {
    __m256i e1 = _mm256_cmpgt_epi8   (vv, _mm256_set1_epi8(-1));
    __m256i e2 = _mm256_cmpgt_epi8   (vv, _mm256_set1_epi8(31));
    __m256i e3 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('"'));
    __m256i e4 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('\\'));
    __m256i r1 = _mm256_andnot_si256 (e2, e1);
    __m256i r2 = _mm256_or_si256     (e3, e4);
    __m256i rv = _mm256_or_si256     (r1, r2);
    return rv;
}
#endif

static inline ssize_t memcchr_quote(const char *sp, ssize_t nb, char *dp, ssize_t dn) {
    uint32_t     mm;
    const char * ss = sp;

#if USE_AVX2
    /* 32-byte loop, full store */
    while (nb >= 32 && dn >= 32) {
        __m256i vv = _mm256_loadu_si256  ((const void *)sp);
        __m256i rv = _mm256_find_quote   (vv);
                     _mm256_storeu_si256 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm256_movemask_epi8(rv)) != 0) {
            return sp - ss + __builtin_ctz(mm);
        }

        /* move to next block */
        sp += 32;
        dp += 32;
        nb -= 32;
        dn -= 32;
    }

    /* 32-byte test, partial store */
    if (nb >= 32) {
        __m256i  vv = _mm256_loadu_si256   ((const void *)sp);
        __m256i  rv = _mm256_find_quote    (vv);
        uint32_t mv = _mm256_movemask_epi8 (rv);
        uint32_t fv = __builtin_ctzll      ((uint64_t)mv | 0x0100000000);

        /* copy at most `dn` characters */
        if (fv <= dn) {
            memcpy_p32(dp, sp, fv);
            return sp - ss + fv;
        } else {
            memcpy_p32(dp, sp, dn);
            return -(sp - ss + dn) - 1;
        }
    }

    /* clear upper half to avoid AVX-SSE transition penalty */
    _mm256_zeroupper();
#endif

    /* 16-byte loop, full store */
    while (nb >= 16 && dn >= 16) {
        __m128i vv = _mm_loadu_si128  ((const void *)sp);
        __m128i rv = _mm_find_quote   (vv);
                     _mm_storeu_si128 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm_movemask_epi8(rv)) != 0) {
            return sp - ss + __builtin_ctz(mm);
        }

        /* move to next block */
        sp += 16;
        dp += 16;
        nb -= 16;
        dn -= 16;
    }

    /* 16-byte test, partial store */
    if (nb >= 16) {
        __m128i  vv = _mm_loadu_si128   ((const void *)sp);
        __m128i  rv = _mm_find_quote    (vv);
        uint32_t mv = _mm_movemask_epi8 (rv);
        uint32_t fv = __builtin_ctz     (mv | 0x010000);

        /* copy at most `dn` characters */
        if (fv <= dn) {
            memcpy_p16(dp, sp, fv);
            return sp - ss + fv;
        } else {
            memcpy_p16(dp, sp, dn);
            return -(sp - ss + dn) - 1;
        }
    }

    /* handle the remaining bytes with scalar code */
    while (nb > 0 && dn > 0) {
        if (_SingleQuoteTab[*(uint8_t *)sp].n) {
            return sp - ss;
        } else {
            dn--, nb--;
            *dp++ = *sp++;
        }
    }

    /* check for dest buffer */
    if (nb == 0) {
        return sp - ss;
    } else {
        return -(sp - ss) - 1;
    }
}

static const bool _EscTab[256] = {
    1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 0x00-0x0F
    1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 0x10-0x1F
    //   '"'
    0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0x20-0x2F
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0x30-0x3F
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0x40-0x4F
    //                                 '""
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, // 0x50-0x5F
    // 0x60-0xFF are zeroes
};

static inline uint8_t escape_mask4(const char *sp) {
    return _EscTab[*(uint8_t *)(sp)] | (_EscTab[*(uint8_t *)(sp + 1)] << 1) | (_EscTab[*(uint8_t *)(sp + 2)] << 2) | (_EscTab[*(uint8_t *)(sp + 3)]  << 3);
}

static inline ssize_t memcchr_quote_unsafe(const char *sp, ssize_t nb, char *dp, const quoted_t * tab) {
    uint32_t     mm;
    const char * ds = dp;
    size_t cn = 0;

simd_copy:

    if (nb < 16) goto scalar_copy;

#if USE_AVX2
    /* 32-byte loop, full store */
    while (nb >= 32) {
        __m256i vv = _mm256_loadu_si256  ((const void *)sp);
        __m256i rv = _mm256_find_quote   (vv);
                     _mm256_storeu_si256 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm256_movemask_epi8(rv)) != 0) {
            cn = __builtin_ctz(mm);
            sp += cn;
            nb -= cn;
            dp += cn;
            goto escape;
        }

        /* move to next block */
        sp += 32;
        dp += 32;
        nb -= 32;
    }

    /* clear upper half to avoid AVX-SSE transition penalty */
    _mm256_zeroupper();
#endif

    /* 16-byte loop, full store */
    while (nb >= 16) {
        __m128i vv = _mm_loadu_si128  ((const void *)sp);
        __m128i rv = _mm_find_quote   (vv);
                     _mm_storeu_si128 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm_movemask_epi8(rv)) != 0) {
            cn =  __builtin_ctz(mm);
            sp += cn;
            nb -= cn;
            dp += cn;
            goto escape;
        }

        /* move to next block */
        sp += 16;
        dp += 16;
        nb -= 16;
    }

    /* handle the remaining bytes with scalar code */
    // while (nb > 0) {
    //     if (_EscTab[*(uint8_t *)sp]) {
    //         goto escape;
    //     } else {
    //         nb--;
    //         *dp++ = *sp++;
    //     }
    // }
    // optimize: loop unrolling here

scalar_copy:
    if (nb >= 8) {
        uint8_t mask1 = escape_mask4(sp);
        *(uint64_t *)dp = *(const uint64_t *)sp;
        if (unlikely(mask1)) {
            cn =  __builtin_ctz(mask1);
            sp += cn;
            nb -= cn;
            dp += cn;
            goto escape;
        }
        uint8_t mask2 = escape_mask4(sp + 4);
        if (unlikely(mask2)) {
            cn =  __builtin_ctz(mask2);
            sp += cn + 4;
            nb -= cn + 4;
            dp += cn + 4;
            goto escape;
        }
        dp += 8, sp += 8, nb -= 8;
    }

    if (nb >= 4) {
        uint8_t mask2 = escape_mask4(sp);
        *(uint32_t *)dp = *(const uint32_t *)sp;
        if (unlikely(mask2)) {
            cn =  __builtin_ctz(mask2);
            sp += cn;
            nb -= cn;
            dp += cn;
            goto escape;
        }
        dp += 4, sp += 4, nb -= 4;
    }

    while (nb > 0) {
        if (unlikely(_EscTab[*(uint8_t *)(sp)])) goto escape;
        *dp++ = *sp++, nb--;
    }
    /* all quote done */
    return dp - ds;
escape:
     /* get the escape entry, handle consecutive quotes */
     do {
        uint8_t ch = *(uint8_t *)sp;
        int nc = tab[ch].n;
        /* copy the quoted value.
         * Note: dp always has at least 8 bytes (MAX_ESCAPED_BYTES) here.
         * so, we not use memcpy_p8(dp, tab[ch].s, nc);
         */
        *(uint64_t *)dp = *(const uint64_t *)tab[ch].s;
        sp++;
        nb--;
        dp += nc;
        if (nb <= 0) break;
        /* copy and find escape chars */
        if (_EscTab[*(uint8_t *)(sp)] == 0) {
            goto simd_copy;
        }
    } while (true);
    return dp - ds;
}

ssize_t quote(const char *sp, ssize_t nb, char *dp, ssize_t *dn, uint64_t flags) {
    ssize_t          nd = *dn;
    const char *     ds = dp;
    const char *     ss = sp;
    const quoted_t * tab;

    /* select quoting table */
    if (!(flags & F_DBLUNQ)) {
        tab = _SingleQuoteTab;
    } else {
        tab = _DoubleQuoteTab;
    }

    if (*dn >= nb * MAX_ESCAPED_BYTES) {
        *dn = memcchr_quote_unsafe(sp, nb, dp, tab);
        return nb;
    }

    /* find the special characters, copy on the fly */
    while (nb != 0) {
        int     nc;
        uint8_t ch;
        ssize_t rb = memcchr_quote(sp, nb, dp, nd);

        /* not enough buffer space */
        if (rb < 0) {
            *dn = dp - ds - rb - 1;
            return -(sp - ss - rb - 1) - 1;
        }

        /* skip already copied bytes */
        sp += rb;
        dp += rb;
        nb -= rb;
        nd -= rb;

        /* get the escape entry, handle consecutive quotes */
        while (nb != 0) {
            ch = *(uint8_t *)sp;
            nc = tab[ch].n;

            /* check for escape character */
            if (nc == 0) {
                break;
            }

            /* check for buffer space */
            if (nc > nd) {
                *dn = dp - ds;
                return -(sp - ss) - 1;
            }

            /* copy the quoted value */
            memcpy_p8(dp, tab[ch].s, nc);
            sp++;
            nb--;
            dp += nc;
            nd -= nc;
        }
    }

    /* all done */
    *dn = dp - ds;
    return sp - ss;
}

/** String Unquoting **/

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

static inline ssize_t memcchr_p32(const char *s, ssize_t nb, char *p) {
    int64_t      r;
    ssize_t      n = nb;
    const char * q = s;

#if USE_AVX2
    __m256i u;
    __m256i v;
    __m256i b = _mm256_set1_epi8('\\');

    /* process every 32 bytes */
    while (n >= 32) {
        u = _mm256_loadu_si256  ((const void *)s);
        v = _mm256_cmpeq_epi8   (u, b);
            _mm256_storeu_si256 ((void *)p, u);

        /* check for matches */
        if ((r = _mm256_movemask_epi8(v)) != 0) {
            return s - q + __builtin_ctzll(r);
        }

        /* move to the next 32 bytes */
        s += 32;
        p += 32;
        n -= 32;
    }

    /* clear upper half to avoid AVX-SSE transition penalty */
    _mm256_zeroupper();
#endif

    /* initialze with '\\' */
    __m128i x;
    __m128i y;
    __m128i a = _mm_set1_epi8('\\');

    /* process every 16 bytes */
    while (n >= 16) {
        x = _mm_loadu_si128  ((const void *)s);
        y = _mm_cmpeq_epi8   (x, a);
            _mm_storeu_si128 ((void *)p, x);

        /* check for matches */
        if ((r = _mm_movemask_epi8(y)) != 0) {
            return s - q + __builtin_ctzll(r);
        }

        /* move to the next 16 bytes */
        s += 16;
        p += 16;
        n -= 16;
    }

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

    /* from line 598 */
    retry_decode:

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
                if (likely(flags & F_UNIREP)) {
                    unirep(&dp);
                    continue;
                } else {
                    *ep = x;
                    return -ERR_EOF;
                }
            } else {
                if (sp[0] == '\\') {
                    nb--;
                    sp++;
                } else if (likely(flags & F_UNIREP)) {
                    unirep(&dp);
                    continue;
                } else {
                    *ep = sp - s - 4;
                    return -ERR_UNICODE;
                }
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
            for (int i = 2; i < 6 && ishex(sp[i]); i++) ++*ep;
            return -ERR_INVAL;
        }

        /* decode the second code-point */
        r1 = unhex16_fast(sp + 2);
        sp += 6;
        nb -= 6;

        /* it must be the other half */
        if (r1 < 0xdc00 || r1 > 0xdfff) {
            if (unlikely(!(flags & F_UNIREP))) {
                *ep = sp - s - 4;
                return -ERR_UNICODE;
            } else {
                r0 = r1;
                unirep(&dp);
                goto retry_decode;
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

static inline __m128i _mm_find_html(__m128i vv) {
    __m128i e1 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('<'));
    __m128i e2 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('>'));
    __m128i e3 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('&'));
    __m128i e4 = _mm_cmpeq_epi8   (vv, _mm_set1_epi8('\xe2'));
    __m128i r1 = _mm_or_si128     (e1, e2);
    __m128i r2 = _mm_or_si128     (e3, e4);
    __m128i rv = _mm_or_si128     (r1, r2);
    return rv;
}

#if USE_AVX2
static inline __m256i _mm256_find_html(__m256i vv) {
    __m256i e1 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('<'));
    __m256i e2 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('>'));
    __m256i e3 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('&'));
    __m256i e4 = _mm256_cmpeq_epi8   (vv, _mm256_set1_epi8('\xe2'));
    __m256i r1 = _mm256_or_si256     (e1, e2);
    __m256i r2 = _mm256_or_si256     (e3, e4);
    __m256i rv = _mm256_or_si256     (r1, r2);
    return rv;
}
#endif

static inline ssize_t memcchr_html_quote(const char *sp, ssize_t nb, char *dp, ssize_t dn) {
    uint32_t     mm;
    const char * ss = sp;

#if USE_AVX2
    /* 32-byte loop, full store */
    while (nb >= 32 && dn >= 32) {
        __m256i vv = _mm256_loadu_si256  ((const void *)sp);
        __m256i rv = _mm256_find_html    (vv);
                     _mm256_storeu_si256 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm256_movemask_epi8(rv)) != 0) {
            return sp - ss + __builtin_ctz(mm);
        }

        /* move to next block */
        sp += 32;
        dp += 32;
        nb -= 32;
        dn -= 32;
    }

    /* 32-byte test, partial store */
    if (nb >= 32) {
        __m256i  vv = _mm256_loadu_si256   ((const void *)sp);
        __m256i  rv = _mm256_find_html     (vv);
        uint32_t mv = _mm256_movemask_epi8 (rv);
        uint32_t fv = __builtin_ctzll      ((uint64_t)mv | 0x0100000000);

        /* copy at most `dn` characters */
        if (fv <= dn) {
            memcpy_p32(dp, sp, fv);
            return sp - ss + fv;
        } else {
            memcpy_p32(dp, sp, dn);
            return -(sp - ss + dn) - 1;
        }
    }

    /* clear upper half to avoid AVX-SSE transition penalty */
    _mm256_zeroupper();
#endif

    /* 16-byte loop, full store */
    while (nb >= 16 && dn >= 16) {
        __m128i vv = _mm_loadu_si128  ((const void *)sp);
        __m128i rv =  _mm_find_html   (vv);
                     _mm_storeu_si128 ((void *)dp, vv);

        /* check for matches */
        if ((mm = _mm_movemask_epi8(rv)) != 0) {
            return sp - ss + __builtin_ctz(mm);
        }

        /* move to next block */
        sp += 16;
        dp += 16;
        nb -= 16;
        dn -= 16;
    }

    /* 16-byte test, partial store */
    if (nb >= 16) {
        __m128i  vv = _mm_loadu_si128   ((const void *)sp);
        __m128i  rv =  _mm_find_html    (vv);
        uint32_t mv = _mm_movemask_epi8 (rv);
        uint32_t fv = __builtin_ctz     (mv | 0x010000);

        /* copy at most `dn` characters */
        if (fv <= dn) {
            memcpy_p16(dp, sp, fv);
            return sp - ss + fv;
        } else {
            memcpy_p16(dp, sp, dn);
            return -(sp - ss + dn) - 1;
        }
    }

    /* handle the remaining bytes with scalar code */
    while (nb > 0 && dn > 0) {
        if (*sp == '<' || *sp == '>' || *sp == '&' || *sp == '\xe2') {
            return sp - ss;
        } else {
            dn--, nb--;
            *dp++ = *sp++;
        }
    }

    /* check for dest buffer */
    if (nb == 0) {
        return sp - ss;
    } else {
        return -(sp - ss) - 1;
    }
}

ssize_t html_escape(const char *sp, ssize_t nb, char *dp, ssize_t *dn) {
    ssize_t          nd  = *dn;
    const char     * ds  = dp;
    const char     * ss  = sp;
    const quoted_t * tab = _HtmlQuoteTab;

    /* find the special characters, copy on the fly */
    while (nb > 0) {
        int     nc = 0;
        uint8_t ch = 0;
        ssize_t rb = 0;
        const char * cur = 0;

        /* not enough buffer space */
        if (nd <= 0) {
            return -(sp - ss) - 1;
        }

        /* find and copy */
        if ((rb = memcchr_html_quote(sp, nb, dp, nd)) < 0) {
            *dn = dp - ds - rb - 1;
            return -(sp - ss - rb - 1) - 1;
        }

        /* skip already copied bytes */
        sp += rb;
        dp += rb;
        nb -= rb;
        nd -= rb;

        /* stop if already finished */
        if (nb <= 0) {
            break;
        }

        /* mark cur postion */
        cur = sp;

        /* check for \u2028 and \u2029, binary is \xe2\x80\xa8 and \xe2\x80\xa9 */
        if (unlikely(*sp == '\xe2')) {
            if (nb >= 3 && *(sp+1) == '\x80' && (*(sp+2) == '\xa8' || *(sp+2) == '\xa9')) {
                sp += 2, nb -= 2;
            } else if (nd > 0) {
                *dp++ = *sp++;
                nb--, nd--;
                continue;
            } else {
                return -(sp - ss) - 1;
            }
        }

        /* get the escape entry, handle consecutive quotes */
        ch = * (uint8_t*) sp;
        nc = tab[ch].n;


        /* check for buffer space */
        if (nd < nc) {
            *dn = dp - ds;
            return -(cur - ss) - 1;
        }

        /* copy the quoted value */
        memcpy_p8(dp, tab[ch].s, nc);
        sp++;
        nb--;
        dp += nc;
        nd -= nc;
    }

    /* all done */
    *dn = dp - ds;
    return sp - ss;
}

#undef MAX_ESCAPED_BYTES

static inline long unescape(const char** src, const char* end, char* dp) {
    const char* sp = *src;
    long nb = end - sp;
    char cc = 0;
    uint32_t r0, r1;

    if (nb <= 0) return -ERR_EOF;

    if ((cc = _UnquoteTab[(uint8_t)sp[1]]) == 0) {
        *src += 1;
        return -ERR_ESCAPE;
    }

    if (cc != -1) {
        *dp = cc;
        *src += 2;
        return 1;
    }

    if (nb < 4) {
        *src += 1;
        return -ERR_EOF;
    }

    /* check for hexadecimal characters */
    if (!unhex16_is(sp + 2)) {
        *src += 2;
        return -ERR_INVAL;
    }

    /* decode the code-point */
    r0 = unhex16_fast(sp + 2);
    sp += 6;
    *src = sp;

    /* ASCII characters, unlikely */
    if (unlikely(r0 <= 0x7f)) {
        *dp++ = (char)r0;
        return 1;
    }

    /* latin-1 characters, unlikely */
    if (unlikely(r0 <= 0x07ff)) {
        *dp++ = (char)(0xc0 | (r0 >> 6));
        *dp++ = (char)(0x80 | (r0 & 0x3f));
        return 2;
    }

    /* 3-byte characters, likely */
    if (likely(r0 < 0xd800 || r0 > 0xdfff)) {
        *dp++ = (char)(0xe0 | ((r0 >> 12)       ));
        *dp++ = (char)(0x80 | ((r0 >>  6) & 0x3f));
        *dp++ = (char)(0x80 | ((r0      ) & 0x3f));
        return 3;
    }

    /* surrogate half, must follows by the other half */
    if (nb < 6 || r0 > 0xdbff || sp[0] != '\\' || sp[1] != 'u') {
        return -ERR_UNICODE;
    }

    /* check the hexadecimal escape */
    if (!unhex16_is(sp + 2)) {
        *src += 2;
        return -ERR_INVAL;
    }

    /* decode the second code-point */
    r1 = unhex16_fast(sp + 2);

    /* it must be the other half */
    if (r1 < 0xdc00 || r1 > 0xdfff) {
        *src += 2;
        return -ERR_UNICODE;
    }

    /* merge two surrogates */
    r0 = (r0 - 0xd800) << 10;
    r1 = (r1 - 0xdc00) + 0x010000;
    r0 += r1;

    /* encode the character */
    *dp++ = (char)(0xf0 | ((r0 >> 18)       ));
    *dp++ = (char)(0x80 | ((r0 >> 12) & 0x3f));
    *dp++ = (char)(0x80 | ((r0 >>  6) & 0x3f));
    *dp++ = (char)(0x80 | ((r0      ) & 0x3f));
    *src = sp + 6;
    return 4;
}