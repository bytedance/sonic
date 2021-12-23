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
#include "ryu_tab.h"

/* use 128-bit type for performance */
typedef __uint128_t uint128_t;

/* Returns e == 0 ? 1 : ceil(log_2(5^e)) */
static inline int32_t pow5bits(const int32_t e) {
    return (int32_t) (((((uint32_t) e) * 1217359) >> 19) + 1);
}

/* Returns floor(log_10(2^e)) */
static inline uint32_t log10pow2(const int32_t e) {
    return (((uint32_t) e) * 78913) >> 18;
}

/* Returns floor(log_10(5^e)) */
static inline uint32_t log10pow5(const int32_t e) {
    return (((uint32_t) e) * 732923) >> 20;
}

static inline uint32_t pow5factor(uint64_t v) {
    uint64_t m_inv5 = 14757395258967641293u; // *5 = 1(mod 2^64)
    uint64_t n_div5 = 3689348814741910323u;  // =2^64 / 5
    uint32_t cnt = 0;
    for (;;) {
        v *= m_inv5;
        if (v > n_div5)
            break;
        ++cnt;
    }
    return cnt;
}

/* Returns true if value is divisible by 5^p */
static inline bool ispow5(const uint64_t v, const uint32_t p) {
    return pow5factor(v) >= p;
}

/* Returns true if value is divisible by 2^p */
static inline bool ispow2(const uint64_t v, const uint32_t p) {
    return (v & ((1ull << p) - 1)) == 0;
}

/* Requires 0 <= v < v < 100000000000000000L */
static inline uint32_t ctz10(const uint64_t v) {
    if (v <                10) return 1;
    if (v <               100) return 2;
    if (v <              1000) return 3;
    if (v <             10000) return 4;
    if (v <            100000) return 5;
    if (v <           1000000) return 6;
    if (v <          10000000) return 7;
    if (v <         100000000) return 8;
    if (v <        1000000000) return 9;
    if (v <       10000000000) return 10;
    if (v <      100000000000) return 11;
    if (v <     1000000000000) return 12;
    if (v <    10000000000000) return 13;
    if (v <   100000000000000) return 14;
    if (v <  1000000000000000) return 15;
    if (v < 10000000000000000) return 16;
                               return 17;
}

/* Best case: use 128-bit type */
static inline uint64_t mulshift(const uint64_t m, const uint64_t *mul, const int32_t j) {
    uint128_t lo = ((uint128_t) m) * mul[0];
    uint128_t hi = ((uint128_t) m) * mul[1];
    return (uint64_t) (((lo >> 64) + hi) >> (j - 64));
}

#define mul_shift_all(m, mul, j, shift)       \
    vp = mulshift(4 * m + 2, mul, j);         \
    vm = mulshift(4 * m - 1 - shift, mul, j); \
    vr = mulshift(4 * m, mul, j);

#define copy_two_digs(dst, src) \
    *(dst) = *(src);            \
    *(dst+1) = *(src+1);

/* A floating decimal representing man * 10^exp */
typedef struct f64_d {
    uint64_t man;
    int32_t exp;
} f64_d;

static inline f64_d f64tod(const uint64_t man,const uint32_t exp) {
    int32_t e2;
    uint64_t m2;
    if (exp == 0) {
        /* subtract 2 so that the bounds computation has 2 additional bits */
        e2 = 1 - 1023 - 52 - 2;
        m2 = man;
    } else {
        e2 = (int32_t) exp - 1023 - 52 - 2;
        m2 = (1ull << 52) | man;
    }
    bool even = (m2 & 1) == 0;

    /* Step 2: Determine the interval of valid decimal representations */
    uint64_t mv = 4 * m2;
    /* Implicit bool -> int conversion. True is 1, false is 0 */
    uint32_t shift = man != 0 || exp <= 1;

    /* Step 3: Convert to a decimal power base using 128-bit arithmetic */
    uint64_t vr, vp, vm;
    int32_t e10;
    bool vmzeros = false;
    bool vrzeros = false;
    if (e2 >= 0) {

        uint32_t q = log10pow2(e2) - (e2 > 3); // max(0, log10pow2(e2) - 1)
        e10 = (int32_t) q;
        int32_t k = DOUBLE_POW5_INV_BITCOUNT + pow5bits((int32_t) q) - 1;
        int32_t i = -e2 + (int32_t) q + k;
        uint64_t *mul = (uint64_t*)DOUBLE_POW5_INV_SPLIT[q];
        /* {vm, vr, vp} * 10^e10 = {mm, mv, mp} * 2^e2
         * mp = 4 * m2 + 2
         * mm = mv - 1 - shift
         */
        mul_shift_all(m2, mul, i, shift)

        if (q <= 21) {
            if (mv % 5 == 0) {
                vrzeros = ispow5(mv, q);
            } else if (even) {
                /* Same as min(e2 + (~mm & 1), pow5Factor(mm)) >= q
                 * <=> e2 + (~mm & 1) >= q && pow5Factor(mm) >= q
                 * <=> true && pow5Factor(mm) >= q, since e2 >= q.
                 */
                vmzeros = ispow5(mv - 1 - shift, q);
            } else {
                /* Same as min(e2 + 1, pow5Factor(mp)) >= q. */
                vp -= ispow5(mv + 2, q);
            }
        }
    } else {
        uint32_t q = log10pow5(-e2) - (-e2 > 1); // max(0, log10pow5(-e2) - 1)
        e10 = (int32_t) q + e2;
        int32_t i = -e2 - (int32_t) q;
        int32_t k = pow5bits(i) - DOUBLE_POW5_BITCOUNT;
        int32_t j = (int32_t) q - k;
        uint64_t *mul = (uint64_t*)DOUBLE_POW5_SPLIT[i];
        /* {vm, vr, vp} * 10^e10 = {mm, mv, mp} * 2^e2 */
        mul_shift_all(m2, mul, j, shift)

        if (q <= 1) {
            /* {vr,vp,vm} is trailing zeros if {mv,mp,mm} has at least q trailing
             * 0 bits. mv = 4 * m2, so it always has at least two trailing 0 bits.
             */
            vrzeros = true;
            if (even) {
                /* mm = mv - 1 - shift, so it has 1 trailing 0 bit if shift = 1 */
                vmzeros = shift == 1;
            } else {
                /* mp = mv + 2, so it always has at least one trailing 0 bit */
                --vp;
            }
        } else if (q < 63) {
            vrzeros = ispow2(mv, q);
        }
    }

    /* Step 4: Find the shortest decimal representation in the interval */
    int32_t removed = 0;
    uint8_t lastrmdig = 0;
    uint64_t dman;
    /* On average, we remove ~2 DIGs */
    if (vmzeros || vrzeros) {
        /* General case, which happens rarely (~0.7%) */
        for (;;) {
            uint64_t vpdiv10 = vp / 10;
            uint64_t vmdiv10 = vm / 10;
            if (vpdiv10 <= vmdiv10) {
                break;
            }
            uint32_t vmmod10 = ((uint32_t) vm) - 10 * ((uint32_t) vmdiv10);
            uint64_t vrdiv10 = vr / 10;
            uint32_t vrmod10 = ((uint32_t) vr) - 10 * ((uint32_t) vrdiv10);
            vmzeros &= vmmod10 == 0;
            vrzeros &= lastrmdig == 0;
            lastrmdig = (uint8_t) vrmod10;
            vr = vrdiv10;
            vp = vpdiv10;
            vm = vmdiv10;
            ++removed;
        }
        if (vmzeros) {
            for (;;) {
                uint64_t vmdiv10 = vm / 10;
                uint32_t vmmod10 = ((uint32_t) vm) - 10 * ((uint32_t) vmdiv10);
                if (vmmod10 != 0) {
                    break;
                }
                uint64_t vpdiv10 = vp / 10;
                uint64_t vrdiv10 = vr / 10;
                uint32_t vrmod10 = ((uint32_t) vr) - 10 * ((uint32_t) vrdiv10);
                vrzeros &= lastrmdig == 0;
                lastrmdig = (uint8_t) vrmod10;
                vr = vrdiv10;
                vp = vpdiv10;
                vm = vmdiv10;
                ++removed;
            }
        }
        if (vrzeros && lastrmdig == 5 && vr % 2 == 0) {
            /* Round even if the exact number is .....50..0 */
            lastrmdig = 4;
        }
        /* We need to take vr + 1 if vr is outside bounds or we need to round up */
        dman = vr + ((vr == vm && (!even || !vmzeros)) || lastrmdig >= 5);
    } else {
        /* Specialized for the common case (~99.3%). Percentages below are relative to this */
        bool roundup= false;
        uint64_t vpdiv100 = vp / 100;
        uint64_t vmdiv100 = vm / 100;
        if (vpdiv100 > vmdiv100) { // Optimization: remove two DIGs at a time (~86.2%).
            uint64_t vrdiv100 = vr / 100;
            uint32_t vrmod100 = ((uint32_t) vr) - 100 * ((uint32_t) vrdiv100);
            roundup = vrmod100 >= 50;
            vr = vrdiv100;
            vp = vpdiv100;
            vm = vmdiv100;
            removed += 2;
        }
        /* Loop iterations below (approximately), without optimization above:
         * 0: 0.03%, 1: 13.8%, 2: 70.6%, 3: 14.0%, 4: 1.40%, 5: 0.14%, 6+: 0.02%
         * Loop iterations below (approximately), with optimization above:
         * 0: 70.6%, 1: 27.8%, 2: 1.40%, 3: 0.14%, 4+: 0.02%
         */
        for (;;) {
            uint64_t vpdiv10 = vp / 10;
            uint64_t vmdiv10 = vm / 10;
            if (vpdiv10 <= vmdiv10) {
                break;
            }
            uint64_t vrdiv10 = vr / 10;
            uint32_t vrmod10 = ((uint32_t) vr) - 10 * ((uint32_t) vrdiv10);
            roundup = vrmod10 >= 5;
            vr = vrdiv10;
            vp = vpdiv10;
            vm = vmdiv10;
            ++removed;
        }
        /* We need to take vr + 1 if vr is outside bounds or we need to round up */
        dman = vr + (vr == vm || roundup);
    }
    f64_d fd = {
        .exp = e10 + removed,
        .man = dman,
    };
    return fd;
}

/* Print the decimal DIGs from mantissa */
static always_inline void print_mantissa(uint64_t man, char *out, int mlen) {
    /* We have at most 17 DIGs, and uint32_t can store 9 DIGs.
     * If man doesn't fit into uint32_t, we cut off 8 DIGs,
     * so the rest will fit into uint32_t.
     */
    char *r = out + mlen;
    if (man < 10) {}
    if ((man >> 32) != 0) {
        /* Expensive 64-bit division */
        uint64_t q = man / 100000000;
        uint32_t man2 = ((uint32_t) man) - 100000000 * ((uint32_t) q);
        man = q;

        uint32_t c  = man2 % 10000;
        man2 /= 10000;
        uint32_t d  = man2 % 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        uint32_t d0 = (d % 100) << 1;
        uint32_t d1 = (d / 100) << 1;
        copy_two_digs(r - 2, DIG_TAB + c0)
        copy_two_digs(r - 4, DIG_TAB + c1)
        copy_two_digs(r - 6, DIG_TAB + d0)
        copy_two_digs(r - 8, DIG_TAB + d1)
        r -= 8;
    }
    uint32_t man2 = (uint32_t) man;
    while (man2 >= 10000) {
#ifdef __clang__ // https://bugs.llvm.org/show_bug.cgi?id=38217
        uint32_t c = man2 - 10000 * (man2 / 10000);
#else
        uint32_t c = man2 % 10000;
#endif
        man2 /= 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        copy_two_digs(r - 2, DIG_TAB + c0)
        copy_two_digs(r - 4, DIG_TAB + c1)
        r -= 4;
    }
    if (man2 >= 100) {
        uint32_t c = (man2 % 100) << 1;
        man2 /= 100;
        copy_two_digs(r - 2, DIG_TAB + c)
        r -= 2;
    }
    if (man2 >= 10) {
        uint32_t c = man2 << 1;
        copy_two_digs(r - 2, DIG_TAB + c)
    } else {
        *out = (char) ('0' + man2);
    }
}

static inline int print_exponent(f64_d v, char *out, int mlen) {
    int idx = 0;

    print_mantissa(v.man, out + idx + 1, mlen);

    /* Print decimal point if needed */
    out[idx] = out[idx + 1];
    if (mlen > 1) {
        out[idx + 1] = '.';
        idx += mlen + 1;
    } else {
        ++idx;
    }

    /* Print the exponent */
    out[idx++] = 'e';
    int32_t exp = v.exp + (int32_t) mlen - 1;
    if (exp < 0) {
        out[idx++] = '-';
        exp = -exp;
    }

    if (exp >= 100) {
        int32_t c = exp % 10;
        copy_two_digs(out + idx, DIG_TAB + 2 * (exp / 10))
        out[idx + 2] = (char) ('0' + c);
        idx += 3;
    } else if (exp >= 10) {
        copy_two_digs(out + idx, DIG_TAB + 2 * exp)
        idx += 2;
    } else {
        out[idx++] = (char) ('0' + exp);
    }

    return idx;
}

static inline int print_decimal(const f64_d v, char *out, int mlen) {
    int idx    = 0;
    int lzeros = 0;
    int rzeros = 0;
    int point  = 0;
    int exp10  = mlen - 1 + v.exp;

    /* parse the point idx and additional zeros */
    if (exp10 < 0) {
        lzeros = -exp10;
        point  = 1;
    } else if (exp10 < mlen - 1) {
        point  = 1 + exp10;
    } else {
        rzeros = exp10 - mlen + 1;
    }

    int i = 0;
    /* add left zeros */
    if (lzeros) {
        out[idx++] = '0';
        out[idx++] = '.';
        point = 0;
    }
    for (i = 1; i < lzeros; ++i) {
        out[idx++] = '0';
    }

    /* add the mantissa DIGs */
    print_mantissa(v.man, out + idx, mlen);
    if (point) {
        for (i = idx + mlen; i > idx + point; --i) {
            out[i] = out[i-1];
        }
        out[idx + point] = '.';
        idx += 1;
    }

    /* add right zeros */
    idx += mlen;
    for (i = 0; i < rzeros; ++i) {
        out[idx++] = '0';
    }

    return idx;
}

static inline bool f64tod_exct_int(const uint64_t man, const uint32_t exp,
  f64_d* v) {
    uint64_t m2 = (1ull << 52) | man; // implicit 1
    int32_t e2 = (int32_t) exp - 1023 - 52;

    if (e2 > 0 || e2 < -52) {
        return false;
    }

    uint64_t mask = (1ull << -e2) - 1;
    if ((m2 & mask) != 0) { // with fraction
        return false;
    }

    v->man = m2 >> -e2;
    v->exp = 0;
    return true;
}

static int inline ryu(uint64_t bits, char *out) {
    /* Step 1: Decode the floating-point number */
    uint64_t man = bits & ((1ull << 52) - 1);
    uint32_t exp = (uint32_t) ((bits >> 52) & ((1u << 11) - 1));

    /* Skip when Infinity */
    if (exp == ((1u << 11) - 1u)) {
        return 0;
    }

    f64_d v;
    /* for integer from [1, 2^53], can be resepensated exactly by double */
    bool is_exact_int = f64tod_exct_int(man, exp, &v);
    if (!is_exact_int){  // find the shortest decimal representation
        v = f64tod(man, exp);
    }

    /* Step 5: Print the decimal representation */
    int idx = 0;
    uint32_t mlen = ctz10(v.man);
    int exp10 = mlen - 1 + v.exp;
    /* The format as Go encoding/json package */
    bool isexp = exp10 < -6 || exp10 >= 21;

    if (isexp) // exponent format
        idx += print_exponent(v, out + idx, mlen);
    else      // decimal format
        idx += print_decimal(v, out + idx, mlen);

    return idx;
}

int f64toa(char *out, double val) {
    int   i = 0;
    char *p = out;
    uint64_t uval = *(uint64_t *)&val;

    /* negative numbers */
    if (unlikely(uval >> 63) == 1) {
        i    = 1;
        uval &= ((1ull << 63) - 1);
        *p++ = '-';
    }

    /* simple case of 0.0 */
    if (uval ==  0) {
        *p = '0';
        return i + 1;
    }

    /* print the number with Ryu algorithm */
    int n = ryu(uval, p);
    return n + i;
}
