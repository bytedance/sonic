#include "scanning.h"

INLINE_ALL long skip_one(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    return skip_one_1(src, p, m, flags);
}
