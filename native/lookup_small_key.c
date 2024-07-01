#include "native.h"
 #include "simd.h"
 #include "test/xprintf.h"

// Fast Lookup Table -------------------------------------------------------

static const int _HDR_SLOTS = 33;
static const int _HDR_SLOT_SIZE  = 5;
static const int _HDR_SIZE  = _HDR_SLOTS * _HDR_SLOT_SIZE;

static const long _NOT_FOUND         = 1;

/**
 * @brief uppercase to lowercase
 * @param[in] src original string
 * @param[in, out] dst lowercase string
 * @param[in] len length of the string
*/
static always_inline void to_lower(uint8_t* dst, const uint8_t* src, size_t len) {
#if defined(__AVX2__)
    const __m256i _A = _mm256_set1_epi8('A' - 1);
    const __m256i Z_ = _mm256_set1_epi8('Z');
    const __m256i delta = _mm256_set1_epi8('a' - 'A');
    uint8_t* q = dst;

    while (len >= 32){
        __m256i op = _mm256_loadu_si256((__m256i*)src);
        __m256i gt = _mm256_cmpgt_epi8(op, _A);
        __m256i lt = _mm256_cmpgt_epi8(Z_, op);
        __m256i mingle = _mm256_and_si256(gt, lt);
        __m256i add = _mm256_and_si256(mingle, delta);
        __m256i lower = _mm256_add_epi8(op, add);
        _mm256_storeu_si256((__m256i *)q, lower);
        src += 32;
        q += 32;
        len -= 32;
    };
#else 
    const __m128i _A = _mm_set1_epi8('A' - 1);
    const __m128i Z_ = _mm_set1_epi8('Z' + 1);
    const __m128i delta = _mm_set1_epi8('a' - 'A');
    uint8_t* q = dst;

    while (len >= 16){
        __m128i op = _mm_loadu_si128((__m128i*)src);
        __m128i gt = _mm_cmpgt_epi8(op, _A);
        __m128i lt = _mm_cmplt_epi8(op, Z_);
        __m128i mingle = _mm_and_si128(gt, lt);
        __m128i add = _mm_and_si128(mingle, delta);
        __m128i lower = _mm_add_epi8(op, add);
        _mm_storeu_si128((__m128i *)q, lower);
        src += 16;
        q += 16;
        len -= 16;
    };

#endif 

    while(len > 0) {
        len--;
        bool isUpper = (*src >= 'A' && *src <= 'Z');
        *q = isUpper? (*src + 'a' - 'A'): *src;
        src++;
        q++;
    }
}

static always_inline ssize_t compre_8(const uint8_t* a, const uint8_t* b, size_t len) {
    int i = 0;
    for (; i < 8; i++) {
        if (a[i] != b[i]) {
            break;
        }
    }
    // matched
    if (i >= len) {
        return 0;
    }
    return -1;
}

static always_inline ssize_t compre_32(const uint8_t* a, const uint8_t* b, size_t len) {
    v256u va = v256_loadu(a);
    v256u vb = v256_loadu(b);

    uint32_t m = mask256_tobitmask(v256_eq(va, vb));
    m = (~m) & ((1ull << len) - 1);
    // matched
    if (m == 0) {
        return 0;
    }
    return -1;
}

// NOTE:
// key amd table must be padding to 32 bytes at last to make it friendly to SIMD.
// key length should be le 32 bytes.
long lookup_small_key(GoString* key, GoSlice* table, long lower_off) {
    uint8_t len = key->len;
    uint8_t* p = table->buf;
    uint8_t* meta = p + len * _HDR_SLOT_SIZE;
    uint8_t* kp = (uint8_t*)key->buf;
    uint8_t  cnt = *meta;
    ssize_t  offset = *(uint32_t*)(meta + 1) + _HDR_SIZE;
    uint8_t  lower[32] = {0};

    p += offset;

    if (cnt == 0) {
        return -_NOT_FOUND;
    }

    if (len <= 8) {
        for (int i = 0; i < cnt; i++) {
            if (compre_8(p, kp, len) == 0) {
                return p[len];
            } else {
                p += len + 1;
            }
        }
        goto case_sensitve;
    }

    for (int i = 0; i < cnt; i++) {
        if (compre_32(p, kp, len) == 0) {
            return p[len];
        } else {
            p += len + 1;
        }
    }

case_sensitve:
    offset = lower_off + *(uint32_t*)(meta + 1);
    p = table->buf + offset;
    to_lower(lower, (uint8_t*)(key->buf), 32);

    if (len <= 8) {
        for (int i = 0; i < cnt; i++) {
            if (compre_8(p, lower, len) == 0) {
                return p[len];
            } else {
                p += len + 1;
            }
        }
        return -_NOT_FOUND;
    }

    for (int i = 0; i < cnt; i++) {
        if (compre_32(p, lower, len) == 0) {
            return p[len];
        } else {
            p += len + 1;
        }
    }
    return -_NOT_FOUND;
}