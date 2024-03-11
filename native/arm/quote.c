#include "parsing.h"

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
