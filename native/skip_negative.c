#include "scanning.h"

long skip_negative(const GoString *src, long *p) {
    return skip_negative_1(src, p);
}