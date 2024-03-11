#include "scanning.h"

long skip_one(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    return skip_one_1(src, p, m, flags);
}
