
#include "parsing.h"

ssize_t html_escape(const char *sp, ssize_t nb, char *dp, ssize_t *dn) {
    ssize_t          nd  = *dn;
    const char     * ds  = dp;
    const char     * ss  = sp;
    const quoted_t * tab = _HtmlQuoteTab;

    /* find the special characters, copy on the fly */
    while (nb > 0) {
        int     nc = 0;
        uint8_t ch = 0;
        ssize_t rb = 0;
        const char * cur = 0;

        /* not enough buffer space */
        if (nd <= 0) {
            return -(sp - ss) - 1;
        }

        /* find and copy */
        if ((rb = memcchr_html_quote(sp, nb, dp, nd)) < 0) {
            *dn = dp - ds - rb - 1;
            return -(sp - ss - rb - 1) - 1;
        }

        /* skip already copied bytes */
        sp += rb;
        dp += rb;
        nb -= rb;
        nd -= rb;

        /* stop if already finished */
        if (nb <= 0) {
            break;
        }

        /* mark cur postion */
        cur = sp;

        /* check for \u2028 and \u2029, binary is \xe2\x80\xa8 and \xe2\x80\xa9 */
        if (unlikely(*sp == '\xe2')) {
            if (nb >= 3 && *(sp+1) == '\x80' && (*(sp+2) == '\xa8' || *(sp+2) == '\xa9')) {
                sp += 2, nb -= 2;
            } else if (nd > 0) {
                *dp++ = *sp++;
                nb--, nd--;
                continue;
            } else {
                return -(sp - ss) - 1;
            }
        }

        /* get the escape entry, handle consecutive quotes */
        ch = * (uint8_t*) sp;
        nc = tab[ch].n;


        /* check for buffer space */
        if (nd < nc) {
            *dn = dp - ds;
            return -(cur - ss) - 1;
        }

        /* copy the quoted value */
        memcpy_p8(dp, tab[ch].s, nc);
        sp++;
        nb--;
        dp += nc;
        nd -= nc;
    }

    /* all done */
    *dn = dp - ds;
    return sp - ss;
}
