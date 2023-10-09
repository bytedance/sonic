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

/* decimical shift witout overflow, e.g. 9 << 61 overflow */
#define MAX_SHIFT 60

/* Decimal represent the integer or float
 * example 1: 1.1   {"11", 2, 1, 0}
 * example 2: -0.1  {"1", 1, 0, 1}
 * example 3: 999   {"999", 3, 3, 0}
 */
typedef struct Decimal {
    char*  d;
    size_t cap;
    int    nd;
    int    dp;
    int    neg;
    int    trunc;
} Decimal;

/* decimal power of ten to binary power of two.
 * For example: POW_TAB[1]: 10 ** 1 ~ 2 ** 3
 */
static const int POW_TAB[9] = {1, 3, 6, 9, 13, 16, 19, 23, 26};

/* Left shift information for decimal.
 * For example, {2, "625"}.  That means that it will add 2 digits to the new decimal
 * when the prefix of decimal is from "625" to "999", and 1 digit from "0" to "624".
 */
typedef struct lshift_cheat  {
    int   delta;                             // number of added digits when left shift
    const char  cutoff[100];                 // minus one digit if under the half(cutoff).
} lshift_cheat;

/* Look up for the decimal shift information by binary shift bits.
 * idx is shift bits for binary.
 * value is the shift information for decimal.
 * For example, idx is 4, the value is {2, "625"}.
 * That means the binary shift 4 bits left, will cause add 2 digits to the decimal
 * if the prefix of decimal is under "625".
 */
const static lshift_cheat LSHIFT_TAB[61];

static inline void decimal_init(Decimal *d, char *dbuf, size_t cap) {
    d->d = dbuf;
    d->cap = cap;
    for (int i = 0; i < d->cap; ++i) {
        d->d[i] = 0;
    }
    d->dp    = 0;
    d->nd    = 0;
    d->neg   = 0;
    d->trunc = 0;
}

static inline void decimal_set(Decimal *d, const char *s, ssize_t len, char *dbuf, ssize_t cap) {
    int i = 0;

    decimal_init(d, dbuf, cap);
    if (s[i] == '-') {
        i++;
        d->neg = 1;
    }

    int saw_dot = 0;
    for (; i < len; i++) {
        if ('0' <= s[i] && s[i] <= '9') {
            if (s[i] == '0' && d->nd == 0) { // ignore leading zeros
                d->dp--;
                continue;
            }
            if (d->nd < d->cap) {
                d->d[d->nd] = s[i];
                d->nd++;
            } else if (s[i] != '0') {
                /* truncat the remaining digits */
                d->trunc = 1;
            }
        } else if (s[i] == '.') {
            saw_dot = 1;
            d->dp = d->nd;
        } else {
            break;
        }
    }

    /* integer */
    if (saw_dot == 0) {
        d->dp = d->nd;
    }

    /* exponent */
    if (i < len && (s[i] == 'e' || s[i] == 'E')) {
        int exp = 0;
        int esgn = 1;

        i++;
        if (s[i] == '+') {
            i++;
        } else if (s[i] == '-') {
            i++;
            esgn = -1;
        }

        for (; i < len && ('0' <= s[i] && s[i] <= '9') && exp < 10000; i++) {
                exp = exp * 10 + (s[i] - '0');
        }
        d->dp += exp * esgn;
    }

    return;
}

/* trim trailing zeros from number */
static inline void trim(Decimal *d) {
    while (d->nd > 0 && d->d[d->nd - 1] == '0') {
        d->nd--;
    }
    if (d->nd == 0) {
        d->dp = 0;
    }
}

/* Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow */
static inline void right_shift(Decimal *d, uint32_t k) {
    int      r = 0; // read pointer
    int      w = 0; // write pointer
    uint64_t n = 0;

    /* Pick up enough leading digits to cover first shift */
    for (; n >> k == 0; r++) {
        if (r >= d->nd) {
            if (n == 0) {
                d->nd = 0; // no digits for this num
                return;
            }
            /* until n has enough bits for right shift */
            while (n >> k == 0) {
                n *= 10;
                r++;
            }
            break;
        }
        n = n * 10 + d->d[r] - '0'; // read the value from d.d
    }
    d->dp -= r - 1; // point shift left

    uint64_t mask = (1ull << k) - 1;
    uint64_t dig = 0;

    /* Pick up a digit, put down a digit */
    for (; r < d->nd; r++) {
        dig = n >> k;
        n &= mask;
        d->d[w++] = (char)(dig + '0');
        n = n * 10 + d->d[r] - '0';
    }

    /* Put down extra digits */
    while (n > 0) {
        dig = n >> k;
        n &= mask;
        if (w < d->cap) {
            d->d[w] = (char)(dig + '0');
            w++;
        } else if (dig > 0) {
            /* truncated */
            d->trunc = 1;
        }
        n *= 10;
    }

    d->nd = w;
    trim(d);
}

/* Compare the leading prefix, if b is lexicographically less, return 0 */
static inline bool prefix_is_less(const char *b, const char *s, uint64_t bn) {
    int i = 0;
    for (; i < bn; i++) {
        if (s[i] == '\0') {
            return false;
        }
        if (b[i] != s[i]) {
            return b[i] < s[i];
        }
    }
    return s[i] != '\0';
}

/* Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow */
static inline void left_shift(Decimal *d, uint32_t k) {
    int delta = LSHIFT_TAB[k].delta;

    if (prefix_is_less(d->d, LSHIFT_TAB[k].cutoff, d->nd)){
        delta--;
    }

    int r = d->nd;         // read index
    int w = d->nd + delta; // write index
    uint64_t n = 0;
    uint64_t quo = 0;
    uint64_t rem = 0;

    /* Pick up a digit, put down a digit */
    for (r--; r >= 0; r--) {
        n += (uint64_t)(d->d[r] - '0') << k;
        quo = n / 10;
        rem = n - 10 * quo;
        w--;
        if (w < d->cap) {
            d->d[w] = (char)(rem + '0');
        } else if (rem != 0) {
            /* truncated */
            d->trunc = 1;
        }
        n = quo;
    }

    /* Put down extra digits */
    while (n > 0) {
        quo = n / 10;
        rem = n - 10 * quo;
        w--;
        if (w < d->cap) {
            d->d[w] = (char)(rem + '0');
        } else if (rem != 0) {
            /* truncated */
            d->trunc = 1;
        }
        n = quo;
    }

    d->nd += delta;
    if (d->nd >= d->cap) {
        d->nd = d->cap;
    }
    d->dp += delta;
    trim(d);
}

static inline void decimal_shift(Decimal *d, int k) {
    if (d->nd == 0 || k == 0) {
        return;
    }

    if (k > 0) {
        while (k > MAX_SHIFT) {
            left_shift(d, MAX_SHIFT);
            k -= MAX_SHIFT;
        }
        if (k) {
            left_shift(d, k);
        }
    }

    if (k < 0) {
        while (k < -MAX_SHIFT) {
            right_shift(d, MAX_SHIFT);
            k += MAX_SHIFT;
        }
        if (k) {
            right_shift(d, -k);
        }
    }

}

static inline int should_roundup(Decimal *d, int nd) {
    if (nd < 0 || nd >= d->nd) {
        return 0;
    }

    /* Exactly halfway - round to even */
    if (d->d[nd] == '5' && nd+1 == d->nd) {
        if (d->trunc) {
            return 1;
        }
        return nd > 0 && (d->d[nd-1]-'0')%2 != 0;
    }

    /* not halfway - round to the nearest */
    return d->d[nd] >= '5';
}

/* Extract integer part, rounded appropriately */
static inline uint64_t rounded_integer(Decimal *d) {
    if (d->dp > 20) { // overflow
        return 0xFFFFFFFFFFFFFFFF; //64 bits
    }

    int i = 0;
    uint64_t n = 0;
    for (i = 0; i < d->dp && i < d->nd; i++) {
        n = n * 10 + (d->d[i] - '0');
    }
    for (; i < d->dp; i++) {
        n *= 10;
    }
    if (should_roundup(d, d->dp)) {
        n++;
    }
    return n;
}

int decimal_to_f64(Decimal *d, double *val) {
    int exp2 = 0;
    uint64_t mant = 0;
    uint64_t bits = 0;

    /* d is zero */
    if (d->nd == 0) {
        mant = 0;
        exp2 = -1023;
        goto out;
    }

    /* Overflow, return inf/INF */
    if (d->dp > 310) {
        goto overflow;
    }
    /* Underflow, return zero */
    if (d->dp < -330) {
        mant = 0;
        exp2 = -1023;
        goto out;
    }

    /* Scale by powers of two until in range [0.5, 1.0) */
    int n = 0;
    while (d->dp > 0) { // d >= 1
        if (d->dp >= 9) {
            n = 27;
        } else {
            n = POW_TAB[d->dp];
        }
        decimal_shift(d, -n); // shift right
        exp2 += n;
    }
    while ((d->dp < 0) || ((d->dp == 0) && (d->d[0] < '5'))) { // d < 0.5
        if (-d->dp >= 9) {
            n = 27;
        } else {
            n = POW_TAB[-d->dp];
        }
        decimal_shift(d, n); // shift left
        exp2 -= n;
    }

    /* Our range is [0.5,1) but floating point range is [1,2) */
    exp2 --;

    /* Minimum exp2 for doulbe is -1022.
     * If the exponent is smaller, move it up and
     * adjust d accordingly.
     */
    if (exp2 < -1022) {
        n = -1022 - exp2;
        decimal_shift(d, -n); // shift right
        exp2 += n;
    }

    /* Exp2 too large */
    if ((exp2 + 1023) >= 0x7FF) {
        goto overflow;
    }

    /* Extract 53 bits. */
    decimal_shift(d, 53);  // shift left
    mant = rounded_integer(d);

    /* Rounding might have added a bit; shift down. */
    if (mant == (2ull << 52)) { // mant has 54 bits
        mant >>= 1;
        exp2 ++;
        if ((exp2 + 1023) >= 0x7FF) {
            goto overflow;
        }
    }

    /* Denormalized? */
    if ((mant & (1ull << 52)) == 0) {
        exp2 = -1023;
    }
    goto out;

overflow:
    /* ±INF/inf */
    mant = 0;
    exp2 = 0x7FF - 1023;

out:
    /* Assemble bits. */
    bits = mant & 0x000FFFFFFFFFFFFF;
    bits |= (uint64_t)((exp2 + 1023) & 0x7FF) << 52;
    if (d->neg) {
        bits |= 1ull << 63;
    }
    *(uint64_t*)val = bits;
    return 0;
}

double atof_native(const char *sp, ssize_t nb, char* dbuf, ssize_t cap) {
    Decimal d;
    double val = 0;
    decimal_set(&d, sp, nb, dbuf, cap);
    decimal_to_f64(&d, &val);
    return val;
}

#undef MAX_SHIFT

const static lshift_cheat LSHIFT_TAB[61] = {
    // Leading digits of 1/2^i = 5^i.
    // 5^23 is not an exact 64-bit floating point number,
    // so have to use bc for the math.
    // Go up to 60 to be large enough for 32bit and 64bit platforms.
    /*
        seq 60 | sed 's/^/5^/' | bc |
        awk 'BEGIN{ print "\t{ 0, \"\" }," }
        {
            log2 = log(2)/log(10)
            printf("\t{ %d, \"%s\" },\t// * %d\n",
                int(log2*NR+1), $0, 2**NR)
        }'
    */
    {0, ""},
    {1, "5"},                                           // * 2
    {1, "25"},                                          // * 4
    {1, "125"},                                         // * 8
    {2, "625"},                                         // * 16
    {2, "3125"},                                        // * 32
    {2, "15625"},                                       // * 64
    {3, "78125"},                                       // * 128
    {3, "390625"},                                      // * 256
    {3, "1953125"},                                     // * 512
    {4, "9765625"},                                     // * 1024
    {4, "48828125"},                                    // * 2048
    {4, "244140625"},                                   // * 4096
    {4, "1220703125"},                                  // * 8192
    {5, "6103515625"},                                  // * 16384
    {5, "30517578125"},                                 // * 32768
    {5, "152587890625"},                                // * 65536
    {6, "762939453125"},                                // * 131072
    {6, "3814697265625"},                               // * 262144
    {6, "19073486328125"},                              // * 524288
    {7, "95367431640625"},                              // * 1048576
    {7, "476837158203125"},                             // * 2097152
    {7, "2384185791015625"},                            // * 4194304
    {7, "11920928955078125"},                           // * 8388608
    {8, "59604644775390625"},                           // * 16777216
    {8, "298023223876953125"},                          // * 33554432
    {8, "1490116119384765625"},                         // * 67108864
    {9, "7450580596923828125"},                         // * 134217728
    {9, "37252902984619140625"},                        // * 268435456
    {9, "186264514923095703125"},                       // * 536870912
    {10, "931322574615478515625"},                      // * 1073741824
    {10, "4656612873077392578125"},                     // * 2147483648
    {10, "23283064365386962890625"},                    // * 4294967296
    {10, "116415321826934814453125"},                   // * 8589934592
    {11, "582076609134674072265625"},                   // * 17179869184
    {11, "2910383045673370361328125"},                  // * 34359738368
    {11, "14551915228366851806640625"},                 // * 68719476736
    {12, "72759576141834259033203125"},                 // * 137438953472
    {12, "363797880709171295166015625"},                // * 274877906944
    {12, "1818989403545856475830078125"},               // * 549755813888
    {13, "9094947017729282379150390625"},               // * 1099511627776
    {13, "45474735088646411895751953125"},              // * 2199023255552
    {13, "227373675443232059478759765625"},             // * 4398046511104
    {13, "1136868377216160297393798828125"},            // * 8796093022208
    {14, "5684341886080801486968994140625"},            // * 17592186044416
    {14, "28421709430404007434844970703125"},           // * 35184372088832
    {14, "142108547152020037174224853515625"},          // * 70368744177664
    {15, "710542735760100185871124267578125"},          // * 140737488355328
    {15, "3552713678800500929355621337890625"},         // * 281474976710656
    {15, "17763568394002504646778106689453125"},        // * 562949953421312
    {16, "88817841970012523233890533447265625"},        // * 1125899906842624
    {16, "444089209850062616169452667236328125"},       // * 2251799813685248
    {16, "2220446049250313080847263336181640625"},      // * 4503599627370496
    {16, "11102230246251565404236316680908203125"},     // * 9007199254740992
    {17, "55511151231257827021181583404541015625"},     // * 18014398509481984
    {17, "277555756156289135105907917022705078125"},    // * 36028797018963968
    {17, "1387778780781445675529539585113525390625"},   // * 72057594037927936
    {18, "6938893903907228377647697925567626953125"},   // * 144115188075855872
    {18, "34694469519536141888238489627838134765625"},  // * 288230376151711744
    {18, "173472347597680709441192448139190673828125"}, // * 576460752303423488
    {19, "867361737988403547205962240695953369140625"}, // * 1152921504606846976
};
