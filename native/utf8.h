/*
 * Copyright (C) 2019 Yaoyuan <ibireme@gmail.com>.
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
 *
 * Copyright 2018-2023 The simdjson authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * 
 * This file may have been modified by ByteDance authors. All ByteDance
 * Modifications are Copyright 2022 ByteDance Authors.
 */

#pragma once

#include "native.h"
#include "utils.h"
#include "test/xassert.h"
#include "test/xprintf.h"

static inline ssize_t valid_utf8_4byte(uint32_t ubin) {
    /*
     Each unicode code point is encoded as 1 to 4 bytes in UTF-8 encoding,
     we use 4-byte mask and pattern value to validate UTF-8 byte sequence,
     this requires the input data to have 4-byte zero padding.
     ---------------------------------------------------
     1 byte
     unicode range [U+0000, U+007F]
     unicode min   [.......0]
     unicode max   [.1111111]
     bit pattern   [0.......]
     ---------------------------------------------------
     2 byte
     unicode range [U+0080, U+07FF]
     unicode min   [......10 ..000000]
     unicode max   [...11111 ..111111]
     bit require   [...xxxx. ........] (1E 00)
     bit mask      [xxx..... xx......] (E0 C0)
     bit pattern   [110..... 10......] (C0 80)
     // 1101 0100 10110000
     // 0001 1110
     ---------------------------------------------------
     3 byte
     unicode range [U+0800, U+FFFF]
     unicode min   [........ ..100000 ..000000]
     unicode max   [....1111 ..111111 ..111111]
     bit require   [....xxxx ..x..... ........] (0F 20 00)
     bit mask      [xxxx.... xx...... xx......] (F0 C0 C0)
     bit pattern   [1110.... 10...... 10......] (E0 80 80)
     ---------------------------------------------------
     3 byte invalid (reserved for surrogate halves)
     unicode range [U+D800, U+DFFF]
     unicode min   [....1101 ..100000 ..000000]
     unicode max   [....1101 ..111111 ..111111]
     bit mask      [....xxxx ..x..... ........] (0F 20 00)
     bit pattern   [....1101 ..1..... ........] (0D 20 00)
     ---------------------------------------------------
     4 byte
     unicode range [U+10000, U+10FFFF]
     unicode min   [........ ...10000 ..000000 ..000000]
     unicode max   [.....100 ..001111 ..111111 ..111111]
     bit err0      [.....100 ........ ........ ........] (04 00 00 00)
     bit err1      [.....011 ..110000 ........ ........] (03 30 00 00)
     bit require   [.....xxx ..xx.... ........ ........] (07 30 00 00)
     bit mask      [xxxxx... xx...... xx...... xx......] (F8 C0 C0 C0)
     bit pattern   [11110... 10...... 10...... 10......] (F0 80 80 80)
     ---------------------------------------------------
     */
    const uint32_t b2_mask = 0x0000C0E0UL;
    const uint32_t b2_patt = 0x000080C0UL;
    const uint32_t b2_requ = 0x0000001EUL;
    const uint32_t b3_mask = 0x00C0C0F0UL;
    const uint32_t b3_patt = 0x008080E0UL;
    const uint32_t b3_requ = 0x0000200FUL;
    const uint32_t b3_erro = 0x0000200DUL;
    const uint32_t b4_mask = 0xC0C0C0F8UL;
    const uint32_t b4_patt = 0x808080F0UL;
    const uint32_t b4_requ = 0x00003007UL;
    const uint32_t b4_err0 = 0x00000004UL;
    const uint32_t b4_err1 = 0x00003003UL;

#define is_valid_seq_2(uni) ( \
    ((uni & b2_mask) == b2_patt) && \
    ((uni & b2_requ)) \
)
    
#define is_valid_seq_3(uni) ( \
    ((uni & b3_mask) == b3_patt) && \
    ((tmp = (uni & b3_requ))) && \
    ((tmp != b3_erro)) \
)
    
#define is_valid_seq_4(uni) ( \
    ((uni & b4_mask) == b4_patt) && \
    ((tmp = (uni & b4_requ))) && \
    ((tmp & b4_err0) == 0 || (tmp & b4_err1) == 0) \
)
    uint32_t tmp = 0;
   
    if (is_valid_seq_3(ubin)) return 3;
    if (is_valid_seq_2(ubin)) return 2;
    if (is_valid_seq_4(ubin)) return 4;
    return 0;
}

static always_inline long write_error(int pos, StateMachine *m, size_t msize) {
    if (m->sp >= msize) {
        return -1;
    }
    m->vt[m->sp++] = pos;
    return 0;
}

// scalar code, error position should excesss 4096
static always_inline long validate_utf8_with_errors(const char *src, long len, long *p, StateMachine *m) {
    const char* start = src + *p;
    const char* end = src + len;
    while (start < end - 3) {
        uint32_t u = (*(uint32_t*)(start));
        if ((unsigned)(*start) < 0x80) {
            start += 1;
            continue;
        }
        size_t n = valid_utf8_4byte(u);
        if (n != 0) { // valid utf
            start += n;
            continue;
        }
        long err = write_error(start - src, m, MAX_RECURSE);
        if (err) {
            *p = start - src;
            return err;
        }
        start += 1;
    }
    while (start < end) {
        if ((unsigned)(*start) < 0x80) {
            start += 1;
            continue;
        }
        uint32_t u = 0;
        memcpy_p4(&u, start, end - start);
        size_t n = valid_utf8_4byte(u);
        if (n != 0) { // valid utf
            start += n;
            continue;
        }
        long err = write_error(start - src, m, MAX_RECURSE);
        if (err) {
            *p = start - src;
            return err;
        }
        start += 1;
    }
    *p = start - src;
    return 0;
}

// validate_utf8_errors returns zero if valid, otherwise, the error position.
static always_inline long validate_utf8_errors(const GoString* s) {
    const char* start = s->buf;
    const char* end = s->buf + s->len;
    while (start < end - 3) {
        uint32_t u = (*(uint32_t*)(start));
        if ((unsigned)(*start) < 0x80) {
            start += 1;
            continue;
        }
        size_t n = valid_utf8_4byte(u);
        if (n == 0) { // invalid utf
            return -(start - s->buf) - 1;
        }
        start += n;
    }
    while (start < end) {
        if ((unsigned)(*start) < 0x80) {
            start += 1;
            continue;
        }
        uint32_t u = 0;
        memcpy_p4(&u, start, end - start);
        size_t n = valid_utf8_4byte(u);
        if (n == 0) { // invalid utf
            return -(start - s->buf) - 1;
        }
        start += n;
    }
    return 0;
}

// SIMD implementation
#if USE_AVX2

    static always_inline __m256i simd256_shr(const __m256i input, const int shift) {
        __m256i shifted = _mm256_srli_epi16(input, shift);
        __m256i mask = _mm256_set1_epi8(0xFFu >> shift);
        return _mm256_and_si256(shifted, mask);
    }

#define simd256_prev(input, prev, N) _mm256_alignr_epi8(input, _mm256_permute2x128_si256(prev, input, 0x21), 16 - (N));

    static always_inline __m256i must_be_2_3_continuation(const __m256i prev2, const __m256i prev3) {
        __m256i is_third_byte  = _mm256_subs_epu8(prev2, _mm256_set1_epi8(0b11100000u-1)); // Only 111_____ will be > 0
        __m256i is_fourth_byte = _mm256_subs_epu8(prev3, _mm256_set1_epi8(0b11110000u-1)); // Only 1111____ will be > 0
        // Caller requires a bool (all 1's). All values resulting from the subtraction will be <= 64, so signed comparison is fine.
        __m256i or = _mm256_or_si256(is_third_byte, is_fourth_byte);
        return _mm256_cmpgt_epi8(or, _mm256_set1_epi8(0));;
    }

    static always_inline __m256i simd256_lookup16(const __m256i input, const uint8_t* table) {
        return _mm256_shuffle_epi8(_mm256_setr_epi8(table[0], table[1], table[2], table[3], table[4], table[5], table[6], table[7], table[8], table[9], table[10], table[11], table[12], table[13], table[14], table[15], table[0], table[1], table[2], table[3], table[4], table[5], table[6], table[7], table[8], table[9], table[10], table[11], table[12], table[13], table[14], table[15]), input);
    }

  //
  // Return nonzero if there are incomplete multibyte characters at the end of the block:
  // e.g. if there is a 4-byte character, but it's 3 bytes from the end.
  //
      static always_inline  __m256i is_incomplete(const __m256i input) {
    // If the previous input's last 3 bytes match this, they're too short (they ended at EOF):
    // ... 1111____ 111_____ 11______
      const uint8_t tab[32] = {
      255, 255, 255, 255, 255, 255, 255, 255,
      255, 255, 255, 255, 255, 255, 255, 255,
      255, 255, 255, 255, 255, 255, 255, 255,
      255, 255, 255, 255, 255, 0b11110000u-1, 0b11100000u-1, 0b11000000u-1};
        const __m256i max_value = _mm256_loadu_si256((const __m256i_u *)(&tab[0]));
        return _mm256_subs_epu8(input, max_value);
    }

  static always_inline __m256i check_special_cases(const __m256i input, const __m256i prev1) {
    // Bit 0 = Too Short (lead byte/ASCII followed by lead byte/ASCII)
    // Bit 1 = Too Long (ASCII followed by continuation)
    // Bit 2 = Overlong 3-byte
    // Bit 4 = Surrogate
    // Bit 5 = Overlong 2-byte
    // Bit 7 = Two Continuations
     const uint8_t TOO_SHORT   = 1<<0; // 11______ 0_______
                                                // 11______ 11______
     const uint8_t TOO_LONG    = 1<<1; // 0_______ 10______
     const uint8_t OVERLONG_3  = 1<<2; // 11100000 100_____
     const uint8_t SURROGATE   = 1<<4; // 11101101 101_____
     const uint8_t OVERLONG_2  = 1<<5; // 1100000_ 10______
     const uint8_t TWO_CONTS   = 1<<7; // 10______ 10______
     const uint8_t TOO_LARGE   = 1<<3; // 11110100 1001____
                                                // 11110100 101_____
                                                // 11110101 1001____
                                                // 11110101 101_____
                                                // 1111011_ 1001____
                                                // 1111011_ 101_____
                                                // 11111___ 1001____
                                                // 11111___ 101_____
     const uint8_t TOO_LARGE_1000 = 1<<6;
                                                // 11110101 1000____
                                                // 1111011_ 1000____
                                                // 11111___ 1000____
     const uint8_t OVERLONG_4  = 1<<6; // 11110000 1000____

    const __m256i prev1_shr4 = simd256_shr(prev1, 4);
    static const uint8_t tab1[16] = {
              // 0_______ ________ <ASCII in byte 1>
      TOO_LONG, TOO_LONG, TOO_LONG, TOO_LONG,
      TOO_LONG, TOO_LONG, TOO_LONG, TOO_LONG,
      // 10______ ________ <continuation in byte 1>
      TWO_CONTS, TWO_CONTS, TWO_CONTS, TWO_CONTS,
      // 1100____ ________ <two byte lead in byte 1>
      TOO_SHORT | OVERLONG_2,
      // 1101____ ________ <two byte lead in byte 1>
      TOO_SHORT,
      // 1110____ ________ <three byte lead in byte 1>
      TOO_SHORT | OVERLONG_3 | SURROGATE,
      // 1111____ ________ <four+ byte lead in byte 1>
      TOO_SHORT | TOO_LARGE | TOO_LARGE_1000 | OVERLONG_4,
    };
    __m256i byte_1_high = simd256_lookup16(prev1_shr4, tab1);
    

    const uint8_t CARRY = TOO_SHORT | TOO_LONG | TWO_CONTS; // These all have ____ in byte 1 .
    __m256i prev1_low = _mm256_and_si256(prev1, _mm256_set1_epi8(0x0F));
    static const uint8_t tab2[16] = {
      // ____0000 ________
      CARRY | OVERLONG_3 | OVERLONG_2 | OVERLONG_4,
      // ____0001 ________
      CARRY | OVERLONG_2,
      // ____001_ ________
      CARRY,
      CARRY,

      // ____0100 ________
      CARRY | TOO_LARGE,
      // ____0101 ________
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      // ____011_ ________
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000,

      // ____1___ ________
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      // ____1101 ________
      CARRY | TOO_LARGE | TOO_LARGE_1000 | SURROGATE,
      CARRY | TOO_LARGE | TOO_LARGE_1000,
      CARRY | TOO_LARGE | TOO_LARGE_1000
    };
    __m256i byte_1_low = simd256_lookup16(prev1_low, tab2);
    

    const __m256i input_shr4 = simd256_shr(input, 4);
    static const uint8_t tab3[16] = {
      // ________ 0_______ <ASCII in byte 2>
      TOO_SHORT, TOO_SHORT, TOO_SHORT, TOO_SHORT,
      TOO_SHORT, TOO_SHORT, TOO_SHORT, TOO_SHORT,

      // ________ 1000____
      TOO_LONG | OVERLONG_2 | TWO_CONTS | OVERLONG_3 | TOO_LARGE_1000 | OVERLONG_4,
      // ________ 1001____
      TOO_LONG | OVERLONG_2 | TWO_CONTS | OVERLONG_3 | TOO_LARGE,
      // ________ 101_____
      TOO_LONG | OVERLONG_2 | TWO_CONTS | SURROGATE  | TOO_LARGE,
      TOO_LONG | OVERLONG_2 | TWO_CONTS | SURROGATE  | TOO_LARGE,

      // ________ 11______
      TOO_SHORT, TOO_SHORT, TOO_SHORT, TOO_SHORT
    };
    __m256i byte_2_high = simd256_lookup16(input_shr4, tab3);
     

    return _mm256_and_si256(_mm256_and_si256(byte_1_high, byte_1_low), byte_2_high);
  }

    static always_inline __m256i check_multibyte_lengths(const __m256i input, const __m256i prev_input, const __m256i sc) {
    __m256i prev2 = simd256_prev(input, prev_input, 2);
    __m256i prev3 = simd256_prev(input, prev_input, 3);
    
    
    __m256i must23 = must_be_2_3_continuation(prev2, prev3);
    
    __m256i must23_80 = _mm256_and_si256(must23, _mm256_set1_epi8(0x80));
    
    return _mm256_xor_si256(must23_80, sc);
  }


    // Check whether the current bytes are valid UTF-8.
    static always_inline __m256i check_utf8_bytes(const __m256i input, const __m256i prev_input) {
        // Flip prev1...prev3 so we can easily determine if they are 2+, 3+ or 4+ lead bytes
        // (2, 3, 4-byte leads become large positive numbers instead of small negative numbers)
        __m256i prev1 = simd256_prev(input, prev_input, 1);
        __m256i sc    = check_special_cases(input, prev1);
        __m256i ret  = check_multibyte_lengths(input, prev_input, sc);
        return ret;
    }

    static always_inline bool is_ascii(const __m256i input) {
      return _mm256_movemask_epi8(input) == 0;
    }

    typedef struct {
        // If this is nonzero, there has been a UTF-8 error.
        __m256i error;
        // The last input we received
        __m256i prev_input_block;
        // Whether the last input we received was incomplete (used for ASCII fast path)
        __m256i prev_incomplete;
    } utf8_checker;

    static always_inline void utf8_checker_init(utf8_checker* checker) {
        checker->error = _mm256_setzero_si256();
        checker->prev_input_block = _mm256_setzero_si256();
        checker->prev_incomplete = _mm256_setzero_si256();
    }
    
    static always_inline bool check_error(utf8_checker* checker) {
        return !_mm256_testz_si256(checker->error, checker->error);
    }

    static always_inline void check64_utf(utf8_checker* checker, const uint8_t* start) {
        __m256i input = _mm256_loadu_si256((__m256i*)start);
        __m256i input2 = _mm256_loadu_si256((__m256i*)(start + 32));
        // check utf-8 chars
        __m256i error1 = check_utf8_bytes(input, checker->prev_input_block);
        __m256i error2 = check_utf8_bytes(input2, input);
        checker->error = _mm256_or_si256(checker->error, _mm256_or_si256(error1, error2));
        checker->prev_input_block = input2;
        checker->prev_incomplete = is_incomplete(input2);
    }

    static always_inline void check64(utf8_checker* checker, const uint8_t* start) {
        // fast path for contiguous ASCII
        __m256i input = _mm256_loadu_si256((__m256i*)start);
        __m256i input2 = _mm256_loadu_si256((__m256i*)(start + 32));
        __m256i reducer = _mm256_or_si256(input, input2);
        // check utf-8
        if (likely(is_ascii(reducer))) {
            checker->error = _mm256_or_si256(checker->error, checker->prev_incomplete);
            return;
        }
        check64_utf(checker, start);
    }

    static always_inline void check128(utf8_checker* checker, const uint8_t* start) {
        // fast path for contiguous ASCII
        __m256i input = _mm256_loadu_si256((__m256i*)start);
        __m256i input2 = _mm256_loadu_si256((__m256i*)(start + 32));
        __m256i input3 = _mm256_loadu_si256((__m256i*)(start + 64));
        __m256i input4 = _mm256_loadu_si256((__m256i*)(start + 96));
        
        __m256i reducer1 = _mm256_or_si256(input, input2);
        __m256i reducer2 = _mm256_or_si256(input3, input4);
        __m256i reducer  = _mm256_or_si256(reducer1, reducer2);

        // full 128 bytes are ascii
        if (likely(is_ascii(reducer))) {
            checker->error = _mm256_or_si256(checker->error, checker->prev_incomplete);
            return;
        }

        // frist 64 bytes is ascii, next 64 bytes must be utf8
        if (likely(is_ascii(reducer1))) {
            checker->error = _mm256_or_si256(checker->error, checker->prev_incomplete);
            check64_utf(checker, start + 64);
            return;
        }

        // frist 64 bytes has utf8, next 64 bytes 
        check64_utf(checker, start);
        if (unlikely(is_ascii(reducer2))) {
            checker->error = _mm256_or_si256(checker->error, checker->prev_incomplete);
        } else {
            check64_utf(checker, start + 64);
        }
    }

    static always_inline void check_eof(utf8_checker* checker) {
        checker->error = _mm256_or_si256(checker->error, checker->prev_incomplete);
    }

    static always_inline void check_remain(utf8_checker* checker, const uint8_t* start, const uint8_t* end) {
        uint8_t buffer[64] = {0};
        int i = 0;
        while (start < end) {
            buffer[i++] = *(start++);
        };
        check64(checker, buffer);
        check_eof(checker);
    }

    static always_inline long validate_utf8_avx2(const GoString* s) {
        xassert(s->buf != NULL || s->len != 0);
        const uint8_t* start = (const uint8_t*)(s->buf);
        const uint8_t* end   = (const uint8_t*)(s->buf + s->len);
        /* check eof */
        if (s->len == 0) {
            return 0;
        }
        utf8_checker checker;
        utf8_checker_init(&checker);
        while (start < (end - 128)) {
            check128(&checker, start);
            if (check_error(&checker)) {
            }
            start += 128;
        };
        while (start < end - 64) {
            check64(&checker, start);
            start += 64;
        }
        check_remain(&checker, start, end);
        return check_error(&checker) ? -1 : 0;
    }
#endif
