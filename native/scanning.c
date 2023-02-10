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
#include "utf8.h"
#include "utils.h"

static const uint64_t ODD_MASK  = 0xaaaaaaaaaaaaaaaa;
static const uint64_t EVEN_MASK = 0x5555555555555555;

// NOTE: mask referenced from decoder/decoder.go
static const uint64_t MASK_VALIDATE_STRING = 1ull << 5;
static const uint64_t MASK_ALLOW_CONTROL   = 1ull << 31;

static const double P10_TAB[23] = {
    /* <= the connvertion to double is not exact when less than 1 => */     1e-000,
    1e+001, 1e+002, 1e+003, 1e+004, 1e+005, 1e+006, 1e+007, 1e+008, 1e+009, 1e+010,
    1e+011, 1e+012, 1e+013, 1e+014, 1e+015, 1e+016, 1e+017, 1e+018, 1e+019, 1e+020,
    1e+021, 1e+022 /* <= the connvertion to double is not exact when larger,  => */
};

static inline uint64_t add32(uint64_t v1, uint64_t v2, uint64_t *vo) {
    uint32_t v;
    uint32_t c = __builtin_uadd_overflow((uint32_t)v1, (uint32_t)v2, &v);

    /* set the carry */
    *vo = c;
    return v;
}

static inline uint64_t add64(uint64_t v1, uint64_t v2, uint64_t *vo) {
    unsigned long long v;
    uint64_t c = __builtin_uaddll_overflow(v1, v2, &v);

    /* set the carry */
    *vo = c;
    return v;
}

static inline char isspace(char ch) {
    return ch == ' ' || ch == '\r' || ch == '\n' | ch == '\t';
}

const int MASK_USE_NUMBER = 1<<1;

static inline void vdigits(const GoString *src, long *p, JsonState *ret, uint64_t flag) {
    --*p;
    if (flag & MASK_USE_NUMBER) {
        long i = skip_number(src, p);
        if (i < 0) {
            ret->vt = i;
            return;
        }
        ret->vt = V_DOUBLE;
        ret->ep = i;
        return;
    }
    vnumber(src, p, ret);
}

static inline char advance_ns(const GoString *src, long *p) {
    size_t       vi = *p;
    size_t       nb = src->len;
    const char * sp = src->buf;

    /* it's likely to run into non-spaces within a few
     * characters, so test up to 4 characters manually */
    if (vi < nb && !isspace(sp[vi])) goto nospace; else vi++;
    if (vi < nb && !isspace(sp[vi])) goto nospace; else vi++;
    if (vi < nb && !isspace(sp[vi])) goto nospace; else vi++;
    if (vi < nb && !isspace(sp[vi])) goto nospace; else vi++;

    /* check EOF */
    if (vi >= nb) {
        *p = vi;
        return 0;
    }

    /* too many spaces, use SIMD to search for characters */
    if ((vi = lspace(sp, nb, vi)) >= nb) {
        return 0;
    }

nospace:
    *p = vi + 1;
    return src->buf[vi];
}

static inline int64_t advance_dword(const GoString *src, long *p, long dec, int64_t ret, uint32_t val) {
    if (*p > src->len + dec - 4) {
        *p = src->len;
        return -ERR_EOF;
    } else if (*(uint32_t *)(src->buf + *p - dec) == val) {
        *p += 4 - dec;
        return ret;
    } else {
        *p -= dec;
        for (int i = 0; src->buf[*p] == (val & 0xff) && i < 4; i++, ++*p) { val >>= 8; }
        return -ERR_INVAL;
    }
}

static inline ssize_t advance_string_default(const GoString *src, long p, int64_t *ep) {
    char     ch;
    uint64_t es;
    uint64_t fe;
    uint64_t os;
    uint64_t m0;
    uint64_t m1;
    uint64_t cr = 0;

    /* prevent out-of-bounds accessing */
    if (unlikely(src->len == p)) {
        return -ERR_EOF;
    }

    /* buffer pointers */
    size_t       nb = src->len;
    const char * sp = src->buf;
    const char * ss = src->buf;

#define ep_init()   *ep = -1;
#define ep_setc()   ep_setx(sp - ss - 1)
#define ep_setx(x)  if (*ep == -1) { *ep = (x); }

    /* seek to `p` */
    nb -= p;
    sp += p;
    ep_init()

#if USE_AVX2
    /* initialize vectors */
    __m256i v0;
    __m256i v1;
    __m256i q0;
    __m256i q1;
    __m256i x0;
    __m256i x1;
    __m256i cq = _mm256_set1_epi8('"');
    __m256i cx = _mm256_set1_epi8('\\');

    /* partial masks */
    uint32_t s0;
    uint32_t s1;
    uint32_t t0;
    uint32_t t1;
#else
    /* initialize vectors */
    __m128i v0;
    __m128i v1;
    __m128i v2;
    __m128i v3;
    __m128i q0;
    __m128i q1;
    __m128i q2;
    __m128i q3;
    __m128i x0;
    __m128i x1;
    __m128i x2;
    __m128i x3;
    __m128i cq = _mm_set1_epi8('"');
    __m128i cx = _mm_set1_epi8('\\');

    /* partial masks */
    uint32_t s0;
    uint32_t s1;
    uint32_t s2;
    uint32_t s3;
    uint32_t t0;
    uint32_t t1;
    uint32_t t2;
    uint32_t t3;
#endif

#define m0_mask(add)                \
    m1 &= ~cr;                      \
    fe  = (m1 << 1) | cr;           \
    os  = (m1 & ~fe) & ODD_MASK;    \
    es  = add(os, m1, &cr) << 1;    \
    m0 &= ~(fe & (es ^ EVEN_MASK));

    /* 64-byte SIMD loop */
    while (likely(nb >= 64)) {
#if USE_AVX2
        v0 = _mm256_loadu_si256   ((const void *)(sp +  0));
        v1 = _mm256_loadu_si256   ((const void *)(sp + 32));
        q0 = _mm256_cmpeq_epi8    (v0, cq);
        q1 = _mm256_cmpeq_epi8    (v1, cq);
        x0 = _mm256_cmpeq_epi8    (v0, cx);
        x1 = _mm256_cmpeq_epi8    (v1, cx);
        s0 = _mm256_movemask_epi8 (q0);
        s1 = _mm256_movemask_epi8 (q1);
        t0 = _mm256_movemask_epi8 (x0);
        t1 = _mm256_movemask_epi8 (x1);
        m0 = ((uint64_t)s1 << 32) | (uint64_t)s0;
        m1 = ((uint64_t)t1 << 32) | (uint64_t)t0;
#else
        v0 = _mm_loadu_si128   ((const void *)(sp +  0));
        v1 = _mm_loadu_si128   ((const void *)(sp + 16));
        v2 = _mm_loadu_si128   ((const void *)(sp + 32));
        v3 = _mm_loadu_si128   ((const void *)(sp + 48));
        q0 = _mm_cmpeq_epi8    (v0, cq);
        q1 = _mm_cmpeq_epi8    (v1, cq);
        q2 = _mm_cmpeq_epi8    (v2, cq);
        q3 = _mm_cmpeq_epi8    (v3, cq);
        x0 = _mm_cmpeq_epi8    (v0, cx);
        x1 = _mm_cmpeq_epi8    (v1, cx);
        x2 = _mm_cmpeq_epi8    (v2, cx);
        x3 = _mm_cmpeq_epi8    (v3, cx);
        s0 = _mm_movemask_epi8 (q0);
        s1 = _mm_movemask_epi8 (q1);
        s2 = _mm_movemask_epi8 (q2);
        s3 = _mm_movemask_epi8 (q3);
        t0 = _mm_movemask_epi8 (x0);
        t1 = _mm_movemask_epi8 (x1);
        t2 = _mm_movemask_epi8 (x2);
        t3 = _mm_movemask_epi8 (x3);
        m0 = ((uint64_t)s3 << 48) | ((uint64_t)s2 << 32) | ((uint64_t)s1 << 16) | (uint64_t)s0;
        m1 = ((uint64_t)t3 << 48) | ((uint64_t)t2 << 32) | ((uint64_t)t1 << 16) | (uint64_t)t0;
#endif
        /** update first quote position */
        if (unlikely(m1 != 0)) {
            ep_setx(sp - ss + __builtin_ctzll(m1))
        }

        /** mask all the escaped quotes */
        if (unlikely(m1 != 0 || cr != 0)) {
            m0_mask(add64)
        }
       
        /* check for end quote */
        if (m0 != 0) {
            return sp - ss + __builtin_ctzll(m0) + 1;
        }

        /* move to the next block */
        sp += 64;
        nb -= 64;
    }

    /* 32-byte SIMD round */
    if (likely(nb >= 32)) {
#if USE_AVX2
        v0 = _mm256_loadu_si256   ((const void *)sp);
        q0 = _mm256_cmpeq_epi8    (v0, cq);
        x0 = _mm256_cmpeq_epi8    (v0, cx);
        s0 = _mm256_movemask_epi8 (q0);
        t0 = _mm256_movemask_epi8 (x0);
        m0 = (uint64_t)s0;
        m1 = (uint64_t)t0;
#else
        v0 = _mm_loadu_si128   ((const void *)(sp +  0));
        v1 = _mm_loadu_si128   ((const void *)(sp + 16));
        q0 = _mm_cmpeq_epi8    (v0, cq);
        q1 = _mm_cmpeq_epi8    (v1, cq);
        x0 = _mm_cmpeq_epi8    (v0, cx);
        x1 = _mm_cmpeq_epi8    (v1, cx);
        s0 = _mm_movemask_epi8 (q0);
        s1 = _mm_movemask_epi8 (q1);
        t0 = _mm_movemask_epi8 (x0);
        t1 = _mm_movemask_epi8 (x1);
        m0 = ((uint64_t)s1 << 16) | (uint64_t)s0;
        m1 = ((uint64_t)t1 << 16) | (uint64_t)t0;
#endif
       
        /** update first quote position */
        if (unlikely(m1 != 0)) {
            ep_setx(sp - ss + __builtin_ctzll(m1))
        }

        /** mask all the escaped quotes */
        if (unlikely(m1 != 0 || cr != 0)) {
            m0_mask(add32)
        }
       
        /* check for end quote */
        if (m0 != 0) {
            return sp - ss + __builtin_ctzll(m0) + 1;
        }

        /* move to the next block */
        sp += 32;
        nb -= 32;
    }

    /* check for carry */
    if (unlikely(cr != 0)) {
        if (nb == 0) {
            return -ERR_EOF;
        } else {
            ep_setc()
            sp++, nb--;
        }
    }

    /* handle the remaining bytes with scalar code */
    while (nb-- > 0 && (ch = *sp++) != '"') {
        if (unlikely(ch == '\\')) {
            if (nb == 0) {
                return -ERR_EOF;
            } else {
                ep_setc()
                sp++, nb--;
            }
        }
    }

#undef ep_init
#undef ep_setc
#undef ep_setx
#undef m0_mask

    /* check for quotes */
    if (ch == '"') {
        return sp - ss;
    } else {
        return -ERR_EOF;
    }
}

#if USE_AVX2

static inline int _mm256_get_mask(__m256i v, __m256i t) {
    return _mm256_movemask_epi8(_mm256_cmpeq_epi8(v, t));
}

// contrl char: 0x00 ~ 0x1F
static inline int _mm256_cchars_mask(__m256i v) {
    __m256i e1 = _mm256_cmpgt_epi8 (v, _mm256_set1_epi8(-1));
    __m256i e2 = _mm256_cmpgt_epi8 (v, _mm256_set1_epi8(31));
    return    _mm256_movemask_epi8 (_mm256_andnot_si256 (e2, e1));
}

// ascii: 0x00 ~ 0x7F
static inline int _mm256_nonascii_mask(__m256i v) {
    return _mm256_movemask_epi8(v);
}

#endif

static inline int _mm_get_mask(__m128i v, __m128i t) {
    return _mm_movemask_epi8(_mm_cmpeq_epi8(v, t));
}

// contrl char: 0x00 ~ 0x1F
static inline int _mm_cchars_mask(__m128i v) {
    __m128i e1 = _mm_cmpgt_epi8 (v, _mm_set1_epi8(-1));
    __m128i e2 = _mm_cmpgt_epi8 (v, _mm_set1_epi8(31));
    return    _mm_movemask_epi8 (_mm_andnot_si128 (e2, e1));
}

// ascii: 0x00 ~ 0x7F
static inline int _mm_nonascii_mask(__m128i v) {
    return _mm_movemask_epi8(v);
}

static inline ssize_t advance_string_validate(const GoString *src, long p, int64_t *ep) {
    char     ch;
    uint64_t m0, m1, m2;
    uint64_t es, fe, os;
    uint64_t cr = 0;
    long     qp = 0;
    long     np = 0;

    /* buffer pointers */
    size_t       nb = src->len;
    const char * sp = src->buf;
    const char * ss = src->buf;

    /* prevent out-of-bounds accessing */
    if (unlikely(nb == p)) {
        return -ERR_EOF;
    }

#define ep_init()    *ep = -1;
#define ep_setc()    ep_setx(sp - ss - 1)
#define ep_setx(x)   if (*ep == -1) { *ep = (x); }
#define ep_seterr(x)  *ep = (x);

    /* seek to `p` */
    nb -= p;
    sp += p;
    ep_init()

#if USE_AVX2
    /* initialize vectors */
    __m256i v0;
    __m256i v1;
    __m256i cq = _mm256_set1_epi8('"');
    __m256i cx = _mm256_set1_epi8('\\');

    /* partial masks */
    uint32_t s0, s1;
    uint32_t t0, t1;
    uint32_t c0, c1;
#else
    /* initialize vectors */
    __m128i v0;
    __m128i v1;
    __m128i v2;
    __m128i v3;
    __m128i cq = _mm_set1_epi8('"');
    __m128i cx = _mm_set1_epi8('\\');

    /* partial masks */
    uint32_t s0, s1, s2, s3;
    uint32_t t0, t1, t2, t3;
    uint32_t c0, c1, c2, c3;
#endif

#define m0_mask(add)                \
    m1 &= ~cr;                      \
    fe  = (m1 << 1) | cr;           \
    os  = (m1 & ~fe) & ODD_MASK;    \
    es  = add(os, m1, &cr) << 1;    \
    m0 &= ~(fe & (es ^ EVEN_MASK));

    /* 64-byte SIMD loop */
    while (likely(nb >= 64)) {
#if USE_AVX2
        v0 = _mm256_loadu_si256   ((const void *)(sp +  0));
        v1 = _mm256_loadu_si256   ((const void *)(sp + 32));
        s0 = _mm256_get_mask(v0, cq);
        s1 = _mm256_get_mask(v1, cq);
        t0 = _mm256_get_mask(v0, cx);
        t1 = _mm256_get_mask(v1, cx);
        c0 = _mm256_cchars_mask(v0);
        c1 = _mm256_cchars_mask(v1);
        m0 = ((uint64_t)s1 << 32) | (uint64_t)s0;
        m1 = ((uint64_t)t1 << 32) | (uint64_t)t0;
        m2 = ((uint64_t)c1 << 32) | (uint64_t)c0;
#else
        v0 = _mm_loadu_si128   ((const void *)(sp +  0));
        v1 = _mm_loadu_si128   ((const void *)(sp + 16));
        v2 = _mm_loadu_si128   ((const void *)(sp + 32));
        v3 = _mm_loadu_si128   ((const void *)(sp + 48));
        s0 = _mm_get_mask(v0, cq);
        s1 = _mm_get_mask(v1, cq);
        s2 = _mm_get_mask(v2, cq);
        s3 = _mm_get_mask(v3, cq);
        t0 = _mm_get_mask(v0, cx);
        t1 = _mm_get_mask(v1, cx);
        t2 = _mm_get_mask(v2, cx);
        t3 = _mm_get_mask(v3, cx);
        c0 = _mm_cchars_mask(v0);
        c1 = _mm_cchars_mask(v1);
        c2 = _mm_cchars_mask(v2);
        c3 = _mm_cchars_mask(v3);
        m0 = ((uint64_t)s3 << 48) | ((uint64_t)s2 << 32) | ((uint64_t)s1 << 16) | (uint64_t)s0;
        m1 = ((uint64_t)t3 << 48) | ((uint64_t)t2 << 32) | ((uint64_t)t1 << 16) | (uint64_t)t0;
        m2 = ((uint64_t)c3 << 48) | ((uint64_t)c2 << 32) | ((uint64_t)c1 << 16) | (uint64_t)c0;

#endif
       
        /** update first quote position */
        if (unlikely(m1 != 0)) {
            ep_setx(sp - ss + __builtin_ctzll(m1))
        }

        /** mask all the escaped quotes */
        if (unlikely(m1 != 0 || cr != 0)) {
            m0_mask(add64)
        }

        qp = m0 ? __builtin_ctzll(m0) : 64;
        np = m2 ? __builtin_ctzll(m2) : 64;
       
        /* get the position of end quote */
        if (m0 != 0) {
            /* check control chars in JSON string */
            if (unlikely(np < qp)) {
                ep_seterr(sp - ss + np)
               
                return -ERR_INVAL;
            }
            return sp - ss + qp + 1;
        }

        /* check control chars in JSON string */
        if (unlikely(m2 != 0)) {
            ep_setx(sp - ss + np)
           
            return -ERR_INVAL;
        }

        /* move to the next block */
        sp += 64;
        nb -= 64;
    }

    /* 32-byte SIMD round */
    if (likely(nb >= 32)) {
#if USE_AVX2
        v0 = _mm256_loadu_si256   ((const void *)sp);
        s0 = _mm256_get_mask (v0, cq);
        t0 = _mm256_get_mask (v0, cx);
        c0 = _mm256_cchars_mask(v0);
        m0 = (uint64_t)s0;
        m1 = (uint64_t)t0;
        m2 = (uint64_t)c0;
#else
        v0 = _mm_loadu_si128   ((const void *)(sp +  0));
        v1 = _mm_loadu_si128   ((const void *)(sp + 16));
        s0 = _mm_get_mask(v0, cq);
        s1 = _mm_get_mask(v1, cq);
        t0 = _mm_get_mask(v0, cx);
        t1 = _mm_get_mask(v1, cx);
        c0 = _mm_cchars_mask(v0);
        c1 = _mm_cchars_mask(v1);
        m0 = ((uint64_t)s1 << 16) | (uint64_t)s0;
        m1 = ((uint64_t)t1 << 16) | (uint64_t)t0;
        m2 = ((uint64_t)c1 << 16) | (uint64_t)c0;
#endif
       
        /** update first quote position */
        if (unlikely(m1 != 0)) {
            ep_setx(sp - ss + __builtin_ctzll(m1))
        }

        /** mask all the escaped quotes */
        if (unlikely(m1 != 0 || cr != 0)) {
            m0_mask(add32)
        }
       
        qp = m0 ? __builtin_ctzll(m0) : 64;
        np = m2 ? __builtin_ctzll(m2) : 64;
       
        /* get the position of end quote */
        if (m0 != 0) {
            if (unlikely(np < qp)) {
                ep_seterr(sp - ss + np)
                return -ERR_INVAL;
            }
            return sp - ss + qp + 1;
        }

        /* check control chars in JSON string */
        if (unlikely(m2 != 0)) {
            ep_seterr(sp - ss + __builtin_ctzll(m2))
            return -ERR_INVAL;
        }

        /* move to the next block */
        sp += 32;
        nb -= 32;
    }

    /* check for carry */
    if (unlikely(cr != 0)) {
        if (nb == 0) {
            return -ERR_EOF;
        } else {
            ep_setc()
            sp++, nb--;
        }
    }

    /* handle the remaining bytes with scalar code */
    while (nb > 0) {
        ch = *sp;
        if (ch == '"') {
           
            return sp - ss + 1;
        }

        /* valid the escaped chars */
        if (unlikely(ch == '\\')) {
            if (nb == 1) {
                return -ERR_EOF;
            }
            ep_setx(sp - ss)
            sp += 2, nb -= 2;
            continue;
        }

        /* valid unescaped chars */
        if (unlikely( ch >= 0 && ch <= 0x1f)) { // control chars
            ep_seterr(sp - ss)
            return -ERR_INVAL;
        }

        sp++, nb--;
    }
    return -ERR_EOF;
#undef ep_init
#undef ep_setc
#undef ep_setx
#undef ep_seterr
#undef m0_mask
}

static inline ssize_t advance_string(const GoString *src, long p, int64_t *ep, uint64_t flags) {
    if ((flags & MASK_VALIDATE_STRING) != 0) {
        return advance_string_validate(src, p, ep);
    } else {
        return advance_string_default(src, p, ep);
    }
}

/** Value Scanning Routines **/

long value(const char *s, size_t n, long p, JsonState *ret, uint64_t flags) {
    long     q = p;
    GoString m = {.buf = s, .len = n};
    bool allow_control = (flags & MASK_ALLOW_CONTROL) != 0;
    /* parse the next identifier, q is UNSAFE, may cause out-of-bounds accessing */
    switch (advance_ns(&m, &q)) {
        case '-' : /* fallthrough */
        case '0' : /* fallthrough */
        case '1' : /* fallthrough */
        case '2' : /* fallthrough */
        case '3' : /* fallthrough */
        case '4' : /* fallthrough */
        case '5' : /* fallthrough */
        case '6' : /* fallthrough */
        case '7' : /* fallthrough */
        case '8' : /* fallthrough */
        case '9' : vdigits(&m, &q, ret, flags)                          ; return q;
        case '"' : vstring(&m, &q, ret, flags)                          ; return q;
        case 'n' : ret->vt = advance_dword(&m, &q, 1, V_NULL, VS_NULL)  ; return q;
        case 't' : ret->vt = advance_dword(&m, &q, 1, V_TRUE, VS_TRUE)  ; return q;
        case 'f' : ret->vt = advance_dword(&m, &q, 0, V_FALSE, VS_ALSE) ; return q;
        case '[' : ret->vt = V_ARRAY                                    ; return q;
        case '{' : ret->vt = V_OBJECT                                   ; return q;
        case ':' : ret->vt = allow_control ? V_KEY_SEP : -ERR_INVAL     ; return allow_control ? q : q - 1;
        case ',' : ret->vt = allow_control ? V_ELEM_SEP : -ERR_INVAL    ; return allow_control ? q : q - 1;
        case ']' : ret->vt = allow_control ? V_ARRAY_END : -ERR_INVAL   ; return allow_control ? q : q - 1;
        case '}' : ret->vt = allow_control ? V_OBJECT_END : -ERR_INVAL  ; return allow_control ? q : q - 1;
        case  0  : ret->vt = V_EOF                                      ; return q;
        default  : ret->vt = -ERR_INVAL                                 ; return q - 1;
    }
}

void vstring(const GoString *src, long *p, JsonState *ret, uint64_t flags) {
    int64_t v = -1;
    int64_t i = *p;
    ssize_t e = advance_string(src, i, &v, flags);
   
    /* check for errors */
    if (e < 0) {
        *p = src->len;
        ret->vt = e;
        return;
    }

    /* update the result, and fix the escape position (escaping past the end of string) */
    *p = e;
    ret->iv = i;
    ret->vt = V_STRING;
    ret->ep = v >= e ? -1 : v;
}

#define set_vt(t)   \
    ret->vt = t;

#define init_ret(t) \
    ret->vt = t;    \
    ret->dv = 0.0;  \
    ret->iv = 0;    \
    ret->ep = *p;

#define check_eof()         \
    if (i >= n) {           \
        *p = n;             \
        ret->vt = -ERR_EOF; \
        return;             \
    }

#define check_sign(on_neg)  \
    if (s[i] == '-') {      \
        i++;                \
        on_neg;             \
        check_eof()         \
    }

#define check_digit()               \
    if (s[i] < '0' || s[i] > '9') { \
        *p = i;                     \
        ret->vt = -ERR_INVAL;       \
        return;                     \
    }

#define check_leading_zero()                                                                    \
    if (s[i] == '0' && (i >= n || (s[i + 1] != '.' && s[i + 1] != 'e' && s[i + 1] != 'E'))) {   \
        *p = ++i;                                                                               \
        return;                                                                                 \
    }

#define parse_sign(sgn)                 \
    if (s[i] == '+' || s[i] == '-') {   \
        sgn = s[i++] == '+' ? 1 : -1;   \
        check_eof()                     \
    }

#define is_digit(val) \
    '0' <= val && val <= '9'

#define add_integer_to_mantissa(man, man_nd, exp10, dig) \
    if (man_nd < 19) {                                   \
        man = man * 10 + dig;                            \
        man_nd++;                                        \
    } else {                                             \
        exp10++;                                         \
    }

#define add_float_to_mantissa(man, man_nd, exp10, dig) \
    man = man * 10 + dig;                              \
    man_nd++;                                          \
    exp10--;

#define parse_float_digits(val, sgn, ...)                       \
    while (i < n && s[i] >= '0' && s[i] <= '9' __VA_ARGS__) {   \
        val *= 10;                                              \
        val += sgn * (s[i++] - '0');                            \
    }

#define parse_integer_digits(val, sgn, ovf)                     \
    while (i < n && s[i] >= '0' && s[i] <= '9') {               \
        if (add_digit_overflow(val, sgn * (s[i++] - '0'))) {    \
            ovf = 1;                                            \
            break;                                              \
        }                                                       \
    }

#define add_digit_overflow(val, chr) (          \
    __builtin_mul_overflow(val, 10, &val) ||    \
    __builtin_add_overflow(val, chr, &val)      \
)

#define vinteger(type, sgn, on_neg)                     \
    int  ovf = 0;                                       \
    type val = 0;                                       \
                                                        \
    /* initial buffer pointers */                       \
    long         i = *p;                                \
    size_t       n = src->len;                          \
    const char * s = src->buf;                          \
                                                        \
    /* initialize the result, and check for '-' */      \
    init_ret(V_INTEGER)                                 \
    check_eof()                                         \
    check_sign(on_neg)                                  \
                                                        \
    /* check for leading zero or any digits */          \
    check_digit()                                       \
    check_leading_zero()                                \
    parse_integer_digits(val, sgn, ovf)                 \
                                                        \
    /* check for overflow */                            \
    if (ovf) {                                          \
        *p = i - 1;                                     \
        ret->vt = -ERR_OVERFLOW;                        \
        return;                                         \
    }                                                   \
                                                        \
    /* check for the decimal part */                    \
    if (i < n && s[i] == '.') {                         \
        *p = i;                                         \
        ret->vt = -ERR_NUMBER_FMT;                      \
        return;                                         \
    }                                                   \
                                                        \
    /* check for the exponent part */                   \
    if (i < n && (s[i] == 'e' || s[i] == 'E')) {        \
        *p = i;                                         \
        ret->vt = -ERR_NUMBER_FMT;                      \
        return;                                         \
    }                                                   \
                                                        \
    /* update the result */                             \
    *p = i;                                             \
    ret->iv = val;

/** check whether float can represent the val exactly **/
static inline bool is_atof_exact(uint64_t man, int exp, int sgn, double *val) {
    *val = (double)man;

    if (man >> 52 != 0) {
        return false;
    }

    /* equal to if (sgn == -1) { *val *= -1; } */
    *(uint64_t *)val |= ((uint64_t)(sgn) >> 63 << 63);

    if (exp == 0 || man == 0) {
        return true;
    } else if (exp > 0 && exp <= 15+22) {
        /* uint64 integers: accurate range <= 10^15          *
         * Powers of 10: accurate range <= 10^22, as P10_TAB *
         * Example: man 1, exp 36, is ok                     */
        if (exp > 22) {
            *val *= P10_TAB[exp-22];
            exp = 22;
        }

        /* f is not accurate when too larger */
        if (*val > 1e15 || *val < -1e15) {
            return false;
        }

        *val *= P10_TAB[exp];
        return true;
    } else if (exp < 0 && exp >= -22) {
        *val /=  P10_TAB[-exp];
        return true;
    }

    return false;
}

static inline double atof_fast(uint64_t man, int exp, int sgn, int trunc, double *val) {
    double val_up = 0.0;

    /* look-up for fast atof if the conversion can be exactly */
    if (is_atof_exact(man, exp, sgn, val)) {
        return true;
    }

    /* A fast atof algorithm for high percison */
    if (atof_eisel_lemire64(man, exp, sgn, val)) {
        if (!trunc || (atof_eisel_lemire64(man+1, exp, sgn, &val_up) && val_up == *val)) {
            return true;
        }
    }

    return false;
}

static bool inline is_overflow(uint64_t man, int sgn, int exp10) {
    /* the former exp10 != 0 means man has overflowed
     * the later euqals to man*sgn < INT64_MIN or > INT64_MAX */
    return exp10 != 0 ||
        ((man >> 63) == 1 && ((uint64_t)sgn & man) != (1ull << 63));
}

void vnumber(const GoString *src, long *p, JsonState *ret) {
    int      sgn = 1;
    uint64_t man = 0; // mantissa for double (float64)
    int   man_nd = 0; // # digits of mantissa, 10 ^ 19 fits uint64_t
    int    exp10 = 0; // val = sgn * man * 10 ^ exp10
    int    trunc = 0;
    double   val = 0;

    /* initial buffer pointers */
    long         i = *p;
    size_t       n = src->len;
    const char * s = src->buf;
    char     *dbuf = ret->dbuf;
    ssize_t   dcap = ret->dcap;

    /* initialize the result, and check for EOF */
    init_ret(V_INTEGER)
    check_eof()
    check_sign(sgn = -1)

    /* check for leading zero */
    check_digit()
    check_leading_zero()

    /* parse the integer part */
    while (i < n && is_digit(s[i])) {
        add_integer_to_mantissa(man, man_nd, exp10, (s[i] - '0'))
        i++;
    }

    if (exp10 > 0) {
        trunc = 1;
    }

    /* check for decimal points */
    if (i < n && s[i] == '.') {
        i++;
        set_vt(V_DOUBLE)
        check_eof()
        check_digit()
    }

    /* skip the leading zeros of 0.000xxxx */
    if (man == 0 && exp10 == 0) {
        while (i < n && s[i] == '0') {
            i++;
            exp10--;
        }
        man = 0;
        man_nd = 0;
    }

    /* the fractional part (uint64_t mantissa can represent at most 19 digits) */
    while (i < n && man_nd < 19 && is_digit(s[i])) {
        add_float_to_mantissa(man, man_nd, exp10, (s[i] - '0'))
        i++;
    }

     /* skip the remaining digits */
    while (i < n && is_digit(s[i])) {
        trunc = 1;
        i++;
    }

    /* check for exponent */
    if (i < n && (s[i] == 'e' || s[i] == 'E')) {
        int esm = 1;
        int exp = 0;

        /* check for the '+' or '-' sign, and parse the power */
        i++;
        set_vt(V_DOUBLE)
        check_eof()
        parse_sign(esm)
        check_digit()
        while (i < n && is_digit(s[i])) {
            if (exp < 10000) {
                exp = exp * 10 + (s[i] - '0');
            }
            i++;
        }
        exp10 += exp * esm;
        goto parse_float;
    }

    if (ret->vt == V_INTEGER) {
        if (!is_overflow(man, sgn, exp10)) {
            ret->iv = (int64_t)man * sgn;

            /* following lines equal to ret->dv = (double)(man) * sgn */
            ret->dv = (double)(man);
            *(uint64_t *)&ret->dv |= ((uint64_t)(sgn) >> 63 << 63);

            *p = i;
            return;
        }
        set_vt(V_DOUBLE)
    }

parse_float:
    /* when fast algorithms failed, use slow fallback.*/
    if(!atof_fast(man, exp10, sgn, trunc, &val)) {
        val = atof_native(s + *p, i - *p, dbuf, dcap);
    }

    /* check parsed double val */
    if (is_infinity(val)) {
        ret->vt = -ERR_FLOAT_INF;
    }

    /* update the result */
    ret->dv = val;
    *p = i;
}

void vsigned(const GoString *src, long *p, JsonState *ret) {
    int64_t sgn = 1;
    vinteger(int64_t, sgn, sgn = -1)
}

void vunsigned(const GoString *src, long *p, JsonState *ret) {
    vinteger(uint64_t, 1, {
        *p = i - 1;
        ret->vt = -ERR_NUMBER_FMT;
        return;
    })
}

#undef init_ret
#undef check_eof
#undef check_digit
#undef check_leading_zero
#undef parse_sign
#undef is_digit
#undef add_integer_to_mantissa
#undef add_float_to_mantissa
#undef parse_float_digits
#undef parse_integer_digits
#undef add_digit_overflow
#undef vinteger

/** Value Skipping FSM **/

#define FSM_VAL         0
#define FSM_ARR         1
#define FSM_OBJ         2
#define FSM_KEY         3
#define FSM_ELEM        4
#define FSM_ARR_0       5
#define FSM_OBJ_0       6

#define FSM_DROP(v)     (v)->sp--
#define FSM_REPL(v, t)  (v)->vt[(v)->sp - 1] = (t)

#define FSM_CHAR(c)     do { if (ch != (c)) return -ERR_INVAL; } while (0)
#define FSM_XERR(v)     do { long r = (v); if (r < 0) return r; } while (0)

static inline void fsm_init(StateMachine *self, int vt) {
    self->sp = 1;
    self->vt[0] = vt;
}

static inline long fsm_push(StateMachine *self, int vt) {
    if (self->sp >= MAX_RECURSE) {
        return -ERR_RECURSE_MAX;
    } else {
        self->vt[self->sp++] = vt;
        return 0;
    }
}

static inline long fsm_exec(StateMachine *self, const GoString *src, long *p, uint64_t flags) {
    int  vt;
    char ch;
    long vi = -1;

    /* run until no more nested values */
    while (self->sp) {
        ch = advance_ns(src, p);
        if (ch  == 0) {
            return -ERR_EOF;
        }
        vt = self->vt[self->sp - 1];

        /* set the start address if any */
        if (vi == -1) {
            vi = *p - 1;
        }

        /* check for special types */
        switch (vt) {
            default: {
                FSM_DROP(self);
                break;
            }

            /* arrays */
            case FSM_ARR: {
                switch (ch) {
                    case ']' : FSM_DROP(self);                    continue;
                    case ',' : FSM_XERR(fsm_push(self, FSM_VAL)); continue;
                    default  : return -ERR_INVAL;
                }
            }

            /* objects */
            case FSM_OBJ: {
                switch (ch) {
                    case '}' : FSM_DROP(self);                    continue;
                    case ',' : FSM_XERR(fsm_push(self, FSM_KEY)); continue;
                    default  : return -ERR_INVAL;
                }
            }

            /* object keys */
            case FSM_KEY: {
                FSM_CHAR('"');
                FSM_REPL(self, FSM_ELEM);
                FSM_XERR(skip_string(src, p, flags));
                continue;
            }

            /* object element */
            case FSM_ELEM: {
                FSM_CHAR(':');
                FSM_REPL(self, FSM_VAL);
                continue;
            }

            /* arrays, first element */
            case FSM_ARR_0: {
                if (ch == ']') {
                    FSM_DROP(self);
                    continue;
                } else {
                    FSM_REPL(self, FSM_ARR);
                    break;
                }
            }

            /* objects, first pair */
            case FSM_OBJ_0: {
                switch (ch) {
                    default: {
                        return -ERR_INVAL;
                    }

                    /* empty object */
                    case '}': {
                        FSM_DROP(self);
                        continue;
                    }

                    /* the quote of the first key */
                    case '"': {
                        FSM_REPL(self, FSM_OBJ);
                        FSM_XERR(skip_string(src, p, flags));
                        FSM_XERR(fsm_push(self, FSM_ELEM));
                        continue;
                    }
                }
            }
        }

        /* simple values */
        switch (ch) {
            case '0' : /* fallthrough */
            case '1' : /* fallthrough */
            case '2' : /* fallthrough */
            case '3' : /* fallthrough */
            case '4' : /* fallthrough */
            case '5' : /* fallthrough */
            case '6' : /* fallthrough */
            case '7' : /* fallthrough */
            case '8' : /* fallthrough */
            case '9' : FSM_XERR(skip_positive(src, p));                     break;
            case '-' : FSM_XERR(skip_negative(src, p));                     break;
            case 'n' : FSM_XERR(advance_dword(src, p, 1, *p - 1, VS_NULL)); break;
            case 't' : FSM_XERR(advance_dword(src, p, 1, *p - 1, VS_TRUE)); break;
            case 'f' : FSM_XERR(advance_dword(src, p, 0, *p - 1, VS_ALSE)); break;
            case '[' : FSM_XERR(fsm_push(self, FSM_ARR_0));                 break;
            case '{' : FSM_XERR(fsm_push(self, FSM_OBJ_0));                 break;
            case '"' : FSM_XERR(skip_string(src, p, flags));                       break;
            case  0  : return -ERR_EOF;
            default  : return -ERR_INVAL;
        }
    }

    /* all done */
    return vi;
}

#undef FSM_DROP
#undef FSM_REPL
#undef FSM_CHAR
#undef FSM_XERR

#define check_bits(mv)                              \
    if (unlikely((v = mv & (mv - 1)) != 0)) {       \
        return -(sp - ss + __builtin_ctz(v) + 1);   \
    }

#define check_sidx(iv)          \
    if (likely(iv == -1)) {     \
        iv = sp - ss - 1;       \
    } else {                    \
        return -(sp - ss);      \
    }

#define check_vidx(iv, mv)                              \
    if (mv != 0) {                                      \
        if (likely(iv == -1)) {                         \
            iv = sp - ss + __builtin_ctz(mv);           \
        } else {                                        \
            return -(sp - ss + __builtin_ctz(mv) + 1);  \
        }                                               \
    }

static inline long do_skip_number(const char *sp, size_t nb) {
    long         di = -1;
    long         ei = -1;
    long         si = -1;
    const char * ss = sp;

    /* check for EOF */
    if (nb == 0) {
        return -1;
    }

    /* special case of '0' */
    if (*sp == '0' && (nb == 1 || (sp[1] != '.' && sp[1] != 'e' && sp[1] != 'E'))) {
        return 1;
    }

#if USE_AVX2
    /* can do with AVX-2 */
    if (likely(nb >= 32)) {
        __m256i d9 = _mm256_set1_epi8('9');
        __m256i ds = _mm256_set1_epi8('/');
        __m256i dp = _mm256_set1_epi8('.');
        __m256i el = _mm256_set1_epi8('e');
        __m256i eu = _mm256_set1_epi8('E');
        __m256i xp = _mm256_set1_epi8('+');
        __m256i xm = _mm256_set1_epi8('-');

        /* 32-byte loop */
        do {
            __m256i sb = _mm256_loadu_si256  ((const void *)sp);
            __m256i i0 = _mm256_cmpgt_epi8   (sb, ds);
            __m256i i9 = _mm256_cmpgt_epi8   (sb, d9);
            __m256i id = _mm256_cmpeq_epi8   (sb, dp);
            __m256i il = _mm256_cmpeq_epi8   (sb, el);
            __m256i iu = _mm256_cmpeq_epi8   (sb, eu);
            __m256i ip = _mm256_cmpeq_epi8   (sb, xp);
            __m256i im = _mm256_cmpeq_epi8   (sb, xm);
            __m256i iv = _mm256_andnot_si256 (i9, i0);
            __m256i ie = _mm256_or_si256     (il, iu);
            __m256i is = _mm256_or_si256     (ip, im);
            __m256i rt = _mm256_or_si256     (iv, id);
            __m256i ru = _mm256_or_si256     (ie, is);
            __m256i rv = _mm256_or_si256     (rt, ru);

            /* exponent and sign position */
            uint32_t md = _mm256_movemask_epi8(id);
            uint32_t me = _mm256_movemask_epi8(ie);
            uint32_t ms = _mm256_movemask_epi8(is);
            uint32_t mr = _mm256_movemask_epi8(rv);

            /* mismatch position */
            uint32_t v;
            uint32_t i = __builtin_ctzll(~(uint64_t)mr | 0x0100000000);

            /* mask out excess characters */
            if (i != 32) {
                md &= (1 << i) - 1;
                me &= (1 << i) - 1;
                ms &= (1 << i) - 1;
            }

            /* check & update decimal point, exponent and sign index */
            check_bits(md)
            check_bits(me)
            check_bits(ms)
            check_vidx(di, md)
            check_vidx(ei, me)
            check_vidx(si, ms)

            /* check for valid number */
            if (i != 32) {
                sp += i;
                _mm256_zeroupper();
                goto check_index;
            }

            /* move to next block */
            sp += 32;
            nb -= 32;
        } while (nb >= 32);

        /* clear the upper half to prevent AVX-SSE transition penalty */
        _mm256_zeroupper();
    }
#endif

    /* can do with SSE */
    if (likely(nb >= 16)) {
        __m128i dc = _mm_set1_epi8(':');
        __m128i ds = _mm_set1_epi8('/');
        __m128i dp = _mm_set1_epi8('.');
        __m128i el = _mm_set1_epi8('e');
        __m128i eu = _mm_set1_epi8('E');
        __m128i xp = _mm_set1_epi8('+');
        __m128i xm = _mm_set1_epi8('-');

        /* 16-byte loop */
        do {
            __m128i sb = _mm_loadu_si128 ((const void *)sp);
            __m128i i0 = _mm_cmpgt_epi8  (sb, ds);
            __m128i i9 = _mm_cmplt_epi8  (sb, dc);
            __m128i id = _mm_cmpeq_epi8  (sb, dp);
            __m128i il = _mm_cmpeq_epi8  (sb, el);
            __m128i iu = _mm_cmpeq_epi8  (sb, eu);
            __m128i ip = _mm_cmpeq_epi8  (sb, xp);
            __m128i im = _mm_cmpeq_epi8  (sb, xm);
            __m128i iv = _mm_and_si128   (i9, i0);
            __m128i ie = _mm_or_si128    (il, iu);
            __m128i is = _mm_or_si128    (ip, im);
            __m128i rt = _mm_or_si128    (iv, id);
            __m128i ru = _mm_or_si128    (ie, is);
            __m128i rv = _mm_or_si128    (rt, ru);

            /* exponent and sign position */
            uint32_t md = _mm_movemask_epi8(id);
            uint32_t me = _mm_movemask_epi8(ie);
            uint32_t ms = _mm_movemask_epi8(is);
            uint32_t mr = _mm_movemask_epi8(rv);

            /* mismatch position */
            uint32_t v;
            uint32_t i = __builtin_ctzll(~mr | 0x00010000);

            /* mask out excess characters */
            if (i != 16) {
                md &= (1 << i) - 1;
                me &= (1 << i) - 1;
                ms &= (1 << i) - 1;
            }

            /* check & update exponent and sign index */
            check_bits(md)
            check_bits(me)
            check_bits(ms)
            check_vidx(di, md)
            check_vidx(ei, me)
            check_vidx(si, ms)

            /* check for valid number */
            if (i != 16) {
                sp += i;
                goto check_index;
            }

            /* move to next block */
            sp += 16;
            nb -= 16;
        } while (nb >= 16);
    }

    /* remaining bytes, do with scalar code */
    while (likely(nb-- > 0)) {
        switch (*sp++) {
            case '0' : /* fallthrough */
            case '1' : /* fallthrough */
            case '2' : /* fallthrough */
            case '3' : /* fallthrough */
            case '4' : /* fallthrough */
            case '5' : /* fallthrough */
            case '6' : /* fallthrough */
            case '7' : /* fallthrough */
            case '8' : /* fallthrough */
            case '9' : break;
            case '.' : check_sidx(di); break;
            case 'e' : /* fallthrough */
            case 'E' : check_sidx(ei); break;
            case '+' : /* fallthrough */
            case '-' : check_sidx(si); break;
            default  : sp--; goto check_index;
        }
    }
check_index:
    if (di == 0 || si == 0 || ei == 0) {
        return -1;
    } else if (di == sp - ss - 1|| si == sp - ss - 1 || ei == sp - ss - 1) {
        return -(sp - ss);
    } else if (si > 0 && ei != si - 1) {
        return -si - 1;
    } else if (di >= 0 && ei >= 0 && di > ei - 1) {
        return -di - 1;
    } else if (di >= 0 && ei >= 0 && di == ei - 1) {
        return -ei - 1;
    } else {
        return sp - ss;
    }
}

#undef check_bits
#undef check_sidx
#undef check_vidx

long skip_array(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    fsm_init(m, FSM_ARR_0);
    return fsm_exec(m, src, p, flags);
}

long skip_object(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    fsm_init(m, FSM_OBJ_0);
    return fsm_exec(m, src, p, flags);
}

long skip_string(const GoString *src, long *p, uint64_t flags) {
    int64_t v;
    ssize_t q = *p - 1;
    ssize_t e = advance_string(src, *p, &v, flags);

    /* check for errors, and update the position */
    if (e >= 0) {
        *p = e;
        return q;
    } else {
        *p = v;
        return e;
    }
}

long skip_negative(const GoString *src, long *p) {
    long i = *p;
    long r = do_skip_number(src->buf + i, src->len - i);

    /* check for errors */
    if (r < 0) {
        *p -= r + 1;
        return -ERR_INVAL;
    }

    /* update value pointer */
    *p += r;
    return i - 1;
}

long skip_positive(const GoString *src, long *p) {
    long i = *p - 1;
    long r = do_skip_number(src->buf + i, src->len - i);

    /* check for errors */
    if (r < 0) {
        *p -= r + 2;
        return -ERR_INVAL;
    }

    /* update value pointer */
    *p += r - 1;
    return i;
}

long skip_number(const GoString *src, long *p) {
    const char* ss = src->buf;
    const char* sp = src->buf + *p;
    size_t nb = src->len;
    long i = *p;
    long r;
    bool neg = *sp == '-';

    sp += neg;
    nb -= neg;
    if (unlikely(nb <= 0)) {
        *p = sp - ss;
        return -ERR_EOF;
    }

    if (unlikely(i < nb && (*sp > '9' || *sp < '0'))) {
        *p = sp - ss;
        return -ERR_INVAL;
    }

    r = do_skip_number(sp, nb);
    if (unlikely(r < 0)) {
        *p = sp - (r + 1) - ss;
        return -ERR_INVAL;
    }
    *p = sp + r - ss;
    return i;
}

long skip_one(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    fsm_init(m, FSM_VAL);
    return fsm_exec(m, src, p, flags);
}

long validate_one(const GoString *src, long *p, StateMachine *m) {
    fsm_init(m, FSM_VAL);
    return fsm_exec(m, src, p, MASK_VALIDATE_STRING);
}

/* Faster skip api for sonic.ast */

static always_inline uint64_t get_maskx64(const char *s, char c) {
#if USE_AVX2
    __m256i v0 = _mm256_loadu_si256((__m256i const *)s);
    __m256i v1 = _mm256_loadu_si256((__m256i const *)(s + 32));
    uint32_t m0 = _mm256_movemask_epi8(_mm256_cmpeq_epi8(v0, _mm256_set1_epi8(c)));
    uint32_t m1 = _mm256_movemask_epi8(_mm256_cmpeq_epi8(v1, _mm256_set1_epi8(c))); 
    return ((uint64_t)(m1) << 32) | (uint64_t)(m0);
#else
    __m128i v0 = _mm_loadu_si128((__m128i const*)s);
    __m128i v1 = _mm_loadu_si128((__m128i const*)(s + 16));
    __m128i v2 = _mm_loadu_si128((__m128i const*)(s + 32));
    __m128i v3 = _mm_loadu_si128((__m128i const*)(s + 48));
    uint32_t m0 = _mm_movemask_epi8(_mm_cmpeq_epi8(v0, _mm_set1_epi8(c)));
    uint32_t m1 = _mm_movemask_epi8(_mm_cmpeq_epi8(v1, _mm_set1_epi8(c)));
    uint32_t m2 = _mm_movemask_epi8(_mm_cmpeq_epi8(v2, _mm_set1_epi8(c)));
    uint32_t m3 = _mm_movemask_epi8(_mm_cmpeq_epi8(v3, _mm_set1_epi8(c)));
    return ((uint64_t)(m3) << 48) | ((uint64_t)(m2)  << 32) | ((uint64_t)(m1) << 16) | (uint64_t)(m0);
#endif
}

static always_inline uint64_t get_maskx32(const char *s, char c) {
#if USE_AVX2
    __m256i v0 = _mm256_loadu_si256((__m256i const *)s);
    uint64_t m0 = (unsigned)_mm256_movemask_epi8(_mm256_cmpeq_epi8(v0, _mm256_set1_epi8(c)));
    return m0;
#else
    __m128i v0 = _mm_loadu_si128((__m128i const*)s);
    __m128i v1 = _mm_loadu_si128((__m128i const*)(s + 16));
    uint64_t m0 = (unsigned)_mm_movemask_epi8(_mm_cmpeq_epi8(v0, _mm_set1_epi8(c)));
    uint64_t m1 = (unsigned)_mm_movemask_epi8(_mm_cmpeq_epi8(v1, _mm_set1_epi8(c)));
    return m0 | (m1 << 16);
#endif
}

// get the string (besides in quote) mask
static always_inline uint64_t get_string_maskx64(const char *s, uint64_t *prev_inquote, uint64_t *prev_bs) {
    uint64_t escaped = *prev_bs;
    uint64_t quote_mask = 0, bs_mask = 0;

    /* read and get the quote or backslash bitmask */
    quote_mask = get_maskx64(s, '"');
    bs_mask = get_maskx64(s, '\\');

    /* get the escaped bitmask */
    if (bs_mask || *prev_bs) {
        bs_mask &= ~(*prev_bs);
        uint64_t follow_bs = (bs_mask << 1) | *prev_bs;
        uint64_t bs_start = bs_mask & ~follow_bs;
        uint64_t odd_start = bs_start & ODD_MASK;
        uint64_t even_or_oc = add64(odd_start, bs_mask, prev_bs);
        uint64_t even_or_escaped = (even_or_oc << 1) ^ EVEN_MASK;
        escaped = follow_bs & even_or_escaped;
    } else {
        *prev_bs = 0;
    }
    quote_mask &= ~escaped;

    /* get the inquote bitmask */
    uint64_t inquote = _mm_cvtsi128_si64(_mm_clmulepi64_si128(_mm_set_epi64x(0, quote_mask), _mm_set1_epi8('\xFF'), 0));
    inquote ^= *prev_inquote;
    *prev_inquote = (uint64_t)(((int64_t)(inquote)) >> 63);
    return inquote;
}

// get the next json structural, '}', ']' or ','。
#if USE_AVX2
static always_inline int get_structural_maskx32(const char *s) {
    __m256i v = _mm256_loadu_si256((const void *)s);
    __m256i e1 = _mm256_cmpeq_epi8(v, _mm256_set1_epi8('}'));
    __m256i e2 = _mm256_cmpeq_epi8(v, _mm256_set1_epi8(']'));
    __m256i e3 = _mm256_cmpeq_epi8(v, _mm256_set1_epi8(','));
    __m256i sv = _mm256_or_si256(_mm256_or_si256(e1, e2), e3);
    return _mm256_movemask_epi8(sv);
}
#endif

static always_inline int get_structural_maskx16(const char *s) {
    __m128i v = _mm_loadu_si128((const void *)s);
    __m128i e1 = _mm_cmpeq_epi8(v, _mm_set1_epi8('}'));
    __m128i e2 = _mm_cmpeq_epi8(v, _mm_set1_epi8(']'));
    __m128i e3 = _mm_cmpeq_epi8(v, _mm_set1_epi8(','));
    __m128i sv = _mm_or_si128(_mm_or_si128(e1, e2), e3);
    return _mm_movemask_epi8(sv);
}

// skip the number at the next '}', ']' or ',' or the ending of json.
static always_inline long skip_number_fast(const GoString *src, long *p) {
    size_t nb = src->len - *p;
    const char *s = src->buf + *p;
    long vi = *p - 1;
    int m = 0;

#if USE_AVX2
    while (likely(nb >= 32)) {
        if ((m = get_structural_maskx32(s))) {
            *p = s - src->buf + __builtin_ctzll(m);
            return vi;
        }
        s += 32, nb -= 32;
    }
#endif

    while (likely(nb >= 16)) {
        if ((m = get_structural_maskx16(s))) {
            *p = s - src->buf + __builtin_ctzll(m);
            return vi;
        }
        s += 16, nb -= 16;
    }

    while (likely(nb > 0)) {
        if (*s == '}' || *s == ']' || *s == ',') {
            *p = s - src->buf;
            return vi;
        }
        s++, nb--;
    }
    *p = s - src->buf;
    return vi;
}

static always_inline long skip_container_fast(const GoString *src, long *p, char lc, char rc) {
    long nb = src->len - *p;
    const char *s = src->buf + *p;
    long vi = *p - 1;

    uint64_t prev_inquote = 0, prev_bs = 0;
    uint64_t lbrace = 0, rbrace = 0;
    size_t lnum = 0, rnum = 0, last_lnum = 0;
    uint64_t inquote = 0;

    while (likely(nb >= 64)) {
skip:
        inquote = get_string_maskx64(s, &prev_inquote, &prev_bs);
        lbrace = get_maskx64(s, lc) & ~inquote;
        rbrace = get_maskx64(s, rc) & ~inquote;

        /* traverse each right brace */
        last_lnum = lnum;
        while (rbrace > 0) {
            uint64_t lbrace_first = (rbrace - 1) & lbrace;
            lnum = last_lnum + __builtin_popcountll((int64_t)lbrace_first);
            bool is_closed = lnum <= rnum;
            if (is_closed) {
                *p = src->len - nb + __builtin_ctzll(rbrace) + 1;
                // *p is out-of-bound access here
                if (*p > src->len) {
                    *p = src->len;
                    return -ERR_EOF;
                }
                return vi;
            }
            rbrace &= (rbrace - 1); // clear the lowest right brace
            rnum ++;
        }
        lnum = last_lnum + __builtin_popcountll((int64_t)lbrace);
        s += 64, nb -= 64;
    }

    if (nb <= 0) {
        *p = src->len;
        return -ERR_EOF;
    }

    char tbuf[64] = {0};
    bool cross_page = vec_cross_page(s, 64);
    if (cross_page) {
        memcpy_p64(tbuf, s, nb);
        s = tbuf;
    }
    goto skip;
}

static always_inline long skip_object_fast(const GoString *src, long *p) {
    return skip_container_fast(src, p, '{', '}');
}

static always_inline long skip_array_fast(const GoString *src, long *p) {
    return skip_container_fast(src, p, '[', ']');
}

static always_inline long skip_string_fast(const GoString *src, long *p) {
    const char* s = src->buf + *p;
    size_t nb = src->len - *p;
    long vi = *p - 1;
    uint64_t prev_bs = 0, escaped;

    while (likely(nb >= 32)) {
        uint32_t quote = get_maskx32(s, '"');
        uint32_t bs_mask = get_maskx32(s, '\\');
        if (bs_mask || prev_bs) {
            bs_mask &= ~prev_bs;
            uint64_t follow_bs = (bs_mask << 1) | prev_bs;
            uint64_t bs_start = bs_mask & ~follow_bs;
            uint64_t odd_start = bs_start & ODD_MASK;
            uint64_t even_or_oc = add32(odd_start, bs_mask, &prev_bs);
            uint64_t even_or_escaped = (even_or_oc << 1) ^ EVEN_MASK;
            escaped = follow_bs & even_or_escaped;
            quote &= ~escaped;
        }
        if (quote) {
            *p = s + __builtin_ctzll(quote) + 1 - src->buf;
            return vi;
        }
        s += 32;
        nb -= 32;
    }

    if (unlikely(prev_bs != 0)) {
        if (nb == 0) return -ERR_EOF;
        s++, nb--;
    }

    while (likely(nb > 0)) {
        if (*s == '\\') {
            s += 2, nb -= 2;
            continue;
        }
        if (*s == '"') {
            *p = s - src->buf + 1;
            return vi;
        }
        s++, nb--;
    }
    return -ERR_EOF;
}

long skip_one_fast(const GoString *src, long *p) {
    char c = advance_ns(src, p);
    /* set the start address */
    long vi = *p - 1;
    switch (c) {
        case '[': return skip_array_fast(src, p);
        case '{': return skip_object_fast(src, p);
        case '"': return skip_string_fast(src, p);
        case '-': case '0' ... '9': return skip_number_fast(src, p);
        case 't': case 'n': { if (*p + 3 < src->len) { *p += 3; } else { return -ERR_EOF; } }; break;
        case 'f': { if (*p + 4 < src->len) { *p += 4; } else { return -ERR_EOF; } }; break;
        case  0 : return -ERR_EOF;
        default : *p -= 1; return -ERR_INVAL; // backward error position
    }
    return vi;
}


static always_inline GoKind kind(const GoIface* iface) {
    return (iface->type->kind_flags) &  GO_KIND_MASK;
}

static always_inline bool is_int(const GoIface* iface) {
    return kind(iface) == Int;
}

static always_inline bool is_str(const GoIface* iface) {
    return kind(iface) == String;
}

static always_inline GoString get_str(const GoIface* iface) {
    return *(GoString*)(iface->value);
}

static always_inline int64_t get_int(const GoIface* iface) {
    return *(int64_t*)(iface->value);
}

// xmemcmpeq return true if s1 and s2 is equal for the n bytes, otherwise, return false.
static always_inline bool xmemcmpeq(const char * s1, const char * s2, size_t n) {
    bool c1, c2;
#if USE_AVX2
    while (n >= 32) {
        __m256i  v1   = _mm256_loadu_si256((const void *)s1);
        __m256i  v2   = _mm256_loadu_si256((const void *)s2);
        uint32_t mask = ~((uint32_t)_mm256_movemask_epi8(_mm256_cmpeq_epi8(v1, v2)));
        if (mask) return false;
        s1 += 32;
        s2 += 32;
        n  -= 32;
    };
    c1 = vec_cross_page(s1, 32);
    c2 = vec_cross_page(s2, 32);
    // not cross page
    if (!c1 && !c2) {
        __m256i  v1   = _mm256_loadu_si256((const void *)s1);
        __m256i  v2   = _mm256_loadu_si256((const void *)s2);
        uint32_t mask = ~((uint32_t)_mm256_movemask_epi8(_mm256_cmpeq_epi8(v1, v2)));
        bool eq = (mask == 0) || (__builtin_ctzll(mask) >= n);
        return eq;
    }
#endif
    while (n >= 16) {
        __m128i  v1   = _mm_loadu_si128((const void *)s1);
        __m128i  v2   = _mm_loadu_si128((const void *)s2);
        uint16_t mask = ~((uint16_t)_mm_movemask_epi8(_mm_cmpeq_epi8(v1, v2)));
        if (mask != 0) return false;
        s1 += 16;
        s2 += 16;
        n  -= 16;
    };
    c1 = vec_cross_page(s1, 16);
    c2 = vec_cross_page(s2, 16);
    // not cross page
    if (!c1 && !c2) {
        __m128i  v1   = _mm_loadu_si128((const void *)s1);
        __m128i  v2   = _mm_loadu_si128((const void *)s2);
        uint16_t mask = ~((uint16_t)_mm_movemask_epi8(_mm_cmpeq_epi8(v1, v2)));
        bool eq = (mask == 0) || (__builtin_ctzll(mask) >= n);
        return eq;
    }
    // cross page
    while (n > 0 && *s1++ == *s2++) n--;
    return n == 0;
}

// match_key return negative if errors, zero if not matched, one if matched.
static always_inline long match_key(const GoString *src, long *p, const GoString key) {
    static const long not_match = 0;
    int64_t v = -1;
    long si = *p;
    long se = advance_string_default(src, *p, &v);
    if (unlikely(se < 0)) {
        *p = src->len;
        return -ERR_EOF;
    }

    /* update position */
    *p = se;

    /* compare non-escaped strings */
    if (likely(v == -1 || v > se)) {
        long sn = se - si - 1;
        return sn == key.len && xmemcmpeq(src->buf + si, key.buf, key.len);
    }

    /* deal with escaped strings */
    char buf[8] = {0}; // escaped buffer
    const char* sp = src->buf + si;
    const char* end = src->buf + se - 1;
    const char* kp = key.buf;
    const char* ke = key.buf + key.len;
    while (sp < end && kp < ke) {
        if (*sp == '\\') {
            long en = unescape(&sp, end, buf);
            if (en < 0) {
                *p = sp - src->buf;
                return en;
            }
            const char* ee = buf + en;
            const char* ep = buf;
            while (kp < ke && ep < ee && *kp == *ep) kp++, ep++;
            if (ep != ee) {
                return not_match;
            }
        } else if (*sp == *kp) {
            sp++, kp++; 
        } else {
            return not_match;
        }
    };
    return sp == end && kp == ke;
}

long get_by_path(const GoString *src, long *p, const GoSlice *path) {
    GoIface *ps = (GoIface*)(path->buf);
    GoIface *pe = (GoIface*)(path->buf) + path->len;
    char c = 0;
    int64_t index;
    long found;

query:
    /* if empty path, skip the whole json */
    if (ps == pe) {
        return skip_one_fast(src, p);
    }
    /* match type: should query key in object, query index in array */
    c = advance_ns(src, p);
    if (is_str(ps)) {
        if (c != '{') goto err_inval;
        goto skip_in_obj;
    } else if (is_int(ps)) {
        if (c != '[') goto err_inval;
        goto skip_in_arr;
    } else {
        goto err_inval;
    }

skip_in_obj:
    c = advance_ns(src, p);
    if (c != '"') goto err_inval;
    found = match_key(src, p, get_str(ps));
    if (found < 0) return found; // parse string errors

    /* value should after : */
    c = advance_ns(src, p);
    if (c != ':') goto err_inval;
    if (found) {
        ps++;
        goto query;
    } else {
        skip_one_fast(src, p);
        c = advance_ns(src, p);
        if (c != ',') goto err_inval; // not found key
        goto skip_in_obj;
    }

skip_in_arr:
    index = get_int(ps);
    /* skip array elem one by one */
    while (index-- > 0) {
        skip_one_fast(src, p);
        c = advance_ns(src, p);
        if (c != ',') goto err_inval; // out of range
    }
    ps++;
    goto query;

err_inval:
    *p -= 1; // backward error position
    return -ERR_INVAL;
}

// 
long validate_utf8(const GoString *src, long *p, StateMachine *m) {
    xassert(*p >= 0 && src->len > *p);
    return validate_utf8_with_errors(src->buf, src->len, p, m);
}

// validate_utf8_fast returns zero if valid, otherwise, the error position.
long validate_utf8_fast(const GoString *s) {
#if USE_AVX2
    /* fast path for valid utf8 */
    if (validate_utf8_avx2(s) == 0) {
        return 0;
    }
#endif
    return validate_utf8_errors(s);
}