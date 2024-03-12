#include "parsing.h"


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