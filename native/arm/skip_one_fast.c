#include "scanning.h"

long skip_one_fast(const GoString *src, long *p) {
    return skip_one_fast_1(src, p);
}