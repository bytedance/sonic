#include "scanning.h"

static always_inline long skip_string(const GoString *src, long *p, uint64_t flags) {
    return skip_string_1(src, p, flags);
}
