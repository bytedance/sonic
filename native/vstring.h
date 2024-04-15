
#pragma once

#include "scanning.h"

static always_inline void vstring_1(const GoString *src, long *p, JsonState *ret, uint64_t flags) {
    int64_t v = -1;
    int64_t i = *p;
    ssize_t e = advance_string(src, i, &v, flags);
   
    /* check for errors */
    if (e < 0) {
        *p = src->len;
        ret->vt = e;
        return;
    }

    /* update the result, and fix the escape position (escaping past the end of string) */
    *p = e;
    ret->iv = i;
    ret->vt = V_STRING;
    ret->ep = v >= e ? -1 : v;
}
