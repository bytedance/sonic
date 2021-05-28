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

/** This is the Florian's Grisu2 Algorithm implemented in C.
 *  See https://legacy.cs.indiana.edu/~dyb/pubs/FP-Printing-PLDI96.pdf for more info.
 */

#include "native.h"

struct f64_t {
    int32_t  e;
    uint64_t f;
};

#define FP_SSIZE    64
#define DP_SSIZE    52
#define DP_1_LG10   0.30102999566398114     // = 1 / log2(10)

#define F64_EBIAS   1075
#define F64_EXMIN   -F64_EBIAS

#define F64_HBIT    0x0010000000000000
#define F64_EMASK   0x7ff0000000000000
#define F64_SMASK   0x000fffffffffffff

static const int16_t TabPowE[87] = {
    -1220, -1193, -1166, -1140, -1113, -1087, -1060, -1034, -1007,  -980,
     -954,  -927,  -901,  -874,  -847,  -821,  -794,  -768,  -741,  -715,
     -688,  -661,  -635,  -608,  -582,  -555,  -529,  -502,  -475,  -449,
     -422,  -396,  -369,  -343,  -316,  -289,  -263,  -236,  -210,  -183,
     -157,  -130,  -103,   -77,   -50,   -24,     3,    30,    56,    83,
      109,   136,   162,   189,   216,   242,   269,   295,   322,   348,
      375,   402,   428,   455,   481,   508,   534,   561,   588,   614,
      641,   667,   694,   720,   747,   774,   800,   827,   853,   880,
      907,   933,   960,   986,  1013,  1039,  1066
};

static const uint64_t TabPowF[87] = {
    0xfa8fd5a0081c0288, 0xbaaee17fa23ebf76,
    0x8b16fb203055ac76, 0xcf42894a5dce35ea,
    0x9a6bb0aa55653b2d, 0xe61acf033d1a45df,
    0xab70fe17c79ac6ca, 0xff77b1fcbebcdc4f,
    0xbe5691ef416bd60c, 0x8dd01fad907ffc3c,
    0xd3515c2831559a83, 0x9d71ac8fada6c9b5,
    0xea9c227723ee8bcb, 0xaecc49914078536d,
    0x823c12795db6ce57, 0xc21094364dfb5637,
    0x9096ea6f3848984f, 0xd77485cb25823ac7,
    0xa086cfcd97bf97f4, 0xef340a98172aace5,
    0xb23867fb2a35b28e, 0x84c8d4dfd2c63f3b,
    0xc5dd44271ad3cdba, 0x936b9fcebb25c996,
    0xdbac6c247d62a584, 0xa3ab66580d5fdaf6,
    0xf3e2f893dec3f126, 0xb5b5ada8aaff80b8,
    0x87625f056c7c4a8b, 0xc9bcff6034c13053,
    0x964e858c91ba2655, 0xdff9772470297ebd,
    0xa6dfbd9fb8e5b88f, 0xf8a95fcf88747d94,
    0xb94470938fa89bcf, 0x8a08f0f8bf0f156b,
    0xcdb02555653131b6, 0x993fe2c6d07b7fac,
    0xe45c10c42a2b3b06, 0xaa242499697392d3,
    0xfd87b5f28300ca0e, 0xbce5086492111aeb,
    0x8cbccc096f5088cc, 0xd1b71758e219652c,
    0x9c40000000000000, 0xe8d4a51000000000,
    0xad78ebc5ac620000, 0x813f3978f8940984,
    0xc097ce7bc90715b3, 0x8f7e32ce7bea5c70,
    0xd5d238a4abe98068, 0x9f4f2726179a2245,
    0xed63a231d4c4fb27, 0xb0de65388cc8ada8,
    0x83c7088e1aab65db, 0xc45d1df942711d9a,
    0x924d692ca61be758, 0xda01ee641a708dea,
    0xa26da3999aef774a, 0xf209787bb47d6b85,
    0xb454e4a179dd1877, 0x865b86925b9bc5c2,
    0xc83553c5c8965d3d, 0x952ab45cfa97a0b3,
    0xde469fbd99a05fe3, 0xa59bc234db398c25,
    0xf6c69a72a3989f5c, 0xb7dcbf5354e9bece,
    0x88fcf317f22241e2, 0xcc20ce9bd35c78a5,
    0x98165af37b2153df, 0xe2a0b5dc971f303a,
    0xa8d9d1535ce3b396, 0xfb9b7cd9a4a7443c,
    0xbb764c4ca7a44410, 0x8bab8eefb6409c1a,
    0xd01fef10a657842c, 0x9b10a4e5e9913129,
    0xe7109bfba19c0c9d, 0xac2820d9623bf429,
    0x80444b5e7aa7cf85, 0xbf21e44003acdd2d,
    0x8e679c2f5e44ff8f, 0xd433179d9c8cb841,
    0x9e19db92b4e31ba9, 0xeb96bf6ebadf77d9,
    0xaf87023b9bf0ee6b
};

static const uint64_t TabPow10[10] = {
    1,
    10,
    100,
    1000,
    10000,
    100000,
    1000000,
    10000000,
    100000000,
    1000000000
};

/** FP-64 Helper **/

static inline void f64_set(struct f64_t *r, double v) {
    uint64_t bv = *(uint64_t *)&v;
    uint64_t sv = bv & F64_SMASK;
    int32_t  ev = (bv & F64_EMASK) >> DP_SSIZE;

    /* check for denormalized values */
    if (ev == 0) {
        r->f = sv;
        r->e = F64_EXMIN + 1;
    } else {
        r->f = sv + F64_HBIT;
        r->e = ev - F64_EBIAS;
    }
}

static inline void f64_raw(struct f64_t *r, uint64_t f, int32_t e) {
    r->e = e;
    r->f = f;
}

static inline void f64_sub(struct f64_t *r, const struct f64_t *a, const struct f64_t *b) {
    r->e = a->e;
    r->f = a->f - b->f;
}

static inline void f64_mul(struct f64_t *r, const struct f64_t *a, const struct f64_t *b) {
    __int128_t v0 = a->f;
    __int128_t v1 = b->f;
    __int128_t v2 = v0 * v1;
    uint64_t   vh = v2 >> 64;
    uint64_t   vl = (uint64_t)v2;

    /* rounding */
    if (vl & (1ull << 63)) {
        vh++;
    }

    /* save the result */
    r->f = vh;
    r->e = a->e + b->e + 64;
}

static inline void f64_norm(struct f64_t *r, const struct f64_t *v) {
    uint64_t f = v->f;
    uint32_t s = __builtin_clzll(f);

    /* remove the leading zeros, and adjust the exponent */
    r->f = f << s;
    r->e = v->e - s;
}

static inline void f64_normb(struct f64_t *m, struct f64_t *p, const struct f64_t *v) {
    int32_t  dv = v->f != F64_HBIT ? 1 : 2;
    int32_t  e0 = v->e - 1;
    int32_t  e1 = v->e - dv;
    uint64_t f0 = (v->f << 1) + 1;
    uint64_t f1 = (v->f << dv) - 1;
    uint32_t sh = __builtin_clzll(f0);

    /* calculate the m+ */
    p->e = e0 - sh;
    p->f = f0 << sh;

    /* calculate the m- */
    m->e = p->e;
    m->f = f1 << (e1 - p->e);
}

static inline void f64_power(struct f64_t *v, int e, int *k) {
    double   dk = (-61 - e) * DP_1_LG10 + 347;
	int32_t  ik = (int32_t)dk;
    uint32_t id;

    /* ceil the logrithmic result */
	if (dk - ik > 0.0) {
        ik++;
    }

    /* calculate the K value */
	id = (uint32_t)(ik >> 3) + 1;
	*k = 348 - (int32_t)(id << 3);

    /* lookup the power */
    v->e = TabPowE[id];
    v->f = TabPowF[id];
}

/** Florian's Grisu2 Algorithm **/

static const char TabDigits[200] = {
    '0', '0', '0', '1', '0', '2', '0', '3', '0', '4', '0', '5', '0', '6', '0', '7', '0', '8', '0', '9',
    '1', '0', '1', '1', '1', '2', '1', '3', '1', '4', '1', '5', '1', '6', '1', '7', '1', '8', '1', '9',
    '2', '0', '2', '1', '2', '2', '2', '3', '2', '4', '2', '5', '2', '6', '2', '7', '2', '8', '2', '9',
    '3', '0', '3', '1', '3', '2', '3', '3', '3', '4', '3', '5', '3', '6', '3', '7', '3', '8', '3', '9',
    '4', '0', '4', '1', '4', '2', '4', '3', '4', '4', '4', '5', '4', '6', '4', '7', '4', '8', '4', '9',
    '5', '0', '5', '1', '5', '2', '5', '3', '5', '4', '5', '5', '5', '6', '5', '7', '5', '8', '5', '9',
    '6', '0', '6', '1', '6', '2', '6', '3', '6', '4', '6', '5', '6', '6', '6', '7', '6', '8', '6', '9',
    '7', '0', '7', '1', '7', '2', '7', '3', '7', '4', '7', '5', '7', '6', '7', '7', '7', '8', '7', '9',
    '8', '0', '8', '1', '8', '2', '8', '3', '8', '4', '8', '5', '8', '6', '8', '7', '8', '8', '8', '9',
    '9', '0', '9', '1', '9', '2', '9', '3', '9', '4', '9', '5', '9', '6', '9', '7', '9', '8', '9', '9'
};

static inline int ctz10(uint64_t n) {
    if (n <         10) return 1;
	if (n <        100) return 2;
	if (n <       1000) return 3;
	if (n <      10000) return 4;
	if (n <     100000) return 5;
	if (n <    1000000) return 6;
	if (n <   10000000) return 7;
	if (n <  100000000) return 8;
	if (n < 1000000000) return 9;
	                    return 10;
}

static inline int divmod(uint64_t *p1, int kp) {
    switch (kp) {
        case 10: kp = *p1 / 1000000000; *p1 %= 1000000000; return kp;
        case  9: kp = *p1 /  100000000; *p1 %=  100000000; return kp;
        case  8: kp = *p1 /   10000000; *p1 %=   10000000; return kp;
        case  7: kp = *p1 /    1000000; *p1 %=    1000000; return kp;
        case  6: kp = *p1 /     100000; *p1 %=     100000; return kp;
        case  5: kp = *p1 /      10000; *p1 %=      10000; return kp;
        case  4: kp = *p1 /       1000; *p1 %=       1000; return kp;
        case  3: kp = *p1 /        100; *p1 %=        100; return kp;
        case  2: kp = *p1 /         10; *p1 %=         10; return kp;
        case  1: kp = *p1;              *p1  =          0; return kp;
        default: __builtin_unreachable();
    }
}

static inline void roundg(char *p, uint64_t d, uint64_t r, uint64_t kp10, uint64_t dpw) {
    while (r < dpw && d - r >= kp10 && (r + kp10 < dpw || dpw - r > r + kp10 - dpw)) {
        r += kp10;
        p[-1] -= 1;
    }
}

static inline int digits(char *p, int *k, const struct f64_t *w, const struct f64_t *m, uint64_t d) {
    uint32_t     dv;
    uint64_t     vt;
    struct f64_t dpw;
    struct f64_t one;

    /* initial state */
    f64_sub(&dpw, m, w);
    f64_raw(&one, 1ull << -m->e, m->e);

    /* m+ cutoff */
    uint64_t p1 = m->f >> -one.e;
    uint64_t p2 = m->f & (one.f - 1);

    /* count the integer part length */
    char *  pb = p;
    int32_t kp = ctz10(p1);

    /* small values */
    while (kp > 0) {
        dv = divmod(&p1, kp);
        kp--;

        /* write one digit */
        if (dv || p > pb) {
            *p++ = (char)(dv + '0');
        }

        /* calculate the error */
        vt = p1 << -one.e;
        vt += p2;

        /* check the precision */
        if (vt <= d) {
            *k += kp;
            roundg(p, d, vt, TabPow10[kp] << -one.e, dpw.f);
            return p - pb;
        }
    }

    /* large values (longer than 6 leading digits) */
    for (;;) {
        d  *= 10;
        p2 *= 10;
        dv  = (p2 >> -one.e) & 0xff;
        p2 &= one.f - 1;
        kp--;

        /* write one digit */
        if (dv || p > pb) {
            *p++ = (char)(dv + '0');
        }

        /* check the precision */
		if (p2 < d) {
			*k += kp;
            roundg(p, d, p2, one.f, dpw.f * TabPow10[-kp]);
			return p - pb;
		}
    }
}

static inline int grisu2(char *p, double v, int *k) {
    struct f64_t vv;
    struct f64_t mk;
    struct f64_t wm;
    struct f64_t wp;
    struct f64_t ww;

    /* calculate m+ and m- */
    f64_set   (&vv, v);
    f64_normb (&wm, &wp, &vv);
    f64_power (&mk, wp.e, k);
    f64_norm  (&ww, &vv);
    f64_mul   (&ww, &ww, &mk);
    f64_mul   (&wp, &wp, &mk);
    f64_mul   (&wm, &wm, &mk);

    /* generate the digits */
    wm.f++;
    wp.f--;
    return digits(p, k, &ww, &wp, wp.f - wm.f);
}

static inline void movchar(char *p, int n, int v, int d) {
    for (int i = n + d - 1; i >= v; i--) {
        p[i] = p[i - d];
    }
}

static inline void setchar(char *p, char c, int n) {
    while (n--) {
        *p++ = c;
    }
}

static inline void inschr1(char *p, char c, int n, int v) {
    movchar(p, n, v, 1);
    p[v] = c;
}

static inline void inschr2(char *p, char c0, char c1, int n, int v, int d) {
    movchar(p, n, v, d + 2);
    p[v + 0] = c0;
    p[v + 1] = c1;
}

static inline void setexpo(char *p, int *n, int k) {
    int    ex = k;
    int    n0 = *n;
    char * p0 = p;

    /* negative exponent */
    if (ex < 0) {
        ex   = -ex;
        *p++ = '-';
    }

    /* single digit exponent */
    if (ex < 10) {
        *n = p - p0 + n0 + 1;
        *p = (char)(ex + '0');
        return;
    }

    /* 2-digit exponent */
	if (ex < 100) {
        *n   = p - p0 + n0 + 2;
		*p++ = TabDigits[ex * 2];
		*p++ = TabDigits[ex * 2 + 1];
        return;
	}

    /* 3-digit exponent */
    *n   = p - p0 + n0 + 3;
    *p++ = (char)(ex / 100 + '0');
    *p++ = TabDigits[(ex % 100) * 2];
    *p++ = TabDigits[(ex % 100) * 2 + 1];
}

static inline void normalize(char *p, int *np, int k) {
    int n  = *np;
    int nk = n + k;

    /* case 1: p = "1234", k = 7 -> "12340000000" */
    if (n <= nk && nk <= 21) {
        *np = nk;
        setchar(p + n, '0', k);
        return;
    }

    /* case 2: p = "1234", k = -2 -> "12.34" */
    if (0 < nk && nk <= 21) {
        *np = n + 1;
        inschr1(p, '.', n, nk);
        return;
	}

    /* case 3: p = "1234", k = -6 -> "0.001234" */
    if (-6 < nk && nk <= 0) {
        *np = 2 - k;
        inschr2(p, '0', '.', n, 0, -nk);
        setchar(p + 2, '0', -nk);
        return;
	}

    /* case 4: p = "1", k = 30 -> "1e30" */
    if (n == 1) {
        (*np)++;
        p[1] = 'e';
        setexpo(p + 2, np, nk - 1);
        return;
    }

    /* case 5 (final case): p = "1234", k = 30 -> "1.234e33" */
    *np += 2;
    inschr1(p, '.', n, 1);
    setchar(p + n + 1, 'e', 1);
    setexpo(p + n + 2, np, nk - 1);
}

int f64toa(char *out, double val) {
    int    i = 0;
    char * p = out;

    /* simple case of 0.0 */
    if (val == 0.0) {
        *p = '0';
        return 1;
    }

    /* negative numbers */
    if (val < 0.0) {
        i = 1;
        val = -val;
        *p++ = '-';
    }

    /* print the number with Grisu2 algorithm */
    int k;
    int n = grisu2(p, val, &k);

    /* normalize the output, and adjust the length */
    normalize(p, &n, k);
    return n + i;
}
