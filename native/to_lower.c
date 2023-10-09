#include "native.h"

/**
 * @brief uppercase to lowercase
 * @param[in] src original string
 * @param[in, out] dst lowercase string
 * @param[in] len length of the string
*/
void to_lower(char* dst, const char* src, size_t len) {
    const __m128i _A = _mm_set1_epi8('A' - 1);
    const __m128i Z_ = _mm_set1_epi8('Z' + 1);
    const __m128i delta = _mm_set1_epi8('a' - 'A');
    char* q = dst;

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

    while(len > 0) {
        len--;
        bool isUpper = (*src >= 'A' && *src <= 'Z');
        *q = isUpper? (*src + 32): *src;
        src++;
        q++;
    }
}