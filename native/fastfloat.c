/* Copyright 2018 Ulf Adams.
 * Modifications copyright 2021 ByteDance Inc.
 *
 * The contents of this file may be used under the terms of the Apache License,
 * Version 2.0.
 *
 *    (See accompanying file LICENSE-Apache or copy at
 *     http: *www.apache.org/licenses/LICENSE-2.0)
 *
 * Alternatively, the contents of this file may be used under the terms of
 * the Boost Software License, Version 1.0.
 *    (See accompanying file LICENSE-Boost or copy at
 *     https: *www.boost.org/LICENSE_1_0.txt)
 *
 * Unless required by applicable law or agreed to in writing, this software
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.
 */
#include "native.h"
#include "tab.h"
#include "test/xassert.h"

#define xprintf(...)
// #define xassert(...)

#define F64_BITS         64
#define F64_EXP_BITS     11
#define F64_SIG_BITS     52
#define F64_EXP_MASK     0x7FF0000000000000ull // middle 11 bits
#define F64_SIG_MASK     0x000FFFFFFFFFFFFFull // lower 52 bits
#define F64_EXP_BIAS     1023
#define F64_INF_NAN_EXP  0x7FF
#define F64_HIDDEN_BIT   0x0010000000000000ull

struct f64_dec {
    uint64_t sig;
    int64_t exp;
};
typedef struct f64_dec f64_dec;

typedef __uint128_t uint128_t;

static inline uint32_t ctz10(const uint64_t v) {
    xassert(0 <= v && v < 100000000000000000ull);
    if (v >= 10000000000) {
        if (v <      100000000000) return 11;
        if (v <     1000000000000) return 12;
        if (v <    10000000000000) return 13;
        if (v <   100000000000000) return 14;
        if (v <  1000000000000000) return 15;
        if (v < 10000000000000000) return 16;
                                   return 17;
    }
    if (v <                10) return 1;
    if (v <               100) return 2;
    if (v <              1000) return 3;
    if (v <             10000) return 4;
    if (v <            100000) return 5;
    if (v <           1000000) return 6;
    if (v <          10000000) return 7;
    if (v <         100000000) return 8;
    if (v <        1000000000) return 9;
                               return 10;

}

bool is_div_pow2(uint64_t val, int32_t e) {
    xassert(e >= 0 && e <= 63);
    uint64_t mask = (1ull << e) - 1;
    return (val & mask) == 0;
}

static inline char* utoa2(char* p, uint32_t val) {
    p[0] = Digits[val];
    p[1] = Digits[val + 1];
    return p + 2;
}

static inline void copy_two_digs(char* dst, const char* src) {
    *(dst) = *(src);
    *(dst + 1) = *(src + 1);
}

static inline char* print_mantissa(uint64_t man, char *out, int mlen) {
    char *r = out + mlen;
    int ctz = 0;
    if (man < 10) {}
    if ((man >> 32) != 0) {
        /* Expensive 64-bit division */
        uint64_t q = man / 100000000;
        uint32_t man2 = ((uint32_t) man) - 100000000 * ((uint32_t) q);
        man = q;
        if (man2 != 0) {
            uint32_t c  = man2 % 10000;
            man2 /= 10000;
            uint32_t d  = man2 % 10000;
            uint32_t c0 = (c % 100) << 1;
            uint32_t c1 = (c / 100) << 1;
            uint32_t d0 = (d % 100) << 1;
            uint32_t d1 = (d / 100) << 1;
            copy_two_digs(r - 2, Digits + c0);
            copy_two_digs(r - 4, Digits + c1);
            copy_two_digs(r - 6, Digits + d0);
            copy_two_digs(r - 8, Digits + d1);
        } else {
            ctz += 8;
        }
        r -= 8;
    }
    uint32_t man2 = (uint32_t) man;
    while (man2 >= 10000) {
        // c = man2 % 10000;
        uint32_t c = man2 - 10000 * (man2 / 10000);
        man2 /= 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        copy_two_digs(r - 2, Digits + c0);
        copy_two_digs(r - 4, Digits + c1);
        r -= 4;
    }
    if (man2 >= 100) {
        uint32_t c = (man2 % 100) << 1;
        man2 /= 100;
        copy_two_digs(r - 2, Digits + c);
        r -= 2;
    }
    if (man2 >= 10) {
        uint32_t c = man2 << 1;
        copy_two_digs(r - 2, Digits + c);
    } else {
        *out = (char) ('0' + man2);
    }
    return out + mlen - ctz;
}

static inline char* print_exponent(f64_dec v, char *out, int mlen) {
    char* p = out + 1;
    char* end = print_mantissa(v.sig, p, mlen);
    while (*(end - 1) == '0') end--;

    /* Print decimal point if needed */
    *out = *p;
    if (end - p > 1) {
        *p = '.';
    } else {
        end--;
    }

    /* Print the exponent */
    *end++ = 'e';
    int32_t exp = v.exp + (int32_t) mlen - 1;
    if (exp < 0) {
        *end++ = '-';
        exp = -exp;
    }

    if (exp >= 100) {
        int32_t c = exp % 10;
        copy_two_digs(end, Digits + 2 * (exp / 10));
        end[2] = (char) ('0' + c);
        end += 3;
    } else if (exp >= 10) {
        copy_two_digs(end, Digits + 2 * exp);
        end += 2;
    } else {
        *end++ = (char) ('0' + exp);
    }
    return end;
}

static inline char* print_decimal(f64_dec v, char* out, int mlen) {
    char* p = out;
    char* end;
    int exp10  = mlen - 1 + v.exp;
    int point = mlen + v.exp;

    /* print leading zeros if fp < 1 */
    if (point <= 0) {
        *p++ = '0', *p++ = '.';
        for (int i = 0; i < -point; i++) {
            *p++ = '0';
        }
    }

    /* add the remaining digits */
    end = print_mantissa(v.sig, p, mlen);
    while (*(end - 1) == '0') end--;
    if (point <= 0) {
        return end;
    }

    /* insert point or add trailing zeros */
    int digs = end - p;
    if (digs > point) {
        for (int i = 0; i < digs - point; i++) {
            *(end - i) = *(end - i - 1);
        }
        p[point] = '.';
        end++;
    } else {
        for (int i = 0; i < point - digs; i++) {
            *end++ = '0';
        }
    }
    return end;
}

static inline char* write_dec(f64_dec dec, char* p) {
    int32_t count = ctz10(dec.sig);
    int32_t dot = count + dec.exp;
    int32_t sci_exp = dot - 1;
    bool exp_fmt = sci_exp < -6 || sci_exp > 20;

    if (exp_fmt) {
        return print_exponent(dec, p, count);
    }
    return print_decimal(dec, p, count);
}

static inline uint64_t f64toraw(double fp) {
    union {
        uint64_t u64;
        double   f64;
    } uval;
    uval.f64 = fp;
    return uval.u64;
}

static inline uint64_t round_odd(uint64x2 g, uint64_t cp) {
    const uint128_t x = ((uint128_t)cp) * g.lo;
    const uint128_t y = ((uint128_t)cp) * g.hi + ((uint64_t)(x >> 64));

    const uint64_t y0 = ((uint64_t)y);
    const uint64_t y1 = ((uint64_t)(y >> 64));
    return y1 | (y0 > 1);
}

/**
 Rendering float point number into decimal.
 The function used Schubfach algorithm, reference:
 The Schubfach way to render doubles, Raffaello Giulietti, 2022-03-20.
 https://drive.google.com/file/d/1gp5xv4CAa78SVgCeWfGqqI4FfYYYuNFb
 https://mail.openjdk.java.net/pipermail/core-libs-dev/2021-November/083536.html
 https://github.com/openjdk/jdk/pull/3402 (Java implementation)
 https://github.com/abolz/Drachennest (C++ implementation)
 */
static inline f64_dec f64todec(uint64_t rsig, int32_t rexp) {
    uint64_t c, cbl, cb, cbr, vbl, vb, vbr, lower, upper, s, sp;
    int32_t q, k, h;
    bool even, irregular, w_inside, u_inside;
    f64_dec dec;

    if (likely(rexp != 0)) {
        /* double is normal */
        c = rsig | F64_HIDDEN_BIT;
        q = rexp - F64_EXP_BIAS - F64_SIG_BITS;

        /* fast path for integer */
        if (q <= 0 && q >= -F64_SIG_BITS && is_div_pow2(c, -q)) {
            dec.sig = c >> -q;
            dec.exp = 0;
            return dec;
        }
    } else {
        c = rsig;
        q = 1 - F64_EXP_BIAS - F64_SIG_BITS;
    }

    even = !(c & 1);
    irregular = rsig == 0 && rexp > 1;

    cbl = 4 * c - 2 + irregular;
    cb = 4 * c;
    cbr = 4 * c + 2;

    // q: [-1500, 1500]
    // k = irregular ? floor(log_10(3/4 2^q)) : floor(log10(pow(2^q)))
    // floor(log10(3/4 2^q)) = (q * 1262611 - 524031) >> 22
    // floor(log10(pow(2^q))) = (q * 1262611) >> 22
    k = (q * 1262611 - (irregular ? 524031 : 0)) >> 22;

    // k: [-1233, 1233]
    // s = floor(V) = floor(c*2^q*10^-k), s is the significand of decimal
    // vb = 4V = 4*c(2^q)*(10^-k)
    // let (g-1)*2^r <= 10^-k < g*2^r, then:
    // vb = cb*(2^q)*(g*2^r) = (cb*g)*2^(q + floor(log2(10^-k)) - 127)
    // let h = 128 + q + floor(log2(10^-k)) - 127, h in [1, 4], then:
    // vb = g*(cb<<h)*2^(-128), becomes the 128-bit multiplications.
    h = q + ((-k) * 1741647 >> 19) + 1;
    uint64x2 pow10 = pow10_ceil_sig(-k);
    vbl = round_odd(pow10, cbl << h);
    vb = round_odd(pow10, cb << h);
    vbr = round_odd(pow10, cbr << h);

    // R_v interval:
    // if c is even: [vl, vr], then (vl, vr)
    lower = vbl + !even;
    upper = vbr - !even;

    s = vb / 4;
    if (s >= 10) {
        // R_k+1 interval contains at most one: up or wp
        uint64_t sp = s / 10;
        bool up_inside = lower <= (40 * sp);
        bool wp_inside = (40 * sp + 40) <= upper;
        if (up_inside != wp_inside) {
            dec.sig = sp + wp_inside;
            dec.exp = k + 1;
            return dec;
        }
    }

    // R_k interval contains at least one: u or w
    u_inside = lower <= (4 * s);
    w_inside = (4 * s + 4) <= upper;
    if (u_inside != w_inside) {
        dec.sig = s + w_inside;
        dec.exp = k;
        return dec;
    }

    // R_k interval contains both u or w
    // let t = s + 1, and then: 
    // vb > mid => t is closer
    // vb = 2(s + t) = 4*s + 2 && s is odd => should round to t
    uint64_t mid = 4 * s + 2;
    bool round_up = vb > mid || (vb == mid && (s & 1) != 0);
    dec.sig = s + round_up;
    dec.exp = k;
    return dec;
}

int f64toa(char *out, double fp) {
    char* p = out;
    uint64_t raw = f64toraw(fp);
    bool neg;
    uint64_t rsig, bsig, dsig;
    int32_t rexp, bexp, dexp;

    neg = ((raw >> (F64_BITS - 1)) != 0);
    rsig = raw & F64_SIG_MASK;
    rexp = (int32_t)((raw & F64_EXP_MASK) >> F64_SIG_BITS);

    /* check infinity and nan */
    if (unlikely(rexp == F64_INF_NAN_EXP)) {
        return 0;
    }

    /* check negative numbers */
    *p = '-';
    p += neg;

    /* simple case of 0.0 */
    if ((raw << 1) ==  0) {
        *p++ = '0';
        return p - out;
    }

    /* use xxx algorithm */
    f64_dec dec = f64todec(rsig, rexp);
    p = write_dec(dec, p);
    return p - out;
}