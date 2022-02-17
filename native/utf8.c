
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

//      code point    | first byte | second byte | thrid byte | fourth byte
//  U+0000  -  U+007F | 0___ ____
//  U+0080  -  U+07FF | 110_ ____  | 10__ ____
//  U+0800  -  U+FFFF | 1110 ____  | 10__ ____   | 10__ ____
// U+10000 - U+10FFFF | 1111 0___  | 10__ ____   | 10__ ____  | 10__ ____
// checks non-ascii characters, and returns the utf-8 length
static inline ssize_t is_utf8(const uint8_t* noascii, size_t n) {
#define is_next(b) (((b) >> 6) == 2)
    uint8_t b0 = *noascii;
    uint8_t b1 = *(noascii + 1);
    uint8_t b2 = *(noascii + 2);
    uint8_t b3 = *(noascii + 3);
    uint32_t r;

    /* 2-byte */
    if (unlikely(n < 2 || !is_next(b1))) {
        return 0;
    }
    if (unlikely(b0 < 0xe0 && b0 >= 0xc0)) {
        r = ((uint32_t)(b0 & 0x1f) << 6) | (b1 & 0x3f);
        if (likely(r >= 0x0080 && r <= 0x07ff)) {
            return 2;
        }
        return 0;
    }

    /* 3-byte */
    if (unlikely(n < 3 || !is_next(b2))) {
        return 0;
    }
    if (likely(b0 < 0xf0)) {
        r = ((uint32_t)(b0 & 0x0f) << 12) | ((uint32_t)(b1 & 0x3f) << 6) | (b2 & 0x3f);
        if (likely(r >= 0x0800u && r <= 0xffffu)) {
            return 3;
        }
        return 0;
    }

    /* 4-byte */
    if (unlikely(n < 4 || !is_next(b3))) {
        return 0;
    }
    if (likely(b0 < 0xf8)) {
        r = ((uint32_t)(b0 & 0x07) << 18) | ((uint32_t)(b1 & 0x3f) << 12) | ((uint32_t)(b2 & 0x3f) << 6) | (b3 & 0x3f);
        if (likely(r >= 0x10000u && r <= 0x10ffffu)) {
            return 4;
        }
    }
    return 0;
#undef is_next
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
        if (unlikely((b = is_utf8(p, nb)) == 0)) {
            return p - s;
        }

        nb -= b;
        rb -= b;
        p  += b;
    }

    return -1;
}