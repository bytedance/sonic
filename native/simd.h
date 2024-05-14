

#include "native.h"

// Portable SIMD Helper -------------------------------------------------------

typedef __m128i v128u;
typedef __m128i mask128;

static always_inline v128u v128_loadu(const uint8_t* ptr) {
    return _mm_loadu_si128((__m128i*)ptr);
}

static always_inline void v128_storeu(const v128u v, uint8_t* ptr) {
    _mm_storeu_si128((__m128i*)ptr, v);
}

static always_inline mask128 v128_eq(const v128u v1, const v128u v2) {
    return _mm_cmpeq_epi8(v1, v2);
}

static always_inline mask128 v128_le(const v128u v1, const v128u v2) {
    __m128i max = _mm_max_epu8(v1, v2);
    return _mm_cmpeq_epi8(max, v2);
}

static always_inline mask128 v128_gt(const v128u v1, const v128u v2) {
    __m128i sub = _mm_subs_epu8(v1, v2); 
    return _mm_xor_si128(_mm_cmpeq_epi8(sub, _mm_setzero_si128()), _mm_set1_epi8('\xff'));
}

static always_inline v128u v128_splat(uint8_t ch) {
    return _mm_set1_epi8((char)ch);
}

static always_inline uint16_t mask128_tobitmask(mask128 mask) {
    return  (uint16_t)_mm_movemask_epi8(mask);
}

static always_inline mask128 mask128_and(mask128 mask1, mask128 mask2) {
    return _mm_and_si128(mask1, mask2);
}

static always_inline mask128 mask128_or(mask128 mask1, mask128 mask2) {
    return _mm_or_si128(mask1, mask2);
}

#ifdef __AVX2__

typedef __m256i v256u;
typedef __m256i mask256;

static inline v256u v256_loadu(const uint8_t* ptr) {
    return _mm256_loadu_si256((__m256i*)ptr);
}

static inline void v256_storeu(const v256u v, uint8_t* ptr) {
    _mm256_storeu_si256((__m256i*)ptr, v);
}

static inline mask256 v256_eq(const v256u v1, const v256u v2) {
    return _mm256_cmpeq_epi8(v1, v2);
}

static inline mask256 v256_le(const v256u v1, const v256u v2) {
    __m256i max = _mm256_max_epu8(v1, v2);
    return _mm256_cmpeq_epi8(max, v2);
}

static inline mask256 v256_gt(const v256u v1, const v256u v2) {
    __m256i sub = _mm256_subs_epu8(v1, v2); 
    return _mm256_xor_si256(_mm256_cmpeq_epi8(sub, _mm256_setzero_si256()), _mm256_set1_epi8('\xff'));
}

static inline v256u v256_splat(uint8_t ch) {
    return _mm256_set1_epi8((char)ch);
}

static inline uint32_t mask256_tobitmask(mask256 mask) {
    return _mm256_movemask_epi8(mask);
}

static inline mask256 mask256_and(mask256 mask1, mask256 mask2) {
    return _mm256_and_si256(mask1, mask2);
}

static inline mask256 mask256_or(mask256 mask1, mask256 mask2) {
    return _mm256_or_si256(mask1, mask2);
}

#else

typedef struct {
    v128u lo;
    v128u hi;   
} v256u;

typedef struct {
    mask128 lo;
    mask128 hi;   
} mask256;

static inline v256u v256_loadu(const uint8_t* ptr) {
    return (v256u){_mm_loadu_si128((__m128i*)ptr), _mm_loadu_si128((__m128i*)(ptr + 16))};
}

static inline void v256_storeu(const v256u v, uint8_t* ptr) {
    _mm_storeu_si128((__m128i*)ptr, v.lo);
    _mm_storeu_si128((__m128i*)(ptr + 16), v.hi);
}

static inline mask256 v256_eq(const v256u v1, const v256u v2) {
    return (mask256){ v128_eq(v1.lo, v2.lo), v128_eq(v1.hi, v2.hi)};
}

static inline mask256 v256_le(const v256u v1, const v256u v2) {
    return (mask256){ v128_le(v1.lo, v2.lo), v128_le(v1.hi, v2.hi)};
}

static inline mask256 v256_gt(const v256u v1, const v256u v2) {
    return (mask256){ v128_gt(v1.lo, v2.lo), v128_gt(v1.hi, v2.hi)};
}

static inline v256u v256_splat(uint8_t ch) {
    return (v256u){v128_splat(ch), v128_splat(ch)};
}

static inline uint32_t mask256_tobitmask(mask256 mask) {
    uint32_t lo = mask128_tobitmask(mask.lo);
    uint32_t hi = mask128_tobitmask(mask.hi);
    return lo | (hi << 16);
}

static inline mask256 mask256_and(mask256 mask1, mask256 mask2) {
    return (mask256){mask128_and(mask1.lo, mask2.lo), mask128_and(mask1.hi, mask2.hi)};
}

static inline mask256 mask256_or(mask256 mask1, mask256 mask2) {
    return (mask256){mask128_or(mask1.lo, mask2.lo), mask128_or(mask1.hi, mask2.hi)};
}

#endif

// TODO: support NEON, AVX512