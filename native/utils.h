#pragma once

#include <immintrin.h>
#include "native.h"

static always_inline bool vec_cross_page(const void * p, size_t n) {
#define PAGE_SIZE 4096
    return (((size_t)(p)) & (PAGE_SIZE - 1)) > (PAGE_SIZE - n);
#undef PAGE_SIZE
}

static inline void memcpy_p4(char *dp, const char *sp, size_t nb) {
    if (nb >= 2) { *(uint16_t *)dp = *(const uint16_t *)sp; sp += 2, dp += 2, nb -= 2; }
    if (nb >= 1) { *dp = *sp; }
}

static inline void memcpy_p8(char *dp, const char *sp, ssize_t nb) {
    if (nb >= 4) { *(uint32_t *)dp = *(const uint32_t *)sp; sp += 4, dp += 4, nb -= 4; }
    if (nb >= 2) { *(uint16_t *)dp = *(const uint16_t *)sp; sp += 2, dp += 2, nb -= 2; }
    if (nb >= 1) { *dp = *sp; }
}

static inline void memcpy_p16(char *dp, const char *sp, size_t nb) {
    if (nb >= 8) { *(uint64_t *)dp = *(const uint64_t *)sp; sp += 8, dp += 8, nb -= 8; }
    if (nb >= 4) { *(uint32_t *)dp = *(const uint32_t *)sp; sp += 4, dp += 4, nb -= 4; }
    if (nb >= 2) { *(uint16_t *)dp = *(const uint16_t *)sp; sp += 2, dp += 2, nb -= 2; }
    if (nb >= 1) { *dp = *sp; }
}

static inline void memcpy_p32(char *dp, const char *sp, size_t nb) {
    if (nb >= 16) { _mm_storeu_si128((void *)dp, _mm_loadu_si128((const void *)sp)); sp += 16, dp += 16, nb -= 16; }
    if (nb >=  8) { *(uint64_t *)dp = *(const uint64_t *)sp;                         sp +=  8, dp +=  8, nb -=  8; }
    if (nb >=  4) { *(uint32_t *)dp = *(const uint32_t *)sp;                         sp +=  4, dp +=  4, nb -=  4; }
    if (nb >=  2) { *(uint16_t *)dp = *(const uint16_t *)sp;                         sp +=  2, dp +=  2, nb -=  2; }
    if (nb >=  1) { *dp = *sp; }
}

static always_inline void memcpy_p64(char * restrict dp, const char * restrict sp, size_t n) {
    long nb = n;
#if USE_AVX2
    if (nb >= 32) { _mm256_storeu_si256((void *)dp, _mm256_loadu_si256((const void *)sp)); sp += 32, dp += 32, nb -= 32; }
#endif
    while (nb >= 16) { _mm_storeu_si128((void *)dp, _mm_loadu_si128((const void *)sp)); sp += 16, dp += 16, nb -= 16; }
    if (nb >=  8) { *(uint64_t *)dp = *(const uint64_t *)sp;                         sp +=  8, dp +=  8, nb -=  8; }
    if (nb >=  4) { *(uint32_t *)dp = *(const uint32_t *)sp;                         sp +=  4, dp +=  4, nb -=  4; }
    if (nb >=  2) { *(uint16_t *)dp = *(const uint16_t *)sp;                         sp +=  2, dp +=  2, nb -=  2; }
    if (nb >=  1) { *dp = *sp; }
}

static always_inline void memcpy8 (char *dp, const char *sp) {
    ((uint64_t *)dp)[0] = ((const uint64_t *)sp)[0];
}

static always_inline void memcpy32(char *dp, const char *sp) {
#if USE_AVX2
    _mm256_storeu_si256((void *)dp, _mm256_loadu_si256((const void *)sp));
#else
    _mm_storeu_si128((void *)(dp),      _mm_loadu_si128((const void *)(sp)));
    _mm_storeu_si128((void *)(dp + 16), _mm_loadu_si128((const void *)(sp + 16)));
#endif
}

static always_inline void memcpy64(char *dp, const char *sp) {
    memcpy32(dp, sp);
    memcpy32(dp + 32, sp + 32);
}