#include "scanning.h"

long validate_utf8(const GoString *src, long *p, StateMachine *m) {
    xassert(*p >= 0 && src->len > *p);
    return validate_utf8_with_errors(src->buf, src->len, p, m);
}