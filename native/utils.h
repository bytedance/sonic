/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include <immintrin.h>
#include <string.h>
#include "native.h"

static always_inline bool vec_cross_page(const void * p, size_t n) {
#define PAGE_SIZE 4096
    return (((size_t)(p)) & (PAGE_SIZE - 1)) > (PAGE_SIZE - n);
#undef PAGE_SIZE
}

static always_inline void memcpy4 (void *__restrict dp, const void *__restrict sp) {
    ((uint32_t *)dp)[0] = ((const uint32_t *)sp)[0];
}

static always_inline void memcpy8 (void *__restrict dp, const void *__restrict sp) {
    ((uint64_t *)dp)[0] = ((const uint64_t *)sp)[0];
}

static always_inline void memcpy16 (void *__restrict dp, const void *__restrict sp) {
    _mm_storeu_si128((void *)(dp), _mm_loadu_si128((const void *)(sp)));
}

static always_inline void memcpy32(void *__restrict dp, const void *__restrict sp) {
#if USE_AVX2
    _mm256_storeu_si256((void *)dp,     _mm256_loadu_si256((const void *)sp));
#else
    _mm_storeu_si128((void *)(dp),      _mm_loadu_si128((const void *)(sp)));
    _mm_storeu_si128((void *)(dp + 16), _mm_loadu_si128((const void *)(sp + 16)));
#endif
}

static always_inline void memcpy64(void *__restrict dp, const void *__restrict sp) {
    memcpy32(dp, sp);
    memcpy32(dp + 32, sp + 32);
}

static always_inline void memcpy_p4(void *__restrict dp, const void *__restrict sp, size_t nb) {
    if (nb >= 2) { *(uint16_t *)dp = *(const uint16_t *)sp; sp += 2, dp += 2, nb -= 2; }
    if (nb >= 1) { *(uint8_t *) dp = *(const uint8_t *)sp; }
}

static always_inline void memcpy_p8(void *__restrict dp, const void *__restrict sp, ssize_t nb) {
    if (nb >= 4) { memcpy4(dp, sp); sp += 4, dp += 4, nb -= 4; }
    memcpy_p4(dp, sp, nb);
}

static always_inline void memcpy_p16(void *__restrict dp, const void *__restrict sp, size_t nb) {
    if (nb >= 8) { memcpy8(dp, sp); sp += 8, dp += 8, nb -= 8; }
    memcpy_p8(dp, sp, nb);
}

static always_inline void memcpy_p32(void *__restrict dp, const void *__restrict sp, size_t nb) {
    if (nb >= 16) { memcpy16(dp, sp); sp += 16, dp += 16, nb -= 16; }
    memcpy_p16(dp, sp, nb);
}

static always_inline void memcpy_p64(void *__restrict dp, const void *__restrict sp, size_t nb) {
    if (nb >= 32) { memcpy32(dp, sp); sp += 32, dp += 32, nb -= 32; }
    memcpy_p32(dp, sp, nb);
}
