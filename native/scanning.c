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

static const char *CS_ARRAY  = "[]{},\"[]{},\"[]{}";
static const char *CS_OBJECT = "[]{},:\"[]{}:,\"[]";

static const double P10_TAB[632] = {
    /* <================= -Inf ================= */ 1e-323, 1e-322, 1e-321, 1e-320,
    1e-319, 1e-318, 1e-317, 1e-316, 1e-315, 1e-314, 1e-313, 1e-312, 1e-311, 1e-310,
    1e-309, 1e-308, 1e-307, 1e-306, 1e-305, 1e-304, 1e-303, 1e-302, 1e-301, 1e-300,
    1e-299, 1e-298, 1e-297, 1e-296, 1e-295, 1e-294, 1e-293, 1e-292, 1e-291, 1e-290,
    1e-289, 1e-288, 1e-287, 1e-286, 1e-285, 1e-284, 1e-283, 1e-282, 1e-281, 1e-280,
    1e-279, 1e-278, 1e-277, 1e-276, 1e-275, 1e-274, 1e-273, 1e-272, 1e-271, 1e-270,
    1e-269, 1e-268, 1e-267, 1e-266, 1e-265, 1e-264, 1e-263, 1e-262, 1e-261, 1e-260,
    1e-259, 1e-258, 1e-257, 1e-256, 1e-255, 1e-254, 1e-253, 1e-252, 1e-251, 1e-250,
    1e-249, 1e-248, 1e-247, 1e-246, 1e-245, 1e-244, 1e-243, 1e-242, 1e-241, 1e-240,
    1e-239, 1e-238, 1e-237, 1e-236, 1e-235, 1e-234, 1e-233, 1e-232, 1e-231, 1e-230,
    1e-229, 1e-228, 1e-227, 1e-226, 1e-225, 1e-224, 1e-223, 1e-222, 1e-221, 1e-220,
    1e-219, 1e-218, 1e-217, 1e-216, 1e-215, 1e-214, 1e-213, 1e-212, 1e-211, 1e-210,
    1e-209, 1e-208, 1e-207, 1e-206, 1e-205, 1e-204, 1e-203, 1e-202, 1e-201, 1e-200,
    1e-199, 1e-198, 1e-197, 1e-196, 1e-195, 1e-194, 1e-193, 1e-192, 1e-191, 1e-190,
    1e-189, 1e-188, 1e-187, 1e-186, 1e-185, 1e-184, 1e-183, 1e-182, 1e-181, 1e-180,
    1e-179, 1e-178, 1e-177, 1e-176, 1e-175, 1e-174, 1e-173, 1e-172, 1e-171, 1e-170,
    1e-169, 1e-168, 1e-167, 1e-166, 1e-165, 1e-164, 1e-163, 1e-162, 1e-161, 1e-160,
    1e-159, 1e-158, 1e-157, 1e-156, 1e-155, 1e-154, 1e-153, 1e-152, 1e-151, 1e-150,
    1e-149, 1e-148, 1e-147, 1e-146, 1e-145, 1e-144, 1e-143, 1e-142, 1e-141, 1e-140,
    1e-139, 1e-138, 1e-137, 1e-136, 1e-135, 1e-134, 1e-133, 1e-132, 1e-131, 1e-130,
    1e-129, 1e-128, 1e-127, 1e-126, 1e-125, 1e-124, 1e-123, 1e-122, 1e-121, 1e-120,
    1e-119, 1e-118, 1e-117, 1e-116, 1e-115, 1e-114, 1e-113, 1e-112, 1e-111, 1e-110,
    1e-109, 1e-108, 1e-107, 1e-106, 1e-105, 1e-104, 1e-103, 1e-102, 1e-101, 1e-100,
    1e-099, 1e-098, 1e-097, 1e-096, 1e-095, 1e-094, 1e-093, 1e-092, 1e-091, 1e-090,
    1e-089, 1e-088, 1e-087, 1e-086, 1e-085, 1e-084, 1e-083, 1e-082, 1e-081, 1e-080,
    1e-079, 1e-078, 1e-077, 1e-076, 1e-075, 1e-074, 1e-073, 1e-072, 1e-071, 1e-070,
    1e-069, 1e-068, 1e-067, 1e-066, 1e-065, 1e-064, 1e-063, 1e-062, 1e-061, 1e-060,
    1e-059, 1e-058, 1e-057, 1e-056, 1e-055, 1e-054, 1e-053, 1e-052, 1e-051, 1e-050,
    1e-049, 1e-048, 1e-047, 1e-046, 1e-045, 1e-044, 1e-043, 1e-042, 1e-041, 1e-040,
    1e-039, 1e-038, 1e-037, 1e-036, 1e-035, 1e-034, 1e-033, 1e-032, 1e-031, 1e-030,
    1e-029, 1e-028, 1e-027, 1e-026, 1e-025, 1e-024, 1e-023, 1e-022, 1e-021, 1e-020,
    1e-019, 1e-018, 1e-017, 1e-016, 1e-015, 1e-014, 1e-013, 1e-012, 1e-011, 1e-010,
    1e-009, 1e-008, 1e-007, 1e-006, 1e-005, 1e-004, 1e-003, 1e-002, 1e-001, 1e-000,
    1e+001, 1e+002, 1e+003, 1e+004, 1e+005, 1e+006, 1e+007, 1e+008, 1e+009, 1e+010,
    1e+011, 1e+012, 1e+013, 1e+014, 1e+015, 1e+016, 1e+017, 1e+018, 1e+019, 1e+020,
    1e+021, 1e+022, 1e+023, 1e+024, 1e+025, 1e+026, 1e+027, 1e+028, 1e+029, 1e+030,
    1e+031, 1e+032, 1e+033, 1e+034, 1e+035, 1e+036, 1e+037, 1e+038, 1e+039, 1e+040,
    1e+041, 1e+042, 1e+043, 1e+044, 1e+045, 1e+046, 1e+047, 1e+048, 1e+049, 1e+050,
    1e+051, 1e+052, 1e+053, 1e+054, 1e+055, 1e+056, 1e+057, 1e+058, 1e+059, 1e+060,
    1e+061, 1e+062, 1e+063, 1e+064, 1e+065, 1e+066, 1e+067, 1e+068, 1e+069, 1e+070,
    1e+071, 1e+072, 1e+073, 1e+074, 1e+075, 1e+076, 1e+077, 1e+078, 1e+079, 1e+080,
    1e+081, 1e+082, 1e+083, 1e+084, 1e+085, 1e+086, 1e+087, 1e+088, 1e+089, 1e+090,
    1e+091, 1e+092, 1e+093, 1e+094, 1e+095, 1e+096, 1e+097, 1e+098, 1e+099, 1e+100,
    1e+101, 1e+102, 1e+103, 1e+104, 1e+105, 1e+106, 1e+107, 1e+108, 1e+109, 1e+110,
    1e+111, 1e+112, 1e+113, 1e+114, 1e+115, 1e+116, 1e+117, 1e+118, 1e+119, 1e+120,
    1e+121, 1e+122, 1e+123, 1e+124, 1e+125, 1e+126, 1e+127, 1e+128, 1e+129, 1e+130,
    1e+131, 1e+132, 1e+133, 1e+134, 1e+135, 1e+136, 1e+137, 1e+138, 1e+139, 1e+140,
    1e+141, 1e+142, 1e+143, 1e+144, 1e+145, 1e+146, 1e+147, 1e+148, 1e+149, 1e+150,
    1e+151, 1e+152, 1e+153, 1e+154, 1e+155, 1e+156, 1e+157, 1e+158, 1e+159, 1e+160,
    1e+161, 1e+162, 1e+163, 1e+164, 1e+165, 1e+166, 1e+167, 1e+168, 1e+169, 1e+170,
    1e+171, 1e+172, 1e+173, 1e+174, 1e+175, 1e+176, 1e+177, 1e+178, 1e+179, 1e+180,
    1e+181, 1e+182, 1e+183, 1e+184, 1e+185, 1e+186, 1e+187, 1e+188, 1e+189, 1e+190,
    1e+191, 1e+192, 1e+193, 1e+194, 1e+195, 1e+196, 1e+197, 1e+198, 1e+199, 1e+200,
    1e+201, 1e+202, 1e+203, 1e+204, 1e+205, 1e+206, 1e+207, 1e+208, 1e+209, 1e+210,
    1e+211, 1e+212, 1e+213, 1e+214, 1e+215, 1e+216, 1e+217, 1e+218, 1e+219, 1e+220,
    1e+221, 1e+222, 1e+223, 1e+224, 1e+225, 1e+226, 1e+227, 1e+228, 1e+229, 1e+230,
    1e+231, 1e+232, 1e+233, 1e+234, 1e+235, 1e+236, 1e+237, 1e+238, 1e+239, 1e+240,
    1e+241, 1e+242, 1e+243, 1e+244, 1e+245, 1e+246, 1e+247, 1e+248, 1e+249, 1e+250,
    1e+251, 1e+252, 1e+253, 1e+254, 1e+255, 1e+256, 1e+257, 1e+258, 1e+259, 1e+260,
    1e+261, 1e+262, 1e+263, 1e+264, 1e+265, 1e+266, 1e+267, 1e+268, 1e+269, 1e+270,
    1e+271, 1e+272, 1e+273, 1e+274, 1e+275, 1e+276, 1e+277, 1e+278, 1e+279, 1e+280,
    1e+281, 1e+282, 1e+283, 1e+284, 1e+285, 1e+286, 1e+287, 1e+288, 1e+289, 1e+290,
    1e+291, 1e+292, 1e+293, 1e+294, 1e+295, 1e+296, 1e+297, 1e+298, 1e+299, 1e+300,
    1e+301, 1e+302, 1e+303, 1e+304, 1e+305, 1e+306, 1e+307, 1e+308, /* = +Inf => */
};

static inline double pow10(double v, int p) {
    if (p < -323) {
        return 0.0;
    } else if (p > 308) {
        return __builtin_inf();
    } else {
        return v * P10_TAB[p + 323];
    }
}

static inline void vdigits(const GoString *src, long *p, JsonState *ret) {
    --*p;
    vnumber(src, p, ret);
}

static inline char advance(const GoString *src, long *p) {
    if (*p >= src->len) {
        return 0;
    } else {
        return src->buf[(*p)++];
    }
}

static inline char advance_ns(const GoString *src, long *p) {
    *p = lspace(src->buf, src->len, *p);
    return advance(src, p);
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
        for (int i = 0; src->buf[*p] == (val & 0xff); i++, ++*p) { val >>= 8; }
        return -ERR_INVAL;
    }
}

static inline ssize_t advance_string(const GoString *src, long *p, int64_t *ep) {
    ssize_t e;
    ssize_t i;

    /* check for end of string */
    if ((e = strchr2(src, *p, '"', '\\')) < 0) {
        *p = src->len;
        return -ERR_EOF;
    }

    /* encounters a '"' at the first scan, it's an unquoted string */
    if (src->buf[e] == '"') {
        *ep = -1;
        return e;
    }

    /* search for the next double quote */
    i = e;
    e = strchr1(src, e + 1, '"');

    /* seek to the end of string */
    while (e >= 0) {
        int          n = 0;
        const char * q = src->buf + e;

        /* counting backslashes */
        while (*--q == '\\') {
            n++;
        }

        /* pairs of backslashes cancel each other out */
        if ((n & 1) == 0) {
            break;
        }

        /* find the next double quote */
        e = e + 1;
        e = strchr1(src, e, '"');
    }

    /* check for end of string */
    if (e < 0) {
        *p = src->len;
        return -ERR_EOF;
    }

    /* update the result */
    *ep = i;
    return e;
}

static inline int64_t advance_number(const GoString *src, long *p, long i, JsonState *ret, size_t sp) {
    size_t       n = src->len;
    const char * s = src->buf;

    /* check for EOF */
    if (i >= n) {
        *p = n;
        return -ERR_EOF;
    }

    /* base factors */
    int     esm = 1;
    int     exp = 0;
    int     ovf = 0;
    int     rem = 0;
    double  val = 0;
    int64_t idx = 0;
    int64_t i64 = 0;
    int64_t rvt = V_INTEGER;

    /* initial state */
    ret->iv = 0;
    ret->ep = sp;

    /* check for the special case of '0' */
    if (s[i] == '0') {
        if (++i >= n) {
            *p = i;
            return V_INTEGER;
        }
    } else {
        if (s[i] < '0' || s[i] > '9') {
            *p = i;
            return -ERR_INVAL;
        } else {
            while (!(ovf = __builtin_smulll_overflow((idx = i64), 10, &i64)) &&
                   !(ovf = __builtin_saddll_overflow(i64, s[i] - '0', &i64)) &&
                   !(++i >= n || !(s[i] >= '0' && s[i] <= '9')));
        }
    }

    /* set the integer part */
    ret->iv = i64;
    ret->dv = i64;

    /* check for integer overflow, in such case
     * the number must be represented by double */
    if (ovf) {
        val = idx;
        rvt = V_DOUBLE;

        /* convert the remaining digits */
        do {
            val *= 10;
            val += s[i++] - '0';
        } while (i < n && s[i] >= '0' && s[i] <= '9');

        /* set the integer part to INT64_MAX to indicate an overflow */
        ret->dv = val;
        ret->iv = INT64_MAX;
    }

    /* check for decimal points */
    if (i < n && s[i] == '.') {
        idx = ++i;
        rvt = V_DOUBLE;

        /* check for EOF */
        if (i >= n) {
            *p = n;
            return -ERR_EOF;
        }

        /* should be a digit */
        if (s[i] < '0' || s[i] > '9') {
            *p = i;
            return -ERR_INVAL;
        }

        /* convert the fractional part */
        do {
            rem *= 10;
            rem += s[i++] - '0';
        } while (i < n && s[i] >= '0' && s[i] <= '9');

        /* combine with the integer part */
        idx -= i;
        ret->dv += pow10(rem, idx);
    }

    /* check for exponent */
    if (i < n && (s[i] == 'e' || s[i] == 'E')) {
        i++;
        rvt = V_DOUBLE;

        /* check for EOF */
        if (i >= n) {
            *p = n;
            return -ERR_EOF;
        }

        /* check for the '+' or '-' sign */
        if (s[i] == '+' || s[i] == '-') {
            if (i >= n - 1) {
                *p = i;
                return -ERR_EOF;
            } else {
                if (s[i++] == '+') {
                    esm = 1;
                } else {
                    esm = -1;
                }
            }
        }

        /* should be a digit */
        if (s[i] < '0' || s[i] > '9') {
            *p = i;
            return -ERR_INVAL;
        }

        /* convert the power */
        do {
            exp *= 10;
            exp += s[i++] - '0';
        } while (i < n && s[i] >= '0' && s[i] <= '9');

        /* apply the power */
        exp *= esm;
        ret->dv = pow10(ret->dv, exp);
    }

    /* calculate the offset */
    *p = i;
    return rvt;
}

/** Value Scanning Routines **/

long value(const char *s, size_t n, long p, JsonState *ret) {
    long     q = p;
    GoString m = {.buf = s, .len = n};

    /* parse the next identifier */
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
        case '9' : vdigits(&m, &q, ret)                                 ; return q;
        case '"' : vstring(&m, &q, ret)                                 ; return q;
        case 'n' : ret->vt = advance_dword(&m, &q, 1, V_NULL, VS_NULL)  ; return q;
        case 't' : ret->vt = advance_dword(&m, &q, 1, V_TRUE, VS_TRUE)  ; return q;
        case 'f' : ret->vt = advance_dword(&m, &q, 0, V_FALSE, VS_ALSE) ; return q;
        case '[' : ret->vt = V_ARRAY                                    ; return q;
        case '{' : ret->vt = V_OBJECT                                   ; return q;
        case  0  : ret->vt = V_EOF                                      ; return q;
        default  : ret->vt = -ERR_INVAL                                 ; return q;
    }
}

void vstring(const GoString *src, long *p, JsonState *ret) {
    int64_t i = *p;
    ssize_t e = advance_string(src, p, &ret->ep);

    /* check for errors */
    if (e < 0) {
        ret->vt = e;
        return;
    }

    /* update the result */
    *p = e + 1;
    ret->iv = i;
    ret->vt = V_STRING;
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

void vnumber(const GoString *src, long *p, JsonState *ret) {
    int     dig;
    int     ovf = 0;
    int     sgn = 1;
    double  val = 0;
    int64_t dec = 0;

    /* initial buffer pointers */
    long         i = *p;
    size_t       n = src->len;
    const char * s = src->buf;

    /* initialize the result, and check for EOF */
    init_ret(V_INTEGER)
    check_eof()
    check_sign(sgn = -1)

    /* check for leading zero */
    check_digit()
    check_leading_zero()

    /* parse the integer part */
    while (i < n && s[i] >= '0' && s[i] <= '9') {
        val = dec;
        dig = sgn * (s[i++] - '0');

        /* add the digit to the integer part if not overflowed */
        if (add_digit_overflow(dec, dig)) {
            ovf = 1;
            val = (val * 10.0) + dig;
            set_vt(V_DOUBLE)
            break;
        }
    }

    /* after overflow, we can only continue with floating point values */
    if (!ovf) {
        val = dec;
    }

    /* after overflow, we can only continue with floating point values */
    while (ovf && i < n && s[i] >= '0' && s[i] <= '9') {
        val *= 10.0;
        val += sgn * (s[i++] - '0');
    }

    /* check for decimal points */
    if (i < n && s[i] == '.') {
        int     idx = ++i;
        int64_t rem = 0;

        /* convert the fractional part (float64 can represent at most 16 digits) */
        set_vt(V_DOUBLE)
        check_eof()
        check_digit()
        parse_float_digits(rem, sgn, && i <= idx + 17)

        /* combine with the decimal part */
        idx -= i;
        val += pow10(rem, idx);

        /* skip the remaining digits */
        while (i < n && s[i] >= '0' && s[i] <= '9') {
            i++;
        }
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
        parse_float_digits(exp, esm)

        /* scale with */
        if (exp != 1) {
            val = pow10(val, exp);
        }
    }

    /* check for integer overflow */
    if (!ovf) {
        ret->iv = dec;
    }

    /* update the result */
    *p = i;
    ret->dv = val;
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

static inline long fsm_exec(StateMachine *self, const GoString *src, long *p) {
    int  vt;
    char ch;
    long vi = -1;

    /* run until no more nested values */
    while (self->sp) {
        ch = advance_ns(src, p);
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
                FSM_XERR(skip_string(src, p));
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
                        FSM_XERR(skip_string(src, p));
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
            case '"' : FSM_XERR(skip_string(src, p));                       break;
            case '[' : FSM_XERR(fsm_push(self, FSM_ARR_0));                 break;
            case '{' : FSM_XERR(fsm_push(self, FSM_OBJ_0));                 break;
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

long skip_one(const GoString *src, long *p, StateMachine *m) {
    fsm_init(m, FSM_VAL);
    return fsm_exec(m, src, p);
}

long skip_array(const GoString *src, long *p, StateMachine *m) {
    fsm_init(m, FSM_ARR_0);
    return fsm_exec(m, src, p);
}

long skip_object(const GoString *src, long *p, StateMachine *m) {
    fsm_init(m, FSM_OBJ_0);
    return fsm_exec(m, src, p);
}

long skip_string(const GoString *src, long *p) {
    int64_t v;
    ssize_t q = *p - 1;
    ssize_t e = advance_string(src, p, &v);

    /* check for errors */
    if (e < 0) {
        return e;
    }

    /* update the position */
    *p = e + 1;
    return q;
}

long skip_negative(const GoString *src, long *p) {
    long      q = *p - 1;
    int64_t   r;
    JsonState v;

    /* skip the number */
    if ((r = advance_number(src, p, *p, &v, q)) < 0) {
        return r;
    } else {
        return q;
    }
}

long skip_positive(const GoString *src, long *p) {
    long      q = *p - 1;
    int64_t   r;
    JsonState v;

    /* skip the number */
    if ((r = advance_number(src, p, q, &v, q)) < 0) {
        return r;
    } else {
        return q;
    }
}
