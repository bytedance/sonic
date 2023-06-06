#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <immintrin.h>


// src 是输入的字符串，可能包含大写字符，dst是转换后的小写字符串，dst位置已经预分配长度为len的内存
void to_lower(char* dst, const char* src, size_t len) {
    const __m128i _A = _mm_set1_epi8('A' - 1);
    const __m128i Z_ = _mm_set1_epi8('Z' + 1);
    const __m128i delta = _mm_set1_epi8('a' - 'A');
    char* q = dst;
    do {
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
    } while (len >= 16);

    while(len > 0) {
        bool isUpper = (*src >= 'A' && *src <= 'Z');
        *q = isUpper? (*src + 32): *src;
        src++;
        q++;
    }
}