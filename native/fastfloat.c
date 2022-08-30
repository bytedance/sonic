/* Copyright 2020 Alexander Bolz
 * 
 * Boost Software License - Version 1.0 - August 17th, 2003
 * 
 * Permission is hereby granted, free of charge, to any person or organization
 * obtaining a copy of the software and accompanying documentation covered by
 * this license (the "Software") to use, reproduce, display, distribute,
 * execute, and transmit the Software, and to prepare derivative works of the
 * Software, and to permit third-parties to whom the Software is furnished to
 * do so, all subject to the following:
 * 
 * The copyright notices in the Software and this entire statement, including
 * the above license grant, this restriction and the following disclaimer,
 * must be included in all copies of the Software, in whole or in part, and
 * all derivative works of the Software, unless such copies or derivative
 * works are solely in the form of machine-executable object code generated by
 * a source language processor.
 *
 * Unless required by applicable law or agreed to in writing, this software
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.
 * 
 * This file may have been modified by ByteDance authors. All ByteDance
 * Modifications are Copyright 2022 ByteDance Authors.
 */

#include "native.h"
#include "tab.h"
#include "test/xassert.h"

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

static inline unsigned ctz10(const uint64_t v) {
    xassert(0 <= v && v < 100000000000000000ull);
    if (v >= 10000000000ull) {
        if (v <      100000000000ull) return 11;
        if (v <     1000000000000ull) return 12;
        if (v <    10000000000000ull) return 13;
        if (v <   100000000000000ull) return 14;
        if (v <  1000000000000000ull) return 15;
        if (v < 10000000000000000ull) return 16;
                                      return 17;
    }
        if (v <                10ull) return 1;
        if (v <               100ull) return 2;
        if (v <              1000ull) return 3;
        if (v <             10000ull) return 4;
        if (v <            100000ull) return 5;
        if (v <           1000000ull) return 6;
        if (v <          10000000ull) return 7;
        if (v <         100000000ull) return 8;
        if (v <        1000000000ull) return 9;
                                      return 10;

}

static inline char* format_significand(uint64_t sig, char *out, int cnt) {
    char *p = out + cnt;
    int ctz = 0;

    if ((sig >> 32) != 0) {
        uint64_t q = sig / 100000000;
        uint32_t r = ((uint32_t)sig) - 100000000 * ((uint32_t) q);
        sig = q;
        if (r != 0) {
            uint32_t c  = r % 10000;
            r /= 10000;
            uint32_t d  = r % 10000;
            uint32_t c0 = (c % 100) << 1;
            uint32_t c1 = (c / 100) << 1;
            uint32_t d0 = (d % 100) << 1;
            uint32_t d1 = (d / 100) << 1;
            copy_two_digs(p - 2, Digits + c0);
            copy_two_digs(p - 4, Digits + c1);
            copy_two_digs(p - 6, Digits + d0);
            copy_two_digs(p - 8, Digits + d1);
        } else {
            ctz += 8;
        }
        p -= 8;
    }

    uint32_t sig2 = (uint32_t)sig;
    while (sig2 >= 10000) {
        uint32_t c = sig2 - 10000 * (sig2 / 10000);
        sig2 /= 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        copy_two_digs(p - 2, Digits + c0);
        copy_two_digs(p - 4, Digits + c1);
        p -= 4;
    }
    if (sig2 >= 100) {
        uint32_t c = (sig2 % 100) << 1;
        sig2 /= 100;
        copy_two_digs(p - 2, Digits + c);
        p -= 2;
    }
    if (sig2 >= 10) {
        uint32_t c = sig2 << 1;
        copy_two_digs(p - 2, Digits + c);
    } else {
        *out = (char) ('0' + sig2);
    }
    return out + cnt - ctz;
}

static inline char* format_integer(uint64_t sig, char *out, unsigned cnt) {
    char *p = out + cnt;
    if ((sig >> 32) != 0) {
        uint64_t q = sig / 100000000;
        uint32_t r = ((uint32_t)sig) - 100000000 * ((uint32_t) q);
        sig = q;
        uint32_t c  = r % 10000;
        r /= 10000;
        uint32_t d  = r % 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        uint32_t d0 = (d % 100) << 1;
        uint32_t d1 = (d / 100) << 1;
        copy_two_digs(p - 2, Digits + c0);
        copy_two_digs(p - 4, Digits + c1);
        copy_two_digs(p - 6, Digits + d0);
        copy_two_digs(p - 8, Digits + d1);
        p -= 8;
    }

    uint32_t sig2 = (uint32_t)sig;
    while (sig2 >= 10000) {
        uint32_t c = sig2 - 10000 * (sig2 / 10000);
        sig2 /= 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        copy_two_digs(p - 2, Digits + c0);
        copy_two_digs(p - 4, Digits + c1);
        p -= 4;
    }
    if (sig2 >= 100) {
        uint32_t c = (sig2 % 100) << 1;
        sig2 /= 100;
        copy_two_digs(p - 2, Digits + c);
        p -= 2;
    }
    if (sig2 >= 10) {
        uint32_t c = sig2 << 1;
        copy_two_digs(p - 2, Digits + c);
    } else {
        *out = (char) ('0' + sig2);
    }
    return out + cnt;
}

static inline char* format_exponent(f64_dec v, char *out, unsigned cnt) {
    char* p = out + 1;
    char* end = format_significand(v.sig, p, cnt);
    while (*(end - 1) == '0') end--;

    /* print decimal point if needed */
    *out = *p;
    if (end - p > 1) {
        *p = '.';
    } else {
        end--;
    }

    /* print the exponent */
    *end++ = 'e';
    int32_t exp = v.exp + (int32_t) cnt - 1;
    if (exp < 0) {
        *end++ = '-';
        exp = -exp;
    } else {
        *end++ = '+';
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

static inline char* format_decimal(f64_dec v, char* out, unsigned cnt) {
    char* p = out;
    char* end;
    int point = cnt + v.exp;

    /* print leading zeros if fp < 1 */
    if (point <= 0) {
        *p++ = '0', *p++ = '.';
        for (int i = 0; i < -point; i++) {
            *p++ = '0';
        }
    }

    /* add the remaining digits */
    end = format_significand(v.sig, p, cnt);
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
    int cnt = ctz10(dec.sig);
    int dot = cnt + dec.exp;
    int sci_exp = dot - 1;
    bool exp_fmt = sci_exp < -6 || sci_exp > 20;
    bool has_dot = dot < cnt;

    if (exp_fmt) {
        return format_exponent(dec, p, cnt);
    }
    if (has_dot) {
        return format_decimal(dec, p, cnt);
    }

    char* end = p + dot;
    p = format_integer(dec.sig, p, cnt);
    while (p < end) *p++ = '0';
    return end;
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
static inline f64_dec f64todec(uint64_t rsig, int32_t rexp, uint64_t c, int32_t q) {
    uint64_t cbl, cb, cbr, vbl, vb, vbr, lower, upper, s;
    int32_t k, h;
    bool even, irregular, w_inside, u_inside;
    f64_dec dec;

    even = !(c & 1);
    irregular = rsig == 0 && rexp > 1;

    cbl = 4 * c - 2 + irregular;
    cb = 4 * c;
    cbr = 4 * c + 2;

    /*  q is in [-1500, 1500]
        k = irregular ? floor(log_10(3/4 * 2^q)) : floor(log10(pow(2^q)))
        floor(log10(3/4 * 2^q))  = (q * 1262611 - 524031) >> 22
        floor(log10(pow(2^q))) = (q * 1262611) >> 22 */
    k = (q * 1262611 - (irregular ? 524031 : 0)) >> 22;

    /*  k is in [-1233, 1233]
        s = floor(V) = floor(c * 2^q * 10^-k)
        vb = 4V = 4 * c * 2^q * 10^-k */
    h = q + ((-k) * 1741647 >> 19) + 1;
    uint64x2 pow10 = pow10_ceil_sig(-k);
    vbl = round_odd(pow10, cbl << h);
    vb = round_odd(pow10, cb << h);
    vbr = round_odd(pow10, cbr << h);

    lower = vbl + !even;
    upper = vbr - !even;

    s = vb / 4;
    if (s >= 10) {
        /* R_k+1 interval contains at most one: up or wp */
        uint64_t sp = s / 10;
        bool up_inside = lower <= (40 * sp);
        bool wp_inside = (40 * sp + 40) <= upper;
        if (up_inside != wp_inside) {
            dec.sig = sp + wp_inside;
            dec.exp = k + 1;
            return dec;
        }
    }

    /* R_k interval contains at least one: u or w */
    u_inside = lower <= (4 * s);
    w_inside = (4 * s + 4) <= upper;
    if (u_inside != w_inside) {
        dec.sig = s + w_inside;
        dec.exp = k;
        return dec;
    }

    /* R_k interval contains both u or w */
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
    uint64_t rsig, c;
    int32_t rexp, q;

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

    /* fp = c * 2^q */
    if (likely(rexp != 0)) {
        /* double is normal */
        c = rsig | F64_HIDDEN_BIT;
        q = rexp - F64_EXP_BIAS - F64_SIG_BITS;

        /* fast path for integer */
        if (q <= 0 && q >= -F64_SIG_BITS && is_div_pow2(c, -q)) {
            uint64_t u = c >> -q;
            p = format_integer(u, p, ctz10(u));
            return p - out;
        }

    } else {
        c = rsig;
        q = 1 - F64_EXP_BIAS - F64_SIG_BITS;
    }

    f64_dec dec = f64todec(rsig, rexp, c, q);
    p = write_dec(dec, p);
    return p - out;
}

#undef F64_BITS       
#undef F64_EXP_BITS   
#undef F64_SIG_BITS   
#undef F64_EXP_MASK   
#undef F64_SIG_MASK   
#undef F64_EXP_BIAS   
#undef F64_INF_NAN_EXP
#undef F64_HIDDEN_BIT 