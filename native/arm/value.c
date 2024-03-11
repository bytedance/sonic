#include "scanning.h"
#include "vstring.h"

long value(const char *s, size_t n, long p, JsonState *ret, uint64_t flags) {
    long     q = p;
    GoString m = {.buf = s, .len = n};
    bool allow_control = (flags & MASK_ALLOW_CONTROL) != 0;
    /* parse the next identifier, q is UNSAFE, may cause out-of-bounds accessing */
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
        case '9' : vdigits(&m, &q, ret, flags)                          ; return q;
        case '"' : vstring_1(&m, &q, ret, flags)                          ; return q;
        case 'n' : ret->vt = advance_dword(&m, &q, 1, V_NULL, VS_NULL)  ; return q;
        case 't' : ret->vt = advance_dword(&m, &q, 1, V_TRUE, VS_TRUE)  ; return q;
        case 'f' : ret->vt = advance_dword(&m, &q, 0, V_FALSE, VS_ALSE) ; return q;
        case '[' : ret->vt = V_ARRAY                                    ; return q;
        case '{' : ret->vt = V_OBJECT                                   ; return q;
        case ':' : ret->vt = allow_control ? V_KEY_SEP : -ERR_INVAL     ; return allow_control ? q : q - 1;
        case ',' : ret->vt = allow_control ? V_ELEM_SEP : -ERR_INVAL    ; return allow_control ? q : q - 1;
        case ']' : ret->vt = allow_control ? V_ARRAY_END : -ERR_INVAL   ; return allow_control ? q : q - 1;
        case '}' : ret->vt = allow_control ? V_OBJECT_END : -ERR_INVAL  ; return allow_control ? q : q - 1;
        case  0  : ret->vt = V_EOF                                      ; return q;
        default  : ret->vt = -ERR_INVAL                                 ; return q - 1;
    }
}
